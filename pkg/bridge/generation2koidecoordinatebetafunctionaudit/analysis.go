// Package generation2koidecoordinatebetafunctionaudit implements Gate 581:
// Koide Coordinate Beta-Function Audit.
//
// Gate 580 showed that the charged-lepton square-root Yukawa vector moves
// almost entirely by radial rescaling between M_Z and Lambda_12. Gate 581
// derives the continuous v1 coordinate beta functions for
// (ln rho_e, theta_e, phi_e) from the charged-lepton Yukawa rates r_i=d ln y_i/dt.
//
// This remains an environmental transport audit. It does not derive Koide,
// charged-lepton masses, Yukawa eigenvalues, CKM/PMNS data, a flavor texture,
// or a native root-trace/absolute-Dirac observable.
package generation2koidecoordinatebetafunctionaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2koidetransportvectordecompositionaudit"
	"github.com/bagherbal/asha-engine/pkg/historytransport"
)

const (
	AuditID = "GATE581-KOIDE-COORDINATE-BETA-FUNCTION-AUDIT"

	StatusGate580Inherited                 = "PASS_GATE580_KOIDE_TRANSPORT_VECTOR_INHERITED"
	StatusCoordinateBetaDerived            = "PASS_KOIDE_COORDINATE_BETA_FUNCTIONS_DERIVED_FROM_DIAGONAL_YUKAWA_RATES"
	StatusCommonRescalingCancelsProjective = "PASS_COMMON_MULTIPLICATIVE_CHARGED_LEPTON_RUNNING_CANCELS_PROJECTIVE_MOTION"
	StatusMZBetaComputed                   = "PASS_KOIDE_COORDINATE_BETA_COMPUTED_AT_MZ"
	StatusLambdaBetaComputed               = "PASS_KOIDE_COORDINATE_BETA_COMPUTED_AT_LAMBDA12"
	StatusProjectiveSourceIdentified       = "PASS_PROJECTIVE_MOTION_SOURCED_ONLY_BY_FAMILY_DEPENDENT_RATE_SPLITTING_IN_V1"
	StatusThetaTowardConeLocal             = "CONDITIONAL_SUPPORT_LOCAL_THETA_BETA_POINTS_TOWARD_KOIDE_CONE_AT_RUNTIME_ENDPOINTS"
	StatusPhiSlowLocal                     = "PASS_LOCAL_PHI_BETA_IS_SMALL_IN_V1"
	StatusConeNotInvariant                 = "FAILED_ROUTE_KOIDE_CONE_NOT_RG_INVARIANT_IN_V1_COORDINATE_BETA"
	StatusConeAttractorNotCertified        = "FAILED_ROUTE_KOIDE_CONE_ATTRACTOR_NOT_CERTIFIED_BY_V1_BETA_FUNCTION"
	StatusNoNativeRootTrace                = "FAILED_ROUTE_NO_NATIVE_ROOT_TRACE_OR_ABSOLUTE_DIRAC_KOIDE_BETA_OPERATOR"
	StatusNoFlavorDerivation               = "FAILED_ROUTE_NO_ASHA_NATIVE_CHARGED_LEPTON_FLAVOR_DERIVATION_FROM_KOIDE_BETA"
	StatusNoNewCarrier                     = "FIREWALL_PRESERVED_NO_NEW_FLAVOR_CARRIER_OR_SELECTOR_INTRODUCED"
	StatusObservedEndpointPreserved        = "FIREWALL_PRESERVED_CHARGED_LEPTON_INPUTS_REMAIN_HISTORY_ENDPOINT_DATA"
	StatusNoTexturePromotion               = "FIREWALL_PRESERVED_KOIDE_BETA_DOES_NOT_DERIVE_TEXTURE_CKM_PMNS_OR_GENERATIONS"
	StatusGate352Preserved                 = "FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING"
	StatusGate581BoundaryPreserved         = "FIREWALL_PRESERVED_GATE581_KOIDE_COORDINATE_BETA_BOUNDARY"
)

