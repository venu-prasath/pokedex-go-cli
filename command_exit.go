package main

import "os"

func commandExit(_ *config, args ...string) error {
	os.Exit(0)
	return nil
}
