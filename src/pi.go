package main

import (
	"math"
	"math/big"
)

// CalcPi calculate Pi for n number of digits
func CalcPi(digits float64) (*big.Float, uint) {
	/**
	* 	This is an implementation for https://en.wikipedia.org/wiki/Chudnovsky_algorithm
	*	it can be improved using binary splitting http://numbers.computation.free.fr/Constants/Algorithms/splitting.html
	* 	if we split it into two independent parts and simplify the formula for more details https://www.craig-wood.com/nick/articles/pi-chudnovsky/
	 */
	n := int64(2 + int(float64(digits)/14.181647462))
	prec := uint(int(math.Ceil(math.Log2(10)*digits)) + int(math.Ceil(math.Log10(digits))) + 2)

	c := new(big.Float).Mul(
		big.NewFloat(float64(426880)),
		new(big.Float).SetPrec(prec).Sqrt(big.NewFloat(float64(10005))),
	)

	k := big.NewInt(int64(6))
	k12 := big.NewInt(int64(12))
	l := big.NewFloat(float64(13591409))
	lc := big.NewFloat(float64(545140134))
	x := big.NewFloat(float64(1))
	xc := big.NewFloat(float64(-262537412640768000))
	m := big.NewFloat(float64(1))
	sum := big.NewFloat(float64(13591409))

	pi := big.NewFloat(0)

	x.SetPrec(prec)
	m.SetPrec(prec)
	sum.SetPrec(prec)
	pi.SetPrec(prec)

	bigI := big.NewInt(0)
	bigOne := big.NewInt(1)

	for ; n > 0; n-- {
		// L calculation
		l.Add(l, lc)

		// X calculation
		x.Mul(x, xc)

		// M calculation
		kpower3 := big.NewInt(0)
		kpower3.Exp(k, big.NewInt(3), nil)
		ktimes16 := new(big.Int).Mul(k, big.NewInt(16))
		mtop := big.NewFloat(0).SetPrec(prec)
		mtop.SetInt(new(big.Int).Sub(kpower3, ktimes16))
		mbot := big.NewFloat(0).SetPrec(prec)
		mbot.SetInt(new(big.Int).Exp(new(big.Int).Add(bigI, bigOne), big.NewInt(3), nil))
		mtmp := big.NewFloat(0).SetPrec(prec)
		mtmp.Quo(mtop, mbot)
		m.Mul(m, mtmp)

		// Sum calculation
		t := big.NewFloat(0).SetPrec(prec)
		t.Mul(m, l)
		t.Quo(t, x)
		sum.Add(sum, t)

		// Pi calculation
		pi.Quo(c, sum)
		k.Add(k, k12)
		bigI.Add(bigI, bigOne)
	}
	return pi, prec
}
