package main

import (
	"fmt"
	"os"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/region"
)

func main() {
	ak := os.Getenv("HUAWEI_ACCESS_KEY_ID")
	sk := os.Getenv("HUAWEI_SECRET_ACCESS_KEY_ID")

	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client := ecs.NewEcsClient(
		ecs.EcsClientBuilder().
			WithRegion(region.ValueOf("cn-east-3")).
			WithCredential(auth).
			Build())

	request := &model.CreateServersRequest{}
	chargingModeExtendparamPrePaidServerExtendParam := model.GetPrePaidServerExtendParamChargingModeEnum().PRE_PAID
	periodTypeExtendparamPrePaidServerExtendParam := model.GetPrePaidServerExtendParamPeriodTypeEnum().MONTH
	periodNumExtendparamPrePaidServerExtendParam := int32(1)
	isAutoRenewExtendparamPrePaidServerExtendParam := model.GetPrePaidServerExtendParamIsAutoRenewEnum().TRUE
	isAutoPayExtendparamPrePaidServerExtendParam := model.GetPrePaidServerExtendParamIsAutoPayEnum().TRUE
	enterpriseProjectIdExtendparamPrePaidServerExtendParam := "eb39b1d2-5565-4249-ac69-814b39f67da4"
	extendparamServer := &model.PrePaidServerExtendParam{
		ChargingMode:        &chargingModeExtendparamPrePaidServerExtendParam,
		PeriodType:          &periodTypeExtendparamPrePaidServerExtendParam,
		PeriodNum:           &periodNumExtendparamPrePaidServerExtendParam,
		IsAutoRenew:         &isAutoRenewExtendparamPrePaidServerExtendParam,
		IsAutoPay:           &isAutoPayExtendparamPrePaidServerExtendParam,
		EnterpriseProjectId: &enterpriseProjectIdExtendparamPrePaidServerExtendParam,
	}
	idSecurityGroupsPrePaidServerSecurityGroup := "76928403-3bec-481c-ac98-a532c08aec91"
	var listSecurityGroupsServer = []model.PrePaidServerSecurityGroup{
		{
			Id: &idSecurityGroupsPrePaidServerSecurityGroup,
		},
	}
	var listDataVolumesServer = []model.PrePaidServerDataVolume{
		{
			Volumetype: model.GetPrePaidServerDataVolumeVolumetypeEnum().SAS,
			Size:       int32(60),
		},
	}
	sizeRootVolumePrePaidServerRootVolume := int32(40)
	rootVolumeServer := &model.PrePaidServerRootVolume{
		Volumetype: model.GetPrePaidServerRootVolumeVolumetypeEnum().SAS,
		Size:       &sizeRootVolumePrePaidServerRootVolume,
	}
	chargingModeExtendparamPrePaidServerEipExtendParam := model.GetPrePaidServerEipExtendParamChargingModeEnum().PRE_PAID
	extendparamEip := &model.PrePaidServerEipExtendParam{
		ChargingMode: &chargingModeExtendparamPrePaidServerEipExtendParam,
	}
	sizeBandwidthPrePaidServerEipBandwidth := int32(5)
	bandwidthEip := &model.PrePaidServerEipBandwidth{
		Size:      &sizeBandwidthPrePaidServerEipBandwidth,
		Sharetype: model.GetPrePaidServerEipBandwidthSharetypeEnum().PER,
	}
	eipPublicip := &model.PrePaidServerEip{
		Iptype:      "5_bgp",
		Bandwidth:   bandwidthEip,
		Extendparam: extendparamEip,
	}
	publicipServer := &model.PrePaidServerPublicip{
		Eip: eipPublicip,
	}
	var listNicsServer = []model.PrePaidServerNic{
		{
			SubnetId: "4d25dbc4-c04c-41ce-a25a-520d2fb905dd",
		},
	}
	serverbody := &model.PrePaidServer{
		ImageRef:       "4cc8a329-32c3-40bd-9ec8-3ea2043ce4cd",
		FlavorRef:      "s6.xlarge.4",
		Name:           "test-aksk",
		Vpcid:          "7399959d-5b73-4199-9878-48866127e1c8",
		Nics:           listNicsServer,
		Publicip:       publicipServer,
		RootVolume:     rootVolumeServer,
		DataVolumes:    &listDataVolumesServer,
		SecurityGroups: &listSecurityGroupsServer,
		Extendparam:    extendparamServer,
	}
	request.Body = &model.CreateServersRequestBody{
		Server: serverbody,
	}
	response, err := client.CreateServers(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
