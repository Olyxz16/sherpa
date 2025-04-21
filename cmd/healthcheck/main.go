package main 

import (
	"net/http"
	"os"
)

func main() {
    url := "http://127.0.0.1:8080/health"
	_, err := http.Get(url)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
