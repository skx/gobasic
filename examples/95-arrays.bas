 10 REM THis program demonstrates the use of arrays
 20 REM
 30 DIM a(10,10)
 40 FOR X = 0 TO 10
 50   FOR Y = 0 TO 10
 60    LET a[X,Y] = X * Y
 70   NEXT Y
 80 NEXT X
100 FOR X = 0 TO 10
110  FOR Y = 0 TO 10
120   PRINT X, "*", Y, "=", a[X,Y], "\n"
130  NEXT Y
140 NEXT X
