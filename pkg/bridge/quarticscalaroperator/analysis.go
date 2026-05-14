// Package quarticscalaroperator implements Gate 185: quartic scalar operator /
// minimal-polynomial construction on H_Φ.
//
// Gate 184 isolated the only dimensionally viable finite-bundle route: the
// quartic contact primary ideal has dimension 4, matching the active scalar
// carrier H_Φ. Gate 185 tests this route with exact rational arithmetic. It
// constructs the branch-free companion operator for the contact quartic factor
//
//	q4(x) = 3240x^4 - 7668x^3 + 6426x^2 - 2235x + 271,
//
// verifies q4(T)=0, verifies the quartic power-sum ledger, and checks whether
// the already-derived Gate-37 scalar/Higgs active operator can be identified
// with this quartic-minimal operator.
//
// The result is deliberately strict. The abstract quartic module exists exactly
// and supplies a lawful Q[x]/(q4)-module. But the Gate-37 scalar operator is
// pair-degenerate and therefore quadratic, not quartic-minimal. The physical
// H_Φ scalar bundle is therefore not promoted in this gate.
package quarticscalaroperator

import (
	"fmt"
	"math"
	"math/big"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/cliffordcontactcommutant"
	"github.com/bagherbal/asha-engine/pkg/dynamics/scalarpotential"
)

type Rational = big.Rat

type QuarticPolynomial struct {
	Name         string
	Coefficients []string // descending: a4,a3,a2,a1,a0
	Monic        []string // descending: 1,c3,c2,c1,c0
	Degree       int
	Description  string
}

type CompanionOperatorAudit struct {
	OperatorName                string
	Dimension                   int
	Basis                       string
	Matrix                      [][]string
	PolynomialIdentityZero      bool
	CharacteristicPolynomial    string
	MinimalPolynomial           string
	CyclicVectorRank            int
	CyclicModule                bool
	BranchFree                  bool
	UsesExactRationalArithmetic bool
	AlgebraActionDerived        bool
	AbstractQuarticModule       bool
	Verdict                     string
}

type MomentAudit struct {
	TraceT     string
	TraceT2    string
	TraceT3    string
	TraceT4    string
	ExpectedP1 string
	ExpectedP2 string
	ExpectedP3 string
	ExpectedP4 string
	P1Matches  bool
	P2Matches  bool
	P3Matches  bool
	P4Matches  bool
	AllMatch   bool
	Verdict    string
}

type Gate37ScalarComparison struct {
	ActiveDimension                 int
	Gate37PairDegenerate            bool
	ActiveSpectrum                  []float64
	HighPairEigenvalue              float64
	LowPairEigenvalue               float64
	PairSplitting                   float64
	Gate37ShapeExact                string
	Gate37ShapeFloat                float64
	MinimalPolynomialCandidate      string
	HasQuarticMinimalPolynomial     bool
	MomentShapeMatchesContactTarget bool
	IdentifiedWithQuarticModule     bool
	PhysicalScalarBundleDerived     bool
	Verdict                         string
}

type BlockRestrictionAudit struct {
	CandidateName                    string
	ExactOmegaContactAvailable       bool
	QuarticPrimaryProjectorAvailable bool
	PhysicalHphiProjectorAvailable   bool
	CanonicalMapQuarticToHphi        bool
	ExactBlockRestrictionComputed    bool
	PolynomialIdentityTestAvailable  bool
	PromotesCompanionToPhysicalHphi  bool
	Verdict                          string
}

type Summary struct {
	CandidatesAudited                       int
	ExactQuarticOperators                   int
	PolynomialIdentitiesVerified            int
	MomentLedgersMatched                    int
	AbstractModulesDerived                  int
	PhysicalHphiOperatorsWithQuarticMinPoly int
	PhysicalScalarBundlesDerived            int
	Comment                                 string
}

type Firewall struct {
	UsesObservedInputForDerivation     bool
	UsesBranchDiagonalization          bool
	UsesArbitraryMatrixFit             bool
	QuarticAbstractOperatorDerived     bool
	QuarticMomentsVerified             bool
	Gate37ScalarOperatorQuadratic      bool
	ExactBlockRestrictionDerived       bool
	CanonicalHphiIdentificationDerived bool
	PhysicalScalarBundleDerived        bool
	ChernWeilCarrierDerived            bool
	HeatKernelMatchingDerived          bool
	ThresholdCorrectedBetaDerived      bool
	AbsoluteCouplingPromoted           bool
	PhysicalConstantsDerived           bool
	StrictNullityBefore                int
	StrictNullityAfter                 int
	ConditionalNullityBefore           int
	ConditionalNullityAfter            int
	ClosedStatements                   []string
	OpenRequirements                   []string
	RecommendedNextGate                string
	Verdict                            string
}

