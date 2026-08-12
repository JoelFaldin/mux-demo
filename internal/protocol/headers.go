package protocol

type Header struct {
	StreamID uint32
	Type     byte
	Length   uint32
}

type messageType int

const (
	FrameTypeNormal messageType = iota
	FrameTypeOpen
	FrameTypeClose
)
