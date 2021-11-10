package rapidapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Extract() ([]Recipe, error) {
	url := "https://spoonacular-recipe-food-nutrition-v1.p.rapidapi.com/recipes/search?query=burger&&number=30&offset=0&type=main%20course"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "spoonacular-recipe-food-nutrition-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "5264f34aafmsh22e0b76f605486cp151bf6jsn587bc64c2eb3")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	type Body struct {
		Result  []Recipe `json:"results"`
		BaseURI string   `json:"baseUri"`
	}

	var b Body

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	if err := json.Unmarshal(body, &b); err != nil {
		return nil, err
	}

	for i, result := range b.Result {
		b.Result[i].Image = b.BaseURI + result.Image
	}

	return b.Result, nil
}
