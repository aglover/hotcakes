package models

import "time"

type Word struct {
	WordId       uint         `gorm:"primaryKey" json:"-"`
	Word         string       `json:"word"`
	PartOfSpeech string       `json:"part_of_speech"`
	Definitions  []Definition `json:"definitions"`
	Synonyms     []Synonym    `json:"synonyms,omitempty"`
	CreatedAt    time.Time    `json:"-"`
	UpdatedAt    time.Time    `json:"-"`
}

type Definition struct {
	DefinitionId    uint      `gorm:"primaryKey" json:"-"`
	WordID          uint      `json:"-"`
	Definition      string    `json:"definition"`
	ExampleSentence string    `json:"example_sentence,omitempty"`
	CreatedAt       time.Time `json:"-"`
	UpdatedAt       time.Time `json:"-"`
}

type Synonym struct {
	SynonymId uint      `gorm:"primaryKey" json:"-"`
	WordID    uint      `json:"-"`
	Synonym   string    `json:"synonym"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
