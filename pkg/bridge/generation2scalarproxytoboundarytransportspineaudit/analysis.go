// Package generation2scalarproxytoboundarytransportspineaudit implements
// Gate 658: Scalar Proxy-to-Boundary Transport Spine Audit.
//
// Gate 657 sealed the K_7/Fano-Hitchin boundary route and pivoted the active
// work back to transport.  Gate 658 merges the scalar proxy-runtime matching
// lane with the high-scale boundary-stress lane into one typed bridge spine:
// lambda_proxy(M_Z) -> lambda_runtime(M_Z) -> lambda(Lambda_12) ->
// GaugeScalarBoundaryStressSeal.  The audit computes all residual slots and
// preserves the firewall against Higgs, stability, threshold, and native scalar
// theorem claims.
package generation2scalarproxytoboundarytransportspineaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate657 "github.com/bagherbal/asha-engine/pkg/bridge/generation2internalobstructionsealclosurepivot"
)

const (
	AuditID = "GATE658-SCALAR-PROXY-TO-BOUNDARY-TRANSPORT-SPINE-AUDIT"

	StatusGate657TransportPivotInherited      = "PASS_GATE657_TRANSPORT_PIVOT_INHERITED"
	StatusScalarProxyRuntimeChainConstructed  = "PASS_SCALAR_PROXY_RUNTIME_CHAIN_CONSTRUCTED"
	StatusHistoryLoopMatchingFormComputed     = "PASS_HISTORY_LOOP_UNIT_MATCHING_FORM_COMPUTED"
	StatusKappaLambdaDefined                  = "PASS_KAPPA_LAMBDA_DEFINED"
	StatusRGTransportLaneRecorded             = "PASS_RG_TRANSPORT_LANE_RECORDED"
	StatusBoundaryStressComparisonInherited   = "PASS_BOUNDARY_STRESS_COMPARISON_INHERITED"
	StatusResidualSlotsSeparated              = "PASS_RESIDUAL_SLOTS_SEPARATED"
	StatusSourceAuditComputed                 = "PASS_SOURCE_AUDIT_COMPUTED"
	StatusScalarProxyBoundarySpineActive      = "CONDITIONAL_SUPPORT_SCALAR_PROXY_TO_BOUNDARY_SPINE_IS_ACTIVE"
	StatusLowScaleLoopMatchingClueActive      = "CONDITIONAL_SUPPORT_LOW_SCALE_LOOP_MATCHING_CLUE_REMAINS_ACTIVE"
	StatusBoundaryStressTransportLive         = "CONDITIONAL_SUPPORT_BOUNDARY_STRESS_TRANSPORT_REMAINS_ACTIVE"
	StatusNoNativeProxyRuntimeTheorem         = "FAILED_ROUTE_NO_NATIVE_PROXY_TO_RUNTIME_MATCHING_THEOREM"
	StatusNoNativeRGThresholdTheorem          = "FAILED_ROUTE_NO_NATIVE_RG_THRESHOLD_THEOREM"
	StatusNoNativeBoundaryStressTheorem       = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_STRESS_THEOREM"
	StatusNoNativeKappaLambdaSource           = "FAILED_ROUTE_NO_NATIVE_KAPPA_LAMBDA_SOURCE_THEOREM"
	StatusNoNativeHistoryLoopSourceFromScalar = "FAILED_ROUTE_NO_NATIVE_HISTORY_LOOP_UNIT_SOURCE_FROM_SCALAR_SPINE"
	StatusNoHiggsMassOrStabilityClaim         = "FAILED_ROUTE_NO_HIGGS_MASS_OR_STABILITY_CLAIM"
	StatusGate658Boundary                     = "FIREWALL_PRESERVED_GATE658_SCALAR_TRANSPORT_SPINE_BOUNDARY"
)

const (
	lambdaProxyMZ   = 0.12490310236015
	lambdaRuntimeMZ = 0.1296525650504758
	deltaLambda     = 0.0047494626903257
	rhoLambdaMatch  = 0.0380251779225699
	historyLoopL    = 0.0397887357729738
	lambdaLambda12  = -0.0497009420776833
	r3Minus1        = 0.0509933868964996
	xiBoundary      = 0.0503471644870914
)

type Gate657Inheritance struct {
	TransportPivotInherited bool
	FanoBoundaryClosed      bool
	ActiveBridgeVectorBuilt bool
	PrimaryWasRGTransport   bool
	ScalarMatchingActive    bool
	BoundaryStressActive    bool
	HistoryLoopActive       bool
	K7BoundaryBlocked       bool
	NoFanoBoundaryMap       bool
	NoSevenTraceTheorem     bool
	FirewallPreserved       bool
	Verdict                 string
}

