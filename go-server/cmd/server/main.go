package main

import (
	"context"
	"fmt"
	"os"

	"github.com/thevzurd/flutter-grpc-tutorial-master/vendor/protocol/grpc"
	"github.com/thevzurd/flutter-grpc-tutorial-master/vendor/service/v1"
)

func main() {
	if err := grpc.RunServer(context.Background(), v1.NewChatServiceServer(), "3000"); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
