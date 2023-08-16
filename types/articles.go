package types

// Article holds article details
type Article struct {
	Title   string `json:"title" bind:"required"`
	Content string `json:"content" bind:"required"`
	Author  string `json:"author" bind:"required"`
}
