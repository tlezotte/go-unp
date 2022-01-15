package main

import (
	"fmt"
	"os"

	"github.com/tlezotte/go-unp/compression"
	"github.com/tlezotte/go-unp/utils"
)

func main() {
	args := os.Args[1:]
	for _, compFile := range args {
		theFile := utils.GetExtension(compFile)
		switch theFile {
		case ".zip", ".egg", ".whl", ".jar":
			compression.RunUnzip(compFile)
		case ".gz":
			fmt.Println("gz")
		case ".tar":
			fmt.Println("tar")
		}
	}
}


//func gunzip() {
//
//func main() {
//	// uncompress data
//	uncompressedData, uncompressedDataErr := gUnzipData(compressedData)
//	if uncompressedDataErr != nil {
//		log.Fatal(uncompressedDataErr)
//	}
//
//	fmt.Println("uncompressed data:", uncompressedData)
//	fmt.Println("uncompressed data len:", len(uncompressedData))
//}
