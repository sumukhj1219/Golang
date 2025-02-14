package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	fmt.Println("🖥️ Displaying system information:")
	fmt.Println("----------------------")
	fmt.Println("Operating System : ", runtime.GOOS)
	fmt.Println("Architecture : ", runtime.GOARCH)

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Hostname : ", hostname)
	}

	cpuInfo, err := cpu.Info()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Cpu : ", cpuInfo[0].ModelName)
		fmt.Println("Cores : ", len(cpuInfo))

	}

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Total RAM: %.2f GB\n", float64(vmStat.Total)/1e9)
		fmt.Printf("Available RAM: %.2f GB\n", float64(vmStat.Available)/1e9)
	}

	// hostStat, err := host.Info()
	// if(err != nil){
	// 	panic(err)
	// }else{
	// 	fmt.Printf("Uptime: %d seconds\n", hostStat.Uptime)
	// 	fmt.Printf("Platform: %s %s\n", hostStat.Platform, hostStat.PlatformVersion)
	// }

	fmt.Println("----------------------")
}
