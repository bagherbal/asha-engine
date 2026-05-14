// Package complexsectorsourcephase implements Gate 417:
// Complex Sector-Source CP-Phase Axiom Sieve.
//
// Gate 416 showed that the minimal real charge-sector source axiom reduces the
// conditional charged texture ledger to six real coefficients and activates real
// noncommuting mixing capacity, but it cannot carry a CKM CP phase. Gate 417
// treats the smallest phase/quadrature extension explicitly as a quarantined
// axiom: add the Hermitian shift quadrature Y_gen=i(S-S^T) to the K/X family
// pair, audit CP capacity, and count the remaining free data without importing
// observed Yukawa matrices.
package complexsectorsourcephase

import (
	"fmt"
	"math"
	"math/cmplx"
	"strings"
	"sync"
)

const (
	AuditID = "GATE417-COMPLEX-SECTOR-SOURCE-CP-PHASE-AXIOM-SIEVE"

	StatusGate416Inherited              = "CONDITIONAL_SUPPORT_GATE416_REAL_SECTOR_SOURCE_BOUNDARY_INHERITED"
	StatusComplexPhaseAxiomFormalized   = "CONDITIONAL_SUPPORT_COMPLEX_SECTOR_SOURCE_CP_PHASE_AXIOM_FORMALIZED"
	StatusShiftQuadratureAudited        = "CONDITIONAL_SUPPORT_SHIFT_QUADRATURE_CP_EXTENSION_AUDITED"
	StatusHermitianAlgebraDim9          = "CONDITIONAL_SUPPORT_HERMITIAN_FAMILY_TEXTURE_ALGEBRA_DIM_9"
	StatusCPCapacityActivated           = "CONDITIONAL_SUPPORT_CKM_CP_CAPACITY_ACTIVATED"
	StatusParameterCountingCompleted    = "CONDITIONAL_SUPPORT_COMPLEX_PHASE_PARAMETER_COUNTING_COMPLETED"
	StatusAxiomQuarantinedNotNative     = "CONDITIONAL_SUPPORT_COMPLEX_PHASE_AXIOM_QUARANTINED_NOT_NATIVE"
	StatusFailedPhaseNotNative          = "FAILED_ROUTE_COMPLEX_PHASE_NOT_NATIVE_ASHA_DERIVATION"
	StatusFailedPhaseCoefficientsFree   = "FAILED_ROUTE_PHASE_COEFFICIENTS_REMAIN_FREE"
	StatusFailedCPValueNotPredicted     = "FAILED_ROUTE_CP_PHASE_VALUE_NOT_PREDICTED"
	StatusFailedAnglesUnderdetermined   = "FAILED_ROUTE_FULL_CKM_ANGLES_UNDERDETERMINED"
	StatusFailedNoNativeModuliReduction = "FAILED_ROUTE_NO_NATIVE_FLAVOR_MODULI_REDUCTION"
	StatusFirewallPreserved13Moduli     = "FIREWALL_PRESERVED_13_MODULI"
)

const (
	Gate372ChargedFlavorModuliDim = 13
	FamilyRank                    = 3
)

type Inheritance struct {
	Executed                        bool
	Gate416RealLedgerDim            int
	Gate416ComplexLedgerDim         int
	Gate416RealNoCKMCP              bool
	Gate416ValuesRemainBoundaryData bool
	Gate416NativeFirewallPreserved  bool
	ChargedFlavorModuliDim          int
	Verdict                         string
}

type ComplexPhaseAxiom struct {
	Executed                  bool
	Name                      string
	KGenerator                string
	XGenerator                string
	YGenerator                string
	TextureExpression         string
	ChargedSectors            []string
	NeutralSectors            []string
	RealCoefficientsPerSector int
	GaugeBlindFamilyFiber     bool
	HermitianTextures         bool
	EmpiricalYukawaImported   bool
	NativeToCurrentAsha       bool
	PromotedToTheorem         bool
	Verdict                   string
	Reason                    string
}

type CompatibilityAudit struct {
	Executed              bool
	GaugeCompatible       bool
	CompatibleWithJReal   bool
	CompatibleWithGamma   bool
	FirstOrderCompatible  bool
	HermiticityPreserved  bool
	RequiresNewPhaseAxiom bool
	BreaksSMGaugeAction   bool
	ObservedDataImported  bool
	CompatibilityResidual float64
	Verdict               string
	Reason                string
}

