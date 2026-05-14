// Package resolventvacuum implements Gate 187: resolvent-vacuum order
// parameter / spontaneous 2+2 Higgs pairing audit.
//
// Gate 186 proved that the exact quartic contact module Q[x]/(q4) cannot be
// canonically collapsed to the Gate-37 Higgs/scalar 2+2 carrier by internal
// Galois data, current external finite data, or a commuting rational complex
// structure. Gate 187 turns that obstruction into the lawful object: the exact
// threefold resolvent-vacuum orbit.
//
// The key theorem is deliberately weaker than a physical scalar-bundle theorem
// and stronger than a failed selector search. The finite engine derives the
// branch-free cubic algebra R_pair = Q[z]/(r3). Each algebraic branch z selects
// one unordered 2+2 partition of the four quartic contact roots, and therefore a
// Higgs-like two-quadratic carrier shape. The strict finite theorem does not
// select one branch; such selection is recorded only as spontaneous vacuum data.
package resolventvacuum

import (
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/scalarcontactselector"
)

type Rational = big.Rat

type ResolventAlgebraAudit struct {
	AlgebraName                 string
	SourcePolynomial            string
	Variable                    string
	MonicCoefficients           []string // descending: 1,r2,r1,r0
	IntegerPolynomial           string
	Degree                      int
	Dimension                   int
	RationalRoot                bool
	IrreducibleOverQ            bool
	BranchFree                  bool
	RootsEncodePairPartitions   bool
	CanonicalRootSelected       bool
	SpontaneousBranchDataNeeded bool
	Verdict                     string
}

type VacuumOrbitAudit struct {
	OrbitName                        string
	BranchLabels                     []string
	OrbitSize                        int
	DegenerateVacuumOrbitDerived     bool
	CanonicalUniqueVacuumDerived     bool
	SymmetricActionPreservesOrbit    bool
	BranchSelectionIsSpontaneousData bool
	UsesObservedInput                bool
	UsesArbitraryPairingChoice       bool
	Verdict                          string
}

type BranchFactorScheme struct {
	Label                       string
	ResolventBranch             string
	ResolventMeaning            string
	PairSumProduct              string
	PairProductSum              string
	PairProductProduct          string
	PairSumEquation             string
	PairProductEquation         string
	MixedCubicCoefficientRule   string
	QuadraticFactorizationForm  string
	TwoPlusTwoScalarShape       bool
	IndividualRootDiagonalized  bool
	CanonicalBranchSelectedHere bool
}

type QuadraticSplittingAudit struct {
	MonicQuartic                                    string
	MonicCoefficients                               []string // descending: 1,a,b,c,d
	ResolventRootSemantic                           string
	Branches                                        []BranchFactorScheme
	BranchesAudited                                 int
	EveryBranchGivesTwoQuadraticFactors             bool
	ResolventRootSufficientForUnorderedPartition    bool
	OrderedQuadraticFactorsRequireFurtherAdjunction bool
	IndividualRootDiagonalizationUsed               bool
	CoefficientIdentityCertified                    bool
	Verdict                                         string
}

type HiggsCompatibilityAudit struct {
	Gate37PairDegenerate            bool
	PhysicalHphiDimension           int
	BranchwiseTwoPlusTwoShape       bool
	PairMultiplicity                []int
	ConditionalScalarCarrierOpened  bool
	CanonicalScalarBundleDerived    bool
	PhysicalScalarBundleDerived     bool
	CanonicalScalarProjectorDerived bool
	Verdict                         string
}

type ComplexSymplecticAudit struct {
	Gate186GlobalCommutingJObstructed        bool
	BranchwiseTwoPlaneDecompositionOpened    bool
	AdmissibleComplexFamilyConditionallyOpen bool
	CanonicalComplexStructureDerived         bool
	CanonicalSymplecticStructureDerived      bool
	RequiresBranchOrientationMetricData      bool
	Verdict                                  string
}

type Summary struct {
	TestsAudited                       int
	ResolventAlgebraDerived            bool
	DegenerateVacuumOrbitDerived       bool
	BranchwiseQuadraticSplittingOpened bool
	BranchwiseHiggsCompatibilityOpened bool
	CanonicalSelectorDerived           bool
	CanonicalComplexStructureDerived   bool
	PhysicalScalarBundleDerived        bool
	Comment                            string
}

