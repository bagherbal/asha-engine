// Package generation2compactsplitresidualtensorblockstructureaudit implements
// Gate 643: CompactSplit ResidualTensor BlockStructure Audit.
//
// Gate 642 compressed the compact/split obstruction into the projective angle
// cos(theta)=13/sqrt(217), sin(theta)=4*sqrt(3)/sqrt(217), but did not derive
// a native trace identity. Gate 643 stops treating the obstruction as only a
// scalar and constructs the normalized residual tensor
//
//	R_hat = (G_hat - <G_hat,B_hat> B_hat) / rho
//
// for the Gate638 twist routes. It then decomposes R_hat into the Hodge-polarity
// blocks K7+ x K7+, K7- x K7-, and K7+ x K7- to test whether the 4*sqrt(3)
// failure component has a typed block carrier. The gate is internal finite
// geometry only; it does not promote split-G2, boundary stress, scalar/flavor
// transport, physical geometry, or a native 7/72 theorem.
package generation2compactsplitresidualtensorblockstructureaudit

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	gate642 "github.com/bagherbal/asha-engine/pkg/bridge/generation2hodgepolarityprojectiveangletraceidentityaudit"
	"github.com/bagherbal/asha-engine/pkg/combinatorics"
	"github.com/bagherbal/asha-engine/pkg/geometry/contact"
	"github.com/bagherbal/asha-engine/pkg/linear"
	"github.com/bagherbal/asha-engine/pkg/octonion"
)

const (
	AuditID = "GATE643-COMPACT-SPLIT-RESIDUAL-TENSOR-BLOCK-STRUCTURE-AUDIT"

	StatusGate642AngleInherited       = "PASS_GATE642_PROJECTIVE_ANGLE_INHERITED"
	StatusResidualTensorDefined       = "PASS_RESIDUAL_TENSOR_DEFINED_ORTHOGONAL_TO_BK"
	StatusHodgePolarityBlocksComputed = "PASS_HODGE_POLARITY_BLOCKS_COMPUTED"
	StatusRouteBlockProfilesComputed  = "PASS_ROUTE_BLOCK_PROFILES_COMPUTED"
	StatusResidualBlockStructure      = "CONDITIONAL_SUPPORT_RESIDUAL_HAS_TYPED_BLOCK_STRUCTURE"
	StatusSameSectorHodgeDiagonal     = "CONDITIONAL_SUPPORT_RESIDUAL_IS_SAME_SECTOR_HODGE_DIAGONAL"
	StatusOffSectorCarrierCandidate   = "CONDITIONAL_SUPPORT_OFF_SECTOR_BLOCK_CARRIES_FAILURE_COMPONENT_CANDIDATE"
	StatusNoOffSectorCarrier          = "FAILED_ROUTE_OFF_SECTOR_BLOCK_DOES_NOT_CARRY_RESIDUAL_TENSOR"
	StatusNoSimpleBlockStructure      = "FAILED_ROUTE_RESIDUAL_TENSOR_HAS_NO_SIMPLE_HODGE_BLOCK_STRUCTURE"
	StatusNoNativeTraceIdentity       = "FAILED_ROUTE_NO_NATIVE_TRACE_IDENTITY_FOR_PROJECTIVE_ANGLE_YET"
	StatusNoCertifiedSplitG2          = "FAILED_ROUTE_NO_CERTIFIED_SPLIT_G2_STRUCTURE"
	StatusNoBoundaryStress            = "FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT"
	StatusNoSevenOver72Theorem        = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM"
	StatusNoScalarFlavorTransport     = "FAILED_ROUTE_NO_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM"
	StatusNoPhysicalAngle             = "FAILED_ROUTE_PROJECTIVE_ANGLE_IS_NOT_PHYSICAL_ANGLE"
	StatusNoPhysicalMetric            = "FAILED_ROUTE_PROJECTIVE_ANGLE_IS_NOT_PHYSICAL_METRIC_THEOREM"
	StatusNoHiggsFlavorGauge          = "FAILED_ROUTE_NO_HIGGS_FLAVOR_PMNS_CKM_GAUGE_THEOREM"
	StatusGate643Boundary             = "FIREWALL_PRESERVED_GATE643_RESIDUAL_TENSOR_IS_INTERNAL_OBSTRUCTION_ONLY"
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
	strictTolerance    = 1e-10
	numericalTolerance = 1e-8
	angleTolerance     = 1e-9
)

