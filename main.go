package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/pradeepg26/sifuconf/parser"
	"os"
	"fmt"
)

type TreeShapeListener struct {
	*parser.BaseSifuListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("enter %v\n", ctx.GetText())
}

func (this *TreeShapeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("exit %v\n", ctx.GetText())
}

func (this *TreeShapeListener) VisitErrorNode(err antlr.ErrorNode) {
	fmt.Printf("%v\n", err)
}

func main() {
	input := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewSifuLexer(input)
	stream := antlr.NewCommonTokenStream(lexer,0)
	p := parser.NewSifuParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Sifu()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
