package fuzzer

import (
	"fmt"
	"net/http"
	"time"
)

type Fuzzer struct {
	Target   string
	Wordlist []string
	Client   *http.Client
}

func New(target string, wordlist []string) *Fuzzer {
	return &Fuzzer{
		Target:   target,
		Wordlist: wordlist,
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (f *Fuzzer) Run() {
	for _, path := range f.Wordlist {
		url := f.Target + "/" + path
		resp, err := f.Client.Head(url)
		if err != nil {
			fmt.Printf("[ERROR] %s - %v\n", url, err)
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode != 404 {
			fmt.Printf("[FOUND] %s -> Status: %d\n", url, resp.StatusCode)
		}
	}
}