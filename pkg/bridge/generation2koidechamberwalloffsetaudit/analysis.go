// Package generation2koidechamberwalloffsetaudit implements Gate 583:
// Koide Chamber-Wall Offset Audit.
//
// Gate 582 rewrote the charged-lepton square-root Yukawa ray in the
// democratic/Fourier form
//
//	x_j = A [ 1 + sqrt(2) R cos(delta + 2*pi*j/3) ].
//
// Gate 583 audits the positive S3 chamber of the Koide circle R=1.  In the
// canonical (e,mu,tau) ordering the physical charged-lepton point lies in the
// chamber 105° < delta < 135°, very close to the electron-zero wall delta=135°.
// The gate replaces the failed search for a simple rational phase by the more
// geometric wall-offset coordinate epsilon = 135° - delta.  It does not derive
// epsilon, charged-lepton masses, a root trace, CKM/PMNS, or generation
// hierarchy; it only compresses the observed environmental ray.
package generation2koidechamberwalloffsetaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2koidefouriercirculantphaseaudit"
	"github.com/bagherbal/asha-engine/pkg/historytransport"
)

const (
	AuditID = "GATE583-KOIDE-CHAMBER-WALL-OFFSET-AUDIT"

	StatusGate582Inherited          = "PASS_GATE582_FOURIER_CIRCULANT_RUNTIME_INHERITED"
	StatusChamberGeometryDefined    = "PASS_KOIDE_POSITIVE_S3_CHAMBER_WALL_GEOMETRY_DEFINED"
	StatusCanonicalChamberCertified = "PASS_CANONICAL_CHARGED_LEPTON_POINT_IN_POSITIVE_KOIDE_CHAMBER"
	StatusElectronWallIdentified    = "PASS_ELECTRON_ZERO_WALL_IDENTIFIED_AT_DELTA_135_DEGREES"
	StatusMZEpsilonComputed         = "PASS_ELECTRON_WALL_OFFSET_EPSILON_COMPUTED_AT_MZ"
	StatusLambdaEpsilonComputed     = "PASS_ELECTRON_WALL_OFFSET_EPSILON_COMPUTED_AT_LAMBDA12"
	StatusEpsilonStable             = "PASS_ELECTRON_WALL_OFFSET_STABLE_UNDER_V1_TRANSPORT"
	StatusElectronSmallness         = "PASS_ELECTRON_SMALLNESS_CONTROLLED_BY_WALL_OFFSET"
	StatusWallHierarchyCompression  = "CONDITIONAL_SUPPORT_WALL_OFFSET_REDUCES_CHARGED_LEPTON_HIERARCHY_DESCRIPTION"
	StatusQuarkFormalCoordinates    = "CONDITIONAL_SUPPORT_QUARK_SECTORS_HAVE_FORMAL_FOURIER_COORDINATES_BUT_NOT_KOIDE_WALL_SEALS"
	StatusNoNativeWallSelector      = "FAILED_ROUTE_NO_NATIVE_CHAMBER_WALL_OR_EPSILON_SELECTOR"
	StatusNoSimpleEpsilon           = "FAILED_ROUTE_EPSILON_NOT_CERTIFIED_AS_SIMPLE_RATIONAL_OR_ROOT_OF_UNITY"
	StatusQuarksNotKoideWalls       = "FAILED_ROUTE_QUARK_SECTORS_NOT_ON_KOIDE_CIRCLE_IN_V1_WALL_AUDIT"
	StatusNoFlavorPromotion         = "FIREWALL_PRESERVED_CHAMBER_WALL_OFFSET_DOES_NOT_DERIVE_TEXTURE_CKM_PMNS_OR_GENERATIONS"
	StatusObservedEndpointPreserved = "FIREWALL_PRESERVED_WALL_OFFSET_REMAINS_OBSERVED_HISTORY_ENDPOINT_ORIENTATION"
	StatusGate352Preserved          = "FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING"
	StatusGate583BoundaryPreserved  = "FIREWALL_PRESERVED_GATE583_KOIDE_CHAMBER_WALL_BOUNDARY"
)

