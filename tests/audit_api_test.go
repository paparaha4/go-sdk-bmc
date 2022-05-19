package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/phoenixnap/go-sdk-bmc/auditapi"
	"github.com/stretchr/testify/suite"
)

type AuditApiTestSuite struct {
	suite.Suite
	ctx           context.Context
	configuration *auditapi.Configuration
}

func (suite *AuditApiTestSuite) SetupSuite() {
	// Set configuration
	suite.configuration = auditapi.NewConfiguration()
	suite.configuration.Host = "127.0.0.1:1080"
	suite.configuration.Scheme = "http"
	// Set the context background
	suite.ctx = context.WithValue(context.Background(), "accessToken", "ACCESSTOKENSTRING")
	suite.ctx = context.WithValue(suite.ctx, "serverIndex", nil)
	fmt.Println(">>> From SetupSuite")
}

// this function executes after all tests executed
func (suite *AuditApiTestSuite) TearDownSuite() {
	fmt.Println(">>> From TearDownSuite")
	TestUtilsImpl{}.reset_expectations()
}

func verify_called_once(suite *AuditApiTestSuite, expectationId string) {
	// Result retrieved from server's verification
	// Verifying expectation matched exactly once.
	verifyResult := TestUtilsImpl{}.verify_expectation_matched_times(expectationId, 1)

	// Asserts a successful result
	suite.Equal(202, verifyResult.StatusCode)
}

func (suite *AuditApiTestSuite) Test_get_events_all_query_params() {

	// Generate payload
	request, response := TestUtilsImpl{}.generate_payloads_from("auditapi/events_get", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setup_expectation(request, response, 1)

	// Fetch a map of query parameters
	qpMap := TestUtilsImpl{}.generate_query_params(request)

	// Converting to int32
	limitstring, ok := qpMap["limit"].(string)
	if !ok {
		panic(ok)
	}

	from, _ := time.Parse(time.RFC3339, request.QueryStringParameters[0].Values[0]) // time.Time | From the date and time (inclusive) to filter event log records by. (optional)
	to, _ := time.Parse(time.RFC3339, request.QueryStringParameters[1].Values[0])   // time.Time | To the date and time (inclusive) to filter event log records by. (optional)
	limit64, _ := strconv.ParseInt(limitstring, 10, 32)
	limit := int32(limit64)                          // int32 | Limit the number of records returned. (optional)
	order := fmt.Sprintf("%v", qpMap["order"])       // string | Ordering of the event's time. SortBy can be introduced later on. (optional) (default to "ASC")
	username := fmt.Sprintf("%v", qpMap["username"]) // string | The username that did the actions. (optional)
	verb := fmt.Sprintf("%v", qpMap["verb"])         // string | The HTTP verb corresponding to the action. (optional)
	uri := fmt.Sprintf("%v", qpMap["uri"])

	// New  ApiClient
	apiClient := auditapi.NewAPIClient(suite.configuration)

	// Operation Execution
	result, r, err := apiClient.EventsApi.EventsGet(suite.ctx).From(from).To(to).Limit(limit).Order(order).Username(username).Verb(verb).Uri(uri).Execute() //.From(from).To(to).Limit(limit).Order(order).Username(username).Verb(verb).Uri(uri).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	verify_called_once(suite, expectationId)
}

func TestAuditApiTestSuite(t *testing.T) {
	suite.Run(t, new(AuditApiTestSuite))
}
