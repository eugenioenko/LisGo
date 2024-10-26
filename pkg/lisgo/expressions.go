package lisgo

type Expression interface {
	Accept(visitor VisitorExpression) LisgoData
}

type VisitorExpression interface {
	VisitExpressionList(expr *ExpressionList) LisgoData
	VisitExpressionAtom(expr *ExpressionAtom) LisgoData
}

type ExpressionList struct {
	List []Expression
}

func NewExpressionList(value []Expression) *ExpressionList {
	return &ExpressionList{value}
}

func (expr *ExpressionList) Accept(visitor VisitorExpression) LisgoData {
	return visitor.VisitExpressionList(expr)
}

type ExpressionAtom struct {
	Atom Token
}

func NewExpressionAtom(value Token) *ExpressionAtom {
	return &ExpressionAtom{value}
}

func (expr *ExpressionAtom) Accept(visitor VisitorExpression) LisgoData {
	return visitor.VisitExpressionAtom(expr)
}
