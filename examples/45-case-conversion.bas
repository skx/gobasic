10 REM
20 REM This is a horrid script which converts a string to lower-case
40 REM

100 LET A="STEVE IS LOWER-CASE"
110 LET L=LEN A - 1
120 FOR I=0 TO L
130   LET A$ = MID$ A, I, 1
140   IF A$ >= "A" AND A$ <= "Z" THEN GOSUB 8000
150   PRINT A$
160 NEXT I
170 PRINT "\n"
180 END


8000 REM REM
8010 REM Converts the character in A$ to lower-case
8020 REM NOTE: Destroys variable 'b'
8030 REM
8040 LET b = CODE A$ + 32
8050 LET A$ = CHR$ b
8060 RETURN



9000 REM
9010 REM Converts the character in A$ to upper-case
9020 REM NOTE: Destroys variable 'b'
9030 REM
9040 LET b = CODE A$ - 32
9050 LET A$ = CHR$ b
9060 RETURN