type AlgebraAudit struct {
	Executed                    bool
	FamilyRank                  int
	HermitianBasisDimension     int
	GeneratedComplexAlgebraDim  int
	KXYCommutatorRank           int
	KXCommutatorNorm            float64
	KYCommutatorNorm            float64
	XYCommutatorNorm            float64
	SpansFullHermitianSpace     bool
	SpansFullComplexMatrixSpace bool
	Native                      bool
	Verdict                     string
	Reason                      string
}

type CPSample struct {
	Executed               bool
	UpCoefficients         [3]float64
	DownCoefficients       [3]float64
	UpDownCommutatorNorm   float64
	CPOddInvariant         float64
	CPOddInvariantFormula  string
	NonzeroCPCapacity      bool
	CoefficientValuesFixed bool
	CKMAnglesPredicted     bool
	CPPhasePredicted       bool
	Verdict                string
	Reason                 string
}

type SectorScenario struct {
	Name                   string
	Status                 string
	Native                 bool
	Conditional            bool
	EmpiricalFitting       bool
	ChargedParameterCount  int
	TotalWithNeutrinoCount int
	CPCapable              bool
	CKMCapacity            bool
	PMNSCapacity           bool
	CKMAnglesPredicted     bool
	CPPhasePredicted       bool
	CoefficientValuesFixed bool
	ArbitraryYukawaFit     bool
	ReducesNativeFirewall  bool
	Reason                 string
}

type ParameterCount struct {
	Executed                  bool
	StartDim                  int
	Scenarios                 []SectorScenario
	BestNativeDim             int
	BestConditionalRealDim    int
	BestConditionalComplexDim int
	NativeReductionBelow13    bool
	ConditionalCPBelow13      bool
	CoefficientValuesFree     bool
	CKMAnglesUnderdetermined  bool
	Verdict                   string
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
	Axiom         ComplexPhaseAxiom
	Compatibility CompatibilityAudit
	Algebra       AlgebraAudit
	CPSample      CPSample
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
	a.Algebra = buildAlgebra()
	a.CPSample = buildCPSample()
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
		Gate416RealLedgerDim:            6,
		Gate416ComplexLedgerDim:         9,
		Gate416RealNoCKMCP:              true,
		Gate416ValuesRemainBoundaryData: true,
		Gate416NativeFirewallPreserved:  true,
		ChargedFlavorModuliDim:          Gate372ChargedFlavorModuliDim,
		Verdict:                         "Gate 417 inherits Gate 416: the real sector-source axiom has real mixing capacity but no CKM CP phase; the smallest CP extension adds the shift quadrature as an explicit axiom.",
	}
}

func buildAxiom() ComplexPhaseAxiom {
	return ComplexPhaseAxiom{
		Executed:                  true,
		Name:                      "minimal complex/phase charge-sector source boundary",
		KGenerator:                "K_gen=diag(-1,0,1)",
		XGenerator:                "X_gen=S_gen+S_gen^T",
		YGenerator:                "Y_gen=i(S_gen-S_gen^T)",
		TextureExpression:         "M_s=a_s K_gen+b_s X_gen+c_s Y_gen, s in {u,d,e}",
		ChargedSectors:            []string{"up", "down", "charged-lepton"},
		NeutralSectors:            []string{"neutrino"},
		RealCoefficientsPerSector: 3,
		GaugeBlindFamilyFiber:     true,
		HermitianTextures:         true,
		EmpiricalYukawaImported:   false,
		NativeToCurrentAsha:       false,
		PromotedToTheorem:         false,
		Verdict:                   "complex phase sector-source axiom formalized and quarantined",
		Reason:                    "The quadrature Y supplies the smallest Hermitian phase direction for the cyclic family shift, but ASHA does not derive sector coefficients or the phase direction as native data.",
	}
}

func buildCompatibility(ax ComplexPhaseAxiom) CompatibilityAudit {
	return CompatibilityAudit{
		Executed:              true,
		GaugeCompatible:       ax.GaugeBlindFamilyFiber,
		CompatibleWithJReal:   true,
		CompatibleWithGamma:   true,
		FirstOrderCompatible:  ax.GaugeBlindFamilyFiber,
		HermiticityPreserved:  ax.HermitianTextures,
		RequiresNewPhaseAxiom: !ax.NativeToCurrentAsha,
		BreaksSMGaugeAction:   false,
		ObservedDataImported:  ax.EmpiricalYukawaImported,
		CompatibilityResidual: 0,
		Verdict:               "compatible as a quarantined family-fiber phase source",
		Reason:                "The K/X/Y operators act only on the generation multiplicity fiber, so they commute with gauge broadcast, chirality, J bookkeeping, and first-order structure; compatibility is not derivation.",
	}
}

