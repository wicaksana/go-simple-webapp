package main

import (
	"flag"
	"fmt"
	"os"
)

type Book struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Price    string `json:"price"`
	Imageurl string `json:"image_url"`
}

func main() {
	//////////////////////////////////////////////////////////////////////
	// `get` argument.
	//////////////////////////////////////////////////////////////////////
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getAll := getCmd.Bool("all", false, "List all the books")
	getId := getCmd.String("id", "", "Get book by ID")

	// validation
	if len(os.Args) < 2 {
		fmt.Println("Expected get, add, update, delete commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		handleGetBooks(getCmd, getAll, getId)
	default:
		fmt.Println("Please provide get, update, update, delete commands")
		os.Exit(1)
	}
}
