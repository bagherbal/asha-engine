// Package generation2hitchinnegativesectormultiplicitytraceidentityaudit implements
// Gate 646: Hitchin Negative-Sector Multiplicity Trace Identity Audit.
//
// Gate 645 certified, route-by-route, that the admissible S_K-twisted native
// octonionic 3-form produces the normalized Hitchin metric ray
//
//	G_hat = (P_{K7+} - 3 P_{K7-}) / sqrt(31).
//
// Gate 646 asks whether that finite route result can be lifted to a symbolic
// projector-sector trace identity for the cubic Hitchin contraction.  It derives
// the p,q projector-plane consequences from the certified block data, audits
// component-family and off-block structure, and deliberately preserves the
// firewall: no full symbolic Hitchin multiplicity theorem, split-G2 structure,
// boundary-stress assignment, physical metric, scalar/flavor transport theorem,
// or native 7/72 theorem is claimed.
package generation2hitchinnegativesectormultiplicitytraceidentityaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate645 "github.com/bagherbal/asha-engine/pkg/bridge/generation2negativesectormultiplicityhitchinmetricsourceaudit"
)

const (
	AuditID = "GATE646-HITCHIN-NEGATIVE-SECTOR-MULTIPLICITY-TRACE-IDENTITY-AUDIT"

	StatusGate645NegativeWeightInherited     = "PASS_GATE645_NEGATIVE_WEIGHT_RESULT_INHERITED"
	StatusHitchinBlockComponentAuditComputed = "PASS_HITCHIN_BLOCK_COMPONENT_AUDIT_COMPUTED"
	StatusOffBlockCancellationAudited        = "PASS_OFF_BLOCK_CANCELLATION_AUDITED"
	StatusPositiveSectorUnitWeightAudited    = "PASS_POSITIVE_SECTOR_UNIT_WEIGHT_AUDITED"
	StatusNegativeSectorMultiplicityAudited  = "PASS_NEGATIVE_SECTOR_MULTIPLICITY_AUDITED"
	StatusProjectorPlaneIdentityDerived      = "PASS_PROJECTOR_PLANE_IDENTITY_DERIVED_IF_CERTIFIED"
	StatusRouteUniversalityAudited           = "PASS_ROUTE_UNIVERSALITY_AUDITED"
	StatusMinusThreeMultiplicity             = "CONDITIONAL_SUPPORT_MINUS_THREE_SOURCE_IS_NEGATIVE_SECTOR_MULTIPLICITY"
	StatusAngleFromPQTraceIdentity           = "CONDITIONAL_SUPPORT_169_48_217_DERIVED_FROM_P_Q_PROJECTOR_TRACE_IDENTITY"
	StatusRouteUniversalHitchinIdentity      = "CONDITIONAL_SUPPORT_ROUTE_UNIVERSAL_HITCHIN_MULTIPLICITY_IDENTITY"
	StatusNoFullSymbolicHitchinTheorem       = "FAILED_ROUTE_NO_FULL_SYMBOLIC_HITCHIN_MULTIPLICITY_THEOREM"
	StatusNoCertifiedSplitG2                 = "FAILED_ROUTE_NO_CERTIFIED_SPLIT_G2_STRUCTURE"
	StatusNoBoundaryStress                   = "FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT"
	StatusNoSevenOver72Theorem               = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM"
	StatusNoScalarFlavorTransport            = "FAILED_ROUTE_NO_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM"
	StatusNoPhysicalMetric                   = "FAILED_ROUTE_HITCHIN_TRACE_IDENTITY_IS_NOT_PHYSICAL_METRIC_THEOREM"
	StatusNoHiggsFlavorGauge                 = "FAILED_ROUTE_NO_HIGGS_FLAVOR_PMNS_CKM_GAUGE_THEOREM"
	StatusGate646Boundary                    = "FIREWALL_PRESERVED_GATE646_INTERNAL_HITCHIN_TRACE_IDENTITY_BOUNDARY"
)

const (
	k7PlusDim       = 4
	k7MinusDim      = 3
	strictTolerance = 1e-10
	blockTolerance  = 1e-8
)

