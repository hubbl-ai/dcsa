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

// TransportDocumentsAPIService TransportDocumentsAPI service
type TransportDocumentsAPIService service

type APIV1TransportDocumentsSummariesGetRequest struct {
	ctx                     context.Context
	APIService              *TransportDocumentsAPIService
	limit                   *int32
	cursor                  *string
	carrierBookingReference *string
	documentStatus          *common.DocumentStatus
	equipmentReference      *string
	aPIVersion              *string
}

// Maximum number of items to return.
func (r APIV1TransportDocumentsSummariesGetRequest) Limit(limit int32) APIV1TransportDocumentsSummariesGetRequest {
	r.limit = &limit
	return r
}

// A server generated value to specify a specific point in a collection result, used for pagination.
func (r APIV1TransportDocumentsSummariesGetRequest) Cursor(cursor string) APIV1TransportDocumentsSummariesGetRequest {
	r.cursor = &cursor
	return r
}

// A set of unique characters provided by carrier to identify a booking.  Specifying this filter will only return events related to this particular carrierBookingReference.
func (r APIV1TransportDocumentsSummariesGetRequest) CarrierBookingReference(carrierBookingReference string) APIV1TransportDocumentsSummariesGetRequest {
	r.carrierBookingReference = &carrierBookingReference
	return r
}

// Filter by the status of the document in the process. Possible values are: - RECE (Received) - DRFT (Drafted) - PENA (Pending Approval) - PENU (Pending Update) - REJE (Rejected) - APPR (Approved) - ISSU (Issued) - SURR (Surrendered) - SUBM (Submitted) - VOID (Void)
func (r APIV1TransportDocumentsSummariesGetRequest) DocumentStatus(documentStatus common.DocumentStatus) APIV1TransportDocumentsSummariesGetRequest {
	r.documentStatus = &documentStatus
	return r
}

// Filter by the unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. According to ISO 6346, a container identification code consists of a 4-letter prefix and a 7-digit number (composed of a 3-letter owner code, a category identifier, a serial number, and a check-digit). If a container does not comply with ISO 6346, it is suggested to follow Recommendation #2 “Container with non-ISO identification” from SMDG.
func (r APIV1TransportDocumentsSummariesGetRequest) EquipmentReference(equipmentReference string) APIV1TransportDocumentsSummariesGetRequest {
	r.equipmentReference = &equipmentReference
	return r
}

// An API-Version header MAY be added to the request (optional); if added it MUST only contain MAJOR version. API-Version header MUST be aligned with the URI version.
func (r APIV1TransportDocumentsSummariesGetRequest) APIVersion(aPIVersion string) APIV1TransportDocumentsSummariesGetRequest {
	r.aPIVersion = &aPIVersion
	return r
}

func (r APIV1TransportDocumentsSummariesGetRequest) Execute() ([]TransportDocumentHeader, *http.Response, error) {
	return r.APIService.V1TransportDocumentsSummariesGetExecute(r)
}

