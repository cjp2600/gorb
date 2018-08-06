package gorb

type (
	Gorb struct {
		provider string
		table    string
	}

	Option func(*Gorb) error
)

// NewGorb
//
func NewGorb(options ...Option) (*Gorb, error) {

	obj := &Gorb{
		provider: "",
		table:    "",
	}

	for _, option := range options {
		if err := option(obj); err != nil {
			return nil, err
		}
	}

	return obj, nil
}

// Set storage provider table name
//
func SetTableName(name string) Option {
	return func(obj *Gorb) error {
		obj.table = name
		return nil
	}
}

// Set provider
//
func SetProvider(provider string) Option {
	return func(obj *Gorb) error {
		obj.provider = provider
		return nil
	}
}
