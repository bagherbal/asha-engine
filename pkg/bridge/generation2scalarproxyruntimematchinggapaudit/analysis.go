// Package generation2scalarproxyruntimematchinggapaudit implements
// Gate 621: Scalar Tree-Proxy to Runtime Matching Gap Audit.
//
// Gate 620 separated the positive spectral/tree scalar proxy lane
// lambda_proxy=(3/8)(b/a^2) from the RG-transported runtime Standard Model
// quartic lambda_runtime(mu). Gate 621 audits the low-scale matching gap
// lambda_runtime(M_Z)-lambda_proxy(M_Z) as a typed bridge correction. It does
// not derive a Higgs mass, scalar stability, a lambda-zero boundary, gauge
// unification, or a native spectral scalar theorem.
package generation2scalarproxyruntimematchinggapaudit

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2ba2onethirdrigidityspectralquarticproxyaudit"
)

const (
	AuditID = "GATE621-SCALAR-TREE-PROXY-TO-RUNTIME-MATCHING-GAP-AUDIT"

	StatusGate620Inherited          = "PASS_GATE620_PROXY_RUNTIME_SEPARATION_INHERITED"
	StatusMatchingGapComputed       = "PASS_LOW_SCALE_MATCHING_GAP_COMPUTED"
	StatusEffectiveCLambdaComputed  = "PASS_EFFECTIVE_C_LAMBDA_CORRECTION_COMPUTED"
	StatusHiggsProxyGapComputed     = "PASS_HIGGS_PROXY_GAP_DIAGNOSTIC_COMPUTED"
	StatusProxyCloseAtMZ            = "CONDITIONAL_SUPPORT_LAMBDA_PROXY_CLOSE_TO_RUNTIME_LAMBDA_AT_MZ"
	StatusPositiveMatchingRequired  = "CONDITIONAL_SUPPORT_POSITIVE_MATCHING_CORRECTION_REQUIRED"
	StatusProxyRuntimeChainDefined  = "CONDITIONAL_SUPPORT_PROXY_TO_RUNTIME_CHAIN_DEFINED"
	StatusNoMatchingTheorem         = "FAILED_ROUTE_NO_NATIVE_MATCHING_CORRECTION_THEOREM"
	StatusNoThreeEighthsTheorem     = "FAILED_ROUTE_NO_NATIVE_C_LAMBDA_THREE_EIGHTHS_THEOREM"
	StatusNoHiggsPoleTheorem        = "FAILED_ROUTE_NO_NATIVE_HIGGS_MASS_OR_POLE_THEOREM"
	StatusNoProxyRuntimeTheorem     = "FAILED_ROUTE_NO_NATIVE_PROXY_TO_RUNTIME_MATCHING_THEOREM"
	StatusNoNeutrinoTraceCompletion = "FAILED_ROUTE_NO_NATIVE_NEUTRINO_TRACE_COMPLETION_THEOREM"
	StatusGate621Boundary           = "FIREWALL_PRESERVED_GATE621_SCALAR_PROXY_MATCHING_BOUNDARY"
)

const (
	cProxy     = 3.0 / 8.0
	lambdaL12  = -0.049700942077683274
	vRuntime   = 246.21965079413738
	ba2MZ      = 0.33307493962706697
	ba2L12     = 0.3330764110541872
	lambdaMZ   = 0.1296525650504758
	proxyMZ    = cProxy * ba2MZ
	proxyL12   = cProxy * ba2L12
	mZLabel    = "M_Z"
	lambda12ID = "Lambda_12"
)

type Inherited struct {
	LambdaProxyMZ          float64
	LambdaRuntimeMZ        float64
	LambdaProxyLambda12    float64
	LambdaRuntimeLambda12  float64
	BA2MZ                  float64
	BA2Lambda12            float64
	PreviousSeparation     string
	PreviousStressVerdict  string
	PreviousProxyClose     string
	PreviousHighScaleBlock string
	Verdict                string
}

type MatchingGapTable struct {
	Scale                         string
	LambdaProxy                   float64
	LambdaRuntime                 float64
	DeltaLambdaMatch              float64
	RelativeToProxy               float64
	RelativeToRuntime             float64
	PositiveCorrectionRequired    bool
	BridgeMatchingCorrectionClaim bool
	Verdict                       string
}

type EffectiveCLambdaCorrection struct {
	BOverA2        float64
	CProxy         float64
	CNeededMZ      float64
	DeltaC         float64
	RelativeDeltaC float64
	Statement      string
	Verdict        string
}