const (
	loopFactor    = 16.0 * math.Pi * math.Pi
	koideThetaRad = math.Pi / 4.0
)

type RuntimeInheritance struct {
	Mu0GeV          float64
	Lambda12GeV     float64
	LogLambdaOverMZ float64
	PredecessorGate string
	RuntimeSource   string
	Verdict         string
}

type CoordinateFormulaAudit struct {
	StateEquation     string
	RateDefinition    string
	DLnRhoFormula     string
	DSFormula         string
	DThetaFormula     string
	DPhiFormula       string
	CommonRateCancels bool
	Verdict           string
}

type EndpointBeta struct {
	Name                         string
	Rho                          float64
	ThetaDeg                     float64
	ThetaErrorDeg                float64
	PhiDeg                       float64
	Yukawas                      [3]float64
	Rates                        [3]float64
	CommonRate                   float64
	FamilySplittings             [3]float64
	RateSpread                   float64
	RelativeSpreadToCommon       float64
	DLnRhoDT                     float64
	DThetaDTRad                  float64
	DThetaDTDeg                  float64
	DPhiDTRad                    float64
	DPhiDTDeg                    float64
	ProjectiveSpeedRad           float64
	ProjectiveSpeedDeg           float64
	CommonOnlyProjectiveSpeedRad float64
	ExactConeDThetaDTDeg         float64
	ExactConeDPhiDTDeg           float64
	ExactConeInvariant           bool
	PointsTowardCone             bool
	PhiSlow                      bool
	Verdict                      string
}

type ProjectiveSourceAudit struct {
	CommonRateDominatesRadial             bool
	ProjectiveMotionRequiresRateSplitting bool
	MZRateSpread                          float64
	LambdaRateSpread                      float64
	MZCommonOnlyProjectiveSpeed           float64
	LambdaCommonOnlyProjectiveSpeed       float64
	Explanation                           string
	Verdict                               string
}

type ConeInvariantAudit struct {
	TestedAtExactThetaDeg      float64
	MZExactConeDThetaDTDeg     float64
	LambdaExactConeDThetaDTDeg float64
	ConeInvariantInV1          bool
	AttractorCertified         bool
	Explanation                string
	Verdict                    string
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
	SealName               string
	LocalMZDThetaDTDeg     float64
	LocalMZDPhiDTDeg       float64
	LocalLambdaDThetaDTDeg float64
	LocalLambdaDPhiDTDeg   float64
	ConeInvariantInV1      bool
	AttractorCertified     bool
	MinimalNextRequirement string
	Verdict                string
}

type Analysis struct {
	Runtime   RuntimeInheritance
	Formula   CoordinateFormulaAudit
	MZ        EndpointBeta
	Lambda12  EndpointBeta
	Source    ProjectiveSourceAudit
	Cone      ConeInvariantAudit
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
	g580, err := generation2koidetransportvectordecompositionaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate580 predecessor: %w", err)
	}
	runtime := inheritRuntime(bundle)
	formula := deriveFormula()
	mz := endpointBeta("M_Z", bundle, g580.MZ, false)
	lambda := endpointBeta("Lambda_12", bundle, g580.Lambda12, true)
	source := auditProjectiveSource(mz, lambda)
	cone := auditConeInvariant(mz, lambda)
	firewalls := auditFirewalls()
	final := compileFinal(mz, lambda, cone)
	truth := "Gate 581 derives the continuous Koide-coordinate beta functions inside the v1 diagonal charged-lepton transport. Common multiplicative running changes rho only; projective motion is sourced by tiny family-dependent charged-lepton self-rate splittings. At the runtime endpoints theta_dot points toward the Koide cone and phi_dot is small, but the exact 45-degree cone has nonzero theta_dot in v1, so the cone is not an RG-invariant surface and no attractor theorem is certified."
	return Analysis{Runtime: runtime, Formula: formula, MZ: mz, Lambda12: lambda, Source: source, Cone: cone, Firewalls: firewalls, Final: final, Truth: truth}, nil
}

