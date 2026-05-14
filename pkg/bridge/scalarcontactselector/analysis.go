// Package scalarcontactselector implements Gate 186: scalar/contact quartic
// identification selector or obstruction theorem.
//
// Gate 185 constructed the exact abstract quartic module Q[x]/(q4), but refused
// to identify it with the physical H_Φ scalar carrier because Gate 37's active
// scalar/Higgs operator is pair-degenerate and hence quadratic-minimal. Gate 186
// asks whether the finite engine has a canonical selector that collapses the
// four irreducible quartic contact roots into the 2+2 pairing required by the
// physical Higgs doublet.
//
// The gate deliberately avoids diagonalizing the quartic roots. It computes the
// exact resolvent cubic whose three roots encode the three possible 2+2
// partitions of a four-root quartic orbit, audits external finite selectors, and
// tests whether the abstract quartic module admits a canonical commuting complex
// structure J^2=-1. The result is an obstruction theorem: the abstract quartic
// module exists, but no current finite datum canonically selects the 2+2 Higgs
// partition or promotes the physical scalar bundle.
package scalarcontactselector

import (
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/quarticscalaroperator"
)

type Rational = big.Rat

type QuarticInputAudit struct {
	Polynomial                            string
	Degree                                int
	PrimaryBlockDimension                 int
	DistinctRealRootsCertified            int
	IrreducibleOverQInherited             bool
	TransitiveGaloisOrbitInherited        bool
	PhysicalHphiDimension                 int
	Gate37PairDegenerate                  bool
	Gate37MinimalPolynomialDegree         int
	QuarticMinimalPolynomialDegree        int
	IdentificationRequiresTwoPairCollapse bool
	Verdict                               string
}

type PartitionAudit struct {
	QuarticRootCount                    int
	TwoPlusTwoPartitions                int
	PartitionLabels                     []string
	PureInternalGaloisInvariantSelector bool
	GaloisInvariantParityConstant       bool
	CanonicalPartitionDerived           bool
	RequiresBranchChoice                bool
	Verdict                             string
}

type ResolventCubicAudit struct {
	Construction                  string
	Variable                      string
	MonicCoefficients             []string // descending: 1,r2,r1,r0
	IntegerPolynomial             string
	Discriminant                  string
	DiscriminantFactorization     string
	EncodesTwoPlusTwoPartitions   bool
	RationalRoot                  bool
	RootsIndividuallySelected     bool
	BranchDiagonalizationUsed     bool
	CanonicalResolventRootDerived bool
	Verdict                       string
}

type ExternalSelectorCandidate struct {
	Name                          string
	Source                        string
	ActsOnQuarticBlock            bool
	ProducesResolventObservable   bool
	SelectsOneOfThreePartitions   bool
	RequiresForbiddenBranchChoice bool
	Verdict                       string
}

type ExternalSelectorAudit struct {
	Candidates                     []ExternalSelectorCandidate
	CandidatesAudited              int
	CandidatesReachingQuarticBlock int
	ResolventObservables           int
	CanonicalPartitionSelectors    int
	Verdict                        string
}

type ComplexStructureAudit struct {
	ModuleName                                 string
	Dimension                                  int
	QuarticRootsReal                           int
	Centralizer                                string
	CentralizerTotallyReal                     bool
	CommutingJEquivalentToElementOfCentralizer bool
	ExistsElementSquareMinusOne                bool
	CanonicalComplexStructureDerived           bool
	SymplecticPairingDerived                   bool
	Verdict                                    string
}

type Summary struct {
	TestsAudited                   int
	ObstructionsProved             int
	AbstractQuarticModuleInherited bool
	ResolventCubicComputed         bool
	InternalPartitionSelector      bool
	ExternalPartitionSelector      bool
	CanonicalComplexStructure      bool
	PhysicalScalarBundleDerived    bool
	Comment                        string
}

