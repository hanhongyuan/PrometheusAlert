package market

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

// InstanceItem is a nested struct in market response
type InstanceItem struct {
	InstanceId     int64  `json:"InstanceId" xml:"InstanceId"`
	OrderId        int64  `json:"OrderId" xml:"OrderId"`
	SupplierName   string `json:"SupplierName" xml:"SupplierName"`
	ProductCode    string `json:"ProductCode" xml:"ProductCode"`
	ProductSkuCode string `json:"ProductSkuCode" xml:"ProductSkuCode"`
	ProductName    string `json:"ProductName" xml:"ProductName"`
	ProductType    string `json:"ProductType" xml:"ProductType"`
	Status         string `json:"Status" xml:"Status"`
	BeganOn        int64  `json:"BeganOn" xml:"BeganOn"`
	EndOn          int64  `json:"EndOn" xml:"EndOn"`
	CreatedOn      int64  `json:"CreatedOn" xml:"CreatedOn"`
	ExtendJson     string `json:"ExtendJson" xml:"ExtendJson"`
	HostJson       string `json:"HostJson" xml:"HostJson"`
	AppJson        string `json:"AppJson" xml:"AppJson"`
	ApiJson        string `json:"ApiJson" xml:"ApiJson"`
	ImageJson      string `json:"ImageJson" xml:"ImageJson"`
	IdaasJson      string `json:"IdaasJson" xml:"IdaasJson"`
}
