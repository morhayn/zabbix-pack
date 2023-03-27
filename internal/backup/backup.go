package backup

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// Check time create backup file. old file and new file in range min and max variable
func CheckTime(min, max int, path string) {
	//Dir exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("0")
		return
	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("0")
		return
	}
	maxTime := files[0].ModTime()
	minTime := files[0].ModTime()
	// size := files[0].Size()
	// countInDir := len(files)
	for _, file := range files {
		if maxTime.After(file.ModTime()) {
			maxTime = file.ModTime()
		}
		if minTime.Before(file.ModTime()) {
			minTime = file.ModTime()
		}
		// size := file.Size()
	}
	// time last file
	diffMin := int(time.Now().Sub(minTime).Hours())
	//  time older file
	diffMax := int(time.Now().Sub(maxTime).Hours())
	if diffMin > min {
		//  file older min value. Backup not create?
		fmt.Println("0")
	}
	if diffMax > max {
		// oldest backup file not deleted
		fmt.Println("0")
	}

	// check size file
	// result processing
	// deciding on result
	fmt.Println("1")
}
