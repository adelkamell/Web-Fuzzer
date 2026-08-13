package fuzzer

import "time"

type Options struct {
	BaseURL    string
	Wordlist   string
	Threads    int
	Extensions []string
	Rate       int
	Timeout    time.Duration
	UserAgent  string
}