package mq

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"sync"
)

// 消息发送的目标Topic名称，需要提前在RocketMQ中创建。
// 一般一个应用对应一个TOPIC.
const TOPIC = "SmartClassroomTopic"

type Producer struct {
}

var (
	producer     *Producer
	producerOnce sync.Once
)

// DefaultXXX 是单例模式写法
func DefaultProducer() *Producer {
	producerOnce.Do(func() {
		producer = &Producer{}
	})
	return producer
}

// 初始化
func InitProducer() {
	ctx := context.Background()

	_ = DefaultProducer()
	klog.CtxInfof(ctx, "初始化mq_producer成功")
}

func (receiver *Producer) Publish(ctx context.Context, tags string, keys string, body string) error {
	klog.CtxInfof(ctx, "模拟发送MQ，topic=%v tags=%v keys=%v body=%v", TOPIC, tags, keys, body)
	return nil
}
