package main

import "fmt"

type taxSystem interface {
	calculateTax() int
}

type turkeyTax struct {
	taxPercentage int
	income        int
}

func (t *turkeyTax) calculateTax() int {
	tax := t.income * t.taxPercentage / 100
	return tax
}

type russiaTax struct {
	taxPercentage int
	income        int
}

func (r *russiaTax) calculateTax() int {
	tax := r.income * r.taxPercentage / 100
	return tax
}

type usaTax struct {
	taxPercentage int
	income        int
}

func (i *usaTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

func main() {
	turkeyTax := &turkeyTax{
		taxPercentage: 18,
		income:        2000,
	}

	russiaTax := &russiaTax{
		taxPercentage: 8,
		income:        3000,
	}

	usaTax := &usaTax{
		taxPercentage: 15,
		income:        4000,
	}
	taxSystems := []taxSystem{turkeyTax, russiaTax, usaTax}
	totalTax := calculateTotalTax(taxSystems)

	fmt.Printf("Total Tax is %d\n", totalTax)
}

func calculateTotalTax(taxSystem []taxSystem) int {
	totalTax := 0

	for _, t := range taxSystem {
		totalTax += t.calculateTax()
	}

	return totalTax
}
