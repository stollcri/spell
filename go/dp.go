/**
 * spell -- utility suggest the correct spelling for words (groups of letters)
 *  https://github.com/stollcri/spell
 *
 *
 * Copyright (c) 2017, Christopher Stoll
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 * * Redistributions of source code must retain the above copyright notice, this
 *   list of conditions and the following disclaimer.
 *
 * * Redistributions in binary form must reproduce the above copyright notice,
 *   this list of conditions and the following disclaimer in the documentation
 *   and/or other materials provided with the distribution.
 *
 * * Neither the name of spell nor the names of its
 *   contributors may be used to endorse or promote products derived from
 *   this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
 * SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
 * CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
 * OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

package main

import (
	"fmt"
)

var DEBUG_PRINT_MATRIX int = 1
var DEBUG_PRINT_ARGPOS int = 0
var DEBUG_PRINT_ARGVAL int = 0
var DEBUG_PRINT_SCORES int = 1

var MOVE_COST int = 0
var SCORE_MATCH int = 10
var SCORE_WILDCARDint =  8
var SCORE_MISMATCH int = -4
var SCORE_SIMILARITY int = -2
var PENALTY_GAP int = -3
var PENALTY_MISMATCH int = -5
var PENALTY_TRANSPOSE int = -2

var EXPECTED_BOX_SCORE int =  7
var MINIMUM_GOOD_SCORE int = 75

func scoreFactor(wordA string, wordB string) float64 {
	result := 0.0
	lengthA := float64(len(wordA))
	lengthB := float64(len(wordB))
	if lengthA > lengthB {
		if lengthA != 0 {
			result = 1
		}
	} else if lengthA > lengthB {
		result = lengthB / lengthA
	} else {
		result = lengthA / lengthB
	}
	return result
}

func gappedScore(currentxchar string, currentychar string) int {
	return 1
}

func characterScore(currentxchar string, currentychar string) int {
	return 1
}

func transposescore(currentxchar string, currentychar string) int {
	return 1
}

func fillMatrix(xsize int, ysize int, xString string, yString string) []int {
	currentxchar := ""
	currentychar := ""
	gapscore := 0
	charscore := 0
	currentindex := 0
	indexdiag := 0
	indexleft := 0
	indexabove := 0
	dpmatrix := make([]int, xsize * ysize)
	if DEBUG_PRINT_MATRIX == 1 { fmt.Println("\n", xString, yString) }
	for y := 0; y < ysize; y++ {
		for x := 0; x < xsize; x++ {
			if x > 0 {
				currentxchar = string([]rune(xString)[x - 1])
			}
			if y > 0 {
				currentychar = string([]rune(yString)[y - 1])
			}

			// TODO: start both of the above loops at 1
			// 		 and remove this condition
			// 		 and remove the print condition below
			if (x > 0) && (y > 0) {
				gapscore = gappedScore(currentxchar, currentychar);
				charscore = characterScore(currentxchar, currentychar);

				currentindex = (y * xsize) + x;
				indexdiag = ((y - 1) * xsize) + (x - 1);
				indexleft = (y * xsize) + (x - 1);
				indexabove = ((y - 1) * xsize) + x;

				// the value of the current pixels is:
				// - the maximum of:
				//   - the diagonal value and the current score (there is a series of matches)
				//   - 0 (just in case all scores are below zero)
				//   - the above value and the gap score (there is a gap in the matches)
				//   - the left value and the gap score (there is a gap in the matches)
				// - minus the cost of movement
				dpmatrix[currentindex] = max4(
					dpmatrix[indexdiag] + charscore,
					0,
					dpmatrix[indexabove] + gapscore,
					dpmatrix[indexleft] + gapscore) + MOVE_COST;

				// values above and left are higher than diagonal, suggesting a transpose
				if ((dpmatrix[indexabove] > dpmatrix[indexdiag]) && (dpmatrix[indexleft] > dpmatrix[indexdiag])) {
					// give the diagonal the maximum value of the other neighbors
					dpmatrix[indexdiag] = max(
						dpmatrix[indexabove],
						dpmatrix[indexleft])
					// give current space the value of transposed caharacters, if less than current value
					dpmatrix[currentindex] = max(
						dpmatrix[indexdiag] + transposescore(currentxchar, currentychar),
						dpmatrix[currentindex])
				}
			}

			if DEBUG_PRINT_MATRIX == 1 {
				if ((x == 0) || (y == 0)) {
					if ((x == 0) && (y == 0)) {
						fmt.Printf("  ");
					}
					if ((x != 0) && (y == 0)) {
						fmt.Printf("__%v ", currentxchar);
					}
					if ((x == 0) && (y != 0)) {
						fmt.Printf("%v ", currentychar);
					}
				} else {
					fmt.Printf("%3d ", dpmatrix[currentindex]);
				}
			}
		}
		if DEBUG_PRINT_MATRIX == 1 { fmt.Println() }
	}
	return dpmatrix
}

func backTrack(dpmatrix []int, xsize int, ysize int, word string) float64 {
	return 1.0
}

func score(wordA string, wordB string) float64 {
	// add one to each dimension for the padding row/collumn
	xsize := len([]rune(wordA)) + 1
	ysize := len([]rune(wordB)) + 1
	dpmatrix := fillMatrix(xsize, ysize, wordA, wordB)
	btresults := backTrack(dpmatrix, xsize, ysize, wordA)
	if DEBUG_PRINT_SCORES == 1 { fmt.Printf("score: %#v %#v\n", btresults, wordB) }
	return btresults;
}

func bestMatch(word string, wordList []string) string {
	sfactor := 0.0
	dpscore := 0.0
	tmpscore := 0.0
	maxscore := 0.0
	wordmatch := ""
	for i := 0; i < len(wordList); i++ {
		if len(wordList[i]) > 0 {
			sfactor = scoreFactor(word, wordList[i])
			dpscore = score(word, wordList[i])
			tmpscore = dpscore * sfactor
			if (tmpscore >= maxscore) {
				wordmatch = wordList[i]
				maxscore = tmpscore
			}
		}
	}
	return wordmatch
}
