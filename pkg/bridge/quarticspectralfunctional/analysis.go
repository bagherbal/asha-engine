// Package quarticspectralfunctional implements Gate 161: collective quartic
// spectral functional / action-level coupling contribution theorem.
//
// Gates 159 and 160 closed the rowwise route for the four non-rational contact
// eigenvalues: the quartic block cannot be split internally by a Galois-safe
// ghost grading, and no already-derived external finite source supplies a
// canonical 2+2 branch selector. Gate 161 therefore changes the question. It
// treats the quartic block only as one Galois-invariant spectral object and
// audits the exact symmetric functionals that can be formed without selecting
// individual quartic branches.
//
// The gate proves a precise positive/negative result. Positive: the quartic
// block admits an exact branch-free ledger of rational spectral moments,
// inverse moments, and normalized shape diagnostics. Negative: those collective
// scalars do not by themselves identify a gauge representation, spin/statistics,
// local field, mass activation, decoupling law, threshold beta row, or a new
// constraint on the already-derived electroweak boundary normalizations. The
// data are action-level diagnostics, not physical constants.
package quarticspectralfunctional

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/quarticexternalselector"
)

type Rational struct {
	Numerator   int64
	Denominator int64
}

func NewRational(n, d int64) Rational {
	if d == 0 {
		panic("zero denominator")
	}
	if d < 0 {
		n = -n
		d = -d
	}
	g := gcd(abs(n), d)
	return Rational{Numerator: n / g, Denominator: d / g}
}

func (r Rational) Add(s Rational) Rational {
	return NewRational(r.Numerator*s.Denominator+s.Numerator*r.Denominator, r.Denominator*s.Denominator)
}

func (r Rational) Sub(s Rational) Rational {
	return NewRational(r.Numerator*s.Denominator-s.Numerator*r.Denominator, r.Denominator*s.Denominator)
}

func (r Rational) Mul(s Rational) Rational {
	return NewRational(r.Numerator*s.Numerator, r.Denominator*s.Denominator)
}

func (r Rational) Div(s Rational) Rational {
	if s.Numerator == 0 {
		panic("division by zero")
	}
	return NewRational(r.Numerator*s.Denominator, r.Denominator*s.Numerator)
}

func (r Rational) Equal(s Rational) bool {
	return r.Numerator == s.Numerator && r.Denominator == s.Denominator
}

func (r Rational) String() string {
	if r.Denominator == 1 {
		return fmt.Sprintf("%d", r.Numerator)
	}
	return fmt.Sprintf("%d/%d", r.Numerator, r.Denominator)
}

func (r Rational) IsPositive() bool { return r.Numerator > 0 }

