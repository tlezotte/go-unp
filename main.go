package main

import (
	"fmt"
	"os"

	"github.com/tlezotte/go-unp/compression"
	"github.com/tlezotte/go-unp/utils"
)

func main() {
	var (
		args         = os.Args[1:]
		maxCharCount = 0
		finalList    = []string{}
		zipExt       = []string{".zip", ".jar", ".egg", ".whl"}
		gzExt        = []string{".gz", ".tgz"}
		tarExt       = []string{".tar"}
	)
	extList := append(zipExt, gzExt...)
	extList = append(extList, tarExt...)

	// Make slice of compressed files
	for _, filename := range args {
		theFile := utils.GetExtension(filename)

		if utils.Contains(extList, theFile) {
			fileLen := len(filename)
			if fileLen > maxCharCount {
				maxCharCount = fileLen
			}
			finalList = append(finalList, filename)
		}
	}

	// Uncompress files
	finalListLen := len(finalList)
	for i, filename := range finalList {
		spinner := fmt.Sprintf("[%d/%d]", i+1, finalListLen)
		fmt.Printf("%s ", utils.Fade(spinner))
		fileExt := utils.GetExtension(filename)

		switch {
		case utils.Contains(zipExt, fileExt):
			compression.RunUnzip(filename, maxCharCount)
		case utils.Contains(gzExt, fileExt):
			compression.RunGunzip(filename, maxCharCount)
		case utils.Contains(tarExt, fileExt):
			compression.RunUntar(filename, maxCharCount)
		}
	}
}
