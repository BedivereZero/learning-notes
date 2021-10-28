package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	elb "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v3"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v3/model"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v3/region"
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

	response, err := client.ListHealthMonitors(new(model.ListHealthMonitorsRequest))
	if err != nil {
		log.Fatal(err)
	}

	if err := json.NewEncoder(os.Stdout).Encode(response.Healthmonitors); err != nil {
		log.Fatal(err)
	}
}
