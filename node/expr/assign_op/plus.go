package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Plus struct {
	AssignOp
}

func NewPlus(variable node.Node, expression node.Node) node.Node {
	return Plus{
		AssignOp{
			"AssignPlus",
			map[string]interface{}{},
			variable,
			expression,
		},
	}
}

func (n Plus) Name() string {
	return "Plus"
}

func (n Plus) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Plus) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Plus) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Plus) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	if n.expression != nil {
		vv := v.GetChildrenVisitor("expression")
		n.expression.Walk(vv)
	}

	v.LeaveNode(n)
}