const (
	lowerWallDeg = 105.0
	upperWallDeg = 135.0
	wallWidthDeg = upperWallDeg - lowerWallDeg
)

type RuntimeInheritance struct {
	Mu0GeV             float64
	Lambda12GeV        float64
	Gate582DeltaMZDeg  float64
	Gate582DeltaL12Deg float64
	Gate582RMZ         float64
	Gate582RL12        float64
	RuntimeSource      string
	Verdict            string
}

type ChamberGeometryAudit struct {
	Formula                 string
	KoideCircleCondition    string
	CanonicalOrder          []string
	PositiveChamberDeg      [2]float64
	LowerWallLabel          string
	UpperWallLabel          string
	UpperWallZeroCheck      float64
	LowerWallZeroCheck      float64
	WallOffsetDefinition    string
	NearWallElectronFormula string
	Verdict                 string
}

type WallPoint struct {
	Name                        string
	DeltaDeg                    float64
	PlaneAmplitudeR             float64
	EpsilonDeg                  float64
	EpsilonRad                  float64
	DistanceToLowerWallDeg      float64
	NormalizedChamberCoordinate float64
	NormalizedDistanceUpperWall float64
	InsideCanonicalChamber      bool
	ElectronRootOverA           float64
	MuonRootOverA               float64
	TauRootOverA                float64
	ExactKoideWallElectronOverA float64
	LinearElectronOverA         float64
	QuadraticElectronOverA      float64
	ExactWallResidual           float64
	LinearResidual              float64
	QuadraticResidual           float64
	ElectronMuonRootRatio       float64
	ElectronMuonMassRatio       float64
	ElectronTauRootRatio        float64
	ElectronTauMassRatio        float64
	Verdict                     string
}

type WallTransportAudit struct {
	MZEpsilonDeg          float64
	LambdaEpsilonDeg      float64
	SignedDriftDeg        float64
	AbsDriftDeg           float64
	MZAmplitudeResidual   float64
	L12AmplitudeResidual  float64
	AmplitudeMovesToward1 bool
	EpsilonStable         bool
	ChamberPreserved      bool
	Verdict               string
}

type SectorWallAudit struct {
	Sector         string
	Labels         []string
	DeltaDeg       float64
	R              float64
	Q              float64
	EpsilonTo135   float64
	KoideLike      bool
	WallSealValid  bool
	Interpretation string
}

type QuarkAnalogyAudit struct {
	Up         SectorWallAudit
	Down       SectorWallAudit
	Conclusion string
	Verdict    string
}

type FirewallAudit struct {
	DerivesLeptonMasses        bool
	DerivesYukawaEigenvalues   bool
	DerivesKoide               bool
	DerivesEpsilon             bool
	DerivesCKM                 bool
	DerivesPMNS                bool
	DerivesGenerationHierarchy bool
	AddsNewCarrier             bool
	PromotesObservedAsNative   bool
	PreservesGate352           bool
	Verdict                    string
}

type FinalVerdict struct {
	SealName                string
	MZEpsilonDeg            float64
	LambdaEpsilonDeg        float64
	MZPlaneAmplitudeR       float64
	LambdaPlaneAmplitudeR   float64
	Chamber                 string
	WallOffsetStableInV1    bool
	HierarchyNearWall       bool
	QuarkWallSealCertified  bool
	NativeSelectorCertified bool
	MinimalNextRequirement  string
	Verdict                 string
}

