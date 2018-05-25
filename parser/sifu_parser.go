// Generated from Sifu.g4 by ANTLR 4.6.

package parser // Sifu

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 1072, 54993, 33286, 44333, 17431, 44785, 36224, 43741, 3, 8, 10, 4,
	2, 9, 2, 3, 2, 6, 2, 6, 10, 2, 13, 2, 14, 2, 7, 3, 2, 2, 2, 3, 2, 2, 2,
	9, 2, 5, 3, 2, 2, 2, 4, 6, 7, 5, 2, 2, 5, 4, 3, 2, 2, 2, 6, 7, 3, 2, 2,
	2, 7, 5, 3, 2, 2, 2, 7, 8, 3, 2, 2, 2, 8, 3, 3, 2, 2, 2, 3, 7,
}

var deserializer = antlr.NewATNDeserializer(nil)

var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames []string

var symbolicNames = []string{
	"", "VARIABLE", "NUMBER", "ITEM", "ARRAY", "BOOLEAN", "WS",
}

var ruleNames = []string{
	"sifu",
}

type SifuParser struct {
	*antlr.BaseParser
}

func NewSifuParser(input antlr.TokenStream) *SifuParser {
	var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	var sharedContextCache = antlr.NewPredictionContextCache()

	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}

	this := new(SifuParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, sharedContextCache)
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Sifu.g4"

	return this
}

// SifuParser tokens.
const (
	SifuParserEOF      = antlr.TokenEOF
	SifuParserVARIABLE = 1
	SifuParserNUMBER   = 2
	SifuParserITEM     = 3
	SifuParserARRAY    = 4
	SifuParserBOOLEAN  = 5
	SifuParserWS       = 6
)

// SifuParserRULE_sifu is the SifuParser rule.
const SifuParserRULE_sifu = 0

// ISifuContext is an interface to support dynamic dispatch.
type ISifuContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSifuContext differentiates from other interfaces.
	IsSifuContext()
}

type SifuContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySifuContext() *SifuContext {
	var p = new(SifuContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SifuParserRULE_sifu
	return p
}

func (*SifuContext) IsSifuContext() {}

func NewSifuContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SifuContext {
	var p = new(SifuContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SifuParserRULE_sifu

	return p
}

func (s *SifuContext) GetParser() antlr.Parser { return s.parser }

func (s *SifuContext) AllITEM() []antlr.TerminalNode {
	return s.GetTokens(SifuParserITEM)
}

func (s *SifuContext) ITEM(i int) antlr.TerminalNode {
	return s.GetToken(SifuParserITEM, i)
}

func (s *SifuContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SifuContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SifuContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SifuListener); ok {
		listenerT.EnterSifu(s)
	}
}

func (s *SifuContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SifuListener); ok {
		listenerT.ExitSifu(s)
	}
}

func (p *SifuParser) Sifu() (localctx ISifuContext) {
	localctx = NewSifuContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SifuParserRULE_sifu)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(3)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SifuParserITEM {
		{
			p.SetState(2)
			p.Match(SifuParserITEM)
		}

		p.SetState(5)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
