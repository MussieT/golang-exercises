package garbagecollector

// go run -gcflags -m main.go
// optimization, since each variable that escapes requires an extra memory allocation. (since it is in the heap - check by running the above command)
func One() {
	NewDuck()
}

type Duck struct {}
var global *int

// go:noinline
func NewDuck() *Duck{
	x := 1
	global = &x
	d := &Duck{}
	return d
}