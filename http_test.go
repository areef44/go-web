package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World")
}

// gunanya package ini adalah untuk menjalankan testing handler lagsung tanpa menjalankan aplikasi webnya
func TestHttp(t *testing.T) {
	//5.buat new request yang berisi method,url body
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)

	//6.tambahkan recorder gunanya untuk menangkap response body response status dsb
	recorder := httptest.NewRecorder()

	//7.panggil function yang telah dibuat
	HelloHandler(recorder, request)

	//8.untuk mengecek hasilnya tambahkan code ini
	response := recorder.Result()
	//pake io.ReadAll untuk membaca semua response body
	body, _ := io.ReadAll(response.Body)
	//konversi menjadi string hasilnya
	bodyString := string(body)

	fmt.Println(bodyString)
}
