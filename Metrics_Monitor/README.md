Monitor PC/Server Metrics Using Golang:
Get CPU, Disk and Memory usage

Line 1: This line declares that this is the main package, which is required for an executable Go program.
Lines 3-10: These lines import the necessary packages (log, os, time, github.com/shirou/gopsutil/cpu, github.com/shirou/gopsutil/disk, and github.com/shirou/gopsutil/mem) that are used in the code for monitoring CPU, memory, and disk usage.
Line 14: The main function is the entry point of the program. It runs when the program starts.

Line 17: monitoringInterval is a variable that defines the time interval (in seconds) between each monitoring iteration.

Line 20-42: This block of code is the main monitoring loop. It runs indefinitely and repeatedly collects and logs the system metrics.

Line 23-28: The CPU usage is obtained using the cpu.Percent() function from the github.com/shirou/gopsutil/cpu package. It returns the percentage of CPU usage for each CPU core. If there is an error, it is logged. Otherwise, the CPU usage is logged.

Line 31-36: The memory usage is obtained using the mem.VirtualMemory() function from the github.com/shirou/gopsutil/mem package. It returns a mem.VirtualMemoryStat struct containing various memory usage statistics. If there is an error, it is logged. Otherwise, the memory usage percentage is logged.

Line 39-46: The disk usage is obtained using the disk.Usage() function from the github.com/shirou/gopsutil/disk package. It returns a disk.UsageStat struct containing disk usage statistics for the specified path (in this case, the root path "/"). If there is an error, it is logged. Otherwise, the total disk space, used space, free space, and usage percentage are logged.

Line 49-50: The hostname of the system is obtained using the os.Hostname() function from the os package. It fetches the hostname and assigns it to the hostname variable. It is then logged.

Line 53: The program sleeps for the specified monitoring interval using time.Sleep() to pause execution before the next iteration of the monitoring loop.

This code continuously monitors the CPU usage, memory usage, disk usage, and hostname of the system at the specified interval and logs the metrics.

go mod init metrics-monitor - Initializes a new Go module in the current directory
go run Metrics.go to run and see the live metrics of CPU, Disk and memory
go build Metrics.go to generate an executable file.

