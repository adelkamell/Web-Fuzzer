// web fuzzer - V3
// Enhancements: rate limiting with time.Tick, support for extensions (php, asp, ...), and custom User-Agent.
// Athor : https://github.com/adelkamell


package main

import (
    "bufio"
    "flag"
    "fmt"
    "net/http"
    "os"
    "strings"
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

func worker(base string, client *http.Client, jobs <-chan Job, wg *sync.WaitGroup, results chan<- Result, rateLimiter <-chan time.Time) {
    defer wg.Done()
    for job := range jobs {
        <-rateLimiter // wait for tick
        url := base + "/" + job.Path + job.Ext
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            continue
        }
        req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; rv:102.0) Gecko/20100101 Firefox/102.0")
        // Additional custom headers could be added from flags (omitted for brevity)
        resp, err := client.Do(req)
        if err != nil {
            continue
        }
        resp.Body.Close()
        if resp.StatusCode >= 200 && resp.StatusCode < 400 {
            results <- Result{FullPath: url, Status: resp.StatusCode, Size: resp.ContentLength}
        }
    }
}

func main() {
    urlFlag := flag.String("u", "", "Base URL (e.g. http://target.com)")
    wordlist := flag.String("w", "wordlist.txt", "Wordlist file")
    threads := flag.Int("t", 20, "Number of threads")
    extensions := flag.String("x", "", "Comma-separated extensions (e.g. php,asp)")
    rate := flag.Int("rate", 10, "Requests per second")
    timeout := flag.Int("timeout", 5, "HTTP timeout in seconds")
    flag.Parse()

    if *urlFlag == "" {
        fmt.Println("Usage: go run main.go -u <BASE_URL> -w wordlist.txt")
        return
    }

    // Prepare extensions
    var exts []string
    if *extensions != "" {
        for _, e := range strings.Split(*extensions, ",") {
            exts = append(exts, "."+strings.TrimSpace(e))
        }
    } else {
        exts = append(exts, "") // no extension
    }

    // HTTP client with timeout and no redirects
    client := &http.Client{
        Timeout: time.Duration(*timeout) * time.Second,
        CheckRedirect: func(req *http.Request, via []*http.Request) error {
            return http.ErrUseLastResponse
        },
    }

    file, err := os.Open(*wordlist)
    if err != nil {
        fmt.Printf("Error opening wordlist: %v\n", err)
        return
    }
    defer file.Close()

    jobs := make(chan Job, *threads*2)
    results := make(chan Result, *threads*2)
    var wg sync.WaitGroup

    // Rate limiter: tick every 1/rate seconds
    rateLimiter := time.Tick(time.Second / time.Duration(*rate))

    // Start workers
    for i := 0; i < *threads; i++ {
        wg.Add(1)
        go worker(*urlFlag, client, jobs, &wg, results, rateLimiter)
    }

    // Feed jobs: each word combined with each extension
    go func() {
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            word := scanner.Text()
            for _, ext := range exts {
                jobs <- Job{Path: word, Ext: ext}
            }
        }
        close(jobs)
    }()

    // Close results after workers finish
    go func() {
        wg.Wait()
        close(results)
    }()

    // Print results
    for res := range results {
        fmt.Printf("[+] %-60s [%d] Size: %d\n", res.FullPath, res.Status, res.Size)
    }
}