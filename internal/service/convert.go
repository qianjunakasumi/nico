// Copyright (c) 2023 qianjunakasumi (https://github.com/qianjunakasumi)
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package service

import "strings"

// UID2URLBody converts integer to BaseX
func UID2URLBody(alphabet string, id int64) string {
	length := int64(len(alphabet))
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(6)
	// [Comment]
	// 6 maybe usual max length of the encoded string, but in different alphabet and different uid
	// it may be different, so we need add system config to set it dynamic
	// calculate the length of the encoded string could compute in front end
	for ; id > 0; id = id / length {
		encodedBuilder.WriteByte(alphabet[(id % length)])
	}

	return encodedBuilder.String()
}
