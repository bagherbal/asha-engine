package branchprojector

import (
	"fmt"
	"math/big"
	"strings"
)

// z satisfies z^3 - 119/60 z^2 + 8411/6480 z - 1637467/5832000 = 0.
var (
	alpha = rat("-71/30")
	beta  = rat("1071/540")
	gamma = rat("-149/216")
	dd    = rat("271/3240")
)

type zpoly [3]*big.Rat // c0 + c1 z + c2 z^2

type elem struct {
	a zpoly // a(z)
	b zpoly // b(z) eta
}

type xpoly []elem // low-to-high x coefficients

type projectorCertificate struct {
	factorA, factorB       xpoly
	bezoutA, bezoutB       xpoly
	projectorA, projectorB xpoly
	traceA, traceB         elem
	factorizationVerified  bool
	bezoutVerified         bool
	projectorAIdempotent   bool
	projectorBIdempotent   bool
	projectorSumIdentity   bool
	projectorsOrthogonal   bool
}

func buildProjectorCertificate() (projectorCertificate, error) {
	q4 := quarticPoly()
	qA, qB, err := quadraticFactors()
	if err != nil {
		return projectorCertificate{}, err
	}
	prod := mulX(qA, qB)
	factorizationVerified := equalX(prod, q4)
	if !factorizationVerified {
		return projectorCertificate{}, fmt.Errorf("Gate 188 factorization check failed: qA*qB=%s q4=%s", prod.String(), q4.String())
	}

	bezA, bezB, gcd, err := extendedGCD(qA, qB)
	if err != nil {
		return projectorCertificate{}, err
	}
	bezoutVerified := equalX(addX(mulX(bezA, qA), mulX(bezB, qB)), oneX()) && equalX(gcd, oneX())
	if !bezoutVerified {
		return projectorCertificate{}, fmt.Errorf("Gate 188 Bezout check failed: A*qA+B*qB=%s gcd=%s", addX(mulX(bezA, qA), mulX(bezB, qB)).String(), gcd.String())
	}

	// e_A is 1 mod q_A and 0 mod q_B; e_B is 0 mod q_A and 1 mod q_B.
	pA := modQ4(mulX(bezB, qB))
	pB := modQ4(mulX(bezA, qA))
	pA2 := modQ4(mulX(pA, pA))
	pB2 := modQ4(mulX(pB, pB))
	sum := modQ4(addX(pA, pB))
	orth := modQ4(mulX(pA, pB))

	return projectorCertificate{
		factorA:               qA,
		factorB:               qB,
		bezoutA:               bezA,
		bezoutB:               bezB,
		projectorA:            pA,
		projectorB:            pB,
		traceA:                traceMultiplication(pA),
		traceB:                traceMultiplication(pB),
		factorizationVerified: factorizationVerified,
		bezoutVerified:        bezoutVerified,
		projectorAIdempotent:  equalX(pA2, pA),
		projectorBIdempotent:  equalX(pB2, pB),
		projectorSumIdentity:  equalX(sum, oneX()),
		projectorsOrthogonal:  equalX(orth, zeroX()),
	}, nil
}

func quadraticFactors() (xpoly, xpoly, error) {
	z := zE()
	eta := etaE()
	two := intE(2)
	// n+q=z and q-n=eta.
	n := divE(subE(z, eta), two)
	q := divE(addE(z, eta), two)
	// q-n=eta, so m=(gamma-alpha*n)/eta.
	m := divE(subE(baseE(gamma), mulE(baseE(alpha), n)), eta)
	p := subE(baseE(alpha), m)
	qA := trimX(xpoly{n, m, oneE()})
	qB := trimX(xpoly{q, p, oneE()})
	return qA, qB, nil
}

func quarticPoly() xpoly {
	return trimX(xpoly{baseE(dd), baseE(gamma), baseE(beta), baseE(alpha), oneE()})
}

func factorSwapPreserved(c projectorCertificate) bool {
	return equalX(swapEtaX(c.factorA), c.factorB) && equalX(swapEtaX(c.factorB), c.factorA) && equalX(swapEtaX(c.projectorA), c.projectorB) && equalX(swapEtaX(c.projectorB), c.projectorA)
}

