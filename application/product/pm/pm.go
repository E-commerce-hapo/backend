package pm

import (
	"fmt"
	"log"

	categorying "github.com/kiem-toan/application/category"
	"github.com/kiem-toan/pkg/event/dispatcher"

	category_domain "github.com/kiem-toan/domain/service/category"
)

type ProductManager struct {
	dispatcher *dispatcher.Dispatcher
	cateQ      *categorying.CategoryQuery
}

func New(cateQ *categorying.CategoryQuery, dispatcher *dispatcher.Dispatcher) *ProductManager {
	m := &ProductManager{}
	m.cateQ = cateQ
	m.dispatcher = dispatcher
	err := dispatcher.AddEventListner(m, category_domain.CreatedCategoryEvent{})
	if err != nil {
		log.Fatal(err)
	}
	return m
}

func (u *ProductManager) Listen(event interface{}) {
	switch event := event.(type) {
	case category_domain.CreatedCategoryEvent:
		fmt.Println("CreatedCategoryEvent: ", event.ID, event.Time)
	default:
		log.Printf("registered an invalid user event: %T\n", event)
	}
}