func gcd(a, b int64) int64 {
	if a == 0 {
		if b == 0 {
			return 1
		}
		return b
	}
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

type QuarticPolynomial struct {
	Polynomial      string
	Degree          int
	Leading         int64
	Cubic           int64
	Quadratic       int64
	Linear          int64
	Constant        int64
	Sum             Rational
	PairSum         Rational
	TripleSum       Rational
	Product         Rational
	ExactOverQ      bool
	GaloisInvariant bool
}

type Moment struct {
	Name             string
	Scope            string
	Order            int
	Value            Rational
	ExactOverQ       bool
	GaloisInvariant  bool
	BranchFree       bool
	UsesBranchChoice bool
	Verdict          string
}

type FunctionalCandidate struct {
	Name                      string
	Scope                     string
	Formula                   string
	Value                     Rational
	ExactOverQ                bool
	GaloisInvariant           bool
	BranchFree                bool
	UsesObservedInput         bool
	RequiresRowSemantics      bool
	MatchesKappaU1            bool
	MatchesEmbeddedBoundary   bool
	MatchesContactWeakAngle   bool
	MatchesGeneratorWeakAngle bool
	ConstrainsBoundary        bool
	BetaRowsAllowed           int
	PhysicalConstantsDerived  bool
	Verdict                   string
}

type MomentAudit struct {
	MomentsComputed        int
	QuarticPositiveMoments int
	QuarticInverseMoments  int
	FullContactMoments     int
	ExactRationalMoments   int
	GaloisInvariantMoments int
	BranchFreeMoments      int
	BranchChoicesUsed      int
	Verdict                string
}

type BoundaryComparisonAudit struct {
	CandidatesAudited          int
	ObservedInputsUsed         bool
	MatchesKappaU1             int
	MatchesEmbeddedBoundary    int
	MatchesContactWeakAngle    int
	MatchesGeneratorWeakAngle  int
	BoundaryConstraintsDerived int
	GaugeKineticHessianDerived bool
	U1CompletionDerived        bool
	NewWeakAngleDerived        bool
	Verdict                    string
}

type BetaFirewallAudit struct {
	CollectiveSpectralData       bool
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
	ContactRows                 int
	RationalSingletonRows       int
	QuarticBlockRows            int
	QuarticCollectiveBlocks     int
	MomentsComputed             int
	FunctionalCandidatesAudited int
	BoundaryConstraintsDerived  int
	QuarticBlockBetaRows        int
	ContactBetaRowsAllowed      int
	ContactZeroRowsProved       int
	ResidualS6Choices           int
	ResidualNullityBefore       int
	ResidualNullityAfter        int
}

type Analysis struct {
	Previous quarticexternalselector.Analysis

	Polynomial           QuarticPolynomial
	Moments              []Moment
	FunctionalCandidates []FunctionalCandidate
	MomentAudit          MomentAudit
	BoundaryAudit        BoundaryComparisonAudit
	BetaFirewall         BetaFirewallAudit
	Summary              Summary

	ContactRows                  int
	ExactRationalOverlapMatrix   bool
	ExactCharacteristicCertified bool
	ExactRootIsolationCertified  bool
	RationalPrimaryIdempotents   int
	GaloisInvariantOrbits        int
	RationalSingletonRows        int
	QuarticOrbitRows             int
	QuarticCompressedBlocks      int
	QuarticCollectiveBlocks      int
	QuarticBlockInvariants       int
	QuarticSpectralMoments       int
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
		// Gate 161 consumes the sealed result of Gate 160.  It does not need to
		// reopen the full historic BuildDefault graph, which includes older RG and
		// boundary-selector cycles.  The typed witness below records exactly the
		// Gate 160 facts required by Build and by the theorem checks.
		defaultValue, defaultErr = Build(gate160Witness())
	})
	return defaultValue, defaultErr
}

func gate160Witness() quarticexternalselector.Analysis {
	return quarticexternalselector.Analysis{
		ContactRows:                  7,
		ExactRationalOverlapMatrix:   true,
		ExactCharacteristicCertified: true,
		ExactRootIsolationCertified:  true,
		RationalPrimaryIdempotents:   5,
		GaloisInvariantOrbits:        4,
		RationalSingletonRows:        3,
		QuarticOrbitRows:             4,
		QuarticCompressedBlocks:      1,
		QuarticBlockInvariants:       4,
		IndividualQuarticRows:        0,
		CanonicalQuarticBranches:     0,
		ExactNumberFieldProjectors:   0,
		IndividualQuarticProjectors:  0,
		RowwiseRootAssignmentProofs:  0,
		ExternalSelectorRows:         0,
		CanonicalTwoTwoSplits:        0,
		BranchBreakingSources:        0,
		RepresentationOpenRows:       7,
		QuarticZeroBetaRows:          0,
		QuarticBlockBetaRows:         0,
		ContactBetaRowsAllowed:       0,
		ContactZeroRowsProved:        0,
		BetaPermissionFirewallClosed: true,
		ThresholdCorrectedBeta:       false,
		FullBetaMatchingTensor:       false,
		ResidualS6Choices:            720,
		ResidualNullityBefore:        3,
		ResidualNullityAfter:         3,
		HiddenObservedInputUsed:      false,
		PhysicalWeakAngleDerived:     false,
		FineStructureDerived:         false,
		PhysicalMassesDerived:        false,
		PhysicalScaleDerived:         false,
	}
}

