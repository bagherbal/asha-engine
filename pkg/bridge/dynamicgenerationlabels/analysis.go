// Package dynamicgenerationlabels implements Gate 395:
// Representation-Origin Search for Dynamic Generation Labels.
//
// Gate 394 proved that current native support, Morita edge, and one-form
// ledgers broadcast centrally over the trivial generation copy C^3_gen. Gate
// 395 therefore stops treating C^3_gen as an input and audits whether the
// native representation theory of Cl(1,7)/Spin(8) itself creates dynamic,
// noncentral generation labels. The gate is intentionally conservative: the
// 16-dimensional spinor decomposes as 8_s + 8_c, while Spin(8) triality is a
// category-level symmetry among 8_v, 8_s, and 8_c. A category-level triple is
// not promoted to a physical generation carrier unless a native functor into
// finite-Dirac flavor space is explicitly found.
package dynamicgenerationlabels

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
)

const (
	AuditID = "GATE395-REPRESENTATION-ORIGIN-DYNAMIC-GENERATION-LABELS"

	StatusGate394Inherited = "CONDITIONAL_SUPPORT_GATE394_CENTRALITY_FIREWALL_INHERITED"
	StatusGate393Inherited = "CONDITIONAL_SUPPORT_GATE393_TRIALITY_DOMAIN_OBSTRUCTION_INHERITED"
	StatusGate372Inherited = "CONDITIONAL_SUPPORT_GATE372_THIRTEEN_MODULI_FIREWALL_INHERITED"
	StatusGate247Inherited = "CONDITIONAL_SUPPORT_GATE247_TRIALITY_FUNCTOR_OBSTRUCTION_INHERITED"

	StatusSpinorDecompositionAudited      = "CONDITIONAL_SUPPORT_CL17_SPINOR_DECOMPOSITION_AUDITED"
	StatusTrialityCategoryAudited         = "CONDITIONAL_SUPPORT_SPIN8_TRIALITY_CATEGORY_AUDITED"
	StatusDynamicLabelSieveAudited        = "CONDITIONAL_SUPPORT_DYNAMIC_GENERATION_LABEL_SIEVE_AUDITED"
	StatusOperatorCapacityAudited         = "CONDITIONAL_SUPPORT_DYNAMIC_OPERATOR_CAPACITY_AUDITED"
	StatusModuliImpactAudited             = "CONDITIONAL_SUPPORT_DYNAMIC_LABEL_MODULI_IMPACT_AUDITED"
	StatusConditionalTrialityArena        = "CONDITIONAL_SUPPORT_TRIALITY_REPRESENTATION_ARENA"
	StatusConditionalSealedBranchCapacity = "CONDITIONAL_SUPPORT_SEALED_BRANCH_OPERATOR_CAPACITY"

	StatusTensionSixteenSplitsTwoNotThree        = "CONDITIONAL_TENSION_CL17_SPINOR_SPLIT_GIVES_TWO_CHIRAL_HALVES_NOT_THREE_GENERATIONS"
	StatusTensionVectorRepMissingFromSpinorSplit = "CONDITIONAL_TENSION_8V_IS_NOT_CONTAINED_IN_NATIVE_SPINOR_SPLIT"
	StatusTensionTrialityCategoryNotFlavorSpace  = "CONDITIONAL_TENSION_TRIALITY_PERMUTES_REPRESENTATION_TYPES_NOT_GENERATION_COPIES"
	StatusTensionNeedFunctorToFiniteDirac        = "CONDITIONAL_TENSION_NEED_FUNCTOR_FROM_TRIALITY_CATEGORY_TO_FINITE_DIRAC_FLAVOR_SPACE"
	StatusTensionSealedCapacityNotNative         = "CONDITIONAL_TENSION_SEALED_BRANCH_CAPACITY_DOES_NOT_REWRITE_NATIVE_FIREWALL"
	StatusTensionNeedTwoNoncommutingOperators    = "CONDITIONAL_TENSION_CKM_PMNS_REQUIRES_TWO_NATIVE_NONCOMMUTING_TEXTURE_OPERATORS"

	StatusVerifiedDynamicGenerationLabels    = "VERIFIED_DYNAMIC_GENERATION_LABELS_DERIVED"
	StatusConditionalDynamicGenerationLabels = "CONDITIONAL_SUPPORT_DYNAMIC_GENERATION_LABELS_DERIVED"
	StatusConditionalCKMCapacity             = "CONDITIONAL_SUPPORT_CKM_MIXING_CAPACITY_ACTIVATED"
	StatusFailedSpinorTwoNotThree            = "FAILED_ROUTE_SPINOR_DECOMPOSITION_IS_TWO_SECTOR_NOT_THREE_GENERATION"
	StatusFailedTrialityOnlyCategory         = "FAILED_ROUTE_TRIALITY_IS_REPRESENTATION_CATEGORY_NOT_GENERATION_CARRIER"
	StatusFailedNoNativeDynamicLabels        = "FAILED_ROUTE_NO_NATIVE_DYNAMIC_GENERATION_LABELS"
	StatusFailedNoNativeNoncommutingPair     = "FAILED_ROUTE_NO_NATIVE_NONCOMMUTING_TEXTURE_PAIR"
	StatusFailedNoModuliReduction            = "FAILED_ROUTE_NO_NATIVE_MODULI_REDUCTION"
	StatusFirewallPreserved13Moduli          = "FIREWALL_PRESERVED_13_MODULI"
)

