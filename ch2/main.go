package main

import (
	"log"
	"os"

	_ "github.com/achristie/goinaction/ch2/matchers"
	"github.com/achristie/goinaction/ch2/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("discharge")
}
