// Copyright (c) 2023 qianjunakasumi (https://github.com/qianjunakasumi)
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package uid

type SenjUID struct {
	grpcUrl string
	name    string
	// grpc instance
}

func (s *SenjUID) NewProvider() (IUIDService, error) {
	//TODO implement me
	//init grpc client
	panic("implement me")
}

func (s *SenjUID) Generate() (UID, error) {
	//TODO implement me
	//todo connect RPC service to get snowflake id
	panic("implement me")
}