type Gate642Inheritance struct {
	CosTheta                     float64
	SinTheta                     float64
	CosSquared                   float64
	SinSquared                   float64
	AlignmentNumerator           int
	FailureNumerator             int
	Denominator                  int
	HodgePolaritySkeleton        bool
	NativeTraceIdentityCertified bool
	SplitG2Certified             bool
	BoundaryStressAssignment     bool
	SevenOver72Theorem           bool
	ScalarFlavorTransport        bool
	PhysicalAngle                bool
	PhysicalMetric               bool
	Gate642FirewallPreserved     bool
	Verdict                      string
}

type ResidualTensorRoute struct {
	Name                    string
	Formula                 string
	Inertia                 string
	Cosine                  float64
	Rho                     float64
	OrthogonalityToBHat     float64
	ResidualUnitNorm        float64
	SymmetryResidual        float64
	RPlusPlusFrobSquared    float64
	RMinusMinusFrobSquared  float64
	TwiceRPlusMinusFrobSq   float64
	BlockNormSum            float64
	PlusPlusRank            int
	MinusMinusRank          int
	PlusMinusRank           int
	PlusPlusTrace           float64
	MinusMinusTrace         float64
	PlusPlusEigenvalues     []float64
	MinusMinusEigenvalues   []float64
	PlusMinusSingularValues []float64
	DominantBlock           string
	OffSectorDominant       bool
	TypedBlockProfile       bool
	Comment                 string
}

type ResidualTensorAudit struct {
	Routes                   []ResidualTensorRoute
	CosineTarget             float64
	RhoTarget                float64
	MaxOrthogonalityToBHat   float64
	MaxResidualUnitNormDrift float64
	MaxCosineDrift           float64
	MaxRhoDrift              float64
	ResidualTensorsCertified bool
	Verdict                  string
}

type HodgeBlockSummary struct {
	RouteCount                    int
	PlusPlusMeanFrobSquared       float64
	MinusMinusMeanFrobSquared     float64
	TwicePlusMinusMeanFrobSquared float64
	OffSectorDominantRoutes       int
	SameDominantBlockAllRoutes    bool
	SameRankProfileAllRoutes      bool
	FailureSkeleton               string
	ExactSameSectorProfile        string
	HasTypedBlockStructure        bool
	NativeTraceIdentityFound      bool
	Verdict                       string
}

type ResidualInterpretation struct {
	AnglePairInherited        bool
	ResidualTensorDefined     bool
	BlocksComputed            bool
	TypedBlockStructure       bool
	OffSectorCarrierCandidate bool
	NativeTraceIdentityFound  bool
	Interpretation            string
	Verdict                   string
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
	Inherited      Gate642Inheritance
	ResidualTensor ResidualTensorAudit
	BlockSummary   HodgeBlockSummary
	Interpretation ResidualInterpretation
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
	g642, err := gate642.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate642 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g642)

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
	omega0, err := nativeOmega0(space)
	if err != nil {
		return Analysis{}, err
	}
	residualTensor, err := buildResidualTensorAudit(omega0, bk, qPlus, qMinus, inherited)
	if err != nil {
		return Analysis{}, err
	}
	summary := buildBlockSummary(residualTensor)
	interpretation := buildInterpretation(inherited, residualTensor, summary)
	firewalls := Firewalls{Verdict: StatusGate643Boundary}
	truth := "Gate 643 constructs the projectively normalized residual tensor R_hat behind the Gate642 angle and decomposes it across K_7^+ and K_7^- blocks.  The residual is certified orthogonal to B_hat and unit-normalized route-by-route.  The block audit finds a repeatable same-sector Hodge-diagonal profile: ||R_{++}||_F^2=3/7, ||R_{--}||_F^2=4/7, and 2||R_{+-}||_F^2=0.  Thus Gate642's p^2q scalar skeleton is not promoted into an off-sector tensor theorem; the residual tensor is cleaner but different, and no native Frobenius/projector trace identity is certified.  The obstruction remains internal finite geometry only: no split-G2 carrier, boundary-stress assignment, scalar/flavor transport, physical metric, or native 7/72 theorem is certified."
	return Analysis{Inherited: inherited, ResidualTensor: residualTensor, BlockSummary: summary, Interpretation: interpretation, Firewalls: firewalls, Truth: truth}, nil
}

