package main

import (
	"github.com/go-liquor/liquor/commands"
	"github.com/go-liquor/liquor/internal/message"
)

func main() {
	if err := commands.ExecuteLiquor(); err != nil {
		message.Fatal("😭 failed to execute liquor: %v", err)
	}
}
