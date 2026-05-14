// Package threeobjectsource implements Gate 396:
// Endogenous Three-Object Source Search beyond Spinor Chirality.
//
// Gates 393-395 showed that neither trivial C^3 generation copies, nor a
// direct generation-to-triality relabeling, nor the native Cl(1,7) spinor split
// derive physical generation labels. Gate 396 therefore moves one level earlier:
// it audits native ASHA structures for any endogenous three-object source at all,
// then tests whether such a source is admissible as a finite-Dirac generation
// address. The gate is intentionally a sieve: finding three objects is not enough.
// A candidate must also carry generation semantics, act noncentrally on C^3_gen,
// and remain compatible with A_F, J, the first-order condition, and electroweak
// charges before it can reduce the Gate-372 flavor firewall.
package threeobjectsource

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
)

const (
	AuditID = "GATE396-ENDOGENOUS-THREE-OBJECT-SOURCE-BEYOND-SPINOR-CHIRALITY"

	StatusGate395Inherited = "CONDITIONAL_SUPPORT_GATE395_SPINOR_CHIRALITY_OBSTRUCTION_INHERITED"
	StatusGate394Inherited = "CONDITIONAL_SUPPORT_GATE394_CENTRAL_GENERATION_BROADCAST_INHERITED"
	StatusGate371Inherited = "CONDITIONAL_SUPPORT_GATE371_INFORMATION_NUMBER_LADDER_CAPACITY_INHERITED"
	StatusGate365Inherited = "CONDITIONAL_SUPPORT_GATE365_KMS_TAU_CAPACITY_INHERITED"
	StatusGate151Inherited = "CONDITIONAL_SUPPORT_GATE151_CONTACT_RATIONAL_IDEMPOTENT_LEDGER_INHERITED"
	StatusGate184Inherited = "CONDITIONAL_SUPPORT_GATE184_CONTACT_IDEMPOTENT_ACTION_OBSTRUCTION_INHERITED"
	StatusGate372Inherited = "CONDITIONAL_SUPPORT_GATE372_THIRTEEN_MODULI_FIREWALL_INHERITED"

	StatusSourceSieveAudited       = "CONDITIONAL_SUPPORT_ENDOGENOUS_THREE_OBJECT_SOURCE_SIEVE_AUDITED"
	StatusNativeThreeSourceFound   = "CONDITIONAL_SUPPORT_NATIVE_THREE_OBJECT_SOURCE_FOUND"
	StatusContactSingletonsFound   = "CONDITIONAL_SUPPORT_CONTACT_RATIONAL_SINGLETON_THREE_SOURCE_FOUND"
	StatusFockColorTripletFound    = "CONDITIONAL_SUPPORT_FOCK_SPATIAL_COLOR_TRIPLET_FOUND"
	StatusFanoTriplesAudited       = "CONDITIONAL_SUPPORT_FANO_TRIPLE_FAMILY_AUDITED"
	StatusModularThreeSlotsAudited = "CONDITIONAL_SUPPORT_MODULAR_THREE_SLOT_CAPACITY_AUDITED"
	StatusOperatorCapacityAudited  = "CONDITIONAL_SUPPORT_THREE_SOURCE_OPERATOR_CAPACITY_AUDITED"
	StatusModuliImpactAudited      = "CONDITIONAL_SUPPORT_THREE_SOURCE_MODULI_IMPACT_AUDITED"

	StatusTensionThreeObjectsNotGeneration       = "CONDITIONAL_TENSION_THREE_OBJECTS_ARE_NOT_AUTOMATIC_GENERATIONS"
	StatusTensionContactSemanticsNotFlavor       = "CONDITIONAL_TENSION_CONTACT_ROOT_SEMANTICS_NOT_FINITE_DIRAC_FLAVOR_SEMANTICS"
	StatusTensionColorTripletNotGeneration       = "CONDITIONAL_TENSION_THREE_SPATIAL_FOCK_MODES_ARE_COLOR_NOT_GENERATION"
	StatusTensionFanoFamilyNeedsSelector         = "CONDITIONAL_TENSION_FANO_TRIPLES_FORM_A_FAMILY_AND_REQUIRE_A_SELECTOR"
	StatusTensionTauIsScalarTraceNotEndomorphism = "CONDITIONAL_TENSION_TAU_ETA_IS_SCALAR_TRACE_NOT_GENERATION_ENDOMORPHISM"
	StatusTensionNeedActionFunctor               = "CONDITIONAL_TENSION_NEED_ACTION_FUNCTOR_TO_FINITE_DIRAC_GENERATION_SPACE"
	StatusTensionNeedTwoNativeNoncommutingOps    = "CONDITIONAL_TENSION_NEED_TWO_NATIVE_NONCOMMUTING_GENERATION_TEXTURE_OPERATORS"

	StatusVerifiedNativeGenerationSource     = "VERIFIED_NATIVE_THREE_OBJECT_GENERATION_SOURCE_DERIVED"
	StatusConditionalDynamicLabelsDerived    = "CONDITIONAL_SUPPORT_DYNAMIC_GENERATION_LABELS_DERIVED"
	StatusConditionalCKMCapacityActivated    = "CONDITIONAL_SUPPORT_CKM_MIXING_CAPACITY_ACTIVATED"
	StatusFailedPrimitiveIdempotentsNotThree = "FAILED_ROUTE_PRIMITIVE_IDEMPOTENTS_NOT_THREE_GENERATIONS"
	StatusFailedThreeSourceNotGeneration     = "FAILED_ROUTE_THREE_OBJECT_SOURCE_NOT_GENERATION_ADDRESS"
	StatusFailedContactNoFlavorFunctor       = "FAILED_ROUTE_CONTACT_SINGLETONS_LACK_FINITE_DIRAC_FLAVOR_FUNCTOR"
	StatusFailedFanoSelectorMissing          = "FAILED_ROUTE_FANO_TRIPLES_REQUIRE_SELECTOR"
	StatusFailedColorTripletIsColor          = "FAILED_ROUTE_FOCK_SPATIAL_TRIPLET_IS_COLOR_NOT_GENERATION"
	StatusFailedTauOrNStillCircular          = "FAILED_ROUTE_TAU_OR_N_THREE_SLOT_OPERATOR_REMAINS_CIRCULAR"
	StatusFailedNoNativeNoncommutingPair     = "FAILED_ROUTE_NO_NATIVE_NONCOMMUTING_TEXTURE_PAIR"
	StatusFailedNoNativeModuliReduction      = "FAILED_ROUTE_NO_NATIVE_MODULI_REDUCTION"
	StatusFirewallPreserved13Moduli          = "FIREWALL_PRESERVED_13_MODULI"
)