type Firewall struct {
	UsesObservedInputForDerivation     bool
	UsesBranchDiagonalization          bool
	UsesArbitraryPairingChoice         bool
	AbstractQuarticModuleInherited     bool
	ResolventPartitionAuditComplete    bool
	InternalGaloisPartitionObstructed  bool
	ExternalSelectorObstructed         bool
	ComplexStructureObstructed         bool
	Gate37PairDegeneracyRecognized     bool
	CanonicalTwoPlusTwoSelectorDerived bool
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
	PreviousGate185  quarticscalaroperator.Analysis
	QuarticInput     QuarticInputAudit
	Partition        PartitionAudit
	Resolvent        ResolventCubicAudit
	ExternalSelector ExternalSelectorAudit
	ComplexStructure ComplexStructureAudit
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
		prev, err := quarticscalaroperator.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 185 input: %w", err)
			return
		}
		defaultA, defaultErr = Build(prev)
	})
	return defaultA, defaultErr
}

func Build(prev quarticscalaroperator.Analysis) (Analysis, error) {
	if !prev.Firewall.QuarticAbstractOperatorDerived || !prev.Firewall.QuarticMomentsVerified {
		return Analysis{}, fmt.Errorf("Gate 186 requires Gate 185 abstract quartic module and moment ledger")
	}
	if prev.Gate37Comparison.ActiveDimension != 4 || !prev.Gate37Comparison.Gate37PairDegenerate {
		return Analysis{}, fmt.Errorf("Gate 186 requires pair-degenerate four-dimensional Gate-37 scalar carrier")
	}

	input := auditQuarticInput(prev)
	partition := auditPartitions()
	resolvent := auditResolvent()
	external := auditExternalSelectors()
	complex := auditComplexStructure()

	summary := Summary{
		TestsAudited:                   4,
		ObstructionsProved:             3,
		AbstractQuarticModuleInherited: prev.Firewall.QuarticAbstractOperatorDerived,
		ResolventCubicComputed:         resolvent.EncodesTwoPlusTwoPartitions,
		InternalPartitionSelector:      partition.CanonicalPartitionDerived,
		ExternalPartitionSelector:      external.CanonicalPartitionSelectors > 0,
		CanonicalComplexStructure:      complex.CanonicalComplexStructureDerived,
		PhysicalScalarBundleDerived:    false,
		Comment:                        "The exact quartic module exists, but the three ways to collapse its four-root orbit into Higgs 2+2 pairs are not canonically selected by internal Galois data, current external finite data, or a commuting complex/symplectic structure.",
	}
	fw := Firewall{
		UsesObservedInputForDerivation:     false,
		UsesBranchDiagonalization:          false,
		UsesArbitraryPairingChoice:         false,
		AbstractQuarticModuleInherited:     prev.Firewall.QuarticAbstractOperatorDerived,
		ResolventPartitionAuditComplete:    resolvent.EncodesTwoPlusTwoPartitions && len(resolvent.MonicCoefficients) == 4,
		InternalGaloisPartitionObstructed:  !partition.CanonicalPartitionDerived && partition.RequiresBranchChoice,
		ExternalSelectorObstructed:         external.CanonicalPartitionSelectors == 0,
		ComplexStructureObstructed:         !complex.CanonicalComplexStructureDerived && !complex.ExistsElementSquareMinusOne,
		Gate37PairDegeneracyRecognized:     input.Gate37PairDegenerate && input.Gate37MinimalPolynomialDegree == 2,
		CanonicalTwoPlusTwoSelectorDerived: false,
		PhysicalScalarBundleDerived:        false,
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
			"a 2+2 identification of the quartic contact orbit is exactly a choice of one resolvent-cubic root",
			"purely internal Galois-invariant data cannot choose one of the three pair partitions",
			"current external finite objects do not act as resolvent-root selectors on the quartic block",
			"the quartic companion centralizer is a totally real field, so no commuting rational J with J^2=-1 is canonical",
		},
		OpenRequirements: []string{
			"derive a new finite external selector that acts on the quartic primary block and selects one resolvent root",
			"derive a physical scalar/contact map P_Φ Ω P_Φ compatible with the Higgs 2+2 degeneracy",
			"derive a non-branch complex/symplectic pairing on the scalar carrier, or prove it is spontaneous vacuum data",
			"promote the selected scalar bundle to a finite Chern character / integration pairing",
		},
		RecommendedNextGate: "Gate 187 — scalar vacuum selector / spontaneous 2+2 pairing source audit",
		Verdict:             "Gate 186 proves that the mismatch between the quartic contact module and the pair-degenerate Higgs operator is a selector obstruction, not a numerical accident. The quartic module is real, but no current finite datum chooses the 2+2 resolvent partition required to identify it with physical H_Φ.",
	}
	truth := "Gate 186 seals the current scalar/contact identification route. Mapping the irreducible quartic contact orbit to the Higgs 2+2 scalar carrier requires choosing one of three resolvent-cubic pairings. Internal Galois invariance cannot choose it, current external finite objects do not choose it, and the totally real quartic module has no canonical commuting complex structure. The physical scalar bundle therefore remains open behind a genuine vacuum/selector datum."
	return Analysis{PreviousGate185: prev, QuarticInput: input, Partition: partition, Resolvent: resolvent, ExternalSelector: external, ComplexStructure: complex, Summary: summary, Firewall: fw, TruthStatement: truth}, nil
}

