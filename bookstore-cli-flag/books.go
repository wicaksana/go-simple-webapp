package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error happened ", err)
		os.Exit(1)
	}
}

func getBooks() (books []Book) {
	booksBytes, err := os.ReadFile("./books.json")
	checkError(err)

	err = json.Unmarshal(booksBytes, &books)
	checkError(err)

	return books
}

func handleGetBooks(getCmd *flag.FlagSet, all *bool, id *string) {
	getCmd.Parse(os.Args[2:])

	// checking for `all` or `id`
	if !*all && *id == "" {
		fmt.Println("subcommand --all or --id is needed")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	// if 'all'
	if *all {
		books := getBooks()
		fmt.Printf("Id \t Title \t Author \t Price \t ImageURL \n")
		for _, book := range books {
			fmt.Printf("%v \t %v \t %v \t %v \t %v \n", book.Id, book.Title, book.Author, book.Price, book.Imageurl)
		}
	}

	// if `id`, return only that book. Throw error if book not found.
	if *id != "" {
		books := getBooks()
		fmt.Printf("Id \t Title \t Author \t Price \t ImageURL \n")
		// to check a book exist or not
		var bookFound bool
		for _, book := range books {
			if *id == book.Id {
				bookFound = true
				fmt.Printf("%v \t %v \t %v \t %v \t %v \n", book.Id, book.Title, book.Author, book.Price, book.Imageurl)
			}
		}
		// if no book found with mentioned id throws an error
		if !bookFound {
			fmt.Println("Book not found")
			os.Exit(1)
		}
	}
}
