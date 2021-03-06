package edas

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

// Application is a nested struct in edas response
type Application struct {
	Name                 string `json:"Name" xml:"Name"`
	ClusterId            string `json:"ClusterId" xml:"ClusterId"`
	ApplicationId        string `json:"ApplicationId" xml:"ApplicationId"`
	Port                 int    `json:"Port" xml:"Port"`
	Email                string `json:"Email" xml:"Email"`
	Memory               int    `json:"Memory" xml:"Memory"`
	LaunchTime           int    `json:"LaunchTime" xml:"LaunchTime"`
	Owner                string `json:"Owner" xml:"Owner"`
	CreateTime           int    `json:"CreateTime" xml:"CreateTime"`
	Dockerize            bool   `json:"Dockerize" xml:"Dockerize"`
	BuildPackageId       int    `json:"BuildPackageId" xml:"BuildPackageId"`
	InstanceCount        int    `json:"InstanceCount" xml:"InstanceCount"`
	RegionId             string `json:"RegionId" xml:"RegionId"`
	HealthCheckUrl       string `json:"HealthCheckUrl" xml:"HealthCheckUrl"`
	UserId               string `json:"UserId" xml:"UserId"`
	Cpu                  int    `json:"Cpu" xml:"Cpu"`
	Phone                string `json:"Phone" xml:"Phone"`
	RunningInstanceCount int    `json:"RunningInstanceCount" xml:"RunningInstanceCount"`
}
