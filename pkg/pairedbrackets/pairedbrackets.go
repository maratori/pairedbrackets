package pairedbrackets

import (
	"flag"
	"go/ast"
	"go/token"
	"go/types"
	"regexp"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/types/typeutil"
)

type Name string

const (
	Parenthesis Name = "parenthesis" // ()
	Bracket     Name = "bracket"     // []
	Brace       Name = "brace"       // {}
)

type Element string

const (
	Unknown          Element = "element"
	Argument         Element = "argument"
	CompositeElement Element = "composite element"
	Constant         Element = "constant"
	Expression       Element = "expression"
	Field            Element = "field"
	Import           Element = "import"
	Index            Element = "index"
	Method           Element = "method"
	Parameter        Element = "parameter"
	Receiver         Element = "receiver"
	Result           Element = "result"
	Statement        Element = "statement"
	Type             Element = "type"
	TypeParameter    Element = "type parameter"
	Variable         Element = "variable"
)

// Fully qualified function examples:
//   - github.com/stretchr/testify/require.Equal
//   - (*github.com/stretchr/testify/assert.Assertions).Equal
const (
	IgnoreFuncCallsFlagName    = "ignore-func-calls"
	IgnoreFuncCallsFlagUsage   = "comma separated list of regexp patterns of fully qualified function calls to ignore"
	IgnoreFuncCallsFlagDefault = "github.com/stretchr/testify/assert,github.com/stretchr/testify/require"
)

// NewAnalyzer returns Analyzer that checks formatting of paired brackets.
func NewAnalyzer() *analysis.Analyzer {
	var (
		ignoreFuncCalls commaListRegexpFlag
		fs              flag.FlagSet
	)

	_ = ignoreFuncCalls.Set(IgnoreFuncCallsFlagDefault)
	fs.Var(&ignoreFuncCalls, IgnoreFuncCallsFlagName, IgnoreFuncCallsFlagUsage)

	return &analysis.Analyzer{
		Name:     "pairedbrackets",
		Doc:      "linter checks formatting of paired brackets",
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    fs,
		Run:      runner(&ignoreFuncCalls),
	}
}

