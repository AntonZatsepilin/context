package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second*3)
	ctx = context.WithValue(ctx, "id", 1)

	// ctx, cancel := context.WithTimeout(ctx, time.Second*3)

	// go func() {
	// 	time.Sleep(time.Millisecond * 100)
	// 	cancel()
	// }()

	parse(ctx)
}

func parse(ctx context.Context) {
	id := ctx.Value("id")
	fmt.Println("parsing with id", id.(int))
	for {
		select {
		case <-time.After(time.Second * 2):
			fmt.Println("parse complete")
			return
		case <-ctx.Done():
			fmt.Println("dedline exceeded")
			return
		}
	}
}
