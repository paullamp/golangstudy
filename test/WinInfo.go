package main

import (
	// "bytes"
	"fmt"
	"math"
	"net"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	// "github.com/golang/text/encoding/simplifiedchinese"
)

type Inters struct {
	Inter_Name string
	Inter_MAC  string
	Inter_IP   string
}
type SystemInfo struct {
	Sn_Number      string
	HostName       string
	DiskDeviceName string
	DiskDeviceSize float64
	OsInfo         string
	Memory         float64
	CpuType        string
	CpuNum         int
	NetInterface   []Inters
}

//采集系统SN
func (si *SystemInfo) GetSnNumber() {
	//1. 采集本机的SN号信息
	cmd := exec.Command("CMD", "/C", "wmic bios get SerialNumber /value")
	sn_number, err := cmd.Output()
	if err != nil {
		fmt.Println("err in run cmd.Output:", err)
		return
	}
	for _, value := range strings.Split(string(sn_number), "\n") {
		if strings.TrimSpace(value) != "" {
			si.Sn_Number = strings.Split(strings.TrimSpace(value), "=")[1]
			// fmt.Println(si.Sn_Number)
		}
	}
}

//采集网卡信息
func (si *SystemInfo) GetNetWork() {
	inets, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces错误:", err)
		return
	}

	//定义包含169.254.0.0/16位的网段
	_, nodhcp, _ := net.ParseCIDR("169.254.0.0/16")
	for _, value := range inets {
		v, _ := value.Addrs()

		if len(v) > 0 {
			if ipnet, ok := v[len(v)-1].(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				//去除非正常的ipv4地址，　去除所有dhcp错误的ip
				if ipnet.IP.To4() != nil && !nodhcp.Contains(ipnet.IP) {
					netint := Inters{}
					//网卡名，网卡mac地址
					// fmt.Println(value.Index)
					netint.Inter_Name = value.Name
					netint.Inter_MAC = value.HardwareAddr.String()
					netint.Inter_IP = v[len(v)-1].String()
					fmt.Println(value.Name)
					fmt.Println(value.HardwareAddr)
					fmt.Println("IPAddr:", v[len(v)-1])
					fmt.Println("-----------feng-ge-xian--------------")
					si.NetInterface = append(si.NetInterface, netint)
				}
			}
		}
	}
}

//获取操作系统信息，主机名，运行的平台(x86_64)
func (si *SystemInfo) GetSys() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("获取主机名错误:", err)
		return
	}
	si.HostName = hostname
	fmt.Println(si.HostName)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
}

func (hi *SystemInfo) ReadSysinfoFromFile() {
	//从文件读取cpu的配置信息
	f, err := os.Open("c:/sysinfo2019.txt")
	if err != nil {
		fmt.Println("打开文件出错", err)
		return
	}
	defer f.Close()
	buf := make([]byte, 4096)

	n, errRead := f.Read(buf)
	if errRead != nil {
		fmt.Println("读取文件内容出错:", errRead)
		return
	}
	//将信息从GB18030转码成UTF-8
	var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(buf[:n])
	resultString := string(decodeBytes)
	sysinfoSlice := strings.Split(resultString, "\n")
	//设置标志位，只读取一次OS 版本", 忽略"BIOS 版本"
	flag := false
	for key, value := range sysinfoSlice {
		if strings.TrimSpace(value) != "" {
			//读取主机名所在行
			if strings.Contains(value, "主机名") {
				hi.Hostname = strings.TrimSpace(strings.Split(value, ":")[1])
			}
			//读取OS信息以及明细版本行

			if strings.Contains(value, "OS 名称") {
				// fmt.Println("Get os 名称", value)
				osname := strings.TrimSpace(strings.Split(value, ":")[1])
				// fmt.Println("osname:", osname)
				hi.OsInfo += osname
			}
			if strings.Contains(value, "OS 版本") {
				// fmt.Println("Get os版本", value)
				osversion := strings.TrimSpace(strings.Split(value, ":")[1])
				// fmt.Println("osversion", osversion)
				if !flag {
					hi.OsInfo += osversion
					flag = true
				}

			}
			if strings.Contains(value, "物理内存总量") {
				memory := strings.TrimSpace(strings.Split(value, ":")[1])
				memoryStr := strings.Replace(strings.Split(memory, " ")[0], ",", "", -1)
				// mem, _ := strconv.Atoi(memoryStr)
				mem, _ := strconv.ParseFloat(memoryStr, 2)
				hi.Memory = math.Ceil(mem / 1024)

			}
			if strings.Contains(value, "处理器") {
				//此行获得的信息为： 处理器:           安装了 1 个处理器。
				//通过正则将获取数字：1
				re := regexp.MustCompile(".*([0-9]+).*")
				ss := re.FindAllStringSubmatch(value, 1)
				cpunum, _ := strconv.Atoi(ss[0][1])
				hi.CpuNum = cpunum
				//下一行为cpu的具体型号信息；格式为：[01]: Intel64 Family 6 Model 58 Stepping 9 GenuineIntel ~3901 Mhz
				cpuString := sysinfoSlice[key+1]
				hi.CpuType = strings.TrimSpace(strings.Split(cpuString, ":")[1])
			}
		}

	}
}

