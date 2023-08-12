

package util

import "github.com/miekg/dns"

// ResponseInterceptor xxx
type ResponseInterceptor struct {
	dns.ResponseWriter
	Msg *dns.Msg
}

// NewResponseInterceptor returns a pointer to a new ResponseReverter.
func NewResponseInterceptor(w dns.ResponseWriter) *ResponseInterceptor {
	return &ResponseInterceptor{ResponseWriter: w}
}

// WriteMsg records the status code and calls the
// underlying ResponseWriter's WriteMsg method.
func (r *ResponseInterceptor) WriteMsg(res *dns.Msg) error {
	r.Msg = res
	return nil
}

// Write is a wrapper that records the size of the message that gets written.
func (r *ResponseInterceptor) Write(buf []byte) (int, error) {
	return len(buf), nil
}

// Hijack implements dns.Hijacker. It simply wraps the underlying
// ResponseWriter's Hijack method if there is one, or returns an error.
func (r *ResponseInterceptor) Hijack() {
	return
}
