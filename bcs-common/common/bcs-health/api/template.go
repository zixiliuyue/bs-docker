

package api

const healthTemplate string = `*** BCS Module Status Change Alarm ***
	Module:    {{.Module}}
	Kind:      {{.Kind}}
	IP:        {{.IP}}
	ClusterID: {{.ClusterID}}
	Message:   {{.Message}}
`

// Writer write implementation
type Writer struct {
	buf []byte
}

// Write interface implementation
func (w *Writer) Write(p []byte) (int, error) {
	w.buf = append(w.buf, p...)
	return len(p), nil
}

// String 用于打印
func (w *Writer) String() string {
	return string(w.buf)
}