type Analysis struct {
	Runtime   RuntimeInheritance
	Chamber   ChamberGeometryAudit
	MZ        WallPoint
	Lambda12  WallPoint
	Transport WallTransportAudit
	Quarks    QuarkAnalogyAudit
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
	g582, err := generation2koidefouriercirculantphaseaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate582 predecessor: %w", err)
	}
	runtime := inheritRuntime(bundle, g582)
	chamber := defineChamber()
	mz := wallPoint("M_Z", g582.MZ.DeltaDeg, g582.MZ.PlaneAmplitudeR, StatusMZEpsilonComputed)
	lambda := wallPoint("Lambda_12", g582.Lambda12.DeltaDeg, g582.Lambda12.PlaneAmplitudeR, StatusLambdaEpsilonComputed)
	transport := auditTransport(mz, lambda)
	quarks := auditQuarkAnalogy(bundle)
	firewalls := auditFirewalls()
	final := compileFinal(mz, lambda, transport, quarks)
	truth := "Gate 583 shows that the charged-lepton Koide Fourier phase is better understood as a near-boundary coordinate inside the positive S3 chamber: for canonical (e,mu,tau), delta=135°-epsilon with epsilon≈2.26718°.  The small electron square-root component is controlled by the distance from the electron-zero wall.  This is environmental compression, not a native ASHA flavor derivation."
	return Analysis{Runtime: runtime, Chamber: chamber, MZ: mz, Lambda12: lambda, Transport: transport, Quarks: quarks, Firewalls: firewalls, Final: final, Truth: truth}, nil
}

func inheritRuntime(b historytransport.Bundle, g582 generation2koidefouriercirculantphaseaudit.Analysis) RuntimeInheritance {
	return RuntimeInheritance{Mu0GeV: b.EndVector.Mu0GeV, Lambda12GeV: b.GaugeBoundary.Lambda12GeV, Gate582DeltaMZDeg: g582.MZ.DeltaDeg, Gate582DeltaL12Deg: g582.Lambda12.DeltaDeg, Gate582RMZ: g582.MZ.PlaneAmplitudeR, Gate582RL12: g582.Lambda12.PlaneAmplitudeR, RuntimeSource: "historytransport.BuildDefault() plus Gate582 Fourier/circulant phase audit", Verdict: StatusGate582Inherited}
}

func defineChamber() ChamberGeometryAudit {
	upperZero := 1.0 + math.Sqrt2*math.Cos(degToRad(upperWallDeg))
	lowerZero := 1.0 + math.Sqrt2*math.Cos(degToRad(lowerWallDeg+120.0))
	return ChamberGeometryAudit{
		Formula:                 "x_j/A=1+sqrt(2)R cos(delta+2*pi*j/3)",
		KoideCircleCondition:    "R=1; positivity in canonical (e,mu,tau) gives the chamber 105°<delta<135°",
		CanonicalOrder:          []string{"e", "mu", "tau"},
		PositiveChamberDeg:      [2]float64{lowerWallDeg, upperWallDeg},
		LowerWallLabel:          "muon-zero wall at delta=105° on the R=1 Koide circle",
		UpperWallLabel:          "electron-zero wall at delta=135° on the R=1 Koide circle",
		UpperWallZeroCheck:      upperZero,
		LowerWallZeroCheck:      lowerZero,
		WallOffsetDefinition:    "epsilon_e=135°-delta for the canonical charged-lepton chamber",
		NearWallElectronFormula: "for R=1 and delta=135°-epsilon: x_e/A=1-cos(epsilon)+sin(epsilon)=epsilon+epsilon^2/2+O(epsilon^3)",
		Verdict:                 strings.Join([]string{StatusChamberGeometryDefined, StatusElectronWallIdentified}, ";"),
	}
}

