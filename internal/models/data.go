package models

import (
	"sync"
)

type Data struct {
	Data      map[uint32]chan []byte
	DataMutex sync.Mutex
}

// Used in server
// Given streamId, searchs in models.Data map
// if theres any existent frame with that id.
func (d *Data) SearchFrame(streamId uint32) (chan []byte, bool) {
	d.DataMutex.Lock()

	find, ok := d.Data[streamId]
	if !ok {
		return nil, false
	} else {
		return find, true
	}
}
