package resume_upload

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/eventials/go-tus"
	"github.com/eventials/go-tus/leveldbstore"
)

const (
	TusLevelDBPath = "./___tus___.upload.db"
	MissMatch      = "mismatch"
)

type Client struct {
	l       sync.Mutex
	url     string
	backoff Backoffer
	c       *tus.Client

	// connected bool
	// store     *leveldbstore.LeveldbStore
}

func DefaultTusConfig() (*tus.Config, error) {
	store, err := leveldbstore.NewLeveldbStore(TusLevelDBPath)
	if err != nil {
		return nil, fmt.Errorf("创建leveldb存储失败，%v", err)
	}

	return &tus.Config{
		ChunkSize:           2 * 1024 * 1024,
		Resume:              true,
		OverridePatchMethod: false,
		Store:               store,
		Header:              make(http.Header),
		HttpClient:          nil,
	}, nil
}

func NewClient(url string, cfg *tus.Config, backoff Backoffer) (*Client, error) {
	var (
		err error
	)

	if cfg == nil {
		cfg, err = DefaultTusConfig()
		if err != nil {
			return nil, err
		}
	}

	client, err := tus.NewClient(url, cfg)
	if err != nil {
		return nil, err
	}

	if backoff == nil {
		backoff = DefaultBackoff
	}

	return &Client{
		url:     url,
		c:       client,
		backoff: backoff,
	}, nil
}

func (client *Client) Upload(path string, ch chan<- struct{}) error {
	if info, err := os.Stat(path); err != nil {
		return fmt.Errorf("文件%v无法读取，%v", path, err)
	} else if info.IsDir() {
		return fmt.Errorf("上传的目标不能是文件夹，%v", path)
	}

	f, err := os.Open(fmt.Sprintf("/home/liulei/Downloads/研发需求资料.rar"))
	if err != nil {
		return err
	}

	defer f.Close()

	upload, err := tus.NewUploadFromFile(f)
	if err != nil {
		return err
	}

	uploader, err := client.c.CreateOrResumeUpload(upload)
	if err != nil {
		return err
	}

	err = uploader.Upload()
	n := 0
	for err != nil {
		if client.backoff != nil {
			time.Sleep(client.backoff.Backoff(int(n)))
		}

		if strings.Contains(err.Error(), MissMatch) {
			uploader, err = client.c.CreateOrResumeUpload(upload)
		}

		err = uploader.Upload()
	}

	ch <- struct{}{}

	return nil
}
