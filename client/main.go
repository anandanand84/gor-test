package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	for _ = range time.Tick(time.Second * 2) {
		resp, err := http.Get("http://server:8080/")
		if err == nil {
			fmt.Println("Received response from server %v ", resp.Body)
		} else {
			fmt.Println("Received Error from server %v", err)
		}
	}
}
