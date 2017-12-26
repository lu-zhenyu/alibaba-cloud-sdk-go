package ecs

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

func (client *Client) ModifyVRouterAttribute(request *ModifyVRouterAttributeRequest) (response *ModifyVRouterAttributeResponse, err error) {
	response = CreateModifyVRouterAttributeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyVRouterAttributeWithChan(request *ModifyVRouterAttributeRequest) (<-chan *ModifyVRouterAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyVRouterAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVRouterAttribute(request)
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

func (client *Client) ModifyVRouterAttributeWithCallback(request *ModifyVRouterAttributeRequest, callback func(response *ModifyVRouterAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVRouterAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyVRouterAttribute(request)
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

type ModifyVRouterAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	VRouterId            string           `position:"Query" name:"VRouterId"`
	Description          string           `position:"Query" name:"Description"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	VRouterName          string           `position:"Query" name:"VRouterName"`
}

type ModifyVRouterAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateModifyVRouterAttributeRequest() (request *ModifyVRouterAttributeRequest) {
	request = &ModifyVRouterAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyVRouterAttribute", "", "")
	return
}

func CreateModifyVRouterAttributeResponse() (response *ModifyVRouterAttributeResponse) {
	response = &ModifyVRouterAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
