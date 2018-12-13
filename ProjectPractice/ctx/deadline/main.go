package main

import (
	"time"
	"context"
	"fmt"
)

func main() {
	d := time.Now().Add(50 * time.Microsecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	select {
	case <- time.After(1* time.Microsecond):
		fmt.Printf("the One")
	case <- ctx.Done():
		fmt.Println(ctx.Err())
	}
}