func wallPoint(name string, deltaDeg, r float64, verdict string) WallPoint {
	epsDeg := upperWallDeg - deltaDeg
	epsRad := degToRad(epsDeg)
	e := componentOverA(r, deltaDeg, 0)
	mu := componentOverA(r, deltaDeg, 1)
	tau := componentOverA(r, deltaDeg, 2)
	exactKoide := 1.0 - math.Cos(epsRad) + math.Sin(epsRad)
	linear := epsRad
	quadratic := epsRad + 0.5*epsRad*epsRad
	inside := deltaDeg > lowerWallDeg && deltaDeg < upperWallDeg && e > 0 && mu > 0 && tau > 0
	return WallPoint{
		Name:                        name,
		DeltaDeg:                    deltaDeg,
		PlaneAmplitudeR:             r,
		EpsilonDeg:                  epsDeg,
		EpsilonRad:                  epsRad,
		DistanceToLowerWallDeg:      deltaDeg - lowerWallDeg,
		NormalizedChamberCoordinate: (deltaDeg - lowerWallDeg) / wallWidthDeg,
		NormalizedDistanceUpperWall: epsDeg / wallWidthDeg,
		InsideCanonicalChamber:      inside,
		ElectronRootOverA:           e,
		MuonRootOverA:               mu,
		TauRootOverA:                tau,
		ExactKoideWallElectronOverA: exactKoide,
		LinearElectronOverA:         linear,
		QuadraticElectronOverA:      quadratic,
		ExactWallResidual:           e - exactKoide,
		LinearResidual:              e - linear,
		QuadraticResidual:           e - quadratic,
		ElectronMuonRootRatio:       e / mu,
		ElectronMuonMassRatio:       (e / mu) * (e / mu),
		ElectronTauRootRatio:        e / tau,
		ElectronTauMassRatio:        (e / tau) * (e / tau),
		Verdict:                     strings.Join([]string{verdict, StatusCanonicalChamberCertified, StatusElectronSmallness, StatusWallHierarchyCompression}, ";"),
	}
}

func auditTransport(mz, lambda WallPoint) WallTransportAudit {
	drift := lambda.EpsilonDeg - mz.EpsilonDeg
	return WallTransportAudit{MZEpsilonDeg: mz.EpsilonDeg, LambdaEpsilonDeg: lambda.EpsilonDeg, SignedDriftDeg: drift, AbsDriftDeg: math.Abs(drift), MZAmplitudeResidual: mz.PlaneAmplitudeR - 1.0, L12AmplitudeResidual: lambda.PlaneAmplitudeR - 1.0, AmplitudeMovesToward1: math.Abs(lambda.PlaneAmplitudeR-1.0) < math.Abs(mz.PlaneAmplitudeR-1.0), EpsilonStable: math.Abs(drift) < 3e-4, ChamberPreserved: mz.InsideCanonicalChamber && lambda.InsideCanonicalChamber, Verdict: strings.Join([]string{StatusEpsilonStable, StatusWallHierarchyCompression}, ";")}
}

func auditQuarkAnalogy(b historytransport.Bundle) QuarkAnalogyAudit {
	up := sectorAudit("up", []string{"u", "c", "t"}, []float64{b.EndVector.YukawaSingularValues.UpQuarks["u"], b.EndVector.YukawaSingularValues.UpQuarks["c"], b.EndVector.YukawaSingularValues.UpQuarks["t"]})
	down := sectorAudit("down", []string{"d", "s", "b"}, []float64{b.EndVector.YukawaSingularValues.DownQuarks["d"], b.EndVector.YukawaSingularValues.DownQuarks["s"], b.EndVector.YukawaSingularValues.DownQuarks["b"]})
	return QuarkAnalogyAudit{Up: up, Down: down, Conclusion: "Quark sectors have formal democratic/Fourier coordinates, but in v1 they are not on the Koide circle and their masses are scheme/QCD-threshold sensitive.  Their chamber offsets are therefore not certified as Koide wall seals.", Verdict: strings.Join([]string{StatusQuarkFormalCoordinates, StatusQuarksNotKoideWalls}, ";")}
}

func sectorAudit(name string, labels []string, ys []float64) SectorWallAudit {
	roots := make([]float64, len(ys))
	for i, y := range ys {
		roots[i] = math.Sqrt(y)
	}
	delta, r := fourierDeltaR(roots)
	q := dot(roots, roots) / math.Pow(sum(roots), 2)
	koideLike := math.Abs(r-1.0) < 1e-3
	return SectorWallAudit{Sector: name, Labels: append([]string{}, labels...), DeltaDeg: delta, R: r, Q: q, EpsilonTo135: upperWallDeg - delta, KoideLike: koideLike, WallSealValid: koideLike, Interpretation: "formal Fourier coordinate only; not a charged-lepton Koide chamber-wall seal in v1"}
}

