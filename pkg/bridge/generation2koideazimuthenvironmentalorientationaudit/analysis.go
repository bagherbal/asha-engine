// Package generation2koideazimuthenvironmentalorientationaudit implements Gate 578:
// Charged-Lepton Koide Azimuth Environmental Orientation Audit.
//
// Gate 577 converted the charged-lepton Yukawa magnitudes into square-root cone
// geometry: radius plus azimuth plus the Koide cone constraint. Gate 578 audits
// the remaining angular datum. It computes the charged-lepton Koide azimuth in a
// fixed democratic-axis frame, tests simple rational/spectral/CKM phase candidates,
// and quarantines the result as an environmental orientation seal unless a native
// ASHA carrier/operator is supplied. The gate does not derive lepton masses,
// flavor textures, CKM/PMNS, generation hierarchy, or observed data.
package generation2koideazimuthenvironmentalorientationaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2koideyukawasquarerootconesealaudit"
	"github.com/bagherbal/asha-engine/pkg/historytransport"
)

const (
	AuditID = "GATE578-KOIDE-AZIMUTH-ENVIRONMENTAL-ORIENTATION-AUDIT"

	StatusRuntimeInherited                = "PASS_GATE577_AND_HISTORY_TRANSPORT_RUNTIME_INHERITED"
	StatusBasisOrthonormalCertified       = "PASS_DEMOCRATIC_AXIS_AZIMUTH_FRAME_ORTHONORMAL_CERTIFIED"
	StatusAzimuthComputedMZ               = "PASS_CHARGED_LEPTON_KOIDE_AZIMUTH_COMPUTED_AT_MZ"
	StatusAzimuthComputedLambda12         = "PASS_CHARGED_LEPTON_KOIDE_AZIMUTH_COMPUTED_AT_LAMBDA12"
	StatusAzimuthTransportStable          = "PASS_KOIDE_AZIMUTH_STABLE_UNDER_V1_TRANSPORT"
	StatusConeCircleDecomposition         = "PASS_CHARGED_LEPTON_YE_REDUCED_TO_RADIUS_PLUS_AZIMUTH_ON_KOIDE_CONE"
	StatusFiveSeventhsNearButNotCertified = "CONDITIONAL_SUPPORT_NEAREST_SIMPLE_RATIONAL_IS_FIVE_SEVENTHS_TURN_BUT_NOT_CERTIFIED"
	StatusAzimuthOrientationSealCandidate = "CONDITIONAL_SUPPORT_CHARGED_LEPTON_KOIDE_AZIMUTH_ENVIRONMENTAL_ORIENTATION_SEAL_CANDIDATE"
	StatusNoSimpleRationalPhaseCertified  = "FAILED_ROUTE_NO_SIMPLE_RATIONAL_OR_ROOT_OF_UNITY_PHASE_MATCH_CERTIFIED"
	StatusNoCKMPhaseIdentification        = "FAILED_ROUTE_KOIDE_AZIMUTH_NOT_IDENTIFIED_WITH_CKM_PHASE_OR_JARLSKOG_ORIENTATION"
	StatusNoPMNSRuntimeCandidate          = "FAILED_ROUTE_NO_PMNS_RUNTIME_INPUT_FOR_KOIDE_AZIMUTH_IDENTIFICATION"
	StatusGate352ObstructionInherited     = "FAILED_ROUTE_GATE352_ROOT_TRACE_OBSTRUCTION_STILL_BLOCKS_NATIVE_KOIDE_AZIMUTH_OPERATOR"
	StatusNoNativeAzimuthDerivation       = "FAILED_ROUTE_NO_ASHA_NATIVE_KOIDE_AZIMUTH_DERIVATION"
	StatusNoFlavorPromotion               = "FIREWALL_PRESERVED_KOIDE_AZIMUTH_DOES_NOT_DERIVE_FLAVOR_TEXTURE_CKM_PMNS_OR_GENERATIONS"
	StatusObservedEndpointQuarantined     = "FIREWALL_PRESERVED_KOIDE_AZIMUTH_REMAINS_OBSERVED_HISTORY_ENDPOINT_ORIENTATION"
	StatusGate578BoundaryPreserved        = "FIREWALL_PRESERVED_GATE578_KOIDE_AZIMUTH_ENVIRONMENTAL_SEAL_BOUNDARY"
)

