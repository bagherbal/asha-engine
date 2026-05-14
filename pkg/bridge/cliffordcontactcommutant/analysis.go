// Package cliffordcontactcommutant implements Gate 184: Clifford-contact
// spectral idempotent / commutant obstruction or construction.
//
// Gate 183 found canonical pre-actions from K7 to H_Fock and an abstract
// quartic scalar module, but no multiplicative action of the commutative
// contact spectral algebra C[Ω_contact] on the physical carriers. Gate 184
// sharpens the question by auditing three highly constrained algebraic routes:
//
//  1. a unital faithful seven-idempotent action of C^7 on the 16D Fock space;
//  2. an embedding through a maximal commutative Cartan/commutant of the
//     Clifford-spinor action;
//  3. the 4D quartic-primary ideal as a rank-one scalar module candidate.
//
// The result is a precise Fock no-go and a real but still abstract quartic
// scalar escape hatch. No arbitrary representation map is used.
package cliffordcontactcommutant

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactmoduleaction"
)

type Route string

const (
	FockRankRoute      Route = "seven-point-fock-rank-obstruction"
	CartanRoute        Route = "clifford-cartan-commutant-search"
	QuarticScalarRoute Route = "quartic-scalar-rank-one-module"
)

type FockRankObstructionAudit struct {
	Route                                   Route
	ContactPointCount                       int
	FockDimension                           int
	UnitalFaithfulActionRequiresIdempotents int
	UniformRankRequiredByTransitiveSymmetry bool
	UniformRankInteger                      bool
	RemainderModuloPoints                   int
	ExampleNearUniformRanks                 []int
	NonUniformRanksExist                    bool
	NonUniformRanksRequireSelector          bool
	CanonicalContactPointSelectorAvailable  bool
	CliffordVectorActionAvailable           bool
	CliffordActionMultiplicativeForC7       bool
	FaithfulMultiplicativeFockActionDerived bool
	Verdict                                 string
}

type CartanCommutantAudit struct {
	Route                               Route
	SpinorDimension                     int
	CliffordOddGeneratorCount           int
	MaximalCommutingGeneratorRank       int
	PrimitiveCartanIdempotents          int
	ContactPointCount                   int
	DimensionalEmbeddingPossible        bool
	RequiresChoiceOfCartan              bool
	CanonicalCartanSelectorDerived      bool
	RequiresDeleteOrMergeOneCartanPoint bool
	ContactSpectralOrderPreserved       bool
	EmbeddingIntoCommutantDerived       bool
	Verdict                             string
}

type QuarticScalarModuleAudit struct {
	Route                                     Route
	QuarticPrimaryDim                         int
	ScalarCarrierDim                          int
	QuarticPolynomial                         string
	RankOneDimensionMatch                     bool
	IntegerRankObstruction                    bool
	GaloisSafePrimaryIdeal                    bool
	AbstractRegularModuleDerived              bool
	CompanionAlgebraActionAvailable           bool
	ActsOnPhysicalHphi                        bool
	CanonicalHphiBasisOrOperatorDerived       bool
	ScalarOperatorHasQuarticMinimalPolynomial bool
	PhysicalScalarBundleDerived               bool
	ChernWeilReady                            bool
	Verdict                                   string
}

type Candidate struct {
	Name                string
	Route               Route
	Domain              string
	Target              string
	DomainDim           int
	TargetDim           int
	CanonicalPredata    bool
	DimensionCompatible bool
	RequiresSelector    bool
	AlgebraHomomorphism bool
	PhysicalAction      bool
	Verdict             string
}

type Summary struct {
	CandidatesAudited             int
	DimensionCompatibleCandidates int
	SelectorBlockedCandidates     int
	AlgebraHomomorphisms          int
	PhysicalFockActions           int
	AbstractQuarticModules        int
	PhysicalScalarActions         int
	CompletePhysicalBundleMaps    int
	Comment                       string
}

