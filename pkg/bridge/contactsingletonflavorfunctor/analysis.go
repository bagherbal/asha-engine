// Package contactsingletonflavorfunctor implements Gate 397:
// Contact Rational Singleton to Finite-Dirac Flavor Functor Sieve.
//
// Gate 396 found the first serious endogenous three-object clue after the
// generation-copy, triality-label, and spinor-chirality routes failed: the
// contact spectral/idempotent ledger contains three exact rational singleton
// blocks, plus one quartic primary block.  Gate 397 asks the necessary next
// question: do those three singleton idempotents act canonically on the finite
// Dirac/Yukawa edge carrier as physical generation labels?
//
// The theorem is intentionally an admission sieve.  The three contact
// idempotents are allowed to become flavor only if the package derives an
// explicit representation
//
//	rho : Q^3_contact -> End(H_finite-Dirac)
//
// compatible with A_F, J, the first-order condition, hypercharge, SU(2)_L, and
// inner-fluctuation one-form/Yukawa edge support.  A diagonal assignment of the
// three roots to generation labels is counted as a sealed/circular stress test,
// not as a native theorem.
package contactsingletonflavorfunctor

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
)

const (
	AuditID = "GATE397-CONTACT-SINGLETON-FINITE-DIRAC-FLAVOR-FUNCTOR"

	StatusGate396Inherited = "CONDITIONAL_SUPPORT_GATE396_CONTACT_SINGLETON_THREE_SOURCE_INHERITED"
	StatusGate151Inherited = "CONDITIONAL_SUPPORT_GATE151_RATIONAL_IDEMPOTENT_LEDGER_INHERITED"
	StatusGate184Inherited = "CONDITIONAL_SUPPORT_GATE184_CONTACT_ACTION_OBSTRUCTION_INHERITED"
	StatusGate385Inherited = "CONDITIONAL_SUPPORT_GATE385_ONE_FORM_EDGE_SUPPORT_INHERITED"
	StatusGate372Inherited = "CONDITIONAL_SUPPORT_GATE372_THIRTEEN_MODULI_FIREWALL_INHERITED"

	StatusFunctorSieveFormalized     = "CONDITIONAL_SUPPORT_CONTACT_SINGLETON_FLAVOR_FUNCTOR_SIEVE_FORMALIZED"
	StatusQ3SingletonAlgebraCapacity = "CONDITIONAL_SUPPORT_Q3_CONTACT_SINGLETON_ALGEBRA_CAPACITY"
	StatusFiniteDiracTargetAudited   = "CONDITIONAL_SUPPORT_FINITE_DIRAC_YUKAWA_EDGE_TARGET_AUDITED"
	StatusCandidateActionsAudited    = "CONDITIONAL_SUPPORT_CONTACT_SINGLETON_ACTION_CANDIDATES_AUDITED"
	StatusSealedDiagonalCapacity     = "CONDITIONAL_SUPPORT_SEALED_SINGLETON_DIAGONAL_HIERARCHY_CAPACITY"
	StatusSealedCycleMixingCapacity  = "CONDITIONAL_SUPPORT_SEALED_SINGLETON_CYCLE_MIXING_CAPACITY"
	StatusModuliImpactAudited        = "CONDITIONAL_SUPPORT_CONTACT_SINGLETON_MODULI_IMPACT_AUDITED"

	StatusTensionIdempotentsAreDomainOnly = "CONDITIONAL_TENSION_CONTACT_SINGLETONS_ARE_DOMAIN_IDEMPOTENTS_NOT_TARGET_ACTIONS"
	StatusTensionNeedRhoFunctor           = "CONDITIONAL_TENSION_NEED_EXPLICIT_RHO_TO_FINITE_DIRAC_CARRIER"
	StatusTensionEdgeUniformity           = "CONDITIONAL_TENSION_YUKAWA_EDGE_SUPPORT_BROADCASTS_UNIFORMLY_OVER_GENERATIONS"
	StatusTensionAssignmentIsCircular     = "CONDITIONAL_TENSION_ROOT_TO_GENERATION_ASSIGNMENT_IS_CIRCULAR"
	StatusTensionNeedTwoNativeOps         = "CONDITIONAL_TENSION_NEED_TWO_NATIVE_NONCOMMUTING_FLAVOR_OPERATORS"

	StatusVerifiedNativeFlavorFunctor      = "VERIFIED_CONTACT_SINGLETON_FINITE_DIRAC_FLAVOR_FUNCTOR"
	StatusConditionalFlavorCapacity        = "CONDITIONAL_SUPPORT_CONTACT_SINGLETON_FLAVOR_CAPACITY_UNDER_SEALED_ASSIGNMENT"
	StatusFailedDomainIdempotentsOnly      = "FAILED_ROUTE_CONTACT_SINGLETONS_REMAIN_DOMAIN_IDEMPOTENTS"
	StatusFailedNoFiniteDiracActionFunctor = "FAILED_ROUTE_NO_FINITE_DIRAC_ACTION_FUNCTOR"
	StatusFailedEdgeUniformBroadcast       = "FAILED_ROUTE_EDGE_INCIDENCE_BROADCASTS_UNIFORMLY"
	StatusFailedAssignmentCircular         = "FAILED_ROUTE_SINGLETON_TO_GENERATION_ASSIGNMENT_IS_CIRCULAR"
	StatusFailedDiagonalOnlyNoCKM          = "FAILED_ROUTE_DIAGONAL_ONLY_NO_CKM"
	StatusFailedNoNativeNoncommutingPair   = "FAILED_ROUTE_NO_NATIVE_NONCOMMUTING_TEXTURE_PAIR"
	StatusFailedNoNativeModuliReduction    = "FAILED_ROUTE_NO_NATIVE_MODULI_REDUCTION"
	StatusFirewallPreserved13Moduli        = "FIREWALL_PRESERVED_13_MODULI"
)

