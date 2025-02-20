// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20190904

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AcListsData struct {

	// 规则id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 访问源
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`

	// 访问目的
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetIp *string `json:"TargetIp,omitempty" name:"TargetIp"`

	// 协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *string `json:"Port,omitempty" name:"Port"`

	// 策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	Strategy *uint64 `json:"Strategy,omitempty" name:"Strategy"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail *string `json:"Detail,omitempty" name:"Detail"`

	// 命中次数
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 执行顺序
	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`

	// 告警规则id
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogId *string `json:"LogId,omitempty" name:"LogId"`
}

type AssetZone struct {

	// 地域
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 地域英文
	ZoneEng *string `json:"ZoneEng,omitempty" name:"ZoneEng"`
}

type AssociatedInstanceInfo struct {

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例类型，3是cvm实例,4是clb实例,5是eni实例,6是云数据库
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 私有网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 公网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`

	// 内网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 关联安全组数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroupCount *uint64 `json:"SecurityGroupCount,omitempty" name:"SecurityGroupCount"`
}

type CfwNatDnatRule struct {

	// 网络协议，可选值：TCP、UDP。
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// 弹性IP。
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" name:"PublicIpAddress"`

	// 公网端口。
	PublicPort *int64 `json:"PublicPort,omitempty" name:"PublicPort"`

	// 内网地址。
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`

	// 内网端口。
	PrivatePort *int64 `json:"PrivatePort,omitempty" name:"PrivatePort"`

	// NAT防火墙转发规则描述。
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateAcRulesRequest struct {
	*tchttp.BaseRequest

	// 创建规则数据
	Data []*RuleInfoData `json:"Data,omitempty" name:"Data"`

	// 0：添加（默认），1：插入
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 边id
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// 访问控制规则状态
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 0：添加，1：覆盖
	Overwrite *uint64 `json:"Overwrite,omitempty" name:"Overwrite"`

	// NAT实例ID, 参数Area存在的时候这个必传
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// portScan: 来自于端口扫描, patchImport: 来自于批量导入
	From *string `json:"From,omitempty" name:"From"`

	// NAT地域
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *CreateAcRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAcRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Data")
	delete(f, "Type")
	delete(f, "EdgeId")
	delete(f, "Enable")
	delete(f, "Overwrite")
	delete(f, "InstanceId")
	delete(f, "From")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAcRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAcRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态值，0:操作成功
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 返回多余的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Info *string `json:"Info,omitempty" name:"Info"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAcRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAcRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateChooseVpcsRequest struct {
	*tchttp.BaseRequest

	// vpc列表
	VpcList []*string `json:"VpcList,omitempty" name:"VpcList"`

	// zone列表
	AllZoneList []*VpcZoneData `json:"AllZoneList,omitempty" name:"AllZoneList"`
}

func (r *CreateChooseVpcsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateChooseVpcsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcList")
	delete(f, "AllZoneList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateChooseVpcsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateChooseVpcsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateChooseVpcsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateChooseVpcsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNatFwInstanceRequest struct {
	*tchttp.BaseRequest

	// 防火墙实例名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 带宽
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 模式 1：接入模式；0：新增模式
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// 新增模式传递参数，其中NewModeItems和NatgwList至少传递一种。
	NewModeItems *NewModeItems `json:"NewModeItems,omitempty" name:"NewModeItems"`

	// 接入模式接入的nat网关列表，其中NewModeItems和NatgwList至少传递一种。
	NatGwList []*string `json:"NatGwList,omitempty" name:"NatGwList"`

	// 主可用区，为空则选择默认可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 备可用区，为空则选择默认可用区
	ZoneBak *string `json:"ZoneBak,omitempty" name:"ZoneBak"`

	// 异地灾备 1：使用异地灾备；0：不使用异地灾备
	CrossAZone *int64 `json:"CrossAZone,omitempty" name:"CrossAZone"`
}

func (r *CreateNatFwInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNatFwInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Width")
	delete(f, "Mode")
	delete(f, "NewModeItems")
	delete(f, "NatGwList")
	delete(f, "Zone")
	delete(f, "ZoneBak")
	delete(f, "CrossAZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNatFwInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateNatFwInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNatFwInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNatFwInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupApiRulesRequest struct {
	*tchttp.BaseRequest

	// 创建规则数据
	Data []*SecurityGroupApiRuleData `json:"Data,omitempty" name:"Data"`

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// 插入类型，0：后插，1：前插，2：中插
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 腾讯云地域的英文简写
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *CreateSecurityGroupApiRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupApiRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Data")
	delete(f, "Direction")
	delete(f, "Type")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityGroupApiRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupApiRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态值，0:添加成功，非0：添加失败
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSecurityGroupApiRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupApiRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupRulesRequest struct {
	*tchttp.BaseRequest

	// 添加的企业安全组规则数据
	Data []*SecurityGroupListData `json:"Data,omitempty" name:"Data"`

	// 方向，0：出站，1：入站，默认1
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// 0：后插，1：前插，2：中插，默认0
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 添加后是否启用规则，0：不启用，1：启用，默认1
	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
}

func (r *CreateSecurityGroupRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Data")
	delete(f, "Direction")
	delete(f, "Type")
	delete(f, "Enable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityGroupRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态值，0：添加成功，非0：添加失败
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSecurityGroupRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAcRuleRequest struct {
	*tchttp.BaseRequest

	// 删除规则对应的id值, 对应获取规则列表接口的Id 值
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// EdgeId值两个vpc间的边id
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// NAT地域， 如ap-shanghai/ap-guangzhou/ap-chongqing等
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DeleteAcRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAcRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Direction")
	delete(f, "EdgeId")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAcRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAcRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态值 0: 删除成功, !0: 删除失败
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 返回多余的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Info *string `json:"Info,omitempty" name:"Info"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAcRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAllAccessControlRuleRequest struct {
	*tchttp.BaseRequest

	// 方向，0：出站，1：入站  默认值是 0
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// VPC间防火墙开关ID  全部删除 EdgeId和Area只填写一个，不填写则不删除vpc间防火墙开关 ，默认值为‘’
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// nat地域 全部删除 EdgeId和Area只填写一个，不填写则不删除nat防火墙开关 默认值为‘’
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DeleteAllAccessControlRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAllAccessControlRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Direction")
	delete(f, "EdgeId")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAllAccessControlRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAllAccessControlRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态值 0: 修改成功, !0: 修改失败
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 删除了几条访问控制规则
	// 注意：此字段可能返回 null，表示取不到有效值。
		Info *int64 `json:"Info,omitempty" name:"Info"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAllAccessControlRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAllAccessControlRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNatFwInstanceRequest struct {
	*tchttp.BaseRequest

	// 防火墙实例id
	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`
}

func (r *DeleteNatFwInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNatFwInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CfwInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNatFwInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNatFwInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNatFwInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNatFwInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteResourceGroupRequest struct {
	*tchttp.BaseRequest

	// 组id
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteResourceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteResourceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteResourceGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteResourceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupAllRuleRequest struct {
	*tchttp.BaseRequest

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// 腾讯云地域的英文简写
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DeleteSecurityGroupAllRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityGroupAllRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Direction")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecurityGroupAllRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupAllRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 0: 操作成功，非0：操作失败
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 返回数据的json字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
		Info *int64 `json:"Info,omitempty" name:"Info"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSecurityGroupAllRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityGroupAllRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupRuleRequest struct {
	*tchttp.BaseRequest

	// 所需要删除规则的ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 腾讯云地域的英文简写
	Area *string `json:"Area,omitempty" name:"Area"`

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// 是否删除反向规则，0：否，1：是
	IsDelReverse *uint64 `json:"IsDelReverse,omitempty" name:"IsDelReverse"`
}

func (r *DeleteSecurityGroupRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityGroupRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Area")
	delete(f, "Direction")
	delete(f, "IsDelReverse")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecurityGroupRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态值，0：成功，非0：失败
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 返回多余的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Info *string `json:"Info,omitempty" name:"Info"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSecurityGroupRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityGroupRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcInstanceRequest struct {
	*tchttp.BaseRequest
}

func (r *DeleteVpcInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVpcInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVpcInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAcListsRequest struct {
	*tchttp.BaseRequest

	// 协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 策略
	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`

	// 搜索值
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`

	// 每页条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 出站还是入站，0：入站，1：出站
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// EdgeId值
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// 规则是否开启，'0': 未开启，'1': 开启, 默认为'0'
	Status *string `json:"Status,omitempty" name:"Status"`

	// 地域
	Area *string `json:"Area,omitempty" name:"Area"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeAcListsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAcListsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Protocol")
	delete(f, "Strategy")
	delete(f, "SearchValue")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Direction")
	delete(f, "EdgeId")
	delete(f, "Status")
	delete(f, "Area")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAcListsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAcListsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总条数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 访问控制列表数据
		Data []*AcListsData `json:"Data,omitempty" name:"Data"`

		// 不算筛选条数的总条数
		AllTotal *uint64 `json:"AllTotal,omitempty" name:"AllTotal"`

		// 访问控制规则全部启用/全部停用
	// 注意：此字段可能返回 null，表示取不到有效值。
		Enable *uint64 `json:"Enable,omitempty" name:"Enable"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAcListsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAcListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssociatedInstanceListRequest struct {
	*tchttp.BaseRequest

	// 列表偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页记录条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 地域代码（例：ap-guangzhou）,支持腾讯云全地域
	Area *string `json:"Area,omitempty" name:"Area"`

	// 额外检索条件（JSON字符串）
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式（asc:升序,desc:降序）
	Order *string `json:"Order,omitempty" name:"Order"`

	// 安全组ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 实例类型,'3'是cvm实例,'4'是clb实例,'5'是eni实例,'6'是云数据库
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeAssociatedInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssociatedInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Area")
	delete(f, "SearchValue")
	delete(f, "By")
	delete(f, "Order")
	delete(f, "SecurityGroupId")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssociatedInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssociatedInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*AssociatedInstanceInfo `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssociatedInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssociatedInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockByIpTimesListRequest struct {
	*tchttp.BaseRequest

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// ip查询条件
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 地域
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 方向
	Direction *string `json:"Direction,omitempty" name:"Direction"`

	// 来源
	Source *string `json:"Source,omitempty" name:"Source"`
}

func (r *DescribeBlockByIpTimesListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlockByIpTimesListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Ip")
	delete(f, "Zone")
	delete(f, "Direction")
	delete(f, "Source")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBlockByIpTimesListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockByIpTimesListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data []*IpStatic `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBlockByIpTimesListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlockByIpTimesListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockStaticListRequest struct {
	*tchttp.BaseRequest

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 类型
	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`

	// top数
	Top *int64 `json:"Top,omitempty" name:"Top"`

	// 查询条件
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
}

func (r *DescribeBlockStaticListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlockStaticListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "QueryType")
	delete(f, "Top")
	delete(f, "SearchValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBlockStaticListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockStaticListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 无
		Data []*StaticInfo `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBlockStaticListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlockStaticListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCfwEipsRequest struct {
	*tchttp.BaseRequest

	// 0：cfw新增模式，1：cfw接入模式
	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`

	// ALL：查询所有弹性公网ip; nat-xxxxx：接入模式场景指定网关的弹性公网ip
	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`

	// 防火墙实例id
	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`
}

func (r *DescribeCfwEipsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfwEipsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Mode")
	delete(f, "NatGatewayId")
	delete(f, "CfwInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCfwEipsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCfwEipsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回值信息
		NatFwEipList []*NatFwEipsInfo `json:"NatFwEipList,omitempty" name:"NatFwEipList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCfwEipsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfwEipsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGuideScanInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeGuideScanInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGuideScanInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGuideScanInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGuideScanInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 扫描信息
		Data *ScanInfo `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGuideScanInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGuideScanInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatFwInfoCountRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeNatFwInfoCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwInfoCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatFwInfoCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatFwInfoCountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回参数
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`

		// 当前租户的nat实例个数
	// 注意：此字段可能返回 null，表示取不到有效值。
		NatFwInsCount *int64 `json:"NatFwInsCount,omitempty" name:"NatFwInsCount"`

		// 当前租户接入子网个数
	// 注意：此字段可能返回 null，表示取不到有效值。
		SubnetCount *int64 `json:"SubnetCount,omitempty" name:"SubnetCount"`

		// 打开开关个数
	// 注意：此字段可能返回 null，表示取不到有效值。
		OpenSwitchCount *int64 `json:"OpenSwitchCount,omitempty" name:"OpenSwitchCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatFwInfoCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwInfoCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatFwInstanceRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeNatFwInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatFwInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatFwInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例数组
		NatinsLst []*NatFwInstance `json:"NatinsLst,omitempty" name:"NatinsLst"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatFwInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatFwInstanceWithRegionRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeNatFwInstanceWithRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwInstanceWithRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatFwInstanceWithRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatFwInstanceWithRegionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		NatinsLst []*NatFwInstance `json:"NatinsLst,omitempty" name:"NatinsLst"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatFwInstanceWithRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwInstanceWithRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatFwInstancesInfoRequest struct {
	*tchttp.BaseRequest

	// 获取实例列表过滤字段
	Filter []*NatFwFilter `json:"Filter,omitempty" name:"Filter"`

	// 第几页
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页长度
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeNatFwInstancesInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwInstancesInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatFwInstancesInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatFwInstancesInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例卡片信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		NatinsLst []*NatInstanceInfo `json:"NatinsLst,omitempty" name:"NatinsLst"`

		// 总共多少条数
	// 注意：此字段可能返回 null，表示取不到有效值。
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatFwInstancesInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwInstancesInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatFwVpcDnsLstRequest struct {
	*tchttp.BaseRequest

	// natfw 防火墙实例id
	NatFwInsId *string `json:"NatFwInsId,omitempty" name:"NatFwInsId"`

	// natfw 过滤，以','分隔
	NatInsIdFilter *string `json:"NatInsIdFilter,omitempty" name:"NatInsIdFilter"`

	// 分页页数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页最多个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeNatFwVpcDnsLstRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwVpcDnsLstRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatFwInsId")
	delete(f, "NatInsIdFilter")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatFwVpcDnsLstRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatFwVpcDnsLstResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// nat防火墙vpc dns 信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		VpcDnsSwitchLst []*VpcDnsInfo `json:"VpcDnsSwitchLst,omitempty" name:"VpcDnsSwitchLst"`

		// 返回参数
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`

		// 开关总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatFwVpcDnsLstResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwVpcDnsLstResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatRuleOverviewRequest struct {
	*tchttp.BaseRequest

	// 方向，0：出站，1：入站 默认值：0
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// NAT地域  这个是必填项，填入相关的英文，'ap-beijing-fsi': '北京金融',
	//         'ap-beijing': '北京',
	//         'ap-changsha-ec': '长沙EC',
	//         'ap-chengdu': '成都',
	//         'ap-chongqing': '重庆',
	//         'ap-fuzhou-ec': '福州EC',
	//         'ap-guangzhou-open': '广州Open',
	//         'ap-guangzhou': '广州',
	//         'ap-hangzhou-ec': '杭州EC',
	//         'ap-jinan-ec': '济南EC',
	//         'ap-nanjing': '南京',
	//         'ap-shanghai-fsi': '上海金融',
	//         'ap-shanghai': '上海',
	//         'ap-shenzhen-fsi': '深圳金融',
	//         'ap-shenzhen': '深圳',
	//         'ap-wuhan-ec': '武汉EC'
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeNatRuleOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatRuleOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Direction")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatRuleOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatRuleOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例ID
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 实例名称
		InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

		// 弹性IP列表
		EipList []*string `json:"EipList,omitempty" name:"EipList"`

		// 端口转发规则数量
		DnatNum *int64 `json:"DnatNum,omitempty" name:"DnatNum"`

		// 访问控制规则总数
		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 访问控制规则剩余配额
		RemainNum *int64 `json:"RemainNum,omitempty" name:"RemainNum"`

		// 阻断规则条数
		BlockNum *int64 `json:"BlockNum,omitempty" name:"BlockNum"`

		// 启用规则条数
		EnableNum *int64 `json:"EnableNum,omitempty" name:"EnableNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatRuleOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatRuleOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceGroupRequest struct {
	*tchttp.BaseRequest

	// 查询类型 网络结构 vpc，业务识别- resource ，资源标签-tag
	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`

	// 资产组id  全部传0
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeResourceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueryType")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回树形结构
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleOverviewRequest struct {
	*tchttp.BaseRequest

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
}

func (r *DescribeRuleOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Direction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuleOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		AllTotal *uint64 `json:"AllTotal,omitempty" name:"AllTotal"`

		// 阻断策略规则数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		StrategyNum *uint64 `json:"StrategyNum,omitempty" name:"StrategyNum"`

		// 启用规则数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		StartRuleNum *uint64 `json:"StartRuleNum,omitempty" name:"StartRuleNum"`

		// 停用规则数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		StopRuleNum *uint64 `json:"StopRuleNum,omitempty" name:"StopRuleNum"`

		// 剩余配额
	// 注意：此字段可能返回 null，表示取不到有效值。
		RemainingNum *int64 `json:"RemainingNum,omitempty" name:"RemainingNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRuleOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupListRequest struct {
	*tchttp.BaseRequest

	// 0: 出站规则，1：入站规则
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// 地域代码（例: ap-guangzhou),支持腾讯云全部地域
	Area *string `json:"Area,omitempty" name:"Area"`

	// 搜索值
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`

	// 每页条数，默认为10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移值，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 状态，'': 全部，'0'：筛选停用规则，'1'：筛选启用规则
	Status *string `json:"Status,omitempty" name:"Status"`

	// 0: 不过滤，1：过滤掉正常规则，保留下发异常规则
	Filter *uint64 `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeSecurityGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Direction")
	delete(f, "Area")
	delete(f, "SearchValue")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Status")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 列表当前规则总条数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 安全组规则列表数据
		Data []*SecurityGroupListData `json:"Data,omitempty" name:"Data"`

		// 不算筛选条数的总条数
		AllTotal *uint64 `json:"AllTotal,omitempty" name:"AllTotal"`

		// 访问控制规则全部启用/全部停用
	// 注意：此字段可能返回 null，表示取不到有效值。
		Enable *uint64 `json:"Enable,omitempty" name:"Enable"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSourceAssetRequest struct {
	*tchttp.BaseRequest

	// 模糊查询
	FuzzySearch *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`

	// 资产类型 1公网 2内网
	InsType *string `json:"InsType,omitempty" name:"InsType"`

	// 是否未分组 1是
	ChooseType *string `json:"ChooseType,omitempty" name:"ChooseType"`

	// 地域
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeSourceAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSourceAssetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FuzzySearch")
	delete(f, "InsType")
	delete(f, "ChooseType")
	delete(f, "Zone")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSourceAssetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSourceAssetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 地域集合
		ZoneList []*AssetZone `json:"ZoneList,omitempty" name:"ZoneList"`

		// 数据
		Data []*InstanceInfo `json:"Data,omitempty" name:"Data"`

		// 10
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSourceAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSourceAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSwitchListsRequest struct {
	*tchttp.BaseRequest

	// 防火墙状态  0: 关闭，1：开启
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 资产类型 CVM/NAT/VPN/CLB/其它
	Type *string `json:"Type,omitempty" name:"Type"`

	// 地域 上海/重庆/广州，等等
	Area *string `json:"Area,omitempty" name:"Area"`

	// 搜索值  例子："{"common":"106.54.189.45"}"
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`

	// 条数  默认值:10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移值 默认值: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序，desc：降序，asc：升序
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段 PortTimes(风险端口数)
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeSwitchListsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSwitchListsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "Type")
	delete(f, "Area")
	delete(f, "SearchValue")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSwitchListsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSwitchListsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总条数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 列表数据
		Data []*SwitchListsData `json:"Data,omitempty" name:"Data"`

		// 区域列表
		AreaLists []*string `json:"AreaLists,omitempty" name:"AreaLists"`

		// 打开个数
	// 注意：此字段可能返回 null，表示取不到有效值。
		OnNum *uint64 `json:"OnNum,omitempty" name:"OnNum"`

		// 关闭个数
	// 注意：此字段可能返回 null，表示取不到有效值。
		OffNum *uint64 `json:"OffNum,omitempty" name:"OffNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSwitchListsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSwitchListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSyncAssetStatusRequest struct {
	*tchttp.BaseRequest

	// 0: 互联网防火墙开关，1：vpc 防火墙开关
	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeSyncAssetStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSyncAssetStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSyncAssetStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSyncAssetStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 1-更新中 2-更新完成 3、4-更新失败
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSyncAssetStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSyncAssetStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTLogInfoRequest struct {
	*tchttp.BaseRequest

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 类型 1 告警 2阻断
	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`

	// 查询条件
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
}

func (r *DescribeTLogInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTLogInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "QueryType")
	delete(f, "SearchValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTLogInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTLogInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 无
		Data *TLogInfo `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTLogInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTLogInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTLogIpListRequest struct {
	*tchttp.BaseRequest

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 类型 1 告警 2阻断
	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`

	// top数
	Top *int64 `json:"Top,omitempty" name:"Top"`

	// 查询条件
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
}

func (r *DescribeTLogIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTLogIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "QueryType")
	delete(f, "Top")
	delete(f, "SearchValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTLogIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTLogIpListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据集合
		Data []*StaticInfo `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTLogIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTLogIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTableStatusRequest struct {
	*tchttp.BaseRequest

	// EdgeId值两个vpc间的边id vpc填Edgeid，不要填Area；
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// 状态值，0：检查表的状态 确实只有一个默认值
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Nat所在地域 NAT填Area，不要填Edgeid；
	Area *string `json:"Area,omitempty" name:"Area"`

	// 方向，0：出站，1：入站 默认值为 0
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
}

func (r *DescribeTableStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeId")
	delete(f, "Status")
	delete(f, "Area")
	delete(f, "Direction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTableStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 0：正常，其它：不正常
	// 注意：此字段可能返回 null，表示取不到有效值。
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTableStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnHandleEventTabListRequest struct {
	*tchttp.BaseRequest

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询示例ID
	AssetID *string `json:"AssetID,omitempty" name:"AssetID"`
}

func (r *DescribeUnHandleEventTabListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnHandleEventTabListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AssetID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUnHandleEventTabListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnHandleEventTabListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 租户伪攻击链未处置事件
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *UnHandleEvent `json:"Data,omitempty" name:"Data"`

		// 错误码，0成功 非0错误
		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`

		// 返回信息 success成功
		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUnHandleEventTabListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnHandleEventTabListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcRuleOverviewRequest struct {
	*tchttp.BaseRequest

	// EdgeId值两个vpc间的边id  不是必填项可以为空，就是所有vpc间的访问控制规则
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
}

func (r *DescribeVpcRuleOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcRuleOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcRuleOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcRuleOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 阻断策略规则数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		StrategyNum *uint64 `json:"StrategyNum,omitempty" name:"StrategyNum"`

		// 启用规则数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		StartRuleNum *uint64 `json:"StartRuleNum,omitempty" name:"StartRuleNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcRuleOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcRuleOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DnsVpcSwitch struct {

	// vpc id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 0：设置为关闭 1:设置为打开
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type ExpandCfwVerticalRequest struct {
	*tchttp.BaseRequest

	// nat：nat防火墙，ew：东西向防火墙
	FwType *string `json:"FwType,omitempty" name:"FwType"`

	// 带宽值
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 防火墙实例id
	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`
}

func (r *ExpandCfwVerticalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExpandCfwVerticalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FwType")
	delete(f, "Width")
	delete(f, "CfwInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExpandCfwVerticalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ExpandCfwVerticalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExpandCfwVerticalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExpandCfwVerticalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceInfo struct {

	// appid信息
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// vpcid信息
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// vpc名称
	VPCName *string `json:"VPCName,omitempty" name:"VPCName"`

	// 子网id
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 资产id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 资产名
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 资产类型
	//  3是cvm实例,4是clb实例,5是eni实例,6是mysql,7是redis,8是NAT,9是VPN,10是ES,11是MARIADB,12是KAFKA 13 NATFW
	InsType *int64 `json:"InsType,omitempty" name:"InsType"`

	// 公网ip
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`

	// 内网ip
	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`

	// 端口数
	PortNum *string `json:"PortNum,omitempty" name:"PortNum"`

	// 漏洞数
	LeakNum *string `json:"LeakNum,omitempty" name:"LeakNum"`

	// 1，公网 2内网
	InsSource *string `json:"InsSource,omitempty" name:"InsSource"`

	// [a,b]
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourcePath []*string `json:"ResourcePath,omitempty" name:"ResourcePath"`
}

type IocListData struct {

	// 待处置IP地址，IP/Domain字段二选一
	IP *string `json:"IP,omitempty" name:"IP"`

	// 只能为0或者1   0代表出站 1代表入站
	Direction *int64 `json:"Direction,omitempty" name:"Direction"`

	// 待处置域名，IP/Domain字段二选一
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type IpStatic struct {

	// 值
	Num *int64 `json:"Num,omitempty" name:"Num"`

	// 折线图横坐标时间
	StatTime *string `json:"StatTime,omitempty" name:"StatTime"`
}

type ModifyAcRuleRequest struct {
	*tchttp.BaseRequest

	// 规则数组
	Data []*RuleInfoData `json:"Data,omitempty" name:"Data"`

	// EdgeId值
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// 访问规则状态
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// NAT地域
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *ModifyAcRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAcRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Data")
	delete(f, "EdgeId")
	delete(f, "Enable")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAcRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAcRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态值，0:操作成功，非0：操作失败
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 返回多余的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Info *string `json:"Info,omitempty" name:"Info"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAcRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAllPublicIPSwitchStatusRequest struct {
	*tchttp.BaseRequest

	// 状态，0：关闭，1：开启
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 选中的防火墙开关Id
	FireWallPublicIPs []*string `json:"FireWallPublicIPs,omitempty" name:"FireWallPublicIPs"`
}

func (r *ModifyAllPublicIPSwitchStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAllPublicIPSwitchStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "FireWallPublicIPs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAllPublicIPSwitchStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAllPublicIPSwitchStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 接口返回信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`

		// 接口返回错误码，0请求成功  非0失败
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAllPublicIPSwitchStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAllPublicIPSwitchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAllRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 状态，0：全部停用，1：全部启用
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// Edge ID值
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// NAT地域
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *ModifyAllRuleStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAllRuleStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "Direction")
	delete(f, "EdgeId")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAllRuleStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAllRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 0: 修改成功, 其他: 修改失败
	// 注意：此字段可能返回 null，表示取不到有效值。
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAllRuleStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAllRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAllSwitchStatusRequest struct {
	*tchttp.BaseRequest

	// 状态，0：关闭，1：开启
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 0: 互联网边界防火墙开关，1：vpc防火墙开关
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 选中的防火墙开关Id
	Ids []*string `json:"Ids,omitempty" name:"Ids"`

	// NAT开关切换类型，1,单个子网，2，同开同关，3，全部
	ChangeType *int64 `json:"ChangeType,omitempty" name:"ChangeType"`

	// NAT实例所在地域
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *ModifyAllSwitchStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAllSwitchStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "Type")
	delete(f, "Ids")
	delete(f, "ChangeType")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAllSwitchStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAllSwitchStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 修改成功与否的状态值 0：修改成功，非 0：修改失败
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAllSwitchStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAllSwitchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAllVPCSwitchStatusRequest struct {
	*tchttp.BaseRequest

	// 状态，0：关闭，1：开启
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 选中的防火墙开关Id
	FireWallVpcIds []*string `json:"FireWallVpcIds,omitempty" name:"FireWallVpcIds"`
}

func (r *ModifyAllVPCSwitchStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAllVPCSwitchStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "FireWallVpcIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAllVPCSwitchStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAllVPCSwitchStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAllVPCSwitchStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAllVPCSwitchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAssetScanRequest struct {
	*tchttp.BaseRequest

	// 扫描范围：1端口, 2端口+漏扫
	ScanRange *int64 `json:"ScanRange,omitempty" name:"ScanRange"`

	// 扫描深度：'heavy', 'medium', 'light'
	ScanDeep *string `json:"ScanDeep,omitempty" name:"ScanDeep"`

	// 扫描类型：1立即扫描 2 周期任务
	RangeType *int64 `json:"RangeType,omitempty" name:"RangeType"`

	// RangeType为2 是必须添加，定时任务时间
	ScanPeriod *string `json:"ScanPeriod,omitempty" name:"ScanPeriod"`

	// 立即扫描这个字段传过滤的扫描集合
	ScanFilterIp []*string `json:"ScanFilterIp,omitempty" name:"ScanFilterIp"`

	// 1全量2单个
	ScanType *int64 `json:"ScanType,omitempty" name:"ScanType"`
}

func (r *ModifyAssetScanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAssetScanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScanRange")
	delete(f, "ScanDeep")
	delete(f, "RangeType")
	delete(f, "ScanPeriod")
	delete(f, "ScanFilterIp")
	delete(f, "ScanType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAssetScanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAssetScanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 接口返回信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`

		// 接口返回错误码，0请求成功  非0失败
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`

		// 状态值 0：成功，1 执行扫描中,其他：失败
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAssetScanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAssetScanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBlockIgnoreListRequest struct {
	*tchttp.BaseRequest

	// 1拦截列表 2 忽略列表
	RuleType *int64 `json:"RuleType,omitempty" name:"RuleType"`

	// IP、Domain二选一，不能同时为空
	IOC []*IocListData `json:"IOC,omitempty" name:"IOC"`

	// 可选值：delete（删除）、edit（编辑）、add（添加）  其他值无效
	IocAction *string `json:"IocAction,omitempty" name:"IocAction"`

	// 时间格式：yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 时间格式：yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *ModifyBlockIgnoreListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBlockIgnoreListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleType")
	delete(f, "IOC")
	delete(f, "IocAction")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBlockIgnoreListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBlockIgnoreListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 接口返回信息
		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`

		// 接口返回错误码，0请求成功  非0失败
		ReturnCode *uint64 `json:"ReturnCode,omitempty" name:"ReturnCode"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBlockIgnoreListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBlockIgnoreListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBlockTopRequest struct {
	*tchttp.BaseRequest

	// 记录id
	UniqueId *string `json:"UniqueId,omitempty" name:"UniqueId"`

	// 操作类型 1 置顶 0取消
	OpeType *string `json:"OpeType,omitempty" name:"OpeType"`
}

func (r *ModifyBlockTopRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBlockTopRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UniqueId")
	delete(f, "OpeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBlockTopRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBlockTopResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBlockTopResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBlockTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyItemSwitchStatusRequest struct {
	*tchttp.BaseRequest

	// id值
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 状态值，0: 关闭 ,1:开启
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 0: 互联网边界边界防火墙开关，1：vpc防火墙开关
	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *ModifyItemSwitchStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyItemSwitchStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Status")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyItemSwitchStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyItemSwitchStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 修改成功与否状态值 0：修改成功，非 0：修改失败
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyItemSwitchStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyItemSwitchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatFwReSelectRequest struct {
	*tchttp.BaseRequest

	// 模式 1：接入模式；0：新增模式
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// 防火墙实例id
	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`

	// 接入模式重新接入的nat网关列表，其中NatGwList和VpcList只能传递一个。
	NatGwList []*string `json:"NatGwList,omitempty" name:"NatGwList"`

	// 新增模式重新接入的vpc列表，其中NatGwList和NatgwList只能传递一个。
	VpcList []*string `json:"VpcList,omitempty" name:"VpcList"`
}

func (r *ModifyNatFwReSelectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatFwReSelectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Mode")
	delete(f, "CfwInstance")
	delete(f, "NatGwList")
	delete(f, "VpcList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNatFwReSelectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatFwReSelectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNatFwReSelectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatFwReSelectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatFwSwitchRequest struct {
	*tchttp.BaseRequest

	// 开关，0：关闭，1：开启
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 防火墙实例id列表，其中CfwInsIdList，SubnetIdList和RouteTableIdList只能传递一种。
	CfwInsIdList []*string `json:"CfwInsIdList,omitempty" name:"CfwInsIdList"`

	// 子网id列表，其中CfwInsIdList，SubnetIdList和RouteTableIdList只能传递一种。
	SubnetIdList []*string `json:"SubnetIdList,omitempty" name:"SubnetIdList"`

	// 路由表id列表，其中CfwInsIdList，SubnetIdList和RouteTableIdList只能传递一种。
	RouteTableIdList []*string `json:"RouteTableIdList,omitempty" name:"RouteTableIdList"`
}

func (r *ModifyNatFwSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatFwSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Enable")
	delete(f, "CfwInsIdList")
	delete(f, "SubnetIdList")
	delete(f, "RouteTableIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNatFwSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatFwSwitchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNatFwSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatFwSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatFwVpcDnsSwitchRequest struct {
	*tchttp.BaseRequest

	// nat 防火墙 id
	NatFwInsId *string `json:"NatFwInsId,omitempty" name:"NatFwInsId"`

	// DNS 开关切换列表
	DnsVpcSwitchLst []*DnsVpcSwitch `json:"DnsVpcSwitchLst,omitempty" name:"DnsVpcSwitchLst"`
}

func (r *ModifyNatFwVpcDnsSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatFwVpcDnsSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatFwInsId")
	delete(f, "DnsVpcSwitchLst")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNatFwVpcDnsSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatFwVpcDnsSwitchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 修改成功
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNatFwVpcDnsSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatFwVpcDnsSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPublicIPSwitchStatusRequest struct {
	*tchttp.BaseRequest

	// 公网IP
	FireWallPublicIP *string `json:"FireWallPublicIP,omitempty" name:"FireWallPublicIP"`

	// 状态值，0: 关闭 ,1:开启
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyPublicIPSwitchStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPublicIPSwitchStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FireWallPublicIP")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPublicIPSwitchStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPublicIPSwitchStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 接口返回信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`

		// 接口返回错误码，0请求成功  非0失败
		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPublicIPSwitchStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPublicIPSwitchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResourceGroupRequest struct {
	*tchttp.BaseRequest

	// 组id
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 上级组id
	ParentId *string `json:"ParentId,omitempty" name:"ParentId"`
}

func (r *ModifyResourceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "GroupName")
	delete(f, "ParentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResourceGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyResourceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupAllRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 列表规则状态，0：全部停用，1：全部启用
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// Edge ID值
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// NAT地域, 腾讯云地域的英文简写
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *ModifySecurityGroupAllRuleStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupAllRuleStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "Direction")
	delete(f, "EdgeId")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityGroupAllRuleStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupAllRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 0: 修改成功, 其他: 修改失败
	// 注意：此字段可能返回 null，表示取不到有效值。
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySecurityGroupAllRuleStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupAllRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupItemRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 方向，0：出站，1：入站，默认1
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// 是否开关开启，0：未开启，1：开启
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 更改的企业安全组规则的执行顺序
	RuleSequence *uint64 `json:"RuleSequence,omitempty" name:"RuleSequence"`
}

func (r *ModifySecurityGroupItemRuleStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupItemRuleStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Direction")
	delete(f, "Status")
	delete(f, "RuleSequence")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityGroupItemRuleStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupItemRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态值，0：修改成功，非0：修改失败
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySecurityGroupItemRuleStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupItemRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupRuleRequest struct {
	*tchttp.BaseRequest

	// 方向，0：出站，1：入站，默认1
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// 编辑后是否启用规则，0：不启用，1：启用，默认1
	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`

	// 编辑的企业安全组规则数据
	Data []*SecurityGroupListData `json:"Data,omitempty" name:"Data"`

	// 编辑的企业安全组规则的原始执行顺序
	SgRuleOriginSequence *uint64 `json:"SgRuleOriginSequence,omitempty" name:"SgRuleOriginSequence"`
}

func (r *ModifySecurityGroupRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Direction")
	delete(f, "Enable")
	delete(f, "Data")
	delete(f, "SgRuleOriginSequence")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityGroupRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态值，0：编辑成功，非0：编辑失败
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// 编辑后新生成规则的Id
		NewRuleId *uint64 `json:"NewRuleId,omitempty" name:"NewRuleId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySecurityGroupRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupSequenceRulesRequest struct {
	*tchttp.BaseRequest

	// 方向，0：出站，1：入站，默认1
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// 企业安全组规则快速排序数据
	Data []*SecurityGroupOrderIndexData `json:"Data,omitempty" name:"Data"`
}

func (r *ModifySecurityGroupSequenceRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupSequenceRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Direction")
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityGroupSequenceRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupSequenceRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态值，0：修改成功，非0：修改失败
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySecurityGroupSequenceRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupSequenceRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySequenceRulesRequest struct {
	*tchttp.BaseRequest

	// 边Id值
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// 修改数据
	Data []*SequenceData `json:"Data,omitempty" name:"Data"`

	// NAT地域
	Area *string `json:"Area,omitempty" name:"Area"`

	// 方向，0：出向，1：入向
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
}

func (r *ModifySequenceRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySequenceRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeId")
	delete(f, "Data")
	delete(f, "Area")
	delete(f, "Direction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySequenceRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySequenceRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 0: 修改成功, 非0: 修改失败
	// 注意：此字段可能返回 null，表示取不到有效值。
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySequenceRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySequenceRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTableStatusRequest struct {
	*tchttp.BaseRequest

	// EdgeId值两个vpc间的边id
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// 状态值，1：锁表，2：解锁表
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Nat所在地域
	Area *string `json:"Area,omitempty" name:"Area"`

	// 0： 出向，1：入向
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
}

func (r *ModifyTableStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTableStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeId")
	delete(f, "Status")
	delete(f, "Area")
	delete(f, "Direction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTableStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTableStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 0：正常，-1：不正常
	// 注意：此字段可能返回 null，表示取不到有效值。
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTableStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTableStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVPCSwitchStatusRequest struct {
	*tchttp.BaseRequest

	// 公网IP
	FirewallVpcId *string `json:"FirewallVpcId,omitempty" name:"FirewallVpcId"`

	// 状态值，0: 关闭 ,1:开启
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyVPCSwitchStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVPCSwitchStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FirewallVpcId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVPCSwitchStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVPCSwitchStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 接口返回信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`

		// 接口返回错误码，0请求成功  非0失败
		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVPCSwitchStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVPCSwitchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NatFwEipsInfo struct {

	// 弹性公网ip
	Eip *string `json:"Eip,omitempty" name:"Eip"`

	// 所属的Nat网关Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`

	// Nat网关名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatGatewayName *string `json:"NatGatewayName,omitempty" name:"NatGatewayName"`
}

type NatFwFilter struct {

	// 过滤的类型，例如实例id
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`

	// 过滤的内容，以',' 分隔
	FilterContent *string `json:"FilterContent,omitempty" name:"FilterContent"`
}

type NatFwInstance struct {

	// nat实例id
	NatinsId *string `json:"NatinsId,omitempty" name:"NatinsId"`

	// nat实例名称
	NatinsName *string `json:"NatinsName,omitempty" name:"NatinsName"`

	// 实例所在地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 0:新增模式，1:接入模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	FwMode *int64 `json:"FwMode,omitempty" name:"FwMode"`

	// 0:正常状态， 1: 正在创建
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type NatInstanceInfo struct {

	// nat实例id
	NatinsId *string `json:"NatinsId,omitempty" name:"NatinsId"`

	// nat实例名称
	NatinsName *string `json:"NatinsName,omitempty" name:"NatinsName"`

	// 实例所在地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 0: 新增模式，1:接入模式
	FwMode *int64 `json:"FwMode,omitempty" name:"FwMode"`

	// 实例带宽大小 Mbps
	BandWidth *int64 `json:"BandWidth,omitempty" name:"BandWidth"`

	// 入向带宽峰值 bps
	InFlowMax *int64 `json:"InFlowMax,omitempty" name:"InFlowMax"`

	// 出向带宽峰值 bps
	OutFlowMax *uint64 `json:"OutFlowMax,omitempty" name:"OutFlowMax"`

	// 地域中文信息
	RegionZh *string `json:"RegionZh,omitempty" name:"RegionZh"`

	// 公网ip数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	EipAddress []*string `json:"EipAddress,omitempty" name:"EipAddress"`

	// 内外使用ip数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcIp []*string `json:"VpcIp,omitempty" name:"VpcIp"`

	// 实例关联子网数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Subnets []*string `json:"Subnets,omitempty" name:"Subnets"`

	// 0 :正常 1：正在初始化
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type NewModeItems struct {

	// 新增模式下接入的vpc列表
	VpcList []*string `json:"VpcList,omitempty" name:"VpcList"`

	// 新增模式下绑定的出口弹性公网ip列表，其中Eips和AddCount至少传递一个。
	Eips []*string `json:"Eips,omitempty" name:"Eips"`

	// 新增模式下新增绑定的出口弹性公网ip个数，其中Eips和AddCount至少传递一个。
	AddCount *int64 `json:"AddCount,omitempty" name:"AddCount"`
}

type RuleInfoData struct {

	// 执行顺序
	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`

	// 访问源
	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`

	// 访问目的
	TargetIp *string `json:"TargetIp,omitempty" name:"TargetIp"`

	// 协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 策略
	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`

	// 访问源类型，1是IP，3是域名，4是IP地址模版，5是域名地址模版
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// 描述
	Detail *string `json:"Detail,omitempty" name:"Detail"`

	// 访问目的类型，1是IP，3是域名，4是IP地址模版，5是域名地址模版
	TargetType *uint64 `json:"TargetType,omitempty" name:"TargetType"`

	// 端口
	Port *string `json:"Port,omitempty" name:"Port"`

	// id值
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 日志id，从告警处创建必传，其它为空
	LogId *string `json:"LogId,omitempty" name:"LogId"`

	// 城市Code
	City *uint64 `json:"City,omitempty" name:"City"`

	// 国家Code
	Country *uint64 `json:"Country,omitempty" name:"Country"`

	// 云厂商，支持多个，以逗号分隔， 1:腾讯云（仅中国香港及海外）,2:阿里云,3:亚马逊云,4:华为云,5:微软云
	CloudCode *string `json:"CloudCode,omitempty" name:"CloudCode"`

	// 是否为地域
	IsRegion *uint64 `json:"IsRegion,omitempty" name:"IsRegion"`

	// 城市名
	CityName *string `json:"CityName,omitempty" name:"CityName"`

	// 国家名
	CountryName *string `json:"CountryName,omitempty" name:"CountryName"`
}

type RunSyncAssetRequest struct {
	*tchttp.BaseRequest

	// 0: 互联网防火墙开关，1：vpc 防火墙开关
	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *RunSyncAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunSyncAssetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunSyncAssetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RunSyncAssetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 0：同步成功，1：资产更新中，2：后台同步调用失败
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunSyncAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunSyncAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanInfo struct {

	// 扫描结果信息
	ScanResultInfo *ScanResultInfo `json:"ScanResultInfo,omitempty" name:"ScanResultInfo"`

	// 扫描状态 0扫描中 1完成   2没赠送过扫描显示开启界面
	ScanStatus *int64 `json:"ScanStatus,omitempty" name:"ScanStatus"`

	// 进度
	ScanPercent *float64 `json:"ScanPercent,omitempty" name:"ScanPercent"`

	// 预计完成时间
	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`
}

type ScanResultInfo struct {

	// 暴漏漏洞数量
	LeakNum *uint64 `json:"LeakNum,omitempty" name:"LeakNum"`

	// 防护ip数量
	IPNum *uint64 `json:"IPNum,omitempty" name:"IPNum"`

	// 暴漏端口数量
	PortNum *uint64 `json:"PortNum,omitempty" name:"PortNum"`

	// 是否开启防护
	IPStatus *bool `json:"IPStatus,omitempty" name:"IPStatus"`

	// 是否拦截攻击
	IdpStatus *bool `json:"IdpStatus,omitempty" name:"IdpStatus"`

	// 是否禁封端口
	BanStatus *bool `json:"BanStatus,omitempty" name:"BanStatus"`
}

type SecurityGroupApiRuleData struct {

	// 访问源，入站时为Ip/Cidr，默认为0.0.0.0/0； 出站时当RuleType为1时，支持内网Ip/Cidr, 当RuleType为2时，填实例ID
	SourceId *string `json:"SourceId,omitempty" name:"SourceId"`

	// 访问目的，出站时为Ip/Cidr，默认为0.0.0.0/0；入站时当RuleType为1时，支持内网Ip/Cidr, 当RuleType为2时，填实例ID
	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`

	// 协议，支持ANY/TCP/UDP/ICMP
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 端口, 当Protocol为ANY或ICMP时，Port为-1/-1
	Port *string `json:"Port,omitempty" name:"Port"`

	// 策略, 1：阻断，2：放行
	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`

	// 描述
	Detail *string `json:"Detail,omitempty" name:"Detail"`

	// 规则类型，1：VpcId+Ip/Cidr, 2: 实例ID，入站时为访问目的类型，出站时为访问源类型
	RuleType *uint64 `json:"RuleType,omitempty" name:"RuleType"`

	// 执行顺序，中间插入必传，前插、后插非必传
	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`

	// 私有网络ID，当RuleType为1时必传
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type SecurityGroupBothWayInfo struct {

	// 执行顺序
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`

	// 访问源
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceId *string `json:"SourceId,omitempty" name:"SourceId"`

	// 访问源类型，默认为0，0: IP, 1: VPC, 2: SUBNET, 3: CVM, 4: CLB, 5: ENI, 6: CDB, 7: 参数模板, 100: 资产分组
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// 访问目的
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`

	// 访问目的类型，默认为0，0: IP, 1: VPC, 2: SUBNET, 3: CVM, 4: CLB, 5: ENI, 6: CDB, 7: 参数模板, 100: 资产分组
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetType *uint64 `json:"TargetType,omitempty" name:"TargetType"`

	// 协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 目的端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *string `json:"Port,omitempty" name:"Port"`

	// 策略, 1：阻断，2：放行
	// 注意：此字段可能返回 null，表示取不到有效值。
	Strategy *uint64 `json:"Strategy,omitempty" name:"Strategy"`

	// 方向，0：出站，1：入站，默认1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail *string `json:"Detail,omitempty" name:"Detail"`

	// 是否开关开启，0：未开启，1：开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 是否是正常规则，0：正常，1：异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNew *uint64 `json:"IsNew,omitempty" name:"IsNew"`

	// 单/双向下发，0:单向下发，1：双向下发
	// 注意：此字段可能返回 null，表示取不到有效值。
	BothWay *uint64 `json:"BothWay,omitempty" name:"BothWay"`

	// 私有网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 公网IP，多个以英文逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`

	// 内网IP，多个以英文逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`

	// 掩码地址，多个以英文逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`

	// 端口协议类型参数模板id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceTemplateId *string `json:"ServiceTemplateId,omitempty" name:"ServiceTemplateId"`

	// 是否使用端口协议模板，0：否，1：是
	ProtocolPortType *uint64 `json:"ProtocolPortType,omitempty" name:"ProtocolPortType"`
}

type SecurityGroupListData struct {

	// 执行顺序
	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`

	// 访问源
	SourceId *string `json:"SourceId,omitempty" name:"SourceId"`

	// 访问源类型，默认为0，1: VPC, 2: SUBNET, 3: CVM, 4: CLB, 5: ENI, 6: CDB, 7: 参数模板, 100: 资源组
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// 访问目的
	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`

	// 访问目的类型，默认为0，1: VPC, 2: SUBNET, 3: CVM, 4: CLB, 5: ENI, 6: CDB, 7: 参数模板, 100:资源组
	TargetType *uint64 `json:"TargetType,omitempty" name:"TargetType"`

	// 协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 目的端口
	Port *string `json:"Port,omitempty" name:"Port"`

	// 策略, 1：阻断，2：放行
	Strategy *uint64 `json:"Strategy,omitempty" name:"Strategy"`

	// 描述
	Detail *string `json:"Detail,omitempty" name:"Detail"`

	// 单/双向下发，0:单向下发，1：双向下发
	BothWay *uint64 `json:"BothWay,omitempty" name:"BothWay"`

	// 规则ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 是否开关开启，0：未开启，1：开启
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 是否是正常规则，0：正常，1：异常
	IsNew *uint64 `json:"IsNew,omitempty" name:"IsNew"`

	// 私有网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 公网IP，多个以英文逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`

	// 内网IP，多个以英文逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`

	// 掩码地址，多个以英文逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`

	// 端口协议类型参数模板id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceTemplateId *string `json:"ServiceTemplateId,omitempty" name:"ServiceTemplateId"`

	// 生成双向下发规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	BothWayInfo []*SecurityGroupBothWayInfo `json:"BothWayInfo,omitempty" name:"BothWayInfo"`

	// 方向，0：出站，1：入站，默认1
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// 是否使用端口协议模板，0：否，1：是
	ProtocolPortType *uint64 `json:"ProtocolPortType,omitempty" name:"ProtocolPortType"`
}

type SecurityGroupOrderIndexData struct {

	// 企业安全组规则当前执行顺序
	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`

	// 企业安全组规则更新目标执行顺序
	NewOrderIndex *uint64 `json:"NewOrderIndex,omitempty" name:"NewOrderIndex"`
}

type SequenceData struct {

	// 规则Id值
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 修改前执行顺序
	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`

	// 修改后执行顺序
	NewOrderIndex *uint64 `json:"NewOrderIndex,omitempty" name:"NewOrderIndex"`
}

type SetNatFwDnatRuleRequest struct {
	*tchttp.BaseRequest

	// 0：cfw新增模式，1：cfw接入模式。
	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`

	// 操作类型，可选值：add，del，modify。
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// 防火墙实例id，该字段必须传递。
	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`

	// 添加或删除操作的Dnat规则列表。
	AddOrDelDnatRules []*CfwNatDnatRule `json:"AddOrDelDnatRules,omitempty" name:"AddOrDelDnatRules"`

	// 修改操作的原始Dnat规则
	OriginDnat *CfwNatDnatRule `json:"OriginDnat,omitempty" name:"OriginDnat"`

	// 修改操作的新的Dnat规则
	NewDnat *CfwNatDnatRule `json:"NewDnat,omitempty" name:"NewDnat"`
}

func (r *SetNatFwDnatRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetNatFwDnatRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Mode")
	delete(f, "OperationType")
	delete(f, "CfwInstance")
	delete(f, "AddOrDelDnatRules")
	delete(f, "OriginDnat")
	delete(f, "NewDnat")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetNatFwDnatRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetNatFwDnatRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetNatFwDnatRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetNatFwDnatRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetNatFwEipRequest struct {
	*tchttp.BaseRequest

	// bind：绑定eip；unbind：解绑eip；newAdd：新增防火墙弹性公网ip
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// 防火墙实例id
	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`

	// 当OperationType 为bind或unbind操作时，使用该字段。
	EipList []*string `json:"EipList,omitempty" name:"EipList"`
}

func (r *SetNatFwEipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetNatFwEipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OperationType")
	delete(f, "CfwInstance")
	delete(f, "EipList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetNatFwEipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetNatFwEipResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetNatFwEipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetNatFwEipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StaticInfo struct {

	// 数
	Num *int64 `json:"Num,omitempty" name:"Num"`

	// 端口
	Port *string `json:"Port,omitempty" name:"Port"`

	// ip信息
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 地址
	Address *string `json:"Address,omitempty" name:"Address"`

	// 资产id
	InsID *string `json:"InsID,omitempty" name:"InsID"`

	// 资产名称
	InsName *string `json:"InsName,omitempty" name:"InsName"`
}

type StopSecurityGroupRuleDispatchRequest struct {
	*tchttp.BaseRequest
}

func (r *StopSecurityGroupRuleDispatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopSecurityGroupRuleDispatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopSecurityGroupRuleDispatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StopSecurityGroupRuleDispatchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true代表成功，false代表错误
	// 注意：此字段可能返回 null，表示取不到有效值。
		Status *bool `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopSecurityGroupRuleDispatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopSecurityGroupRuleDispatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchListsData struct {

	// 公网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`

	// 内网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntranetIp *string `json:"IntranetIp,omitempty" name:"IntranetIp"`

	// 实例名
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 资产类型
	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Area *string `json:"Area,omitempty" name:"Area"`

	// 防火墙开关
	Switch *int64 `json:"Switch,omitempty" name:"Switch"`

	// id值
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 公网 IP 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIpType *uint64 `json:"PublicIpType,omitempty" name:"PublicIpType"`

	// 风险端口数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PortTimes *uint64 `json:"PortTimes,omitempty" name:"PortTimes"`

	// 最近扫描时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`

	// 扫描深度
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanMode *string `json:"ScanMode,omitempty" name:"ScanMode"`

	// 扫描状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanStatus *uint64 `json:"ScanStatus,omitempty" name:"ScanStatus"`
}

type TLogInfo struct {

	// 失陷主机
	OutNum *int64 `json:"OutNum,omitempty" name:"OutNum"`

	// 待处置告警
	HandleNum *int64 `json:"HandleNum,omitempty" name:"HandleNum"`

	// 漏洞攻击
	VulNum *int64 `json:"VulNum,omitempty" name:"VulNum"`

	// 网络探测
	NetworkNum *int64 `json:"NetworkNum,omitempty" name:"NetworkNum"`

	// 封禁列表
	BanNum *int64 `json:"BanNum,omitempty" name:"BanNum"`

	// 暴力破解
	BruteForceNum *int64 `json:"BruteForceNum,omitempty" name:"BruteForceNum"`
}

type UnHandleEvent struct {

	// 伪攻击链类型
	EventTableListStruct []*UnHandleEventDetail `json:"EventTableListStruct,omitempty" name:"EventTableListStruct"`

	// 1 是  0否
	BaseLineUser *uint64 `json:"BaseLineUser,omitempty" name:"BaseLineUser"`

	// 1 打开 0 关闭
	BaseLineInSwitch *uint64 `json:"BaseLineInSwitch,omitempty" name:"BaseLineInSwitch"`

	// 1 打开 0 关闭
	BaseLineOutSwitch *uint64 `json:"BaseLineOutSwitch,omitempty" name:"BaseLineOutSwitch"`
}

type UnHandleEventDetail struct {

	// 安全事件名称
	EventName *string `json:"EventName,omitempty" name:"EventName"`

	// 未处置事件数量
	Total *uint64 `json:"Total,omitempty" name:"Total"`
}

type VpcDnsInfo struct {

	// vpc id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// vpc 名称
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// nat 防火墙模式 0：新增模式， 1: 接入模式
	FwMode *int64 `json:"FwMode,omitempty" name:"FwMode"`

	// vpc ipv4网段范围 CIDR（Classless Inter-Domain Routing，无类域间路由选择）
	VpcIpv4Cidr *string `json:"VpcIpv4Cidr,omitempty" name:"VpcIpv4Cidr"`

	// 外网弹性ip，防火墙 dns解析地址
	DNSEip *string `json:"DNSEip,omitempty" name:"DNSEip"`

	// nat网关id
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatInsId *string `json:"NatInsId,omitempty" name:"NatInsId"`

	// nat网关名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatInsName *string `json:"NatInsName,omitempty" name:"NatInsName"`

	// 0：开关关闭 ， 1: 开关打开
	SwitchStatus *int64 `json:"SwitchStatus,omitempty" name:"SwitchStatus"`
}

type VpcZoneData struct {

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// vpc节点地域
	Region *string `json:"Region,omitempty" name:"Region"`
}