type Firewall struct {
	UsesObservedInputForDerivation     bool
	ArbitraryLinearMapUsed             bool
	ContactBaseInherited               bool
	CliffordPreactionInherited         bool
	FockRankObstructionProved          bool
	CartanCommutantObstructionProved   bool
	QuarticScalarAbstractModuleDerived bool
	CanonicalFockActionDerived         bool
	CanonicalScalarActionDerived       bool
	PhysicalBundleMapDerived           bool
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
	PreviousGate183 contactmoduleaction.Analysis
	FockRank        FockRankObstructionAudit
	Cartan          CartanCommutantAudit
	QuarticScalar   QuarticScalarModuleAudit
	Candidates      []Candidate
	Summary         Summary
	Firewall        Firewall
	TruthStatement  string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		defaultA, defaultErr = buildDefault()
	})
	return defaultA, defaultErr
}

func buildDefault() (Analysis, error) {
	prev, err := contactmoduleaction.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate 183 input: %w", err)
	}

	pointCount := 7
	fockDim := 16
	nearUniform := []int{3, 3, 2, 2, 2, 2, 2}
	fock := FockRankObstructionAudit{
		Route:                                   FockRankRoute,
		ContactPointCount:                       pointCount,
		FockDimension:                           fockDim,
		UnitalFaithfulActionRequiresIdempotents: pointCount,
		UniformRankRequiredByTransitiveSymmetry: true,
		UniformRankInteger:                      fockDim%pointCount == 0,
		RemainderModuloPoints:                   fockDim % pointCount,
		ExampleNearUniformRanks:                 nearUniform,
		NonUniformRanksExist:                    sum(nearUniform) == fockDim,
		NonUniformRanksRequireSelector:          true,
		CanonicalContactPointSelectorAvailable:  false,
		CliffordVectorActionAvailable:           prev.CliffordSpinor.LinearK7ToEndFockMapDerived,
		CliffordActionMultiplicativeForC7:       false,
		FaithfulMultiplicativeFockActionDerived: false,
		Verdict:                                 "A faithful unital C^7 action on H_Fock would be seven orthogonal spectral idempotents with integer ranks summing to 16. Contact-mode symmetry would require equal ranks, but 16 mod 7 = 2. Non-uniform ranks such as 3,3,2,2,2,2,2 exist only after choosing which contact points get larger fibers, and no canonical contact-point selector is available.",
	}

	cartan := CartanCommutantAudit{
		Route:                               CartanRoute,
		SpinorDimension:                     fockDim,
		CliffordOddGeneratorCount:           7,
		MaximalCommutingGeneratorRank:       3,
		PrimitiveCartanIdempotents:          8,
		ContactPointCount:                   pointCount,
		DimensionalEmbeddingPossible:        true,
		RequiresChoiceOfCartan:              true,
		CanonicalCartanSelectorDerived:      false,
		RequiresDeleteOrMergeOneCartanPoint: true,
		ContactSpectralOrderPreserved:       false,
		EmbeddingIntoCommutantDerived:       false,
		Verdict:                             "A Clifford Cartan/commutant has eight primitive idempotents, so seven contact points can fit only after choosing a Cartan and then deleting, merging, or otherwise selecting one of eight primitive cells. The engine has no canonical Cartan gauge selector or seven-of-eight embedding compatible with the contact spectrum.",
	}

	quartic := QuarticScalarModuleAudit{
		Route:                                     QuarticScalarRoute,
		QuarticPrimaryDim:                         4,
		ScalarCarrierDim:                          4,
		QuarticPolynomial:                         "3240x^4 - 7668x^3 + 6426x^2 - 2235x + 271",
		RankOneDimensionMatch:                     true,
		IntegerRankObstruction:                    false,
		GaloisSafePrimaryIdeal:                    prev.QuarticScalar.GaloisSafePrimaryIdeal,
		AbstractRegularModuleDerived:              prev.QuarticScalar.AbstractRankOneModuleOverQuartic,
		CompanionAlgebraActionAvailable:           prev.QuarticScalar.CompanionRepresentationAvailable,
		ActsOnPhysicalHphi:                        false,
		CanonicalHphiBasisOrOperatorDerived:       false,
		ScalarOperatorHasQuarticMinimalPolynomial: false,
		PhysicalScalarBundleDerived:               false,
		ChernWeilReady:                            false,
		Verdict:                                   "The quartic primary ideal has exactly the right 4D rank-one regular-module size for H_Φ, so the 16 mod 7 obstruction vanishes. This is a genuine algebraic escape hatch, but it remains abstract: no canonical scalar operator/basis on physical H_Φ has been derived whose minimal polynomial is the quartic contact factor.",
	}

	candidates := buildCandidates(fock, cartan, quartic)
	summary := auditCandidates(candidates)
	firewall := Firewall{
		UsesObservedInputForDerivation:     false,
		ArbitraryLinearMapUsed:             false,
		ContactBaseInherited:               prev.Firewall.ContactBaseInherited,
		CliffordPreactionInherited:         prev.Firewall.CliffordSpinorPreactionDerived,
		FockRankObstructionProved:          !fock.FaithfulMultiplicativeFockActionDerived && !fock.UniformRankInteger && fock.NonUniformRanksRequireSelector,
		CartanCommutantObstructionProved:   !cartan.EmbeddingIntoCommutantDerived && cartan.RequiresChoiceOfCartan && !cartan.CanonicalCartanSelectorDerived,
		QuarticScalarAbstractModuleDerived: quartic.AbstractRegularModuleDerived && quartic.RankOneDimensionMatch,
		CanonicalFockActionDerived:         false,
		CanonicalScalarActionDerived:       false,
		PhysicalBundleMapDerived:           false,
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
			"faithful multiplicative C^7 action on 16D H_Fock is blocked by 16 mod 7 plus missing contact-point selector",
			"Clifford vector multiplication is not a commutative spectral-idempotent action",
			"Cartan/commutant embedding requires a noncanonical Cartan and seven-of-eight cell choice",
			"quartic scalar primary ideal is a branch-free 4D abstract rank-one module candidate",
		},
		OpenRequirements: []string{
			"canonical scalar operator on H_Φ with the quartic contact minimal polynomial",
			"canonical identification H_Φ ≅ Q[x]/q4(x) as a physical scalar module",
			"compatibility of the quartic scalar module with SU(2)_L × U(1)_Y scalar action and real/chiral structures",
			"finite integration/Chern-character pairing for any promoted scalar bundle",
		},
		RecommendedNextGate: "Gate 185 — quartic scalar operator/minimal-polynomial construction on H_Φ",
		Verdict:             "Fock contact-idempotent action and Clifford Cartan embedding are obstructed; the quartic scalar 4→4 route remains the only viable finite physical-bundle target, but is not yet promoted to H_Φ.",
	}
	truth := "Gate 184 proves the 7-point contact base cannot act canonically and faithfully on the 16D Fock space by spectral idempotents: equal ranks are impossible because 16 mod 7 = 2, while non-uniform ranks require a forbidden selector for contact points. The Clifford Cartan/commutant route is also obstructed by the need to choose a Cartan and a seven-of-eight idempotent embedding. The only surviving route is the quartic scalar escape hatch: the 4D quartic primary ideal has an abstract rank-one module of exactly the scalar dimension. This is a real algebraic target, but not yet a physical H_Φ bundle until a canonical quartic-minimal scalar operator or equivalent identification is derived."
	return Analysis{PreviousGate183: prev, FockRank: fock, Cartan: cartan, QuarticScalar: quartic, Candidates: candidates, Summary: summary, Firewall: firewall, TruthStatement: truth}, nil
}

