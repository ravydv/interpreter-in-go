package ast

import (
	"bytes"
	"strings"

	"github.com/ravydv/interpreter-in-go/token"
)

// Node every node in ast implement this interface.
type Node interface {
	TokenLiteral() string
	String() string
}

/*
The AST consists solely of Nodes that are connected to
each other - it’s a tree after all. Some of these nodes implement the Statement and some the
Expression interface. These interfaces only contain dummy methods called statementNode and
expressionNode respectively. They are not strictly necessary but help us by guiding the Go
compiler and possibly causing it to throw errors when we use a Statement where an Expression
should’ve been used, and vice versa.
*/
type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Program root node of every AST our parser produces. Every valid Monkey program is a series of statements.
type Program struct {
	Statements []Statement // slice of AST nodes that implement the Statement interface.
}

// TokenLiteral return literal value of the token associated with AST node.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// LetStatement is variable binding statements of the form of: let <identifier> = <expression>;, an AST node
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier // Name to hold the identifier of the binding
	Value Expression  // Value for the expression that produces the value
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral return literal value of the token associated with AST node.
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// Identifier identifier AST node hold the identifier binding, ex: the x in let x = 5;
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {

}

// TokenLiteral return literal value of the token associated with AST node.
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

// ReturnStatement an AST node for return statement consist of keyword return and an expression ex: return <expression>
type ReturnStatement struct {
	Token       token.Token // the token.RETURN token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral return literal value of the token associated with AST node.
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// ExpressionStatement expression statement AST node, an expression statement is not really a distinct statement; it’s a statement that consists solely of one expression.
// ex: let x = 5; //let statement
// x + 10; // expression statement
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral return literal value of the token associated with AST node.
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// IntegerLiteral integer literal AST node
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral return integer literal value  associated with AST node.
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

//PrefixExpression an AST node for a prefix operator expression , <operator><expression> ex: !5
type PrefixExpression struct {
	Token    token.Token // the prefix token, ex: !, -
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral return prefix operator literal value of AST node
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

// InfixExpression an AST node for infix expression of the form of: <expression> <operator> <expression>, ex: 10 + 10
type InfixExpression struct {
	Token    token.Token // the operator token, ex: +, -, *
	Left     Expression
	Operator string
	Right    Expression
}

func (in *InfixExpression) expressionNode() {}

// TokenLiteral return infix expression literal associated with the AST node
func (in *InfixExpression) TokenLiteral() string {
	return in.Token.Literal
}
func (in *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(in.Left.String())
	out.WriteString(" " + in.Operator + " ")
	out.WriteString(in.Right.String())
	out.WriteString(")")
	return out.String()
}

// Boolean an AST node for boolean
type Boolean struct {
	Token token.Token // the true, false token
	Value bool
}

func (b *Boolean) expressionNode() {}

// TokenLiteral return boolean literal
func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}
func (b *Boolean) String() string {
	return b.Token.Literal
}

// IfExpression an if statement AST node of the form of: if(<condition>) <consequence> else <alternative>
type IfExpression struct {
	Token       token.Token // the 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}

// TokenLiteral return literal value of the AST node
func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

// BlockStatement is list of statement
type BlockStatement struct {
	Token      token.Token // the { token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}
func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// FunctionLiteral function literal AST node of the form of: fn <parameters> <blockstatement>. where parameters: (<identifier>,<identifier> ..)
type FunctionLiteral struct {
	Token      token.Token // the fn token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {}
func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer
	params := []string{}

	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())
	return out.String()
}

// CallExpression call expression AST node of the form of: <expression>(<comma seperated expression>)
type CallExpression struct {
	Token     token.Token // the '(' token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer
	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}
