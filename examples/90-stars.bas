10 REM
20 REM Example adapted from Wikiepedia:
30 REM   https://en.wikipedia.org/wiki/BASIC#Syntax
40 REM


100 INPUT "What is your name: ", U$
110 PRINT "Hello ", U$

200 INPUT "How many stars do you want: ", N

300 LET S$ = ""
310 FOR I = 1 TO N
320   S$ = S$ + "*"
330 NEXT I
340 PRINT S$

400 INPUT "Do you want more stars? ", A$
410 IF LEN(A$) = 0 THEN GOTO 400
420 A$ = LEFT$(A$, 1)
430 IF A$ = "Y" THEN GOTO 200
440 IF A$ = "y" THEN GOTO 200

500 PRINT "Goodbye ", U$
510 END
