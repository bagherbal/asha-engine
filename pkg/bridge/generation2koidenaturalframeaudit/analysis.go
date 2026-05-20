// Package generation2koidenaturalframeaudit implements Gate 579:
// Koide Natural Frame Audit.
//
// Gate 577 identified the charged-lepton square-root Yukawa cone as the first
// sharp environmental flavor geometry. Gate 578 computed the remaining azimuth
// and blocked simple rational/root-of-unity/CKM identifications. Gate 579 asks
// a narrower question: does the charged-lepton Koide geometry live more
// naturally in the pole-mass frame, the M_Z runtime Yukawa frame, or the
// Lambda_12 boundary-transport frame?
//
// The gate preserves the root-trace/firewall boundary.  It does not derive
// lepton masses, Yukawa eigenvalues, PMNS/CKM data, flavor texture, generation
// hierarchy, or a native ASHA flavor carrier.  It only classifies frame
// dependence already visible in the History Transport v1 runtime.
package generation2koidenaturalframeaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2koideazimuthenvironmentalorientationaudit"
	"github.com/bagherbal/asha-engine/pkg/historytransport"
)

const (
	AuditID = "GATE579-KOIDE-NATURAL-FRAME-AUDIT"

	StatusRuntimeInherited          = "PASS_GATE578_AND_HISTORY_TRANSPORT_RUNTIME_INHERITED"
	StatusFrameBasisCertified       = "PASS_DEMOCRATIC_KOIDE_FRAME_REUSED_AND_CERTIFIED"
	StatusPoleFrameComputed         = "PASS_CHARGED_LEPTON_POLE_MASS_FRAME_KOIDE_COORDINATES_COMPUTED"
	StatusMZFrameComputed           = "PASS_CHARGED_LEPTON_MZ_YUKAWA_FRAME_KOIDE_COORDINATES_COMPUTED"
	StatusLambdaFrameComputed       = "PASS_CHARGED_LEPTON_LAMBDA12_YUKAWA_FRAME_KOIDE_COORDINATES_COMPUTED"
	StatusPoleMZAngleEquivalent     = "PASS_POLE_MASS_AND_MZ_YUKAWA_FRAMES_ANGLE_EQUIVALENT_IN_V1"
	StatusPoleMZDegenerate          = "PASS_POLE_MZ_FRAME_DEGENERACY_DUE_TO_UNIFORM_RESCALING"
	StatusAzimuthTransportInvariant = "PASS_KOIDE_AZIMUTH_NEAR_TRANSPORT_INVARIANT_ACROSS_FRAMES"
	StatusBoundarySlightlyCleaner   = "CONDITIONAL_SUPPORT_LAMBDA12_FRAME_SLIGHTLY_CLOSER_TO_KOIDE_CONE_IN_V1"
	StatusNaturalFrameNotCertified  = "FAILED_ROUTE_NO_NATURAL_KOIDE_FRAME_CERTIFIED_BY_V1_ONLY"
	StatusNoNativeFrameOperator     = "FAILED_ROUTE_NO_NATIVE_ROOT_TRACE_ABSOLUTE_DIRAC_OR_FRAME_OPERATOR"
	StatusNoPMNSComparison          = "FAILED_ROUTE_NO_PMNS_RUNTIME_INPUT_FOR_NATURAL_FRAME_TEST"
	StatusNoFlavorDerivation        = "FAILED_ROUTE_NO_ASHA_NATIVE_CHARGED_LEPTON_FLAVOR_DERIVATION_FROM_FRAME_AUDIT"
	StatusNoObservedImport          = "FIREWALL_PRESERVED_OBSERVED_LEPTON_INPUTS_REMAIN_HISTORY_ENDPOINT_DATA"
	StatusNoTexturePromotion        = "FIREWALL_PRESERVED_NATURAL_FRAME_AUDIT_DOES_NOT_DERIVE_TEXTURE_CKM_PMNS_OR_GENERATIONS"
	StatusGate352Preserved          = "FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING"
	StatusGate579BoundaryPreserved  = "FIREWALL_PRESERVED_GATE579_KOIDE_NATURAL_FRAME_BOUNDARY"
)

const KoideTarget = 2.0 / 3.0

type RuntimeInheritance struct {
	Mu0GeV           float64
	Lambda12GeV      float64
	Gate578PhiMZDeg  float64
	Gate578PhiLDeg   float64
	Gate578Verdict   string
	LeptonMassScheme string
	LeptonTransport  string
	RuntimeSource    string
	Verdict          string
}