func buildCandidates(f FockRankObstructionAudit, c CartanCommutantAudit, q QuarticScalarModuleAudit) []Candidate {
	return []Candidate{
		{Name: "faithful seven-idempotent action on H_Fock", Route: f.Route, Domain: "C[Ω] ≅ C^7", Target: "H_Fock", DomainDim: f.ContactPointCount, TargetDim: f.FockDimension, CanonicalPredata: true, DimensionCompatible: false, RequiresSelector: true, AlgebraHomomorphism: false, PhysicalAction: false, Verdict: f.Verdict},
		{Name: "non-uniform seven-fiber rank pattern on H_Fock", Route: f.Route, Domain: "C^7 idempotents", Target: "rank pattern 3,3,2,2,2,2,2", DomainDim: f.ContactPointCount, TargetDim: f.FockDimension, CanonicalPredata: true, DimensionCompatible: true, RequiresSelector: true, AlgebraHomomorphism: false, PhysicalAction: false, Verdict: "dimension-compatible only after choosing which two contact points receive rank 3; no canonical selector exists"},
		{Name: "Clifford Cartan commutant embedding", Route: c.Route, Domain: "Cartan idempotents in Cl(7) spinor action", Target: "seven contact spectral cells", DomainDim: c.PrimitiveCartanIdempotents, TargetDim: c.ContactPointCount, CanonicalPredata: true, DimensionCompatible: c.DimensionalEmbeddingPossible, RequiresSelector: true, AlgebraHomomorphism: false, PhysicalAction: false, Verdict: c.Verdict},
		{Name: "quartic primary regular module", Route: q.Route, Domain: "Q[x]/q4(x)", Target: "abstract 4D scalar module", DomainDim: q.QuarticPrimaryDim, TargetDim: q.ScalarCarrierDim, CanonicalPredata: q.GaloisSafePrimaryIdeal, DimensionCompatible: q.RankOneDimensionMatch, RequiresSelector: false, AlgebraHomomorphism: q.CompanionAlgebraActionAvailable, PhysicalAction: false, Verdict: "abstract rank-one quartic module is derived, but not yet identified with physical H_Φ"},
		{Name: "quartic primary action on physical H_Φ", Route: q.Route, Domain: "Q[x]/q4(x)", Target: "H_Φ", DomainDim: q.QuarticPrimaryDim, TargetDim: q.ScalarCarrierDim, CanonicalPredata: q.GaloisSafePrimaryIdeal, DimensionCompatible: q.RankOneDimensionMatch, RequiresSelector: false, AlgebraHomomorphism: false, PhysicalAction: q.PhysicalScalarBundleDerived, Verdict: q.Verdict},
	}
}

