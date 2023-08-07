// Copyright (c) 2023 qianjunakasumi (https://github.com/qianjunakasumi)
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// initMongoConnection
func initMongoConnection(url string) (*mongo.Database, error) {
	// Connect database
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	// ping
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	dataBase := client.Database("senjunico")
	return dataBase, nil
}

// todo init mongo collections and documents (need data model)
