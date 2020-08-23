package main

import (
	"fmt"
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
    n, err := io.Copy(os.Stdout, response.Body)
    if err != nil {
        log.Fatal(err)
    } else{
    fmt.Println(string(n));
    }
   defer response.Body.Close()
}


func main() {
    makeRequest("https://www.instagram.com/firstteamalljerry/?hl=en")
}
