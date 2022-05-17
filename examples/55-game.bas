01 REM
02 REM This is a simple guessing game.
03 REM
04 REM The computer picks a random number, and you have to guess it.
05 REM
06 REM Inspired by the code found here:
07 REM
08 REM     http://www.worldofspectrum.org/ZXBasicManual/zxmanchap3.html
09 REM

 10 LET b=RND 100
 20 LET count=1
 30 PRINT "I have picked a random number (1-100), please guess it!!\n"
 40 INPUT "Enter your choice:", a
 60 IF b = a THEN GOTO 2000 ELSE PRINT "Your choice was ":
 70 IF a < b THEN PRINT "too low!\n\n":
 80 IF a > b THEN PRINT "too high!\n\n":
 90 LET count = count + 1
100 GOTO 40


2000 PRINT "\n\nYou guessed my number!\n"
2010 PRINT "You took", count, "attempts.\n"
2020 END
