package main

import "fmt"

type Payable interface {
	CalculatePay() float64
}

type SalariedEmployee struct {
	hour int64
	wages int64
}

type CommercialEmployee struct {
	
}

func (c *SalariedEmployee) CalculatePay() int64 {
	return c.wages * c.hour
}

func (c *SalariedEmployee) String() string {
	return fmt.Sprintf("Wages: %s", c.wages)
}

func main() {

}