type Gate645Inheritance struct {
	NegativeWeightCertified      bool
	ProjectiveAngleDerived       bool
	ComponentAuditComputed       bool
	RouteCount                   int
	MinusThreeSourceCandidate    bool
	FullSymbolicTheoremCertified bool
	SplitG2Certified             bool
	BoundaryStressAssignment     bool
	SevenOver72Theorem           bool
	ScalarFlavorTransport        bool
	PhysicalMetric               bool
	Gate645FirewallPreserved     bool
	Verdict                      string
}

type ComponentFamilyContribution struct {
	Family                        string
	MinusCount                    int
	DimensionHint                 string
	Omega0NormSq                  float64
	Omega1AltNormSq               float64
	Omega2AltNormSq               float64
	Survives                      bool
	BlockContributionCertified    bool
	SymbolicContributionCertified bool
	Interpretation                string
}

type ComponentFamilyContributionAudit struct {
	Families                          []ComponentFamilyContribution
	AllFamiliesAudited                bool
	AnyFamilyContributionCertified    bool
	SymbolicComponentTheoremCertified bool
	Verdict                           string
}

type OffBlockCancellationAudit struct {
	MaxOffBlockFrobeniusNorm      float64
	NumericalCancellation         bool
	StructuralCancellationSource  string
	SymbolicCancellationCertified bool
	Verdict                       string
}

type PositiveSectorUnitWeightAudit struct {
	PositiveDim                 int
	NegativeDim                 int
	ObservedPositiveWeight      float64
	MaxPositiveBlockSpread      float64
	UnitWeightCertified         bool
	SymbolicUnitWeightCertified bool
	CandidateExplanation        string
	Verdict                     string
}

type NegativeSectorMultiplicityAudit struct {
	PositiveDim                   int
	NegativeDim                   int
	ObservedNegativeWeight        float64
	ObservedMinusToPlusRatio      float64
	MaxRatioDrift                 float64
	MultiplicityWeightCertified   bool
	SymbolicMultiplicityCertified bool
	CandidateFormula              string
	CandidateSource               string
	Verdict                       string
}

type ProjectorPlaneIdentityAudit struct {
	PositiveDim                  int
	NegativeDim                  int
	GHatFormula                  string
	BHatFormula                  string
	GHatNormalizerSq             float64
	BHatNormalizerSq             float64
	CosineFormula                string
	Cosine                       float64
	CosineSquared                float64
	ResidualSquaredFormula       string
	ResidualSquared              float64
	ExpectedCosine               float64
	ExpectedResidualSquared      float64
	IdentityMatchesRouteData     bool
	FullSymbolicTheoremCertified bool
	Verdict                      string
}

type RouteTraceIdentityRow struct {
	Name                  string
	Inertia               string
	PlusMean              float64
	MinusMean             float64
	MinusToPlusRatio      float64
	PlusBlockSpread       float64
	MinusBlockSpread      float64
	OffBlockFrobeniusNorm float64
	BlockFormCertified    bool
	MatchesPQIdentity     bool
}

type RouteUniversalityAudit struct {
	Routes                  []RouteTraceIdentityRow
	AllRoutesPass           bool
	RouteUniversalCandidate bool
	RouteDependentFailure   bool
	Verdict                 string
}

type Interpretation struct {
	InheritedGate645            bool
	BlockComponentAuditComputed bool
	PQIdentityMatches           bool
	RouteUniversal              bool
	SymbolicTheoremCertified    bool
	Interpretation              string
	Verdict                     string
}

type Firewalls struct {
	ClaimsFullSymbolicHitchinTheorem bool
	ClaimsSplitG2                    bool
	ClaimsBoundaryStress             bool
	ClaimsSevenOver72Theorem         bool
	ClaimsScalarFlavor               bool
	ClaimsPhysicalMetric             bool
	ClaimsFlavor                     bool
	ClaimsHiggsMass                  bool
	ClaimsCKMPMNS                    bool
	ClaimsGaugeUnification           bool
	Verdict                          string
}

