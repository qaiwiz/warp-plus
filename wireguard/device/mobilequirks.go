// DisableSomeRoamingForBrokenMobileSemantics is a method on the Device struct that is used to disable roaming for broken mobile semantics.
// This function should ideally be called before peers are created, but it will attempt to handle the situation and race conditions
// if called after peers have already been created.
//
// The function works by setting the `brokenRoaming` field on the `net` field of the `Device` struct to `true`.
// It then locks the `peers` field using a read lock and iterates over the `keyMap` field of the `peers` field.
// For each peer in the `keyMap`, the function locks the `endpoint` field, sets the `disableRoaming` field to `true` if the `val` field is not `nil`, and then unlocks the `endpoint` field.
//
// After all peers have been processed, the function releases the read lock on the `peers` field.
//
// This function is intended to be used in situations where the mobile semantics are broken and roaming needs to be disabled.
// It is important to note that this function should be called as early as possible to avoid race conditions and ensure proper operation.
func (device *Device) DisableSomeRoamingForBrokenMobileSemantics() {
	device.net.brokenRoaming = true
	device.peers.RLock()
	for _, peer := range device.peers.keyMap {
		peer.endpoint.Lock()
	
