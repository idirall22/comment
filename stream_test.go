package comment

import (
	"context"
	"testing"
	"time"

	"github.com/idirall22/comment/brokers/memory"
)

// Test register client to stream
func testSubscribeClientStream(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())

	clientsLenght := testService.broker.(*memory.Memory).GetClientsLength

	testService.subscribeClientStream(ctx, 1)

	if clientsLenght() != 1 {
		t.Errorf("Clients length should be 1 but got %d", clientsLenght())
	}

	cancel()

	time.Sleep(time.Millisecond)

	defer func() {
		if clientsLenght() != 0 {
			t.Errorf("Clients length should be 0 but got %d", clientsLenght())
		}
	}()
}