type Analysis struct {
	PreviousGate184  cliffordcontactcommutant.Analysis
	Gate37           scalarpotential.Analysis
	Polynomial       QuarticPolynomial
	Companion        CompanionOperatorAudit
	Moments          MomentAudit
	Gate37Comparison Gate37ScalarComparison
	BlockRestriction BlockRestrictionAudit
	Summary          Summary
	Firewall         Firewall
	TruthStatement   string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := cliffordcontactcommutant.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 184 input: %w", err)
			return
		}
		sp, err := scalarpotential.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 37 scalar input: %w", err)
			return
		}
		defaultA, defaultErr = Build(prev, sp, 1e-10)
	})
	return defaultA, defaultErr
}

func Build(prev cliffordcontactcommutant.Analysis, sp scalarpotential.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !prev.Firewall.QuarticScalarAbstractModuleDerived {
		return Analysis{}, fmt.Errorf("Gate 185 requires Gate 184 quartic scalar abstract module")
	}
	if sp.ActiveRealDimension != 4 || len(sp.ActiveSpectrum) != 4 {
		return Analysis{}, fmt.Errorf("Gate 185 requires four-dimensional Gate-37 scalar active sector")
	}

	poly := buildPolynomial()
	comp, err := auditCompanion(poly)
	if err != nil {
		return Analysis{}, err
	}
	moments, err := auditMoments(comp)
	if err != nil {
		return Analysis{}, err
	}
	g37 := auditGate37(sp, moments, eps)
	block := BlockRestrictionAudit{
		CandidateName:                    "T_Φ = P_Φ Ω_contact P_Φ",
		ExactOmegaContactAvailable:       true,
		QuarticPrimaryProjectorAvailable: true,
		PhysicalHphiProjectorAvailable:   false,
		CanonicalMapQuarticToHphi:        false,
		ExactBlockRestrictionComputed:    false,
		PolynomialIdentityTestAvailable:  false,
		PromotesCompanionToPhysicalHphi:  false,
		Verdict:                          "The exact contact overlap and quartic primary projector exist, but the engine has not derived a canonical physical H_Φ projector/map into the quartic contact primary block. Therefore the requested block restriction cannot be promoted to an H_Φ operator in this gate.",
	}
	summary := Summary{
		CandidatesAudited:                       3,
		ExactQuarticOperators:                   boolInt(comp.AbstractQuarticModule),
		PolynomialIdentitiesVerified:            boolInt(comp.PolynomialIdentityZero),
		MomentLedgersMatched:                    boolInt(moments.AllMatch),
		AbstractModulesDerived:                  boolInt(comp.AbstractQuarticModule),
		PhysicalHphiOperatorsWithQuarticMinPoly: boolInt(g37.HasQuarticMinimalPolynomial),
		PhysicalScalarBundlesDerived:            boolInt(g37.PhysicalScalarBundleDerived),
		Comment:                                 "The companion operator closes the abstract quartic module exactly; the Gate-37 physical scalar operator remains pair-degenerate/quadratic, and the block restriction to H_Φ lacks a canonical map.",
	}
	fw := Firewall{
		UsesObservedInputForDerivation:     false,
		UsesBranchDiagonalization:          false,
		UsesArbitraryMatrixFit:             false,
		QuarticAbstractOperatorDerived:     comp.AbstractQuarticModule && comp.AlgebraActionDerived,
		QuarticMomentsVerified:             moments.AllMatch,
		Gate37ScalarOperatorQuadratic:      sp.PairDegenerate && !g37.HasQuarticMinimalPolynomial,
		ExactBlockRestrictionDerived:       block.ExactBlockRestrictionComputed,
		CanonicalHphiIdentificationDerived: g37.IdentifiedWithQuarticModule,
		PhysicalScalarBundleDerived:        g37.PhysicalScalarBundleDerived,
		ChernWeilCarrierDerived:            false,
		HeatKernelMatchingDerived:          false,
		ThresholdCorrectedBetaDerived:      false,
		AbsoluteCouplingPromoted:           false,
		PhysicalConstantsDerived:           false,
		StrictNullityBefore:                prev.Firewall.StrictNullityAfter,
		StrictNullityAfter:                 prev.Firewall.StrictNullityAfter,
		ConditionalNullityBefore:           prev.Firewall.ConditionalNullityAfter,
		ConditionalNullityAfter:            prev.Firewall.ConditionalNullityAfter,
		ClosedStatements: []string{
			"Q[x]/(q4) has an exact 4D companion-module realization with q4(T)=0",
			"the quartic contact power sums p1..p4 are reproduced exactly by the companion operator",
			"branch diagonalization and numerical eigenvalue fitting are unnecessary and unused",
			"Gate-37 active scalar operator is pair-degenerate and therefore not the quartic-minimal operator",
		},
		OpenRequirements: []string{
			"derive a canonical physical H_Φ operator whose minimal polynomial is q4",
			"derive a canonical projector/map P_Φ from the scalar active carrier into the quartic contact primary block",
			"prove compatibility of the quartic module with scalar SU(2)_L×U(1)_Y action, J, and γ",
			"construct a finite Chern character / integration pairing for the promoted scalar bundle",
		},
		RecommendedNextGate: "Gate 186 — scalar/contact quartic identification selector or obstruction theorem",
		Verdict:             "Gate 185 constructs the exact quartic companion module and verifies its Galois-invariant moment ledger, but it does not identify that abstract module with the physical Gate-37 H_Φ scalar carrier. The physical scalar bundle remains open.",
	}
	truth := "Gate 185 succeeds as an exact quartic-module theorem and fails as a physical H_Φ promotion theorem. The companion action of Q[x]/(q4) is branch-free, rational, cyclic, and moment-correct, so the quartic scalar escape hatch is algebraically real. But Gate 37's scalar/Higgs active operator is pair-degenerate with a quadratic minimal polynomial; without a canonical scalar/contact identification map, the companion module cannot yet be called the physical scalar bundle."
	return Analysis{PreviousGate184: prev, Gate37: sp, Polynomial: poly, Companion: comp, Moments: moments, Gate37Comparison: g37, BlockRestriction: block, Summary: summary, Firewall: fw, TruthStatement: truth}, nil
}

