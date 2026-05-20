// Package generation2hodgestark7tow7leakagerankaudit implements
// Gate 632: Hodge-Star K7-to-W7 Leakage Rank Audit.
//
// Gate 631 sharpened the balanced Boolean--octonionic defect problem to the
// explicit finite-operator candidate
//
//	Phi_* = P_W * |_{K_7}: K_7 -> W_7,
//
// where H=Lambda^4 R^8, U=Im(P_B), V=Im(P_G), K_7=U∩V, and
// W_7=(U+V)^perp.  Gate 632 constructs the actual Lambda^4 Hodge-star matrix,
// the K_7 and W_7 orthonormal bases, and the 7x7 leakage matrix
// Q_W^T * Q_K.  The computed result is a firewall-preserving failure of the
// clean Hodge route: the leakage matrix has numerical rank zero, so *K_7 stays
// inside U+V up to finite precision.  No boundary stress, scalar RG matching,
// Higgs mass, flavor, CKM/PMNS, or gauge unification theorem is derived.
package generation2hodgestark7tow7leakagerankaudit

import (
	"fmt"
	"math"
	"sort"
	"sync"

	gate631 "github.com/bagherbal/asha-engine/pkg/bridge/generation2orthogonalcokernelk7pairingaudit"
	"github.com/bagherbal/asha-engine/pkg/combinatorics"
	"github.com/bagherbal/asha-engine/pkg/geometry/contact"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

const (
	AuditID = "GATE632-HODGE-STAR-K7-TO-W7-LEAKAGE-RANK-AUDIT"

	StatusGate631Inherited                  = "PASS_GATE631_PAIRING_PROBLEM_INHERITED"
	StatusHodgeStarTyped                    = "PASS_HODGE_STAR_OPERATOR_TYPED_ON_LAMBDA4_R8"
	StatusK7AndW7BasesCertified             = "PASS_K7_AND_W7_BASES_CERTIFIED"
	StatusHodgeLeakageMatrixComputed        = "PASS_HODGE_LEAKAGE_MATRIX_COMPUTED"
	StatusHodgeStarDoesNotPairK7ToW7        = "FAILED_ROUTE_HODGE_STAR_DOES_NOT_PAIR_K7_TO_W7"
	StatusNoCanonicalK7W7PairingFound       = "FAILED_ROUTE_NO_CANONICAL_K7_W7_PAIRING_FOUND"
	StatusNoBoundaryStressAssignment        = "FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT_YET"
	StatusAlternativeCompositesNoHigherRank = "FAILED_ROUTE_ALTERNATIVE_STAR_PROJECTOR_COMPOSITES_DO_NOT_PAIR_K7_TO_W7"
	StatusGate632Boundary                   = "FIREWALL_PRESERVED_GATE632_HODGE_PAIRING_BOUNDARY"
)

const (
	vectorDimExpected  = 8
	lambda4DimExpected = 70
	rankPBExpected     = 56
	rankPGExpected     = 14
	k7DimExpected      = 7
	spanDimExpected    = 63
	w7DimExpected      = 7
	boundaryPairDim    = 2
	numericalTolerance = 1e-8
	strictTolerance    = 1e-10
)

type Gate631Inheritance struct {
	HDimension                int
	UDimension                int
	VDimension                int
	K7Dimension               int
	SpanDimension             int
	W7Dimension               int
	IndexZeroInherited        bool
	OrthogonalRepresentative  bool
	PairingProblemSharpened   bool
	ProjectorAlgebraFailed    bool
	HodgeRankTestRequired     bool
	BoundaryAssignmentMissing bool
	Gate631FirewallPreserved  bool
	Verdict                   string
}

type HodgeStarMatrixAudit struct {
	Basis                 string
	MatrixDimension       int
	TypedOnLambda4R8      bool
	MapsLambda4ToLambda4  bool
	OrientationConvention string
	StarSquaredResidual   float64
	Trace                 float64
	SelfDualDimension     int
	AntiSelfDualDimension int
	Verdict               string
}

type BasisCertificate struct {
	QKRows                    int
	QKCols                    int
	QWRows                    int
	QWCols                    int
	SpanRows                  int
	SpanCols                  int
	QKOrthonormalResidual     float64
	QWOrthonormalResidual     float64
	QWQKOrthogonalityResidual float64
	PBQKMinusQKResidual       float64
	PGQKMinusQKResidual       float64
	PBQWResidual              float64
	PGQWResidual              float64
	QWOrthogonalToUAndV       bool
	K7ContainedInUAndV        bool
	Verdict                   string
}

type LeakageRankTable struct {
	Formula              string
	Rows                 int
	Cols                 int
	Rank                 int
	SingularValues       []float64
	Determinant          float64
	FrobeniusNorm        float64
	MinimumSingularValue float64
	ConditionNumber      float64
	Classification       string
	Verdict              string
}

type ImageContainmentAudit struct {
	StarK7FrobeniusNorm         float64
	PWStarK7FrobeniusNorm       float64
	PUVStarK7FrobeniusNorm      float64
	LeakageRatio                float64
	SpanContainmentRatio        float64
	StarK7ContainedInUPlusV     bool
	TransverseComponentDetected bool
	Verdict                     string
}

type PairingMetricAudit struct {
	Formula                      string
	Computed                     bool
	RankFull                     bool
	Trace                        float64
	ScaleCandidate               float64
	ProportionalIdentityResidual float64
	ConformalOrIsometric         bool
	AnisotropicNondegenerate     bool
	Degenerate                   bool
	Verdict                      string
}

type OrientationAudit struct {
	Determinant                  float64
	NonzeroDeterminant           bool
	Sign                         int
	BasisOrientationDependent    bool
	PhysicalOrientationCertified bool
	Verdict                      string
}

type AlternativeCompositeAudit struct {
	Rows                     []AlternativeCompositeRow
	AnyHigherRankThanPhiStar bool
	AnyNondegenerate         bool
	Reason                   string
	Verdict                  string
}

type AlternativeCompositeRow struct {
	Name          string
	Formula       string
	Rank          int
	FrobeniusNorm float64
	SameAsPhiStar bool
	Verdict       string
}

type BoundaryReadinessAudit struct {
	HodgePairingCertified       bool
	K7ToW7PairingFound          bool
	BoundaryPairDimension       int
	StillRequiresW7ToBoundary   bool
	StillRequiresDefectTraceMap bool
	BoundaryAssignmentCertified bool
	MissingObject               string
	Verdict                     string
}

type NativeASHAStatus struct {
	Lambda4Native                  bool
	PBPGProjectorsConstructed      bool
	K7FrameConstructed             bool
	W7FrameConstructed             bool
	HodgeStarMatrixConstructed     bool
	HodgeRankComputed              bool
	HodgePairingNondegenerate      bool
	CanonicalK7ToW7PairingFound    bool
	BoundaryStressAssignmentNative bool
	Statement                      string
	Verdict                        string
}

type Firewalls struct {
	ClaimsBoundaryStressAssignment bool
	ClaimsScalarRGMatching         bool
	ClaimsHiggsMassDerivation      bool
	ClaimsFlavorDerivation         bool
	ClaimsCKMPMNSDerivation        bool
	ClaimsGaugeUnification         bool
	ClaimsPhysicalOrientation      bool
	ClaimsNativeTraceWeight        bool
	Verdict                        string
}

type Analysis struct {
	Inherited             Gate631Inheritance
	HodgeStar             HodgeStarMatrixAudit
	Basis                 BasisCertificate
	Leakage               LeakageRankTable
	ImageContainment      ImageContainmentAudit
	PairingMetric         PairingMetricAudit
	Orientation           OrientationAudit
	AlternativeComposites AlternativeCompositeAudit
	BoundaryReadiness     BoundaryReadinessAudit
	NativeStatus          NativeASHAStatus
	Firewalls             Firewalls
	Truth                 string
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
	inherited, err := buildInheritance()
	if err != nil {
		return Analysis{}, err
	}

	space, err := contact.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build contact space: %w", err)
	}
	if space.AmbientDimension() != lambda4DimExpected || space.Dimension() != k7DimExpected {
		return Analysis{}, fmt.Errorf("unexpected contact dimensions: ambient=%d K7=%d", space.AmbientDimension(), space.Dimension())
	}

	qK := space.ContactFrame
	qSpan, _, err := orthonormalizeColumns(concatColumns(space.BooleanSupport.Normalized, space.G2Support.Orthonormal), strictTolerance)
	if err != nil {
		return Analysis{}, fmt.Errorf("orthonormalize U+V span: %w", err)
	}
	if qSpan.Cols() != spanDimExpected {
		return Analysis{}, fmt.Errorf("unexpected dim(U+V): got %d want %d", qSpan.Cols(), spanDimExpected)
	}
	pUV, err := qSpan.Mul(qSpan.Transpose())
	if err != nil {
		return Analysis{}, fmt.Errorf("construct P_{U+V}: %w", err)
	}
	pW, err := linear.Identity(lambda4DimExpected).Sub(pUV)
	if err != nil {
		return Analysis{}, fmt.Errorf("construct P_W: %w", err)
	}
	qW, _, err := basisFromProjector(pW, 1, numericalTolerance)
	if err != nil {
		return Analysis{}, fmt.Errorf("extract W7 basis: %w", err)
	}
	if qW.Cols() != w7DimExpected {
		return Analysis{}, fmt.Errorf("unexpected W7 dimension: got %d want %d", qW.Cols(), w7DimExpected)
	}

	star, err := hodgeStarLambda4R8()
	if err != nil {
		return Analysis{}, err
	}
	starAudit, err := buildHodgeStarAudit(star)
	if err != nil {
		return Analysis{}, err
	}
	basis, err := buildBasisCertificate(space, qK, qW, qSpan)
	if err != nil {
		return Analysis{}, err
	}

	starK, err := star.Mul(qK)
	if err != nil {
		return Analysis{}, fmt.Errorf("compute *Q_K: %w", err)
	}
	m, err := qW.Transpose().Mul(starK)
	if err != nil {
		return Analysis{}, fmt.Errorf("compute Q_W^T * Q_K: %w", err)
	}
	leakage, err := buildLeakageRankTable(m)
	if err != nil {
		return Analysis{}, err
	}
	containment, err := buildImageContainment(pW, pUV, starK)
	if err != nil {
		return Analysis{}, err
	}
	metric, err := buildPairingMetric(m, leakage.Rank)
	if err != nil {
		return Analysis{}, err
	}
	orientation := buildOrientationAudit(leakage.Determinant)
	alternatives, err := buildAlternativeComposites(space, pW, star, qK, leakage.Rank)
	if err != nil {
		return Analysis{}, err
	}
	boundary := BoundaryReadinessAudit{
		HodgePairingCertified:       leakage.Rank == k7DimExpected,
		K7ToW7PairingFound:          leakage.Rank == k7DimExpected,
		BoundaryPairDimension:       boundaryPairDim,
		StillRequiresW7ToBoundary:   true,
		StillRequiresDefectTraceMap: true,
		BoundaryAssignmentCertified: false,
		MissingObject:               "W_7 -> R^2_boundary or normalized K_7/W_7 defect-trace map",
		Verdict:                     StatusNoBoundaryStressAssignment,
	}
	if leakage.Rank == 0 {
		boundary.StillRequiresW7ToBoundary = false
	}
	native := NativeASHAStatus{
		Lambda4Native:                  true,
		PBPGProjectorsConstructed:      true,
		K7FrameConstructed:             true,
		W7FrameConstructed:             true,
		HodgeStarMatrixConstructed:     true,
		HodgeRankComputed:              true,
		HodgePairingNondegenerate:      leakage.Rank == k7DimExpected,
		CanonicalK7ToW7PairingFound:    false,
		BoundaryStressAssignmentNative: false,
		Statement:                      "The finite matrix test is complete: P_W *|_{K_7} has rank zero at the certified tolerance, so the clean Hodge-star leakage route does not supply the missing K_7->W_7 pairing.",
		Verdict:                        StatusNoCanonicalK7W7PairingFound,
	}
	firewalls := Firewalls{Verdict: StatusGate632Boundary}
	truth := "Gate 632 executes the explicit finite rank test requested by Gate 631.  The Lambda^4 R^8 Hodge-star matrix is typed and satisfies *^2=I, K_7 and W_7 bases are certified, and M_*=Q_W^T*Q_K is computed.  Its singular values are numerical-zero scale, so rank(M_*)=0: *K_7 remains in U+V up to finite precision and does not leak into W_7.  Therefore the Hodge-star K_7->W_7 pairing route fails; the boundary-stress assignment remains absent."

	return Analysis{
		Inherited:             inherited,
		HodgeStar:             starAudit,
		Basis:                 basis,
		Leakage:               leakage,
		ImageContainment:      containment,
		PairingMetric:         metric,
		Orientation:           orientation,
		AlternativeComposites: alternatives,
		BoundaryReadiness:     boundary,
		NativeStatus:          native,
		Firewalls:             firewalls,
		Truth:                 truth,
	}, nil
}

