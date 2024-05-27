package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	titleKey       = "Title: "
	descriptionKey = "Description: "
	tagsKey        = "Tags: "
)

func NewPost(r io.Reader) (Post, error) {
	scanner := bufio.NewScanner(r)

	readStringVal := func(key string) string {
		scanner.Scan()
		return scanner.Text()[len(key):]
	}

	Title := readStringVal(titleKey)
	Description := readStringVal(descriptionKey)
	Tags := strings.Split(readStringVal(tagsKey), ", ")
	// ignore the `---` seperator
	scanner.Scan()

	Body := readBody(scanner)

	return Post{Title, Description, Tags, Body}, nil
}

func readBody(scanner *bufio.Scanner) string {
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}
