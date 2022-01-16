package compression

import (
	"fmt"

	"github.com/tlezotte/go-unp/utils"
)

func RunUntar(compFile string, padding int) {
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
