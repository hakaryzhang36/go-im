package utils

import (
	"context"
	"fmt"
)

func Publish(ctx context.Context, channel string, msg string) error {
	err := Rdb.Publish(ctx, channel, msg).Err()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Pub %s : %s\n", channel, msg)
	}
	return err
}

func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := Rdb.PSubscribe(ctx, channel)
	rec, err := sub.ReceiveMessage(ctx)
	if err != nil {
		fmt.Println(err)
		return "", err
	} else {
		fmt.Printf("Sub %s : %s\n", channel, rec.Payload)
		return rec.Payload, err
	}
}