const eps = 1e-10

type Inheritance struct {
	Executed                         bool
	Gate395SpinorTwoNotThree         bool
	Gate395TrialityCategoryOnly      bool
	Gate394CentralBroadcast          bool
	Gate371NumberLadderNonNative     bool
	Gate365TauKMSNonNative           bool
	Gate151ContactRationalSingletons int
	Gate151QuarticPrimaryBlocks      int
	Gate184FockContactActionBlocked  bool
	Gate372ChargedModuliDim          int
	NoEmpiricalFlavorValuesImported  bool
	Verdict                          string
}

type SourceCandidate struct {
	Name                         string
	Source                       string
	Native                       bool
	Endogenous                   bool
	Sealed                       bool
	CircularIfPromoted           bool
	ObjectCount                  int
	ExactlyThreeObjects          bool
	FamilyCount                  int
	RequiresSelector             bool
	GenerationSemantics          bool
	ContactSemantics             bool
	ColorSemantics               bool
	ScalarTraceSemantics         bool
	CompatibleWithFiniteDirac    bool
	CompatibleWithJ              bool
	CompatibleWithFirstOrder     bool
	CompatibleWithElectroweak    bool
	NoncentralInOwnAlgebra       bool
	NoncentralOnGenerationSpace  bool
	DiagonalOnly                 bool
	MixingCapacity               bool
	Operator                     [][]float64
	Spectrum                     []float64
	CommutantDimension           int
	PromotableAsGenerationSource bool
	Reason                       string
	Verdict                      string
}

