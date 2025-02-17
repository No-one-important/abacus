// Code generated from /home/vikimaster2/Projects/Go/abacus/Abacus.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Abacus

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAbacusListener is a complete listener for a parse tree produced by AbacusParser.
type BaseAbacusListener struct{}

var _ AbacusListener = &BaseAbacusListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAbacusListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAbacusListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAbacusListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAbacusListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseAbacusListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseAbacusListener) ExitRoot(ctx *RootContext) {}

// EnterVariableDeclaration is called when production VariableDeclaration is entered.
func (s *BaseAbacusListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production VariableDeclaration is exited.
func (s *BaseAbacusListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterLambdaDeclaration is called when production LambdaDeclaration is entered.
func (s *BaseAbacusListener) EnterLambdaDeclaration(ctx *LambdaDeclarationContext) {}

// ExitLambdaDeclaration is called when production LambdaDeclaration is exited.
func (s *BaseAbacusListener) ExitLambdaDeclaration(ctx *LambdaDeclarationContext) {}

// EnterNot is called when production Not is entered.
func (s *BaseAbacusListener) EnterNot(ctx *NotContext) {}

// ExitNot is called when production Not is exited.
func (s *BaseAbacusListener) ExitNot(ctx *NotContext) {}

// EnterGreaterOrEqualComparison is called when production GreaterOrEqualComparison is entered.
func (s *BaseAbacusListener) EnterGreaterOrEqualComparison(ctx *GreaterOrEqualComparisonContext) {}

// ExitGreaterOrEqualComparison is called when production GreaterOrEqualComparison is exited.
func (s *BaseAbacusListener) ExitGreaterOrEqualComparison(ctx *GreaterOrEqualComparisonContext) {}

// EnterGreaterComparison is called when production GreaterComparison is entered.
func (s *BaseAbacusListener) EnterGreaterComparison(ctx *GreaterComparisonContext) {}

// ExitGreaterComparison is called when production GreaterComparison is exited.
func (s *BaseAbacusListener) ExitGreaterComparison(ctx *GreaterComparisonContext) {}

// EnterLessOrEqualComparison is called when production LessOrEqualComparison is entered.
func (s *BaseAbacusListener) EnterLessOrEqualComparison(ctx *LessOrEqualComparisonContext) {}

// ExitLessOrEqualComparison is called when production LessOrEqualComparison is exited.
func (s *BaseAbacusListener) ExitLessOrEqualComparison(ctx *LessOrEqualComparisonContext) {}

// EnterBooleanAtom is called when production BooleanAtom is entered.
func (s *BaseAbacusListener) EnterBooleanAtom(ctx *BooleanAtomContext) {}

// ExitBooleanAtom is called when production BooleanAtom is exited.
func (s *BaseAbacusListener) ExitBooleanAtom(ctx *BooleanAtomContext) {}

// EnterLessComparison is called when production LessComparison is entered.
func (s *BaseAbacusListener) EnterLessComparison(ctx *LessComparisonContext) {}

// ExitLessComparison is called when production LessComparison is exited.
func (s *BaseAbacusListener) ExitLessComparison(ctx *LessComparisonContext) {}

// EnterParenthesesBoolean is called when production ParenthesesBoolean is entered.
func (s *BaseAbacusListener) EnterParenthesesBoolean(ctx *ParenthesesBooleanContext) {}

// ExitParenthesesBoolean is called when production ParenthesesBoolean is exited.
func (s *BaseAbacusListener) ExitParenthesesBoolean(ctx *ParenthesesBooleanContext) {}

// EnterAndOrXor is called when production AndOrXor is entered.
func (s *BaseAbacusListener) EnterAndOrXor(ctx *AndOrXorContext) {}

// ExitAndOrXor is called when production AndOrXor is exited.
func (s *BaseAbacusListener) ExitAndOrXor(ctx *AndOrXorContext) {}

// EnterEqualComparison is called when production EqualComparison is entered.
func (s *BaseAbacusListener) EnterEqualComparison(ctx *EqualComparisonContext) {}

// ExitEqualComparison is called when production EqualComparison is exited.
func (s *BaseAbacusListener) ExitEqualComparison(ctx *EqualComparisonContext) {}

// EnterBoolAtom is called when production boolAtom is entered.
func (s *BaseAbacusListener) EnterBoolAtom(ctx *BoolAtomContext) {}

// ExitBoolAtom is called when production boolAtom is exited.
func (s *BaseAbacusListener) ExitBoolAtom(ctx *BoolAtomContext) {}

// EnterVariablesLambda is called when production VariablesLambda is entered.
func (s *BaseAbacusListener) EnterVariablesLambda(ctx *VariablesLambdaContext) {}

// ExitVariablesLambda is called when production VariablesLambda is exited.
func (s *BaseAbacusListener) ExitVariablesLambda(ctx *VariablesLambdaContext) {}

// EnterNullArityLambda is called when production NullArityLambda is entered.
func (s *BaseAbacusListener) EnterNullArityLambda(ctx *NullArityLambdaContext) {}

// ExitNullArityLambda is called when production NullArityLambda is exited.
func (s *BaseAbacusListener) ExitNullArityLambda(ctx *NullArityLambdaContext) {}

// EnterSignedExpr is called when production SignedExpr is entered.
func (s *BaseAbacusListener) EnterSignedExpr(ctx *SignedExprContext) {}

// ExitSignedExpr is called when production SignedExpr is exited.
func (s *BaseAbacusListener) ExitSignedExpr(ctx *SignedExprContext) {}

// EnterMod is called when production Mod is entered.
func (s *BaseAbacusListener) EnterMod(ctx *ModContext) {}

// ExitMod is called when production Mod is exited.
func (s *BaseAbacusListener) ExitMod(ctx *ModContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseAbacusListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseAbacusListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseAbacusListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseAbacusListener) ExitAddSub(ctx *AddSubContext) {}

// EnterPercent is called when production Percent is entered.
func (s *BaseAbacusListener) EnterPercent(ctx *PercentContext) {}

// ExitPercent is called when production Percent is exited.
func (s *BaseAbacusListener) ExitPercent(ctx *PercentContext) {}

// EnterLambdaExpr is called when production LambdaExpr is entered.
func (s *BaseAbacusListener) EnterLambdaExpr(ctx *LambdaExprContext) {}

// ExitLambdaExpr is called when production LambdaExpr is exited.
func (s *BaseAbacusListener) ExitLambdaExpr(ctx *LambdaExprContext) {}

// EnterPow is called when production Pow is entered.
func (s *BaseAbacusListener) EnterPow(ctx *PowContext) {}

// ExitPow is called when production Pow is exited.
func (s *BaseAbacusListener) ExitPow(ctx *PowContext) {}

// EnterAtomExpr is called when production AtomExpr is entered.
func (s *BaseAbacusListener) EnterAtomExpr(ctx *AtomExprContext) {}

// ExitAtomExpr is called when production AtomExpr is exited.
func (s *BaseAbacusListener) ExitAtomExpr(ctx *AtomExprContext) {}

// EnterParentheses is called when production Parentheses is entered.
func (s *BaseAbacusListener) EnterParentheses(ctx *ParenthesesContext) {}

// ExitParentheses is called when production Parentheses is exited.
func (s *BaseAbacusListener) ExitParentheses(ctx *ParenthesesContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseAbacusListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseAbacusListener) ExitParameter(ctx *ParameterContext) {}

// EnterRecursionParameters is called when production recursionParameters is entered.
func (s *BaseAbacusListener) EnterRecursionParameters(ctx *RecursionParametersContext) {}

// ExitRecursionParameters is called when production recursionParameters is exited.
func (s *BaseAbacusListener) ExitRecursionParameters(ctx *RecursionParametersContext) {}

// EnterMixedTuple is called when production mixedTuple is entered.
func (s *BaseAbacusListener) EnterMixedTuple(ctx *MixedTupleContext) {}

// ExitMixedTuple is called when production mixedTuple is exited.
func (s *BaseAbacusListener) ExitMixedTuple(ctx *MixedTupleContext) {}

// EnterTuple is called when production tuple is entered.
func (s *BaseAbacusListener) EnterTuple(ctx *TupleContext) {}

// ExitTuple is called when production tuple is exited.
func (s *BaseAbacusListener) ExitTuple(ctx *TupleContext) {}

// EnterLambdaArguments is called when production lambdaArguments is entered.
func (s *BaseAbacusListener) EnterLambdaArguments(ctx *LambdaArgumentsContext) {}

// ExitLambdaArguments is called when production lambdaArguments is exited.
func (s *BaseAbacusListener) ExitLambdaArguments(ctx *LambdaArgumentsContext) {}

// EnterVariablesTuple is called when production variablesTuple is entered.
func (s *BaseAbacusListener) EnterVariablesTuple(ctx *VariablesTupleContext) {}

// ExitVariablesTuple is called when production variablesTuple is exited.
func (s *BaseAbacusListener) ExitVariablesTuple(ctx *VariablesTupleContext) {}

// EnterFuncExpr is called when production FuncExpr is entered.
func (s *BaseAbacusListener) EnterFuncExpr(ctx *FuncExprContext) {}

// ExitFuncExpr is called when production FuncExpr is exited.
func (s *BaseAbacusListener) ExitFuncExpr(ctx *FuncExprContext) {}

// EnterConstant is called when production Constant is entered.
func (s *BaseAbacusListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production Constant is exited.
func (s *BaseAbacusListener) ExitConstant(ctx *ConstantContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseAbacusListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseAbacusListener) ExitNumber(ctx *NumberContext) {}

// EnterVariable is called when production Variable is entered.
func (s *BaseAbacusListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production Variable is exited.
func (s *BaseAbacusListener) ExitVariable(ctx *VariableContext) {}

// EnterPlusSign is called when production PlusSign is entered.
func (s *BaseAbacusListener) EnterPlusSign(ctx *PlusSignContext) {}

// ExitPlusSign is called when production PlusSign is exited.
func (s *BaseAbacusListener) ExitPlusSign(ctx *PlusSignContext) {}

// EnterMinusSign is called when production MinusSign is entered.
func (s *BaseAbacusListener) EnterMinusSign(ctx *MinusSignContext) {}

// ExitMinusSign is called when production MinusSign is exited.
func (s *BaseAbacusListener) ExitMinusSign(ctx *MinusSignContext) {}

// EnterSqrtFunction is called when production SqrtFunction is entered.
func (s *BaseAbacusListener) EnterSqrtFunction(ctx *SqrtFunctionContext) {}

// ExitSqrtFunction is called when production SqrtFunction is exited.
func (s *BaseAbacusListener) ExitSqrtFunction(ctx *SqrtFunctionContext) {}

// EnterCbrtFunction is called when production CbrtFunction is entered.
func (s *BaseAbacusListener) EnterCbrtFunction(ctx *CbrtFunctionContext) {}

// ExitCbrtFunction is called when production CbrtFunction is exited.
func (s *BaseAbacusListener) ExitCbrtFunction(ctx *CbrtFunctionContext) {}

// EnterLnFunction is called when production LnFunction is entered.
func (s *BaseAbacusListener) EnterLnFunction(ctx *LnFunctionContext) {}

// ExitLnFunction is called when production LnFunction is exited.
func (s *BaseAbacusListener) ExitLnFunction(ctx *LnFunctionContext) {}

// EnterLogDefFunction is called when production LogDefFunction is entered.
func (s *BaseAbacusListener) EnterLogDefFunction(ctx *LogDefFunctionContext) {}

// ExitLogDefFunction is called when production LogDefFunction is exited.
func (s *BaseAbacusListener) ExitLogDefFunction(ctx *LogDefFunctionContext) {}

// EnterLog2Function is called when production Log2Function is entered.
func (s *BaseAbacusListener) EnterLog2Function(ctx *Log2FunctionContext) {}

// ExitLog2Function is called when production Log2Function is exited.
func (s *BaseAbacusListener) ExitLog2Function(ctx *Log2FunctionContext) {}

// EnterLog10Function is called when production Log10Function is entered.
func (s *BaseAbacusListener) EnterLog10Function(ctx *Log10FunctionContext) {}

// ExitLog10Function is called when production Log10Function is exited.
func (s *BaseAbacusListener) ExitLog10Function(ctx *Log10FunctionContext) {}

// EnterFloorFunction is called when production FloorFunction is entered.
func (s *BaseAbacusListener) EnterFloorFunction(ctx *FloorFunctionContext) {}

// ExitFloorFunction is called when production FloorFunction is exited.
func (s *BaseAbacusListener) ExitFloorFunction(ctx *FloorFunctionContext) {}

// EnterCeilFunction is called when production CeilFunction is entered.
func (s *BaseAbacusListener) EnterCeilFunction(ctx *CeilFunctionContext) {}

// ExitCeilFunction is called when production CeilFunction is exited.
func (s *BaseAbacusListener) ExitCeilFunction(ctx *CeilFunctionContext) {}

// EnterExpFunction is called when production ExpFunction is entered.
func (s *BaseAbacusListener) EnterExpFunction(ctx *ExpFunctionContext) {}

// ExitExpFunction is called when production ExpFunction is exited.
func (s *BaseAbacusListener) ExitExpFunction(ctx *ExpFunctionContext) {}

// EnterSinFunction is called when production SinFunction is entered.
func (s *BaseAbacusListener) EnterSinFunction(ctx *SinFunctionContext) {}

// ExitSinFunction is called when production SinFunction is exited.
func (s *BaseAbacusListener) ExitSinFunction(ctx *SinFunctionContext) {}

// EnterCosFunction is called when production CosFunction is entered.
func (s *BaseAbacusListener) EnterCosFunction(ctx *CosFunctionContext) {}

// ExitCosFunction is called when production CosFunction is exited.
func (s *BaseAbacusListener) ExitCosFunction(ctx *CosFunctionContext) {}

// EnterTanFunction is called when production TanFunction is entered.
func (s *BaseAbacusListener) EnterTanFunction(ctx *TanFunctionContext) {}

// ExitTanFunction is called when production TanFunction is exited.
func (s *BaseAbacusListener) ExitTanFunction(ctx *TanFunctionContext) {}

// EnterRoundDefFunction is called when production RoundDefFunction is entered.
func (s *BaseAbacusListener) EnterRoundDefFunction(ctx *RoundDefFunctionContext) {}

// ExitRoundDefFunction is called when production RoundDefFunction is exited.
func (s *BaseAbacusListener) ExitRoundDefFunction(ctx *RoundDefFunctionContext) {}

// EnterSignFunction is called when production SignFunction is entered.
func (s *BaseAbacusListener) EnterSignFunction(ctx *SignFunctionContext) {}

// ExitSignFunction is called when production SignFunction is exited.
func (s *BaseAbacusListener) ExitSignFunction(ctx *SignFunctionContext) {}

// EnterAbsFunction is called when production AbsFunction is entered.
func (s *BaseAbacusListener) EnterAbsFunction(ctx *AbsFunctionContext) {}

// ExitAbsFunction is called when production AbsFunction is exited.
func (s *BaseAbacusListener) ExitAbsFunction(ctx *AbsFunctionContext) {}

// EnterRound2Function is called when production Round2Function is entered.
func (s *BaseAbacusListener) EnterRound2Function(ctx *Round2FunctionContext) {}

// ExitRound2Function is called when production Round2Function is exited.
func (s *BaseAbacusListener) ExitRound2Function(ctx *Round2FunctionContext) {}

// EnterLogFunction is called when production LogFunction is entered.
func (s *BaseAbacusListener) EnterLogFunction(ctx *LogFunctionContext) {}

// ExitLogFunction is called when production LogFunction is exited.
func (s *BaseAbacusListener) ExitLogFunction(ctx *LogFunctionContext) {}

// EnterMinFunction is called when production MinFunction is entered.
func (s *BaseAbacusListener) EnterMinFunction(ctx *MinFunctionContext) {}

// ExitMinFunction is called when production MinFunction is exited.
func (s *BaseAbacusListener) ExitMinFunction(ctx *MinFunctionContext) {}

// EnterMaxFunction is called when production MaxFunction is entered.
func (s *BaseAbacusListener) EnterMaxFunction(ctx *MaxFunctionContext) {}

// ExitMaxFunction is called when production MaxFunction is exited.
func (s *BaseAbacusListener) ExitMaxFunction(ctx *MaxFunctionContext) {}

// EnterAvgFunction is called when production AvgFunction is entered.
func (s *BaseAbacusListener) EnterAvgFunction(ctx *AvgFunctionContext) {}

// ExitAvgFunction is called when production AvgFunction is exited.
func (s *BaseAbacusListener) ExitAvgFunction(ctx *AvgFunctionContext) {}

// EnterUntilFunction is called when production UntilFunction is entered.
func (s *BaseAbacusListener) EnterUntilFunction(ctx *UntilFunctionContext) {}

// ExitUntilFunction is called when production UntilFunction is exited.
func (s *BaseAbacusListener) ExitUntilFunction(ctx *UntilFunctionContext) {}

// EnterFromFunction is called when production FromFunction is entered.
func (s *BaseAbacusListener) EnterFromFunction(ctx *FromFunctionContext) {}

// ExitFromFunction is called when production FromFunction is exited.
func (s *BaseAbacusListener) ExitFromFunction(ctx *FromFunctionContext) {}

// EnterReverseFunction is called when production ReverseFunction is entered.
func (s *BaseAbacusListener) EnterReverseFunction(ctx *ReverseFunctionContext) {}

// ExitReverseFunction is called when production ReverseFunction is exited.
func (s *BaseAbacusListener) ExitReverseFunction(ctx *ReverseFunctionContext) {}

// EnterNthFunction is called when production NthFunction is entered.
func (s *BaseAbacusListener) EnterNthFunction(ctx *NthFunctionContext) {}

// ExitNthFunction is called when production NthFunction is exited.
func (s *BaseAbacusListener) ExitNthFunction(ctx *NthFunctionContext) {}
