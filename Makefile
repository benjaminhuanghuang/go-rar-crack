unziptest:
	./unar -f -o test -p 123456 50.rar

crack:
	go run ./main/main.go
	