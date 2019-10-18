package comment

import (
	"context"
	"errors"

	"github.com/idirall22/comment/models"
	u "github.com/idirall22/user"
)

// Add a comment
func (s *Service) addComment(ctx context.Context, form CForm) (*models.Comment, error) {

	if !form.ValidateForm() {
		return nil, ErrorForm
	}

	userID, ok := ctx.Value(u.IDCtx).(int64)
	if !ok {
		return nil, errors.New("Error user id not valid")
	}

	c, err := s.provider.New(ctx, form.Content, userID, form.PostID)

	if err != nil {
		return nil, err
	}

	return c, nil
}

// Update a comment
func (s *Service) updateComment(ctx context.Context, id int64, form CForm) (*models.Comment, error) {

	if !form.ValidateForm() {
		return nil, ErrorForm
	}

	userID, ok := ctx.Value(u.IDCtx).(int64)
	if !ok {
		return nil, errors.New("Error user id not valid")
	}

	return s.provider.Update(ctx, userID, id, form.Content)
}

// Delete a comment
func (s *Service) deleteComment(ctx context.Context, commentID int64) error {

	userID, ok := ctx.Value(u.IDCtx).(int64)
	if !ok {
		return errors.New("Error user id not valid")
	}

	return s.provider.Delete(ctx, userID, commentID)
}
