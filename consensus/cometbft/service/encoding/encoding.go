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

package encoding

import (
	"fmt"

	ctypes "github.com/berachain/beacon-kit/consensus-types/types"
	datypes "github.com/berachain/beacon-kit/da/types"
	"github.com/berachain/beacon-kit/primitives/common"
	"github.com/berachain/beacon-kit/primitives/encoding/ssz"
)

// ExtractBlobsAndBlockFromRequest extracts the blobs and block from an ABCI
// request.
func ExtractBlobsAndBlockFromRequest(
	req ABCIRequest,
	beaconBlkIndex uint,
	blobSidecarsIndex uint,
	forkVersion common.Version,
) (*ctypes.SignedBeaconBlock, datypes.BlobSidecars, error) {
	if req == nil {
		return nil, nil, ErrNilABCIRequest
	}

	blk, err := UnmarshalBeaconBlockFromABCIRequest(
		req.GetTxs(),
		beaconBlkIndex,
		forkVersion,
	)
	if err != nil {
		return nil, nil, err
	}

	blobs, err := UnmarshalBlobSidecarsFromABCIRequest(
		req.GetTxs(),
		blobSidecarsIndex,
	)

	return blk, blobs, err
}

// UnmarshalBeaconBlockFromABCIRequest extracts a beacon block from an ABCI
// request.
func UnmarshalBeaconBlockFromABCIRequest(
	txs [][]byte,
	bzIndex uint,
	forkVersion common.Version,
) (*ctypes.SignedBeaconBlock, error) {
	var signedBlk *ctypes.SignedBeaconBlock
	lenTxs := uint(len(txs))

	// Ensure there are transactions in the request and that the request is
	// valid.
	if txs == nil || lenTxs == 0 {
		return signedBlk, ErrNoBeaconBlockInRequest
	}
	if bzIndex >= lenTxs {
		return signedBlk, ErrBzIndexOutOfBounds
	}

	// Extract the beacon block from the ABCI request.
	blkBz := txs[bzIndex]
	if blkBz == nil {
		return signedBlk, ErrNilBeaconBlockInRequest
	}

	block, err := ctypes.NewEmptySignedBeaconBlockWithVersion(forkVersion)
	if err != nil {
		return nil, fmt.Errorf("attempt at building block with wrong version %s: %w", forkVersion, err)
	}
	if err = ssz.Unmarshal(blkBz, block); err != nil {
		return nil, err
	}
	return block, nil
}

// UnmarshalBlobSidecarsFromABCIRequest extracts blob sidecars from an ABCI
// request.
func UnmarshalBlobSidecarsFromABCIRequest(
	txs [][]byte,
	bzIndex uint,
) (datypes.BlobSidecars, error) {
	if len(txs) == 0 || bzIndex >= uint(len(txs)) {
		return nil, ErrNoBlobSidecarInRequest
	}

	sidecarBz := txs[bzIndex]
	if sidecarBz == nil {
		return nil, ErrNilBlobSidecarInRequest
	}

	var sidecars datypes.BlobSidecars
	if err := ssz.Unmarshal(sidecarBz, &sidecars); err != nil {
		return nil, err
	}
	return sidecars, nil
}