const eps = 1e-10

type Inheritance struct {
	Executed                         bool
	Gate394CentralityFirewall        bool
	Gate394NativeNoncentralOperators int
	Gate394NativeNoncommutingPairs   int
	Gate393DomainAdmitted            bool
	Gate247TrialityFunctorDerived    bool
	Gate372ChargedModuliDim          int
	NoEmpiricalFlavorValuesImported  bool
	Verdict                          string
}

type SpinorDecompositionAudit struct {
	Executed                                 bool
	CliffordAlgebra                          string
	FullSpinorRealDimension                  int
	ChiralSplit                              []int
	NativeChiralSectorCount                  int
	HasThreeNativeSectors                    bool
	HasVectorRepresentationInsideSpinorSplit bool
	GenerationLabelsDerived                  bool
	ChiralityOperatorSpectrum                []float64
	Verdict                                  string
}

type TrialityAudit struct {
	Executed                         bool
	RepresentationTypes              []string
	AutomorphismGroup                string
	CategoryLevelTriple              bool
	ActsOnPhysicalGenerationCopies   bool
	ExplicitThetaOnFiniteDiracFlavor bool
	NativeFunctorToC3Gen             bool
	VectorRepSuppliedNatively        bool
	Verdict                          string
}

type LabelCandidate struct {
	Name                         string
	Source                       string
	Native                       bool
	Sealed                       bool
	Circular                     bool
	Dimension                    int
	SectorCount                  int
	Operator                     [][]float64
	Central                      bool
	NonCentral                   bool
	DiagonalOnly                 bool
	Mixing                       bool
	Spectrum                     []float64
	CommutantDimension           int
	GenerationLabelsDerived      bool
	CompatibleWithFiniteDirac    bool
	CompatibleWithJ              bool
	CompatibleWithFirstOrder     bool
	CompatibleWithHyperchargeSU2 bool
	Reason                       string
	Verdict                      string
}

type DynamicLabelAudit struct {
	Executed                   bool
	Candidates                 []LabelCandidate
	NativeCandidateCount       int
	NativeGenerationLabelCount int
	NativeNoncentralCount      int
	SealedNoncentralCount      int
	Verdict                    string
}

type PairAudit struct {
	Name           string
	Left           string
	Right          string
	NativePair     bool
	SealedPair     bool
	CommutatorNorm float64
	Noncommuting   bool
	CKMCapacity    bool
	Reason         string
	Verdict        string
}

