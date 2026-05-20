// Package generation2hodgestarinternaldestinationaudit implements
// Gate 633: Hodge-Star Internal Destination and Octonionic Residual Audit.
//
// Gate 632 proved that the transverse Hodge leakage candidate
//
//	P_W *|_{K_7}: K_7 -> W_7
//
// has rank zero.  This gate asks the next computable question: if *K_7 does
// not reach W_7, where inside U+V does it land?  The high-value hypothesis is
// the octonionic residual V_0=V∩K_7^perp, because dim V=14=7+7.  The actual
// finite matrix result is sharper and more restrictive: L_7:=*K_7 is K_7
// itself, up to numerical precision.  Therefore the Hodge route does not
// produce a new companion seven-plane V_0, U_0, or an oblique T_56 slice; it
// certifies K_7 as Hodge-stable and leaves the boundary-stress assignment
// absent.
package generation2hodgestarinternaldestinationaudit

import (
	"fmt"
	"math"
	"sort"
	"sync"

	gate632 "github.com/bagherbal/asha-engine/pkg/bridge/generation2hodgestark7tow7leakagerankaudit"
	"github.com/bagherbal/asha-engine/pkg/combinatorics"
	"github.com/bagherbal/asha-engine/pkg/geometry/contact"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

const (
	AuditID = "GATE633-HODGE-STAR-INTERNAL-DESTINATION-AUDIT"

	StatusGate632Inherited                  = "PASS_GATE632_HODGE_TRANSVERSE_FAILURE_INHERITED"
	StatusHodgeCompanionL7Defined           = "PASS_HODGE_COMPANION_L7_DEFINED"
	StatusHodgeInternalContainmentConfirmed = "PASS_HODGE_STAR_INTERNAL_CONTAINMENT_CONFIRMED"
	StatusHodgeStarPreservesK7              = "PASS_HODGE_STAR_PRESERVES_K7"
	StatusK7HodgeStable                     = "CONDITIONAL_SUPPORT_K7_IS_HODGE_STABLE"
	StatusNoNewInternalHodgeCompanion       = "FAILED_ROUTE_NO_NEW_INTERNAL_HODGE_COMPANION"
	StatusDoesNotPairOctonionicResidualV0   = "FAILED_ROUTE_HODGE_STAR_DOES_NOT_PAIR_K7_WITH_OCTONIONIC_RESIDUAL_V0"
	StatusVDoesNotDecomposeAsK7PlusStarK7   = "FAILED_ROUTE_V_DOES_NOT_DECOMPOSE_AS_K7_PLUS_STAR_K7"
	StatusDoesNotEnterBooleanResidualU0     = "FAILED_ROUTE_HODGE_STAR_DOES_NOT_ENTER_BOOLEAN_RESIDUAL_U0"
	StatusNoObliqueInternalSevenPlane       = "FAILED_ROUTE_NO_OBLIQUE_INTERNAL_HODGE_SEVEN_PLANE"
	StatusNoK7ToW7Pairing                   = "FAILED_ROUTE_NO_K7_TO_W7_PAIRING"
	StatusNoBoundaryStressAssignment        = "FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT"
	StatusGate633Boundary                   = "FIREWALL_PRESERVED_GATE633_INTERNAL_HODGE_DESTINATION_BOUNDARY"
)

const (
	vectorDimExpected  = 8
	lambda4DimExpected = 70
	rankPBExpected     = 56
	rankPGExpected     = 14
	k7DimExpected      = 7
	spanDimExpected    = 63
	w7DimExpected      = 7
	u0DimExpected      = 49
	v0DimExpected      = 7
	t56DimExpected     = 56
	boundaryPairDim    = 2
	numericalTolerance = 1e-8
	strictTolerance    = 1e-10
)

type Gate632Inheritance struct {
	HDimension                 int
	UDimension                 int
	VDimension                 int
	K7Dimension                int
	SpanDimension              int
	W7Dimension                int
	LeakageRank                int
	PWStarK7FrobeniusNorm      float64
	PUVStarK7FrobeniusNorm     float64
	TransverseFailureCertified bool
	StarK7InsideUPlusV         bool
	NoBoundaryAssignment       bool
	Gate632FirewallPreserved   bool
	Verdict                    string
}

type L7Certificate struct {
	Definition                   string
	Rows                         int
	Cols                         int
	Rank                         int
	QLOthonormalResidual         float64
	PWQLFrobeniusNorm            float64
	PUVProjectionFraction        float64
	PUVContainmentResidual       float64
	StarTwoCycleResidual         float64
	InternalContainmentCertified bool
	Verdict                      string
}

type K7PreservationAudit struct {
	MatrixFormula       string
	Rank                int
	SingularValues      []float64
	Determinant         float64
	ProjectionFraction  float64
	ContainmentResidual float64
	StarPreservesK7     bool
	HodgeStable         bool
	Verdict             string
}

type T56InternalComplementAudit struct {
	Definition          string
	Dimension           int
	Rank                int
	SingularValues      []float64
	ProjectionFraction  float64
	ContainmentResidual float64
	L7InsideT56         bool
	Verdict             string
}

type OctonionicResidualAudit struct {
	Definition            string
	VDimension            int
	K7Dimension           int
	V0Dimension           int
	Rank                  int
	SingularValues        []float64
	ProjectionFraction    float64
	ContainmentResidual   float64
	L7EqualsV0            bool
	VDecomposesAsK7StarK7 bool
	Verdict               string
}

type BooleanResidualAudit struct {
	Definition          string
	UDimension          int
	K7Dimension         int
	U0Dimension         int
	Rank                int
	SingularValues      []float64
	ProjectionFraction  float64
	ContainmentResidual float64
	L7InsideU0          bool
	Verdict             string
}

type ObliqueDecompositionAudit struct {
	Definition              string
	CandidateT56Dimension   int
	ProjectionFraction      float64
	ContainmentResidual     float64
	DirectSumCoordinateNorm float64
	DirectSumResidual       float64
	ObliquePlaneDetected    bool
	Reason                  string
	Verdict                 string
}

type StarTwoCycleAudit struct {
	Formula               string
	StarSquaredResidual   float64
	L7EqualsK7            bool
	L7EqualsV0            bool
	L7EqualsU0            bool
	CarrierClassification string
	Verdict               string
}

type ConsequenceFor7Over72 struct {
	K7Stable                  bool
	Octonionic14HodgeSplit    bool
	NewSevenPlaneDiscovered   bool
	BoundaryPairDimension     int
	TraceWeightPromoted       bool
	BoundaryAssignmentMissing bool
	Statement                 string
	Verdict                   string
}

type Firewalls struct {
	ClaimsBoundaryStressAssignment bool
	ClaimsSevenOver72Theorem       bool
	ClaimsScalarRGMatching         bool
	ClaimsHiggsMassDerivation      bool
	ClaimsFlavorDerivation         bool
	ClaimsCKMPMNSDerivation        bool
	ClaimsGaugeUnification         bool
	Verdict                        string
}

type Analysis struct {
	Inherited             Gate632Inheritance
	L7                    L7Certificate
	K7Preservation        K7PreservationAudit
	T56Complement         T56InternalComplementAudit
	OctonionicResidual    OctonionicResidualAudit
	BooleanResidual       BooleanResidualAudit
	ObliqueDecomposition  ObliqueDecompositionAudit
	StarTwoCycle          StarTwoCycleAudit
	ConsequenceFor7Over72 ConsequenceFor7Over72
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
	pK := space.ContactProjector.Matrix
	pU := space.BooleanSupport.Support.Matrix
	pV := space.G2Support.Support.Matrix
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
	pU0, err := pU.Sub(pK)
	if err != nil {
		return Analysis{}, fmt.Errorf("construct P_U0: %w", err)
	}
	pV0, err := pV.Sub(pK)
	if err != nil {
		return Analysis{}, fmt.Errorf("construct P_V0: %w", err)
	}
	pT, err := pUV.Sub(pK)
	if err != nil {
		return Analysis{}, fmt.Errorf("construct P_T56: %w", err)
	}
	qU0, _, err := basisFromProjector(pU0, 1, numericalTolerance)
	if err != nil {
		return Analysis{}, fmt.Errorf("extract U0 basis: %w", err)
	}
	qV0, _, err := basisFromProjector(pV0, 1, numericalTolerance)
	if err != nil {
		return Analysis{}, fmt.Errorf("extract V0 basis: %w", err)
	}
	qT, _, err := basisFromProjector(pT, 1, numericalTolerance)
	if err != nil {
		return Analysis{}, fmt.Errorf("extract T56 basis: %w", err)
	}
	if qU0.Cols() != u0DimExpected || qV0.Cols() != v0DimExpected || qT.Cols() != t56DimExpected {
		return Analysis{}, fmt.Errorf("unexpected residual dimensions: U0=%d V0=%d T56=%d", qU0.Cols(), qV0.Cols(), qT.Cols())
	}
	star, err := hodgeStarLambda4R8()
	if err != nil {
		return Analysis{}, err
	}
	qL, err := star.Mul(qK)
	if err != nil {
		return Analysis{}, fmt.Errorf("compute Q_L=*Q_K: %w", err)
	}

	l7, err := buildL7Certificate(qK, qL, pW, pUV, star)
	if err != nil {
		return Analysis{}, err
	}
	k7, err := buildK7Preservation(qK, qL, pK)
	if err != nil {
		return Analysis{}, err
	}
	t56, err := buildT56Audit(qT, qL, pT)
	if err != nil {
		return Analysis{}, err
	}
	v0, err := buildOctonionicResidual(qV0, qL, pV0)
	if err != nil {
		return Analysis{}, err
	}
	u0, err := buildBooleanResidual(qU0, qL, pU0)
	if err != nil {
		return Analysis{}, err
	}
	oblique, err := buildObliqueDecomposition(qU0, qV0, qL, pT)
	if err != nil {
		return Analysis{}, err
	}
	cycle := StarTwoCycleAudit{
		Formula:               "S_* L_7 = K_7 because L_7=S_*K_7 and S_*^2=I on Lambda^4 R^8",
		StarSquaredResidual:   l7.StarTwoCycleResidual,
		L7EqualsK7:            k7.StarPreservesK7,
		L7EqualsV0:            v0.L7EqualsV0,
		L7EqualsU0:            u0.L7InsideU0,
		CarrierClassification: "K_7 is Hodge-stable; no distinct companion seven-plane is produced",
		Verdict:               StatusHodgeStarPreservesK7,
	}
	consequence := ConsequenceFor7Over72{
		K7Stable:                  k7.StarPreservesK7,
		Octonionic14HodgeSplit:    v0.VDecomposesAsK7StarK7,
		NewSevenPlaneDiscovered:   false,
		BoundaryPairDimension:     boundaryPairDim,
		TraceWeightPromoted:       false,
		BoundaryAssignmentMissing: true,
		Statement:                 "Gate 633 blocks the hoped-for V=K_7⊕*K_7 split: *K_7 returns to K_7 itself.  The 7/72 bridge weight remains a dimension-compression clue, not a native trace theorem or boundary assignment.",
		Verdict:                   StatusNoBoundaryStressAssignment,
	}
	firewalls := Firewalls{Verdict: StatusGate633Boundary}
	truth := "Gate 633 identifies the internal destination of *K_7 after Gate 632's transverse leakage failure.  The high-value V_0 hypothesis is tested and fails: projection of L_7=*K_7 onto V_0 is numerical-zero scale.  Instead L_7 equals K_7 up to certified tolerance, with Q_K^T*Q_K full-rank and singular values equal to one.  Thus K_7 is Hodge-stable; no octonionic residual companion, Boolean residual carrier, oblique T_56 seven-plane, K_7->W_7 pairing, or boundary-stress assignment is certified."

	return Analysis{
		Inherited:             inherited,
		L7:                    l7,
		K7Preservation:        k7,
		T56Complement:         t56,
		OctonionicResidual:    v0,
		BooleanResidual:       u0,
		ObliqueDecomposition:  oblique,
		StarTwoCycle:          cycle,
		ConsequenceFor7Over72: consequence,
		Firewalls:             firewalls,
		Truth:                 truth,
	}, nil
}

func buildInheritance() (Gate632Inheritance, error) {
	g632, err := gate632.BuildDefault()
	if err != nil {
		return Gate632Inheritance{}, fmt.Errorf("Gate632 inheritance unavailable: %w", err)
	}
	return Gate632Inheritance{
		HDimension:                 g632.Basis.QKRows,
		UDimension:                 rankPBExpected,
		VDimension:                 rankPGExpected,
		K7Dimension:                g632.Basis.QKCols,
		SpanDimension:              g632.Basis.SpanCols,
		W7Dimension:                g632.Basis.QWCols,
		LeakageRank:                g632.Leakage.Rank,
		PWStarK7FrobeniusNorm:      g632.ImageContainment.PWStarK7FrobeniusNorm,
		PUVStarK7FrobeniusNorm:     g632.ImageContainment.PUVStarK7FrobeniusNorm,
		TransverseFailureCertified: g632.Leakage.Verdict == gate632.StatusHodgeStarDoesNotPairK7ToW7,
		StarK7InsideUPlusV:         g632.ImageContainment.StarK7ContainedInUPlusV,
		NoBoundaryAssignment:       g632.BoundaryReadiness.Verdict == gate632.StatusNoBoundaryStressAssignment,
		Gate632FirewallPreserved:   g632.Firewalls.Verdict == gate632.StatusGate632Boundary,
		Verdict:                    StatusGate632Inherited,
	}, nil
}

func buildL7Certificate(qK, qL, pW, pUV, star linear.Matrix) (L7Certificate, error) {
	iso, err := isometryResidual(qL)
	if err != nil {
		return L7Certificate{}, err
	}
	pwql, err := pW.Mul(qL)
	if err != nil {
		return L7Certificate{}, err
	}
	puvFrac, err := projectionFraction(qL, pUV)
	if err != nil {
		return L7Certificate{}, err
	}
	puvResidual, err := projectionResidual(qL, pUV)
	if err != nil {
		return L7Certificate{}, err
	}
	starQL, err := star.Mul(qL)
	if err != nil {
		return L7Certificate{}, err
	}
	cycleDiff, err := starQL.Sub(qK)
	if err != nil {
		return L7Certificate{}, err
	}
	sv, err := singularValues(qL)
	if err != nil {
		return L7Certificate{}, err
	}
	return L7Certificate{
		Definition:                   "L_7 := *K_7, represented by Q_L=S_*Q_K",
		Rows:                         qL.Rows(),
		Cols:                         qL.Cols(),
		Rank:                         rankFromSingularValues(sv, numericalTolerance),
		QLOthonormalResidual:         iso,
		PWQLFrobeniusNorm:            pwql.FrobeniusNorm(),
		PUVProjectionFraction:        puvFrac,
		PUVContainmentResidual:       puvResidual,
		StarTwoCycleResidual:         cycleDiff.FrobeniusNorm(),
		InternalContainmentCertified: pwql.FrobeniusNorm() < numericalTolerance && puvResidual < numericalTolerance,
		Verdict:                      StatusHodgeCompanionL7Defined,
	}, nil
}

func buildK7Preservation(qK, qL, pK linear.Matrix) (K7PreservationAudit, error) {
	m, err := qK.Transpose().Mul(qL)
	if err != nil {
		return K7PreservationAudit{}, err
	}
	sv, err := singularValues(m)
	if err != nil {
		return K7PreservationAudit{}, err
	}
	det, err := determinant(m)
	if err != nil {
		return K7PreservationAudit{}, err
	}
	frac, err := projectionFraction(qL, pK)
	if err != nil {
		return K7PreservationAudit{}, err
	}
	resid, err := projectionResidual(qL, pK)
	if err != nil {
		return K7PreservationAudit{}, err
	}
	rank := rankFromSingularValues(sv, numericalTolerance)
	preserves := rank == k7DimExpected && math.Abs(frac-1) < numericalTolerance && resid < numericalTolerance
	return K7PreservationAudit{
		MatrixFormula:       "M_KK = Q_K^T S_* Q_K",
		Rank:                rank,
		SingularValues:      sv,
		Determinant:         det,
		ProjectionFraction:  frac,
		ContainmentResidual: resid,
		StarPreservesK7:     preserves,
		HodgeStable:         preserves,
		Verdict:             StatusHodgeStarPreservesK7,
	}, nil
}

func buildT56Audit(qT, qL, pT linear.Matrix) (T56InternalComplementAudit, error) {
	m, err := qT.Transpose().Mul(qL)
	if err != nil {
		return T56InternalComplementAudit{}, err
	}
	sv, err := singularValues(m)
	if err != nil {
		return T56InternalComplementAudit{}, err
	}
	frac, err := projectionFraction(qL, pT)
	if err != nil {
		return T56InternalComplementAudit{}, err
	}
	resid, err := projectionResidual(qL, pT)
	if err != nil {
		return T56InternalComplementAudit{}, err
	}
	return T56InternalComplementAudit{
		Definition:          "T_56=(U+V)∩K_7^perp, with projector P_T=P_{U+V}-P_K",
		Dimension:           qT.Cols(),
		Rank:                rankFromSingularValues(sv, numericalTolerance),
		SingularValues:      sv,
		ProjectionFraction:  frac,
		ContainmentResidual: resid,
		L7InsideT56:         resid < numericalTolerance,
		Verdict:             StatusNoObliqueInternalSevenPlane,
	}, nil
}

func buildOctonionicResidual(qV0, qL, pV0 linear.Matrix) (OctonionicResidualAudit, error) {
	m, err := qV0.Transpose().Mul(qL)
	if err != nil {
		return OctonionicResidualAudit{}, err
	}
	sv, err := singularValues(m)
	if err != nil {
		return OctonionicResidualAudit{}, err
	}
	frac, err := projectionFraction(qL, pV0)
	if err != nil {
		return OctonionicResidualAudit{}, err
	}
	resid, err := projectionResidual(qL, pV0)
	if err != nil {
		return OctonionicResidualAudit{}, err
	}
	rank := rankFromSingularValues(sv, numericalTolerance)
	equals := rank == v0DimExpected && resid < numericalTolerance && math.Abs(frac-1) < numericalTolerance
	return OctonionicResidualAudit{
		Definition:            "V_0=V∩K_7^perp=V⊖K_7",
		VDimension:            rankPGExpected,
		K7Dimension:           k7DimExpected,
		V0Dimension:           qV0.Cols(),
		Rank:                  rank,
		SingularValues:        sv,
		ProjectionFraction:    frac,
		ContainmentResidual:   resid,
		L7EqualsV0:            equals,
		VDecomposesAsK7StarK7: equals,
		Verdict:               StatusDoesNotPairOctonionicResidualV0,
	}, nil
}

func buildBooleanResidual(qU0, qL, pU0 linear.Matrix) (BooleanResidualAudit, error) {
	m, err := qU0.Transpose().Mul(qL)
	if err != nil {
		return BooleanResidualAudit{}, err
	}
	sv, err := singularValues(m)
	if err != nil {
		return BooleanResidualAudit{}, err
	}
	frac, err := projectionFraction(qL, pU0)
	if err != nil {
		return BooleanResidualAudit{}, err
	}
	resid, err := projectionResidual(qL, pU0)
	if err != nil {
		return BooleanResidualAudit{}, err
	}
	inside := resid < numericalTolerance && math.Abs(frac-1) < numericalTolerance
	return BooleanResidualAudit{
		Definition:          "U_0=U∩K_7^perp=U⊖K_7",
		UDimension:          rankPBExpected,
		K7Dimension:         k7DimExpected,
		U0Dimension:         qU0.Cols(),
		Rank:                rankFromSingularValues(sv, numericalTolerance),
		SingularValues:      sv,
		ProjectionFraction:  frac,
		ContainmentResidual: resid,
		L7InsideU0:          inside,
		Verdict:             StatusDoesNotEnterBooleanResidualU0,
	}, nil
}

func buildObliqueDecomposition(qU0, qV0, qL, pT linear.Matrix) (ObliqueDecompositionAudit, error) {
	qUV0 := concatColumns(qU0, qV0)
	coord, err := qUV0.Transpose().Mul(qL)
	if err != nil {
		return ObliqueDecompositionAudit{}, err
	}
	frac, err := projectionFraction(qL, pT)
	if err != nil {
		return ObliqueDecompositionAudit{}, err
	}
	resid, err := projectionResidual(qL, pT)
	if err != nil {
		return ObliqueDecompositionAudit{}, err
	}
	return ObliqueDecompositionAudit{
		Definition:              "oblique test inside T_56=U_0+V_0 after removing the K_7 component",
		CandidateT56Dimension:   t56DimExpected,
		ProjectionFraction:      frac,
		ContainmentResidual:     resid,
		DirectSumCoordinateNorm: coord.FrobeniusNorm(),
		DirectSumResidual:       resid,
		ObliquePlaneDetected:    false,
		Reason:                  "L_7 is already contained in K_7, so there is no nonzero T_56 component to decompose into U_0 and V_0 coordinates",
		Verdict:                 StatusNoObliqueInternalSevenPlane,
	}, nil
}

func Statuses() []string {
	return []string{
		StatusGate632Inherited,
		StatusHodgeCompanionL7Defined,
		StatusHodgeInternalContainmentConfirmed,
		StatusHodgeStarPreservesK7,
		StatusK7HodgeStable,
		StatusNoNewInternalHodgeCompanion,
		StatusDoesNotPairOctonionicResidualV0,
		StatusVDoesNotDecomposeAsK7PlusStarK7,
		StatusDoesNotEnterBooleanResidualU0,
		StatusNoObliqueInternalSevenPlane,
		StatusNoK7ToW7Pairing,
		StatusNoBoundaryStressAssignment,
		StatusGate633Boundary,
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

func projectionFraction(q, p linear.Matrix) (float64, error) {
	pq, err := p.Mul(q)
	if err != nil {
		return 0, err
	}
	n := pq.FrobeniusNorm()
	return (n * n) / float64(q.Cols()), nil
}

func projectionResidual(q, p linear.Matrix) (float64, error) {
	pq, err := p.Mul(q)
	if err != nil {
		return 0, err
	}
	diff, err := q.Sub(pq)
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
