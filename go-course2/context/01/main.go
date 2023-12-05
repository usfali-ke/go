package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("Context:", ctx)
	fmt.Printf("Context type: %T\t\n", ctx)
	fmt.Println("Error:", ctx.Err())
	fmt.Println("-------------------------")

	ctx, cancel := context.WithCancel(ctx)
	fmt.Println("Context:", ctx)
	fmt.Printf("Context type: %T\t\n", ctx)
	fmt.Println("cancel:\t", cancel)
	fmt.Println("Error:", ctx.Err())
	fmt.Println("-------------------------")

	cancel()

	fmt.Println("Context:", ctx)
	fmt.Printf("Context type: %T\t\n", ctx)
	fmt.Println("cancel:\t", cancel)
	fmt.Println("Error:", ctx.Err())
	fmt.Println("-------------------------")

}
