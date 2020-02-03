package main

import (
	"reflect"
	"testing"
)

type interfaceA interface {
}

type A struct {
	interfaceA
}

func getInterface() interfaceA {
	var a *A = nil
	return a
}

func TestPrint(t *testing.T)  {
	t.Log(getInterface()==nil)
	t.Log("getInterface(): ", getInterface())
	b := getInterface()
	t.Log("type: ", reflect.TypeOf(b))
	t.Logf("%#v", getInterface())
}
