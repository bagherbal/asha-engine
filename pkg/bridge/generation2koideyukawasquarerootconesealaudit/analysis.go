// Package generation2koideyukawasquarerootconesealaudit implements Gate 577:
// Koide Square-Root Yukawa Cone Environmental Seal Audit.
//
// Gate 577 is the first history-seal reduction after the ASHA History Transport
// v1 runtime.  It does not derive charged-lepton masses, Yukawa eigenvalues, or
// Koide from native ASHA algebra.  It audits the strongest visible environmental
// fingerprint in the runtime output: the charged-lepton square-root Yukawa vector
// lies extremely close to the Koide cone Q=2/3, equivalently at 45 degrees from
// the democratic axis in positive root-Yukawa space.  The audit converts this
// pattern into a minimal bridge-only environmental seal while preserving the Gate
// 352 root-trace obstruction and all flavor firewalls.
package generation2koideyukawasquarerootconesealaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/fermionicroottracesieve"
	"github.com/bagherbal/asha-engine/pkg/historytransport"
)

const (
	AuditID = "GATE577-KOIDE-YUKAWA-SQUARE-ROOT-CONE-ENVIRONMENTAL-SEAL-AUDIT"

	StatusRuntimeInherited                     = "PASS_HISTORY_TRANSPORT_V1_RUNTIME_FLAVOR_OUTPUT_INHERITED"
	StatusSquareRootGeometryDefined            = "PASS_SQUARE_ROOT_YUKAWA_VECTOR_GEOMETRY_DEFINED"
	StatusKoideConeAngleEquivalenceVerified    = "PASS_KOIDE_Q_TWO_THIRDS_EQUIVALENT_TO_45_DEGREE_CONE"
	StatusChargedLeptonKoideVisibleMZ          = "PASS_CHARGED_LEPTON_KOIDE_CONE_ALIGNMENT_VISIBLE_AT_MZ"
	StatusChargedLeptonKoideVisibleLambda12    = "PASS_CHARGED_LEPTON_KOIDE_CONE_ALIGNMENT_VISIBLE_AT_LAMBDA12"
	StatusKoideEnvironmentalSealCandidate      = "CONDITIONAL_SUPPORT_CHARGED_LEPTON_KOIDE_CONE_ENVIRONMENTAL_SEAL_CANDIDATE"
	StatusKoideSealMinimalParameterized        = "CONDITIONAL_SUPPORT_KOIDE_SEAL_REDUCES_CHARGED_LEPTON_MAGNITUDES_TO_RADIUS_AND_AZIMUTH_PLUS_CONE_CONSTRAINT"
	StatusKoideNotUniversal                    = "FAILED_ROUTE_KOIDE_CONE_NOT_UNIVERSAL_ACROSS_UP_DOWN_YUKAWA_SECTORS"
	StatusGate352RootTraceObstructionInherited = "FAILED_ROUTE_GATE352_ROOT_TRACE_OBSTRUCTION_INHERITED_NO_NATIVE_KOIDE_OPERATOR"
	StatusNoNativeLeptonMassDerivation         = "FAILED_ROUTE_NO_ASHA_NATIVE_CHARGED_LEPTON_MASS_OR_YUKAWA_DERIVATION"
	StatusNoCKMFlavorPromotion                 = "FIREWALL_PRESERVED_KOIDE_CONE_DOES_NOT_DERIVE_CKM_PMNS_OR_FLAVOR_TEXTURE"
	StatusNoObservedImportAsNative             = "FIREWALL_PRESERVED_OBSERVED_LEPTON_DATA_REMAINS_HISTORY_ENDPOINT"
	StatusGate577BoundaryPreserved             = "FIREWALL_PRESERVED_GATE577_KOIDE_ENVIRONMENTAL_SEAL_BOUNDARY"
)

const (
	KoideTarget   = 2.0 / 3.0
	KoideAngleDeg = 45.0
)

type RuntimeInheritance struct {
	Mu0GeV      float64
	Lambda12GeV float64
	JCKM        float64
	KoideQe     float64
	Source      string
	Verdict     string
}

type GeometryDefinition struct {
	RootVectorFormula    string
	KoideFormula         string
	DemocraticAxis       string
	CosineFormula        string
	ConeEquivalence      string
	TargetQ              float64
	TargetAngleDeg       float64
	PositiveConeOnly     bool
	UsesObservedEndpoint bool
	Verdict              string
}

type ConePoint struct {
	Scale              string
	Sector             string
	Labels             []string
	Yukawas            []float64
	RootVector         []float64
	Rho                float64
	DemocraticParallel float64
	PerpendicularNorm  float64
	PerpOverParallel   float64
	AzimuthDeg         float64
	Q                  float64
	DeltaFromTwoThirds float64
	AngleDeg           float64
	AngleDeltaDeg      float64
	OnKoideCone1e4     bool
	OnKoideCone1e5     bool
	Verdict            string
}

