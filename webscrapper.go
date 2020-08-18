package main

import (
    "io"
    "log"
    "net/http"
    "os"
    "time"
)

func makeRequest(url string){
  // Create HTTP client with timeout
    client := &http.Client{
        Timeout: 30 * time.Second,
    }

    // Make request
    response, err := client.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    // Copy data from the response to standard output
    n, err := io.Copy(os.Stdout, response.Body)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Number of bytes copied to STDOUT:", n)
}


func main() {
    makeRequest("https://www.instagram.com/firstteamalljerry/?hl=en")
}
