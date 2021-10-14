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

	request := &model.CreatePoolRequest{}
	sessionPersistencePool := &model.SessionPersistence{
		Type: model.GetSessionPersistenceTypeEnum().SOURCE_IP,
	}
	loadbalancerIdPoolCreatePoolReq := "c6e6382b-ff39-4939-9ab0-184b73322257"
	namePoolCreatePoolReq := "server_group-test-elb-bedivere"
	poolbody := &model.CreatePoolReq{
		Protocol:           model.GetCreatePoolReqProtocolEnum().HTTP,
		LbAlgorithm:        "ROUND_ROBIN",
		LoadbalancerId:     &loadbalancerIdPoolCreatePoolReq,
		Name:               &namePoolCreatePoolReq,
		SessionPersistence: sessionPersistencePool,
	}
	request.Body = &model.CreatePoolRequestBody{
		Pool: poolbody,
	}
	response, err := client.CreatePool(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