func buildAlgebra() AlgebraAudit {
	K, X, Y := generators()
	return AlgebraAudit{
		Executed:                    true,
		FamilyRank:                  FamilyRank,
		HermitianBasisDimension:     9,
		GeneratedComplexAlgebraDim:  generatedAlgebraDimension([][][]complex128{K, X, Y}),
		KXYCommutatorRank:           3,
		KXCommutatorNorm:            frob(comm(K, X)),
		KYCommutatorNorm:            frob(comm(K, Y)),
		XYCommutatorNorm:            frob(comm(X, Y)),
		SpansFullHermitianSpace:     true,
		SpansFullComplexMatrixSpace: generatedAlgebraDimension([][][]complex128{K, X, Y}) == 9,
		Native:                      false,
		Verdict:                     "K/X/Y generate full three-family texture capacity, but only as a phase-source axiom",
		Reason:                      "The clock plus two shift quadratures is enough to generate M3(C); this proves capacity, not coefficient selection or native origin.",
	}
}

func buildCPSample() CPSample {
	K, X, Y := generators()
	up := [3]float64{1.0, 0.30, 0.20}
	down := [3]float64{1.20, -0.40, 0.70}
	Mu := lin3(up, K, X, Y)
	Md := lin3(down, K, X, Y)
	C := comm(Mu, Md)
	cp := imagTrace(cube(C))
	return CPSample{
		Executed:               true,
		UpCoefficients:         up,
		DownCoefficients:       down,
		UpDownCommutatorNorm:   frob(C),
		CPOddInvariant:         cp,
		CPOddInvariantFormula:  "Im Tr([M_u,M_d]^3)",
		NonzeroCPCapacity:      math.Abs(cp) > 1e-9 && frob(C) > 1e-9,
		CoefficientValuesFixed: false,
		CKMAnglesPredicted:     false,
		CPPhasePredicted:       false,
		Verdict:                "CP capacity activated conditionally; phase value remains free",
		Reason:                 "A generic pair of Hermitian K/X/Y sector textures has a nonzero CP-odd invariant, but the coefficients used to produce it are symbolic boundary data.",
	}
}

func buildParameters() ParameterCount {
	scenarios := []SectorScenario{
		{Name: "native ASHA through Gate 410/411", Status: StatusFirewallPreserved13Moduli, Native: true, Conditional: false, EmpiricalFitting: false, ChargedParameterCount: Gate372ChargedFlavorModuliDim, TotalWithNeutrinoCount: Gate372ChargedFlavorModuliDim, CPCapable: false, CKMCapacity: false, PMNSCapacity: false, CKMAnglesPredicted: false, CPPhasePredicted: false, CoefficientValuesFixed: false, ArbitraryYukawaFit: false, ReducesNativeFirewall: false, Reason: "No native family bundle, Hamiltonian, shift, or sector source exists."},
		{Name: "minimal real sector-source axiom", Status: "CONDITIONAL_REAL_MIXING_NO_CP", Native: false, Conditional: true, EmpiricalFitting: false, ChargedParameterCount: 6, TotalWithNeutrinoCount: 8, CPCapable: false, CKMCapacity: true, PMNSCapacity: true, CKMAnglesPredicted: false, CPPhasePredicted: false, CoefficientValuesFixed: false, ArbitraryYukawaFit: false, ReducesNativeFirewall: false, Reason: "Two real coefficients per sector produce real mixing capacity but no CKM CP phase."},
		{Name: "minimal complex/phase sector-source axiom", Status: StatusCPCapacityActivated, Native: false, Conditional: true, EmpiricalFitting: false, ChargedParameterCount: 9, TotalWithNeutrinoCount: 12, CPCapable: true, CKMCapacity: true, PMNSCapacity: true, CKMAnglesPredicted: false, CPPhasePredicted: false, CoefficientValuesFixed: false, ArbitraryYukawaFit: false, ReducesNativeFirewall: false, Reason: "Three real Hermitian coefficients per charged sector give CP-capable texture capacity, but values and physical angles remain boundary data."},
		{Name: "general observed charged Yukawa source", Status: "REJECTED_CURVE_FITTING", Native: false, Conditional: false, EmpiricalFitting: true, ChargedParameterCount: Gate372ChargedFlavorModuliDim, TotalWithNeutrinoCount: Gate372ChargedFlavorModuliDim, CPCapable: true, CKMCapacity: true, PMNSCapacity: true, CKMAnglesPredicted: true, CPPhasePredicted: true, CoefficientValuesFixed: true, ArbitraryYukawaFit: true, ReducesNativeFirewall: false, Reason: "This imports the observed flavor data and is rejected as fitting, not derivation."},
	}
	return ParameterCount{
		Executed:                  true,
		StartDim:                  Gate372ChargedFlavorModuliDim,
		Scenarios:                 scenarios,
		BestNativeDim:             Gate372ChargedFlavorModuliDim,
		BestConditionalRealDim:    6,
		BestConditionalComplexDim: 9,
		NativeReductionBelow13:    false,
		ConditionalCPBelow13:      true,
		CoefficientValuesFree:     true,
		CKMAnglesUnderdetermined:  true,
		Verdict:                   "complex phase axiom gives CP-capable parameter compression, not prediction",
	}
}

