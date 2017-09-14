package serial

type TypedMessageMarshalI interface {
	MarshalText(target *TypedMessage) (text []byte, err error)
	UnmarshalText(text []byte, target *TypedMessage) error
}

var typedMessageMarshaler TypedMessageMarshalI

func SetTypedMessageMarshaler(tmmi TypedMessageMarshalI) {
	typedMessageMarshaler = tmmi
}

func (v *TypedMessage) MarshalText() (text []byte, err error) {
	return typedMessageMarshaler.MarshalText(v)
}

func (v *TypedMessage) UnmarshalText(text []byte) error {
	return typedMessageMarshaler.UnmarshalText(text, v)
}
