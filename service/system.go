package service

import (
	"encoding/json"
	"log"
	"time"

	"github.com/manojown/connector/model"
	cpu "github.com/shirou/gopsutil/v3/cpu"
	disk "github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

func CpuUsage() model.CPUStatus {
	var cpuStatus model.CPUStatus
	cpuStatus.Cores, _ = cpu.Counts(true)
	usage, err := cpu.Percent(0, false)
	cpuStatus.Usage = usage[0]
	if err != nil {
		log.Println("Issue while get Percentage.")
	}
	return cpuStatus
}

func Polling(config model.Config) {
	log.Println("Polling Started", config)
	var server model.Server
	var pollUrl string = config.URL + "/" + "connector"
	for {
		time.Sleep(10 * time.Second)
		server.Token = config.Token
		server.Port = config.Port
		server.CPU = CpuUsage()
		server.DiskSpace = DiskUsage()
		server.RAM = MemStat()
		server.LastConnected = time.Now().Unix()
		dataTosent, _ := json.Marshal(server)
		go APICall(pollUrl, "POST", dataTosent)
	}

	// APICall("http://localhost:8080/connector", "POST")
}

func DiskUsage() model.DiskStatus {
	var diskSpace model.DiskStatus
	d, _ := disk.Usage("/")
	diskSpace.All = d.Free
	diskSpace.Free = d.Total
	diskSpace.Used = d.Used
	return diskSpace
}

func MemStat() model.MemStatus {
	//Occupied by itself
	var memory model.MemStatus
	v, err := mem.VirtualMemory()

	if err == nil {
		memory.All = v.Total
		memory.Free = v.Free
		memory.Used = v.Used
	}
	return memory
}
