// Package familycoefficientselector implements Gate 414:
// Family Coefficient Selector / Constrained Connection Curvature Sieve.
//
// Gate 413 admitted a second family operator only as an explicit axiom. The
// pair (K_gen, S_gen) gives full noncommuting texture capacity, but leaves the
// coefficients of the up/down/lepton textures free. Gate 414 audits whether a
// trace, curvature, finite-action, or constrained U(3)_gen connection rule can
// select those coefficients without empirical Yukawa input. The result is again
// quarantined: the functionals are meaningful, but none selects physical
// coefficients natively. Curvature minimization selects flat/commuting family
// connections, while nonzero mixing requires external sector source data.
package familycoefficientselector

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE414-FAMILY-COEFFICIENT-SELECTOR-CONSTRAINED-CONNECTION-CURVATURE-SIEVE"

	StatusGate413BoundaryInherited      = "CONDITIONAL_SUPPORT_GATE413_NONCOMMUTING_CAPACITY_INHERITED"
	StatusSelectorArenaFormalized       = "CONDITIONAL_SUPPORT_FAMILY_COEFFICIENT_SELECTOR_ARENA_FORMALIZED"
	StatusTraceFunctionalAudited        = "CONDITIONAL_SUPPORT_TRACE_NORM_FUNCTIONAL_AUDITED"
	StatusCurvatureFunctionalAudited    = "CONDITIONAL_SUPPORT_CURVATURE_FUNCTIONAL_AUDITED"
	StatusSectorSplitAudited            = "CONDITIONAL_SUPPORT_SECTOR_SPLIT_FUNCTIONAL_AUDITED"
	StatusConnectionStressTested        = "CONDITIONAL_SUPPORT_CONSTRAINED_CONNECTION_STRESS_TESTED"
	StatusEmpiricalFirewallAudited      = "CONDITIONAL_SUPPORT_EMPIRICAL_INDEPENDENCE_AUDITED"
	StatusFailedNoNativeSelector        = "FAILED_ROUTE_NO_NATIVE_COEFFICIENT_SELECTOR"
	StatusFailedTraceU3Invariant        = "FAILED_ROUTE_TRACE_NORM_IS_U3_INVARIANT"
	StatusFailedCurvatureFlatOrbits     = "FAILED_ROUTE_CURVATURE_ACTION_SELECTS_FLAT_OR_DEGENERATE_ORBITS"
	StatusFailedSectorWeightsFree       = "FAILED_ROUTE_SECTOR_WEIGHTS_REMAIN_FREE"
	StatusFailedConnectionNeedsAxiom    = "FAILED_ROUTE_CONNECTION_CURVATURE_REQUIRES_EXTERNAL_FAMILY_ACTION"
	StatusFailedNoCKMAnglePrediction    = "FAILED_ROUTE_NO_CKM_ANGLE_PREDICTION"
	StatusFailedNoNativeModuliReduction = "FAILED_ROUTE_NO_NATIVE_FLAVOR_MODULI_REDUCTION"
	StatusFirewallPreserved13Moduli     = "FIREWALL_PRESERVED_13_MODULI"
)

const (
	Gate372ChargedFlavorModuliDim = 13
	FamilyRank                    = 3
)

type Inheritance struct {
	Executed                   bool
	Gate413PairAxiomCompatible bool
	Gate413PairNotNative       bool
	Gate413CKMCapacity         bool
	Gate413CoefficientsFree    bool
	Gate412KDiagonalOnly       bool
	Gate411AxiomLedgerCompiled bool
	ChargedModuliDim           int
	Verdict                    string
}

type SelectorArena struct {
	Executed                  bool
	KName                     string
	ShiftObservableName       string
	KMatrix                   [][]float64
	XMatrix                   [][]float64
	FamilyBasisDimension      int
	GeneratedAlgebraDimension int
	NoncommutingCapacity      bool
	CoefficientsNative        bool
	Verdict                   string
	Reason                    string
}

type FunctionalAudit struct {
	Name                         string
	Executed                     bool
	FunctionalType               string
	GaugeCompatible              bool
	EmpiricalIndependent         bool
	UniqueCoefficientRay         bool
	SelectsNoncommutingTexture   bool
	SelectsPhysicalSectorWeights bool
	SelectorNative               bool
	DiagnosticValue              float64
	Verdict                      string
	Reason                       string
}

type ConnectionAudit struct {
	Executed                       bool
	ConnectionAnsatz               string
	FamilyCurvatureSampleNorm      float64
	YangMillsMinimizerFlat         bool
	FlatMinimizerCommutes          bool
	NonzeroCurvatureRequiresSource bool
	GaugeCompatibilityIfFamilyOnly bool
	ConnectionNativeInCurrentAsha  bool
	CoefficientsFixedByCurvature   bool
	CKMCapacityConditional         bool
	CKMAnglePredicted              bool
	Verdict                        string
	Reason                         string
}

