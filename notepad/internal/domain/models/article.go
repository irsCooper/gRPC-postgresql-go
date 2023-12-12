package models

type Article struct {
	Owner   int64
	Title   string
	Content string
	Tags    []string
}
