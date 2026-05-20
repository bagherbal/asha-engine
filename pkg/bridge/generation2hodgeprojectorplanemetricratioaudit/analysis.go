// Package generation2hodgeprojectorplanemetricratioaudit implements
// Gate 644: HodgeProjector Plane MetricRatio Audit.
//
// Gate 643 showed that the projective compact/split residual tensor is not an
// off-sector leakage term.  It is the same-sector diagonal complement to
// B_hat inside span{P_{K7+}, P_{K7-}}.  Gate 644 reconstructs the normalized
// twist metric itself and audits whether the repeated routes collapse to the
// simpler projector-plane ratio
//
//	G_hat = (P_{K7+} - 3 P_{K7-}) / sqrt(31),
//
// while
//
//	B_hat = (P_{K7+} - P_{K7-}) / sqrt(7).
//
// This derives the Gate642 projective angle from the two diagonal rays
// (1,-1) and (1,-3), but does not certify a native source theorem for the
// -3 weight, a split-G2 structure, a boundary-stress assignment, scalar/flavor
// transport, physical geometry, or a native 7/72 theorem.
package generation2hodgeprojectorplanemetricratioaudit

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	gate643 "github.com/bagherbal/asha-engine/pkg/bridge/generation2compactsplitresidualtensorblockstructureaudit"
	"github.com/bagherbal/asha-engine/pkg/combinatorics"
	"github.com/bagherbal/asha-engine/pkg/geometry/contact"
	"github.com/bagherbal/asha-engine/pkg/linear"
	"github.com/bagherbal/asha-engine/pkg/octonion"
)

const (
	AuditID = "GATE644-HODGE-PROJECTOR-PLANE-METRIC-RATIO-AUDIT"

	StatusGate643ResidualInherited     = "PASS_GATE643_RESIDUAL_TENSOR_BLOCK_STRUCTURE_INHERITED"
	StatusGHatReconstructed            = "PASS_GHAT_RECONSTRUCTED_FROM_BHAT_AND_RHAT"
	StatusProjectorPlaneMetricsDefined = "PASS_HODGE_PROJECTOR_PLANE_METRICS_DEFINED"
	StatusRouteMetricRatiosComputed    = "PASS_ROUTE_METRIC_RATIOS_COMPUTED"
	StatusHodgeDiagonalRatio           = "CONDITIONAL_SUPPORT_GTWIST_HAS_HODGE_DIAGONAL_RATIO_1_TO_MINUS_3"
	StatusProjectorPlaneAngle          = "CONDITIONAL_SUPPORT_PROJECTIVE_ANGLE_DERIVES_FROM_PROJECTOR_PLANE_GEOMETRY"
	StatusMinusThreeSourceCandidate    = "CONDITIONAL_SUPPORT_MINUS_THREE_WEIGHT_MATCHES_NEGATIVE_HODGE_SECTOR_DIMENSION_CANDIDATE"
	StatusNoMinusThreeSource           = "FAILED_ROUTE_NO_NATIVE_SOURCE_FOR_MINUS_THREE_WEIGHT_YET"
	StatusNoNativeTraceIdentity        = "FAILED_ROUTE_NO_NATIVE_TRACE_IDENTITY_FOR_PROJECTOR_PLANE_RATIO_YET"
	StatusNoCertifiedSplitG2           = "FAILED_ROUTE_NO_CERTIFIED_SPLIT_G2_STRUCTURE"
	StatusNoBoundaryStress             = "FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT"
	StatusNoSevenOver72Theorem         = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM"
	StatusNoScalarFlavorTransport      = "FAILED_ROUTE_NO_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM"
	StatusNoPhysicalAngle              = "FAILED_ROUTE_PROJECTOR_PLANE_ANGLE_IS_NOT_PHYSICAL_ANGLE"
	StatusNoPhysicalMetric             = "FAILED_ROUTE_PROJECTOR_PLANE_METRIC_IS_NOT_PHYSICAL_METRIC_THEOREM"
	StatusNoHiggsFlavorGauge           = "FAILED_ROUTE_NO_HIGGS_FLAVOR_PMNS_CKM_GAUGE_THEOREM"
	StatusGate644Boundary              = "FIREWALL_PRESERVED_GATE644_PROJECTOR_PLANE_RATIO_IS_INTERNAL_ONLY"
)

const (
	vectorDimExpected  = 8
	lambda4DimExpected = 70
	k7DimExpected      = 7
	k7PlusDim          = 4
	k7MinusDim         = 3
	g2CopyDim          = 7
	alignmentRoot      = 13
	failureNumerator   = 48
	angleDenominator   = 217
	projectorDenom     = 31
	strictTolerance    = 1e-10
	numericalTolerance = 1e-8
	angleTolerance     = 1e-9
	ratioTolerance     = 1e-8
)

