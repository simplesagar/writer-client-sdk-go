// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CreateTermsRequestFailHandlingEnum string

const (
	CreateTermsRequestFailHandlingEnumAccumulate   CreateTermsRequestFailHandlingEnum = "accumulate"
	CreateTermsRequestFailHandlingEnumValidate     CreateTermsRequestFailHandlingEnum = "validate"
	CreateTermsRequestFailHandlingEnumSkip         CreateTermsRequestFailHandlingEnum = "skip"
	CreateTermsRequestFailHandlingEnumValidateOnly CreateTermsRequestFailHandlingEnum = "validateOnly"
)

func (e CreateTermsRequestFailHandlingEnum) ToPointer() *CreateTermsRequestFailHandlingEnum {
	return &e
}

func (e *CreateTermsRequestFailHandlingEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "accumulate":
		fallthrough
	case "validate":
		fallthrough
	case "skip":
		fallthrough
	case "validateOnly":
		*e = CreateTermsRequestFailHandlingEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateTermsRequestFailHandlingEnum: %v", v)
	}
}

type CreateTermsRequest struct {
	FailHandling *CreateTermsRequestFailHandlingEnum `json:"failHandling,omitempty"`
	Models       []TermCreate                        `json:"models,omitempty"`
}
