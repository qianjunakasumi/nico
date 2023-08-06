// Copyright (c) 2023 qianjunakasumi (https://github.com/qianjunakasumi)
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package uid

import "github.com/qianjunakasumi/nico/pkg/snowflake"

type Snowflake struct {
	nodeId int64
	epoch  int64
	node   *snowflake.Node
}

func (s *Snowflake) NewProvider() (IUIDService, error) {
	node, err := snowflake.NewNode(s.nodeId, s.epoch)
	if err != nil {
		return nil, err
	}

	s.node = node
	return s, nil
}

func (s *Snowflake) Generate() (UID, error) {
	return UID(s.node.Generate()), nil
}
