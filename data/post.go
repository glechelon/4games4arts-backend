package data

type Post struct {
	Id      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Image   string `json:"image"`
	Teaser  string `json:"teaser"`
	Content string `json: "content"`
}
