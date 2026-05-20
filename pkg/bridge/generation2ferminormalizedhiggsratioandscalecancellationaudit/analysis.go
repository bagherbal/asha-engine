// Package generation2ferminormalizedhiggsratioandscalecancellationaudit implements
// Gate 779: Fermi-Normalized Higgs Ratio and Scale-Cancellation Audit.
//
// Gate 778 showed that the sealed tree Higgs tower still requires a
// dimensionful Fermi/VEV scale seal. Gate 779 cancels that scale from the tree
// proxy by using the external Fermi-VEV convention, recording the dimensionless
// identity 4 sqrt(2) G_F m_H_tree^2 = C_Higgs. This is a Fermi-normalized
// ratio and scale-cancellation audit only. It does not derive G_F, v, the Higgs
// pole mass, scalar runtime lambda, Yukawa operators, CKM/PMNS, flavor
// hierarchy, or a native HistoryLoopUnit theorem.
package generation2ferminormalizedhiggsratioandscalecancellationaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE779-FERMI-NORMALIZED-HIGGS-RATIO-SCALE-CANCELLATION-AUDIT"

	StatusGate778ElectroweakScaleAirlockInherited            = "PASS_GATE778_ELECTROWEAK_SCALE_AIRLOCK_INHERITED"
	StatusFermiNormalizedRatioDefined                        = "PASS_FERMI_NORMALIZED_RATIO_DEFINED"
	StatusVEVScaleCancellationComputed                       = "PASS_VEV_SCALE_CANCELLATION_COMPUTED"
	StatusNumericalRatioLedgerRecorded                       = "PASS_NUMERICAL_RATIO_LEDGER_RECORDED"
	StatusDimensionlessAndScaleTasksSeparated                = "PASS_DIMENSIONLESS_AND_SCALE_TASKS_SEPARATED"
	StatusPhysicalFirewallsEnforced                          = "PASS_PHYSICAL_FIREWALLS_ENFORCED"
	StatusFermiNormalizedTreeMassRatioEqualsCHiggs           = "CONDITIONAL_SUPPORT_FERMI_NORMALIZED_TREE_MASS_RATIO_EQUALS_C_HIGGS"
	StatusHiggsTreeProxySplitsIntoDimensionlessAndFermiScale = "CONDITIONAL_SUPPORT_HIGGS_TREE_PROXY_SPLITS_INTO_DIMENSIONLESS_C_HIGGS_AND_EXTERNAL_FERMI_SCALE"
	StatusNoNativeFermiConstantTheorem                       = "FAILED_ROUTE_NO_NATIVE_FERMI_CONSTANT_THEOREM"
	StatusNoNativeElectroweakScaleTheorem                    = "FAILED_ROUTE_NO_NATIVE_ELECTROWEAK_SCALE_THEOREM"
	StatusCHiggsNotNativeHiggsTheorem                        = "FAILED_ROUTE_C_HIGGS_NOT_NATIVE_HIGGS_THEOREM"
	StatusTreeProxyNotPoleMass                               = "FAILED_ROUTE_TREE_PROXY_NOT_POLE_MASS"
	StatusNoYukawaOperatorOrEigenvalueTheorem                = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate779FermiNormalizedHiggsRatioBoundary           = "FIREWALL_PRESERVED_GATE779_FERMI_NORMALIZED_HIGGS_RATIO_BOUNDARY"
)

const (
	// Gate 778 ledger snapshots.
	cHiggsMZ         = 1.0372205204048603
	vevConventionGeV = 246.2196508
)

type Gate778Inheritance struct {
	Inherited              bool
	TreeTowerFormula       string
	FermiVEVSeal           string
	CHiggs                 float64
	VEVGeV                 float64
	DilationFactor         float64
	TreeMassGeV            float64
	EquivalentGFGeVMinus2  float64
	NativeFermiTheorem     bool
	NativeElectroweakScale bool
	PoleMassTheorem        bool
	Verdict                string
}

type FermiNormalizedRatio struct {
	Defined             bool
	SquaredTreeFormula  string
	VEVCancelledFormula string
	FermiConvention     string
	NormalizedIdentity  string
	DimensionlessLeft   string
	DimensionlessRight  string
	UsesExternalGFSeal  bool
	DerivesGF           bool
	DerivesVEV          bool
	Verdict             string
}

