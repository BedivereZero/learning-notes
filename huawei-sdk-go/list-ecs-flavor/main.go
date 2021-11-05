package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

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
			Build(),
	)

	response, err := client.ListFlavors(&model.ListFlavorsRequest{})
	if err != nil {
		log.Fatal(err)
	}

	for _, flavor := range *response.Flavors {
		if !IsValidFlavor(flavor, 8, 32) {
			continue
		}
		if err := json.NewEncoder(os.Stdout).Encode(flavor); err != nil {
			log.Fatal(err)
		}
	}
}

func IsValidFlavor(flavor model.Flavor, cpu, ram int) bool {
	return (flavor.Vcpus == strconv.Itoa(cpu)) && (flavor.Ram == int32(ram*1024) && flavor.OsExtraSpecs.EcsinstanceArchitecture == nil)
}
