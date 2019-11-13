package resume_upload

import (
	"net/http"
	"testing"

	"github.com/eventials/go-tus/leveldbstore"
)

func TestNewClient(t *testing.T) {
	url := "http://127.0.0.1:8080/files/"

	store, err := leveldbstore.NewLeveldbStore("./test.upload.db")
	if err != nil {
		t.Fatalf("创建leveldb存储失败，%v", err)
	}

	cfg := &Config{
		ChunkSize:           2 * 1024 * 1024,
		Resume:              true,
		OverridePatchMethod: false,
		Store:               store,
		Header:              make(http.Header),
		HttpClient:          nil,
	}

	_, err = NewClient(url, cfg, nil)
	if err != nil {
		t.Fatal(err)
	}
}