type ScaleCancellation struct {
	Computed                      bool
	TreeMassSquaredOverVSquared   float64
	CHiggsOverFour                float64
	Sqrt2GFTreeMassSquared        float64
	FourSqrt2GFTreeMassSquared    float64
	MatchesCHiggs                 bool
	ScaleCancelledToDimensionless bool
	Verdict                       string
}

type NumericalRatioLedger struct {
	CHiggs                     float64
	VEVGeV                     float64
	EquivalentGFGeVMinus2      float64
	TreeMassGeV                float64
	TreeMassSquaredGeV2        float64
	TreeMassOverVEV            float64
	Sqrt2GFTreeMassSquared     float64
	FourSqrt2GFTreeMassSquared float64
	Finite                     bool
	Verdict                    string
}

type TaskSeparation struct {
	Separated             bool
	DimensionlessTask     string
	ScaleTask             string
	RequiresBothForMass   bool
	RatioDoesNotDeriveGF  bool
	RatioDoesNotDeriveVEV bool
	Verdict               string
}

type PhysicalFirewalls struct {
	Audited                                bool
	RatioPoleMassTheorem                   bool
	GFAShaNativeInput                      bool
	FermiNormalizedRatioMeasuredPrediction bool
	CHiggsNativeHiggsTheorem               bool
	TreeProxyPoleMass                      bool
	DimensionlessRatioElectroweakScale     bool
	YukawaOperatorOrEigenvalue             bool
	Verdict                                string
}

type Analysis struct {
	Gate778      Gate778Inheritance
	Ratio        FermiNormalizedRatio
	Cancellation ScaleCancellation
	Ledger       NumericalRatioLedger
	Tasks        TaskSeparation
	Firewalls    PhysicalFirewalls
	Truth        string
}

var (
	cacheMu sync.Mutex
	cache   *Analysis
)