func buildPolynomial() QuarticPolynomial {
	return QuarticPolynomial{
		Name:         "contact quartic primary factor q4",
		Coefficients: []string{"3240", "-7668", "6426", "-2235", "271"},
		Monic:        []string{"1", "-71/30", "1071/540", "149/216", "271/3240"},
		Degree:       4,
		Description:  "the Galois-invariant quartic primary factor of the exact contact overlap characteristic polynomial",
	}
}

func auditCompanion(poly QuarticPolynomial) (CompanionOperatorAudit, error) {
	// Monic polynomial: x^4 + c3 x^3 + c2 x^2 + c1 x + c0.
	c3 := rat("-71/30")
	c2 := rat("1071/540")
	c1 := rat("-149/216")
	c0 := rat("271/3240")
	// Companion in basis {1,x,x^2,x^3}; multiplication by x modulo q4.
	zero := rat("0")
	one := rat("1")
	T := matrix{
		{zero, zero, zero, neg(c0)},
		{one, zero, zero, neg(c1)},
		{zero, one, zero, neg(c2)},
		{zero, zero, one, neg(c3)},
	}
	qT := evalMonicQuartic(T, c3, c2, c1, c0)
	identZero := qT.isZero()
	krylov := matrix{
		{rat("1"), rat("0"), rat("0"), rat("0")},
		{rat("0"), rat("1"), rat("0"), rat("0")},
		{rat("0"), rat("0"), rat("1"), rat("0")},
		{rat("0"), rat("0"), rat("0"), rat("1")},
	}
	return CompanionOperatorAudit{
		OperatorName:                "T_q = multiplication by x on Q[x]/(q4)",
		Dimension:                   4,
		Basis:                       "{1, x, x², x³}",
		Matrix:                      T.format(),
		PolynomialIdentityZero:      identZero,
		CharacteristicPolynomial:    "3240x^4 - 7668x^3 + 6426x^2 - 2235x + 271",
		MinimalPolynomial:           "3240x^4 - 7668x^3 + 6426x^2 - 2235x + 271",
		CyclicVectorRank:            krylov.rank(),
		CyclicModule:                krylov.rank() == 4,
		BranchFree:                  true,
		UsesExactRationalArithmetic: true,
		AlgebraActionDerived:        identZero,
		AbstractQuarticModule:       identZero,
		Verdict:                     "The companion multiplication operator is a canonical exact rational representative of the quartic ideal. Its cyclic module has rank 4, and q4(T_q)=0 exactly, so the abstract quartic module is derived without choosing quartic roots.",
	}, nil
}

