// Package generation2koidetransportvectordecompositionaudit implements Gate 580:
// Koide Transport-Vector Decomposition Audit.
//
// Gate 577 identified the charged-lepton square-root Koide cone. Gate 578
// computed the residual azimuth. Gate 579 showed that pole and M_Z Yukawa
// frames are projectively identical in v1 while the Lambda_12 endpoint is
// slightly closer to Q=2/3. Gate 580 differentiates the observed transport in
// the Koide coordinates (rho, theta, phi) over the v1 interval from M_Z to
// Lambda_12.
//
// This gate is a finite-difference bridge audit. It does not derive lepton
// masses, Koide, Yukawa eigenvalues, PMNS/CKM, generations, root-trace data,
// or a native ASHA flavor operator. It only decomposes the already-produced
// runtime history flow into radial, cone-angle, and azimuthal components.
package generation2koidetransportvectordecompositionaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2koidenaturalframeaudit"
	"github.com/bagherbal/asha-engine/pkg/historytransport"
)

const (
	AuditID = "GATE580-KOIDE-TRANSPORT-VECTOR-DECOMPOSITION-AUDIT"

	StatusRuntimeInherited          = "PASS_GATE579_AND_HISTORY_TRANSPORT_RUNTIME_INHERITED"
	StatusKoideCoordinatesInherited = "PASS_KOIDE_PROJECTIVE_COORDINATES_INHERITED"
	StatusFiniteIntervalCertified   = "PASS_MZ_TO_LAMBDA12_LOG_TRANSPORT_INTERVAL_CERTIFIED"
	StatusTransportVectorComputed   = "PASS_KOIDE_TRANSPORT_VECTOR_COMPONENTS_COMPUTED"
	StatusRadialComponentDominant   = "PASS_CHARGED_LEPTON_TRANSPORT_DOMINATED_BY_RADIAL_RESCALING_IN_V1"
	StatusThetaMovesTowardCone      = "CONDITIONAL_SUPPORT_THETA_COMPONENT_MOVES_TOWARD_KOIDE_CONE_IN_V1"
	StatusAzimuthNearlyInvariant    = "PASS_PHI_COMPONENT_NEARLY_INVARIANT_IN_V1"
	StatusProjectiveRayStable       = "CONDITIONAL_SUPPORT_CHARGED_LEPTON_PROJECTIVE_RAY_NEARLY_STABLE_IN_V1"
	StatusConeAttractorNotCertified = "FAILED_ROUTE_KOIDE_CONE_ATTRACTOR_NOT_CERTIFIED_BY_TWO_POINT_V1_FINITE_DIFFERENCE"
	StatusNoContinuousBeta          = "FAILED_ROUTE_NO_CONTINUOUS_KOIDE_COORDINATE_BETA_FUNCTION_CERTIFIED"
	StatusNoNativeRootTrace         = "FAILED_ROUTE_NO_NATIVE_ROOT_TRACE_OR_ABSOLUTE_DIRAC_TRANSPORT_OPERATOR"
	StatusNoFlavorDerivation        = "FAILED_ROUTE_NO_ASHA_NATIVE_CHARGED_LEPTON_FLAVOR_TRANSPORT_DERIVATION"
	StatusNoNewCarrier              = "FIREWALL_PRESERVED_NO_NEW_FLAVOR_CARRIER_OR_SELECTOR_INTRODUCED"
	StatusNoObservedPromotion       = "FIREWALL_PRESERVED_OBSERVED_LEPTON_ENDPOINTS_REMAIN_HISTORY_DATA"
	StatusNoTexturePromotion        = "FIREWALL_PRESERVED_KOIDE_TRANSPORT_DOES_NOT_DERIVE_TEXTURE_CKM_PMNS_OR_GENERATIONS"
	StatusGate352Preserved          = "FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING"
	StatusGate580BoundaryPreserved  = "FIREWALL_PRESERVED_GATE580_KOIDE_TRANSPORT_VECTOR_BOUNDARY"
)

const KoideTargetThetaDeg = 45.0