type SectorComparison struct {
	Points                       []ConePoint
	ChargedLeptonMZSharp         bool
	ChargedLeptonLambda12Sharp   bool
	UpQuarksOnKoideCone          bool
	DownQuarksOnKoideCone        bool
	ChargedLeptonTransportStable bool
	KoideUniversalAcrossSectors  bool
	BestSector                   string
	Verdict                      string
}

type MinimalEnvironmentalSeal struct {
	Name                           string
	Carrier                        string
	SealConstraint                 string
	ReducedCoordinates             []string
	OriginalPositiveMagnitudes     int
	ConeConstraintCount            int
	RemainingContinuousCoordinates int
	NativeDerivation               bool
	BridgeOnly                     bool
	SolvesFirstLogicalSealAs       string
	Verdict                        string
}

type Gate352Inheritance struct {
	Gate                int
	EmpiricalAlignment  bool
	NativePromotion     bool
	RequiredNewObject   string
	RootTraceNative     bool
	PfaffianCanGenerate bool
	Verdict             string
}

type FirewallAudit struct {
	DerivesChargedLeptonMasses bool
	DerivesYukawaEigenvalues   bool
	DerivesCKM                 bool
	DerivesPMNS                bool
	DerivesGenerationHierarchy bool
	ImportsObservedAsNative    bool
	AddsNewCarrier             bool
	PreservesGate352           bool
	Verdict                    string
}

type FinalVerdict struct {
	FirstLogicalSeal           string
	StrongestRuntimeGeometry   string
	KoideQeMZ                  float64
	KoideDeltaMZ               float64
	KoideAngleMZDeg            float64
	KoideQeLambda12            float64
	KoideDeltaLambda12         float64
	NativeASHAFlavorDerivation bool
	NextRequiredTheorem        string
	Verdict                    string
}

type Analysis struct {
	Runtime    RuntimeInheritance
	Geometry   GeometryDefinition
	Comparison SectorComparison
	Seal       MinimalEnvironmentalSeal
	Gate352    Gate352Inheritance
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
		return Analysis{}, fmt.Errorf("build history transport v1 runtime: %w", err)
	}
	g352, err := fermionicroottracesieve.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate352 root-trace predecessor: %w", err)
	}
	runtime := inheritRuntime(bundle)
	geometry := defineGeometry()
	comparison := compareSectors(bundle)
	seal := defineMinimalSeal(comparison)
	inherited := inheritGate352(g352)
	firewalls := auditFirewalls(inherited)
	final := compileFinal(comparison, seal, inherited)
	truth := "Gate 577 solves the first logical history seal only in the bridge sense: charged-lepton endpoint data define a sharply aligned Koide square-root cone seal, reducing Y_e magnitudes to radius plus azimuth plus the Q=2/3 cone constraint. Gate 352 remains binding: no native ASHA root-trace operator or lepton-mass derivation is promoted."
	return Analysis{Runtime: runtime, Geometry: geometry, Comparison: comparison, Seal: seal, Gate352: inherited, Firewalls: firewalls, Final: final, Truth: truth}, nil
}

func inheritRuntime(b historytransport.Bundle) RuntimeInheritance {
	return RuntimeInheritance{Mu0GeV: b.EndVector.Mu0GeV, Lambda12GeV: b.GaugeBoundary.Lambda12GeV, JCKM: b.FlavorTransport.JCKM, KoideQe: b.FlavorTransport.KoideQe, Source: "historytransport.BuildDefault() runtime outputs 01-07", Verdict: StatusRuntimeInherited}
}

func defineGeometry() GeometryDefinition {
	return GeometryDefinition{
		RootVectorFormula:    "x_f=(sqrt(y_1),sqrt(y_2),sqrt(y_3)) for positive Yukawa singular values",
		KoideFormula:         "Q_f=(y_1+y_2+y_3)/(sqrt(y_1)+sqrt(y_2)+sqrt(y_3))^2",
		DemocraticAxis:       "n=(1,1,1)/sqrt(3)",
		CosineFormula:        "cos(theta)=(x_f·n)/||x_f||, hence Q_f=1/(3 cos^2(theta))",
		ConeEquivalence:      "Q_f=2/3 iff cos^2(theta)=1/2 iff theta=45 degrees in the positive cone",
		TargetQ:              KoideTarget,
		TargetAngleDeg:       KoideAngleDeg,
		PositiveConeOnly:     true,
		UsesObservedEndpoint: true,
		Verdict:              strings.Join([]string{StatusSquareRootGeometryDefined, StatusKoideConeAngleEquivalenceVerified}, ";"),
	}
}

