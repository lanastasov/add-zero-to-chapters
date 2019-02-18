package main

import (
	"os"
	"path/filepath"
)

func main() {
	files, _ := filepath.Glob("*.*")
	for i := range files {
		if (files[i][0] == 49 && files[i][1] == 46) ||
			(files[i][0] == 50 && files[i][1] == 46) ||
			(files[i][0] == 51 && files[i][1] == 46) ||
			(files[i][0] == 52 && files[i][1] == 46) ||
			(files[i][0] == 53 && files[i][1] == 46) ||
			(files[i][0] == 54 && files[i][1] == 46) ||
			(files[i][0] == 55 && files[i][1] == 46) ||
			(files[i][0] == 56 && files[i][1] == 46) ||
			(files[i][0] == 57 && files[i][1] == 46) {
			_ = os.Rename(files[i], "0"+files[i])
		}
	}

	files, _ = filepath.Glob("*.*")
	for i := range files {
		if filepath.Ext(files[i]) == ".vtt" {
			os.Mkdir("VTT", 077)
			break
		}
	}

	for i := range files {
		if filepath.Ext(files[i]) == ".vtt" {
			_ = os.Rename(files[i], "VTT/"+files[i])
		}
	}

}
