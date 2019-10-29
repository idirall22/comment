package memory

import (
	"testing"

	"github.com/idirall22/comment/models"
)

var testMemory = &Memory{}

func TestBroker(t *testing.T) {

	r := 0
	done := make(chan bool)
	for i := 0; i < 10; i++ {

		c := &models.ClientStream{
			Comment: make(chan *models.Comment, 1),
			UserID:  int64(i),
			GroupID: 1,
		}

		testMemory.NewClient(c)

		go func(cc *models.ClientStream) {
			<-c.Comment
			r++
			testMemory.RemoveClient(c)

			if r >= 9 && testMemory.GetClientsLength() == 0 {
				done <- true
			}
		}(c)
	}

	comment := &models.Comment{ID: 1, Content: "Comment test 1"}
	testMemory.Brodcast(comment)
	<-done
}
