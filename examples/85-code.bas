10 REM
20 REM Prints out ASCII values
30 REM

40 PRINT "'*' is ", CODE "*", "\n"
50 PRINT "' ' is ", CODE " ", "\n"

100 LET A="Steve"
110 LET L=LEN A - 1
120 FOR I=0 TO L
130   LET X = MID$ A, I, 1
140   PRINT "Character ", I, "is", X, "with code", CODE X, "\n"
150 NEXT I