// runner returns run function for analyzer.
//
// Nodes without brackets (alphabet order):
//   - [*ast.AssignStmt]
//   - [*ast.BadDecl]
//   - [*ast.BadExpr]
//   - [*ast.BadStmt]
//   - [*ast.BasicLit]
//   - [*ast.BinaryExpr]
//   - [*ast.BranchStmt]
//   - [*ast.CaseClause]
//   - [*ast.ChanType]
//   - [*ast.CommClause]
//   - [*ast.CommentGroup]
//   - [*ast.Comment]
//   - [*ast.DeclStmt]
//   - [*ast.DeferStmt]
//   - [*ast.Ellipsis]
//   - [*ast.EmptyStmt]
//   - [*ast.ExprStmt]
//   - [*ast.Field]
//   - [*ast.File]
//   - [*ast.GoStmt]
//   - [*ast.Ident]
//   - [*ast.ImportSpec]
//   - [*ast.IncDecStmt]
//   - [*ast.KeyValueExpr]
//   - [*ast.LabeledStmt]
//   - [*ast.Package]
//   - [*ast.ReturnStmt]
//   - [*ast.SelectorExpr]
//   - [*ast.SendStmt]
//   - [*ast.StarExpr]
//   - [*ast.UnaryExpr]
//   - [*ast.ValueSpec]
//
// Ignored because go syntax doesn't allow to start line with right bracket:
//   - [*ast.ArrayType]
//   - [*ast.MapType]
//
// [*ast.BlockStmt] ignored because is checked in other cases:
//   - [*ast.ForStmt]
//   - [*ast.FuncLit]
//   - [*ast.IfStmt]
//   - [*ast.RangeStmt]
//   - [*ast.SelectStmt]
//   - [*ast.SwitchStmt]
//   - [*ast.TypeSwitchStmt]
//
// [*ast.FieldList] ignored because is checked in other cases:
//   - [*ast.FuncDecl]
//   - [*ast.FuncType]
//   - [*ast.InterfaceType]
//   - [*ast.StructType]
//   - [*ast.TypeSpec]
func runner(ignoreFuncCalls *commaListRegexpFlag) func(*analysis.Pass) (interface{}, error) {
	return func(pass *analysis.Pass) (interface{}, error) {
		filter := []ast.Node{ // alphabet order
			&ast.CallExpr{},
			&ast.CompositeLit{},
			&ast.ForStmt{},
			&ast.FuncDecl{},
			&ast.FuncLit{},
			&ast.FuncType{},
			&ast.GenDecl{},
			&ast.IfStmt{},
			&ast.IndexExpr{},
			&ast.IndexListExpr{},
			&ast.InterfaceType{},
			&ast.ParenExpr{},
			&ast.RangeStmt{},
			&ast.SelectStmt{},
			&ast.SliceExpr{},
			&ast.StructType{},
			&ast.SwitchStmt{},
			&ast.TypeAssertExpr{},
			&ast.TypeSpec{},
			&ast.TypeSwitchStmt{},
		}
		astInspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector) //nolint:errcheck // let's panic
		astInspector.Preorder(filter, func(node ast.Node) {
			switch n := node.(type) {
			case *ast.CallExpr:
				validateCall(pass, n, ignoreFuncCalls)
			case *ast.CompositeLit:
				validate(pass, Brace, n.Lbrace, n.Rbrace, n.Elts, CompositeElement)
			case *ast.ForStmt:
				validateBlock(pass, n.Body)
			case *ast.FuncDecl:
				// n.Type is checked separately in [*ast.FuncType]
				validateFieldList(pass, Parenthesis, n.Recv, Receiver)
				validateBlock(pass, n.Body)
			case *ast.FuncLit:
				// n.Type is checked separately in [*ast.FuncType]
				validateBlock(pass, n.Body)
			case *ast.FuncType:
				validateFieldList(pass, Bracket, n.TypeParams, TypeParameter)
				validateFieldList(pass, Parenthesis, n.Params, Parameter)
				validateFieldList(pass, Parenthesis, n.Results, Result)
			case *ast.GenDecl:
				validate(pass, Parenthesis, n.Lparen, n.Rparen, n.Specs, genDeclElement(n))
			case *ast.IfStmt:
				validateBlock(pass, n.Body)
				if block, ok := n.Else.(*ast.BlockStmt); ok {
					validateBlock(pass, block)
				}
			case *ast.IndexExpr:
				validate(pass, Bracket, n.Lbrack, n.Rbrack, []ast.Expr{n.Index}, Unknown) // unknown because it may be several types
			case *ast.IndexListExpr:
				validate(pass, Bracket, n.Lbrack, n.Rbrack, n.Indices, Unknown) // unknown because it may be several types
			case *ast.InterfaceType:
				validateFieldList(pass, Brace, n.Methods, Method)
			case *ast.ParenExpr:
				validate(pass, Parenthesis, n.Lparen, n.Rparen, []ast.Expr{n.X}, Expression)
			case *ast.RangeStmt:
				validateBlock(pass, n.Body)
			case *ast.SelectStmt:
				validateBlock(pass, n.Body)
			case *ast.SliceExpr:
				validate(pass, Bracket, n.Lbrack, n.Rbrack, []ast.Expr{n.Low, n.High, n.Max}, Index)
			case *ast.StructType:
				validateFieldList(pass, Brace, n.Fields, Field)
			case *ast.SwitchStmt:
				validateBlock(pass, n.Body)
			case *ast.TypeAssertExpr:
				validate(pass, Parenthesis, n.Lparen, n.Rparen, []ast.Expr{n.Type}, Type)
			case *ast.TypeSpec:
				validateFieldList(pass, Bracket, n.TypeParams, TypeParameter)
			case *ast.TypeSwitchStmt:
				validateBlock(pass, n.Body)
			}
		})

		return nil, nil
	}
}

func validateCall(pass *analysis.Pass, node *ast.CallExpr, ignoreFuncCalls *commaListRegexpFlag) {
	if callee := typeutil.Callee(pass.TypesInfo, node); callee != nil {
		if fn, ok := callee.(*types.Func); ok {
			if ignoreFuncCalls.Match(fn.FullName()) {
				return
			}
		}
	}

	validate(pass, Parenthesis, node.Lparen, node.Rparen, node.Args, Argument)
}

