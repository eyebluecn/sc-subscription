package config

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api/miscservice"
)

var (
	MiscRpcClient miscservice.Client
)

func InitRpcClient() {

	miscRpcClient, err := miscservice.NewClient("MiscService", client.WithHostPorts(fmt.Sprintf("0.0.0.0:%v", MiscServerPort)))
	if err != nil {
		klog.CtxInfof(context.Background(), "miscservice client init error: %v", err)
		panic(err)
	}

	MiscRpcClient = miscRpcClient

}
