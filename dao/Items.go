package dao

type Restaurant struct {
	name    string
	id      int64
	address string
}

type Desk struct {
	name   string
	id     int64
	active bool
}

func (d *Desk) GetName() string {
	return d.name
}

func (d *Desk) SetName(name string) {
	d.name = name
}
