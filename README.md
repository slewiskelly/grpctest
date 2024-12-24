# grpctest

Package grpctest provides utilities for gRPC testing.

## Usage

```go
package echo

import (
	"context"
	"testing"

	"github.com/slewiskelly/grpctest" // Imports the package for use within the test.

	pb "github.com/slewiskelly/echo/api/v1/echo"
)

func TestEcho(t *testing.T) {
	srv := New( // Initializes a (test) server with the given options.
		Register(pb.RegisterEchoServiceServer, &echo.Server{}), // Registers the service and the implementation to the server.
	)
	defer srv.Stop() // Stops the server, closing all connections, cancelling all active RPCs, and refusing subsequent requests.

	conn := srv.Conn() // Creates a connection to the server.
	defer conn.Close() // Closes the connection to the server.

	e := pb.NewEchoServiceClient(conn) // Creates an Echo service client, using the connection to the server.

	resp, err := e.Echo(context.Background(), &pb.Request{Message: "hello, world"})
	if err != nil {
		t.Errorf("Echo(...) = _, %v", err)
	}

	if got, want := resp.GetMessage(), "hello, world"; got != want {
		t.Errorf("Echo(...) = %q, _; want %q", got, want)
	}
}
```
