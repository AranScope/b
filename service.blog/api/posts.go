package api

import (
	"fmt"
	"github.com/AranScope/b/service.blog/domain"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func HandleGetPosts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, _ = fmt.Fprintf(w, "%v", domain.ReadAllPosts())
}