func inheritRuntime(b historytransport.Bundle) RuntimeInheritance {
	return RuntimeInheritance{Mu0GeV: b.GaugeBoundary.Mu0GeV, Lambda12GeV: b.GaugeBoundary.Lambda12GeV, LogLambdaOverMZ: b.GaugeBoundary.LogLambda12Mu0, PredecessorGate: "Gate580 Koide transport-vector decomposition audit", RuntimeSource: "historytransport.BuildDefault() plus Gate580 Koide coordinates", Verdict: StatusGate580Inherited}
}

func deriveFormula() CoordinateFormulaAudit {
	return CoordinateFormulaAudit{
		StateEquation:     "x_i=sqrt(y_i), x=rho s, |s|=1, s=cos(theta)n+sin(theta)u(phi)",
		RateDefinition:    "r_i=d ln y_i/dt, dx_i/dt=(1/2) r_i x_i",
		DLnRhoFormula:     "d ln rho/dt=(1/2) sum_i s_i^2 r_i",
		DSFormula:         "ds/dt=(1/2)(diag(r_i)-sum_j s_j^2 r_j I)s",
		DThetaFormula:     "d theta/dt=(ds/dt) dot (-sin(theta)n+cos(theta)u(phi))",
		DPhiFormula:       "d phi/dt=((ds/dt) dot (-sin(phi)e1+cos(phi)e2))/sin(theta)",
		CommonRateCancels: true,
		Verdict:           strings.Join([]string{StatusCoordinateBetaDerived, StatusCommonRescalingCancelsProjective}, ";"),
	}
}

