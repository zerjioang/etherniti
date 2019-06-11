package optional

type Trigger func()
type functionalModel uint8

const (
	undefined functionalModel = iota
	equal
)

type Optional struct {
	in   []interface{}
	a    Trigger
	b    Trigger
	mode functionalModel
}

func (o Optional) Map(t Trigger) Optional {
	//save first path
	o.a = t
	return o
}
func (o Optional) OrElse(t Trigger) Optional {
	o.b = t

	switch o.mode {
	case equal:
		if o.in[0] == o.in[1] {
			o.a()
		} else {
			o.b()
		}
	case undefined:
	default:
	}
	return o
}

func Equal(values ...interface{}) Optional {
	return Optional{in: values, mode: equal}
}
