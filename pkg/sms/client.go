package sms

// This is a generated file and is not intended for modification!

import (
	"fmt"
	api "infobip-api-go-client/internal"
)

var (
	// reference imported packages just in case auto-generated code doesn't
	_ = fmt.Sprint
	_ = api.CallAPI
)

type Client api.Client

func (client *Client) GetSentSMSDeliveryReports(query GetSentSMSDeliveryReportsQuery) (*SMSReportResponse, error) {
	path := "/sms/1/reports"
	method := "GET"
	var response SMSReportResponse

	err := api.CallAPI(downcast(client), method, path, nil, query, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (client *Client) SendMultipleTextualSMSAdvanced(request SMSAdvancedTextualRequest) (*SMSResponse, error) {
	path := "/sms/1/text/advanced"
	method := "POST"
	var response SMSResponse

	err := api.CallAPI(downcast(client), method, path, nil, nil, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func downcast(client *Client) *api.Client {
	c := api.Client(*client)
	return &c
}