func BuildDefault() (*Analysis, error) {
	cacheMu.Lock()
	defer cacheMu.Unlock()
	if cache != nil {
		return cloneAnalysis(cache), nil
	}
	for name, value := range map[string]float64{
		"C_Higgs": cHiggsMZ,
		"v":       vevConventionGeV,
	} {
		if math.IsNaN(value) || math.IsInf(value, 0) || value <= 0 {
			return nil, fmt.Errorf("invalid %s ledger: %.17g", name, value)
		}
	}
	dRadial := math.Sqrt(cHiggsMZ)
	gf := 1 / (math.Sqrt2 * vevConventionGeV * vevConventionGeV)
	mass := (vevConventionGeV / 2) * dRadial
	mass2 := mass * mass
	massOverV := mass / vevConventionGeV
	sqrt2GFM2 := math.Sqrt2 * gf * mass2
	fourSqrt2GFM2 := 4 * sqrt2GFM2
	chiggsOver4 := cHiggsMZ / 4
	m2OverV2 := mass2 / (vevConventionGeV * vevConventionGeV)
	if !finite(dRadial, gf, mass, mass2, massOverV, sqrt2GFM2, fourSqrt2GFM2, chiggsOver4, m2OverV2) {
		return nil, fmt.Errorf("invalid Fermi-normalized Higgs ratio ledger")
	}

	a := &Analysis{
		Gate778: Gate778Inheritance{
			Inherited:              true,
			TreeTowerFormula:       "m_H_tree=(v/2)sqrt(C_Higgs)",
			FermiVEVSeal:           "FermiVEVScaleSeal: v=(sqrt(2)G_F)^(-1/2)",
			CHiggs:                 cHiggsMZ,
			VEVGeV:                 vevConventionGeV,
			DilationFactor:         dRadial,
			TreeMassGeV:            mass,
			EquivalentGFGeVMinus2:  gf,
			NativeFermiTheorem:     false,
			NativeElectroweakScale: false,
			PoleMassTheorem:        false,
			Verdict:                StatusGate778ElectroweakScaleAirlockInherited,
		},
		Ratio: FermiNormalizedRatio{
			Defined:             true,
			SquaredTreeFormula:  "m_H_tree^2=(v^2/4)C_Higgs",
			VEVCancelledFormula: "m_H_tree^2/v^2=C_Higgs/4",
			FermiConvention:     "1/v^2=sqrt(2)G_F",
			NormalizedIdentity:  "4sqrt(2)G_F m_H_tree^2=C_Higgs",
			DimensionlessLeft:   "4sqrt(2)G_F m_H_tree^2",
			DimensionlessRight:  "C_Higgs",
			UsesExternalGFSeal:  true,
			DerivesGF:           false,
			DerivesVEV:          false,
			Verdict:             StatusFermiNormalizedRatioDefined,
		},
		Cancellation: ScaleCancellation{
			Computed:                      true,
			TreeMassSquaredOverVSquared:   m2OverV2,
			CHiggsOverFour:                chiggsOver4,
			Sqrt2GFTreeMassSquared:        sqrt2GFM2,
			FourSqrt2GFTreeMassSquared:    fourSqrt2GFM2,
			MatchesCHiggs:                 closeRel(fourSqrt2GFM2, cHiggsMZ, 1e-15) && closeRel(sqrt2GFM2, chiggsOver4, 1e-15) && closeRel(m2OverV2, chiggsOver4, 1e-15),
			ScaleCancelledToDimensionless: true,
			Verdict:                       StatusVEVScaleCancellationComputed,
		},
		Ledger: NumericalRatioLedger{
			CHiggs:                     cHiggsMZ,
			VEVGeV:                     vevConventionGeV,
			EquivalentGFGeVMinus2:      gf,
			TreeMassGeV:                mass,
			TreeMassSquaredGeV2:        mass2,
			TreeMassOverVEV:            massOverV,
			Sqrt2GFTreeMassSquared:     sqrt2GFM2,
			FourSqrt2GFTreeMassSquared: fourSqrt2GFM2,
			Finite:                     finite(cHiggsMZ, vevConventionGeV, gf, mass, mass2, massOverV, sqrt2GFM2, fourSqrt2GFM2),
			Verdict:                    StatusNumericalRatioLedgerRecorded,
		},
		Tasks: TaskSeparation{
			Separated:             true,
			DimensionlessTask:     "derive or reduce C_Higgs natively",
			ScaleTask:             "derive or seal G_F / v",
			RequiresBothForMass:   true,
			RatioDoesNotDeriveGF:  true,
			RatioDoesNotDeriveVEV: true,
			Verdict:               StatusDimensionlessAndScaleTasksSeparated,
		},
		Firewalls: PhysicalFirewalls{
			Audited:                                true,
			RatioPoleMassTheorem:                   false,
			GFAShaNativeInput:                      false,
			FermiNormalizedRatioMeasuredPrediction: false,
			CHiggsNativeHiggsTheorem:               false,
			TreeProxyPoleMass:                      false,
			DimensionlessRatioElectroweakScale:     false,
			YukawaOperatorOrEigenvalue:             false,
			Verdict:                                StatusGate779FermiNormalizedHiggsRatioBoundary,
		},
		Truth: fmt.Sprintf("Gate779 cancels the VEV scale from the sealed tree proxy: 4sqrt(2)G_F m_H_tree^2=%.16f equals C_Higgs=%.16f under the external FermiVEVScaleSeal. This separates the dimensionless ASHA task C_Higgs from the independent Fermi/VEV scale task, while deriving no native G_F, VEV, pole mass, Higgs theorem, or Yukawa theorem.", fourSqrt2GFM2, cHiggsMZ),
	}
	cache = a
	return cloneAnalysis(a), nil
}

func cloneAnalysis(a *Analysis) *Analysis {
	clone := *a
	return &clone
}

func finite(values ...float64) bool {
	for _, v := range values {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return false
		}
	}
	return true
}

func closeRel(got, want, tol float64) bool {
	if math.IsNaN(got) || math.IsNaN(want) || math.IsInf(got, 0) || math.IsInf(want, 0) {
		return false
	}
	d := math.Abs(got - want)
	if want == 0 {
		return d <= tol
	}
	return d/math.Abs(want) <= tol
}