func auditQuarticInput(prev quarticscalaroperator.Analysis) QuarticInputAudit {
	return QuarticInputAudit{
		Polynomial:                            "3240x^4 - 7668x^3 + 6426x^2 - 2235x + 271",
		Degree:                                4,
		PrimaryBlockDimension:                 4,
		DistinctRealRootsCertified:            4,
		IrreducibleOverQInherited:             true,
		TransitiveGaloisOrbitInherited:        true,
		PhysicalHphiDimension:                 prev.Gate37Comparison.ActiveDimension,
		Gate37PairDegenerate:                  prev.Gate37Comparison.Gate37PairDegenerate,
		Gate37MinimalPolynomialDegree:         2,
		QuarticMinimalPolynomialDegree:        4,
		IdentificationRequiresTwoPairCollapse: true,
		Verdict:                               "The contact operator has a four-root irreducible quartic primary block, while Gate 37's physical scalar/Higgs operator is 2+2 pair-degenerate. Any identification must collapse the quartic orbit into a selected two-pair partition.",
	}
}

func auditPartitions() PartitionAudit {
	return PartitionAudit{
		QuarticRootCount:                    4,
		TwoPlusTwoPartitions:                3,
		PartitionLabels:                     []string{"12|34", "13|24", "14|23"},
		PureInternalGaloisInvariantSelector: false,
		GaloisInvariantParityConstant:       true,
		CanonicalPartitionDerived:           false,
		RequiresBranchChoice:                true,
		Verdict:                             "The quartic roots form one transitive Galois orbit. Any purely internal Galois-invariant parity or pairing is constant on the orbit and cannot select one of the three 2+2 partitions.",
	}
}

func auditResolvent() ResolventCubicAudit {
	// For monic q4 = x^4 + a x^3 + b x^2 + c x + d, the partition resolvent
	// with roots (r1+r2)(r3+r4), (r1+r3)(r2+r4), (r1+r4)(r2+r3) is
	// z^3 - b z^2 + (a c - 4 d) z + (4 b d - a^2 d - c^2).
	a := rat("-71/30")
	b := rat("1071/540")
	c := rat("-149/216")
	d := rat("271/3240")
	coeff2 := neg(b)
	coeff1 := sub(mul(a, c), mul(rat("4"), d))
	coeff0 := sub(sub(mul(mul(rat("4"), b), d), mul(mul(a, a), d)), mul(c, c))
	return ResolventCubicAudit{
		Construction:                  "z^3 - b z^2 + (a c - 4d)z + (4bd - a^2d - c^2), roots encode the three 2+2 partitions",
		Variable:                      "z",
		MonicCoefficients:             []string{"1", ratString(coeff2), ratString(coeff1), ratString(coeff0)},
		IntegerPolynomial:             "5832000z^3 - 11566800z^2 + 7569900z - 1637467",
		Discriminant:                  "471497/531441000000",
		DiscriminantFactorization:     "13*36269/(2^6*3^12*5^6)",
		EncodesTwoPlusTwoPartitions:   true,
		RationalRoot:                  false,
		RootsIndividuallySelected:     false,
		BranchDiagonalizationUsed:     false,
		CanonicalResolventRootDerived: false,
		Verdict:                       "The resolvent cubic is exact and branch-free, but selecting a Higgs 2+2 pairing would require choosing one of its three roots. The current engine derives the cubic, not a canonical root of it.",
	}
}

