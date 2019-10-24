package comment

import (
	"context"

	"github.com/idirall22/comment/models"
	"github.com/idirall22/utilities"
)

// Add a comment
func (s *Service) addComment(ctx context.Context, form CForm) (*models.Comment, error) {

	userID, err := utilities.GetUserIDFromContext(ctx)

	if err != nil {
		return nil, err
	}

	if !form.ValidateForm() {
		return nil, ErrorForm
	}

	c, err := s.provider.New(ctx, form.Content, userID, form.PostID)

	if err != nil {
		return nil, err
	}

	return c, nil
}

// Update a comment
func (s *Service) updateComment(ctx context.Context, id int64, form CForm) (*models.Comment, error) {

	userID, err := utilities.GetUserIDFromContext(ctx)

	if err != nil {
		return nil, err
	}

	if !form.ValidateForm() {
		return nil, ErrorForm
	}

	return s.provider.Update(ctx, userID, id, form.Content)
}

// Delete a comment
func (s *Service) deleteComment(ctx context.Context, commentID int64) error {

	userID, err := utilities.GetUserIDFromContext(ctx)

	if err != nil {
		return err
	}

	return s.provider.Delete(ctx, userID, commentID)
}
