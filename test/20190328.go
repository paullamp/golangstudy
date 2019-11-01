package main

import (
	"fmt"
	"os/exec"
)

func getMemory() {
	res := exec.Command("cat /proc/cpuinfo")
	buf, _ := res.Output()
	fmt.Printf("%s\n", string(buf))
}
