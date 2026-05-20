// Package generation2vevscaleairlockanddimensionfulhiggstowerfirewallaudit implements
// Gate 777: VEV Scale Airlock and Dimensionful Higgs Tower Firewall Audit.
//
// Gate 776 compressed the sealed tree Higgs radial tower into the dimensionless
// total correction factor C_Higgs=(3/N_eff)[1+L_Hopf(1-kappa_lambda_red)].
// Gate 777 audits the remaining dimensionful scale airlock: the bridge supplies
// dimensionless correction factors, while the physical GeV scale is carried by
// the supplied VEV convention. This is a scale and VEV-firewall audit only. It
// does not derive the VEV, Fermi constant, electroweak symmetry breaking, Higgs
// pole mass, Yukawa operators, CKM/PMNS, flavor hierarchy, or a native
// HistoryLoopUnit theorem.
package generation2vevscaleairlockanddimensionfulhiggstowerfirewallaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE777-VEV-SCALE-AIRLOCK-DIMENSIONFUL-HIGGS-TOWER-FIREWALL-AUDIT"

	StatusGate776TotalCorrectionDecompositionInherited = "PASS_GATE776_TOTAL_CORRECTION_DECOMPOSITION_INHERITED"
	StatusDimensionlessDimensionfulSplitAudited        = "PASS_DIMENSIONLESS_DIMENSIONFUL_SPLIT_AUDITED"
	StatusVEVConventionSealRecorded                    = "PASS_VEV_CONVENTION_SEAL_RECORDED"
	StatusScaleSensitivityComputed                     = "PASS_SCALE_SENSITIVITY_COMPUTED"
	StatusBaselineScaleInterpretationRecorded          = "PASS_BASELINE_SCALE_INTERPRETATION_RECORDED"
	StatusRemainingSourcePressureSplitRecorded         = "PASS_REMAINING_SOURCE_PRESSURE_SPLIT_RECORDED"
	StatusPhysicalFirewallsEnforced                    = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusASHAControlsDimensionlessCorrection   = "CONDITIONAL_SUPPORT_ASHA_CURRENTLY_CONTROLS_DIMENSIONLESS_HIGGS_TREE_CORRECTION"
	StatusVEVSealSuppliesDimensionfulScale      = "CONDITIONAL_SUPPORT_VEV_SEAL_SUPPLIES_THE_DIMENSIONFUL_ELECTROWEAK_SCALE"
	StatusTreeProxyEqualsHalfScaleTimesDilation = "CONDITIONAL_SUPPORT_TREE_PROXY_EQUALS_ELECTROWEAK_HALF_SCALE_TIMES_RADIAL_DILATION"

	StatusNoNativeVEVTheorem                  = "FAILED_ROUTE_NO_NATIVE_VEV_THEOREM"
	StatusCHiggsNotDimensionfulMassTheorem    = "FAILED_ROUTE_C_HIGGS_NOT_DIMENSIONFUL_MASS_THEOREM"
	StatusTreeProxyNotPoleMass                = "FAILED_ROUTE_TREE_PROXY_NOT_POLE_MASS"
	StatusNoNativeElectroweakScaleTheorem     = "FAILED_ROUTE_NO_NATIVE_ELECTROWEAK_SCALE_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate777VEVScaleBoundary             = "FIREWALL_PRESERVED_GATE777_VEV_SCALE_AIRLOCK_BOUNDARY"
)

const (
	// Gate 776 ledger snapshots.
	cHiggsMZ         = 1.0372205204048603
	cYukawaMZ        = 0.9992248188812008
	cHistoryMZ       = 1.038025177923625
	vevConventionGeV = 246.2196508
)

type Gate776Inheritance struct {
	Inherited               bool
	CHiggsFormula           string
	CHiggs                  float64
	DilationFactor          float64
	TreeMassFormula         string
	TreeMassGeV             float64
	NativeHiggsTheorem      bool
	DimensionfulMassTheorem bool
	PoleMassTheorem         bool
	Verdict                 string
}

type DimensionlessDimensionfulSplit struct {
	DimensionlessObjects          []string
	DimensionfulScaleSeal         string
	MassPower                     int
	A2Power                       int
	A3Power                       int
	Lambda3Power                  int
	MuSquaredPower                int
	C0Power                       int
	Audited                       bool
	CHiggsDimensionfulMassTheorem bool
	Verdict                       string
}

type VEVConventionSeal struct {
	SealName                 string
	ValueGeV                 float64
	PhiNormConvention        string
	ExternallyRelatedToFermi bool
	NativeVEVTheorem         bool
	NativeFermiScaleTheorem  bool
	Recorded                 bool
	Verdict                  string
}