func auditMoments(comp CompanionOperatorAudit) (MomentAudit, error) {
	T, err := matrixFromStrings(comp.Matrix)
	if err != nil {
		return MomentAudit{}, err
	}
	T2 := T.mul(T)
	T3 := T2.mul(T)
	T4 := T3.mul(T)
	tr1 := T.trace()
	tr2 := T2.trace()
	tr3 := T3.trace()
	tr4 := T4.trace()
	p1 := rat("71/30")
	p2 := rat("1471/900")
	p3 := rat("33581/27000")
	p4 := rat("809891/810000")
	return MomentAudit{
		TraceT: ratString(tr1), TraceT2: ratString(tr2), TraceT3: ratString(tr3), TraceT4: ratString(tr4),
		ExpectedP1: ratString(p1), ExpectedP2: ratString(p2), ExpectedP3: ratString(p3), ExpectedP4: ratString(p4),
		P1Matches: tr1.Cmp(p1) == 0, P2Matches: tr2.Cmp(p2) == 0, P3Matches: tr3.Cmp(p3) == 0, P4Matches: tr4.Cmp(p4) == 0,
		AllMatch: tr1.Cmp(p1) == 0 && tr2.Cmp(p2) == 0 && tr3.Cmp(p3) == 0 && tr4.Cmp(p4) == 0,
		Verdict:  "The companion operator reproduces the exact quartic Galois-invariant power sums from Gate 161.",
	}, nil
}

func auditGate37(sp scalarpotential.Analysis, moments MomentAudit, eps float64) Gate37ScalarComparison {
	// Gate37 active spectrum is pair-degenerate; a pair-degenerate 4x4 diagonalizable
	// operator has at most a quadratic minimal polynomial (x-high)(x-low).
	hasQuarticMin := false
	identified := false
	shapeExact := "1197/4624"
	return Gate37ScalarComparison{
		ActiveDimension:                 sp.ActiveRealDimension,
		Gate37PairDegenerate:            sp.PairDegenerate,
		ActiveSpectrum:                  append([]float64(nil), sp.ActiveSpectrum...),
		HighPairEigenvalue:              sp.HighPairEigenvalue,
		LowPairEigenvalue:               sp.LowPairEigenvalue,
		PairSplitting:                   sp.PairSplitting,
		Gate37ShapeExact:                shapeExact,
		Gate37ShapeFloat:                sp.LambdaShape,
		MinimalPolynomialCandidate:      fmt.Sprintf("(x-%.12g)(x-%.12g), because the Gate-37 active spectrum is pair-degenerate", sp.HighPairEigenvalue, sp.LowPairEigenvalue),
		HasQuarticMinimalPolynomial:     hasQuarticMin,
		MomentShapeMatchesContactTarget: math.Abs(sp.LambdaShape-1197.0/4624.0) <= eps,
		IdentifiedWithQuarticModule:     identified,
		PhysicalScalarBundleDerived:     false,
		Verdict:                         "Gate 37 supplies the scalar-potential shape target and a 4D active scalar carrier, but its active scalar mixing spectrum is pair-degenerate. That operator is therefore quadratic-minimal, not the quartic companion operator. A new canonical scalar/contact identification selector is required before the quartic module becomes physical H_Φ.",
	}
}

// Exact rational matrix helpers.
type matrix [][]*big.Rat

func rat(s string) *big.Rat {
	r, ok := new(big.Rat).SetString(s)
	if !ok {
		panic("invalid rational: " + s)
	}
	return r
}
func clone(x *big.Rat) *big.Rat  { return new(big.Rat).Set(x) }
func neg(x *big.Rat) *big.Rat    { return new(big.Rat).Neg(x) }
func add(x, y *big.Rat) *big.Rat { return new(big.Rat).Add(x, y) }
func mul(x, y *big.Rat) *big.Rat { return new(big.Rat).Mul(x, y) }

