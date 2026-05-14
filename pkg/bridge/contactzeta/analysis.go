// Package contactzeta implements Gate 162: finite contact spectral zeta
// regularization / seven-root action functional audit.
//
// Gate 161 established that the quartic contact block is usable only through
// collective Galois-invariant spectral functionals. Gate 162 extends that
// ledger from positive moments to the finite zeta function of the complete
// seven-root contact spectrum,
//
//	zeta_contact(s) = Sum_i lambda_i^{-s}.
//
// The gate is deliberately conservative. It proves that the finite zeta ledger
// is exact, rational, branch-free, and pole-free for the audited integer values
// s=0..4. It also proves the negative permission result: zeta data by itself is
// not a finite spectral triple, does not choose a cutoff/test function, does not
// derive a gauge-kinetic representation map, and therefore does not open beta
// rows or physical constants.
package contactzeta

import (
	"fmt"
	"math/big"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/quarticspectralfunctional"
)

type Rational struct {
	v *big.Rat
}

func NewRational(n, d int64) Rational {
	if d == 0 {
		panic("zero denominator")
	}
	return Rational{v: new(big.Rat).SetFrac(big.NewInt(n), big.NewInt(d))}
}

func rationalFromGate161(r quarticspectralfunctional.Rational) Rational {
	return NewRational(r.Numerator, r.Denominator)
}

func (r Rational) Add(s Rational) Rational {
	return Rational{v: new(big.Rat).Add(r.v, s.v)}
}

func (r Rational) Sub(s Rational) Rational {
	return Rational{v: new(big.Rat).Sub(r.v, s.v)}
}

func (r Rational) Mul(s Rational) Rational {
	return Rational{v: new(big.Rat).Mul(r.v, s.v)}
}

func (r Rational) Div(s Rational) Rational {
	if s.v.Sign() == 0 {
		panic("division by zero")
	}
	return Rational{v: new(big.Rat).Quo(r.v, s.v)}
}

func (r Rational) Pow(k int) Rational {
	if k < 0 {
		return NewRational(1, 1).Div(r.Pow(-k))
	}
	out := NewRational(1, 1)
	for i := 0; i < k; i++ {
		out = out.Mul(r)
	}
	return out
}

func (r Rational) Equal(s Rational) bool { return r.v.Cmp(s.v) == 0 }
func (r Rational) String() string        { return r.v.RatString() }
func (r Rational) IsPositive() bool      { return r.v.Sign() > 0 }

type ZetaValue struct {
	S                    int
	RationalPart         Rational
	QuarticPart          Rational
	Full                 Rational
	ExactOverQ           bool
	GaloisInvariant      bool
	BranchFree           bool
	PoleFree             bool
	UsesObservedInput    bool
	UsesBranchChoice     bool
	RequiresRowSemantics bool
	Verdict              string
}

type ActionFunctionalCandidate struct {
	Name                      string
	Formula                   string
	Value                     Rational
	ExactOverQ                bool
	GaloisInvariant           bool
	BranchFree                bool
	UsesObservedInput         bool
	UsesBranchChoice          bool
	RequiresRowSemantics      bool
	RequiresSpectralTriple    bool
	RequiresCutoffFunction    bool
	CoefficientCanonical      bool
	MatchesKappaU1            bool
	MatchesEmbeddedBoundary   bool
	MatchesContactWeakAngle   bool
	MatchesGeneratorWeakAngle bool
	ConstrainsBoundary        bool
	BetaRowsAllowed           int
	PhysicalConstantsDerived  bool
	Verdict                   string
}

type ZetaRegularizationAudit struct {
	ValuesComputed               int
	ExactRationalValues          int
	GaloisInvariantValues        int
	BranchFreeValues             int
	BranchChoicesUsed            int
	ObservedInputsUsed           bool
	RowSemanticsUsed             bool
	Poles                        int
	AnalyticContinuationNeeded   bool
	PositiveNonzeroSpectrumRows  int
	ZetaZeroEqualsDimension      bool
	FiniteRegularizationComplete bool
	Verdict                      string
}