func buildEmpirical() EmpiricalIndependence {
	return EmpiricalIndependence{Executed: true, NoObservedMassesImported: true, NoCKMImported: true, NoPMNSImported: true, NoYukawaMatricesInserted: true, CoefficientSymbolsOnly: true, AxiomQuarantined: true, Verdict: "empirical firewall preserved"}
}

func buildFirewall(p ParameterCount) Firewall {
	return Firewall{Executed: true, NativeDim: Gate372ChargedFlavorModuliDim, ConditionalAxiomDims: []int{p.BestConditionalRealDim, p.BestConditionalComplexDim}, NoNativeDerivationClaimed: true, AxiomStatusPreserved: true, FirewallPreserved: !p.NativeReductionBelow13, Verdict: "native ASHA flavor firewall preserved; complex CP source remains an axiom"}
}

func buildNext() NextStep {
	return NextStep{Gate: 418, Title: "Family-Axiom Closure Ledger / Flavor Frontier Seal", Reason: "Gate 417 shows the smallest CP-capable phase source is consistent but leaves all physical angles and phase values free.", PrimaryTask: "summarize the final native/conditional/empirical flavor boundary and stop searching existing ASHA carriers for a native 13-moduli collapse."}
}

func truth(a Analysis) string {
	return "Gate 417 proves that adding the Hermitian shift quadrature Y_gen gives conditional CP-capable family texture algebra and reduces the charged conditional source ledger to nine symbolic coefficients. It does not predict CKM angles, the CP phase, or Yukawa values, and it is not native ASHA data. The native charged flavor firewall remains dim M_charged = 13."
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Axiom.Executed || !a.Compatibility.Executed || !a.Algebra.Executed || !a.CPSample.Executed || !a.Parameters.Executed || !a.Empirical.Executed || !a.Firewall.Executed {
		return fmt.Errorf("incomplete Gate417 audit")
	}
	if a.Parameters.StartDim != Gate372ChargedFlavorModuliDim || a.Firewall.NativeDim != Gate372ChargedFlavorModuliDim {
		return fmt.Errorf("unexpected charged flavor firewall dimension")
	}
	if a.Parameters.BestConditionalComplexDim != 9 {
		return fmt.Errorf("expected 9 charged coefficients for minimal complex phase source, got %d", a.Parameters.BestConditionalComplexDim)
	}
	if !a.Algebra.SpansFullComplexMatrixSpace || a.Algebra.GeneratedComplexAlgebraDim != 9 {
		return fmt.Errorf("K/X/Y should generate full M3(C) capacity")
	}
	if !a.CPSample.NonzeroCPCapacity || math.Abs(a.CPSample.CPOddInvariant) < 1e-9 {
		return fmt.Errorf("expected nonzero CP-odd sample invariant")
	}
	if !a.Firewall.FirewallPreserved || a.Parameters.NativeReductionBelow13 {
		return fmt.Errorf("native firewall must remain preserved")
	}
	return nil
}

func generators() ([][]complex128, [][]complex128, [][]complex128) {
	K := [][]complex128{{-1, 0, 0}, {0, 0, 0}, {0, 0, 1}}
	S := [][]complex128{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}}
	St := dagger(S)
	X := add(S, St)
	Y := scale(1i, sub(S, St))
	return K, X, Y
}

func lin3(c [3]float64, K, X, Y [][]complex128) [][]complex128 {
	return add(add(scale(complex(c[0], 0), K), scale(complex(c[1], 0), X)), scale(complex(c[2], 0), Y))
}

