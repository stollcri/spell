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
	"math"
	"unicode/utf8"
	"runtime"
)

var DEBUG_PRINT_MATRIX int = 0
var DEBUG_PRINT_ARGPOS int = 0
var DEBUG_PRINT_ARGVAL int = 0
var DEBUG_PRINT_SCORES int = 0

var MOVE_COST int = 0
var SCORE_MATCH int = 10
var SCORE_WILDCARD int = 8
var SCORE_MISMATCH int = -4
var SCORE_SIMILARITY int = -2
var PENALTY_GAP int = -3
var PENALTY_MISMATCH int = -5
var PENALTY_TRANSPOSE int = -2

var EXPECTED_BOX_SCORE int =  7
var MINIMUM_GOOD_SCORE int = 75

var CHAR_MARIX [93]string

func initCharMatrix() {
// 9 -- same value
// 8 -- difference in case (e.g. a, A -or- b, B)
// 7 --
// 6 -- next to on the keyboard (e.g. q, w -or- q, a)
// 5 -- looks similar (e.g. 1, l, I -or- 0, o, O)
// 4 -- diagonal on the keyboard (e.g. a, z -or- s, x)
// 3 -- can sound similar (e.g. c, k -or- f, ph)
// 2 -- difference in shift state (e.g. 1, ! -or- 2, @)
// 1 -- looks close (e.g. t, f -or- B, D)
// 0 -- no similarity
//                !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz{|}
 CHAR_MARIX[0] = "900000000000000020000000000000000000000000000000000000000000000000000000000000000000000000000" // !
 CHAR_MARIX[1] = "090000800000000000000000000000000000000000000000000000000000000800000000000000000000000000000" // "
 CHAR_MARIX[2] = "009000000000000000200000000000000000000000000000000000000000000000000000000000000000000000000" // #
 CHAR_MARIX[3] = "000900000000000000020000000000000000000000000000000000000000000000000000000000000000000000000" // $
 CHAR_MARIX[4] = "000090000000000000002000000000000000000000000000000000000000000000000000000000000000000000000" // %
 CHAR_MARIX[5] = "000009000000000000000020000000000000000000000000000000000000000000000000000000000000000000000" // &
 CHAR_MARIX[6] = "080000900000000000000000000000000000000000000000000000000000000800000000000000000000000000000" // '
 CHAR_MARIX[7] = "000000090000000000000000200000000000000000000000000000000020000000000000000000000000000000000" // (
 CHAR_MARIX[8] = "000000009000000200000000000000000000000000000000000000000000200000000000000000000000000000000" // )
 CHAR_MARIX[9] = "000000000900000000000002000000000000000000000000000000000000000000000000000000000000000000000" // *
CHAR_MARIX[10] = "000000000090000000000000000020000000000000000000000000000000000000000000000000000000000000000" // +
CHAR_MARIX[11] = "000000000009000000000000000200000000000000000000000000000000000000000000000400000000000000000" // ,
CHAR_MARIX[12] = "000000000000900000000000000000000000000000000000000000000000002000000000000000000000000000000" // -
CHAR_MARIX[13] = "000000000000090000000000000002000000000000000000000000000000000000000000000400000000000000000" // .
CHAR_MARIX[14] = "000000000000009000000000000000200000000000000000000000000000000000000000000000000000000000000" // /
CHAR_MARIX[15] = "000000002000000900000000000000000000000000000050000000000000000000000000000000500000000000000" // 0
CHAR_MARIX[16] = "200000000000000090000000000000000000000050000000000000000000000000000000000500000000000000000" // 1
CHAR_MARIX[17] = "000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000" // 2
CHAR_MARIX[18] = "002000000000000000900000000000000000000000000000000000000000000000000000000000000000000000000" // 3
CHAR_MARIX[19] = "000200000000000000090000000000000000000000000000000000000000000000000000000000000000000000000" // 4
CHAR_MARIX[20] = "000020000000000000009000000000000000000000000000000000000000000000000000000000000000000000000" // 5
CHAR_MARIX[21] = "000000000000000000000900000000000000000000000000000000000000020000000000000000000000000000000" // 6
CHAR_MARIX[22] = "000002000000000000000090000000000000000000000000000000000000000000000000000000000000000000000" // 7
CHAR_MARIX[23] = "000000000200000000000009000000000000000000000000000000000000000000000000000000000000000000000" // 8
CHAR_MARIX[24] = "000000020000000000000000920000000000000000000000000000000000000000000000000000000000000000000" // 9
CHAR_MARIX[25] = "000000000000000000000000292000000000000000060000000000000000000000000000000000000000000000000" // :
CHAR_MARIX[26] = "000000000000000000000000029000000000000000000000000000000000000000000000000600000000000000000" // ;
CHAR_MARIX[27] = "000000000002000000000000000900000000000000000000000000000000000000000000000000000000000000000" // <
CHAR_MARIX[28] = "000000000020000000000000000090000000000000000000000000000000000000000000000000000000000000000" // =
CHAR_MARIX[29] = "000000000000020000000000000009000000000000000000000000000000000000000000000000000000000000000" // >
CHAR_MARIX[30] = "000000000000002000000000000000900000000000000000000000000000000000000000000000000000000000000" // ?
CHAR_MARIX[31] = "000000000000000002000000000000090000000000000000000000000000000000000000000000000000000000000" // @
CHAR_MARIX[32] = "000000000000000000000000000000009000000000000000606000400400000080000000000000000000000000000" // A
CHAR_MARIX[33] = "000000000000000000000000000000000901004400000600010006000000000008000000000000000000000000000" // B
CHAR_MARIX[34] = "000000000000000000000000000000000094041000300010001006060000000000800000000000000000000000000" // C
CHAR_MARIX[35] = "000000000000000000000000000000000149660000000000046300040000000000080000000000000000000000000" // D
CHAR_MARIX[36] = "000000000000000000000000000000000006910000000000064000600000000000008000000000000000000000000" // E
CHAR_MARIX[37] = "000000000000000000000000000000000046196000000003060404000000000000000800000000000000000000000" // F
CHAR_MARIX[38] = "000000000000000000000000000000000410069600000010100604000000000000000080000000000000000000000" // G
CHAR_MARIX[39] = "000000000000000000000000000000000400006906000400000000006000000000000008000000000000000000000" // H
CHAR_MARIX[40] = "000000000000000050000000000000000000000094600060000060000000000000000000800500000000000000000" // I
CHAR_MARIX[41] = "000000000000000000000000000000000000000649604400000060004000000000000000080000000000000000000" // J
CHAR_MARIX[42] = "000000000000000000000000000000000030000066964040000000000000000000000000008000000000000000000" // K
CHAR_MARIX[43] = "000000000000000000000000060000000000000000690064000000003000000000000000000800000000000000000" // L
CHAR_MARIX[44] = "000000000000000000000000000000000000000004409600000000000000000000000000000080000000000000000" // M
CHAR_MARIX[45] = "000000000000000000000000000000000600000404006900000000000000000000000000000008000000000000000" // N
CHAR_MARIX[46] = "000000000000000500000000000000000010001060460096000000000000000000000000000000800000000000000" // O
CHAR_MARIX[47] = "000000000000000000000000000000000000030000040069000000000000000000000000000000080000000000000" // P
CHAR_MARIX[48] = "000000000000000000000000000000006000001000000000900000600000000000000000000000008000000000000" // Q
CHAR_MARIX[49] = "000000000000000000000000000000000104660000000000090600000000000000000000000000000800000000000" // R
CHAR_MARIX[50] = "000000000000000000000000000000006016400000000000009000640400000000000000000000000080000000000" // S
CHAR_MARIX[51] = "000000000000000000000000000000000003046000000000060900006000000000000000000000000008000000000" // T
CHAR_MARIX[52] = "000000000000000000000000000000000000000066000000000091106000000000000000000000000000800000000" // U
CHAR_MARIX[53] = "000000000000000000000000000000000660044000000000000019000000000000000000000000000000080000000" // V
CHAR_MARIX[54] = "000000000000000000000000000000004000600000000000606010900000000000000000000000000000008000000" // W
CHAR_MARIX[55] = "000000000000000000000000000000000064000000000000004000090600000000000000000000000000000800000" // X
CHAR_MARIX[56] = "000000000000000000000000000000000000000604030000000660009000000000000000000000000000000080000" // Y
CHAR_MARIX[57] = "000000000000000000000000000000004000000000000000004000060900000000000000000000000000000008000" // Z
CHAR_MARIX[58] = "000000020000000000000000000000000000000000000000000000000090000000000000000000000000000000200" // [
CHAR_MARIX[59] = "000000000000000000000000000000000000000000000000000000000009000000000000000000000000000000020" // \0
CHAR_MARIX[60] = "000000002000000000000000000000000000000000000000000000000000900000000000000000000000000000002" // ]
CHAR_MARIX[61] = "000000000000000000000200000000000000000000000000000000000000090000000000000000000000000000000" // ^
CHAR_MARIX[62] = "000000000000200000000000000000000000000000000000000000000000009000000000000000000000000000000" // _
CHAR_MARIX[63] = "080000800000000000000000000000000000000000000000000000000000000900000000000000000000000000000" // `
CHAR_MARIX[64] = "000000000000000000000000000000008000000000000000000000000000000090000000000000006060004004000" // a
CHAR_MARIX[65] = "000000000000000000000000000000000800000000000000000000000000000009010044000006000100060000000" // b
CHAR_MARIX[66] = "000000000000000000000000000000000080000000000000000000000000000000940410003000100010060600000" // c
CHAR_MARIX[67] = "000000000000000000000000000000000008000000000000000000000000000001496600000000000463000400000" // d
CHAR_MARIX[68] = "000000000000000000000000000000000000800000000000000000000000000000069100000000000640006000000" // e
CHAR_MARIX[69] = "000000000000000000000000000000000000080000000000000000000000000000461960000000030604040000000" // f
CHAR_MARIX[70] = "000000000000000000000000000000000000008000000000000000000000000004100696000000101006040000000" // g
CHAR_MARIX[71] = "000000000000000000000000000000000000000800000000000000000000000004000069060004000000000060000" // h
CHAR_MARIX[72] = "000000000000000000000000000000000000000080000000000000000000000000000000946000600000600000000" // i
CHAR_MARIX[73] = "000000000000000000000000000000000000000008000000000000000000000000000006496044000000600040000" // j
CHAR_MARIX[74] = "000000000000000000000000000000000000000000800000000000000000000000300000669640400000000000000" // k
CHAR_MARIX[75] = "000000000004040050000000006000000000000050080000000000000000000000000000006900640000000030000" // l
CHAR_MARIX[76] = "000000000000000000000000000000000000000000008000000000000000000000000000044096000000000000000" // m
CHAR_MARIX[77] = "000000000000000000000000000000000000000000000800000000000000000006000004040069000000000000000" // n
CHAR_MARIX[78] = "000000000000000500000000000000000000000000000080000000000000000000100010604600960000000000000" // o
CHAR_MARIX[79] = "000000000000000000000000000000000000000000000008000000000000000000000300000400690000000000000" // p
CHAR_MARIX[80] = "000000000000000000000000000000000000000000000000800000000000000060000010000000009000006000000" // q
CHAR_MARIX[81] = "000000000000000000000000000000000000000000000000080000000000000001046600000000000906000000000" // r
CHAR_MARIX[82] = "000000000000000000000000000000000000000000000000008000000000000060164000000000000090006404000" // s
CHAR_MARIX[83] = "000000000000000000000000000000000000000000000000000800000000000000030460000000000609000060000" // t
CHAR_MARIX[84] = "000000000000000000000000000000000000000000000000000080000000000000000000660000000000911060000" // u
CHAR_MARIX[85] = "000000000000000000000000000000000000000000000000000008000000000006600440000000000000190000000" // v
CHAR_MARIX[86] = "000000000000000000000000000000000000000000000000000000800000000040006000000000006060109000000" // w
CHAR_MARIX[87] = "000000000000000000000000000000000000000000000000000000080000000000640000000000000040000906000" // x
CHAR_MARIX[88] = "000000000000000000000000000000000000000000000000000000008000000000000006040300000006600090000" // y
CHAR_MARIX[89] = "000000000000000000000000000000000000000000000000000000000800000040000000000000000040000609000" // z
CHAR_MARIX[90] = "000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000900" // {
CHAR_MARIX[91] = "000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000090" // |
CHAR_MARIX[92] = "000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000009" // }
}

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