type Gate643Inheritance struct {
	CosTheta                 float64
	SinTheta                 float64
	ResidualTensorCertified  bool
	SameSectorHodgeDiagonal  bool
	OffSectorCarrierRejected bool
	RPlusPlusFrobSquared     float64
	RMinusMinusFrobSquared   float64
	TwiceRPlusMinusFrobSq    float64
	NativeTraceIdentityFound bool
	SplitG2Certified         bool
	BoundaryStressAssignment bool
	SevenOver72Theorem       bool
	ScalarFlavorTransport    bool
	PhysicalAngle            bool
	PhysicalMetric           bool
	Gate643FirewallPreserved bool
	Verdict                  string
}

type ProjectorPlaneDefinition struct {
	BHatFormula                    string
	GHatFormula                    string
	BHatPlusWeight                 float64
	BHatMinusWeight                float64
	GHatPlusWeight                 float64
	GHatMinusWeight                float64
	BHatNorm                       float64
	GHatNorm                       float64
	BHatTargetResidual             float64
	ProjectorPlaneTargetResidual   float64
	ProjectorPlaneMetricsCertified bool
	Verdict                        string
}

type MetricRatioRoute struct {
	Name                         string
	Formula                      string
	Inertia                      string
	Cosine                       float64
	Rho                          float64
	GHatToProjectorPlaneResidual float64
	GHatToReconstructedResidual  float64
	PlusBlockEigenvalues         []float64
	MinusBlockEigenvalues        []float64
	PlusBlockMean                float64
	MinusBlockMean               float64
	PlusBlockSpread              float64
	MinusBlockSpread             float64
	PlusMinusFrobNorm            float64
	ObservedMinusToPlusRatio     float64
	Ratio1ToMinus3Certified      bool
	Comment                      string
}

type MetricRatioAudit struct {
	Routes                    []MetricRatioRoute
	MaxProjectorPlaneResidual float64
	MaxReconstructedResidual  float64
	MaxPlusSpread             float64
	MaxMinusSpread            float64
	MaxOffDiagonalNorm        float64
	MaxRatioDrift             float64
	AllRoutesRatioCertified   bool
	Verdict                   string
}

type ProjectiveAngleFromPlane struct {
	PlusDim                  int
	MinusDim                 int
	BHatWeights              string
	GHatWeights              string
	InnerProductFormula      string
	ComputedCosine           float64
	ExpectedCosine           float64
	ComputedSinSquared       float64
	ExpectedSinSquared       float64
	AngleDerivedFromPlane    bool
	NativeTraceIdentityFound bool
	Verdict                  string
}

type MinusThreeSourceAudit struct {
	NegativeHodgeSectorDim      int
	CandidateFromDimK7Minus     string
	CandidateFromTraceBalance   string
	CandidateFromTwistOperation string
	CertifiedNativeSource       bool
	Verdict                     string
}

type Interpretation struct {
	Gate643Inherited         bool
	GHatReconstructed        bool
	RatioCertified           bool
	AngleFromPlane           bool
	MinusThreeSourceFound    bool
	NativeTraceIdentityFound bool
	Interpretation           string
	Verdict                  string
}

type Firewalls struct {
	ClaimsNativeTraceIdentity bool
	ClaimsSplitG2             bool
	ClaimsBoundaryStress      bool
	ClaimsSevenOver72Theorem  bool
	ClaimsScalarFlavor        bool
	ClaimsPhysicalAngle       bool
	ClaimsPhysicalMetric      bool
	ClaimsFlavor              bool
	ClaimsHiggsMass           bool
	ClaimsCKMPMNS             bool
	ClaimsGaugeUnification    bool
	Verdict                   string
}

