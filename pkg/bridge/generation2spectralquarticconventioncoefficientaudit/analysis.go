// Package generation2spectralquarticconventioncoefficientaudit implements
// Gate 619: Spectral Quartic Convention Coefficient c_lambda Audit.
//
// Gate 618 left the scalar airlock blocked at the convention coefficient
// c_lambda in the formal relation lambda_canon ?= c_lambda * b/a^2.
// Gate 619 audits the convention family, runtime b/a^2 diagnostics, the sign
// obstruction, and the separation between runtime lambda and spectral-action
// boundary quartic. It does not derive a Higgs mass, scalar stability,
// lambda=0 boundary, gauge unification, or native GaugeScalarBoundaryStressSeal.
package generation2spectralquarticconventioncoefficientaudit

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2spectralactionabf0canonicalscalarquarticairlockaudit"
)

const (
	AuditID = "GATE619-SPECTRAL-QUARTIC-CONVENTION-COEFFICIENT-C-LAMBDA-AUDIT"

	StatusGate618Inherited           = "PASS_GATE618_C_LAMBDA_BLOCKER_INHERITED"
	StatusConventionFamilyClassified = "PASS_CONVENTION_FAMILY_CLASSIFIED"
	StatusCLambdaTargetDefined       = "PASS_FORMAL_C_LAMBDA_TARGET_DEFINED"
	StatusBA2DiagnosticComputed      = "PASS_RUNTIME_B_OVER_A_SQUARED_DIAGNOSTIC_COMPUTED"
	StatusSignAudited                = "PASS_B_OVER_A_SQUARED_SIGN_AUDITED"
	StatusRGSeparationAudited        = "PASS_RG_TRANSPORT_SEPARATION_AUDITED"
	StatusSymbolicBOverA2            = "CONDITIONAL_SUPPORT_LAMBDA_CANON_B_OVER_A_SQUARED_FORM_REMAINS_SYMBOLIC"
	StatusNoCLambdaValue             = "FAILED_ROUTE_NO_CERTIFIED_C_LAMBDA_VALUE"
	StatusNegativeRuntimeNotDirect   = "FAILED_ROUTE_NEGATIVE_RUNTIME_LAMBDA_NOT_DIRECT_POSITIVE_B_OVER_A_SQUARED_BOUNDARY"
	StatusNoRuntimeMatchingTheorem   = "FAILED_ROUTE_NO_NATIVE_RUNTIME_MATCHING_THEOREM"
	StatusStressRuntimeShadow        = "FAILED_ROUTE_STRESS_SEAL_SCALAR_SIDE_REMAINS_RUNTIME_SHADOW"
	StatusNoNativeKPhi               = "FAILED_ROUTE_NO_NATIVE_K_PHI_THEOREM"
	StatusNoNativeLambdaPhi          = "FAILED_ROUTE_NO_NATIVE_LAMBDA_PHI_THEOREM"
	StatusNoNativeVEV                = "FAILED_ROUTE_NO_NATIVE_HIGGS_VEV_OR_MATCHING_THEOREM"
	StatusGate619Boundary            = "FIREWALL_PRESERVED_GATE619_C_LAMBDA_BOUNDARY"
)

const (
	// Runtime scalar data inherited through Gates 606-618.
	lambdaRuntimeMZ       = 0.1296525650504758
	lambdaRuntimeLambda12 = -0.049700942077683274
	xiBoundary            = 0.0503471644870914
	r3MinusOne            = 0.0509933868964996

	// Visible v1 runtime Yukawa-trace diagnostics with neutrino sector absent/skipped.
	// These are not native derivations; they are endpoint/history ledgers.
	aMZVisible        = 2.8424095142339083
	bMZVisible        = 2.6910096440382287
	bOverA2MZVisible  = 0.33307493962706697
	aL12Visible       = 0.6941198223775996
	bL12Visible       = 0.16047699018700937
	bOverA2L12Visible = 0.3330764110541872
)