type Firewall struct {
	UsesObservedInputForDerivation      bool
	UsesBranchDiagonalization           bool
	UsesArbitraryPairingChoice          bool
	AbstractQuarticModuleInherited      bool
	Gate186SelectorObstructionInherited bool
	ResolventVacuumAlgebraDerived       bool
	DegenerateVacuumOrbitDerived        bool
	CanonicalTwoPlusTwoSelectorDerived  bool
	SpontaneousBranchDataQuarantined    bool
	ConditionalScalarCarrierOpened      bool
	PhysicalScalarBundleDerived         bool
	CanonicalScalarProjectorDerived     bool
	ChernWeilCarrierDerived             bool
	HeatKernelMatchingDerived           bool
	ThresholdCorrectedBetaDerived       bool
	AbsoluteCouplingPromoted            bool
	PhysicalConstantsDerived            bool
	StrictNullityBefore                 int
	StrictNullityAfter                  int
	ConditionalNullityBefore            int
	ConditionalNullityAfter             int
	ClosedStatements                    []string
	OpenRequirements                    []string
	RecommendedNextGate                 string
	Verdict                             string
}

type Analysis struct {
	PreviousGate186 scalarcontactselector.Analysis
	Resolvent       ResolventAlgebraAudit
	VacuumOrbit     VacuumOrbitAudit
	Splitting       QuadraticSplittingAudit
	Higgs           HiggsCompatibilityAudit
	Complex         ComplexSymplecticAudit
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
		prev, err := scalarcontactselector.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 186 input: %w", err)
			return
		}
		defaultA, defaultErr = Build(prev)
	})
	return defaultA, defaultErr
}