const (
	KoideTarget                  = 2.0 / 3.0
	DriftCertificationMultiplier = 100.0
)

type RuntimeInheritance struct {
	Mu0GeV          float64
	Lambda12GeV     float64
	Gate577QeMZ     float64
	Gate577PhiMZDeg float64
	Gate577PhiLDeg  float64
	Gate577Verdict  string
	Source          string
	Verdict         string
}

type AzimuthFrame struct {
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
	Convention     string
	Verdict        string
}

type AzimuthPoint struct {
	Scale              string
	Labels             []string
	Yukawas            []float64
	RootVector         []float64
	Rho                float64
	Q                  float64
	DeltaFromTwoThirds float64
	ThetaDeg           float64
	ThetaDeltaDeg      float64
	Parallel           float64
	PerpendicularNorm  float64
	CoordE1            float64
	CoordE2            float64
	PhiSignedDeg       float64
	PhiDeg             float64
	Verdict            string
}

type AzimuthTransport struct {
	MZ                  AzimuthPoint
	Lambda12            AzimuthPoint
	DeltaPhiDeg         float64
	AbsDeltaPhiDeg      float64
	DeltaQ              float64
	StableAt1eMinus3Deg bool
	StableAt1eMinus2Deg bool
	Verdict             string
}

type PhaseCandidate struct {
	Name         string
	Formula      string
	CandidateDeg float64
	DistanceDeg  float64
	ThresholdDeg float64
	Certified    bool
	Verdict      string
}

type CandidateAudit struct {
	PointPhiDeg                float64
	DriftScaleDeg              float64
	CertificationThresholdDeg  float64
	NearestRationalDenominator int
	NearestRationalNumerator   int
	NearestRationalTurn        string
	NearestRationalDistanceDeg float64
	Candidates                 []PhaseCandidate
	AnyCertified               bool
	Verdict                    string
}

type AzimuthSeal struct {
	Name                           string
	Carrier                        string
	Coordinates                    []string
	ConeConstraint                 string
	AzimuthDefinition              string
	RadiusMZ                       float64
	PhiMZDeg                       float64
	PhiLambda12Deg                 float64
	DriftDeg                       float64
	OriginalPositiveMagnitudes     int
	Constraints                    int
	RemainingContinuousCoordinates int
	NativeDerivation               bool
	BridgeOnly                     bool
	Verdict                        string
}

type FirewallAudit struct {
	DerivesChargedLeptonMasses        bool
	DerivesYukawaEigenvalues          bool
	DerivesCKM                        bool
	DerivesPMNS                       bool
	DerivesGenerationHierarchy        bool
	IdentifiesWithASHAProjectivePhase bool
	ImportsObservedAsNative           bool
	AddsNewCarrier                    bool
	PreservesGate352                  bool
	Verdict                           string
}

type FinalVerdict struct {
	SealName                   string
	PhiMZDeg                   float64
	PhiSignedMZDeg             float64
	PhiLambda12Deg             float64
	DeltaPhiDeg                float64
	NearestSimpleCandidate     string
	NearestSimpleDistanceDeg   float64
	CertifiedSimplePhase       bool
	NativeASHAFlavorDerivation bool
	NextRequiredTheorem        string
	Verdict                    string
}

type Analysis struct {
	Runtime    RuntimeInheritance
	Frame      AzimuthFrame
	Transport  AzimuthTransport
	Candidates CandidateAudit
	Seal       AzimuthSeal
	Firewalls  FirewallAudit
	Final      FinalVerdict
	Truth      string
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
	gate577, err := generation2koideyukawasquarerootconesealaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate577 Koide cone predecessor: %w", err)
	}
	frame := buildFrame()
	transport := buildAzimuthTransport(bundle, frame)
	candidates := auditCandidates(bundle, transport)
	runtime := inheritRuntime(bundle, gate577, transport)
	seal := defineSeal(transport)
	firewalls := auditFirewalls()
	final := compileFinal(transport, candidates, seal)
	truth := "Gate 578 reduces the charged-lepton environmental seal one step further: after the Koide cone fixes theta≈45 degrees, the remaining endpoint orientation is the azimuth phi_e≈257.267 degrees. The angle is stable under v1 transport, but no simple rational/root-of-unity, CKM, PMNS, ASHA projective phase, or native root-trace source is certified."
	return Analysis{Runtime: runtime, Frame: frame, Transport: transport, Candidates: candidates, Seal: seal, Firewalls: firewalls, Final: final, Truth: truth}, nil
}

