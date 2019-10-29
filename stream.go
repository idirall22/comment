package comment

import (
	"context"

	"github.com/idirall22/comment/models"
)

// RegisterClientStream register a user to comment stream
func (s *Service) subscribeClientStream(ctx context.Context) *models.ClientStream {

	cs := &models.ClientStream{Comment: make(chan *models.Comment)}

	s.broker.NewClient(cs)

	go func() {
		<-ctx.Done()
		s.broker.RemoveClient(cs)
	}()

	return cs
}