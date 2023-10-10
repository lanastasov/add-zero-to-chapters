package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// https://github.com/lanastasov/copy-file-names-empty
// https://github.com/lanastasov/add-zero-to-chapters
// https://github.com/lanastasov/add-zero-to-chapters-subfolder
// https://github.com/lanastasov/generate-mp4-video-index
// 
func main() {
	fmt.Println("--add-zero-to-chapters.go--ver-0.0.4")

	// re1 := regexp.MustCompile(`\D\d{1}\.\s`)
	// re2 := regexp.MustCompile(`\D\d{2}\.\s`)
	// [^Mm] -> lesson1.abceede 1.MP4

	re1 := regexp.MustCompile(`\D\d{1}\.[^Mm]`)
	re2 := regexp.MustCompile(`\D\d{2}\.[^Mm]`)
	re3 := regexp.MustCompile(`^\d{1}\.`)
	re4 := regexp.MustCompile(`^\d{2}\.`)
	re5 := regexp.MustCompile(`^\d{1}\D`)
	re6 := regexp.MustCompile(`^\d{2}\D`)
	re7 := regexp.MustCompile(`\D\d{1}\.`)
	re8 := regexp.MustCompile(`\D\d{2}\.`)
	re9 := regexp.MustCompile(`\d{3}\.\d\s`)

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
			_ = os.Rename(files[i], "00"+files[i])
		}

		v := re1.FindString(files[i])
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

		v = re3.FindString(files[i])
		if v != "" {
			v = strings.TrimLeft(v, string(v[0]))
			v = strings.TrimRight(v, string(v[len(v)-1]))
			_ = os.Rename(files[i], strings.Replace(files[i], v, "00"+v, 1))
		}

		v = re4.FindString(files[i])
		if v != "" {
			_ = os.Rename(files[i], strings.Replace(files[i], v, "0"+v, 1))
		}

		v = re5.FindString(files[i])
		if v != "" {
			v = strings.TrimRight(v, string(v[len(v)-1]))
			_ = os.Rename(files[i], strings.Replace(files[i], v, "00"+v+". ", 1))
		}

		v = re6.FindString(files[i])
		if v != "" {
			v = strings.TrimRight(v, string(v[len(v)-1]))
			_ = os.Rename(files[i], strings.Replace(files[i], v, "0"+v+". ", 1))
		}

		v = re7.FindString(files[i])
		if v != "" {
			v = strings.TrimLeft(v, string(v[0]))
			v = strings.TrimRight(v, string(v[len(v)-1]))
			_ = os.Rename(files[i], strings.Replace(files[i], v, " 00"+v, 1))
		}

		v = re8.FindString(files[i])
		if v != "" {
			v = strings.TrimLeft(v, string(v[0]))
			v = strings.TrimRight(v, string(v[len(v)-1]))
			_ = os.Rename(files[i], strings.Replace(files[i], v, " 0"+v, 1))
		}

		// 038.3 filename.mp4 -> 038.03 filename.mp4
		v = re9.FindString(files[i])
		if v != "" {
			v = strings.TrimLeft(v, string(v[0]))
			v = strings.TrimRight(v, string(v[len(v)-1]))
			_ = os.Rename(files[i],strings.Replace(files[i], v, strings.Replace(v,".",".0",1), 1))
		}
	}
	var input string
	fmt.Println("Enter Text Below:")
	fmt.Scanln(&input)
	fmt.Println(input)
}