type HiggsProxyGapDiagnostic struct {
	VRuntimeGeV                   float64
	LambdaProxyMZ                 float64
	LambdaRuntimeMZ               float64
	MassProxyGeV                  float64
	MassRuntimeGeV                float64
	DeltaMassRuntimeMinusProxyGeV float64
	RelativeMassGap               float64
	ClaimsHiggsDerivation         bool
	ClaimsPoleMassTheorem         bool
	Statement                     string
	Verdict                       string
}

type TypedSourceCandidate struct {
	Name                  string
	CanHavePositiveSign   bool
	RequiresObservedInput bool
	NativeCertified       bool
	Comment               string
	Verdict               string
}

type SignAudit struct {
	ProxyBelowRuntime           bool
	PositiveDeltaLambda         bool
	PositiveSourcesExistAsSlots bool
	NativePositiveCorrection    bool
	Statement                   string
	Verdict                     string
}

type NeutrinoTraceCompletionAudit struct {
	VisibleV1IncludesNeutrinoYukawa bool
	CouldShiftAAndB                 bool
	ValuesInserted                  bool
	NativeCompletionTheorem         bool
	Statement                       string
	Verdict                         string
}

type RuntimeTransportChain struct {
	LambdaProxyMZ              float64
	DeltaLambdaMatch           float64
	LambdaRuntimeMZ            float64
	LambdaRuntimeL12           float64
	ChainClosedByNativeTheorem bool
	Statement                  string
	Verdict                    string
}

type StressSealImpact struct {
	ImprovesScalarLaneArchitecture  bool
	StressStillUsesLambdaRuntimeL12 bool
	CanReplaceStressLambdaWithProxy bool
	LambdaProxyL12                  float64
	LambdaRuntimeL12                float64
	Statement                       string
	Verdict                         string
}

type NativeStatus struct {
	NativeCThreeEighthsTheorem    bool
	NativeMatchingCorrection      bool
	NativePoleMSbarConversion     bool
	NativeNeutrinoTraceCompletion bool
	NativeHiggsPoleTheorem        bool
	NativeProxyRuntimeTheorem     bool
	Statement                     string
	Verdict                       string
}

type Firewalls struct {
	ClaimsHiggsMass           bool
	ClaimsHiggsStability      bool
	ClaimsLambdaZeroBoundary  bool
	ClaimsGaugeUnification    bool
	ClaimsNativeScalarTheorem bool
	ClaimsNativeMatching      bool
	Verdict                   string
}

type Analysis struct {
	Inherited        Inherited
	MatchingGap      MatchingGapTable
	EffectiveCLambda EffectiveCLambdaCorrection
	HiggsProxyGap    HiggsProxyGapDiagnostic
	SourceCandidates []TypedSourceCandidate
	Sign             SignAudit
	NeutrinoTrace    NeutrinoTraceCompletionAudit
	RuntimeChain     RuntimeTransportChain
	StressImpact     StressSealImpact
	NativeStatus     NativeStatus
	Firewalls        Firewalls
	Truth            string
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
	g620, err := generation2ba2onethirdrigidityspectralquarticproxyaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate620 predecessor: %w", err)
	}
	match := buildMatchingGap()
	a := Analysis{
		Inherited:        inherit(g620),
		MatchingGap:      match,
		EffectiveCLambda: buildEffectiveCLambda(),
		HiggsProxyGap:    buildHiggsProxyGap(),
		SourceCandidates: buildSourceCandidates(),
		Sign:             buildSignAudit(match),
		NeutrinoTrace:    buildNeutrinoTraceAudit(),
		RuntimeChain:     buildRuntimeChain(match),
		StressImpact:     buildStressImpact(),
		NativeStatus:     buildNativeStatus(),
		Firewalls:        auditFirewalls(),
		Truth:            "Gate 621 audits the first scalar arrow lambda_proxy -> lambda_runtime(M_Z). The gap is positive and small enough to be a bridge-layer matching/convention/loop-correction slot, while the high-scale negative lambda_runtime(Lambda_12) remains an RG-transported continuum result. The scalar chain is lambda_proxy + Delta_lambda_match -> lambda_runtime(M_Z) -> RG transport -> lambda_runtime(Lambda_12); no Higgs mass, pole theorem, or native proxy-to-runtime matching theorem is derived.",
	}
	return a, nil
}

