package analyzer

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var GoDocPrefix = &analysis.Analyzer{
	Name: "godocprefix",
	Doc:  "Checks godoc for exported fields of exported structs has the prefix of serialized versions.",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	g := &goDocPrefix{
		pass: pass,
	}

	return g.run()
}

type goDocPrefix struct {
	pass *analysis.Pass
}

func (g *goDocPrefix) run() (interface{}, error) {
	for _, file := range g.pass.Files {
		g.inspectFile(file)
	}

	return nil, nil
}

func (g *goDocPrefix) inspectFile(file *ast.File) bool {
	for _, decl := range file.Decls {
		switch decl := decl.(type) {
		case *ast.GenDecl:
			switch decl.Tok {
			case token.TYPE:
				for _, spec := range decl.Specs {
					typeSpec, ok := spec.(*ast.TypeSpec)
					if !ok {
						continue
					}

					g.inspectSpec(typeSpec)
				}
			}
		}
	}

	return false
}

func (g *goDocPrefix) inspectSpec(spec *ast.TypeSpec) {
	switch spec.Type.(type) {
	case *ast.StructType:
		if !ast.IsExported(spec.Name.Name) {
			return
		}

		g.inspectStruct(spec.Type.(*ast.StructType))
	}
}

func (g *goDocPrefix) inspectStruct(structType *ast.StructType) {
	for _, field := range structType.Fields.List {
		// Skip a line with multiple fields.
		if len(field.Names) != 1 {
			continue
		}

		if !ast.IsExported(field.Names[0].Name) {
			continue
		}

		jsonTag := g.jsonTag(field)
		if jsonTag == "" {
			continue
		}

		serializedName := strings.Split(jsonTag, ",")[0]
		if serializedName == "-" || serializedName == "" {
			continue
		}

		var doc string
		pos := field.Pos()
		if field.Doc != nil {
			doc = field.Doc.Text()
			pos = field.Doc.Pos()
		}

		fixablePrefix := field.Names[0].Name + " "
		prefix := serializedName + " "

		if !strings.HasPrefix(doc, prefix) {
			var suggestedFixes []analysis.SuggestedFix

			if pos, end := g.fixableDoc(fixablePrefix, field.Doc); pos != token.NoPos {
				suggestedFixes = append(suggestedFixes, analysis.SuggestedFix{
					Message: "Use the serialized name as a prefix",
					TextEdits: []analysis.TextEdit{
						{
							Pos:     pos,
							End:     end,
							NewText: []byte(prefix),
						},
					},
				})
			}

			g.pass.Report(analysis.Diagnostic{
				Pos:     pos,
				Message: fmt.Sprintf("godoc prefix of %s should be '%s '", field.Names[0].Name, serializedName),

				SuggestedFixes: suggestedFixes,
			})
		}
	}
}

func (g *goDocPrefix) jsonTag(field *ast.Field) string {
	if field.Tag == nil {
		return ""
	}

	unquoted, err := strconv.Unquote(field.Tag.Value)
	if err != nil {
		return ""
	}

	tag := reflect.StructTag(unquoted)

	return tag.Get("json")
}

func (g *goDocPrefix) fixableDoc(fixablePrefix string, doc *ast.CommentGroup) (pos, end token.Pos) {
	if doc == nil {
		return token.NoPos, token.NoPos
	}
	if !strings.HasPrefix(doc.Text(), fixablePrefix) {
		return token.NoPos, token.NoPos
	}

	fixablePrefixBytes := []byte(fixablePrefix)

	for _, comment := range doc.List {
		index := bytes.Index([]byte(comment.Text), fixablePrefixBytes)

		if index == -1 {
			continue
		}

		pos := comment.Pos() + token.Pos(index)

		return pos, pos + token.Pos(len(fixablePrefixBytes))
	}

	return token.NoPos, token.NoPos
}