type SourceAudit struct {
	Executed                        bool
	Candidates                      []SourceCandidate
	NativeCandidateCount            int
	NativeExactlyThreeSourceCount   int
	NativeGenerationSourceCount     int
	PromotableGenerationSourceCount int
	NativeNoncentralOnGeneration    int
	SealedNoncentralOnGeneration    int
	BestNativeThreeSource           string
	Verdict                         string
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
	NoEmpiricalOrderingImported  bool
	NoManualGenerationAssignment bool
	NoContactRootsPromoted       bool
	NoColorModesPromoted         bool
	NoFanoTriplePromoted         bool
	NoTauOrNPromoted             bool
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
	Sources     SourceAudit
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
	sources, err := auditSources(inheritance)
	if err != nil {
		return Analysis{}, err
	}
	operators := auditOperators(sources)
	moduli := auditModuli(inheritance, sources, operators)
	firewall := auditFirewall(sources, operators, moduli)
	next := chooseNextGate(sources, operators)
	truth := buildTruth(sources, operators, moduli, next)
	return Analysis{inheritance, sources, operators, moduli, firewall, next, truth}, nil
}

func inheritPreviousGates() Inheritance {
	return Inheritance{
		Executed:                         true,
		Gate395SpinorTwoNotThree:         true,
		Gate395TrialityCategoryOnly:      true,
		Gate394CentralBroadcast:          true,
		Gate371NumberLadderNonNative:     true,
		Gate365TauKMSNonNative:           true,
		Gate151ContactRationalSingletons: 3,
		Gate151QuarticPrimaryBlocks:      1,
		Gate184FockContactActionBlocked:  true,
		Gate372ChargedModuliDim:          13,
		NoEmpiricalFlavorValuesImported:  true,
		Verdict: join(
			StatusGate395Inherited,
			StatusGate394Inherited,
			StatusGate371Inherited,
			StatusGate365Inherited,
			StatusGate151Inherited,
			StatusGate184Inherited,
			StatusGate372Inherited,
		),
	}
}

