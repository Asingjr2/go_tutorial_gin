package main

// Type created for demo example.
type article struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
}

// Array of articles.  Gin does not come with db integration out the box.
var articleList = []article{
	article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// Simply returns list of articles.  If db was attached we could fetch real db items.
func getAllArticles() []article {
	return articleList
}
