package handlers

import (
	"net/http"
)


func APIModeResponse(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Api mode is active. Static pages are not rendered."))
}
