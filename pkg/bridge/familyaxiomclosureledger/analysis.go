// Package familyaxiomclosureledger implements Gate 418:
// Family-Axiom Closure Ledger / Flavor Frontier Seal.
//
// Gates 411-417 established the honest flavor boundary. Native ASHA data keeps
// the charged flavor moduli sealed at dimension 13. Explicit family axioms can
// add hierarchy, mixing, and CP capacity, but their coefficient values remain
// boundary/environmental data. Gate 418 is therefore a capstone ledger: it does
// not search for another hidden selector; it records the axiom progression,
// parameter compression, empirical firewall, and final flavor frontier seal.
package familyaxiomclosureledger

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE418-FAMILY-AXIOM-CLOSURE-LEDGER-FLAVOR-FRONTIER-SEAL"

	StatusGate417Inherited                  = "CONDITIONAL_SUPPORT_GATE417_CP_PHASE_BOUNDARY_INHERITED"
	StatusAxiomProgressionLedgerCompiled    = "CONDITIONAL_SUPPORT_FAMILY_AXIOM_PROGRESSION_LEDGER_COMPILED"
	StatusParameterReductionSummaryCompiled = "CONDITIONAL_SUPPORT_FLAVOR_PARAMETER_REDUCTION_SUMMARY_COMPILED"
	StatusEnvironmentalSealFormalized       = "FLAVOR_COEFFICIENTS_ENVIRONMENTAL_SEAL"
	StatusProjectFlavorSectorSealedComplete = "PROJECT_FLAVOR_SECTOR_FORMALLY_SEALED_AND_COMPLETE"
	StatusEmpiricalFirewallPreserved        = "CONDITIONAL_SUPPORT_EMPIRICAL_FIREWALL_PRESERVED"
	StatusNoNativeFlavorPrediction          = "FAILED_ROUTE_NO_NATIVE_FLAVOR_COEFFICIENT_PREDICTION"
	StatusNoAxiomPromotedNative             = "FAILED_ROUTE_NO_FAMILY_AXIOM_PROMOTED_TO_NATIVE_THEOREM"
	StatusNineCoefficientsRemainFree        = "FAILED_ROUTE_NINE_SOURCE_COEFFICIENTS_REMAIN_ENVIRONMENTAL"
	StatusCKMPhaseNotPredicted              = "FAILED_ROUTE_CKM_ANGLES_AND_CP_PHASE_NOT_PREDICTED"
	StatusFirewallPreserved13Moduli         = "FIREWALL_PRESERVED_13_MODULI"
)

const (
	Gate372ChargedFlavorModuliDim = 13
	ConditionalComplexChargedDim  = 9
)

type Inheritance struct {
	Executed                     bool
	Gate411AxiomLedgerCompiled   bool
	Gate412HierarchyCapacity     bool
	Gate413MixingCapacity        bool
	Gate414NoCoefficientSelector bool
	Gate415SourceLedgerCompiled  bool
	Gate416RealLedgerDim         int
	Gate416RealNoCP              bool
	Gate417ComplexLedgerDim      int
	Gate417CPCapacity            bool
	Gate417CoefficientValuesFree bool
	ChargedFlavorModuliDim       int
	Verdict                      string
}

type AxiomStep struct {
	Gate                    int
	Name                    string
	Operators               []string
	AddsHierarchy           bool
	AddsRealMixing          bool
	AddsCPCapacity          bool
	GaugeCompatible         bool
	CompatibleWithJGamma    bool
	FirstOrderCompatible    bool
	NativeToCurrentAsha     bool
	PromotedToTheorem       bool
	EmpiricalValuesImported bool
	CoefficientsFixed       bool
	RemainingFreeLedger     int
	Verdict                 string
	Boundary                string
}

type AxiomProgression struct {
	Executed             bool
	Steps                []AxiomStep
	MinimalHierarchyGate int
	MinimalMixingGate    int
	MinimalCPGate        int
	AllAxiomsQuarantined bool
	Verdict              string
}

type ParameterRow struct {
	Name                string
	Status              string
	Native              bool
	Conditional         bool
	EmpiricalFitting    bool
	ChargedCoordinates  int
	HierarchyCapacity   bool
	CKMCapacity         bool
	CPCapacity          bool
	ValuesPredicted     bool
	ValuesEnvironmental bool
	Comments            string
}

type ParameterReduction struct {
	Executed                   bool
	StartDim                   int
	Rows                       []ParameterRow
	ConditionalCompressedDim   int
	NativeDim                  int
	ConditionalCompression     bool
	NativeCompression          bool
	NineCoefficientsSymbolic   bool
	CKMAndPhaseValuesPredicted bool
	Verdict                    string
}

