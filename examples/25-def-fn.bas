 10 REM
 20 REM This program demonstrates the usage
 30 REM of DEF FN, and FN.
 40 REM
 50 DEF FN double(x) = x + x
 60 DEF FN square(x) = x * x
 70 DEF FN cube(x)   = x * x * x
 80 DEF FN quad(x)   = x * x * x * x
 90 PRINT "N\tDoubled\tSquared\tCubed\tQuadded (?)\n"
100 FOR I = 1 TO 10
110   PRINT I, "\t", FN double(I), "\t", FN square(I), "\t", FN cube(I), "\t", FN quad(I), "\n"
120 NEXT I