func auditSources(inh Inheritance) (SourceAudit, error) {
	candidates := []SourceCandidate{
		newSourceCandidate(
			"contact rational singleton idempotent blocks",
			"Gate-151 rational/Galois-safe contact spectral decomposition: three rational singleton blocks plus one quartic primary block",
			true, true, false, false,
			inh.Gate151ContactRationalSingletons, 0, false,
			true, false, true, false, false,
			false, false, false, false,
			true, false, true, false,
			diagFromWeights([]float64{-1, 0, 1}),
			"This is the strongest native three-object source found, but its semantics are contact spectral-root/idempotent semantics, not finite-Dirac generation semantics.",
		),
		newSourceCandidate(
			"Fock spatial color triplet",
			"Gate-13/Gate-17 Fock carrier: three spatial modes with B-L=1/3",
			true, true, false, false,
			3, 0, false,
			true, false, false, true, false,
			true, true, true, true,
			true, false, true, false,
			diagFromWeights([]float64{1.0 / 3.0, 1.0 / 3.0, 1.0 / 3.0}),
			"The source is exactly a native triplet, but the project already uses it as color/spatial charge structure; promoting it to generation would confuse color with flavor.",
		),
		newSourceCandidate(
			"octonionic Fano line triples",
			"G2/Fano octonionic incidence: seven oriented triples/lines in the Fano plane",
			true, true, false, false,
			3, 7, true,
			true, false, true, false, false,
			false, false, false, false,
			true, false, false, true,
			cycle3(),
			"Octonionic triples are native, but there is a family of seven; choosing one as generation order requires an additional selector.",
		),
		newSourceCandidate(
			"sealed chosen three-cycle stress test",
			"non-native stress-test action obtained only after choosing one three-object branch/cycle",
			false, true, true, true,
			3, 0, false,
			true, false, false, false, false,
			false, false, false, false,
			true, true, false, true,
			cycle3(),
			"This cycle demonstrates noncommuting capacity with a diagonal three-slot operator, but the chosen branch carrier is sealed and circular.",
		),
		newSourceCandidate(
			"Clifford/Fock primitive idempotent cells",
			"maximal commuting Clifford/Fock idempotent cells; Gate-184 records an eight-cell Cartan obstruction for contact action",
			true, true, false, false,
			8, 0, true,
			false, false, false, false, false,
			false, false, false, false,
			true, false, true, false,
			diagFromWeights([]float64{0, 1, 2, 3, 4, 5, 6, 7}),
			"Primitive idempotent cells are native, but the canonical counts are eight or sixteen, not three; selecting three cells is exactly the missing selector problem.",
		),
		newSourceCandidate(
			"Morita dimension-three bimodule slots",
			"Gate-272 Morita bimodule summands contain dimension-three fundamental/antifundamental slots",
			true, true, false, false,
			2, 0, false,
			false, false, false, true, false,
			true, true, true, true,
			true, false, false, false,
			identity(3),
			"The relevant slots have dimension three because of the M3(C) color block; they are not three generation objects and native edge incidence remains uniform over generations.",
		),
		newSourceCandidate(
			"modular tau_eta three-slot scalar trace",
			"eta-signed scalar/neutral trace sequence tau_eta=(2,-2,1)",
			false, true, true, true,
			3, 0, false,
			true, false, false, false, true,
			false, false, false, false,
			true, true, true, false,
			diagFromWeights([]float64{2, -2, 1}),
			"The three values have hierarchy capacity, but previous gates classify tau_eta as a scalar trace functional, not a generation-space endomorphism.",
		),
		newSourceCandidate(
			"Schrodinger/Fock information number ladder",
			"Gate-371 bridge-level information ladder N=diag(0,1,2)",
			false, true, true, true,
			3, 0, false,
			true, false, false, false, false,
			false, false, false, false,
			true, true, true, false,
			diagFromWeights([]float64{0, 1, 2}),
			"The number ladder breaks copied U(3) degeneracy as a capacity witness, but the finite ASHA ledger still does not derive it as the generation Hamiltonian.",
		),
	}
	for i := range candidates {
		if err := fillSource(&candidates[i]); err != nil {
			return SourceAudit{}, err
		}
	}
	native, native3, nativeGen, promotable, nativeNoncentral, sealedNoncentral := 0, 0, 0, 0, 0, 0
	best := "none"
	for _, c := range candidates {
		if c.Native {
			native++
		}
		if c.Native && c.ExactlyThreeObjects {
			native3++
			if best == "none" {
				best = c.Name
			}
		}
		if c.Native && c.GenerationSemantics {
			nativeGen++
		}
		if c.PromotableAsGenerationSource {
			promotable++
		}
		if c.Native && c.NoncentralOnGenerationSpace {
			nativeNoncentral++
		}
		if c.Sealed && c.NoncentralOnGenerationSpace {
			sealedNoncentral++
		}
	}
	verdict := join(StatusSourceSieveAudited, StatusNativeThreeSourceFound, StatusContactSingletonsFound, StatusFockColorTripletFound, StatusFanoTriplesAudited, StatusModularThreeSlotsAudited, StatusFailedThreeSourceNotGeneration, StatusTensionThreeObjectsNotGeneration, StatusTensionNeedActionFunctor)
	if promotable > 0 {
		verdict = join(StatusSourceSieveAudited, StatusVerifiedNativeGenerationSource)
	}
	return SourceAudit{true, candidates, native, native3, nativeGen, promotable, nativeNoncentral, sealedNoncentral, best, verdict}, nil
}

func newSourceCandidate(name, source string, native, endogenous, sealed, circular bool, objectCount, familyCount int, requiresSelector bool, exactThree, genSem, contactSem, colorSem, scalarSem bool, finiteDirac, j, firstOrder, ew bool, noncentralOwn, noncentralGen, diagonal, mixing bool, op [][]float64, reason string) SourceCandidate {
	return SourceCandidate{
		Name:                        name,
		Source:                      source,
		Native:                      native,
		Endogenous:                  endogenous,
		Sealed:                      sealed,
		CircularIfPromoted:          circular,
		ObjectCount:                 objectCount,
		ExactlyThreeObjects:         exactThree,
		FamilyCount:                 familyCount,
		RequiresSelector:            requiresSelector,
		GenerationSemantics:         genSem,
		ContactSemantics:            contactSem,
		ColorSemantics:              colorSem,
		ScalarTraceSemantics:        scalarSem,
		CompatibleWithFiniteDirac:   finiteDirac,
		CompatibleWithJ:             j,
		CompatibleWithFirstOrder:    firstOrder,
		CompatibleWithElectroweak:   ew,
		NoncentralInOwnAlgebra:      noncentralOwn,
		NoncentralOnGenerationSpace: noncentralGen,
		DiagonalOnly:                diagonal,
		MixingCapacity:              mixing,
		Operator:                    op,
		Reason:                      reason,
	}
}

