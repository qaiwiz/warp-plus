//go:build ios

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */

package device

// The following variables are defined to fit within memory limits for iOS's Network Extension API,
// which has stricter requirements. These are vars instead of consts, because heavier network
// extensions might want to reduce them further.

// QueueStagedSize is the maximum number of packets that can be stored in the staging queue.
var QueueStagedSize = 128

// QueueOutboundSize is the maximum number of packets that can be stored in the outbound queue.
var QueueOutboundSize = 1024

// QueueInboundSize is the maximum number of packets that can be stored in the inbound queue.
var QueueInboundSize = 1024

// QueueHandshakeSize is the maximum number of packets that can be stored in the handshake queue.
var QueueHandshakeSize = 1024

// PreallocatedBuffersPerPool is the number of preallocated buffers for each buffer pool.
var PreallocatedBuffersPerPool uint32 = 1024

// MaxSegmentSize is the maximum size of a single network segment in bytes.
const MaxSegmentSize = 1700
