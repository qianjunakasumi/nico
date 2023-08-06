// Copyright (c) 2023 qianjunakasumi (https://github.com/qianjunakasumi)
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package uid

// UID is used for most identification codes of nico,
// including url address conversion (core component)
type UID int64

// IUIDService provide abstraction of global number allocator,
// it's around to generate an UID
type IUIDService interface {
	// Generate an UID
	Generate() (UID, error)
}

type IUIDProvider interface {
	// NewProvider return an IUIDService instance which actually serve nico
	NewProvider() (IUIDService, error)
}

// NewProvider create an IUIDService
func NewProvider(p IUIDProvider) (IUIDService, error) {
	return p.NewProvider()
}
