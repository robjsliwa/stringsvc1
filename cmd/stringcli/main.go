package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/robjsliwa/stringsvc1"
	"github.com/robjsliwa/stringsvc1/client/grpc"
	"google.golang.org/grpc"
)

func main() {
	var (
		grpcAddr = flag.String("addr", ":8081",
			"gRPC address")
	)
	flag.Parse()
	ctx := context.Background()
	conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(),
		grpc.WithTimeout(1*time.Second))
	if err != nil {
		log.Fatalln("gRPC dial:", err)
	}
	defer conn.Close()
	stringService := grpcclient.New(conn)
	args := flag.Args()
	var cmd string
	cmd, args = pop(args)
	switch cmd {
	case "uppercase":
		var inputString string
		inputString, _ = pop(args)
		uppercase(ctx, stringService, inputString)
	case "count":
		var inputString string
		inputString, _ = pop(args)
		count(ctx, stringService, inputString)
	default:
		log.Fatalln("unknown command", cmd)
	}
}

func uppercase(ctx context.Context, service stringsvc1.StringService, inputString string) {
	uppercasedString, err := service.Uppercase(ctx, inputString)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(uppercasedString)
}

func count(ctx context.Context, service stringsvc1.StringService, inputString string) {
	count, err := service.Count(ctx, inputString)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(count)
}

func pop(s []string) (string, []string) {
	if len(s) == 0 {
		return "", s
	}
	return s[0], s[1:]
}
