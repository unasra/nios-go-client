/*
IbClient

Testing RecordaAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package dns

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/unasra/nios-go-client/dns"
)

var readableAttributes = "aws_rte53_record_info,cloud_info,comment,creation_time,creator,ddns_principal,ddns_protected,disable,discovered_data,dns_name,extattrs,forbid_reclamation,ipv4addr,last_queried,ms_ad_user_data,name,reclaimable,shared_record_group,ttl,use_ttl,view,zone"

func TestCreateARecord(t *testing.T) {
	apiClient := dns.NewAPIClient()
	RecordA := dns.RecordA{
		//Ipv4addr: dns.RecordAIpv4addr{
		//	RecordAIpv4addrOneOf: &dns.RecordAIpv4addrOneOf{
		//		ObjectFunction: dns.PtrString("next_available_ip"),
		//		Parameters:     map[string]interface{}{},
		//		ResultField:    dns.PtrString("ips"),
		//		Object:         dns.PtrString("network"),
		//		ObjectParameters: map[string]interface{}{
		//			"network":      "85.85.0.0/16",
		//			"network_view": "default",
		//		},
		//	},
		//	//String: dns.PtrString("192.168.1.1"),
		//},
		Name: "test360.example.com",
		Extattrs: map[string]interface{}{
			"Site": map[string]interface{}{
				"value": "Site1",
			},
		},
	}
	resp, httpRes, err := apiClient.RecordaAPI.Post(context.Background()).RecordA(RecordA).ReturnAsObject(1).ReturnFields2(readableAttributes).Execute()

	if err != nil {
		t.Errorf("Error: %v", err)
	}
	require.Nil(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, httpRes.StatusCode)
}

func TestGetARecords(t *testing.T) {
	apiClient := dns.NewAPIClient()
	resp, httpRes, err := apiClient.RecordaAPI.Get(context.Background()).ReturnFields2("comment").ReturnAsObject(1).
		Body(
			map[string]interface{}{
				"*Site": "Site1",
			},
		).Execute()

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.NotEmpty(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}

func TestGetARecordsasObject(t *testing.T) {
	apiClient := dns.NewAPIClient()
	resp, httpRes, err := apiClient.RecordaAPI.Get(context.Background()).Execute()

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	// for i, record := range resp {
	// 	fmt.Println(i, *record.Name, *record.Ipv4addr)
	// }

	//fmt.Println(*resp[0].Name)
	assert.NotEmpty(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}

func TestGetARecordByReference(t *testing.T) {

	apiClient := dns.NewAPIClient()
	recordaReference := "ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmV4YW1wbGUsdGVzdDgsMTkyLjAuMi4x:test8.example.com/default"

	resp, httpRes, err := apiClient.RecordaAPI.RecordaReferenceGet(context.Background(), recordaReference).ReturnAsObject(1).Execute()
	//RecordaReferenceGet(context.Background(), recordaReference).Execute()

	assert.NotEmpty(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
	assert.Nil(t, err)
}

func TestUpdateARecord(t *testing.T) {
	apiClient := dns.NewAPIClient()
	RecordA := dns.RecordA{
		//Name:    "test8.example.com",
		Comment: dns.PtrString("This is newly updated comment"),
	}
	recordaReference := "ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmV4YW1wbGUsdGVzdDgsMTkyLjAuMi4x:test8.example.com/default"

	resp, httpRes, err := apiClient.RecordaAPI.RecordaReferencePut(context.Background(), recordaReference).RecordA(RecordA).ReturnFields2("comment").ReturnAsObject(1).Execute()

	assert.NotEmpty(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
	assert.Nil(t, err)
	fmt.Println(resp)
}

func TestDeleteARecord(t *testing.T) {
	apiClient := dns.NewAPIClient()
	recordaReference := "ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmV4YW1wbGUsdGVzdDgsMTkyLjAuMi4x:test8.example.com/default"

	resp, err := apiClient.RecordaAPI.RecordaReferenceDelete(context.Background(), recordaReference).ReturnAsObject(1).Execute()

	assert.NotEmpty(t, resp)
	//assert.Equal(t, 200, httpRes.StatusCode)
	assert.Nil(t, err)
	fmt.Println(resp)
}

// func TestRecordaAPIService(t *testing.T) {

// 	apiClient := dns.NewAPIClient()

// 	t.Run("Test RecordaAPIService Get", func(t *testing.T) {

// 		t.Skip("skip test") // remove to run test

// 		resp, httpRes, err := apiClient.RecordaAPI.Get(context.Background()).Execute()

// 		require.Nil(t, err)
// 		require.NotNil(t, resp)
// 		assert.Equal(t, 200, httpRes.StatusCode)

// 	})

// 	t.Run("Test RecordaAPIService Post", func(t *testing.T) {

// 		t.Skip("skip test") // remove to run test

// 		resp, httpRes, err := apiClient.RecordaAPI.Post(context.Background()).Execute()

// 		require.Nil(t, err)
// 		require.NotNil(t, resp)
// 		assert.Equal(t, 200, httpRes.StatusCode)

// 	})

// 	t.Run("Test RecordaAPIService RecordaReferenceDelete", func(t *testing.T) {

// 		t.Skip("skip test") // remove to run test

// 		var recordaReference string

// 		resp, httpRes, err := apiClient.RecordaAPI.RecordaReferenceDelete(context.Background(), recordaReference).Execute()

// 		require.Nil(t, err)
// 		require.NotNil(t, resp)
// 		assert.Equal(t, 200, httpRes.StatusCode)

// 	})

// 	t.Run("Test RecordaAPIService RecordaReferenceGet", func(t *testing.T) {

// 		t.Skip("skip test") // remove to run test

// 		var recordaReference string

// 		resp, httpRes, err := apiClient.RecordaAPI.RecordaReferenceGet(context.Background(), recordaReference).Execute()

// 		require.Nil(t, err)
// 		require.NotNil(t, resp)
// 		assert.Equal(t, 200, httpRes.StatusCode)

// 	})

// 	t.Run("Test RecordaAPIService RecordaReferencePut", func(t *testing.T) {

// 		t.Skip("skip test") // remove to run test

// 		var recordaReference string

// 		resp, httpRes, err := apiClient.RecordaAPI.RecordaReferencePut(context.Background(), recordaReference).Execute()

// 		require.Nil(t, err)
// 		require.NotNil(t, resp)
// 		assert.Equal(t, 200, httpRes.StatusCode)

// 	})

// }
