package domain

type Book struct {
	ID          int    `json:"id"`           // شناسه یکتا برای کتاب
	Title       string `json:"title"`        // عنوان کتاب
	Author      string `json:"author"`       // نویسنده کتاب
	PublishedAt string `json:"published_at"` // تاریخ انتشار کتاب
}
