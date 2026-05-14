// Package familybundleaxiomledger implements Gate 411:
// Axiom-Candidate Ledger for Nontrivial Family Bundle Extensions.
//
// Gate 410 proved that current ASHA data do not derive a nontrivial
// fermionic family bundle. Gate 411 is therefore not another hidden flavor
// derivation and not an empirical Yukawa seal. It compiles a quarantined,
// ranked ledger of minimal mathematical extensions that could be tested in
// future gates without promoting any of them as native ASHA theorems.
package familybundleaxiomledger

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE411-AXIOM-CANDIDATE-LEDGER-NONTRIVIAL-FAMILY-BUNDLE-EXTENSIONS"

	StatusGate410Inherited                        = "CONDITIONAL_SUPPORT_GATE410_FAMILY_BUNDLE_EXTENSION_BOUNDARY_INHERITED"
	StatusAxiomLedgerCompiled                     = "CONDITIONAL_SUPPORT_FAMILY_BUNDLE_AXIOM_LEDGER_COMPILED"
	StatusEpistemologicalBoundaryDocumented       = "CONDITIONAL_SUPPORT_EPISTEMOLOGICAL_BOUNDARY_DOCUMENTED"
	StatusAxiomCostRankingAudited                 = "CONDITIONAL_SUPPORT_AXIOM_COST_RANKING_AUDITED"
	StatusCKMPMNSCapacityAudited                  = "CONDITIONAL_SUPPORT_CKM_PMNS_CAPACITY_AUDITED"
	StatusEmpiricalIndependenceAudited            = "CONDITIONAL_SUPPORT_EMPIRICAL_INDEPENDENCE_AUDITED"
	StatusLeastCostCandidateIdentified            = "CONDITIONAL_SUPPORT_LEAST_COST_AXIOM_CANDIDATE_IDENTIFIED"
	StatusModularHamiltonianCandidateQuarantined  = "CONDITIONAL_SUPPORT_MODULAR_HAMILTONIAN_AXIOM_CANDIDATE_QUARANTINED"
	StatusFamilyConnectionCandidateQuarantined    = "CONDITIONAL_SUPPORT_FAMILY_CONNECTION_AXIOM_CANDIDATE_QUARANTINED"
	StatusPrimitiveIdealCandidateQuarantined      = "CONDITIONAL_SUPPORT_PRIMITIVE_IDEAL_EXTENSION_CANDIDATE_QUARANTINED"
	StatusTrialityLocalSystemCandidateQuarantined = "CONDITIONAL_SUPPORT_TRIALITY_LOCAL_SYSTEM_CANDIDATE_QUARANTINED"
	StatusFailedNoAxiomPromoted                   = "FAILED_ROUTE_NO_AXIOM_PROMOTED_TO_NATIVE_THEOREM"
	StatusFailedModularKMSNeedsHamiltonianAxiom   = "FAILED_ROUTE_MODULAR_KMS_NEEDS_HAMILTONIAN_AXIOM"
	StatusFailedFamilyConnectionNeedsAxiom        = "FAILED_ROUTE_FAMILY_CONNECTION_NEEDS_CONNECTION_AXIOM"
	StatusFailedPrimitiveIdealNeedsAlgebraAxiom   = "FAILED_ROUTE_PRIMITIVE_IDEAL_EXTENSION_NEEDS_ALGEBRA_AXIOM"
	StatusFailedTrialityLocalSystemNeedsFunctor   = "FAILED_ROUTE_TRIALITY_LOCAL_SYSTEM_NEEDS_FUNCTOR_AXIOM"
	StatusFailedUnconstrainedSourceIsCurveFitting = "FAILED_ROUTE_UNCONSTRAINED_FAMILY_SOURCE_COLLAPSES_TO_CURVE_FITTING"
	StatusFailedNoNativeCKMPMNS                   = "FAILED_ROUTE_NO_NATIVE_CKM_PMNS_FROM_CURRENT_ASHA"
	StatusFirewallPreserved13Moduli               = "FIREWALL_PRESERVED_13_MODULI"
)

