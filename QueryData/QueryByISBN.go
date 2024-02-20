package QueryData

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func QueryByISBN(isbn string) (isbnMap IsbnMapInfo) {
	url := fmt.Sprintf("%s?from=0&size=10&keyword=%s&categoryId=&withTotal=1", isbnList, isbn)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Access-Token", secret.AccessTokenIsbn)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf(err.Error())
			return
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(body, &isbnMap)
	if err != nil {
		log.Printf(err.Error())
	}
	return
}
