package blogposts

import (
	"io"
	"io/fs"
)

func openPostFile(fs fs.FS, filename string) (io.Reader, error) {
	postFile, err := fs.Open(filename)
	if err != nil {
		return nil, err
	}
	defer postFile.Close()
	return postFile, nil
}

func parse(reader io.Reader) (Post, error) {
	postData, err := io.ReadAll(reader)
	if err != nil {
		return Post{}, err
	}
	return Post{Title: string(postData[7:])}, nil
}
