package utils

import (
	"fmt"
	"testing"
)

func TestInitCPU(t *testing.T) {
	if cpu, err := InitCPU(); err != nil {
		t.Error(err)
	} else {
		fmt.Println("cpu:", cpu)
	}
	if ram, err := InitRAM(); err != nil {
		t.Error(err)
	} else {
		fmt.Println("ram:", ram)
	}
	if disk, err := InitDisk(); err != nil {
		t.Error(err)
	} else {
		fmt.Println("disk:", disk)
	}
}
