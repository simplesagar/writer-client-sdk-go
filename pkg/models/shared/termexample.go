// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type TermExampleType string

const (
	TermExampleTypeGood TermExampleType = "good"
	TermExampleTypeBad  TermExampleType = "bad"
)

func (e TermExampleType) ToPointer() *TermExampleType {
	return &e
}

func (e *TermExampleType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "good":
		fallthrough
	case "bad":
		*e = TermExampleType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TermExampleType: %v", v)
	}
}

type TermExample struct {
	Example    string          `json:"example"`
	ID         *int64          `json:"id,omitempty"`
	TermBankID int64           `json:"termBankId"`
	TermID     int64           `json:"termId"`
	Type       TermExampleType `json:"type"`
}
