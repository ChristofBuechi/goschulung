package main

import (
	"fmt"
	"net/http"
	"crypto/md5"
	"io"
	"strconv"
	"time"
	"html/template"
	"io/ioutil"
)

func upload(w http.ResponseWriter, request *http.Request) {
	fmt.Println("method:", request.Method)
	if request.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		request.ParseMultipartForm(32 << 20)
		file, _, err := request.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		f, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprintf(w, "%s!", process(f))

	}
}

func process(s []byte) string {
	return string(s)
}

//func process(w http.ResponseWriter, request *http.Request) {
//	buf := new(bytes.Buffer)
//	buf.ReadFrom(request.Body)
//	s := buf.String()
//
//	fmt.Fprintf(w, "%s!", s)
//}

func start(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "<html><head><title>Upload file</title></head><body>" +
	"<form enctype=\"multipart/form-data\" action=\"http://127.0.0.1:8080/upload\" method=\"post\">" +
	"<input type=\"file\" name=\"uploadfile\" />" +
	"<input type=\"hidden\" name=\"token\" value=\"{{.}}\"/><input type=\"submit\" value=\"upload\" />" +
	"</form></body></html>")
}

func main() {

	http.HandleFunc("/", start)
	http.HandleFunc("/upload", upload)
	panic(http.ListenAndServe(":8080", nil))
}