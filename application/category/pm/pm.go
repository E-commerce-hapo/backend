package pm

import (
	"log"
)

type ProcessManager struct {
}

func New() *ProcessManager {
	return &ProcessManager{}
}

func (u ProcessManager) Listen(event interface{}) {
	switch event := event.(type) {
	default:
		log.Printf("registered an invalid user event: %T\n", event)
	}
}
