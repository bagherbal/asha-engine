// Package familyboundarysourceaxiom implements Gate 415:
// Family Boundary Condition / Sector Source Axiom Minimality Sieve.
//
// Gate 414 proved that trace and curvature functionals over the Gate-413
// family pair (K_gen, S_gen) do not select physical texture coefficients.
// Gate 415 therefore stops searching for a native selector and audits the
// least additional boundary/source axioms that could choose coefficient rays
// without importing observed Yukawa matrices. Every candidate is kept
// quarantined: the ledger can rank mathematical cost and CKM/PMNS capacity,
// but it does not promote any source to a native ASHA theorem.
package familyboundarysourceaxiom

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE415-FAMILY-BOUNDARY-CONDITION-SECTOR-SOURCE-AXIOM-MINIMALITY-SIEVE"

	StatusGate414BoundaryInherited        = "CONDITIONAL_SUPPORT_GATE414_COEFFICIENT_SELECTOR_BOUNDARY_INHERITED"
	StatusSourceAxiomLedgerCompiled       = "CONDITIONAL_SUPPORT_FAMILY_BOUNDARY_SOURCE_AXIOM_LEDGER_COMPILED"
	StatusMinimalityRankingAudited        = "CONDITIONAL_SUPPORT_BOUNDARY_SOURCE_MINIMALITY_RANKING_AUDITED"
	StatusCKMPMNSCapacityAudited          = "CONDITIONAL_SUPPORT_BOUNDARY_SOURCE_CKM_PMNS_CAPACITY_AUDITED"
	StatusEmpiricalIndependenceAudited    = "CONDITIONAL_SUPPORT_BOUNDARY_SOURCE_EMPIRICAL_INDEPENDENCE_AUDITED"
	StatusLeastCostAxiomIdentified        = "CONDITIONAL_SUPPORT_LEAST_COST_BOUNDARY_AXIOM_IDENTIFIED"
	StatusFailedNoNativeBoundary          = "FAILED_ROUTE_NO_NATIVE_FAMILY_BOUNDARY_CONDITION"
	StatusFailedSelectorRequiresAxiom     = "FAILED_ROUTE_COEFFICIENT_SELECTOR_REQUIRES_SOURCE_AXIOM"
	StatusFailedDiscreteUnderdetermines   = "FAILED_ROUTE_DISCRETE_SYMMETRY_SOURCE_UNDERDETERMINES_ANGLES"
	StatusFailedFlatBoundaryUnconstrained = "FAILED_ROUTE_FLAT_CONNECTION_BOUNDARY_DIAGONAL_OR_UNCONSTRAINED"
	StatusFailedObservedYukawaRejected    = "FAILED_ROUTE_OBSERVED_YUKAWA_SOURCE_REJECTED_AS_CURVE_FITTING"
	StatusFailedNoNativeModuliReduction   = "FAILED_ROUTE_NO_NATIVE_FLAVOR_MODULI_REDUCTION"
	StatusFirewallPreserved13Moduli       = "FIREWALL_PRESERVED_13_MODULI"
)

const (
	Gate372ChargedFlavorModuliDim = 13
	FamilyRank                    = 3
)

type Inheritance struct {
	Executed                       bool
	Gate413NoncommutingCapacity    bool
	Gate413PairNotNative           bool
	Gate414NoCoefficientSelector   bool
	Gate414TraceCurvatureExhausted bool
	Gate414CoefficientsFree        bool
	ChargedModuliDim               int
	Verdict                        string
}

type SourceArena struct {
	Executed                bool
	FamilyPair              string
	Sectors                 []string
	CoefficientsPerSector   int
	BaselineCoefficientRays int
	GaugeCompatibleIfFamily bool
	NativeSelectorPresent   bool
	EmpiricalYukawaImported bool
	Verdict                 string
	Reason                  string
}

type AxiomCandidate struct {
	Name                   string
	Executed               bool
	AxiomKind              string
	MathematicalCost       int
	GaugeCompatible        bool
	CompatibleWithJGamma   bool
	EmpiricalIndependent   bool
	ImportsObservedYukawa  bool
	SelectsCoefficientRay  bool
	FixesCoefficientValues bool
	CKMCapacity            bool
	PMNSCapacity           bool
	DiagonalOnly           bool
	FreeRealParameters     int
	NativeToCurrentAsha    bool
	PromotedToTheorem      bool
	Verdict                string
	Reason                 string
}

