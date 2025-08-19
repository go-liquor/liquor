package main

import (
	"fmt"
	"os"

	"github.com/go-liquor/liquor/v3/internal/commands"
	"github.com/logrusorgru/aurora/v4"
)

func main() {
	if err := commands.Execute(); err != nil {
		fmt.Printf("🚫 %v\n", aurora.Red(err.Error()))
		os.Exit(1)
	}
}
