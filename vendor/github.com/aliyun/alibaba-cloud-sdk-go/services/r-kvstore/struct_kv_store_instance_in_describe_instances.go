package r_kvstore

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

// KVStoreInstanceInDescribeInstances is a nested struct in r_kvstore response
type KVStoreInstanceInDescribeInstances struct {
	VpcId               string                  `json:"VpcId" xml:"VpcId"`
	PrivateIp           string                  `json:"PrivateIp" xml:"PrivateIp"`
	Capacity            int64                   `json:"Capacity" xml:"Capacity"`
	ReplacateId         string                  `json:"ReplacateId" xml:"ReplacateId"`
	CreateTime          string                  `json:"CreateTime" xml:"CreateTime"`
	ConnectionDomain    string                  `json:"ConnectionDomain" xml:"ConnectionDomain"`
	IsRds               bool                    `json:"IsRds" xml:"IsRds"`
	ChargeType          string                  `json:"ChargeType" xml:"ChargeType"`
	ArchitectureType    string                  `json:"ArchitectureType" xml:"ArchitectureType"`
	NetworkType         string                  `json:"NetworkType" xml:"NetworkType"`
	ConnectionMode      string                  `json:"ConnectionMode" xml:"ConnectionMode"`
	Port                int64                   `json:"Port" xml:"Port"`
	SecondaryZoneId     string                  `json:"SecondaryZoneId" xml:"SecondaryZoneId"`
	EngineVersion       string                  `json:"EngineVersion" xml:"EngineVersion"`
	PackageType         string                  `json:"PackageType" xml:"PackageType"`
	Config              string                  `json:"Config" xml:"Config"`
	VpcCloudInstanceId  string                  `json:"VpcCloudInstanceId" xml:"VpcCloudInstanceId"`
	Bandwidth           int64                   `json:"Bandwidth" xml:"Bandwidth"`
	InstanceName        string                  `json:"InstanceName" xml:"InstanceName"`
	ShardCount          int                     `json:"ShardCount" xml:"ShardCount"`
	UserName            string                  `json:"UserName" xml:"UserName"`
	GlobalInstanceId    string                  `json:"GlobalInstanceId" xml:"GlobalInstanceId"`
	QPS                 int64                   `json:"QPS" xml:"QPS"`
	InstanceClass       string                  `json:"InstanceClass" xml:"InstanceClass"`
	DestroyTime         string                  `json:"DestroyTime" xml:"DestroyTime"`
	InstanceId          string                  `json:"InstanceId" xml:"InstanceId"`
	InstanceType        string                  `json:"InstanceType" xml:"InstanceType"`
	HasRenewChangeOrder bool                    `json:"HasRenewChangeOrder" xml:"HasRenewChangeOrder"`
	RegionId            string                  `json:"RegionId" xml:"RegionId"`
	SearchKey           string                  `json:"SearchKey" xml:"SearchKey"`
	EndTime             string                  `json:"EndTime" xml:"EndTime"`
	VSwitchId           string                  `json:"VSwitchId" xml:"VSwitchId"`
	NodeType            string                  `json:"NodeType" xml:"NodeType"`
	Connections         int64                   `json:"Connections" xml:"Connections"`
	ResourceGroupId     string                  `json:"ResourceGroupId" xml:"ResourceGroupId"`
	ZoneId              string                  `json:"ZoneId" xml:"ZoneId"`
	InstanceStatus      string                  `json:"InstanceStatus" xml:"InstanceStatus"`
	ProxyCount          int                     `json:"ProxyCount" xml:"ProxyCount"`
	CloudType           string                  `json:"CloudType" xml:"CloudType"`
	Tags                TagsInDescribeInstances `json:"Tags" xml:"Tags"`
}