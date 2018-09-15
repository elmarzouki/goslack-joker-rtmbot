package joker

import (
	"encoding/json"
	"log"
	"net/http"
)

type Joke struct {
	ID         int         `json:"id"`
	Joke       string      `json:"joke"`
	Categories interface{} `json:"categories"`
}

type icndbResponse struct {
	Value Joke   `json:"value"`
	Type  string `json:"type"`
}

func GetJoke() Joke {
	client := http.Client{}
	randomIcndb := "http://api.icndb.com/jokes/random?limitTo=[nerdy]"
	request, err := http.NewRequest("GET", randomIcndb, nil)
	if err != nil {
		log.Fatalln(err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	var result icndbResponse
	json.NewDecoder(response.Body).Decode(&result)
	log.Println(result)

	return result.Value
}
