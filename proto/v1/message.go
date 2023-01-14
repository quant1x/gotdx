package v1

type Message interface {
	Serialize() ([]byte, error)
	UnSerialize(head interface{}, in []byte) error
}
