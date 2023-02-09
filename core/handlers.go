package core

import "net/http"

func helloHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello!"))
}

func (d *Daemon) RegisterHandlers() {
	http.HandleFunc("/hello", helloHandler)
}
