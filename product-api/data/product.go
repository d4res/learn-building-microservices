package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
    ID          int     `json:"is"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float32 `json:"price"`
    SKU         string  `json:"sku"`
    // using - to complete ignore one filed in josn
    // another one is omitempty which will ignore the field when empty
    CreatedOn   string  `json:"-"`
    UpdatedOn   string  `json:"-"`
    DeletedOn   string    `json:"-"`
}


type Products []*Product 

func (p *Products) ToJSON(w io.Writer) error {
    e := json.NewEncoder(w)
    return e.Encode(p)
}

func GetProducts() Products {
    return productList
}

var productList = []*Product{
    {
        ID:          1,
        Name:        "Latte",
        Description: "Frothy milky coffee",
        Price:       2.45,
        SKU:         "abc323",
        CreatedOn:   time.Now().UTC().String(),
        UpdatedOn:   time.Now().UTC().String(),
    },
    {
        ID:          2,
        Name:        "Espresso",
        Description: "Short and strong coffee without milk",
        Price:       1.99,
        SKU:         "fjd34",
        CreatedOn:   time.Now().UTC().String(),
        UpdatedOn:   time.Now().UTC().String(),
    },
}