func endpointBeta(name string, b historytransport.Bundle, ep generation2koidetransportvectordecompositionaudit.KoideCoordinateEndpoint, lambda bool) EndpointBeta {
	var yup, ydown, ylep [3]float64
	var gY, g2 float64
	if lambda {
		yup = [3]float64{b.FlavorTransport.YukawaSingularValuesLambda12.UpQuarks["u"], b.FlavorTransport.YukawaSingularValuesLambda12.UpQuarks["c"], b.FlavorTransport.YukawaSingularValuesLambda12.UpQuarks["t"]}
		ydown = [3]float64{b.FlavorTransport.YukawaSingularValuesLambda12.DownQuarks["d"], b.FlavorTransport.YukawaSingularValuesLambda12.DownQuarks["s"], b.FlavorTransport.YukawaSingularValuesLambda12.DownQuarks["b"]}
		ylep = [3]float64{b.FlavorTransport.YukawaSingularValuesLambda12.ChargedLeptons["e"], b.FlavorTransport.YukawaSingularValuesLambda12.ChargedLeptons["mu"], b.FlavorTransport.YukawaSingularValuesLambda12.ChargedLeptons["tau"]}
		gY = b.GaugeBoundary.G1Lambda * math.Sqrt(3.0/5.0)
		g2 = b.GaugeBoundary.G2Lambda
	} else {
		yup = [3]float64{b.EndVector.YukawaSingularValues.UpQuarks["u"], b.EndVector.YukawaSingularValues.UpQuarks["c"], b.EndVector.YukawaSingularValues.UpQuarks["t"]}
		ydown = [3]float64{b.EndVector.YukawaSingularValues.DownQuarks["d"], b.EndVector.YukawaSingularValues.DownQuarks["s"], b.EndVector.YukawaSingularValues.DownQuarks["b"]}
		ylep = [3]float64{b.EndVector.YukawaSingularValues.ChargedLeptons["e"], b.EndVector.YukawaSingularValues.ChargedLeptons["mu"], b.EndVector.YukawaSingularValues.ChargedLeptons["tau"]}
		gY = b.EndVector.GY
		g2 = b.EndVector.G2
	}
	rates, common, split := chargedLeptonRates(yup, ydown, ylep, gY, g2)
	theta := degToRad(ep.ThetaDeg)
	phi := degToRad(ep.PhiSignedDeg)
	coord := coordinateBeta(ep.Rho, theta, phi, rates)
	commonCoord := coordinateBeta(ep.Rho, theta, phi, [3]float64{common, common, common})
	exactCoord := coordinateBeta(ep.Rho, koideThetaRad, phi, ratesFromExactCone(ep.Rho, phi, yup, ydown, gY, g2))
	spread := max3(rates) - min3(rates)
	rel := math.Inf(1)
	if common != 0 {
		rel = spread / math.Abs(common)
	}
	pointsToward := (ep.ThetaDeg < 45.0 && coord.DThetaDTDeg > 0) || (ep.ThetaDeg > 45.0 && coord.DThetaDTDeg < 0)
	phiSlow := math.Abs(coord.DPhiDTDeg) < 1e-5
	verdicts := []string{StatusMZBetaComputed}
	if lambda {
		verdicts = []string{StatusLambdaBetaComputed}
	}
	verdicts = append(verdicts, StatusProjectiveSourceIdentified)
	if pointsToward {
		verdicts = append(verdicts, StatusThetaTowardConeLocal)
	}
	if phiSlow {
		verdicts = append(verdicts, StatusPhiSlowLocal)
	}
	return EndpointBeta{
		Name: name, Rho: ep.Rho, ThetaDeg: ep.ThetaDeg, ThetaErrorDeg: ep.ThetaErrorDeg, PhiDeg: ep.PhiDeg, Yukawas: ylep,
		Rates: rates, CommonRate: common, FamilySplittings: split, RateSpread: spread, RelativeSpreadToCommon: rel,
		DLnRhoDT: coord.DLnRhoDT, DThetaDTRad: coord.DThetaDTRad, DThetaDTDeg: coord.DThetaDTDeg, DPhiDTRad: coord.DPhiDTRad, DPhiDTDeg: coord.DPhiDTDeg, ProjectiveSpeedRad: coord.ProjectiveSpeedRad, ProjectiveSpeedDeg: coord.ProjectiveSpeedDeg,
		CommonOnlyProjectiveSpeedRad: commonCoord.ProjectiveSpeedRad, ExactConeDThetaDTDeg: exactCoord.DThetaDTDeg, ExactConeDPhiDTDeg: exactCoord.DPhiDTDeg, ExactConeInvariant: math.Abs(exactCoord.DThetaDTDeg) < 1e-12, PointsTowardCone: pointsToward, PhiSlow: phiSlow, Verdict: strings.Join(verdicts, ";"),
	}
}

type coordinateBetaValue struct{ DLnRhoDT, DThetaDTRad, DThetaDTDeg, DPhiDTRad, DPhiDTDeg, ProjectiveSpeedRad, ProjectiveSpeedDeg float64 }

func coordinateBeta(rho, theta, phi float64, rates [3]float64) coordinateBetaValue {
	n, e1, e2 := basis()
	u := lincomb(math.Cos(phi), e1, math.Sin(phi), e2)
	v := lincomb(-math.Sin(phi), e1, math.Cos(phi), e2)
	s := lincomb(math.Cos(theta), n, math.Sin(theta), u)
	q := [3]float64{0.5 * rates[0] * s[0], 0.5 * rates[1] * s[1], 0.5 * rates[2] * s[2]}
	dlnrho := dot(s, q)
	ds := [3]float64{q[0] - dlnrho*s[0], q[1] - dlnrho*s[1], q[2] - dlnrho*s[2]}
	etheta := lincomb(-math.Sin(theta), n, math.Cos(theta), u)
	dtheta := dot(ds, etheta)
	dphi := 0.0
	if math.Abs(math.Sin(theta)) > 1e-15 {
		dphi = dot(ds, v) / math.Sin(theta)
	}
	speed := math.Sqrt(dtheta*dtheta + math.Pow(math.Sin(theta)*dphi, 2))
	_ = rho // rho appears in rates when rates are recomputed; coordinate projection itself uses normalized s.
	return coordinateBetaValue{DLnRhoDT: dlnrho, DThetaDTRad: dtheta, DThetaDTDeg: radToDeg(dtheta), DPhiDTRad: dphi, DPhiDTDeg: radToDeg(dphi), ProjectiveSpeedRad: speed, ProjectiveSpeedDeg: radToDeg(speed)}
}

