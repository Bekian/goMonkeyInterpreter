
# Nov 14th, '23, v0.0.2; bottom of page 22

    - Init up to chapter 1.4
    - Added tokens: 
        - ASSIGN       =
	    - PLUS         +
	    - MINUS        -
	    - BANG         !
	    - ASTERISK     *
	    - SLASH        / 
        - Less than    <
        - Greater than >
        - True         true
        - False        false
        - If           if
        - Else         else 
        - Return       return
    - Changed lexer test string and map to match
    - Refactored error log in lexer_test to match source cause its not one line and easier to read
    - Refactored two functions in the lexer to be laid out in a way that is more reflective of when they're called in the test
    - Added corresponding the above new token mappings to the lexer

# Nov 15th '23, v0.1.0
    - Added peekChar function to the lexer which allows you to read forwards without changing the position of the current character
    - Added equal to (==) and not equal to (!=) conditional tokens
    - Added a more comprehensive test to verify all the current tokens and keywords are funcitoning
    - Added a repl which allows the user to repeatedly enter characters and symbols that will be converted to tokens
    - Added a main.go file which just prints some nice text and runs the repl

## v0.1.1 - bottom of 34
    - Added beginnings of the ast

# Nov 17th '23, v0.1.2
    - Added beginnings of the parser (need to finish documenting and maybe reread)