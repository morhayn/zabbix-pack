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
	// duration time last file in hours
	diffMin := int(time.Now().Sub(minTime).Hours())
	// duration time older file in hours
	diffMax := int(time.Now().Sub(maxTime).Hours())
	if diffMin > min {
		// last file older min value. Backup not create?
		fmt.Println("0")
	}
	if diffMax > max {
		// oldest backup file oldest max value. Its not deleted?
		fmt.Println("0")
	}
	// check size file
	// result processing
	// deciding on result
	fmt.Println("1")
}
