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

	request := &model.ResizeServerRequest{}
	request.ServerId = os.Getenv("HUAWEI_SERVER_ID")
	modeResizeResizePrePaidServerOption := "withStopServer"
	resizebody := &model.ResizePrePaidServerOption{
		FlavorRef: "t6.medium.2",
		Mode:      &modeResizeResizePrePaidServerOption,
	}
	request.Body = &model.ResizeServerRequestBody{
		Resize: resizebody,
	}

	response, err := client.ResizeServer(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
