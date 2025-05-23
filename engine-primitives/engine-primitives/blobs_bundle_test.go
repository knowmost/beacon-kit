// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2025, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package engineprimitives_test

import (
	"testing"

	engineprimitives "github.com/berachain/beacon-kit/engine-primitives/engine-primitives"
	"github.com/berachain/beacon-kit/primitives/eip4844"
	"github.com/stretchr/testify/require"
)

func TestBlobsBundleV1(t *testing.T) {
	t.Parallel()
	bundle := &engineprimitives.BlobsBundleV1{
		Commitments: []eip4844.KZGCommitment{{1, 2, 3}, {4, 5, 6}},
		Proofs:      []eip4844.KZGProof{{7, 8, 9}, {10, 11, 12}},
		Blobs:       []*eip4844.Blob{{13, 14, 15}, {16, 17, 18}},
	}

	commitments := bundle.GetCommitments()
	require.Equal(t, bundle.Commitments, commitments)

	proofs := bundle.GetProofs()
	require.Equal(t, bundle.Proofs, proofs)

	blobs := bundle.GetBlobs()
	require.Equal(t, bundle.Blobs, blobs)
}
