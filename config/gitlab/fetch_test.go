package gitlab_test

import (
	"testing"

	"web-layout/config/gitlab"
)

type TestConfig struct {
	A string   `env:"A,required"`
	B int      `env:"B" envDefault:"3000"`
	C []string `env:"C" envSeparator:":"`
}

func TestFetch(t *testing.T) {
	// TODO: set envs below, and then run the test
	//os.Setenv("GITLAB_TOKEN", "xxx")
	//os.Setenv("GITLAB_PID", "xxx")
	//os.Setenv("GITLAB_PATH", "xxx")
	c := &TestConfig{}

	if err := gitlab.Fetch(c); err != nil {
		t.Error(err)
		return
	}

	t.Log(c)
}