type Inherited struct {
	LambdaRuntimeLambda12 float64
	R3MinusOne            float64
	XiBoundary            float64
	FormalCandidate       string
	Blocker               string
	Verdict               string
}

type ConventionFamilyRow struct {
	Factor           string
	Role             string
	CanChangeCLambda bool
	Certified        bool
	Impact           string
	Verdict          string
}

type CandidateFormula struct {
	Name         string
	Formula      string
	Native       bool
	Certified    bool
	RequiredData []string
	Statement    string
	Verdict      string
}

type BA2Diagnostic struct {
	Scale                  string
	ATrace                 float64
	BTrace                 float64
	BOverA2                float64
	LambdaRuntime          float64
	CLambdaRequiredRuntime float64
	CLambdaRequiredZero    float64
	CLambdaRequiredNegXi   float64
	NeutrinoIncluded       bool
	Complete               bool
	Statement              string
	Verdict                string
}

type SignAudit struct {
	BOverA2NonNegative              bool
	PositiveCLambdaGivesPositive    bool
	LambdaRuntimeAtLambda12Negative bool
	DirectPositiveBoundaryPossible  bool
	Statement                       string
	Verdict                         string
}

type RGTransportSeparation struct {
	LambdaRuntimeBoundaryInitial bool
	LambdaRuntimeRunUpLedger     bool
	LambdaRuntimeSpectralQuartic bool
	RequiresMatchingTheorem      bool
	RequiresLoopThresholdControl bool
	Statement                    string
	Verdict                      string
}

type StressSealImpact struct {
	UsesLambdaRuntimeShadow bool
	CanUseLambdaCanon       bool
	CanUseBA2Directly       bool
	LambdaRuntime           float64
	R3MinusOne              float64
	XiBoundary              float64
	Statement               string
	Verdict                 string
}

type NativeStatus struct {
	NativeCLambda        bool
	NativeKPhi           bool
	NativeLambdaPhi      bool
	NativeBA2Theorem     bool
	NativeRuntimeMatch   bool
	NativeSignConvention bool
	NativeVEV            bool
	Statement            string
	Verdict              string
}

type Firewalls struct {
	ClaimsHiggsMass        bool
	ClaimsHiggsStability   bool
	ClaimsLambdaZero       bool
	ClaimsGaugeUnification bool
	ClaimsNativeStress     bool
	ClaimsNativeCLambda    bool
	Verdict                string
}

type Analysis struct {
	Inherited         Inherited
	Conventions       []ConventionFamilyRow
	CandidateFormulas []CandidateFormula
	Diagnostics       []BA2Diagnostic
	SignAudit         SignAudit
	RGSeparation      RGTransportSeparation
	StressImpact      StressSealImpact
	NativeStatus      NativeStatus
	Firewalls         Firewalls
	Truth             string
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
	g618, err := generation2spectralactionabf0canonicalscalarquarticairlockaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate618 predecessor: %w", err)
	}
	inherited := inherit(g618)
	a := Analysis{
		Inherited:         inherited,
		Conventions:       buildConventionFamily(),
		CandidateFormulas: buildCandidateFormulas(),
		Diagnostics:       buildBA2Diagnostics(),
		SignAudit:         buildSignAudit(),
		RGSeparation:      buildRGSeparation(),
		StressImpact:      buildStressImpact(),
		NativeStatus:      buildNativeStatus(),
		Firewalls:         auditFirewalls(),
		Truth:             "Gate 619 isolates c_lambda as the convention/normalization blocker in the formal lambda_canon ?= c_lambda*b/a^2 scalar airlock. Runtime b/a^2 diagnostics can be computed from observed Yukawa ledgers, but b/a^2 is nonnegative, while lambda_runtime(Lambda_12) is negative; therefore negative runtime lambda cannot be a direct positive spectral-action boundary quartic. Without certified c_lambda, K_phi, Lambda_phi, sign conventions, and runtime matching, the GaugeScalarBoundaryStressSeal scalar side remains a runtime shadow.",
	}
	return a, nil
}

