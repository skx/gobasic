# TODO

Weekend BASIC - Max budget 10 hours.

Rough plan, which might change over time.

* [x] Tokenize input
* [x] Allow a program to be created, just an array of tokens.
* [x] Add a case-statement to switch on token-type.
      * Just ignore LINENO + Newlines
      * If we assume one statement per-line we just have to write a handler for PRINT, a handler for FOR, etc, etc.
* [x] Implement a simple PRINT, which will just output static-strings.
* [x] Implement a simple GOTO.
      * Simplest possible control-flow statement.
      * All we have to do is change the offset we're interpreting from in our array of tokens.
* [x] Implement a simple Stack
* [x] Implement GOSUB / RETURN to use that new stack.
* [x] Update PRINT to handle multiple comma-separated arguments.
      * Allow ints, and strings to be mixed.
* [x] Setup a holder for variables
      * Variables named "BLAH" will hold ints.
      * Variables named "BLAH$" will hold strings.
         * This needs revisiting when we have Objects.
         * Strings can probably be ignored to be honest..
* [x] Handle LET statements.
      * These will store int values in names.
      * Note this also requires expressions.
         * We'll start with ints, but will switch to Objects later.
* [x] Update our PRINT statement to print int-variable contents.
      * This should be trivial :)
         * We'll revisit this when we have Objects.
* [x] Handle FOR loops both with and without step-offsets
      * FOR i=1 TO 10
      * FOR i=0 TO 10 STEP 2
      * FOR i=10 TO 0 STEP -1
* [ ] Handle IF statements and conditionals
  * I'm going to decide "IF COND THEN STATEMENT [ELSE STATEMENT] FI"




## Limitations

We're restricted to int-variables, because factor(), term(), etc, will only
return ints.

I would need an object-holder to solve that problem.  If I do that then it
becomes hard to ensure type-safety, though I guess I could do a check on
vars:get/set




## Enhancements?

Now we have expressions and variables we could imagine:

    10 LET A = 300
    20 GOSUB A

That might be too horrid to imagine; I'm not sure.




## Missing Features

Obvious missing features:

* IF
   * This requires conditionals.

Storing strings in variables is doesn't let them come out again, except
for in PRINT.




## Improvements

* [ ] Consider the use of a function to find line-numbers
      * Ideally we'd scan the program once, at load-time, to find line-numbers.
         * Since we do this for both GOTO & GOSUB.
* [ ] Allow registering functions so we can easily add CHRS$, LEFT$, ABS, RAND etc.   * Of course CHR$, LEFT$, LEN, etc, all require the use of strings.
      * Here we go again.
