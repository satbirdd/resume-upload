package resume_upload

import (
	"net/http"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg, err := DefaultTusConfig()
	if err != nil {
		t.Fatal(err)
	}

	if cfg == nil {
		t.Fatal("DefaultTusConfig is nill")
	}
}

func TestDefaultConfigWithHeader(t *testing.T) {

	cfg, err := DefaultTusConfig()
	if err != nil {
		t.Fatal(err)
	}

	if cfg == nil {
		t.Fatal("DefaultTusConfig is nill")
	}

	header := http.Header{
		"Authorization": []string{"Bearer user:password"},
	}
	cfg, err = DefaultTusConfigWithHeader(header)
	if err != nil {
		t.Fatal(err)
	}

	if cfg == nil {
		t.Fatal("DefaultTusConfigWithHeader is nil")
	}

	if cfg.Header["Authorization"][0] != header["Authorization"][0] {
		t.Fatal("Set DefaultTusConfigWithHeader failed")
	}

	cfg, err = DefaultTusConfig()
	if err != nil {
		t.Fatal(err)
	}

	if cfg == nil {
		t.Fatal("DefaultTusConfig is nill")
	}

	if len(cfg.Header["Authorization"]) != 0 {
		t.Fatal("DefaultTusConfig has been changed")
	}
}
