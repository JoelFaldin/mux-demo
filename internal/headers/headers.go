package headers

type Header struct {
	StreamID uint32
	Type     byte
	Length   uint32
}

type messageType int

const (
	normal messageType = iota
	open
	close
)
