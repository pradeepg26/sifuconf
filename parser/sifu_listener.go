// Generated from Sifu.g4 by ANTLR 4.6.

package parser // Sifu

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SifuListener is a complete listener for a parse tree produced by SifuParser.
type SifuListener interface {
	antlr.ParseTreeListener

	// EnterSifu is called when entering the sifu production.
	EnterSifu(c *SifuContext)

	// ExitSifu is called when exiting the sifu production.
	ExitSifu(c *SifuContext)
}
