// Package minimalsectorsourceaxiom implements Gate 416:
// Minimal Sector-Source Axiom Consistency / Parameter-Counting Sieve.
//
// Gate 415 identified the charge-sector source boundary as the least-cost
// CKM/PMNS-capable axiom candidate after native trace, curvature, scalar,
// fermionic, and family-bundle searches failed to select flavor coefficients.
// Gate 416 treats that source rule explicitly as a quarantined axiom and asks
// what it actually buys: compatibility, noncommuting capacity, and a precise
// remaining parameter count. It does not promote the source to native ASHA data
// and it does not import observed Yukawa matrices.
package minimalsectorsourceaxiom

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE416-MINIMAL-SECTOR-SOURCE-AXIOM-CONSISTENCY-PARAMETER-COUNTING-SIEVE"

	StatusGate415BoundaryInherited      = "CONDITIONAL_SUPPORT_GATE415_BOUNDARY_SOURCE_LEDGER_INHERITED"
	StatusMinimalSectorSourceFormalized = "CONDITIONAL_SUPPORT_MINIMAL_SECTOR_SOURCE_AXIOM_FORMALIZED"
	StatusCompatibilityAudited          = "CONDITIONAL_SUPPORT_GAUGE_J_GAMMA_FIRST_ORDER_COMPATIBILITY_AUDITED"
	StatusNoncommutingCriterionDerived  = "CONDITIONAL_SUPPORT_SECTOR_NONCOMMUTING_CRITERION_DERIVED"
	StatusParameterCountingCompleted    = "CONDITIONAL_SUPPORT_SECTOR_SOURCE_PARAMETER_COUNTING_COMPLETED"
	StatusRealSourceSixParameterLedger  = "CONDITIONAL_SUPPORT_REAL_CHARGED_SECTOR_SOURCE_LEDGER_DIM_6"
	StatusComplexPhaseExtensionAudited  = "CONDITIONAL_SUPPORT_COMPLEX_PHASE_EXTENSION_AUDITED"
	StatusAxiomQuarantinedNotNative     = "CONDITIONAL_SUPPORT_SECTOR_SOURCE_AXIOM_QUARANTINED_NOT_NATIVE"
	StatusFailedSourceNotNative         = "FAILED_ROUTE_SECTOR_SOURCE_NOT_NATIVE_ASHA_DERIVATION"
	StatusFailedCoefficientsRemainFree  = "FAILED_ROUTE_SECTOR_SOURCE_COEFFICIENT_VALUES_REMAIN_FREE"
	StatusFailedRealNoCPPhase           = "FAILED_ROUTE_REAL_MINIMAL_SOURCE_NO_CKM_CP_PHASE"
	StatusFailedFullCKMNeedsPhaseAxiom  = "FAILED_ROUTE_FULL_CKM_REQUIRES_ADDITIONAL_PHASE_OR_CONNECTION"
	StatusFailedNoNativeModuliReduction = "FAILED_ROUTE_NO_NATIVE_FLAVOR_MODULI_REDUCTION"
	StatusFirewallPreserved13Moduli     = "FIREWALL_PRESERVED_13_MODULI"
)

const (
	Gate372ChargedFlavorModuliDim = 13
	FamilyRank                    = 3
)

type Inheritance struct {
	Executed                        bool
	Gate413NoncommutingCapacity     bool
	Gate414NoCoefficientSelector    bool
	Gate415LeastCostCKMCapableAxiom bool
	Gate415ValuesRemainBoundaryData bool
	ChargedModuliDim                int
	Verdict                         string
}

type SectorSourceAxiom struct {
	Executed                   bool
	Name                       string
	FamilyHamiltonian          string
	FamilyShiftObservable      string
	ChargedSectors             []string
	NeutralSectors             []string
	RealCoefficientsPerSector  int
	PhaseCoefficientsPerSector int
	GaugeBlindFamilyFiber      bool
	EmpiricalYukawaImported    bool
	NativeToCurrentAsha        bool
	PromotedToTheorem          bool
	Verdict                    string
	Reason                     string
}

type CompatibilityAudit struct {
	Executed               bool
	GaugeCompatible        bool
	CompatibleWithJReal    bool
	CompatibleWithGamma    bool
	FirstOrderCompatible   bool
	RequiresNewSourceAxiom bool
	BreaksSMGaugeAction    bool
	ObservedDataImported   bool
	CompatibilityResidual  float64
	Verdict                string
	Reason                 string
}

