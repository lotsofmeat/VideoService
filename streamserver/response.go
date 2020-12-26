package main 

import (
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, sc int, errMsg string) {
	w.WriteHeader(sc)
	//Return the error message to client
	io.WriteString(w, errMsg)
}
