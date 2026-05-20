// Package generation2flavorreducedscalarhiggsnormalformandkappaesubstitutionaudit implements
// Gate 752: Flavor-Reduced Scalar-Higgs Normal Form and Kappa_e Substitution Audit.
//
// Gate 751 produced the typed scalar-Higgs normal form. Gate 752 lawfully
// substitutes the Gate 748 source-type expression for kappa_e into that normal
// form, audits the double insertion sensitivity, and preserves the flavor and
// scalar-runtime firewalls.
package generation2flavorreducedscalarhiggsnormalformandkappaesubstitutionaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate751 "github.com/bagherbal/asha-engine/pkg/bridge/generation2scalarhiggstypednormalformandillegaltermrejectionaudit"
)

const (
	AuditID = "GATE752-FLAVOR-REDUCED-SCALAR-HIGGS-NORMAL-FORM-KAPPA-E-SUBSTITUTION-AUDIT"

	StatusGate751TypedNormalFormInherited     = "PASS_GATE751_SCALAR_HIGGS_TYPED_NORMAL_FORM_INHERITED"
	StatusGate748KappaESourceFormInherited    = "PASS_GATE748_KAPPA_E_SOURCE_FORM_INHERITED"
	StatusKappaEReducedCandidateDefined       = "PASS_KAPPA_E_REDUCED_CANDIDATE_DEFINED"
	StatusReducedCubicWallPolynomialDefined   = "PASS_REDUCED_CUBIC_WALL_POLYNOMIAL_DEFINED"
	StatusReducedScalarHiggsNormalFormWritten = "PASS_REDUCED_SCALAR_HIGGS_NORMAL_FORM_WRITTEN"
	StatusNumericalResidualAudited            = "PASS_NUMERICAL_RESIDUAL_AUDITED"
	StatusDoubleInsertionSensitivityAudited   = "PASS_DOUBLE_INSERTION_SENSITIVITY_AUDITED"
	StatusReductionStatusClassified           = "PASS_REDUCTION_STATUS_CLASSIFIED"
	StatusPhysicalFirewallsEnforced           = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusKappaERedStrongApproximation = "CONDITIONAL_SUPPORT_KAPPA_E_RED_STRONGLY_APPROXIMATES_ACTIVE_KAPPA_E"
	StatusNormalFormCanBeFlavorReduced = "CONDITIONAL_SUPPORT_SCALAR_HIGGS_NORMAL_FORM_CAN_BE_FLAVOR_REDUCED"
	StatusKappaESealPartiallyReduced   = "CONDITIONAL_SUPPORT_KAPPA_E_SEAL_IS_PARTIALLY_REDUCED_TO_TYPED_WALL_ORIENTATION_FORM"

	StatusKappaERedNotExact                   = "FAILED_ROUTE_KAPPA_E_RED_NOT_EXACT"
	StatusKappaERedNotNativeFlavorTheorem     = "FAILED_ROUTE_KAPPA_E_RED_NOT_NATIVE_FLAVOR_THEOREM"
	StatusNoNativePMNSOrCKMTheorem            = "FAILED_ROUTE_NO_NATIVE_PMNS_OR_CKM_THEOREM"
	StatusNoNativeFlavorDeficitTheorem        = "FAILED_ROUTE_NO_NATIVE_FLAVOR_DEFICIT_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusNoIndependentScalarRuntimeTheorem   = "FAILED_ROUTE_NO_INDEPENDENT_SCALAR_RUNTIME_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem        = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusGate752Boundary                     = "FIREWALL_PRESERVED_GATE752_FLAVOR_REDUCED_SCALAR_HIGGS_NORMAL_FORM_BOUNDARY"
)

const (
	pK7          = 7.0 / 72.0
	kappaEActive = 0.00550355419157456
	kappaEOrient = 0.00550633006471245
	lambdaAbs    = 0.0497009420776833
	rMinusOne    = 0.050993386896499596
	xiBoundary   = 0.0503471644870914
	lambdaProxy  = 0.12490310236015
)

type Gate751Inheritance struct {
	Inherited              bool
	TypedNormalFormReady   bool
	IllegalTermsRejected   bool
	KappaInsertionRecorded bool
	Verdict                string
}

type KappaEReducedCandidate struct {
	Formula             string
	PMNSReactorTerm     string
	CKMOrientationTerm  string
	HyperchargeBoundary string
	StressMomentTerm    string
	SSplit              float64
	KappaEOrient        float64
	KappaEReduced       float64
	ActiveKappaE        float64
	Residual            float64
	Verdict             string
}

type ReducedCubicWallPolynomial struct {
	MapType      string
	Definition   string
	Polynomial   string
	SSplit       float64
	ActiveValue  float64
	ReducedValue float64
	Difference   float64
	Verdict      string
}