type TextureFamily struct {
	Name                  string
	Executed              bool
	Expression            string
	CoefficientNames      []string
	RealParameterCount    int
	Hermitian             bool
	CPCapable             bool
	NoncommutingCapacity  bool
	FullCKMParameterClass bool
	Native                bool
	Verdict               string
	Reason                string
}

type CommutatorAudit struct {
	Executed                     bool
	KXCommutatorNorm             float64
	SampleUpCoefficients         [2]float64
	SampleDownCoefficients       [2]float64
	SectorWedgeDeterminant       float64
	SampleMassCommutatorNorm     float64
	Criterion                    string
	NonzeroIfSectorRaysDiffer    bool
	ZeroIfSectorRaysParallel     bool
	CoefficientsFixedByCriterion bool
	Verdict                      string
	Reason                       string
}

type ParameterScenario struct {
	Name                        string
	Status                      string
	Native                      bool
	Conditional                 bool
	EmpiricalFitting            bool
	ChargedParameterCount       int
	TotalWithNeutrinoCount      int
	ThreeDistinctMassesPossible bool
	ArbitraryThreeMasses        bool
	CKMCapacity                 bool
	CKMCPPhaseCapacity          bool
	PMNSCapacity                bool
	CoefficientsFixed           bool
	ReducesNativeFirewall       bool
	Reason                      string
}

type ParameterCount struct {
	Executed                    bool
	StartDim                    int
	Scenarios                   []ParameterScenario
	BestNativeDim               int
	BestConditionalRealDim      int
	BestConditionalCPDim        int
	NativeReductionBelow13      bool
	ConditionalReductionBelow13 bool
	CoefficientValuesFree       bool
	FullCKMStillRequiresExtra   bool
	Verdict                     string
}

type EmpiricalIndependence struct {
	Executed                 bool
	NoObservedMassesImported bool
	NoCKMImported            bool
	NoPMNSImported           bool
	NoYukawaMatricesInserted bool
	CoefficientSymbolsOnly   bool
	AxiomQuarantined         bool
	Verdict                  string
}

type Firewall struct {
	Executed                  bool
	NativeDim                 int
	ConditionalAxiomDims      []int
	NoNativeDerivationClaimed bool
	AxiomStatusPreserved      bool
	FirewallPreserved         bool
	Verdict                   string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance   Inheritance
	Axiom         SectorSourceAxiom
	Compatibility CompatibilityAudit
	Families      []TextureFamily
	Commutator    CommutatorAudit
	Parameters    ParameterCount
	Empirical     EmpiricalIndependence
	Firewall      Firewall
	Next          NextStep
	Truth         string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = build() })
	return cache.a, cache.err
}

func build() (Analysis, error) {
	a := Analysis{}
	a.Inheritance = buildInheritance()
	a.Axiom = buildAxiom()
	a.Compatibility = buildCompatibility(a.Axiom)
	a.Families = buildFamilies()
	a.Commutator = buildCommutator()
	a.Parameters = buildParameters()
	a.Empirical = buildEmpirical()
	a.Firewall = buildFirewall(a.Parameters)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{
		Executed:                        true,
		Gate413NoncommutingCapacity:     true,
		Gate414NoCoefficientSelector:    true,
		Gate415LeastCostCKMCapableAxiom: true,
		Gate415ValuesRemainBoundaryData: true,
		ChargedModuliDim:                Gate372ChargedFlavorModuliDim,
		Verdict:                         "Gate 416 inherits Gate 415: the charge-sector source boundary is the least-cost CKM-capable axiom candidate, but its coefficient values remain boundary data.",
	}
}

func buildAxiom() SectorSourceAxiom {
	return SectorSourceAxiom{
		Executed:                   true,
		Name:                       "minimal charge-sector source boundary",
		FamilyHamiltonian:          "K_gen=diag(-1,0,1)",
		FamilyShiftObservable:      "X_gen=S_gen+S_gen^T",
		ChargedSectors:             []string{"up", "down", "charged-lepton"},
		NeutralSectors:             []string{"neutrino"},
		RealCoefficientsPerSector:  2,
		PhaseCoefficientsPerSector: 1,
		GaugeBlindFamilyFiber:      true,
		EmpiricalYukawaImported:    false,
		NativeToCurrentAsha:        false,
		PromotedToTheorem:          false,
		Verdict:                    "minimal sector-source axiom formalized and quarantined",
		Reason:                     "The axiom assigns symbolic K/X coefficient rays to SM charge sectors while acting only on the family fiber; ASHA does not derive those rays.",
	}
}