type OperatorCapacityAudit struct {
	Executed                  bool
	NativeOperators           int
	NativeGenerationOperators int
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
	NoTrialityLabelsPromoted     bool
	NoNPromoted                  bool
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
	Spinor      SpinorDecompositionAudit
	Triality    TrialityAudit
	Labels      DynamicLabelAudit
	Operators   OperatorCapacityAudit
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
	spinor := auditSpinorDecomposition()
	triality := auditTrialityCategory(spinor)
	labels, err := auditDynamicLabels(spinor, triality)
	if err != nil {
		return Analysis{}, err
	}
	operators := auditOperatorCapacity(labels)
	moduli := auditModuli(inheritance, spinor, triality, labels, operators)
	firewall := auditFirewall(labels, operators, moduli)
	next := chooseNextGate(spinor, triality, labels, operators)
	truth := buildTruth(spinor, triality, labels, operators, moduli, next)
	return Analysis{inheritance, spinor, triality, labels, operators, moduli, firewall, next, truth}, nil
}

func inheritPreviousGates() Inheritance {
	return Inheritance{
		Executed:                         true,
		Gate394CentralityFirewall:        true,
		Gate394NativeNoncentralOperators: 0,
		Gate394NativeNoncommutingPairs:   0,
		Gate393DomainAdmitted:            false,
		Gate247TrialityFunctorDerived:    false,
		Gate372ChargedModuliDim:          13,
		NoEmpiricalFlavorValuesImported:  true,
		Verdict:                          join(StatusGate394Inherited, StatusGate393Inherited, StatusGate247Inherited, StatusGate372Inherited),
	}
}

func auditSpinorDecomposition() SpinorDecompositionAudit {
	chirality := append(repeat(1, 8), repeat(-1, 8)...)
	return SpinorDecompositionAudit{
		Executed:                                 true,
		CliffordAlgebra:                          "Cℓ(1,7) full real spinor S has dim_R=16; even/chiral Spin(8) half-spinors split S=S+⊕S- with dimensions 8+8",
		FullSpinorRealDimension:                  16,
		ChiralSplit:                              []int{8, 8},
		NativeChiralSectorCount:                  2,
		HasThreeNativeSectors:                    false,
		HasVectorRepresentationInsideSpinorSplit: false,
		GenerationLabelsDerived:                  false,
		ChiralityOperatorSpectrum:                toFloat64(chirality),
		Verdict:                                  join(StatusSpinorDecompositionAudited, StatusTensionSixteenSplitsTwoNotThree, StatusFailedSpinorTwoNotThree),
	}
}

func auditTrialityCategory(spinor SpinorDecompositionAudit) TrialityAudit {
	return TrialityAudit{
		Executed:                         true,
		RepresentationTypes:              []string{"8_v vector", "8_s left half-spinor", "8_c right half-spinor"},
		AutomorphismGroup:                "Out(Spin(8)) ≅ S3",
		CategoryLevelTriple:              true,
		ActsOnPhysicalGenerationCopies:   false,
		ExplicitThetaOnFiniteDiracFlavor: false,
		NativeFunctorToC3Gen:             false,
		VectorRepSuppliedNatively:        spinor.HasVectorRepresentationInsideSpinorSplit,
		Verdict:                          join(StatusTrialityCategoryAudited, StatusConditionalTrialityArena, StatusTensionVectorRepMissingFromSpinorSplit, StatusTensionTrialityCategoryNotFlavorSpace, StatusFailedTrialityOnlyCategory),
	}
}

