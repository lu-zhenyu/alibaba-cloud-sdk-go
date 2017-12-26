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

func (client *Client) MigrateMarket(request *MigrateMarketRequest) (response *MigrateMarketResponse, err error) {
	response = CreateMigrateMarketResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) MigrateMarketWithChan(request *MigrateMarketRequest) (<-chan *MigrateMarketResponse, <-chan error) {
	responseChan := make(chan *MigrateMarketResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MigrateMarket(request)
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

func (client *Client) MigrateMarketWithCallback(request *MigrateMarketRequest, callback func(response *MigrateMarketResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MigrateMarketResponse
		var err error
		defer close(result)
		response, err = client.MigrateMarket(request)
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

type MigrateMarketRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	FromType             requests.Integer `position:"Query" name:"FromType"`
	Version              string           `position:"Query" name:"Version"`
}

type MigrateMarketResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateMigrateMarketRequest() (request *MigrateMarketRequest) {
	request = &MigrateMarketRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "MigrateMarket", "", "")
	return
}

func CreateMigrateMarketResponse() (response *MigrateMarketResponse) {
	response = &MigrateMarketResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
