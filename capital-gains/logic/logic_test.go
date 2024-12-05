package logic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewLogic(t *testing.T) {
	l := NewLogic()
	assert.NotNil(t, l, "NewLogic() should return a non-nil instance")
}