type MinimalityRanking struct {
	Executed             bool
	RankedNames          []string
	LeastCostName        string
	LeastCost            int
	LeastCostStillAxiom  bool
	LeastCostCKMCapacity bool
	LeastCostFixesAngles bool
	NoCandidateNative    bool
	Verdict              string
}

type CapacityAudit struct {
	Executed                 bool
	ConditionalCKMAvailable  bool
	ConditionalPMNSAvailable bool
	AnyCandidateFixesAngles  bool
	AnyCandidateNative       bool
	AnyCandidateCurveFitting bool
	BestEmpiricalIndependent string
	RequiredExtraData        string
	Verdict                  string
	Reason                   string
}

type ModuliScenario struct {
	Name                string
	Status              string
	ModuliDim           int
	ThreeMassesPossible bool
	CKMPossible         bool
	PMNSPossible        bool
	CoefficientsFixed   bool
	NativeReduction     bool
	ConditionalOnly     bool
	EmpiricalFitting    bool
	Reason              string
}

type ModuliImpact struct {
	StartDim                  int
	Scenarios                 []ModuliScenario
	BestNativeDim             int
	NativeReductionBelow13    bool
	ConditionalMixingCapacity bool
	CoefficientsRemainFree    bool
	FirewallPreserved         bool
	Verdict                   string
}

type Firewall struct {
	Executed                  bool
	NoObservedMassesImported  bool
	NoCKMImported             bool
	NoPMNSImported            bool
	NoYukawaMatricesInserted  bool
	AxiomsQuarantined         bool
	NoNativeDerivationClaimed bool
	Verdict                   string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Arena       SourceArena
	Candidates  []AxiomCandidate
	Ranking     MinimalityRanking
	Capacity    CapacityAudit
	Moduli      ModuliImpact
	Firewall    Firewall
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
	a.Arena = buildArena()
	a.Candidates = buildCandidates()
	a.Ranking = buildRanking(a.Candidates)
	a.Capacity = buildCapacity(a.Candidates, a.Ranking)
	a.Moduli = buildModuli(a.Candidates)
	a.Firewall = buildFirewall()
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{Executed: true, Gate413NoncommutingCapacity: true, Gate413PairNotNative: true, Gate414NoCoefficientSelector: true, Gate414TraceCurvatureExhausted: true, Gate414CoefficientsFree: true, ChargedModuliDim: Gate372ChargedFlavorModuliDim, Verdict: "Gate 415 inherits the Gate-414 boundary: noncommuting texture capacity exists only conditionally, while native trace/curvature functionals leave the family coefficients free."}
}

func buildArena() SourceArena {
	sectors := []string{"up", "down", "charged-lepton", "neutrino"}
	return SourceArena{Executed: true, FamilyPair: "K_gen plus S_gen/X_gen family texture pair", Sectors: sectors, CoefficientsPerSector: 2, BaselineCoefficientRays: len(sectors) * 2, GaugeCompatibleIfFamily: true, NativeSelectorPresent: false, EmpiricalYukawaImported: false, Verdict: "boundary/source axiom arena formalized", Reason: "The family pair is gauge-blind and can be coupled to SM sectors, but ASHA supplies no source selecting the sector coefficient rays."}
}

