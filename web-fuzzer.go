//mini version of a web fuzzer - V1
// support a tiny list for now
//future feature's : worldlist , ...
// Athor : https://github.com/adelkamell


package main

import (
    "fmt"
    "net/http"
)

func main() {
    base := "http://example.com"
    words := []string{"admin", "login", "config", "backup"}
    for _, word := range words {
        url := base + "/" + word
        resp, err := http.Head(url)
        if err != nil {
            continue
        }
        if resp.StatusCode == 200 {
            fmt.Printf("[+] Found: /%s\n", word)
        }
    }
}