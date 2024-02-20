package Sever

import (
	"encoding/json"
	"github.com/KrxkGit/SelectBook/QueryData"
	"log"
	"net/http"
	"strconv"
)

func HandleListMajors(w http.ResponseWriter, r *http.Request) {
	QueryData.ReadMajorList()
	b, err := json.Marshal(QueryData.Majors)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	w.Write(b)
}

func HandleQueryBookList(w http.ResponseWriter, r *http.Request) {
	// 接受 Post 请求
	grade, err := strconv.Atoi(r.FormValue("grade"))
	if err != nil {
		log.Printf(err.Error())
	}
	majorId := r.FormValue("majorId")
	log.Printf("%v %v\n", grade, majorId)

	// 设置查询参数
	QueryData.MySelect.SetQueryGrade(grade)
	QueryData.MySelect.SetQueryMajorId(majorId)

	// 查询书籍
	QueryData.ReadBookList()
	b, err := json.Marshal(QueryData.Books)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	w.Write(b)

}
func HandleQueryIsbn(w http.ResponseWriter, r *http.Request) {
	// 接受 Get 请求
	query := r.URL.Query()
	isbn := query.Get("isbn")
	//fmt.Println(isbn)

	// 查询ISBN 对应的信息
	isbnMap := QueryData.QueryByISBN(isbn)
	b, err := json.Marshal(isbnMap)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	w.Write(b)
}
