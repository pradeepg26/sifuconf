package lexer

import (
  // "fmt"
  "testing"
  "github.com/stretchr/testify/assert"
)

func noMoreTokens(t *testing.T, l *Lexer) {
  for tok := range l.Items {
    t.Errorf("Unexpected token received '%s'", tok)
  }
}

func TestAssignBoolean(t *testing.T) {
  t.Parallel()
  assert := assert.New(t)

  l := Lex("TestAssignBoolean", `hello = true`)

  token := <-l.Items
  assert.Equal(VARIABLE, token.ItemType)
  assert.Equal("hello", token.ItemValue)

  token = <-l.Items
  assert.Equal(ASSIGNMENT, token.ItemType)
  assert.Equal("=", token.ItemValue)

  token = <-l.Items
  assert.Equal(BOOLEAN, token.ItemType)
  assert.Equal("true", token.ItemValue)

  noMoreTokens(t, l)
}

func TestAssignInteger(t *testing.T) {
  t.Parallel()
  assert := assert.New(t)

  l := Lex("TestAssignInteger", `hello = 10`)

  token := <-l.Items
  assert.Equal(VARIABLE, token.ItemType)
  assert.Equal("hello", token.ItemValue)

  token = <-l.Items
  assert.Equal(ASSIGNMENT, token.ItemType)
  assert.Equal("=", token.ItemValue)

  token = <-l.Items
  assert.Equal(NUMBER, token.ItemType)
  assert.Equal("10", token.ItemValue)

  noMoreTokens(t, l)
}

func TestAssignFloat(t *testing.T) {
  t.Parallel()
  assert := assert.New(t)

  l := Lex("TestAssignFloat", `hello = 1.0`)

  token := <-l.Items
  assert.Equal(VARIABLE, token.ItemType)
  assert.Equal("hello", token.ItemValue)

  token = <-l.Items
  assert.Equal(ASSIGNMENT, token.ItemType)
  assert.Equal("=", token.ItemValue)

  token = <-l.Items
  assert.Equal(NUMBER, token.ItemType)
  assert.Equal("1.0", token.ItemValue)

  noMoreTokens(t, l)
}

func TestAssignString(t *testing.T) {
  t.Parallel()
  assert := assert.New(t)

  l := Lex("TestAssignString", `hello = "world"`)

  token := <-l.Items
  assert.Equal(VARIABLE, token.ItemType)
  assert.Equal("hello", token.ItemValue)

  token = <-l.Items
  assert.Equal(ASSIGNMENT, token.ItemType)
  assert.Equal("=", token.ItemValue)

  token = <-l.Items
  assert.Equal(STRING, token.ItemType)
  assert.Equal(`"world"`, token.ItemValue)

  noMoreTokens(t, l)
}

func TestCommentAfterAssignBoolean(t *testing.T) {
  t.Parallel()
  assert := assert.New(t)

  l := Lex("TestCommentAfterAssignBoolean", `hello = false // this is a bool`)

  token := <-l.Items
  assert.Equal(VARIABLE, token.ItemType)
  assert.Equal("hello", token.ItemValue)

  token = <-l.Items
  assert.Equal(ASSIGNMENT, token.ItemType)
  assert.Equal("=", token.ItemValue)

  token = <-l.Items
  assert.Equal(BOOLEAN, token.ItemType)
  assert.Equal("false", token.ItemValue)

  token = <-l.Items
  assert.Equal(COMMENT, token.ItemType)
  assert.Equal("// this is a bool", token.ItemValue)

  noMoreTokens(t, l)
}

func TestCommentAfterAssignInteger(t *testing.T) {
  t.Parallel()
  assert := assert.New(t)

  l := Lex("TestCommentAfterAssignInteger", `hello = 10 // this is an int`)

  token := <-l.Items
  assert.Equal(VARIABLE, token.ItemType)
  assert.Equal("hello", token.ItemValue)

  token = <-l.Items
  assert.Equal(ASSIGNMENT, token.ItemType)
  assert.Equal("=", token.ItemValue)

  token = <-l.Items
  assert.Equal(NUMBER, token.ItemType)
  assert.Equal("10", token.ItemValue)

  token = <-l.Items
  assert.Equal(COMMENT, token.ItemType)
  assert.Equal("// this is an int", token.ItemValue)

  noMoreTokens(t, l)
}

func TestCommentAfterAssignFloat(t *testing.T) {
  t.Parallel()
  assert := assert.New(t)

  l := Lex("TestCommentAfterAssignFloat", `hello = 1.0 // this is a float`)

  token := <-l.Items
  assert.Equal(VARIABLE, token.ItemType)
  assert.Equal("hello", token.ItemValue)

  token = <-l.Items
  assert.Equal(ASSIGNMENT, token.ItemType)
  assert.Equal("=", token.ItemValue)

  token = <-l.Items
  assert.Equal(NUMBER, token.ItemType)
  assert.Equal("1.0", token.ItemValue)

  token = <-l.Items
  assert.Equal(COMMENT, token.ItemType)
  assert.Equal("// this is a float", token.ItemValue)

  noMoreTokens(t, l)
}