const (
	Gate372ChargedFlavorModuliDim = 13
	FamilyRank                    = 3
)

type Inheritance struct {
	Executed                          bool
	Gate410NoNativeFamilyBundle       bool
	Gate410KOTwistOnlySigns           bool
	Gate410KMSRequiresExternalK       bool
	Gate410PrimitiveIdealsWrongDomain bool
	Gate410RequiresNewAxiom           bool
	Gate409TrivialU3Multiplicity      bool
	Gate408ScalarFlavorBlind          bool
	Gate372ChargedModuliDim           int
	NoEmpiricalInputsImported         bool
	Verdict                           string
}

type AxiomCandidate struct {
	Name                           string
	Kind                           string
	MinimalAdditionalData          string
	MathematicalCost               int
	NativeInCurrentAsha            bool
	PromotedToTheorem              bool
	ChangesFamilyCarrier           bool
	ReplacesTrivialC3              bool
	ProvidesThreeFamilies          bool
	ProvidesDiagonalOperator       bool
	ProvidesTwoNoncommutingOps     bool
	CKMCapacity                    bool
	PMNSCapacity                   bool
	PureGeometricFormulable        bool
	EmpiricalIndependentFormulable bool
	CurveFittingRisk               string
	PreservesAF                    bool
	PreservesJ                     bool
	PreservesGamma                 bool
	PreservesFirstOrder            bool
	PreservesGaugeCharges          bool
	RequiresNewAxiom               bool
	RequiresExternalHamiltonian    bool
	RequiresAlgebraExtension       bool
	RequiresFunctor                bool
	RequiresConnection             bool
	RequiresEmpiricalYukawas       bool
	Verdict                        string
	Reason                         string
}

type AxiomLedger struct {
	Executed                  bool
	Candidates                []AxiomCandidate
	CandidatesAudited         int
	PromotedAxioms            int
	PureGeometricCandidates   int
	EmpiricalIndependentCount int
	CurveFittingRiskHigh      int
	CKMCapableCandidates      int
	PMNSCapableCandidates     int
	LowestCost                int
	LeastCostNames            []string
	Verdict                   string
}

type CapacityAudit struct {
	Executed                     bool
	CandidatesWithDiagonalOnly   int
	CandidatesWithNoncommuting   int
	NativeNoncommutingPairs      int
	ConditionalNoncommutingPairs int
	CKMNative                    bool
	CKMConditional               bool
	PMNSNative                   bool
	PMNSConditional              bool
	Verdict                      string
	Reason                       string
}

type EmpiricalIndependenceAudit struct {
	Executed                      bool
	NoObservedMassesImported      bool
	NoCKMImported                 bool
	NoPMNSImported                bool
	NoYukawaMatricesInserted      bool
	CandidatesCanBePureRules      int
	CandidatesCollapseToFitting   int
	UnconstrainedFamilySourceRisk bool
	Verdict                       string
	Reason                        string
}

type CostRanking struct {
	Executed bool
	Rows     []CostRow
	Verdict  string
}

type CostRow struct {
	Rank     int
	Name     string
	Cost     int
	Benefit  string
	MainRisk string
	NextTest string
}

type Boundary struct {
	Executed                    bool
	LawSpaceNative              bool
	FamilyBundleNative          bool
	NewAxiomRequiredForFamilies bool
	CurrentASHAFlavorComplete   bool
	Statement                   string
	Verdict                     string
}

type ModuliScenario struct {
	Name                        string
	Status                      string
	ModuliDim                   int
	ThreeDistinctMassesPossible bool
	CKMPossible                 bool
	PMNSPossible                bool
	Reason                      string
}

type ModuliImpact struct {
	StartDim                    int
	Scenarios                   []ModuliScenario
	BestNativeDim               int
	NativeReductionBelow13      bool
	ConditionalReductionBelow13 bool
	FirewallPreserved           bool
	Verdict                     string
}

