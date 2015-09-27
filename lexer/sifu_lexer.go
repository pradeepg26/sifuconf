package lexer

import (
  "fmt"
  "strings"
  "unicode"
  "unicode/utf8"
)

type ItemType int

const (
  KEYWORD ItemType = iota
  VARIABLE
  TYPE
  ASSIGNMENT
  BOOLEAN
  NUMBER
  STRING
  LIST_START
  LIST_END
  COMMENT
  ERROR
)

const eof = -1

type Token struct {
  ItemType ItemType
  Pos int
  ItemValue string
}

func (t ItemType) String() string {
  switch t {
  case KEYWORD:
    return "KEYWORD"
  case VARIABLE:
    return "VARIABLE"
  case TYPE:
    return "TYPE"
  case ASSIGNMENT:
    return "ASSIGNMENT"
  case BOOLEAN:
    return "BOOLEAN"
  case NUMBER:
    return "NUMBER"
  case STRING:
    return "STRING"
  case COMMENT:
    return "COMMENT"
  case ERROR:
    return "ERROR"
  default:
    return "UNKNOWN_TYPE"
  }
}

func (t Token) String() string {
  switch t.ItemType {
  case ERROR:
    return fmt.Sprintf("{type: '%s', value: '%s'}", t.ItemType, t.ItemValue)
  case COMMENT:
    return fmt.Sprintf("{type: '%s', value: '%.20s...'}", t.ItemType, t.ItemValue)
  default:
    return fmt.Sprintf("{type: '%s', value: '%s'}", t.ItemType, t.ItemValue)
  }
}

type Lexer struct {
  Name string
  Input string
  Start int
  Pos int
  Width int
  Items chan Token
}

func Lex(name, data string) (*Lexer) {
  l := &Lexer{
    Name: name,
    Input: data,
    Items: make(chan Token),
  }
  go l.run()
  return l
}

// ------------------------- Lexer Methods -------------------------
func (l *Lexer) run() {
  for state := lexLine; state != nil; {
    state = state(l)
  }
  close(l.Items)
}

func (l *Lexer) emit(t ItemType) {
  token := Token{t, l.Start, l.Input[l.Start : l.Pos]}
  // fmt.Printf("Emit: (%s) - %s\n", l.Name, token)
  l.Items <- token
  l.Start = l.Pos
}

// next returns the next rune in the input.
func (l *Lexer) next() rune {
  if l.Pos >= len(l.Input) {
    l.Width = 0
    return eof
  }
  r, w := utf8.DecodeRuneInString(l.Input[l.Pos:])
  l.Width = w
  l.Pos += l.Width
  return r
}

// peek returns but does not consume the next rune in the input.
func (l *Lexer) peek() rune {
  r := l.next()
  l.backup()
  return r
}

// backup steps back one rune. Can only be called once per call of next.
func (l *Lexer) backup() {
  l.Pos -= l.Width
}

// accept consumes the next rune if it's from the valid set.
func (l *Lexer) accept(valid string) bool {
  if strings.IndexRune(valid, l.next()) >= 0 {
    return true
  }
  l.backup()
  return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *Lexer) acceptRun(valid string) {
  for strings.IndexRune(valid, l.next()) >= 0 {
  }
  l.backup()
}

// ignore skips over the pending input before this point.
func (l *Lexer) ignore() {
  l.Start = l.Pos
}

// lineNumber reports which line we're on, based on the Position of
// the previous item returned by nextItem. Doing it this way
// means we don't have to worry about peek double counting.
// func (l lexer) lineNumber() int {
  // return 1 + strings.Count(l.Input[:l.lastPos], "\n")
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

func isAsciiAlpha(r rune) bool {
    return r >= 'a' && r <= 'z'
}

// ------------------------- State Functions -------------------------
type stateFn func(*Lexer) stateFn

func lexLine(l *Lexer) stateFn {
  l.consumeSpaces() // Eat any whitespace at the beginning of the line

  switch r := l.next(); {
  case r == eof:
    return nil
  case isEndOfLine(r):
    return lexLine
  }
  l.backup()

  // If you see a "//", lex a comment string
  if l.scanComment() {
    l.emit(COMMENT)
    l.next() // Eat EOL
    return lexLine
  }

  if isAsciiAlpha(l.peek()) {
    return lexKeyword
  }

  return l.errorf("Must be a comment, keyword or variable")
}

func lexKeyword(l *Lexer) stateFn {
  l.consumeSpaces()

  keywords := []string{"final", "override", "required"}
  if !l.accept("abcdefghijklmnopqrstuvwxyz") {
    // Keywords and Variable names must start with a letter
    return l.errorf("Illegal character encountered. " +
       "Expecting keyword or variable")
  }

  l.backup() // backup and see if it's a keyword
  for _, k := range keywords {
    if strings.HasPrefix(l.Input[l.Pos:], k) {
      l.Pos += len(k)
      // Next character must be a space for it to be a keyword
      if l.peek() != ' ' {
        // Not a keyword, back up and try to parse a variable name
        l.Pos -= len(k)
        return lexVariableName
      }
      l.emit(KEYWORD)
      return lexKeyword
    }
  }

  // Not a keyword... but after a keyword a variable name must follow
  return lexVariableName
}

func lexVariableName(l *Lexer) stateFn {
  l.consumeSpaces()
  l.acceptRun("abcdefghijklmnopqrstuvwxyz01234567890_.")
  if l.peek() != ' ' {
    return l.errorf("A variable name must be terminated with a space. " +
       "Encountered '%s'", string(l.peek()))
  }
  l.emit(VARIABLE)
  return lexAssign
}

