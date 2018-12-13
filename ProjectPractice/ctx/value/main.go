package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "id", 123)
	ctx = context.WithValue(ctx, "session_id", "jophdbjkfsojhjefsdfsfd")
	process(ctx)
}

func process(ctx context.Context) {
	id, ok := ctx.Value("id").(int)
	if !ok {
		fmt.Printf("convert format is err")
		id = 10203
	}
	session_id, _ := ctx.Value("session_id").(string)
	fmt.Printf("id: %d, sessid: %v", id, session_id)
}
