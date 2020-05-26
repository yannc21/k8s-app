package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/sha3"
)

type results struct {
	Result []aQuote
}

type aQuote struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

func quoteOfTheDay(w http.ResponseWriter, req *http.Request) {

	var res results
	//var input io.Reader
	var toReturn string

	// workaround for CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if os.Getenv("OFFLINE_MODE") == "true" {
		// read data.json
		input, err := os.Open("./data.json")

		if err != nil {
			log.Fatal("couldn't opend data.json")
		}

		defer input.Close()

		body, _ := ioutil.ReadAll(input)

		json.Unmarshal(body, &res)
		thequote, _ := json.Marshal(res.Result[0])
		toReturn = string(thequote)
	} else {
		uid := os.Getenv("API_UID")
		token := os.Getenv("API_TOKENID")
		query := "abraham+lincoln"
		searchtype := "AUTHOR"
		url := "https://www.stands4.com/services/v2/quotes.php?uid=" + uid + "&tokenid=" + token + "&searchtype=" + searchtype + "&query=" + query + "&format=json"
		resp, err := http.Get(url)
		fmt.Println(url)

		if err != nil {
			log.Fatal("couldn't get result from API server : ")
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &res)
		thequote, _ := json.Marshal(res.Result[0])
		toReturn = string(thequote)
	}
	fmt.Println(toReturn)
	fmt.Fprintf(w, toReturn)

}

func main() {

	envVar := []string{"API_UID", "API_TOKENID"}
	fqdn := "0.0.0.0"
	port := "8090"

	fmt.Println(sha3.New512())

	if tmp := os.Getenv("FQDN"); tmp != "" {
		fqdn = tmp
	}

	if tmp := os.Getenv("PORT"); tmp != "" {
		port = tmp
	}

	if os.Getenv("OFFLINE_MODE") == "true" {
		fmt.Println("OFFLINE mode enabled !!")
	} else {
		for _, e := range envVar {
			if os.Getenv(e) == "" {
				log.Fatal(e + " environment variable not set !!")
			}
		}
	}
	fmt.Println("will listen to : " + fqdn + ":" +port)
	http.HandleFunc("/quote", quoteOfTheDay)
	http.ListenAndServe(fqdn+":"+port, nil)
}
