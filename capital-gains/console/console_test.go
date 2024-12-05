package console

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

type mocks struct {
	processor *Mockprocessor
}

func createMocks(t *testing.T) (*console, *mocks) {
	ctrl := gomock.NewController(t)
	m := &mocks{
		processor: NewMockprocessor(ctrl),
	}
	l := NewConsole(m.processor)
	return l, m
}

func Test_NewConsole(t *testing.T) {
	console, mocks := createMocks(t)
	assert.NotNil(t, console, "NewConsole() should return a non-nil instance")
	assert.Equal(t, mocks, console.processor, "processor should be set correctly")
}