type Analysis struct {
	Inherited      Gate643Inheritance
	Definitions    ProjectorPlaneDefinition
	MetricRatio    MetricRatioAudit
	AngleFromPlane ProjectiveAngleFromPlane
	MinusThree     MinusThreeSourceAudit
	Interpretation Interpretation
	Firewalls      Firewalls
	Truth          string
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
	g643, err := gate643.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate643 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g643)

	space, err := contact.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build contact space: %w", err)
	}
	if space.AmbientDimension() != lambda4DimExpected || space.Dimension() != k7DimExpected {
		return Analysis{}, fmt.Errorf("unexpected contact dimensions: ambient=%d K7=%d", space.AmbientDimension(), space.Dimension())
	}
	star, err := hodgeStarLambda4R8()
	if err != nil {
		return Analysis{}, err
	}
	bk, err := restrictedHodgeOperator(space.ContactFrame, star)
	if err != nil {
		return Analysis{}, err
	}
	qPlus, qMinus, err := hodgePolarityBases(bk)
	if err != nil {
		return Analysis{}, err
	}
	pPlus, err := projector(qPlus)
	if err != nil {
		return Analysis{}, err
	}
	pMinus, err := projector(qMinus)
	if err != nil {
		return Analysis{}, err
	}
	omega0, err := nativeOmega0(space)
	if err != nil {
		return Analysis{}, err
	}
	defs, err := buildDefinitions(bk, pPlus, pMinus)
	if err != nil {
		return Analysis{}, err
	}
	ratioAudit, err := buildMetricRatioAudit(omega0, bk, pPlus, pMinus, qPlus, qMinus, defs, inherited)
	if err != nil {
		return Analysis{}, err
	}
	angle := buildAngleFromPlane(defs)
	minusThree := MinusThreeSourceAudit{
		NegativeHodgeSectorDim:      k7MinusDim,
		CandidateFromDimK7Minus:     "-3 matches -dim(K_7^-), the negative Hodge sector dimension in Gate634.",
		CandidateFromTraceBalance:   "The ratio (1,-3) is trace-asymmetric across 4|3 sectors; it is compatible with but not derived from tr(S_K)=+1.",
		CandidateFromTwistOperation: "The antisymmetrized Gate638 twists repeatedly yield the same normalized diagonal ray, but the finite code has not derived why the twist must produce the -3 weight.",
		CertifiedNativeSource:       false,
		Verdict:                     join(StatusMinusThreeSourceCandidate, StatusNoMinusThreeSource),
	}
	interp := buildInterpretation(inherited, ratioAudit, angle, minusThree)
	firewalls := Firewalls{Verdict: StatusGate644Boundary}
	truth := "Gate 644 reconstructs the Gate643 projective residual into the normalized twist metric itself.  Across omega_1_alt, omega_2_alt, and omega_B_alt, G_hat collapses to the Hodge-projector plane ray (P_{K7+}-3P_{K7-})/sqrt(31), while B_hat=(P_{K7+}-P_{K7-})/sqrt(7).  The Gate642 angle then follows from the diagonal pair (1,-1) versus (1,-3): <G_hat,B_hat>=(4+9)/sqrt(31*7)=13/sqrt(217), and the residual square is 48/217.  This is a cleaner internal projector-plane geometry, but it still lacks a native source theorem for the -3 weight and does not certify split-G2, boundary stress, scalar/flavor transport, physical geometry, or a native 7/72 theorem."
	return Analysis{Inherited: inherited, Definitions: defs, MetricRatio: ratioAudit, AngleFromPlane: angle, MinusThree: minusThree, Interpretation: interp, Firewalls: firewalls, Truth: truth}, nil
}

func buildInheritance(g643 gate643.Analysis) Gate643Inheritance {
	pp, mm, pm := g643.BlockSummary.PlusPlusMeanFrobSquared, g643.BlockSummary.MinusMinusMeanFrobSquared, g643.BlockSummary.TwicePlusMinusMeanFrobSquared
	return Gate643Inheritance{
		CosTheta:                 g643.Inherited.CosTheta,
		SinTheta:                 g643.Inherited.SinTheta,
		ResidualTensorCertified:  g643.ResidualTensor.ResidualTensorsCertified,
		SameSectorHodgeDiagonal:  strings.Contains(g643.BlockSummary.Verdict, gate643.StatusSameSectorHodgeDiagonal),
		OffSectorCarrierRejected: strings.Contains(g643.BlockSummary.Verdict, gate643.StatusNoOffSectorCarrier),
		RPlusPlusFrobSquared:     pp,
		RMinusMinusFrobSquared:   mm,
		TwiceRPlusMinusFrobSq:    pm,
		NativeTraceIdentityFound: g643.BlockSummary.NativeTraceIdentityFound,
		SplitG2Certified:         g643.Firewalls.ClaimsSplitG2,
		BoundaryStressAssignment: g643.Firewalls.ClaimsBoundaryStress,
		SevenOver72Theorem:       g643.Firewalls.ClaimsSevenOver72Theorem,
		ScalarFlavorTransport:    g643.Firewalls.ClaimsScalarFlavor,
		PhysicalAngle:            g643.Firewalls.ClaimsPhysicalAngle,
		PhysicalMetric:           g643.Firewalls.ClaimsPhysicalMetric,
		Gate643FirewallPreserved: g643.Firewalls.Verdict == gate643.StatusGate643Boundary,
		Verdict:                  StatusGate643ResidualInherited,
	}
}