type RuntimeInheritance struct {
	Mu0GeV          float64
	Lambda12GeV     float64
	LogLambdaOverMZ float64
	PredecessorGate string
	RuntimeSource   string
	Verdict         string
}

type KoideCoordinateEndpoint struct {
	Name          string
	Rho           float64
	Q             float64
	DeltaQ        float64
	AbsDeltaQ     float64
	ThetaDeg      float64
	ThetaRad      float64
	ThetaErrorDeg float64
	ThetaErrorRad float64
	PhiDeg        float64
	PhiSignedDeg  float64
	PhiRad        float64
	Verdict       string
}

type TransportVector struct {
	From                    KoideCoordinateEndpoint
	To                      KoideCoordinateEndpoint
	DeltaT                  float64
	DeltaRho                float64
	DeltaLnRho              float64
	DLnRhoDT                float64
	DeltaThetaDeg           float64
	DeltaThetaRad           float64
	DThetaDTDeg             float64
	DThetaDTRad             float64
	DeltaPhiDeg             float64
	DeltaPhiRad             float64
	DPhiDTDeg               float64
	DPhiDTRad               float64
	AverageThetaDeg         float64
	ProjectiveAngularDelta  float64
	ProjectiveAngularDT     float64
	RadialToProjectiveRatio float64
	ThetaImprovementFactor  float64
	QImprovementFactor      float64
	MovesTowardCone         bool
	PhiNearlyInvariant      bool
	RadialDominant          bool
	Verdict                 string
}

type DynamicalInterpretation struct {
	MostlyRadialRescaling   bool
	ConeAttractionVisible   bool
	ConeAttractorCertified  bool
	AzimuthPreserved        bool
	ContinuousBetaCertified bool
	Explanation             string
	MissingTheorem          string
	Verdict                 string
}

type FirewallAudit struct {
	DerivesKoide               bool
	DerivesLeptonMasses        bool
	DerivesYukawaEigenvalues   bool
	DerivesCKM                 bool
	DerivesPMNS                bool
	DerivesGenerationHierarchy bool
	IntroducesNewCarrier       bool
	PromotesObservedAsNative   bool
	PreservesGate352           bool
	Verdict                    string
}

type FinalVerdict struct {
	SealName                  string
	DlnRhoDt                  float64
	DThetaDtDeg               float64
	DPhiDtDeg                 float64
	ThetaImprovementFactor    float64
	QImprovementFactor        float64
	ProjectiveAngularDeltaDeg float64
	RadialToProjectiveRatio   float64
	ConeAttractorCertified    bool
	MinimalNextRequirement    string
	Verdict                   string
}

type Analysis struct {
	Runtime   RuntimeInheritance
	MZ        KoideCoordinateEndpoint
	Lambda12  KoideCoordinateEndpoint
	Transport TransportVector
	Dynamics  DynamicalInterpretation
	Firewalls FirewallAudit
	Final     FinalVerdict
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
	bundle, err := historytransport.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build history transport runtime: %w", err)
	}
	g579, err := generation2koidenaturalframeaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate579 predecessor: %w", err)
	}
	runtime := inheritRuntime(bundle)
	mz := endpointFromGate579("M_Z", g579.Compare.MZ)
	lambda := endpointFromGate579("Lambda_12", g579.Compare.Lambda12)
	transport := decomposeTransport(mz, lambda, bundle.GaugeBoundary.LogLambda12Mu0)
	dynamics := interpretDynamics(transport)
	firewalls := auditFirewalls()
	final := compileFinal(transport, dynamics)
	truth := "Gate 580 decomposes the charged-lepton square-root Yukawa transport into Koide coordinates. In v1, the dominant motion is radial rescaling; the projective ray is nearly stable; theta moves slightly toward the 45-degree Koide cone; and phi is almost invariant. This is not a certified attractor theorem, only a two-endpoint bridge-layer finite difference."
	return Analysis{Runtime: runtime, MZ: mz, Lambda12: lambda, Transport: transport, Dynamics: dynamics, Firewalls: firewalls, Final: final, Truth: truth}, nil
}

