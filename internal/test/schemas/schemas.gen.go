// Package schemas provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package schemas

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// N5StartsWithNumber defines model for 5StartsWithNumber.
type N5StartsWithNumber map[string]interface{}

// AnyType1 defines model for AnyType1.
type AnyType1 interface{}

// AnyType2 defines model for AnyType2.
type AnyType2 interface{}

// CustomStringType defines model for CustomStringType.
type CustomStringType string

// GenericObject defines model for GenericObject.
type GenericObject map[string]interface{}

// NullableProperties defines model for NullableProperties.
type NullableProperties struct {
	Optional            *string `json:"optional,omitempty"`
	OptionalAndNullable *string `json:"optionalAndNullable"`
	Required            string  `json:"required" validate:"required"`
	RequiredAndNullable *string `json:"requiredAndNullable" validate:"required"`
}

// StringInPath defines model for StringInPath.
type StringInPath string

// Issue185JSONBody defines parameters for Issue185.
type Issue185JSONBody NullableProperties

// Issue9JSONBody defines parameters for Issue9.
type Issue9JSONBody interface{}

// Issue9Params defines parameters for Issue9.
type Issue9Params struct {
	Foo string `json:"foo" validate:"required"`
}

// Issue185RequestBody defines body for Issue185 for application/json ContentType.
type Issue185JSONRequestBody Issue185JSONBody

// Issue9RequestBody defines body for Issue9 for application/json ContentType.
type Issue9JSONRequestBody Issue9JSONBody

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestEditor RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = http.DefaultClient
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditor = fn
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// EnsureEverythingIsReferenced request
	EnsureEverythingIsReferenced(ctx context.Context) (*http.Response, error)

	// Issue127 request
	Issue127(ctx context.Context) (*http.Response, error)

	// Issue185 request  with any body
	Issue185WithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	Issue185(ctx context.Context, body Issue185JSONRequestBody) (*http.Response, error)

	// Issue209 request
	Issue209(ctx context.Context, str StringInPath) (*http.Response, error)

	// Issue30 request
	Issue30(ctx context.Context, pFallthrough string) (*http.Response, error)

	// Issue41 request
	Issue41(ctx context.Context, n1param N5StartsWithNumber) (*http.Response, error)

	// Issue9 request  with any body
	Issue9WithBody(ctx context.Context, params *Issue9Params, contentType string, body io.Reader) (*http.Response, error)

	Issue9(ctx context.Context, params *Issue9Params, body Issue9JSONRequestBody) (*http.Response, error)
}

