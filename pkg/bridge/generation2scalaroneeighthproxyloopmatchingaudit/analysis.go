// Package generation2scalaroneeighthproxyloopmatchingaudit implements
// Gate 622: Scalar One-Eighth Proxy and Loop-Matching Correction Audit.
//
// Gate 621 typed the low-scale scalar matching gap between the positive
// spectral/tree proxy lambda_proxy=(3/8)(b/a^2) and the runtime canonical
// Standard Model quartic lambda_runtime(M_Z). Gate 622 audits whether this
// gap is loop-sized, especially relative to 1/(8*pi), while preserving all
// Higgs, scalar-stability, and native-theorem firewalls.
package generation2scalaroneeighthproxyloopmatchingaudit

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2scalarproxyruntimematchinggapaudit"
)

const (
	AuditID = "GATE622-SCALAR-ONE-EIGHTH-PROXY-LOOP-MATCHING-CORRECTION-AUDIT"

	StatusGate621Inherited         = "PASS_GATE621_MATCHING_GAP_INHERITED"
	StatusOneEighthProxyAudited    = "PASS_ONE_EIGHTH_PROXY_AUDITED"
	StatusRelativeLoopComputed     = "PASS_RELATIVE_LOOP_CORRECTION_COMPUTED"
	StatusAbsoluteLoopComputed     = "PASS_ABSOLUTE_LOOP_CORRECTION_COMPUTED"
	StatusLoopSized                = "CONDITIONAL_SUPPORT_MATCHING_GAP_IS_LOOP_SIZED"
	StatusOneOver8PiClose          = "CONDITIONAL_SUPPORT_ONE_OVER_8PI_CLOSE_TO_RELATIVE_MATCHING_GAP"
	StatusOneOver64PiClose         = "CONDITIONAL_SUPPORT_ONE_OVER_64PI_CLOSE_TO_ABSOLUTE_MATCHING_GAP"
	StatusPositiveLoopRequired     = "CONDITIONAL_SUPPORT_POSITIVE_LOOP_MATCHING_CORRECTION_REQUIRED"
	StatusRefinedLoopProxyComputed = "PASS_REFINED_LOOP_PROXY_DIAGNOSTIC_COMPUTED"
	StatusRuntimeChainDefined      = "CONDITIONAL_SUPPORT_PROXY_LOOP_MATCHING_RUNTIME_CHAIN_DEFINED"
	StatusNoLoopMatchingTheorem    = "FAILED_ROUTE_NO_NATIVE_LOOP_MATCHING_THEOREM"
	StatusNoOneEighthScalarTheorem = "FAILED_ROUTE_NO_NATIVE_ONE_EIGHTH_SCALAR_THEOREM"
	StatusNoHiggsPoleTheorem       = "FAILED_ROUTE_NO_NATIVE_HIGGS_POLE_THEOREM"
	StatusNoProxyRuntimeTheorem    = "FAILED_ROUTE_NO_NATIVE_PROXY_TO_RUNTIME_MATCHING_THEOREM"
	StatusNoBA2OneThirdTheorem     = "FAILED_ROUTE_NO_NATIVE_BA2_ONE_THIRD_THEOREM"
	StatusNoThreeEighthsTheorem    = "FAILED_ROUTE_NO_NATIVE_C_LAMBDA_THREE_EIGHTHS_THEOREM"
	StatusGate622Boundary          = "FIREWALL_PRESERVED_GATE622_SCALAR_LOOP_MATCHING_BOUNDARY"
)

const (
	ba2MZ         = 0.33307493962706697
	lambdaProxy   = (3.0 / 8.0) * ba2MZ
	lambdaRuntime = 0.1296525650504758
	vRuntime      = 246.21965079413738
	ytMZ          = 0.973191904392486
	g2MZ          = 0.6527521238927322
	gYMZ          = 0.3500756885970262
)

type Inherited struct {
	LambdaProxyMZ        float64
	LambdaRuntimeMZ      float64
	DeltaLambdaMatch     float64
	RelativeToProxy      float64
	RelativeToRuntime    float64
	BA2MZ                float64
	MProxyGeV            float64
	MRuntimeGeV          float64
	PreviousGapVerdict   string
	PreviousChainVerdict string
	Verdict              string
}

type OneEighthProxyAudit struct {
	LambdaProxy              float64
	LambdaRuntime            float64
	OneEighth                float64
	ProxyMinusOneEighth      float64
	RuntimeMinusOneEighth    float64
	ProxyRelativeDeviation   float64
	RuntimeRelativeDeviation float64
	ProxyFromBA2             string
	NativeClaim              bool
	Verdict                  string
}

