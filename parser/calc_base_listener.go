// Generated from calc.g4 by ANTLR 4.7.

package parser // calc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasecalcListener is a complete listener for a parse tree produced by calcParser.
type BasecalcListener struct{}

var _ calcListener = &BasecalcListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasecalcListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasecalcListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasecalcListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasecalcListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BasecalcListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BasecalcListener) ExitStart(ctx *StartContext) {}

// EnterNumber is called when production Number is entered.
func (s *BasecalcListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BasecalcListener) ExitNumber(ctx *NumberContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BasecalcListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BasecalcListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BasecalcListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BasecalcListener) ExitAddSub(ctx *AddSubContext) {}