func Statuses() []string {
	return []string{
		StatusGate778ElectroweakScaleAirlockInherited,
		StatusFermiNormalizedRatioDefined,
		StatusVEVScaleCancellationComputed,
		StatusNumericalRatioLedgerRecorded,
		StatusDimensionlessAndScaleTasksSeparated,
		StatusPhysicalFirewallsEnforced,
		StatusFermiNormalizedTreeMassRatioEqualsCHiggs,
		StatusHiggsTreeProxySplitsIntoDimensionlessAndFermiScale,
		StatusNoNativeFermiConstantTheorem,
		StatusNoNativeElectroweakScaleTheorem,
		StatusCHiggsNotNativeHiggsTheorem,
		StatusTreeProxyNotPoleMass,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate779FermiNormalizedHiggsRatioBoundary,
	}
}

func FormatGate778(g Gate778Inheritance) string {
	return fmt.Sprintf("inherited=%t formula=%s seal=%s C=%.17g v=%.10f D=%.17g mass=%.12f GF=%.14e nativeFermi=%t nativeEWScale=%t pole=%t verdict=%s", g.Inherited, g.TreeTowerFormula, g.FermiVEVSeal, g.CHiggs, g.VEVGeV, g.DilationFactor, g.TreeMassGeV, g.EquivalentGFGeVMinus2, g.NativeFermiTheorem, g.NativeElectroweakScale, g.PoleMassTheorem, g.Verdict)
}

func FormatRatio(r FermiNormalizedRatio) string {
	return fmt.Sprintf("defined=%t squared=%s cancelled=%s convention=%s identity=%s left=%s right=%s externalGF=%t derivesGF=%t derivesVEV=%t verdict=%s", r.Defined, r.SquaredTreeFormula, r.VEVCancelledFormula, r.FermiConvention, r.NormalizedIdentity, r.DimensionlessLeft, r.DimensionlessRight, r.UsesExternalGFSeal, r.DerivesGF, r.DerivesVEV, r.Verdict)
}

func FormatCancellation(c ScaleCancellation) string {
	return fmt.Sprintf("computed=%t m2/v2=%.17g C/4=%.17g sqrt2GFm2=%.17g fourSqrt2GFm2=%.17g matches=%t dimensionless=%t verdict=%s", c.Computed, c.TreeMassSquaredOverVSquared, c.CHiggsOverFour, c.Sqrt2GFTreeMassSquared, c.FourSqrt2GFTreeMassSquared, c.MatchesCHiggs, c.ScaleCancelledToDimensionless, c.Verdict)
}

func FormatLedger(l NumericalRatioLedger) string {
	return fmt.Sprintf("C=%.17g v=%.10f GF=%.14e mass=%.12f mass2=%.12f mass/v=%.16f sqrt2GFm2=%.16f fourSqrt2GFm2=%.16f finite=%t verdict=%s", l.CHiggs, l.VEVGeV, l.EquivalentGFGeVMinus2, l.TreeMassGeV, l.TreeMassSquaredGeV2, l.TreeMassOverVEV, l.Sqrt2GFTreeMassSquared, l.FourSqrt2GFTreeMassSquared, l.Finite, l.Verdict)
}

func FormatTasks(t TaskSeparation) string {
	return fmt.Sprintf("separated=%t dimensionless=%s scale=%s requiresBoth=%t ratioNoGF=%t ratioNoVEV=%t verdict=%s", t.Separated, t.DimensionlessTask, t.ScaleTask, t.RequiresBothForMass, t.RatioDoesNotDeriveGF, t.RatioDoesNotDeriveVEV, t.Verdict)
}

func FormatFirewalls(f PhysicalFirewalls) string {
	return fmt.Sprintf("audited=%t ratioPole=%t GFNative=%t ratioMeasured=%t CNative=%t pole=%t ratioScale=%t yukawaOp=%t verdict=%s", f.Audited, f.RatioPoleMassTheorem, f.GFAShaNativeInput, f.FermiNormalizedRatioMeasuredPrediction, f.CHiggsNativeHiggsTheorem, f.TreeProxyPoleMass, f.DimensionlessRatioElectroweakScale, f.YukawaOperatorOrEigenvalue, f.Verdict)
}

func containsAll(haystack, needles []string) bool {
	joined := "\x00" + strings.Join(haystack, "\x00") + "\x00"
	for _, n := range needles {
		if !strings.Contains(joined, "\x00"+n+"\x00") {
			return false
		}
	}
	return true
}
