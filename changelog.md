# Nov 14th, '23, v0.0.2
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