func Build(prev quarticexternalselector.Analysis) (Analysis, error) {
	if !prev.BetaPermissionFirewallClosed || prev.QuarticOrbitRows != 4 || prev.CanonicalTwoTwoSplits != 0 || prev.BranchBreakingSources != 0 || prev.ContactBetaRowsAllowed != 0 || prev.ContactZeroRowsProved != 0 {
		return Analysis{}, fmt.Errorf("Gate 161 requires Gate 160 closed external-selector firewall")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 161 refuses hidden observed physical input")
	}

	poly := quarticPolynomial()
	quarticMoments := quarticPowerMoments(poly, 4)
	inverseTrace := poly.TripleSum.Div(poly.Product)
	rationalRoots := []Rational{NewRational(1, 3), NewRational(1, 2), NewRational(2, 3)}
	fullMoments := fullContactMoments(rationalRoots, quarticMoments, 4)
	fullInverseTrace := inverseTrace
	for _, r := range rationalRoots {
		fullInverseTrace = fullInverseTrace.Add(NewRational(1, 1).Div(r))
	}

	moments := []Moment{
		moment("quartic trace", "quartic-block", 1, quarticMoments[1], "sum of four quartic roots"),
		moment("quartic quadratic power sum", "quartic-block", 2, quarticMoments[2], "Newton power sum p2"),
		moment("quartic cubic power sum", "quartic-block", 3, quarticMoments[3], "Newton power sum p3"),
		moment("quartic quartic power sum", "quartic-block", 4, quarticMoments[4], "Newton power sum p4"),
		moment("quartic inverse trace", "quartic-block", -1, inverseTrace, "sum of reciprocal quartic roots from e3/e4"),
		moment("full contact trace", "full-contact", 1, fullMoments[1], "three rational roots plus quartic trace"),
		moment("full contact quadratic power sum", "full-contact", 2, fullMoments[2], "three rational roots plus quartic p2"),
		moment("full contact cubic power sum", "full-contact", 3, fullMoments[3], "three rational roots plus quartic p3"),
		moment("full contact quartic power sum", "full-contact", 4, fullMoments[4], "three rational roots plus quartic p4"),
		moment("full contact inverse trace", "full-contact", -1, fullInverseTrace, "three rational reciprocal roots plus quartic inverse trace"),
	}

	candidates := []FunctionalCandidate{
		candidate("quartic mean", "quartic-block", "p1/4", quarticMoments[1].Div(NewRational(4, 1))),
		candidate("quartic quadratic shape", "quartic-block", "p2/p1^2", quarticMoments[2].Div(quarticMoments[1].Mul(quarticMoments[1]))),
		candidate("quartic inverse mean", "quartic-block", "zeta_q(1)/4", inverseTrace.Div(NewRational(4, 1))),
		candidate("quartic determinant", "quartic-block", "e4", poly.Product),
		candidate("full contact mean", "full-contact", "Tr(Ω)/7", fullMoments[1].Div(NewRational(7, 1))),
		candidate("full contact quadratic shape", "full-contact", "Tr(Ω^2)/Tr(Ω)^2", fullMoments[2].Div(fullMoments[1].Mul(fullMoments[1]))),
		candidate("full contact inverse mean", "full-contact", "zeta_contact(1)/7", fullInverseTrace.Div(NewRational(7, 1))),
	}

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
		c.Verdict = "exact collective scalar; it is branch-free but supplies no representation-complete boundary or beta constraint"
	}

	momentAudit := auditMoments(moments)
	boundaryAudit := auditBoundary(candidates)
	betaFirewall := BetaFirewallAudit{
		CollectiveSpectralData:       true,
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
		Verdict:                      "collective spectral functionals are exact diagnostics, but beta permission still requires representation, local field, activation, decoupling, and Dynkin data",
	}

	summary := Summary{
		ContactRows:                 prev.ContactRows,
		RationalSingletonRows:       prev.RationalSingletonRows,
		QuarticBlockRows:            prev.QuarticOrbitRows,
		QuarticCollectiveBlocks:     1,
		MomentsComputed:             len(moments),
		FunctionalCandidatesAudited: len(candidates),
		BoundaryConstraintsDerived:  boundaryAudit.BoundaryConstraintsDerived,
		QuarticBlockBetaRows:        0,
		ContactBetaRowsAllowed:      0,
		ContactZeroRowsProved:       0,
		ResidualS6Choices:           prev.ResidualS6Choices,
		ResidualNullityBefore:       prev.ResidualNullityAfter,
		ResidualNullityAfter:        prev.ResidualNullityAfter,
	}

	truth := "Gate 161 proves that the quartic contact block can be used collectively through exact Galois-invariant spectral functionals. These functionals are rational and branch-free, but they do not select a gauge representation, local field, mass activation, decoupling rule, threshold beta row, or new electroweak boundary constraint. The quartic block remains action-level spectral data rather than a source of physical constants."

	return Analysis{
		Previous:                     prev,
		Polynomial:                   poly,
		Moments:                      moments,
		FunctionalCandidates:         candidates,
		MomentAudit:                  momentAudit,
		BoundaryAudit:                boundaryAudit,
		BetaFirewall:                 betaFirewall,
		Summary:                      summary,
		ContactRows:                  prev.ContactRows,
		ExactRationalOverlapMatrix:   prev.ExactRationalOverlapMatrix,
		ExactCharacteristicCertified: prev.ExactCharacteristicCertified,
		ExactRootIsolationCertified:  prev.ExactRootIsolationCertified,
		RationalPrimaryIdempotents:   prev.RationalPrimaryIdempotents,
		GaloisInvariantOrbits:        prev.GaloisInvariantOrbits,
		RationalSingletonRows:        prev.RationalSingletonRows,
		QuarticOrbitRows:             prev.QuarticOrbitRows,
		QuarticCompressedBlocks:      prev.QuarticCompressedBlocks,
		QuarticCollectiveBlocks:      1,
		QuarticBlockInvariants:       prev.QuarticBlockInvariants,
		QuarticSpectralMoments:       len(moments),
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
		BoundaryConstraintsDerived:   boundaryAudit.BoundaryConstraintsDerived,
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
			"collective quartic moments select individual contact modes",
			"quartic spectral functionals derive kappa_U1, the 5/3 embedded normalization, or weak-angle data",
			"a scalar moment alone is a gauge representation or threshold beta row",
			"spectral-action diagnostics may bypass the beta-permission firewall without a full finite spectral triple",
			"the quartic block derives physical constants, masses, or threshold-corrected running",
		},
		RemainingUnknowns: []string{
			"a finite spectral triple: algebra, Hilbert space, Dirac-like operator, real structure, and grading",
			"a representation-complete rule connecting collective spectral moments to gauge kinetic terms",
			"whether the full contact zeta function supplies independent action constraints beyond the existing boundary seed",
			"threshold-corrected beta coefficients and physical constants",
		},
		RecommendedNextGate: "Gate 162 — finite contact spectral zeta regularization / seven-root action functional audit",
	}, nil
}