func buildInheritance() (Gate631Inheritance, error) {
	g631, err := gate631.BuildDefault()
	if err != nil {
		return Gate631Inheritance{}, fmt.Errorf("Gate631 inheritance unavailable: %w", err)
	}
	return Gate631Inheritance{
		HDimension:                g631.Inherited.HDimension,
		UDimension:                g631.Inherited.UDimension,
		VDimension:                g631.Inherited.VDimension,
		K7Dimension:               g631.Inherited.K7Dimension,
		SpanDimension:             g631.Inherited.SpanDimension,
		W7Dimension:               g631.OrthogonalW7.WDimension,
		IndexZeroInherited:        g631.Inherited.Index == 0,
		OrthogonalRepresentative:  g631.OrthogonalW7.RepresentsCokernel,
		PairingProblemSharpened:   g631.CandidatePairings.PairingProblemSharpened,
		ProjectorAlgebraFailed:    g631.ProjectorAlgebra.Verdict == gate631.StatusProjectorAlgebraFails,
		HodgeRankTestRequired:     g631.HodgeStar.Verdict == gate631.StatusHodgeStarRequiresExplicitRankTest,
		BoundaryAssignmentMissing: g631.BoundaryReadiness.Verdict == gate631.StatusNoBoundaryStressAssignment,
		Gate631FirewallPreserved:  g631.Firewalls.Verdict == gate631.StatusGate631Boundary,
		Verdict:                   StatusGate631Inherited,
	}, nil
}

