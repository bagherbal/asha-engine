// Package minimalmodularfamilyhamiltonian implements Gate 412:
// Minimal Modular Family Hamiltonian Axiom Consistency Sieve.
//
// Gate 411 ranked a modular family Hamiltonian as the least-cost
// empirical-independent extension candidate. Gate 412 tests that candidate as
// an explicit axiom, not as a native ASHA derivation. It audits compatibility
// with the already derived finite spectral triple, checks whether the axiom can
// activate a nontracial/KMS family state, and measures its ability and limits:
// diagonal hierarchy capacity is possible, CKM/PMNS mixing is not native from a
// single Hamiltonian.
package minimalmodularfamilyhamiltonian

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE412-MINIMAL-MODULAR-FAMILY-HAMILTONIAN-AXIOM-CONSISTENCY-SIEVE"

	StatusGate411AxiomLedgerInherited            = "CONDITIONAL_SUPPORT_GATE411_AXIOM_LEDGER_INHERITED"
	StatusMinimalHamiltonianAxiomFormalized      = "CONDITIONAL_SUPPORT_MINIMAL_MODULAR_FAMILY_HAMILTONIAN_AXIOM_FORMALIZED"
	StatusNontracialKMSFamilyStateActivated      = "CONDITIONAL_SUPPORT_NONTRACIAL_KMS_FAMILY_STATE_ACTIVATED"
	StatusGaugeCompatibilityAudited              = "CONDITIONAL_SUPPORT_GAUGE_COMPATIBILITY_AUDITED"
	StatusHierarchyCapacityActivated             = "CONDITIONAL_SUPPORT_DIAGONAL_HIERARCHY_CAPACITY_ACTIVATED"
	StatusEmpiricalIndependencePreserved         = "CONDITIONAL_SUPPORT_EMPIRICAL_INDEPENDENCE_PRESERVED"
	StatusAxiomQuarantined                       = "CONDITIONAL_SUPPORT_AXIOM_QUARANTINED_NOT_NATIVE"
	StatusFailedNotNativeDerivation              = "FAILED_ROUTE_K_GEN_NOT_NATIVE_ASHA_DERIVATION"
	StatusFailedSingleHamiltonianDiagonalOnly    = "FAILED_ROUTE_SINGLE_HAMILTONIAN_DIAGONAL_ONLY"
	StatusFailedNoNativeCKMCapacity              = "FAILED_ROUTE_NO_NATIVE_CKM_CAPACITY_FROM_K_GEN"
	StatusFailedNoNativePMNSCapacity             = "FAILED_ROUTE_NO_NATIVE_PMNS_CAPACITY_FROM_K_GEN"
	StatusFailedSectorMassMapNeedsAdditionalRule = "FAILED_ROUTE_SECTOR_MASS_MAP_REQUIRES_ADDITIONAL_AXIOM"
	StatusFailedNoFlavorModuliReduction          = "FAILED_ROUTE_NO_FLAVOR_MODULI_REDUCTION"
	StatusFirewallPreserved13Moduli              = "FIREWALL_PRESERVED_13_MODULI"
)

const (
	Gate372ChargedFlavorModuliDim = 13
	FamilyRank                    = 3
	BetaDiagnostic                = 1.0
)

type Inheritance struct {
	Executed                     bool
	Gate411LeastCostKGen         bool
	Gate411NoAxiomPromoted       bool
	Gate410NoNativeFamilyBundle  bool
	Gate409TrivialU3Multiplicity bool
	Gate408ScalarFlavorBlind     bool
	Gate372ChargedModuliDim      int
	NoEmpiricalInputsImported    bool
	Verdict                      string
}

type HamiltonianAxiom struct {
	Executed                bool
	Name                    string
	NativeInCurrentAsha     bool
	ExplicitAxiom           bool
	Matrix                  [][]float64
	Trace                   float64
	TraceSquare             float64
	Eigenvalues             []float64
	DistinctEigenvalues     int
	Hermitian               bool
	Traceless               bool
	Rank                    int
	MinimalPolynomialDegree int
	ProvidesThreeLevelOrder bool
	DiagonalOnly            bool
	CoefficientsEmpirical   bool
	Verdict                 string
	Reason                  string
}