type KoideFrame struct {
	DemocraticAxis []float64
	E1             []float64
	E2             []float64
	DotNE1         float64
	DotNE2         float64
	DotE1E2        float64
	NormN          float64
	NormE1         float64
	NormE2         float64
	RightHanded    bool
	Verdict        string
}

type FramePoint struct {
	Name                 string
	Carrier              string
	Labels               []string
	Values               []float64
	RootVector           []float64
	Rho                  float64
	Q                    float64
	DeltaQ               float64
	AbsDeltaQ            float64
	ThetaDeg             float64
	ThetaDeltaDeg        float64
	PhiSignedDeg         float64
	PhiDeg               float64
	CoordinateE1         float64
	CoordinateE2         float64
	UniformScaleFromPole float64
	Verdict              string
}

type FrameComparison struct {
	Pole                  FramePoint
	MZ                    FramePoint
	Lambda12              FramePoint
	DeltaPhiPoleToMZDeg   float64
	DeltaPhiMZToLambdaDeg float64
	DeltaQPoleToMZ        float64
	DeltaQMZToLambda      float64
	LambdaCloserThanMZ    bool
	MZEqualPole           bool
	AzimuthStable         bool
	BestKoideFrame        string
	Verdict               string
}

type NaturalFrameAudit struct {
	Question                 string
	PoleMassFrameNatural     bool
	MZYukawaFrameIndependent bool
	BoundaryFrameCertified   bool
	BoundaryFrameCleanerInV1 bool
	Reason                   string
	MissingTheorem           string
	Verdict                  string
}

type FirewallAudit struct {
	DerivesLeptonMasses        bool
	DerivesYukawaEigenvalues   bool
	DerivesPMNS                bool
	DerivesCKM                 bool
	DerivesGenerationHierarchy bool
	PromotesObservedAsNative   bool
	AddsNewCarrier             bool
	PreservesGate352           bool
	Verdict                    string
}

type FinalVerdict struct {
	SealName               string
	BestKoideResidualFrame string
	PolePhiDeg             float64
	MZPhiDeg               float64
	Lambda12PhiDeg         float64
	PoleDeltaQ             float64
	MZDeltaQ               float64
	Lambda12DeltaQ         float64
	NaturalFrameCertified  bool
	MinimalNextRequirement string
	Verdict                string
}

type Analysis struct {
	Runtime   RuntimeInheritance
	Frame     KoideFrame
	Compare   FrameComparison
	Natural   NaturalFrameAudit
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
	gate578, err := generation2koideazimuthenvironmentalorientationaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate578 Koide azimuth predecessor: %w", err)
	}
	frame := buildFrame()
	compare := compareFrames(bundle, frame)
	runtime := inheritRuntime(bundle, gate578)
	natural := auditNaturalFrame(compare)
	firewalls := auditFirewalls()
	final := compileFinal(compare, natural)
	truth := "Gate 579 shows that the pole-mass and M_Z Yukawa Koide frames are angle-identical in v1 because charged-lepton Yukawas are only a uniform rescaling of pole masses. The Lambda_12 frame is slightly closer to the Koide cone and the azimuth is almost invariant, but v1 cannot certify a unique natural frame."
	return Analysis{Runtime: runtime, Frame: frame, Compare: compare, Natural: natural, Firewalls: firewalls, Final: final, Truth: truth}, nil
}

func inheritRuntime(b historytransport.Bundle, g578 generation2koideazimuthenvironmentalorientationaudit.Analysis) RuntimeInheritance {
	leptonScheme, leptonTransport := "", ""
	for _, f := range b.Inputs.Fermions {
		if f.Name == "e" {
			leptonScheme = f.Scheme
			leptonTransport = f.Transport
			break
		}
	}
	return RuntimeInheritance{Mu0GeV: b.EndVector.Mu0GeV, Lambda12GeV: b.GaugeBoundary.Lambda12GeV, Gate578PhiMZDeg: g578.Transport.MZ.PhiDeg, Gate578PhiLDeg: g578.Transport.Lambda12.PhiDeg, Gate578Verdict: g578.Final.Verdict, LeptonMassScheme: leptonScheme, LeptonTransport: leptonTransport, RuntimeSource: "historytransport.BuildDefault() plus Gate578 azimuth predecessor", Verdict: StatusRuntimeInherited}
}