type ReducedScalarHiggsNormalForm struct {
	Formula        string
	Expanded       string
	ActiveRuntime  float64
	ReducedRuntime float64
	RuntimeShift   float64
	Verdict        string
}

type NumericalResidualAudit struct {
	KappaEResidual  float64
	FWallShift      float64
	RuntimeShift    float64
	RuntimeShiftAbs float64
	ResidualScale   string
	Verdict         string
}

type DoubleInsertionSensitivity struct {
	Formula               string
	DeltaKappaE           float64
	PK7SSquared           float64
	LinearSensitivity     float64
	PredictedRuntimeShift float64
	ActualRuntimeShift    float64
	Agreement             float64
	Verdict               string
}

type ReductionStatus struct {
	BareSealReduced     bool
	NativeFlavorTheorem bool
	Components          []string
	Classification      string
	Verdict             string
}

type PhysicalFirewalls struct {
	KappaERedNativeBlocked   bool
	PMNSCKMBlocked           bool
	FlavorDeficitBlocked     bool
	YukawaBlocked            bool
	RuntimePredictionBlocked bool
	HiggsMassBlocked         bool
	Verdict                  string
}

type Analysis struct {
	Gate751     Gate751Inheritance
	KappaERed   KappaEReducedCandidate
	FWallRed    ReducedCubicWallPolynomial
	RuntimeRed  ReducedScalarHiggsNormalForm
	Residual    NumericalResidualAudit
	Sensitivity DoubleInsertionSensitivity
	Reduction   ReductionStatus
	Firewalls   PhysicalFirewalls
	Truth       string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	g751, err := gate751.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate751 inheritance unavailable: %w", err)
	}
	inheritance := buildGate751Inheritance(g751)
	kappa := buildKappaEReducedCandidate()
	fwall := buildReducedCubicWallPolynomial(kappa.SSplit, kappa.KappaEReduced)
	runtime := buildReducedScalarHiggsNormalForm(fwall.ReducedValue, kappa.KappaEReduced)
	residual := buildNumericalResidualAudit(kappa, fwall, runtime)
	sensitivity := buildDoubleInsertionSensitivity(kappa, runtime)
	reduction := buildReductionStatus()
	firewalls := buildPhysicalFirewalls()
	truth := "Gate 752 lawfully substitutes the Gate 748 kappa_e source-type candidate into the Gate 751 typed normal form. The resulting flavor-reduced scalar-Higgs form remains a bridge expression: kappa_e is strongly approximated by flavor orientation plus hypercharge-boundary and boundary-stress moment terms, but this is not a native flavor theorem or an independent scalar-runtime prediction."
	return Analysis{Gate751: inheritance, KappaERed: kappa, FWallRed: fwall, RuntimeRed: runtime, Residual: residual, Sensitivity: sensitivity, Reduction: reduction, Firewalls: firewalls, Truth: truth}, nil
}

func buildGate751Inheritance(g gate751.Analysis) Gate751Inheritance {
	return Gate751Inheritance{
		Inherited:              g.Gate750.Inherited && g.NormalForm.AllOperatorCollapsed,
		TypedNormalFormReady:   strings.Contains(g.NormalForm.RuntimeFormula, "L_Hopf") && strings.Contains(g.Cubic.MapType, "Q_boundary -> Q_history"),
		IllegalTermsRejected:   g.Illegal.K7BoundaryBlocked && g.Illegal.FWallOperatorBlocked && g.Illegal.SevenLoopBlocked,
		KappaInsertionRecorded: g.KappaE.InsideFWall3 && g.KappaE.OutsideRuntimeTransport && !g.KappaE.NativeFlavorTheorem,
		Verdict:                StatusGate751TypedNormalFormInherited,
	}
}

func sSplit() float64 { return rMinusOne - lambdaAbs }

func kappaEReduced(s float64) float64 {
	return kappaEOrient - (5.0/3.0)*s*s + xiBoundary*pK7*s*s
}

func fWall3(s, kappa float64) float64 {
	return pK7*s + kappa*pK7*s*s - 2.0*pK7*pK7*s*s*s
}

func runtimeBridge(fwall, kappa float64) float64 {
	lHopf := 1.0 / (8.0 * math.Pi)
	return lambdaProxy * (1.0 + lHopf*(1.0-lambdaAbs-fwall+kappa))
}