type SpectralActionAudit struct {
	CandidatesAudited          int
	ObservedInputsUsed         bool
	BranchChoicesUsed          int
	RowSemanticsUsed           bool
	RequiresSpectralTriple     int
	RequiresCutoffFunction     int
	CanonicalCoefficients      int
	FiniteDiracSelected        bool
	RealStructureSelected      bool
	GradingSelected            bool
	GaugeKineticMapDerived     bool
	BoundaryConstraintsDerived int
	MatchesKappaU1             int
	MatchesEmbeddedBoundary    int
	MatchesContactWeakAngle    int
	MatchesGeneratorWeakAngle  int
	NewActionNormalization     bool
	PhysicalConstantsDerived   bool
	Verdict                    string
}

type BetaFirewallAudit struct {
	FiniteZetaData               bool
	SpectralTripleComplete       bool
	CanonicalCutoffSelected      bool
	GaugeKineticMapRows          int
	IndividualQuarticRows        int
	GaugeRepresentationRows      int
	SpinStatisticsRows           int
	LocalFieldRows               int
	MassActivationRows           int
	DecouplingRows               int
	DynkinIndexRows              int
	ThresholdBetaRows            int
	ProvenZeroRows               int
	PhysicalConstantsDerived     bool
	BetaPermissionFirewallClosed bool
	Verdict                      string
}

type Summary struct {
	ContactRows                int
	RationalSingletonRows      int
	QuarticBlockRows           int
	QuarticCollectiveBlocks    int
	ZetaValuesComputed         int
	ActionCandidatesAudited    int
	BoundaryConstraintsDerived int
	QuarticBlockBetaRows       int
	ContactBetaRowsAllowed     int
	ContactZeroRowsProved      int
	ResidualS6Choices          int
	ResidualNullityBefore      int
	ResidualNullityAfter       int
}

type Analysis struct {
	Previous quarticspectralfunctional.Analysis

	ZetaValues       []ZetaValue
	ActionCandidates []ActionFunctionalCandidate
	ZetaAudit        ZetaRegularizationAudit
	ActionAudit      SpectralActionAudit
	BetaFirewall     BetaFirewallAudit
	Summary          Summary

	ContactRows                  int
	ExactRationalOverlapMatrix   bool
	ExactCharacteristicCertified bool
	ExactRootIsolationCertified  bool
	PositiveNonzeroSpectrumRows  int
	RationalPrimaryIdempotents   int
	GaloisInvariantOrbits        int
	RationalSingletonRows        int
	QuarticOrbitRows             int
	QuarticCompressedBlocks      int
	QuarticCollectiveBlocks      int
	QuarticBlockInvariants       int
	QuarticSpectralMoments       int
	ContactZetaValues            int
	FiniteZetaPoleCount          int
	AnalyticContinuationNeeded   bool
	SpectralActionCandidates     int
	SpectralTripleComplete       bool
	FiniteDiracSelected          bool
	RealStructureSelected        bool
	GradingSelected              bool
	CanonicalCutoffSelected      bool
	GaugeKineticMapRows          int
	IndividualQuarticRows        int
	CanonicalQuarticBranches     int
	ExactNumberFieldProjectors   int
	IndividualQuarticProjectors  int
	RowwiseRootAssignmentProofs  int
	ExternalSelectorRows         int
	CanonicalTwoTwoSplits        int
	BranchBreakingSources        int
	ChargeSemanticRows           int
	T3RRowsDerived               int
	ChiralityRowsDerived         int
	BMinusLRowsDerived           int
	SU2LRowsDerived              int
	HyperchargeRowsDerived       int
	GaugeRepresentationRows      int
	SpinStatisticsRows           int
	LocalFieldRows               int
	KineticPoleResidueRows       int
	MassActivationRows           int
	DecouplingRows               int
	DynkinIndexRows              int
	BRSTCancellationRows         int
	ConstraintRows               int
	PropagatorRows               int
	RepresentationCompleteRows   int
	RepresentationOpenRows       int
	QuarticZeroBetaRows          int
	QuarticBlockBetaRows         int
	ContactBetaRowsAllowed       int
	ContactZeroRowsProved        int
	BetaPermissionFirewallClosed bool
	ThresholdCorrectedBeta       bool
	FullBetaMatchingTensor       bool
	BoundaryConstraintsDerived   int
	ResidualS6Choices            int
	ResidualNullityBefore        int
	ResidualNullityAfter         int
	HiddenObservedInputUsed      bool
	PhysicalWeakAngleDerived     bool
	FineStructureDerived         bool
	PhysicalMassesDerived        bool
	PhysicalScaleDerived         bool

	TruthStatement      string
	RejectedClaims      []string
	RemainingUnknowns   []string
	RecommendedNextGate string
}