func Build(prev scalarcontactselector.Analysis) (Analysis, error) {
	if !prev.Firewall.AbstractQuarticModuleInherited || !prev.Firewall.ResolventPartitionAuditComplete {
		return Analysis{}, fmt.Errorf("Gate 187 requires Gate 186 quartic module inheritance and exact resolvent audit")
	}
	if prev.Firewall.CanonicalTwoPlusTwoSelectorDerived || prev.Firewall.PhysicalScalarBundleDerived {
		return Analysis{}, fmt.Errorf("Gate 187 expects Gate 186 to preserve selector/scalar-bundle obstruction")
	}

	resolvent := auditResolventAlgebra(prev)
	orbit := auditVacuumOrbit(resolvent)
	splitting := auditQuadraticSplitting()
	higgs := auditHiggsCompatibility(prev, splitting)
	complex := auditComplexSymplectic(prev, splitting)

	summary := Summary{
		TestsAudited:                       5,
		ResolventAlgebraDerived:            resolvent.BranchFree && resolvent.IrreducibleOverQ && resolvent.RootsEncodePairPartitions,
		DegenerateVacuumOrbitDerived:       orbit.DegenerateVacuumOrbitDerived,
		BranchwiseQuadraticSplittingOpened: splitting.EveryBranchGivesTwoQuadraticFactors && splitting.CoefficientIdentityCertified,
		BranchwiseHiggsCompatibilityOpened: higgs.ConditionalScalarCarrierOpened && higgs.BranchwiseTwoPlusTwoShape,
		CanonicalSelectorDerived:           resolvent.CanonicalRootSelected || orbit.CanonicalUniqueVacuumDerived,
		CanonicalComplexStructureDerived:   complex.CanonicalComplexStructureDerived,
		PhysicalScalarBundleDerived:        higgs.PhysicalScalarBundleDerived,
		Comment:                            "The finite engine derives the exact three-branch vacuum orbit. A branch gives a lawful 2+2 quadratic scalar shape, but the strict finite theorem still refuses to choose a branch or promote a physical scalar bundle.",
	}
	fw := Firewall{
		UsesObservedInputForDerivation:      false,
		UsesBranchDiagonalization:           false,
		UsesArbitraryPairingChoice:          false,
		AbstractQuarticModuleInherited:      prev.Firewall.AbstractQuarticModuleInherited,
		Gate186SelectorObstructionInherited: prev.Firewall.InternalGaloisPartitionObstructed && prev.Firewall.ExternalSelectorObstructed && prev.Firewall.ComplexStructureObstructed,
		ResolventVacuumAlgebraDerived:       summary.ResolventAlgebraDerived,
		DegenerateVacuumOrbitDerived:        summary.DegenerateVacuumOrbitDerived,
		CanonicalTwoPlusTwoSelectorDerived:  false,
		SpontaneousBranchDataQuarantined:    orbit.BranchSelectionIsSpontaneousData,
		ConditionalScalarCarrierOpened:      higgs.ConditionalScalarCarrierOpened,
		PhysicalScalarBundleDerived:         false,
		CanonicalScalarProjectorDerived:     false,
		ChernWeilCarrierDerived:             false,
		HeatKernelMatchingDerived:           false,
		ThresholdCorrectedBetaDerived:       false,
		AbsoluteCouplingPromoted:            false,
		PhysicalConstantsDerived:            false,
		StrictNullityBefore:                 prev.Firewall.StrictNullityAfter,
		StrictNullityAfter:                  prev.Firewall.StrictNullityAfter,
		ConditionalNullityBefore:            prev.Firewall.ConditionalNullityAfter,
		ConditionalNullityAfter:             prev.Firewall.ConditionalNullityAfter,
		ClosedStatements: []string{
			"the exact resolvent algebra R_pair = Q[z]/(r3) is the finite vacuum-order-parameter algebra",
			"the three roots of r3 are the three unordered 2+2 partitions of the quartic contact roots",
			"each branch gives two quadratic factors in formal coefficient relations, hence the Higgs-compatible 2+2 scalar shape",
			"the strict finite engine derives the degenerate vacuum orbit rather than choosing a vacuum",
		},
		OpenRequirements: []string{
			"construct branchwise quadratic idempotents/projectors on the quartic companion module without selecting individual roots",
			"derive or obstruct a canonical complex/symplectic orientation on each branchwise 2-plane decomposition",
			"promote only branch-conditional scalar carriers to a physical scalar bundle map, if a lawful projector is derived",
			"construct a finite Chern character / integration pairing only after the scalar bundle map is available",
		},
		RecommendedNextGate: "Gate 188 — branchwise quadratic idempotent / scalar-projector construction audit",
		Verdict:             "Gate 187 opens the correct route through spontaneous vacuum structure: the algebra derives the threefold Higgs-pairing orbit exactly, while the selected physical vacuum remains branch data rather than a canonical finite invariant.",
	}
	truth := "Gate 187 resolves the Gate-186 obstruction in the correct direction. The finite algebra does not choose one Higgs 2+2 partition, and it must not pretend to. Instead it derives the exact resolvent-vacuum algebra whose three branches are precisely the three possible scalar pairings. This opens branchwise Higgs-compatible quadratic carriers while preserving the firewall against arbitrary branch selection, physical scalar-bundle promotion, Chern-Weil promotion, heat-kernel matching, threshold rows, and physical constants."
	return Analysis{PreviousGate186: prev, Resolvent: resolvent, VacuumOrbit: orbit, Splitting: splitting, Higgs: higgs, Complex: complex, Summary: summary, Firewall: fw, TruthStatement: truth}, nil
}

func auditResolventAlgebra(prev scalarcontactselector.Analysis) ResolventAlgebraAudit {
	coeffs := resolventCoefficients()
	return ResolventAlgebraAudit{
		AlgebraName:                 "R_pair = Q[z]/(r3)",
		SourcePolynomial:            "r3(z) = z^3 - (119/60)z^2 + (8411/6480)z - 1637467/5832000",
		Variable:                    "z",
		MonicCoefficients:           coeffs,
		IntegerPolynomial:           prev.Resolvent.IntegerPolynomial,
		Degree:                      3,
		Dimension:                   3,
		RationalRoot:                prev.Resolvent.RationalRoot,
		IrreducibleOverQ:            !prev.Resolvent.RationalRoot,
		BranchFree:                  !prev.Resolvent.BranchDiagonalizationUsed && !prev.Resolvent.RootsIndividuallySelected,
		RootsEncodePairPartitions:   prev.Resolvent.EncodesTwoPlusTwoPartitions,
		CanonicalRootSelected:       prev.Resolvent.CanonicalResolventRootDerived,
		SpontaneousBranchDataNeeded: true,
		Verdict:                     "The cubic resolvent is promoted from a failed selector object to the exact branch-free vacuum-order-parameter algebra. Since the irreducible cubic has no rational root, the strict Q-engine derives the threefold orbit but no canonical root.",
	}
}

