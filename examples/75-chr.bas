10 REM
20 REM Prints out characters.
30 REM

40 PRINT "\nThis program outputs some printable ASCII characters:\n\n"
50 FOR a=32 TO 126
60   PRINT CHR$ a ,
70 IF ( a % 8) =  7  THEN PRINT "\n":
80 NEXT a
90 PRINT "\n\n"