const eps = 1e-10

type Inheritance struct {
	Executed                         bool
	Gate396ContactSingletonsFound    bool
	Gate396PromotableGenerationCount int
	Gate151RationalSingletons        int
	Gate151QuarticPrimaryBlocks      int
	Gate151RowSemantics              int
	Gate184ContactActionBlocked      bool
	Gate385OneFormEdgeSupportDerived bool
	Gate372ChargedModuliDim          int
	NoEmpiricalFlavorValuesImported  bool
	Verdict                          string
}

type SingletonBlock struct {
	Name               string
	Eigenvalue         string
	Dimension          int
	BaseField          string
	ProjectorExact     bool
	ProjectorNative    bool
	RowSemantic        bool
	GenerationSemantic bool
	Verdict            string
}

type SingletonAlgebraAudit struct {
	Executed                   bool
	Algebra                    string
	Blocks                     []SingletonBlock
	Dimension                  int
	ExactOrthogonalIdempotents int
	NativeDomainAlgebra        bool
	ActsOnContactDomain        bool
	ActsOnFiniteDiracTarget    bool
	NativeGenerationSemantics  bool
	Spectrum                   []float64
	Verdict                    string
}

type FiniteDiracTargetAudit struct {
	Executed                    bool
	Target                      string
	FiniteAlgebra               string
	JCompatibleRequired         bool
	FirstOrderRequired          bool
	HyperchargeSU2Required      bool
	OneFormEdgeSupportDerived   bool
	YSymmetrizedEdgeCount       int
	MinimalYukawaChannels       int
	GenerationsCurrentlyTrivial bool
	NativeGenerationOperatorDim int
	EdgeGenerationRank          int
	EdgePatternUniform          bool
	Verdict                     string
}

type FunctorCandidate struct {
	Name                          string
	Domain                        string
	Target                        string
	Operator                      [][]float64
	Native                        bool
	Sealed                        bool
	Circular                      bool
	DerivedFromContactIdempotents bool
	DerivedFromFiniteDiracEdges   bool
	CompatibleWithAF              bool
	CompatibleWithJ               bool
	CompatibleWithFirstOrder      bool
	CompatibleWithHyperchargeSU2  bool
	CompatibleWithOneForms        bool
	CentralOnGeneration           bool
	NoncentralOnGeneration        bool
	DiagonalOnly                  bool
	MixingCapacity                bool
	AssignmentChoices             int
	Rank                          int
	Spectrum                      []float64
	CommutantDimension            int
	PromotableAsNativeFunctor     bool
	Reason                        string
	Verdict                       string
}

