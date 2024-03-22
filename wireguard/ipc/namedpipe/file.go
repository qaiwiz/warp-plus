// Named pipe package provides a named pipe implementation for the Windows platform.
package namedpipe

import (
	// ... (imports omitted for brevity)
)

// A timeoutChan is a channel that is used to signal timeouts for IO operations.
type timeoutChan chan struct{}

// ioInitOnce is used to ensure that the ioCompletionPort is initialized only once.
var ioInitOnce sync.Once

// ioCompletionPort is the Windows IO completion port used for asynchronous IO operations.
var ioCompletionPort windows.Handle

// ioResult contains the result of an asynchronous IO operation, including the number of bytes transferred and any error.
type ioResult struct {
	bytes uint32 // Number of bytes transferred.
	err   error  // Error, if any.
}

// ioOperation represents an outstanding asynchronous Win32 IO operation.
type ioOperation struct {
	o       windows.Overlapped // Overlapped structure used for asynchronous IO.
	ch      chan ioResult     // Channel to receive the result of the IO operation.
}

// initIo initializes the ioCompletionPort used for asynchronous IO operations.
func initIo() {
	// ... (code omitted for brevity)
}

// file implements the io.Reader, io.Writer, and io.Closer interfaces for a Win32 handle without blocking in a syscall.
type file struct {
	handle        windows.Handle // The Win32 handle.
	wg            sync.WaitGroup // Wait group used to track outstanding IO operations.
	wgLock        sync.RWMutex   // Mutex used to synchronize access to the wait group.
	closing       atomic.Bool   // Flag indicating whether the file is being closed.
	socket        bool           // Flag indicating whether the handle is a socket.
	readDeadline  deadlineHandler // Deadline handler for read operations.
	writeDeadline deadlineHandler // Deadline handler for write operations.
}

// deadlineHandler handles deadlines for IO operations.
type deadlineHandler struct {
	setLock     sync.Mutex     // Mutex used to synchronize access to the channel and timer.
	channel     timeoutChan    // Channel used to signal timeouts.
	channelLock sync.RWMutex   // Mutex used to synchronize access to the channel.
	timer       *time.Timer    // Timer used to wait for the deadline.
	timedout    atomic.Bool    // Flag indicating whether the deadline has been exceeded.
}

// makeFile creates a new file from an existing Win32 handle.
func makeFile(h windows.Handle) (*file, error) {
	// ... (code omitted for brevity)
}

// closeHandle closes the resources associated with a Win32 handle.
func (f *file) closeHandle() {
	// ... (code omitted for brevity)
}

// Close closes the file and releases any associated resources.
func (f *file) Close() error {
	// ... (code omitted for brevity)
}

// prepareIo prepares for a new IO operation and returns an ioOperation structure that can be used to track the operation.
func (f *file) prepareIo() (*ioOperation, error) {
	// ... (code omitted for brevity)
}

// ioCompletionProcessor processes completed async IOs forever.
func ioCompletionProcessor(h windows.Handle) {
	// ... (code omitted for brevity)
}

// asyncIo processes the return value from ReadFile or WriteFile, blocking until the operation has actually completed.
func (f *file) asyncIo(c *ioOperation, d *deadlineHandler, bytes uint32, err error) (int, error) {
	// ... (code omitted for brevity)
}

// Read reads from a file handle.
func (f *file) Read(b []byte) (int, error) {
	// ... (code omitted for brevity)
}

// Write writes to a file handle.
func (f *file) Write(b []byte) (int, error) {
	// ... (code omitted for brevity)
}

// SetReadDeadline sets the read deadline for the file.
func (f *file) SetReadDeadline(deadline time.Time) error {
	// ... (code omitted for brevity)
}

// SetWriteDeadline sets the write deadline for the file.
func (f *file) SetWriteDeadline(deadline time.Time) error {
	// ... (code omitted for brevity)
}

// Flush flushes any buffered data for the file.
func (f *file) Flush() error {
	// ... (code omitted for brevity)
}

// Fd returns the file descriptor for the file.
func (f *file) Fd() uintptr {
	// ... (code omitted for brevity)
}

// set sets the deadline for the deadlineHandler.
func (d *deadlineHandler) set(deadline time.Time) error {
	// ... (code omitted for brevity)
}
