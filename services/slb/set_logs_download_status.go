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

func (client *Client) SetLogsDownloadStatus(request *SetLogsDownloadStatusRequest) (response *SetLogsDownloadStatusResponse, err error) {
	response = CreateSetLogsDownloadStatusResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SetLogsDownloadStatusWithChan(request *SetLogsDownloadStatusRequest) (<-chan *SetLogsDownloadStatusResponse, <-chan error) {
	responseChan := make(chan *SetLogsDownloadStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetLogsDownloadStatus(request)
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

func (client *Client) SetLogsDownloadStatusWithCallback(request *SetLogsDownloadStatusRequest, callback func(response *SetLogsDownloadStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetLogsDownloadStatusResponse
		var err error
		defer close(result)
		response, err = client.SetLogsDownloadStatus(request)
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

type SetLogsDownloadStatusRequest struct {
	*requests.RpcRequest
	Tags                 string           `position:"Query" name:"Tags"`
	RoleName             string           `position:"Query" name:"RoleName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	LogsDownloadStatus   string           `position:"Query" name:"LogsDownloadStatus"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type SetLogsDownloadStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateSetLogsDownloadStatusRequest() (request *SetLogsDownloadStatusRequest) {
	request = &SetLogsDownloadStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "SetLogsDownloadStatus", "", "")
	return
}

func CreateSetLogsDownloadStatusResponse() (response *SetLogsDownloadStatusResponse) {
	response = &SetLogsDownloadStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
