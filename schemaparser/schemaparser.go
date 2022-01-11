package schemaparser

import (
	"github.com/jensneuse/graphql-go-tools/pkg/astparser"
)

func ParseTypes(schema []byte) ([]string, error) {
	document, report := astparser.ParseGraphqlDocumentBytes(schema)
	if report.HasErrors() {
		return nil, report
	}

	visitor := newVisitor()
	visitor.Walk(&document, nil, &report)
	if report.HasErrors() {
		return nil, report
	}

	return visitor.typeNames, nil
}

func CountEnumValues(schema []byte) (int, error) {
	document, report := astparser.ParseGraphqlDocumentBytes(schema)
	if report.HasErrors() {
		return 0, report
	}

	visitor := newVisitor()
	visitor.Walk(&document, nil, &report)
	if report.HasErrors() {
		return 0, report
	}

	return len(visitor.enumTypes), nil
}
