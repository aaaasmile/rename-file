package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println("Rename files")
	// NOTE:folder_03 and folder_01 are synch with the sd-card of the player
	dirToScan := `D:\Projects\go-lang\rename-file\folder_07`
	files, err := ioutil.ReadDir(dirToScan)
	if err != nil {
		panic(err)
	}
	// The mp3 filename format for Yx5300 Chip is:
	// 001xxx.mp3
	// use the inial filename instead of xxx
	renum := regexp.MustCompile("^[0-9]+")
	for ix, ffInfo := range files {
		if !ffInfo.IsDir() {
			fmt.Println("File: ", ix, ffInfo.Name())
			src := filepath.Join(dirToScan, ffInfo.Name())

			part := ffInfo.Name()[:3]
			arr := renum.FindAllString(part, -1)
			if len(arr) == 1 && len(arr[0]) == 3 {
				fmt.Println("already processed item ", arr[0])
				part = ffInfo.Name()[3:6]
			} else if len(arr) == 1 && len(arr[0]) == 2 {
				fmt.Println("Number in the filename ", arr[0])
				arr := strings.Split(ffInfo.Name(), "-") // something like '11 - My song.mp3'
				if len(arr) > 1 {
					arr[1] = strings.Trim(arr[1], " ")
					part = arr[1][:3]
				} else {
					part = ffInfo.Name()[2:5]
				}
			}
			partascii := StringToAsciiBytes(part) // avoid accent characters trouble in filename
			fmt.Println("ascii: ", partascii)

			part = strings.ReplaceAll(string(partascii), " ", "x")
			dst := fmt.Sprintf("%03d%s.mp3", ix+1, part)
			dst = filepath.Join(dirToScan, dst)
			err = os.Rename(src, dst)
			if err != nil {
				panic(err)
			}

			fmt.Println("Rename : ", src, dst)
		}
	}

}

func StringToAsciiBytes(s string) []byte {
	t := make([]byte, utf8.RuneCountInString(s))
	i := 0
	for _, r := range s {
		t[i] = byte(r)
		if t[i] < 48 ||
			((t[i] > 57) && (t[i] < 65)) ||
			((t[i] > 90) && (t[i] < 97)) ||
			(t[i] > 122) {
			t[i] = 120 // if it is not a valid alpha numeric character (i.e. 'è'), then is x
		}
		i++
	}
	return t
}