type FirewallAudit struct {
	Executed                      bool
	NoObservedMassesImported      bool
	NoCKMImported                 bool
	NoPMNSImported                bool
	NoYukawaAmplitudesInserted    bool
	NoAxiomPromoted               bool
	NoExternalHamiltonianPromoted bool
	NoFamilyConnectionPromoted    bool
	NoAlgebraExtensionPromoted    bool
	NoFunctorPromoted             bool
	NoModuliReductionClaimed      bool
	Verdict                       string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance           Inheritance
	Ledger                AxiomLedger
	Capacity              CapacityAudit
	EmpiricalIndependence EmpiricalIndependenceAudit
	Ranking               CostRanking
	Boundary              Boundary
	Moduli                ModuliImpact
	Firewall              FirewallAudit
	Next                  NextStep
	Truth                 string
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
	a.Ledger = buildLedger()
	a.Capacity = buildCapacity(a.Ledger)
	a.EmpiricalIndependence = buildEmpiricalIndependence(a.Ledger)
	a.Ranking = buildRanking(a.Ledger)
	a.Boundary = buildBoundary(a.Ledger)
	a.Moduli = buildModuli(a.Ledger, a.Capacity)
	a.Firewall = buildFirewall(a.Moduli)
	a.Next = buildNext(a.Ranking)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate410NoNativeFamilyBundle || a.Inheritance.Gate372ChargedModuliDim != Gate372ChargedFlavorModuliDim {
		return fmt.Errorf("Gate410 inheritance missing")
	}
	if !a.Ledger.Executed || a.Ledger.CandidatesAudited < 5 || a.Ledger.PromotedAxioms != 0 {
		return fmt.Errorf("axiom ledger incomplete or promoted an axiom")
	}
	if a.Capacity.NativeNoncommutingPairs != 0 || a.Capacity.CKMNative || a.Capacity.PMNSNative {
		return fmt.Errorf("native CKM/PMNS capacity was promoted")
	}
	if !a.EmpiricalIndependence.NoObservedMassesImported || !a.EmpiricalIndependence.NoYukawaMatricesInserted {
		return fmt.Errorf("empirical data imported")
	}
	if !a.Boundary.NewAxiomRequiredForFamilies || a.Boundary.FamilyBundleNative {
		return fmt.Errorf("epistemological boundary not preserved")
	}
	if a.Moduli.StartDim != Gate372ChargedFlavorModuliDim || a.Moduli.BestNativeDim != Gate372ChargedFlavorModuliDim || !a.Moduli.FirewallPreserved {
		return fmt.Errorf("charged flavor firewall broken incorrectly")
	}
	if !a.Firewall.NoAxiomPromoted || !a.Firewall.NoModuliReductionClaimed {
		return fmt.Errorf("firewall violation")
	}
	return nil
}

func buildInheritance() Inheritance {
	return Inheritance{Executed: true, Gate410NoNativeFamilyBundle: true, Gate410KOTwistOnlySigns: true, Gate410KMSRequiresExternalK: true, Gate410PrimitiveIdealsWrongDomain: true, Gate410RequiresNewAxiom: true, Gate409TrivialU3Multiplicity: true, Gate408ScalarFlavorBlind: true, Gate372ChargedModuliDim: Gate372ChargedFlavorModuliDim, NoEmpiricalInputsImported: true, Verdict: StatusGate410Inherited}
}

