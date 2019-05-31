 10 REM
 20 REM THis program demonstrates the use of arrays
 30 REM It creates a 10x10 array, full of random numbers,
 40 REM then prints it out - as hex

 50 REM Setup hex-table
 60 GOSUB 2000

100 REM
110 REM Generate a 10x10 array and populate it with random numbers
120 REM
130 DIM a(10,10)
140 FOR X = 0 TO 10
150   FOR Y = 0 TO 10
160    LET a[X,Y] = RND 255
170   NEXT Y
180 NEXT X

200 REM
210 REM Now print the contents - as hex values
220 REM
230 FOR X = 0 TO 10
240  FOR Y = 0 TO 10
250   LET v = a[X,Y]
260   GOSUB 1000
270   PRINT " "
280  NEXT Y
290  PRINT "\n"
300 NEXT X

400 END


1000 REM
1010 REM Print the value in "v" as a two-digit Hex number
1020 REM
1030 LET a1 = INT(v / 16)
1040 LET b1 = v - ( a1 * 16 )
1050 LET x = hex[a1] + hex[b1]
1060 PRINT x
1070 RETURN


2000 REM
2010 REM Setup a hex-table, via the DATA statements later.
2020 REM
2030 DIM hex(16)
2040 FOR I = 0 TO 15
2050  READ x
2060   hex[I] = CHR$ x
2070 NEXT I
2080 RETURN


10000 REM 0-9
10010 DATA 48, 49, 50, 51, 52, 53, 55, 55, 56, 57
10020 REM A-F
10030 DATA 65, 66, 67, 68, 69, 70