func traceMultiplication(p xpoly) elem {
	basis := []xpoly{{oneE()}, {zeroE(), oneE()}, {zeroE(), zeroE(), oneE()}, {zeroE(), zeroE(), zeroE(), oneE()}}
	tr := zeroE()
	for j, b := range basis {
		col := modQ4(mulX(p, b))
		if j < len(col) {
			tr = addE(tr, col[j])
		}
	}
	return tr
}

func modQ4(p xpoly) xpoly {
	p = trimX(copyX(p))
	for p.degree() >= 4 {
		k := p.degree()
		lc := p[k]
		shift := k - 4
		// subtract lc*x^shift*q4, with q4 monic. This cancels x^k.
		p[shift+4] = subE(p[shift+4], lc)
		p[shift+3] = subE(p[shift+3], mulE(lc, baseE(alpha)))
		p[shift+2] = subE(p[shift+2], mulE(lc, baseE(beta)))
		p[shift+1] = subE(p[shift+1], mulE(lc, baseE(gamma)))
		p[shift] = subE(p[shift], mulE(lc, baseE(dd)))
		p = trimX(p)
	}
	return p
}

func extendedGCD(a, b xpoly) (xpoly, xpoly, xpoly, error) {
	r0, r1 := trimX(copyX(a)), trimX(copyX(b))
	s0, s1 := oneX(), zeroX()
	t0, t1 := zeroX(), oneX()
	for !r1.isZero() {
		q, r, err := divmodX(r0, r1)
		if err != nil {
			return nil, nil, nil, err
		}
		r0, r1 = r1, r
		s0, s1 = s1, subX(s0, mulX(q, s1))
		t0, t1 = t1, subX(t0, mulX(q, t1))
	}
	lc := r0.leading()
	invLC := invE(lc)
	r0 = scalarMulX(invLC, r0)
	s0 = scalarMulX(invLC, s0)
	t0 = scalarMulX(invLC, t0)
	return trimX(s0), trimX(t0), trimX(r0), nil
}

func divmodX(a, b xpoly) (xpoly, xpoly, error) {
	a = trimX(copyX(a))
	b = trimX(copyX(b))
	if b.isZero() {
		return nil, nil, fmt.Errorf("division by zero polynomial")
	}
	q := make(xpoly, max(1, a.degree()-b.degree()+1))
	for i := range q {
		q[i] = zeroE()
	}
	r := a
	invLead := invE(b.leading())
	for !r.isZero() && r.degree() >= b.degree() {
		shift := r.degree() - b.degree()
		coef := mulE(r.leading(), invLead)
		q[shift] = addE(q[shift], coef)
		term := shiftX(scalarMulX(coef, b), shift)
		r = trimX(subX(r, term))
	}
	return trimX(q), trimX(r), nil
}

func oneX() xpoly  { return xpoly{oneE()} }
func zeroX() xpoly { return xpoly{} }

func copyX(p xpoly) xpoly {
	out := make(xpoly, len(p))
	copy(out, p)
	return out
}

func trimX(p xpoly) xpoly {
	for len(p) > 0 && p[len(p)-1].IsZero() {
		p = p[:len(p)-1]
	}
	return p
}

func (p xpoly) degree() int {
	p = trimX(p)
	if len(p) == 0 {
		return -1
	}
	return len(p) - 1
}
func (p xpoly) isZero() bool { return len(trimX(p)) == 0 }
func (p xpoly) leading() elem {
	p = trimX(p)
	if len(p) == 0 {
		return zeroE()
	}
	return p[len(p)-1]
}

func addX(a, b xpoly) xpoly {
	n := max(len(a), len(b))
	out := make(xpoly, n)
	for i := 0; i < n; i++ {
		out[i] = addE(getX(a, i), getX(b, i))
	}
	return trimX(out)
}
func subX(a, b xpoly) xpoly {
	n := max(len(a), len(b))
	out := make(xpoly, n)
	for i := 0; i < n; i++ {
		out[i] = subE(getX(a, i), getX(b, i))
	}
	return trimX(out)
}
func mulX(a, b xpoly) xpoly {
	if len(a) == 0 || len(b) == 0 {
		return zeroX()
	}
	out := make(xpoly, len(a)+len(b)-1)
	for i := range out {
		out[i] = zeroE()
	}
	for i := range a {
		for j := range b {
			out[i+j] = addE(out[i+j], mulE(a[i], b[j]))
		}
	}
	return trimX(out)
}
func scalarMulX(c elem, p xpoly) xpoly {
	out := make(xpoly, len(p))
	for i := range p {
		out[i] = mulE(c, p[i])
	}
	return trimX(out)
}
func shiftX(p xpoly, n int) xpoly {
	if len(p) == 0 {
		return zeroX()
	}
	out := make(xpoly, len(p)+n)
	for i := 0; i < n; i++ {
		out[i] = zeroE()
	}
	copy(out[n:], p)
	return trimX(out)
}
func getX(p xpoly, i int) elem {
	if i < 0 || i >= len(p) {
		return zeroE()
	}
	return p[i]
}
func equalX(a, b xpoly) bool {
	a, b = trimX(a), trimX(b)
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !a[i].Equal(b[i]) {
			return false
		}
	}
	return true
}

