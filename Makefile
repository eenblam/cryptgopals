dirs=./encode ./analysis ./xor ./cipher ./rand

fmt:
	gofmt -w */*.go

test:
	go test -cover $(dirs)

testv:
	go test -v -cover $(dirs)

cover:
	go test -v -coverprofile=coverage.out $(dirs)
	go tool cover -html=coverage.out
	rm *.out

# Heat maps for coverage. Only need atomic if we start parallelizing things.
# https://blog.golang.org/cover
heat:
	go test -v -covermode=count -coverprofile=count.out $(dirs)
	go tool cover -html=count.out
	rm *.out

atomic:
	go test -v -covermode=atomic -coverprofile=atomic.out $(dirs)
	go tool cover -html=atomic.out
	rm *.out
