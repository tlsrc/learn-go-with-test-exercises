package blogposts

import (
	"io/fs"
)

type Post struct {
	Title string
}

func NewPostsFromFS(filesystem fs.FS) ([]Post, error) {
	dir, _ := fs.ReadDir(filesystem, ".")

	var posts []Post
	for _, file := range dir {
		reader, _ := openPostFile(filesystem, file.Name())
		post, _ := parse(reader)
		posts = append(posts, post)
	}
	return posts, nil
}
