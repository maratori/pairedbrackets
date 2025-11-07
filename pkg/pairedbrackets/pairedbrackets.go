package pairedbrackets

import (
	"flag"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
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

// NewAnalyzer returns Analyzer that checks formatting of paired brackets.
func NewAnalyzer() *analysis.Analyzer {
	var (
		fs flag.FlagSet
	)

	return &analysis.Analyzer{
		Name:     "pairedbrackets",
		Doc:      "linter checks formatting of paired brackets",
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    fs,
		Run: func(pass *analysis.Pass) (any, error) {
			return run(runner{
				pass: pass,
			})
		},
	}
}

type runner struct {
	pass *analysis.Pass
}

// Run runs the analyzer.
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
func run(r runner) (any, error) {
	filter := []ast.Node{ // alphabet order
		new(ast.CallExpr),
		new(ast.CompositeLit),
		new(ast.ForStmt),
		new(ast.FuncDecl),
		new(ast.FuncLit),
		new(ast.FuncType),
		new(ast.GenDecl),
		new(ast.IfStmt),
		new(ast.IndexExpr),
		new(ast.IndexListExpr),
		new(ast.InterfaceType),
		new(ast.ParenExpr),
		new(ast.RangeStmt),
		new(ast.SelectStmt),
		new(ast.SliceExpr),
		new(ast.StructType),
		new(ast.SwitchStmt),
		new(ast.TypeAssertExpr),
		new(ast.TypeSpec),
		new(ast.TypeSwitchStmt),
	}
	astInspector := r.pass.ResultOf[inspect.Analyzer].(*inspector.Inspector) //nolint:errcheck // let's panic
	astInspector.Preorder(filter, func(node ast.Node) {
		switch n := node.(type) {
		case *ast.CallExpr:
			validate(r, Parenthesis, n.Lparen, n.Rparen, n.Args, Argument)
		case *ast.CompositeLit:
			validate(r, Brace, n.Lbrace, n.Rbrace, n.Elts, CompositeElement)
		case *ast.ForStmt:
			validateBlock(r, n.Body)
		case *ast.FuncDecl:
			// n.Type is checked separately in [*ast.FuncType]
			validateFieldList(r, Parenthesis, n.Recv, Receiver)
			validateBlock(r, n.Body)
		case *ast.FuncLit:
			// n.Type is checked separately in [*ast.FuncType]
			validateBlock(r, n.Body)
		case *ast.FuncType:
			validateFieldList(r, Bracket, n.TypeParams, TypeParameter)
			validateFieldList(r, Parenthesis, n.Params, Parameter)
			validateFieldList(r, Parenthesis, n.Results, Result)
		case *ast.GenDecl:
			validate(r, Parenthesis, n.Lparen, n.Rparen, n.Specs, genDeclElement(n))
		case *ast.IfStmt:
			validateBlock(r, n.Body)
			if block, ok := n.Else.(*ast.BlockStmt); ok {
				validateBlock(r, block)
			}
		case *ast.IndexExpr:
			validate(r, Bracket, n.Lbrack, n.Rbrack, []ast.Expr{n.Index}, Unknown) // unknown because it may be several types
		case *ast.IndexListExpr:
			validate(r, Bracket, n.Lbrack, n.Rbrack, n.Indices, Unknown) // unknown because it may be several types
		case *ast.InterfaceType:
			validateFieldList(r, Brace, n.Methods, Method)
		case *ast.ParenExpr:
			validate(r, Parenthesis, n.Lparen, n.Rparen, []ast.Expr{n.X}, Expression)
		case *ast.RangeStmt:
			validateBlock(r, n.Body)
		case *ast.SelectStmt:
			validateBlock(r, n.Body)
		case *ast.SliceExpr:
			validate(r, Bracket, n.Lbrack, n.Rbrack, []ast.Expr{n.Low, n.High, n.Max}, Index)
		case *ast.StructType:
			validateFieldList(r, Brace, n.Fields, Field)
		case *ast.SwitchStmt:
			validateBlock(r, n.Body)
		case *ast.TypeAssertExpr:
			validate(r, Parenthesis, n.Lparen, n.Rparen, []ast.Expr{n.Type}, Type)
		case *ast.TypeSpec:
			validateFieldList(r, Bracket, n.TypeParams, TypeParameter)
		case *ast.TypeSwitchStmt:
			validateBlock(r, n.Body)
		}
	})

	return nil, nil //nolint:nilnil // the linter has no result, it should return nil
}

func validateFieldList(r runner, name Name, node *ast.FieldList, element Element) {
	if node == nil {
		return
	}

	validate(r, name, node.Opening, node.Closing, node.List, element)
}

func validateBlock(r runner, node *ast.BlockStmt) {
	if node == nil {
		return
	}

	validate(r, Brace, node.Lbrace, node.Rbrace, node.List, Statement)
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

func validate[N ast.Node](r runner, bracket Name, left, right token.Pos, list []N, element Element) {
	if left == token.NoPos || right == token.NoPos {
		return // no brackets - nothing to check
	}

	leftLine := r.pass.Fset.Position(left).Line
	rightLine := r.pass.Fset.Position(right).Line

	if leftLine == rightLine {
		return
	}

	// aaa, bbb, ccccccccc
	// ^         ^       ^
	firstPos, lastPos, lastEnd, ok := boundaries(list)

	if !ok {
		return // list is empty
	}

	firstPosLine := r.pass.Fset.Position(firstPos).Line
	lastPosLine := firstPosLine // optimisation
	if lastPos != firstPos {
		lastPosLine = r.pass.Fset.Position(lastPos).Line
	}
	lastEndLine := r.pass.Fset.Position(lastEnd).Line

	switch {
	case leftLine == lastPosLine: // left bracket is ok
		if lastEndLine != rightLine {
			r.pass.Reportf(right, "right %s should be on the previous line", bracket)
		}
	case leftLine != firstPosLine: // left bracket is ok
		if lastEndLine == rightLine {
			r.pass.Reportf(right, "right %s should be on the next line", bracket)
		}
	default:
		r.pass.Reportf(
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
