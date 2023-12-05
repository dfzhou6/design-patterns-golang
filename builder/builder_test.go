package builder

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	macBuilder := NewMacComputerBuilder()
	macBuilder.SetCPU("Apple M1")
	macBuilder.SetRAM("Samsung RAM")
	macBuilder.SetStorage("WestData Storage")
	mac := macBuilder.GetComputer()
	fmt.Printf("mac computer is %+v\n", *mac)

	huaweiBuilder := NewHuaweiComputerBuilder()
	huaweiBuilder.SetCPU("Huawei CPU")
	huaweiBuilder.SetRAM("Huawei RAM")
	huaweiBuilder.SetStorage("Huawei Storage")
	huawei := huaweiBuilder.GetComputer()
	fmt.Printf("huawei computer is %+v\n", *huawei)
}
