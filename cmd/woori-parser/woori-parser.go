package main

import (
	"fmt"

	"github.com/kangseokgyu/auto-gagebu/internal/pkg/wooricard"
)

func main() {
	receipts := wooricard.GetReceipts()
	fmt.Println(receipts)
}
