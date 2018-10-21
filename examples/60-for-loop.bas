10 REM This program tests for-loops a little.
15 REM

20 PRINT "IN ONES\n"
30 FOR I = 1 to 10 STEP 1
40 PRINT "",I, "\n"
50 NEXT I

100 PRINT "IN TWOS\n"
110 FOR I = 0 to 10 STEP 2
120 PRINT "",I,"\n"
130 NEXT I


500 PRINT "Backwards\n"
510 FOR I = 10 to 0 STEP -1
520 PRINT "",I,"\n"
530 NEXT I

1000 PRINT "With a variable\n"
1010 LET term=4
1020 FOR I = 1 TO term STEP 1
1030   PRINT "", I, "\n"
1040 NEXT I
