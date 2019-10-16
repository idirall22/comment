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

// UForm sturcture
type UForm struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
}

// ValidateForm validate the form
func (f *UForm) ValidateForm() bool {
	if f.Content == "" || f.ID <= 0 {
		return false
	}
	return true
}