func auditDynamicLabels(spinor SpinorDecompositionAudit, triality TrialityAudit) (DynamicLabelAudit, error) {
	labels := []LabelCandidate{
		newLabelCandidate("spinor chirality split", "native Cℓ(1,7) spinor decomposition S=8_s⊕8_c", true, false, false, 16, spinor.NativeChiralSectorCount, diagFromWeights(toFloat64(append(repeat(1, 8), repeat(-1, 8)...))), "native and meaningful for chirality, but it has two sectors, not three generation labels"),
		newLabelCandidate("triality representation-type triple", "category-level labels {8_v,8_s,8_c}", false, true, true, 24, 3, cycle3(), "triality gives a threefold representation arena only after adjoining 8_v and treating representation types as labels"),
		newLabelCandidate("sealed branch number operator", "N_branch=diag(0,1,2) on {8_v,8_s,8_c} labels", false, true, true, 3, 3, diagFromWeights([]float64{0, 1, 2}), "hierarchy-capable only as a sealed branch-label operator; not derived inside finite Dirac flavor space"),
		newLabelCandidate("finite-Dirac generation broadcast", "Gate-394 native Morita/one-form generation lift", true, false, false, 3, 1, identity3(), "current native finite-Dirac operators still factor through I3"),
	}
	for i := range labels {
		if err := fillLabelCandidate(&labels[i]); err != nil {
			return DynamicLabelAudit{}, err
		}
	}
	native, nativeGen, nativeNoncentral, sealedNoncentral := 0, 0, 0, 0
	for _, l := range labels {
		if l.Native {
			native++
		}
		if l.Native && l.GenerationLabelsDerived {
			nativeGen++
		}
		if l.Native && l.NonCentral {
			nativeNoncentral++
		}
		if l.Sealed && l.NonCentral {
			sealedNoncentral++
		}
	}
	verdict := join(StatusDynamicLabelSieveAudited, StatusFailedNoNativeDynamicLabels, StatusTensionNeedFunctorToFiniteDirac, StatusTensionSealedCapacityNotNative)
	if nativeGen > 0 {
		verdict = join(StatusDynamicLabelSieveAudited, StatusVerifiedDynamicGenerationLabels)
	}
	_ = triality
	return DynamicLabelAudit{true, labels, native, nativeGen, nativeNoncentral, sealedNoncentral, verdict}, nil
}

func newLabelCandidate(name, source string, native, sealed, circular bool, dim, sectors int, op [][]float64, reason string) LabelCandidate {
	return LabelCandidate{
		Name:                         name,
		Source:                       source,
		Native:                       native,
		Sealed:                       sealed,
		Circular:                     circular,
		Dimension:                    dim,
		SectorCount:                  sectors,
		Operator:                     op,
		CompatibleWithFiniteDirac:    native && sectors == 3,
		CompatibleWithJ:              true,
		CompatibleWithFirstOrder:     true,
		CompatibleWithHyperchargeSU2: true,
		Reason:                       reason,
	}
}

func fillLabelCandidate(c *LabelCandidate) error {
	if len(c.Operator) == 0 || len(c.Operator) != len(c.Operator[0]) {
		return fmt.Errorf("candidate %s operator must be square", c.Name)
	}
	n := len(c.Operator)
	tr := trace(c.Operator) / float64(n)
	centralResidual := frob(sub(c.Operator, scale(identity(n), tr)))
	c.Central = centralResidual < eps
	c.NonCentral = !c.Central
	c.DiagonalOnly = isDiagonal(c.Operator, eps)
	c.Mixing = c.NonCentral && !c.DiagonalOnly
	c.Spectrum = sortedDiagonalOrProxy(c.Operator)
	c.CommutantDimension = commutantDimBySpectrum(c.Spectrum)
	c.GenerationLabelsDerived = c.Native && c.SectorCount == 3 && c.NonCentral && c.CompatibleWithFiniteDirac && !c.Circular
	switch {
	case c.GenerationLabelsDerived:
		c.Verdict = StatusVerifiedDynamicGenerationLabels
	case c.Native && c.SectorCount == 2:
		c.Verdict = join(StatusFailedSpinorTwoNotThree, StatusTensionSixteenSplitsTwoNotThree)
	case c.Native && c.Central:
		c.Verdict = join(StatusFailedNoNativeDynamicLabels, StatusGate394Inherited)
	case c.Sealed && c.NonCentral && c.Mixing:
		c.Verdict = join(StatusConditionalSealedBranchCapacity, StatusConditionalCKMCapacity, StatusTensionSealedCapacityNotNative)
	case c.Sealed && c.NonCentral && c.DiagonalOnly:
		c.Verdict = join(StatusConditionalSealedBranchCapacity, StatusTensionSealedCapacityNotNative)
	default:
		c.Verdict = StatusDynamicLabelSieveAudited
	}
	return nil
}