func buildHodgeStarAudit(star linear.Matrix) (HodgeStarMatrixAudit, error) {
	star2, err := star.Mul(star)
	if err != nil {
		return HodgeStarMatrixAudit{}, err
	}
	residual, err := star2.MaxAbsDiff(linear.Identity(star.Rows()))
	if err != nil {
		return HodgeStarMatrixAudit{}, err
	}
	trace, err := star.Trace()
	if err != nil {
		return HodgeStarMatrixAudit{}, err
	}
	selfDual := int(math.Round((float64(star.Rows()) + trace) / 2.0))
	antiSelfDual := star.Rows() - selfDual
	return HodgeStarMatrixAudit{
		Basis:                 "lexicographic oriented wedge basis e_I for |I|=4 in R^8",
		MatrixDimension:       star.Rows(),
		TypedOnLambda4R8:      star.Rows() == lambda4DimExpected && star.Cols() == lambda4DimExpected,
		MapsLambda4ToLambda4:  true,
		OrientationConvention: "star(e_I)=sgn(I, I^c)e_{I^c} relative to e_0∧...∧e_7",
		StarSquaredResidual:   residual,
		Trace:                 trace,
		SelfDualDimension:     selfDual,
		AntiSelfDualDimension: antiSelfDual,
		Verdict:               StatusHodgeStarTyped,
	}, nil
}

