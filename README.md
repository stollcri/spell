# spell

A utility to find the spelling of a word you are thinking of but have no idea how to spell. Give it a bunch of letters and it will tell you which word you are most likely looking for.

This utility is slower than regular spelling suggestion programs, but that is because it is more inventive with its matching. If you are looking for something that is fast or sticks to giving more traditional results, then this is not the utlity for you.

```
$ spell asdfasdf
headfast
$ spell butthole
bunghole
```

## TODO

- allow for multiple words to be returned
- improve results by adding probability checking (presently the last match is geven when there are multiple matches of the same quality)

Spelling wordlist from https://github.com/dwyl/english-words
