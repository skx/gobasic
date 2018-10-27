10 REM
20 REM This is a horrid script which converts a string to lower-case
40 REM

100 LET A="STEVE IS LOWER-CASE"
110 LET L=( LEN A ) - 1
120 FOR I=0 TO L
130   LET A$ = MID$ A, I, 1
140   IF A$ >= "A" AND A$ <= "Z" THEN GOSUB 8000
150   PRINT A$
160 NEXT I
170 PRINT "\n"


200 LET A="steve is in upper-case, now!"
210 LET L=(LEN A) - 1
220 FOR I=0 TO L
230   LET A$ = MID$ A, I, 1
240   IF A$ >= "a" AND A$ <= "z" THEN GOSUB 9000
250   PRINT A$
260 NEXT I
270 PRINT "\n"

300 END


8000 REM REM
8010 REM Converts the character in A$ to lower-case
8020 REM
8030 LET A$ = CHR$ ( ( CODE A$ ) + 32 )
8045 RETURN



9000 REM
9010 REM Converts the character in A$ to upper-case
9020 REM
9030 LET A$ = CHR$ ( ( CODE A$ ) - 32 )
9040 RETURN