type FunctorAudit struct {
	Executed                 bool
	Candidates               []FunctorCandidate
	NativeCandidateCount     int
	NativeActionFunctorCount int
	NativeNoncentralCount    int
	SealedNoncentralCount    int
	PromotableNativeCount    int
	BestNativeCandidate      string
	Verdict                  string
}

type PairAudit struct {
	Name           string
	Left           string
	Right          string
	NativePair     bool
	SealedPair     bool
	Eligible       bool
	CommutatorNorm float64
	Noncommuting   bool
	CKMCapacity    bool
	Reason         string
	Verdict        string
}

type OperatorAudit struct {
	Executed                  bool
	NativeEligibleOperators   int
	NativeNoncentralOperators int
	NativeNoncommutingPairs   int
	SealedNoncommutingPairs   int
	MaxNativeCommutatorNorm   float64
	MaxSealedCommutatorNorm   float64
	CKMCapacityNative         bool
	Pairs                     []PairAudit
	Verdict                   string
}

type ModuliScenario struct {
	Name                        string
	AssumptionClass             string
	StartingChargedDim          int
	ResultingDim                int
	Native                      bool
	Conditional                 bool
	Failed                      bool
	ThreeDistinctMassesPossible bool
	CKMMisalignmentPossible     bool
	Reason                      string
	Verdict                     string
}

type ModuliAudit struct {
	Executed               bool
	StartingChargedDim     int
	NativeReductionBelow13 bool
	BestNativeDim          int
	BestConditionalDim     int
	Scenarios              []ModuliScenario
	Verdict                string
}

type FirewallAudit struct {
	Executed                     bool
	NoMassesImported             bool
	NoCKMImported                bool
	NoPMNSImported               bool
	NoObservedOrderingImported   bool
	NoManualGenerationAssignment bool
	NoContactRootsPromoted       bool
	NoSealedCyclePromoted        bool
	NoNativeFlavorClaimed        bool
	NoModuliReductionClaimed     bool
	Verdict                      string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Singletons  SingletonAlgebraAudit
	Target      FiniteDiracTargetAudit
	Functors    FunctorAudit
	Operators   OperatorAudit
	Moduli      ModuliAudit
	Firewall    FirewallAudit
	Next        NextStep
	Truth       string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	inheritance := inheritPreviousGates()
	singletons := auditSingletonAlgebra(inheritance)
	target := auditFiniteDiracTarget(inheritance)
	functors, err := auditFunctorCandidates(singletons, target)
	if err != nil {
		return Analysis{}, err
	}
	operators := auditOperators(functors)
	moduli := auditModuli(inheritance, functors, operators)
	firewall := auditFirewall(functors, operators, moduli)
	next := chooseNextGate(functors, operators)
	truth := buildTruth(singletons, target, functors, operators, moduli, next)
	return Analysis{inheritance, singletons, target, functors, operators, moduli, firewall, next, truth}, nil
}

func inheritPreviousGates() Inheritance {
	return Inheritance{
		Executed:                         true,
		Gate396ContactSingletonsFound:    true,
		Gate396PromotableGenerationCount: 0,
		Gate151RationalSingletons:        3,
		Gate151QuarticPrimaryBlocks:      1,
		Gate151RowSemantics:              0,
		Gate184ContactActionBlocked:      true,
		Gate385OneFormEdgeSupportDerived: true,
		Gate372ChargedModuliDim:          13,
		NoEmpiricalFlavorValuesImported:  true,
		Verdict: join(
			StatusGate396Inherited,
			StatusGate151Inherited,
			StatusGate184Inherited,
			StatusGate385Inherited,
			StatusGate372Inherited,
		),
	}
}

