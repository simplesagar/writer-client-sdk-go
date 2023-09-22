// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package writerclientsdkgo

import (
	"fmt"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"github.com/writerai/writer-client-sdk-go/pkg/utils"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"https://enterprise-api.writer.com",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          *shared.Security
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	Globals           map[string]map[string]map[string]interface{}
	RetryConfig       *utils.RetryConfig
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

type Writer struct {
	// Methods related to AI Content Detector
	AIContentDetector *aiContentDetector
	// Methods related to Billing
	Billing *billing
	// Methods related to CoWrite
	CoWrite *coWrite
	// Methods related to Completions
	Completions *completions
	// Methods related to Content
	Content *content
	// Methods related to Download the customized model
	DownloadTheCustomizedModel *downloadTheCustomizedModel
	// Methods related to Files
	Files *files
	// Methods related to Model Customization
	ModelCustomization *modelCustomization
	// Methods related to Model
	Models *models
	// Methods related to Snippets
	Snippet *snippet
	// Methods related to Styleguide
	Styleguide *styleguide
	// Methods related to Terminology
	Terminology *terminology
	// Methods related to User
	User *user
	// Methods related to document
	Document *document

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*Writer)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Writer) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Writer) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *Writer) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Writer) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *Writer) {
		sdk.sdkConfiguration.Security = &security
	}
}

// WithOrganizationID allows setting the OrganizationID parameter for all supported operations
func WithOrganizationID(organizationID int64) SDKOption {
	return func(sdk *Writer) {
		if _, ok := sdk.sdkConfiguration.Globals["parameters"]["pathParam"]; !ok {
			sdk.sdkConfiguration.Globals["parameters"]["pathParam"] = map[string]interface{}{}
		}

		sdk.sdkConfiguration.Globals["parameters"]["pathParam"]["OrganizationID"] = organizationID
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *Writer) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Writer {
	sdk := &Writer{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.7",
			SDKVersion:        "0.17.0",
			GenVersion:        "2.125.1",
			Globals: map[string]map[string]map[string]interface{}{
				"parameters": {},
			},
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.AIContentDetector = newAIContentDetector(sdk.sdkConfiguration)

	sdk.Billing = newBilling(sdk.sdkConfiguration)

	sdk.CoWrite = newCoWrite(sdk.sdkConfiguration)

	sdk.Completions = newCompletions(sdk.sdkConfiguration)

	sdk.Content = newContent(sdk.sdkConfiguration)

	sdk.DownloadTheCustomizedModel = newDownloadTheCustomizedModel(sdk.sdkConfiguration)

	sdk.Files = newFiles(sdk.sdkConfiguration)

	sdk.ModelCustomization = newModelCustomization(sdk.sdkConfiguration)

	sdk.Models = newModels(sdk.sdkConfiguration)

	sdk.Snippet = newSnippet(sdk.sdkConfiguration)

	sdk.Styleguide = newStyleguide(sdk.sdkConfiguration)

	sdk.Terminology = newTerminology(sdk.sdkConfiguration)

	sdk.User = newUser(sdk.sdkConfiguration)

	sdk.Document = newDocument(sdk.sdkConfiguration)

	return sdk
}