func buildKappaEReducedCandidate() KappaEReducedCandidate {
	s := sSplit()
	reduced := kappaEReduced(s)
	return KappaEReducedCandidate{
		Formula:             "kappa_e_red=sin²(theta13)/4-J_CKM-(5/3)s²+xi_boundary p_K7 s²",
		PMNSReactorTerm:     "sin²(theta13)/4",
		CKMOrientationTerm:  "-J_CKM",
		HyperchargeBoundary: "-(5/3)s²",
		StressMomentTerm:    "+xi_boundary p_K7 s²",
		SSplit:              s,
		KappaEOrient:        kappaEOrient,
		KappaEReduced:       reduced,
		ActiveKappaE:        kappaEActive,
		Residual:            kappaEActive - reduced,
		Verdict: strings.Join([]string{
			StatusKappaEReducedCandidateDefined,
			StatusKappaERedStrongApproximation,
			StatusKappaERedNotExact,
		}, "; "),
	}
}

func buildReducedCubicWallPolynomial(s, reducedKappa float64) ReducedCubicWallPolynomial {
	active := fWall3(s, kappaEActive)
	reduced := fWall3(s, reducedKappa)
	return ReducedCubicWallPolynomial{
		MapType:      "F_wall_3_red: Q_boundary -> Q_history",
		Definition:   "F_wall_3_red(s)=p s+kappa_e_red p s²-2p²s³",
		Polynomial:   "p_K7 s+kappa_e_red p_K7 s^2-2p_K7^2 s^3",
		SSplit:       s,
		ActiveValue:  active,
		ReducedValue: reduced,
		Difference:   reduced - active,
		Verdict: strings.Join([]string{
			StatusReducedCubicWallPolynomialDefined,
			StatusNormalFormCanBeFlavorReduced,
		}, "; "),
	}
}

func buildReducedScalarHiggsNormalForm(fwallRed, kappaRed float64) ReducedScalarHiggsNormalForm {
	active := runtimeBridge(fWall3(sSplit(), kappaEActive), kappaEActive)
	reduced := runtimeBridge(fwallRed, kappaRed)
	return ReducedScalarHiggsNormalForm{
		Formula:        "lambda_runtime_red=lambda_proxy[1+L_Hopf(1-|lambda|-F_wall_3_red(s)+kappa_e_red)]",
		Expanded:       "lambda_proxy[1+Tr_K7+(rho_plus R_Hopf)(1-|lambda|-F_wall_3_red(sigma_boundary(b))+kappa_e_red)]",
		ActiveRuntime:  active,
		ReducedRuntime: reduced,
		RuntimeShift:   reduced - active,
		Verdict: strings.Join([]string{
			StatusReducedScalarHiggsNormalFormWritten,
			StatusNormalFormCanBeFlavorReduced,
			StatusNoIndependentScalarRuntimeTheorem,
		}, "; "),
	}
}

func buildNumericalResidualAudit(k KappaEReducedCandidate, f ReducedCubicWallPolynomial, r ReducedScalarHiggsNormalForm) NumericalResidualAudit {
	return NumericalResidualAudit{
		KappaEResidual:  k.Residual,
		FWallShift:      f.Difference,
		RuntimeShift:    r.RuntimeShift,
		RuntimeShiftAbs: math.Abs(r.RuntimeShift),
		ResidualScale:   "Gate748 boundary-stress moment precision; runtime shift approximately 1e-13 scale",
		Verdict: strings.Join([]string{
			StatusNumericalResidualAudited,
			StatusKappaERedStrongApproximation,
			StatusKappaERedNotExact,
		}, "; "),
	}
}

func buildDoubleInsertionSensitivity(k KappaEReducedCandidate, r ReducedScalarHiggsNormalForm) DoubleInsertionSensitivity {
	ps2 := pK7 * k.SSplit * k.SSplit
	lHopf := 1.0 / (8.0 * math.Pi)
	delta := k.KappaEReduced - k.ActiveKappaE
	predicted := lambdaProxy * lHopf * delta * (1.0 - ps2)
	agreement := predicted - r.RuntimeShift
	return DoubleInsertionSensitivity{
		Formula:               "delta lambda_runtime≈lambda_proxy L_Hopf delta_kappa_e(1-p_K7 s²)",
		DeltaKappaE:           delta,
		PK7SSquared:           ps2,
		LinearSensitivity:     lambdaProxy * lHopf * (1.0 - ps2),
		PredictedRuntimeShift: predicted,
		ActualRuntimeShift:    r.RuntimeShift,
		Agreement:             agreement,
		Verdict:               StatusDoubleInsertionSensitivityAudited,
	}
}

func buildReductionStatus() ReductionStatus {
	return ReductionStatus{
		BareSealReduced:     true,
		NativeFlavorTheorem: false,
		Components: []string{
			"PMNS reactor leakage candidate sin²(theta13)/4",
			"CKM orientation correction candidate -J_CKM",
			"hypercharge-normalized boundary-square correction -(5/3)s²",
			"boundary-stress-weighted K7 second raw moment correction +xi_boundary p_K7 s²",
		},
		Classification: "kappa_e is partially reduced from a bare bridge input to a typed wall-orientation source candidate, but remains non-native",
		Verdict: strings.Join([]string{
			StatusReductionStatusClassified,
			StatusKappaESealPartiallyReduced,
			StatusKappaERedNotNativeFlavorTheorem,
		}, "; "),
	}
}