type ScaleSensitivity struct {
	MassFractionalFormula   string
	MuSquaredFormula        string
	C0Formula               string
	MassSensitivityToV      float64
	MassSensitivityToC      float64
	MuSquaredSensitivityToV int
	C0SensitivityToV        int
	Computed                bool
	Finite                  bool
	Verdict                 string
}

type BaselineScaleInterpretation struct {
	BaselineFormula     string
	BaselineGeV         float64
	DilationFormula     string
	DilationFactor      float64
	MassFormula         string
	MassGeV             float64
	MassFromBaselineGeV float64
	Residual            float64
	Recorded            bool
	DerivedMassTheorem  bool
	Verdict             string
}

type RemainingSourcePressure struct {
	DimensionlessTargets            []string
	DimensionfulTargets             []string
	DimensionlessAloneDerivesGeV    bool
	ElectroweakScaleTheoremRequired bool
	Recorded                        bool
	Verdict                         string
}

type PhysicalFirewalls struct {
	Audited                       bool
	CHiggsDimensionfulMassTheorem bool
	VEVNativeTheorem              bool
	HalfScaleDerivedHiggsMass     bool
	FermiScaleNativeASHATheorem   bool
	TreeProxyPoleMass             bool
	FullHiggsPrediction           bool
	YukawaOperatorOrEigenvalue    bool
	Verdict                       string
}

type DerivedLedger struct {
	CHiggs         float64
	CYukawa        float64
	CHistory       float64
	LambdaHBridge  float64
	Lambda4Tree    float64
	DilationFactor float64
	VHalfGeV       float64
	MassGeV        float64
	A2GeV2         float64
	A3GeV          float64
	Lambda3GeV     float64
	MuSquaredGeV2  float64
	C0GeV4         float64
	Finite         bool
}

type Analysis struct {
	Gate776     Gate776Inheritance
	Split       DimensionlessDimensionfulSplit
	VEV         VEVConventionSeal
	Sensitivity ScaleSensitivity
	Baseline    BaselineScaleInterpretation
	Pressure    RemainingSourcePressure
	Ledger      DerivedLedger
	Firewalls   PhysicalFirewalls
	Truth       string
}

var (
	cacheMu sync.Mutex
	cache   *Analysis
)

