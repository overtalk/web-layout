package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const file = "CONF_PATH"

func Fetch(c interface{}) error {
	path, isExist := os.LookupEnv(file)
	if !isExist {
		return fmt.Errorf("env `%s` is ablent", file)
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, c)
}