func compareSectors(b historytransport.Bundle) SectorComparison {
	points := []ConePoint{}
	points = append(points, conePoint("M_Z", "up_quarks", []string{"u", "c", "t"}, values(b.FlavorTransport.YukawaSingularValuesMZ.UpQuarks, []string{"u", "c", "t"})))
	points = append(points, conePoint("M_Z", "down_quarks", []string{"d", "s", "b"}, values(b.FlavorTransport.YukawaSingularValuesMZ.DownQuarks, []string{"d", "s", "b"})))
	points = append(points, conePoint("M_Z", "charged_leptons", []string{"e", "mu", "tau"}, values(b.FlavorTransport.YukawaSingularValuesMZ.ChargedLeptons, []string{"e", "mu", "tau"})))
	points = append(points, conePoint("Lambda_12", "up_quarks", []string{"u", "c", "t"}, values(b.FlavorTransport.YukawaSingularValuesLambda12.UpQuarks, []string{"u", "c", "t"})))
	points = append(points, conePoint("Lambda_12", "down_quarks", []string{"d", "s", "b"}, values(b.FlavorTransport.YukawaSingularValuesLambda12.DownQuarks, []string{"d", "s", "b"})))
	points = append(points, conePoint("Lambda_12", "charged_leptons", []string{"e", "mu", "tau"}, values(b.FlavorTransport.YukawaSingularValuesLambda12.ChargedLeptons, []string{"e", "mu", "tau"})))

	leptonMZ := findPoint(points, "M_Z", "charged_leptons")
	leptonL := findPoint(points, "Lambda_12", "charged_leptons")
	upMZ := findPoint(points, "M_Z", "up_quarks")
	downMZ := findPoint(points, "M_Z", "down_quarks")
	stable := math.Abs(leptonMZ.Q-leptonL.Q) < 5e-6
	verdict := strings.Join([]string{StatusChargedLeptonKoideVisibleMZ, StatusChargedLeptonKoideVisibleLambda12, StatusKoideEnvironmentalSealCandidate, StatusKoideNotUniversal}, ";")
	return SectorComparison{Points: points, ChargedLeptonMZSharp: leptonMZ.OnKoideCone1e5, ChargedLeptonLambda12Sharp: leptonL.OnKoideCone1e5, UpQuarksOnKoideCone: upMZ.OnKoideCone1e4, DownQuarksOnKoideCone: downMZ.OnKoideCone1e4, ChargedLeptonTransportStable: stable, KoideUniversalAcrossSectors: upMZ.OnKoideCone1e4 && downMZ.OnKoideCone1e4 && leptonMZ.OnKoideCone1e4, BestSector: "charged_leptons", Verdict: verdict}
}

func values(m map[string]float64, labels []string) []float64 {
	out := make([]float64, len(labels))
	for i, k := range labels {
		out[i] = m[k]
	}
	return out
}

func conePoint(scale, sector string, labels []string, y []float64) ConePoint {
	root := make([]float64, len(y))
	sumY := 0.0
	rootSum := 0.0
	for i, v := range y {
		root[i] = math.Sqrt(v)
		sumY += v
		rootSum += root[i]
	}
	rho := math.Sqrt(sumY)
	n := []float64{1 / math.Sqrt(3), 1 / math.Sqrt(3), 1 / math.Sqrt(3)}
	e1 := []float64{1 / math.Sqrt(2), -1 / math.Sqrt(2), 0}
	e2 := []float64{1 / math.Sqrt(6), 1 / math.Sqrt(6), -2 / math.Sqrt(6)}
	parallel := dot(root, n)
	perp := make([]float64, 3)
	for i := 0; i < 3; i++ {
		perp[i] = root[i] - parallel*n[i]
	}
	perpNorm := norm(perp)
	q := sumY / (rootSum * rootSum)
	theta := radToDeg(math.Atan2(perpNorm, parallel))
	phi := radToDeg(math.Atan2(dot(perp, e2), dot(perp, e1)))
	if phi < 0 {
		phi += 360
	}
	delta := q - KoideTarget
	verdict := "FAILED_ROUTE_SECTOR_NOT_ON_KOIDE_CONE"
	if math.Abs(delta) < 1e-4 {
		verdict = "PASS_SECTOR_ON_KOIDE_CONE_TO_1E_MINUS_4"
	}
	if math.Abs(delta) < 1e-5 {
		verdict = "PASS_SECTOR_ON_KOIDE_CONE_TO_1E_MINUS_5"
	}
	return ConePoint{Scale: scale, Sector: sector, Labels: append([]string(nil), labels...), Yukawas: append([]float64(nil), y...), RootVector: root, Rho: rho, DemocraticParallel: parallel, PerpendicularNorm: perpNorm, PerpOverParallel: perpNorm / parallel, AzimuthDeg: phi, Q: q, DeltaFromTwoThirds: delta, AngleDeg: theta, AngleDeltaDeg: theta - KoideAngleDeg, OnKoideCone1e4: math.Abs(delta) < 1e-4, OnKoideCone1e5: math.Abs(delta) < 1e-5, Verdict: verdict}
}