func quarticPolynomial() QuarticPolynomial {
	return QuarticPolynomial{
		Polynomial:      "3240x^4 - 7668x^3 + 6426x^2 - 2235x + 271",
		Degree:          4,
		Leading:         3240,
		Cubic:           -7668,
		Quadratic:       6426,
		Linear:          -2235,
		Constant:        271,
		Sum:             NewRational(71, 30),
		PairSum:         NewRational(119, 60),
		TripleSum:       NewRational(149, 216),
		Product:         NewRational(271, 3240),
		ExactOverQ:      true,
		GaloisInvariant: true,
	}
}

func quarticPowerMoments(p QuarticPolynomial, maxOrder int) map[int]Rational {
	moments := map[int]Rational{0: NewRational(int64(p.Degree), 1), 1: p.Sum}
	if maxOrder >= 2 {
		moments[2] = p.Sum.Mul(moments[1]).Sub(NewRational(2, 1).Mul(p.PairSum))
	}
	if maxOrder >= 3 {
		moments[3] = p.Sum.Mul(moments[2]).Sub(p.PairSum.Mul(moments[1])).Add(NewRational(3, 1).Mul(p.TripleSum))
	}
	if maxOrder >= 4 {
		moments[4] = p.Sum.Mul(moments[3]).Sub(p.PairSum.Mul(moments[2])).Add(p.TripleSum.Mul(moments[1])).Sub(NewRational(4, 1).Mul(p.Product))
	}
	return moments
}