type EnvironmentalSeal struct {
	Executed                   bool
	Name                       string
	Statement                  string
	NativeLawSpaceComplete     bool
	FlavorCapacityComplete     bool
	CoefficientValuesPredicted bool
	CoefficientsEnvironmental  bool
	RequiresHistoricalData     bool
	NoEmpiricalFitting         bool
	NoNativeCollapseClaimed    bool
	Verdict                    string
}

type EmpiricalFirewall struct {
	Executed                 bool
	NoObservedMassesImported bool
	NoCKMImported            bool
	NoPMNSImported           bool
	NoYukawaMatricesInserted bool
	SymbolicCoefficientsOnly bool
	RejectsCurveFitting      bool
	Verdict                  string
}

type FinalSeal struct {
	Executed                      bool
	NativeChargedDim              int
	ConditionalComplexDim         int
	FlavorSectorFormallySealed    bool
	ProjectFlavorCompleteAsLedger bool
	NoNativePredictionClaimed     bool
	FirewallPreserved             bool
	FinalStatus                   string
	Verdict                       string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Progression AxiomProgression
	Parameters  ParameterReduction
	Seal        EnvironmentalSeal
	Empirical   EmpiricalFirewall
	Final       FinalSeal
	Next        NextStep
	Truth       string
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
	a.Progression = buildProgression()
	a.Parameters = buildParameters(a.Progression)
	a.Seal = buildEnvironmentalSeal(a.Parameters)
	a.Empirical = buildEmpirical()
	a.Final = buildFinalSeal(a.Parameters, a.Seal)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{
		Executed:                     true,
		Gate411AxiomLedgerCompiled:   true,
		Gate412HierarchyCapacity:     true,
		Gate413MixingCapacity:        true,
		Gate414NoCoefficientSelector: true,
		Gate415SourceLedgerCompiled:  true,
		Gate416RealLedgerDim:         6,
		Gate416RealNoCP:              true,
		Gate417ComplexLedgerDim:      ConditionalComplexChargedDim,
		Gate417CPCapacity:            true,
		Gate417CoefficientValuesFree: true,
		ChargedFlavorModuliDim:       Gate372ChargedFlavorModuliDim,
		Verdict:                      "Gate 418 inherits Gates 411-417: minimal family axioms give hierarchy, real mixing, and CP capacity, but no native coefficient selector or value prediction.",
	}
}

func buildProgression() AxiomProgression {
	steps := []AxiomStep{
		{
			Gate: 411, Name: "family-bundle axiom ledger", Operators: []string{"candidate axiom classes"}, AddsHierarchy: false, AddsRealMixing: false, AddsCPCapacity: false,
			GaugeCompatible: true, CompatibleWithJGamma: true, FirstOrderCompatible: true, NativeToCurrentAsha: false, PromotedToTheorem: false, EmpiricalValuesImported: false, CoefficientsFixed: false, RemainingFreeLedger: Gate372ChargedFlavorModuliDim,
			Verdict: "least-cost axiom candidates ranked", Boundary: "No candidate promoted as native ASHA derivation.",
		},
		{
			Gate: 412, Name: "minimal modular family Hamiltonian", Operators: []string{"K_gen=diag(-1,0,1)"}, AddsHierarchy: true, AddsRealMixing: false, AddsCPCapacity: false,
			GaugeCompatible: true, CompatibleWithJGamma: true, FirstOrderCompatible: true, NativeToCurrentAsha: false, PromotedToTheorem: false, EmpiricalValuesImported: false, CoefficientsFixed: false, RemainingFreeLedger: Gate372ChargedFlavorModuliDim,
			Verdict: "hierarchy capacity activated", Boundary: "A single diagonal Hamiltonian cannot produce CKM/PMNS mixing.",
		},
		{
			Gate: 413, Name: "noncommuting modular pair", Operators: []string{"K_gen", "S_gen", "X_gen=S+S^T"}, AddsHierarchy: true, AddsRealMixing: true, AddsCPCapacity: false,
			GaugeCompatible: true, CompatibleWithJGamma: true, FirstOrderCompatible: true, NativeToCurrentAsha: false, PromotedToTheorem: false, EmpiricalValuesImported: false, CoefficientsFixed: false, RemainingFreeLedger: Gate372ChargedFlavorModuliDim,
			Verdict: "real CKM/PMNS capacity activated", Boundary: "Roots of unity and noncommutation do not fix physical angles or coefficients.",
		},
		{
			Gate: 416, Name: "minimal real charge-sector source", Operators: []string{"K_gen", "X_gen"}, AddsHierarchy: true, AddsRealMixing: true, AddsCPCapacity: false,
			GaugeCompatible: true, CompatibleWithJGamma: true, FirstOrderCompatible: true, NativeToCurrentAsha: false, PromotedToTheorem: false, EmpiricalValuesImported: false, CoefficientsFixed: false, RemainingFreeLedger: 6,
			Verdict: "conditional charged ledger compressed to six real coefficients", Boundary: "Real textures have no CKM CP phase and values remain boundary data.",
		},
		{
			Gate: 417, Name: "minimal complex/phase sector source", Operators: []string{"K_gen", "X_gen", "Y_gen=i(S-S^T)"}, AddsHierarchy: true, AddsRealMixing: true, AddsCPCapacity: true,
			GaugeCompatible: true, CompatibleWithJGamma: true, FirstOrderCompatible: true, NativeToCurrentAsha: false, PromotedToTheorem: false, EmpiricalValuesImported: false, CoefficientsFixed: false, RemainingFreeLedger: ConditionalComplexChargedDim,
			Verdict: "conditional CP-capable charged ledger compressed to nine symbolic coefficients", Boundary: "The nine coefficients, CKM angles, and CP phase are not predicted.",
		},
	}
	return AxiomProgression{
		Executed: true, Steps: steps, MinimalHierarchyGate: 412, MinimalMixingGate: 413, MinimalCPGate: 417, AllAxiomsQuarantined: allQuarantined(steps), Verdict: "family axiom progression compiled without promoting any axiom to native theorem",
	}
}

