/* pointer receivers are used to overcome a limitation of using
receiver types.
the limitation is that the argument or data is passed by call by value
hence the method cannot modify the argument or data.
NOTE
A common rule of thumb is
- While using pointer receivers use only that
- while using receiver types use only that

*/

package main

import "fmt"

type Point struct {
	x float64
	y float64
}

func (p *Point) offsetX(x1 float64) float64 {
	p.x += x1
	return p.x
}

func main() {
	p1 := Point{
		x: 2,
		y: 4,
	}

	fmt.Println(p1)
	fmt.Println(p1.offsetX(1))
	fmt.Println(p1)
}
