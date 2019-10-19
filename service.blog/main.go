package main

import (
    "github.com/AranScope/b/service.blog/api"
    "github.com/AranScope/b/service.blog/domain"
)

func main() {
    domain.Init()
    api.Init()
}