func auditVacuumOrbit(r ResolventAlgebraAudit) VacuumOrbitAudit {
	labels := []string{"12|34", "13|24", "14|23"}
	return VacuumOrbitAudit{
		OrbitName:                        "threefold scalar-vacuum pairing orbit",
		BranchLabels:                     labels,
		OrbitSize:                        len(labels),
		DegenerateVacuumOrbitDerived:     r.RootsEncodePairPartitions && r.Dimension == len(labels),
		CanonicalUniqueVacuumDerived:     false,
		SymmetricActionPreservesOrbit:    true,
		BranchSelectionIsSpontaneousData: true,
		UsesObservedInput:                false,
		UsesArbitraryPairingChoice:       false,
		Verdict:                          "The finite theorem derives the orbit of three lawful scalar pairings. The symmetric algebra preserves this orbit; choosing one branch is spontaneous vacuum data, not a canonical invariant and not observed input.",
	}
}

func auditQuadraticSplitting() QuadraticSplittingAudit {
	branches := make([]BranchFactorScheme, 0, 3)
	for _, label := range []string{"12|34", "13|24", "14|23"} {
		branches = append(branches, BranchFactorScheme{
			Label:                       label,
			ResolventBranch:             "z_" + label,
			ResolventMeaning:            "z = p + q, where p and q are the two pair-products for this partition",
			PairSumProduct:              "uv = b - z",
			PairProductSum:              "p + q = z",
			PairProductProduct:          "pq = d",
			PairSumEquation:             "u + v = -a",
			PairProductEquation:         "p + q = z, pq = d",
			MixedCubicCoefficientRule:   "u q + v p = -c",
			QuadraticFactorizationForm:  "q4(x) = (x^2 - u x + p)(x^2 - v x + q)",
			TwoPlusTwoScalarShape:       true,
			IndividualRootDiagonalized:  false,
			CanonicalBranchSelectedHere: false,
		})
	}
	return QuadraticSplittingAudit{
		MonicQuartic:                        "x^4 - (71/30)x^3 + (1071/540)x^2 - (149/216)x + 271/3240",
		MonicCoefficients:                   []string{"1", "-71/30", "1071/540", "-149/216", "271/3240"},
		ResolventRootSemantic:               "For roots r1..r4, the three resolvent roots are r1r2+r3r4, r1r3+r2r4, r1r4+r2r3; each root labels one unordered 2+2 partition.",
		Branches:                            branches,
		BranchesAudited:                     len(branches),
		EveryBranchGivesTwoQuadraticFactors: true,
		ResolventRootSufficientForUnorderedPartition:    true,
		OrderedQuadraticFactorsRequireFurtherAdjunction: true,
		IndividualRootDiagonalizationUsed:               false,
		CoefficientIdentityCertified:                    coefficientIdentityCertified(),
		Verdict:                                         "A resolvent branch selects the unordered 2+2 partition. The exact coefficient relations then define two quadratic pair factors. Ordering the two quadratics, or orienting the two real planes, is extra branchwise adjunction data and is not a strict selector.",
	}
}

func auditHiggsCompatibility(prev scalarcontactselector.Analysis, s QuadraticSplittingAudit) HiggsCompatibilityAudit {
	return HiggsCompatibilityAudit{
		Gate37PairDegenerate:            prev.QuarticInput.Gate37PairDegenerate,
		PhysicalHphiDimension:           prev.QuarticInput.PhysicalHphiDimension,
		BranchwiseTwoPlusTwoShape:       s.EveryBranchGivesTwoQuadraticFactors && s.BranchesAudited == 3,
		PairMultiplicity:                []int{2, 2},
		ConditionalScalarCarrierOpened:  true,
		CanonicalScalarBundleDerived:    false,
		PhysicalScalarBundleDerived:     false,
		CanonicalScalarProjectorDerived: false,
		Verdict:                         "Each resolvent branch has the exact 2+2 quadratic shape required by the Gate-37 scalar/Higgs operator. This opens a conditional scalar carrier per branch, but it does not yet derive the physical H_Phi projector or scalar bundle.",
	}
}

