package dto

type BookRequest struct {
	Isbn            string  `json:"isbn"`
	Title           string  `json:"title"`
	Genre           string  `json:"genre"`
	PublicationYear string  `json:"publicationYear"`
	CopiesAvailable uint    `json:"copiesAvailable"`
	Price           float32 `json:"price"`
	PublisherId     uint    `json:"publisherId"`
	AuthorIds       []uint  `json:"authorIds"`
}
