package goweb

import (
	"net/http"
	"testing"
)

// 1.buat server menggunakan net/http
func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
	}

	// 2.jalankan server menggunakan listen and serve
	err := server.ListenAndServe()

	// 3.cek err
	if err != nil {
		panic(err)
	}
}