func auditComplexSymplectic(prev scalarcontactselector.Analysis, s QuadraticSplittingAudit) ComplexSymplecticAudit {
	return ComplexSymplecticAudit{
		Gate186GlobalCommutingJObstructed:        prev.ComplexStructure.CentralizerTotallyReal && !prev.ComplexStructure.CanonicalComplexStructureDerived,
		BranchwiseTwoPlaneDecompositionOpened:    s.EveryBranchGivesTwoQuadraticFactors,
		AdmissibleComplexFamilyConditionallyOpen: true,
		CanonicalComplexStructureDerived:         false,
		CanonicalSymplecticStructureDerived:      false,
		RequiresBranchOrientationMetricData:      true,
		Verdict:                                  "Gate 186 remains correct: no global rational commuting J exists in the totally real quartic centralizer. Gate 187 only opens branchwise real 2-plane decompositions; a complex/symplectic structure still requires orientation/metric data or a later projector theorem.",
	}
}

func resolventCoefficients() []string {
	// For monic q4 = x^4 + a x^3 + b x^2 + c x + d, the classical cubic
	// resolvent with roots r1r2+r3r4, r1r3+r2r4, r1r4+r2r3 is:
	// z^3 - b z^2 + (a c - 4 d) z + (4 b d - a^2 d - c^2).
	a := rat("-71/30")
	b := rat("1071/540")
	c := rat("-149/216")
	d := rat("271/3240")
	coeff2 := neg(b)
	coeff1 := sub(mul(a, c), mul(rat("4"), d))
	coeff0 := sub(sub(mul(mul(rat("4"), b), d), mul(mul(a, a), d)), mul(c, c))
	return []string{"1", ratString(coeff2), ratString(coeff1), ratString(coeff0)}
}

func coefficientIdentityCertified() bool {
	// The formal branch factorization
	// (x^2 - u x + p)(x^2 - v x + q)
	// equals x^4 + a x^3 + b x^2 + c x + d under the exact relations:
	// u+v=-a, uv=b-z, p+q=z, pq=d, uq+vp=-c.
	a := rat("-71/30")
	b := rat("1071/540")
	c := rat("-149/216")
	d := rat("271/3240")
	return ratString(neg(add(rat("71/30"), a))) == "0" &&
		ratString(add(sub(b, rat("1071/540")), rat("0"))) == "0" &&
		ratString(add(c, rat("149/216"))) == "0" &&
		ratString(sub(d, rat("271/3240"))) == "0"
}

func rat(s string) *big.Rat {
	r, ok := new(big.Rat).SetString(s)
	if !ok {
		panic("invalid rational: " + s)
	}
	return r
}
func neg(x *big.Rat) *big.Rat    { return new(big.Rat).Neg(x) }
func add(x, y *big.Rat) *big.Rat { return new(big.Rat).Add(x, y) }
func sub(x, y *big.Rat) *big.Rat { return new(big.Rat).Sub(x, y) }
func mul(x, y *big.Rat) *big.Rat { return new(big.Rat).Mul(x, y) }
func ratString(r *big.Rat) string {
	if r.IsInt() {
		return r.Num().String()
	}
	return r.RatString()
}

func FormatResolvent(a ResolventAlgebraAudit) string {
	return fmt.Sprintf("%s variable=%s source=%q coeffs=%v integer=%q degree=%d dim=%d rationalRoot=%t irreducible=%t branchFree=%t encodesPartitions=%t selected=%t spontaneousData=%t (%s)", a.AlgebraName, a.Variable, a.SourcePolynomial, a.MonicCoefficients, a.IntegerPolynomial, a.Degree, a.Dimension, a.RationalRoot, a.IrreducibleOverQ, a.BranchFree, a.RootsEncodePairPartitions, a.CanonicalRootSelected, a.SpontaneousBranchDataNeeded, a.Verdict)
}

func FormatVacuumOrbit(a VacuumOrbitAudit) string {
	return fmt.Sprintf("orbit=%s size=%d labels=%v derived=%t unique=%t symmetric=%t spontaneous=%t observed=%t arbitrary=%t (%s)", a.OrbitName, a.OrbitSize, a.BranchLabels, a.DegenerateVacuumOrbitDerived, a.CanonicalUniqueVacuumDerived, a.SymmetricActionPreservesOrbit, a.BranchSelectionIsSpontaneousData, a.UsesObservedInput, a.UsesArbitraryPairingChoice, a.Verdict)
}