func TestCommentAfterAssignString(t *testing.T) {
  t.Parallel()
  assert := assert.New(t)

  l := Lex("TestCommentAfterAssignFloat", `hello = "string"// this is a string`)

  token := <-l.Items
  assert.Equal(VARIABLE, token.ItemType)
  assert.Equal("hello", token.ItemValue)

  token = <-l.Items
  assert.Equal(ASSIGNMENT, token.ItemType)
  assert.Equal("=", token.ItemValue)

  token = <-l.Items
  assert.Equal(STRING, token.ItemType)
  assert.Equal(`"string"`, token.ItemValue)

  token = <-l.Items
  assert.Equal(COMMENT, token.ItemType)
  assert.Equal("// this is a string", token.ItemValue)

  noMoreTokens(t, l)
}

func TestQuotesInStringValue(t *testing.T) {
  t.Parallel()
  assert := assert.New(t)

  l := Lex("TestCommentAfterAssignFloat", `hello = "st\"r\"ing"`)

  token := <-l.Items
  assert.Equal(VARIABLE, token.ItemType)
  assert.Equal("hello", token.ItemValue)

  token = <-l.Items
  assert.Equal(ASSIGNMENT, token.ItemType)
  assert.Equal("=", token.ItemValue)

  token = <-l.Items
  assert.Equal(STRING, token.ItemType)
  assert.Equal(`"st\"r\"ing"`, token.ItemValue)

  noMoreTokens(t, l)
}

func TestSingleLineListInts(t *testing.T) {
  t.Parallel()
  assert := assert.New(t)

  l := Lex("TestSingleLineListInts", "list = [1, 2, 3, 4]")

  token := <-l.Items
  assert.Equal(VARIABLE, token.ItemType)
  assert.Equal("list", token.ItemValue)

  token = <-l.Items
  assert.Equal(ASSIGNMENT, token.ItemType)
  assert.Equal("=", token.ItemValue)

  token = <-l.Items
  assert.Equal(LIST_START, token.ItemType)
  assert.Equal("[", token.ItemValue)

  token = <-l.Items
  assert.Equal(NUMBER, token.ItemType)
  assert.Equal("1", token.ItemValue)

  token = <-l.Items
  assert.Equal(NUMBER, token.ItemType)
  assert.Equal("2", token.ItemValue)

  token = <-l.Items
  assert.Equal(NUMBER, token.ItemType)
  assert.Equal("3", token.ItemValue)

  token = <-l.Items
  assert.Equal(NUMBER, token.ItemType)
  assert.Equal("4", token.ItemValue)

  token = <-l.Items
  assert.Equal(LIST_END, token.ItemType)
  assert.Equal("]", token.ItemValue)

  noMoreTokens(t, l)
}

func TestSingleLineListStrings(t *testing.T) {
  t.Parallel()
  assert := assert.New(t)

  l := Lex("TestSingleLineListStrings", `list = ["12", "23", "34", "45"]`)

  token := <-l.Items
  assert.Equal(VARIABLE, token.ItemType)
  assert.Equal("list", token.ItemValue)

  token = <-l.Items
  assert.Equal(ASSIGNMENT, token.ItemType)
  assert.Equal("=", token.ItemValue)

  token = <-l.Items
  assert.Equal(LIST_START, token.ItemType)
  assert.Equal("[", token.ItemValue)

  token = <-l.Items
  assert.Equal(STRING, token.ItemType)
  assert.Equal(`"12"`, token.ItemValue)

  token = <-l.Items
  assert.Equal(STRING, token.ItemType)
  assert.Equal(`"23"`, token.ItemValue)

  token = <-l.Items
  assert.Equal(STRING, token.ItemType)
  assert.Equal(`"34"`, token.ItemValue)

  token = <-l.Items
  assert.Equal(STRING, token.ItemType)
  assert.Equal(`"45"`, token.ItemValue)

  token = <-l.Items
  assert.Equal(LIST_END, token.ItemType)
  assert.Equal("]", token.ItemValue)

  noMoreTokens(t, l)
}

