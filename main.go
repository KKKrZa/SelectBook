package main

import (
	"github.com/KrxkGit/SelectBook/QueryData"
	"github.com/KrxkGit/SelectBook/Sever"
)

func main() {
	QueryData.LoadSecret()
	//QueryData.ReadMajorList()
	//QueryData.ReadBookList()
	Sever.Server()
	//isbnMap := QueryData.QueryByISBN("9787112060221")
	//fmt.Println(isbnMap.Data.Books[0].Title)
}
