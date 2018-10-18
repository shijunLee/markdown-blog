package common

import (
	"fmt"
	"testing"
)

func TestGetPost(t *testing.T) {
	a := GetPost("test")
	fmt.Println(a)
}
