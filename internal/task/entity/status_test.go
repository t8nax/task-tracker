package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseStatus_ReturnsStatus_WhenStatusIsValid(t *testing.T) {
	in := map[Status]string{StatusNone: "", StatusToDo: "todo", StatusInProgress: "in-progress", StatusDone: "done"}

	for expStatus, s := range in {
		status, err := ParseStatus(s)
		assert.Nil(t, err)
		assert.Equal(t, expStatus, status)
	} 
}

func TestParseStatus_ReturnsErr_WhenStatusIsInvalid(t *testing.T) {
	s := "invalid_status"

	status, err := ParseStatus(s)

	assert.Equal(t, StatusNone, status)
	assert.Error(t, err)
	assert.EqualError(t, err, GetErrInvalidStatus(s).Error())
}