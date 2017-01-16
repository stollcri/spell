#
# spell -- Find the correct spelling for a bunch of letters (a word)
#  Copyright (c) 2017 Christopher Stoll (https://github.com/stollcri)
#

.PHONY: default spell test run install

default: debug
test:
	mix test

debug: spell run

run:
	time ./spell synthesis Pokémon

spell:
	mix escript.build

install:
	cp spell /usr/local/bin/spell
	chown root:wheel /usr/local/bin/spell

	cp .spell ~/.spell