func fillSource(c *SourceCandidate) error {
	if len(c.Operator) > 0 {
		if err := validateSquare(c.Operator); err != nil {
			return fmt.Errorf("%s operator: %w", c.Name, err)
		}
		c.Spectrum = diagonalSpectrumIfDiagonal(c.Operator)
		c.CommutantDimension = commutantDimensionSnapshot(c)
	}
	c.PromotableAsGenerationSource = c.Native && c.Endogenous && c.ExactlyThreeObjects && c.GenerationSemantics && c.CompatibleWithFiniteDirac && c.CompatibleWithJ && c.CompatibleWithFirstOrder && c.CompatibleWithElectroweak && c.NoncentralOnGenerationSpace && !c.RequiresSelector && !c.CircularIfPromoted

	verdicts := []string{}
	if c.Native && c.ExactlyThreeObjects {
		verdicts = append(verdicts, StatusNativeThreeSourceFound)
	}
	if c.Name == "contact rational singleton idempotent blocks" {
		verdicts = append(verdicts, StatusContactSingletonsFound, StatusTensionContactSemanticsNotFlavor, StatusFailedContactNoFlavorFunctor)
	}
	if c.Name == "Fock spatial color triplet" {
		verdicts = append(verdicts, StatusFockColorTripletFound, StatusTensionColorTripletNotGeneration, StatusFailedColorTripletIsColor)
	}
	if c.Name == "octonionic Fano line triples" {
		verdicts = append(verdicts, StatusFanoTriplesAudited, StatusTensionFanoFamilyNeedsSelector, StatusFailedFanoSelectorMissing)
	}
	if c.Name == "Clifford/Fock primitive idempotent cells" {
		verdicts = append(verdicts, StatusFailedPrimitiveIdempotentsNotThree)
	}
	if c.ScalarTraceSemantics || c.Name == "Schrodinger/Fock information number ladder" {
		verdicts = append(verdicts, StatusModularThreeSlotsAudited, StatusTensionTauIsScalarTraceNotEndomorphism, StatusFailedTauOrNStillCircular)
	}
	if !c.PromotableAsGenerationSource {
		verdicts = append(verdicts, StatusFailedThreeSourceNotGeneration)
	} else {
		verdicts = append(verdicts, StatusVerifiedNativeGenerationSource)
	}
	c.Verdict = join(verdicts...)
	return nil
}

func commutantDimensionSnapshot(c *SourceCandidate) int {
	if !c.NoncentralOnGenerationSpace {
		return 9 // full End(C^3) still commutes with the absent/central generation action.
	}
	if c.DiagonalOnly {
		return 3
	}
	if c.MixingCapacity {
		return 1
	}
	return 9
}

