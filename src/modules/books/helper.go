package books

import "server-go/src/share/exception"

func TestPanic() {
	panic(exception.BadRequestError("TEST", "TEST"))
}
