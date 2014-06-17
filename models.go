package main

type Post struct {
	Id      int
	Title   string
	Link    string
	Author  string
	Date    string
	Summary string
	Body    string
	Tags    []Tag
}

type Tag struct {
	Name string
}