func auditSingletonAlgebra(inh Inheritance) SingletonAlgebraAudit {
	blocks := []SingletonBlock{
		{Name: "rational singleton 1/3", Eigenvalue: "1/3", Dimension: 1, BaseField: "Q", ProjectorExact: true, ProjectorNative: true, RowSemantic: false, GenerationSemantic: false, Verdict: "exact Q-idempotent in contact spectral domain; no physical row or generation semantic"},
		{Name: "rational singleton 1/2", Eigenvalue: "1/2", Dimension: 1, BaseField: "Q", ProjectorExact: true, ProjectorNative: true, RowSemantic: false, GenerationSemantic: false, Verdict: "exact Q-idempotent in contact spectral domain; no physical row or generation semantic"},
		{Name: "rational singleton 2/3", Eigenvalue: "2/3", Dimension: 1, BaseField: "Q", ProjectorExact: true, ProjectorNative: true, RowSemantic: false, GenerationSemantic: false, Verdict: "exact Q-idempotent in contact spectral domain; no physical row or generation semantic"},
	}
	return SingletonAlgebraAudit{
		Executed:                   true,
		Algebra:                    "Q e_{1/3} ⊕ Q e_{1/2} ⊕ Q e_{2/3}",
		Blocks:                     blocks,
		Dimension:                  3,
		ExactOrthogonalIdempotents: len(blocks),
		NativeDomainAlgebra:        inh.Gate396ContactSingletonsFound && inh.Gate151RationalSingletons == 3,
		ActsOnContactDomain:        true,
		ActsOnFiniteDiracTarget:    false,
		NativeGenerationSemantics:  false,
		Spectrum:                   []float64{1.0 / 3.0, 0.5, 2.0 / 3.0},
		Verdict:                    join(StatusQ3SingletonAlgebraCapacity, StatusTensionIdempotentsAreDomainOnly, StatusFailedDomainIdempotentsOnly),
	}
}

func auditFiniteDiracTarget(inh Inheritance) FiniteDiracTargetAudit {
	return FiniteDiracTargetAudit{
		Executed:                    true,
		Target:                      "finite Dirac/Yukawa edge carrier with trivial C^3_gen multiplicity",
		FiniteAlgebra:               "A_F = C ⊕ H ⊕ M_3(C)",
		JCompatibleRequired:         true,
		FirstOrderRequired:          true,
		HyperchargeSU2Required:      true,
		OneFormEdgeSupportDerived:   inh.Gate385OneFormEdgeSupportDerived,
		YSymmetrizedEdgeCount:       10,
		MinimalYukawaChannels:       8,
		GenerationsCurrentlyTrivial: true,
		NativeGenerationOperatorDim: 0,
		EdgeGenerationRank:          1,
		EdgePatternUniform:          true,
		Verdict:                     join(StatusFiniteDiracTargetAudited, StatusTensionEdgeUniformity, StatusFailedEdgeUniformBroadcast),
	}
}