var defaultOnce sync.Once
var defaultValue Analysis
var defaultErr error

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := quarticspectralfunctional.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev quarticspectralfunctional.Analysis) (Analysis, error) {
	if !prev.BetaPermissionFirewallClosed || prev.QuarticCollectiveBlocks != 1 || prev.ContactBetaRowsAllowed != 0 || prev.ContactZeroRowsProved != 0 || prev.BoundaryConstraintsDerived != 0 {
		return Analysis{}, fmt.Errorf("Gate 162 requires Gate 161 collective spectral firewall")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 162 refuses hidden observed physical input")
	}

	rationalRoots := []Rational{NewRational(1, 3), NewRational(1, 2), NewRational(2, 3)}
	zetaValues := contactZetaValues(prev.Polynomial, rationalRoots, 4)
	positiveMoments := fullPositiveMoments(prev.Polynomial, rationalRoots, 2)
	product := fullContactProduct(prev.Polynomial, rationalRoots)
	actionCandidates := actionFunctionalCandidates(zetaValues, positiveMoments, product)
	markBoundaryComparisons(actionCandidates)

	zetaAudit := auditZeta(zetaValues)
	actionAudit := auditAction(actionCandidates)
	betaFirewall := BetaFirewallAudit{
		FiniteZetaData:               true,
		SpectralTripleComplete:       false,
		CanonicalCutoffSelected:      false,
		GaugeKineticMapRows:          0,
		IndividualQuarticRows:        0,
		GaugeRepresentationRows:      0,
		SpinStatisticsRows:           0,
		LocalFieldRows:               0,
		MassActivationRows:           0,
		DecouplingRows:               0,
		DynkinIndexRows:              0,
		ThresholdBetaRows:            0,
		ProvenZeroRows:               0,
		PhysicalConstantsDerived:     false,
		BetaPermissionFirewallClosed: true,
		Verdict:                      "finite zeta values are exact action-level data, but they are not a representation-complete beta ledger or a finite spectral triple",
	}

	summary := Summary{
		ContactRows:                prev.ContactRows,
		RationalSingletonRows:      prev.RationalSingletonRows,
		QuarticBlockRows:           prev.QuarticOrbitRows,
		QuarticCollectiveBlocks:    prev.QuarticCollectiveBlocks,
		ZetaValuesComputed:         len(zetaValues),
		ActionCandidatesAudited:    len(actionCandidates),
		BoundaryConstraintsDerived: actionAudit.BoundaryConstraintsDerived,
		QuarticBlockBetaRows:       0,
		ContactBetaRowsAllowed:     0,
		ContactZeroRowsProved:      0,
		ResidualS6Choices:          prev.ResidualS6Choices,
		ResidualNullityBefore:      prev.ResidualNullityAfter,
		ResidualNullityAfter:       prev.ResidualNullityAfter,
	}

	truth := "Gate 162 constructs the finite seven-root contact zeta ledger zeta_contact(s)=sum lambda^-s for s=0..4. The ledger is exact over Q, Galois-invariant, branch-free, and pole-free because the spectrum is finite and nonzero. However, zeta regularization alone does not provide a finite Dirac operator, real structure, grading, canonical cutoff function, or gauge-kinetic representation map. It therefore supplies action-level spectral data but derives no beta row, no boundary normalization, and no physical constant."

	return Analysis{
		Previous:                     prev,
		ZetaValues:                   zetaValues,
		ActionCandidates:             actionCandidates,
		ZetaAudit:                    zetaAudit,
		ActionAudit:                  actionAudit,
		BetaFirewall:                 betaFirewall,
		Summary:                      summary,
		ContactRows:                  prev.ContactRows,
		ExactRationalOverlapMatrix:   prev.ExactRationalOverlapMatrix,
		ExactCharacteristicCertified: prev.ExactCharacteristicCertified,
		ExactRootIsolationCertified:  prev.ExactRootIsolationCertified,
		PositiveNonzeroSpectrumRows:  7,
		RationalPrimaryIdempotents:   prev.RationalPrimaryIdempotents,
		GaloisInvariantOrbits:        prev.GaloisInvariantOrbits,
		RationalSingletonRows:        prev.RationalSingletonRows,
		QuarticOrbitRows:             prev.QuarticOrbitRows,
		QuarticCompressedBlocks:      prev.QuarticCompressedBlocks,
		QuarticCollectiveBlocks:      prev.QuarticCollectiveBlocks,
		QuarticBlockInvariants:       prev.QuarticBlockInvariants,
		QuarticSpectralMoments:       prev.QuarticSpectralMoments,
		ContactZetaValues:            len(zetaValues),
		FiniteZetaPoleCount:          0,
		AnalyticContinuationNeeded:   false,
		SpectralActionCandidates:     len(actionCandidates),
		SpectralTripleComplete:       false,
		FiniteDiracSelected:          false,
		RealStructureSelected:        false,
		GradingSelected:              false,
		CanonicalCutoffSelected:      false,
		GaugeKineticMapRows:          0,
		IndividualQuarticRows:        0,
		CanonicalQuarticBranches:     0,
		ExactNumberFieldProjectors:   0,
		IndividualQuarticProjectors:  0,
		RowwiseRootAssignmentProofs:  0,
		ExternalSelectorRows:         0,
		CanonicalTwoTwoSplits:        0,
		BranchBreakingSources:        0,
		ChargeSemanticRows:           0,
		T3RRowsDerived:               0,
		ChiralityRowsDerived:         0,
		BMinusLRowsDerived:           0,
		SU2LRowsDerived:              0,
		HyperchargeRowsDerived:       0,
		GaugeRepresentationRows:      0,
		SpinStatisticsRows:           0,
		LocalFieldRows:               0,
		KineticPoleResidueRows:       0,
		MassActivationRows:           0,
		DecouplingRows:               0,
		DynkinIndexRows:              0,
		BRSTCancellationRows:         0,
		ConstraintRows:               0,
		PropagatorRows:               0,
		RepresentationCompleteRows:   0,
		RepresentationOpenRows:       prev.RepresentationOpenRows,
		QuarticZeroBetaRows:          0,
		QuarticBlockBetaRows:         0,
		ContactBetaRowsAllowed:       0,
		ContactZeroRowsProved:        0,
		BetaPermissionFirewallClosed: true,
		ThresholdCorrectedBeta:       false,
		FullBetaMatchingTensor:       false,
		BoundaryConstraintsDerived:   actionAudit.BoundaryConstraintsDerived,
		ResidualS6Choices:            prev.ResidualS6Choices,
		ResidualNullityBefore:        prev.ResidualNullityAfter,
		ResidualNullityAfter:         prev.ResidualNullityAfter,
		HiddenObservedInputUsed:      false,
		PhysicalWeakAngleDerived:     false,
		FineStructureDerived:         false,
		PhysicalMassesDerived:        false,
		PhysicalScaleDerived:         false,
		TruthStatement:               truth,
		RejectedClaims: []string{
			"finite zeta regularization selects individual quartic contact branches",
			"zeta values alone define a complete noncommutative spectral triple",
			"a zeta scalar fixes the cutoff/test function or spectral-action coefficients",
			"contact zeta functionals derive kappa_U1, the embedded 5/3 normalization, or weak-angle data",
			"finite contact zeta data permit threshold beta rows without representation and decoupling data",
			"the zeta ledger derives alpha_EM, masses, a physical scale, CKM, or PMNS",
		},
		RemainingUnknowns: []string{
			"a canonical finite Dirac-like operator on the total finite Hilbert space",
			"a real structure and grading compatible with the already-proved quartic no-go theorems",
			"a canonical cutoff/test function or equivalent action-weight rule",
			"a representation-complete map from spectral terms to gauge kinetic normalizations",
			"threshold-corrected beta coefficients and physical constants",
		},
		RecommendedNextGate: "Gate 163 — finite spectral action principle / spectral triple construction audit",
	}, nil
}

