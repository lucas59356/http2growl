# Automação para compilação do programa

binpath = D:\Programas\bin

build: 
	go build -o main.exe

clean:
	rm -rf *.exe

win_install: build
	mv main.exe $(binpath)/http2growl.exe