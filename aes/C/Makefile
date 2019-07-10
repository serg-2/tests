# My First Makefile

all: encrypt_decrypt.o
	gcc -o result encrypt_decrypt.o -lssl -lcrypto

encrypt_decrypt.o: encrypt_decrypt.c
	gcc -c encrypt_decrypt.c

clean:
	rm -rf *.o
	rm -rf result