type LoopCorrectionCandidate struct {
	Name             string
	Value            float64
	Residual         float64
	RelativeResidual float64
	Typed            bool
	NativeCertified  bool
	Comment          string
}

type RelativeLoopCorrectionTable struct {
	RhoLambdaMatch  float64
	Candidates      []LoopCorrectionCandidate
	ClosestName     string
	ClosestResidual float64
	Verdict         string
}

type AbsoluteLoopCorrectionTable struct {
	DeltaLambdaMatch float64
	Candidates       []LoopCorrectionCandidate
	ClosestName      string
	ClosestResidual  float64
	Verdict          string
}

type SignAudit struct {
	LambdaProxyBelowRuntime bool
	PositiveCorrection      bool
	SignCompatibleSlots     []string
	NativeSourceCertified   bool
	Statement               string
	Verdict                 string
}

type HiggsProxyRefinementDiagnostic struct {
	LoopUnitOneOver8Pi        float64
	LambdaProxy               float64
	LambdaAnsatz              float64
	LambdaRuntime             float64
	AnsatzMinusRuntime        float64
	RelativeAnsatzResidual    float64
	MassProxyGeV              float64
	MassAnsatzGeV             float64
	MassRuntimeGeV            float64
	DeltaMassAnsatzRuntimeGeV float64
	ClaimsHiggsPrediction     bool
	Statement                 string
	Verdict                   string
}

type RuntimeTransportChain struct {
	LambdaProxy         float64
	DeltaLambdaLoopSlot float64
	LambdaRuntimeMZ     float64
	LambdaRuntimeL12    float64
	ChainNative         bool
	Statement           string
	Verdict             string
}

type NativeStatus struct {
	NativeBA2OneThirdTheorem       bool
	NativeCThreeEighthsTheorem     bool
	NativeOneOver8PiScalarMatching bool
	NativeProxyRuntimeTheorem      bool
	NativeHiggsPoleTheorem         bool
	Statement                      string
	Verdict                        string
}

type Firewalls struct {
	ClaimsHiggsMass           bool
	ClaimsHiggsStability      bool
	ClaimsLambdaZeroBoundary  bool
	ClaimsGaugeUnification    bool
	ClaimsNativeScalarTheorem bool
	Verdict                   string
}

type Analysis struct {
	Inherited       Inherited
	OneEighth       OneEighthProxyAudit
	RelativeLoops   RelativeLoopCorrectionTable
	AbsoluteLoops   AbsoluteLoopCorrectionTable
	Sign            SignAudit
	HiggsDiagnostic HiggsProxyRefinementDiagnostic
	RuntimeChain    RuntimeTransportChain
	NativeStatus    NativeStatus
	Firewalls       Firewalls
	Truth           string
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
	g621, err := generation2scalarproxyruntimematchinggapaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate621 predecessor: %w", err)
	}
	inherited := inherit(g621)
	rel := buildRelativeLoopTable(inherited.RelativeToProxy)
	abs := buildAbsoluteLoopTable(inherited.DeltaLambdaMatch)
	a := Analysis{
		Inherited:       inherited,
		OneEighth:       buildOneEighthAudit(),
		RelativeLoops:   rel,
		AbsoluteLoops:   abs,
		Sign:            buildSignAudit(),
		HiggsDiagnostic: buildHiggsDiagnostic(),
		RuntimeChain:    buildRuntimeChain(inherited),
		NativeStatus:    buildNativeStatus(),
		Firewalls:       auditFirewalls(),
		Truth:           "Gate 622 audits the scalar chain lambda_proxy≈1/8, plus a positive loop-sized matching correction, into lambda_runtime(M_Z). The relative matching gap is close to 1/(8*pi), and the absolute gap is close to 1/(64*pi), but no native loop matching theorem, one-eighth scalar theorem, Higgs pole theorem, or proxy-to-runtime theorem is derived.",
	}
	return a, nil
}

func inherit(a generation2scalarproxyruntimematchinggapaudit.Analysis) Inherited {
	return Inherited{
		LambdaProxyMZ:        a.MatchingGap.LambdaProxy,
		LambdaRuntimeMZ:      a.MatchingGap.LambdaRuntime,
		DeltaLambdaMatch:     a.MatchingGap.DeltaLambdaMatch,
		RelativeToProxy:      a.MatchingGap.RelativeToProxy,
		RelativeToRuntime:    a.MatchingGap.RelativeToRuntime,
		BA2MZ:                a.EffectiveCLambda.BOverA2,
		MProxyGeV:            a.HiggsProxyGap.MassProxyGeV,
		MRuntimeGeV:          a.HiggsProxyGap.MassRuntimeGeV,
		PreviousGapVerdict:   a.MatchingGap.Verdict,
		PreviousChainVerdict: a.RuntimeChain.Verdict,
		Verdict:              StatusGate621Inherited,
	}
}

