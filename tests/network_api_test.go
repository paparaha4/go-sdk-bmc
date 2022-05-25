package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/networkapi"
	"github.com/stretchr/testify/suite"
)

type NetworkApiTestSuite struct {
	suite.Suite
	ctx           context.Context
	configuration *networkapi.Configuration
	apiClient     *networkapi.APIClient
}

// this function executes before each test
func (suite *NetworkApiTestSuite) SetupSuite() {
	// Set configuration
	suite.configuration = networkapi.NewConfiguration()
	suite.configuration.Host = "127.0.0.1:1080"
	suite.configuration.Scheme = "http"
	// Set the context background
	suite.ctx = context.WithValue(context.Background(), "accessToken", "ACCESSTOKENSTRING")
	suite.ctx = context.WithValue(suite.ctx, "serverIndex", nil)
	fmt.Println(">>> From SetupSuite")
	suite.apiClient = networkapi.NewAPIClient(suite.configuration)
}

// this function executes after each test
func (suite *NetworkApiTestSuite) TearDownTest() {
	fmt.Println(">>> From TearDownSuite")
	defer TestUtilsImpl{}.resetExpectations()
}

func (suite *NetworkApiTestSuite) verifyCalledOnce(expectationId string) {
	// Result retrieved from server's verification
	// Verifying expectation matched exactly once.
	verifyResult := TestUtilsImpl{}.verifyExpectationMatchedTimes(expectationId, 1)

	// Asserts a successful result
	suite.Equal(202, verifyResult.StatusCode)
}

func (suite *NetworkApiTestSuite) TestGetPrivateNetworks() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkapi/networks_get", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Fetch a map of query parameters
	qpMap := TestUtilsImpl{}.generateQueryParams(request)

	loc := fmt.Sprintf("%v", qpMap["location"])

	// Operation Execution
	result, _, _ := suite.apiClient.PrivateNetworksApi.PrivateNetworksGet(suite.ctx).Location(loc).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *NetworkApiTestSuite) TestCreatePrivateNetwork() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkapi/networks_post", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	byteData := TestUtilsImpl{}.extractRequestBody(request)
	var networkCreate networkapi.PrivateNetworkCreate
	json.Unmarshal(byteData, &networkCreate)

	// Operation Execution
	result, _, _ := suite.apiClient.PrivateNetworksApi.PrivateNetworksPost(suite.ctx).PrivateNetworkCreate(networkCreate).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	suite.verifyCalledOnce(expectationId)
}

// func (suite *NetworkApiTestSuite) TestGetPrivateNetworkById() {

// }

// func (suite *NetworkApiTestSuite) TestPutPrivateNetworkById() {

// }

// func (suite *NetworkApiTestSuite) TestDeletePrivateNetworkById() {

// }

// func (suite *NetworkApiTestSuite) TestGetPublicNetworks() {

// }

// func (suite *NetworkApiTestSuite) TestCreatePublicNetwork() {

// }

// func (suite *NetworkApiTestSuite) TestGetPublicNetworkById() {

// }

// func (suite *NetworkApiTestSuite) TestPatchPublicNetworkById() {

// }

// func (suite *NetworkApiTestSuite) TestDeletePublicNetworkById() {

// }

// func (suite *NetworkApiTestSuite) TestPostPublicNetworkIpBlockByPublicNetworkId() {

// }

// func (suite *NetworkApiTestSuite) TestDeletePublicNetworkIpBlocksByPublicNetworkIdAndIpBlockID() {

// }

func TestNetworkApiTestSuite(t *testing.T) {
	suite.Run(t, new(NetworkApiTestSuite))
}
