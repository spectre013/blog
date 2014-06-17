package main

import (
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"net/http"
	//"strings"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tags := make([]Tag, 0)
	tags = append(tags, Tag{Name: "GOLANG"})

	posts := new(Post)
	posts.Title = "This is a Test"
	posts.Link = "this-test"
	posts.Author = "Brian"
	posts.Date = "Sunday, June 16, 2014"
	posts.Id = 1
	posts.Summary = "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation test link ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate another link velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit..."
	posts.Tags = tags
	Render(w, "index", posts)
}

func adminIndex(w http.ResponseWriter, r *http.Request) {
	Render(w, "index", "Test")
}

/* UTILITY FUNCATIONS */
func Json(object interface{}) []byte {
	j, err := json.Marshal(object)
	if err != nil {
		fmt.Println("error:", err)
	}
	return j
}