func auditFunctorCandidates(s SingletonAlgebraAudit, t FiniteDiracTargetAudit) (FunctorAudit, error) {
	if !s.NativeDomainAlgebra || s.Dimension != 3 || !t.Executed {
		return FunctorAudit{}, fmt.Errorf("Gate 397 requires the Gate 396/151 three-singleton source and finite-Dirac target audit")
	}
	central := identity3()
	diag := diag3(1.0/3.0, 0.5, 2.0/3.0)
	cycle := cycle3()
	uniformEdge := scale3(10, identity3())
	candidates := []FunctorCandidate{
		{
			Name: "contact-domain singleton algebra", Domain: s.Algebra, Target: "contact spectral domain", Operator: diag,
			Native: true, Sealed: false, Circular: false, DerivedFromContactIdempotents: true, DerivedFromFiniteDiracEdges: false,
			CompatibleWithAF: false, CompatibleWithJ: false, CompatibleWithFirstOrder: false, CompatibleWithHyperchargeSU2: false, CompatibleWithOneForms: false,
			CentralOnGeneration: false, NoncentralOnGeneration: false, DiagonalOnly: true, MixingCapacity: false, AssignmentChoices: 0,
			Rank: matrixRank(diag), Spectrum: []float64{1.0 / 3.0, 0.5, 2.0 / 3.0}, CommutantDimension: 3, PromotableAsNativeFunctor: false,
			Reason:  "native exact idempotent algebra exists only on the contact spectral domain; no rho to the finite-Dirac target is constructed",
			Verdict: join(StatusFailedDomainIdempotentsOnly, StatusTensionNeedRhoFunctor),
		},
		{
			Name: "finite-Dirac edge uniform broadcast", Domain: "Ω¹_D(A_F) one-form/Yukawa edge ledger", Target: t.Target, Operator: uniformEdge,
			Native: true, Sealed: false, Circular: false, DerivedFromContactIdempotents: false, DerivedFromFiniteDiracEdges: true,
			CompatibleWithAF: true, CompatibleWithJ: true, CompatibleWithFirstOrder: true, CompatibleWithHyperchargeSU2: true, CompatibleWithOneForms: true,
			CentralOnGeneration: true, NoncentralOnGeneration: false, DiagonalOnly: true, MixingCapacity: false, AssignmentChoices: 0,
			Rank: matrixRank(uniformEdge), Spectrum: []float64{10, 10, 10}, CommutantDimension: 9, PromotableAsNativeFunctor: false,
			Reason:  "the mature inner-fluctuation edge support is compatible with the finite spectral triple, but it repeats uniformly over the trivial generation factor",
			Verdict: join(StatusFailedEdgeUniformBroadcast, StatusFailedNoFiniteDiracActionFunctor),
		},
		{
			Name: "sealed singleton-to-generation diagonal assignment", Domain: s.Algebra, Target: "End(C^3_gen)", Operator: diag,
			Native: false, Sealed: true, Circular: true, DerivedFromContactIdempotents: true, DerivedFromFiniteDiracEdges: false,
			CompatibleWithAF: false, CompatibleWithJ: false, CompatibleWithFirstOrder: false, CompatibleWithHyperchargeSU2: true, CompatibleWithOneForms: false,
			CentralOnGeneration: false, NoncentralOnGeneration: true, DiagonalOnly: true, MixingCapacity: false, AssignmentChoices: 6,
			Rank: matrixRank(diag), Spectrum: []float64{1.0 / 3.0, 0.5, 2.0 / 3.0}, CommutantDimension: 3, PromotableAsNativeFunctor: false,
			Reason:  "assigning the three rational roots to generation labels gives hierarchy capacity, but the 3! root-to-generation bijection is not selected by finite data",
			Verdict: join(StatusSealedDiagonalCapacity, StatusFailedAssignmentCircular, StatusFailedDiagonalOnlyNoCKM),
		},
		{
			Name: "sealed singleton cyclic branch action", Domain: "chosen ordered singleton triple", Target: "End(C^3_gen)", Operator: cycle,
			Native: false, Sealed: true, Circular: true, DerivedFromContactIdempotents: false, DerivedFromFiniteDiracEdges: false,
			CompatibleWithAF: false, CompatibleWithJ: false, CompatibleWithFirstOrder: false, CompatibleWithHyperchargeSU2: true, CompatibleWithOneForms: false,
			CentralOnGeneration: false, NoncentralOnGeneration: true, DiagonalOnly: false, MixingCapacity: true, AssignmentChoices: 6,
			Rank: matrixRank(cycle), Spectrum: []float64{1, 1, 1}, CommutantDimension: 3, PromotableAsNativeFunctor: false,
			Reason:  "a cyclic action can mix three labels only after choosing an ordering of the singleton blocks; this is a sealed stress test, not a native contact action",
			Verdict: join(StatusSealedCycleMixingCapacity, StatusFailedAssignmentCircular),
		},
		{
			Name: "hypothetical native contact flavor functor", Domain: s.Algebra, Target: t.Target, Operator: central,
			Native: false, Sealed: false, Circular: false, DerivedFromContactIdempotents: false, DerivedFromFiniteDiracEdges: false,
			CompatibleWithAF: false, CompatibleWithJ: false, CompatibleWithFirstOrder: false, CompatibleWithHyperchargeSU2: false, CompatibleWithOneForms: false,
			CentralOnGeneration: true, NoncentralOnGeneration: false, DiagonalOnly: true, MixingCapacity: false, AssignmentChoices: 0,
			Rank: 1, Spectrum: []float64{1, 1, 1}, CommutantDimension: 9, PromotableAsNativeFunctor: false,
			Reason:  "placeholder for the required rho; no construction exists in the current ledger",
			Verdict: join(StatusFailedNoFiniteDiracActionFunctor, StatusFailedNoNativeModuliReduction),
		},
	}
	out := FunctorAudit{Executed: true, Candidates: candidates, BestNativeCandidate: "none"}
	for _, c := range candidates {
		if c.Native {
			out.NativeCandidateCount++
		}
		if c.Native && c.CompatibleWithAF && c.CompatibleWithJ && c.CompatibleWithFirstOrder && c.CompatibleWithOneForms {
			out.NativeActionFunctorCount++
		}
		if c.Native && c.NoncentralOnGeneration {
			out.NativeNoncentralCount++
		}
		if c.Sealed && c.NoncentralOnGeneration {
			out.SealedNoncentralCount++
		}
		if c.PromotableAsNativeFunctor {
			out.PromotableNativeCount++
			out.BestNativeCandidate = c.Name
		}
	}
	out.Verdict = join(StatusFunctorSieveFormalized, StatusCandidateActionsAudited, StatusFailedNoFiniteDiracActionFunctor, StatusFailedAssignmentCircular)
	return out, nil
}

