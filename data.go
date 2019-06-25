package goaugmented

type dataPtr struct {
	Id  uint64
	Dat interface{}
}

func (d *dataPtr) ID() uint64 {
	return d.Id
}

func (d *dataPtr) Data() interface{} {
	return d.Dat
}

func (d *dataPtr) SetID(id uint64) {
	d.Id = id
}
