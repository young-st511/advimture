package editor

// Register holds yanked/deleted text content.
type Register struct {
	Content  string
	Linewise bool
}

// NewRegister creates an empty register.
func NewRegister() *Register {
	return &Register{}
}

// Set stores content in the register.
func (r *Register) Set(content string, linewise bool) {
	r.Content = content
	r.Linewise = linewise
}
