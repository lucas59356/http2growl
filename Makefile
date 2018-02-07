# Automação para compilação do programa

binpath = D:\Programas\bin

build: 
	go build

clean:
	rm -rf *.exe

win_install: build
	mv http2growl.exe $(binpath)/http2growl.exe

install: build
	install http2growl /usr/bin/

clean:
	rm http2growl http2growl.exe