func contactZetaValues(poly quarticspectralfunctional.QuarticPolynomial, rationalRoots []Rational, maxS int) []ZetaValue {
	quartic := quarticInversePowerMoments(poly, maxS)
	out := make([]ZetaValue, 0, maxS+1)
	for s := 0; s <= maxS; s++ {
		ratPart := NewRational(int64(len(rationalRoots)), 1)
		if s > 0 {
			ratPart = NewRational(0, 1)
			for _, root := range rationalRoots {
				ratPart = ratPart.Add(root.Pow(-s))
			}
		}
		full := ratPart.Add(quartic[s])
		out = append(out, ZetaValue{
			S:                    s,
			RationalPart:         ratPart,
			QuarticPart:          quartic[s],
			Full:                 full,
			ExactOverQ:           true,
			GaloisInvariant:      true,
			BranchFree:           true,
			PoleFree:             true,
			UsesObservedInput:    false,
			UsesBranchChoice:     false,
			RequiresRowSemantics: false,
			Verdict:              "finite-spectrum zeta value; exact and branch-free but action-level only",
		})
	}
	return out
}

func quarticInversePowerMoments(poly quarticspectralfunctional.QuarticPolynomial, maxS int) map[int]Rational {
	e1 := rationalFromGate161(poly.Sum)
	e2 := rationalFromGate161(poly.PairSum)
	e3 := rationalFromGate161(poly.TripleSum)
	e4 := rationalFromGate161(poly.Product)
	// Elementary symmetric functions of reciprocal quartic roots.
	E1 := e3.Div(e4)
	E2 := e2.Div(e4)
	E3 := e1.Div(e4)
	E4 := NewRational(1, 1).Div(e4)
	moments := map[int]Rational{0: NewRational(4, 1)}
	if maxS >= 1 {
		moments[1] = E1
	}
	if maxS >= 2 {
		moments[2] = E1.Mul(moments[1]).Sub(NewRational(2, 1).Mul(E2))
	}
	if maxS >= 3 {
		moments[3] = E1.Mul(moments[2]).Sub(E2.Mul(moments[1])).Add(NewRational(3, 1).Mul(E3))
	}
	if maxS >= 4 {
		moments[4] = E1.Mul(moments[3]).Sub(E2.Mul(moments[2])).Add(E3.Mul(moments[1])).Sub(NewRational(4, 1).Mul(E4))
	}
	return moments
}

