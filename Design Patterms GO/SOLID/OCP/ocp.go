package main

import "fmt"

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (spec *ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == spec.color
}

type SizeSpecification struct {
	size Size
}

func (spec *SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == spec.size
}

func SpecFilter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}
func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, small}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	fmt.Printf("Green Products (OLD): \n")
	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	fmt.Print("Green Products (NEW): \n")
	greenSpec := ColorSpecification{green}
	for _, v := range SpecFilter(products, &greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}
}