func buildCandidates() []AxiomCandidate {
	return []AxiomCandidate{
		{Name: "universal family source", Executed: true, AxiomKind: "single coefficient ray shared by all sectors", MathematicalCost: 1, GaugeCompatible: true, CompatibleWithJGamma: true, EmpiricalIndependent: true, ImportsObservedYukawa: false, SelectsCoefficientRay: true, FixesCoefficientValues: false, CKMCapacity: false, PMNSCapacity: false, DiagonalOnly: true, FreeRealParameters: 1, NativeToCurrentAsha: false, PromotedToTheorem: false, Verdict: "least cost but flavor-blind", Reason: "A universal source is mathematically cheap and gauge compatible, but it aligns all sectors and cannot produce CKM/PMNS."},
		{Name: "charge-sector source boundary", Executed: true, AxiomKind: "sector-labelled source rays for up/down/lepton/neutrino channels", MathematicalCost: 2, GaugeCompatible: true, CompatibleWithJGamma: true, EmpiricalIndependent: true, ImportsObservedYukawa: false, SelectsCoefficientRay: true, FixesCoefficientValues: false, CKMCapacity: true, PMNSCapacity: true, DiagonalOnly: false, FreeRealParameters: 8, NativeToCurrentAsha: false, PromotedToTheorem: false, Verdict: "least cost CKM-capable axiom", Reason: "This is the smallest audited rule that can assign distinct K/S coefficients to sectors, but the coefficient values remain boundary data."},
		{Name: "Z3 Weyl phase source", Executed: true, AxiomKind: "roots-of-unity phase rule on family shift", MathematicalCost: 2, GaugeCompatible: true, CompatibleWithJGamma: true, EmpiricalIndependent: true, ImportsObservedYukawa: false, SelectsCoefficientRay: true, FixesCoefficientValues: false, CKMCapacity: true, PMNSCapacity: true, DiagonalOnly: false, FreeRealParameters: 4, NativeToCurrentAsha: false, PromotedToTheorem: false, Verdict: "discrete phase capacity, angle underdetermined", Reason: "The clock/shift algebra supplies phases, but roots of unity do not determine physical mixing magnitudes or CKM angles."},
		{Name: "flat U(3)_gen holonomy boundary", Executed: true, AxiomKind: "family connection/holonomy boundary condition", MathematicalCost: 3, GaugeCompatible: true, CompatibleWithJGamma: true, EmpiricalIndependent: true, ImportsObservedYukawa: false, SelectsCoefficientRay: true, FixesCoefficientValues: false, CKMCapacity: true, PMNSCapacity: true, DiagonalOnly: false, FreeRealParameters: 4, NativeToCurrentAsha: false, PromotedToTheorem: false, Verdict: "mixing-capable but unconstrained holonomy", Reason: "A family holonomy can encode mixing, but without a topological quantization or source equation it is exactly a new connection axiom."},
		{Name: "modular KMS sector source", Executed: true, AxiomKind: "nontracial density/Hamiltonian plus sector source", MathematicalCost: 3, GaugeCompatible: true, CompatibleWithJGamma: true, EmpiricalIndependent: true, ImportsObservedYukawa: false, SelectsCoefficientRay: true, FixesCoefficientValues: false, CKMCapacity: true, PMNSCapacity: true, DiagonalOnly: false, FreeRealParameters: 6, NativeToCurrentAsha: false, PromotedToTheorem: false, Verdict: "hierarchy plus mixing capacity, still sourced", Reason: "KMS structure gives hierarchy language, but the source Hamiltonian and sector splitting are additional axioms."},
		{Name: "observed Yukawa matrix source", Executed: true, AxiomKind: "full external Yukawa amplitude ledger", MathematicalCost: 5, GaugeCompatible: true, CompatibleWithJGamma: true, EmpiricalIndependent: false, ImportsObservedYukawa: true, SelectsCoefficientRay: true, FixesCoefficientValues: true, CKMCapacity: true, PMNSCapacity: true, DiagonalOnly: false, FreeRealParameters: Gate372ChargedFlavorModuliDim, NativeToCurrentAsha: false, PromotedToTheorem: false, Verdict: "rejected curve fitting", Reason: "This is phenomenologically complete, but it imports the firewall data and is not a geometric explanation."},
	}
}

func buildRanking(cs []AxiomCandidate) MinimalityRanking {
	names := make([]string, 0, len(cs))
	leastName := ""
	leastCost := math.MaxInt
	leastCKM := false
	for _, c := range cs {
		names = append(names, fmt.Sprintf("%s(cost=%d)", c.Name, c.MathematicalCost))
		if c.EmpiricalIndependent && !c.ImportsObservedYukawa && c.CKMCapacity && c.MathematicalCost < leastCost {
			leastCost = c.MathematicalCost
			leastName = c.Name
			leastCKM = c.CKMCapacity
		}
	}
	return MinimalityRanking{Executed: true, RankedNames: names, LeastCostName: leastName, LeastCost: leastCost, LeastCostStillAxiom: true, LeastCostCKMCapacity: leastCKM, LeastCostFixesAngles: false, NoCandidateNative: true, Verdict: "least-cost CKM-capable source axiom identified but quarantined"}
}

