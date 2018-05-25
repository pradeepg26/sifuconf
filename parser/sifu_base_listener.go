// Generated from Sifu.g4 by ANTLR 4.6.

package parser // Sifu

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSifuListener is a complete listener for a parse tree produced by SifuParser.
type BaseSifuListener struct{}

var _ SifuListener = &BaseSifuListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSifuListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSifuListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSifuListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSifuListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSifu is called when production sifu is entered.
func (s *BaseSifuListener) EnterSifu(ctx *SifuContext) {}

// ExitSifu is called when production sifu is exited.
func (s *BaseSifuListener) ExitSifu(ctx *SifuContext) {}
