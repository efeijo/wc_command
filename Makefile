build:
	go build -o cwc cmd/main.go

diff: build
	./cwc test.txt > cwc_out.txt
	wc test.txt > wc_out.txt
	diff cwc_out.txt wc_out.txt

clean:
	rm *_out.txt
	rm cwc