func (c *Client) EnsureEverythingIsReferenced(ctx context.Context) (*http.Response, error) {
	req, err := NewEnsureEverythingIsReferencedRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) Issue127(ctx context.Context) (*http.Response, error) {
	req, err := NewIssue127Request(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) Issue185WithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewIssue185RequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) Issue185(ctx context.Context, body Issue185JSONRequestBody) (*http.Response, error) {
	req, err := NewIssue185Request(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) Issue209(ctx context.Context, str StringInPath) (*http.Response, error) {
	req, err := NewIssue209Request(c.Server, str)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) Issue30(ctx context.Context, pFallthrough string) (*http.Response, error) {
	req, err := NewIssue30Request(c.Server, pFallthrough)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) Issue41(ctx context.Context, n1param N5StartsWithNumber) (*http.Response, error) {
	req, err := NewIssue41Request(c.Server, n1param)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) Issue9WithBody(ctx context.Context, params *Issue9Params, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewIssue9RequestWithBody(c.Server, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) Issue9(ctx context.Context, params *Issue9Params, body Issue9JSONRequestBody) (*http.Response, error) {
	req, err := NewIssue9Request(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

// NewEnsureEverythingIsReferencedRequest generates requests for EnsureEverythingIsReferenced
func NewEnsureEverythingIsReferencedRequest(server string) (*http.Request, error) {
	var err error

	var queryUrl *url.URL
	queryUrl, err = url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/ensure-everything-is-referenced")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	req, err = http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewIssue127Request generates requests for Issue127
func NewIssue127Request(server string) (*http.Request, error) {
	var err error

	var queryUrl *url.URL
	queryUrl, err = url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/issues/127")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	req, err = http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewIssue185Request calls the generic Issue185 builder with application/json body
func NewIssue185Request(server string, body Issue185JSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewIssue185RequestWithBody(server, "application/json", bodyReader)
}

// NewIssue185RequestWithBody generates requests for Issue185 with any type of body
func NewIssue185RequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var queryUrl *url.URL
	queryUrl, err = url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/issues/185")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	req, err = http.NewRequest("GET", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewIssue209Request generates requests for Issue209
func NewIssue209Request(server string, str StringInPath) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "str", str)
	if err != nil {
		return nil, err
	}

	var queryUrl *url.URL
	queryUrl, err = url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/issues/209/$%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	req, err = http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewIssue30Request generates requests for Issue30
func NewIssue30Request(server string, pFallthrough string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "fallthrough", pFallthrough)
	if err != nil {
		return nil, err
	}

	var queryUrl *url.URL
	queryUrl, err = url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/issues/30/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	req, err = http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewIssue41Request generates requests for Issue41
func NewIssue41Request(server string, n1param N5StartsWithNumber) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "1param", n1param)
	if err != nil {
		return nil, err
	}

	var queryUrl *url.URL
	queryUrl, err = url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/issues/41/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	req, err = http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewIssue9Request calls the generic Issue9 builder with application/json body
func NewIssue9Request(server string, params *Issue9Params, body Issue9JSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewIssue9RequestWithBody(server, params, "application/json", bodyReader)
}

// NewIssue9RequestWithBody generates requests for Issue9 with any type of body
func NewIssue9RequestWithBody(server string, params *Issue9Params, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var queryUrl *url.URL
	queryUrl, err = url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/issues/9")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	var queryFrag string
	var parsed url.Values
	var queryParamBuf []byte
	_ = queryFrag
	_ = parsed
	_ = queryParamBuf

	if queryFrag, err = runtime.StyleParam("form", true, "foo", params.Foo); err != nil {
		return nil, err
	} else if parsed, err = url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryUrl.RawQuery = queryValues.Encode()

	var req *http.Request
	req, err = http.NewRequest("GET", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// EnsureEverythingIsReferenced request
	EnsureEverythingIsReferencedWithResponse(ctx context.Context) (*EnsureEverythingIsReferencedResponse, error)

	// Issue127 request
	Issue127WithResponse(ctx context.Context) (*Issue127Response, error)

	// Issue185 request  with any body
	Issue185WithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*Issue185Response, error)

	Issue185WithResponse(ctx context.Context, body Issue185JSONRequestBody) (*Issue185Response, error)

	// Issue209 request
	Issue209WithResponse(ctx context.Context, str StringInPath) (*Issue209Response, error)

	// Issue30 request
	Issue30WithResponse(ctx context.Context, pFallthrough string) (*Issue30Response, error)

	// Issue41 request
	Issue41WithResponse(ctx context.Context, n1param N5StartsWithNumber) (*Issue41Response, error)

	// Issue9 request  with any body
	Issue9WithBodyWithResponse(ctx context.Context, params *Issue9Params, contentType string, body io.Reader) (*Issue9Response, error)

	Issue9WithResponse(ctx context.Context, params *Issue9Params, body Issue9JSONRequestBody) (*Issue9Response, error)
}

type EnsureEverythingIsReferencedResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		AnyType1 *AnyType1 `json:"anyType1,omitempty"`

		// This should be an interface{}
		AnyType2         *AnyType2         `json:"anyType2,omitempty"`
		CustomStringType *CustomStringType `json:"customStringType,omitempty" validate:"custom"`
	}
}

// Status returns HTTPResponse.Status
func (r EnsureEverythingIsReferencedResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r EnsureEverythingIsReferencedResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type Issue127Response struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GenericObject
	XML200       *GenericObject
	YAML200      *GenericObject
	JSONDefault  *GenericObject
}

// Status returns HTTPResponse.Status
func (r Issue127Response) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r Issue127Response) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type Issue185Response struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r Issue185Response) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r Issue185Response) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type Issue209Response struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r Issue209Response) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r Issue209Response) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type Issue30Response struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r Issue30Response) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r Issue30Response) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type Issue41Response struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r Issue41Response) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r Issue41Response) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type Issue9Response struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r Issue9Response) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r Issue9Response) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// EnsureEverythingIsReferencedWithResponse request returning *EnsureEverythingIsReferencedResponse
func (c *ClientWithResponses) EnsureEverythingIsReferencedWithResponse(ctx context.Context) (*EnsureEverythingIsReferencedResponse, error) {
	rsp, err := c.EnsureEverythingIsReferenced(ctx)
	if err != nil {
		return nil, err
	}
	return ParseEnsureEverythingIsReferencedResponse(rsp)
}

