package monketype

import (
	"testing"
)

func acceptMonkeType(it MonkeType) {
	return
}

func Test_MonkeType(test *testing.T) {
	acceptMonkeType(Content{})
	acceptMonkeType(User{})
}
