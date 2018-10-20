10 REM This program demonstrates our in-progress IF support
20 REM
30 REM For the moment we skip a single token, and allow a single
40 REM expression between THEN+ELSE, or ELSE+NEWLINE
40 REM

100 IF 1 THEN PRINT "OK1\n" : ELSE PRINT "FAIL1\n"
110 IF 1 THEN PRINT "OK2\n"

120 REM
130 REM Prove execution keeps going.
140 REM

150 LET a = 3
160 PRINT "A is", a, "\n"