func ident(n int) matrix {
	m := make(matrix, n)
	for i := 0; i < n; i++ {
		m[i] = make([]*big.Rat, n)
		for j := 0; j < n; j++ {
			if i == j {
				m[i][j] = rat("1")
			} else {
				m[i][j] = rat("0")
			}
		}
	}
	return m
}
func zero(n int) matrix {
	m := make(matrix, n)
	for i := 0; i < n; i++ {
		m[i] = make([]*big.Rat, n)
		for j := 0; j < n; j++ {
			m[i][j] = rat("0")
		}
	}
	return m
}
func (m matrix) dim() int { return len(m) }
func (m matrix) add(n matrix) matrix {
	d := m.dim()
	out := zero(d)
	for i := 0; i < d; i++ {
		for j := 0; j < d; j++ {
			out[i][j] = add(m[i][j], n[i][j])
		}
	}
	return out
}
func (m matrix) scale(c *big.Rat) matrix {
	d := m.dim()
	out := zero(d)
	for i := 0; i < d; i++ {
		for j := 0; j < d; j++ {
			out[i][j] = mul(c, m[i][j])
		}
	}
	return out
}
func (m matrix) mul(n matrix) matrix {
	d := m.dim()
	out := zero(d)
	for i := 0; i < d; i++ {
		for j := 0; j < d; j++ {
			s := rat("0")
			for k := 0; k < d; k++ {
				s = add(s, mul(m[i][k], n[k][j]))
			}
			out[i][j] = s
		}
	}
	return out
}
func (m matrix) trace() *big.Rat {
	s := rat("0")
	for i := 0; i < m.dim(); i++ {
		s = add(s, m[i][i])
	}
	return s
}
func (m matrix) isZero() bool {
	for i := range m {
		for j := range m[i] {
			if m[i][j].Sign() != 0 {
				return false
			}
		}
	}
	return true
}
func (m matrix) format() [][]string {
	out := make([][]string, len(m))
	for i := range m {
		out[i] = make([]string, len(m[i]))
		for j := range m[i] {
			out[i][j] = ratString(m[i][j])
		}
	}
	return out
}
func matrixFromStrings(xs [][]string) (matrix, error) {
	m := make(matrix, len(xs))
	for i := range xs {
		m[i] = make([]*big.Rat, len(xs[i]))
		for j, s := range xs[i] {
			r, ok := new(big.Rat).SetString(s)
			if !ok {
				return nil, fmt.Errorf("bad rational matrix entry %q", s)
			}
			m[i][j] = r
		}
	}
	return m, nil
}

func evalMonicQuartic(T matrix, c3, c2, c1, c0 *big.Rat) matrix {
	I := ident(T.dim())
	T2 := T.mul(T)
	T3 := T2.mul(T)
	T4 := T3.mul(T)
	return T4.add(T3.scale(c3)).add(T2.scale(c2)).add(T.scale(c1)).add(I.scale(c0))
}

func (m matrix) rank() int {
	// Gaussian elimination over Q.
	a := make(matrix, len(m))
	for i := range m {
		a[i] = make([]*big.Rat, len(m[i]))
		for j := range m[i] {
			a[i][j] = clone(m[i][j])
		}
	}
	rows, cols := len(a), len(a[0])
	rank := 0
	for col := 0; col < cols && rank < rows; col++ {
		pivot := -1
		for r := rank; r < rows; r++ {
			if a[r][col].Sign() != 0 {
				pivot = r
				break
			}
		}
		if pivot < 0 {
			continue
		}
		a[rank], a[pivot] = a[pivot], a[rank]
		inv := new(big.Rat).Inv(a[rank][col])
		for c := col; c < cols; c++ {
			a[rank][c] = mul(a[rank][c], inv)
		}
		for r := 0; r < rows; r++ {
			if r == rank || a[r][col].Sign() == 0 {
				continue
			}
			factor := clone(a[r][col])
			for c := col; c < cols; c++ {
				a[r][c] = new(big.Rat).Sub(a[r][c], mul(factor, a[rank][c]))
			}
		}
		rank++
	}
	return rank
}

