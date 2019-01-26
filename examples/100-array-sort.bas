10 REM
20 REM This program demonstrates sorting an array.
30 REM
40 REM The sort is a bubble-sort.  Yeah.
50 REM


100 REM Get the numbers into the array
110 DIM A(10)
120 FOR I = 1 TO 10
130   READ A[I]
140 NEXT I


200 REM Print initial array
210 PRINT "Initial array:\n"
220 FOR I = 1 TO 10
230  PRINT A[I], " "
240 NEXT I
250 PRINT "\n"


300 REM SortyMcSortFace
310 FOR I = 1 TO 10
320  FOR J = 0 TO 9
330    LET N=J+1
340    IF A[I] < A[N] THEN SWAP A[I], A[N]
350  NEXT J
360 NEXT I


300 REM Print the sorted array
310 PRINT "Sorted array:\n"
320 FOR I = 1 TO 10
330  PRINT A[I], " "
340 NEXT I
350 PRINT "\n"



1000 DATA 32, 4, 11
1001 DATA 2, 93, 3
1002 DATA 102, 5, -1
1003 DATA 1
