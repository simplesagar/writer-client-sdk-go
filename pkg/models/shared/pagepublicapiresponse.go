// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/writerai/writer-client-sdk-go/pkg/utils"
	"time"
)

type PagePublicAPIResponseStatus string

const (
	PagePublicAPIResponseStatusLive    PagePublicAPIResponseStatus = "live"
	PagePublicAPIResponseStatusOffline PagePublicAPIResponseStatus = "offline"
)

func (e PagePublicAPIResponseStatus) ToPointer() *PagePublicAPIResponseStatus {
	return &e
}

func (e *PagePublicAPIResponseStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "live":
		fallthrough
	case "offline":
		*e = PagePublicAPIResponseStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PagePublicAPIResponseStatus: %v", v)
	}
}

type PagePublicAPIResponse struct {
	CreatedAt time.Time                   `json:"createdAt"`
	ID        int64                       `json:"id"`
	Order     int64                       `json:"order"`
	Section   *SectionInfo                `json:"section,omitempty"`
	Status    PagePublicAPIResponseStatus `json:"status"`
	Title     string                      `json:"title"`
	UpdatedAt time.Time                   `json:"updatedAt"`
	UpdatedBy *SimpleUser                 `json:"updatedBy,omitempty"`
	URL       string                      `json:"url"`
}

func (p PagePublicAPIResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PagePublicAPIResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PagePublicAPIResponse) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *PagePublicAPIResponse) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *PagePublicAPIResponse) GetOrder() int64 {
	if o == nil {
		return 0
	}
	return o.Order
}

func (o *PagePublicAPIResponse) GetSection() *SectionInfo {
	if o == nil {
		return nil
	}
	return o.Section
}

func (o *PagePublicAPIResponse) GetStatus() PagePublicAPIResponseStatus {
	if o == nil {
		return PagePublicAPIResponseStatus("")
	}
	return o.Status
}

func (o *PagePublicAPIResponse) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *PagePublicAPIResponse) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *PagePublicAPIResponse) GetUpdatedBy() *SimpleUser {
	if o == nil {
		return nil
	}
	return o.UpdatedBy
}

func (o *PagePublicAPIResponse) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}
