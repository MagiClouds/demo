package services

import (
	"fmt"
	"testing"
)

func TestInt32(t *testing.T)  {
	fmt.Println(int32(0))
}

func TestMap(t *testing.T)  {
	m := map[string]interface{}{"a":1,
		"b":"haha"}

	for k := range m {
		fmt.Println(k)
	}

}