func validateFieldList(pass *analysis.Pass, name Name, node *ast.FieldList, element Element) {
	if node == nil {
		return
	}

	validate(pass, name, node.Opening, node.Closing, node.List, element)
}

func validateBlock(pass *analysis.Pass, node *ast.BlockStmt) {
	if node == nil {
		return
	}

	validate(pass, Brace, node.Lbrace, node.Rbrace, node.List, Statement)
}

func genDeclElement(node *ast.GenDecl) Element {
	switch node.Tok { //nolint:exhaustive // only 4 options are possible, see doc for [*ast.GenDecl]
	case token.IMPORT:
		return Import
	case token.CONST:
		return Constant
	case token.TYPE:
		return Type
	case token.VAR:
		return Variable
	default:
		return Unknown
	}
}

func validate[N ast.Node](pass *analysis.Pass, bracket Name, left, right token.Pos, list []N, element Element) {
	if left == token.NoPos || right == token.NoPos {
		return // no brackets - nothing to check
	}

	leftLine := pass.Fset.Position(left).Line
	rightLine := pass.Fset.Position(right).Line

	if leftLine == rightLine {
		return
	}

	// aaa, bbb, ccccccccc
	// ^         ^       ^
	firstPos, lastPos, lastEnd, ok := boundaries(list)

	if !ok {
		return // list is empty
	}

	firstPosLine := pass.Fset.Position(firstPos).Line
	lastPosLine := firstPosLine // optimisation
	if lastPos != firstPos {
		lastPosLine = pass.Fset.Position(lastPos).Line
	}
	lastEndLine := pass.Fset.Position(lastEnd).Line

	switch {
	case leftLine == lastPosLine: // left bracket is ok
		if lastEndLine != rightLine {
			pass.Reportf(right, "right %s should be on the previous line", bracket)
		}
	case leftLine != firstPosLine: // left bracket is ok
		if lastEndLine == rightLine {
			pass.Reportf(right, "right %s should be on the next line", bracket)
		}
	default:
		pass.Reportf(
			left,
			"left %s should either be the last character on a line or be on the same line with the last %s",
			bracket,
			element,
		)
	}
}

// boundaries returns:
//   - position of the first character of the first node
//   - position of the first character of the last node
//   - position of the last character of the last node
//   - false, if list doesn't contain non-nil nodes
//
// Example:
//
//	aaaaaa, bbbbbb, cccccc
//	^               ^    ^
func boundaries[N ast.Node](list []N) (token.Pos, token.Pos, token.Pos, bool) {
	ok := false
	firstPos, lastPos, lastEnd := token.NoPos, token.NoPos, token.NoPos
	var last N
	for _, n := range list {
		if any(n) == nil {
			continue
		}

		p := n.Pos()

		if !ok {
			ok = true
			firstPos, lastPos = p, p
			last = n
			continue
		}

		if p < firstPos { // seems like impossible, but let's be conservative
			firstPos = p
		}

		if p > lastPos {
			lastPos = p
			last = n
		}
	}

	if ok {
		lastEnd = last.End() - 1 // -1 because we need position of node's last character
	}

	return firstPos, lastPos, lastEnd, ok
}

type commaListRegexpFlag struct {
	original string
	regs     []*regexp.Regexp
}

func (f *commaListRegexpFlag) String() string {
	if f == nil {
		return ""
	}
	return f.original
}

func (f *commaListRegexpFlag) Set(value string) error {
	regs := make([]*regexp.Regexp, 0, len(value))
	for _, pattern := range strings.Split(value, ",") {
		re, err := regexp.Compile(pattern)
		if err != nil {
			return err
		}
		regs = append(regs, re)
	}
	f.original = value
	f.regs = regs
	return nil
}

func (f *commaListRegexpFlag) Match(s string) bool {
	for _, r := range f.regs {
		if r.MatchString(s) {
			return true
		}
	}
	return false
}
