package prodservice

import "strconv"

type ProdModel struct {
	ProdID   int    `json:"pid"`
	ProdName string `json:pname`
}

func NewProd(id int, name string) *ProdModel {
	return &ProdModel{
		ProdID:   id,
		ProdName: name,
	}
}

func NewProdList(n int) []*ProdModel {
	ret := make([]*ProdModel, 0)
	for i := 0; i < n; i++ {
		ret = append(ret, NewProd(100+i, "prodname"+strconv.Itoa(i)))
	}
	return ret
}
