package comment

// CForm sturcture
type CForm struct {
	Content string `json:"content"`
	PostID  int64  `json:"post_id"`
}

// ValidateForm validate the form
func (f *CForm) ValidateForm() bool {
	if f.Content == "" || f.PostID <= 0 {
		return false
	}
	return true
}
