package _3_visitor

/*
Объекты реализующие интерфейс shape для связи с visitor
*/

type square struct {
	side int
}

func (s *square) accept(v visitor) {
	v.visitForSquare(s)
}

type circle struct {
	radius int
}

func (c *circle) accept(v visitor) {
	v.visitForCircle(c)
}

type rectangle struct {
	l int
	b int
}

func (t *rectangle) accept(v visitor) {
	v.visitForRectangle(t)
}