func gappedScore(charX rune, charY rune) int {
	if ((charX == rune(' ')) || (charY == rune(' '))) {
		return PENALTY_GAP
	} else {
		return PENALTY_MISMATCH
	}
}

func similarityScore(charX rune, charY rune) int {
	charXval, charXsize := utf8.DecodeRuneInString(string(charX))
	var charXpos int = int(charXval) - (32 + 1)

	if (charXsize == 1) && (charXpos >= 0) && (charXpos <= 92) {
		charYval, charYsize := utf8.DecodeRuneInString(string(charY))
		var charYpos int = int(charYval) - (32 + 1)

		if (charYsize == 1) && (charYpos >= 0) && (charYpos <= 92) {
			currentrow := CHAR_MARIX[charXpos]
			var currentcol int = (int(currentrow[charYpos]) - 48)

			if currentcol > 0 {
				return currentcol + SCORE_SIMILARITY
			} else {
				return SCORE_MISMATCH
			}
		} else {
			return SCORE_MISMATCH
		}
	} else {
		return SCORE_MISMATCH
	}
}

func characterScore(charX rune, charY rune) int {
	// exact match
	if(charX == charY){
		return SCORE_MATCH
	// wildcard in the possibility
	// (could probably also be handled by character similarity matrix, this is a shortcut)
	} else if (charY == '*') {
		return SCORE_WILDCARD
	// check character similarity
	} else {
		return similarityScore(charX, charY)
	}
}

