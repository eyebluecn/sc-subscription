// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc/src/common/config"
	"github.com/eyebluecn/sc-misc/src/util"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api/subscriptionservice"
	"log"
	"time"

	"github.com/cloudwego/kitex/client"
)

func main() {
	c, err := subscriptionservice.NewClient("SubscriptionService", client.WithHostPorts(fmt.Sprintf("0.0.0.0:%v", config.SubscriptionServerPort)))
	if err != nil {
		log.Fatal(err)
	}
	for {
		req := &sc_subscription_api.SubscriptionPageRequest{
			PageNum:  1,
			PageSize: 20,
		}
		resp, err := c.SubscriptionPage(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		klog.Info(util.ToJSON(resp))
		time.Sleep(time.Second)
	}
}