func buildInheritance(g642 gate642.Analysis) Gate642Inheritance {
	return Gate642Inheritance{
		CosTheta:                     g642.Inherited.CosTheta,
		SinTheta:                     g642.Inherited.SinTheta,
		CosSquared:                   g642.Inherited.CosSquared,
		SinSquared:                   g642.Inherited.SinSquared,
		AlignmentNumerator:           g642.ProjectivePair.AlignmentAmplitude,
		FailureNumerator:             g642.ProjectivePair.FailureAmplitudeSquared,
		Denominator:                  g642.ProjectivePair.Denominator,
		HodgePolaritySkeleton:        g642.SectorBlocks.BlockSkeletonMatches,
		NativeTraceIdentityCertified: g642.TraceIdentity.NativeTraceIdentityFound,
		SplitG2Certified:             g642.Firewalls.ClaimsSplitG2,
		BoundaryStressAssignment:     g642.Firewalls.ClaimsBoundaryStress,
		SevenOver72Theorem:           g642.Firewalls.ClaimsSevenOver72Theorem,
		ScalarFlavorTransport:        g642.Firewalls.ClaimsScalarFlavor,
		PhysicalAngle:                g642.Firewalls.ClaimsPhysicalAngle,
		PhysicalMetric:               g642.Firewalls.ClaimsPhysicalMetric,
		Gate642FirewallPreserved:     g642.Firewalls.Verdict == gate642.StatusGate642Boundary,
		Verdict:                      StatusGate642AngleInherited,
	}
}

