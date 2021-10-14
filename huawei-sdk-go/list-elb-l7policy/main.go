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

	request := &model.ListL7policiesRequest{}
	nameRequest := "policy-test-elb-bedivere"
	request.Name = &nameRequest
	response, err := client.ListL7policies(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
