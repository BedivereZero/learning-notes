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

	request := &model.CreateHealthMonitorRequest{}
	healthmonitorbody := &model.CreateHealthMonitorOption{
		// AdminStateUp:   new(bool),
		Delay: int32(5),
		// DomainName:     new(string),
		// ExpectedCodes:  new(string),
		// HttpMethod:     new(string),
		MaxRetries: int32(3),
		MaxRetriesDown: func() *int32 {
			n := int32(3)
			return &n
		}(),
		// MonitorPort:    new(int32),
		// Name:           new(string),
		PoolId: "bf1b723d-9d95-47c2-a775-e949ffc459ba",
		// ProjectId: new(string),
		Timeout: int32(3),
		Type:    "TCP",
		// UrlPath: new(string),
	}
	request.Body = &model.CreateHealthMonitorRequestBody{
		Healthmonitor: healthmonitorbody,
	}
	response, err := client.CreateHealthMonitor(request)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.NewEncoder(os.Stdout).Encode(response.Healthmonitor); err != nil {
		log.Fatal(err)
	}
}