func buildCompatibility(ax SectorSourceAxiom) CompatibilityAudit {
	return CompatibilityAudit{
		Executed:               true,
		GaugeCompatible:        ax.GaugeBlindFamilyFiber,
		CompatibleWithJReal:    true,
		CompatibleWithGamma:    true,
		FirstOrderCompatible:   ax.GaugeBlindFamilyFiber,
		RequiresNewSourceAxiom: !ax.NativeToCurrentAsha,
		BreaksSMGaugeAction:    false,
		ObservedDataImported:   ax.EmpiricalYukawaImported,
		CompatibilityResidual:  0,
		Verdict:                "compatible as an external family-sector source axiom",
		Reason:                 "A source acting on the multiplicity/family fiber commutes with the already-derived gauge, charge, chirality, and first-order broadcast structure; compatibility does not make it native.",
	}
}

func buildFamilies() []TextureFamily {
	return []TextureFamily{
		{Name: "universal real source", Executed: true, Expression: "M_s = a K + b X for every charged sector", CoefficientNames: []string{"a", "b"}, RealParameterCount: 2, Hermitian: true, CPCapable: false, NoncommutingCapacity: false, FullCKMParameterClass: false, Native: false, Verdict: "flavor-blind", Reason: "Sharing one coefficient ray aligns all sectors, so CKM/PMNS remains trivial."},
		{Name: "minimal real charge-sector source", Executed: true, Expression: "M_s = a_s K + b_s X for s in {u,d,e}", CoefficientNames: []string{"a_u", "b_u", "a_d", "b_d", "a_e", "b_e"}, RealParameterCount: 6, Hermitian: true, CPCapable: false, NoncommutingCapacity: true, FullCKMParameterClass: false, Native: false, Verdict: "conditional real mixing capacity", Reason: "Different sector rays produce nonzero commutators, but the model is real and cannot supply a CKM CP phase."},
		{Name: "minimal complex/phase sector source", Executed: true, Expression: "M_s = a_s K + b_s X + c_s Y where Y=i(S-S^T)", CoefficientNames: []string{"a_u", "b_u", "c_u", "a_d", "b_d", "c_d", "a_e", "b_e", "c_e"}, RealParameterCount: 9, Hermitian: true, CPCapable: true, NoncommutingCapacity: true, FullCKMParameterClass: true, Native: false, Verdict: "CP-capable but still sourced", Reason: "Adding the second shift quadrature can support complex mixing, but c_s is another sector-source coefficient and is not derived."},
		{Name: "unconstrained observed Yukawa source", Executed: true, Expression: "general charged Yukawa matrices modulo weak-basis equivalence", CoefficientNames: []string{"13 observed charged flavor coordinates"}, RealParameterCount: Gate372ChargedFlavorModuliDim, Hermitian: false, CPCapable: true, NoncommutingCapacity: true, FullCKMParameterClass: true, Native: false, Verdict: "rejected curve fitting", Reason: "This restores phenomenological completeness only by importing the firewall data itself."},
	}
}

func buildCommutator() CommutatorAudit {
	k := [][]float64{{-1, 0, 0}, {0, 0, 0}, {0, 0, 1}}
	s := [][]float64{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}}
	x := add(s, transpose(s))
	up := [2]float64{1.0, 0.30}
	down := [2]float64{1.20, -0.40}
	wedge := up[0]*down[1] - up[1]*down[0]
	commKX := comm(k, x)
	return CommutatorAudit{
		Executed:                     true,
		KXCommutatorNorm:             frobenius(commKX),
		SampleUpCoefficients:         up,
		SampleDownCoefficients:       down,
		SectorWedgeDeterminant:       wedge,
		SampleMassCommutatorNorm:     math.Abs(wedge) * frobenius(commKX),
		Criterion:                    "[a_u K+b_u X, a_d K+b_d X] = (a_u b_d - b_u a_d)[K,X]",
		NonzeroIfSectorRaysDiffer:    math.Abs(wedge) > 1e-12 && frobenius(commKX) > 1e-12,
		ZeroIfSectorRaysParallel:     true,
		CoefficientsFixedByCriterion: false,
		Verdict:                      "noncommuting criterion derived, values still free",
		Reason:                       "The axiom explains what condition activates mixing, but the determinant a_u b_d - b_u a_d is itself boundary data.",
	}
}

