// web fuzzer - V2
// improvement compare to first version : goroutine and channel, reading from file
// Athor : https://github.com/adelkamell


package main

import (
    "bufio"
    "flag"
    "fmt"
    "net/http"
    "os"
    "sync"
)

func worker(base string, jobs <-chan string, wg *sync.WaitGroup, results chan<- string) {
    defer wg.Done()
    for word := range jobs {
        resp, err := http.Head(base + "/" + word)
        if err != nil {
            continue
        }
        if resp.StatusCode >= 200 && resp.StatusCode < 400 {
            results <- word
        }
    }
}

func main() {
    urlFlag := flag.String("u", "", "Base URL (e.g. http://target.com)")
    wordlist := flag.String("w", "wordlist.txt", "Path to wordlist")
    threads := flag.Int("t", 20, "Number of concurrent threads")
    flag.Parse()

    if *urlFlag == "" {
        fmt.Println("Usage: go run main.go -u <BASE_URL>")
        return
    }

    file, err := os.Open(*wordlist)
    if err != nil {
        fmt.Printf("Error opening wordlist: %v\n", err)
        return
    }
    defer file.Close()

    jobs := make(chan string, *threads)
    results := make(chan string, *threads)
    var wg sync.WaitGroup

    // Start workers
    for i := 0; i < *threads; i++ {
        wg.Add(1)
        go worker(*urlFlag, jobs, &wg, results)
    }

    // Send words to jobs channel
    go func() {
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            jobs <- scanner.Text()
        }
        close(jobs)
    }()

    // Close results when all workers done
    go func() {
        wg.Wait()
        close(results)
    }()

    // Print found paths
    for path := range results {
        fmt.Println("[+] /" + path)
    }
}