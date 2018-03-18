package main

import "path/filepath"

func relative(path1, path2 string) string {
	return filepath.Join(filepath.Dir(path1), path2)
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
