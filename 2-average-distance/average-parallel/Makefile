CC=gcc
FLAGS=-pthread
OUT=amd

all: amd

amd: amd.c
	$(CC) $(FLAGS) -o $(OUT) amd.c

debug: amd.c
	$(CC) $(FLAGES) -o $(OUT).debug amd.c -DDEBUG=1 -Wall

clean:
	rm -f $(OUT) $(OUT).debug
