fmt:
	gofmt -w */*.go

test:
	go test -v -cover ./encode ./freq ./xor

cover:
	go test -v -coverprofile=coverage.out ./encode ./freq ./xor
	go tool cover -html=coverage.out
	rm *.out

# Heat maps for coverage. Only need atomic if we start parallelizing things.
# https://blog.golang.org/cover
heat:
	go test -v -covermode=count -coverprofile=count.out ./encode ./freq ./xor
	go tool cover -html=count.out
	rm *.out

atomic:
	go test -v -covermode=atomic -coverprofile=atomic.out ./encode ./freq ./xor
	go tool cover -html=atomic.out
	rm *.out