func inherit(a generation2ba2onethirdrigidityspectralquarticproxyaudit.Analysis) Inherited {
	return Inherited{
		LambdaProxyMZ:          a.ProxyRows[0].LambdaProxy,
		LambdaRuntimeMZ:        a.ProxyRows[0].LambdaRuntime,
		LambdaProxyLambda12:    a.ProxyRows[1].LambdaProxy,
		LambdaRuntimeLambda12:  a.ProxyRows[1].LambdaRuntime,
		BA2MZ:                  a.ProxyRows[0].BOverA2,
		BA2Lambda12:            a.ProxyRows[1].BOverA2,
		PreviousSeparation:     a.Separation.Verdict,
		PreviousStressVerdict:  a.StressImpact.Verdict,
		PreviousProxyClose:     a.ProxyRows[0].Verdict,
		PreviousHighScaleBlock: a.ProxyRows[1].Verdict,
		Verdict:                StatusGate620Inherited,
	}
}

func buildMatchingGap() MatchingGapTable {
	d := lambdaMZ - proxyMZ
	return MatchingGapTable{
		Scale:                         mZLabel,
		LambdaProxy:                   proxyMZ,
		LambdaRuntime:                 lambdaMZ,
		DeltaLambdaMatch:              d,
		RelativeToProxy:               d / proxyMZ,
		RelativeToRuntime:             d / lambdaMZ,
		PositiveCorrectionRequired:    d > 0,
		BridgeMatchingCorrectionClaim: false,
		Verdict:                       StatusMatchingGapComputed,
	}
}

func buildEffectiveCLambda() EffectiveCLambdaCorrection {
	cNeeded := lambdaMZ / ba2MZ
	deltaC := cNeeded - cProxy
	return EffectiveCLambdaCorrection{
		BOverA2:        ba2MZ,
		CProxy:         cProxy,
		CNeededMZ:      cNeeded,
		DeltaC:         deltaC,
		RelativeDeltaC: deltaC / cProxy,
		Statement:      "The low-scale runtime quartic would require c_lambda slightly above 3/8 if one forced lambda_runtime(M_Z)=c_lambda*b/a^2. This is an effective bridge diagnostic, not a certified convention coefficient.",
		Verdict:        StatusEffectiveCLambdaComputed,
	}
}

func buildHiggsProxyGap() HiggsProxyGapDiagnostic {
	massProxy := math.Sqrt(2*proxyMZ) * vRuntime
	massRuntime := math.Sqrt(2*lambdaMZ) * vRuntime
	delta := massRuntime - massProxy
	return HiggsProxyGapDiagnostic{
		VRuntimeGeV:                   vRuntime,
		LambdaProxyMZ:                 proxyMZ,
		LambdaRuntimeMZ:               lambdaMZ,
		MassProxyGeV:                  massProxy,
		MassRuntimeGeV:                massRuntime,
		DeltaMassRuntimeMinusProxyGeV: delta,
		RelativeMassGap:               delta / massProxy,
		ClaimsHiggsDerivation:         false,
		ClaimsPoleMassTheorem:         false,
		Statement:                     "m_H_proxy and m_H_runtime are diagnostic tree-level translations through the same v ledger. They are not a Higgs pole-mass prediction or native scalar theorem.",
		Verdict:                       StatusHiggsProxyGapComputed,
	}
}

func buildSourceCandidates() []TypedSourceCandidate {
	return []TypedSourceCandidate{
		{Name: "pole/MSbar matching", CanHavePositiveSign: true, RequiresObservedInput: true, NativeCertified: false, Comment: "Lawful scalar matching slot, but no native theorem or calculation is certified here.", Verdict: StatusNoMatchingTheorem},
		{Name: "one-loop scalar threshold correction", CanHavePositiveSign: true, RequiresObservedInput: true, NativeCertified: false, Comment: "Sign-compatible as a bridge slot; no threshold spectrum is provided.", Verdict: StatusNoMatchingTheorem},
		{Name: "top-loop correction", CanHavePositiveSign: true, RequiresObservedInput: true, NativeCertified: false, Comment: "Typed loop source candidate; no native top-loop matching theorem is supplied.", Verdict: StatusNoMatchingTheorem},
		{Name: "gauge-loop correction", CanHavePositiveSign: true, RequiresObservedInput: true, NativeCertified: false, Comment: "Typed loop source candidate; no certified scalar matching calculation is present.", Verdict: StatusNoMatchingTheorem},
		{Name: "two-loop RG improvement", CanHavePositiveSign: true, RequiresObservedInput: true, NativeCertified: false, Comment: "May shift the runtime ledger, but it is not a low-scale native proxy theorem.", Verdict: StatusNoMatchingTheorem},
		{Name: "scalar field normalization convention", CanHavePositiveSign: true, RequiresObservedInput: false, NativeCertified: false, Comment: "Could move effective c_lambda; Gate 619 did not certify the convention coefficient.", Verdict: StatusNoThreeEighthsTheorem},
		{Name: "missing neutrino/Yukawa trace contribution", CanHavePositiveSign: true, RequiresObservedInput: true, NativeCertified: false, Comment: "Can alter a,b,b/a^2 if supplied; no values are inserted in v1.", Verdict: StatusNoNeutrinoTraceCompletion},
		{Name: "spectral-action convention c_lambda", CanHavePositiveSign: true, RequiresObservedInput: false, NativeCertified: false, Comment: "Equivalent to adjusting c_lambda; still convention-sealed.", Verdict: StatusNoThreeEighthsTheorem},
	}
}