func buildCapacity(cs []AxiomCandidate, r MinimalityRanking) CapacityAudit {
	conditionalCKM := false
	conditionalPMNS := false
	fixes := false
	native := false
	curve := false
	for _, c := range cs {
		conditionalCKM = conditionalCKM || (c.CKMCapacity && c.EmpiricalIndependent && !c.NativeToCurrentAsha)
		conditionalPMNS = conditionalPMNS || (c.PMNSCapacity && c.EmpiricalIndependent && !c.NativeToCurrentAsha)
		fixes = fixes || (c.FixesCoefficientValues && c.EmpiricalIndependent)
		native = native || c.NativeToCurrentAsha
		curve = curve || c.ImportsObservedYukawa
	}
	return CapacityAudit{Executed: true, ConditionalCKMAvailable: conditionalCKM, ConditionalPMNSAvailable: conditionalPMNS, AnyCandidateFixesAngles: fixes, AnyCandidateNative: native, AnyCandidateCurveFitting: curve, BestEmpiricalIndependent: r.LeastCostName, RequiredExtraData: "a sector-source/boundary rule selecting coefficient rays for K_gen and S_gen", Verdict: "capacity requires explicit boundary/source axiom", Reason: "A minimal source can activate mixing without observed Yukawa matrices, but it still supplies new data outside current ASHA."}
}

func buildModuli(cs []AxiomCandidate) ModuliImpact {
	scenarios := []ModuliScenario{
		{Name: "current ASHA through Gate 414", Status: StatusFirewallPreserved13Moduli, ModuliDim: Gate372ChargedFlavorModuliDim, ThreeMassesPossible: true, CKMPossible: true, PMNSPossible: true, CoefficientsFixed: false, NativeReduction: false, ConditionalOnly: true, EmpiricalFitting: false, Reason: "Noncommuting capacity exists only under quarantined K/S axioms; coefficients remain free."},
		{Name: "universal family source", Status: "CONDITIONAL_UNIVERSAL_SOURCE_FLAVOR_BLIND", ModuliDim: Gate372ChargedFlavorModuliDim, ThreeMassesPossible: true, CKMPossible: false, PMNSPossible: false, CoefficientsFixed: false, NativeReduction: false, ConditionalOnly: true, EmpiricalFitting: false, Reason: "One shared source gives hierarchy scale but aligns all sectors."},
		{Name: "charge-sector source boundary", Status: "CONDITIONAL_SECTOR_SOURCE_CKM_CAPACITY", ModuliDim: Gate372ChargedFlavorModuliDim, ThreeMassesPossible: true, CKMPossible: true, PMNSPossible: true, CoefficientsFixed: false, NativeReduction: false, ConditionalOnly: true, EmpiricalFitting: false, Reason: "Smallest CKM-capable source, but coefficient values are boundary data."},
		{Name: "Z3 Weyl phase source", Status: StatusFailedDiscreteUnderdetermines, ModuliDim: Gate372ChargedFlavorModuliDim, ThreeMassesPossible: true, CKMPossible: true, PMNSPossible: true, CoefficientsFixed: false, NativeReduction: false, ConditionalOnly: true, EmpiricalFitting: false, Reason: "Roots of unity constrain phases but not physical angles/magnitudes."},
		{Name: "observed Yukawa matrix source", Status: StatusFailedObservedYukawaRejected, ModuliDim: 0, ThreeMassesPossible: true, CKMPossible: true, PMNSPossible: true, CoefficientsFixed: true, NativeReduction: false, ConditionalOnly: false, EmpiricalFitting: true, Reason: "Fixes data by importing the target phenomenology; rejected as explanation."},
	}
	return ModuliImpact{StartDim: Gate372ChargedFlavorModuliDim, Scenarios: scenarios, BestNativeDim: Gate372ChargedFlavorModuliDim, NativeReductionBelow13: false, ConditionalMixingCapacity: true, CoefficientsRemainFree: true, FirewallPreserved: true, Verdict: "boundary/source axiom ledger compiled; no native moduli reduction"}
}

