package _2_builder

/*
Director позволяет лучше изолилировать код от клиента, его использование опционально
 */

type Director struct {
	collector Collector
}

func NewDirector(c Collector) *Director {
	return &Director{
		collector: c,
	}
}

// SetCollector меняет сборщик
func (d *Director) SetCollector(c Collector) {
	d.collector = c
}

//  BuildPhone собирает объект методами сборщика, решает их порядок
func (d *Director) BuildPhone() modelPhone {
	d.collector.SetCPU()
	d.collector.SetRAM()
	d.collector.SetMB()
	d.collector.SetCorpus()
	return d.collector.BuildPhone()
}
