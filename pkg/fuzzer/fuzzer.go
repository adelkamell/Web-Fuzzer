package fuzzer

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	//"strings"
	"sync"
	"time"
)

type Job struct {
	Path string
	Ext  string
}

type Result struct {
	FullPath string
	Status   int
	Size     int64
}

type Fuzzer struct {
	Options    *Options
	Client     *http.Client
	Results    chan Result
	RateLimiter <-chan time.Time
}

func NewFuzzer(opts *Options) *Fuzzer {
	client := &http.Client{
		Timeout: opts.Timeout * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	return &Fuzzer{
		Options:    opts,
		Client:     client,
		Results:    make(chan Result, opts.Threads*2),
		RateLimiter: time.Tick(time.Second / time.Duration(opts.Rate)),
	}
}

func (f *Fuzzer) worker(jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		<-f.RateLimiter // Rate limiting
		url := f.Options.BaseURL + "/" + job.Path + job.Ext

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			continue
		}
		req.Header.Set("User-Agent", f.Options.UserAgent)

		resp, err := f.Client.Do(req)
		if err != nil {
			continue
		}
		resp.Body.Close()

		if resp.StatusCode >= 200 && resp.StatusCode < 400 {
			f.Results <- Result{
				FullPath: url,
				Status:   resp.StatusCode,
				Size:     resp.ContentLength,
			}
		}
	}
}

func (f *Fuzzer) Run() {
	file, err := os.Open(f.Options.Wordlist)
	if err != nil {
		fmt.Printf("Error opening wordlist: %v\n", err)
		return
	}
	defer file.Close()

	jobs := make(chan Job, f.Options.Threads*2)
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < f.Options.Threads; i++ {
		wg.Add(1)
		go f.worker(jobs, &wg)
	}

	// Feed jobs
	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			word := scanner.Text()
			for _, ext := range f.Options.Extensions {
				jobs <- Job{Path: word, Ext: ext}
			}
		}
		close(jobs)
	}()

	// Close results after workers finish
	go func() {
		wg.Wait()
		close(f.Results)
	}()

	// Print results
	for res := range f.Results {
		fmt.Printf("[+] %-60s [%d] Size: %d\n", res.FullPath, res.Status, res.Size)
	}
}