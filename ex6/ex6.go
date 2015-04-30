package main

// BEGIN IMPORTS OMIT
import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

// BEGIN APKEY OMIT
const APIKEY string = "123"

// END APKEY OMIT

// END IMPORTS OMIT

// BEGIN MAIN OMIT
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/lookup", lookupHandler)

	http.ListenAndServe("0.0.0.0:9999", mux)
}

// END MAIN OMIT

// BEGIN ROOT OMIT
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Rendering root")
	rootTemplate := template.Must(template.ParseFiles("./template/root.html"))
	if err := rootTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

// END ROOT OMIT

// BEGIN LOOKUP OMIT
func lookupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Rendering Search Results")

	r.ParseForm()
	query := r.Form["q"][0]
	results := SearchFor(query)

	data := struct {
		Results []Result
	}{results}

	resultTemplate := template.Must(template.ParseFiles("./template/result.html"))
	if err := resultTemplate.Execute(w, data); err != nil {
		panic(err)
	}
}

// END LOOKUP OMIT

// BEGIN RESULTWRAP OMIT
type ResultSet struct {
	Entries []Result `json:"handelsnaam"`
}

type Result struct {
	Name  string `json:"text"`
	Extra struct {
		id string `json:"id"`
	} `json:"extra"`
}

// END RESULTWRAP OMIT

// BEGIN API OMIT
func SearchFor(q string) []Result {
	url := "https://overheid.io/suggest/kvk/" + q

	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("ovio-api-key", APIKEY)

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	results := new(ResultSet)

	json.NewDecoder(resp.Body).Decode(results)

	return results.Entries
}

// END API OMIT
