package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	url = "https://od-api.oxforddictionaries.com/api/v1/entries/en/"
)

func main() {
	wordToDefine := extractWordFromInput()
	dfn := queryDefintion(wordToDefine)
	fmt.Println(dfn.Results[0].LexicalEntries[0].Entries[0].Senses[0].Definitions[0])
}

func extractWordFromInput() (wordToDefine string) {
	if len(os.Args) > 1 {
		wordToDefine = os.Args[1]
	} else {
		log.Fatal(fmt.Sprintf("usage: %s <word-to-define>", os.Args[0]))
	}
	return
}

func queryDefintion(wordToDefine string) definition {
	query := url + strings.ToLower(wordToDefine)

	definerClient := http.Client{}

	req, err := http.NewRequest(http.MethodGet, query, nil)
	handleErr(err)

	// APP_ID and APP_KEY
	// are initialised in local (non-VCS) file
	req.Header.Set("app_id", appID)
	req.Header.Set("app_key", appKey)

	res, err := definerClient.Do(req)
	handleErr(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	handleErr(err)

	dfn := definition{}
	handleErr(json.Unmarshal(body, &dfn))
	return dfn
}

func handleErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
