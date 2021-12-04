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

// ToDo: online mode. what number of window shift width

func (sr SR) InferThreshold(){

}

func (sr SR) saliencyMap() {

}

func (sr SR) computeGrads() {

}

func (sr SR) add_est_points() {

}

func (sr SR) Score() {

}