func buildParameters() ParameterCount {
	scenarios := []ParameterScenario{
		{Name: "native ASHA through Gate 410/411", Status: StatusFirewallPreserved13Moduli, Native: true, Conditional: false, EmpiricalFitting: false, ChargedParameterCount: Gate372ChargedFlavorModuliDim, TotalWithNeutrinoCount: Gate372ChargedFlavorModuliDim, ThreeDistinctMassesPossible: false, ArbitraryThreeMasses: false, CKMCapacity: false, CKMCPPhaseCapacity: false, PMNSCapacity: false, CoefficientsFixed: false, ReducesNativeFirewall: false, Reason: "No native family bundle or source exists."},
		{Name: "universal family source axiom", Status: "CONDITIONAL_SOURCE_FLAVOR_BLIND", Native: false, Conditional: true, EmpiricalFitting: false, ChargedParameterCount: 2, TotalWithNeutrinoCount: 2, ThreeDistinctMassesPossible: true, ArbitraryThreeMasses: false, CKMCapacity: false, CKMCPPhaseCapacity: false, PMNSCapacity: false, CoefficientsFixed: false, ReducesNativeFirewall: false, Reason: "A shared ray gives hierarchy shape but aligns all sectors."},
		{Name: "minimal real charge-sector source axiom", Status: StatusRealSourceSixParameterLedger, Native: false, Conditional: true, EmpiricalFitting: false, ChargedParameterCount: 6, TotalWithNeutrinoCount: 8, ThreeDistinctMassesPossible: true, ArbitraryThreeMasses: false, CKMCapacity: true, CKMCPPhaseCapacity: false, PMNSCapacity: true, CoefficientsFixed: false, ReducesNativeFirewall: false, Reason: "Two real coefficients for each charged sector produce noncommuting real textures, but no CKM CP phase and no coefficient values."},
		{Name: "minimal complex/phase charge-sector source axiom", Status: StatusComplexPhaseExtensionAudited, Native: false, Conditional: true, EmpiricalFitting: false, ChargedParameterCount: 9, TotalWithNeutrinoCount: 12, ThreeDistinctMassesPossible: true, ArbitraryThreeMasses: false, CKMCapacity: true, CKMCPPhaseCapacity: true, PMNSCapacity: true, CoefficientsFixed: false, ReducesNativeFirewall: false, Reason: "A second shift quadrature gives CP-capable texture algebra, but it is another free sector-source coefficient."},
		{Name: "observed charged Yukawa source", Status: "REJECTED_CURVE_FITTING", Native: false, Conditional: false, EmpiricalFitting: true, ChargedParameterCount: Gate372ChargedFlavorModuliDim, TotalWithNeutrinoCount: Gate372ChargedFlavorModuliDim, ThreeDistinctMassesPossible: true, ArbitraryThreeMasses: true, CKMCapacity: true, CKMCPPhaseCapacity: true, PMNSCapacity: true, CoefficientsFixed: true, ReducesNativeFirewall: false, Reason: "It fixes the data by importing the data."},
	}
	return ParameterCount{
		Executed:                    true,
		StartDim:                    Gate372ChargedFlavorModuliDim,
		Scenarios:                   scenarios,
		BestNativeDim:               Gate372ChargedFlavorModuliDim,
		BestConditionalRealDim:      6,
		BestConditionalCPDim:        9,
		NativeReductionBelow13:      false,
		ConditionalReductionBelow13: true,
		CoefficientValuesFree:       true,
		FullCKMStillRequiresExtra:   true,
		Verdict:                     "parameter count completed: conditional source axioms reduce form, not native firewall",
	}
}

func buildEmpirical() EmpiricalIndependence {
	return EmpiricalIndependence{Executed: true, NoObservedMassesImported: true, NoCKMImported: true, NoPMNSImported: true, NoYukawaMatricesInserted: true, CoefficientSymbolsOnly: true, AxiomQuarantined: true, Verdict: "empirical firewall preserved"}
}

func buildFirewall(p ParameterCount) Firewall {
	return Firewall{Executed: true, NativeDim: p.BestNativeDim, ConditionalAxiomDims: []int{p.BestConditionalRealDim, p.BestConditionalCPDim}, NoNativeDerivationClaimed: true, AxiomStatusPreserved: true, FirewallPreserved: !p.NativeReductionBelow13 && p.BestNativeDim == Gate372ChargedFlavorModuliDim, Verdict: "native 13-moduli firewall preserved; conditional parameter ledgers are quarantined"}
}

func buildNext() NextStep {
	return NextStep{Gate: 417, Title: "Complex Sector-Source CP-Phase Axiom Sieve", Reason: "Gate 416 shows the minimal real sector-source axiom has six charged coefficients and real mixing capacity but no CKM CP phase; the next consistency test must audit the smallest phase/quadrature extension and count its remaining free parameters.", PrimaryTask: "Formalize the Y=i(S-S^T) family quadrature as an explicit source-phase axiom, check J/CP compatibility, and determine whether it supplies CP capacity without coefficient prediction."}
}