func buildLedger() AxiomLedger {
	candidates := []AxiomCandidate{
		{Name: "minimal modular family Hamiltonian axiom", Kind: "modular/KMS state", MinimalAdditionalData: "a geometric rule selecting traceless Hermitian K_gen on the family fiber", MathematicalCost: 2, ChangesFamilyCarrier: false, ReplacesTrivialC3: false, ProvidesThreeFamilies: true, ProvidesDiagonalOperator: true, ProvidesTwoNoncommutingOps: false, CKMCapacity: false, PMNSCapacity: false, PureGeometricFormulable: true, EmpiricalIndependentFormulable: true, CurveFittingRisk: "medium", PreservesAF: true, PreservesJ: false, PreservesGamma: true, PreservesFirstOrder: false, PreservesGaugeCharges: true, RequiresNewAxiom: true, RequiresExternalHamiltonian: true, Verdict: StatusFailedModularKMSNeedsHamiltonianAxiom, Reason: "Lowest-cost hierarchy candidate: it can make a three-level state, but alone it is diagonal and still needs a native Hamiltonian axiom."},
		{Name: "nontrivial U(3)_gen family connection axiom", Kind: "family bundle connection", MinimalAdditionalData: "connection/curvature on the generation fiber with constrained holonomy", MathematicalCost: 3, ChangesFamilyCarrier: true, ReplacesTrivialC3: true, ProvidesThreeFamilies: true, ProvidesDiagonalOperator: true, ProvidesTwoNoncommutingOps: true, CKMCapacity: true, PMNSCapacity: true, PureGeometricFormulable: true, EmpiricalIndependentFormulable: true, CurveFittingRisk: "high unless curvature is quantized", PreservesAF: true, PreservesJ: false, PreservesGamma: true, PreservesFirstOrder: false, PreservesGaugeCharges: true, RequiresNewAxiom: true, RequiresConnection: true, Verdict: StatusFailedFamilyConnectionNeedsAxiom, Reason: "This has the right CKM/PMNS capacity, but it is exactly the missing family-bundle connection, not a native consequence."},
		{Name: "primitive ideal family algebra extension", Kind: "finite algebra extension", MinimalAdditionalData: "new primitive family summands or a finite algebra whose irreps are the families", MathematicalCost: 4, ChangesFamilyCarrier: true, ReplacesTrivialC3: true, ProvidesThreeFamilies: true, ProvidesDiagonalOperator: true, ProvidesTwoNoncommutingOps: true, CKMCapacity: true, PMNSCapacity: true, PureGeometricFormulable: true, EmpiricalIndependentFormulable: true, CurveFittingRisk: "medium-high", PreservesAF: false, PreservesJ: false, PreservesGamma: true, PreservesFirstOrder: false, PreservesGaugeCharges: false, RequiresNewAxiom: true, RequiresAlgebraExtension: true, Verdict: StatusFailedPrimitiveIdealNeedsAlgebraAxiom, Reason: "Can encode families structurally, but changes the finite algebra and must reprove first-order/J/gauge compatibility."},
		{Name: "triality local-system functor axiom", Kind: "Spin(8) representation functor", MinimalAdditionalData: "typed functor from finite-Dirac family states to 8v, 8s, 8c with a native breaking datum", MathematicalCost: 3, ChangesFamilyCarrier: true, ReplacesTrivialC3: true, ProvidesThreeFamilies: true, ProvidesDiagonalOperator: false, ProvidesTwoNoncommutingOps: false, CKMCapacity: false, PMNSCapacity: false, PureGeometricFormulable: true, EmpiricalIndependentFormulable: true, CurveFittingRisk: "medium", PreservesAF: false, PreservesJ: false, PreservesGamma: true, PreservesFirstOrder: false, PreservesGaugeCharges: false, RequiresNewAxiom: true, RequiresFunctor: true, Verdict: StatusFailedTrialityLocalSystemNeedsFunctor, Reason: "Triality is the correct threefold arena but exact triality degenerates; a functor and a breaking datum are still axiomatic."},
		{Name: "contact singleton family-label axiom", Kind: "contact-to-family functor", MinimalAdditionalData: "functor from the three rational contact singleton blocks to finite-Dirac family labels", MathematicalCost: 3, ChangesFamilyCarrier: true, ReplacesTrivialC3: true, ProvidesThreeFamilies: true, ProvidesDiagonalOperator: true, ProvidesTwoNoncommutingOps: false, CKMCapacity: false, PMNSCapacity: false, PureGeometricFormulable: true, EmpiricalIndependentFormulable: true, CurveFittingRisk: "medium", PreservesAF: false, PreservesJ: false, PreservesGamma: false, PreservesFirstOrder: false, PreservesGaugeCharges: false, RequiresNewAxiom: true, RequiresFunctor: true, Verdict: StatusFailedTrialityLocalSystemNeedsFunctor, Reason: "The three contact singleton blocks are native, but Gate 397 rejected their finite-Dirac flavor functor; using them requires a new functor axiom."},
		{Name: "unconstrained external Yukawa/source matrix", Kind: "empirical source", MinimalAdditionalData: "arbitrary complex Yukawa matrices or observed masses and CKM/PMNS angles", MathematicalCost: 5, ChangesFamilyCarrier: false, ReplacesTrivialC3: false, ProvidesThreeFamilies: true, ProvidesDiagonalOperator: true, ProvidesTwoNoncommutingOps: true, CKMCapacity: true, PMNSCapacity: true, PureGeometricFormulable: false, EmpiricalIndependentFormulable: false, CurveFittingRisk: "maximal", PreservesAF: true, PreservesJ: false, PreservesGamma: true, PreservesFirstOrder: false, PreservesGaugeCharges: true, RequiresEmpiricalYukawas: true, Verdict: StatusFailedUnconstrainedSourceIsCurveFitting, Reason: "This reproduces flavor by definition and is therefore excluded from the axiom ledger as a theorem route."},
	}
	out := AxiomLedger{Executed: true, Candidates: candidates, CandidatesAudited: len(candidates), LowestCost: 1 << 30, Verdict: StatusAxiomLedgerCompiled}
	for _, c := range candidates {
		if c.PromotedToTheorem {
			out.PromotedAxioms++
		}
		if c.PureGeometricFormulable {
			out.PureGeometricCandidates++
		}
		if c.EmpiricalIndependentFormulable {
			out.EmpiricalIndependentCount++
		}
		if strings.Contains(c.CurveFittingRisk, "high") || strings.Contains(c.CurveFittingRisk, "maximal") {
			out.CurveFittingRiskHigh++
		}
		if c.CKMCapacity {
			out.CKMCapableCandidates++
		}
		if c.PMNSCapacity {
			out.PMNSCapableCandidates++
		}
		if c.EmpiricalIndependentFormulable && c.MathematicalCost < out.LowestCost {
			out.LowestCost = c.MathematicalCost
			out.LeastCostNames = []string{c.Name}
		} else if c.EmpiricalIndependentFormulable && c.MathematicalCost == out.LowestCost {
			out.LeastCostNames = append(out.LeastCostNames, c.Name)
		}
	}
	return out
}

