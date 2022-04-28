package testecdsa

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
	"time"
)

func BenchmarkSign(b *testing.B) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10000; j++ {
			go SendSign()
		}
		time.Sleep(time.Second)
	}
}

func SendSign() {
	client := http.Client{Timeout: time.Second * 2}
	res, err := client.Get("http://localhost:8080/sign")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", body)
}