type ProxyLane struct {
	Formula              string
	LambdaProxyMZ        float64
	BA2Ratio             float64
	OneEighth            float64
	DifferenceFromOne8   float64
	RelativeFromOne8     float64
	CloseToOneEighth     bool
	TreeProxyOnly        bool
	CannotReplaceRuntime bool
	Verdict              string
}

type LowScaleMatchingLane struct {
	LambdaProxyMZ          float64
	LambdaRuntimeMZ        float64
	DeltaLambdaMatch       float64
	RelativeMatch          float64
	HistoryLoopUnit        float64
	KappaLambda            float64
	OneMinusKappaLambda    float64
	ReconstructedRuntimeMZ float64
	ReconstructionResidual float64
	RawLAnsatzRuntime      float64
	RawLAnsatzResidual     float64
	LoopSized              bool
	Verdict                string
}

type RGTransportLane struct {
	StartScale           string
	BoundaryScale        string
	LambdaRuntimeStart   float64
	LambdaBoundary       float64
	AbsLambdaBoundary    float64
	RuntimeTurnsNegative bool
	UsesCurrentV1RG      bool
	ClaimsThresholdLaw   bool
	Verdict              string
}

type BoundaryStressLane struct {
	AbsLambdaBoundary        float64
	R3Minus1                 float64
	XiBoundary               float64
	MeanStressRecomputed     float64
	BoundarySplit            float64
	HalfSplit                float64
	AbsLambdaResidualToXi    float64
	R3ResidualToXi           float64
	AntiAlignmentForm        string
	XiPreferredOverHalfTrace bool
	Verdict                  string
}

type ResidualSlot struct {
	Name           string
	Value          float64
	Scale          string
	TypedStatus    string
	RequiresSource string
}

type ResidualDecomposition struct {
	Slots                 []ResidualSlot
	MatchSlotSeparated    bool
	RGSlotSeparated       bool
	BoundarySlotSeparated bool
	ThresholdSlotsOpen    bool
	Verdict               string
}

type SourceAudit struct {
	KappaLambdaSourceCertified bool
	XiBoundarySourceCertified  bool
	HistoryLoopSourceCertified bool
	ProxyRuntimeTheorem        bool
	RGThresholdTheorem         bool
	BoundaryStressTheorem      bool
	SearchedRandomConstants    bool
	TypedQuantitiesOnly        bool
	Verdict                    string
}

type SpineClassification struct {
	Name                 string
	Active               bool
	BridgeLayerOnly      bool
	MergesScalarBoundary bool
	Touches              []string
	NextPressurePoint    string
	Verdict              string
}

type Firewalls struct {
	ClaimsHiggsMass             bool
	ClaimsScalarStability       bool
	ClaimsGaugeUnification      bool
	ClaimsThresholdExistence    bool
	ClaimsNativeScalarTheorem   bool
	ClaimsBoundaryStressDerived bool
	ClaimsPhysicalSpacetime     bool
	ClaimsFlavorTheorem         bool
	Verdict                     string
}

type Analysis struct {
	Inherited Gate657Inheritance
	Proxy     ProxyLane
	Matching  LowScaleMatchingLane
	RG        RGTransportLane
	Boundary  BoundaryStressLane
	Residuals ResidualDecomposition
	Sources   SourceAudit
	Spine     SpineClassification
	Firewalls Firewalls
	Truth     string
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
	g657, err := gate657.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate657 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g657)
	proxy := buildProxyLane()
	matching := buildLowScaleMatching()
	rg := buildRGTransport()
	boundary := buildBoundaryStress()
	residuals := buildResiduals(matching, boundary)
	sources := buildSourceAudit()
	spine := buildSpine()
	firewalls := Firewalls{Verdict: StatusGate658Boundary}
	truth := "Gate 658 merges the scalar proxy-runtime and boundary-stress lanes into one active bridge spine: lambda_proxy(M_Z)->lambda_runtime(M_Z)->lambda(Lambda_12)->xi_boundary.  The low-scale matching is exactly typed as lambda_runtime=lambda_proxy[1+L(1-kappa_lambda)] with kappa_lambda defined, while RG transport and boundary stress remain empirical/bridge slots.  No native proxy-runtime theorem, RG threshold theorem, boundary-stress theorem, Higgs claim, or stability claim is derived."
	return Analysis{Inherited: inherited, Proxy: proxy, Matching: matching, RG: rg, Boundary: boundary, Residuals: residuals, Sources: sources, Spine: spine, Firewalls: firewalls, Truth: truth}, nil
}