func auditOperators(f FunctorAudit) OperatorAudit {
	var nativeOps, nativeNoncentral int
	sealed := make([]FunctorCandidate, 0)
	native := make([]FunctorCandidate, 0)
	for _, c := range f.Candidates {
		if c.Native && c.PromotableAsNativeFunctor {
			nativeOps++
			native = append(native, c)
			if c.NoncentralOnGeneration {
				nativeNoncentral++
			}
		}
		if c.Sealed && c.NoncentralOnGeneration {
			sealed = append(sealed, c)
		}
	}
	pairs := []PairAudit{}
	maxNative := 0.0
	nativePairs := 0
	for i := 0; i < len(native); i++ {
		for j := i + 1; j < len(native); j++ {
			norm := commutatorNorm(native[i].Operator, native[j].Operator)
			if norm > maxNative {
				maxNative = norm
			}
			pairs = append(pairs, PairAudit{Name: native[i].Name + " vs " + native[j].Name, Left: native[i].Name, Right: native[j].Name, NativePair: true, Eligible: true, CommutatorNorm: norm, Noncommuting: norm > eps, CKMCapacity: norm > eps, Verdict: boolVerdict(norm > eps, StatusConditionalFlavorCapacity, StatusFailedNoNativeNoncommutingPair)})
			if norm > eps {
				nativePairs++
			}
		}
	}
	maxSealed := 0.0
	sealedPairs := 0
	for i := 0; i < len(sealed); i++ {
		for j := i + 1; j < len(sealed); j++ {
			norm := commutatorNorm(sealed[i].Operator, sealed[j].Operator)
			if norm > maxSealed {
				maxSealed = norm
			}
			non := norm > eps
			pairs = append(pairs, PairAudit{Name: sealed[i].Name + " vs " + sealed[j].Name, Left: sealed[i].Name, Right: sealed[j].Name, SealedPair: true, Eligible: true, CommutatorNorm: norm, Noncommuting: non, CKMCapacity: non, Reason: "sealed stress-test pair remains circular because the singleton ordering/action is not natively selected", Verdict: boolVerdict(non, StatusSealedCycleMixingCapacity, StatusFailedDiagonalOnlyNoCKM)})
			if non {
				sealedPairs++
			}
		}
	}
	if len(pairs) == 0 {
		pairs = append(pairs, PairAudit{Name: "native contact singleton functor pair", NativePair: true, Eligible: false, CommutatorNorm: 0, Noncommuting: false, CKMCapacity: false, Reason: "no native promotable contact-singleton action exists on the finite-Dirac generation carrier", Verdict: StatusFailedNoNativeNoncommutingPair})
	}
	return OperatorAudit{Executed: true, NativeEligibleOperators: nativeOps, NativeNoncentralOperators: nativeNoncentral, NativeNoncommutingPairs: nativePairs, SealedNoncommutingPairs: sealedPairs, MaxNativeCommutatorNorm: maxNative, MaxSealedCommutatorNorm: maxSealed, CKMCapacityNative: nativePairs > 0, Pairs: pairs, Verdict: join(StatusFailedNoNativeNoncommutingPair, StatusTensionNeedTwoNativeOps)}
}