func auditOperators(src SourceAudit) OperatorAudit {
	eligibleNative := []SourceCandidate{}
	sealed := []SourceCandidate{}
	for _, c := range src.Candidates {
		if c.PromotableAsGenerationSource && c.NoncentralOnGenerationSpace {
			eligibleNative = append(eligibleNative, c)
		}
		if c.Sealed && c.NoncentralOnGenerationSpace {
			sealed = append(sealed, c)
		}
	}
	pairs := []PairAudit{}
	maxNative, maxSealed := 0.0, 0.0
	nativePairs, sealedPairs := 0, 0
	for i := 0; i < len(eligibleNative); i++ {
		for j := i + 1; j < len(eligibleNative); j++ {
			p := buildPair(eligibleNative[i], eligibleNative[j], true, false)
			if p.CommutatorNorm > maxNative {
				maxNative = p.CommutatorNorm
			}
			if p.Noncommuting {
				nativePairs++
			}
			pairs = append(pairs, p)
		}
	}
	for i := 0; i < len(sealed); i++ {
		for j := i + 1; j < len(sealed); j++ {
			p := buildPair(sealed[i], sealed[j], false, true)
			if p.CommutatorNorm > maxSealed {
				maxSealed = p.CommutatorNorm
			}
			if p.Noncommuting {
				sealedPairs++
			}
			pairs = append(pairs, p)
		}
	}
	if len(pairs) == 0 {
		pairs = append(pairs, PairAudit{Name: "native eligible operator pair", NativePair: true, Eligible: false, Reason: "no promotable native generation operators were found", Verdict: StatusFailedNoNativeNoncommutingPair})
	}
	return OperatorAudit{
		Executed:                  true,
		NativeEligibleOperators:   len(eligibleNative),
		NativeNoncentralOperators: src.NativeNoncentralOnGeneration,
		NativeNoncommutingPairs:   nativePairs,
		SealedNoncommutingPairs:   sealedPairs,
		MaxNativeCommutatorNorm:   maxNative,
		MaxSealedCommutatorNorm:   maxSealed,
		CKMCapacityNative:         nativePairs > 0,
		Pairs:                     pairs,
		Verdict:                   join(StatusOperatorCapacityAudited, StatusFailedNoNativeNoncommutingPair, StatusTensionNeedTwoNativeNoncommutingOps),
	}
}

func buildPair(a, b SourceCandidate, native, sealed bool) PairAudit {
	norm := 0.0
	eligible := len(a.Operator) > 0 && len(b.Operator) > 0 && len(a.Operator) == len(b.Operator)
	if eligible {
		norm = frobNorm(sub(matMul(a.Operator, b.Operator), matMul(b.Operator, a.Operator)))
	}
	noncommuting := norm > eps
	verdict := StatusFailedNoNativeNoncommutingPair
	if sealed && noncommuting {
		verdict = StatusConditionalCKMCapacityActivated + " | " + StatusFailedTauOrNStillCircular
	}
	if native && noncommuting {
		verdict = StatusConditionalCKMCapacityActivated
	}
	return PairAudit{
		Name:           a.Name + " × " + b.Name,
		Left:           a.Name,
		Right:          b.Name,
		NativePair:     native,
		SealedPair:     sealed,
		Eligible:       eligible,
		CommutatorNorm: norm,
		Noncommuting:   noncommuting,
		CKMCapacity:    noncommuting,
		Reason:         pairReason(native, sealed, noncommuting, eligible),
		Verdict:        verdict,
	}
}

func pairReason(native, sealed, noncommuting, eligible bool) string {
	if !eligible {
		return "operators do not act on the same eligible generation carrier"
	}
	if native && noncommuting {
		return "native noncommuting pair would activate CKM/PMNS capacity"
	}
	if sealed && noncommuting {
		return "sealed stress-test pair is noncommuting, but its carrier is circular/non-native"
	}
	return "pair commutes or lacks eligible generation semantics"
}