func transposescore(charX rune, charY rune) int {
	// TODO: do something here
	return PENALTY_TRANSPOSE;
}

func fillMatrix(xSize int, ySize int, xString string, yString string) []int {
	var currentXChar rune
	var currentYChar rune
	gapScore := 0
	charScore := 0
	currentIndex := 0
	indexDiag := 0
	indexLeft := 0
	indexAbove := 0
	dpMatrix := make([]int, xSize * ySize)
	if DEBUG_PRINT_MATRIX == 1 { fmt.Println("\n", xString, yString) }
	for y := 0; y < ySize; y++ {
		for x := 0; x < xSize; x++ {
			if x > 0 {
				currentXChar = []rune(xString)[x - 1]
			}
			if y > 0 {
				currentYChar = []rune(yString)[y - 1]
			}

			// TODO: start both of the above loops at 1
			// 		 and remove this condition
			// 		 and remove the print condition below
			if (x > 0) && (y > 0) {
				gapScore = gappedScore(currentXChar, currentYChar)
				charScore = characterScore(currentXChar, currentYChar)

				currentIndex = (y * xSize) + x
				indexDiag = ((y - 1) * xSize) + (x - 1)
				indexLeft = (y * xSize) + (x - 1)
				indexAbove = ((y - 1) * xSize) + x

				// the value of the current pixels is:
				// - the maximum of:
				//   - the diagonal value and the current score (there is a series of matches)
				//   - 0 (just in case all scores are below zero)
				//   - the above value and the gap score (there is a gap in the matches)
				//   - the left value and the gap score (there is a gap in the matches)
				// - minus the cost of movement
				dpMatrix[currentIndex] = max4(
					dpMatrix[indexDiag] + charScore,
					0,
					dpMatrix[indexAbove] + gapScore,
					dpMatrix[indexLeft] + gapScore) + MOVE_COST

				// values above and left are higher than diagonal, suggesting a transpose
				if ((dpMatrix[indexAbove] > dpMatrix[indexDiag]) && (dpMatrix[indexLeft] > dpMatrix[indexDiag])) {
					// give the diagonal the maximum value of the other neighbors
					dpMatrix[indexDiag] = max(
						dpMatrix[indexAbove],
						dpMatrix[indexLeft])
					// give current space the value of transposed caharacters, if less than current value
					dpMatrix[currentIndex] = max(
						dpMatrix[indexDiag] + transposescore(currentXChar, currentYChar),
						dpMatrix[currentIndex])
				}
			}

			if DEBUG_PRINT_MATRIX == 1 {
				if ((x == 0) || (y == 0)) {
					if ((x == 0) && (y == 0)) {
						fmt.Printf("  ")
					}
					if ((x != 0) && (y == 0)) {
						fmt.Printf("__%v ", string(currentXChar))
					}
					if ((x == 0) && (y != 0)) {
						fmt.Printf("%v ", string(currentYChar))
					}
				} else {
					fmt.Printf("%3d ", dpMatrix[currentIndex])
				}
			}
		}
		if DEBUG_PRINT_MATRIX == 1 { fmt.Println() }
	}
	return dpMatrix
}

