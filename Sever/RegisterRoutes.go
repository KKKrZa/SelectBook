package Sever

import "net/http"

func RegisterRoutes() {
	http.HandleFunc("/listMajors", HandleListMajors)
	http.HandleFunc("/queryBookList", HandleQueryBookList)
	http.HandleFunc("/queryISBN", HandleQueryIsbn)
}
