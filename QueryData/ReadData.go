package QueryData

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func ReadMajorList() {
	url := fmt.Sprintf("%s?accessToken=%s", majorList, secret.AccessTokenQu)
	method := "POST"

	payload := strings.NewReader(`{"termCode":23242}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf(err.Error())
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 手动清空原来(若返回请求体无该项，json 不会扫描该项而导致不会更新)
	Majors.Data = nil
	err = json.Unmarshal(body, &Majors)
	if err != nil {
		log.Printf(err.Error())
	}
	//fmt.Println(Majors.Data[1].MajorName, Majors.Data[1].MajorId, Majors.Data[1].Grade)
}

func ReadBookList() {
	url := fmt.Sprintf("%s?accessToken=%s", bookList, secret.AccessTokenQu)
	method := "POST"

	mySelectByte, _ := json.Marshal(MySelect)

	payload := bytes.NewReader(mySelectByte)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf(err.Error())
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 手动清空原来(若返回请求体无该项，json 不会扫描该项而导致不会更新)
	Books.Data.List = nil
	err = json.Unmarshal(body, &Books)
	if err != nil {
		log.Printf(err.Error())
	}
	//for key, value := range Books.Data.List {
	//	fmt.Println(key, value.CourseName, value.PressName)
	//}
}
