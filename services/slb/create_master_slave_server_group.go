package slb

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

func (client *Client) CreateMasterSlaveServerGroup(request *CreateMasterSlaveServerGroupRequest) (response *CreateMasterSlaveServerGroupResponse, err error) {
	response = CreateCreateMasterSlaveServerGroupResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateMasterSlaveServerGroupWithChan(request *CreateMasterSlaveServerGroupRequest) (<-chan *CreateMasterSlaveServerGroupResponse, <-chan error) {
	responseChan := make(chan *CreateMasterSlaveServerGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMasterSlaveServerGroup(request)
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

func (client *Client) CreateMasterSlaveServerGroupWithCallback(request *CreateMasterSlaveServerGroupRequest, callback func(response *CreateMasterSlaveServerGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMasterSlaveServerGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateMasterSlaveServerGroup(request)
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

type CreateMasterSlaveServerGroupRequest struct {
	*requests.RpcRequest
	Tags                       string           `position:"Query" name:"Tags"`
	MasterSlaveBackendServers  string           `position:"Query" name:"MasterSlaveBackendServers"`
	MasterSlaveServerGroupName string           `position:"Query" name:"MasterSlaveServerGroupName"`
	ResourceOwnerAccount       string           `position:"Query" name:"ResourceOwnerAccount"`
	AccessKeyId                string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId            requests.Integer `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId             string           `position:"Query" name:"LoadBalancerId"`
	OwnerAccount               string           `position:"Query" name:"OwnerAccount"`
	OwnerId                    requests.Integer `position:"Query" name:"OwnerId"`
}

type CreateMasterSlaveServerGroupResponse struct {
	*responses.BaseResponse
	RequestId                 string `json:"RequestId" xml:"RequestId"`
	MasterSlaveServerGroupId  string `json:"MasterSlaveServerGroupId" xml:"MasterSlaveServerGroupId"`
	MasterSlaveBackendServers struct {
		MasterSlaveBackendServer []struct {
			ServerId   string           `json:"ServerId" xml:"ServerId"`
			Port       requests.Integer `json:"Port" xml:"Port"`
			Weight     requests.Integer `json:"Weight" xml:"Weight"`
			ServerType string           `json:"ServerType" xml:"ServerType"`
		} `json:"MasterSlaveBackendServer" xml:"MasterSlaveBackendServer"`
	} `json:"MasterSlaveBackendServers" xml:"MasterSlaveBackendServers"`
}

func CreateCreateMasterSlaveServerGroupRequest() (request *CreateMasterSlaveServerGroupRequest) {
	request = &CreateMasterSlaveServerGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "CreateMasterSlaveServerGroup", "", "")
	return
}

func CreateCreateMasterSlaveServerGroupResponse() (response *CreateMasterSlaveServerGroupResponse) {
	response = &CreateMasterSlaveServerGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
