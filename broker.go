package comment

import "github.com/idirall22/comment/models"

// Broker interface
type Broker interface {
	NewClient(c *models.ClientStream)
	RemoveClient(c *models.ClientStream)
	Brodcast(comment *models.Comment)
}
