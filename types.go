package main

type definition struct {
	Metadata metadata `json:"metadata"`
	Results  []result `json:"results"`
}

type metadata struct {
	Provider string `json:"provider"`
}

type result struct {
	Word           string         `json:"word"`
	LexicalEntries []lexicalEntry `json:"lexicalEntries"`
}

type lexicalEntry struct {
	Entries  []entry `json:"entries"`
	Category string  `json:"lexicalCategory"`
}

type entry struct {
	Senses []sense `json:"senses"`
}

type sense struct {
	Definitions      []string `json:"definitions"`
	ShortDefinitions []string `json:"short_definitions"`
}
