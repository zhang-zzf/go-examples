package gostub

import (
	"github.com/prashantv/gostub"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

var aVar = 0

func TestGoStub(t *testing.T) {
	mockVal := 1
	defer gostub.
		Stub(&aVar, mockVal).
		// Stub() can be chained
		//Stub(&aVar, 2).
		Reset()
	// defer stub.Reset()
	assert.Equal(t, aVar, mockVal)
}

var aFunc = func(str string) string {
	return str
}

func TestGoStubFunc(t *testing.T) {
	assert.Equal(t, 0, aVar)
	mockVal := "MockVal"
	defer gostub.
		StubFunc(&aFunc, mockVal).
		Reset()
	assert.Equal(t, aFunc("anyValue"), mockVal)
}

func TestGoStubEnv(t *testing.T) {
	assert.Equal(t, "admin", os.Getenv("USER"))
	stubs := gostub.New()
	stubs.SetEnv("USER", "zzf")
	defer stubs.Reset()
	//
	assert.Equal(t, "zzf", os.Getenv("USER"))
	// test manual reset can reset env
	stubs.Reset()
	assert.Equal(t, "admin", os.Getenv("USER"))
}

type illegalArgumentError struct {
	msg string
}

func (r illegalArgumentError) Error() string {
	return r.msg
}

func TestStructMethod(t *testing.T) {
	msg := "Hello"
	err := &illegalArgumentError{msg}
	assert.Equal(t, msg, err.Error())
	defer gostub.StubFunc(err.Error, "mockVal").Reset()
	assert.Equal(t, "mockVal", err.Error())
}

func TestGoStubMethod(t *testing.T) {
	gostub.StubFunc(time.Now, time.Now())

}