func buildCapacity(l AxiomLedger) CapacityAudit {
	out := CapacityAudit{Executed: true, Verdict: StatusCKMPMNSCapacityAudited, Reason: "CKM/PMNS needs two noncommuting family operators; current ASHA provides none natively, while family connection/algebra/source axioms would have only conditional capacity."}
	for _, c := range l.Candidates {
		if c.ProvidesDiagonalOperator && !c.ProvidesTwoNoncommutingOps {
			out.CandidatesWithDiagonalOnly++
		}
		if c.ProvidesTwoNoncommutingOps {
			out.CandidatesWithNoncommuting++
			if !c.NativeInCurrentAsha {
				out.ConditionalNoncommutingPairs++
			}
		}
		if c.NativeInCurrentAsha && c.ProvidesTwoNoncommutingOps {
			out.NativeNoncommutingPairs++
		}
		if c.NativeInCurrentAsha && c.CKMCapacity {
			out.CKMNative = true
		}
		if !c.NativeInCurrentAsha && c.CKMCapacity && !c.RequiresEmpiricalYukawas {
			out.CKMConditional = true
		}
		if c.NativeInCurrentAsha && c.PMNSCapacity {
			out.PMNSNative = true
		}
		if !c.NativeInCurrentAsha && c.PMNSCapacity && !c.RequiresEmpiricalYukawas {
			out.PMNSConditional = true
		}
	}
	return out
}