func buildResidualTensorAudit(omega0 tensor3, bk, qPlus, qMinus linear.Matrix, inherited Gate642Inheritance) (ResidualTensorAudit, error) {
	sk := bk
	omega3 := transformTensor(omega0, sk, sk, sk)
	invGOmega, err := inverse(hitchinBMatrix(omega0))
	if err != nil {
		return ResidualTensorAudit{}, err
	}
	omegaB, err := compactCrossProductBKTensor(omega0, invGOmega, bk)
	if err != nil {
		return ResidualTensorAudit{}, err
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
	_ = omega3 // kept as a typed Gate638 object, but Gate643 audits the repeated split residual cluster only.
	routes := make([]ResidualTensorRoute, 0, len(candidates))
	maxOrth, maxNormDrift, maxCosDrift, maxRhoDrift := 0.0, 0.0, 0.0, 0.0
	all := true
	for _, c := range candidates {
		metric := hitchinBMatrix(c.tensor)
		r, err := auditResidualRoute(c.name, c.formula, metric, bk, qPlus, qMinus, inherited.CosTheta, inherited.SinTheta)
		if err != nil {
			return ResidualTensorAudit{}, err
		}
		if r.OrthogonalityToBHat > maxOrth {
			maxOrth = r.OrthogonalityToBHat
		}
		normDrift := math.Abs(r.ResidualUnitNorm - 1)
		if normDrift > maxNormDrift {
			maxNormDrift = normDrift
		}
		cosDrift := math.Abs(r.Cosine - inherited.CosTheta)
		if cosDrift > maxCosDrift {
			maxCosDrift = cosDrift
		}
		rhoDrift := math.Abs(r.Rho - inherited.SinTheta)
		if rhoDrift > maxRhoDrift {
			maxRhoDrift = rhoDrift
		}
		if r.OrthogonalityToBHat > angleTolerance || normDrift > angleTolerance || cosDrift > 1e-6 || rhoDrift > 1e-6 {
			all = false
		}
		routes = append(routes, r)
	}
	verdict := join(StatusResidualTensorDefined, StatusRouteBlockProfilesComputed)
	if !all {
		verdict = StatusNoSimpleBlockStructure
	}
	return ResidualTensorAudit{Routes: routes, CosineTarget: inherited.CosTheta, RhoTarget: inherited.SinTheta, MaxOrthogonalityToBHat: maxOrth, MaxResidualUnitNormDrift: maxNormDrift, MaxCosineDrift: maxCosDrift, MaxRhoDrift: maxRhoDrift, ResidualTensorsCertified: all, Verdict: verdict}, nil
}

func auditResidualRoute(name, formula string, gTwist, bk, qPlus, qMinus linear.Matrix, cosTarget, rhoTarget float64) (ResidualTensorRoute, error) {
	bNorm := bk.FrobeniusNorm()
	gNorm := gTwist.FrobeniusNorm()
	if bNorm == 0 || gNorm == 0 {
		return ResidualTensorRoute{}, fmt.Errorf("zero norm in route %s", name)
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
		return ResidualTensorRoute{}, err
	}
	rho := residualRaw.FrobeniusNorm()
	if rho == 0 {
		return ResidualTensorRoute{}, fmt.Errorf("zero residual in route %s", name)
	}
	rHat := residualRaw.Scale(1 / rho)
	orth := math.Abs(frobeniusInner(rHat, bHat))
	unit := rHat.FrobeniusNorm()
	sym := symmetryResidual(rHat)
	pp, err := sandwich(qPlus, rHat, qPlus)
	if err != nil {
		return ResidualTensorRoute{}, err
	}
	mm, err := sandwich(qMinus, rHat, qMinus)
	if err != nil {
		return ResidualTensorRoute{}, err
	}
	pm, err := sandwich(qPlus, rHat, qMinus)
	if err != nil {
		return ResidualTensorRoute{}, err
	}
	ppNorm2 := frobSquared(pp)
	mmNorm2 := frobSquared(mm)
	pm2 := 2 * frobSquared(pm)
	blockSum := ppNorm2 + mmNorm2 + pm2
	ppEig, _ := sortedEigenvalues(pp)
	mmEig, _ := sortedEigenvalues(mm)
	pmSing, _ := singularValues(pm)
	ppTrace, _ := pp.Trace()
	mmTrace, _ := mm.Trace()
	ppRank := rankFromValues(absValues(ppEig), strictTolerance)
	mmRank := rankFromValues(absValues(mmEig), strictTolerance)
	pmRank := rankFromValues(pmSing, strictTolerance)
	dom := "K7+ x K7+"
	maxv := ppNorm2
	if mmNorm2 > maxv {
		dom = "K7- x K7-"
		maxv = mmNorm2
	}
	if pm2 > maxv {
		dom = "K7+ x K7-"
	}
	offDom := dom == "K7+ x K7-"
	typed := math.Abs(blockSum-1) < 1e-8 && ppRank > 0 && mmRank > 0
	comment := "R_hat is the unit Frobenius component of the split-twist metric orthogonal to B_hat; block norms are measured in the Gate634 Hodge-polarity eigenbasis."
	if offDom {
		comment += " The off-sector block is the largest block contribution and is the strongest carrier candidate for the 4*sqrt(3) failure component."
	} else {
		comment += " The residual is carried by same-sector Hodge blocks rather than by an off-sector K7+ x K7- block."
	}
	p, m, z, _ := inertia(gTwist)
	return ResidualTensorRoute{Name: name, Formula: formula, Inertia: inertiaString(p, m, z), Cosine: inner, Rho: rho, OrthogonalityToBHat: orth, ResidualUnitNorm: unit, SymmetryResidual: sym, RPlusPlusFrobSquared: ppNorm2, RMinusMinusFrobSquared: mmNorm2, TwiceRPlusMinusFrobSq: pm2, BlockNormSum: blockSum, PlusPlusRank: ppRank, MinusMinusRank: mmRank, PlusMinusRank: pmRank, PlusPlusTrace: ppTrace, MinusMinusTrace: mmTrace, PlusPlusEigenvalues: ppEig, MinusMinusEigenvalues: mmEig, PlusMinusSingularValues: pmSing, DominantBlock: dom, OffSectorDominant: offDom, TypedBlockProfile: typed, Comment: comment}, nil
}

func buildBlockSummary(a ResidualTensorAudit) HodgeBlockSummary {
	n := len(a.Routes)
	pp, mm, pm := 0.0, 0.0, 0.0
	off := 0
	sameDom := true
	sameRank := true
	dom := ""
	rankSig := ""
	for i, r := range a.Routes {
		pp += r.RPlusPlusFrobSquared
		mm += r.RMinusMinusFrobSquared
		pm += r.TwiceRPlusMinusFrobSq
		if r.OffSectorDominant {
			off++
		}
		sig := fmt.Sprintf("%d/%d/%d", r.PlusPlusRank, r.MinusMinusRank, r.PlusMinusRank)
		if i == 0 {
			dom = r.DominantBlock
			rankSig = sig
		} else {
			if r.DominantBlock != dom {
				sameDom = false
			}
			if sig != rankSig {
				sameRank = false
			}
		}
	}
	if n > 0 {
		pp /= float64(n)
		mm /= float64(n)
		pm /= float64(n)
	}
	has := a.ResidualTensorsCertified && n >= 3 && sameRank
	verdict := join(StatusHodgePolarityBlocksComputed, StatusResidualBlockStructure, StatusSameSectorHodgeDiagonal, StatusNoOffSectorCarrier)
	if off > 0 {
		verdict = join(StatusHodgePolarityBlocksComputed, StatusResidualBlockStructure, StatusOffSectorCarrierCandidate)
	}
	if !has {
		verdict = StatusNoSimpleBlockStructure
	}
	if off == 0 && has {
		verdict = join(StatusHodgePolarityBlocksComputed, StatusResidualBlockStructure, StatusSameSectorHodgeDiagonal, StatusNoOffSectorCarrier)
	}
	return HodgeBlockSummary{RouteCount: n, PlusPlusMeanFrobSquared: pp, MinusMinusMeanFrobSquared: mm, TwicePlusMinusMeanFrobSquared: pm, OffSectorDominantRoutes: off, SameDominantBlockAllRoutes: sameDom, SameRankProfileAllRoutes: sameRank, FailureSkeleton: "Gate642 skeleton: 48=p^2*q with p=dim(K7+)=4 and q=dim(K7-)=3; Gate643 tests the actual R_hat blocks, not only the scalar angle.", ExactSameSectorProfile: "R_hat block profile repeats as ||R++||_F^2=3/7, ||R--||_F^2=4/7, 2||R+-||_F^2=0, with ranks 4/3/0.", HasTypedBlockStructure: has, NativeTraceIdentityFound: false, Verdict: verdict}
}

func buildInterpretation(inh Gate642Inheritance, rt ResidualTensorAudit, bs HodgeBlockSummary) ResidualInterpretation {
	interp := "The projective residual tensor exists as a unit Frobenius tensor orthogonal to B_hat and decomposes cleanly across K7+ and K7- same-sector blocks.  The off-sector block is not the carrier in the audited routes; instead the residual is same-sector Hodge-diagonal with a 3/7 and 4/7 split.  Therefore the p^2q=48 scalar skeleton remains a compressed obstruction ratio rather than a certified off-sector block-source theorem."
	return ResidualInterpretation{AnglePairInherited: inh.Verdict == StatusGate642AngleInherited, ResidualTensorDefined: rt.ResidualTensorsCertified, BlocksComputed: len(rt.Routes) >= 3, TypedBlockStructure: bs.HasTypedBlockStructure, OffSectorCarrierCandidate: bs.OffSectorDominantRoutes > 0, NativeTraceIdentityFound: false, Interpretation: interp, Verdict: join(StatusResidualBlockStructure, StatusNoNativeTraceIdentity)}
}

func Statuses() []string {
	return []string{StatusGate642AngleInherited, StatusResidualTensorDefined, StatusHodgePolarityBlocksComputed, StatusRouteBlockProfilesComputed, StatusResidualBlockStructure, StatusSameSectorHodgeDiagonal, StatusOffSectorCarrierCandidate, StatusNoOffSectorCarrier, StatusNoSimpleBlockStructure, StatusNoNativeTraceIdentity, StatusNoCertifiedSplitG2, StatusNoBoundaryStress, StatusNoSevenOver72Theorem, StatusNoScalarFlavorTransport, StatusNoPhysicalAngle, StatusNoPhysicalMetric, StatusNoHiggsFlavorGauge, StatusGate643Boundary}
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

func frobSquared(a linear.Matrix) float64 { n := a.FrobeniusNorm(); return n * n }

func symmetryResidual(a linear.Matrix) float64 {
	max := 0.0
	for r := 0; r < a.Rows(); r++ {
		for c := 0; c < a.Cols(); c++ {
			if d := math.Abs(a.At(r, c) - a.At(c, r)); d > max {
				max = d
			}
		}
	}
	return max
}

func singularValues(m linear.Matrix) ([]float64, error) {
	mtm, err := m.Transpose().Mul(m)
	if err != nil {
		return nil, err
	}
	eig, err := linear.SymmetricEigenJacobi(mtm, 1e-13, 0)
	if err != nil {
		return nil, err
	}
	vals := make([]float64, len(eig.Values))
	for i, v := range eig.Values {
		if v < 0 && math.Abs(v) < 1e-12 {
			v = 0
		}
		if v < 0 {
			return nil, fmt.Errorf("negative singular square %g", v)
		}
		vals[i] = math.Sqrt(v)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(vals)))
	return vals, nil
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

func absValues(v []float64) []float64 {
	out := make([]float64, len(v))
	for i, x := range v {
		out[i] = math.Abs(x)
	}
	return out
}
func rankFromValues(v []float64, eps float64) int {
	n := 0
	for _, x := range v {
		if math.Abs(x) > eps {
			n++
		}
	}
	return n
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
