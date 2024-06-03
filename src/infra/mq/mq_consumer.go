package mq

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"sync"
)

// MQ消费者，实际情况是一个topic对应一个消费者。这里把所有的消费者都写到这里了。
type Consumer struct {
	listeners map[string][]ConsumerListener
}

var (
	consumer     *Consumer
	consumerOnce sync.Once
)

// DefaultXXX 是单例模式写法
func DefaultConsumer() *Consumer {
	consumerOnce.Do(func() {
		consumer = &Consumer{
			listeners: make(map[string][]ConsumerListener),
		}
	})
	return consumer
}

// 初始化
func InitConsumer() {
	ctx := context.Background()

	_ = DefaultConsumer()
	klog.CtxInfof(ctx, "初始化mq_consumer成功")
}

// infra作为底层组件，为了让依赖关系不变，这里使用注册listener的方式。
func (receiver *Consumer) Register(tags string, listener ConsumerListener) {
	if tags == "" {
		panic("tags can not be empty")
	}

	if receiver.listeners[tags] == nil {
		receiver.listeners[tags] = []ConsumerListener{listener}
	} else {
		receiver.listeners[tags] = append(receiver.listeners[tags], listener)
	}
}

// 模拟Mq消费者，收到消息。
func (receiver *Consumer) Receive(ctx context.Context, topic string, tags string, keys string, body string) error {
	klog.CtxInfof(ctx, "模拟接收MQ，topic=%v tags=%v keys=%v body=%v", topic, tags, keys, body)

	//将消息分发给listeners.
	if listeners, ok := receiver.listeners[tags]; ok {
		klog.CtxErrorf(ctx, "注册了有 %v 个方法能够处理 tags=%v", len(listeners), tags)
		for _, listener := range listeners {
			err := listener(ctx, keys, body)
			if err != nil {
				//仅打印错误即可。
				klog.CtxErrorf(ctx, "在处理消息到达时出错了 %v", err)
			}
		}
	} else {
		klog.CtxErrorf(ctx, "没有注册方法能够处理 tags=%v", tags)
	}
	return nil
}
