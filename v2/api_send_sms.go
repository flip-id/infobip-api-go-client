/*
 * Infobip Client API Libraries OpenAPI Specification
 *
 * OpenAPI specification containing public endpoints supported in client API libraries.
 *
 * API version: 1.0.157
 * Contact: support@infobip.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package infobip

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"reflect"
)

// Linger please
var (
	_ _context.Context
)

// SendSmsApiService SendSmsApi service
type SendSmsApiService service

type ApiGetOutboundSmsMessageDeliveryReportsRequest struct {
	ctx        _context.Context
	ApiService *SendSmsApiService
	bulkId     *string
	messageId  *string
	limit      *int32
}

func (r ApiGetOutboundSmsMessageDeliveryReportsRequest) BulkId(bulkId string) ApiGetOutboundSmsMessageDeliveryReportsRequest {
	r.bulkId = &bulkId
	return r
}
func (r ApiGetOutboundSmsMessageDeliveryReportsRequest) MessageId(messageId string) ApiGetOutboundSmsMessageDeliveryReportsRequest {
	r.messageId = &messageId
	return r
}
func (r ApiGetOutboundSmsMessageDeliveryReportsRequest) Limit(limit int32) ApiGetOutboundSmsMessageDeliveryReportsRequest {
	r.limit = &limit
	return r
}

func (r ApiGetOutboundSmsMessageDeliveryReportsRequest) Execute() (SmsDeliveryResult, *_nethttp.Response, error) {
	return r.ApiService.GetOutboundSmsMessageDeliveryReportsExecute(r)
}

/*
 * GetOutboundSmsMessageDeliveryReports Get outbound SMS message delivery reports
 * If you are for any reason unable to receive real time delivery reports on your endpoint, you can use this API method to learn if and when the message has been delivered to the recipient. Each request will return a batch of delivery reports - only once. The following API request will return only new reports that arrived since the last API request.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetOutboundSmsMessageDeliveryReportsRequest
 */
