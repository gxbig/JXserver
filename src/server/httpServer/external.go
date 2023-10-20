package httpServer

import (
	"context"
	"net/http"
)

var Mux *http.ServeMux
var ctx = context.Background()

func init() {
	Mux = http.NewServeMux()

}
