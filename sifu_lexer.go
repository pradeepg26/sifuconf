package main

import (
  "fmt"
  "strings"
  "unicode"
	"unicode/utf8"
)

type itemType int

const (
  KEYWORD itemType = iota
  VARIABLE
  ASSIGNMENT
  NUMBER
  STRING
  COMMENT
  ERROR
  NEWLINE
  EOF
)

const eof = -1

type token struct {
  itemType itemType
  pos int
  itemValue string
}

func (t itemType) String() string {
  switch t {
  case KEYWORD:
    return "KEYWORD"
  case VARIABLE:
    return "VARIABLE"
  case ASSIGNMENT:
    return "ASSIGNMENT"
  case NUMBER:
    return "NUMBER"
  case STRING:
    return "STRING"
  case COMMENT:
    return "COMMENT"
  case ERROR:
    return "ERROR"
  case NEWLINE:
    return "NEWLINE"
  case EOF:
    return "EOF"
  default:
    return "UNKNOWN_TYPE"
  }
}

func (t token) String() string {
  switch t.itemType {
  case EOF:
    return "EOF"
  case NEWLINE:
    return "NEWLINE"
    // return fmt.Sprintf("{type: '%s', value: '\\n'}", t.itemType)
  case ERROR:
    return fmt.Sprintf("{type: '%s', value: '%s'}", t.itemType, t.itemValue)
  default:
    if len(t.itemValue) > 10 {
      return fmt.Sprintf("{type: '%s', value: %.10q...}", t.itemType, t.itemValue)
    } else {
      return fmt.Sprintf("{type: '%s', value: '%s'}", t.itemType, t.itemValue)
    }
  }
}

type lexer struct {
  name string
  input string
  start int
  pos int
  width int
  items chan token
}

func lex(name, data string) (*lexer) {
  l := &lexer{
    name: name,
    input: data,
    items: make(chan token),
  }
  go l.run()
  return l
}

// ------------------------- Lexer Methods -------------------------
func (l *lexer) run() {
    for state := lexMain; state != nil; {
      state = state(l)
    }
    close(l.items)
}

func (l *lexer) emit(t itemType) {
  l.items <- token{t, l.start, l.input[l.start : l.pos]}
  l.start = l.pos
}

// next returns the next rune in the input.
func (l *lexer) next() rune {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = w
	l.pos += l.width
	return r
}

// peek returns but does not consume the next rune in the input.
func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// backup steps back one rune. Can only be called once per call of next.
func (l *lexer) backup() {
	l.pos -= l.width
}