func lexAssign(l *Lexer) stateFn {
  l.consumeSpaces()
  if l.accept("=") {
    l.emit(ASSIGNMENT)
    return lexValue
  }
  return l.errorf("Expecting '='. Encountered '%s'", string(l.peek()))
}

func lexValue(l *Lexer) stateFn {
  l.consumeSpaces()
  switch r := l.peek(); {
  case r == '"':
    return lexStringValue
  case r == 't' || r == 'f':
    if l.scanBoolean() {
      l.emit(BOOLEAN)
      return lexPostValue
    }
    // Not a boolean, so fallthrough
    fallthrough
  case isAsciiAlpha(r):
    return lexTypeValue
  case r == '[':
    return lexList
  case unicode.IsDigit(r) || r == '.':
    return lexNumberValue
  default:
    return l.errorf("Illegal character")
  }
}

func (l *Lexer) scanBoolean() bool {
  if strings.HasPrefix(l.Input[l.Pos:], "true") {
    l.Pos += len("true")
    if isAlphaNumeric(l.peek()) {
      // After a boolean, next character must not be an alphanumeric
      // Backup
      l.Pos -= len("true")
      return false
    }
    return true
  } else if strings.HasPrefix(l.Input[l.Pos:], "false") {
    l.Pos += len("false")
    if isAlphaNumeric(l.peek()) {
      l.Pos -= len("false")
      return false
    }
    return true
  }
  return false
}

func lexTypeValue(l *Lexer) stateFn {
  // types := []string{"string", "integer", "float", "boolean"}
  return lexPostValue
}

// lexString scans a quoted string.
func lexStringValue(l *Lexer) stateFn {
  if l.scanString() {
    l.emit(STRING)
    return lexPostValue
  } else {
    return l.errorf("Unterminated string found")
  }
}

func (l *Lexer) scanString() bool {
  if l.next() != '"' {
    return false
  }

  for {
    switch l.next() {
    case '\\':
      if r := l.next(); r != eof && !isEndOfLine(r) {
        break
      }
      fallthrough
    case eof, '\n':
      return false
    case '"':
      return true
    }
  }
}

// lexNumberValue scans a number: decimal, or float.
func lexNumberValue(l *Lexer) stateFn {
  if !l.scanNumber() {
    return l.errorf("bad number syntax: %q", l.Input[l.Start:l.Pos])
  }
  l.emit(NUMBER)
  return lexPostValue
}

func (l *Lexer) scanNumber() bool {
  // Optional leading sign.
  l.accept("+-")
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

func lexList(l *Lexer) stateFn {
  l.consumeSpaces()
  if !l.accept("[") {
    l.errorf("Unexpected character. Expecting LIST_START")
  }

  for r := l.peek(); r != ']'; {
    switch r := r; {
    case r == '"':
      if l.scanString() {
        l.emit(STRING)
        // Try to accept a ','
        if !l.accept(",") {
          return l.errorf("List elements must be terminated by a comma ','")
        }
        // Scan for an optional comment
      }
    case r == 't' || r == 'f':
      if l.scanBoolean() {
        l.emit(BOOLEAN)
        return lexPostValue
      }
      // Not a boolean, so fallthrough
      fallthrough
    case isAsciiAlpha(r):
      return lexTypeValue
    case r == '[':
      return lexList
    case unicode.IsDigit(r) || r == '.':
      return lexNumberValue
    default:
      return l.errorf("Illegal character")
    }
  }
  return lexPostValue
}

func lexListElement(l *Lexer) stateFn {
  switch r := l.next(); {
  case r == '"':
      if l.scanString() {
        l.emit(STRING)
        return lexPostListElement
      }
      return l.errorf("Unterminated string")
    case r == 't' || r == 'f':
      if l.scanBoolean() {
        l.emit(BOOLEAN)
        return lexPostValue
      }
      // Not a boolean, so fallthrough
      fallthrough
    case unicode.IsDigit(r) || r == '.':
      if l.scanNumber() {
        l.emit(NUMBER)
        return lexPostListElement
      }
      return l.errorf("Invalid number character")
    default:
      return l.errorf("Illegal character")
  }
}

func lexPostListElement(l *Lexer) stateFn {
  l.consumeSpaces()
  // List has been closed
  if l.accept("]") {
    l.emit(LIST_END)
    return lexPostValue
  }

  if !l.accept(",") {
    return l.errorf("List elements must be terminated by a comma ','")
  }

  return lexListElement
}

func lexPostValue(l *Lexer) stateFn {
  l.consumeSpaces()

  if l.scanComment() {
    l.emit(COMMENT)
    l.next() // Eat EOL
    return lexLine
  }

  if r := l.peek(); !isEndOfLine(r) && r != eof {
    return l.errorf("Expecting a comment or a newline or EOF. Received '%s'", string(r))
  }

  return lexLine
}

func (l *Lexer) scanComment() bool {
  if strings.HasPrefix(l.Input[l.Pos:], "//") {
    // Comments are always terminated at EOL or EOF
    for !isEndOfLine(l.peek()) {
      if (l.next() == eof) {
        break
      }
    }
    return true
  }
  return false
}

// Consumes all whitespace characters
func (l *Lexer) consumeSpaces() {
  for isSpace(l.next()) {
  }
  l.backup()
  l.ignore()
}

// errorf returns an error Token and terminates the scan by passing
// back a nil pointer that will be the next state, terminating l.nextItem.
func (l *Lexer) errorf(format string, args ...interface{}) stateFn {
  l.Items <- Token{ERROR, l.Start, fmt.Sprintf(format, args...)}
  return nil
}