func TestMultiLineListStrings(t *testing.T) {
  t.Parallel()
  assert := assert.New(t)

  l := Lex("TestMultiLineListStrings", `multiline_list = [
                                          "e1",
                                          "e2",
                                          "e3"
                                        ] // Multiline lists`);

  token := <-l.Items
  assert.Equal(VARIABLE, token.ItemType)
  assert.Equal("multiline_list", token.ItemValue)

  token = <-l.Items
  assert.Equal(ASSIGNMENT, token.ItemType)
  assert.Equal("=", token.ItemValue)

  token = <-l.Items
  assert.Equal(LIST_START, token.ItemType)
  assert.Equal("[", token.ItemValue)

  token = <-l.Items
  assert.Equal(STRING, token.ItemType)
  assert.Equal(`"e1"`, token.ItemValue)

  token = <-l.Items
  assert.Equal(STRING, token.ItemType)
  assert.Equal(`"e2"`, token.ItemValue)

  token = <-l.Items
  assert.Equal(STRING, token.ItemType)
  assert.Equal(`"e3"`, token.ItemValue)

  token = <-l.Items
  assert.Equal(LIST_END, token.ItemType)
  assert.Equal("]", token.ItemValue)

  token = <-l.Items
  assert.Equal(COMMENT, token.ItemType)
  assert.Equal("// Multiline lists", token.ItemValue)

  noMoreTokens(t, l)
}

func TestMultiLineListWithMultipleElements(t *testing.T) {
  t.Parallel()
  assert := assert.New(t)

  l := Lex("TestMultiLineListWithMultipleElements",
     `multiline_list = [
        "e1", "e2", // Multiple elements on one line
        "e3"
      ] // Multiline lists`);

  token := <-l.Items
  assert.Equal(VARIABLE, token.ItemType)
  assert.Equal("multiline_list", token.ItemValue)

  token = <-l.Items
  assert.Equal(ASSIGNMENT, token.ItemType)
  assert.Equal("=", token.ItemValue)

  token = <-l.Items
  assert.Equal(LIST_START, token.ItemType)
  assert.Equal("[", token.ItemValue)

  token = <-l.Items
  assert.Equal(STRING, token.ItemType)
  assert.Equal(`"e1"`, token.ItemValue)

  token = <-l.Items
  assert.Equal(STRING, token.ItemType)
  assert.Equal(`"e2"`, token.ItemValue)

  token = <-l.Items
  assert.Equal(COMMENT, token.ItemType)
  assert.Equal(`// Multiple elements on one line`, token.ItemValue)

  token = <-l.Items
  assert.Equal(STRING, token.ItemType)
  assert.Equal(`"e3"`, token.ItemValue)

  token = <-l.Items
  assert.Equal(LIST_END, token.ItemType)
  assert.Equal("]", token.ItemValue)

  token = <-l.Items
  assert.Equal(COMMENT, token.ItemType)
  assert.Equal("// Multiline lists", token.ItemValue)

  noMoreTokens(t, l)
}

func TestMultiLineListStringsWithComments(t *testing.T) {
  t.Parallel()
  assert := assert.New(t)

  l := Lex("TestMultiLineListStringsWithComments",
     `multiline_list = [
        "e1", // First element
        "e2", // Second element
        "e3" // Last element
      ] // Multiline lists`);

  token := <-l.Items
  assert.Equal(VARIABLE, token.ItemType)
  assert.Equal("multiline_list", token.ItemValue)

  token = <-l.Items
  assert.Equal(ASSIGNMENT, token.ItemType)
  assert.Equal("=", token.ItemValue)

  token = <-l.Items
  assert.Equal(LIST_START, token.ItemType)
  assert.Equal("[", token.ItemValue)

  token = <-l.Items
  assert.Equal(STRING, token.ItemType)
  assert.Equal(`"e1"`, token.ItemValue)

  token = <-l.Items
  assert.Equal(COMMENT, token.ItemType)
  assert.Equal(`// First element`, token.ItemValue)

  token = <-l.Items
  assert.Equal(STRING, token.ItemType)
  assert.Equal(`"e2"`, token.ItemValue)

  token = <-l.Items
  assert.Equal(COMMENT, token.ItemType)
  assert.Equal(`// Second element`, token.ItemValue)

  token = <-l.Items
  assert.Equal(STRING, token.ItemType)
  assert.Equal(`"e3"`, token.ItemValue)

  token = <-l.Items
  assert.Equal(COMMENT, token.ItemType)
  assert.Equal(`// Last element`, token.ItemValue)

  token = <-l.Items
  assert.Equal(LIST_END, token.ItemType)
  assert.Equal("]", token.ItemValue)

  token = <-l.Items
  assert.Equal(COMMENT, token.ItemType)
  assert.Equal("// Multiline lists", token.ItemValue)

  noMoreTokens(t, l)
}

