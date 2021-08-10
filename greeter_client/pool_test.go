package greeter_client

import (
	"context"
	"testing"

	pool "github.com/processout/grpc-go-pool"
	"google.golang.org/grpc"
)

func TestNew(t *testing.T) {

	p, err := pool.New(func() (*grpc.ClientConn, error) {
		return grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	}, 1, 3, 0)

	if err != nil {
		t.Errorf("The pool returned an error: %s", err.Error())
	}

	// Get a client
	client, err := p.Get(context.Background())
	if err != nil {
		t.Errorf("Get returned an error: %s", err.Error())
	}
	if client == nil {
		t.Error("client was nil")
	}

	// Return the client
	err = client.Close()
	if err != nil {
		t.Errorf("Close returned an error: %s", err.Error())
	}

	if err != pool.ErrAlreadyClosed {
		t.Errorf("Expected error \"%s\" but got \"%s\"",
			pool.ErrAlreadyClosed.Error(), err.Error())
	}

}