func (a *SendSmsApiService) GetOutboundSmsMessageDeliveryReports(ctx _context.Context) ApiGetOutboundSmsMessageDeliveryReportsRequest {
	return ApiGetOutboundSmsMessageDeliveryReportsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return SmsDeliveryResult
 */
func (a *SendSmsApiService) GetOutboundSmsMessageDeliveryReportsExecute(r ApiGetOutboundSmsMessageDeliveryReportsRequest) (SmsDeliveryResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SmsDeliveryResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SendSmsApiService.GetOutboundSmsMessageDeliveryReports")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sms/1/reports"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.bulkId != nil {
		localVarQueryParams.Add("bulkId", parameterToString(*r.bulkId, ""))
	}
	if r.messageId != nil {
		localVarQueryParams.Add("messageId", parameterToString(*r.messageId, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode <= 499 {
			var v SmsApiException
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 && localVarHTTPResponse.StatusCode <= 599 {
			var v SmsApiException
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v SmsDeliveryResult
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetOutboundSmsMessageLogsRequest struct {
	ctx           _context.Context
	ApiService    *SendSmsApiService
	from          *string
	to            *string
	bulkId        *[]string
	messageId     *[]string
	generalStatus *string
	sentSince     *Time
	sentUntil     *Time
	limit         *int32
	mcc           *string
	mnc           *string
}

func (r ApiGetOutboundSmsMessageLogsRequest) From(from string) ApiGetOutboundSmsMessageLogsRequest {
	r.from = &from
	return r
}
func (r ApiGetOutboundSmsMessageLogsRequest) To(to string) ApiGetOutboundSmsMessageLogsRequest {
	r.to = &to
	return r
}
func (r ApiGetOutboundSmsMessageLogsRequest) BulkId(bulkId []string) ApiGetOutboundSmsMessageLogsRequest {
	r.bulkId = &bulkId
	return r
}
func (r ApiGetOutboundSmsMessageLogsRequest) MessageId(messageId []string) ApiGetOutboundSmsMessageLogsRequest {
	r.messageId = &messageId
	return r
}
func (r ApiGetOutboundSmsMessageLogsRequest) GeneralStatus(generalStatus string) ApiGetOutboundSmsMessageLogsRequest {
	r.generalStatus = &generalStatus
	return r
}
func (r ApiGetOutboundSmsMessageLogsRequest) SentSince(sentSince Time) ApiGetOutboundSmsMessageLogsRequest {
	r.sentSince = &sentSince
	return r
}
func (r ApiGetOutboundSmsMessageLogsRequest) SentUntil(sentUntil Time) ApiGetOutboundSmsMessageLogsRequest {
	r.sentUntil = &sentUntil
	return r
}
func (r ApiGetOutboundSmsMessageLogsRequest) Limit(limit int32) ApiGetOutboundSmsMessageLogsRequest {
	r.limit = &limit
	return r
}
func (r ApiGetOutboundSmsMessageLogsRequest) Mcc(mcc string) ApiGetOutboundSmsMessageLogsRequest {
	r.mcc = &mcc
	return r
}
func (r ApiGetOutboundSmsMessageLogsRequest) Mnc(mnc string) ApiGetOutboundSmsMessageLogsRequest {
	r.mnc = &mnc
	return r
}

func (r ApiGetOutboundSmsMessageLogsRequest) Execute() (SmsLogsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetOutboundSmsMessageLogsExecute(r)
}

/*
 * GetOutboundSmsMessageLogs Get outbound SMS message logs
 * You should use this method for displaying logs in the user interface or for some other less frequent usage. See [message delivery reports](#channels/sms/get-outbound-sms-message-delivery-reports) if your use case is to verify message delivery.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetOutboundSmsMessageLogsRequest
 */
func (a *SendSmsApiService) GetOutboundSmsMessageLogs(ctx _context.Context) ApiGetOutboundSmsMessageLogsRequest {
	return ApiGetOutboundSmsMessageLogsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return SmsLogsResponse
 */
func (a *SendSmsApiService) GetOutboundSmsMessageLogsExecute(r ApiGetOutboundSmsMessageLogsRequest) (SmsLogsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SmsLogsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SendSmsApiService.GetOutboundSmsMessageLogs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sms/1/logs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.from != nil {
		localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	}
	if r.to != nil {
		localVarQueryParams.Add("to", parameterToString(*r.to, ""))
	}
	if r.bulkId != nil {
		t := *r.bulkId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("bulkId", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("bulkId", parameterToString(t, "multi"))
		}
	}
	if r.messageId != nil {
		t := *r.messageId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("messageId", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("messageId", parameterToString(t, "multi"))
		}
	}
	if r.generalStatus != nil {
		localVarQueryParams.Add("generalStatus", parameterToString(*r.generalStatus, ""))
	}
	if r.sentSince != nil {
		localVarQueryParams.Add("sentSince", parameterToString(*r.sentSince, ""))
	}
	if r.sentUntil != nil {
		localVarQueryParams.Add("sentUntil", parameterToString(*r.sentUntil, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.mcc != nil {
		localVarQueryParams.Add("mcc", parameterToString(*r.mcc, ""))
	}
	if r.mnc != nil {
		localVarQueryParams.Add("mnc", parameterToString(*r.mnc, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode <= 499 {
			var v SmsApiException
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 && localVarHTTPResponse.StatusCode <= 599 {
			var v SmsApiException
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v SmsLogsResponse
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPreviewSmsMessageRequest struct {
	ctx               _context.Context
	ApiService        *SendSmsApiService
	smsPreviewRequest *SmsPreviewRequest
}

func (r ApiPreviewSmsMessageRequest) SmsPreviewRequest(smsPreviewRequest SmsPreviewRequest) ApiPreviewSmsMessageRequest {
	r.smsPreviewRequest = &smsPreviewRequest
	return r
}

func (r ApiPreviewSmsMessageRequest) Execute() (SmsPreviewResponse, *_nethttp.Response, error) {
	return r.ApiService.PreviewSmsMessageExecute(r)
}

/*
 * PreviewSmsMessage Preview SMS message
 * Avoid unpleasant surprises and check how different message configurations will affect your message text, number of characters and message parts.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPreviewSmsMessageRequest
 */
func (a *SendSmsApiService) PreviewSmsMessage(ctx _context.Context) ApiPreviewSmsMessageRequest {
	return ApiPreviewSmsMessageRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return SmsPreviewResponse
 */
func (a *SendSmsApiService) PreviewSmsMessageExecute(r ApiPreviewSmsMessageRequest) (SmsPreviewResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SmsPreviewResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SendSmsApiService.PreviewSmsMessage")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sms/1/preview"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/xml"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.smsPreviewRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode <= 499 {
			var v SmsApiException
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 && localVarHTTPResponse.StatusCode <= 599 {
			var v SmsApiException
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v SmsPreviewResponse
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSendBinarySmsMessageRequest struct {
	ctx                      _context.Context
	ApiService               *SendSmsApiService
	smsAdvancedBinaryRequest *SmsAdvancedBinaryRequest
}

func (r ApiSendBinarySmsMessageRequest) SmsAdvancedBinaryRequest(smsAdvancedBinaryRequest SmsAdvancedBinaryRequest) ApiSendBinarySmsMessageRequest {
	r.smsAdvancedBinaryRequest = &smsAdvancedBinaryRequest
	return r
}

func (r ApiSendBinarySmsMessageRequest) Execute() (SmsResponse, *_nethttp.Response, error) {
	return r.ApiService.SendBinarySmsMessageExecute(r)
}

/*
 * SendBinarySmsMessage Send binary SMS message
 * Send single or multiple binary messages to one or more destination address.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiSendBinarySmsMessageRequest
 */
func (a *SendSmsApiService) SendBinarySmsMessage(ctx _context.Context) ApiSendBinarySmsMessageRequest {
	return ApiSendBinarySmsMessageRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return SmsResponse
 */
func (a *SendSmsApiService) SendBinarySmsMessageExecute(r ApiSendBinarySmsMessageRequest) (SmsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SmsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SendSmsApiService.SendBinarySmsMessage")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sms/2/binary/advanced"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/xml"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if r.smsAdvancedBinaryRequest != nil {
		r.smsAdvancedBinaryRequest.Default()
	}
	localVarPostBody = r.smsAdvancedBinaryRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode <= 499 {
			var v SmsApiException
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 && localVarHTTPResponse.StatusCode <= 599 {
			var v SmsApiException
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v SmsResponse
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSendSmsMessageRequest struct {
	ctx                       _context.Context
	ApiService                *SendSmsApiService
	smsAdvancedTextualRequest *SmsAdvancedTextualRequest
}

func (r ApiSendSmsMessageRequest) SmsAdvancedTextualRequest(smsAdvancedTextualRequest SmsAdvancedTextualRequest) ApiSendSmsMessageRequest {
	r.smsAdvancedTextualRequest = &smsAdvancedTextualRequest
	return r
}

func (r ApiSendSmsMessageRequest) Execute() (SmsResponse, *_nethttp.Response, error) {
	return r.ApiService.SendSmsMessageExecute(r)
}

/*
 * SendSmsMessage Send SMS message
 * 99% of all use cases can be achieved by using this API method. Everything from sending a simple single message to a single destination, up to batch sending of personalized messages to the thousands of recipients with a single API request. Language, transliteration, scheduling and every advanced feature you can think of is supported.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiSendSmsMessageRequest
 */
func (a *SendSmsApiService) SendSmsMessage(ctx _context.Context) ApiSendSmsMessageRequest {
	return ApiSendSmsMessageRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return SmsResponse
 */
func (a *SendSmsApiService) SendSmsMessageExecute(r ApiSendSmsMessageRequest) (SmsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SmsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SendSmsApiService.SendSmsMessage")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sms/2/text/advanced"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/xml"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if r.smsAdvancedTextualRequest != nil {
		r.smsAdvancedTextualRequest.Default()
	}
	localVarPostBody = r.smsAdvancedTextualRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode <= 499 {
			var v SmsApiException
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 && localVarHTTPResponse.StatusCode <= 599 {
			var v SmsApiException
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v SmsResponse
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
