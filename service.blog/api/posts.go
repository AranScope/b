package api

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func HandleGetPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, _ = fmt.Fprint(w, "Welcome!\n")
}