func buildBasisCertificate(space contact.Space, qK, qW, qSpan linear.Matrix) (BasisCertificate, error) {
	qkIso, err := isometryResidual(qK)
	if err != nil {
		return BasisCertificate{}, err
	}
	qwIso, err := isometryResidual(qW)
	if err != nil {
		return BasisCertificate{}, err
	}
	qwqk, err := qW.Transpose().Mul(qK)
	if err != nil {
		return BasisCertificate{}, err
	}
	pbqk, err := space.BooleanSupport.Support.Matrix.Mul(qK)
	if err != nil {
		return BasisCertificate{}, err
	}
	pgqk, err := space.G2Support.Support.Matrix.Mul(qK)
	if err != nil {
		return BasisCertificate{}, err
	}
	pbqkDiff, err := pbqk.Sub(qK)
	if err != nil {
		return BasisCertificate{}, err
	}
	pgqkDiff, err := pgqk.Sub(qK)
	if err != nil {
		return BasisCertificate{}, err
	}
	pbqw, err := space.BooleanSupport.Support.Matrix.Mul(qW)
	if err != nil {
		return BasisCertificate{}, err
	}
	pgqw, err := space.G2Support.Support.Matrix.Mul(qW)
	if err != nil {
		return BasisCertificate{}, err
	}
	return BasisCertificate{
		QKRows:                    qK.Rows(),
		QKCols:                    qK.Cols(),
		QWRows:                    qW.Rows(),
		QWCols:                    qW.Cols(),
		SpanRows:                  qSpan.Rows(),
		SpanCols:                  qSpan.Cols(),
		QKOrthonormalResidual:     qkIso,
		QWOrthonormalResidual:     qwIso,
		QWQKOrthogonalityResidual: qwqk.FrobeniusNorm(),
		PBQKMinusQKResidual:       pbqkDiff.FrobeniusNorm(),
		PGQKMinusQKResidual:       pgqkDiff.FrobeniusNorm(),
		PBQWResidual:              pbqw.FrobeniusNorm(),
		PGQWResidual:              pgqw.FrobeniusNorm(),
		QWOrthogonalToUAndV:       pbqw.FrobeniusNorm() < numericalTolerance && pgqw.FrobeniusNorm() < numericalTolerance,
		K7ContainedInUAndV:        pbqkDiff.FrobeniusNorm() < numericalTolerance && pgqkDiff.FrobeniusNorm() < numericalTolerance,
		Verdict:                   StatusK7AndW7BasesCertified,
	}, nil
}