func inherit(a generation2spectralactionabf0canonicalscalarquarticairlockaudit.Analysis) Inherited {
	return Inherited{
		LambdaRuntimeLambda12: a.Inherited.LambdaRuntime,
		R3MinusOne:            a.Inherited.R3MinusOne,
		XiBoundary:            a.Inherited.XiBoundary,
		FormalCandidate:       a.RatioAudit.CandidateFormula,
		Blocker:               StatusNoCLambdaValue,
		Verdict:               StatusGate618Inherited,
	}
}

func buildConventionFamily() []ConventionFamilyRow {
	v := StatusNoCLambdaValue
	return []ConventionFamilyRow{
		{Factor: "real versus complex scalar normalization", Role: "changes scalar component counting and quartic normalization", CanChangeCLambda: true, Certified: false, Impact: "can multiply c_lambda by convention factors", Verdict: v},
		{Factor: "Higgs doublet normalization", Role: "maps spectral scalar variable to SM Higgs doublet", CanChangeCLambda: true, Certified: false, Impact: "changes K_phi and quartic normalization", Verdict: v},
		{Factor: "|H|^4 versus (H†H)^2 convention", Role: "potential notation", CanChangeCLambda: true, Certified: false, Impact: "can shift factors of 2 or 4", Verdict: v},
		{Factor: "Euclidean to Lorentzian sign transfer", Role: "action/sign convention", CanChangeCLambda: true, Certified: false, Impact: "controls sign interpretation of pre-canonical scalar coefficient", Verdict: v},
		{Factor: "spectral-action f0 normalization", Role: "heat-kernel coefficient normalization", CanChangeCLambda: true, Certified: false, Impact: "sets common prefactor before canonical rescaling", Verdict: v},
		{Factor: "trace normalization and representation multiplicities", Role: "finite trace convention", CanChangeCLambda: true, Certified: false, Impact: "sets a,b coefficient normalization", Verdict: v},
		{Factor: "scalar field rescaling by K_phi", Role: "canonical normalization", CanChangeCLambda: true, Certified: false, Impact: "lambda_canon=Lambda_phi/K_phi^2 up to convention", Verdict: v},
		{Factor: "Standard Model potential convention", Role: "runtime lambda convention", CanChangeCLambda: true, Certified: false, Impact: "matches lambda_runtime=m_H^2/(2v^2) convention", Verdict: v},
	}
}

func buildCandidateFormulas() []CandidateFormula {
	return []CandidateFormula{
		{
			Name:         "formal spectral-action scalar ratio",
			Formula:      "lambda_canon ?= c_lambda * b/a^2",
			Native:       false,
			Certified:    false,
			RequiredData: []string{"c_lambda", "K_phi", "Lambda_phi", "f0 convention", "scalar normalization", "potential convention"},
			Statement:    "The b/a^2 shape remains the lawful symbolic target, but the normalization coefficient is not certified.",
			Verdict:      StatusSymbolicBOverA2,
		},
		{
			Name:         "runtime matching target",
			Formula:      "lambda_runtime(mu) ?= RG_mu[lambda_canon(Lambda)]",
			Native:       false,
			Certified:    false,
			RequiredData: []string{"boundary scale", "matching theorem", "loop order", "thresholds", "scheme conversion"},
			Statement:    "No theorem identifies the runtime transported SM quartic with the spectral-action boundary quartic.",
			Verdict:      StatusNoRuntimeMatchingTheorem,
		},
	}
}

func buildBA2Diagnostics() []BA2Diagnostic {
	return []BA2Diagnostic{
		makeDiagnostic("M_Z", aMZVisible, bMZVisible, bOverA2MZVisible, lambdaRuntimeMZ),
		makeDiagnostic("Lambda_12", aL12Visible, bL12Visible, bOverA2L12Visible, lambdaRuntimeLambda12),
	}
}