func auditModuli(inh Inheritance, sources SourceAudit, ops OperatorAudit) ModuliAudit {
	start := inh.Gate372ChargedModuliDim
	scenarios := []ModuliScenario{
		{Name: "native Gate396 ledger", AssumptionClass: "native", StartingChargedDim: start, ResultingDim: start, Native: true, ThreeDistinctMassesPossible: false, CKMMisalignmentPossible: false, Reason: "no promotable native generation source and no native noncommuting pair", Verdict: StatusFirewallPreserved13Moduli},
		{Name: "contact rational singleton source without flavor functor", AssumptionClass: "native three-object source, wrong semantics", StartingChargedDim: start, ResultingDim: start, Native: true, Failed: true, ThreeDistinctMassesPossible: false, CKMMisalignmentPossible: false, Reason: "three contact spectral idempotents are native but do not act on finite-Dirac generation space", Verdict: join(StatusContactSingletonsFound, StatusFailedContactNoFlavorFunctor, StatusFirewallPreserved13Moduli)},
		{Name: "Fock spatial/color triplet as generation", AssumptionClass: "forbidden semantic relabeling", StartingChargedDim: start, ResultingDim: start, Native: true, Failed: true, ThreeDistinctMassesPossible: false, CKMMisalignmentPossible: false, Reason: "the three Fock spatial modes already carry color/spatial charge semantics", Verdict: join(StatusFailedColorTripletIsColor, StatusFirewallPreserved13Moduli)},
		{Name: "sealed tau_eta or N diagonal ladder", AssumptionClass: "sealed diagonal hierarchy", StartingChargedDim: start, ResultingDim: 9, Conditional: true, Failed: true, ThreeDistinctMassesPossible: true, CKMMisalignmentPossible: false, Reason: "a diagonal three-slot operator can split eigenvalues but cannot create CKM/PMNS misalignment", Verdict: join(StatusFailedTauOrNStillCircular, StatusFailedNoNativeNoncommutingPair)},
		{Name: "sealed diagonal plus cyclic stress test", AssumptionClass: "sealed noncommuting capacity", StartingChargedDim: start, ResultingDim: start, Conditional: true, Failed: true, ThreeDistinctMassesPossible: true, CKMMisalignmentPossible: true, Reason: "noncommuting capacity exists only after importing circular branch labels, so no native moduli quotient is allowed", Verdict: join(StatusConditionalCKMCapacityActivated, StatusFailedTauOrNStillCircular, StatusFailedNoNativeModuliReduction)},
	}
	bestNative, bestCond := start, start
	for _, s := range scenarios {
		if s.Native && s.ResultingDim < bestNative {
			bestNative = s.ResultingDim
		}
		if s.Conditional && s.ResultingDim < bestCond {
			bestCond = s.ResultingDim
		}
	}
	_ = sources
	_ = ops
	return ModuliAudit{true, start, false, bestNative, bestCond, scenarios, join(StatusModuliImpactAudited, StatusFailedNoNativeModuliReduction, StatusFirewallPreserved13Moduli)}
}

func auditFirewall(src SourceAudit, ops OperatorAudit, mod ModuliAudit) FirewallAudit {
	return FirewallAudit{
		Executed:                     true,
		NoMassesImported:             true,
		NoCKMImported:                true,
		NoPMNSImported:               true,
		NoEmpiricalOrderingImported:  true,
		NoManualGenerationAssignment: true,
		NoContactRootsPromoted:       src.PromotableGenerationSourceCount == 0,
		NoColorModesPromoted:         true,
		NoFanoTriplePromoted:         true,
		NoTauOrNPromoted:             true,
		NoNativeFlavorClaimed:        !ops.CKMCapacityNative && src.PromotableGenerationSourceCount == 0,
		NoModuliReductionClaimed:     !mod.NativeReductionBelow13,
		Verdict:                      join(StatusFirewallPreserved13Moduli, StatusFailedThreeSourceNotGeneration, StatusFailedNoNativeNoncommutingPair),
	}
}

func chooseNextGate(src SourceAudit, ops OperatorAudit) NextStep {
	_ = ops
	if src.NativeExactlyThreeSourceCount > 0 && src.PromotableGenerationSourceCount == 0 {
		return NextStep{
			Gate:        397,
			Title:       "Contact Rational Singleton to Finite-Dirac Flavor Functor Sieve",
			Reason:      "Gate 396 found a genuine native three-object source in the rational contact singleton idempotents, but it does not yet act on finite-Dirac generation space.",
			PrimaryTask: "Test whether the three rational contact singleton blocks admit a canonical A_F/J/first-order compatible module action on the finite Dirac/Yukawa edge carrier without choosing quartic branches, color modes, or empirical flavor data.",
		}
	}
	return NextStep{Gate: 397, Title: "Three-Source Selector Search", Reason: "Gate 396 found no usable source", PrimaryTask: "Search for a selector before flavor texture work."}
}

