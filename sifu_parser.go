package main

import (
  "github.com/fatih/color"
  "github.com/pradeepg26/sifuconf/lexer"
  "fmt"
  "io/ioutil"
)

func main() {
  dat, _ := ioutil.ReadFile("sample.sifu")
  conf := string(dat)

  l := lexer.Lex("sample.sifu", conf)

  cyan := color.New(color.FgCyan).SprintFunc()
  red := color.New(color.FgRed).SprintFunc()

  for t := range l.Items {
    if (t.ItemType == lexer.ERROR) {
      fmt.Printf("%s%s\n", cyan(l.Input[0:l.Pos]), red(l.Input[l.Pos:]))
    }
  }
}
