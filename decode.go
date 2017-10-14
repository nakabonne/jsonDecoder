package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Request struct {
	Numbers []string `json:"numbers"`
}

func decode(r *http.Request) *Request {
	bytes, _ := ioutil.ReadAll(r.Body)                      // io.ReadCloser型のボディをバイト列に変換
	request := new(Request)                                 // デコード後に格納する箱
	if err := json.Unmarshal(bytes, &request); err != nil { // デコードする
		log.Println(err)
	}
	return request
}
