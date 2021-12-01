package tssr

import (
	"fmt"
)

const (
	EPSILON = 1e-8
)

type SR struct {
	Threshold float64
	WindowAmp int
	WindowLocal int
	NumEstpoints int
	NumGradPoints int
}


func NewSR(threshold float64, windowAmp int, windowLocal int, numEstPoints int, numGradPoints int) *SR {
	fmt.Println("hoge")
	fmt.Printf("%T\n", EPSILON)

	// ToDo: param validation
	
	return &SR{
		Threshold: threshold,
		WindowAmp: windowAmp,
		WindowLocal: windowLocal,
		NumEstpoints: numEstPoints, 
		NumGradPoints: numGradPoints,
	}
}


func (sr SR) InferThreshold(){

}