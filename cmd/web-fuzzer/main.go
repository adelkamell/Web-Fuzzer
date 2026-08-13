package main

import (
	"fmt"
	"os"

	"github.com/adelkamell/Web-Fuzzer/pkg/fuzzer"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: web-fuzzer <url> <wordlist_file>")
		os.Exit(1)
	}
	// در اینجا کد خواندن فایل wordlist را اضافه کنید
	wordlist := []string{"admin", "login", "test"} // نمونه
	f := fuzzer.New(os.Args[1], wordlist)
	f.Run()
}