package lexer

import (
	"encoding/json"
	"fmt"
	"strings"
)

type TokenType int

type Token struct {
	Type     TokenType
	Value    string
	Location *Location
}

const (
	None                 TokenType = iota
	Identifier                     // identifier
	WhiteSpace                     // ' '
	NewLine                        // '\n'
	SingleLineComment              // '//'
	MultiLineComment               // '/*' -> '*/'
	OpenBrace                      // '{'
	CloseBrace                     // '}'
	OpenParen                      // '('
	CloseParen                     // ')'
	OpenBracket                    // '['
	CloseBracket                   // ']'
	Number                         // '3',  '.15', '-2', '-6.2'
	Boolean                        // 'true', 'false'
	String                         // '"' -> some text -> '"'
	Assignment                     // operator? -> '='
	ImplicitDeclaration            // ':='
	Increment                      // '++'
	Decrement                      // '--'
	Addition                       // '+'
	Subtraction                    // '-'
	Multiplication                 // '*'
	Division                       // '/'
	Exponentiation                 // '^'
	Modulus                        // '%'
	Or                             // '|'
	And                            // '&'
	Nor                            // '!|'
	Nand                           // '!&'
	Not                            // '!'
	Is                             // 'is'
	EqualTo                        // '=='
	NotEqualTo                     // '!='
	GreaterThan                    // '>'
	LessThan                       // '<'
	GreaterThanOrEqualTo           // '>='
	LessThanOrEqualTo              // '<='
	Delimiter                      // ','
	Return                         // 'return'
	Pipe                           // '|>'
	Break                          // 'break'
	Continue                       // 'continue'
	Fallthrough                    // 'fallthrough'
	My                             // 'my'
	Class                          // 'class'
	New                            // 'new'
	Operator                       // 'operator'
	To                             // 'to'
	Extends                        // 'extends'
	From                           // 'from'
	Enum                           // 'enum'
	For                            // 'for'
	In                             // 'in'
	While                          // 'while'
	Do                             // 'do'
	If                             // 'if'
	Else                           // 'else'
	Match                          // 'match'
	Yield                          // 'yield'
	QuestionMark                   // '?'
	Colon                          // ':'
	Null                           // 'null'
	Nullish                        // '??'
	Spread                         // '...'
	Semicolon                      // ';'
	Import                         // 'import'
	Export                         // 'export'
	Test                           // 'test'
	Assert                         // 'assert'
	Throws                         // 'throws'
	Catch                          // 'catch'
	Okay                           // 'okay'
	EOF
)

var Keywords = map[string]TokenType{
	"true":        Boolean,
	"false":       Boolean,
	"is":          Is,
	"return":      Return,
	"break":       Break,
	"continue":    Continue,
	"fallthrough": Fallthrough,
	"my":          My,
	"class":       Class,
	"new":         New,
	"operator":    Operator,
	"to":          To,
	"extends":     Extends,
	"from":        From,
	"enum":        Enum,
	"for":         For,
	"in":          In,
	"while":       While,
	"do":          Do,
	"if":          If,
	"else":        Else,
	"match":       Match,
	"yield":       Yield,
	"null":        Null,
	"import":      Import,
	"export":      Export,
	"test":        Test,
	"assert":      Assert,
	"throws":      Throws,
	"catch":       Catch,
	"okay":        Okay,
}

var Symbols = map[string]TokenType{
	"\n":  NewLine,
	"{":   OpenBrace,
	"}":   CloseBrace,
	"(":   OpenParen,
	")":   CloseParen,
	"[":   OpenBracket,
	"]":   CloseBracket,
	":=":  ImplicitDeclaration,
	"+":   Addition,
	"-":   Subtraction,
	"*":   Multiplication,
	"/":   Division,
	"^":   Exponentiation,
	"%":   Modulus,
	"++":  Increment,
	"--":  Decrement,
	"|":   Or,
	"&":   And,
	"!|":  Nor,
	"!&":  Nand,
	"!":   Not,
	"==":  EqualTo,
	"!=":  NotEqualTo,
	">":   GreaterThan,
	"<":   LessThan,
	">=":  GreaterThanOrEqualTo,
	"<=":  LessThanOrEqualTo,
	",":   Delimiter,
	"|>":  Pipe,
	"?":   QuestionMark,
	":":   Colon,
	"??":  Nullish,
	"...": Spread,
	";":   Semicolon,
}

