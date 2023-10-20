package util

import (
	"github.com/name5566/leaf/log"
	"net/http"
)

func Write(w http.ResponseWriter, data []byte) {
	_, err := w.Write(data)
	if err != nil {
		log.Error(err.Error())
	}
}