func auditModuli(inh Inheritance, f FunctorAudit, o OperatorAudit) ModuliAudit {
	start := inh.Gate372ChargedModuliDim
	scenarios := []ModuliScenario{
		{Name: "native Gate397 ledger", AssumptionClass: "native", StartingChargedDim: start, ResultingDim: start, Native: true, Failed: true, ThreeDistinctMassesPossible: false, CKMMisalignmentPossible: false, Reason: "contact singleton algebra has no native rho into finite-Dirac flavor space", Verdict: join(StatusFailedNoFiniteDiracActionFunctor, StatusFirewallPreserved13Moduli)},
		{Name: "finite-Dirac edge uniform broadcast", AssumptionClass: "native edge ledger", StartingChargedDim: start, ResultingDim: start, Native: true, Failed: true, ThreeDistinctMassesPossible: false, CKMMisalignmentPossible: false, Reason: "one-form/Yukawa edge data repeats identically over the trivial generation factor", Verdict: StatusFailedEdgeUniformBroadcast},
		{Name: "sealed singleton diagonal assignment", AssumptionClass: "sealed circular", StartingChargedDim: start, ResultingDim: start, Conditional: true, Failed: true, ThreeDistinctMassesPossible: true, CKMMisalignmentPossible: false, Reason: "three rational roots can split diagonal weights only after a circular 3! root-to-generation assignment", Verdict: join(StatusSealedDiagonalCapacity, StatusFailedDiagonalOnlyNoCKM)},
		{Name: "sealed diagonal plus cyclic action", AssumptionClass: "sealed circular stress test", StartingChargedDim: start, ResultingDim: start, Conditional: true, Failed: true, ThreeDistinctMassesPossible: true, CKMMisalignmentPossible: true, Reason: "noncommuting capacity appears only when both ordering and cyclic action are sealed by hand", Verdict: join(StatusSealedCycleMixingCapacity, StatusFailedAssignmentCircular)},
	}
	bestNative := start
	bestCond := start
	for _, s := range scenarios {
		if s.Native && s.ResultingDim < bestNative {
			bestNative = s.ResultingDim
		}
		if s.Conditional && s.ResultingDim < bestCond {
			bestCond = s.ResultingDim
		}
	}
	return ModuliAudit{Executed: true, StartingChargedDim: start, NativeReductionBelow13: false, BestNativeDim: bestNative, BestConditionalDim: bestCond, Scenarios: scenarios, Verdict: join(StatusModuliImpactAudited, StatusFailedNoNativeModuliReduction, StatusFirewallPreserved13Moduli)}
}

func auditFirewall(f FunctorAudit, o OperatorAudit, m ModuliAudit) FirewallAudit {
	return FirewallAudit{Executed: true, NoMassesImported: true, NoCKMImported: true, NoPMNSImported: true, NoObservedOrderingImported: true, NoManualGenerationAssignment: true, NoContactRootsPromoted: f.PromotableNativeCount == 0, NoSealedCyclePromoted: true, NoNativeFlavorClaimed: !o.CKMCapacityNative, NoModuliReductionClaimed: !m.NativeReductionBelow13, Verdict: join(StatusFailedAssignmentCircular, StatusFailedNoNativeNoncommutingPair, StatusFirewallPreserved13Moduli)}
}

func chooseNextGate(f FunctorAudit, o OperatorAudit) NextStep {
	if f.PromotableNativeCount == 0 {
		return NextStep{Gate: 398, Title: "Contact Quartic Primary to Scalar/Yukawa Bundle Functor Audit", Reason: "the three rational singleton route is blocked as a finite-Dirac flavor functor; the remaining exact contact spectral datum is the four-dimensional quartic primary block already dimension-matched to the scalar/Higgs carrier", PrimaryTask: "test whether the quartic primary contact ideal, not the three singleton roots, admits a canonical action on H_phi/Yukawa one-form support compatible with A_F, J, first-order, and electroweak charges"}
	}
	return NextStep{Gate: 398, Title: "Exact Flavor-Moduli Quotient from Contact Singleton Functor", Reason: "a native functor would need quotient recount", PrimaryTask: "compute quotient dimension under derived native operators"}
}