func auditCandidates(xs []Candidate) Summary {
	s := Summary{CandidatesAudited: len(xs)}
	for _, x := range xs {
		if x.DimensionCompatible {
			s.DimensionCompatibleCandidates++
		}
		if x.RequiresSelector {
			s.SelectorBlockedCandidates++
		}
		if x.AlgebraHomomorphism {
			s.AlgebraHomomorphisms++
		}
		if x.Target == "H_Fock" && x.PhysicalAction {
			s.PhysicalFockActions++
		}
		if x.Route == QuarticScalarRoute && x.AlgebraHomomorphism && !x.PhysicalAction {
			s.AbstractQuarticModules++
		}
		if x.Target == "H_Φ" && x.PhysicalAction {
			s.PhysicalScalarActions++
		}
	}
	s.CompletePhysicalBundleMaps = s.PhysicalFockActions + s.PhysicalScalarActions
	s.Comment = "Fock and Cartan routes are selector-blocked; quartic scalar route is dimension-compatible and algebraically real only abstractly"
	return s
}

func sum(xs []int) int {
	total := 0
	for _, x := range xs {
		total += x
	}
	return total
}

func FormatFockRank(a FockRankObstructionAudit) string {
	return fmt.Sprintf("route=%s points=%d fockDim=%d idempotents=%d symmetryUniform=%t uniformInteger=%t rem=%d ranks=%v nonuniform=%t selector=%t contactSelector=%t clifford=%t cliffordC7=%t action=%t (%s)", a.Route, a.ContactPointCount, a.FockDimension, a.UnitalFaithfulActionRequiresIdempotents, a.UniformRankRequiredByTransitiveSymmetry, a.UniformRankInteger, a.RemainderModuloPoints, a.ExampleNearUniformRanks, a.NonUniformRanksExist, a.NonUniformRanksRequireSelector, a.CanonicalContactPointSelectorAvailable, a.CliffordVectorActionAvailable, a.CliffordActionMultiplicativeForC7, a.FaithfulMultiplicativeFockActionDerived, a.Verdict)
}

