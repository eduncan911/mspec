package examples

// this file holds the example models under test.
//

const normalColor = "brown"

type paint struct {
	color      string
	iswashable bool
}

type dog struct {
	color  string
	paint  *paint
	washed bool
}

func BirthDog() *dog {
	return &dog{
		color:  normalColor,
		paint:  nil,
		washed: false,
	}
}

func (d *dog) Paint(p *paint) {
	d.color = p.color
	d.paint = p
	d.washed = false
}

func (d *dog) Wash() {
	d.color = normalColor
	d.paint = nil
	d.washed = true

}
