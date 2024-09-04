package comment_dto

type UpdateCommentRequestDTO struct {
	Content  string `json:"content"`
	ImageURL string `json:"image_url"`
}
