 10 REM For all your enterprise needs
 20 FOR A = 1 TO 100
 30   let f = A % 3
 40   let b = A % 5
 50   let both = A % 15
 60   LET msg = A
 70   IF f = 0 THEN LET msg = "Fizz"
 80   IF b = 0 THEN LET msg = "Buzz"
 90   IF both = 0 THEN LET msg = "FizzBuzz"
110   PRINT msg, "\n"
120 NEXT A