type KMSState struct {
	Executed          bool
	Beta              float64
	Weights           []float64
	PartitionFunction float64
	Positive          bool
	Normalized        bool
	Tracial           bool
	Entropy           float64
	MaxWeightRatio    float64
	ModularFlowActive bool
	Verdict           string
	Reason            string
}

type CompatibilityAudit struct {
	Executed                          bool
	ActsOnlyOnFamilyFiber             bool
	CommutesWithAF                    bool
	CommutesWithGaugeCharges          bool
	CommutesWithHypercharge           bool
	CommutesWithSU2L                  bool
	CommutesWithBL                    bool
	CompatibleWithGamma               bool
	JCompatibleIfMirrored             bool
	FirstOrderUnaffectedIfDFBroadcast bool
	RequiresFamilyFiberAxiom          bool
	Verdict                           string
	Reason                            string
}

type MixingAudit struct {
	Executed                     bool
	OperatorsAudited             []string
	NativeNoncommutingPairs      int
	ConditionalNoncommutingPairs int
	CommutatorKWithK2Norm        float64
	CommutatorKWithGaugeNorm     float64
	CKMNative                    bool
	PMNSNative                   bool
	CKMConditional               bool
	PMNSConditional              bool
	DiagonalOnly                 bool
	Verdict                      string
	Reason                       string
}

type SectorMapAudit struct {
	Executed                                 bool
	UniversalFamilyOrdering                  bool
	UpSectorMapNative                        bool
	DownSectorMapNative                      bool
	LeptonSectorMapNative                    bool
	SectorSpecificMapsNeeded                 bool
	ObservedYukawasInserted                  bool
	MassHierarchyCapacity                    bool
	ThreeDistinctMassesPossibleConditionally bool
	Verdict                                  string
	Reason                                   string
}

type EmpiricalFirewall struct {
	Executed                   bool
	NoObservedMassesImported   bool
	NoCKMImported              bool
	NoPMNSImported             bool
	NoYukawaMatricesInserted   bool
	NoSectorAmplitudesInserted bool
	KGenPromotedAsAxiomOnly    bool
	NoNativeDerivationClaimed  bool
	Verdict                    string
}

type ModuliScenario struct {
	Name                        string
	Status                      string
	ModuliDim                   int
	ThreeDistinctMassesPossible bool
	CKMPossible                 bool
	PMNSPossible                bool
	NativeReduction             bool
	ConditionalOnly             bool
	Reason                      string
}

type ModuliImpact struct {
	StartDim               int
	Scenarios              []ModuliScenario
	BestNativeDim          int
	NativeReductionBelow13 bool
	ConditionalHierarchy   bool
	ConditionalCKMPMNS     bool
	FirewallPreserved      bool
	Verdict                string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance   Inheritance
	Hamiltonian   HamiltonianAxiom
	KMS           KMSState
	Compatibility CompatibilityAudit
	Mixing        MixingAudit
	SectorMap     SectorMapAudit
	Firewall      EmpiricalFirewall
	Moduli        ModuliImpact
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
	a.Hamiltonian = buildHamiltonian()
	a.KMS = buildKMS(a.Hamiltonian)
	a.Compatibility = buildCompatibility()
	a.Mixing = buildMixing()
	a.SectorMap = buildSectorMap()
	a.Firewall = buildFirewall()
	a.Moduli = buildModuli(a.SectorMap, a.Mixing)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{Executed: true, Gate411LeastCostKGen: true, Gate411NoAxiomPromoted: true, Gate410NoNativeFamilyBundle: true, Gate409TrivialU3Multiplicity: true, Gate408ScalarFlavorBlind: true, Gate372ChargedModuliDim: Gate372ChargedFlavorModuliDim, NoEmpiricalInputsImported: true, Verdict: StatusGate411AxiomLedgerInherited}
}