func buildDefinitions(bk, pPlus, pMinus linear.Matrix) (ProjectorPlaneDefinition, error) {
	bNorm := bk.FrobeniusNorm()
	if bNorm == 0 {
		return ProjectorPlaneDefinition{}, fmt.Errorf("zero B_K norm")
	}
	bHat := bk.Scale(1 / bNorm)
	targetB, err := pPlus.Sub(pMinus)
	if err != nil {
		return ProjectorPlaneDefinition{}, err
	}
	targetB = targetB.Scale(1 / math.Sqrt(float64(k7DimExpected)))
	bResidual, err := bHat.Sub(targetB)
	if err != nil {
		return ProjectorPlaneDefinition{}, err
	}
	targetG, err := pPlus.Sub(pMinus.Scale(3))
	if err != nil {
		return ProjectorPlaneDefinition{}, err
	}
	targetG = targetG.Scale(1 / math.Sqrt(float64(projectorDenom)))
	gNorm := targetG.FrobeniusNorm()
	return ProjectorPlaneDefinition{
		BHatFormula:                    "B_hat=(P_{K7+}-P_{K7-})/sqrt(7)",
		GHatFormula:                    "G_hat=(P_{K7+}-3P_{K7-})/sqrt(31)",
		BHatPlusWeight:                 1 / math.Sqrt(float64(k7DimExpected)),
		BHatMinusWeight:                -1 / math.Sqrt(float64(k7DimExpected)),
		GHatPlusWeight:                 1 / math.Sqrt(float64(projectorDenom)),
		GHatMinusWeight:                -3 / math.Sqrt(float64(projectorDenom)),
		BHatNorm:                       bHat.FrobeniusNorm(),
		GHatNorm:                       gNorm,
		BHatTargetResidual:             bResidual.FrobeniusNorm(),
		ProjectorPlaneTargetResidual:   math.Abs(gNorm - 1),
		ProjectorPlaneMetricsCertified: bResidual.FrobeniusNorm() < strictTolerance && math.Abs(gNorm-1) < strictTolerance,
		Verdict:                        StatusProjectorPlaneMetricsDefined,
	}, nil
}

func buildMetricRatioAudit(omega0 tensor3, bk, pPlus, pMinus, qPlus, qMinus linear.Matrix, defs ProjectorPlaneDefinition, inherited Gate643Inheritance) (MetricRatioAudit, error) {
	sk := bk
	invGOmega, err := inverse(hitchinBMatrix(omega0))
	if err != nil {
		return MetricRatioAudit{}, err
	}
	omegaB, err := compactCrossProductBKTensor(omega0, invGOmega, bk)
	if err != nil {
		return MetricRatioAudit{}, err
	}
	raw1 := transformTensor(omega0, sk, linear.Identity(k7DimExpected), linear.Identity(k7DimExpected))
	raw2 := transformTensor(omega0, sk, sk, linear.Identity(k7DimExpected))
	candidates := []struct {
		name, formula string
		tensor        tensor3
	}{
		{"omega_1_alt", "Alt[Ω_0(S_K x,y,z)]", alternateTensor(raw1)},
		{"omega_2_alt", "Alt[Ω_0(S_K x,S_K y,z)]", alternateTensor(raw2)},
		{"omega_B_alt", "Alt[B_K(x ×_{Ω_0} y,z)]", alternateTensor(omegaB)},
	}

	targetG, err := pPlus.Sub(pMinus.Scale(3))
	if err != nil {
		return MetricRatioAudit{}, err
	}
	targetG = targetG.Scale(1 / math.Sqrt(float64(projectorDenom)))
	maxPlane, maxReconstruct, maxPSpread, maxMSpread, maxOff, maxRatio := 0.0, 0.0, 0.0, 0.0, 0.0, 0.0
	routes := make([]MetricRatioRoute, 0, len(candidates))
	all := true
	for _, c := range candidates {
		metric := hitchinBMatrix(c.tensor)
		r, err := auditMetricRatioRoute(c.name, c.formula, metric, bk, targetG, qPlus, qMinus, inherited.CosTheta, inherited.SinTheta)
		if err != nil {
			return MetricRatioAudit{}, err
		}
		if r.GHatToProjectorPlaneResidual > maxPlane {
			maxPlane = r.GHatToProjectorPlaneResidual
		}
		if r.GHatToReconstructedResidual > maxReconstruct {
			maxReconstruct = r.GHatToReconstructedResidual
		}
		if r.PlusBlockSpread > maxPSpread {
			maxPSpread = r.PlusBlockSpread
		}
		if r.MinusBlockSpread > maxMSpread {
			maxMSpread = r.MinusBlockSpread
		}
		if r.PlusMinusFrobNorm > maxOff {
			maxOff = r.PlusMinusFrobNorm
		}
		ratioDrift := math.Abs(r.ObservedMinusToPlusRatio + 3)
		if ratioDrift > maxRatio {
			maxRatio = ratioDrift
		}
		if !r.Ratio1ToMinus3Certified {
			all = false
		}
		routes = append(routes, r)
	}
	verdict := join(StatusGHatReconstructed, StatusRouteMetricRatiosComputed, StatusHodgeDiagonalRatio)
	if !all {
		verdict = StatusNoNativeTraceIdentity
	}
	return MetricRatioAudit{Routes: routes, MaxProjectorPlaneResidual: maxPlane, MaxReconstructedResidual: maxReconstruct, MaxPlusSpread: maxPSpread, MaxMinusSpread: maxMSpread, MaxOffDiagonalNorm: maxOff, MaxRatioDrift: maxRatio, AllRoutesRatioCertified: all, Verdict: verdict}, nil
}

