<div align="center">
    <picture>
        <source srcset="https://user-images.githubusercontent.com/6267663/223574357-9a053550-02f9-49f1-b453-1b11db148d7b.svg" media="(prefers-color-scheme: dark)" width="500">
        <img src="https://user-images.githubusercontent.com/6267663/223574369-77805bfe-6d95-44e8-ac48-f9494101e1dc.svg" width="500">
    </picture>
    <h1>Go SDK</h1>
   <p>AI for everyone.</p>
   <a href="https://dev.writer.com/docs"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=000000&style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
  <a href="https://github.com/writerai/writer-client-sdk-go/releases"><img src="https://img.shields.io/github/v/release/writerai/writer-client-sdk-go?sort=semver&style=for-the-badge" /></a>
</div>


<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/writerai/writer-client-sdk-go
```
<!-- End SDK Installation -->

## Authentication

Writer authenticates your API requests using your account’s API keys. If you do not include your key when making an API request, or use one that is incorrect or outdated, Writer returns an error.

Your API keys are available in the account dashboard. We include randomly generated API keys in our code examples if you are not logged in. Replace these with your own or log in to see code examples populated with your own API keys.

<img width="1070" alt="writer-auth" src="https://user-images.githubusercontent.com/6267663/223578295-89087c24-c55a-48bf-b74a-5f057e21e14f.png">

If you cannot see your secret API keys in the Dashboard, this means you do not have access to them. Contact your Writer account owner and ask to be added to their team as a developer.

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"log"
)

func main() {
	s := writerclientsdkgo.New(
		writerclientsdkgo.WithSecurity(""),
		writerclientsdkgo.WithOrganizationID(850421),
	)

	ctx := context.Background()
	res, err := s.Billing.GetSubscriptionDetails(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.SubscriptionPublicResponseAPI != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [.Billing](docs/sdks/billing/README.md)

* [GetSubscriptionDetails](docs/sdks/billing/README.md#getsubscriptiondetails) - Get your organization subscription details

### [.AIContentDetector](docs/sdks/aicontentdetector/README.md)

* [Detect](docs/sdks/aicontentdetector/README.md#detect) - Content detector api

### [.Content](docs/sdks/content/README.md)

* [Check](docs/sdks/content/README.md#check) - Check your content against your preset styleguide.
* [Correct](docs/sdks/content/README.md#correct) - Apply the style guide suggestions directly to your content.

### [.CoWrite](docs/sdks/cowrite/README.md)

* [GenerateContent](docs/sdks/cowrite/README.md#generatecontent) - Generate content using predefined templates
* [ListTemplates](docs/sdks/cowrite/README.md#listtemplates) - Get a list of your existing CoWrite templates

### [.Files](docs/sdks/files/README.md)

* [Delete](docs/sdks/files/README.md#delete) - Delete file
* [Get](docs/sdks/files/README.md#get) - Get file
* [List](docs/sdks/files/README.md#list) - List files
* [Upload](docs/sdks/files/README.md#upload) - Upload file

### [.Models](docs/sdks/models/README.md)

* [List](docs/sdks/models/README.md#list) - List available LLM models

### [.Completions](docs/sdks/completions/README.md)

* [Create](docs/sdks/completions/README.md#create) - Create completion for LLM model
* [CreateModelCustomizationCompletion](docs/sdks/completions/README.md#createmodelcustomizationcompletion) - Create completion for LLM customization model

### [.ModelCustomization](docs/sdks/modelcustomization/README.md)

* [Create](docs/sdks/modelcustomization/README.md#create) - Create model customization
* [Delete](docs/sdks/modelcustomization/README.md#delete) - Delete Model customization
* [Get](docs/sdks/modelcustomization/README.md#get) - Get model customization
* [List](docs/sdks/modelcustomization/README.md#list) - List model customizations

### [.DownloadTheCustomizedModel](docs/sdks/downloadthecustomizedmodel/README.md)

* [FetchFile](docs/sdks/downloadthecustomizedmodel/README.md#fetchfile) - Download your fine-tuned model (available only for Palmyra Base and Palmyra Large)

### [.Document](docs/sdks/document/README.md)

* [Get](docs/sdks/document/README.md#get) - Get document details
* [List](docs/sdks/document/README.md#list) - List team documents

### [.Snippet](docs/sdks/snippet/README.md)

* [Delete](docs/sdks/snippet/README.md#delete) - Delete snippets
* [Find](docs/sdks/snippet/README.md#find) - Find snippets
* [Update](docs/sdks/snippet/README.md#update) - Update snippets

### [.Styleguide](docs/sdks/styleguide/README.md)

* [Get](docs/sdks/styleguide/README.md#get) - Page details
* [ListPages](docs/sdks/styleguide/README.md#listpages) - List your styleguide pages

### [.Terminology](docs/sdks/terminology/README.md)

* [Add](docs/sdks/terminology/README.md#add) - Add terms
* [Delete](docs/sdks/terminology/README.md#delete) - Delete terms
* [Find](docs/sdks/terminology/README.md#find) - Find terms
* [Update](docs/sdks/terminology/README.md#update) - Update terms

### [.User](docs/sdks/user/README.md)

* [List](docs/sdks/user/README.md#list) - List users
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Pagination -->
# Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
<!-- End Pagination -->



<!-- Start Go Types -->

<!-- End Go Types -->



<!-- Start Global Parameters -->
# Global Parameters

A parameter is configured globally. This parameter must be set on the SDK client instance itself during initialization. When configured as an option during SDK initialization, This global value will be used as the default on the operations that use it. When such operations are called, there is a place in each to override the global value, if needed.

For example, you can set `organizationId` to `99895` at SDK initialization and then you do not have to pass the same value on calls to operations like `Detect`. But if you want to do so you may, which will locally override the global setting. See the example code below for a demonstration.


## Available Globals

The following global parameter is available. The required parameter must be set when you initialize the SDK client.

| Name | Type | Required | Description |
| ---- | ---- |:--------:| ----------- |
| organizationId | int64 | ✔️ | The organizationId parameter. |



## Example

```go
package main

