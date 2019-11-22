package memorable

// New for create an instance of the builder pattern
func New() CodeBuilder {
	return &codeBuilder{}
}
