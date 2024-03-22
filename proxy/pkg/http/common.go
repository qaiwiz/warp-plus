// WriteHeader implements the http.ResponseWriter.WriteHeader method.
// It writes the HTTP status code to the response if it hasn't already been written.
func (w *responseWriter) WriteHeader(status int) {
    if !w.written {
        w.status = status
        w.written = true
        return
    }
    // If the headers have already been written, WriteHeader will panic.
    // However, we can still set the status code using http.Flush() if it hasn't been set yet.
    if w.status == 0 {
        w.status = status
    } else if status != w.status {
        http.Error(w, "header already written", http.StatusInternalServerError)
    }
}

// Write implements the http.ResponseWriter.Write method.
// It writes the data to the underlying net.Conn object.
func (w *responseWriter) Write(data []byte) (int, error)