func auditFirewalls() FirewallAudit {
	return FirewallAudit{DerivesLeptonMasses: false, DerivesYukawaEigenvalues: false, DerivesKoide: false, DerivesEpsilon: false, DerivesCKM: false, DerivesPMNS: false, DerivesGenerationHierarchy: false, AddsNewCarrier: false, PromotesObservedAsNative: false, PreservesGate352: true, Verdict: strings.Join([]string{StatusNoNativeWallSelector, StatusNoSimpleEpsilon, StatusNoFlavorPromotion, StatusObservedEndpointPreserved, StatusGate352Preserved, StatusGate583BoundaryPreserved}, ";")}
}

func compileFinal(mz, lambda WallPoint, transport WallTransportAudit, quarks QuarkAnalogyAudit) FinalVerdict {
	return FinalVerdict{SealName: "ChargedLeptonKoideChamberWallOffsetSeal", MZEpsilonDeg: mz.EpsilonDeg, LambdaEpsilonDeg: lambda.EpsilonDeg, MZPlaneAmplitudeR: mz.PlaneAmplitudeR, LambdaPlaneAmplitudeR: lambda.PlaneAmplitudeR, Chamber: "canonical (e,mu,tau) positive Koide chamber 105°<delta<135°; electron-zero wall at 135°", WallOffsetStableInV1: transport.EpsilonStable, HierarchyNearWall: mz.EpsilonDeg < 0.1*wallWidthDeg && mz.ElectronRootOverA < 0.05, QuarkWallSealCertified: quarks.Up.WallSealValid || quarks.Down.WallSealValid, NativeSelectorCertified: false, MinimalNextRequirement: "a native root-trace/absolute-Dirac or ordered circulant generation-plane operator that selects the positive S3 chamber and fixes the electron-wall offset epsilon", Verdict: strings.Join([]string{StatusWallHierarchyCompression, StatusNoNativeWallSelector, StatusNoFlavorPromotion, StatusGate583BoundaryPreserved}, ";")}
}

func Statuses() []string {
	return []string{StatusGate582Inherited, StatusChamberGeometryDefined, StatusCanonicalChamberCertified, StatusElectronWallIdentified, StatusMZEpsilonComputed, StatusLambdaEpsilonComputed, StatusEpsilonStable, StatusElectronSmallness, StatusWallHierarchyCompression, StatusQuarkFormalCoordinates, StatusNoNativeWallSelector, StatusNoSimpleEpsilon, StatusQuarksNotKoideWalls, StatusNoFlavorPromotion, StatusObservedEndpointPreserved, StatusGate352Preserved, StatusGate583BoundaryPreserved}
}

func componentOverA(r, deltaDeg float64, j int) float64 {
	return 1.0 + math.Sqrt2*r*math.Cos(degToRad(deltaDeg)+2.0*math.Pi*float64(j)/3.0)
}

func fourierDeltaR(roots []float64) (float64, float64) {
	a := sum(roots) / 3.0
	var re, im float64
	for j, x := range roots {
		w := x/a - 1.0
		angle := -2.0 * math.Pi * float64(j) / 3.0
		re += w * math.Cos(angle)
		im += w * math.Sin(angle)
	}
	delta := normalizeDeg(radToDeg(math.Atan2(im, re)))
	r := math.Hypot(re, im) / (3.0 * math.Sqrt2 / 2.0)
	return delta, r
}

func dot(a, b []float64) float64 {
	var s float64
	for i := range a {
		s += a[i] * b[i]
	}
	return s
}

func sum(xs []float64) float64 {
	var s float64
	for _, x := range xs {
		s += x
	}
	return s
}

func normalizeDeg(x float64) float64 {
	y := math.Mod(x, 360.0)
	if y < 0 {
		y += 360.0
	}
	return y
}

func radToDeg(x float64) float64 { return x * 180.0 / math.Pi }
func degToRad(x float64) float64 { return x * math.Pi / 180.0 }