func auditMetricRatioRoute(name, formula string, gTwist, bk, targetG, qPlus, qMinus linear.Matrix, cosTarget, rhoTarget float64) (MetricRatioRoute, error) {
	bNorm := bk.FrobeniusNorm()
	gNorm := gTwist.FrobeniusNorm()
	if bNorm == 0 || gNorm == 0 {
		return MetricRatioRoute{}, fmt.Errorf("zero norm in route %s", name)
	}
	bHat := bk.Scale(1 / bNorm)
	gHat := gTwist.Scale(1 / gNorm)
	inner := frobeniusInner(gHat, bHat)
	if inner < 0 {
		gHat = gHat.Scale(-1)
		inner = -inner
	}
	residualRaw, err := gHat.Sub(bHat.Scale(inner))
	if err != nil {
		return MetricRatioRoute{}, err
	}
	rho := residualRaw.FrobeniusNorm()
	if rho == 0 {
		return MetricRatioRoute{}, fmt.Errorf("zero residual in route %s", name)
	}
	rHat := residualRaw.Scale(1 / rho)
	reconstructed, err := bHat.Scale(cosTarget).Add(rHat.Scale(rhoTarget))
	if err != nil {
		return MetricRatioRoute{}, err
	}
	reconResidual, err := gHat.Sub(reconstructed)
	if err != nil {
		return MetricRatioRoute{}, err
	}
	planeResidual, err := gHat.Sub(targetG)
	if err != nil {
		return MetricRatioRoute{}, err
	}
	pp, err := sandwich(qPlus, gHat, qPlus)
	if err != nil {
		return MetricRatioRoute{}, err
	}
	mm, err := sandwich(qMinus, gHat, qMinus)
	if err != nil {
		return MetricRatioRoute{}, err
	}
	pm, err := sandwich(qPlus, gHat, qMinus)
	if err != nil {
		return MetricRatioRoute{}, err
	}
	ppEig, _ := sortedEigenvalues(pp)
	mmEig, _ := sortedEigenvalues(mm)
	ppMean := mean(ppEig)
	mmMean := mean(mmEig)
	ppSpread := maxAbsDeviation(ppEig, 1/math.Sqrt(float64(projectorDenom)))
	mmSpread := maxAbsDeviation(mmEig, -3/math.Sqrt(float64(projectorDenom)))
	ratio := math.NaN()
	if math.Abs(ppMean) > 1e-14 {
		ratio = mmMean / ppMean
	}
	cert := planeResidual.FrobeniusNorm() < ratioTolerance && reconResidual.FrobeniusNorm() < ratioTolerance && ppSpread < ratioTolerance && mmSpread < ratioTolerance && pm.FrobeniusNorm() < ratioTolerance && math.Abs(ratio+3) < ratioTolerance
	p, m, z, _ := inertia(gTwist)
	comment := "After sign alignment with B_hat, the normalized twist metric is tested against the projector-plane ray (P_{K7+}-3P_{K7-})/sqrt(31)."
	if cert {
		comment += " The route certifies the diagonal 1:-3 Hodge-sector metric ratio."
	} else {
		comment += " The route does not certify the exact 1:-3 projector-plane ratio at tolerance."
	}
	return MetricRatioRoute{Name: name, Formula: formula, Inertia: inertiaString(p, m, z), Cosine: inner, Rho: rho, GHatToProjectorPlaneResidual: planeResidual.FrobeniusNorm(), GHatToReconstructedResidual: reconResidual.FrobeniusNorm(), PlusBlockEigenvalues: ppEig, MinusBlockEigenvalues: mmEig, PlusBlockMean: ppMean, MinusBlockMean: mmMean, PlusBlockSpread: ppSpread, MinusBlockSpread: mmSpread, PlusMinusFrobNorm: pm.FrobeniusNorm(), ObservedMinusToPlusRatio: ratio, Ratio1ToMinus3Certified: cert, Comment: comment}, nil
}

