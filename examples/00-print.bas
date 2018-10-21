10 REM This program demonstrates printing Ints & strings

20 PRINT "Hello, world\n"

30 LET a = 3
40 PRINT "The contents of the variable 'a' are", a, "\n"

50 LET a$ = "My name is String"
60 PRINT "The contents of the variable 'a$' are", a$, "\n"


80 LET a$ = "Steve"
90 LET b = LEN a$

100 PRINT "String '" a$ "' is ", b, "characters long\n"
110 PRINT "'Steve' is STILL ", LEN "Steve", "characters long\n"
