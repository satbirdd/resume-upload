package resume_upload

import (
	"net/http"
	"sync"

	log "github.com/dsoprea/go-logging"
	"github.com/eventials/go-tus/leveldbstore"
)

var (
	config  *Config
	once    sync.Once
	initErr error
)

type Config struct {
	// ChunkSize divide the file into chunks.
	ChunkSize int64
	// Resume enables resumable upload.
	Resume bool
	// OverridePatchMethod allow to by pass proxies sendind a POST request instead of PATCH.
	OverridePatchMethod bool
	// Store map an upload's fingerprint with the corresponding upload URL.
	// If Resume is true the Store is required.
	Store Store
	// Set custom header values used in all requests.
	Header http.Header
	// HTTP Client
	HttpClient *http.Client
}

func DefaultTusConfig() (*Config, error) {
	once.Do(func() {
		store, err := leveldbstore.NewLeveldbStore(TusLevelDBPath)
		if err != nil {
			initErr = log.Errorf("创建leveldb存储失败，%v", err)
			return
		}

		config = &Config{
			ChunkSize:           2 * 1024 * 1024,
			Resume:              true,
			OverridePatchMethod: false,
			Store:               store,
			Header:              make(http.Header),
			HttpClient:          nil,
		}
	})

	return config, initErr
}

func DefaultTusConfigWithHeader(header http.Header) (*Config, error) {
	config, err := DefaultTusConfig()

	if config != nil {
		config.Header = header
	}

	return config, err
}
