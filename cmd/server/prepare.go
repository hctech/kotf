package main

import (
	"github.com/KubeOperator/kotf/pkg/constant"
	"os"
)

func prepareStart() error {
	funcs := []func() error{
		makeDataDir,
	}
	for _, f := range funcs {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}

func makeDataDir() error {
	err := os.MkdirAll(constant.ProjectDir, 0755)
	if err != nil {
		return err
	}
	return nil
}
