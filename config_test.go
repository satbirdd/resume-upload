package resume_upload

import "testing"

func TestDefaultConfig(t *testing.T) {
	cfg, err := DefaultTusConfig()
	if err != nil {
		t.Fatal(err)
	}

	if cfg == nil {
		t.Fatal("DefaultTusConfig is nill")
	}
}
