package blogposts

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("Test failing FS always fails")
}

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("Title: Post1")},
		"hello-world2.md": {Data: []byte("Title: Post2")},
	}

	posts, err := NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}
	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d", len(posts), len(fs))
	}

	assertPost(Post{Title: "Post1"}, posts[0], t)
	assertPost(Post{Title: "Post2"}, posts[1], t)
}

func assertPost(expected Post, actual Post, t *testing.T) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Wanted %v but got %v", expected, actual)
	}
}

func TestPropagatesError(t *testing.T) {
	fs := StubFailingFS{}

	_, err := NewPostsFromFS(fs)

	if err != nil {
		t.Errorf("Should have propagated error but got nil")
	}
}
