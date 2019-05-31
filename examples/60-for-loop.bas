10 REM
20 REM This program tests for-loops a little.
30 REM

100 PRINT "IN ONES\n"
110 FOR I = 1 to 10 STEP 1
120   PRINT "\t",I, "\n"
130 NEXT I

200 PRINT "IN TWOS\n"
210 FOR I = 0 to 10 STEP 2
220   PRINT "\t",I,"\n"
230 NEXT I

300 PRINT "Backwards\n"
310 FOR I = 10 to 0 STEP -1
320   PRINT "\t",I,"\n"
330 NEXT I

400 PRINT "With a variable\n"
410 LET term=4
420 FOR I = 1 TO term STEP 1
430   PRINT "\t", I, "\n"
440 NEXT I

500 PRINT "With an expression\n"
510 FOR I = 1 TO 3 * 4 + 5
520   PRINT "\t", I, "\n"
530 NEXT I
