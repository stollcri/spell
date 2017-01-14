#
# spell -- Find the correct spelling for a bunch of letters (a word)
#  Copyright (c) 2017 Christopher Stoll (https://github.com/stollcri)
#

CC = cc

CFLAGS = -O3 -Wall -D NDEBUG -Wall
AOFILE = ./bin/spell
CFILES = spell.c dp.c

CFLAGS_DBG = -O0 -g -Wall
AOFILE_DBG = ./bin/spell
CFILES_DBG = $(CFILES)

# make sure scan-build is available
SCAN_BUILD_EXISTS := @true
SCAN_BUILD_WHICH := $(shell which scan-build)
ifeq ($(strip $(SCAN_BUILD_WHICH)),)
	SCAN_BUILD_EXISTS := @echo "\nERROR:\n scan-build not found.\n See: http://clang-analyzer.llvm.org/installation\n" && exit 1
endif

.PHONY: default spell test debug run analysis findtrouble clean install

default: test
test: debug run

debug:
	@mkdir -p ./bin/
	$(CC) $(CFLAGS_DBG) -o $(AOFILE_DBG) $(CFILES_DBG)

analysis:
	echo "Variables should be checked for null after *alloc ... check the following ..."
	grep -inlR -e alloc\( . | xargs grep -inR -A1 -e alloc\(
	echo
	echo "Variables should be nulled after free ... check the following output ..."
	grep -inlR -e free\( . | xargs grep -inR -A1 -e free\(
	echo
	$(SCAN_BUILD_EXISTS)
	scan-build -o tst/out -v -V make debug

findtrouble:
	grep -inR -e alloc\( .
	grep -inR -e strcpy\( -e strcat\( -e scanf\( -e gets\( .

run:
	time bin/spell synthesis

clean:
	rm -r tst/out/20*

spell:
	@mkdir -p ./bin/
	$(CC) $(CFLAGS) -o $(AOFILE) $(CFILES)

install:
	cp bin/spell /usr/local/bin/spell
	chown root:wheel /usr/local/bin/spell

	mkdir -p /usr/local/man/
	cp doc/spell /usr/local/man/spell.1
	chown root:wheel /usr/local/man/spell.1

	cp .spell ~/.spell
