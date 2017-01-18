package foo

import (
	"gopkg.in/jarcoal/httpmock.v1"
	"testing"
)

func TestSomething(t *testing.T) {
	if addsTwo(10) != 12 {
		t.Errorf("Can't sum!")
	}
}

func TestMock(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "https://api.mybiz.com/articles.json"
	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(200, `[{"id": 1, "name": "My Great Article"}]`))

	article, err := getDataFromJson(url)
	if err != nil {
		t.Error(err)
	}
	if article != "My Great Article" {
		t.Errorf("Wrong article!")
	}
}