func swapEtaX(p xpoly) xpoly {
	out := make(xpoly, len(p))
	for i := range p {
		out[i] = elem{a: p[i].a, b: negZ(p[i].b)}
	}
	return trimX(out)
}

func (p xpoly) String() string {
	p = trimX(p)
	if len(p) == 0 {
		return "0"
	}
	parts := []string{}
	for i := len(p) - 1; i >= 0; i-- {
		if p[i].IsZero() {
			continue
		}
		coef := p[i].String()
		suffix := ""
		switch i {
		case 0:
			suffix = ""
		case 1:
			suffix = "*x"
		default:
			suffix = fmt.Sprintf("*x^%d", i)
		}
		parts = append(parts, coef+suffix)
	}
	return strings.Join(parts, " + ")
}

func zeroZ() zpoly          { return zpoly{new(big.Rat), new(big.Rat), new(big.Rat)} }
func oneZ() zpoly           { return zpoly{rat("1"), new(big.Rat), new(big.Rat)} }
func zZ() zpoly             { return zpoly{new(big.Rat), rat("1"), new(big.Rat)} }
func z2Z() zpoly            { return zpoly{new(big.Rat), new(big.Rat), rat("1")} }
func ratZ(r *big.Rat) zpoly { return zpoly{cloneRat(r), new(big.Rat), new(big.Rat)} }

func addZ(a, b zpoly) zpoly { return zpoly{addR(a[0], b[0]), addR(a[1], b[1]), addR(a[2], b[2])} }
func subZ(a, b zpoly) zpoly { return zpoly{subR(a[0], b[0]), subR(a[1], b[1]), subR(a[2], b[2])} }
func negZ(a zpoly) zpoly    { return zpoly{negR(a[0]), negR(a[1]), negR(a[2])} }

func mulZ(a, b zpoly) zpoly {
	raw := make([]*big.Rat, 5)
	for i := range raw {
		raw[i] = new(big.Rat)
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			raw[i+j] = addR(raw[i+j], mulR(a[i], b[j]))
		}
	}
	// z^3 = 119/60 z^2 - 8411/6480 z + 1637467/5832000
	for k := 4; k >= 3; k-- {
		coef := raw[k]
		if coef.Sign() == 0 {
			continue
		}
		raw[k] = new(big.Rat)
		raw[k-1] = addR(raw[k-1], mulR(coef, rat("119/60")))
		raw[k-2] = subR(raw[k-2], mulR(coef, rat("8411/6480")))
		raw[k-3] = addR(raw[k-3], mulR(coef, rat("1637467/5832000")))
	}
	return zpoly{raw[0], raw[1], raw[2]}
}

func invZ(a zpoly) zpoly {
	if isZeroZ(a) {
		panic("inverse of zero in Q[z]/r3")
	}
	basis := []zpoly{oneZ(), zZ(), z2Z()}
	mat := make([][]*big.Rat, 3)
	for row := 0; row < 3; row++ {
		mat[row] = make([]*big.Rat, 4)
		for col := 0; col < 3; col++ {
			prod := mulZ(a, basis[col])
			mat[row][col] = cloneRat(prod[row])
		}
		mat[row][3] = new(big.Rat)
	}
	mat[0][3] = rat("1")
	sol := solveRat(mat)
	return zpoly{sol[0], sol[1], sol[2]}
}

func isZeroZ(a zpoly) bool { return a[0].Sign() == 0 && a[1].Sign() == 0 && a[2].Sign() == 0 }
func equalZ(a, b zpoly) bool {
	return a[0].Cmp(b[0]) == 0 && a[1].Cmp(b[1]) == 0 && a[2].Cmp(b[2]) == 0
}