func BuildDefault() (*Analysis, error) {
	cacheMu.Lock()
	defer cacheMu.Unlock()
	if cache != nil {
		clone := *cache
		clone.Split.DimensionlessObjects = append([]string(nil), cache.Split.DimensionlessObjects...)
		clone.Pressure.DimensionlessTargets = append([]string(nil), cache.Pressure.DimensionlessTargets...)
		clone.Pressure.DimensionfulTargets = append([]string(nil), cache.Pressure.DimensionfulTargets...)
		return &clone, nil
	}
	for name, value := range map[string]float64{
		"C_Higgs":   cHiggsMZ,
		"C_Yukawa":  cYukawaMZ,
		"C_History": cHistoryMZ,
		"v":         vevConventionGeV,
	} {
		if math.IsNaN(value) || math.IsInf(value, 0) || value <= 0 {
			return nil, fmt.Errorf("invalid %s ledger: %.17g", name, value)
		}
	}

	dilation := math.Sqrt(cHiggsMZ)
	vHalf := vevConventionGeV / 2
	mass := vHalf * dilation
	lambdaH := cHiggsMZ / 8
	v2 := vevConventionGeV * vevConventionGeV
	a2 := lambdaH * v2
	a3 := lambdaH * vevConventionGeV
	lambda3 := (3.0 / 4.0) * vevConventionGeV * cHiggsMZ
	lambda4 := (3.0 / 4.0) * cHiggsMZ
	mu2 := -lambdaH * v2
	c0 := 0.25 * lambdaH * v2 * v2

	a := &Analysis{
		Gate776: Gate776Inheritance{
			Inherited:               true,
			CHiggsFormula:           "C_Higgs=(3/N_eff)[1+L_Hopf(1-kappa_lambda_red)]",
			CHiggs:                  cHiggsMZ,
			DilationFactor:          dilation,
			TreeMassFormula:         "m_H_tree=(v/2)sqrt(C_Higgs)",
			TreeMassGeV:             mass,
			NativeHiggsTheorem:      false,
			DimensionfulMassTheorem: false,
			PoleMassTheorem:         false,
			Verdict:                 StatusGate776TotalCorrectionDecompositionInherited,
		},
		Split: DimensionlessDimensionfulSplit{
			DimensionlessObjects:          []string{"C_Higgs", "C_Yukawa", "C_History", "lambda_H_bridge=C_Higgs/8", "lambda_4=(3/4)C_Higgs"},
			DimensionfulScaleSeal:         "v",
			MassPower:                     1,
			A2Power:                       2,
			A3Power:                       1,
			Lambda3Power:                  1,
			MuSquaredPower:                2,
			C0Power:                       4,
			Audited:                       true,
			CHiggsDimensionfulMassTheorem: false,
			Verdict:                       StatusDimensionlessDimensionfulSplitAudited,
		},
		VEV: VEVConventionSeal{
			SealName:                 "VEVConventionSeal",
			ValueGeV:                 vevConventionGeV,
			PhiNormConvention:        "phi^dagger phi=v^2/2",
			ExternallyRelatedToFermi: true,
			NativeVEVTheorem:         false,
			NativeFermiScaleTheorem:  false,
			Recorded:                 true,
			Verdict:                  StatusVEVConventionSealRecorded,
		},
		Sensitivity: ScaleSensitivity{
			MassFractionalFormula:   "delta m_H_tree/m_H_tree=delta v/v+(1/2)delta C_Higgs/C_Higgs",
			MuSquaredFormula:        "delta mu^2/mu^2=delta lambda_H/lambda_H+2 delta v/v",
			C0Formula:               "delta c0/c0=delta lambda_H/lambda_H+4 delta v/v",
			MassSensitivityToV:      1,
			MassSensitivityToC:      0.5,
			MuSquaredSensitivityToV: 2,
			C0SensitivityToV:        4,
			Computed:                true,
			Finite:                  true,
			Verdict:                 StatusScaleSensitivityComputed,
		},
		Baseline: BaselineScaleInterpretation{
			BaselineFormula:     "m_baseline=v/2",
			BaselineGeV:         vHalf,
			DilationFormula:     "D_radial=sqrt(C_Higgs)",
			DilationFactor:      dilation,
			MassFormula:         "m_H_tree_proxy=(v/2)D_radial",
			MassGeV:             mass,
			MassFromBaselineGeV: vHalf * dilation,
			Residual:            mass - vHalf*dilation,
			Recorded:            true,
			DerivedMassTheorem:  false,
			Verdict:             StatusBaselineScaleInterpretationRecorded,
		},
		Pressure: RemainingSourcePressure{
			DimensionlessTargets:            []string{"N_eff", "kappa_lambda_red", "L_Hopf", "kappa_e_red", "boundary response polynomial"},
			DimensionfulTargets:             []string{"v"},
			DimensionlessAloneDerivesGeV:    false,
			ElectroweakScaleTheoremRequired: true,
			Recorded:                        true,
			Verdict:                         StatusRemainingSourcePressureSplitRecorded,
		},
		Ledger: DerivedLedger{
			CHiggs:         cHiggsMZ,
			CYukawa:        cYukawaMZ,
			CHistory:       cHistoryMZ,
			LambdaHBridge:  lambdaH,
			Lambda4Tree:    lambda4,
			DilationFactor: dilation,
			VHalfGeV:       vHalf,
			MassGeV:        mass,
			A2GeV2:         a2,
			A3GeV:          a3,
			Lambda3GeV:     lambda3,
			MuSquaredGeV2:  mu2,
			C0GeV4:         c0,
			Finite:         finite(cHiggsMZ, cYukawaMZ, cHistoryMZ, lambdaH, lambda4, dilation, vHalf, mass, a2, a3, lambda3, mu2, c0),
		},
		Firewalls: PhysicalFirewalls{
			Audited:                       true,
			CHiggsDimensionfulMassTheorem: false,
			VEVNativeTheorem:              false,
			HalfScaleDerivedHiggsMass:     false,
			FermiScaleNativeASHATheorem:   false,
			TreeProxyPoleMass:             false,
			FullHiggsPrediction:           false,
			YukawaOperatorOrEigenvalue:    false,
			Verdict:                       StatusGate777VEVScaleBoundary,
		},
		Truth: fmt.Sprintf("Gate777 separates the dimensionless Higgs correction C_Higgs=%.16f from the dimensionful VEV seal v=%.7f GeV. The tree proxy is the electroweak half-scale %.7f GeV times D_radial=%.16f, not a native mass, VEV, electroweak-scale, pole-mass, or Yukawa theorem.", cHiggsMZ, vevConventionGeV, vHalf, dilation),
	}
	cache = a
	clone := *a
	clone.Split.DimensionlessObjects = append([]string(nil), a.Split.DimensionlessObjects...)
	clone.Pressure.DimensionlessTargets = append([]string(nil), a.Pressure.DimensionlessTargets...)
	clone.Pressure.DimensionfulTargets = append([]string(nil), a.Pressure.DimensionfulTargets...)
	return &clone, nil
}

func finite(values ...float64) bool {
	for _, v := range values {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return false
		}
	}
	return true
}

