package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/wcsiu/gc-tuning/testecdsa"
)

var ballast []byte

func main() {
	// Create a large heap allocation of 10 GiB
	ballast = make([]byte, 10<<30)

	var err error
	testecdsa.PK, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	http.Handle("/sign", testecdsa.Sign{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
