package schemaparser

import (
	"github.com/jensneuse/graphql-go-tools/pkg/ast"
	"github.com/jensneuse/graphql-go-tools/pkg/astvisitor"
)

type visitor struct {
	astvisitor.Walker
	typeNames []string
	enumTypes []string
}

func newVisitor() *visitor {
	w := astvisitor.NewWalker(48)
	v := &visitor{
		Walker:    w,
		typeNames: []string{},
	}

	v.RegisterEnterDocumentVisitor(v)
	return v
}

func (v *visitor) EnterDocument(operation, definition *ast.Document) {
	for _, r := range operation.RootNodes {
		switch r.Kind {

		/*
			This block comment is for task 2
		*/
		// case ast.NodeKindInterfaceTypeDefinition:
		// 	name := operation.InterfaceTypeDefinitionNameString(r.Ref)
		// 	v.typeNames = append(v.typeNames, name)
		case ast.NodeKindObjectTypeDefinition:
			name := operation.ObjectTypeDefinitionNameString(r.Ref)
			v.typeNames = append(v.typeNames, name)
		case ast.NodeKindEnumTypeDefinition: // task 3
			enumNode := operation.EnumTypeDefinitions[r.Ref]
			refs := enumNode.EnumValuesDefinition.Refs
			for _, ref := range refs {
				x := operation.EnumValueDefinitionNameString(ref)
				v.enumTypes = append(v.enumTypes, x)
			}
		default:
			continue
		}
	}
}