func zeroE() elem           { return elem{a: zeroZ(), b: zeroZ()} }
func oneE() elem            { return elem{a: oneZ(), b: zeroZ()} }
func intE(n int64) elem     { return baseE(new(big.Rat).SetInt64(n)) }
func baseE(r *big.Rat) elem { return elem{a: ratZ(r), b: zeroZ()} }
func zE() elem              { return elem{a: zZ(), b: zeroZ()} }
func etaE() elem            { return elem{a: zeroZ(), b: oneZ()} }

func etaSquareZ() zpoly { return subZ(z2Z(), ratZ(rat("271/810"))) }

func addE(a, b elem) elem { return elem{a: addZ(a.a, b.a), b: addZ(a.b, b.b)} }
func subE(a, b elem) elem { return elem{a: subZ(a.a, b.a), b: subZ(a.b, b.b)} }
func negE(a elem) elem    { return elem{a: negZ(a.a), b: negZ(a.b)} }
func mulE(a, b elem) elem {
	ac := mulZ(a.a, b.a)
	bd := mulZ(mulZ(a.b, b.b), etaSquareZ())
	adbc := addZ(mulZ(a.a, b.b), mulZ(a.b, b.a))
	return elem{a: addZ(ac, bd), b: adbc}
}
func divE(a, b elem) elem { return mulE(a, invE(b)) }

func invE(e elem) elem {
	// (a+b eta)^-1=(a-b eta)/(a^2-b^2 eta^2)
	den := subZ(mulZ(e.a, e.a), mulZ(mulZ(e.b, e.b), etaSquareZ()))
	invDen := invZ(den)
	return elem{a: mulZ(e.a, invDen), b: negZ(mulZ(e.b, invDen))}
}

func (e elem) Equal(o elem) bool { return equalZ(e.a, o.a) && equalZ(e.b, o.b) }
func (e elem) IsZero() bool      { return isZeroZ(e.a) && isZeroZ(e.b) }
func (e elem) String() string {
	if e.IsZero() {
		return "0"
	}
	parts := []string{}
	if !isZeroZ(e.a) {
		parts = append(parts, zString(e.a))
	}
	if !isZeroZ(e.b) {
		parts = append(parts, "("+zString(e.b)+")*eta")
	}
	return strings.Join(parts, " + ")
}

func zString(p zpoly) string {
	terms := []string{}
	names := []string{"", "z", "z^2"}
	for i := 2; i >= 0; i-- {
		if p[i].Sign() == 0 {
			continue
		}
		s := p[i].RatString()
		if i == 0 {
			terms = append(terms, s)
		} else {
			terms = append(terms, s+"*"+names[i])
		}
	}
	if len(terms) == 0 {
		return "0"
	}
	return strings.Join(terms, " + ")
}

func solveRat(a [][]*big.Rat) []*big.Rat {
	n := len(a)
	for col := 0; col < n; col++ {
		pivot := -1
		for row := col; row < n; row++ {
			if a[row][col].Sign() != 0 {
				pivot = row
				break
			}
		}
		if pivot < 0 {
			panic("singular rational system")
		}
		if pivot != col {
			a[pivot], a[col] = a[col], a[pivot]
		}
		inv := new(big.Rat).Inv(a[col][col])
		for j := col; j <= n; j++ {
			a[col][j] = mulR(a[col][j], inv)
		}
		for row := 0; row < n; row++ {
			if row == col || a[row][col].Sign() == 0 {
				continue
			}
			factor := cloneRat(a[row][col])
			for j := col; j <= n; j++ {
				a[row][j] = subR(a[row][j], mulR(factor, a[col][j]))
			}
		}
	}
	out := make([]*big.Rat, n)
	for i := 0; i < n; i++ {
		out[i] = cloneRat(a[i][n])
	}
	return out
}

func rat(s string) *big.Rat {
	r, ok := new(big.Rat).SetString(s)
	if !ok {
		panic("invalid rational: " + s)
	}
	return r
}
func cloneRat(x *big.Rat) *big.Rat { return new(big.Rat).Set(x) }
func addR(x, y *big.Rat) *big.Rat  { return new(big.Rat).Add(x, y) }
func subR(x, y *big.Rat) *big.Rat  { return new(big.Rat).Sub(x, y) }
func mulR(x, y *big.Rat) *big.Rat  { return new(big.Rat).Mul(x, y) }
func negR(x *big.Rat) *big.Rat     { return new(big.Rat).Neg(x) }
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
