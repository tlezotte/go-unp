package compression

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"

	"github.com/tlezotte/go-unp/utils"
)

func gUnzipData(data []byte) (resData []byte, err error) {
	b := bytes.NewBuffer(data)

	var r io.Reader
	r, err = gzip.NewReader(b)
	if err != nil {
		return
	}

	var resB bytes.Buffer
	_, err = resB.ReadFrom(r)
	if err != nil {
		return
	}

	resData = resB.Bytes()

	return
}

func RunGunzip(compFile string, padding int) {
	//output := utils.GetFilename(compFile)
	fmt.Printf("Extracting: %*s", padding, compFile)
	//err := UnzipSource(compFile, output)
	//if err != nil {
	//	//log.Fatal(err)
	//	//fmt.Printf("  %s %s\n", color.RedString("⚠"), err)
  //      fmt.Printf("  %s\n", utils.Danger("⚠"))
	//} else {
        fmt.Printf("  %s\n", utils.Info("✓"))
  //  }
}