func auditOperatorCapacity(labels DynamicLabelAudit) OperatorCapacityAudit {
	var pairs []PairAudit
	maxNative, maxSealed := 0.0, 0.0
	nativePairs, sealedPairs := 0, 0
	for i := 0; i < len(labels.Candidates); i++ {
		for j := i + 1; j < len(labels.Candidates); j++ {
			a, b := labels.Candidates[i], labels.Candidates[j]
			if len(a.Operator) != len(b.Operator) {
				continue
			}
			cn := frob(commutator(a.Operator, b.Operator))
			non := cn > eps
			nativePair := a.Native && b.Native
			sealedPair := (a.Sealed || b.Sealed) && non
			ckm := nativePair && non && a.GenerationLabelsDerived && b.GenerationLabelsDerived
			verdict := StatusFailedNoNativeNoncommutingPair
			reason := "native pair is commuting, dimension-mismatched, or does not produce generation labels"
			if sealedPair {
				verdict = join(StatusConditionalSealedBranchCapacity, StatusConditionalCKMCapacity, StatusTensionSealedCapacityNotNative)
				reason = "noncommutation appears only after sealed/circular branch-label operators are admitted"
				sealedPairs++
				if cn > maxSealed {
					maxSealed = cn
				}
			}
			if nativePair {
				if cn > maxNative {
					maxNative = cn
				}
				if non {
					nativePairs++
				}
			}
			pairs = append(pairs, PairAudit{
				Name:           a.Name + " :: " + b.Name,
				Left:           a.Name,
				Right:          b.Name,
				NativePair:     nativePair,
				SealedPair:     sealedPair,
				CommutatorNorm: cn,
				Noncommuting:   non,
				CKMCapacity:    ckm,
				Reason:         reason,
				Verdict:        verdict,
			})
		}
	}
	nativeOps, nativeGen, nativeNoncentral := 0, 0, 0
	for _, c := range labels.Candidates {
		if c.Native {
			nativeOps++
		}
		if c.Native && c.GenerationLabelsDerived {
			nativeGen++
		}
		if c.Native && c.NonCentral {
			nativeNoncentral++
		}
	}
	verdict := join(StatusOperatorCapacityAudited, StatusFailedNoNativeNoncommutingPair, StatusTensionNeedTwoNoncommutingOperators)
	if nativePairs > 0 && nativeGen >= 2 {
		verdict = join(StatusOperatorCapacityAudited, StatusConditionalCKMCapacity)
	}
	return OperatorCapacityAudit{true, nativeOps, nativeGen, nativeNoncentral, nativePairs, sealedPairs, maxNative, maxSealed, false, pairs, verdict}
}

