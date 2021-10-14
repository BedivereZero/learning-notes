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

	request := &model.CreateL7policyRequest{}
	var listRulesL7policy = []model.CreateL7ruleReqInPolicy{
		{
			Type:        model.GetCreateL7ruleReqInPolicyTypeEnum().HOST_NAME,
			CompareType: "EQUAL_TO",
			Value:       "test-elb-bedivere.test.eisoo.com",
		},
	}
	nameL7policyCreateL7policyReq := "policy-test-elb-bedivere"
	redirectPoolIdL7policyCreateL7policyReq := "49d7de21-dcee-4f42-8d3c-053f4340e10c"
	l7policybody := &model.CreateL7policyReq{
		Name:           &nameL7policyCreateL7policyReq,
		Action:         model.GetCreateL7policyReqActionEnum().REDIRECT_TO_POOL,
		ListenerId:     "ed88dfc5-82b7-419b-9858-2d70ddab88db",
		RedirectPoolId: &redirectPoolIdL7policyCreateL7policyReq,
		Rules:          &listRulesL7policy,
	}
	request.Body = &model.CreateL7policyRequestBody{
		L7policy: l7policybody,
	}
	response, err := client.CreateL7policy(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
