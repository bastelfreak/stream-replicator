// Copyright (c) 2021, R.I. Pienaar and the Project contributors
//
// SPDX-License-Identifier: Apache-2.0

package election

import (
	"time"

	"github.com/nats-io/nats.go"
)

// Option configures the election system
type Option func(o *options)

type options struct {
	name       string
	key        string
	bucket     nats.KeyValue
	ttl        time.Duration
	cInterval  time.Duration
	replicator string
	wonCb      func()
	lostCb     func()
	campaignCb func(s State)
	bo         Backoff
	debug      func(format string, a ...any)
}

// WithBackoff will use the provided Backoff timer source to decrease campaign intervals over time
func WithBackoff(bo Backoff) Option {
	return func(o *options) { o.bo = bo }
}

// OnWon is a callback called when winning an election
func OnWon(cb func()) Option {
	return func(o *options) { o.wonCb = cb }
}

// OnLost is a callback called when losing an election
func OnLost(cb func()) Option {
	return func(o *options) { o.lostCb = cb }
}

// OnCampaign is called each time a campaign is done by the leader or a candidate
func OnCampaign(cb func(s State)) Option {
	return func(o *options) { o.campaignCb = cb }
}

// WithDebug sets a function to do debug logging with
func WithDebug(cb func(format string, a ...any)) Option {
	return func(o *options) { o.debug = cb }
}

// WithReplicator sets the replicator name for monitoring
func WithReplicator(r string) Option {
	return func(o *options) { o.replicator = r }
}

// SkipTTLValidateForTests turns off Bucket TTL validation for testing
func SkipTTLValidateForTests() {
	skipValidate = true
}
