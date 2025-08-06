package web

type ArticleUpdateRequest struct {
	Id      int    `json:"id"`
	Title   string `validate:"required,min=1,max=255" json:"title"`
	Content string `validate:"required,min=1,max=5000" json:"content"`
}
