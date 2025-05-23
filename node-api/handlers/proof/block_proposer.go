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

package proof

import (
	"github.com/berachain/beacon-kit/node-api/handlers"
	"github.com/berachain/beacon-kit/node-api/handlers/proof/merkle"
	"github.com/berachain/beacon-kit/node-api/handlers/proof/types"
	"github.com/berachain/beacon-kit/node-api/handlers/utils"
)

// GetBlockProposer returns the block proposer pubkey for the given timestamp
// id along with a merkle proof that can be verified against the beacon block
// root. It also returns the merkle proof of the proposer index.
func (h *Handler) GetBlockProposer(c handlers.Context) (any, error) {
	params, err := utils.BindAndValidate[types.BlockProposerRequest](c, h.Logger())
	if err != nil {
		return nil, err
	}
	slot, beaconState, blockHeader, err := h.resolveTimestampID(params.TimestampID)
	if err != nil {
		return nil, err
	}

	h.Logger().Info("Generating block proposer proofs", "slot", slot)

	// Generate the proof (along with the "correct" beacon block root to verify against) for the
	// proposer validator's pubkey.
	bsm, err := beaconState.GetMarshallable()
	if err != nil {
		return nil, err
	}
	pubkeyProof, beaconBlockRoot, err := merkle.ProveProposerPubkeyInBlock(blockHeader, bsm)
	if err != nil {
		return nil, err
	}

	// Generate the proof for the proposer index.
	proposerIndexProof, _, err := merkle.ProveProposerIndexInBlock(blockHeader)
	if err != nil {
		return nil, err
	}

	// Get the pubkey of the proposer validator.
	proposerValidator, err := beaconState.ValidatorByIndex(blockHeader.GetProposerIndex())
	if err != nil {
		return nil, err
	}

	return types.BlockProposerResponse{
		BeaconBlockHeader:    blockHeader,
		BeaconBlockRoot:      beaconBlockRoot,
		ValidatorPubkey:      proposerValidator.GetPubkey(),
		ValidatorPubkeyProof: pubkeyProof,
		ProposerIndexProof:   proposerIndexProof,
	}, nil
}
