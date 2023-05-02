package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

// 4.buat handler test
func TestHandler(t *testing.T) {

	//5.buat handler menggunakan anonymous function
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		//6. buat login webnya
		fmt.Fprint(writer, "Hello World")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func TestServeMux(t *testing.T) {

	//5. buat serve mux nya dlu
	mux := http.NewServeMux()

	//6.tentukan endpoint untuk page nya
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "index page")
	})
	mux.HandleFunc("/about", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "about page")
	})
	mux.HandleFunc("/products/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "product page")
	})
	mux.HandleFunc("/products/details/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "details page")
	})

	//7.buat servernya dan tambahkan mux nya
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	//8.tambahkan listen and servenya
	err := server.ListenAndServe()

	//9.cek apabila error
	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	//5.tambahkan handler
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		//6.cetak request method
		fmt.Fprint(writer, request.Method)
		//7.cetak request uri
		fmt.Fprint(writer, request.RequestURI)
	}

	//8.tambahkan listen and servenya
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	//9.cek apabila error
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