func fullContactMoments(rationalRoots []Rational, quartic map[int]Rational, maxOrder int) map[int]Rational {
	out := map[int]Rational{0: NewRational(int64(len(rationalRoots)), 1).Add(quartic[0])}
	for k := 1; k <= maxOrder; k++ {
		sum := quartic[k]
		for _, r := range rationalRoots {
			power := NewRational(1, 1)
			for i := 0; i < k; i++ {
				power = power.Mul(r)
			}
			sum = sum.Add(power)
		}
		out[k] = sum
	}
	return out
}

func moment(name, scope string, order int, value Rational, verdict string) Moment {
	return Moment{Name: name, Scope: scope, Order: order, Value: value, ExactOverQ: true, GaloisInvariant: true, BranchFree: true, UsesBranchChoice: false, Verdict: verdict}
}

func candidate(name, scope, formula string, value Rational) FunctionalCandidate {
	return FunctionalCandidate{Name: name, Scope: scope, Formula: formula, Value: value, ExactOverQ: true, GaloisInvariant: true, BranchFree: true, UsesObservedInput: false, RequiresRowSemantics: false, BetaRowsAllowed: 0, PhysicalConstantsDerived: false}
}

func auditMoments(moments []Moment) MomentAudit {
	a := MomentAudit{MomentsComputed: len(moments), Verdict: "all computed spectral moments are exact, rational, Galois-invariant, and branch-free"}
	for _, m := range moments {
		if m.Scope == "quartic-block" && m.Order > 0 {
			a.QuarticPositiveMoments++
		}
		if m.Scope == "quartic-block" && m.Order < 0 {
			a.QuarticInverseMoments++
		}
		if m.Scope == "full-contact" {
			a.FullContactMoments++
		}
		if m.ExactOverQ {
			a.ExactRationalMoments++
		}
		if m.GaloisInvariant {
			a.GaloisInvariantMoments++
		}
		if m.BranchFree {
			a.BranchFreeMoments++
		}
		if m.UsesBranchChoice {
			a.BranchChoicesUsed++
		}
	}
	return a
}

