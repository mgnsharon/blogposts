package blogposts_test

import (
	"bytes"
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/mgnsharon/blogposts"
)

const (
	firstBody = `Title: Post 1
Description: Description 1
Tags: t1, t2
---
The body is here it

is very informative.`
	secondBody = `Title: Post 2
Description: Description 2
Tags: tag1, tag2
---
Here is body number 2`
)

var tfs = fstest.MapFS{
	"hello world.md":  {Data: []byte(firstBody)},
	"hello-world2.md": {Data: []byte(secondBody)},
}

func TestNewBlogPosts(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		posts, err := blogposts.NewPostsFromFS(tfs)
		if err != nil {
			t.Fatal(err)
		}

		assertPost(t, posts[0], blogposts.Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"t1", "t2"},
			Body: `The body is here it

is very informative.`,
		})
	})

	t.Run("filesystem error", func(t *testing.T) {
		_, err := blogposts.NewPostsFromFS(StubFailingFS{})
		if err == nil {
			t.Error("expected an error to be returned")
		}
	})
}

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("i am constantly failing")
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func BenchmarkNewPost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		blogposts.NewPost(bytes.NewReader([]byte(firstBody)))
	}
}
