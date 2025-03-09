package state

type Observer interface {
	Update(old, new Model)
}

type RuntimeObserver struct {
	update func(new Model, old Model)
}

func NewRuntimeObserver(update func(Model, Model)) *RuntimeObserver {
	return &RuntimeObserver{
		update: update,
	}
}

func (o *RuntimeObserver) Update(new Model, old Model) {
	o.update(new, old)
}
