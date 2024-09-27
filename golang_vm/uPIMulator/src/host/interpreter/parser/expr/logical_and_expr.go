package expr

import (
	"errors"
)

type LogicalAndExpr struct {
	loperand *Expr
	roperand *Expr
}

func (this *LogicalAndExpr) Init(loperand *Expr, roperand *Expr) {
	if loperand.ExprType() != PRIMARY && loperand.ExprType() != POSTFIX &&
		loperand.ExprType() != UNARY &&
		loperand.ExprType() != MULTIPLICATIVE &&
		loperand.ExprType() != ADDITIVE &&
		loperand.ExprType() != SHIFT &&
		loperand.ExprType() != RELATIONAL &&
		loperand.ExprType() != EQUALITY &&
		loperand.ExprType() != BITWISE_AND &&
		loperand.ExprType() != BITWISE_XOR &&
		loperand.ExprType() != BITWISE_OR &&
		loperand.ExprType() != LOGICAL_AND {
		err := errors.New("loperand expr type is wrong")
		panic(err)
	} else if roperand.ExprType() != PRIMARY && roperand.ExprType() != POSTFIX &&
		roperand.ExprType() != UNARY &&
		roperand.ExprType() != MULTIPLICATIVE &&
		roperand.ExprType() != ADDITIVE &&
		roperand.ExprType() != SHIFT &&
		roperand.ExprType() != RELATIONAL &&
		roperand.ExprType() != EQUALITY &&
		roperand.ExprType() != BITWISE_AND &&
		roperand.ExprType() != BITWISE_XOR &&
		roperand.ExprType() != BITWISE_OR &&
		roperand.ExprType() != LOGICAL_AND {
		err := errors.New("roperand expr type is wrong")
		panic(err)
	}

	this.loperand = loperand
	this.roperand = roperand
}

func (this *LogicalAndExpr) Loperand() *Expr {
	return this.loperand
}

func (this *LogicalAndExpr) Roperand() *Expr {
	return this.roperand
}
