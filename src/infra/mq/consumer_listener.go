package mq

import "context"

// mq消费者的函数类型
type ConsumerListener func(ctx context.Context, keys string, body string) error
