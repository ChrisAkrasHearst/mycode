/* Alta3 Research | RZFeeser
   HTTP GET requests  */
   
package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "encoding/json"
)

type Magic struct {
	Card struct {
		Name          string   `json:"name"`
		ManaCost      string   `json:"manaCost"`
		Cmc           float64  `json:"cmc"`
		Colors        []string `json:"colors"`
		ColorIdentity []string `json:"colorIdentity"`
		Type          string   `json:"type"`
		Supertypes    []string `json:"supertypes"`
		Types         []string `json:"types"`
		Subtypes      []string `json:"subtypes"`
		Rarity        string   `json:"rarity"`
		Set           string   `json:"set"`
		SetName       string   `json:"setName"`
		Text          string   `json:"text"`
		Artist        string   `json:"artist"`
		Number        string   `json:"number"`
		Power         string   `json:"power"`
		Toughness     string   `json:"toughness"`
		Layout        string   `json:"layout"`
		Multiverseid  string   `json:"multiverseid"`
		ImageURL      string   `json:"imageUrl"`
		Watermark     string   `json:"watermark"`
		Rulings       []struct {
			Date string `json:"date"`
			Text string `json:"text"`
		} `json:"rulings"`
		ForeignNames []struct {
			Name        string      `json:"name"`
			Text        string      `json:"text"`
			Type        string      `json:"type"`
			Flavor      interface{} `json:"flavor"`
			ImageURL    string      `json:"imageUrl"`
			Language    string      `json:"language"`
			Identifiers struct {
				ScryfallID   string `json:"scryfallId"`
				MultiverseID int    `json:"multiverseId"`
			} `json:"identifiers"`
			Multiverseid int `json:"multiverseid"`
		} `json:"foreignNames"`
		Printings    []string `json:"printings"`
		OriginalText string   `json:"originalText"`
		OriginalType string   `json:"originalType"`
		Legalities   []struct {
			Format   string `json:"format"`
			Legality string `json:"legality"`
		} `json:"legalities"`
		ID string `json:"id"`
	} `json:"card"`
}



func main() {

    // create a GET request to the page
    resp, err := http.Get("https://api.magicthegathering.io/v1/cards/386616")

    // check for errors
    if err != nil {
        log.Fatal(err)
    }

    // the client must close the response body when finished
    defer resp.Body.Close()

    // read the content of the body with "ReadAll()"
    //body, err := io.ReadAll(resp.Body)

    //if err != nil {
    //    log.Fatal(err)
    //}

    // print received data to the console
    //fmt.Println(string(body))

    var Mymagic Magic

    byteValue, _ := io.ReadAll(resp.Body)

    json.Unmarshal(byteValue, &Mymagic)

    fmt.Println(Mymagic)

}