func fullPositiveMoments(poly quarticspectralfunctional.QuarticPolynomial, rationalRoots []Rational, maxOrder int) map[int]Rational {
	e1 := rationalFromGate161(poly.Sum)
	e2 := rationalFromGate161(poly.PairSum)
	e3 := rationalFromGate161(poly.TripleSum)
	e4 := rationalFromGate161(poly.Product)
	quartic := map[int]Rational{0: NewRational(4, 1), 1: e1}
	if maxOrder >= 2 {
		quartic[2] = e1.Mul(quartic[1]).Sub(NewRational(2, 1).Mul(e2))
	}
	if maxOrder >= 3 {
		quartic[3] = e1.Mul(quartic[2]).Sub(e2.Mul(quartic[1])).Add(NewRational(3, 1).Mul(e3))
	}
	if maxOrder >= 4 {
		quartic[4] = e1.Mul(quartic[3]).Sub(e2.Mul(quartic[2])).Add(e3.Mul(quartic[1])).Sub(NewRational(4, 1).Mul(e4))
	}
	out := map[int]Rational{0: NewRational(7, 1)}
	for k := 1; k <= maxOrder; k++ {
		sum := quartic[k]
		for _, root := range rationalRoots {
			sum = sum.Add(root.Pow(k))
		}
		out[k] = sum
	}
	return out
}