func buildEmpiricalIndependence(l AxiomLedger) EmpiricalIndependenceAudit {
	out := EmpiricalIndependenceAudit{Executed: true, NoObservedMassesImported: true, NoCKMImported: true, NoPMNSImported: true, NoYukawaMatricesInserted: true, Verdict: StatusEmpiricalIndependenceAudited}
	for _, c := range l.Candidates {
		if c.PureGeometricFormulable && c.EmpiricalIndependentFormulable {
			out.CandidatesCanBePureRules++
		}
		if c.RequiresEmpiricalYukawas || !c.EmpiricalIndependentFormulable {
			out.CandidatesCollapseToFitting++
		}
	}
	out.UnconstrainedFamilySourceRisk = out.CandidatesCollapseToFitting > 0
	out.Reason = "The ledger separates pure axiom candidates from unconstrained source matrices; no observed Yukawa, CKM, PMNS, or mass data are imported."
	return out
}

func buildRanking(l AxiomLedger) CostRanking {
	rows := []CostRow{
		{Rank: 1, Name: "minimal modular family Hamiltonian axiom", Cost: 2, Benefit: "hierarchy/three-level density capacity", MainRisk: "diagonal-only unless paired with a second source", NextTest: "derive or constrain K_gen from a pure topological rule"},
		{Rank: 2, Name: "nontrivial U(3)_gen family connection axiom", Cost: 3, Benefit: "native-style CKM/PMNS capacity if curvature is constrained", MainRisk: "unconstrained connection is curve-fitting", NextTest: "quantized holonomy/curvature consistency sieve"},
		{Rank: 3, Name: "triality/contact functor axioms", Cost: 3, Benefit: "threefold semantics linked to existing ASHA structures", MainRisk: "wrong domain or exact triality degeneracy", NextTest: "typed functor compatibility with A_F,J,first-order"},
		{Rank: 4, Name: "primitive ideal family algebra extension", Cost: 4, Benefit: "structural family irreps", MainRisk: "changes A_F and requires rebuilding spectral triple", NextTest: "minimal algebra-extension compatibility audit"},
		{Rank: 5, Name: "unconstrained external Yukawa/source matrix", Cost: 5, Benefit: "phenomenological completeness", MainRisk: "pure curve-fitting", NextTest: "quarantine only; not a derivation route"},
	}
	return CostRanking{Executed: true, Rows: rows, Verdict: StatusAxiomCostRankingAudited}
}

func buildBoundary(l AxiomLedger) Boundary {
	return Boundary{Executed: true, LawSpaceNative: true, FamilyBundleNative: false, NewAxiomRequiredForFamilies: true, CurrentASHAFlavorComplete: false, Verdict: StatusEpistemologicalBoundaryDocumented, Statement: "Current ASHA derives the law-space/gauge-Higgs scaffold but does not derive a nontrivial family bundle. Any reduction of the charged flavor moduli now requires a new explicit axiom or extension, not another hidden use of existing carriers."}
}

