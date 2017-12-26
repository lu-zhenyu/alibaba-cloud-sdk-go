package dm

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

func (client *Client) DeleteTemplate(request *DeleteTemplateRequest) (response *DeleteTemplateResponse, err error) {
	response = CreateDeleteTemplateResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DeleteTemplateWithChan(request *DeleteTemplateRequest) (<-chan *DeleteTemplateResponse, <-chan error) {
	responseChan := make(chan *DeleteTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteTemplate(request)
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

func (client *Client) DeleteTemplateWithCallback(request *DeleteTemplateRequest, callback func(response *DeleteTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteTemplateResponse
		var err error
		defer close(result)
		response, err = client.DeleteTemplate(request)
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

type DeleteTemplateRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	TemplateId           requests.Integer `position:"Query" name:"TemplateId"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	FromType             requests.Integer `position:"Query" name:"FromType"`
}

type DeleteTemplateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateDeleteTemplateRequest() (request *DeleteTemplateRequest) {
	request = &DeleteTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "DeleteTemplate", "", "")
	return
}

func CreateDeleteTemplateResponse() (response *DeleteTemplateResponse) {
	response = &DeleteTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