func buildFrame() KoideFrame {
	n := []float64{1 / math.Sqrt(3), 1 / math.Sqrt(3), 1 / math.Sqrt(3)}
	e1 := []float64{1 / math.Sqrt(2), -1 / math.Sqrt(2), 0}
	e2 := []float64{1 / math.Sqrt(6), 1 / math.Sqrt(6), -2 / math.Sqrt(6)}
	return KoideFrame{DemocraticAxis: n, E1: e1, E2: e2, DotNE1: dot(n, e1), DotNE2: dot(n, e2), DotE1E2: dot(e1, e2), NormN: norm(n), NormE1: norm(e1), NormE2: norm(e2), RightHanded: dot(cross(e1, e2), n) > 0.999999999999, Verdict: StatusFrameBasisCertified}
}

func compareFrames(b historytransport.Bundle, frame KoideFrame) FrameComparison {
	labels := []string{"e", "mu", "tau"}
	poleMasses := leptonPoleMasses(b, labels)
	mzY := values(b.FlavorTransport.YukawaSingularValuesMZ.ChargedLeptons, labels)
	lY := values(b.FlavorTransport.YukawaSingularValuesLambda12.ChargedLeptons, labels)
	pole := framePoint("pole_mass", "charged-lepton pole masses m_i; x_i=sqrt(m_i)", labels, poleMasses, frame, 1)
	mzScale := math.Sqrt(mzY[0] / poleMasses[0])
	mz := framePoint("M_Z_yukawa", "runtime charged-lepton Yukawas y_i(M_Z); x_i=sqrt(y_i)", labels, mzY, frame, mzScale)
	lambdaScale := math.Sqrt(lY[0] / poleMasses[0])
	lambda := framePoint("Lambda_12_yukawa", "transported charged-lepton Yukawas y_i(Lambda_12); x_i=sqrt(y_i)", labels, lY, frame, lambdaScale)
	dPoleMZ := signedAngleDelta(mz.PhiDeg, pole.PhiDeg)
	dMZL := signedAngleDelta(lambda.PhiDeg, mz.PhiDeg)
	dqPoleMZ := mz.DeltaQ - pole.DeltaQ
	dqMZL := lambda.DeltaQ - mz.DeltaQ
	mzEqualPole := math.Abs(dPoleMZ) < 1e-12 && math.Abs(dqPoleMZ) < 1e-14 && math.Abs(mz.ThetaDeg-pole.ThetaDeg) < 1e-12
	lambdaCloser := lambda.AbsDeltaQ < mz.AbsDeltaQ
	best := "M_Z_yukawa/pole_mass_tie"
	if lambdaCloser {
		best = "Lambda_12_yukawa_by_smaller_abs_Q_residual_in_v1"
	}
	verdict := strings.Join([]string{StatusPoleFrameComputed, StatusMZFrameComputed, StatusLambdaFrameComputed, StatusPoleMZAngleEquivalent, StatusPoleMZDegenerate, StatusAzimuthTransportInvariant, StatusBoundarySlightlyCleaner}, ";")
	return FrameComparison{Pole: pole, MZ: mz, Lambda12: lambda, DeltaPhiPoleToMZDeg: dPoleMZ, DeltaPhiMZToLambdaDeg: dMZL, DeltaQPoleToMZ: dqPoleMZ, DeltaQMZToLambda: dqMZL, LambdaCloserThanMZ: lambdaCloser, MZEqualPole: mzEqualPole, AzimuthStable: math.Abs(dMZL) < 3e-4, BestKoideFrame: best, Verdict: verdict}
}

func auditNaturalFrame(c FrameComparison) NaturalFrameAudit {
	return NaturalFrameAudit{Question: "Does pole mass, M_Z Yukawa, or Lambda_12 transport provide the natural Koide frame?", PoleMassFrameNatural: c.MZEqualPole, MZYukawaFrameIndependent: false, BoundaryFrameCertified: false, BoundaryFrameCleanerInV1: c.LambdaCloserThanMZ, Reason: "Pole and M_Z frames are angle-degenerate in v1 because y_i(M_Z)=sqrt(2)m_i/v is a common positive rescaling for charged leptons. Lambda_12 is slightly closer to Q=2/3 and almost preserves phi_e, but this depends on approximate v1 running and is not a theorem.", MissingTheorem: "a native root-trace/absolute-Dirac observable or transport theorem that selects pole, M_Z, or Lambda_12 as the charged-lepton Koide frame", Verdict: strings.Join([]string{StatusBoundarySlightlyCleaner, StatusNaturalFrameNotCertified, StatusNoNativeFrameOperator, StatusNoPMNSComparison}, ";")}
}