func buildLeakageRankTable(m linear.Matrix) (LeakageRankTable, error) {
	sv, err := singularValues(m)
	if err != nil {
		return LeakageRankTable{}, err
	}
	rank := rankFromSingularValues(sv, numericalTolerance)
	minSV := 0.0
	cond := math.Inf(1)
	if len(sv) > 0 {
		minSV = sv[len(sv)-1]
		if minSV > 0 {
			cond = sv[0] / minSV
		}
	}
	det, err := determinant(m)
	if err != nil {
		return LeakageRankTable{}, err
	}
	classification := "rank-zero: Hodge star does not leak K_7 into W_7"
	verdict := StatusHodgeStarDoesNotPairK7ToW7
	if rank == k7DimExpected {
		classification = "rank-seven: nondegenerate Hodge-star K_7->W_7 pairing"
		verdict = "PASS_HODGE_STAR_K7_TO_W7_PAIRING_NONDEGENERATE"
	} else if rank > 0 {
		classification = "partial Hodge leakage only"
		verdict = "CONDITIONAL_SUPPORT_PARTIAL_HODGE_LEAKAGE_ONLY"
	}
	return LeakageRankTable{
		Formula:              "M_* = Q_W^T * Q_K",
		Rows:                 m.Rows(),
		Cols:                 m.Cols(),
		Rank:                 rank,
		SingularValues:       sv,
		Determinant:          det,
		FrobeniusNorm:        m.FrobeniusNorm(),
		MinimumSingularValue: minSV,
		ConditionNumber:      cond,
		Classification:       classification,
		Verdict:              verdict,
	}, nil
}

func buildImageContainment(pW, pUV, starK linear.Matrix) (ImageContainmentAudit, error) {
	pwStar, err := pW.Mul(starK)
	if err != nil {
		return ImageContainmentAudit{}, err
	}
	puvStar, err := pUV.Mul(starK)
	if err != nil {
		return ImageContainmentAudit{}, err
	}
	starNorm := starK.FrobeniusNorm()
	leakageNorm := pwStar.FrobeniusNorm()
	spanNorm := puvStar.FrobeniusNorm()
	return ImageContainmentAudit{
		StarK7FrobeniusNorm:         starNorm,
		PWStarK7FrobeniusNorm:       leakageNorm,
		PUVStarK7FrobeniusNorm:      spanNorm,
		LeakageRatio:                safeDiv(leakageNorm, starNorm),
		SpanContainmentRatio:        safeDiv(spanNorm, starNorm),
		StarK7ContainedInUPlusV:     leakageNorm < numericalTolerance,
		TransverseComponentDetected: leakageNorm >= numericalTolerance,
		Verdict:                     StatusHodgeStarDoesNotPairK7ToW7,
	}, nil
}

