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
	feedback(dfn)
	// temporary feedback mechanism
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

	// appID and appKey
	// is initialised in local (non-VCS) file
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

func feedback(dfn definition) {
	// logPath
	// is initialised in local (non-VCS) file
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0640)
	handleErr(err)
	defer f.Close()

	log.SetOutput(f)
	log.Printf("%-20s: %s", dfn.Results[0].Word,
		dfn.Results[0].LexicalEntries[0].Entries[0].Senses[0].Definitions[0])
}

func handleErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
