package _2_builder

/*
Строитель — это порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово.
Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов.


Минусы паттерна:
Усложняет код программы из-за введения дополнительных классов;
Клиент будет привязан к конкретным структурам строителей.
*/

const (
	XIAOMI = "xiaomi"
	HONOR  = "honor"
)

// Структура создаваемых объектов
type modelPhone struct {
	CPU    int
	RAM    int
	MB     string
	Corpus string
}

type Collector interface {
	SetCPU()
	SetRAM()
	SetMB()
	SetCorpus()

	BuildPhone() modelPhone
}

// Метод GetCollector возвращает ссылку на конкретную структуру сборщика
func GetCollector(collectorType string) Collector {
	switch collectorType {
	case HONOR:
		return &HonorBuilder{}
	case XIAOMI:
		return &XiaomiBuilder{}
	default:
		return nil
	}
}
