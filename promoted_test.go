package promoted

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// interface
type I interface {
	Name()
	Zoo()
	Foo()
}

// Base
type Base struct {
}

func (b Base) Name() {}

func (b Base) Foo() {}

func (b Base) Zoo() {}

// Sub struct
type Sub struct {
	Base
}

func (s Sub) Foo() {}

// SubSub
type SubSub struct {
	Sub
}

func (s SubSub) Zoo() {}

// test case

func TestGetOverriddenMethods(t *testing.T) {
	var i *I
	var methods, expected []string

	methods = GetOverriddenMethods((*Sub)(nil), (*Base)(nil), i)
	expected = []string{"Foo"}
	_ = assert.EqualValuesf(t, expected, methods, "GetOverriddenMethods test Sub fail")

	methods = GetOverriddenMethods((*SubSub)(nil), (*Sub)(nil), i)
	expected = []string{"Zoo"}
	_ = assert.EqualValuesf(t, expected, methods, "GetOverriddenMethods test SubSub fail")

	methods = GetOverriddenMethods((*SubSub)(nil), (*Base)(nil), i)
	expected = []string{"Zoo", "Foo"}
	sort.Strings(methods)
	sort.Strings(expected)
	_ = assert.EqualValuesf(t, expected, methods, "GetOverriddenMethods test SubSub fail")
}