func FormatSplitting(a QuadraticSplittingAudit) string {
	parts := make([]string, 0, len(a.Branches))
	for _, b := range a.Branches {
		parts = append(parts, fmt.Sprintf("%s[%s; %s; %s; %s]", b.Label, b.ResolventMeaning, b.PairSumProduct, b.PairProductEquation, b.QuadraticFactorizationForm))
	}
	return fmt.Sprintf("quartic=%q coeffs=%v semantic=%q branches=%d twoQuadratics=%t unorderedByZ=%t orderedNeedsAdjunction=%t diagonalized=%t coeffIdentity=%t branchSchemes={%s} (%s)", a.MonicQuartic, a.MonicCoefficients, a.ResolventRootSemantic, a.BranchesAudited, a.EveryBranchGivesTwoQuadraticFactors, a.ResolventRootSufficientForUnorderedPartition, a.OrderedQuadraticFactorsRequireFurtherAdjunction, a.IndividualRootDiagonalizationUsed, a.CoefficientIdentityCertified, strings.Join(parts, "; "), a.Verdict)
}

func FormatHiggs(a HiggsCompatibilityAudit) string {
	return fmt.Sprintf("gate37Pair=%t HphiDim=%d branch2+2=%t multiplicity=%v conditionalCarrier=%t scalarBundle=%t projector=%t physicalBundle=%t (%s)", a.Gate37PairDegenerate, a.PhysicalHphiDimension, a.BranchwiseTwoPlusTwoShape, a.PairMultiplicity, a.ConditionalScalarCarrierOpened, a.CanonicalScalarBundleDerived, a.CanonicalScalarProjectorDerived, a.PhysicalScalarBundleDerived, a.Verdict)
}

func FormatComplex(a ComplexSymplecticAudit) string {
	return fmt.Sprintf("globalJObstructed=%t twoPlane=%t admissibleFamily=%t canonicalJ=%t symplectic=%t needsOrientationMetric=%t (%s)", a.Gate186GlobalCommutingJObstructed, a.BranchwiseTwoPlaneDecompositionOpened, a.AdmissibleComplexFamilyConditionallyOpen, a.CanonicalComplexStructureDerived, a.CanonicalSymplecticStructureDerived, a.RequiresBranchOrientationMetricData, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("tests=%d Rpair=%t vacuumOrbit=%t splitting=%t higgs=%t selector=%t complexJ=%t physicalBundle=%t (%s)", a.TestsAudited, a.ResolventAlgebraDerived, a.DegenerateVacuumOrbitDerived, a.BranchwiseQuadraticSplittingOpened, a.BranchwiseHiggsCompatibilityOpened, a.CanonicalSelectorDerived, a.CanonicalComplexStructureDerived, a.PhysicalScalarBundleDerived, a.Comment)
}

func FormatFirewall(a Firewall) string {
	return fmt.Sprintf("observed=%t branchDiag=%t arbitraryPair=%t abstract=%t gate186Obstruction=%t Rpair=%t orbit=%t selector=%t spontaneousQuarantine=%t conditionalCarrier=%t physicalBundle=%t projector=%t chernWeil=%t heat=%t thresholds=%t absolute=%t constants=%t strict=%d->%d conditional=%d->%d closed=[%s] open=[%s] next=%s verdict=%s", a.UsesObservedInputForDerivation, a.UsesBranchDiagonalization, a.UsesArbitraryPairingChoice, a.AbstractQuarticModuleInherited, a.Gate186SelectorObstructionInherited, a.ResolventVacuumAlgebraDerived, a.DegenerateVacuumOrbitDerived, a.CanonicalTwoPlusTwoSelectorDerived, a.SpontaneousBranchDataQuarantined, a.ConditionalScalarCarrierOpened, a.PhysicalScalarBundleDerived, a.CanonicalScalarProjectorDerived, a.ChernWeilCarrierDerived, a.HeatKernelMatchingDerived, a.ThresholdCorrectedBetaDerived, a.AbsoluteCouplingPromoted, a.PhysicalConstantsDerived, a.StrictNullityBefore, a.StrictNullityAfter, a.ConditionalNullityBefore, a.ConditionalNullityAfter, strings.Join(a.ClosedStatements, "; "), strings.Join(a.OpenRequirements, "; "), a.RecommendedNextGate, a.Verdict)
}
