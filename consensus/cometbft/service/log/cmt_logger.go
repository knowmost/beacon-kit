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

package log

import (
	"github.com/berachain/beacon-kit/log"
	"github.com/berachain/beacon-kit/log/phuslu"
	cmtlog "github.com/cometbft/cometbft/libs/log"
)

type CometLogger struct {
	log.AdvancedLogger[*phuslu.Logger]
}

func WrapCometLogger(logger *phuslu.Logger) *CometLogger {
	return &CometLogger{
		AdvancedLogger: logger,
	}
}

func (l *CometLogger) With(keyVals ...any) cmtlog.Logger {
	return &CometLogger{
		AdvancedLogger: l.AdvancedLogger.With(keyVals...),
	}
}

func (l *CometLogger) Impl() any {
	return l.AdvancedLogger
}