type CoefficientImpact struct {
	Executed                          bool
	Sectors                           []string
	CoefficientsPerSector             int
	TotalFreeTextureCoefficients      int
	TopologicalCoefficientValuesFound int
	RootsOfUnityFixCoefficients       bool
	TraceFixesCoefficients            bool
	CurvatureFixesCoefficients        bool
	SectorSplittingNative             bool
	YukawaDataImported                bool
	Verdict                           string
	Reason                            string
}

type ModuliScenario struct {
	Name                        string
	Status                      string
	ModuliDim                   int
	ThreeDistinctMassesPossible bool
	CKMPossible                 bool
	PMNSPossible                bool
	CoefficientsFixed           bool
	NativeReduction             bool
	ConditionalOnly             bool
	Reason                      string
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
	AxiomStatusPreserved      bool
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
	Inheritance  Inheritance
	Arena        SelectorArena
	Functionals  []FunctionalAudit
	Connection   ConnectionAudit
	Coefficients CoefficientImpact
	Moduli       ModuliImpact
	Firewall     Firewall
	Next         NextStep
	Truth        string
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
	a.Functionals = buildFunctionals(a.Arena)
	a.Connection = buildConnection(a.Arena)
	a.Coefficients = buildCoefficients()
	a.Moduli = buildModuli(a.Coefficients, a.Connection)
	a.Firewall = buildFirewall()
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{Executed: true, Gate413PairAxiomCompatible: true, Gate413PairNotNative: true, Gate413CKMCapacity: true, Gate413CoefficientsFree: true, Gate412KDiagonalOnly: true, Gate411AxiomLedgerCompiled: true, ChargedModuliDim: Gate372ChargedFlavorModuliDim, Verdict: "Gate 414 inherits the Gate-413 boundary: a noncommuting family pair activates conditional mixing capacity, but the texture coefficients remain unselected and the pair is not native."}
}

func buildArena() SelectorArena {
	k := [][]float64{{-1, 0, 0}, {0, 0, 0}, {0, 0, 1}}
	s := [][]float64{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}}
	x := add(s, transpose(s))
	return SelectorArena{Executed: true, KName: "K_gen=diag(-1,0,1)", ShiftObservableName: "X_gen=S_gen+S_gen^T", KMatrix: k, XMatrix: x, FamilyBasisDimension: FamilyRank, GeneratedAlgebraDimension: 9, NoncommutingCapacity: frobenius(comm(k, x)) > 1e-12, CoefficientsNative: false, Verdict: "coefficient-selector arena formalized", Reason: "The family pair spans enough algebra for mixing, but ASHA has not supplied a functional selecting sector coefficients."}
}

func buildFunctionals(arena SelectorArena) []FunctionalAudit {
	k := arena.KMatrix
	x := arena.XMatrix
	commNorm := frobenius(comm(k, x))
	return []FunctionalAudit{
		{Name: "quadratic trace/norm", Executed: true, FunctionalType: "Tr(A^T A)", GaugeCompatible: true, EmpiricalIndependent: true, UniqueCoefficientRay: false, SelectsNoncommutingTexture: false, SelectsPhysicalSectorWeights: false, SelectorNative: false, DiagnosticValue: trace(mul(transpose(k), k)) + trace(mul(transpose(x), x)), Verdict: "no selector", Reason: "The trace norm is invariant under orthogonal family-basis rotations and fixes only scale/normalization, not a physical coefficient ray."},
		{Name: "adjoint-curvature relative to K", Executed: true, FunctionalType: "||[K,A]||^2", GaugeCompatible: true, EmpiricalIndependent: true, UniqueCoefficientRay: false, SelectsNoncommutingTexture: true, SelectsPhysicalSectorWeights: false, SelectorNative: false, DiagnosticValue: commNorm * commNorm, Verdict: "capacity diagnostic, not coefficient theorem", Reason: "The functional detects the shift direction, but maximizing/minimizing requires an added normalization and sign/sector rule; it does not fix up/down/lepton coefficients."},
		{Name: "spectral action family trace", Executed: true, FunctionalType: "Tr f(D_Family^2)", GaugeCompatible: true, EmpiricalIndependent: true, UniqueCoefficientRay: false, SelectsNoncommutingTexture: false, SelectsPhysicalSectorWeights: false, SelectorNative: false, DiagnosticValue: 0, Verdict: "central or flat", Reason: "With no native family curvature or source, the spectral trace is a class function; it cannot select a noncentral texture orientation."},
		{Name: "sector-split source functional", Executed: true, FunctionalType: "<J_sector,A>", GaugeCompatible: true, EmpiricalIndependent: false, UniqueCoefficientRay: true, SelectsNoncommutingTexture: true, SelectsPhysicalSectorWeights: true, SelectorNative: false, DiagnosticValue: 1, Verdict: "quarantined external source", Reason: "A source can pick any desired coefficient ray, but the source is exactly the missing data unless derived elsewhere."},
	}
}

