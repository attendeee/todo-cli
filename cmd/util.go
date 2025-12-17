package cmd

import (
	"fmt"
	"task/repository"

	"gopkg.in/yaml.v3"
)

func MustMarshalFileToBytes(tasks []repository.Task) []byte {
	b, err := yaml.Marshal(tasks)
	if err != nil {
		panic(fmt.Errorf("Error: unable to marshal yaml file"))
	}

	return b

}