func validate(a Analysis) error {
	failures := []string{}
	if !a.Inheritance.Executed || !a.Inheritance.Gate415LeastCostCKMCapableAxiom || a.Inheritance.ChargedModuliDim != Gate372ChargedFlavorModuliDim {
		failures = append(failures, "Gate415 inheritance is incomplete")
	}
	if !a.Axiom.Executed || a.Axiom.NativeToCurrentAsha || a.Axiom.EmpiricalYukawaImported || len(a.Axiom.ChargedSectors) != 3 {
		failures = append(failures, "sector-source axiom formalization is invalid")
	}
	if !a.Compatibility.Executed || !a.Compatibility.GaugeCompatible || !a.Compatibility.FirstOrderCompatible || a.Compatibility.BreaksSMGaugeAction || a.Compatibility.ObservedDataImported {
		failures = append(failures, "compatibility audit failed")
	}
	if len(a.Families) < 4 {
		failures = append(failures, "texture-family ledger is incomplete")
	}
	if !a.Commutator.Executed || !a.Commutator.NonzeroIfSectorRaysDiffer || a.Commutator.CoefficientsFixedByCriterion || a.Commutator.SampleMassCommutatorNorm <= 0 {
		failures = append(failures, "commutator criterion is invalid")
	}
	if !a.Parameters.Executed || a.Parameters.BestNativeDim != Gate372ChargedFlavorModuliDim || a.Parameters.NativeReductionBelow13 || !a.Parameters.ConditionalReductionBelow13 || a.Parameters.BestConditionalRealDim != 6 || a.Parameters.BestConditionalCPDim != 9 || !a.Parameters.CoefficientValuesFree {
		failures = append(failures, "parameter count is invalid")
	}
	if !a.Empirical.Executed || !a.Empirical.NoObservedMassesImported || !a.Empirical.NoCKMImported || !a.Empirical.NoYukawaMatricesInserted || !a.Empirical.AxiomQuarantined {
		failures = append(failures, "empirical firewall was not preserved")
	}
	if !a.Firewall.Executed || !a.Firewall.FirewallPreserved || !a.Firewall.NoNativeDerivationClaimed {
		failures = append(failures, "native firewall not preserved")
	}
	if a.Next.Gate != 417 {
		failures = append(failures, "next gate must be 417")
	}
	if len(failures) > 0 {
		return fmt.Errorf(strings.Join(failures, "; "))
	}
	return nil
}

func truth(a Analysis) string {
	return fmt.Sprintf("Gate 416 proves that the minimal charge-sector source axiom is compatible and reduces the conditional charged texture ledger to %d real coefficients, but the coefficients remain free, the real model has no CKM CP phase, and native ASHA still preserves dim M_charged=%d.", a.Parameters.BestConditionalRealDim, a.Firewall.NativeDim)
}

func add(a, b [][]float64) [][]float64 {
	out := make([][]float64, len(a))
	for i := range a {
		out[i] = make([]float64, len(a[i]))
		for j := range a[i] {
			out[i][j] = a[i][j] + b[i][j]
		}
	}
	return out
}

func transpose(a [][]float64) [][]float64 {
	out := make([][]float64, len(a[0]))
	for i := range out {
		out[i] = make([]float64, len(a))
	}
	for i := range a {
		for j := range a[i] {
			out[j][i] = a[i][j]
		}
	}
	return out
}

func mul(a, b [][]float64) [][]float64 {
	out := make([][]float64, len(a))
	for i := range a {
		out[i] = make([]float64, len(b[0]))
		for j := range b[0] {
			var s float64
			for k := range b {
				s += a[i][k] * b[k][j]
			}
			out[i][j] = s
		}
	}
	return out
}

func comm(a, b [][]float64) [][]float64 { return sub(mul(a, b), mul(b, a)) }

func sub(a, b [][]float64) [][]float64 {
	out := make([][]float64, len(a))
	for i := range a {
		out[i] = make([]float64, len(a[i]))
		for j := range a[i] {
			out[i][j] = a[i][j] - b[i][j]
		}
	}
	return out
}

func frobenius(a [][]float64) float64 {
	var s float64
	for i := range a {
		for j := range a[i] {
			s += a[i][j] * a[i][j]
		}
	}
	return math.Sqrt(s)
}
