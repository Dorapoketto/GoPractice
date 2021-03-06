package try

import (
	"fmt"
	"os"
	"runtime"
)

func getOsInfo() {
	var goos string = runtime.GOOS
	fmt.Printf("The operatoring system is %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}
