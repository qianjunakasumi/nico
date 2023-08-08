// Copyright (c) 2023 qianjunakasumi (https://github.com/qianjunakasumi)
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package database

import (
	"context"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// InitMongoConnection returns a mongo client and database connection
func InitMongoConnection(ctx context.Context, url string) (client *mongo.Client, database *mongo.Database, err error) {
	log.Debug().Str("target", url).Msg("connecting mongodb")

	client, err = mongo.Connect(
		ctx,
		options.Client().ApplyURI(url),
	)
	if err != nil {
		return
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return
	}

	database = client.Database("senjunico")
	log.Debug().Msg("connected mongodb and apply database to senjunico")

	return
}

// DisconnectMongo disconnects the mongo client
func DisconnectMongo(ctx context.Context, mongo *mongo.Client) error {
	return mongo.Disconnect(ctx)
}

// todo init mongo collections and documents (need data model)