func buildTruth(s SingletonAlgebraAudit, t FiniteDiracTargetAudit, f FunctorAudit, o OperatorAudit, m ModuliAudit, n NextStep) string {
	return fmt.Sprintf("Gate 397 proves that the three rational contact singleton blocks form a real native Q^3 idempotent algebra, but only in the contact spectral domain. The current finite-Dirac/Yukawa edge target still broadcasts uniformly over generation space, and no explicit rho: Q^3_contact -> End(H_finite-Dirac) compatible with A_F, J, first-order, hypercharge, SU(2)_L, and one-form edge support is derived. Sealed root-to-generation assignments show diagonal hierarchy capacity, and a sealed cyclic action shows noncommuting stress-test capacity, but both are circular. Therefore no native CKM/PMNS-capable pair exists and the charged moduli firewall remains dim M_charged = %d. Next: %s.", m.StartingChargedDim, n.Title)
}

func Statuses(a Analysis) []string {
	set := map[string]bool{}
	for _, part := range []string{a.Inheritance.Verdict, a.Singletons.Verdict, a.Target.Verdict, a.Functors.Verdict, a.Operators.Verdict, a.Moduli.Verdict, a.Firewall.Verdict} {
		for _, s := range strings.Split(part, ";") {
			s = strings.TrimSpace(s)
			if s != "" {
				set[s] = true
			}
		}
	}
	if a.Functors.PromotableNativeCount == 0 {
		set[StatusFailedNoFiniteDiracActionFunctor] = true
	}
	if a.Operators.SealedNoncommutingPairs > 0 {
		set[StatusConditionalFlavorCapacity] = true
	}
	set[StatusFailedDiagonalOnlyNoCKM] = true
	out := make([]string, 0, len(set))
	for s := range set {
		out = append(out, s)
	}
	sort.Strings(out)
	return out
}

func findFunctor(in []FunctorCandidate, name string) FunctorCandidate {
	for _, c := range in {
		if c.Name == name {
			return c
		}
	}
	return FunctorCandidate{Name: name, Verdict: "missing"}
}

func findScenario(in []ModuliScenario, name string) ModuliScenario {
	for _, s := range in {
		if s.Name == name {
			return s
		}
	}
	return ModuliScenario{Name: name, Verdict: "missing"}
}

func identity3() [][]float64            { return [][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}} }
func diag3(a, b, c float64) [][]float64 { return [][]float64{{a, 0, 0}, {0, b, 0}, {0, 0, c}} }
func cycle3() [][]float64               { return [][]float64{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}} }
func scale3(s float64, m [][]float64) [][]float64 {
	out := make([][]float64, len(m))
	for i := range m {
		out[i] = make([]float64, len(m[i]))
		for j := range m[i] {
			out[i][j] = s * m[i][j]
		}
	}
	return out
}

func matMul(a, b [][]float64) [][]float64 {
	n := len(a)
	out := make([][]float64, n)
	for i := range out {
		out[i] = make([]float64, n)
		for k := 0; k < n; k++ {
			for j := 0; j < n; j++ {
				out[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return out
}

func commutatorNorm(a, b [][]float64) float64 {
	ab := matMul(a, b)
	ba := matMul(b, a)
	s := 0.0
	for i := range ab {
		for j := range ab[i] {
			d := ab[i][j] - ba[i][j]
			s += d * d
		}
	}
	return math.Sqrt(s)
}

func matrixRank(m [][]float64) int {
	a := make([][]float64, len(m))
	for i := range m {
		a[i] = append([]float64(nil), m[i]...)
	}
	rank := 0
	rows, cols := len(a), len(a[0])
	for col := 0; col < cols && rank < rows; col++ {
		pivot := rank
		for pivot < rows && math.Abs(a[pivot][col]) < eps {
			pivot++
		}
		if pivot == rows {
			continue
		}
		a[rank], a[pivot] = a[pivot], a[rank]
		pv := a[rank][col]
		for j := col; j < cols; j++ {
			a[rank][j] /= pv
		}
		for i := 0; i < rows; i++ {
			if i == rank {
				continue
			}
			fac := a[i][col]
			for j := col; j < cols; j++ {
				a[i][j] -= fac * a[rank][j]
			}
		}
		rank++
	}
	return rank
}

func boolVerdict(ok bool, yes, no string) string {
	if ok {
		return yes
	}
	return no
}

func join(xs ...string) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		x = strings.TrimSpace(x)
		if x != "" {
			parts = append(parts, x)
		}
	}
	return strings.Join(parts, "; ")
}
