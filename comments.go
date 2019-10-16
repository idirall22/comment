package comment

import (
	"context"

	"github.com/idirall22/comment/models"
)

// Add a comment
func (s *Service) addComment(ctx context.Context, form CForm) (*models.Comment, error) {

	if !form.ValidateForm() {
		return nil, ErrorForm
	}
	// TODO: get user id from context
	c, err := s.provider.New(ctx, form.Content, 1, form.PostID)

	if err != nil {
		return nil, err
	}

	return c, nil
}