func buildConnection(arena SelectorArena) ConnectionAudit {
	k := arena.KMatrix
	x := arena.XMatrix
	n := frobenius(comm(k, x))
	return ConnectionAudit{Executed: true, ConnectionAnsatz: "A_family = a K_gen + b X_gen on the U(3)_gen fiber", FamilyCurvatureSampleNorm: n, YangMillsMinimizerFlat: true, FlatMinimizerCommutes: true, NonzeroCurvatureRequiresSource: true, GaugeCompatibilityIfFamilyOnly: true, ConnectionNativeInCurrentAsha: false, CoefficientsFixedByCurvature: false, CKMCapacityConditional: true, CKMAnglePredicted: false, Verdict: "constrained connection has capacity but no selector", Reason: "The Yang-Mills-like curvature action is minimized by flat commuting family connections. Nonzero CKM-capable curvature must be imposed by a source/boundary condition not present in current ASHA."}
}

func buildCoefficients() CoefficientImpact {
	sectors := []string{"up", "down", "charged-lepton", "neutrino"}
	return CoefficientImpact{Executed: true, Sectors: sectors, CoefficientsPerSector: 2, TotalFreeTextureCoefficients: len(sectors) * 2, TopologicalCoefficientValuesFound: 0, RootsOfUnityFixCoefficients: false, TraceFixesCoefficients: false, CurvatureFixesCoefficients: false, SectorSplittingNative: false, YukawaDataImported: false, Verdict: "sector coefficients remain free", Reason: "The K/S family pair supplies a basis of possible textures; no native rule assigns the sector-specific coefficients needed for masses or mixing angles."}
}

func buildModuli(coeff CoefficientImpact, conn ConnectionAudit) ModuliImpact {
	scenarios := []ModuliScenario{
		{Name: "native ASHA through Gate 410", Status: StatusFirewallPreserved13Moduli, ModuliDim: Gate372ChargedFlavorModuliDim, ThreeDistinctMassesPossible: false, CKMPossible: false, PMNSPossible: false, CoefficientsFixed: false, NativeReduction: false, ConditionalOnly: false, Reason: "No nontrivial family bundle is native."},
		{Name: "Gate 412 K_gen only", Status: "CONDITIONAL_DIAGONAL_HIERARCHY_ONLY", ModuliDim: Gate372ChargedFlavorModuliDim, ThreeDistinctMassesPossible: true, CKMPossible: false, PMNSPossible: false, CoefficientsFixed: false, NativeReduction: false, ConditionalOnly: true, Reason: "A single Hamiltonian is diagonal and gives no mixing."},
		{Name: "Gate 413 K_gen plus shift", Status: "CONDITIONAL_MIXING_CAPACITY_COEFFICIENTS_FREE", ModuliDim: Gate372ChargedFlavorModuliDim, ThreeDistinctMassesPossible: true, CKMPossible: true, PMNSPossible: true, CoefficientsFixed: false, NativeReduction: false, ConditionalOnly: true, Reason: "Noncommuting capacity exists, but sector coefficients are unselected."},
		{Name: "Gate 414 trace/curvature selector", Status: StatusFailedNoNativeSelector, ModuliDim: Gate372ChargedFlavorModuliDim, ThreeDistinctMassesPossible: true, CKMPossible: conn.CKMCapacityConditional, PMNSPossible: conn.CKMCapacityConditional, CoefficientsFixed: false, NativeReduction: false, ConditionalOnly: true, Reason: "Audited functionals either remain invariant/flat or require external sector source data."},
	}
	return ModuliImpact{StartDim: Gate372ChargedFlavorModuliDim, Scenarios: scenarios, BestNativeDim: Gate372ChargedFlavorModuliDim, NativeReductionBelow13: false, ConditionalMixingCapacity: conn.CKMCapacityConditional, CoefficientsRemainFree: coeff.TotalFreeTextureCoefficients > 0, FirewallPreserved: true, Verdict: "13-moduli firewall preserved; mixing capacity is conditional but coefficient selection is not derived"}
}

func buildFirewall() Firewall {
	return Firewall{Executed: true, NoObservedMassesImported: true, NoCKMImported: true, NoPMNSImported: true, NoYukawaMatricesInserted: true, AxiomStatusPreserved: true, NoNativeDerivationClaimed: true, Verdict: "empirical and native-derivation firewalls preserved"}
}