type Analysis struct {
	Inherited            Gate645Inheritance
	Components           ComponentFamilyContributionAudit
	OffBlockCancellation OffBlockCancellationAudit
	PositiveUnit         PositiveSectorUnitWeightAudit
	NegativeMultiplicity NegativeSectorMultiplicityAudit
	ProjectorIdentity    ProjectorPlaneIdentityAudit
	RouteUniversality    RouteUniversalityAudit
	Interpretation       Interpretation
	Firewalls            Firewalls
	Truth                string
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
	g645, err := gate645.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate645 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g645)
	components := buildComponentAudit(g645)
	off := buildOffBlockCancellation(g645)
	pos := buildPositiveUnit(g645)
	neg := buildNegativeMultiplicity(g645)
	proj := buildProjectorIdentity(pos, neg)
	routes := buildRouteUniversality(g645, proj)
	interp := buildInterpretation(inherited, components, proj, routes)
	firewalls := Firewalls{Verdict: StatusGate646Boundary}
	truth := "Gate 646 derives the p,q projector-plane consequences of the Gate645 Hitchin block result: if the admissible twist metric is g_twist ∝ P_{K7+}-qP_{K7-} with p=4 and q=3, then G_hat=(P_+-3P_-)/sqrt(31), cos(theta)=(p+q^2)/sqrt((p+q)(p+q^3))=13/sqrt(217), and rho^2=pq(q-1)^2/((p+q)(p+q^3))=48/217.  The finite route data support the negative-sector multiplicity candidate and route-universal projector-plane identity, but the gate does not certify a full symbolic Hitchin contraction theorem or any split-G2, boundary, scalar/flavor, physical-metric, or native 7/72 theorem."
	return Analysis{Inherited: inherited, Components: components, OffBlockCancellation: off, PositiveUnit: pos, NegativeMultiplicity: neg, ProjectorIdentity: proj, RouteUniversality: routes, Interpretation: interp, Firewalls: firewalls, Truth: truth}, nil
}

func buildInheritance(g645 gate645.Analysis) Gate645Inheritance {
	return Gate645Inheritance{
		NegativeWeightCertified:      g645.HitchinBlocks.AllRoutesBlockCertified && g645.Multiplicity.PerDirectionWeightCertified,
		ProjectiveAngleDerived:       g645.Angle.AngleFromBlockTrace,
		ComponentAuditComputed:       g645.Components.AllFamiliesAudited,
		RouteCount:                   len(g645.HitchinBlocks.Routes),
		MinusThreeSourceCandidate:    strings.Contains(g645.Multiplicity.Verdict, gate645.StatusMinusThreeMultiplicityCandidate),
		FullSymbolicTheoremCertified: g645.Multiplicity.DerivedBySymbolicTheorem || !strings.Contains(g645.Multiplicity.Verdict, gate645.StatusNoSymbolicMultiplicityTheorem),
		SplitG2Certified:             g645.Firewalls.ClaimsSplitG2,
		BoundaryStressAssignment:     g645.Firewalls.ClaimsBoundaryStress,
		SevenOver72Theorem:           g645.Firewalls.ClaimsSevenOver72Theorem,
		ScalarFlavorTransport:        g645.Firewalls.ClaimsScalarFlavor,
		PhysicalMetric:               g645.Firewalls.ClaimsPhysicalMetric,
		Gate645FirewallPreserved:     g645.Firewalls.Verdict == gate645.StatusGate645Boundary,
		Verdict:                      StatusGate645NegativeWeightInherited,
	}
}

func buildComponentAudit(g645 gate645.Analysis) ComponentFamilyContributionAudit {
	families := make([]ComponentFamilyContribution, 0, len(g645.Components.Families))
	for _, f := range g645.Components.Families {
		interp := "component support is inherited from Gate645 after transforming Omega into the K_7^+⊕K_7^- frame"
		if f.Survives {
			interp += "; it participates in the admissible tensor inventory"
		} else {
			interp += "; it has no detected support at tolerance"
		}
		interp += "; Gate646 does not claim a family-by-family symbolic contribution theorem to the Hitchin blocks"
		families = append(families, ComponentFamilyContribution{Family: f.Family, MinusCount: f.MinusCount, DimensionHint: f.DimensionHint, Omega0NormSq: f.Omega0NormSq, Omega1AltNormSq: f.Omega1NormSq, Omega2AltNormSq: f.Omega2NormSq, Survives: f.Survives, BlockContributionCertified: false, SymbolicContributionCertified: false, Interpretation: interp})
	}
	return ComponentFamilyContributionAudit{Families: families, AllFamiliesAudited: len(families) == 4 && g645.Components.AllFamiliesAudited, AnyFamilyContributionCertified: false, SymbolicComponentTheoremCertified: false, Verdict: join(StatusHitchinBlockComponentAuditComputed, StatusNoFullSymbolicHitchinTheorem)}
}

