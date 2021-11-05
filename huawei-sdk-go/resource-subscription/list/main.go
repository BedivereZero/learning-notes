package main

import (
	"fmt"
	"os"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/global"
	bss "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2/model"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2/region"
)

func main() {
	ak := os.Getenv("HUAWEI_ACCESS_KEY_ID")
	sk := os.Getenv("HUAWEI_SECRET_ACCESS_KEY_ID")

	auth := global.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client := bss.NewBssClient(
		bss.BssClientBuilder().
			WithRegion(region.ValueOf("cn-north-1")).
			WithCredential(auth).
			Build())

	request := &model.ListPayPerUseCustomerResourcesRequest{}
	response, err := client.ListPayPerUseCustomerResources(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