//获取操作系统信息，通过执行systeminfo命令， 然后解析输出
func (si *SystemInfo) GetSysInfoCmd() {
	cmd := exec.Command("CMD", "/C", "SYSTEMINFO > c:/sysinfo2019.txt")
	cmdResult, err := cmd.Output()
	if err != nil {
		fmt.Println("运行SYSTEMINFO命令出错:", err)
		return
	}
	fmt.Println(cmdResult)
	// fmt.Println(string(cmdResult))
	// data, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader(cmdResult), simplifiedchinese.GB18030.NewEncoder()))
	// fmt.Println(data)
	// fmt.Println(string(data))
}

//获取系统总磁盘大小
func (si *SystemInfo) GetDisk() {
	//通过执行命令：wmic diskdrive list brief , 并且解析输出
	diskcmd := exec.Command("CMD", "/C", "wmic diskdrive list brief")
	cmdResult, err := diskcmd.Output()
	if err != nil {
		fmt.Println("执行获取磁盘大小命令失败:", err)
		return
	}
	// fmt.Println(string(cmdResult))
	for _, value := range strings.Split(string(cmdResult), "\n") {
		if strings.TrimSpace(value) == "" || strings.Contains(strings.TrimSpace(value), "Caption") {
			continue
		}
		diskSlice := strings.Split(strings.TrimSpace(value), " ")
		diskDeviceName := diskSlice[0]
		diskDeviceSize, _ := strconv.ParseFloat(diskSlice[len(diskSlice)-1], 2)
		// fmt.Println(diskDeviceName, math.Ceil(diskDeviceSize/1024/1024/1024))
		si.DiskDeviceName = diskDeviceName
		si.DiskDeviceSize = math.Ceil(diskDeviceSize / 1024 / 1024 / 1024)
	}
}

func (si *SystemInfo) AllInfo() {
	/*
			type SystemInfo struct {
			Sn_Number      string
			HostName       string
			DiskDeviceName string
			DiskDeviceSize float64
			OsInfo         string
			Memory         float64
			CpuType        string
			CpuNum         int
			NetInterface   []Inters
		}
	*/
	fmt.Println("Sn_Number:", si.Sn_Number)
	fmt.Println("HostName:", si.HostName)
	fmt.Println("DiskDeviceName:", si.DiskDeviceName)
	fmt.Println("DiskDeviceSize:", si.DiskDeviceSize)
	fmt.Println("OsInfo:", si.OsInfo)
	fmt.Println("Memory:", si.Memory)
	fmt.Println("CpuType:", si.CpuType)
	fmt.Println("CpuNum:", si.CpuNum)
}

func main() {
	si := SystemInfo{}
	si.GetSnNumber()
	// fmt.Println(si)
	si.GetNetWork()
	si.GetSys()
	si.GetSysInfoCmd()
	si.GetDisk()
	si.ReadSysinfoFromFile()
	fmt.Println(si)

}