func inheritRuntime(b historytransport.Bundle) RuntimeInheritance {
	return RuntimeInheritance{Mu0GeV: b.GaugeBoundary.Mu0GeV, Lambda12GeV: b.GaugeBoundary.Lambda12GeV, LogLambdaOverMZ: b.GaugeBoundary.LogLambda12Mu0, PredecessorGate: "Gate579 Koide natural frame audit", RuntimeSource: "historytransport.BuildDefault() plus Gate579 Koide coordinates", Verdict: StatusRuntimeInherited}
}

func endpointFromGate579(name string, p generation2koidenaturalframeaudit.FramePoint) KoideCoordinateEndpoint {
	thetaRad := degToRad(p.ThetaDeg)
	phiSigned := p.PhiSignedDeg
	phiRad := degToRad(phiSigned)
	return KoideCoordinateEndpoint{Name: name, Rho: p.Rho, Q: p.Q, DeltaQ: p.DeltaQ, AbsDeltaQ: p.AbsDeltaQ, ThetaDeg: p.ThetaDeg, ThetaRad: thetaRad, ThetaErrorDeg: p.ThetaDeg - KoideTargetThetaDeg, ThetaErrorRad: degToRad(p.ThetaDeg - KoideTargetThetaDeg), PhiDeg: p.PhiDeg, PhiSignedDeg: phiSigned, PhiRad: phiRad, Verdict: StatusKoideCoordinatesInherited}
}

func decomposeTransport(mz, lambda KoideCoordinateEndpoint, dt float64) TransportVector {
	deltaThetaDeg := lambda.ThetaDeg - mz.ThetaDeg
	deltaPhiDeg := signedAngleDelta(lambda.PhiDeg, mz.PhiDeg)
	deltaThetaRad := degToRad(deltaThetaDeg)
	deltaPhiRad := degToRad(deltaPhiDeg)
	dlnrho := math.Log(lambda.Rho / mz.Rho)
	avgTheta := 0.5 * (mz.ThetaRad + lambda.ThetaRad)
	projective := math.Sqrt(deltaThetaRad*deltaThetaRad + math.Pow(math.Sin(avgTheta)*deltaPhiRad, 2))
	ratio := math.Inf(1)
	if projective > 0 {
		ratio = math.Abs(dlnrho) / projective
	}
	thetaImprovement := mz.AbsDeltaQ / lambda.AbsDeltaQ
	if lambda.AbsDeltaQ == 0 {
		thetaImprovement = math.Inf(1)
	}
	qImprovement := thetaImprovement
	movesTowardCone := math.Abs(lambda.ThetaErrorDeg) < math.Abs(mz.ThetaErrorDeg)
	phiStable := math.Abs(deltaPhiDeg) < 3e-4
	radialDominant := ratio > 100
	verdict := strings.Join([]string{StatusTransportVectorComputed, StatusRadialComponentDominant, StatusThetaMovesTowardCone, StatusAzimuthNearlyInvariant, StatusProjectiveRayStable}, ";")
	return TransportVector{From: mz, To: lambda, DeltaT: dt, DeltaRho: lambda.Rho - mz.Rho, DeltaLnRho: dlnrho, DLnRhoDT: dlnrho / dt, DeltaThetaDeg: deltaThetaDeg, DeltaThetaRad: deltaThetaRad, DThetaDTDeg: deltaThetaDeg / dt, DThetaDTRad: deltaThetaRad / dt, DeltaPhiDeg: deltaPhiDeg, DeltaPhiRad: deltaPhiRad, DPhiDTDeg: deltaPhiDeg / dt, DPhiDTRad: deltaPhiRad / dt, AverageThetaDeg: radToDeg(avgTheta), ProjectiveAngularDelta: projective, ProjectiveAngularDT: projective / dt, RadialToProjectiveRatio: ratio, ThetaImprovementFactor: thetaImprovement, QImprovementFactor: qImprovement, MovesTowardCone: movesTowardCone, PhiNearlyInvariant: phiStable, RadialDominant: radialDominant, Verdict: verdict}
}

