package daemon

import (
	"kkokey/distribute-compute-module/dcm/core/types"

	"fmt"
	"runtime"
	"strings"
)

func Runner(msg chan string) {
	strBuilder := strings.Builder{}
	_, err := strBuilder.WriteString("daemon running... local ip addr: ")
	if err != nil {
		fmt.Print(err)
	}

	distributeModule := types.GetModule()
	_, err = strBuilder.WriteString(distributeModule.IPAddr.String())
	if err != nil {
		fmt.Print(err)
	}
	msg <- strBuilder.String()
}

func PrintMemUsage() {
	var size = "Kb"
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v %s", bToKb(m.Alloc), size)
	fmt.Printf("\tTotalAlloc = %v %s", bToKb(m.TotalAlloc), size)
	fmt.Printf("\tSys = %v %s", bToKb(m.Sys), size)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToKb(b uint64) uint64 {
	return b / 1024
}
