package comment

import (
	"context"

	"github.com/idirall22/comment/models"
)

//Provider interface
type Provider interface {
	New(context.Context, string, int64, int64) (*models.Comment, error)
	List(context.Context, int64, uint, uint) ([]*models.Comment, error)
	Update(context.Context, *models.Comment) error
	Delete(context.Context, int64) error
}
