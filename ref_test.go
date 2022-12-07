package ref_test

import (
	"testing"

	"github.com/0x5ab/ref"
	"github.com/stretchr/testify/assert"
)

type Test struct {
	TestVar int
}

func TestRef(t *testing.T) {
	assert.Equal(t, 3, *ref.V(3))
	assert.Equal(t, "hi", *ref.V("hi"))
	assert.Equal(t, 1.0, *ref.V(1.0))
	assert.Equal(t, Test{TestVar: 3}, *ref.V(Test{TestVar: 3}))

	assert.Equal(t, 3, ***ref.V(ref.V(ref.V(3))))
	assert.Equal(t, Test{TestVar: 3}, **ref.V(&Test{TestVar: 3}))

	s := &Test{TestVar: 10}
	assert.Equal(t, s, *ref.V(s))
}
