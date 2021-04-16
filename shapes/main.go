package main

import "fmt"


type shape interface {
	getArea() float64
}


type triangle struct{}
type square struct{}


func main (){
	tr := triangle {}
	sq := square {}


	setArea(tr)
	setArea(sq)

}

func setArea(s shape)  {
	
	fmt.Println(s.getArea() )
}


func (triangle) getArea(base float64, height float64) float64 {

	return 0.5*base*height

}

func (square) getArea (sideLength float64)  float64 {

	return sideLength*sideLength
}  
	