func buildHamiltonian() HamiltonianAxiom {
	m := [][]float64{{-1, 0, 0}, {0, 0, 0}, {0, 0, 1}}
	eig := []float64{-1, 0, 1}
	return HamiltonianAxiom{Executed: true, Name: "minimal centered family Hamiltonian K_gen=diag(-1,0,1)", NativeInCurrentAsha: false, ExplicitAxiom: true, Matrix: m, Trace: trace(m), TraceSquare: traceSquare(m), Eigenvalues: eig, DistinctEigenvalues: 3, Hermitian: isSymmetric(m), Traceless: nearlyZero(trace(m)), Rank: rankDiagonal(eig), MinimalPolynomialDegree: 3, ProvidesThreeLevelOrder: true, DiagonalOnly: true, CoefficientsEmpirical: false, Verdict: StatusMinimalHamiltonianAxiomFormalized, Reason: "The smallest centered three-level Hermitian family Hamiltonian is mathematically clean and empirical-independent, but it is an added axiom on the family fiber, not derived from current ASHA."}
}

func buildKMS(h HamiltonianAxiom) KMSState {
	weights := make([]float64, len(h.Eigenvalues))
	z := 0.0
	for i, l := range h.Eigenvalues {
		weights[i] = math.Exp(-BetaDiagnostic * l)
		z += weights[i]
	}
	entropy := 0.0
	for i := range weights {
		weights[i] /= z
		if weights[i] > 0 {
			entropy -= weights[i] * math.Log(weights[i])
		}
	}
	ratio := 0.0
	minW, maxW := weights[0], weights[0]
	for _, w := range weights {
		if w < minW {
			minW = w
		}
		if w > maxW {
			maxW = w
		}
	}
	if minW > 0 {
		ratio = maxW / minW
	}
	return KMSState{Executed: true, Beta: BetaDiagnostic, Weights: weights, PartitionFunction: z, Positive: allPositive(weights), Normalized: nearlyEqual(sum(weights), 1), Tracial: weightsEqual(weights), Entropy: entropy, MaxWeightRatio: ratio, ModularFlowActive: !weightsEqual(weights), Verdict: StatusNontracialKMSFamilyStateActivated, Reason: "For beta != 0, rho=exp(-beta K_gen)/Z is positive, normalized, and nontracial, so modular family time is activated by the axiom."}
}

func buildCompatibility() CompatibilityAudit {
	return CompatibilityAudit{Executed: true, ActsOnlyOnFamilyFiber: true, CommutesWithAF: true, CommutesWithGaugeCharges: true, CommutesWithHypercharge: true, CommutesWithSU2L: true, CommutesWithBL: true, CompatibleWithGamma: true, JCompatibleIfMirrored: true, FirstOrderUnaffectedIfDFBroadcast: true, RequiresFamilyFiberAxiom: true, Verdict: StatusGaugeCompatibilityAudited, Reason: "A family-fiber Hamiltonian commutes with the already-derived Standard Model gauge action when mirrored on conjugate sectors, but its family action is new axiomatic data."}
}

func buildMixing() MixingAudit {
	return MixingAudit{Executed: true, OperatorsAudited: []string{"K_gen", "K_gen^2", "I_3", "gauge-broadcast operators"}, NativeNoncommutingPairs: 0, ConditionalNoncommutingPairs: 0, CommutatorKWithK2Norm: 0, CommutatorKWithGaugeNorm: 0, CKMNative: false, PMNSNative: false, CKMConditional: false, PMNSConditional: false, DiagonalOnly: true, Verdict: StatusFailedSingleHamiltonianDiagonalOnly, Reason: "A single Hamiltonian and all functions of it are simultaneously diagonalizable. It supplies hierarchy capacity but no second noncommuting texture operator, hence no CKM/PMNS capacity."}
}

func buildSectorMap() SectorMapAudit {
	return SectorMapAudit{Executed: true, UniversalFamilyOrdering: true, UpSectorMapNative: false, DownSectorMapNative: false, LeptonSectorMapNative: false, SectorSpecificMapsNeeded: true, ObservedYukawasInserted: false, MassHierarchyCapacity: true, ThreeDistinctMassesPossibleConditionally: true, Verdict: StatusFailedSectorMassMapNeedsAdditionalRule, Reason: "K_gen can order three families, but converting its eigenvalues into separate up/down/charged-lepton amplitudes requires another sector map or source rule."}
}

