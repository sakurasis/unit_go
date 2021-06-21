package main

import (
	"code.local/json/githubIssue"
	"fmt"
	"log"
	"os"
)

func main() {
	issues, err := githubIssue.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d  issues:\n", issues.TotalCount)
	for _, item := range issues.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

