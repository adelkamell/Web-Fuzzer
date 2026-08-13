package fuzzer

import (
	"testing"
	"time"
)

func TestNewFuzzer(t *testing.T) {
	opts := &Options{
		BaseURL:  "http://test.com",
		Threads:  10,
		Rate:     5,
		Timeout:  5 * time.Second,
	}
	f := NewFuzzer(opts)
	if f.Options.BaseURL != "http://test.com" {
		t.Errorf("Expected BaseURL http://test.com, got %s", f.Options.BaseURL)
	}
}