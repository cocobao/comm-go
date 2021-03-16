package service

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func SetupMgo(DB, addr, username, password string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	op := &options.ClientOptions{}
	if len(username) > 0 && len(password) > 0 {
		op.SetAuth(options.Credential{
			AuthMechanism: "SCRAM-SHA-1",
			AuthSource:    DB,
			Username:      username,
			Password:      password,
		})
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(addr), op)
	if err != nil {
		fmt.Println("create mongo client fail,", err)
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Printf("ping to mongo fail, err:%+v\n, addr:%s, username:%s, password:%s\n", err, addr, username, password)
		return nil, err
	}
	return client, nil
}
