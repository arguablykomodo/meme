package main

import "path/filepath"

// Returns path if absolute, if else then resolves correct path
func resolvePath(path, source string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(source, path)
}

// Util function for keeping things DRY
func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