// Issue127WithResponse request returning *Issue127Response
func (c *ClientWithResponses) Issue127WithResponse(ctx context.Context) (*Issue127Response, error) {
	rsp, err := c.Issue127(ctx)
	if err != nil {
		return nil, err
	}
	return ParseIssue127Response(rsp)
}

// Issue185WithBodyWithResponse request with arbitrary body returning *Issue185Response
func (c *ClientWithResponses) Issue185WithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*Issue185Response, error) {
	rsp, err := c.Issue185WithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseIssue185Response(rsp)
}

func (c *ClientWithResponses) Issue185WithResponse(ctx context.Context, body Issue185JSONRequestBody) (*Issue185Response, error) {
	rsp, err := c.Issue185(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseIssue185Response(rsp)
}

// Issue209WithResponse request returning *Issue209Response
func (c *ClientWithResponses) Issue209WithResponse(ctx context.Context, str StringInPath) (*Issue209Response, error) {
	rsp, err := c.Issue209(ctx, str)
	if err != nil {
		return nil, err
	}
	return ParseIssue209Response(rsp)
}

// Issue30WithResponse request returning *Issue30Response
func (c *ClientWithResponses) Issue30WithResponse(ctx context.Context, pFallthrough string) (*Issue30Response, error) {
	rsp, err := c.Issue30(ctx, pFallthrough)
	if err != nil {
		return nil, err
	}
	return ParseIssue30Response(rsp)
}

// Issue41WithResponse request returning *Issue41Response
func (c *ClientWithResponses) Issue41WithResponse(ctx context.Context, n1param N5StartsWithNumber) (*Issue41Response, error) {
	rsp, err := c.Issue41(ctx, n1param)
	if err != nil {
		return nil, err
	}
	return ParseIssue41Response(rsp)
}

// Issue9WithBodyWithResponse request with arbitrary body returning *Issue9Response
func (c *ClientWithResponses) Issue9WithBodyWithResponse(ctx context.Context, params *Issue9Params, contentType string, body io.Reader) (*Issue9Response, error) {
	rsp, err := c.Issue9WithBody(ctx, params, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseIssue9Response(rsp)
}

func (c *ClientWithResponses) Issue9WithResponse(ctx context.Context, params *Issue9Params, body Issue9JSONRequestBody) (*Issue9Response, error) {
	rsp, err := c.Issue9(ctx, params, body)
	if err != nil {
		return nil, err
	}
	return ParseIssue9Response(rsp)
}

// ParseEnsureEverythingIsReferencedResponse parses an HTTP response from a EnsureEverythingIsReferencedWithResponse call
func ParseEnsureEverythingIsReferencedResponse(rsp *http.Response) (*EnsureEverythingIsReferencedResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &EnsureEverythingIsReferencedResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			AnyType1 *AnyType1 `json:"anyType1,omitempty"`

			// This should be an interface{}
			AnyType2         *AnyType2         `json:"anyType2,omitempty"`
			CustomStringType *CustomStringType `json:"customStringType,omitempty" validate:"custom"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseIssue127Response parses an HTTP response from a Issue127WithResponse call
func ParseIssue127Response(rsp *http.Response) (*Issue127Response, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &Issue127Response{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GenericObject
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.YAML200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest GenericObject
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "xml") && rsp.StatusCode == 200:
		var dest GenericObject
		if err := xml.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.YAML200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "yaml") && rsp.StatusCode == 200:
		var dest GenericObject
		if err := yaml.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.YAML200 = &dest

	case rsp.StatusCode == 200:
	// Content-type (text/markdown) unsupported

	case true:
		// Content-type (text/markdown) unsupported

	}

	return response, nil
}

// ParseIssue185Response parses an HTTP response from a Issue185WithResponse call
func ParseIssue185Response(rsp *http.Response) (*Issue185Response, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &Issue185Response{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ParseIssue209Response parses an HTTP response from a Issue209WithResponse call
func ParseIssue209Response(rsp *http.Response) (*Issue209Response, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &Issue209Response{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ParseIssue30Response parses an HTTP response from a Issue30WithResponse call
func ParseIssue30Response(rsp *http.Response) (*Issue30Response, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &Issue30Response{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ParseIssue41Response parses an HTTP response from a Issue41WithResponse call
func ParseIssue41Response(rsp *http.Response) (*Issue41Response, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &Issue41Response{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ParseIssue9Response parses an HTTP response from a Issue9WithResponse call
func ParseIssue9Response(rsp *http.Response) (*Issue9Response, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &Issue9Response{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /ensure-everything-is-referenced)
	EnsureEverythingIsReferenced(ctx EnsureEverythingIsReferencedContext) error

	// (GET /issues/127)
	Issue127(ctx Issue127Context) error

	// (GET /issues/185)
	Issue185(ctx Issue185Context) error

	// (GET /issues/209/${str})
	Issue209(ctx Issue209Context, str StringInPath) error

	// (GET /issues/30/{fallthrough})
	Issue30(ctx Issue30Context, pFallthrough string) error

	// (GET /issues/41/{1param})
	Issue41(ctx Issue41Context, n1param N5StartsWithNumber) error

	// (GET /issues/9)
	Issue9(ctx Issue9Context, params Issue9Params) error
}

type EnsureEverythingIsReferencedContext struct {
	echo.Context
}

func (c *EnsureEverythingIsReferencedContext) JSON200(resp struct {
	AnyType1 *AnyType1 `json:"anyType1,omitempty"`

	// This should be an interface{}
	AnyType2         *AnyType2         `json:"anyType2,omitempty"`
	CustomStringType *CustomStringType `json:"customStringType,omitempty" validate:"custom"`
}) error {
	err := c.Validate(resp)
	if err != nil {
		return fmt.Errorf("response validation failed: %s", err)
	}
	return c.JSON(200, resp)
}

type Issue127Context struct {
	echo.Context
}

func (c *Issue127Context) JSON200(resp GenericObject) error {
	err := c.Validate(resp)
	if err != nil {
		return fmt.Errorf("response validation failed: %s", err)
	}
	return c.JSON(200, resp)
}

func (c *Issue127Context) XML200(resp GenericObject) error {
	err := c.Validate(resp)
	if err != nil {
		return fmt.Errorf("response validation failed: %s", err)
	}
	return c.XML(200, resp)
}

func (c *Issue127Context) YAML200(resp GenericObject) error {
	err := c.Validate(resp)
	if err != nil {
		return fmt.Errorf("response validation failed: %s", err)
	}
	var out []byte
	out, err = yaml.Marshal(resp)
	if err != nil {
		return err
	}
	return c.Blob(200, "text/yaml", out)
}

type Issue185Context struct {
	echo.Context
}

func (c *Issue185Context) BindJSON() (*Issue185JSONBody, error) {
	var err error

	// optional
	if c.Request().ContentLength == 0 {
		return nil, nil
	}

	ctype := c.Request().Header.Get(echo.HeaderContentType)
	if ctype != "application/json" {
		err = errors.New(fmt.Sprintf("incorrect content type: %s", ctype))
		return nil, err
	}

	var result Issue185JSONBody
	if err = c.Bind(&result); err != nil {
		return nil, err
	}

	if err = c.Validate(result); err != nil {
		return nil, &echo.HTTPError{
			Code:     http.StatusBadRequest,
			Message:  fmt.Sprintf("request validation failed: %s", err.Error()),
			Internal: err,
		}
	}

	return &result, nil
}

type Issue209Context struct {
	echo.Context
}

type Issue30Context struct {
	echo.Context
}

type Issue41Context struct {
	echo.Context
}

type Issue9Context struct {
	echo.Context
}

func (c *Issue9Context) BindJSON() (*Issue9JSONBody, error) {
	var err error

	// optional
	if c.Request().ContentLength == 0 {
		return nil, nil
	}

	ctype := c.Request().Header.Get(echo.HeaderContentType)
	if ctype != "application/json" {
		err = errors.New(fmt.Sprintf("incorrect content type: %s", ctype))
		return nil, err
	}

	var result Issue9JSONBody
	if err = c.Bind(&result); err != nil {
		return nil, err
	}

	if err = c.Validate(result); err != nil {
		return nil, &echo.HTTPError{
			Code:     http.StatusBadRequest,
			Message:  fmt.Sprintf("request validation failed: %s", err.Error()),
			Internal: err,
		}
	}

	return &result, nil
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler func(echo.Context) ServerInterface
}

// EnsureEverythingIsReferenced converts echo context to params.
func (w *ServerInterfaceWrapper) EnsureEverythingIsReferenced(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler(ctx).EnsureEverythingIsReferenced(EnsureEverythingIsReferencedContext{ctx})
	return err
}

// Issue127 converts echo context to params.
func (w *ServerInterfaceWrapper) Issue127(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler(ctx).Issue127(Issue127Context{ctx})
	return err
}

// Issue185 converts echo context to params.
func (w *ServerInterfaceWrapper) Issue185(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler(ctx).Issue185(Issue185Context{ctx})
	return err
}

// Issue209 converts echo context to params.
func (w *ServerInterfaceWrapper) Issue209(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "str" -------------
	var str StringInPath

	err = runtime.BindStyledParameter("simple", false, "str", ctx.Param("str"), &str)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter str: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler(ctx).Issue209(Issue209Context{ctx}, str)
	return err
}

// Issue30 converts echo context to params.
func (w *ServerInterfaceWrapper) Issue30(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "fallthrough" -------------
	var pFallthrough string

	err = runtime.BindStyledParameter("simple", false, "fallthrough", ctx.Param("fallthrough"), &pFallthrough)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter fallthrough: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler(ctx).Issue30(Issue30Context{ctx}, pFallthrough)
	return err
}

// Issue41 converts echo context to params.
func (w *ServerInterfaceWrapper) Issue41(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "1param" -------------
	var n1param N5StartsWithNumber

	err = runtime.BindStyledParameter("simple", false, "1param", ctx.Param("1param"), &n1param)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter 1param: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler(ctx).Issue41(Issue41Context{ctx}, n1param)
	return err
}

// Issue9 converts echo context to params.
func (w *ServerInterfaceWrapper) Issue9(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params Issue9Params
	// ------------- Required query parameter "foo" -------------

	err = runtime.BindQueryParameter("form", true, true, "foo", ctx.QueryParams(), &params.Foo)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter foo: %s", err))
	}

	// Validate params
	err = ctx.Validate(params)
	if err != nil {
		return &echo.HTTPError{
			Code:     http.StatusBadRequest,
			Message:  fmt.Sprintf("request validation failed: %s", err.Error()),
			Internal: err,
		}
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler(ctx).Issue9(Issue9Context{ctx}, params)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface, pathPrefix string) {

	wrapper := ServerInterfaceWrapper{
		Handler: func(echo.Context) ServerInterface {
			return si
		},
	}
	wrapper.RegisterHandlers(router, pathPrefix)

}

func (wrapper ServerInterfaceWrapper) RegisterHandlers(router EchoRouter, pathPrefix string) {
	router.GET(path.Join(pathPrefix, "/ensure-everything-is-referenced"), wrapper.EnsureEverythingIsReferenced)
	router.GET(path.Join(pathPrefix, "/issues/127"), wrapper.Issue127)
	router.GET(path.Join(pathPrefix, "/issues/185"), wrapper.Issue185)
	router.GET(path.Join(pathPrefix, "/issues/209/$:str"), wrapper.Issue209)
	router.GET(path.Join(pathPrefix, "/issues/30/:fallthrough"), wrapper.Issue30)
	router.GET(path.Join(pathPrefix, "/issues/41/:1param"), wrapper.Issue41)
	router.GET(path.Join(pathPrefix, "/issues/9"), wrapper.Issue9)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7RX32/bNhD+Vw5cgb24lu00WKO3rCiGPKwNmgB7aPJAi2eLjUSy5CmJYOh/H46yLXmW",
	"vGbpniKZuh/fd98dLxuR2dJZg4aCSDfCSS9LJPTx7Ya8Nusrcy0p53eFIfPakbZGpOISQjwHJymHvaWY",
	"CM3H/KuYCCNLFKkIxAcev1faoxIp+QonImQ5lpJdU+22n2mzFk3T7A5jIuc3JD2FvzTln6pyif44m9tc",
	"B2hNgGNCiCbwpCkHCaY1m+wC2eU3zEg0E3Fp6tva4Vykm+5tMRYgt1WhYIkgDWhD6Fcyw03Djj5UgWzZ",
	"cnYbo2zEyvpSkkhFFg+7+DugE/EHGvQ6+9wm1FHRZfipKgq5LPDaW4eeNLalOnizMU1ZDHA52R9eGrXz",
	"xd+Z/XNbjSO7rlyb8cOXOT3w+rV7HvZ3f1QvdqDNyg7UBwNBJgMGWFkPj9JrWwXQIVTxp8oosI/ogXSJ",
	"U7guUAYEqRRIoJ0tm94ZaWpYVmtY6WdU0zvDZdPEmNooN+gfo5ge0Yc2+nw6m85artFIp0Uqzqaz6VxM",
	"YiPEGiVoQuXxLT6irynXZv1Wh7ceV+jRZC3Na6QR6aFRzmpDgM86UIBggXJJ0DUwZNKwNDOPklCBNkC5",
	"DncmOMxAGgXGEn/gfGVQRVysIclhrpRIxceY4Md9flfhS5cdlyg4a0KruMVsxn8yawhNTFo6V+gseku+",
	"Bc580+vwQ73KruvEG48rkYpfkg5Ksm3+ZN+dzWRns/hBmwXbZANNecr2qIlZcEcabKIOk1ZbyXzx22jp",
	"/pQPCEwqVCZUzlnPlYmkPROw4wDKml8JnEcsHUH3VTydDpTpiuNy1FeW5BQRh2OJ4fZ9PZfFa1wx+KSU",
	"/kHZJ/NqR7V8TTbsRuFKVgX9j+T9JMT/VN778/GhUTuENdtHBPCUo4HdTZDspi10bQnSI+zG97js3p9v",
	"hzUG+t2q+qeRNnDNtWh7Guf0+gQsZhfJm00g34zy8CHH7CGAXnUrSgtVYVbIjoKiHga8mF2I4xwmB6vS",
	"12Fk3SfJwSrV3PcgnM2SzUoWBeXeVuu8OUbwBQNfOAoesH6yXvXXEOcx3lI87PnKYwLj/rMdHFtKBnCd",
	"zX4E1sAq10v2RStdH/S7ebKZx1DjhbveZdLb53jdjBvdfp8bQPauvXX/DUcb/ySEU3I93kmb5v6kWC/G",
	"NVpoNNQKNMS5D9pk1nvMqKj5uagUqrjYbFuvpWFpVc03+53p8I627sUILd8r9HWvvta+rK7/eRxsZ2+f",
	"ic/bARWRiaHm5/8M4v7VIqh8IVKRE7k0SbbLF69zU4XoSummUnO//R0AAP//Wanza+kMAAA=",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