func inheritRuntime(b historytransport.Bundle, g577 generation2koideyukawasquarerootconesealaudit.Analysis, t AzimuthTransport) RuntimeInheritance {
	return RuntimeInheritance{
		Mu0GeV:          b.EndVector.Mu0GeV,
		Lambda12GeV:     b.GaugeBoundary.Lambda12GeV,
		Gate577QeMZ:     g577.Final.KoideQeMZ,
		Gate577PhiMZDeg: t.MZ.PhiDeg,
		Gate577PhiLDeg:  t.Lambda12.PhiDeg,
		Gate577Verdict:  g577.Final.Verdict,
		Source:          "historytransport.BuildDefault() plus Gate577 Koide cone predecessor",
		Verdict:         StatusRuntimeInherited,
	}
}

func buildFrame() AzimuthFrame {
	n := []float64{1 / math.Sqrt(3), 1 / math.Sqrt(3), 1 / math.Sqrt(3)}
	e1 := []float64{1 / math.Sqrt(2), -1 / math.Sqrt(2), 0}
	e2 := []float64{1 / math.Sqrt(6), 1 / math.Sqrt(6), -2 / math.Sqrt(6)}
	return AzimuthFrame{
		DemocraticAxis: n,
		E1:             e1,
		E2:             e2,
		DotNE1:         dot(n, e1),
		DotNE2:         dot(n, e2),
		DotE1E2:        dot(e1, e2),
		NormN:          norm(n),
		NormE1:         norm(e1),
		NormE2:         norm(e2),
		RightHanded:    dot(cross(e1, e2), n) > 0.999999,
		Convention:     "phi=atan2((x-(x·n)n)·e2,(x-(x·n)n)·e1); reported in signed degrees and [0,360) degrees",
		Verdict:        StatusBasisOrthonormalCertified,
	}
}

func buildAzimuthTransport(b historytransport.Bundle, frame AzimuthFrame) AzimuthTransport {
	labels := []string{"e", "mu", "tau"}
	mz := azimuthPoint("M_Z", labels, values(b.FlavorTransport.YukawaSingularValuesMZ.ChargedLeptons, labels), frame)
	l := azimuthPoint("Lambda_12", labels, values(b.FlavorTransport.YukawaSingularValuesLambda12.ChargedLeptons, labels), frame)
	dphi := signedAngleDelta(l.PhiDeg, mz.PhiDeg)
	verdict := strings.Join([]string{StatusAzimuthComputedMZ, StatusAzimuthComputedLambda12, StatusAzimuthTransportStable, StatusConeCircleDecomposition}, ";")
	return AzimuthTransport{MZ: mz, Lambda12: l, DeltaPhiDeg: dphi, AbsDeltaPhiDeg: math.Abs(dphi), DeltaQ: l.Q - mz.Q, StableAt1eMinus3Deg: math.Abs(dphi) < 1e-3, StableAt1eMinus2Deg: math.Abs(dphi) < 1e-2, Verdict: verdict}
}

