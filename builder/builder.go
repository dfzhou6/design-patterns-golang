package builder

type Computer struct {
	CPU     string
	RAM     string
	Storage string
}

type ComputerBuilder interface {
	SetCPU(cpu string)
	SetRAM(ram string)
	SetStorage(storage string)
	GetComputer() *Computer
}

type MacComputerBuilder struct {
	computer *Computer
}

func NewMacComputerBuilder() *MacComputerBuilder {
	return &MacComputerBuilder{
		computer: &Computer{},
	}
}

func (impl *MacComputerBuilder) SetCPU(cpu string) {
	impl.computer.CPU = cpu
}

func (impl *MacComputerBuilder) SetRAM(ram string) {
	impl.computer.RAM = ram
}

func (impl *MacComputerBuilder) SetStorage(storage string) {
	impl.computer.Storage = storage
}

func (impl *MacComputerBuilder) GetComputer() *Computer {
	return impl.computer
}

type HuaweiComputerBuilder struct {
	computer *Computer
}

func NewHuaweiComputerBuilder() *HuaweiComputerBuilder {
	return &HuaweiComputerBuilder{
		computer: &Computer{},
	}
}

func (impl *HuaweiComputerBuilder) SetCPU(cpu string) {
	impl.computer.CPU = cpu
}

func (impl *HuaweiComputerBuilder) SetRAM(ram string) {
	impl.computer.RAM = ram
}

func (impl *HuaweiComputerBuilder) SetStorage(storage string) {
	impl.computer.Storage = storage
}

func (impl *HuaweiComputerBuilder) GetComputer() *Computer {
	return impl.computer
}