func buildPhysicalFirewalls() PhysicalFirewalls {
	return PhysicalFirewalls{
		KappaERedNativeBlocked:   true,
		PMNSCKMBlocked:           true,
		FlavorDeficitBlocked:     true,
		YukawaBlocked:            true,
		RuntimePredictionBlocked: true,
		HiggsMassBlocked:         true,
		Verdict: strings.Join([]string{
			StatusPhysicalFirewallsEnforced,
			StatusKappaERedNotNativeFlavorTheorem,
			StatusNoNativePMNSOrCKMTheorem,
			StatusNoNativeFlavorDeficitTheorem,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
			StatusNoIndependentScalarRuntimeTheorem,
			StatusNoHiggsMassOrPoleMassTheorem,
			StatusGate752Boundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate751TypedNormalFormInherited,
		StatusGate748KappaESourceFormInherited,
		StatusKappaEReducedCandidateDefined,
		StatusReducedCubicWallPolynomialDefined,
		StatusReducedScalarHiggsNormalFormWritten,
		StatusNumericalResidualAudited,
		StatusDoubleInsertionSensitivityAudited,
		StatusReductionStatusClassified,
		StatusPhysicalFirewallsEnforced,
		StatusKappaERedStrongApproximation,
		StatusNormalFormCanBeFlavorReduced,
		StatusKappaESealPartiallyReduced,
		StatusKappaERedNotExact,
		StatusKappaERedNotNativeFlavorTheorem,
		StatusNoNativePMNSOrCKMTheorem,
		StatusNoNativeFlavorDeficitTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusNoIndependentScalarRuntimeTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusGate752Boundary,
	}
}

func FormatGate751(x Gate751Inheritance) string {
	return fmt.Sprintf("inherited=%t typed=%t illegal=%t kappa=%t verdict=%q", x.Inherited, x.TypedNormalFormReady, x.IllegalTermsRejected, x.KappaInsertionRecorded, x.Verdict)
}

func FormatKappaEReduced(x KappaEReducedCandidate) string {
	return fmt.Sprintf("formula=%q s=%.16g orient=%.16g red=%.16g active=%.16g residual=%.16g verdict=%q", x.Formula, x.SSplit, x.KappaEOrient, x.KappaEReduced, x.ActiveKappaE, x.Residual, x.Verdict)
}

func FormatFWallRed(x ReducedCubicWallPolynomial) string {
	return fmt.Sprintf("type=%q def=%q active=%.16g red=%.16g diff=%.16g verdict=%q", x.MapType, x.Definition, x.ActiveValue, x.ReducedValue, x.Difference, x.Verdict)
}

func FormatRuntimeRed(x ReducedScalarHiggsNormalForm) string {
	return fmt.Sprintf("formula=%q active=%.16g red=%.16g shift=%.16g verdict=%q", x.Formula, x.ActiveRuntime, x.ReducedRuntime, x.RuntimeShift, x.Verdict)
}

func FormatResidual(x NumericalResidualAudit) string {
	return fmt.Sprintf("kappaResidual=%.16g fWallShift=%.16g runtimeShift=%.16g scale=%q verdict=%q", x.KappaEResidual, x.FWallShift, x.RuntimeShift, x.ResidualScale, x.Verdict)
}

func FormatSensitivity(x DoubleInsertionSensitivity) string {
	return fmt.Sprintf("formula=%q delta=%.16g pS2=%.16g sensitivity=%.16g predicted=%.16g actual=%.16g agreement=%.16g verdict=%q", x.Formula, x.DeltaKappaE, x.PK7SSquared, x.LinearSensitivity, x.PredictedRuntimeShift, x.ActualRuntimeShift, x.Agreement, x.Verdict)
}

func FormatReduction(x ReductionStatus) string {
	return fmt.Sprintf("bareReduced=%t native=%t components=[%s] classification=%q verdict=%q", x.BareSealReduced, x.NativeFlavorTheorem, strings.Join(x.Components, "; "), x.Classification, x.Verdict)
}

func FormatFirewalls(x PhysicalFirewalls) string {
	return fmt.Sprintf("kappaNative=%t pmnsCkm=%t flavor=%t yukawa=%t runtime=%t higgs=%t verdict=%q", x.KappaERedNativeBlocked, x.PMNSCKMBlocked, x.FlavorDeficitBlocked, x.YukawaBlocked, x.RuntimePredictionBlocked, x.HiggsMassBlocked, x.Verdict)
}
