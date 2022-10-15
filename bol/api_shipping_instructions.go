/*
DCSA OpenAPI specification for Electronic Bill of Lading

API specification issued by DCSA.org.  For explanation to specific values or objects please refer to the <a href='https://dcsa.org/wp-content/uploads/2020/12/20201208-DCSA-P1-DCSA-Information-Model-v3.0-FINAL.pdf'>Information Model v3.0</a>  It is possible to use this API as a standalone API. In order to do so it is necessary to use the poll-endPoint - /v1/events  in order to poll event information.  It is recomended to implement the <a href='https://app.swaggerhub.com/apis/dcsaorg/DOCUMENTATION_EVENT_HUB'>DCSA Documentation Event Hub</a> in order to use the push model. Here events are pushed as they occur.

API version: 1.0.0
Contact: info@dcsa.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bol

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/hubbl-ai/dcsa/common"
)

// ShippingInstructionsAPIService ShippingInstructionsAPI service
type ShippingInstructionsAPIService service

type APIV1ShippingInstructionsPostRequest struct {
	ctx                     context.Context
	APIService              *ShippingInstructionsAPIService
	shippingInstructionNoID *ShippingInstructionNoID
	aPIVersion              *string
}

// Parameters used to create the shipping instruction
func (r APIV1ShippingInstructionsPostRequest) ShippingInstructionNoID(shippingInstructionNoID ShippingInstructionNoID) APIV1ShippingInstructionsPostRequest {
	r.shippingInstructionNoID = &shippingInstructionNoID
	return r
}

// An API-Version header MAY be added to the request (optional); if added it MUST only contain MAJOR version. API-Version header MUST be aligned with the URI version.
func (r APIV1ShippingInstructionsPostRequest) APIVersion(aPIVersion string) APIV1ShippingInstructionsPostRequest {
	r.aPIVersion = &aPIVersion
	return r
}

func (r APIV1ShippingInstructionsPostRequest) Execute() (*ShippingInstruction, *http.Response, error) {
	return r.APIService.V1ShippingInstructionsPostExecute(r)
}

/*
V1ShippingInstructionsPost Post shipping instruction

# Creates a new shipping instruction

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return APIV1ShippingInstructionsPostRequest
*/
func (a *ShippingInstructionsAPIService) V1ShippingInstructionsPost(ctx context.Context) APIV1ShippingInstructionsPostRequest {
	return APIV1ShippingInstructionsPostRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ShippingInstruction
func (a *ShippingInstructionsAPIService) V1ShippingInstructionsPostExecute(r APIV1ShippingInstructionsPostRequest) (*ShippingInstruction, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ShippingInstruction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingInstructionsAPIService.V1ShippingInstructionsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/shipping-instructions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.shippingInstructionNoID == nil {
		return localVarReturnValue, nil, reportError("shippingInstructionNoID is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.aPIVersion != nil {
		localVarHeaderParams["API-Version"] = parameterToString(*r.aPIVersion, "")
	}
	// body params
	localVarPostBody = r.shippingInstructionNoID
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ModelError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type APIV1ShippingInstructionsShippingInstructionIDGetRequest struct {
	ctx                   context.Context
	APIService            *ShippingInstructionsAPIService
	shippingInstructionID string
	aPIVersion            *string
}

// An API-Version header MAY be added to the request (optional); if added it MUST only contain MAJOR version. API-Version header MUST be aligned with the URI version.
func (r APIV1ShippingInstructionsShippingInstructionIDGetRequest) APIVersion(aPIVersion string) APIV1ShippingInstructionsShippingInstructionIDGetRequest {
	r.aPIVersion = &aPIVersion
	return r
}

func (r APIV1ShippingInstructionsShippingInstructionIDGetRequest) Execute() (*ShippingInstruction, *http.Response, error) {
	return r.APIService.V1ShippingInstructionsShippingInstructionIDGetExecute(r)
}

/*
V1ShippingInstructionsShippingInstructionIDGet Get shipping instruction

Retrieves the shipping instruction with the ID in the path.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param shippingInstructionID
	@return APIV1ShippingInstructionsShippingInstructionIDGetRequest
*/
func (a *ShippingInstructionsAPIService) V1ShippingInstructionsShippingInstructionIDGet(ctx context.Context, shippingInstructionID string) APIV1ShippingInstructionsShippingInstructionIDGetRequest {
	return APIV1ShippingInstructionsShippingInstructionIDGetRequest{
		APIService:            a,
		ctx:                   ctx,
		shippingInstructionID: shippingInstructionID,
	}
}

// Execute executes the request
//
//	@return ShippingInstruction
func (a *ShippingInstructionsAPIService) V1ShippingInstructionsShippingInstructionIDGetExecute(r APIV1ShippingInstructionsShippingInstructionIDGetRequest) (*ShippingInstruction, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ShippingInstruction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingInstructionsAPIService.V1ShippingInstructionsShippingInstructionIDGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/shipping-instructions/{shippingInstructionID}"
	localVarPath = strings.Replace(localVarPath, "{"+"shippingInstructionID"+"}", url.PathEscape(parameterToString(r.shippingInstructionID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.shippingInstructionID) > 100 {
		return localVarReturnValue, nil, reportError("shippingInstructionID must have less than 100 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.aPIVersion != nil {
		localVarHeaderParams["API-Version"] = parameterToString(*r.aPIVersion, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ModelError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type APIV1ShippingInstructionsShippingInstructionIDPutRequest struct {
	ctx                   context.Context
	APIService            *ShippingInstructionsAPIService
	shippingInstructionID string
	shippingInstruction   *ShippingInstruction
	aPIVersion            *string
}

// Parameters used to update the shipping instruction
func (r APIV1ShippingInstructionsShippingInstructionIDPutRequest) ShippingInstruction(shippingInstruction ShippingInstruction) APIV1ShippingInstructionsShippingInstructionIDPutRequest {
	r.shippingInstruction = &shippingInstruction
	return r
}

// An API-Version header MAY be added to the request (optional); if added it MUST only contain MAJOR version. API-Version header MUST be aligned with the URI version.
func (r APIV1ShippingInstructionsShippingInstructionIDPutRequest) APIVersion(aPIVersion string) APIV1ShippingInstructionsShippingInstructionIDPutRequest {
	r.aPIVersion = &aPIVersion
	return r
}

func (r APIV1ShippingInstructionsShippingInstructionIDPutRequest) Execute() (*ShippingInstruction, *http.Response, error) {
	return r.APIService.V1ShippingInstructionsShippingInstructionIDPutExecute(r)
}

/*
V1ShippingInstructionsShippingInstructionIDPut Put shipping instruction

# Updates a shipping instruction

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param shippingInstructionID
	@return APIV1ShippingInstructionsShippingInstructionIDPutRequest
*/
func (a *ShippingInstructionsAPIService) V1ShippingInstructionsShippingInstructionIDPut(ctx context.Context, shippingInstructionID string) APIV1ShippingInstructionsShippingInstructionIDPutRequest {
	return APIV1ShippingInstructionsShippingInstructionIDPutRequest{
		APIService:            a,
		ctx:                   ctx,
		shippingInstructionID: shippingInstructionID,
	}
}

// Execute executes the request
//
//	@return ShippingInstruction
func (a *ShippingInstructionsAPIService) V1ShippingInstructionsShippingInstructionIDPutExecute(r APIV1ShippingInstructionsShippingInstructionIDPutRequest) (*ShippingInstruction, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ShippingInstruction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingInstructionsAPIService.V1ShippingInstructionsShippingInstructionIDPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/shipping-instructions/{shippingInstructionID}"
	localVarPath = strings.Replace(localVarPath, "{"+"shippingInstructionID"+"}", url.PathEscape(parameterToString(r.shippingInstructionID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.shippingInstructionID) > 100 {
		return localVarReturnValue, nil, reportError("shippingInstructionID must have less than 100 elements")
	}
	if r.shippingInstruction == nil {
		return localVarReturnValue, nil, reportError("shippingInstruction is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.aPIVersion != nil {
		localVarHeaderParams["API-Version"] = parameterToString(*r.aPIVersion, "")
	}
	// body params
	localVarPostBody = r.shippingInstruction
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ModelError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type APIV1ShippingInstructionsSummariesGetRequest struct {
	ctx                     context.Context
	APIService              *ShippingInstructionsAPIService
	carrierBookingReference *string
	documentStatus          *common.DocumentStatus
	equipmentReference      *string
	limit                   *int32
	cursor                  *string
	aPIVersion              *string
}

// A set of unique characters provided by carrier to identify a booking.  Specifying this filter will only return events related to this particular carrierBookingReference.
func (r APIV1ShippingInstructionsSummariesGetRequest) CarrierBookingReference(carrierBookingReference string) APIV1ShippingInstructionsSummariesGetRequest {
	r.carrierBookingReference = &carrierBookingReference
	return r
}

// Filter by the status of the document in the process. Possible values are: - RECE (Received) - DRFT (Drafted) - PENA (Pending Approval) - PENU (Pending Update) - REJE (Rejected) - APPR (Approved) - ISSU (Issued) - SURR (Surrendered) - SUBM (Submitted) - VOID (Void)
func (r APIV1ShippingInstructionsSummariesGetRequest) DocumentStatus(documentStatus common.DocumentStatus) APIV1ShippingInstructionsSummariesGetRequest {
	r.documentStatus = &documentStatus
	return r
}

// Filter by the unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. According to ISO 6346, a container identification code consists of a 4-letter prefix and a 7-digit number (composed of a 3-letter owner code, a category identifier, a serial number, and a check-digit). If a container does not comply with ISO 6346, it is suggested to follow Recommendation #2 “Container with non-ISO identification” from SMDG.
func (r APIV1ShippingInstructionsSummariesGetRequest) EquipmentReference(equipmentReference string) APIV1ShippingInstructionsSummariesGetRequest {
	r.equipmentReference = &equipmentReference
	return r
}

// Maximum number of items to return.
func (r APIV1ShippingInstructionsSummariesGetRequest) Limit(limit int32) APIV1ShippingInstructionsSummariesGetRequest {
	r.limit = &limit
	return r
}

// A server generated value to specify a specific point in a collection result, used for pagination.
func (r APIV1ShippingInstructionsSummariesGetRequest) Cursor(cursor string) APIV1ShippingInstructionsSummariesGetRequest {
	r.cursor = &cursor
	return r
}

// An API-Version header MAY be added to the request (optional); if added it MUST only contain MAJOR version. API-Version header MUST be aligned with the URI version.
func (r APIV1ShippingInstructionsSummariesGetRequest) APIVersion(aPIVersion string) APIV1ShippingInstructionsSummariesGetRequest {
	r.aPIVersion = &aPIVersion
	return r
}

func (r APIV1ShippingInstructionsSummariesGetRequest) Execute() ([]ShippingInstructionHeader, *http.Response, error) {
	return r.APIService.V1ShippingInstructionsSummariesGetExecute(r)
}

/*
V1ShippingInstructionsSummariesGet Get shipping instruction

Retrieves the shipping instruction &apos;metadata&apos; with the ID in the path.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return APIV1ShippingInstructionsSummariesGetRequest
*/
func (a *ShippingInstructionsAPIService) V1ShippingInstructionsSummariesGet(ctx context.Context) APIV1ShippingInstructionsSummariesGetRequest {
	return APIV1ShippingInstructionsSummariesGetRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []ShippingInstructionHeader
func (a *ShippingInstructionsAPIService) V1ShippingInstructionsSummariesGetExecute(r APIV1ShippingInstructionsSummariesGetRequest) ([]ShippingInstructionHeader, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ShippingInstructionHeader
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingInstructionsAPIService.V1ShippingInstructionsSummariesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/shipping-instructions-summaries"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.carrierBookingReference != nil {
		localVarQueryParams.Add("carrierBookingReference", parameterToString(*r.carrierBookingReference, ""))
	}
	if r.documentStatus != nil {
		localVarQueryParams.Add("documentStatus", parameterToString(*r.documentStatus, ""))
	}
	if r.equipmentReference != nil {
		localVarQueryParams.Add("equipmentReference", parameterToString(*r.equipmentReference, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.aPIVersion != nil {
		localVarHeaderParams["API-Version"] = parameterToString(*r.aPIVersion, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ModelError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}