func buildSignAudit(match MatchingGapTable) SignAudit {
	return SignAudit{
		ProxyBelowRuntime:           proxyMZ < lambdaMZ,
		PositiveDeltaLambda:         match.DeltaLambdaMatch > 0,
		PositiveSourcesExistAsSlots: true,
		NativePositiveCorrection:    false,
		Statement:                   "lambda_proxy(M_Z) is below lambda_runtime(M_Z), so any low-scale proxy-to-runtime bridge requires a positive scalar matching/convention correction.",
		Verdict:                     StatusPositiveMatchingRequired,
	}
}

func buildNeutrinoTraceAudit() NeutrinoTraceCompletionAudit {
	return NeutrinoTraceCompletionAudit{
		VisibleV1IncludesNeutrinoYukawa: false,
		CouldShiftAAndB:                 true,
		ValuesInserted:                  false,
		NativeCompletionTheorem:         false,
		Statement:                       "The visible v1 a,b ledgers omit a concrete neutrino Yukawa or Weinberg-proxy contribution. Such data could shift b/a^2, but no neutrino values or native trace-completion theorem are supplied.",
		Verdict:                         StatusNoNeutrinoTraceCompletion,
	}
}

func buildRuntimeChain(match MatchingGapTable) RuntimeTransportChain {
	return RuntimeTransportChain{
		LambdaProxyMZ:              proxyMZ,
		DeltaLambdaMatch:           match.DeltaLambdaMatch,
		LambdaRuntimeMZ:            lambdaMZ,
		LambdaRuntimeL12:           lambdaL12,
		ChainClosedByNativeTheorem: false,
		Statement:                  "Scalar chain: lambda_proxy + Delta_lambda_match -> lambda_runtime(M_Z) -> RG transport -> lambda_runtime(Lambda_12). This separates the positive proxy from the negative high-scale runtime quartic.",
		Verdict:                    StatusProxyRuntimeChainDefined,
	}
}

func buildStressImpact() StressSealImpact {
	return StressSealImpact{
		ImprovesScalarLaneArchitecture:  true,
		StressStillUsesLambdaRuntimeL12: true,
		CanReplaceStressLambdaWithProxy: false,
		LambdaProxyL12:                  proxyL12,
		LambdaRuntimeL12:                lambdaL12,
		Statement:                       "The matching-gap audit improves the scalar architecture, but the GaugeScalarBoundaryStressSeal still pairs R_3-1 with lambda_runtime(Lambda_12), not lambda_proxy.",
		Verdict:                         StatusProxyRuntimeChainDefined,
	}
}

func buildNativeStatus() NativeStatus {
	return NativeStatus{
		NativeCThreeEighthsTheorem:    false,
		NativeMatchingCorrection:      false,
		NativePoleMSbarConversion:     false,
		NativeNeutrinoTraceCompletion: false,
		NativeHiggsPoleTheorem:        false,
		NativeProxyRuntimeTheorem:     false,
		Statement:                     "No native theorem currently proves c_lambda=3/8, the scalar matching correction, pole/MSbar conversion, neutrino trace completion, Higgs pole mass, or proxy-to-runtime equality.",
		Verdict:                       StatusNoProxyRuntimeTheorem,
	}
}

func auditFirewalls() Firewalls {
	return Firewalls{Verdict: StatusGate621Boundary}
}

func Statuses() []string {
	return []string{
		StatusGate620Inherited,
		StatusMatchingGapComputed,
		StatusEffectiveCLambdaComputed,
		StatusHiggsProxyGapComputed,
		StatusProxyCloseAtMZ,
		StatusPositiveMatchingRequired,
		StatusProxyRuntimeChainDefined,
		StatusNoMatchingTheorem,
		StatusNoThreeEighthsTheorem,
		StatusNoHiggsPoleTheorem,
		StatusNoProxyRuntimeTheorem,
		StatusNoNeutrinoTraceCompletion,
		StatusGate621Boundary,
	}
}
