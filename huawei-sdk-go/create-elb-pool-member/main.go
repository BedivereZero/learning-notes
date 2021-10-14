package main

import (
	"fmt"
	"os"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	elb "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v2/model"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v2/region"
)

func main() {
	ak := os.Getenv("HUAWEI_ACCESS_KEY_ID")
	sk := os.Getenv("HUAWEI_SECRET_ACCESS_KEY_ID")

	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client := elb.NewElbClient(
		elb.ElbClientBuilder().
			WithRegion(region.ValueOf("cn-east-3")).
			WithCredential(auth).
			Build())

	request := &model.CreateMemberRequest{}
	request.PoolId = "49d7de21-dcee-4f42-8d3c-053f4340e10c"
	nameMemberCreateMemberReq := "member-test-elb-bedivere"
	memberbody := &model.CreateMemberReq{
		Name:         &nameMemberCreateMemberReq,
		ProtocolPort: int32(80),
		SubnetId:     "ed42b5ab-3f89-4021-bff2-eb4b836210c0",
		Address:      "192.168.0.220",
	}
	request.Body = &model.CreateMemberRequestBody{
		Member: memberbody,
	}
	response, err := client.CreateMember(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