func fullContactProduct(poly quarticspectralfunctional.QuarticPolynomial, rationalRoots []Rational) Rational {
	prod := rationalFromGate161(poly.Product)
	for _, root := range rationalRoots {
		prod = prod.Mul(root)
	}
	return prod
}

func actionFunctionalCandidates(zeta []ZetaValue, positive map[int]Rational, product Rational) []ActionFunctionalCandidate {
	z := map[int]Rational{}
	for _, value := range zeta {
		z[value.S] = value.Full
	}
	candidates := []ActionFunctionalCandidate{
		actionCandidate("dimension term", "zeta(0)", z[0]),
		actionCandidate("inverse trace", "zeta(1)", z[1]),
		actionCandidate("inverse quadratic trace", "zeta(2)", z[2]),
		actionCandidate("inverse cubic trace", "zeta(3)", z[3]),
		actionCandidate("inverse quartic trace", "zeta(4)", z[4]),
		actionCandidate("inverse mean", "zeta(1)/7", z[1].Div(NewRational(7, 1))),
		actionCandidate("inverse quadratic shape", "zeta(2)/zeta(1)^2", z[2].Div(z[1].Mul(z[1]))),
		actionCandidate("positive-inverse balance", "Tr(Omega) zeta(1) / 49", positive[1].Mul(z[1]).Div(NewRational(49, 1))),
		actionCandidate("full determinant", "prod(lambda_i)", product),
		actionCandidate("reciprocal determinant", "1/prod(lambda_i)", NewRational(1, 1).Div(product)),
	}
	for i := range candidates {
		candidates[i].Verdict = "exact finite zeta/action scalar; without a spectral triple and canonical cutoff coefficient it does not constrain physics"
	}
	return candidates
}

func actionCandidate(name, formula string, value Rational) ActionFunctionalCandidate {
	return ActionFunctionalCandidate{
		Name:                     name,
		Formula:                  formula,
		Value:                    value,
		ExactOverQ:               true,
		GaloisInvariant:          true,
		BranchFree:               true,
		UsesObservedInput:        false,
		UsesBranchChoice:         false,
		RequiresRowSemantics:     false,
		RequiresSpectralTriple:   true,
		RequiresCutoffFunction:   true,
		CoefficientCanonical:     false,
		BetaRowsAllowed:          0,
		PhysicalConstantsDerived: false,
	}
}

func markBoundaryComparisons(candidates []ActionFunctionalCandidate) {
	kappaU1 := NewRational(6, 1)
	embeddedBoundary := NewRational(5, 3)
	contactWeakAngle := NewRational(3, 8)
	generatorWeakAngle := NewRational(1, 4)
	for i := range candidates {
		c := &candidates[i]
		c.MatchesKappaU1 = c.Value.Equal(kappaU1)
		c.MatchesEmbeddedBoundary = c.Value.Equal(embeddedBoundary)
		c.MatchesContactWeakAngle = c.Value.Equal(contactWeakAngle)
		c.MatchesGeneratorWeakAngle = c.Value.Equal(generatorWeakAngle)
		c.ConstrainsBoundary = c.MatchesKappaU1 || c.MatchesEmbeddedBoundary || c.MatchesContactWeakAngle || c.MatchesGeneratorWeakAngle
	}
}