func buildAngleFromPlane(defs ProjectorPlaneDefinition) ProjectiveAngleFromPlane {
	cos := (float64(k7PlusDim)*1*1 + float64(k7MinusDim)*(-3)*(-1)) / math.Sqrt(float64(projectorDenom*k7DimExpected))
	sin2 := 1 - cos*cos
	expectedCos := float64(alignmentRoot) / math.Sqrt(float64(angleDenominator))
	expectedSin2 := float64(failureNumerator) / float64(angleDenominator)
	ok := math.Abs(cos-expectedCos) < strictTolerance && math.Abs(sin2-expectedSin2) < strictTolerance && defs.ProjectorPlaneMetricsCertified
	return ProjectiveAngleFromPlane{PlusDim: k7PlusDim, MinusDim: k7MinusDim, BHatWeights: "(1,-1)/sqrt(7)", GHatWeights: "(1,-3)/sqrt(31)", InnerProductFormula: "<G_hat,B_hat>=(4*1*1 + 3*(-3)*(-1))/sqrt(31*7)=13/sqrt(217)", ComputedCosine: cos, ExpectedCosine: expectedCos, ComputedSinSquared: sin2, ExpectedSinSquared: expectedSin2, AngleDerivedFromPlane: ok, NativeTraceIdentityFound: false, Verdict: join(StatusProjectorPlaneAngle, StatusNoNativeTraceIdentity)}
}

func buildInterpretation(inh Gate643Inheritance, ratio MetricRatioAudit, angle ProjectiveAngleFromPlane, minus MinusThreeSourceAudit) Interpretation {
	interp := "The Gate643 residual tensor lets the full normalized twist metric be reconstructed inside the two-dimensional Hodge-projector plane.  Route-by-route, G_hat is the diagonal ray (P_{K7+}-3P_{K7-})/sqrt(31), while B_hat is (P_{K7+}-P_{K7-})/sqrt(7).  This reduces the previous 13:48:217 angle to the projector-plane comparison (1,-1) versus (1,-3).  The -3 weight matches the negative Hodge-sector dimension as a candidate, but no native theorem derives that weight from the twist operator."
	return Interpretation{Gate643Inherited: inh.Verdict == StatusGate643ResidualInherited, GHatReconstructed: ratio.AllRoutesRatioCertified, RatioCertified: ratio.AllRoutesRatioCertified, AngleFromPlane: angle.AngleDerivedFromPlane, MinusThreeSourceFound: minus.CertifiedNativeSource, NativeTraceIdentityFound: false, Interpretation: interp, Verdict: join(StatusHodgeDiagonalRatio, StatusProjectorPlaneAngle, StatusNoMinusThreeSource, StatusNoNativeTraceIdentity)}
}

func Statuses() []string {
	return []string{StatusGate643ResidualInherited, StatusGHatReconstructed, StatusProjectorPlaneMetricsDefined, StatusRouteMetricRatiosComputed, StatusHodgeDiagonalRatio, StatusProjectorPlaneAngle, StatusMinusThreeSourceCandidate, StatusNoMinusThreeSource, StatusNoNativeTraceIdentity, StatusNoCertifiedSplitG2, StatusNoBoundaryStress, StatusNoSevenOver72Theorem, StatusNoScalarFlavorTransport, StatusNoPhysicalAngle, StatusNoPhysicalMetric, StatusNoHiggsFlavorGauge, StatusGate644Boundary}
}

// tensor3 stores components Ω_abc in the Q_K basis.
type tensor3 [k7DimExpected][k7DimExpected][k7DimExpected]float64

func nativeOmega0(space contact.Space) (tensor3, error) {
	coords, err := g2RawCoordinates(space.G2Support.RawColumns, space.ContactFrame)
	if err != nil {
		return tensor3{}, err
	}
	tCoords, _ := splitG2Coordinates(coords)
	phi := octonion.StandardAssociativeForm()
	return pullbackAssociativeTensor(phi, tCoords), nil
}

func g2RawCoordinates(raw, qK linear.Matrix) (linear.Matrix, error) {
	gram, err := raw.Transpose().Mul(raw)
	if err != nil {
		return linear.Matrix{}, err
	}
	inv, err := inverse(gram)
	if err != nil {
		return linear.Matrix{}, fmt.Errorf("invert raw P_G Gram: %w", err)
	}
	pinv, err := inv.Mul(raw.Transpose())
	if err != nil {
		return linear.Matrix{}, err
	}
	return pinv.Mul(qK)
}

func splitG2Coordinates(coords linear.Matrix) (linear.Matrix, linear.Matrix) {
	t := linear.NewMatrix(g2CopyDim, k7DimExpected)
	s := linear.NewMatrix(g2CopyDim, k7DimExpected)
	for r := 0; r < g2CopyDim; r++ {
		for c := 0; c < k7DimExpected; c++ {
			t.Set(r, c, coords.At(r, c))
			s.Set(r, c, coords.At(g2CopyDim+r, c))
		}
	}
	return t, s
}

func pullbackAssociativeTensor(phi octonion.AssociativeForm, coords linear.Matrix) tensor3 {
	var out tensor3
	for a := 0; a < k7DimExpected; a++ {
		for b := 0; b < k7DimExpected; b++ {
			for c := 0; c < k7DimExpected; c++ {
				sum := 0.0
				for i := 0; i < g2CopyDim; i++ {
					for j := 0; j < g2CopyDim; j++ {
						for k := 0; k < g2CopyDim; k++ {
							sum += phi.Value(i, j, k) * coords.At(i, a) * coords.At(j, b) * coords.At(k, c)
						}
					}
				}
				out[a][b][c] = sum
			}
		}
	}
	return out
}

