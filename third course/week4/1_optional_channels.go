/*
	let's assume we have 3 channels namely T1, T2 and T3

channel T3 computes the product of the numbers received form T1
and T2. hence both values are needed for T3 to carryout its operation
and we cannot proceed without them.
There are other cases in which T3 only needs data from one of the
channels to carryout its intended operation. For such cases we can
use the select statement as shown in the example below
*/
package main

import "fmt"

func prod(val1, val2 int, c chan int) {
	c <- val1 * val2
}

/*
*Abort Channel
The select case is also useful for a case in which we need to
receive data until a certain signal is passed/received (abort signal).

the example below shows how this may be used

for {
	select {
	case a := <- c:
		fmt.Println(a)
	case <- abort:
		return
	}
}

*/

func main() {
	var cah chan int = make(chan int)
	var ca chan int = make(chan int)

	go prod(3, 4, cah)
	go prod(4, 4, ca)

	select {
	case a := <-cah:
		fmt.Println(a)
	case b := <-ca:
		fmt.Println(b)
	}

	/*
	*NOTE: select can also be used for sending operations.
	*The example above only shows for 2 receiving operations
	 */

}