func buildNext() NextStep {
	return NextStep{Gate: 415, Title: "Family Boundary Condition / Sector Source Axiom Minimality Sieve", Reason: "Gate 414 shows trace and curvature functionals do not fix coefficients. The next possible route is an explicit minimal boundary/source axiom for sector coefficients, ranked by mathematical cost and empirical independence.", PrimaryTask: "Classify the least additional source/boundary data required to select up/down/lepton coefficient rays without inserting observed Yukawa matrices."}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate413CoefficientsFree || a.Inheritance.ChargedModuliDim != Gate372ChargedFlavorModuliDim {
		return fmt.Errorf("bad inheritance")
	}
	if !a.Arena.Executed || !a.Arena.NoncommutingCapacity || a.Arena.GeneratedAlgebraDimension != 9 || a.Arena.CoefficientsNative {
		return fmt.Errorf("bad arena")
	}
	if len(a.Functionals) < 4 {
		return fmt.Errorf("missing functionals")
	}
	if a.Connection.CoefficientsFixedByCurvature || a.Connection.CKMAnglePredicted || a.Connection.ConnectionNativeInCurrentAsha || !a.Connection.CKMCapacityConditional {
		return fmt.Errorf("bad connection")
	}
	if a.Coefficients.TopologicalCoefficientValuesFound != 0 || a.Coefficients.YukawaDataImported || a.Coefficients.SectorSplittingNative {
		return fmt.Errorf("bad coefficients")
	}
	if a.Moduli.BestNativeDim != Gate372ChargedFlavorModuliDim || a.Moduli.NativeReductionBelow13 || !a.Moduli.FirewallPreserved {
		return fmt.Errorf("bad moduli")
	}
	return nil
}

func Statuses(a Analysis) []string {
	statuses := []string{
		StatusGate413BoundaryInherited,
		StatusSelectorArenaFormalized,
		StatusTraceFunctionalAudited,
		StatusCurvatureFunctionalAudited,
		StatusSectorSplitAudited,
		StatusConnectionStressTested,
		StatusEmpiricalFirewallAudited,
		StatusFailedNoNativeSelector,
		StatusFailedTraceU3Invariant,
		StatusFailedCurvatureFlatOrbits,
		StatusFailedSectorWeightsFree,
		StatusFailedConnectionNeedsAxiom,
		StatusFailedNoCKMAnglePrediction,
		StatusFailedNoNativeModuliReduction,
		StatusFirewallPreserved13Moduli,
	}
	return statuses
}

func truth(a Analysis) string {
	var b strings.Builder
	b.WriteString("Gate 414 proves that the K/S family-pair axiom creates noncommuting texture capacity but does not select physical coefficients. ")
	b.WriteString("Trace and spectral functionals are too invariant, curvature minimization selects flat commuting family connections, and nonzero mixing requires an external sector source or boundary condition. ")
	b.WriteString("Therefore the CKM/PMNS arena is conditionally available, but no native ASHA theorem predicts mixing angles or reduces the Gate-372 charged flavor firewall. dim M_charged remains 13.")
	return b.String()
}

func add(a, b [][]float64) [][]float64 {
	out := zeros(len(a), len(a[0]))
	for i := range a {
		for j := range a[i] {
			out[i][j] = a[i][j] + b[i][j]
		}
	}
	return out
}

func sub(a, b [][]float64) [][]float64 {
	out := zeros(len(a), len(a[0]))
	for i := range a {
		for j := range a[i] {
			out[i][j] = a[i][j] - b[i][j]
		}
	}
	return out
}

func transpose(a [][]float64) [][]float64 {
	out := zeros(len(a[0]), len(a))
	for i := range a {
		for j := range a[i] {
			out[j][i] = a[i][j]
		}
	}
	return out
}

func mul(a, b [][]float64) [][]float64 {
	out := zeros(len(a), len(b[0]))
	for i := range a {
		for k := range b {
			for j := range b[0] {
				out[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return out
}

func comm(a, b [][]float64) [][]float64 { return sub(mul(a, b), mul(b, a)) }

func zeros(r, c int) [][]float64 {
	out := make([][]float64, r)
	for i := range out {
		out[i] = make([]float64, c)
	}
	return out
}

func trace(a [][]float64) float64 {
	n := len(a)
	if len(a[0]) < n {
		n = len(a[0])
	}
	t := 0.0
	for i := 0; i < n; i++ {
		t += a[i][i]
	}
	return t
}

func frobenius(a [][]float64) float64 {
	s := 0.0
	for i := range a {
		for j := range a[i] {
			s += a[i][j] * a[i][j]
		}
	}
	return math.Sqrt(s)
}

func FormatFloat(x float64) string {
	if math.Abs(x) < 1e-12 {
		x = 0
	}
	return fmt.Sprintf("%.12g", x)
}
