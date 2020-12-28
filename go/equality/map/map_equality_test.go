package map_test

import (
	"fmt"
	"reflect"
	"testing"
)

type example struct {
}

func TestValueEqual(t *testing.T) {
	m1 := map[string]interface{}{
		"key1": 1,
		"key2": "value2",
	}

	m2 := map[string]interface{}{
		"key1": 1,
		"key2": "value2",
	}

	if reflect.DeepEqual(m1, m2) {
		fmt.Println("reflect.DeepEqual(m1, m2) same")
	} else {
		fmt.Println("reflect.DeepEqual(m1, m2) different")
	}
}

func TestPointerEqual(t *testing.T) {
	m1 := map[string]interface{}{
		"key1": 1,
		"key2": "value2",
	}

	m2 := map[string]interface{}{
		"key1": 1,
		"key2": "value2",
	}

	if reflect.ValueOf(m1).Pointer() == reflect.ValueOf(m2).Pointer() {
		fmt.Println("pointer same")
	} else {
		fmt.Println("pointer different")
	}
}