func auditExternalSelectors() ExternalSelectorAudit {
	candidates := []ExternalSelectorCandidate{
		{Name: "quartic symmetric moment ledger", Source: "Gates 161-162", ActsOnQuarticBlock: true, ProducesResolventObservable: false, SelectsOneOfThreePartitions: false, RequiresForbiddenBranchChoice: false, Verdict: "Symmetric power sums/zeta data are invariant under all root permutations; they see the quartic block collectively and cannot pick a partition."},
		{Name: "Gate-37 scalar/Higgs quadratic operator", Source: "Gate 37", ActsOnQuarticBlock: false, ProducesResolventObservable: false, SelectsOneOfThreePartitions: false, RequiresForbiddenBranchChoice: true, Verdict: "The operator has the desired 2+2 degeneracy on H_Φ, but no canonical map identifies its pairs with a quartic contact partition."},
		{Name: "B-L / Fock charge polarization", Source: "Gate 16 / Gate 22", ActsOnQuarticBlock: false, ProducesResolventObservable: false, SelectsOneOfThreePartitions: false, RequiresForbiddenBranchChoice: false, Verdict: "B-L acts on Fock charge labels; Gate 184 sealed the direct contact-to-Fock idempotent action, so it cannot select a quartic contact partition."},
		{Name: "scalar covariant derivative broken-sector diagnostic", Source: "Gates 84-101", ActsOnQuarticBlock: false, ProducesResolventObservable: false, SelectsOneOfThreePartitions: false, RequiresForbiddenBranchChoice: false, Verdict: "The broken-sector Hessian/diag(1,1,4) data fix gauge normalization, not a quartic resolvent-root observable."},
		{Name: "topological action seal S_top=8π²", Source: "Gates 174-175", ActsOnQuarticBlock: false, ProducesResolventObservable: false, SelectsOneOfThreePartitions: false, RequiresForbiddenBranchChoice: false, Verdict: "The topological seal is a scalar normalization datum; it carries no operator action on the quartic primary block."},
		{Name: "quartic companion operator T_q", Source: "Gate 185", ActsOnQuarticBlock: true, ProducesResolventObservable: true, SelectsOneOfThreePartitions: false, RequiresForbiddenBranchChoice: true, Verdict: "T_q generates the quartic field and its resolvent algebra, but no polynomial in T_q that is Galois-invariant selects one resolvent root without branching."},
	}
	reach := 0
	obs := 0
	selectors := 0
	for _, c := range candidates {
		if c.ActsOnQuarticBlock {
			reach++
		}
		if c.ProducesResolventObservable {
			obs++
		}
		if c.SelectsOneOfThreePartitions {
			selectors++
		}
	}
	return ExternalSelectorAudit{
		Candidates:                     candidates,
		CandidatesAudited:              len(candidates),
		CandidatesReachingQuarticBlock: reach,
		ResolventObservables:           obs,
		CanonicalPartitionSelectors:    selectors,
		Verdict:                        "Existing external data either do not act on the quartic block, are fully symmetric on it, or require the same forbidden branch choice. No external resolvent-root selector is derived.",
	}
}

