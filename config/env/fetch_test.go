package env_test

import (
	"os"
	"testing"

	"web-layout/config/env"
)

type TestConfig struct {
	A string   `env:"A,required"`
	B int      `env:"B" envDefault:"3000"`
	C []string `env:"C" envSeparator:":"`
}

func TestFetch1(t *testing.T) {
	os.Setenv("A", "aaaa")
	os.Setenv("B", "12")

	c := &TestConfig{}

	if err := env.Fetch(c); err != nil {
		t.Error(err)
		return
	}

	t.Log(c)
}

func TestFetch2(t *testing.T) {
	os.Setenv("A", "aaaa")
	os.Setenv("C", "a:b:C:D")

	c := &TestConfig{}

	if err := env.Fetch(c); err != nil {
		t.Error(err)
		return
	}

	t.Log(c)
}
