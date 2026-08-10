package models

import (
	"sync"
)

type Data struct {
	Data      map[uint32]chan []byte
	DataMutex sync.Mutex
}

func NewData() *Data {
	return &Data{
		Data: make(map[uint32]chan []byte),
	}
}

// Used in server
// Given streamId, searchs in models.Data map
// if theres any existent frame with that id.
func (d *Data) GetFrame(streamId uint32) chan []byte {
	d.DataMutex.Lock()
	defer d.DataMutex.Unlock()

	find, ok := d.Data[streamId]
	if !ok {
		ch := make(chan []byte, 10)
		d.Data[streamId] = ch
		return ch
	} else {
		return find
	}
}