func backTrack(dpMatrix []int, xSize int, ySize int, word string) float64 {
	// 1. Find the local maximum, in the bottom row or right column
	//    This determines the starting point for the backtrack

	localMax := 0
	localMaxindex := 0
	currentIndex := 0
	argumentXCut := 0

	// check bottom-most row for a maximum
	for x := 1; x < xSize; x++ {
		currentIndex = (ySize * xSize) - x
		if dpMatrix[currentIndex] > localMax {
			localMaxindex = currentIndex
			localMax = dpMatrix[localMaxindex]
			argumentXCut = xSize - x - 1
			if DEBUG_PRINT_ARGPOS == 1 { fmt.Printf("=1=> %d : %d (%d, %d)\n", localMax, argumentXCut, currentIndex, xSize) }
		}
	}

	// check right-most column for a maximum
	for y := (ySize - 1); y > 0; y-- {
		currentIndex = (y * xSize) + (xSize - 1)
		if dpMatrix[currentIndex] > localMax {
			localMaxindex = currentIndex
			localMax = dpMatrix[localMaxindex]
			argumentXCut = 1
			if DEBUG_PRINT_ARGPOS == 1 { fmt.Printf("=2=> %d : %d (%d, %d)\n", localMax, argumentXCut, currentIndex, xSize) }
		}
	}

	// 2. Find the global maximum by backtracking from local max

	currentIndex = localMaxindex
	globalMax := localMax

	indexDiag := 0
	indexLeft := 0
	indexAbove := 0
	valueDiag := 0
	valueLeft := 0
	valueAbove := 0

	continueTesting := 1
	for continueTesting > 0 {
		if dpMatrix[currentIndex] > globalMax {
			globalMax = dpMatrix[currentIndex]
			argumentXCut = int(math.Abs(math.Remainder(float64(currentIndex), float64(xSize))))
			if DEBUG_PRINT_ARGPOS == 1 { fmt.Printf("=3=> %d : %d (%d, %d)\n", globalMax, argumentXCut, currentIndex, xSize) }
		}

		indexLeft = currentIndex - 1
		indexAbove = currentIndex - xSize
		indexDiag = indexAbove - 1
		// make sure we don't go outside the matrix bounds
		valueDiag = 0
		if (indexLeft < 0) || (indexAbove < 0) || (indexDiag < 0) {
			valueLeft = math.MinInt64
			valueAbove = math.MinInt64
			valueDiag = math.MinInt64
		} else {
			if indexLeft % xSize != 0 {
				valueLeft = dpMatrix[indexLeft]
			} else {
				valueLeft = math.MinInt64
				valueDiag = math.MinInt64
			}
			if indexAbove > xSize {
				valueAbove = dpMatrix[indexAbove]
			} else {
				valueAbove = math.MinInt64
				valueDiag = math.MinInt64
			}
			if valueDiag != math.MinInt64 {
				valueDiag = dpMatrix[indexDiag]
			}
		}

		// no equals, so biased up (ties hit else)
		if valueLeft > valueAbove {
			// no equals, so biased diagonal (ties hit else)
			if valueLeft > valueDiag {
				// go left
				currentIndex = indexLeft
			} else {
				// go diag
				currentIndex = indexDiag
			}
		} else {
			// no equals, so biased diagonal (ties hit else)
			if valueAbove > valueDiag {
				// go up
				currentIndex = indexAbove
			} else {
				// go diag
				currentIndex = indexDiag
			}
		}

		if (valueDiag == math.MinInt64) && (valueAbove == math.MinInt64) && (valueLeft == math.MinInt64) {
			continueTesting = 0
		}
	}

	scorePercent := (float64(globalMax) / float64(max((xSize - 1), (ySize - 1)) * 10)) * 100
	return scorePercent
}

func score(wordA string, wordB string) float64 {
	// add one to each dimension for the padding row/collumn
	xSize := len([]rune(wordA)) + 1
	ySize := len([]rune(wordB)) + 1
	dpMatrix := fillMatrix(xSize, ySize, wordA, wordB)
	btResults := backTrack(dpMatrix, xSize, ySize, wordA)
	sFactor := scoreFactor(wordA,wordB)
	finalScore := btResults * sFactor
	if DEBUG_PRINT_SCORES == 1 { fmt.Printf("score: %#v %#v\n", btResults, wordB) }
	return finalScore
}

func bestMatch(word string, wordList []string) string {
	initCharMatrix()
	wordScore := 0.0
	maxScore := 0.0
	wordMatch := ""
	garbageCollectCounter := 0
	for i := 0; i < len(wordList); i++ {
		garbageCollectCounter++
		if len(wordList[i]) > 0 {
			wordScore = score(word, wordList[i])
			if (wordScore >= maxScore) {
				wordMatch = wordList[i]
				maxScore = wordScore
			}
			if (garbageCollectCounter % 2) == 0 {
				runtime.GC()
			}
		}
	}
	return wordMatch
}