import (
	"context"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"log"
)

func main() {
	s := writerclientsdkgo.New(
		writerclientsdkgo.WithSecurity(""),
		writerclientsdkgo.WithOrganizationID(496531),
	)

	contentDetectorRequest := shared.ContentDetectorRequest{
		Input: "string",
	}

	var organizationID *int64 = 592237

	ctx := context.Background()
	res, err := s.AIContentDetector.Detect(ctx, contentDetectorRequest, organizationID)
	if err != nil {
		log.Fatal(err)
	}

	if res.Classes != nil {
		// handle response
	}
}

```
<!-- End Global Parameters -->



<!-- Start Error Handling -->
# Error Handling

Handling errors in your SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.


## Example

```go
package main

import (
	"context"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"log"
)

func main() {
	s := writerclientsdkgo.New(
		writerclientsdkgo.WithSecurity(""),
		writerclientsdkgo.WithOrganizationID(850421),
	)

	ctx := context.Background()
	res, err := s.Billing.GetSubscriptionDetails(ctx)
	if err != nil {

		var e *sdkerrors.FailResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

	}
}

```
<!-- End Error Handling -->



<!-- Start Server Selection -->
# Server Selection

## Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://enterprise-api.writer.com` | None |

For example:

```go
package main

import (
	"context"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"log"
)

func main() {
	s := writerclientsdkgo.New(
		writerclientsdkgo.WithServerIndex(0),
		writerclientsdkgo.WithSecurity(""),
		writerclientsdkgo.WithOrganizationID(850421),
	)

	ctx := context.Background()
	res, err := s.Billing.GetSubscriptionDetails(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.SubscriptionPublicResponseAPI != nil {
		// handle response
	}
}

```


## Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:

```go
package main

import (
	"context"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"log"
)

func main() {
	s := writerclientsdkgo.New(
		writerclientsdkgo.WithServerURL("https://enterprise-api.writer.com"),
		writerclientsdkgo.WithSecurity(""),
		writerclientsdkgo.WithOrganizationID(850421),
	)

	ctx := context.Background()
	res, err := s.Billing.GetSubscriptionDetails(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.SubscriptionPublicResponseAPI != nil {
		// handle response
	}
}

```
<!-- End Server Selection -->



<!-- Start Custom HTTP Client -->
# Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client -->



<!-- Start Authentication -->

# Authentication

## Per-Client Security Schemes

Your SDK supports the following security scheme globally:

| Name     | Type     | Scheme   |
| -------- | -------- | -------- |
| `APIKey` | apiKey   | API key  |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:

```go
package main

import (
	"context"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"log"
)

func main() {
	s := writerclientsdkgo.New(
		writerclientsdkgo.WithSecurity(""),
		writerclientsdkgo.WithOrganizationID(850421),
	)

	ctx := context.Background()
	res, err := s.Billing.GetSubscriptionDetails(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.SubscriptionPublicResponseAPI != nil {
		// handle response
	}
}

```
<!-- End Authentication -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