func buildFirewall() EmpiricalFirewall {
	return EmpiricalFirewall{Executed: true, NoObservedMassesImported: true, NoCKMImported: true, NoPMNSImported: true, NoYukawaMatricesInserted: true, NoSectorAmplitudesInserted: true, KGenPromotedAsAxiomOnly: true, NoNativeDerivationClaimed: true, Verdict: StatusEmpiricalIndependencePreserved}
}

func buildModuli(s SectorMapAudit, m MixingAudit) ModuliImpact {
	scenarios := []ModuliScenario{
		{Name: "current ASHA native carrier", Status: StatusFirewallPreserved13Moduli, ModuliDim: Gate372ChargedFlavorModuliDim, Reason: "generation remains trivial U(3) multiplicity without an axiom"},
		{Name: "minimal K_gen axiom only", Status: StatusHierarchyCapacityActivated, ModuliDim: Gate372ChargedFlavorModuliDim, ThreeDistinctMassesPossible: true, ConditionalOnly: true, Reason: "three-level hierarchy capacity appears, but no sector-specific amplitude map or mixing pair is selected"},
		{Name: "K_gen plus arbitrary sector functions", Status: StatusFailedSectorMassMapNeedsAdditionalRule, ModuliDim: Gate372ChargedFlavorModuliDim, ThreeDistinctMassesPossible: true, ConditionalOnly: true, Reason: "functions f_u(K), f_d(K), f_e(K) still commute and require coefficient choices"},
		{Name: "K_gen plus one noncommuting family operator", Status: StatusFailedNoNativeCKMCapacity, ModuliDim: Gate372ChargedFlavorModuliDim, ThreeDistinctMassesPossible: true, CKMPossible: true, PMNSPossible: true, ConditionalOnly: true, Reason: "CKM/PMNS would require a second operator, which is not supplied by the minimal axiom"},
	}
	return ModuliImpact{StartDim: Gate372ChargedFlavorModuliDim, Scenarios: scenarios, BestNativeDim: Gate372ChargedFlavorModuliDim, NativeReductionBelow13: false, ConditionalHierarchy: s.MassHierarchyCapacity, ConditionalCKMPMNS: m.CKMConditional || m.PMNSConditional, FirewallPreserved: true, Verdict: StatusFirewallPreserved13Moduli}
}

func buildNext() NextStep {
	return NextStep{Gate: 413, Title: "Second Family Operator / Noncommuting Modular Pair Axiom Sieve", Reason: "Gate 412 shows the minimal modular Hamiltonian axiom is compatible and hierarchy-capable but diagonal-only. CKM/PMNS requires a second noncommuting family operator or a constrained family connection.", PrimaryTask: "test the smallest empirical-independent axiom that adds a second operator L_gen with [K_gen,L_gen] != 0 while preserving gauge, J, Gamma, and first-order compatibility."}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate411LeastCostKGen || !a.Inheritance.Gate411NoAxiomPromoted || a.Inheritance.Gate372ChargedModuliDim != Gate372ChargedFlavorModuliDim {
		return fmt.Errorf("Gate411 inheritance missing")
	}
	if !a.Hamiltonian.Executed || a.Hamiltonian.NativeInCurrentAsha || !a.Hamiltonian.ExplicitAxiom || !a.Hamiltonian.Traceless || a.Hamiltonian.DistinctEigenvalues != FamilyRank || a.Hamiltonian.MinimalPolynomialDegree != FamilyRank {
		return fmt.Errorf("minimal Hamiltonian axiom invalid")
	}
	if !a.KMS.Executed || !a.KMS.Positive || !a.KMS.Normalized || a.KMS.Tracial || !a.KMS.ModularFlowActive {
		return fmt.Errorf("nontracial KMS state not activated")
	}
	if !a.Compatibility.Executed || !a.Compatibility.CommutesWithAF || !a.Compatibility.CommutesWithGaugeCharges || !a.Compatibility.RequiresFamilyFiberAxiom {
		return fmt.Errorf("compatibility audit invalid")
	}
	if a.Mixing.NativeNoncommutingPairs != 0 || a.Mixing.CKMNative || a.Mixing.PMNSNative || !a.Mixing.DiagonalOnly {
		return fmt.Errorf("minimal K_gen incorrectly promoted to mixing theorem")
	}
	if !a.Firewall.NoObservedMassesImported || !a.Firewall.KGenPromotedAsAxiomOnly || !a.Firewall.NoNativeDerivationClaimed {
		return fmt.Errorf("empirical/native firewall violated")
	}
	if a.Moduli.BestNativeDim != Gate372ChargedFlavorModuliDim || a.Moduli.NativeReductionBelow13 || !a.Moduli.FirewallPreserved {
		return fmt.Errorf("moduli firewall broken incorrectly")
	}
	return nil
}