func formatValue(value string) string {
	return strings.ReplaceAll(value, "\n", "\\n")
}

func CreateToken(tokentype TokenType, value string, row, col, idx int) *Token {
	return &Token{
		tokentype,
		value,
		CreateLocation(row, col, idx),
	}
}

func (tt TokenType) String() string {
	switch tt {
	case Identifier:
		return "Identifier"
	case WhiteSpace:
		return "WhiteSpace"
	case NewLine:
		return "NewLine"
	case SingleLineComment:
		return "SingleLineComment"
	case MultiLineComment:
		return "MultiLineComment"
	case OpenBrace:
		return "OpenBrace"
	case CloseBrace:
		return "CloseBrace"
	case OpenParen:
		return "OpenParen"
	case CloseParen:
		return "CloseParen"
	case OpenBracket:
		return "OpenBracket"
	case CloseBracket:
		return "CloseBracket"
	case Number:
		return "Number"
	case Boolean:
		return "Boolean"
	case String:
		return "String"
	case Assignment:
		return "Assignment"
	case ImplicitDeclaration:
		return "ImplicitDeclaration"
	case Increment:
		return "Increment"
	case Decrement:
		return "Decrement"
	case Addition:
		return "Addition"
	case Subtraction:
		return "Subtraction"
	case Multiplication:
		return "Multiplication"
	case Division:
		return "Division"
	case Exponentiation:
		return "Exponentiation"
	case Modulus:
		return "Modulus"
	case Or:
		return "Or"
	case And:
		return "And"
	case Nor:
		return "Nor"
	case Nand:
		return "Nand"
	case Not:
		return "Not"
	case Is:
		return "Is"
	case EqualTo:
		return "EqualTo"
	case NotEqualTo:
		return "NotEqualTo"
	case GreaterThan:
		return "GreaterThan"
	case LessThan:
		return "LessThan"
	case GreaterThanOrEqualTo:
		return "GreaterThanOrEqualTo"
	case LessThanOrEqualTo:
		return "LessThanOrEqualTo"
	case Delimiter:
		return "Delimiter"
	case Return:
		return "Return"
	case Pipe:
		return "Pipe"
	case Break:
		return "Break"
	case Continue:
		return "Continue"
	case Fallthrough:
		return "Fallthrough"
	case My:
		return "My"
	case Class:
		return "Class"
	case New:
		return "New"
	case Operator:
		return "Operator"
	case To:
		return "To"
	case Extends:
		return "Extends"
	case From:
		return "From"
	case Enum:
		return "Enum"
	case For:
		return "For"
	case In:
		return "In"
	case While:
		return "While"
	case Do:
		return "Do"
	case If:
		return "If"
	case Else:
		return "Else"
	case Match:
		return "Match"
	case Yield:
		return "Yield"
	case QuestionMark:
		return "QuestionMark"
	case Colon:
		return "Colon"
	case Null:
		return "Null"
	case Nullish:
		return "Nullish"
	case Spread:
		return "Spread"
	case Semicolon:
		return "Semicolon"
	case Import:
		return "Import"
	case Export:
		return "Export"
	case Test:
		return "Test"
	case Assert:
		return "Assert"
	case Throws:
		return "Throws"
	case Catch:
		return "Catch"
	case Okay:
		return "Okay"
	case EOF:
		return "EOF"
	default:
		return "Unknown"
	}
}

func (tt TokenType) MarshalJSON() ([]byte, error) {
	return json.Marshal(tt.String())
}

func (t Token) String() string {
	row, col, idx := t.Location.Get()

	return "Token{" + t.Type.String() + " '" + formatValue(t.Value) + "' at " +
		fmt.Sprint(row) + ":" + fmt.Sprint(col) + ", #" + fmt.Sprint(idx) + "}"
}