func auditCandidates(b historytransport.Bundle, t AzimuthTransport) CandidateAudit {
	drift := math.Max(t.AbsDeltaPhiDeg, 1e-6)
	threshold := DriftCertificationMultiplier * drift
	if threshold < 1e-3 {
		threshold = 1e-3
	}
	phi := t.MZ.PhiDeg
	num, den, dist := nearestRationalTurn(phi, 72)
	cands := []PhaseCandidate{
		candidate("nearest simple rational turn", fmt.Sprintf("%d/%d of full turn", num, den), 360*float64(num)/float64(den), phi, threshold),
		candidate("third-root / SU(3) phase grid", "multiples of 120 degrees", nearestMultiple(phi, 120), phi, threshold),
		candidate("fourth-root / quadrant phase grid", "multiples of 90 degrees", nearestMultiple(phi, 90), phi, threshold),
		candidate("sixth-root phase grid", "multiples of 60 degrees", nearestMultiple(phi, 60), phi, threshold),
		candidate("eighth-root / 45-degree grid", "multiples of 45 degrees", nearestMultiple(phi, 45), phi, threshold),
		candidate("CKM delta phase", "delta_CKM from v1 input, degrees", radToDeg(b.Inputs.CKM.Delta), phi, threshold),
		candidate("CKM supplementary phase", "180+delta_CKM, degrees", normalizeDeg(180+radToDeg(b.Inputs.CKM.Delta)), phi, threshold),
		candidate("CKM complement phase", "360-delta_CKM, degrees", normalizeDeg(360-radToDeg(b.Inputs.CKM.Delta)), phi, threshold),
	}
	any := false
	for _, c := range cands {
		any = any || c.Certified
	}
	verdicts := []string{StatusFiveSeventhsNearButNotCertified, StatusNoSimpleRationalPhaseCertified, StatusNoCKMPhaseIdentification, StatusNoPMNSRuntimeCandidate}
	if any {
		verdicts = []string{"CONDITIONAL_SUPPORT_PHASE_CANDIDATE_WITHIN_GATE578_NUMERICAL_THRESHOLD_REQUIRES_NEW_THEOREM"}
	}
	return CandidateAudit{PointPhiDeg: phi, DriftScaleDeg: drift, CertificationThresholdDeg: threshold, NearestRationalDenominator: den, NearestRationalNumerator: num, NearestRationalTurn: fmt.Sprintf("%d/%d", num, den), NearestRationalDistanceDeg: dist, Candidates: cands, AnyCertified: any, Verdict: strings.Join(verdicts, ";")}
}

func defineSeal(t AzimuthTransport) AzimuthSeal {
	return AzimuthSeal{
		Name:                           "ChargedLeptonKoideAzimuthSeal",
		Carrier:                        "positive charged-lepton square-root Yukawa cone around democratic axis",
		Coordinates:                    []string{"rho_e=||sqrt(Y_e)||", "theta_e≈45 degrees from Koide cone", "phi_e=azimuth around n in {e1,e2} frame"},
		ConeConstraint:                 "Q_e≈2/3 fixes theta_e≈45 degrees",
		AzimuthDefinition:              "phi_e=atan2((x_e-(x_e·n)n)·e2,(x_e-(x_e·n)n)·e1)",
		RadiusMZ:                       t.MZ.Rho,
		PhiMZDeg:                       t.MZ.PhiDeg,
		PhiLambda12Deg:                 t.Lambda12.PhiDeg,
		DriftDeg:                       t.DeltaPhiDeg,
		OriginalPositiveMagnitudes:     3,
		Constraints:                    1,
		RemainingContinuousCoordinates: 2,
		NativeDerivation:               false,
		BridgeOnly:                     true,
		Verdict:                        strings.Join([]string{StatusAzimuthOrientationSealCandidate, StatusConeCircleDecomposition}, ";"),
	}
}

func auditFirewalls() FirewallAudit {
	return FirewallAudit{DerivesChargedLeptonMasses: false, DerivesYukawaEigenvalues: false, DerivesCKM: false, DerivesPMNS: false, DerivesGenerationHierarchy: false, IdentifiesWithASHAProjectivePhase: false, ImportsObservedAsNative: false, AddsNewCarrier: false, PreservesGate352: true, Verdict: strings.Join([]string{StatusGate352ObstructionInherited, StatusNoNativeAzimuthDerivation, StatusNoFlavorPromotion, StatusObservedEndpointQuarantined, StatusGate578BoundaryPreserved}, ";")}
}

func compileFinal(t AzimuthTransport, c CandidateAudit, s AzimuthSeal) FinalVerdict {
	return FinalVerdict{SealName: s.Name, PhiMZDeg: t.MZ.PhiDeg, PhiSignedMZDeg: t.MZ.PhiSignedDeg, PhiLambda12Deg: t.Lambda12.PhiDeg, DeltaPhiDeg: t.DeltaPhiDeg, NearestSimpleCandidate: c.NearestRationalTurn + " turn", NearestSimpleDistanceDeg: c.NearestRationalDistanceDeg, CertifiedSimplePhase: c.AnyCertified, NativeASHAFlavorDerivation: false, NextRequiredTheorem: "a native root-trace/phase operator or functor that produces phi_e on the Koide cone while preserving Gate352 and flavor firewalls", Verdict: strings.Join([]string{StatusAzimuthOrientationSealCandidate, StatusNoSimpleRationalPhaseCertified, StatusNoNativeAzimuthDerivation, StatusGate578BoundaryPreserved}, ";")}
}

