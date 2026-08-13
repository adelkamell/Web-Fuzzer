package fuzzer

import "testing"

func TestNewFuzzer(t *testing.T) {
	f := New("http://test.com", []string{"admin"})
	if f.Target != "http://test.com" {
		t.Errorf("Target mismatch")
	}
}