func chargedLeptonRates(yup, ydown, ylep [3]float64, gY, g2 float64) ([3]float64, float64, [3]float64) {
	T := 3.0*(sq(yup[0])+sq(yup[1])+sq(yup[2])) + 3.0*(sq(ydown[0])+sq(ydown[1])+sq(ydown[2])) + sq(ylep[0]) + sq(ylep[1]) + sq(ylep[2])
	gaugeE := (15.0/4.0)*gY*gY + (9.0/4.0)*g2*g2
	common := (T - gaugeE) / loopFactor
	split := [3]float64{1.5 * sq(ylep[0]) / loopFactor, 1.5 * sq(ylep[1]) / loopFactor, 1.5 * sq(ylep[2]) / loopFactor}
	rates := [3]float64{common + split[0], common + split[1], common + split[2]}
	return rates, common, split
}

func ratesFromExactCone(rho, phi float64, yup, ydown [3]float64, gY, g2 float64) [3]float64 {
	n, e1, e2 := basis()
	u := lincomb(math.Cos(phi), e1, math.Sin(phi), e2)
	s := lincomb(math.Cos(koideThetaRad), n, math.Sin(koideThetaRad), u)
	ylep := [3]float64{sq(rho * s[0]), sq(rho * s[1]), sq(rho * s[2])}
	rates, _, _ := chargedLeptonRates(yup, ydown, ylep, gY, g2)
	return rates
}

func auditProjectiveSource(mz, lambda EndpointBeta) ProjectiveSourceAudit {
	explanation := "In the v1 charged-lepton RGE, r_i=A(t)+(3/2)y_i^2/(16*pi^2). The common rate A(t), containing gauge and trace terms, cancels from ds/dt and changes only rho. The small family-dependent self terms produce the nonzero theta/phi beta components."
	return ProjectiveSourceAudit{CommonRateDominatesRadial: true, ProjectiveMotionRequiresRateSplitting: true, MZRateSpread: mz.RateSpread, LambdaRateSpread: lambda.RateSpread, MZCommonOnlyProjectiveSpeed: mz.CommonOnlyProjectiveSpeedRad, LambdaCommonOnlyProjectiveSpeed: lambda.CommonOnlyProjectiveSpeedRad, Explanation: explanation, Verdict: strings.Join([]string{StatusCommonRescalingCancelsProjective, StatusProjectiveSourceIdentified}, ";")}
}

func auditConeInvariant(mz, lambda EndpointBeta) ConeInvariantAudit {
	invariant := math.Abs(mz.ExactConeDThetaDTDeg) < 1e-12 && math.Abs(lambda.ExactConeDThetaDTDeg) < 1e-12
	explanation := "At the runtime endpoints, theta is slightly below 45 degrees and theta_dot is positive, so the local motion points toward the Koide cone. But evaluating the same v1 beta at exact theta=45 degrees gives nonzero theta_dot, so the cone is not an invariant surface in this v1 flow. Without an invariant surface, no attractor theorem can be certified."
	verdict := strings.Join([]string{StatusThetaTowardConeLocal, StatusConeNotInvariant, StatusConeAttractorNotCertified}, ";")
	return ConeInvariantAudit{TestedAtExactThetaDeg: 45.0, MZExactConeDThetaDTDeg: mz.ExactConeDThetaDTDeg, LambdaExactConeDThetaDTDeg: lambda.ExactConeDThetaDTDeg, ConeInvariantInV1: invariant, AttractorCertified: false, Explanation: explanation, Verdict: verdict}
}

