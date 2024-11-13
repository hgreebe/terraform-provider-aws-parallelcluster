/*
ParallelCluster

ParallelCluster API

API version: 3.11.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)


// ImageLogsAPIService ImageLogsAPI service
type ImageLogsAPIService service

type ImageLogsAPIGetImageLogEventsRequest struct {
	ctx context.Context
	ApiService *ImageLogsAPIService
	imageId string
	logStreamName string
	region *string
	nextToken *string
	startFromHead *bool
	limit *int32
	startTime *time.Time
	endTime *time.Time
}

// AWS Region that the operation corresponds to.
func (r ImageLogsAPIGetImageLogEventsRequest) Region(region string) ImageLogsAPIGetImageLogEventsRequest {
	r.region = &region
	return r
}

// Token to use for paginated requests.
func (r ImageLogsAPIGetImageLogEventsRequest) NextToken(nextToken string) ImageLogsAPIGetImageLogEventsRequest {
	r.nextToken = &nextToken
	return r
}

// If the value is true, the earliest log events are returned first. If the value is false, the latest log events are returned first. (Defaults to &#39;false&#39;.)
func (r ImageLogsAPIGetImageLogEventsRequest) StartFromHead(startFromHead bool) ImageLogsAPIGetImageLogEventsRequest {
	r.startFromHead = &startFromHead
	return r
}

// The maximum number of log events returned. If you don&#39;t specify a value, the maximum is as many log events as can fit in a response size of 1 MB, up to 10,000 log events.
func (r ImageLogsAPIGetImageLogEventsRequest) Limit(limit int32) ImageLogsAPIGetImageLogEventsRequest {
	r.limit = &limit
	return r
}

// The start of the time range, expressed in ISO 8601 format (e.g. &#39;2021-01-01T20:00:00Z&#39;). Events with a timestamp equal to this time or later than this time are included.
func (r ImageLogsAPIGetImageLogEventsRequest) StartTime(startTime time.Time) ImageLogsAPIGetImageLogEventsRequest {
	r.startTime = &startTime
	return r
}

// The end of the time range, expressed in ISO 8601 format (e.g. &#39;2021-01-01T20:00:00Z&#39;). Events with a timestamp equal to or later than this time are not included.
func (r ImageLogsAPIGetImageLogEventsRequest) EndTime(endTime time.Time) ImageLogsAPIGetImageLogEventsRequest {
	r.endTime = &endTime
	return r
}

func (r ImageLogsAPIGetImageLogEventsRequest) Execute() (*GetImageLogEventsResponseContent, *http.Response, error) {
	return r.ApiService.GetImageLogEventsExecute(r)
}

/*
GetImageLogEvents Method for GetImageLogEvents

Retrieve the events associated with an image build.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param imageId Id of the image.
 @param logStreamName Name of the log stream.
 @return ImageLogsAPIGetImageLogEventsRequest
*/
func (a *ImageLogsAPIService) GetImageLogEvents(ctx context.Context, imageId string, logStreamName string) ImageLogsAPIGetImageLogEventsRequest {
	return ImageLogsAPIGetImageLogEventsRequest{
		ApiService: a,
		ctx: ctx,
		imageId: imageId,
		logStreamName: logStreamName,
	}
}

// Execute executes the request
//  @return GetImageLogEventsResponseContent
func (a *ImageLogsAPIService) GetImageLogEventsExecute(r ImageLogsAPIGetImageLogEventsRequest) (*GetImageLogEventsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetImageLogEventsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImageLogsAPIService.GetImageLogEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/images/custom/{imageId}/logstreams/{logStreamName}"
	localVarPath = strings.Replace(localVarPath, "{"+"imageId"+"}", url.PathEscape(parameterValueToString(r.imageId, "imageId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"logStreamName"+"}", url.PathEscape(parameterValueToString(r.logStreamName, "logStreamName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.region != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "region", r.region, "form", "")
	}
	if r.nextToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nextToken", r.nextToken, "form", "")
	}
	if r.startFromHead != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startFromHead", r.startFromHead, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.startTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["aws.auth.sigv4"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v BadRequestExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v UnauthorizedClientErrorResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v NotFoundExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v LimitExceededExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InternalServiceExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
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