func buildModuli(l AxiomLedger, c CapacityAudit) ModuliImpact {
	scenarios := []ModuliScenario{
		{Name: "current ASHA native carrier", Status: StatusFirewallPreserved13Moduli, ModuliDim: Gate372ChargedFlavorModuliDim, Reason: "trivial U(3)_gen multiplicity remains unselected"},
		{Name: "minimal modular Hamiltonian axiom", Status: StatusFailedModularKMSNeedsHamiltonianAxiom, ModuliDim: Gate372ChargedFlavorModuliDim, ThreeDistinctMassesPossible: true, Reason: "conditional diagonal hierarchy only; no native K_gen and no CKM pair"},
		{Name: "nontrivial family connection axiom", Status: StatusFailedFamilyConnectionNeedsAxiom, ModuliDim: Gate372ChargedFlavorModuliDim, ThreeDistinctMassesPossible: true, CKMPossible: true, PMNSPossible: true, Reason: "could reduce moduli if constrained, but no connection axiom is promoted"},
		{Name: "primitive ideal algebra extension", Status: StatusFailedPrimitiveIdealNeedsAlgebraAxiom, ModuliDim: Gate372ChargedFlavorModuliDim, ThreeDistinctMassesPossible: true, CKMPossible: true, PMNSPossible: true, Reason: "requires changing A_F and revalidating the finite triple"},
		{Name: "unconstrained empirical source", Status: StatusFailedUnconstrainedSourceIsCurveFitting, ModuliDim: Gate372ChargedFlavorModuliDim, ThreeDistinctMassesPossible: true, CKMPossible: true, PMNSPossible: true, Reason: "phenomenological fit is quarantined and not counted as derivation"},
	}
	return ModuliImpact{StartDim: Gate372ChargedFlavorModuliDim, Scenarios: scenarios, BestNativeDim: Gate372ChargedFlavorModuliDim, NativeReductionBelow13: false, ConditionalReductionBelow13: c.CKMConditional || c.PMNSConditional, FirewallPreserved: true, Verdict: StatusFirewallPreserved13Moduli}
}

func buildFirewall(m ModuliImpact) FirewallAudit {
	return FirewallAudit{Executed: true, NoObservedMassesImported: true, NoCKMImported: true, NoPMNSImported: true, NoYukawaAmplitudesInserted: true, NoAxiomPromoted: true, NoExternalHamiltonianPromoted: true, NoFamilyConnectionPromoted: true, NoAlgebraExtensionPromoted: true, NoFunctorPromoted: true, NoModuliReductionClaimed: !m.NativeReductionBelow13, Verdict: StatusFirewallPreserved13Moduli}
}

func buildNext(r CostRanking) NextStep {
	return NextStep{Gate: 412, Title: "Minimal Modular Family Hamiltonian Axiom Consistency Sieve", Reason: "Gate 411 ranks the modular family Hamiltonian as the lowest-cost empirical-independent axiom candidate, but it must be tested as an explicit axiom, not promoted as native.", PrimaryTask: "formulate the smallest K_gen axiom, check compatibility with A_F,J,Gamma,first-order,gauge charges, and determine whether diagonal hierarchy alone can be made non-fitting."}
}

func truth(a Analysis) string {
	return "Gate 411 compiles the family-bundle axiom ledger without promoting any new structure. The least-cost empirical-independent candidate is a modular family Hamiltonian axiom, while true CKM/PMNS capacity requires a nontrivial family connection or algebra extension. All such routes are explicit extensions, not native consequences of current ASHA. Therefore the epistemological boundary is documented and the 13 charged flavor moduli firewall remains preserved."
}

func Statuses(a Analysis) []string {
	statuses := []string{StatusGate410Inherited, StatusAxiomLedgerCompiled, StatusEpistemologicalBoundaryDocumented, StatusAxiomCostRankingAudited, StatusCKMPMNSCapacityAudited, StatusEmpiricalIndependenceAudited, StatusLeastCostCandidateIdentified}
	for _, c := range a.Ledger.Candidates {
		addStatus(&statuses, c.Verdict)
	}
	for _, s := range []string{StatusModularHamiltonianCandidateQuarantined, StatusFamilyConnectionCandidateQuarantined, StatusPrimitiveIdealCandidateQuarantined, StatusTrialityLocalSystemCandidateQuarantined, StatusFailedNoAxiomPromoted, StatusFailedNoNativeCKMPMNS, StatusFirewallPreserved13Moduli} {
		addStatus(&statuses, s)
	}
	return statuses
}

func addStatus(xs *[]string, s string) {
	if s != "" && !contains(*xs, s) {
		*xs = append(*xs, s)
	}
}
func contains(xs []string, v string) bool {
	for _, x := range xs {
		if x == v {
			return true
		}
	}
	return false
}
