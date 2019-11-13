package resume_upload

// import (
// 	"fmt"
// 	"os"
// 	"sync"

// 	"github.com/eventials/go-tus"
// 	"github.com/syndtr/goleveldb/leveldb"
// )

// const (
// 	TusLevelDBPath  = "./___tus___.upload.db"
// 	TaskLevelDBPath = "./___ru___.task.db"
// )

// var (
// 	client        *Client
// 	once          sync.Once
// 	leveldbClient *leveldb.Client
// 	tusClient     *tus.Client
// )

// func init() {
// 	var err error
// 	leveldbClient, err = leveldb.OpenFile(TaskLevelDBPath, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// type Client struct {
// 	l         sync.Mutex
// 	url       string
// 	connected bool
// 	// c         tus.Client
// 	// store     *leveldbstore.LeveldbStore
// }

// func GetClient(url string) (*Client, error) {
// 	once.Do(func() {
// 		// store, _ := leveldbstore.NewLeveldbStore(LevelDBPath)
// 		// client, _ := tus.NewClient(url, &tus.Config{
// 		// 	ChunkSize:           2 * 1024 * 1024,
// 		// 	Resume:              true,
// 		// 	OverridePatchMethod: false,
// 		// 	Store:               store,
// 		// 	Header:              make(http.Header),
// 		// 	HttpClient:          nil,
// 		// })

// 		// taskStore, _ := leveldb.OpenFile(TaskLevelDBPath, nil)
// 		client = &Client{
// 			url: url,
// 			// store:     taskStore,
// 			// c:         client,
// 			connected: true,
// 		}

// 		client.StartBGTask()
// 	})

// 	return client, nil
// }

// func (client *Cilent) SetConnected(val bool) {
// 	client.l.Lock()

// 	defer client.l.Unlock()

// 	client.connected = val
// }

// func (client *Client) StartBGTask() {
// 	go func() {
// 		for {

// 		}
// 	}()
// }

// // 对于加入的任务，仅仅在leveldb上做个标记
// // 然后由后台任务处理
// func (client *Client) AddUpload(path string) error {
// 	if info, err := os.Stat(path); err != nil {
// 		return fmt.Errorf("文件%v无法读取，%v", err)
// 	}

// 	// if client.connected {
// 	// 	f, err := os.Open(path)
// 	// 	if err != nil {
// 	// 		return fmt.Errorf("文件%v无法读取，%v", err)
// 	// 	}

// 	// 	defer f.Close()

// 	// 	upload, err := tus.NewUploadFromFile(f)
// 	// 	if err != nil {
// 	// 		panic(err)
// 	// 	}

// 	// 	uploader, err := client.CreateOrResumeUpload(upload)
// 	// 	if err != nil {
// 	// 		panic(err)
// 	// 	}

// 	// 	err = uploader.Upload()
// 	// 	if err != nil {
// 	// 		client.store.Put([]byte(path), []byte("1"), nil)
// 	// 	}
// 	// } else {
// 	levelDBClient().Put([]byte(path), []byte("1"), &leveldb.WriteOptions{
// 		Sync: true,
// 	})
// 	// }
// }