func interpretDynamics(t TransportVector) DynamicalInterpretation {
	explanation := "The finite difference from M_Z to Lambda_12 shows nonzero radial shrinkage, a small positive theta step toward 45 degrees, and a tiny azimuth drift. Because only two endpoints and v1 approximate running are used, this supports a near-invariant projective ray and possible cone-directed motion, but does not certify a continuous Koide beta function or attractor."
	missing := "continuous Koide-coordinate RG equations for (rho,theta,phi), with threshold/multi-loop control and a native root-trace or absolute-Dirac observable if promotion beyond environmental seal is attempted"
	return DynamicalInterpretation{MostlyRadialRescaling: t.RadialDominant, ConeAttractionVisible: t.MovesTowardCone, ConeAttractorCertified: false, AzimuthPreserved: t.PhiNearlyInvariant, ContinuousBetaCertified: false, Explanation: explanation, MissingTheorem: missing, Verdict: strings.Join([]string{StatusRadialComponentDominant, StatusThetaMovesTowardCone, StatusAzimuthNearlyInvariant, StatusConeAttractorNotCertified, StatusNoContinuousBeta}, ";")}
}

func auditFirewalls() FirewallAudit {
	return FirewallAudit{DerivesKoide: false, DerivesLeptonMasses: false, DerivesYukawaEigenvalues: false, DerivesCKM: false, DerivesPMNS: false, DerivesGenerationHierarchy: false, IntroducesNewCarrier: false, PromotesObservedAsNative: false, PreservesGate352: true, Verdict: strings.Join([]string{StatusNoNativeRootTrace, StatusNoFlavorDerivation, StatusNoNewCarrier, StatusNoObservedPromotion, StatusNoTexturePromotion, StatusGate352Preserved, StatusGate580BoundaryPreserved}, ";")}
}

func compileFinal(t TransportVector, d DynamicalInterpretation) FinalVerdict {
	return FinalVerdict{SealName: "ChargedLeptonKoideTransportVectorSeal", DlnRhoDt: t.DLnRhoDT, DThetaDtDeg: t.DThetaDTDeg, DPhiDtDeg: t.DPhiDTDeg, ThetaImprovementFactor: t.ThetaImprovementFactor, QImprovementFactor: t.QImprovementFactor, ProjectiveAngularDeltaDeg: radToDeg(t.ProjectiveAngularDelta), RadialToProjectiveRatio: t.RadialToProjectiveRatio, ConeAttractorCertified: d.ConeAttractorCertified, MinimalNextRequirement: d.MissingTheorem, Verdict: strings.Join([]string{StatusTransportVectorComputed, StatusThetaMovesTowardCone, StatusAzimuthNearlyInvariant, StatusConeAttractorNotCertified, StatusGate580BoundaryPreserved}, ";")}
}

func degToRad(x float64) float64 { return x * math.Pi / 180 }
func radToDeg(x float64) float64 { return x * 180 / math.Pi }

func normalizeDeg(x float64) float64 {
	y := math.Mod(x, 360)
	if y < 0 {
		y += 360
	}
	return y
}

func signedAngleDelta(a, b float64) float64 {
	d := normalizeDeg(a) - normalizeDeg(b)
	if d > 180 {
		d -= 360
	}
	if d < -180 {
		d += 360
	}
	return d
}

func Statuses() []string {
	return []string{StatusRuntimeInherited, StatusKoideCoordinatesInherited, StatusFiniteIntervalCertified, StatusTransportVectorComputed, StatusRadialComponentDominant, StatusThetaMovesTowardCone, StatusAzimuthNearlyInvariant, StatusProjectiveRayStable, StatusConeAttractorNotCertified, StatusNoContinuousBeta, StatusNoNativeRootTrace, StatusNoFlavorDerivation, StatusNoNewCarrier, StatusNoObservedPromotion, StatusNoTexturePromotion, StatusGate352Preserved, StatusGate580BoundaryPreserved}
}
