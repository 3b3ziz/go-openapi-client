/*
OpenAPI Petstore

Testing StoreAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	openapiclient "github.com/3b3ziz/go-openapi-client/clients/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_StoreAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test StoreAPIService DeleteOrder", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orderId string

		httpRes, err := apiClient.StoreAPI.DeleteOrder(context.Background(), orderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StoreAPIService GetInventory", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.StoreAPI.GetInventory(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StoreAPIService GetOrderById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orderId int64

		resp, httpRes, err := apiClient.StoreAPI.GetOrderById(context.Background(), orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StoreAPIService PlaceOrder", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.StoreAPI.PlaceOrder(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
