package api

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Init() {
	router := httprouter.New()
	router.GET("/posts", HandleGetPosts)
	log.Fatal(http.ListenAndServe(":8080", router))
}
