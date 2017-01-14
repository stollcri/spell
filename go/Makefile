#
# spell -- Find the correct spelling for a bunch of letters (a word)
#  Copyright (c) 2017 Christopher Stoll (https://github.com/stollcri)
#

GOC = go build

CFLAGS = -race
AOFILE = ./bin/spell

CFLAGS_DBG = -race
AOFILE_DBG = $(AOFILE)

export GOPATH = $GOPATH:$(pwd)

.PHONY: default spell test debug run install

default: test
test: debug run

debug:
	@mkdir -p ./bin/
	$(GOC) $(CFLAGS_DBG) -o $(AOFILE_DBG)

run:
	time bin/spell synthesis Pokémon

spell:
	@mkdir -p ./bin/
	$(GOC) $(CFLAGS) -o $(AOFILE)

install:
	cp bin/spell /usr/local/bin/spell
	chown root:wheel /usr/local/bin/spell

	mkdir -p /usr/local/man/
	cp doc/spell /usr/local/man/spell.1
	chown root:wheel /usr/local/man/spell.1

	cp .spell ~/.spell