func buildOneEighthAudit() OneEighthProxyAudit {
	one := 1.0 / 8.0
	return OneEighthProxyAudit{
		LambdaProxy:              lambdaProxy,
		LambdaRuntime:            lambdaRuntime,
		OneEighth:                one,
		ProxyMinusOneEighth:      lambdaProxy - one,
		RuntimeMinusOneEighth:    lambdaRuntime - one,
		ProxyRelativeDeviation:   (lambdaProxy - one) / one,
		RuntimeRelativeDeviation: (lambdaRuntime - one) / one,
		ProxyFromBA2:             "lambda_proxy=(3/8)(b/a^2), with b/a^2≈1/3, so lambda_proxy≈1/8.",
		NativeClaim:              false,
		Verdict:                  StatusOneEighthProxyAudited,
	}
}

func buildRelativeLoopTable(rho float64) RelativeLoopCorrectionTable {
	alpha2 := g2MZ * g2MZ / (4 * math.Pi)
	e := g2MZ * gYMZ / math.Sqrt(g2MZ*g2MZ+gYMZ*gYMZ)
	alphaEM := e * e / (4 * math.Pi)
	yt2 := ytMZ * ytMZ
	candidates := []LoopCorrectionCandidate{
		loopCandidate("1/(8*pi)", 1/(8*math.Pi), rho, "typed loop angular unit; strongest scalar matching clue"),
		loopCandidate("1/(16*pi)", 1/(16*math.Pi), rho, "typed loop unit but too small"),
		loopCandidate("1/(4*pi)", 1/(4*math.Pi), rho, "typed loop unit but too large"),
		loopCandidate("alpha_2(M_Z)", alpha2, rho, "weak coupling loop-normalized endpoint quantity"),
		loopCandidate("alpha_EM(M_Z)", alphaEM, rho, "electromagnetic endpoint coupling; too small for direct relative gap"),
		loopCandidate("y_t(M_Z)^2/(16*pi^2)", yt2/(16*math.Pi*math.Pi), rho, "top one-loop factor, single color/spin coefficient omitted"),
		loopCandidate("3*y_t(M_Z)^2/(16*pi^2)", 3*yt2/(16*math.Pi*math.Pi), rho, "top one-loop factor with color-3 multiplier"),
		loopCandidate("6*y_t(M_Z)^2/(16*pi^2)", 6*yt2/(16*math.Pi*math.Pi), rho, "top one-loop factor with sixfold coefficient; numerically close but coefficient theorem missing"),
	}
	name, res := closest(candidates)
	return RelativeLoopCorrectionTable{RhoLambdaMatch: rho, Candidates: candidates, ClosestName: name, ClosestResidual: res, Verdict: StatusRelativeLoopComputed}
}

func buildAbsoluteLoopTable(delta float64) AbsoluteLoopCorrectionTable {
	candidates := []LoopCorrectionCandidate{
		loopCandidate("1/(64*pi)", 1/(64*math.Pi), delta, "absolute scalar gap expected from (1/8)*(1/(8*pi))"),
		loopCandidate("lambda_proxy/(8*pi)", lambdaProxy/(8*math.Pi), delta, "same typed loop correction using actual proxy rather than exact 1/8"),
		loopCandidate("lambda_proxy*alpha_2(M_Z)", lambdaProxy*(g2MZ*g2MZ/(4*math.Pi)), delta, "weak-coupling multiplicative correction"),
		loopCandidate("lambda_proxy*6*y_t(M_Z)^2/(16*pi^2)", lambdaProxy*(6*ytMZ*ytMZ/(16*math.Pi*math.Pi)), delta, "top-loop multiplicative correction with sixfold coefficient; close but not certified"),
	}
	name, res := closest(candidates)
	return AbsoluteLoopCorrectionTable{DeltaLambdaMatch: delta, Candidates: candidates, ClosestName: name, ClosestResidual: res, Verdict: StatusAbsoluteLoopComputed}
}

func loopCandidate(name string, value, target float64, comment string) LoopCorrectionCandidate {
	residual := value - target
	rr := math.NaN()
	if target != 0 {
		rr = residual / target
	}
	return LoopCorrectionCandidate{Name: name, Value: value, Residual: residual, RelativeResidual: rr, Typed: true, NativeCertified: false, Comment: comment}
}

