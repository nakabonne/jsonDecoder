package main

import "encoding/json"

type Response struct {
	Sum int `json:"sum"`
}

func encode(r Response) ([]byte, error) {
	return json.Marshal(r)
}
