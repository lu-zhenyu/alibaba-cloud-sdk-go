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

func (client *Client) ListPhotos(request *ListPhotosRequest) (response *ListPhotosResponse, err error) {
	response = CreateListPhotosResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListPhotosWithChan(request *ListPhotosRequest) (<-chan *ListPhotosResponse, <-chan error) {
	responseChan := make(chan *ListPhotosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPhotos(request)
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

func (client *Client) ListPhotosWithCallback(request *ListPhotosRequest, callback func(response *ListPhotosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPhotosResponse
		var err error
		defer close(result)
		response, err = client.ListPhotos(request)
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

type ListPhotosRequest struct {
	*requests.RpcRequest
	Cursor    string `position:"Query" name:"Cursor"`
	Size      string `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	State     string `position:"Query" name:"State"`
	Direction string `position:"Query" name:"Direction"`
}

type ListPhotosResponse struct {
	*responses.BaseResponse
	Code       string `json:"Code"`
	Message    string `json:"Message"`
	NextCursor string `json:"NextCursor"`
	TotalCount int    `json:"TotalCount"`
	RequestId  string `json:"RequestId"`
	Action     string `json:"Action"`
	Photos     []struct {
		Id              int64  `json:"Id"`
		Title           string `json:"Title"`
		FileId          string `json:"FileId"`
		State           string `json:"State"`
		Md5             string `json:"Md5"`
		IsVideo         bool   `json:"IsVideo"`
		Remark          string `json:"Remark"`
		Width           int64  `json:"Width"`
		Height          int64  `json:"Height"`
		Ctime           int64  `json:"Ctime"`
		Mtime           int64  `json:"Mtime"`
		TakenAt         int64  `json:"TakenAt"`
		InactiveTime    int64  `json:"InactiveTime"`
		ShareExpireTime int64  `json:"ShareExpireTime"`
	} `json:"Photos"`
}

func CreateListPhotosRequest() (request *ListPhotosRequest) {
	request = &ListPhotosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListPhotos", "", "")
	return
}

func CreateListPhotosResponse() (response *ListPhotosResponse) {
	response = &ListPhotosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