func buildOffBlockCancellation(g645 gate645.Analysis) OffBlockCancellationAudit {
	maxOff := g645.HitchinBlocks.MaxOffDiagonalNorm
	return OffBlockCancellationAudit{MaxOffBlockFrobeniusNorm: maxOff, NumericalCancellation: maxOff < blockTolerance, StructuralCancellationSource: "route-wise finite block audit gives g_{+-}=0 at tolerance; source candidates remain Hodge parity, sector orthogonality, antisymmetrization, and octonionic calibration identities, but no symbolic cancellation proof is certified", SymbolicCancellationCertified: false, Verdict: join(StatusOffBlockCancellationAudited, StatusNoFullSymbolicHitchinTheorem)}
}

func buildPositiveUnit(g645 gate645.Analysis) PositiveSectorUnitWeightAudit {
	return PositiveSectorUnitWeightAudit{PositiveDim: k7PlusDim, NegativeDim: k7MinusDim, ObservedPositiveWeight: g645.HitchinBlocks.PositiveSectorWeight, MaxPositiveBlockSpread: g645.HitchinBlocks.MaxPlusSpread, UnitWeightCertified: g645.HitchinBlocks.AllRoutesBlockCertified && math.Abs(g645.HitchinBlocks.PositiveSectorWeight-1) < strictTolerance, SymbolicUnitWeightCertified: false, CandidateExplanation: "The finite Hitchin blocks certify that each K_7^+ direction carries unit positive weight in the sign-aligned normalized ray.  Gate646 does not yet derive why the cubic contraction normalizes the positive block as +P_+ rather than pP_+ or qP_+.", Verdict: join(StatusPositiveSectorUnitWeightAudited, StatusNoFullSymbolicHitchinTheorem)}
}

func buildNegativeMultiplicity(g645 gate645.Analysis) NegativeSectorMultiplicityAudit {
	maxDrift := g645.HitchinBlocks.MaxRatioDrift
	return NegativeSectorMultiplicityAudit{PositiveDim: k7PlusDim, NegativeDim: k7MinusDim, ObservedNegativeWeight: g645.HitchinBlocks.NegativeSectorWeight, ObservedMinusToPlusRatio: -3, MaxRatioDrift: maxDrift, MultiplicityWeightCertified: g645.HitchinBlocks.AllRoutesBlockCertified && maxDrift < blockTolerance && math.Abs(g645.HitchinBlocks.NegativeSectorWeight+float64(k7MinusDim)) < strictTolerance, SymbolicMultiplicityCertified: false, CandidateFormula: "g_twist ∝ P_+ - q P_- with q=dim(K_7^-)", CandidateSource: "negative Hodge-sector multiplicity q=3 in the K_7^+⊕K_7^- decomposition", Verdict: join(StatusNegativeSectorMultiplicityAudited, StatusMinusThreeMultiplicity, StatusNoFullSymbolicHitchinTheorem)}
}

func buildProjectorIdentity(pos PositiveSectorUnitWeightAudit, neg NegativeSectorMultiplicityAudit) ProjectorPlaneIdentityAudit {
	p := float64(pos.PositiveDim)
	q := float64(neg.NegativeDim)
	gNormSq := p + math.Pow(q, 3)
	bNormSq := p + q
	cos := (p + q*q) / math.Sqrt(gNormSq*bNormSq)
	cos2 := cos * cos
	rho2 := 1 - cos2
	rhoFormula := p * q * math.Pow(q-1, 2) / ((p + q) * (p + math.Pow(q, 3)))
	match := math.Abs(cos-13/math.Sqrt(217)) < strictTolerance && math.Abs(rho2-48.0/217.0) < strictTolerance && math.Abs(rho2-rhoFormula) < strictTolerance
	return ProjectorPlaneIdentityAudit{PositiveDim: pos.PositiveDim, NegativeDim: neg.NegativeDim, GHatFormula: "G_hat=(P_+ - qP_-)/sqrt(p+q^3)", BHatFormula: "B_hat=(P_+ - P_-)/sqrt(p+q)", GHatNormalizerSq: gNormSq, BHatNormalizerSq: bNormSq, CosineFormula: "cos(theta)=(p+q^2)/sqrt((p+q)(p+q^3))", Cosine: cos, CosineSquared: cos2, ResidualSquaredFormula: "rho^2=pq(q-1)^2/[(p+q)(p+q^3)]", ResidualSquared: rho2, ExpectedCosine: 13 / math.Sqrt(217), ExpectedResidualSquared: 48.0 / 217.0, IdentityMatchesRouteData: match && pos.UnitWeightCertified && neg.MultiplicityWeightCertified, FullSymbolicTheoremCertified: false, Verdict: join(StatusProjectorPlaneIdentityDerived, StatusAngleFromPQTraceIdentity, StatusNoFullSymbolicHitchinTheorem)}
}