func transformTensor(t tensor3, a, b, c linear.Matrix) tensor3 {
	var out tensor3
	for i := 0; i < k7DimExpected; i++ {
		for j := 0; j < k7DimExpected; j++ {
			for k := 0; k < k7DimExpected; k++ {
				sum := 0.0
				for p := 0; p < k7DimExpected; p++ {
					for q := 0; q < k7DimExpected; q++ {
						for r := 0; r < k7DimExpected; r++ {
							sum += a.At(p, i) * b.At(q, j) * c.At(r, k) * t[p][q][r]
						}
					}
				}
				out[i][j][k] = sum
			}
		}
	}
	return out
}

func alternateTensor(t tensor3) tensor3 {
	var out tensor3
	perms := [][]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}}
	for a := 0; a < k7DimExpected; a++ {
		for b := 0; b < k7DimExpected; b++ {
			for c := 0; c < k7DimExpected; c++ {
				idx := []int{a, b, c}
				sum := 0.0
				for _, p := range perms {
					sum += float64(paritySign(p)) * t[idx[p[0]]][idx[p[1]]][idx[p[2]]]
				}
				out[a][b][c] = sum / 6.0
			}
		}
	}
	return out
}

func compactCrossProductBKTensor(omega0 tensor3, invG, bk linear.Matrix) (tensor3, error) {
	var omegaB tensor3
	for a := 0; a < k7DimExpected; a++ {
		for b := 0; b < k7DimExpected; b++ {
			for z := 0; z < k7DimExpected; z++ {
				sum := 0.0
				for m := 0; m < k7DimExpected; m++ {
					xabm := 0.0
					for c := 0; c < k7DimExpected; c++ {
						xabm += invG.At(m, c) * omega0[a][b][c]
					}
					sum += bk.At(m, z) * xabm
				}
				omegaB[a][b][z] = sum
			}
		}
	}
	return omegaB, nil
}

func hitchinBMatrix(t tensor3) linear.Matrix {
	b := linear.NewMatrix(k7DimExpected, k7DimExpected)
	perms := permutations7()
	for i := 0; i < k7DimExpected; i++ {
		for j := 0; j < k7DimExpected; j++ {
			sum := 0.0
			for _, p := range perms {
				sum += float64(paritySign(p)) * t[i][p[0]][p[1]] * t[j][p[2]][p[3]] * t[p[4]][p[5]][p[6]]
			}
			b.Set(i, j, sum/144.0)
		}
	}
	return b
}

func hodgePolarityBases(sk linear.Matrix) (linear.Matrix, linear.Matrix, error) {
	eig, err := linear.SymmetricEigenJacobi(sk, 1e-13, 0)
	if err != nil {
		return linear.Matrix{}, linear.Matrix{}, err
	}
	qPlus := linear.NewMatrix(k7DimExpected, k7PlusDim)
	qMinus := linear.NewMatrix(k7DimExpected, k7MinusDim)
	ip, im := 0, 0
	for col, val := range eig.Values {
		if val > 0 {
			if ip >= k7PlusDim {
				return linear.Matrix{}, linear.Matrix{}, fmt.Errorf("too many positive Hodge vectors")
			}
			for r := 0; r < k7DimExpected; r++ {
				qPlus.Set(r, ip, eig.Vectors.At(r, col))
			}
			ip++
		} else if val < 0 {
			if im >= k7MinusDim {
				return linear.Matrix{}, linear.Matrix{}, fmt.Errorf("too many negative Hodge vectors")
			}
			for r := 0; r < k7DimExpected; r++ {
				qMinus.Set(r, im, eig.Vectors.At(r, col))
			}
			im++
		}
	}
	if ip != k7PlusDim || im != k7MinusDim {
		return linear.Matrix{}, linear.Matrix{}, fmt.Errorf("bad Hodge split dimensions: plus=%d minus=%d", ip, im)
	}
	return qPlus, qMinus, nil
}

func restrictedHodgeOperator(qK, star linear.Matrix) (linear.Matrix, error) {
	left, err := qK.Transpose().Mul(star)
	if err != nil {
		return linear.Matrix{}, err
	}
	return left.Mul(qK)
}

func hodgeStarLambda4R8() (linear.Matrix, error) {
	basis, err := combinatorics.Subsets(vectorDimExpected, 4)
	if err != nil {
		return linear.Matrix{}, err
	}
	index := combinatorics.IndexByKey(basis)
	star := linear.NewMatrix(len(basis), len(basis))
	for col, subset := range basis {
		comp := complementSubset(vectorDimExpected, subset)
		row, ok := index[comp.Key()]
		if !ok {
			return linear.Matrix{}, fmt.Errorf("missing complement subset %v", comp)
		}
		full := append(append([]int{}, subset...), comp...)
		star.Set(row, col, float64(paritySign(full)))
	}
	return star, nil
}