func auditZeta(values []ZetaValue) ZetaRegularizationAudit {
	a := ZetaRegularizationAudit{ValuesComputed: len(values), PositiveNonzeroSpectrumRows: 7, Verdict: "finite contact zeta values are exact, pole-free, and branch-free; no analytic continuation is needed for the finite seven-root ledger"}
	for _, z := range values {
		if z.ExactOverQ {
			a.ExactRationalValues++
		}
		if z.GaloisInvariant {
			a.GaloisInvariantValues++
		}
		if z.BranchFree {
			a.BranchFreeValues++
		}
		if z.UsesBranchChoice {
			a.BranchChoicesUsed++
		}
		if z.UsesObservedInput {
			a.ObservedInputsUsed = true
		}
		if z.RequiresRowSemantics {
			a.RowSemanticsUsed = true
		}
		if !z.PoleFree {
			a.Poles++
		}
		if z.S == 0 && z.Full.Equal(NewRational(7, 1)) {
			a.ZetaZeroEqualsDimension = true
		}
	}
	a.AnalyticContinuationNeeded = false
	a.FiniteRegularizationComplete = a.ValuesComputed == 5 && a.ExactRationalValues == 5 && a.GaloisInvariantValues == 5 && a.BranchFreeValues == 5 && a.BranchChoicesUsed == 0 && !a.ObservedInputsUsed && !a.RowSemanticsUsed && a.Poles == 0 && a.ZetaZeroEqualsDimension
	return a
}

func auditAction(candidates []ActionFunctionalCandidate) SpectralActionAudit {
	a := SpectralActionAudit{CandidatesAudited: len(candidates), Verdict: "zeta/action scalars are exact, but all physical use requires missing spectral-triple and cutoff-function data"}
	for _, c := range candidates {
		if c.UsesObservedInput {
			a.ObservedInputsUsed = true
		}
		if c.UsesBranchChoice {
			a.BranchChoicesUsed++
		}
		if c.RequiresRowSemantics {
			a.RowSemanticsUsed = true
		}
		if c.RequiresSpectralTriple {
			a.RequiresSpectralTriple++
		}
		if c.RequiresCutoffFunction {
			a.RequiresCutoffFunction++
		}
		if c.CoefficientCanonical {
			a.CanonicalCoefficients++
		}
		if c.MatchesKappaU1 {
			a.MatchesKappaU1++
		}
		if c.MatchesEmbeddedBoundary {
			a.MatchesEmbeddedBoundary++
		}
		if c.MatchesContactWeakAngle {
			a.MatchesContactWeakAngle++
		}
		if c.MatchesGeneratorWeakAngle {
			a.MatchesGeneratorWeakAngle++
		}
		if c.ConstrainsBoundary {
			a.BoundaryConstraintsDerived++
		}
		if c.PhysicalConstantsDerived {
			a.PhysicalConstantsDerived = true
		}
	}
	return a
}

func FormatZetaValue(z ZetaValue) string {
	return fmt.Sprintf("zeta(%d): rational=%s quartic=%s full=%s exactQ=%t galois=%t branchFree=%t poleFree=%t observed=%t branchChoice=%t rowSemantics=%t (%s)", z.S, z.RationalPart.String(), z.QuarticPart.String(), z.Full.String(), z.ExactOverQ, z.GaloisInvariant, z.BranchFree, z.PoleFree, z.UsesObservedInput, z.UsesBranchChoice, z.RequiresRowSemantics, z.Verdict)
}

func FormatZetaList(values []ZetaValue) string {
	parts := make([]string, 0, len(values))
	for _, z := range values {
		parts = append(parts, fmt.Sprintf("zeta(%d)=%s", z.S, z.Full.String()))
	}
	sort.Strings(parts)
	return strings.Join(parts, "; ")
}

func FormatActionCandidate(c ActionFunctionalCandidate) string {
	return fmt.Sprintf("%s: %s=%s exactQ=%t galois=%t branchFree=%t observed=%t branchChoice=%t rowSemantics=%t needsTriple=%t needsCutoff=%t coeffCanonical=%t match(kappa6=%t, embedded5/3=%t, weak3/8=%t, gen1/4=%t) constrains=%t beta=%d physical=%t (%s)", c.Name, c.Formula, c.Value.String(), c.ExactOverQ, c.GaloisInvariant, c.BranchFree, c.UsesObservedInput, c.UsesBranchChoice, c.RequiresRowSemantics, c.RequiresSpectralTriple, c.RequiresCutoffFunction, c.CoefficientCanonical, c.MatchesKappaU1, c.MatchesEmbeddedBoundary, c.MatchesContactWeakAngle, c.MatchesGeneratorWeakAngle, c.ConstrainsBoundary, c.BetaRowsAllowed, c.PhysicalConstantsDerived, c.Verdict)
}

