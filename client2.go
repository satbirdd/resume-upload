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
	log "github.com/sirupsen/logrus"
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

func init() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	log.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true
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
		log.Warnf("[Resumable Upload]文件 %v 第%v上传失败，%v", path, n+1, err)
		if client.backoff != nil {
			time.Sleep(client.backoff.Backoff(int(n)))
		}

		if strings.Contains(err.Error(), MissMatch) {
			uploader, err = client.c.CreateOrResumeUpload(upload)
		}

		n += 1

		err = uploader.Upload()
	}

	log.Infof("[Resumable Upload]文件 %v 第%v上传成功", path, n+1)

	ch <- struct{}{}

	return nil
}
