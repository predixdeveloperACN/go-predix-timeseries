package filter

type Filter interface {
	Type() string
	Attributes() interface{}
}

type filter struct {
	type_      string
	attributes interface{}
}

func (f *filter) Type() string {
	return f.type_
}

func (f *filter) Attributes() interface{} {
	return f.attributes
}

func New(t string, attributes interface{}) Filter {
	return &filter{t, attributes}
}
