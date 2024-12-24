package grpctest

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/slewiskelly/grpctest/internal/testdata/echo"
	"github.com/slewiskelly/grpctest/internal/testdata/reverse"

	echopb "github.com/slewiskelly/grpctest/internal/testdata/api/echo/v1"
	reversepb "github.com/slewiskelly/grpctest/internal/testdata/api/reverse/v1"
)

func Example() {
	ctx := context.Background()

	srv := New(
		Register(echopb.RegisterEchoServiceServer, &echo.Server{}),
		Register(reversepb.RegisterReverseServiceServer, &reverse.Server{}),
	)
	defer srv.Stop()

	conn := srv.Conn()
	defer conn.Close()

	e, r := echopb.NewEchoServiceClient(conn), reversepb.NewReverseServiceClient(conn)

	resp, err := e.Echo(ctx, &echopb.Request{Message: "hello, world"})
	if err != nil {
		log.Fatal(err)
	}

	rresp, err := r.Reverse(ctx, &reversepb.Request{Message: resp.GetMessage()})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%q reversed is %q.\n", resp.GetMessage(), rresp.GetMessage())
	// Output: "hello, world" reversed is "dlrow ,olleh".
}

func Test(t *testing.T) {
	srv := New(
		Register(echopb.RegisterEchoServiceServer, &echo.Server{}),
	)
	defer srv.Stop()

	conn := srv.Conn()
	defer conn.Close()

	e := echopb.NewEchoServiceClient(conn)

	resp, err := e.Echo(context.Background(), &echopb.Request{Message: "hello, world"})
	if err != nil {
		t.Errorf("Echo(...) = _, %v", err)
	}

	if got, want := resp.GetMessage(), "hello, world"; got != want {
		t.Errorf("Echo(...) = %q, _; want %q", got, want)
	}
}