func closest(rows []LoopCorrectionCandidate) (string, float64) {
	if len(rows) == 0 {
		return "", math.NaN()
	}
	best := rows[0]
	for _, r := range rows[1:] {
		if math.Abs(r.Residual) < math.Abs(best.Residual) {
			best = r
		}
	}
	return best.Name, best.Residual
}

func buildSignAudit() SignAudit {
	return SignAudit{
		LambdaProxyBelowRuntime: lambdaProxy < lambdaRuntime,
		PositiveCorrection:      lambdaRuntime-lambdaProxy > 0,
		SignCompatibleSlots:     []string{"pole/MSbar matching", "one-loop scalar threshold", "top/gauge/scalar loop correction", "two-loop improvement", "field-normalization convention"},
		NativeSourceCertified:   false,
		Statement:               "The proxy lies below the runtime quartic at M_Z; the bridge correction must be positive. Several typed loop/matching slots can have this sign, but none is certified natively.",
		Verdict:                 StatusPositiveLoopRequired,
	}
}

func buildHiggsDiagnostic() HiggsProxyRefinementDiagnostic {
	loop := 1 / (8 * math.Pi)
	ansatz := lambdaProxy * (1 + loop)
	massProxy := math.Sqrt(2*lambdaProxy) * vRuntime
	massAnsatz := math.Sqrt(2*ansatz) * vRuntime
	massRuntime := math.Sqrt(2*lambdaRuntime) * vRuntime
	return HiggsProxyRefinementDiagnostic{
		LoopUnitOneOver8Pi:        loop,
		LambdaProxy:               lambdaProxy,
		LambdaAnsatz:              ansatz,
		LambdaRuntime:             lambdaRuntime,
		AnsatzMinusRuntime:        ansatz - lambdaRuntime,
		RelativeAnsatzResidual:    (ansatz - lambdaRuntime) / lambdaRuntime,
		MassProxyGeV:              massProxy,
		MassAnsatzGeV:             massAnsatz,
		MassRuntimeGeV:            massRuntime,
		DeltaMassAnsatzRuntimeGeV: massAnsatz - massRuntime,
		ClaimsHiggsPrediction:     false,
		Statement:                 "lambda_proxy*(1+1/(8*pi)) is a loop-sized bridge diagnostic and remains distinct from a Higgs pole-mass theorem.",
		Verdict:                   StatusRefinedLoopProxyComputed,
	}
}

func buildRuntimeChain(inh Inherited) RuntimeTransportChain {
	return RuntimeTransportChain{
		LambdaProxy:         inh.LambdaProxyMZ,
		DeltaLambdaLoopSlot: inh.DeltaLambdaMatch,
		LambdaRuntimeMZ:     inh.LambdaRuntimeMZ,
		LambdaRuntimeL12:    -0.049700942077683274,
		ChainNative:         false,
		Statement:           "Scalar chain: lambda_proxy≈1/8 plus a loop-sized positive matching correction gives lambda_runtime(M_Z), which is then RG-transported to negative lambda_runtime(Lambda_12).",
		Verdict:             StatusRuntimeChainDefined,
	}
}

func buildNativeStatus() NativeStatus {
	return NativeStatus{
		NativeBA2OneThirdTheorem:       false,
		NativeCThreeEighthsTheorem:     false,
		NativeOneOver8PiScalarMatching: false,
		NativeProxyRuntimeTheorem:      false,
		NativeHiggsPoleTheorem:         false,
		Statement:                      "No native theorem currently proves b/a^2=1/3, c_lambda=3/8, a 1/(8*pi) scalar matching correction, proxy-to-runtime equality, or a Higgs pole formula.",
		Verdict:                        StatusNoProxyRuntimeTheorem,
	}
}

func auditFirewalls() Firewalls { return Firewalls{Verdict: StatusGate622Boundary} }

func Statuses() []string {
	return []string{
		StatusGate621Inherited,
		StatusOneEighthProxyAudited,
		StatusRelativeLoopComputed,
		StatusAbsoluteLoopComputed,
		StatusLoopSized,
		StatusOneOver8PiClose,
		StatusOneOver64PiClose,
		StatusPositiveLoopRequired,
		StatusRefinedLoopProxyComputed,
		StatusRuntimeChainDefined,
		StatusNoLoopMatchingTheorem,
		StatusNoOneEighthScalarTheorem,
		StatusNoHiggsPoleTheorem,
		StatusNoProxyRuntimeTheorem,
		StatusNoBA2OneThirdTheorem,
		StatusNoThreeEighthsTheorem,
		StatusGate622Boundary,
	}
}
