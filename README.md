# spell

A utility to find the spelling of a word you are thinking of but have no idea how to spell. Give it a bunch of letters and it will tell you which word you are most likely looking for.

This utility is slower than regular spelling suggestion programs, but that is because it is more inventive with its matching. If you are looking for something that is fast or sticks to giving more traditional results, then this is not the utlity for you.

```
$ spell amiright
airtight
$ spell "am i right"
millwright
$ spell asdasdfsadk
sassafrack
$ spell asdfasdf
headfast
$ spell diplococus
diplodocus
$ spell gfkjhgjkj
dogfought
$ spell lksdjfoij
ladyflies
$ spell scichiatrist
psychiatrist
$ spell qwerty
wefty
```

A rediculous method for generating random words near a certain length

```
spell $(cat /dev/urandom | env LC_CTYPE=C tr -dc 'a-zA-Z' | fold -w 8 | head -n 1)
```

```
$ spell $(cat /dev/random | env LC_CTYPE=C tr -dc 'a-zA-Z' | fold -w 6 | head -n 1)
biogen
$ spell $(cat /dev/random | env LC_CTYPE=C tr -dc 'a-zA-Z' | fold -w 8 | head -n 1)
studfish
$ spell $(cat /dev/random | env LC_CTYPE=C tr -dc 'a-zA-Z' | fold -w 10 | head -n 1)
battycake
$ spell $(cat /dev/random | env LC_CTYPE=C tr -dc 'a-zA-Z' | fold -w 12 | head -n 1)
hurtlessness
```

## TODO

- allow for multiple words to be returned
- improve results by adding probability checking (presently the last match is geven when there are multiple matches of the same quality)

Spelling wordlist from https://github.com/dwyl/english-words

## Go Version

tldr: for all intents and purposes it doesn't work

I wanted to implement something in Go, and since I already had the logic for this algorithm worked out, I thought it would make a good candidate. The Go version works, but it is so terribly slow. Some of the slowness could be due to the fact that I have to manually call the garbage collector every couple of itterations, and even that doesn't guarantee that the program will not crash due to some sort of memory related issue. As a new go progammer, I'm sure I did not properly optimize things. But, even if I could improve performance by an order of magnitude it still be considerably slower than the c version. This algorithm does a lot of work and even the c version takes almost 2 seconds to run given a word list with over 350,000 lines. I enjoyed writting the go version, and it was fairly easy to port over the c code, but in the future I will not be using Go for such intensive tasks.

One further note, the Go version handled UTF-8 while the c version does not.
