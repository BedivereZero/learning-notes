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

	request := &model.CancelResourcesSubscriptionRequest{}
	var listResourceIdsbody = []string{
		"57bbea29-6c56-4bb5-b28e-99db65f39384",
	}
	request.Body = &model.UnsubscribeResourcesReq{
		UnsubscribeType: int32(1),
		ResourceIds:     listResourceIdsbody,
	}
	response, err := client.CancelResourcesSubscription(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
