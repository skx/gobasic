10 REM
20 REM Prints out characters.
30 REM

40 FOR a=32 TO 127:
50   PRINT CHR$ a,
60 IF ( a % 8) =  0  THEN PRINT "\n":
70 NEXT a