/*
V1TransportDocumentsSummariesGet Get Transport Documents

Retrieves all Transport Documents

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return APIV1TransportDocumentsSummariesGetRequest
*/
func (a *TransportDocumentsAPIService) V1TransportDocumentsSummariesGet(ctx context.Context) APIV1TransportDocumentsSummariesGetRequest {
	return APIV1TransportDocumentsSummariesGetRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []TransportDocumentHeader
func (a *TransportDocumentsAPIService) V1TransportDocumentsSummariesGetExecute(r APIV1TransportDocumentsSummariesGetRequest) ([]TransportDocumentHeader, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []TransportDocumentHeader
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportDocumentsAPIService.V1TransportDocumentsSummariesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/transport-documents-summaries"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.carrierBookingReference != nil {
		localVarQueryParams.Add("carrierBookingReference", parameterToString(*r.carrierBookingReference, ""))
	}
	if r.documentStatus != nil {
		localVarQueryParams.Add("documentStatus", parameterToString(*r.documentStatus, ""))
	}
	if r.equipmentReference != nil {
		localVarQueryParams.Add("equipmentReference", parameterToString(*r.equipmentReference, ""))
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

type APIV1TransportDocumentsTransportDocumentReferenceApprovalsPostRequest struct {
	ctx                        context.Context
	APIService                 *TransportDocumentsAPIService
	transportDocumentReference string
	documentSignature          *string
	aPIVersion                 *string
}

// The signature of the Transport Document to approve
func (r APIV1TransportDocumentsTransportDocumentReferenceApprovalsPostRequest) DocumentSignature(documentSignature string) APIV1TransportDocumentsTransportDocumentReferenceApprovalsPostRequest {
	r.documentSignature = &documentSignature
	return r
}

// An API-Version header MAY be added to the request (optional); if added it MUST only contain MAJOR version. API-Version header MUST be aligned with the URI version.
func (r APIV1TransportDocumentsTransportDocumentReferenceApprovalsPostRequest) APIVersion(aPIVersion string) APIV1TransportDocumentsTransportDocumentReferenceApprovalsPostRequest {
	r.aPIVersion = &aPIVersion
	return r
}

func (r APIV1TransportDocumentsTransportDocumentReferenceApprovalsPostRequest) Execute() (*TransportDocument, *http.Response, error) {
	return r.APIService.V1TransportDocumentsTransportDocumentReferenceApprovalsPostExecute(r)
}

/*
V1TransportDocumentsTransportDocumentReferenceApprovalsPost Approve a Transport Document.

Approves a Transport Document.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param transportDocumentReference The transportDocumentReference of the transportDocument
	@return APIV1TransportDocumentsTransportDocumentReferenceApprovalsPostRequest
*/
func (a *TransportDocumentsAPIService) V1TransportDocumentsTransportDocumentReferenceApprovalsPost(ctx context.Context, transportDocumentReference string) APIV1TransportDocumentsTransportDocumentReferenceApprovalsPostRequest {
	return APIV1TransportDocumentsTransportDocumentReferenceApprovalsPostRequest{
		APIService:                 a,
		ctx:                        ctx,
		transportDocumentReference: transportDocumentReference,
	}
}

// Execute executes the request
//
//	@return TransportDocument
func (a *TransportDocumentsAPIService) V1TransportDocumentsTransportDocumentReferenceApprovalsPostExecute(r APIV1TransportDocumentsTransportDocumentReferenceApprovalsPostRequest) (*TransportDocument, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TransportDocument
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportDocumentsAPIService.V1TransportDocumentsTransportDocumentReferenceApprovalsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/transport-documents/{transportDocumentReference}/approvals"
	localVarPath = strings.Replace(localVarPath, "{"+"transportDocumentReference"+"}", url.PathEscape(parameterToString(r.transportDocumentReference, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.transportDocumentReference) > 20 {
		return localVarReturnValue, nil, reportError("transportDocumentReference must have less than 20 elements")
	}
	if r.documentSignature == nil {
		return localVarReturnValue, nil, reportError("documentSignature is required and must be specified")
	}

	localVarQueryParams.Add("documentSignature", parameterToString(*r.documentSignature, ""))
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

type APIV1TransportDocumentsTransportDocumentReferenceGetRequest struct {
	ctx                        context.Context
	APIService                 *TransportDocumentsAPIService
	transportDocumentReference string
	aPIVersion                 *string
}

// An API-Version header MAY be added to the request (optional); if added it MUST only contain MAJOR version. API-Version header MUST be aligned with the URI version.
func (r APIV1TransportDocumentsTransportDocumentReferenceGetRequest) APIVersion(aPIVersion string) APIV1TransportDocumentsTransportDocumentReferenceGetRequest {
	r.aPIVersion = &aPIVersion
	return r
}

func (r APIV1TransportDocumentsTransportDocumentReferenceGetRequest) Execute() (*TransportDocument, *http.Response, error) {
	return r.APIService.V1TransportDocumentsTransportDocumentReferenceGetExecute(r)
}

/*
V1TransportDocumentsTransportDocumentReferenceGet Get Transport Document

Retrieves the transport document with the transportDocumentReference in the path.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param transportDocumentReference The transportDocumentReference of the transportDocument
	@return APIV1TransportDocumentsTransportDocumentReferenceGetRequest
*/
func (a *TransportDocumentsAPIService) V1TransportDocumentsTransportDocumentReferenceGet(ctx context.Context, transportDocumentReference string) APIV1TransportDocumentsTransportDocumentReferenceGetRequest {
	return APIV1TransportDocumentsTransportDocumentReferenceGetRequest{
		APIService:                 a,
		ctx:                        ctx,
		transportDocumentReference: transportDocumentReference,
	}
}

// Execute executes the request
//
//	@return TransportDocument
func (a *TransportDocumentsAPIService) V1TransportDocumentsTransportDocumentReferenceGetExecute(r APIV1TransportDocumentsTransportDocumentReferenceGetRequest) (*TransportDocument, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TransportDocument
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportDocumentsAPIService.V1TransportDocumentsTransportDocumentReferenceGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/transport-documents/{transportDocumentReference}"
	localVarPath = strings.Replace(localVarPath, "{"+"transportDocumentReference"+"}", url.PathEscape(parameterToString(r.transportDocumentReference, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.transportDocumentReference) > 20 {
		return localVarReturnValue, nil, reportError("transportDocumentReference must have less than 20 elements")
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