func complementSubset(n int, selected combinatorics.Subset) combinatorics.Subset {
	out := make([]int, 0, n-len(selected))
	for i := 0; i < n; i++ {
		found := false
		for _, v := range selected {
			if i == v {
				found = true
				break
			}
		}
		if !found {
			out = append(out, i)
		}
	}
	return combinatorics.Subset(out)
}

func projector(q linear.Matrix) (linear.Matrix, error) { return q.Mul(q.Transpose()) }

func sandwich(left, middle, right linear.Matrix) (linear.Matrix, error) {
	lm, err := left.Transpose().Mul(middle)
	if err != nil {
		return linear.Matrix{}, err
	}
	return lm.Mul(right)
}

func frobeniusInner(a, b linear.Matrix) float64 {
	sum := 0.0
	for r := 0; r < a.Rows(); r++ {
		for c := 0; c < a.Cols(); c++ {
			sum += a.At(r, c) * b.At(r, c)
		}
	}
	return sum
}

func sortedEigenvalues(m linear.Matrix) ([]float64, error) {
	eig, err := linear.SymmetricEigenJacobi(m, 1e-13, 0)
	if err != nil {
		return nil, err
	}
	vals := append([]float64(nil), eig.Values...)
	sort.Float64s(vals)
	return vals, nil
}

func mean(v []float64) float64 {
	if len(v) == 0 {
		return math.NaN()
	}
	s := 0.0
	for _, x := range v {
		s += x
	}
	return s / float64(len(v))
}

func maxAbsDeviation(v []float64, target float64) float64 {
	m := 0.0
	for _, x := range v {
		if d := math.Abs(x - target); d > m {
			m = d
		}
	}
	return m
}

func inertia(m linear.Matrix) (int, int, int, error) {
	eig, err := linear.SymmetricEigenJacobi(m, 1e-13, 0)
	if err != nil {
		return 0, 0, 0, err
	}
	p, q, z := 0, 0, 0
	for _, v := range eig.Values {
		if v > numericalTolerance {
			p++
		} else if v < -numericalTolerance {
			q++
		} else {
			z++
		}
	}
	return p, q, z, nil
}
func inertiaString(p, m, z int) string { return fmt.Sprintf("(%d,%d,%d)", p, m, z) }

func permutations7() [][]int {
	out := make([][]int, 0, 5040)
	cur := []int{0, 1, 2, 3, 4, 5, 6}
	var walk func(int)
	walk = func(pos int) {
		if pos == len(cur) {
			p := append([]int{}, cur...)
			out = append(out, p)
			return
		}
		for i := pos; i < len(cur); i++ {
			cur[pos], cur[i] = cur[i], cur[pos]
			walk(pos + 1)
			cur[pos], cur[i] = cur[i], cur[pos]
		}
	}
	walk(0)
	return out
}
func paritySign(values []int) int {
	inv := 0
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				inv++
			}
		}
	}
	if inv%2 == 0 {
		return 1
	}
	return -1
}

func inverse(a linear.Matrix) (linear.Matrix, error) {
	if a.Rows() != a.Cols() {
		return linear.Matrix{}, fmt.Errorf("inverse requires square matrix: %dx%d", a.Rows(), a.Cols())
	}
	n := a.Rows()
	aug := make([][]float64, n)
	for i := 0; i < n; i++ {
		aug[i] = make([]float64, 2*n)
		for j := 0; j < n; j++ {
			aug[i][j] = a.At(i, j)
		}
		aug[i][n+i] = 1
	}
	for col := 0; col < n; col++ {
		pivot := col
		maxAbs := math.Abs(aug[col][col])
		for r := col + 1; r < n; r++ {
			if v := math.Abs(aug[r][col]); v > maxAbs {
				maxAbs = v
				pivot = r
			}
		}
		if maxAbs < 1e-14 {
			return linear.Matrix{}, fmt.Errorf("singular matrix at pivot %d", col)
		}
		aug[col], aug[pivot] = aug[pivot], aug[col]
		pv := aug[col][col]
		for j := 0; j < 2*n; j++ {
			aug[col][j] /= pv
		}
		for r := 0; r < n; r++ {
			if r == col {
				continue
			}
			f := aug[r][col]
			for j := 0; j < 2*n; j++ {
				aug[r][j] -= f * aug[col][j]
			}
		}
	}
	out := linear.NewMatrix(n, n)
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			out.Set(r, c, aug[r][n+c])
		}
	}
	return out, nil
}

func join(parts ...string) string { return strings.Join(parts, "; ") }
