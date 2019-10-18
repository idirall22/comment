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

// Update a comment
func (s *Service) updateComment(ctx context.Context, id int64, form CForm) error {

	if !form.ValidateForm() {
		return ErrorForm
	}

	return s.provider.Update(ctx, id, form.Content)
}

// Delete a comment
func (s *Service) deleteComment(ctx context.Context, commentID int64) error {

	return s.provider.Delete(ctx, commentID)
}