func auditModuli(inh Inheritance, spinor SpinorDecompositionAudit, tr TrialityAudit, labels DynamicLabelAudit, ops OperatorCapacityAudit) ModuliAudit {
	start := inh.Gate372ChargedModuliDim
	scenarios := []ModuliScenario{
		{"native Cℓ(1,7) spinor chirality split", "native 8_s⊕8_c two-sector representation", start, start, true, false, true, false, false, "chirality is important but does not create three generation labels", join(StatusFailedSpinorTwoNotThree, StatusFirewallPreserved13Moduli)},
		{"triality representation category", "sealed {8_v,8_s,8_c} branch labels", start, 9, false, true, true, true, false, "branch labels can be stress-tested, but no functor to finite-Dirac flavor space is native", join(StatusConditionalTrialityArena, StatusFailedTrialityOnlyCategory, StatusTensionSealedCapacityNotNative)},
		{"sealed branch N plus triality cycle", "sealed noncommuting branch operators", start, 9, false, true, true, true, false, "noncommuting capacity exists only after circular branch-label insertion", join(StatusConditionalSealedBranchCapacity, StatusConditionalCKMCapacity, StatusTensionSealedCapacityNotNative)},
		{"native dynamic generation labels", "missing prerequisite", start, start, false, false, true, false, false, "no native dynamic labels or two native noncommuting texture operators were derived", join(StatusFailedNoNativeDynamicLabels, StatusFailedNoNativeNoncommutingPair, StatusFirewallPreserved13Moduli)},
	}
	bestNative := start
	bestConditional := start
	for _, s := range scenarios {
		if s.Native && s.ResultingDim < bestNative {
			bestNative = s.ResultingDim
		}
		if s.Conditional && s.ResultingDim < bestConditional {
			bestConditional = s.ResultingDim
		}
	}
	_ = spinor
	_ = tr
	_ = labels
	_ = ops
	return ModuliAudit{true, start, false, bestNative, bestConditional, scenarios, join(StatusModuliImpactAudited, StatusFailedNoModuliReduction, StatusFirewallPreserved13Moduli)}
}

func auditFirewall(labels DynamicLabelAudit, ops OperatorCapacityAudit, mod ModuliAudit) FirewallAudit {
	return FirewallAudit{
		Executed:                     true,
		NoMassesImported:             true,
		NoCKMImported:                true,
		NoPMNSImported:               true,
		NoEmpiricalOrderingImported:  true,
		NoManualGenerationAssignment: true,
		NoTrialityLabelsPromoted:     true,
		NoNPromoted:                  true,
		NoNativeFlavorClaimed:        labels.NativeGenerationLabelCount == 0 && !ops.CKMCapacityNative,
		NoModuliReductionClaimed:     !mod.NativeReductionBelow13,
		Verdict:                      join(StatusFirewallPreserved13Moduli, StatusFailedNoNativeDynamicLabels, StatusFailedNoNativeNoncommutingPair),
	}
}

func chooseNextGate(spinor SpinorDecompositionAudit, tr TrialityAudit, labels DynamicLabelAudit, ops OperatorCapacityAudit) NextStep {
	_ = spinor
	_ = tr
	_ = labels
	_ = ops
	return NextStep{
		Gate:        396,
		Title:       "Endogenous Three-Object Source Search beyond Spinor Chirality",
		Reason:      "Gate 395 shows that native spinor decomposition gives two chiral halves, while triality gives a category-level triple only after adjoining 8_v. The missing theorem is not another texture calculation; it is an endogenous three-object source tied to finite Dirac flavor space.",
		PrimaryTask: "Audit whether primitive idempotents, minimal left ideals, octonionic/Fano triples, or modular/KMS sectors derive exactly three addressable, noncentral generation labels compatible with A_F, J, first-order, and electroweak charges.",
	}
}

func buildTruth(spinor SpinorDecompositionAudit, tr TrialityAudit, labels DynamicLabelAudit, ops OperatorCapacityAudit, mod ModuliAudit, next NextStep) string {
	_ = tr
	_ = labels
	_ = ops
	_ = mod
	return fmt.Sprintf("Gate 395 audits the representation-origin hypothesis and rejects the direct claim that the Cℓ(1,7) spinor split dynamically derives three generations. The native spinor decomposition is %d=%d+%d, giving two chiral half-spinor sectors, not three generation labels. Spin(8) triality supplies a threefold representation arena {8_v,8_s,8_c}, but the vector representation is not contained inside the spinor split and no native functor to finite-Dirac flavor space is derived. Sealed branch operators can be made noncentral and even noncommuting, but they remain circular label actions. The Gate-372 charged 13-moduli firewall is preserved. Next: Gate %d — %s.", spinor.FullSpinorRealDimension, spinor.ChiralSplit[0], spinor.ChiralSplit[1], next.Gate, next.Title)
}

