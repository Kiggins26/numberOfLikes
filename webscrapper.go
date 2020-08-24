package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
    "io/ioutil"
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
    } else{
        body, err1 := ioutil.ReadAll(response.Body)
        if err1 != nil{
        }else{

        bs := body
        fmt.Println(string(bs));
    }
}
   defer response.Body.Close()
}


func main() {
    makeRequest("https://www.instagram.com/firstteamalljerry/?hl=en")
}
