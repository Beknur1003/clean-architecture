package author

type CreateAuthorDTO struct {
	Name string `json:"name"`
	Year int    `json:"year"`
}

type UpdateBookDTO struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Year int    `json:"year"`
}
