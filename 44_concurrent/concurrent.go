package _4_concurrent

// Race 竞争
func Race() {
	var a int
	go func() {
		a++
	}()
	if a == 0 {
		println(a)
	}
}
