package leak

/*
If we were using an unbuffered pipe, the two slower go routines would be blocked trying to
send their responses to a pipe from which no one would ever try to read data.
This situation, which is called a go subroutine leak, would be a bug. Unlike variables,
lost go routines are not automatically collected by the garbage collector, so it is important
to ensure that go routines must terminate on their own when they are no longer needed.

Solve - use buffered chan
responses := make(chan string, 3)
*/

func request(region string) string {
	return "result"
}

func mirroredQuery() string {
	responses := make(chan string)
	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()
	return <-responses // first request value
}