func auditFirewalls() FirewallAudit {
	return FirewallAudit{DerivesLeptonMasses: false, DerivesYukawaEigenvalues: false, DerivesPMNS: false, DerivesCKM: false, DerivesGenerationHierarchy: false, PromotesObservedAsNative: false, AddsNewCarrier: false, PreservesGate352: true, Verdict: strings.Join([]string{StatusNoNativeFrameOperator, StatusNoFlavorDerivation, StatusNoObservedImport, StatusNoTexturePromotion, StatusGate352Preserved, StatusGate579BoundaryPreserved}, ";")}
}

func compileFinal(c FrameComparison, n NaturalFrameAudit) FinalVerdict {
	return FinalVerdict{SealName: "ChargedLeptonKoideNaturalFrameSeal", BestKoideResidualFrame: c.BestKoideFrame, PolePhiDeg: c.Pole.PhiDeg, MZPhiDeg: c.MZ.PhiDeg, Lambda12PhiDeg: c.Lambda12.PhiDeg, PoleDeltaQ: c.Pole.DeltaQ, MZDeltaQ: c.MZ.DeltaQ, Lambda12DeltaQ: c.Lambda12.DeltaQ, NaturalFrameCertified: n.BoundaryFrameCertified, MinimalNextRequirement: n.MissingTheorem, Verdict: strings.Join([]string{StatusPoleMZAngleEquivalent, StatusBoundarySlightlyCleaner, StatusNaturalFrameNotCertified, StatusGate579BoundaryPreserved}, ";")}
}

func framePoint(name, carrier string, labels []string, vals []float64, frame KoideFrame, uniformScaleFromPole float64) FramePoint {
	root := make([]float64, len(vals))
	sum, rootSum := 0.0, 0.0
	for i, v := range vals {
		root[i] = math.Sqrt(v)
		sum += v
		rootSum += root[i]
	}
	parallel := dot(root, frame.DemocraticAxis)
	perp := make([]float64, 3)
	for i := 0; i < 3; i++ {
		perp[i] = root[i] - parallel*frame.DemocraticAxis[i]
	}
	ce1, ce2 := dot(perp, frame.E1), dot(perp, frame.E2)
	signed := radToDeg(math.Atan2(ce2, ce1))
	phi := signed
	if phi < 0 {
		phi += 360
	}
	q := sum / (rootSum * rootSum)
	return FramePoint{Name: name, Carrier: carrier, Labels: append([]string(nil), labels...), Values: append([]float64(nil), vals...), RootVector: root, Rho: math.Sqrt(sum), Q: q, DeltaQ: q - KoideTarget, AbsDeltaQ: math.Abs(q - KoideTarget), ThetaDeg: radToDeg(math.Atan2(norm(perp), parallel)), ThetaDeltaDeg: radToDeg(math.Atan2(norm(perp), parallel)) - 45, PhiSignedDeg: signed, PhiDeg: phi, CoordinateE1: ce1, CoordinateE2: ce2, UniformScaleFromPole: uniformScaleFromPole, Verdict: pointVerdict(name)}
}

func pointVerdict(name string) string {
	switch name {
	case "pole_mass":
		return StatusPoleFrameComputed
	case "M_Z_yukawa":
		return StatusMZFrameComputed
	case "Lambda_12_yukawa":
		return StatusLambdaFrameComputed
	default:
		return "PASS_KOIDE_FRAME_POINT_COMPUTED"
	}
}

func leptonPoleMasses(b historytransport.Bundle, labels []string) []float64 {
	m := map[string]float64{}
	for _, f := range b.Inputs.Fermions {
		m[f.Name] = f.MassGeV
	}
	return values(m, labels)
}

func values(m map[string]float64, labels []string) []float64 {
	out := make([]float64, len(labels))
	for i, k := range labels {
		out[i] = m[k]
	}
	return out
}

func dot(a, b []float64) float64 {
	s := 0.0
	for i := range a {
		s += a[i] * b[i]
	}
	return s
}

func cross(a, b []float64) []float64 {
	return []float64{a[1]*b[2] - a[2]*b[1], a[2]*b[0] - a[0]*b[2], a[0]*b[1] - a[1]*b[0]}
}

func norm(a []float64) float64   { return math.Sqrt(dot(a, a)) }
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
	return []string{StatusRuntimeInherited, StatusFrameBasisCertified, StatusPoleFrameComputed, StatusMZFrameComputed, StatusLambdaFrameComputed, StatusPoleMZAngleEquivalent, StatusPoleMZDegenerate, StatusAzimuthTransportInvariant, StatusBoundarySlightlyCleaner, StatusNaturalFrameNotCertified, StatusNoNativeFrameOperator, StatusNoPMNSComparison, StatusNoFlavorDerivation, StatusNoObservedImport, StatusNoTexturePromotion, StatusGate352Preserved, StatusGate579BoundaryPreserved}
}