func allQuarantined(steps []AxiomStep) bool {
	for _, s := range steps {
		if s.NativeToCurrentAsha || s.PromotedToTheorem || s.EmpiricalValuesImported {
			return false
		}
	}
	return true
}

func buildParameters(p AxiomProgression) ParameterReduction {
	rows := []ParameterRow{
		{Name: "native ASHA flavor frontier", Status: StatusFirewallPreserved13Moduli, Native: true, Conditional: false, EmpiricalFitting: false, ChargedCoordinates: Gate372ChargedFlavorModuliDim, HierarchyCapacity: false, CKMCapacity: false, CPCapacity: false, ValuesPredicted: false, ValuesEnvironmental: true, Comments: "No native family bundle or coefficient selector exists."},
		{Name: "K_gen hierarchy axiom", Status: "CONDITIONAL_HIERARCHY_CAPACITY", Native: false, Conditional: true, EmpiricalFitting: false, ChargedCoordinates: Gate372ChargedFlavorModuliDim, HierarchyCapacity: true, CKMCapacity: false, CPCapacity: false, ValuesPredicted: false, ValuesEnvironmental: true, Comments: "Breaks family levels but remains diagonal-only."},
		{Name: "K/X real sector-source axiom", Status: "CONDITIONAL_REAL_MIXING_NO_CP", Native: false, Conditional: true, EmpiricalFitting: false, ChargedCoordinates: 6, HierarchyCapacity: true, CKMCapacity: true, CPCapacity: false, ValuesPredicted: false, ValuesEnvironmental: true, Comments: "Two real coefficients per charged sector; no CKM CP phase."},
		{Name: "K/X/Y complex sector-source axiom", Status: "CONDITIONAL_CP_CAPABLE_LEDGER", Native: false, Conditional: true, EmpiricalFitting: false, ChargedCoordinates: ConditionalComplexChargedDim, HierarchyCapacity: true, CKMCapacity: true, CPCapacity: true, ValuesPredicted: false, ValuesEnvironmental: true, Comments: "Three symbolic coefficients per charged sector; CP-capable but not predictive."},
		{Name: "observed Yukawa/CKM input", Status: "REJECTED_CURVE_FITTING", Native: false, Conditional: false, EmpiricalFitting: true, ChargedCoordinates: Gate372ChargedFlavorModuliDim, HierarchyCapacity: true, CKMCapacity: true, CPCapacity: true, ValuesPredicted: true, ValuesEnvironmental: false, Comments: "Imports the data it should explain and is rejected as derivation."},
	}
	return ParameterReduction{
		Executed: true, StartDim: Gate372ChargedFlavorModuliDim, Rows: rows, ConditionalCompressedDim: ConditionalComplexChargedDim, NativeDim: Gate372ChargedFlavorModuliDim,
		ConditionalCompression: true, NativeCompression: false, NineCoefficientsSymbolic: true, CKMAndPhaseValuesPredicted: false,
		Verdict: "conditional axiom chain compresses the charged ledger to nine symbolic coefficients, but native ASHA remains at thirteen charged flavor moduli",
	}
}

