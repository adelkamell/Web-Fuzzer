package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/adelkamell/Web-Fuzzer/pkg/fuzzer"
)

func main() {
	urlFlag := flag.String("u", "", "Base URL (e.g. http://target.com)")
	wordlist := flag.String("w", "wordlist.txt", "Wordlist file")
	threads := flag.Int("t", 20, "Number of threads")
	extensions := flag.String("x", "", "Comma-separated extensions (e.g. php,asp)")
	rate := flag.Int("rate", 10, "Requests per second")
	timeout := flag.Int("timeout", 5, "HTTP timeout in seconds")
	userAgent := flag.String("ua", "Mozilla/5.0 (Windows NT 10.0; rv:102.0) Gecko/20100101 Firefox/102.0", "Custom User-Agent")
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

	opts := &fuzzer.Options{
		BaseURL:    *urlFlag,
		Wordlist:   *wordlist,
		Threads:    *threads,
		Extensions: exts,
		Rate:       *rate,
		Timeout:    time.Duration(*timeout),
		UserAgent:  *userAgent,
	}

	f := fuzzer.NewFuzzer(opts)
	f.Run()
}