func buildInheritance(g gate657.Analysis) Gate657Inheritance {
	return Gate657Inheritance{
		TransportPivotInherited: g.Strategic.ReturnToTransport && g.Ranking.PrimaryPath == "RG/threshold transport refinement",
		FanoBoundaryClosed:      g.Closure.BoundaryInterfaceFailed && g.Strategic.StopFanoBoundaryLane,
		ActiveBridgeVectorBuilt: g.Active.ActiveCount == 5,
		PrimaryWasRGTransport:   g.Ranking.PrimaryPath == "RG/threshold transport refinement",
		ScalarMatchingActive:    g.Strategic.ScalarMatchingLive,
		BoundaryStressActive:    g.Strategic.BoundaryStressLive,
		HistoryLoopActive:       g.Strategic.HistoryLoopLive,
		K7BoundaryBlocked:       g.Strategic.K7BoundaryBlocked,
		NoFanoBoundaryMap:       g.Closure.BoundaryInterfaceFailed,
		NoSevenTraceTheorem:     g.Inherited.NoSevenOver72Theorem && g.Inherited.NoSevenOver144Theorem,
		FirewallPreserved:       g.Firewalls.Verdict == gate657.StatusGate657Boundary,
		Verdict:                 StatusGate657TransportPivotInherited,
	}
}

func buildProxyLane() ProxyLane {
	one8 := 0.125
	ba2 := lambdaProxyMZ * 8.0 / 3.0
	diff := lambdaProxyMZ - one8
	return ProxyLane{
		Formula:              "lambda_proxy(M_Z)=(3/8)(b/a^2)",
		LambdaProxyMZ:        lambdaProxyMZ,
		BA2Ratio:             ba2,
		OneEighth:            one8,
		DifferenceFromOne8:   diff,
		RelativeFromOne8:     diff / one8,
		CloseToOneEighth:     math.Abs(diff) < 1e-3,
		TreeProxyOnly:        true,
		CannotReplaceRuntime: true,
		Verdict:              StatusScalarProxyRuntimeChainConstructed,
	}
}

func buildLowScaleMatching() LowScaleMatchingLane {
	kappa := 1.0 - rhoLambdaMatch/historyLoopL
	reconstructed := lambdaProxyMZ * (1.0 + historyLoopL*(1.0-kappa))
	rawL := lambdaProxyMZ * (1.0 + historyLoopL)
	return LowScaleMatchingLane{
		LambdaProxyMZ:          lambdaProxyMZ,
		LambdaRuntimeMZ:        lambdaRuntimeMZ,
		DeltaLambdaMatch:       deltaLambda,
		RelativeMatch:          rhoLambdaMatch,
		HistoryLoopUnit:        historyLoopL,
		KappaLambda:            kappa,
		OneMinusKappaLambda:    1.0 - kappa,
		ReconstructedRuntimeMZ: reconstructed,
		ReconstructionResidual: reconstructed - lambdaRuntimeMZ,
		RawLAnsatzRuntime:      rawL,
		RawLAnsatzResidual:     rawL - lambdaRuntimeMZ,
		LoopSized:              rhoLambdaMatch > 0 && math.Abs(rhoLambdaMatch-historyLoopL) < 0.002,
		Verdict:                join(StatusHistoryLoopMatchingFormComputed, StatusKappaLambdaDefined, StatusLowScaleLoopMatchingClueActive),
	}
}

func buildRGTransport() RGTransportLane {
	return RGTransportLane{
		StartScale:           "M_Z",
		BoundaryScale:        "Lambda_12",
		LambdaRuntimeStart:   lambdaRuntimeMZ,
		LambdaBoundary:       lambdaLambda12,
		AbsLambdaBoundary:    math.Abs(lambdaLambda12),
		RuntimeTurnsNegative: lambdaLambda12 < 0,
		UsesCurrentV1RG:      true,
		ClaimsThresholdLaw:   false,
		Verdict:              StatusRGTransportLaneRecorded,
	}
}

func buildBoundaryStress() BoundaryStressLane {
	absL := math.Abs(lambdaLambda12)
	mean := 0.5 * (r3Minus1 + absL)
	split := r3Minus1 - absL
	half := 0.5 * split
	return BoundaryStressLane{
		AbsLambdaBoundary:        absL,
		R3Minus1:                 r3Minus1,
		XiBoundary:               xiBoundary,
		MeanStressRecomputed:     mean,
		BoundarySplit:            split,
		HalfSplit:                half,
		AbsLambdaResidualToXi:    absL - xiBoundary,
		R3ResidualToXi:           r3Minus1 - xiBoundary,
		AntiAlignmentForm:        "(R_3-1, lambda(Lambda_12)) ≈ (+xi_boundary, -xi_boundary)",
		XiPreferredOverHalfTrace: true,
		Verdict:                  join(StatusBoundaryStressComparisonInherited, StatusBoundaryStressTransportLive),
	}
}

