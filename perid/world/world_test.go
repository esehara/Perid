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

func TestEnvironmetSetEnvVlue (t *testing.T) {
	environ := SetUpTest()
	var usevalue EnvValue
	usevalue = EnvValue{
		Name: "foo",
		Val: "bar",
	}
	environ.Append(usevalue)
	if environ.Len() != 1 {
		t.Errorf("Enviroment array is not access. Enviroment has %d.", environ.Len())
	}
	ok, result := environ.Search("foo")
	if ok && result != "bar" {
		t.Errorf("Enviroment get not 'bar' value")
	}
}

func TestEnvironmentNotSetMultiple(t *testing.T) {
	environ := SetUpTest()
	var usevalue EnvValue
	usevalue = EnvValue{
		Name: "foo",
		Val: "bar",
	}
	environ.SetValue(usevalue)
	environ.SetValue(usevalue)
	if environ.Len() != 1 {
		t.Errorf("Enviroment requires Single Assignment")
	}
}
