package contract

import "errors"
import "fmt"

// Contract ...
type Contract struct {
	contractFunc func() error
}

func (c *Contract) check() error {
	return c.contractFunc()
}

// IsPositive ...
func IsPositive(num float64) *Contract {
	return &Contract{func() error {
		if num < 1 {
			return errors.New("int is not positive")
		}
		return nil
	}}
}

// IsZeroOrPositive ...
func IsZeroOrPositive(num float64) *Contract {
	return &Contract{func() error {
		if num < 0 {
			return errors.New("int is not positive")
		}
		return nil
	}}
}

// IsNegative ...
func IsNegative(num float64) *Contract {
	return &Contract{func() error {
		if num < 1 {
			return errors.New("int is not negative")
		}
		return nil
	}}
}

// NotNil ...
func NotNil(instance interface{}) *Contract {
	return &Contract{func() error {
		if instance == nil {
			return errors.New("instance is nil")
		}
		return nil
	}}
}

// NotDefualtNilValue ...
func NotDefualtNilValue(val interface{}) *Contract {
	return &Contract{func() error {
		err := errors.New("nil value")
		switch t := val.(type) {
		case string:
			if t == "" {
				return err
			}
		case float64:
			fmt.Println(t)
			if t == 0 {
				return err
			}
			//etc.
		}
		return nil
	}}
}

// Wrap ...
func Wrap(contracts ...*Contract) {
	for _, c := range contracts {
		err := c.check()
		if err != nil {
			panic(err)
		}
	}
}

// Check ...
func (c *Contract) Check() {
	if err := c.check(); err != nil {
		panic(err)
	}
}

// WrapWithError ...
func WrapWithError(contracts ...*Contract) error {
	for _, c := range contracts {
		err := c.check()
		if err != nil {
			return err
		}
	}
	return nil
}

// CheckWithError ...
func (c *Contract) CheckWithError() error {
	return c.check()
}