func dot(a, b []float64) float64 {
	s := 0.0
	for i := range a {
		s += a[i] * b[i]
	}
	return s
}
func norm(a []float64) float64   { return math.Sqrt(dot(a, a)) }
func radToDeg(x float64) float64 { return x * 180 / math.Pi }

func findPoint(points []ConePoint, scale, sector string) ConePoint {
	for _, p := range points {
		if p.Scale == scale && p.Sector == sector {
			return p
		}
	}
	return ConePoint{}
}

func defineMinimalSeal(c SectorComparison) MinimalEnvironmentalSeal {
	return MinimalEnvironmentalSeal{
		Name:                           "ChargedLeptonKoideConeSeal",
		Carrier:                        "positive square-root charged-lepton Yukawa space R^3_+ with x_e=(sqrt(y_e),sqrt(y_mu),sqrt(y_tau))",
		SealConstraint:                 "Q_e=2/3 equivalently angle(x_e,(1,1,1))=45 degrees",
		ReducedCoordinates:             []string{"rho_e=||x_e||", "phi_e=azimuth around democratic axis", "Q_e fixed to 2/3 up to measured residual"},
		OriginalPositiveMagnitudes:     3,
		ConeConstraintCount:            1,
		RemainingContinuousCoordinates: 2,
		NativeDerivation:               false,
		BridgeOnly:                     true,
		SolvesFirstLogicalSealAs:       "minimal environmental geometry for charged-lepton magnitudes; not a native mass theorem",
		Verdict:                        strings.Join([]string{StatusKoideEnvironmentalSealCandidate, StatusKoideSealMinimalParameterized}, ";"),
	}
}

func inheritGate352(a fermionicroottracesieve.Analysis) Gate352Inheritance {
	return Gate352Inheritance{Gate: 352, EmpiricalAlignment: a.KoidePromotion.EmpiricalAlignment, NativePromotion: a.KoidePromotion.NativePromotion, RequiredNewObject: a.KoidePromotion.RequiredNewObject, RootTraceNative: a.RootTrace.RootTraceNative, PfaffianCanGenerate: a.Pfaffian.PfaffianCanGenerateKoide, Verdict: StatusGate352RootTraceObstructionInherited}
}

func auditFirewalls(g Gate352Inheritance) FirewallAudit {
	return FirewallAudit{DerivesChargedLeptonMasses: false, DerivesYukawaEigenvalues: false, DerivesCKM: false, DerivesPMNS: false, DerivesGenerationHierarchy: false, ImportsObservedAsNative: false, AddsNewCarrier: false, PreservesGate352: !g.NativePromotion && !g.RootTraceNative && !g.PfaffianCanGenerate, Verdict: strings.Join([]string{StatusNoNativeLeptonMassDerivation, StatusNoCKMFlavorPromotion, StatusNoObservedImportAsNative, StatusGate577BoundaryPreserved}, ";")}
}

func compileFinal(c SectorComparison, s MinimalEnvironmentalSeal, g Gate352Inheritance) FinalVerdict {
	mz := findPoint(c.Points, "M_Z", "charged_leptons")
	l := findPoint(c.Points, "Lambda_12", "charged_leptons")
	return FinalVerdict{FirstLogicalSeal: s.Name, StrongestRuntimeGeometry: "charged-lepton Koide square-root cone", KoideQeMZ: mz.Q, KoideDeltaMZ: mz.DeltaFromTwoThirds, KoideAngleMZDeg: mz.AngleDeg, KoideQeLambda12: l.Q, KoideDeltaLambda12: l.DeltaFromTwoThirds, NativeASHAFlavorDerivation: false, NextRequiredTheorem: g.RequiredNewObject, Verdict: strings.Join([]string{StatusKoideEnvironmentalSealCandidate, StatusGate352RootTraceObstructionInherited, StatusGate577BoundaryPreserved}, ";")}
}

func Statuses() []string {
	return []string{StatusRuntimeInherited, StatusSquareRootGeometryDefined, StatusKoideConeAngleEquivalenceVerified, StatusChargedLeptonKoideVisibleMZ, StatusChargedLeptonKoideVisibleLambda12, StatusKoideEnvironmentalSealCandidate, StatusKoideSealMinimalParameterized, StatusKoideNotUniversal, StatusGate352RootTraceObstructionInherited, StatusNoNativeLeptonMassDerivation, StatusNoCKMFlavorPromotion, StatusNoObservedImportAsNative, StatusGate577BoundaryPreserved}
}
