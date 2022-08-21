package model

import "go-marketplace/entity"

type ListProductData struct {
	Products []entity.Product `json:"products"`
}
type ListProductMeta struct {
	Page           int `json:"page"`
	MaxPage        int `json:"max_page"`
	ContentPerPage int `json:"content_per_page"`
	Total          int `json:"total"`
}
type ListProduct struct {
	Data ListProductData `json:"data"`
	Meta ListProductMeta `json:"meta"`
}

func (lpm ListProductMeta) MaxPageSlice() []map[string]int {
	var slice []map[string]int

	for i := 0; i < lpm.MaxPage; i++ {
		slice = append(slice, map[string]int{"num": i + 1})
	}

	return slice
}

func (lpm ListProductMeta) PagePrev() int {
	var page int
	if lpm.Page <= 1 {
		page = 1
	} else {
		page = lpm.Page - 1
	}
	return page
}

func (lpm ListProductMeta) PageNext() int {
	var page int
	if lpm.Page == lpm.MaxPage {
		page = lpm.Page
	} else {
		page = lpm.Page + 1
	}
	return page
}