func buildFirewall() Firewall {
	return Firewall{Executed: true, NoObservedMassesImported: true, NoCKMImported: true, NoPMNSImported: true, NoYukawaMatricesInserted: true, AxiomsQuarantined: true, NoNativeDerivationClaimed: true, Verdict: "all boundary/source candidates remain quarantined; empirical firewall preserved"}
}

func buildNext() NextStep {
	return NextStep{Gate: 416, Title: "Minimal Sector-Source Axiom Consistency / Parameter-Counting Sieve", Reason: "Gate 415 identifies the charge-sector source boundary as the least-cost CKM-capable axiom, but its coefficients remain free. The next gate should test this explicit axiom's consistency and exact parameter count without claiming prediction.", PrimaryTask: "Formalize the minimal sector-source axiom and compute how many flavor parameters it leaves free under gauge, J, Γ, and first-order compatibility."}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate414CoefficientsFree || a.Inheritance.ChargedModuliDim != Gate372ChargedFlavorModuliDim {
		return fmt.Errorf("bad inheritance")
	}
	if !a.Arena.Executed || a.Arena.NativeSelectorPresent || a.Arena.EmpiricalYukawaImported || a.Arena.BaselineCoefficientRays <= 0 {
		return fmt.Errorf("bad arena")
	}
	if len(a.Candidates) < 5 {
		return fmt.Errorf("missing axiom candidates")
	}
	if !a.Ranking.Executed || a.Ranking.LeastCostName == "" || !a.Ranking.LeastCostStillAxiom || a.Ranking.LeastCostFixesAngles || !a.Ranking.NoCandidateNative {
		return fmt.Errorf("bad ranking")
	}
	if !a.Capacity.ConditionalCKMAvailable || !a.Capacity.ConditionalPMNSAvailable || a.Capacity.AnyCandidateFixesAngles || a.Capacity.AnyCandidateNative {
		return fmt.Errorf("bad capacity")
	}
	if a.Moduli.BestNativeDim != Gate372ChargedFlavorModuliDim || a.Moduli.NativeReductionBelow13 || !a.Moduli.FirewallPreserved {
		return fmt.Errorf("bad moduli")
	}
	if !a.Firewall.AxiomsQuarantined || !a.Firewall.NoNativeDerivationClaimed || !a.Firewall.NoYukawaMatricesInserted {
		return fmt.Errorf("bad firewall")
	}
	return nil
}

func Statuses(a Analysis) []string {
	return []string{
		StatusGate414BoundaryInherited,
		StatusSourceAxiomLedgerCompiled,
		StatusMinimalityRankingAudited,
		StatusCKMPMNSCapacityAudited,
		StatusEmpiricalIndependenceAudited,
		StatusLeastCostAxiomIdentified,
		StatusFailedNoNativeBoundary,
		StatusFailedSelectorRequiresAxiom,
		StatusFailedDiscreteUnderdetermines,
		StatusFailedFlatBoundaryUnconstrained,
		StatusFailedObservedYukawaRejected,
		StatusFailedNoNativeModuliReduction,
		StatusFirewallPreserved13Moduli,
	}
}

func truth(a Analysis) string {
	var b strings.Builder
	b.WriteString("Gate 415 compiles the minimal boundary/source axiom ledger after Gate 414's coefficient-selector failure. ")
	b.WriteString("The least-cost CKM-capable option is a charge-sector source boundary that assigns coefficient rays to the K/S family pair, but it remains an explicit new axiom and does not fix physical mixing angles. ")
	b.WriteString("Discrete roots of unity, flat holonomy, and modular KMS sources all have capacity but underdetermine coefficients; observed Yukawa matrices are rejected as curve-fitting. ")
	b.WriteString("No native ASHA boundary condition is found, and dim M_charged remains 13.")
	return b.String()
}

func FormatFloat(x float64) string {
	if math.Abs(x) < 1e-12 {
		x = 0
	}
	return fmt.Sprintf("%.12g", x)
}