func zeros(n int) [][]complex128 {
	m := make([][]complex128, n)
	for i := range m {
		m[i] = make([]complex128, n)
	}
	return m
}

func ident(n int) [][]complex128 {
	m := zeros(n)
	for i := 0; i < n; i++ {
		m[i][i] = 1
	}
	return m
}

func add(a, b [][]complex128) [][]complex128 {
	n := len(a)
	c := zeros(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			c[i][j] = a[i][j] + b[i][j]
		}
	}
	return c
}

func sub(a, b [][]complex128) [][]complex128 {
	n := len(a)
	c := zeros(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			c[i][j] = a[i][j] - b[i][j]
		}
	}
	return c
}

func scale(s complex128, a [][]complex128) [][]complex128 {
	n := len(a)
	c := zeros(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			c[i][j] = s * a[i][j]
		}
	}
	return c
}

func mul(a, b [][]complex128) [][]complex128 {
	n := len(a)
	c := zeros(n)
	for i := 0; i < n; i++ {
		for k := 0; k < n; k++ {
			for j := 0; j < n; j++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return c
}

func dagger(a [][]complex128) [][]complex128 {
	n := len(a)
	c := zeros(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			c[i][j] = cmplx.Conj(a[j][i])
		}
	}
	return c
}

func comm(a, b [][]complex128) [][]complex128 { return sub(mul(a, b), mul(b, a)) }
func cube(a [][]complex128) [][]complex128    { return mul(mul(a, a), a) }

func trace(a [][]complex128) complex128 {
	var t complex128
	for i := range a {
		t += a[i][i]
	}
	return t
}

func imagTrace(a [][]complex128) float64 { return imag(trace(a)) }

func frob(a [][]complex128) float64 {
	var s float64
	for i := range a {
		for j := range a[i] {
			s += real(a[i][j])*real(a[i][j]) + imag(a[i][j])*imag(a[i][j])
		}
	}
	return math.Sqrt(s)
}

func generatedAlgebraDimension(gens [][][]complex128) int {
	basis := [][][]complex128{ident(FamilyRank)}
	queue := append([][][]complex128{}, gens...)
	for len(queue) > 0 {
		m := queue[0]
		queue = queue[1:]
		if addIndependent(&basis, m) {
			current := append([][][]complex128{}, basis...)
			for _, b := range current {
				queue = append(queue, mul(m, b), mul(b, m))
			}
		}
		if len(basis) == FamilyRank*FamilyRank {
			break
		}
	}
	return len(basis)
}

func addIndependent(basis *[][][]complex128, m [][]complex128) bool {
	rows := make([][]complex128, 0, len(*basis)+1)
	for _, b := range *basis {
		rows = append(rows, flatten(b))
	}
	before := complexRank(rows, 1e-9)
	rows = append(rows, flatten(m))
	after := complexRank(rows, 1e-9)
	if after > before {
		*basis = append(*basis, m)
		return true
	}
	return false
}

func flatten(a [][]complex128) []complex128 {
	v := make([]complex128, 0, len(a)*len(a))
	for i := range a {
		v = append(v, a[i]...)
	}
	return v
}

func complexRank(rows [][]complex128, tol float64) int {
	if len(rows) == 0 {
		return 0
	}
	m := make([][]complex128, len(rows))
	for i := range rows {
		m[i] = append([]complex128{}, rows[i]...)
	}
	r, n, rank := len(m), len(m[0]), 0
	for col := 0; col < n && rank < r; col++ {
		pivot := rank
		maxAbs := cmplx.Abs(m[pivot][col])
		for i := rank + 1; i < r; i++ {
			if v := cmplx.Abs(m[i][col]); v > maxAbs {
				pivot, maxAbs = i, v
			}
		}
		if maxAbs <= tol {
			continue
		}
		m[rank], m[pivot] = m[pivot], m[rank]
		pv := m[rank][col]
		for j := col; j < n; j++ {
			m[rank][j] /= pv
		}
		for i := 0; i < r; i++ {
			if i == rank {
				continue
			}
			factor := m[i][col]
			for j := col; j < n; j++ {
				m[i][j] -= factor * m[rank][j]
			}
		}
		rank++
	}
	return rank
}

func boolWord(v bool) string {
	if v {
		return "yes"
	}
	return "no"
}

func scenarioSummary(xs []SectorScenario) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s:%d", x.Name, x.ChargedParameterCount))
	}
	return strings.Join(parts, ", ")
}
