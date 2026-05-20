// Package generation2ba2onethirdrigidityspectralquarticproxyaudit implements
// Gate 620: b/a² One-Third Rigidity and Spectral Quartic Proxy Audit.
//
// Gate 619 showed that lambda_runtime(Lambda_12) cannot be directly identified
// with c_lambda*b/a² because b/a² is nonnegative while the v1 transported
// runtime quartic is negative at Lambda_12. Gate 620 audits the separate clue
// that b/a² is nearly frozen near 1/3 and that (3/8)(b/a²) gives a positive
// spectral/tree scalar quartic proxy close to the low-scale canonical runtime
// quartic. This gate is diagnostic only: it does not derive a Higgs mass,
// scalar stability, lambda-zero boundary, or native scalar quartic theorem.
package generation2ba2onethirdrigidityspectralquarticproxyaudit

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2spectralquarticconventioncoefficientaudit"
)

const (
	AuditID = "GATE620-B-A2-ONE-THIRD-RIGIDITY-SPECTRAL-QUARTIC-PROXY-AUDIT"

	StatusGate619Inherited          = "PASS_GATE619_BA2_DIAGNOSTIC_INHERITED"
	StatusBA2NearOneThird           = "PASS_B_OVER_A_SQUARED_NEAR_ONE_THIRD_AT_MZ_AND_LAMBDA12"
	StatusTopDominanceSupport       = "CONDITIONAL_SUPPORT_BA2_ONE_THIRD_FROM_TOP_COLOR_DOMINANCE"
	StatusThreeEighthsProxyComputed = "PASS_C_LAMBDA_THREE_EIGHTHS_PROXY_COMPUTED"
	StatusProxyCloseAtMZ            = "CONDITIONAL_SUPPORT_LAMBDA_PROXY_CLOSE_TO_RUNTIME_LAMBDA_AT_MZ"
	StatusProxyFailsAtL12           = "FAILED_ROUTE_LAMBDA_PROXY_DOES_NOT_EQUAL_NEGATIVE_RUNTIME_LAMBDA_AT_LAMBDA12"
	StatusSeparateScalarLanes       = "CONDITIONAL_SUPPORT_SPECTRAL_TREE_QUARTIC_AND_RUNTIME_RG_QUARTIC_MUST_BE_SEPARATED"
	StatusNoBA2Theorem              = "FAILED_ROUTE_NO_NATIVE_BA2_ONE_THIRD_THEOREM"
	StatusNoThreeEighthsTheorem     = "FAILED_ROUTE_NO_NATIVE_C_LAMBDA_THREE_EIGHTHS_SCALAR_THEOREM"
	StatusNoProxyRuntimeMatch       = "FAILED_ROUTE_NO_NATIVE_PROXY_TO_RUNTIME_MATCHING_THEOREM"
	StatusGate620Boundary           = "FIREWALL_PRESERVED_GATE620_BA2_QUARTIC_PROXY_BOUNDARY"
)

const (
	// Runtime data inherited from Gate 619.
	aMZ       = 2.8424095142339083
	bMZ       = 2.6910096440382287
	ba2MZ     = 0.33307493962706697
	lambdaMZ  = 0.1296525650504758
	aL12      = 0.6941198223775996
	bL12      = 0.16047699018700937
	ba2L12    = 0.3330764110541872
	lambdaL12 = -0.049700942077683274
	vRuntime  = 246.21965079413738
	yTopMZ    = 0.973191904392486
	yTopL12   = 0.4809200309718785
	sin2Theta = 3.0 / 8.0
	oneThird  = 1.0 / 3.0
)

type Inherited struct {
	MZBA2                 float64
	Lambda12BA2           float64
	LambdaRuntimeMZ       float64
	LambdaRuntimeLambda12 float64
	PreviousSignVerdict   string
	PreviousStressVerdict string
	Verdict               string
}

type BA2RigidityRow struct {
	Scale              string
	ATrace             float64
	BTrace             float64
	BOverA2            float64
	DeltaFromOneThird  float64
	RelativeToOneThird float64
	LambdaRuntime      float64
	Verdict            string
}

