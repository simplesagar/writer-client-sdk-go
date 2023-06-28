// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type TermMistakePos string

const (
	TermMistakePosNoun      TermMistakePos = "noun"
	TermMistakePosVerb      TermMistakePos = "verb"
	TermMistakePosAdverb    TermMistakePos = "adverb"
	TermMistakePosAdjective TermMistakePos = "adjective"
)

func (e TermMistakePos) ToPointer() *TermMistakePos {
	return &e
}

func (e *TermMistakePos) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "noun":
		fallthrough
	case "verb":
		fallthrough
	case "adverb":
		fallthrough
	case "adjective":
		*e = TermMistakePos(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TermMistakePos: %v", v)
	}
}

type TermMistake struct {
	CaseSensitive bool            `json:"caseSensitive"`
	ID            *int64          `json:"id,omitempty"`
	Mistake       string          `json:"mistake"`
	Pos           *TermMistakePos `json:"pos,omitempty"`
	TermBankID    int64           `json:"termBankId"`
	TermID        int64           `json:"termId"`
}
