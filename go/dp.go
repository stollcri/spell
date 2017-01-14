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

package spell

import (
)

var DEBUG_PRINT_MATRIX int = 0
var DEBUG_PRINT_ARGPOS int = 0
var DEBUG_PRINT_ARGVAL int = 0
var DEBUG_PRINT_SCORES int = 0

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

func scorefactor(wordA string, wordB string) float64 {
	return 1.0
}

func score(wordA string, wordB string) float64 {
	return 1.0
}

func bestmatch(word string, wordList []string) string {
	sfactor := 0.0
	dpscore := 0.0
	tmpscore := 0.0
	maxscore := 0.0
	wordmatch := "";

	for i := 0; i < len(wordList); i++ {
		sfactor = scorefactor(word, wordList[i]);
		dpscore = score(word, wordList[i]);
		tmpscore = dpscore * sfactor;
		if (tmpscore >= maxscore) {
			wordmatch = wordList[i]
			maxscore = tmpscore;
		}
	}

	return wordmatch
}