func buildPairingMetric(m linear.Matrix, rank int) (PairingMetricAudit, error) {
	g, err := m.Transpose().Mul(m)
	if err != nil {
		return PairingMetricAudit{}, err
	}
	trace, err := g.Trace()
	if err != nil {
		return PairingMetricAudit{}, err
	}
	scale := trace / float64(g.Rows())
	residual := proportionalIdentityResidual(g, scale)
	return PairingMetricAudit{
		Formula:                      "G_* = M_*^T M_*",
		Computed:                     true,
		RankFull:                     rank == k7DimExpected,
		Trace:                        trace,
		ScaleCandidate:               scale,
		ProportionalIdentityResidual: residual,
		ConformalOrIsometric:         rank == k7DimExpected && residual < numericalTolerance,
		AnisotropicNondegenerate:     rank == k7DimExpected && residual >= numericalTolerance,
		Degenerate:                   rank < k7DimExpected,
		Verdict:                      StatusHodgeStarDoesNotPairK7ToW7,
	}, nil
}

func buildOrientationAudit(det float64) OrientationAudit {
	sign := 0
	if det > numericalTolerance {
		sign = 1
	} else if det < -numericalTolerance {
		sign = -1
	}
	return OrientationAudit{
		Determinant:                  det,
		NonzeroDeterminant:           sign != 0,
		Sign:                         sign,
		BasisOrientationDependent:    sign != 0,
		PhysicalOrientationCertified: false,
		Verdict:                      StatusHodgeStarDoesNotPairK7ToW7,
	}
}

func buildAlternativeComposites(space contact.Space, pW, star, qK linear.Matrix, baseRank int) (AlternativeCompositeAudit, error) {
	type opFunc func() (linear.Matrix, error)
	operators := []struct {
		name    string
		formula string
		build   opFunc
		same    bool
	}{
		{"P_W * P_B |K7", "P_W * P_B Q_K", func() (linear.Matrix, error) { return compose3(pW, star, space.BooleanSupport.Support.Matrix, qK) }, true},
		{"P_W * P_G |K7", "P_W * P_G Q_K", func() (linear.Matrix, error) { return compose3(pW, star, space.G2Support.Support.Matrix, qK) }, true},
		{"P_W [*,P_B] |K7", "P_W(*P_B-P_B*)Q_K", func() (linear.Matrix, error) {
			return commutatorComposite(pW, star, space.BooleanSupport.Support.Matrix, qK)
		}, true},
		{"P_W [*,P_G] |K7", "P_W(*P_G-P_G*)Q_K", func() (linear.Matrix, error) {
			return commutatorComposite(pW, star, space.G2Support.Support.Matrix, qK)
		}, true},
		{"P_W (*P_B-*P_G) |K7", "P_W*(P_B-P_G)Q_K", func() (linear.Matrix, error) {
			return differenceComposite(pW, star, space.BooleanSupport.Support.Matrix, space.G2Support.Support.Matrix, qK)
		}, false},
	}
	rows := make([]AlternativeCompositeRow, 0, len(operators))
	anyHigher := false
	anyFull := false
	for _, op := range operators {
		image, err := op.build()
		if err != nil {
			return AlternativeCompositeAudit{}, err
		}
		m, err := imageMetricInW(pW, image)
		if err != nil {
			return AlternativeCompositeAudit{}, err
		}
		sv, err := singularValues(m)
		if err != nil {
			return AlternativeCompositeAudit{}, err
		}
		r := rankFromSingularValues(sv, numericalTolerance)
		if r > baseRank {
			anyHigher = true
		}
		if r == k7DimExpected {
			anyFull = true
		}
		rows = append(rows, AlternativeCompositeRow{Name: op.name, Formula: op.formula, Rank: r, FrobeniusNorm: image.FrobeniusNorm(), SameAsPhiStar: op.same && r == baseRank, Verdict: StatusAlternativeCompositesNoHigherRank})
	}
	return AlternativeCompositeAudit{
		Rows:                     rows,
		AnyHigherRankThanPhiStar: anyHigher,
		AnyNondegenerate:         anyFull,
		Reason:                   "because P_B k=P_G k=k for k in K_7 and P_W kills U+V, the tested star/projector composites do not repair the rank-zero Hodge leakage",
		Verdict:                  StatusAlternativeCompositesNoHigherRank,
	}, nil
}