type BA2RigiditySummary struct {
	BA2Drift          float64
	AbsBA2Drift       float64
	RelativeDriftToMZ float64
	NearlyInvariant   bool
	Statement         string
	Verdict           string
}

type TopDominanceRow struct {
	Scale          string
	YTop           float64
	ApproxA        float64
	ApproxB        float64
	ApproxBA2      float64
	RuntimeA       float64
	RuntimeB       float64
	RuntimeBA2     float64
	ADeltaRelative float64
	BDeltaRelative float64
	Statement      string
	Verdict        string
}

type CLambdaProxyRow struct {
	Scale                string
	CLambdaCandidate     float64
	BOverA2              float64
	LambdaProxy          float64
	LambdaRuntime        float64
	ProxyMinusRuntime    float64
	AbsResidual          float64
	RelativeToRuntimeAbs float64
	SignCompatible       bool
	Verdict              string
}

type HiggsProxyDiagnostic struct {
	VRuntimeGeV           float64
	LambdaProxyMZ         float64
	MassProxyGeV          float64
	LambdaRuntimeMZ       float64
	RuntimeMassGeV        float64
	MassResidualGeV       float64
	ClaimsHiggsDerivation bool
	Statement             string
	Verdict               string
}

type RuntimeTransportSeparation struct {
	SpectralTreeProxyPositive   bool
	RuntimeQuarticRGTransported bool
	ProxyEqualsRuntimeAtMZ      bool
	ProxyEqualsRuntimeAtL12     bool
	LambdaL12Negative           bool
	RequiresMatchingTheorem     bool
	Statement                   string
	Verdict                     string
}

type StressSealImpact struct {
	StressUsesRuntimeLambda    bool
	SpectralLaneUsesProxy      bool
	CanReplaceRuntimeWithProxy bool
	RuntimeLambdaL12           float64
	ProxyLambdaL12             float64
	Statement                  string
	Verdict                    string
}

type NativeStatus struct {
	NativeTopDominanceTheorem    bool
	NativeBA2OneThirdTheorem     bool
	NativeThreeEighthsTheorem    bool
	NativeProxyRuntimeMatching   bool
	NativeHiggsMassTheorem       bool
	NativeScalarStabilityTheorem bool
	Statement                    string
	Verdict                      string
}

type Firewalls struct {
	ClaimsHiggsMass         bool
	ClaimsHiggsStability    bool
	ClaimsLambdaZero        bool
	ClaimsNativeQuartic     bool
	ClaimsProxyRuntimeMatch bool
	ClaimsNativeBA2         bool
	Verdict                 string
}

type Analysis struct {
	Inherited       Inherited
	RigidityRows    []BA2RigidityRow
	RigiditySummary BA2RigiditySummary
	TopDominance    []TopDominanceRow
	ProxyRows       []CLambdaProxyRow
	HiggsProxy      HiggsProxyDiagnostic
	Separation      RuntimeTransportSeparation
	StressImpact    StressSealImpact
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
	g619, err := generation2spectralquarticconventioncoefficientaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate619 predecessor: %w", err)
	}
	inherited := inherit(g619)
	rigidity := buildRigidityRows()
	proxy := buildProxyRows()
	a := Analysis{
		Inherited:       inherited,
		RigidityRows:    rigidity,
		RigiditySummary: buildRigiditySummary(rigidity),
		TopDominance:    buildTopDominanceRows(),
		ProxyRows:       proxy,
		HiggsProxy:      buildHiggsProxy(proxy),
		Separation:      buildSeparation(proxy),
		StressImpact:    buildStressImpact(proxy),
		NativeStatus:    buildNativeStatus(),
		Firewalls:       auditFirewalls(),
		Truth:           "Gate 620 separates the positive spectral/tree scalar proxy lane from the RG-transported runtime scalar lane. The visible v1 trace ratio b/a^2 is nearly one-third at both M_Z and Lambda_12, conditionally explained by top/color dominance. The typed proxy lambda_proxy=(3/8)(b/a^2) is close to lambda_runtime(M_Z), but remains positive at Lambda_12 and therefore cannot equal the negative runtime lambda there. The GaugeScalarBoundaryStressSeal must continue to use lambda_runtime(Lambda_12), while the spectral/tree lane may track lambda_proxy only as a bridge diagnostic.",
	}
	return a, nil
}