func TestListWithElementsOnFirstLine(t *testing.T) {
  t.Parallel()
  assert := assert.New(t)

  l := Lex("TestListWithElementsOnFirstLine",
     `multiline_list = [ "e1", // One!
        "e2", "e3" // Last element
      ] // Multiline lists`);

  token := <-l.Items
  assert.Equal(VARIABLE, token.ItemType)
  assert.Equal("multiline_list", token.ItemValue)

  token = <-l.Items
  assert.Equal(ASSIGNMENT, token.ItemType)
  assert.Equal("=", token.ItemValue)

  token = <-l.Items
  assert.Equal(LIST_START, token.ItemType)
  assert.Equal("[", token.ItemValue)

  token = <-l.Items
  assert.Equal(STRING, token.ItemType)
  assert.Equal(`"e1"`, token.ItemValue)

  token = <-l.Items
  assert.Equal(COMMENT, token.ItemType)
  assert.Equal(`// One!`, token.ItemValue)

  token = <-l.Items
  assert.Equal(STRING, token.ItemType)
  assert.Equal(`"e2"`, token.ItemValue)

  token = <-l.Items
  assert.Equal(STRING, token.ItemType)
  assert.Equal(`"e3"`, token.ItemValue)

  token = <-l.Items
  assert.Equal(COMMENT, token.ItemType)
  assert.Equal(`// Last element`, token.ItemValue)

  token = <-l.Items
  assert.Equal(LIST_END, token.ItemType)
  assert.Equal("]", token.ItemValue)

  token = <-l.Items
  assert.Equal(COMMENT, token.ItemType)
  assert.Equal("// Multiline lists", token.ItemValue)

  noMoreTokens(t, l)
}

func TestListFloats(t *testing.T) {
  t.Parallel()
  assert := assert.New(t)

  l := Lex("TestListFloats",
    `list = [0.123, .34,
      4.946, // comment
      1234.0
    ] // Another comment`);

  token := <-l.Items
  assert.Equal(VARIABLE, token.ItemType)
  assert.Equal("list", token.ItemValue)

  token = <-l.Items
  assert.Equal(ASSIGNMENT, token.ItemType)
  assert.Equal("=", token.ItemValue)

  token = <-l.Items
  assert.Equal(LIST_START, token.ItemType)
  assert.Equal("[", token.ItemValue)

  token = <-l.Items
  assert.Equal(NUMBER, token.ItemType)
  assert.Equal(`0.123`, token.ItemValue)

  token = <-l.Items
  assert.Equal(NUMBER, token.ItemType)
  assert.Equal(`.34`, token.ItemValue)

  token = <-l.Items
  assert.Equal(NUMBER, token.ItemType)
  assert.Equal(`4.946`, token.ItemValue)

  token = <-l.Items
  assert.Equal(COMMENT, token.ItemType)
  assert.Equal(`// comment`, token.ItemValue)

  token = <-l.Items
  assert.Equal(NUMBER, token.ItemType)
  assert.Equal(`1234.0`, token.ItemValue)

  token = <-l.Items
  assert.Equal(LIST_END, token.ItemType)
  assert.Equal("]", token.ItemValue)

  token = <-l.Items
  assert.Equal(COMMENT, token.ItemType)
  assert.Equal("// Another comment", token.ItemValue)

}

func TestUnterminatedString(t *testing.T) {
  t.Parallel()
  assert := assert.New(t)

  l := Lex("TestUnterminatedString", `hello = "unterminated string`)

  token := <-l.Items
  assert.Equal(VARIABLE, token.ItemType)
  assert.Equal("hello", token.ItemValue)

  token = <-l.Items
  assert.Equal(ASSIGNMENT, token.ItemType)
  assert.Equal("=", token.ItemValue)

  token = <-l.Items
  assert.Equal(ERROR, token.ItemType)

  noMoreTokens(t, l)
}

func TestEmptyLines(t *testing.T) {
  t.Parallel()
  assert := assert.New(t)

  l := Lex("TestEmptyLines", `hello = "string"

    whosits = 1234
  `)

  token := <-l.Items
  assert.Equal(VARIABLE, token.ItemType)
  assert.Equal("hello", token.ItemValue)

  token = <-l.Items
  assert.Equal(ASSIGNMENT, token.ItemType)
  assert.Equal("=", token.ItemValue)

  token = <-l.Items
  assert.Equal(STRING, token.ItemType)
  assert.Equal(`"string"`, token.ItemValue)

  token = <-l.Items
  assert.Equal(VARIABLE, token.ItemType)
  assert.Equal("whosits", token.ItemValue)

  token = <-l.Items
  assert.Equal(ASSIGNMENT, token.ItemType)
  assert.Equal("=", token.ItemValue)

  token = <-l.Items
  assert.Equal(NUMBER, token.ItemType)
  assert.Equal("1234", token.ItemValue)

  noMoreTokens(t, l)
}

func TestCommentOnly(t *testing.T) {
  t.Parallel()
  assert := assert.New(t)

  l := Lex("TestCommentOnly", "// hello world")

  token := <-l.Items
  assert.Equal(COMMENT, token.ItemType)
  assert.Equal("// hello world", token.ItemValue)

  noMoreTokens(t, l)
}
