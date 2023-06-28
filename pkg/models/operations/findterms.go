// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

// FindTermsPartOfSpeech
type FindTermsPartOfSpeech string

const (
	FindTermsPartOfSpeechNoun      FindTermsPartOfSpeech = "noun"
	FindTermsPartOfSpeechVerb      FindTermsPartOfSpeech = "verb"
	FindTermsPartOfSpeechAdverb    FindTermsPartOfSpeech = "adverb"
	FindTermsPartOfSpeechAdjective FindTermsPartOfSpeech = "adjective"
)

func (e FindTermsPartOfSpeech) ToPointer() *FindTermsPartOfSpeech {
	return &e
}

func (e *FindTermsPartOfSpeech) UnmarshalJSON(data []byte) error {
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
		*e = FindTermsPartOfSpeech(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindTermsPartOfSpeech: %v", v)
	}
}

// FindTermsSortField
type FindTermsSortField string

const (
	FindTermsSortFieldTerm             FindTermsSortField = "term"
	FindTermsSortFieldCreationTime     FindTermsSortField = "creationTime"
	FindTermsSortFieldModificationTime FindTermsSortField = "modificationTime"
	FindTermsSortFieldType             FindTermsSortField = "type"
)

func (e FindTermsSortField) ToPointer() *FindTermsSortField {
	return &e
}

func (e *FindTermsSortField) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "term":
		fallthrough
	case "creationTime":
		fallthrough
	case "modificationTime":
		fallthrough
	case "type":
		*e = FindTermsSortField(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindTermsSortField: %v", v)
	}
}

// FindTermsSortOrder
type FindTermsSortOrder string

const (
	FindTermsSortOrderAsc  FindTermsSortOrder = "asc"
	FindTermsSortOrderDesc FindTermsSortOrder = "desc"
)

func (e FindTermsSortOrder) ToPointer() *FindTermsSortOrder {
	return &e
}

func (e *FindTermsSortOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = FindTermsSortOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindTermsSortOrder: %v", v)
	}
}

// FindTermsType
type FindTermsType string

const (
	FindTermsTypeApproved FindTermsType = "approved"
	FindTermsTypeBanned   FindTermsType = "banned"
	FindTermsTypePending  FindTermsType = "pending"
)

func (e FindTermsType) ToPointer() *FindTermsType {
	return &e
}

func (e *FindTermsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "approved":
		fallthrough
	case "banned":
		fallthrough
	case "pending":
		*e = FindTermsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindTermsType: %v", v)
	}
}

type FindTermsRequest struct {
	Limit          *int64                 `queryParam:"style=form,explode=true,name=limit"`
	Offset         *int64                 `queryParam:"style=form,explode=true,name=offset"`
	OrganizationID *int64                 `pathParam:"style=simple,explode=false,name=organizationId"`
	PartOfSpeech   *FindTermsPartOfSpeech `queryParam:"style=form,explode=true,name=partOfSpeech"`
	SortField      *FindTermsSortField    `queryParam:"style=form,explode=true,name=sortField"`
	SortOrder      *FindTermsSortOrder    `queryParam:"style=form,explode=true,name=sortOrder"`
	Tags           []string               `queryParam:"style=form,explode=true,name=tags"`
	TeamID         int64                  `pathParam:"style=simple,explode=false,name=teamId"`
	Term           *string                `queryParam:"style=form,explode=true,name=term"`
	Type           *FindTermsType         `queryParam:"style=form,explode=true,name=type"`
}

type FindTermsResponse struct {
	ContentType string
	// Bad Request
	FailResponse                    *shared.FailResponse
	Headers                         map[string][]string
	PaginatedResultFullTermWithUser *shared.PaginatedResultFullTermWithUser
	StatusCode                      int
	RawResponse                     *http.Response
}