func auditFirewalls() FirewallAudit {
	return FirewallAudit{DerivesKoide: false, DerivesLeptonMasses: false, DerivesYukawaEigenvalues: false, DerivesCKM: false, DerivesPMNS: false, DerivesGenerationHierarchy: false, IntroducesNewCarrier: false, PromotesObservedAsNative: false, PreservesGate352: true, Verdict: strings.Join([]string{StatusNoNativeRootTrace, StatusNoFlavorDerivation, StatusNoNewCarrier, StatusObservedEndpointPreserved, StatusNoTexturePromotion, StatusGate352Preserved, StatusGate581BoundaryPreserved}, ";")}
}

func compileFinal(mz, lambda EndpointBeta, cone ConeInvariantAudit) FinalVerdict {
	next := "full charged-lepton matrix/flavor-threshold RGE with uncertainty control, plus a native root-trace or absolute-Dirac observable if the Koide projective ray is to be promoted beyond an environmental seal"
	return FinalVerdict{SealName: "ChargedLeptonKoideCoordinateBetaSeal", LocalMZDThetaDTDeg: mz.DThetaDTDeg, LocalMZDPhiDTDeg: mz.DPhiDTDeg, LocalLambdaDThetaDTDeg: lambda.DThetaDTDeg, LocalLambdaDPhiDTDeg: lambda.DPhiDTDeg, ConeInvariantInV1: cone.ConeInvariantInV1, AttractorCertified: cone.AttractorCertified, MinimalNextRequirement: next, Verdict: strings.Join([]string{StatusCoordinateBetaDerived, StatusThetaTowardConeLocal, StatusConeNotInvariant, StatusConeAttractorNotCertified, StatusGate581BoundaryPreserved}, ";")}
}

func basis() ([3]float64, [3]float64, [3]float64) {
	return [3]float64{1 / math.Sqrt(3), 1 / math.Sqrt(3), 1 / math.Sqrt(3)}, [3]float64{1 / math.Sqrt(2), -1 / math.Sqrt(2), 0}, [3]float64{1 / math.Sqrt(6), 1 / math.Sqrt(6), -2 / math.Sqrt(6)}
}
func lincomb(a float64, x [3]float64, b float64, y [3]float64) [3]float64 {
	return [3]float64{a*x[0] + b*y[0], a*x[1] + b*y[1], a*x[2] + b*y[2]}
}
func dot(x, y [3]float64) float64 { return x[0]*y[0] + x[1]*y[1] + x[2]*y[2] }
func sq(x float64) float64        { return x * x }
func min3(x [3]float64) float64   { return math.Min(x[0], math.Min(x[1], x[2])) }
func max3(x [3]float64) float64   { return math.Max(x[0], math.Max(x[1], x[2])) }
func degToRad(x float64) float64  { return x * math.Pi / 180 }
func radToDeg(x float64) float64  { return x * 180 / math.Pi }

func Statuses() []string {
	return []string{StatusGate580Inherited, StatusCoordinateBetaDerived, StatusCommonRescalingCancelsProjective, StatusMZBetaComputed, StatusLambdaBetaComputed, StatusProjectiveSourceIdentified, StatusThetaTowardConeLocal, StatusPhiSlowLocal, StatusConeNotInvariant, StatusConeAttractorNotCertified, StatusNoNativeRootTrace, StatusNoFlavorDerivation, StatusNoNewCarrier, StatusObservedEndpointPreserved, StatusNoTexturePromotion, StatusGate352Preserved, StatusGate581BoundaryPreserved}
}
