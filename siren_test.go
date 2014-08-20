package siren

import (
	"encoding/json"
	"testing"
)

const (
	MarshalResponse = `{"class":["order"],"properties":{"itemCount":3,"orderNumber":42,"status":"pending"},"entities":[{"class":["info","customer"],"rel":["http://x.io/rels/customer"],"properties":{"customerId":"pj123","name":"Peter Joseph"},"links":[{"rel":["self"],"href":"http://api.x.io/customers/pj123"}]},{"class":["items","collection"],"rel":["http://x.io/rels/order-items"],"href":"http://api.x.io/orders/42/items"}],"actions":[{"name":"add-item","title":"Add Item","method":"POST","href":"http://api.x.io/orders/42/items","type":"application/x-www-form-urlencoded","fields":[{"name":"orderNumber","type":"hidden","value":"42"},{"name":"productCode","type":"text"},{"name":"quantity","type":"number"}]}],"links":[{"rel":["self"],"href":"http://api.x.io/orders/42"},{"rel":["previous"],"href":"http://api.x.io/orders/41"},{"rel":["next"],"href":"http://api.x.io/orders/43"}]}`
)

type Order struct {
	OrderNumber int    `json:"orderNumber" siren:"property"`
	ItemCount   int    `json:"itemCount" siren:"property"`
	Status      string `json:"status" siren:"property"`

	Customer Customer `json:"customer"`
}

type Customer struct {
	CustomerId string `json:"customerId" siren:"property"`
	Name       string `json:"name" siren:"property"`
}

func (self Order) MarshalSirenJSON() ([]byte, error) {
	orderProperties, err := ParseProperties(self)
	if err != nil {
		return []byte{}, err
	}

	customerProperties, err := ParseProperties(self.Customer)
	if err != nil {
		return []byte{}, err
	}

	entities := Entities{
		NewEntity(NewRel("http://x.io/rels/customer")).
			WithClass(NewClass("info", "customer")).
			WithProperties(customerProperties).
			WithLinks(NewLink(NewRel("self"), "http://api.x.io/customers/pj123")),
		NewLinkEntity(NewRel("http://x.io/rels/order-items"), "http://api.x.io/orders/42/items").
			WithClass(NewClass("items", "collection")),
	}

	actions := Actions{
		NewAction("add-item", "http://api.x.io/orders/42/items").
			WithTitle("Add Item").
			WithMethod("POST").
			WithContentType("application/x-www-form-urlencoded").
			WithFields(
			Fields{
				NewHiddenField("orderNumber").WithValue("42"),
				NewTextField("productCode"),
				NewNumberField("quantity"),
			}),
	}

	links := Links{
		NewLink(NewRel("self"), "http://api.x.io/orders/42"),
		NewLink(NewRel("previous"), "http://api.x.io/orders/41"),
		NewLink(NewRel("next"), "http://api.x.io/orders/43"),
	}

	doc := NewDocument().
		WithClass(NewClass("order")).
		WithProperties(orderProperties).
		WithEntities(entities).
		WithActions(actions).
		WithLinks(links)

	return json.Marshal(doc)
}

func TestMarshal(t *testing.T) {
	order := &Order{
		OrderNumber: 42,
		ItemCount:   3,
		Status:      "pending",
		Customer: Customer{
			CustomerId: "pj123",
			Name:       "Peter Joseph",
		},
	}

	b, err := Marshal(order)
	if err != nil {
		t.Fatalf("Marshalling failed with error %s", err.Error())
	}

	response := string(b)
	if response != MarshalResponse {
		t.Fatalf("Marshalling failed! Payload is different!")
	}
}