func FormatCartan(a CartanCommutantAudit) string {
	return fmt.Sprintf("route=%s spinorDim=%d odd=%d rank=%d cartanCells=%d contactPoints=%d dimFit=%t chooseCartan=%t canonicalCartan=%t deleteOrMerge=%t preservesContact=%t embedding=%t (%s)", a.Route, a.SpinorDimension, a.CliffordOddGeneratorCount, a.MaximalCommutingGeneratorRank, a.PrimitiveCartanIdempotents, a.ContactPointCount, a.DimensionalEmbeddingPossible, a.RequiresChoiceOfCartan, a.CanonicalCartanSelectorDerived, a.RequiresDeleteOrMergeOneCartanPoint, a.ContactSpectralOrderPreserved, a.EmbeddingIntoCommutantDerived, a.Verdict)
}

func FormatQuarticScalar(a QuarticScalarModuleAudit) string {
	return fmt.Sprintf("route=%s qdim=%d hphi=%d q=%q rank1=%t rankObstruction=%t galois=%t abstract=%t companion=%t actsHphi=%t hphiOperator=%t minPoly=%t physical=%t chernWeil=%t (%s)", a.Route, a.QuarticPrimaryDim, a.ScalarCarrierDim, a.QuarticPolynomial, a.RankOneDimensionMatch, a.IntegerRankObstruction, a.GaloisSafePrimaryIdeal, a.AbstractRegularModuleDerived, a.CompanionAlgebraActionAvailable, a.ActsOnPhysicalHphi, a.CanonicalHphiBasisOrOperatorDerived, a.ScalarOperatorHasQuarticMinimalPolynomial, a.PhysicalScalarBundleDerived, a.ChernWeilReady, a.Verdict)
}

func FormatCandidate(a Candidate) string {
	return fmt.Sprintf("%s[%s: %s -> %s dims=%d->%d predata=%t dim=%t selector=%t hom=%t physical=%t: %s]", a.Name, a.Route, a.Domain, a.Target, a.DomainDim, a.TargetDim, a.CanonicalPredata, a.DimensionCompatible, a.RequiresSelector, a.AlgebraHomomorphism, a.PhysicalAction, a.Verdict)
}

func FormatCandidates(xs []Candidate) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, FormatCandidate(x))
	}
	return strings.Join(parts, " | ")
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("candidates=%d dimCompatible=%d selectorBlocked=%d hom=%d fockActions=%d abstractQuartic=%d scalarActions=%d complete=%d (%s)", a.CandidatesAudited, a.DimensionCompatibleCandidates, a.SelectorBlockedCandidates, a.AlgebraHomomorphisms, a.PhysicalFockActions, a.AbstractQuarticModules, a.PhysicalScalarActions, a.CompletePhysicalBundleMaps, a.Comment)
}

func FormatFirewall(a Firewall) string {
	return fmt.Sprintf("observed=%t arbitrary=%t base=%t clifford=%t fockNoGo=%t cartanNoGo=%t quarticAbstract=%t fockAction=%t scalarAction=%t bundle=%t chernWeil=%t heat=%t thresholds=%t absolute=%t constants=%t strict=%d->%d conditional=%d->%d verdict=%s", a.UsesObservedInputForDerivation, a.ArbitraryLinearMapUsed, a.ContactBaseInherited, a.CliffordPreactionInherited, a.FockRankObstructionProved, a.CartanCommutantObstructionProved, a.QuarticScalarAbstractModuleDerived, a.CanonicalFockActionDerived, a.CanonicalScalarActionDerived, a.PhysicalBundleMapDerived, a.ChernWeilCarrierDerived, a.HeatKernelMatchingDerived, a.ThresholdCorrectedBetaDerived, a.AbsoluteCouplingPromoted, a.PhysicalConstantsDerived, a.StrictNullityBefore, a.StrictNullityAfter, a.ConditionalNullityBefore, a.ConditionalNullityAfter, a.Verdict)
}
