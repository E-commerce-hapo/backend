package pm

import (
	"github.com/E-commerce-hapo/backend/pkg/event/dispatcher"
)

type ProductManager struct {
	dispatcher *dispatcher.Dispatcher
	//cateQ      *categorying.CategoryQuery
}

//func New(cateQ *categorying.CategoryQuery, dispatcher *dispatcher.Dispatcher) *ProductManager {
//	m := &ProductManager{}
//	m.cateQ = cateQ
//	m.dispatcher = dispatcher
//	err := dispatcher.AddEventListner(m, category_domain.CreatedCategoryEvent{})
//	if err != nil {
//		log.Fatal(err)
//	}
//	return m
//}
//
//func (u *ProductManager) Listen(event interface{}) {
//	switch event := event.(type) {
//	case category_domain.CreatedCategoryEvent:
//		fmt.Println("CreatedCategoryEvent: ", event.ID, event.Time)
//	default:
//		log.Printf("registered an invalid user event: %T\n", event)
//	}
//}