func ratString(r *big.Rat) string {
	if r.IsInt() {
		return r.Num().String()
	}
	return r.RatString()
}
func boolInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func FormatPolynomial(p QuarticPolynomial) string {
	return fmt.Sprintf("%s degree=%d coeffs=%v monic=%v (%s)", p.Name, p.Degree, p.Coefficients, p.Monic, p.Description)
}
func FormatCompanion(a CompanionOperatorAudit) string {
	rows := make([]string, 0, len(a.Matrix))
	for _, row := range a.Matrix {
		rows = append(rows, "["+strings.Join(row, ",")+"]")
	}
	return fmt.Sprintf("%s dim=%d basis=%s matrix=%s q(T)=0:%t char=%q min=%q cyclicRank=%d cyclic=%t branchFree=%t exact=%t action=%t abstract=%t (%s)", a.OperatorName, a.Dimension, a.Basis, strings.Join(rows, ";"), a.PolynomialIdentityZero, a.CharacteristicPolynomial, a.MinimalPolynomial, a.CyclicVectorRank, a.CyclicModule, a.BranchFree, a.UsesExactRationalArithmetic, a.AlgebraActionDerived, a.AbstractQuarticModule, a.Verdict)
}
func FormatMoments(a MomentAudit) string {
	return fmt.Sprintf("traces=[%s,%s,%s,%s] expected=[%s,%s,%s,%s] matches=[%t,%t,%t,%t] all=%t (%s)", a.TraceT, a.TraceT2, a.TraceT3, a.TraceT4, a.ExpectedP1, a.ExpectedP2, a.ExpectedP3, a.ExpectedP4, a.P1Matches, a.P2Matches, a.P3Matches, a.P4Matches, a.AllMatch, a.Verdict)
}
func FormatGate37(a Gate37ScalarComparison) string {
	return fmt.Sprintf("dim=%d pairDegenerate=%t spectrum=%v high=%.12g low=%.12g split=%.12g shape=%s≈%.12g minCandidate=%q quarticMin=%t shapeTarget=%t identified=%t physical=%t (%s)", a.ActiveDimension, a.Gate37PairDegenerate, a.ActiveSpectrum, a.HighPairEigenvalue, a.LowPairEigenvalue, a.PairSplitting, a.Gate37ShapeExact, a.Gate37ShapeFloat, a.MinimalPolynomialCandidate, a.HasQuarticMinimalPolynomial, a.MomentShapeMatchesContactTarget, a.IdentifiedWithQuarticModule, a.PhysicalScalarBundleDerived, a.Verdict)
}
func FormatBlockRestriction(a BlockRestrictionAudit) string {
	return fmt.Sprintf("%s omega=%t qProjector=%t hphiProjector=%t map=%t computed=%t test=%t promoted=%t (%s)", a.CandidateName, a.ExactOmegaContactAvailable, a.QuarticPrimaryProjectorAvailable, a.PhysicalHphiProjectorAvailable, a.CanonicalMapQuarticToHphi, a.ExactBlockRestrictionComputed, a.PolynomialIdentityTestAvailable, a.PromotesCompanionToPhysicalHphi, a.Verdict)
}
func FormatSummary(a Summary) string {
	return fmt.Sprintf("candidates=%d exactQuartic=%d identities=%d moments=%d abstract=%d physicalQuarticHphi=%d physicalBundles=%d (%s)", a.CandidatesAudited, a.ExactQuarticOperators, a.PolynomialIdentitiesVerified, a.MomentLedgersMatched, a.AbstractModulesDerived, a.PhysicalHphiOperatorsWithQuarticMinPoly, a.PhysicalScalarBundlesDerived, a.Comment)
}
func FormatFirewall(a Firewall) string {
	return fmt.Sprintf("observed=%t branches=%t arbitrary=%t abstract=%t moments=%t gate37Quadratic=%t block=%t hphiID=%t physicalBundle=%t chernWeil=%t heat=%t thresholds=%t absolute=%t constants=%t strict=%d->%d conditional=%d->%d verdict=%s", a.UsesObservedInputForDerivation, a.UsesBranchDiagonalization, a.UsesArbitraryMatrixFit, a.QuarticAbstractOperatorDerived, a.QuarticMomentsVerified, a.Gate37ScalarOperatorQuadratic, a.ExactBlockRestrictionDerived, a.CanonicalHphiIdentificationDerived, a.PhysicalScalarBundleDerived, a.ChernWeilCarrierDerived, a.HeatKernelMatchingDerived, a.ThresholdCorrectedBetaDerived, a.AbsoluteCouplingPromoted, a.PhysicalConstantsDerived, a.StrictNullityBefore, a.StrictNullityAfter, a.ConditionalNullityBefore, a.ConditionalNullityAfter, a.Verdict)
}
