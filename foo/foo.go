package foo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func addsTwo(x int) int {
	return x + 2
}

func getDataFromJson(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	type Article struct {
		ID   int
		Name string
	}
	var articles []Article
	err = json.Unmarshal(body, &articles)
	if err != nil {
		return "", err
	}
	return articles[0].Name, nil
}