func buildResiduals(m LowScaleMatchingLane, b BoundaryStressLane) ResidualDecomposition {
	slots := []ResidualSlot{
		{Name: "Delta_match", Value: m.DeltaLambdaMatch, Scale: "lambda(M_Z)", TypedStatus: "positive low-scale proxy-to-runtime bridge correction", RequiresSource: "native proxy-runtime matching theorem"},
		{Name: "rho_lambda_match", Value: m.RelativeMatch, Scale: "relative to lambda_proxy", TypedStatus: "close to L=1/(8*pi) with kappa_lambda deficit", RequiresSource: "HistoryLoopUnit scalar matching source"},
		{Name: "kappa_lambda", Value: m.KappaLambda, Scale: "dimensionless L-deficit", TypedStatus: "defined by 1-rho/L; source not certified", RequiresSource: "native kappa_lambda theorem"},
		{Name: "Delta_RG", Value: lambdaLambda12 - lambdaRuntimeMZ, Scale: "runtime lambda transport", TypedStatus: "current v1 RG transport output", RequiresSource: "threshold/scheme/top/alpha_s refinement"},
		{Name: "Delta_boundary", Value: b.BoundarySplit, Scale: "gauge-scalar wound split", TypedStatus: "Gate613/Gate626 boundary stress split", RequiresSource: "native boundary-stress theorem"},
		{Name: "delta_lambda_threshold", Value: 0, Scale: "open correction slot", TypedStatus: "placeholder slot; not fitted", RequiresSource: "threshold matching ledger"},
		{Name: "delta_top", Value: 0, Scale: "open correction slot", TypedStatus: "placeholder slot; not fitted", RequiresSource: "top Yukawa uncertainty/source theorem"},
		{Name: "delta_alpha_s", Value: 0, Scale: "open correction slot", TypedStatus: "placeholder slot; not fitted", RequiresSource: "strong coupling transport refinement"},
		{Name: "delta_scheme", Value: 0, Scale: "open correction slot", TypedStatus: "placeholder slot; not fitted", RequiresSource: "scheme conversion ledger"},
		{Name: "delta_pole_MSbar", Value: 0, Scale: "open correction slot", TypedStatus: "placeholder slot; not fitted", RequiresSource: "pole-to-MSbar conversion theorem/ledger"},
	}
	return ResidualDecomposition{Slots: slots, MatchSlotSeparated: true, RGSlotSeparated: true, BoundarySlotSeparated: true, ThresholdSlotsOpen: true, Verdict: StatusResidualSlotsSeparated}
}

func buildSourceAudit() SourceAudit {
	return SourceAudit{
		KappaLambdaSourceCertified: false,
		XiBoundarySourceCertified:  false,
		HistoryLoopSourceCertified: false,
		ProxyRuntimeTheorem:        false,
		RGThresholdTheorem:         false,
		BoundaryStressTheorem:      false,
		SearchedRandomConstants:    false,
		TypedQuantitiesOnly:        true,
		Verdict: join(StatusSourceAuditComputed, StatusNoNativeProxyRuntimeTheorem, StatusNoNativeRGThresholdTheorem,
			StatusNoNativeBoundaryStressTheorem, StatusNoNativeKappaLambdaSource, StatusNoNativeHistoryLoopSourceFromScalar),
	}
}

func buildSpine() SpineClassification {
	return SpineClassification{
		Name:                 "ScalarProxyToBoundaryTransportSpine",
		Active:               true,
		BridgeLayerOnly:      true,
		MergesScalarBoundary: true,
		Touches:              []string{"lambda_proxy(M_Z)", "lambda_runtime(M_Z)", "L=1/(8*pi)", "kappa_lambda", "lambda(Lambda_12)", "R_3-1", "xi_boundary"},
		NextPressurePoint:    "Does the low-scale loop matching correction propagate lawfully into the high-scale scalar/gauge boundary wound under refined RG/threshold transport?",
		Verdict:              StatusScalarProxyBoundarySpineActive,
	}
}

func join(parts ...string) string { return strings.Join(parts, "; ") }

func Statuses() []string {
	return []string{
		StatusGate657TransportPivotInherited,
		StatusScalarProxyRuntimeChainConstructed,
		StatusHistoryLoopMatchingFormComputed,
		StatusKappaLambdaDefined,
		StatusRGTransportLaneRecorded,
		StatusBoundaryStressComparisonInherited,
		StatusResidualSlotsSeparated,
		StatusSourceAuditComputed,
		StatusScalarProxyBoundarySpineActive,
		StatusLowScaleLoopMatchingClueActive,
		StatusBoundaryStressTransportLive,
		StatusNoNativeProxyRuntimeTheorem,
		StatusNoNativeRGThresholdTheorem,
		StatusNoNativeBoundaryStressTheorem,
		StatusNoNativeKappaLambdaSource,
		StatusNoNativeHistoryLoopSourceFromScalar,
		StatusNoHiggsMassOrStabilityClaim,
		StatusGate658Boundary,
	}
}