type ImageLogsAPIGetImageStackEventsRequest struct {
	ctx context.Context
	ApiService *ImageLogsAPIService
	imageId string
	region *string
	nextToken *string
}

// AWS Region that the operation corresponds to.
func (r ImageLogsAPIGetImageStackEventsRequest) Region(region string) ImageLogsAPIGetImageStackEventsRequest {
	r.region = &region
	return r
}

// Token to use for paginated requests.
func (r ImageLogsAPIGetImageStackEventsRequest) NextToken(nextToken string) ImageLogsAPIGetImageStackEventsRequest {
	r.nextToken = &nextToken
	return r
}

func (r ImageLogsAPIGetImageStackEventsRequest) Execute() (*GetImageStackEventsResponseContent, *http.Response, error) {
	return r.ApiService.GetImageStackEventsExecute(r)
}

/*
GetImageStackEvents Method for GetImageStackEvents

Retrieve the events associated with the stack for a given image build.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param imageId Id of the image.
 @return ImageLogsAPIGetImageStackEventsRequest
*/
func (a *ImageLogsAPIService) GetImageStackEvents(ctx context.Context, imageId string) ImageLogsAPIGetImageStackEventsRequest {
	return ImageLogsAPIGetImageStackEventsRequest{
		ApiService: a,
		ctx: ctx,
		imageId: imageId,
	}
}

// Execute executes the request
//  @return GetImageStackEventsResponseContent
func (a *ImageLogsAPIService) GetImageStackEventsExecute(r ImageLogsAPIGetImageStackEventsRequest) (*GetImageStackEventsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetImageStackEventsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImageLogsAPIService.GetImageStackEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/images/custom/{imageId}/stackevents"
	localVarPath = strings.Replace(localVarPath, "{"+"imageId"+"}", url.PathEscape(parameterValueToString(r.imageId, "imageId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.region != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "region", r.region, "form", "")
	}
	if r.nextToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nextToken", r.nextToken, "form", "")
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["aws.auth.sigv4"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v BadRequestExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v UnauthorizedClientErrorResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v NotFoundExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v LimitExceededExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InternalServiceExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
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

type ImageLogsAPIListImageLogStreamsRequest struct {
	ctx context.Context
	ApiService *ImageLogsAPIService
	imageId string
	region *string
	nextToken *string
}

// Region that the given image belongs to.
func (r ImageLogsAPIListImageLogStreamsRequest) Region(region string) ImageLogsAPIListImageLogStreamsRequest {
	r.region = &region
	return r
}

// Token to use for paginated requests.
func (r ImageLogsAPIListImageLogStreamsRequest) NextToken(nextToken string) ImageLogsAPIListImageLogStreamsRequest {
	r.nextToken = &nextToken
	return r
}

func (r ImageLogsAPIListImageLogStreamsRequest) Execute() (*ListImageLogStreamsResponseContent, *http.Response, error) {
	return r.ApiService.ListImageLogStreamsExecute(r)
}

/*
ListImageLogStreams Method for ListImageLogStreams

Retrieve the list of log streams associated with an image.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param imageId Id of the image.
 @return ImageLogsAPIListImageLogStreamsRequest
*/
func (a *ImageLogsAPIService) ListImageLogStreams(ctx context.Context, imageId string) ImageLogsAPIListImageLogStreamsRequest {
	return ImageLogsAPIListImageLogStreamsRequest{
		ApiService: a,
		ctx: ctx,
		imageId: imageId,
	}
}

// Execute executes the request
//  @return ListImageLogStreamsResponseContent
func (a *ImageLogsAPIService) ListImageLogStreamsExecute(r ImageLogsAPIListImageLogStreamsRequest) (*ListImageLogStreamsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListImageLogStreamsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImageLogsAPIService.ListImageLogStreams")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/images/custom/{imageId}/logstreams"
	localVarPath = strings.Replace(localVarPath, "{"+"imageId"+"}", url.PathEscape(parameterValueToString(r.imageId, "imageId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.region != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "region", r.region, "form", "")
	}
	if r.nextToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nextToken", r.nextToken, "form", "")
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["aws.auth.sigv4"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v BadRequestExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v UnauthorizedClientErrorResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v NotFoundExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v LimitExceededExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InternalServiceExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
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
