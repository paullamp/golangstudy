package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	target_url := "http://10.0.204.105:9999/upload"
	filename := "main.go"
	postFile(filename, target_url)
}

func postFile(filename string, target_url string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("openfile failed")
		return err
	}

	// open file handle
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("open file failed")
		return err
	}
	//iocopy
	_, err := io.Copy(fileWriter, fh)
	if err != nil {
		fmt.Println("file copy failed")
		return
	}

	contenttype := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	// start to post upload
	resp, err := http.Post(target_url, contenttype, bodyBuf)
	if err != nil {
		fmt.Println("post err")
		return err
	}

	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}
