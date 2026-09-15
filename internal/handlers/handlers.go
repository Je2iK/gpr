package handlers

import (
	"net/http"
)

func HelloGo(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello, Go!"))
}