// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type UploadFileGlobals struct {
	OrganizationID int64 `pathParam:"style=simple,explode=false,name=organizationId"`
}

func (o *UploadFileGlobals) GetOrganizationID() int64 {
	if o == nil {
		return 0
	}
	return o.OrganizationID
}

type UploadFileRequest struct {
	UploadModelFileRequest shared.UploadModelFileRequest `request:"mediaType=multipart/form-data"`
	OrganizationID         *int64                        `pathParam:"style=simple,explode=false,name=organizationId"`
}

func (o *UploadFileRequest) GetUploadModelFileRequest() shared.UploadModelFileRequest {
	if o == nil {
		return shared.UploadModelFileRequest{}
	}
	return o.UploadModelFileRequest
}

func (o *UploadFileRequest) GetOrganizationID() *int64 {
	if o == nil {
		return nil
	}
	return o.OrganizationID
}

type UploadFileResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	ModelFile   *shared.ModelFile
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UploadFileResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UploadFileResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *UploadFileResponse) GetModelFile() *shared.ModelFile {
	if o == nil {
		return nil
	}
	return o.ModelFile
}

func (o *UploadFileResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UploadFileResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
