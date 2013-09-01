package parser

type lexer struct {
    text string
    pos int
}

var first int
var pos int

func (l lexer) next() rune {
	if pos >= len([]rune(l.text)) {
		pos += 1
		return 0
	}
	c := []rune(l.text)[pos]
	pos += 1
	return c
}

func (l lexer) emit(lval *NutshSymType) {
	lval.val = string([]rune(l.text)[first:pos-1])
	pos -= 1
	first = pos
}

func (l lexer) skip() {
	first = pos
}

func (l lexer) Lex(lval *NutshSymType) int {
	var c rune

	start:
	c = l.next()
	switch {
	case whitespace(c):
		l.skip()
		goto start
	case c == '"':
		l.skip()
		c = l.next()
		for c != '"' {
			c = l.next()
		}
		l.emit(lval)
		l.next()
		l.skip()
		return STRING
	case c == '=':
		c = l.next()
		if c == '~' {
			c = l.next()
			l.emit(lval)
			return MATCH
		} else {
			panic("Syntax error: Expected ~ after =.")
		}
	case alnum(c):
		for alnum(c) {
			c = l.next()
		}
		l.emit(lval)
		switch lval.val {
		case "if":
			return IF
		case "def":
			return DEF
		case "prompt":
			return PROMPT
		default:
			return IDENTIFIER
		}
    }
	if c != 0 {
		l.next()
		l.emit(lval)
	}
    return int(c)
}

func alnum(r rune) bool {
    return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_'
}

func whitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n'
}

func (l lexer) Error(e string) {
    panic(e)
}