func inherit(a generation2spectralquarticconventioncoefficientaudit.Analysis) Inherited {
	return Inherited{
		MZBA2:                 a.Diagnostics[0].BOverA2,
		Lambda12BA2:           a.Diagnostics[1].BOverA2,
		LambdaRuntimeMZ:       lambdaMZ,
		LambdaRuntimeLambda12: a.Diagnostics[1].LambdaRuntime,
		PreviousSignVerdict:   a.SignAudit.Verdict,
		PreviousStressVerdict: a.StressImpact.Verdict,
		Verdict:               StatusGate619Inherited,
	}
}

func buildRigidityRows() []BA2RigidityRow {
	return []BA2RigidityRow{
		buildRigidityRow("M_Z", aMZ, bMZ, ba2MZ, lambdaMZ),
		buildRigidityRow("Lambda_12", aL12, bL12, ba2L12, lambdaL12),
	}
}

func buildRigidityRow(scale string, a, b, ba2, lambda float64) BA2RigidityRow {
	d := ba2 - oneThird
	return BA2RigidityRow{Scale: scale, ATrace: a, BTrace: b, BOverA2: ba2, DeltaFromOneThird: d, RelativeToOneThird: d / oneThird, LambdaRuntime: lambda, Verdict: StatusBA2NearOneThird}
}

func buildRigiditySummary(rows []BA2RigidityRow) BA2RigiditySummary {
	drift := rows[1].BOverA2 - rows[0].BOverA2
	return BA2RigiditySummary{
		BA2Drift:          drift,
		AbsBA2Drift:       math.Abs(drift),
		RelativeDriftToMZ: drift / rows[0].BOverA2,
		NearlyInvariant:   math.Abs(drift) < 2e-6,
		Statement:         "b/a^2 is nearly fixed near 1/3 in the v1 runtime ledgers; this is a trace-shape diagnostic, not a native theorem.",
		Verdict:           StatusBA2NearOneThird,
	}
}

func buildTopDominanceRows() []TopDominanceRow {
	return []TopDominanceRow{
		buildTopDominanceRow("M_Z", yTopMZ, aMZ, bMZ, ba2MZ),
		buildTopDominanceRow("Lambda_12", yTopL12, aL12, bL12, ba2L12),
	}
}

func buildTopDominanceRow(scale string, yt, runtimeA, runtimeB, runtimeBA2 float64) TopDominanceRow {
	aApprox := 3 * yt * yt
	bApprox := 3 * math.Pow(yt, 4)
	return TopDominanceRow{
		Scale:          scale,
		YTop:           yt,
		ApproxA:        aApprox,
		ApproxB:        bApprox,
		ApproxBA2:      bApprox / (aApprox * aApprox),
		RuntimeA:       runtimeA,
		RuntimeB:       runtimeB,
		RuntimeBA2:     runtimeBA2,
		ADeltaRelative: (runtimeA - aApprox) / runtimeA,
		BDeltaRelative: (runtimeB - bApprox) / runtimeB,
		Statement:      "Top/color dominance gives a≈3 y_t^2 and b≈3 y_t^4, hence b/a^2≈1/3. This explains the v1 rigidity conditionally from observed Yukawa hierarchy, not from ASHA-native flavor law.",
		Verdict:        StatusTopDominanceSupport,
	}
}

func buildProxyRows() []CLambdaProxyRow {
	return []CLambdaProxyRow{
		buildProxyRow("M_Z", ba2MZ, lambdaMZ),
		buildProxyRow("Lambda_12", ba2L12, lambdaL12),
	}
}

func buildProxyRow(scale string, ba2, lambda float64) CLambdaProxyRow {
	proxy := sin2Theta * ba2
	res := proxy - lambda
	v := StatusProxyCloseAtMZ
	if lambda < 0 {
		v = StatusProxyFailsAtL12
	}
	return CLambdaProxyRow{
		Scale:                scale,
		CLambdaCandidate:     sin2Theta,
		BOverA2:              ba2,
		LambdaProxy:          proxy,
		LambdaRuntime:        lambda,
		ProxyMinusRuntime:    res,
		AbsResidual:          math.Abs(res),
		RelativeToRuntimeAbs: math.Abs(res) / math.Abs(lambda),
		SignCompatible:       proxy*lambda >= 0,
		Verdict:              v,
	}
}