func repeat(v, n int) []int {
	out := make([]int, n)
	for i := range out {
		out[i] = v
	}
	return out
}
func toFloat64(xs []int) []float64 {
	out := make([]float64, len(xs))
	for i, v := range xs {
		out[i] = float64(v)
	}
	return out
}
func identity3() [][]float64 { return identity(3) }
func identity(n int) [][]float64 {
	m := make([][]float64, n)
	for i := 0; i < n; i++ {
		m[i] = make([]float64, n)
		m[i][i] = 1
	}
	return m
}
func cycle3() [][]float64 { return [][]float64{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}} }
func diagFromWeights(w []float64) [][]float64 {
	m := make([][]float64, len(w))
	for i := range w {
		m[i] = make([]float64, len(w))
		m[i][i] = w[i]
	}
	return m
}
func trace(a [][]float64) float64 {
	s := 0.0
	for i := range a {
		s += a[i][i]
	}
	return s
}
func scale(a [][]float64, k float64) [][]float64 {
	r := make([][]float64, len(a))
	for i := range a {
		r[i] = make([]float64, len(a[i]))
		for j := range a[i] {
			r[i][j] = k * a[i][j]
		}
	}
	return r
}
func sub(a, b [][]float64) [][]float64 {
	r := make([][]float64, len(a))
	for i := range a {
		r[i] = make([]float64, len(a[i]))
		for j := range a[i] {
			r[i][j] = a[i][j] - b[i][j]
		}
	}
	return r
}
func mul(a, b [][]float64) [][]float64 {
	n := len(a)
	r := make([][]float64, n)
	for i := 0; i < n; i++ {
		r[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			s := 0.0
			for k := 0; k < n; k++ {
				s += a[i][k] * b[k][j]
			}
			r[i][j] = s
		}
	}
	return r
}
func commutator(a, b [][]float64) [][]float64 { return sub(mul(a, b), mul(b, a)) }
func frob(a [][]float64) float64 {
	s := 0.0
	for i := range a {
		for j := range a[i] {
			s += a[i][j] * a[i][j]
		}
	}
	return math.Sqrt(s)
}
func isDiagonal(a [][]float64, tol float64) bool {
	for i := range a {
		for j := range a[i] {
			if i != j && math.Abs(a[i][j]) > tol {
				return false
			}
		}
	}
	return true
}
func sortedDiagonalOrProxy(a [][]float64) []float64 {
	out := make([]float64, len(a))
	if isDiagonal(a, eps) {
		for i := range a {
			out[i] = a[i][i]
		}
	} else {
		for i := range a {
			s := 0.0
			for j := range a[i] {
				s += a[i][j] * a[i][j]
			}
			out[i] = math.Sqrt(s)
		}
	}
	sort.Float64s(out)
	return out
}
func commutantDimBySpectrum(spec []float64) int {
	if len(spec) == 0 {
		return 0
	}
	groups := []int{}
	used := make([]bool, len(spec))
	for i, x := range spec {
		if used[i] {
			continue
		}
		count := 0
		for j, y := range spec {
			if math.Abs(x-y) < 1e-8 {
				used[j] = true
				count++
			}
		}
		groups = append(groups, count)
	}
	dim := 0
	for _, g := range groups {
		dim += g * g
	}
	return dim
}
func join(parts ...string) string {
	nonempty := []string{}
	for _, p := range parts {
		if p != "" {
			nonempty = append(nonempty, p)
		}
	}
	return strings.Join(nonempty, ";")
}
func findLabel(cs []LabelCandidate, name string) LabelCandidate {
	for _, c := range cs {
		if c.Name == name {
			return c
		}
	}
	return LabelCandidate{Name: name, Verdict: "MISSING"}
}
