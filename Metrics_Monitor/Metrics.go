package main

import (
	"log"
	"os"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	// Monitoring interval in seconds
	monitoringInterval := 5

	// Start the monitoring loop
	for {
		// Get CPU usage
		cpuUsage, err := cpu.Percent(time.Second, false)
		if err != nil {
			log.Printf("Error getting CPU usage: %s", err.Error())
		} else {
			log.Printf("CPU Usage: %.2f%%", cpuUsage[0])
		}

		// Get memory usage
		memUsage, err := mem.VirtualMemory()
		if err != nil {
			log.Printf("Error getting memory usage: %s", err.Error())
		} else {
			log.Printf("Memory Usage: %.2f%%", memUsage.UsedPercent)
		}

		// Get disk usage
		rootPath := "/" // Change this to the appropriate disk path if needed
		diskUsage, err := disk.Usage(rootPath)
		if err != nil {
			log.Printf("Error getting disk usage: %s", err.Error())
		} else {
			log.Printf("Disk Usage: Total: %d GB, Used: %d GB, Free: %d GB, Usage: %.2f%%", diskUsage.Total/1024/1024/1024, diskUsage.Used/1024/1024/1024, diskUsage.Free/1024/1024/1024, diskUsage.UsedPercent)
		}

		// To fetch hostname
		hostname, _ := os.Hostname()
		log.Printf("Hostname: %s", hostname)

		// Sleep for the specified monitoring interval
		time.Sleep(time.Duration(monitoringInterval) * time.Second)
	}
}