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

#include "spell.h"
#include <errno.h>
#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "dp.h"

static inline void exitapp(int exitcode) __attribute__ ((noreturn));
static inline void exitapp(int exitcode)
{
	exit(exitcode);
}

static inline void handleinterupt(int signum)
{
	exitapp(EXIT_SUCCESS);
}

static inline void setinterupthandlers()
{
	signal(SIGUSR1, handleinterupt);
	signal(SIGUSR2, handleinterupt);
	signal(SIGINT, handleinterupt);
}

static inline char **readconfigfile(char *configfilename, int *possibilitycount)
{
	FILE *configfile = fopen(configfilename, "r");
	if (configfile == NULL) {
		fprintf(stderr, "Error opening '%s': %s\n", configfilename, strerror(errno));
		exitapp(EXIT_FAILURE);
	}

	char currentchar;
	int stringposition = 0;
	int filelineposition = 0;
	int linelength = FILE_LINE_BLOCK_LENGTH;
	int linescount = FILE_LINES_BLOCK_COUNT;
	char *configline = (char*)malloc(linelength * sizeof(char));
	char **configlines = malloc(linescount * sizeof(char*));

	while ((currentchar = fgetc(configfile)) != EOF) {
		if (currentchar == '\n') {
			// end the last line
			configline[stringposition] = '\0';
			// add the line to the array
			configlines[filelineposition] = configline;
			// increment the line counter
			++filelineposition;
			// resize the line count as needed
			if (filelineposition >= linescount) {
				linescount += FILE_LINES_BLOCK_COUNT;
				configlines = realloc(configlines, linescount * sizeof(char*));
			}
			// reset the position for the next line
			stringposition = 0;
			// reset the line length
			linelength = FILE_LINE_BLOCK_LENGTH;
			// start a new line
			configline = (char*)malloc(linelength * sizeof(char));
		} else {
			configline[stringposition] = currentchar;
			++stringposition;
			// resize the line string as needed
			if (stringposition >= linelength) {
				linelength += FILE_LINE_BLOCK_LENGTH;
				configline = realloc(configline, linelength * sizeof(char));
			}
		}
	}
	if (ferror(configfile)) {
	    fprintf(stderr, "IO error: %s\n", strerror(errno));
	    exitapp(EXIT_FAILURE);
	}
	fclose(configfile);

	*possibilitycount = filelineposition;
	return configlines;
}

//~~~~~~
// MAIN
//~~~~~~

int main(int argc, char **argv)
{
	// first thing, prepare to be interupted
	// (this shouldn't be needed at all, but just in case)
	setinterupthandlers();

	char *argone = argv[1];
	FILE *outfile = stdout;

	int showversion = 0;
	int showhelp = 0;

	if (argone[0] == '-') {
		if (argone[1] == 'v') {
			showversion = 1;
		} else if ((argone[1] == 'h') || (argone[1] == '?')) {
			showhelp = 1;
		}
	}

	if (showversion == 1) {
		printf(PROGRAM_NAME " " PROGRAM_VERS "\n");
		printf(PROGRAM_COPY "\n");
		printf(PROGRAM_URLS "\n");
		exitapp(EXIT_SUCCESS);
	}
	if (showhelp == 1) {
		printf(PROGRAM_HELP "\n");
		exitapp(EXIT_SUCCESS);
	}

	// load possible corrections
	int possibilitycount = 0;
	char **possibilities = NULL;
	if (getenv("SPELL_FILE")) {
		possibilities = readconfigfile(getenv("SPELL_FILE"), &possibilitycount);
	} else {
		possibilities = readconfigfile(".spell", &possibilitycount);
	}

	// write the recomended command string, then close the target file
	char *recomendedstring = NULL;
	recomendedstring = bestmatch(argone, possibilities, possibilitycount);

	fputs(recomendedstring, outfile);
	fputs("\n", outfile);
	fclose(outfile);

	exitapp(EXIT_SUCCESS);
}
