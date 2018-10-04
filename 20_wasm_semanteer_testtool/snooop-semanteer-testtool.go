package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"gitlab.com/snooop/snooop-semanteer-testtool/model"
)

// GitCommitHash is the SHA-1 hash of the git commit.
// The variable will be set by the Go linker tool. See
// the Makefile of this project.
var GitCommitHash string

// GOOS=js GOARCH=wasm go build -o snooop-semanteer-testtool.wasm snooop-semanteer-testtool.go
func main() {

	// Misc flags

	var showBuildInfo bool
	flag.BoolVar(&showBuildInfo, "buildInfo", false,
		"Shows build information and exits.")

	flag.Parse()
	if showBuildInfo {
		printVersionInfoToStdOutAndExit()
	}

	fetchSearchResultFromServer()

}

func printVersionInfoToStdOutAndExit() {
	if GitCommitHash != "" {
		fmt.Printf("Git SHA-1 hash of the commit used to build this binary:\n%v\n", GitCommitHash)
	} else {
		fmt.Printf("This binary has not been built using the provided Makefile.\n" +
			"Therefore the Git SHA-1 hash of the commit used to build this\n" +
			"binary is not available.\n")
	}
	os.Exit(0)
}

func fetchSearchResultFromServer() {
	fmt.Printf("\nfetching search result from server...\n")

	resp, _ := http.Get("https://snooop.net/search?q=blumen&sph.fq=sem_active:true&sort=sem_modification_date%20asc&sph.fq=sem_source_id:*,3,*&rows=16&sph.fq=(eventend:([2018-08-22T00:00:00.000Z%20TO%20*])%20OR%20(eventend:[*%20TO%201970-01-02T00:00:00Z]%20AND%20eventstart:[*%20TO%201970-01-02T00:00:00Z])%20OR%20(eventstart:[2018-08-22T00:00:00.000Z%20TO%20*]%20AND%20eventend:[*%20TO%201970-01-02T00:00:00Z]))")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("\nconverting JSON to object...\n")

	var searchResponse model.SearchResponse
	json.Unmarshal(body, &searchResponse)

	// TODO verify IDs

	// key: UIDs value:sem_contact_name
	expectedUIDs := map[string]string{
		"92379": "Kauningerhof",
		"30037": "Die kleine Gärtnerei",
		"256":   "Blatt und Blüte",
		"175":   "Fictive Tile",
		"195":   "Second Fictive Tile",
	}

	actualUIDs := make(map[string]string)

	// copy actual UIDs to map
	for i := 0; i < len(searchResponse.Response.Docs); i++ {
		// fmt.Printf("\n------\n%+v\n-----\n", searchResponse.Response.Docs[i].ID)
		// fmt.Printf(searchResponse.Response.Docs[i].SemContactName)
		actualUIDs[searchResponse.Response.Docs[i].ID] = searchResponse.Response.Docs[i].SemContactName
	}

	for key, value := range actualUIDs {
		fmt.Printf("%v - %v\n", key, value)
	}

	fmt.Printf("\nIDS verifying...\n")
	for expectedUID := range expectedUIDs {
		if _, expectedUIDExists := actualUIDs[expectedUID]; expectedUIDExists {
			fmt.Printf("Expected UID %v does exist in actual IDs.\n", expectedUID)
		} else {
			fmt.Printf("Expected UID %v does not exist in actual IDs.\n", expectedUID)
		}
	}
}