// accept consumes the next rune if it's from the valid set.
func (l *lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

// ignore skips over the pending input before this point.
func (l *lexer) ignore() {
	l.start = l.pos
}

// lineNumber reports which line we're on, based on the position of
// the previous item returned by nextItem. Doing it this way
// means we don't have to worry about peek double counting.
// func (l lexer) lineNumber() int {
  // return 1 + strings.Count(l.input[:l.lastPos], "\n")
// }

// ------------------------- Utility Functions -------------------------
// isSpace reports whether r is a space character.
func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

// isEndOfLine reports whether r is an end-of-line character.
func isEndOfLine(r rune) bool {
	return r == '\r' || r == '\n'
}

// isAlphaNumeric reports whether r is an alphabetic, digit, or underscore.
func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

// ------------------------- State Functions -------------------------
type stateFn func(*lexer) stateFn

func lexMain(l *lexer) stateFn {
  r := l.next()
  if (r == eof) {
    l.emit(EOF)
    return nil
  } else {
    return lexWord
  }
}

func lexWord(l *lexer) stateFn {
  // Helper internal function to emit a word
  // Keep scanning until you see one of a Space, Assignment or EOL
  for l.pos <= len(l.input) {
    switch n := l.peek(); {
    case isSpace(n):
      emitWord(l)
      consumeSpaces(l)
      return lexWord
    case n == '/':
      emitWord(l)
      return lexComment
    case n == '=':
      emitWord(l)
      return lexAssign
    case isEndOfLine(n):
      return l.errorf("Unexpected NEWLINE")
    case n == eof:
      return l.errorf("Unexpected EOF")
    }
    l.next()
  }
  return l.errorf("Unexpected EOF")
}

func emitWord(l *lexer) {
  word := l.input[l.start:l.pos]
  switch word {
  case "":
    // Ignore empty strings
    return
  case "final", "secret", "abstract", "override":
    l.emit(KEYWORD)
    return
  default:
    l.emit(VARIABLE)
    return
  }
}

// Consumes all whitespace characters
func consumeSpaces(l *lexer) {
  for isSpace(l.peek()) {
    l.next()
  }
  l.ignore()
}

func lexAssign(l *lexer) stateFn {
  l.next()
  l.emit(ASSIGNMENT)
  return lexValue
}

func lexComment(l *lexer) stateFn {
  if l.peek() != '/' { // Check next char
    l.backup()
    return l.errorf("Invalid character '/'")
  }

  for !isEndOfLine(l.peek()) {
    l.next()
  }
  l.emit(COMMENT)

  // Comments are always terminated at EOL
  l.next()
  l.emit(NEWLINE)
  return lexMain
}

func lexValue(l *lexer) stateFn {
  // Consume all spaces after '='
  consumeSpaces(l)
  switch r := l.next(); {
  case r == '+' || r == '-' || r == '.' || ('0' <= r && r <= '9'):
    l.backup()
    return lexNumber
  case r == '"':
    return lexString
  case isEndOfLine(r):
    return l.errorf("Unexpected NEWLINE")
  case r == eof:
    return l.errorf("Unexpected EOF")
  default:
    return l.errorf("Unexpected character '%s'", string(r))
  }
}

func lexAfterValue(l *lexer) stateFn {
  // Just lexed value... next must be either '/' or EOL or EOF
  consumeSpaces(l)
  switch r := l.next(); {
  case r == '/':
    return lexComment
  case isEndOfLine(r):
    l.emit(NEWLINE)
    return lexMain
  case r == eof:
    l.emit(EOF)
    return nil
  default:
    return l.errorf("Bad character after value assignment. '%s'", string(r))
  }
}

// lexString scans a quoted string.
func lexString(l *lexer) stateFn {
  Loop:
	for {
		switch l.next() {
		case '\\':
			if r := l.next(); r != eof && r != '\n' {
				break
			}
			fallthrough
		case eof, '\n':
			return l.errorf("unterminated quoted string")
		case '"':
			break Loop
		}
	}
	l.emit(STRING)
  return lexAfterValue
}

// lexNumber scans a number: decimal, or float. This
// isn't a perfect number scanner - for instance it accepts "." and "0x0.2"
// and "089" - but when it's wrong the input is invalid and the parser (via
// strconv) will notice.
func lexNumber(l *lexer) stateFn {
	if !l.scanNumber() {
		return l.errorf("bad number syntax: %q", l.input[l.start:l.pos])
	}
	l.emit(NUMBER)
	return lexAfterValue
}

func (l *lexer) scanNumber() bool {
	// Optional leading sign.
	l.accept("+-")
	// Is it hex?
	digits := "0123456789"
	l.acceptRun(digits)
	if l.accept(".") {
		l.acceptRun(digits)
	}

	// Next thing mustn't be alphanumeric.
	if isAlphaNumeric(l.peek()) {
		l.next()
		return false
	}
	return true
}

// errorf returns an error token and terminates the scan by passing
// back a nil pointer that will be the next state, terminating l.nextItem.
func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- token{ERROR, l.start, fmt.Sprintf(format, args...)}
	return nil
}

func main() {
  testCase := "final override config_key = \"23\\\"4\\\"5\"// hello world\noverride c = 1234";
	// tokens := make(chan *token, 30)
  l := lex("testLexer", testCase)
  for t := range l.items {
    fmt.Println("Lexed token:", t)
  }
}