func makeDiagnostic(scale string, a, b, ratio, lambda float64) BA2Diagnostic {
	return BA2Diagnostic{
		Scale:                  scale,
		ATrace:                 a,
		BTrace:                 b,
		BOverA2:                ratio,
		LambdaRuntime:          lambda,
		CLambdaRequiredRuntime: lambda / ratio,
		CLambdaRequiredZero:    0,
		CLambdaRequiredNegXi:   -xiBoundary / ratio,
		NeutrinoIncluded:       false,
		Complete:               false,
		Statement:              "visible v1 runtime diagnostic with neutrino sector absent/skipped; useful for sign and scale checks only, not a native a,b theorem",
		Verdict:                StatusBA2DiagnosticComputed,
	}
}

func buildSignAudit() SignAudit {
	return SignAudit{
		BOverA2NonNegative:              true,
		PositiveCLambdaGivesPositive:    true,
		LambdaRuntimeAtLambda12Negative: lambdaRuntimeLambda12 < 0,
		DirectPositiveBoundaryPossible:  false,
		Statement:                       "Because b/a^2 is nonnegative for positive Yukawa singular values, negative lambda_runtime(Lambda_12) cannot be a direct positive-c_lambda spectral boundary quartic. It can only arise after RG transport, sign convention change, or an unproven matching relation.",
		Verdict:                         StatusNegativeRuntimeNotDirect,
	}
}

func buildRGSeparation() RGTransportSeparation {
	return RGTransportSeparation{
		LambdaRuntimeBoundaryInitial: false,
		LambdaRuntimeRunUpLedger:     true,
		LambdaRuntimeSpectralQuartic: false,
		RequiresMatchingTheorem:      true,
		RequiresLoopThresholdControl: true,
		Statement:                    "lambda_runtime(Lambda_12) is the v1 SM canonical quartic run upward from M_Z, not a certified spectral-action boundary coefficient.",
		Verdict:                      StatusRGSeparationAudited,
	}
}

func buildStressImpact() StressSealImpact {
	return StressSealImpact{
		UsesLambdaRuntimeShadow: true,
		CanUseLambdaCanon:       false,
		CanUseBA2Directly:       false,
		LambdaRuntime:           lambdaRuntimeLambda12,
		R3MinusOne:              r3MinusOne,
		XiBoundary:              xiBoundary,
		Statement:               "GaugeScalarBoundaryStressSeal remains (R3-1, lambda_runtime)≈(+xi,-xi); it cannot yet be lifted to (R3-1, c_lambda*b/a^2) or a native spectral-action coefficient relation.",
		Verdict:                 StatusStressRuntimeShadow,
	}
}

func buildNativeStatus() NativeStatus {
	return NativeStatus{
		NativeCLambda:        false,
		NativeKPhi:           false,
		NativeLambdaPhi:      false,
		NativeBA2Theorem:     false,
		NativeRuntimeMatch:   false,
		NativeSignConvention: false,
		NativeVEV:            false,
		Statement:            "The project has not certified c_lambda, K_phi, Lambda_phi, a,b-to-lambda normalization, runtime matching, sign convention closure, or Higgs VEV theorem.",
		Verdict:              StatusNoCLambdaValue,
	}
}

func auditFirewalls() Firewalls {
	return Firewalls{Verdict: StatusGate619Boundary}
}

func Statuses() []string {
	return []string{
		StatusGate618Inherited,
		StatusConventionFamilyClassified,
		StatusCLambdaTargetDefined,
		StatusBA2DiagnosticComputed,
		StatusSignAudited,
		StatusRGSeparationAudited,
		StatusSymbolicBOverA2,
		StatusNoCLambdaValue,
		StatusNegativeRuntimeNotDirect,
		StatusNoRuntimeMatchingTheorem,
		StatusStressRuntimeShadow,
		StatusNoNativeKPhi,
		StatusNoNativeLambdaPhi,
		StatusNoNativeVEV,
		StatusGate619Boundary,
	}
}

func nearly(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
