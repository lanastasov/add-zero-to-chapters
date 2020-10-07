package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// incorrect
// 15.1 2. Help the Cashier (Conditional & looping Section).zip
// 15.1  002. Help the Cashier (Conditional & looping Section).zip

// correct
//15.1 5.1 The Island Treasure.zip.zip
//15.1 5.1 The Island Treasure.zip.zip

func main() {
	fmt.Println("--add-zero-to-chapters.go--")

	re := regexp.MustCompile(`\D\d{1}\.\s`)
	re2 := regexp.MustCompile(`\D\d{2}\.\s`)

	files, _ := filepath.Glob("*.*")
	for i := range files {
		if (files[i][0] == 49 && (files[i][1] == 46 || files[i][1] == 45)) ||
			(files[i][0] == 50 && (files[i][1] == 46 || files[i][1] == 45)) ||
			(files[i][0] == 51 && (files[i][1] == 46 || files[i][1] == 45)) ||
			(files[i][0] == 52 && (files[i][1] == 46 || files[i][1] == 45)) ||
			(files[i][0] == 53 && (files[i][1] == 46 || files[i][1] == 45)) ||
			(files[i][0] == 54 && (files[i][1] == 46 || files[i][1] == 45)) ||
			(files[i][0] == 55 && (files[i][1] == 46 || files[i][1] == 45)) ||
			(files[i][0] == 56 && (files[i][1] == 46 || files[i][1] == 45)) ||
			(files[i][0] == 57 && (files[i][1] == 46 || files[i][1] == 45)) {
			_ = os.Rename(files[i], "0"+files[i])
		}

		v := re.FindString(files[i])
		if v != "" {
			v = strings.TrimLeft(v, string(v[0]))
			v = strings.TrimRight(v, string(v[len(v)-1]))
			_ = os.Rename(files[i], strings.Replace(files[i], v, " 00"+v, 1))
		}

		v = re2.FindString(files[i])
		if v != "" {
			v = strings.TrimLeft(v, string(v[0]))
			v = strings.TrimRight(v, string(v[len(v)-1]))
			_ = os.Rename(files[i], strings.Replace(files[i], v, " 0"+v, 1))
		}

	}
}
