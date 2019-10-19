package domain

import "time"
import "github.com/monzo/gocassa"

type Post struct {
	Id string
	Author string
	Title string
	Body string
	Created time.Time
	Updated time.Time
}

var postsTable gocassa.Table

func Init() {
	keySpace, err := gocassa.ConnectToKeySpace("blog", []string{"cassandra:9042"}, "", "")
	if err != nil {
		panic(err)
	}
	postsTable = keySpace.Table("posts", &Post{}, gocassa.Keys{
		PartitionKeys: []string{"Id"},
	})
}

func ReadAllPosts() []*Post {
	var posts []*Post
	if err := postsTable.Where(gocassa.Eq("Id", "*")).Read(&posts).Run(); err != nil {
		panic(err)
	}
	return posts
}