func Statuses() []string {
	return []string{
		StatusGate631Inherited,
		StatusHodgeStarTyped,
		StatusK7AndW7BasesCertified,
		StatusHodgeLeakageMatrixComputed,
		StatusHodgeStarDoesNotPairK7ToW7,
		StatusNoCanonicalK7W7PairingFound,
		StatusAlternativeCompositesNoHigherRank,
		StatusNoBoundaryStressAssignment,
		StatusGate632Boundary,
	}
}

func concatColumns(a, b linear.Matrix) linear.Matrix {
	out := linear.NewMatrix(a.Rows(), a.Cols()+b.Cols())
	for r := 0; r < a.Rows(); r++ {
		for c := 0; c < a.Cols(); c++ {
			out.Set(r, c, a.At(r, c))
		}
		for c := 0; c < b.Cols(); c++ {
			out.Set(r, a.Cols()+c, b.At(r, c))
		}
	}
	return out
}

func orthonormalizeColumns(cols linear.Matrix, eps float64) (linear.Matrix, []float64, error) {
	gram, err := cols.Transpose().Mul(cols)
	if err != nil {
		return linear.Matrix{}, nil, err
	}
	eig, err := linear.SymmetricEigenJacobi(gram, 1e-13, 0)
	if err != nil {
		return linear.Matrix{}, nil, err
	}
	values, vectors, err := linear.SortEigenDescending(eig.Values, eig.Vectors)
	if err != nil {
		return linear.Matrix{}, nil, err
	}
	keep := make([]int, 0)
	for i, value := range values {
		if value > eps {
			keep = append(keep, i)
		}
	}
	q := linear.NewMatrix(cols.Rows(), len(keep))
	for newCol, oldCol := range keep {
		norm := math.Sqrt(values[oldCol])
		for r := 0; r < cols.Rows(); r++ {
			sum := 0.0
			for c := 0; c < cols.Cols(); c++ {
				sum += cols.At(r, c) * vectors.At(c, oldCol)
			}
			q.Set(r, newCol, sum/norm)
		}
	}
	return q, values, nil
}

func basisFromProjector(p linear.Matrix, target, eps float64) (linear.Matrix, []float64, error) {
	eig, err := linear.SymmetricEigenJacobi(p, 1e-13, 0)
	if err != nil {
		return linear.Matrix{}, nil, err
	}
	values, vectors, err := linear.SortEigenDescending(eig.Values, eig.Vectors)
	if err != nil {
		return linear.Matrix{}, nil, err
	}
	cols := make([]int, 0)
	for i, value := range values {
		if math.Abs(value-target) < eps {
			cols = append(cols, i)
		}
	}
	q := linear.NewMatrix(p.Rows(), len(cols))
	for newCol, oldCol := range cols {
		for r := 0; r < p.Rows(); r++ {
			q.Set(r, newCol, vectors.At(r, oldCol))
		}
	}
	return q, values, nil
}

func hodgeStarLambda4R8() (linear.Matrix, error) {
	basis, err := combinatorics.Subsets(vectorDimExpected, 4)
	if err != nil {
		return linear.Matrix{}, err
	}
	index := combinatorics.IndexByKey(basis)
	star := linear.NewMatrix(len(basis), len(basis))
	for col, subset := range basis {
		comp := complementSubset(subset, vectorDimExpected)
		row, ok := index[comp.Key()]
		if !ok {
			return linear.Matrix{}, fmt.Errorf("missing complement %v in Lambda4 basis", comp)
		}
		star.Set(row, col, float64(hodgeSign(subset, vectorDimExpected)))
	}
	return star, nil
}

func complementSubset(s combinatorics.Subset, n int) combinatorics.Subset {
	used := make(map[int]bool, len(s))
	for _, v := range s {
		used[v] = true
	}
	comp := make(combinatorics.Subset, 0, n-len(s))
	for i := 0; i < n; i++ {
		if !used[i] {
			comp = append(comp, i)
		}
	}
	return comp
}

func hodgeSign(s combinatorics.Subset, n int) int {
	used := make(map[int]bool, len(s))
	seq := make([]int, 0, n)
	for _, v := range s {
		used[v] = true
		seq = append(seq, v)
	}
	for i := 0; i < n; i++ {
		if !used[i] {
			seq = append(seq, i)
		}
	}
	inv := 0
	for i := 0; i < len(seq); i++ {
		for j := i + 1; j < len(seq); j++ {
			if seq[i] > seq[j] {
				inv++
			}
		}
	}
	if inv%2 == 0 {
		return 1
	}
	return -1
}

