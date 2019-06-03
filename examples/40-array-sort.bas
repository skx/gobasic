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

 200 REM
 210 REM Print initial array
 220 REM
 230 PRINT "Initial array:\n\t"
 240 GOSUB 1000

 300 REM
 310 REM SortyMcSortFace
 320 REM
 330 FOR I = 1 TO 10
 340  FOR J = 0 TO 9
 350    LET N=J+1
 360    IF A[I] < A[N] THEN SWAP A[I], A[N]
 370  NEXT J
 380 NEXT I

 400 REM
 410 REM Print the sorted array
 420 REM
 430 PRINT "Sorted array:\n\t"
 440 GOSUB 1000

 500 REM
 510 REM Finished
 520 REM
 540 END

1000 REM Print the array
1010 FOR I = 1 TO 10
1020  PRINT A[I], " "
1030 NEXT I
1040 PRINT "\n"
1050 RETURN

9000 DATA 32, 4, 11
9001 DATA 2, 93, 3
9002 DATA 102, 5, -1
9003 DATA 1
