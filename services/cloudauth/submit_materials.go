package cloudauth

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

// SubmitMaterials invokes the cloudauth.SubmitMaterials API synchronously
// api document: https://help.aliyun.com/api/cloudauth/submitmaterials.html
func (client *Client) SubmitMaterials(request *SubmitMaterialsRequest) (response *SubmitMaterialsResponse, err error) {
	response = CreateSubmitMaterialsResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitMaterialsWithChan invokes the cloudauth.SubmitMaterials API asynchronously
// api document: https://help.aliyun.com/api/cloudauth/submitmaterials.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitMaterialsWithChan(request *SubmitMaterialsRequest) (<-chan *SubmitMaterialsResponse, <-chan error) {
	responseChan := make(chan *SubmitMaterialsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitMaterials(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// SubmitMaterialsWithCallback invokes the cloudauth.SubmitMaterials API asynchronously
// api document: https://help.aliyun.com/api/cloudauth/submitmaterials.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitMaterialsWithCallback(request *SubmitMaterialsRequest, callback func(response *SubmitMaterialsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitMaterialsResponse
		var err error
		defer close(result)
		response, err = client.SubmitMaterials(request)
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

// SubmitMaterialsRequest is the request struct for api SubmitMaterials
type SubmitMaterialsRequest struct {
	*requests.RpcRequest
	SourceIp        string                     `position:"Query" name:"SourceIp"`
	ResourceOwnerId requests.Integer           `position:"Query" name:"ResourceOwnerId"`
	VerifyToken     string                     `position:"Query" name:"VerifyToken"`
	Material        *[]SubmitMaterialsMaterial `position:"Query" name:"Material"  type:"Repeated"`
}

// SubmitMaterialsMaterial is a repeated param struct in SubmitMaterialsRequest
type SubmitMaterialsMaterial struct {
	MaterialType string `name:"MaterialType"`
	Value        string `name:"Value"`
}

// SubmitMaterialsResponse is the response struct for api SubmitMaterials
type SubmitMaterialsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateSubmitMaterialsRequest creates a request to invoke SubmitMaterials API
func CreateSubmitMaterialsRequest() (request *SubmitMaterialsRequest) {
	request = &SubmitMaterialsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2018-05-04", "SubmitMaterials", "cloudauth", "openAPI")
	return
}

// CreateSubmitMaterialsResponse creates a response to parse from SubmitMaterials response
func CreateSubmitMaterialsResponse() (response *SubmitMaterialsResponse) {
	response = &SubmitMaterialsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
