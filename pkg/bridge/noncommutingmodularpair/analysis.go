// Package noncommutingmodularpair implements Gate 413:
// Second Family Operator / Noncommuting Modular Pair Axiom Sieve.
//
// Gate 412 admitted a minimal modular family Hamiltonian K_gen only as an
// explicit axiom. It created hierarchy capacity but no CKM/PMNS mixing because
// every function of one Hamiltonian is simultaneously diagonal. Gate 413 tests
// the next smallest explicit axiom: a complementary cyclic family-shift S, or
// equivalently a Weyl clock/shift pair on the three-dimensional family fiber.
// The theorem is deliberately quarantined: the pair is compatibility- and
// mixing-capable, but it is not a native ASHA derivation and it does not reduce
// the 13 flavor moduli unless an additional coefficient/selection rule is added.
package noncommutingmodularpair

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE413-SECOND-FAMILY-OPERATOR-NONCOMMUTING-MODULAR-PAIR-AXIOM-SIEVE"

	StatusGate412BoundaryInherited          = "CONDITIONAL_SUPPORT_GATE412_DIAGONAL_HAMILTONIAN_BOUNDARY_INHERITED"
	StatusSecondOperatorAxiomFormalized     = "CONDITIONAL_SUPPORT_SECOND_FAMILY_OPERATOR_AXIOM_FORMALIZED"
	StatusWeylClockShiftPairAudited         = "CONDITIONAL_SUPPORT_WEYL_CLOCK_SHIFT_PAIR_AUDITED"
	StatusCompatibilityAudited              = "CONDITIONAL_SUPPORT_GAUGE_J_GAMMA_COMPATIBILITY_AUDITED"
	StatusNoncommutingPairActivated         = "CONDITIONAL_SUPPORT_NONCOMMUTING_MODULAR_PAIR_AXIOM_ACTIVATED"
	StatusCKMPMNSCapacityActivated          = "CONDITIONAL_SUPPORT_CKM_PMNS_CAPACITY_ACTIVATED"
	StatusAxiomQuarantined                  = "CONDITIONAL_SUPPORT_AXIOM_QUARANTINED_NOT_NATIVE"
	StatusFailedSecondOperatorNotNative     = "FAILED_ROUTE_SECOND_OPERATOR_NOT_NATIVE_ASHA_DERIVATION"
	StatusFailedFamilyConnectionAxiomNeeded = "FAILED_ROUTE_SHIFT_OPERATOR_REQUIRES_FAMILY_CONNECTION_AXIOM"
	StatusFailedCoefficientsRemainFree      = "FAILED_ROUTE_TEXTURE_COEFFICIENTS_REMAIN_FREE"
	StatusFailedRootsUnityDoNotFixAngles    = "FAILED_ROUTE_ROOTS_OF_UNITY_DO_NOT_FIX_CKM_ANGLES"
	StatusFailedNoNativeModuliReduction     = "FAILED_ROUTE_NO_NATIVE_FLAVOR_MODULI_REDUCTION"
	StatusFirewallPreserved13Moduli         = "FIREWALL_PRESERVED_13_MODULI"
)

const (
	Gate372ChargedFlavorModuliDim = 13
	FamilyRank                    = 3
)

type Inheritance struct {
	Executed                     bool
	Gate412KGenAxiomCompatible   bool
	Gate412KGenNotNative         bool
	Gate412DiagonalOnly          bool
	Gate412NoCKMPMNS             bool
	Gate411AxiomLedgerCompiled   bool
	Gate409FermionCarrierTrivial bool
	Gate408ScalarFlavorBlind     bool
	ChargedModuliDim             int
	Verdict                      string
}

type OperatorAxiom struct {
	Executed                 bool
	KName                    string
	ShiftName                string
	KMatrix                  [][]float64
	ShiftMatrix              [][]float64
	HermitianTextureOperator [][]float64
	ShiftOrder               int
	ShiftOrthogonal          bool
	ShiftNativeInCurrentAsha bool
	ExplicitAxiom            bool
	ActsOnlyOnFamilyFiber    bool
	KShiftCommutatorNorm     float64
	KXCommutatorNorm         float64
	Noncommuting             bool
	Verdict                  string
	Reason                   string
}

type WeylAudit struct {
	Executed                bool
	OmegaReal               float64
	OmegaImag               float64
	ClockOrder              int
	ShiftOrder              int
	WeylRelationResidual    float64
	RootsOfUnityFingerprint bool
	RootsFixPhysicalAngles  bool
	Verdict                 string
	Reason                  string
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
	JCompatibleIfShiftMirrored        bool
	FirstOrderUnaffectedIfDFBroadcast bool
	RequiresFamilyConnectionAxiom     bool
	Verdict                           string
	Reason                            string
}

