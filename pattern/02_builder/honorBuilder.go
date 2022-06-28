package _2_builder


/*
HonorBuilder реализует интерфейс сборщика, собирая конкретные объекты
*/


type HonorBuilder struct {
	CPU    int
	RAM    int
	MB     string
	Corpus string
}

func NewHonorBuilder() Collector {
	return &HonorBuilder{}
}

func (h *HonorBuilder) SetCPU() {
	h.CPU = 4
}

func (h *HonorBuilder) SetRAM() {
	h.RAM = 6
}

func (h *HonorBuilder) SetMB() {
	h.MB = "miniATX"
}

func (h *HonorBuilder) SetCorpus() {
	h.Corpus = "glass"
}

func (h *HonorBuilder) BuildPhone() modelPhone {
	return modelPhone{
		CPU:    h.CPU,
		RAM:    h.RAM,
		MB:     h.MB,
		Corpus: h.Corpus,
	}
}