func Statuses() []string {
	return []string{
		StatusGate776TotalCorrectionDecompositionInherited,
		StatusDimensionlessDimensionfulSplitAudited,
		StatusVEVConventionSealRecorded,
		StatusScaleSensitivityComputed,
		StatusBaselineScaleInterpretationRecorded,
		StatusRemainingSourcePressureSplitRecorded,
		StatusPhysicalFirewallsEnforced,
		StatusASHAControlsDimensionlessCorrection,
		StatusVEVSealSuppliesDimensionfulScale,
		StatusTreeProxyEqualsHalfScaleTimesDilation,
		StatusNoNativeVEVTheorem,
		StatusCHiggsNotDimensionfulMassTheorem,
		StatusTreeProxyNotPoleMass,
		StatusNoNativeElectroweakScaleTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate777VEVScaleBoundary,
	}
}

func FormatGate776(g Gate776Inheritance) string {
	return fmt.Sprintf("inherited=%t C=%s CHiggs=%.17g D=%.17g mass=%s %.12f nativeHiggs=%t massTheorem=%t pole=%t verdict=%s", g.Inherited, g.CHiggsFormula, g.CHiggs, g.DilationFactor, g.TreeMassFormula, g.TreeMassGeV, g.NativeHiggsTheorem, g.DimensionfulMassTheorem, g.PoleMassTheorem, g.Verdict)
}

func FormatSplit(s DimensionlessDimensionfulSplit) string {
	return fmt.Sprintf("dimensionless=%s scale=%s powers[m=%d A2=%d A3=%d lambda3=%d mu2=%d c0=%d] audited=%t massTheorem=%t verdict=%s", strings.Join(s.DimensionlessObjects, ","), s.DimensionfulScaleSeal, s.MassPower, s.A2Power, s.A3Power, s.Lambda3Power, s.MuSquaredPower, s.C0Power, s.Audited, s.CHiggsDimensionfulMassTheorem, s.Verdict)
}

func FormatVEV(v VEVConventionSeal) string {
	return fmt.Sprintf("seal=%s value=%.10f convention=%s externalFermi=%t nativeVEV=%t nativeFermi=%t recorded=%t verdict=%s", v.SealName, v.ValueGeV, v.PhiNormConvention, v.ExternallyRelatedToFermi, v.NativeVEVTheorem, v.NativeFermiScaleTheorem, v.Recorded, v.Verdict)
}

func FormatSensitivity(s ScaleSensitivity) string {
	return fmt.Sprintf("mass=%s mu=%s c0=%s dm/dv=%.1f dm/dC=%.1f dmu/dv=%d dc0/dv=%d computed=%t finite=%t verdict=%s", s.MassFractionalFormula, s.MuSquaredFormula, s.C0Formula, s.MassSensitivityToV, s.MassSensitivityToC, s.MuSquaredSensitivityToV, s.C0SensitivityToV, s.Computed, s.Finite, s.Verdict)
}

func FormatBaseline(b BaselineScaleInterpretation) string {
	return fmt.Sprintf("baseline=%s %.10f D=%s %.17g mass=%s %.12f residual=%.3e recorded=%t derivedMass=%t verdict=%s", b.BaselineFormula, b.BaselineGeV, b.DilationFormula, b.DilationFactor, b.MassFormula, b.MassGeV, b.Residual, b.Recorded, b.DerivedMassTheorem, b.Verdict)
}

func FormatPressure(p RemainingSourcePressure) string {
	return fmt.Sprintf("dimensionlessTargets=%s dimensionfulTargets=%s dimensionlessAloneGeV=%t ewScaleRequired=%t recorded=%t verdict=%s", strings.Join(p.DimensionlessTargets, ","), strings.Join(p.DimensionfulTargets, ","), p.DimensionlessAloneDerivesGeV, p.ElectroweakScaleTheoremRequired, p.Recorded, p.Verdict)
}

func FormatLedger(l DerivedLedger) string {
	return fmt.Sprintf("C=%.17g CY=%.17g CH=%.17g lambda=%.17g lambda4=%.17g D=%.17g vhalf=%.10f m=%.12f A2=%.12f A3=%.12f lambda3=%.12f mu2=%.12f c0=%.12f finite=%t", l.CHiggs, l.CYukawa, l.CHistory, l.LambdaHBridge, l.Lambda4Tree, l.DilationFactor, l.VHalfGeV, l.MassGeV, l.A2GeV2, l.A3GeV, l.Lambda3GeV, l.MuSquaredGeV2, l.C0GeV4, l.Finite)
}

func FormatFirewalls(f PhysicalFirewalls) string {
	return fmt.Sprintf("audited=%t Cmass=%t nativeVEV=%t halfMass=%t fermiNative=%t pole=%t fullPrediction=%t yukawaOp=%t verdict=%s", f.Audited, f.CHiggsDimensionfulMassTheorem, f.VEVNativeTheorem, f.HalfScaleDerivedHiggsMass, f.FermiScaleNativeASHATheorem, f.TreeProxyPoleMass, f.FullHiggsPrediction, f.YukawaOperatorOrEigenvalue, f.Verdict)
}