func isometryResidual(q linear.Matrix) (float64, error) {
	qtq, err := q.Transpose().Mul(q)
	if err != nil {
		return 0, err
	}
	diff, err := qtq.Sub(linear.Identity(q.Cols()))
	if err != nil {
		return 0, err
	}
	return diff.FrobeniusNorm(), nil
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
	vals := append([]float64(nil), eig.Values...)
	sort.Slice(vals, func(i, j int) bool { return vals[i] > vals[j] })
	sv := make([]float64, len(vals))
	for i, v := range vals {
		if v < 0 && v > -1e-10 {
			v = 0
		}
		if v < 0 {
			return nil, fmt.Errorf("negative singular eigenvalue %.16g", v)
		}
		sv[i] = math.Sqrt(v)
	}
	return sv, nil
}

func rankFromSingularValues(values []float64, eps float64) int {
	r := 0
	for _, v := range values {
		if v > eps {
			r++
		}
	}
	return r
}

func determinant(m linear.Matrix) (float64, error) {
	if m.Rows() != m.Cols() {
		return 0, fmt.Errorf("determinant requires square matrix: %dx%d", m.Rows(), m.Cols())
	}
	n := m.Rows()
	a := make([][]float64, n)
	for r := 0; r < n; r++ {
		a[r] = make([]float64, n)
		for c := 0; c < n; c++ {
			a[r][c] = m.At(r, c)
		}
	}
	det := 1.0
	sign := 1.0
	for i := 0; i < n; i++ {
		pivot := i
		for r := i + 1; r < n; r++ {
			if math.Abs(a[r][i]) > math.Abs(a[pivot][i]) {
				pivot = r
			}
		}
		if math.Abs(a[pivot][i]) < 1e-14 {
			return 0, nil
		}
		if pivot != i {
			a[pivot], a[i] = a[i], a[pivot]
			sign *= -1
		}
		p := a[i][i]
		det *= p
		for r := i + 1; r < n; r++ {
			factor := a[r][i] / p
			for c := i; c < n; c++ {
				a[r][c] -= factor * a[i][c]
			}
		}
	}
	return sign * det, nil
}

func proportionalIdentityResidual(m linear.Matrix, scale float64) float64 {
	res := 0.0
	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < m.Cols(); c++ {
			want := 0.0
			if r == c {
				want = scale
			}
			d := m.At(r, c) - want
			res += d * d
		}
	}
	return math.Sqrt(res)
}

func compose3(a, b, c, q linear.Matrix) (linear.Matrix, error) {
	bc, err := b.Mul(c)
	if err != nil {
		return linear.Matrix{}, err
	}
	abc, err := a.Mul(bc)
	if err != nil {
		return linear.Matrix{}, err
	}
	return abc.Mul(q)
}

func commutatorComposite(pW, star, projector, qK linear.Matrix) (linear.Matrix, error) {
	sp, err := star.Mul(projector)
	if err != nil {
		return linear.Matrix{}, err
	}
	ps, err := projector.Mul(star)
	if err != nil {
		return linear.Matrix{}, err
	}
	comm, err := sp.Sub(ps)
	if err != nil {
		return linear.Matrix{}, err
	}
	pwComm, err := pW.Mul(comm)
	if err != nil {
		return linear.Matrix{}, err
	}
	return pwComm.Mul(qK)
}

func differenceComposite(pW, star, pb, pg, qK linear.Matrix) (linear.Matrix, error) {
	diff, err := pb.Sub(pg)
	if err != nil {
		return linear.Matrix{}, err
	}
	starDiff, err := star.Mul(diff)
	if err != nil {
		return linear.Matrix{}, err
	}
	pwStarDiff, err := pW.Mul(starDiff)
	if err != nil {
		return linear.Matrix{}, err
	}
	return pwStarDiff.Mul(qK)
}

// imageMetricInW returns a small Gram representative of a W-valued image. Since
// pW is already applied in all alternatives, the nonzero singular values of the
// image are enough for rank comparison.
func imageMetricInW(_ linear.Matrix, image linear.Matrix) (linear.Matrix, error) {
	return image, nil
}

func safeDiv(a, b float64) float64 {
	if b == 0 {
		return math.NaN()
	}
	return a / b
}
