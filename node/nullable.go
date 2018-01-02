package node

type Nullable struct {
	name       string
	attributes map[string]interface{}
	position   *Position
	expr       Node
}

func NewNullable(expression Node) Node {
	return Nullable{
		"Nullable",
		map[string]interface{}{},
		nil,
		expression,
	}
}

func (n Nullable) Name() string {
	return "Nullable"
}

func (n Nullable) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Nullable) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Nullable) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Nullable) Position() *Position {
	return n.position
}

func (n Nullable) SetPosition(p *Position) Node {
	n.position = p
	return n
}

func (n Nullable) Walk(v Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}