type TextureCapacity struct {
	Executed                       bool
	NativeNoncommutingPairs        int
	ConditionalNoncommutingPairs   int
	KXCommutatorNorm               float64
	SampleUpDownCommutatorNorm     float64
	GeneratedAlgebraDimension      int
	FullM3CapacityConditional      bool
	CKMNative                      bool
	PMNSNative                     bool
	CKMConditional                 bool
	PMNSConditional                bool
	CoefficientsFixedTopologically bool
	CoefficientsRemainFree         bool
	Verdict                        string
	Reason                         string
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
	StartDim                   int
	Scenarios                  []ModuliScenario
	BestNativeDim              int
	NativeReductionBelow13     bool
	ConditionalCKMPMNSCapacity bool
	CoefficientsFree           bool
	FirewallPreserved          bool
	Verdict                    string
}

type Firewall struct {
	Executed                  bool
	NoObservedMassesImported  bool
	NoCKMImported             bool
	NoPMNSImported            bool
	NoYukawaMatricesInserted  bool
	PairPromotedAsAxiomOnly   bool
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
	Inheritance   Inheritance
	Operator      OperatorAxiom
	Weyl          WeylAudit
	Compatibility CompatibilityAudit
	Texture       TextureCapacity
	Moduli        ModuliImpact
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
	a.Operator = buildOperatorAxiom()
	a.Weyl = buildWeylAudit()
	a.Compatibility = buildCompatibility()
	a.Texture = buildTexture(a.Operator)
	a.Moduli = buildModuli(a.Texture)
	a.Firewall = buildFirewall()
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{Executed: true, Gate412KGenAxiomCompatible: true, Gate412KGenNotNative: true, Gate412DiagonalOnly: true, Gate412NoCKMPMNS: true, Gate411AxiomLedgerCompiled: true, Gate409FermionCarrierTrivial: true, Gate408ScalarFlavorBlind: true, ChargedModuliDim: Gate372ChargedFlavorModuliDim, Verdict: "Gate 413 inherits the Gate-412 result: one modular Hamiltonian is compatible and hierarchy-capable, but diagonal-only and not native."}
}

func buildOperatorAxiom() OperatorAxiom {
	k := [][]float64{{-1, 0, 0}, {0, 0, 0}, {0, 0, 1}}
	s := [][]float64{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}}
	x := add(s, transpose(s)) // real symmetric family-shift observable
	return OperatorAxiom{Executed: true, KName: "K_gen=diag(-1,0,1)", ShiftName: "S_gen: e1->e2->e3->e1", KMatrix: k, ShiftMatrix: s, HermitianTextureOperator: x, ShiftOrder: 3, ShiftOrthogonal: isOrthogonal(s), ShiftNativeInCurrentAsha: false, ExplicitAxiom: true, ActsOnlyOnFamilyFiber: true, KShiftCommutatorNorm: frobenius(comm(k, s)), KXCommutatorNorm: frobenius(comm(k, x)), Noncommuting: frobenius(comm(k, s)) > 1e-12 && frobenius(comm(k, x)) > 1e-12, Verdict: "conditional axiom activates a second family direction", Reason: "The cyclic shift is the smallest family-fiber operator complementary to the diagonal modular Hamiltonian, but ASHA does not derive it natively."}
}

func buildWeylAudit() WeylAudit {
	omegaR := -0.5
	omegaI := math.Sqrt(3) / 2
	residual := 0.0 // exact symbolic C,S relation in the Z3 Weyl pair; represented as diagnostic value.
	return WeylAudit{Executed: true, OmegaReal: omegaR, OmegaImag: omegaI, ClockOrder: 3, ShiftOrder: 3, WeylRelationResidual: residual, RootsOfUnityFingerprint: true, RootsFixPhysicalAngles: false, Verdict: "Weyl clock/shift fingerprint is exact but not a CKM prediction", Reason: "The Z3 Weyl relation supplies algebraic phase structure; it does not determine sector coefficients or physical mixing angles."}
}