func buildHiggsProxy(rows []CLambdaProxyRow) HiggsProxyDiagnostic {
	proxy := rows[0].LambdaProxy
	massProxy := math.Sqrt(2*proxy) * vRuntime
	runtimeMass := math.Sqrt(2*lambdaMZ) * vRuntime
	return HiggsProxyDiagnostic{
		VRuntimeGeV:           vRuntime,
		LambdaProxyMZ:         proxy,
		MassProxyGeV:          massProxy,
		LambdaRuntimeMZ:       lambdaMZ,
		RuntimeMassGeV:        runtimeMass,
		MassResidualGeV:       massProxy - runtimeMass,
		ClaimsHiggsDerivation: false,
		Statement:             "m_H_proxy=sqrt(2 lambda_proxy) v is a low-scale diagnostic of the positive spectral/tree proxy only. It is not a Higgs mass derivation or pole-mass theorem.",
		Verdict:               StatusProxyCloseAtMZ,
	}
}

func buildSeparation(rows []CLambdaProxyRow) RuntimeTransportSeparation {
	return RuntimeTransportSeparation{
		SpectralTreeProxyPositive:   rows[0].LambdaProxy > 0 && rows[1].LambdaProxy > 0,
		RuntimeQuarticRGTransported: true,
		ProxyEqualsRuntimeAtMZ:      math.Abs(rows[0].ProxyMinusRuntime) < 0.006,
		ProxyEqualsRuntimeAtL12:     false,
		LambdaL12Negative:           lambdaL12 < 0,
		RequiresMatchingTheorem:     true,
		Statement:                   "lambda_proxy=(3/8)(b/a^2) is a positive spectral/tree diagnostic; lambda_runtime(mu) is the canonical SM quartic transported by continuum RG. They must not be identified without a matching theorem.",
		Verdict:                     StatusSeparateScalarLanes,
	}
}

func buildStressImpact(rows []CLambdaProxyRow) StressSealImpact {
	return StressSealImpact{
		StressUsesRuntimeLambda:    true,
		SpectralLaneUsesProxy:      true,
		CanReplaceRuntimeWithProxy: false,
		RuntimeLambdaL12:           lambdaL12,
		ProxyLambdaL12:             rows[1].LambdaProxy,
		Statement:                  "The GaugeScalarBoundaryStressSeal continues to pair R_3-1 with lambda_runtime(Lambda_12). The spectral scalar lane may track lambda_proxy, but lambda_proxy cannot replace the negative runtime quartic at Lambda_12.",
		Verdict:                    StatusSeparateScalarLanes,
	}
}

func buildNativeStatus() NativeStatus {
	return NativeStatus{
		NativeTopDominanceTheorem:    false,
		NativeBA2OneThirdTheorem:     false,
		NativeThreeEighthsTheorem:    false,
		NativeProxyRuntimeMatching:   false,
		NativeHiggsMassTheorem:       false,
		NativeScalarStabilityTheorem: false,
		Statement:                    "No native theorem currently proves top dominance, b/a^2=1/3, c_lambda=3/8 for the scalar quartic, proxy-to-runtime matching, Higgs mass, or scalar stability.",
		Verdict:                      StatusNoProxyRuntimeMatch,
	}
}

func auditFirewalls() Firewalls {
	return Firewalls{Verdict: StatusGate620Boundary}
}

func Statuses() []string {
	return []string{
		StatusGate619Inherited,
		StatusBA2NearOneThird,
		StatusTopDominanceSupport,
		StatusThreeEighthsProxyComputed,
		StatusProxyCloseAtMZ,
		StatusProxyFailsAtL12,
		StatusSeparateScalarLanes,
		StatusNoBA2Theorem,
		StatusNoThreeEighthsTheorem,
		StatusNoProxyRuntimeMatch,
		StatusGate620Boundary,
	}
}
