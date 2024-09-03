package main

import "io"

func main() {

	var dd int = 17
	var mm int = 04
	var yy int = 2021
	var str string
	str = smt.sprint("%02d-%02d-%04d", dd, mm, yy)
	io.WriteString(os.stdout, str)

}
