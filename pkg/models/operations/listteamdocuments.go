// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

// ListTeamDocumentsSortField
type ListTeamDocumentsSortField string

const (
	ListTeamDocumentsSortFieldTitle            ListTeamDocumentsSortField = "title"
	ListTeamDocumentsSortFieldCreationTime     ListTeamDocumentsSortField = "creationTime"
	ListTeamDocumentsSortFieldModificationTime ListTeamDocumentsSortField = "modificationTime"
	ListTeamDocumentsSortFieldModifiedByMeTime ListTeamDocumentsSortField = "modifiedByMeTime"
	ListTeamDocumentsSortFieldOpenedByMeTime   ListTeamDocumentsSortField = "openedByMeTime"
)

func (e ListTeamDocumentsSortField) ToPointer() *ListTeamDocumentsSortField {
	return &e
}

func (e *ListTeamDocumentsSortField) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "title":
		fallthrough
	case "creationTime":
		fallthrough
	case "modificationTime":
		fallthrough
	case "modifiedByMeTime":
		fallthrough
	case "openedByMeTime":
		*e = ListTeamDocumentsSortField(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListTeamDocumentsSortField: %v", v)
	}
}

// ListTeamDocumentsSortOrder
type ListTeamDocumentsSortOrder string

const (
	ListTeamDocumentsSortOrderAsc  ListTeamDocumentsSortOrder = "asc"
	ListTeamDocumentsSortOrderDesc ListTeamDocumentsSortOrder = "desc"
)

func (e ListTeamDocumentsSortOrder) ToPointer() *ListTeamDocumentsSortOrder {
	return &e
}

func (e *ListTeamDocumentsSortOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = ListTeamDocumentsSortOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListTeamDocumentsSortOrder: %v", v)
	}
}

type ListTeamDocumentsRequest struct {
	Limit          *int                        `queryParam:"style=form,explode=true,name=limit"`
	Offset         *int64                      `queryParam:"style=form,explode=true,name=offset"`
	OrganizationID *int64                      `pathParam:"style=simple,explode=false,name=organizationId"`
	Search         *string                     `queryParam:"style=form,explode=true,name=search"`
	SortField      *ListTeamDocumentsSortField `queryParam:"style=form,explode=true,name=sortField"`
	SortOrder      *ListTeamDocumentsSortOrder `queryParam:"style=form,explode=true,name=sortOrder"`
	TeamID         int64                       `pathParam:"style=simple,explode=false,name=teamId"`
}

type ListTeamDocumentsResponse struct {
	BriefDocuments *shared.BriefDocuments
	ContentType    string
	// Bad Request
	FailResponse *shared.FailResponse
	Headers      map[string][]string
	StatusCode   int
	RawResponse  *http.Response
}