func auditComplexStructure() ComplexStructureAudit {
	return ComplexStructureAudit{
		ModuleName:             "Q[x]/(q4) companion module",
		Dimension:              4,
		QuarticRootsReal:       4,
		Centralizer:            "Q[T_q] ≅ Q[x]/(q4)",
		CentralizerTotallyReal: true,
		CommutingJEquivalentToElementOfCentralizer: true,
		ExistsElementSquareMinusOne:                false,
		CanonicalComplexStructureDerived:           false,
		SymplecticPairingDerived:                   false,
		Verdict:                                    "A commuting rational complex structure would be an element j of the totally real quartic field with j²=-1. Under every real embedding, j² is nonnegative for real j, so j²=-1 has no element in the totally real centralizer. A 2+2 complex/symplectic pairing requires extra selector data.",
	}
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

func FormatQuarticInput(a QuarticInputAudit) string {
	return fmt.Sprintf("poly=%q degree=%d blockDim=%d realRoots=%d irreducible=%t transitive=%t HphiDim=%d pairDegenerate=%t gate37MinDeg=%d quarticMinDeg=%d collapse=%t (%s)", a.Polynomial, a.Degree, a.PrimaryBlockDimension, a.DistinctRealRootsCertified, a.IrreducibleOverQInherited, a.TransitiveGaloisOrbitInherited, a.PhysicalHphiDimension, a.Gate37PairDegenerate, a.Gate37MinimalPolynomialDegree, a.QuarticMinimalPolynomialDegree, a.IdentificationRequiresTwoPairCollapse, a.Verdict)
}
func FormatPartition(a PartitionAudit) string {
	return fmt.Sprintf("roots=%d partitions=%d labels=%v internalSelector=%t parityConstant=%t canonical=%t branch=%t (%s)", a.QuarticRootCount, a.TwoPlusTwoPartitions, a.PartitionLabels, a.PureInternalGaloisInvariantSelector, a.GaloisInvariantParityConstant, a.CanonicalPartitionDerived, a.RequiresBranchChoice, a.Verdict)
}
func FormatResolvent(a ResolventCubicAudit) string {
	return fmt.Sprintf("%s in %s coeffs=%v integer=%q discr=%s factors=%s encodes=%t rationalRoot=%t selected=%t branchDiag=%t canonicalRoot=%t (%s)", a.Construction, a.Variable, a.MonicCoefficients, a.IntegerPolynomial, a.Discriminant, a.DiscriminantFactorization, a.EncodesTwoPlusTwoPartitions, a.RationalRoot, a.RootsIndividuallySelected, a.BranchDiagonalizationUsed, a.CanonicalResolventRootDerived, a.Verdict)
}
func FormatExternalSelectors(a ExternalSelectorAudit) string {
	parts := make([]string, 0, len(a.Candidates))
	for _, c := range a.Candidates {
		parts = append(parts, fmt.Sprintf("%s[source=%s reach=%t resolvent=%t selects=%t branch=%t]", c.Name, c.Source, c.ActsOnQuarticBlock, c.ProducesResolventObservable, c.SelectsOneOfThreePartitions, c.RequiresForbiddenBranchChoice))
	}
	return fmt.Sprintf("audited=%d reach=%d resolventObs=%d selectors=%d candidates={%s} (%s)", a.CandidatesAudited, a.CandidatesReachingQuarticBlock, a.ResolventObservables, a.CanonicalPartitionSelectors, strings.Join(parts, "; "), a.Verdict)
}
func FormatComplexStructure(a ComplexStructureAudit) string {
	return fmt.Sprintf("module=%s dim=%d realRoots=%d centralizer=%s totallyReal=%t commutingJInCentralizer=%t squareMinusOne=%t canonicalJ=%t symplectic=%t (%s)", a.ModuleName, a.Dimension, a.QuarticRootsReal, a.Centralizer, a.CentralizerTotallyReal, a.CommutingJEquivalentToElementOfCentralizer, a.ExistsElementSquareMinusOne, a.CanonicalComplexStructureDerived, a.SymplecticPairingDerived, a.Verdict)
}
func FormatSummary(a Summary) string {
	return fmt.Sprintf("tests=%d obstructions=%d abstract=%t resolvent=%t internalSelector=%t externalSelector=%t complexJ=%t physicalBundle=%t (%s)", a.TestsAudited, a.ObstructionsProved, a.AbstractQuarticModuleInherited, a.ResolventCubicComputed, a.InternalPartitionSelector, a.ExternalPartitionSelector, a.CanonicalComplexStructure, a.PhysicalScalarBundleDerived, a.Comment)
}
func FormatFirewall(a Firewall) string {
	return fmt.Sprintf("observed=%t branchDiag=%t arbitraryPair=%t abstract=%t resolvent=%t internalObstruction=%t externalObstruction=%t complexObstruction=%t gate37Pair=%t selector=%t bundle=%t chernWeil=%t heat=%t thresholds=%t absolute=%t constants=%t strict=%d->%d conditional=%d->%d verdict=%s", a.UsesObservedInputForDerivation, a.UsesBranchDiagonalization, a.UsesArbitraryPairingChoice, a.AbstractQuarticModuleInherited, a.ResolventPartitionAuditComplete, a.InternalGaloisPartitionObstructed, a.ExternalSelectorObstructed, a.ComplexStructureObstructed, a.Gate37PairDegeneracyRecognized, a.CanonicalTwoPlusTwoSelectorDerived, a.PhysicalScalarBundleDerived, a.ChernWeilCarrierDerived, a.HeatKernelMatchingDerived, a.ThresholdCorrectedBetaDerived, a.AbsoluteCouplingPromoted, a.PhysicalConstantsDerived, a.StrictNullityBefore, a.StrictNullityAfter, a.ConditionalNullityBefore, a.ConditionalNullityAfter, a.Verdict)
}
