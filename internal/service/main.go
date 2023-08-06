// Copyright (c) 2023 qianjunakasumi (https://github.com/qianjunakasumi)
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package service

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type NicoService struct {
	mongo *mongo.Database
}

func NewNicoService(m *mongo.Database) *NicoService {
	return &NicoService{m}
}

func (n *NicoService) Run(ctx context.Context) error {

	// todo pull up gRPC service
	// todo pull up graphQL service

	return nil
}