func buildTruth(src SourceAudit, ops OperatorAudit, mod ModuliAudit, next NextStep) string {
	return fmt.Sprintf("Gate 396 moves before texture algebra and asks whether ASHA has any endogenous three-object source at all. It finds native three-object sources, especially the contact rational singleton idempotent blocks and the Fock spatial/color triplet, but neither is a finite-Dirac generation address: the first has contact-root semantics and no flavor functor, while the second is color/spatial charge structure. Fano triples form a sevenfold family requiring a selector; tau_eta and N remain sealed/circular three-slot capacity witnesses. No promotable native generation source and no native noncommuting operator pair are derived, so the charged flavor firewall remains dim=%d. Next: Gate %d — %s.", mod.BestNativeDim, next.Gate, next.Title)
}

func Statuses(a Analysis) []string {
	set := map[string]bool{}
	fields := []string{a.Inheritance.Verdict, a.Sources.Verdict, a.Operators.Verdict, a.Moduli.Verdict, a.Firewall.Verdict}
	for _, c := range a.Sources.Candidates {
		fields = append(fields, c.Verdict)
	}
	for _, p := range a.Operators.Pairs {
		fields = append(fields, p.Verdict)
	}
	for _, s := range a.Moduli.Scenarios {
		fields = append(fields, s.Verdict)
	}
	for _, f := range fields {
		for _, part := range strings.Split(f, " | ") {
			part = strings.TrimSpace(part)
			if part != "" {
				set[part] = true
			}
		}
	}
	out := make([]string, 0, len(set))
	for s := range set {
		out = append(out, s)
	}
	sort.Strings(out)
	return out
}

func join(parts ...string) string {
	clean := []string{}
	seen := map[string]bool{}
	for _, p := range parts {
		for _, q := range strings.Split(p, " | ") {
			q = strings.TrimSpace(q)
			if q == "" || seen[q] {
				continue
			}
			seen[q] = true
			clean = append(clean, q)
		}
	}
	return strings.Join(clean, " | ")
}

func validateSquare(m [][]float64) error {
	if len(m) == 0 {
		return fmt.Errorf("empty matrix")
	}
	n := len(m)
	for i := range m {
		if len(m[i]) != n {
			return fmt.Errorf("row %d has length %d, expected %d", i, len(m[i]), n)
		}
	}
	return nil
}

func identity(n int) [][]float64 {
	m := make([][]float64, n)
	for i := 0; i < n; i++ {
		m[i] = make([]float64, n)
		m[i][i] = 1
	}
	return m
}

func diagFromWeights(w []float64) [][]float64 {
	m := make([][]float64, len(w))
	for i := range w {
		m[i] = make([]float64, len(w))
		m[i][i] = w[i]
	}
	return m
}

func cycle3() [][]float64 { return [][]float64{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}} }

func diagonalSpectrumIfDiagonal(m [][]float64) []float64 {
	if len(m) == 0 {
		return nil
	}
	out := []float64{}
	for i := range m {
		for j := range m[i] {
			if i != j && math.Abs(m[i][j]) > eps {
				return nil
			}
		}
		out = append(out, m[i][i])
	}
	return out
}

func matMul(a, b [][]float64) [][]float64 {
	n := len(a)
	out := make([][]float64, n)
	for i := 0; i < n; i++ {
		out[i] = make([]float64, n)
		for k := 0; k < n; k++ {
			for j := 0; j < n; j++ {
				out[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return out
}

func sub(a, b [][]float64) [][]float64 {
	n := len(a)
	out := make([][]float64, n)
	for i := 0; i < n; i++ {
		out[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			out[i][j] = a[i][j] - b[i][j]
		}
	}
	return out
}

func frobNorm(m [][]float64) float64 {
	s := 0.0
	for i := range m {
		for j := range m[i] {
			s += m[i][j] * m[i][j]
		}
	}
	return math.Sqrt(s)
}

func findSource(candidates []SourceCandidate, name string) SourceCandidate {
	for _, c := range candidates {
		if c.Name == name {
			return c
		}
	}
	return SourceCandidate{}
}

func scenarioByName(scenarios []ModuliScenario, name string) ModuliScenario {
	for _, s := range scenarios {
		if s.Name == name {
			return s
		}
	}
	return ModuliScenario{}
}