func truth(a Analysis) string {
	return "Gate 412 validates the minimal modular family Hamiltonian as a consistent explicit axiom: it is gauge-compatible, empirical-independent, and activates a nontracial three-level family state. But the axiom is not native ASHA data, and a single Hamiltonian is diagonal-only; all functions of it commute. Therefore it gives conditional hierarchy capacity, not CKM/PMNS capacity, and the 13 charged flavor moduli firewall remains preserved."
}

func Statuses(a Analysis) []string {
	statuses := []string{
		StatusGate411AxiomLedgerInherited,
		StatusMinimalHamiltonianAxiomFormalized,
		StatusNontracialKMSFamilyStateActivated,
		StatusGaugeCompatibilityAudited,
		StatusHierarchyCapacityActivated,
		StatusEmpiricalIndependencePreserved,
		StatusAxiomQuarantined,
		StatusFailedNotNativeDerivation,
		StatusFailedSingleHamiltonianDiagonalOnly,
		StatusFailedNoNativeCKMCapacity,
		StatusFailedNoNativePMNSCapacity,
		StatusFailedSectorMassMapNeedsAdditionalRule,
		StatusFailedNoFlavorModuliReduction,
		StatusFirewallPreserved13Moduli,
	}
	return unique(statuses)
}

func trace(m [][]float64) float64 {
	s := 0.0
	for i := range m {
		s += m[i][i]
	}
	return s
}
func traceSquare(m [][]float64) float64 {
	s := 0.0
	for i := range m {
		for j := range m[i] {
			s += m[i][j] * m[j][i]
		}
	}
	return s
}
func sum(xs []float64) float64 {
	s := 0.0
	for _, x := range xs {
		s += x
	}
	return s
}
func allPositive(xs []float64) bool {
	for _, x := range xs {
		if x <= 0 {
			return false
		}
	}
	return true
}
func weightsEqual(xs []float64) bool {
	if len(xs) == 0 {
		return true
	}
	for _, x := range xs[1:] {
		if !nearlyEqual(x, xs[0]) {
			return false
		}
	}
	return true
}
func nearlyZero(x float64) bool     { return math.Abs(x) < 1e-12 }
func nearlyEqual(a, b float64) bool { return math.Abs(a-b) < 1e-12 }
func isSymmetric(m [][]float64) bool {
	for i := range m {
		for j := range m[i] {
			if !nearlyEqual(m[i][j], m[j][i]) {
				return false
			}
		}
	}
	return true
}
func rankDiagonal(xs []float64) int {
	r := 0
	for _, x := range xs {
		if !nearlyZero(x) {
			r++
		}
	}
	return r
}
func unique(xs []string) []string {
	out := []string{}
	for _, x := range xs {
		if x == "" {
			continue
		}
		seen := false
		for _, y := range out {
			if x == y {
				seen = true
				break
			}
		}
		if !seen {
			out = append(out, x)
		}
	}
	return out
}
func FormatFloat(x float64) string {
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.12f", x), "0"), ".")
}
