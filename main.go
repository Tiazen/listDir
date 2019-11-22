package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
)

var (
	d        = flag.String("d", ".", "Directory to process")
	a        = flag.Bool("a", false, "Print all info")
	h        = flag.Bool("h", false, "Print normal size")
	isSorted = flag.String("sort", "", "Sorts by chosen field")
)

func hrSize(fsize int64) string {
	measures := []string{"KB", "MB", "GB", "TB"}
	measure := 0
	if fsize < 1024 {
		fsize = 1
	}
	for i := fsize; i > 1023; i /= 1024 {
		measure += 1
		fsize = int64(math.Ceil(float64(fsize) / 1024))
	}
	return strconv.Itoa(int(fsize)) + measures[measure]
}

func printAll(file os.FileInfo, isH bool) {
	time := file.ModTime().Format("Jan 06 15:04")
	if isH {
		if int(file.Size()) > 1023 {
			fSize := strconv.Itoa(int(file.Size()) / 1024)
			fmt.Printf("%sKB %s %s \n", fSize, time, file.Name())
		} else {
			fSize := strconv.Itoa(int(file.Size()))
			fmt.Printf("%s %s %s \n", fSize, time, file.Name())
		}
	} else {
		fSize := strconv.Itoa(int(file.Size()))
		fmt.Printf("%s %s %s \n", fSize, time, file.Name())
	}
}

func sortByDate(files []os.FileInfo) []os.FileInfo {
	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime().Unix() > files[j].ModTime().Unix()
	})
	return files
}

func main() {
	flag.Parse()
	files, _ := ioutil.ReadDir(*d)
	if *isSorted == "date" {
		files = sortByDate(files)
	}
	for _, file := range files {
		if *h {
			printAll(file, true)
		} else if *a {
			printAll(file, false)
		} else {
			fmt.Println(file.Name())
		}
	}
}