func buildEnvironmentalSeal(p ParameterReduction) EnvironmentalSeal {
	return EnvironmentalSeal{
		Executed:               true,
		Name:                   StatusEnvironmentalSealFormalized,
		Statement:              "The finite Cℓ(1,7)/spectral-action law-space plus quarantined minimal family axioms can encode hierarchy, mixing, and CP capacity, but no native functional fixes the nine charged source coefficients; those values are environmental boundary coordinates unless a future axiom supplies them.",
		NativeLawSpaceComplete: true, FlavorCapacityComplete: true, CoefficientValuesPredicted: p.CKMAndPhaseValuesPredicted, CoefficientsEnvironmental: !p.CKMAndPhaseValuesPredicted,
		RequiresHistoricalData: true, NoEmpiricalFitting: true, NoNativeCollapseClaimed: !p.NativeCompression, Verdict: "flavor coefficient environmental seal formalized",
	}
}

func buildEmpirical() EmpiricalFirewall {
	return EmpiricalFirewall{Executed: true, NoObservedMassesImported: true, NoCKMImported: true, NoPMNSImported: true, NoYukawaMatricesInserted: true, SymbolicCoefficientsOnly: true, RejectsCurveFitting: true, Verdict: "empirical firewall preserved while sealing the coefficient frontier"}
}

func buildFinalSeal(p ParameterReduction, s EnvironmentalSeal) FinalSeal {
	return FinalSeal{
		Executed: true, NativeChargedDim: Gate372ChargedFlavorModuliDim, ConditionalComplexDim: p.ConditionalCompressedDim,
		FlavorSectorFormallySealed: true, ProjectFlavorCompleteAsLedger: true, NoNativePredictionClaimed: true,
		FirewallPreserved: !p.NativeCompression && s.CoefficientsEnvironmental,
		FinalStatus:       StatusProjectFlavorSectorSealedComplete,
		Verdict:           "family-axiom flavor frontier sealed: capacity classified, coefficient prediction not claimed",
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 419, Title: "Post-Flavor Architecture Consolidation / Final Law-Space Board", Reason: "Gate 418 seals the flavor frontier; the next useful gate is architectural consolidation, not another flavor search.", PrimaryTask: "rebuild the essential ASHA tower after the flavor closure and mark native theorems, quarantined axioms, and environmental coordinates."}
}

func truth(a Analysis) string {
	return "Gate 418 formally seals the flavor frontier. Native ASHA remains at dim M_charged=13. The quarantined K/X/Y family axiom chain supplies hierarchy, mixing, and CP capacity and conditionally compresses the charged source ledger to nine symbolic coefficients, but no native variational or geometric rule fixes those coefficients. The values are environmental boundary coordinates, not current ASHA derivations."
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Progression.Executed || !a.Parameters.Executed || !a.Seal.Executed || !a.Empirical.Executed || !a.Final.Executed {
		return fmt.Errorf("incomplete Gate418 audit")
	}
	if a.Parameters.StartDim != Gate372ChargedFlavorModuliDim || a.Final.NativeChargedDim != Gate372ChargedFlavorModuliDim {
		return fmt.Errorf("unexpected charged flavor dimension")
	}
	if a.Progression.MinimalHierarchyGate != 412 || a.Progression.MinimalMixingGate != 413 || a.Progression.MinimalCPGate != 417 {
		return fmt.Errorf("unexpected axiom progression gates")
	}
	if !a.Progression.AllAxiomsQuarantined {
		return fmt.Errorf("no family axiom may be promoted to native theorem in Gate418")
	}
	if a.Parameters.ConditionalCompressedDim != ConditionalComplexChargedDim || !a.Parameters.NineCoefficientsSymbolic {
		return fmt.Errorf("expected nine symbolic conditional coefficients")
	}
	if a.Parameters.NativeCompression || a.Parameters.CKMAndPhaseValuesPredicted {
		return fmt.Errorf("Gate418 must not claim native compression or CKM prediction")
	}
	if !a.Seal.CoefficientsEnvironmental || a.Seal.CoefficientValuesPredicted {
		return fmt.Errorf("environmental coefficient seal failed")
	}
	if !a.Final.FirewallPreserved || a.Final.FinalStatus != StatusProjectFlavorSectorSealedComplete {
		return fmt.Errorf("final flavor seal failed")
	}
	return nil
}

func Statuses() []string {
	return []string{
		StatusGate417Inherited,
		StatusAxiomProgressionLedgerCompiled,
		StatusParameterReductionSummaryCompiled,
		StatusEnvironmentalSealFormalized,
		StatusProjectFlavorSectorSealedComplete,
		StatusEmpiricalFirewallPreserved,
		StatusNoNativeFlavorPrediction,
		StatusNoAxiomPromotedNative,
		StatusNineCoefficientsRemainFree,
		StatusCKMPhaseNotPredicted,
		StatusFirewallPreserved13Moduli,
	}
}

func joinOps(ops []string) string { return strings.Join(ops, ", ") }
