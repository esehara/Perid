package world

import (
	"testing"
)

func SetUpTest() *envs {
	environ := AccessEnv()
	environ.UnsafeReset()
	return environ
}

func TestEnvironmentIsSingleton (t *testing.T) {
	environ := SetUpTest()
	if environ.Len() != 0 {
		t.Errorf("Enviroment singleton is not initialize")
	}
	
	environ.Append("foobar")
	if environ.Len() != 1 {
		t.Errorf("Enviroment array is not access")
	}

	another_environ := AccessEnv()
	if another_environ.Len() != 1 {
		t.Errorf("Enviroment is not singleton")
	}
}
