package models

import (
	"sync"
)

type Data struct {
	Data      map[uint32]*Stream
	DataMutex sync.Mutex
}

func NewData() *Data {
	return &Data{
		Data: make(map[uint32]*Stream),
	}
}

// Used in server
// Given streamId, searchs in models.Data map
// if theres any existent frame with that id.
func (d *Data) GetFrame(streamId uint32) *Stream {
	d.DataMutex.Lock()
	defer d.DataMutex.Unlock()

	find, ok := d.Data[streamId]
	if !ok {
		s := &Stream{
			StreamID: streamId,
			Chan:     make(chan []byte, 10),
		}
		d.Data[streamId] = s
		return s
	} else {
		return find
	}
}
