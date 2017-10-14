package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

func culc(w http.ResponseWriter, r *http.Request) {
	req := decode(r)
	sum := 0
	for _, n := range req.Numbers {
		i, _ := strconv.Atoi(n)
		sum += i
	}
	res, err := encode(Response(Response{sum}))
	if err != nil {
		log.Fatalln(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(res)
}

func main() {
	port := flag.String("p", ":8080", "ポートを指定して下さい")
	flag.Parse()
	http.HandleFunc("/", culc)
	if err := http.ListenAndServe(*port, nil); err != nil {
		log.Fatalln(err)
	}
}
