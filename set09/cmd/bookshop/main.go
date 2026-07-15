package main

import (
	"fmt"
	"os"

	"learn-go-by-exercises/set09/catalog"
	"learn-go-by-exercises/set09/pricing"
)

func main() {
	c := catalog.New([]catalog.Book{{
		ISBN: "978-0-13-419044-0", Title: "The Go Programming Language",
		Author: "Donovan and Kernighan", PriceCents: 3999,
	}})
	if len(os.Args) != 2 {
		fmt.Println("usage: bookshop <isbn>")
		return
	}
	book, ok := c.Find(os.Args[1])
	if !ok {
		fmt.Println("book not found")
		return
	}
	fmt.Printf("%s: $%.2f\n", book.Title, float64(pricing.FinalPrice(*book, 10))/100)
}
