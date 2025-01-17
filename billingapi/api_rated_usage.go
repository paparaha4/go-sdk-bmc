/*
Billing API

Automate your infrastructure billing with the Bare Metal Cloud Billing API. Reserve your server instances to ensure guaranteed resource availability for 12, 24, and 36 months. Retrieve your server’s rated usage for a given period and enable or disable auto-renewals.<br> <br> <span class='pnap-api-knowledge-base-link'> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/phoenixnap-bare-metal-cloud-billing-models' target='_blank'>here</a> </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/billing/v1/)</b>

API version: 0.1
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package billingapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

type RatedUsageApi interface {

	/*
		RatedUsageGet List the rated usage.

		Retrieves all rated usage for given time period. The information is presented as a list of rated usage records. Every record corresponds to a charge. All date & times are in UTC.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiRatedUsageGetRequest
	*/
	RatedUsageGet(ctx context.Context) ApiRatedUsageGetRequest

	// RatedUsageGetExecute executes the request
	//  @return []RatedUsageGet200ResponseInner
	RatedUsageGetExecute(r ApiRatedUsageGetRequest) ([]RatedUsageGet200ResponseInner, *http.Response, error)

	/*
		RatedUsageMonthToDateGet List the rated usage records for the current calendar month.

		Retrieves all rated usage for the current calendar month. The information is presented as a list of rated usage records. Every record corresponds to a charge. All date & times are in UTC.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiRatedUsageMonthToDateGetRequest
	*/
	RatedUsageMonthToDateGet(ctx context.Context) ApiRatedUsageMonthToDateGetRequest

	// RatedUsageMonthToDateGetExecute executes the request
	//  @return []RatedUsageGet200ResponseInner
	RatedUsageMonthToDateGetExecute(r ApiRatedUsageMonthToDateGetRequest) ([]RatedUsageGet200ResponseInner, *http.Response, error)
}

// RatedUsageApiService RatedUsageApi service
type RatedUsageApiService service

type ApiRatedUsageGetRequest struct {
	ctx             context.Context
	ApiService      RatedUsageApi
	fromYearMonth   *string
	toYearMonth     *string
	productCategory *ProductCategoryEnum
}

// From year month (inclusive) to filter rated usage records by.
func (r ApiRatedUsageGetRequest) FromYearMonth(fromYearMonth string) ApiRatedUsageGetRequest {
	r.fromYearMonth = &fromYearMonth
	return r
}

// To year month (inclusive) to filter rated usage records by.
func (r ApiRatedUsageGetRequest) ToYearMonth(toYearMonth string) ApiRatedUsageGetRequest {
	r.toYearMonth = &toYearMonth
	return r
}

// The product category
func (r ApiRatedUsageGetRequest) ProductCategory(productCategory ProductCategoryEnum) ApiRatedUsageGetRequest {
	r.productCategory = &productCategory
	return r
}

func (r ApiRatedUsageGetRequest) Execute() ([]RatedUsageGet200ResponseInner, *http.Response, error) {
	return r.ApiService.RatedUsageGetExecute(r)
}

/*
RatedUsageGet List the rated usage.

Retrieves all rated usage for given time period. The information is presented as a list of rated usage records. Every record corresponds to a charge. All date & times are in UTC.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRatedUsageGetRequest
*/
func (a *RatedUsageApiService) RatedUsageGet(ctx context.Context) ApiRatedUsageGetRequest {
	return ApiRatedUsageGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return []RatedUsageGet200ResponseInner
func (a *RatedUsageApiService) RatedUsageGetExecute(r ApiRatedUsageGetRequest) ([]RatedUsageGet200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []RatedUsageGet200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RatedUsageApiService.RatedUsageGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rated-usage"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.fromYearMonth == nil {
		return localVarReturnValue, nil, reportError("fromYearMonth is required and must be specified")
	}
	if r.toYearMonth == nil {
		return localVarReturnValue, nil, reportError("toYearMonth is required and must be specified")
	}

	localVarQueryParams.Add("fromYearMonth", parameterToString(*r.fromYearMonth, ""))
	localVarQueryParams.Add("toYearMonth", parameterToString(*r.toYearMonth, ""))
	if r.productCategory != nil {
		localVarQueryParams.Add("productCategory", parameterToString(*r.productCategory, ""))
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
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
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

type ApiRatedUsageMonthToDateGetRequest struct {
	ctx             context.Context
	ApiService      RatedUsageApi
	productCategory *ProductCategoryEnum
}

// The product category
func (r ApiRatedUsageMonthToDateGetRequest) ProductCategory(productCategory ProductCategoryEnum) ApiRatedUsageMonthToDateGetRequest {
	r.productCategory = &productCategory
	return r
}

func (r ApiRatedUsageMonthToDateGetRequest) Execute() ([]RatedUsageGet200ResponseInner, *http.Response, error) {
	return r.ApiService.RatedUsageMonthToDateGetExecute(r)
}

/*
RatedUsageMonthToDateGet List the rated usage records for the current calendar month.

Retrieves all rated usage for the current calendar month. The information is presented as a list of rated usage records. Every record corresponds to a charge. All date & times are in UTC.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRatedUsageMonthToDateGetRequest
*/
func (a *RatedUsageApiService) RatedUsageMonthToDateGet(ctx context.Context) ApiRatedUsageMonthToDateGetRequest {
	return ApiRatedUsageMonthToDateGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return []RatedUsageGet200ResponseInner
func (a *RatedUsageApiService) RatedUsageMonthToDateGetExecute(r ApiRatedUsageMonthToDateGetRequest) ([]RatedUsageGet200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []RatedUsageGet200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RatedUsageApiService.RatedUsageMonthToDateGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rated-usage/month-to-date"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.productCategory != nil {
		localVarQueryParams.Add("productCategory", parameterToString(*r.productCategory, ""))
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
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
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
