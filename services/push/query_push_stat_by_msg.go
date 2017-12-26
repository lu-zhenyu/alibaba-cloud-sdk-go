package push

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) QueryPushStatByMsg(request *QueryPushStatByMsgRequest) (response *QueryPushStatByMsgResponse, err error) {
	response = CreateQueryPushStatByMsgResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryPushStatByMsgWithChan(request *QueryPushStatByMsgRequest) (<-chan *QueryPushStatByMsgResponse, <-chan error) {
	responseChan := make(chan *QueryPushStatByMsgResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryPushStatByMsg(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) QueryPushStatByMsgWithCallback(request *QueryPushStatByMsgRequest, callback func(response *QueryPushStatByMsgResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryPushStatByMsgResponse
		var err error
		defer close(result)
		response, err = client.QueryPushStatByMsg(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type QueryPushStatByMsgRequest struct {
	*requests.RpcRequest
	AppKey    requests.Integer `position:"Query" name:"AppKey"`
	MessageId string           `position:"Query" name:"MessageId"`
}

type QueryPushStatByMsgResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	PushStats struct {
		PushStat []struct {
			MessageId              string           `json:"MessageId" xml:"MessageId"`
			AcceptCount            requests.Integer `json:"AcceptCount" xml:"AcceptCount"`
			SentCount              requests.Integer `json:"SentCount" xml:"SentCount"`
			ReceivedCount          requests.Integer `json:"ReceivedCount" xml:"ReceivedCount"`
			OpenedCount            requests.Integer `json:"OpenedCount" xml:"OpenedCount"`
			DeletedCount           requests.Integer `json:"DeletedCount" xml:"DeletedCount"`
			SmsSentCount           requests.Integer `json:"SmsSentCount" xml:"SmsSentCount"`
			SmsSkipCount           requests.Integer `json:"SmsSkipCount" xml:"SmsSkipCount"`
			SmsFailedCount         requests.Integer `json:"SmsFailedCount" xml:"SmsFailedCount"`
			SmsReceiveSuccessCount requests.Integer `json:"SmsReceiveSuccessCount" xml:"SmsReceiveSuccessCount"`
			SmsReceiveFailedCount  requests.Integer `json:"SmsReceiveFailedCount" xml:"SmsReceiveFailedCount"`
		} `json:"PushStat" xml:"PushStat"`
	} `json:"PushStats" xml:"PushStats"`
}

func CreateQueryPushStatByMsgRequest() (request *QueryPushStatByMsgRequest) {
	request = &QueryPushStatByMsgRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "QueryPushStatByMsg", "", "")
	return
}

func CreateQueryPushStatByMsgResponse() (response *QueryPushStatByMsgResponse) {
	response = &QueryPushStatByMsgResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