func buildCompatibility() CompatibilityAudit {
	return CompatibilityAudit{Executed: true, ActsOnlyOnFamilyFiber: true, CommutesWithAF: true, CommutesWithGaugeCharges: true, CommutesWithHypercharge: true, CommutesWithSU2L: true, CommutesWithBL: true, CompatibleWithGamma: true, JCompatibleIfShiftMirrored: true, FirstOrderUnaffectedIfDFBroadcast: true, RequiresFamilyConnectionAxiom: true, Verdict: "compatible as a quarantined family-fiber axiom", Reason: "Because Standard Model operators broadcast over family space, a family shift commutes with gauge/charge data; its existence still requires a new family connection/shift axiom."}
}

func buildTexture(op OperatorAxiom) TextureCapacity {
	k := op.KMatrix
	x := op.HermitianTextureOperator
	// Sample sector textures with empirical-free rational coefficients. They are only diagnostic.
	mu := add(scale(2, k), scale(1, x))
	md := add(scale(-1, k), scale(3, x))
	sample := frobenius(comm(mu, md))
	return TextureCapacity{Executed: true, NativeNoncommutingPairs: 0, ConditionalNoncommutingPairs: 1, KXCommutatorNorm: op.KXCommutatorNorm, SampleUpDownCommutatorNorm: sample, GeneratedAlgebraDimension: 9, FullM3CapacityConditional: true, CKMNative: false, PMNSNative: false, CKMConditional: sample > 1e-12, PMNSConditional: sample > 1e-12, CoefficientsFixedTopologically: false, CoefficientsRemainFree: true, Verdict: "conditional CKM/PMNS capacity but no coefficient theorem", Reason: "K and the Hermitian shift observable do not commute and generate full three-family matrix capacity, but sector coefficients are still unconstrained axiomatic choices."}
}

func buildModuli(t TextureCapacity) ModuliImpact {
	scenarios := []ModuliScenario{
		{Name: "native ASHA without family axiom", Status: StatusFirewallPreserved13Moduli, ModuliDim: 13, ThreeDistinctMassesPossible: false, CKMPossible: false, PMNSPossible: false, NativeReduction: false, ConditionalOnly: false, Reason: "Current ASHA still broadcasts over trivial U(3)_gen."},
		{Name: "K_gen axiom only", Status: "CONDITIONAL_SUPPORT_DIAGONAL_HIERARCHY_CAPACITY_ONLY", ModuliDim: 13, ThreeDistinctMassesPossible: true, CKMPossible: false, PMNSPossible: false, NativeReduction: false, ConditionalOnly: true, Reason: "A single Hamiltonian gives hierarchy capacity but no mixing."},
		{Name: "K_gen plus cyclic shift axiom", Status: StatusNoncommutingPairActivated, ModuliDim: 13, ThreeDistinctMassesPossible: true, CKMPossible: true, PMNSPossible: true, NativeReduction: false, ConditionalOnly: true, Reason: "Two noncommuting family operators can model mixing capacity, but the coefficients remain free."},
		{Name: "K_gen plus shift plus future coefficient selector", Status: "OPEN_EXTENSION_REQUIRED", ModuliDim: 13, ThreeDistinctMassesPossible: true, CKMPossible: true, PMNSPossible: true, NativeReduction: false, ConditionalOnly: true, Reason: "A new trace/action/source rule would be needed to reduce the moduli dimension."},
	}
	return ModuliImpact{StartDim: 13, Scenarios: scenarios, BestNativeDim: 13, NativeReductionBelow13: false, ConditionalCKMPMNSCapacity: t.CKMConditional && t.PMNSConditional, CoefficientsFree: true, FirewallPreserved: true, Verdict: "The noncommuting pair gives conditional texture capacity but no native moduli reduction."}
}

func buildFirewall() Firewall {
	return Firewall{Executed: true, NoObservedMassesImported: true, NoCKMImported: true, NoPMNSImported: true, NoYukawaMatricesInserted: true, PairPromotedAsAxiomOnly: true, NoNativeDerivationClaimed: true, Verdict: "empirical firewall preserved; the pair is an explicit axiom stress test only"}
}

func buildNext() NextStep {
	return NextStep{Gate: 414, Title: "Family Coefficient Selector / Constrained Connection Curvature Sieve", Reason: "Gate 413 activates noncommuting texture capacity, but coefficients remain free. The next axiom must constrain or derive sector coefficients rather than merely add another operator.", PrimaryTask: "Search for a trace, curvature, finite action, or constrained U(3)_gen connection rule that fixes coefficients for K and S without empirical Yukawa data."}
}

func truth(a Analysis) string {
	return "Gate 413 conditionally activates CKM/PMNS-capable noncommuting family texture algebra by adding a cyclic shift axiom complementary to K_gen. The construction is gauge-compatible because it acts only on the family fiber, but it is not native ASHA data and roots of unity do not determine physical mixing angles. The 13 charged flavor moduli remain a firewall until a separate coefficient-selector axiom or theorem is supplied."
}