func azimuthPoint(scale string, labels []string, y []float64, frame AzimuthFrame) AzimuthPoint {
	root := make([]float64, len(y))
	sumY, rootSum := 0.0, 0.0
	for i, v := range y {
		root[i] = math.Sqrt(v)
		sumY += v
		rootSum += root[i]
	}
	rho := math.Sqrt(sumY)
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
	q := sumY / (rootSum * rootSum)
	theta := radToDeg(math.Atan2(norm(perp), parallel))
	verdict := StatusAzimuthComputedMZ
	if scale == "Lambda_12" {
		verdict = StatusAzimuthComputedLambda12
	}
	return AzimuthPoint{Scale: scale, Labels: append([]string(nil), labels...), Yukawas: append([]float64(nil), y...), RootVector: root, Rho: rho, Q: q, DeltaFromTwoThirds: q - KoideTarget, ThetaDeg: theta, ThetaDeltaDeg: theta - 45, Parallel: parallel, PerpendicularNorm: norm(perp), CoordE1: ce1, CoordE2: ce2, PhiSignedDeg: signed, PhiDeg: phi, Verdict: verdict}
}

func values(m map[string]float64, labels []string) []float64 {
	out := make([]float64, len(labels))
	for i, k := range labels {
		out[i] = m[k]
	}
	return out
}

func candidate(name, formula string, angle, phi, threshold float64) PhaseCandidate {
	d := angularDistanceDeg(phi, angle)
	certified := d <= threshold
	verdict := "FAILED_ROUTE_PHASE_CANDIDATE_OUTSIDE_GATE578_CERTIFICATION_THRESHOLD"
	if certified {
		verdict = "CONDITIONAL_SUPPORT_PHASE_CANDIDATE_WITHIN_GATE578_NUMERICAL_THRESHOLD_REQUIRES_NEW_THEOREM"
	}
	return PhaseCandidate{Name: name, Formula: formula, CandidateDeg: normalizeDeg(angle), DistanceDeg: d, ThresholdDeg: threshold, Certified: certified, Verdict: verdict}
}

func nearestRationalTurn(phi float64, maxDen int) (num, den int, dist float64) {
	bestN, bestD := 0, 1
	best := math.Inf(1)
	for d := 1; d <= maxDen; d++ {
		for n := 0; n < d; n++ {
			angle := 360 * float64(n) / float64(d)
			if dd := angularDistanceDeg(phi, angle); dd < best {
				best, bestN, bestD = dd, n, d
			}
		}
	}
	g := gcd(bestN, bestD)
	if g > 0 {
		bestN /= g
		bestD /= g
	}
	return bestN, bestD, best
}

func nearestMultiple(phi, step float64) float64 {
	return normalizeDeg(math.Round(phi/step) * step)
}

func normalizeDeg(x float64) float64 {
	y := math.Mod(x, 360)
	if y < 0 {
		y += 360
	}
	return y
}

func angularDistanceDeg(a, b float64) float64 {
	d := math.Abs(normalizeDeg(a) - normalizeDeg(b))
	if d > 180 {
		d = 360 - d
	}
	return d
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

func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
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

func Statuses() []string {
	return []string{StatusRuntimeInherited, StatusBasisOrthonormalCertified, StatusAzimuthComputedMZ, StatusAzimuthComputedLambda12, StatusAzimuthTransportStable, StatusConeCircleDecomposition, StatusFiveSeventhsNearButNotCertified, StatusAzimuthOrientationSealCandidate, StatusNoSimpleRationalPhaseCertified, StatusNoCKMPhaseIdentification, StatusNoPMNSRuntimeCandidate, StatusGate352ObstructionInherited, StatusNoNativeAzimuthDerivation, StatusNoFlavorPromotion, StatusObservedEndpointQuarantined, StatusGate578BoundaryPreserved}
}