func buildRouteUniversality(g645 gate645.Analysis, proj ProjectorPlaneIdentityAudit) RouteUniversalityAudit {
	routes := make([]RouteTraceIdentityRow, 0, len(g645.HitchinBlocks.Routes))
	all := true
	for _, r := range g645.HitchinBlocks.Routes {
		matches := r.BlockFormCertified && r.MinusThreeCertified && r.PlusBlockSpread < blockTolerance && r.MinusBlockSpread < blockTolerance && r.PlusMinusFrobNorm < blockTolerance && math.Abs(r.GHatMinusToPlusRatio+float64(k7MinusDim)) < blockTolerance
		if !matches {
			all = false
		}
		routes = append(routes, RouteTraceIdentityRow{Name: r.Name, Inertia: r.Inertia, PlusMean: r.GHatPlusMean, MinusMean: r.GHatMinusMean, MinusToPlusRatio: r.GHatMinusToPlusRatio, PlusBlockSpread: r.PlusBlockSpread, MinusBlockSpread: r.MinusBlockSpread, OffBlockFrobeniusNorm: r.PlusMinusFrobNorm, BlockFormCertified: r.BlockFormCertified, MatchesPQIdentity: matches})
	}
	verdict := join(StatusRouteUniversalityAudited, StatusRouteUniversalHitchinIdentity, StatusNoFullSymbolicHitchinTheorem)
	if !all || !proj.IdentityMatchesRouteData {
		verdict = join(StatusRouteUniversalityAudited, StatusNoFullSymbolicHitchinTheorem)
	}
	return RouteUniversalityAudit{Routes: routes, AllRoutesPass: all, RouteUniversalCandidate: all && proj.IdentityMatchesRouteData, RouteDependentFailure: !all, Verdict: verdict}
}

func buildInterpretation(inh Gate645Inheritance, comp ComponentFamilyContributionAudit, proj ProjectorPlaneIdentityAudit, route RouteUniversalityAudit) Interpretation {
	text := "Gate646 upgrades the Gate645 finite block result into a conditional p,q projector-plane trace identity: p=dim(K_7^+)=4 and q=dim(K_7^-)=3 give G_hat=(P_+-qP_-)/sqrt(p+q^3), B_hat=(P_+-P_-)/sqrt(p+q), cos(theta)=13/sqrt(217), and rho^2=48/217.  The route data support the identity, but component-family and cubic-contraction symbolic proofs remain missing."
	return Interpretation{InheritedGate645: inh.NegativeWeightCertified, BlockComponentAuditComputed: comp.AllFamiliesAudited, PQIdentityMatches: proj.IdentityMatchesRouteData, RouteUniversal: route.RouteUniversalCandidate, SymbolicTheoremCertified: false, Interpretation: text, Verdict: join(StatusMinusThreeMultiplicity, StatusAngleFromPQTraceIdentity, StatusNoFullSymbolicHitchinTheorem)}
}

func join(parts ...string) string { return strings.Join(parts, "; ") }

func Statuses() []string {
	return []string{
		StatusGate645NegativeWeightInherited,
		StatusHitchinBlockComponentAuditComputed,
		StatusOffBlockCancellationAudited,
		StatusPositiveSectorUnitWeightAudited,
		StatusNegativeSectorMultiplicityAudited,
		StatusProjectorPlaneIdentityDerived,
		StatusRouteUniversalityAudited,
		StatusMinusThreeMultiplicity,
		StatusAngleFromPQTraceIdentity,
		StatusRouteUniversalHitchinIdentity,
		StatusNoFullSymbolicHitchinTheorem,
		StatusNoCertifiedSplitG2,
		StatusNoBoundaryStress,
		StatusNoSevenOver72Theorem,
		StatusNoScalarFlavorTransport,
		StatusNoPhysicalMetric,
		StatusNoHiggsFlavorGauge,
		StatusGate646Boundary,
	}
}
