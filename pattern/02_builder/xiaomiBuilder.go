package _2_builder

/*
XiaomiBuilder реализует интерфейс сборщика, собирая конкретные объекты
 */

type XiaomiBuilder struct {
	CPU    int
	RAM    int
	MB     string
	Corpus string
}

func NewXiaomiBuilder() Collector {
	return &XiaomiBuilder{}
}

func (x *XiaomiBuilder) SetCPU() {
	x.CPU = 2
}

func (x *XiaomiBuilder) SetRAM() {
	x.RAM = 4
}

func (x *XiaomiBuilder) SetMB() {
	x.MB = "microATX"
}

func (x *XiaomiBuilder) SetCorpus() {
	x.Corpus = "plastic"
}

func (x *XiaomiBuilder) BuildPhone() modelPhone {
	return modelPhone{
		CPU:    x.CPU,
		RAM:    x.RAM,
		MB:     x.MB,
		Corpus: x.Corpus,
	}
}