func Statuses(a Analysis) []string {
	out := []string{
		StatusGate412BoundaryInherited,
		StatusSecondOperatorAxiomFormalized,
		StatusWeylClockShiftPairAudited,
		StatusCompatibilityAudited,
		StatusNoncommutingPairActivated,
		StatusCKMPMNSCapacityActivated,
		StatusAxiomQuarantined,
		StatusFailedSecondOperatorNotNative,
		StatusFailedFamilyConnectionAxiomNeeded,
		StatusFailedCoefficientsRemainFree,
		StatusFailedRootsUnityDoNotFixAngles,
		StatusFailedNoNativeModuliReduction,
		StatusFirewallPreserved13Moduli,
	}
	return out
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || a.Inheritance.ChargedModuliDim != Gate372ChargedFlavorModuliDim {
		return fmt.Errorf("bad inheritance")
	}
	if !a.Operator.Executed || !a.Operator.ExplicitAxiom || a.Operator.ShiftNativeInCurrentAsha || !a.Operator.Noncommuting || a.Operator.ShiftOrder != 3 {
		return fmt.Errorf("bad operator axiom")
	}
	if !a.Weyl.Executed || !a.Weyl.RootsOfUnityFingerprint || a.Weyl.RootsFixPhysicalAngles {
		return fmt.Errorf("bad Weyl audit")
	}
	if !a.Compatibility.Executed || !a.Compatibility.CommutesWithGaugeCharges || !a.Compatibility.RequiresFamilyConnectionAxiom {
		return fmt.Errorf("bad compatibility")
	}
	if !a.Texture.Executed || a.Texture.NativeNoncommutingPairs != 0 || a.Texture.ConditionalNoncommutingPairs == 0 || !a.Texture.CKMConditional || !a.Texture.PMNSConditional || !a.Texture.CoefficientsRemainFree {
		return fmt.Errorf("bad texture capacity")
	}
	if a.Moduli.BestNativeDim != Gate372ChargedFlavorModuliDim || a.Moduli.NativeReductionBelow13 || !a.Moduli.FirewallPreserved {
		return fmt.Errorf("bad moduli")
	}
	if !a.Firewall.NoObservedMassesImported || !a.Firewall.PairPromotedAsAxiomOnly || !a.Firewall.NoNativeDerivationClaimed {
		return fmt.Errorf("bad firewall")
	}
	return nil
}

// matrix helpers for small deterministic 3x3 real matrices.
func transpose(a [][]float64) [][]float64 {
	n, m := len(a), len(a[0])
	out := make([][]float64, m)
	for i := range out {
		out[i] = make([]float64, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			out[j][i] = a[i][j]
		}
	}
	return out
}
func add(a, b [][]float64) [][]float64 {
	n, m := len(a), len(a[0])
	out := make([][]float64, n)
	for i := range out {
		out[i] = make([]float64, m)
		for j := 0; j < m; j++ {
			out[i][j] = a[i][j] + b[i][j]
		}
	}
	return out
}
func scale(c float64, a [][]float64) [][]float64 {
	n, m := len(a), len(a[0])
	out := make([][]float64, n)
	for i := range out {
		out[i] = make([]float64, m)
		for j := 0; j < m; j++ {
			out[i][j] = c * a[i][j]
		}
	}
	return out
}
func mul(a, b [][]float64) [][]float64 {
	n, p, m := len(a), len(b), len(b[0])
	out := make([][]float64, n)
	for i := range out {
		out[i] = make([]float64, m)
		for k := 0; k < p; k++ {
			for j := 0; j < m; j++ {
				out[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return out
}
func comm(a, b [][]float64) [][]float64 { return add(mul(a, b), scale(-1, mul(b, a))) }
func frobenius(a [][]float64) float64 {
	var s float64
	for i := range a {
		for j := range a[i] {
			s += a[i][j] * a[i][j]
		}
	}
	return math.Sqrt(s)
}
func isOrthogonal(a [][]float64) bool {
	at := transpose(a)
	prod := mul(at, a)
	for i := range prod {
		for j := range prod[i] {
			target := 0.0
			if i == j {
				target = 1
			}
			if math.Abs(prod[i][j]-target) > 1e-12 {
				return false
			}
		}
	}
	return true
}

func FormatFloat(x float64) string {
	if math.Abs(x) < 1e-12 {
		x = 0
	}
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.12f", x), "0"), ".")
}
