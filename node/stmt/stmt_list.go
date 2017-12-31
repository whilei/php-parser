package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type StmtList struct {
	name       string
	attributes map[string]interface{}
	stmts      []node.Node
}

func NewStmtList(stmts []node.Node) node.Node {
	return StmtList{
		"StmtList",
		map[string]interface{}{},
		stmts,
	}
}

func (n StmtList) Name() string {
	return "StmtList"
}

func (n StmtList) Attributes() map[string]interface{} {
	return n.attributes
}

func (n StmtList) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n StmtList) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n StmtList) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.stmts != nil {
		vv := v.GetChildrenVisitor("stmts")
		for _, nn := range n.stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
