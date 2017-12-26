package vpc

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

func (client *Client) AssociatePhysicalConnectionToVirtualBorderRouter(request *AssociatePhysicalConnectionToVirtualBorderRouterRequest) (response *AssociatePhysicalConnectionToVirtualBorderRouterResponse, err error) {
	response = CreateAssociatePhysicalConnectionToVirtualBorderRouterResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) AssociatePhysicalConnectionToVirtualBorderRouterWithChan(request *AssociatePhysicalConnectionToVirtualBorderRouterRequest) (<-chan *AssociatePhysicalConnectionToVirtualBorderRouterResponse, <-chan error) {
	responseChan := make(chan *AssociatePhysicalConnectionToVirtualBorderRouterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AssociatePhysicalConnectionToVirtualBorderRouter(request)
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

func (client *Client) AssociatePhysicalConnectionToVirtualBorderRouterWithCallback(request *AssociatePhysicalConnectionToVirtualBorderRouterRequest, callback func(response *AssociatePhysicalConnectionToVirtualBorderRouterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AssociatePhysicalConnectionToVirtualBorderRouterResponse
		var err error
		defer close(result)
		response, err = client.AssociatePhysicalConnectionToVirtualBorderRouter(request)
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

type AssociatePhysicalConnectionToVirtualBorderRouterRequest struct {
	*requests.RpcRequest
	VlanId               string           `position:"Query" name:"VlanId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	UserCidr             string           `position:"Query" name:"UserCidr"`
	PhysicalConnectionId string           `position:"Query" name:"PhysicalConnectionId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	CircuitCode          string           `position:"Query" name:"CircuitCode"`
	LocalGatewayIp       string           `position:"Query" name:"LocalGatewayIp"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	PeerGatewayIp        string           `position:"Query" name:"PeerGatewayIp"`
	PeeringSubnetMask    string           `position:"Query" name:"PeeringSubnetMask"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	VbrId                string           `position:"Query" name:"VbrId"`
}

type AssociatePhysicalConnectionToVirtualBorderRouterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateAssociatePhysicalConnectionToVirtualBorderRouterRequest() (request *AssociatePhysicalConnectionToVirtualBorderRouterRequest) {
	request = &AssociatePhysicalConnectionToVirtualBorderRouterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "AssociatePhysicalConnectionToVirtualBorderRouter", "", "")
	return
}

func CreateAssociatePhysicalConnectionToVirtualBorderRouterResponse() (response *AssociatePhysicalConnectionToVirtualBorderRouterResponse) {
	response = &AssociatePhysicalConnectionToVirtualBorderRouterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
