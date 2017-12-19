package cloudphoto

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

func (client *Client) ListPhotoStores(request *ListPhotoStoresRequest) (response *ListPhotoStoresResponse, err error) {
	response = CreateListPhotoStoresResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListPhotoStoresWithChan(request *ListPhotoStoresRequest) (<-chan *ListPhotoStoresResponse, <-chan error) {
	responseChan := make(chan *ListPhotoStoresResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPhotoStores(request)
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

func (client *Client) ListPhotoStoresWithCallback(request *ListPhotoStoresRequest, callback func(response *ListPhotoStoresResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPhotoStoresResponse
		var err error
		defer close(result)
		response, err = client.ListPhotoStores(request)
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

type ListPhotoStoresRequest struct {
	*requests.RpcRequest
}

type ListPhotoStoresResponse struct {
	*responses.BaseResponse
	Code        string `json:"Code"`
	Message     string `json:"Message"`
	RequestId   string `json:"RequestId"`
	Action      string `json:"Action"`
	PhotoStores []struct {
		Id               int64  `json:"Id"`
		Name             string `json:"Name"`
		Remark           string `json:"Remark"`
		AutoCleanEnabled bool   `json:"AutoCleanEnabled"`
		AutoCleanDays    int    `json:"AutoCleanDays"`
		DefaultQuota     int64  `json:"DefaultQuota"`
		Ctime            int64  `json:"Ctime"`
		Mtime            int64  `json:"Mtime"`
		Buckets          []struct {
			Name   string `json:"Name"`
			Region string `json:"Region"`
			State  string `json:"State"`
		} `json:"Buckets"`
	} `json:"PhotoStores"`
}

func CreateListPhotoStoresRequest() (request *ListPhotoStoresRequest) {
	request = &ListPhotoStoresRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListPhotoStores", "", "")
	return
}

func CreateListPhotoStoresResponse() (response *ListPhotoStoresResponse) {
	response = &ListPhotoStoresResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