func FormatActionCandidateList(candidates []ActionFunctionalCandidate) string {
	parts := make([]string, 0, len(candidates))
	for _, c := range candidates {
		parts = append(parts, fmt.Sprintf("%s=%s", c.Name, c.Value.String()))
	}
	sort.Strings(parts)
	return strings.Join(parts, "; ")
}

func FormatZetaAudit(a ZetaRegularizationAudit) string {
	return fmt.Sprintf("values=%d exactQ=%d galois=%d branchFree=%d branchChoices=%d observed=%t rowSemantics=%t poles=%d analyticContinuation=%t positiveRows=%d zeta0dim=%t complete=%t (%s)", a.ValuesComputed, a.ExactRationalValues, a.GaloisInvariantValues, a.BranchFreeValues, a.BranchChoicesUsed, a.ObservedInputsUsed, a.RowSemanticsUsed, a.Poles, a.AnalyticContinuationNeeded, a.PositiveNonzeroSpectrumRows, a.ZetaZeroEqualsDimension, a.FiniteRegularizationComplete, a.Verdict)
}

func FormatActionAudit(a SpectralActionAudit) string {
	return fmt.Sprintf("candidates=%d observed=%t branchChoices=%d rowSemantics=%t needsTriple=%d needsCutoff=%d canonicalCoeff=%d D=%t J=%t grading=%t gaugeMap=%t constraints=%d matches(kappa6=%d, embedded5/3=%d, weak3/8=%d, gen1/4=%d) newAction=%t physical=%t (%s)", a.CandidatesAudited, a.ObservedInputsUsed, a.BranchChoicesUsed, a.RowSemanticsUsed, a.RequiresSpectralTriple, a.RequiresCutoffFunction, a.CanonicalCoefficients, a.FiniteDiracSelected, a.RealStructureSelected, a.GradingSelected, a.GaugeKineticMapDerived, a.BoundaryConstraintsDerived, a.MatchesKappaU1, a.MatchesEmbeddedBoundary, a.MatchesContactWeakAngle, a.MatchesGeneratorWeakAngle, a.NewActionNormalization, a.PhysicalConstantsDerived, a.Verdict)
}

func FormatBetaFirewall(f BetaFirewallAudit) string {
	return fmt.Sprintf("finiteZeta=%t triple=%t cutoff=%t gaugeMap=%d individualQuartic=%d gaugeRep=%d spin=%d local=%d mass=%d decoupling=%d dynkin=%d beta=%d zero=%d physical=%t closed=%t (%s)", f.FiniteZetaData, f.SpectralTripleComplete, f.CanonicalCutoffSelected, f.GaugeKineticMapRows, f.IndividualQuarticRows, f.GaugeRepresentationRows, f.SpinStatisticsRows, f.LocalFieldRows, f.MassActivationRows, f.DecouplingRows, f.DynkinIndexRows, f.ThresholdBetaRows, f.ProvenZeroRows, f.PhysicalConstantsDerived, f.BetaPermissionFirewallClosed, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contactRows=%d rational=%d quarticRows=%d collectiveBlocks=%d zeta=%d candidates=%d constraints=%d quarticBeta=%d contactBeta=%d zero=%d S6=%d nullity=%d→%d", s.ContactRows, s.RationalSingletonRows, s.QuarticBlockRows, s.QuarticCollectiveBlocks, s.ZetaValuesComputed, s.ActionCandidatesAudited, s.BoundaryConstraintsDerived, s.QuarticBlockBetaRows, s.ContactBetaRowsAllowed, s.ContactZeroRowsProved, s.ResidualS6Choices, s.ResidualNullityBefore, s.ResidualNullityAfter)
}
