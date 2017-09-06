# Automação para compilação do programa

binpath = D:\Programas\bin

build: 
	go build -i -v main.go mod_gntp.go mod_http.go

clean:
	rm -rf main.exe

win_install: build
	mv main.exe $(binpath)/http2growl.exe