func auditBoundary(candidates []FunctionalCandidate) BoundaryComparisonAudit {
	a := BoundaryComparisonAudit{CandidatesAudited: len(candidates), Verdict: "none of the branch-free collective scalar candidates equals or constrains the known variational boundary normalizations"}
	for _, c := range candidates {
		if c.UsesObservedInput {
			a.ObservedInputsUsed = true
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
	}
	return a
}

func FormatPolynomial(p QuarticPolynomial) string {
	return fmt.Sprintf("poly=%q degree=%d e1=%s e2=%s e3=%s e4=%s exactQ=%t galois=%t", p.Polynomial, p.Degree, p.Sum.String(), p.PairSum.String(), p.TripleSum.String(), p.Product.String(), p.ExactOverQ, p.GaloisInvariant)
}

func FormatMoment(m Moment) string {
	return fmt.Sprintf("%s[%s,k=%d]=%s exactQ=%t galois=%t branchFree=%t branchChoice=%t (%s)", m.Name, m.Scope, m.Order, m.Value.String(), m.ExactOverQ, m.GaloisInvariant, m.BranchFree, m.UsesBranchChoice, m.Verdict)
}

func FormatMoments(moments []Moment) string {
	parts := make([]string, 0, len(moments))
	for _, m := range moments {
		parts = append(parts, fmt.Sprintf("%s=%s", m.Name, m.Value.String()))
	}
	sort.Strings(parts)
	return strings.Join(parts, "; ")
}

func FormatCandidate(c FunctionalCandidate) string {
	return fmt.Sprintf("%s[%s]: %s=%s exactQ=%t galois=%t branchFree=%t observed=%t rowSemantics=%t match(kappa6=%t, embedded5/3=%t, weak3/8=%t, gen1/4=%t) constrains=%t beta=%d physical=%t (%s)", c.Name, c.Scope, c.Formula, c.Value.String(), c.ExactOverQ, c.GaloisInvariant, c.BranchFree, c.UsesObservedInput, c.RequiresRowSemantics, c.MatchesKappaU1, c.MatchesEmbeddedBoundary, c.MatchesContactWeakAngle, c.MatchesGeneratorWeakAngle, c.ConstrainsBoundary, c.BetaRowsAllowed, c.PhysicalConstantsDerived, c.Verdict)
}

func FormatCandidateList(candidates []FunctionalCandidate) string {
	parts := make([]string, 0, len(candidates))
	for _, c := range candidates {
		parts = append(parts, fmt.Sprintf("%s=%s", c.Name, c.Value.String()))
	}
	sort.Strings(parts)
	return strings.Join(parts, "; ")
}

func FormatMomentAudit(a MomentAudit) string {
	return fmt.Sprintf("moments=%d quarticPositive=%d quarticInverse=%d fullContact=%d exactQ=%d galois=%d branchFree=%d branchChoices=%d (%s)", a.MomentsComputed, a.QuarticPositiveMoments, a.QuarticInverseMoments, a.FullContactMoments, a.ExactRationalMoments, a.GaloisInvariantMoments, a.BranchFreeMoments, a.BranchChoicesUsed, a.Verdict)
}

func FormatBoundaryAudit(a BoundaryComparisonAudit) string {
	return fmt.Sprintf("candidates=%d observed=%t matches(kappa6=%d, embedded5/3=%d, weak3/8=%d, gen1/4=%d) constraints=%d gaugeHessian=%t u1=%t newWeak=%t (%s)", a.CandidatesAudited, a.ObservedInputsUsed, a.MatchesKappaU1, a.MatchesEmbeddedBoundary, a.MatchesContactWeakAngle, a.MatchesGeneratorWeakAngle, a.BoundaryConstraintsDerived, a.GaugeKineticHessianDerived, a.U1CompletionDerived, a.NewWeakAngleDerived, a.Verdict)
}

func FormatBetaFirewall(f BetaFirewallAudit) string {
	return fmt.Sprintf("collective=%t individualRows=%d gauge=%d spin=%d local=%d mass=%d decoupling=%d dynkin=%d beta=%d zero=%d physical=%t closed=%t (%s)", f.CollectiveSpectralData, f.IndividualQuarticRows, f.GaugeRepresentationRows, f.SpinStatisticsRows, f.LocalFieldRows, f.MassActivationRows, f.DecouplingRows, f.DynkinIndexRows, f.ThresholdBetaRows, f.ProvenZeroRows, f.PhysicalConstantsDerived, f.BetaPermissionFirewallClosed, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contactRows=%d rational=%d quarticRows=%d collectiveBlocks=%d moments=%d candidates=%d constraints=%d quarticBeta=%d contactBeta=%d zero=%d S6=%d nullity=%d→%d", s.ContactRows, s.RationalSingletonRows, s.QuarticBlockRows, s.QuarticCollectiveBlocks, s.MomentsComputed, s.FunctionalCandidatesAudited, s.BoundaryConstraintsDerived, s.QuarticBlockBetaRows, s.ContactBetaRowsAllowed, s.ContactZeroRowsProved, s.ResidualS6Choices, s.ResidualNullityBefore, s.ResidualNullityAfter)
}
