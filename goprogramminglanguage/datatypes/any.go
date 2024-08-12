package datatypes

// any - lacks expressiveness

type Customer struct {}

type Contract struct {}

type Store struct{}

func (s *Store) Get(id string) (any, error) { // returns any
	return 1, nil
}

func (s *Store) Set(id string, v any) error { // accepts any
	return nil
}

// When any can be useful

// 1. From the standard library encoding/json

// Anything can be marshaled
func Marshal(v any) ([]byte, error) {return nil, nil}

// the other is from database/sql package.
// func (c *Conn) QueryContext(ctx context.Context, query string, args ...any) (*Rows, error) {}
