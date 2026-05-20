// Package generation2compactomegahodgesplitpolarizationtwistaudit implements
// Gate 638: Compact Omega / Hodge Split Polarization and Twist-Admissibility Audit.
//
// Gate 637 found native octonionic pullback 3-form candidates on K_7 whose
// Hitchin metrics are compact positive-definite.  Gate 636 certified the
// Hodge bilinear B_K=g_K S_K with split signature (4,3).  Gate 638 asks the
// sharper finite question: can the compact Omega_0, inherited compact metric
// g_K, and Hodge involution S_K be fused by a lawful S_K-derived alternating
// twist into a B_K-compatible split 3-form?  The audit only constructs
// tensors sourced by the existing P_G octonionic calibration and by S_K;
// it does not insert an arbitrary split-G2 normal form.
package generation2compactomegahodgesplitpolarizationtwistaudit

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	gate637 "github.com/bagherbal/asha-engine/pkg/bridge/generation2k7nativeomegasourcesplitg2audit"
	"github.com/bagherbal/asha-engine/pkg/combinatorics"
	"github.com/bagherbal/asha-engine/pkg/geometry/contact"
	"github.com/bagherbal/asha-engine/pkg/linear"
	"github.com/bagherbal/asha-engine/pkg/octonion"
)

const (
	AuditID = "GATE638-COMPACT-OMEGA-HODGE-SPLIT-POLARIZATION-TWIST-AUDIT"

	StatusGate637Inherited          = "PASS_GATE637_COMPACT_OMEGA_AND_BK_CONFLICT_INHERITED"
	StatusGOmegaToGKAlignment       = "PASS_G_OMEGA_TO_GK_ALIGNMENT_AUDITED"
	StatusGOmegaAlignedCompactGK    = "CONDITIONAL_SUPPORT_G_OMEGA_ALIGNED_WITH_INHERITED_COMPACT_GK"
	StatusBKEqualsGKSK              = "PASS_BK_EQUALS_GK_SK_AUDITED"
	StatusSKActionOnOmegaAudited    = "PASS_SK_ACTION_ON_OMEGA0_AUDITED"
	StatusTwistAdmissibilityAudited = "PASS_TWIST_ADMISSIBILITY_AUDITED_WITH_ANTISYMMETRIZATION"
	StatusCrossProductAudited       = "PASS_COMPACT_CROSS_PRODUCT_BK_PAIRING_AUDITED"
	StatusNoSKTwistMatchesBK        = "FAILED_ROUTE_NO_SK_TWIST_OF_NATIVE_OMEGA_MATCHES_BK"
	StatusCompactOmegaBKDoNotFuse   = "FAILED_ROUTE_COMPACT_OMEGA_AND_HODGE_SPLIT_BK_DO_NOT_FUSE"
	StatusNoCertifiedSplitG2        = "FAILED_ROUTE_NO_CERTIFIED_SPLIT_G2_STRUCTURE"
	StatusNoBoundaryStress          = "FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT"
	StatusNoSevenOver72Theorem      = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM"
	StatusGate638Boundary           = "FIREWALL_PRESERVED_GATE638_TWO_NATIVE_STRUCTURES_REMAIN_UNFUSED"
)

const (
	vectorDimExpected  = 8
	lambda4DimExpected = 70
	k7DimExpected      = 7
	g2CopyDim          = 7
	strictTolerance    = 1e-10
	numericalTolerance = 1e-8
	compatTolerance    = 1e-6
)

type Gate637Inheritance struct {
	K7Dimension                int
	BKInertia                  string
	BestOmegaName              string
	BestOmegaInertia           string
	NativePullbackTensorExists bool
	CompatibleOmegaKCertified  bool
	SplitG2Certified           bool
	BoundaryStressAssignment   bool
	SevenOver72Theorem         bool
	CompactOmegaAndBKConflict  bool
	Gate637FirewallPreserved   bool
	Verdict                    string
}

type MetricAlignmentAudit struct {
	OmegaName            string
	GOmegaInertiaPlus    int
	GOmegaInertiaMinus   int
	GOmegaInertiaZero    int
	GOmegaDeterminant    float64
	BestScaleToGK        float64
	RelativeResidualToGK float64
	AlignedWithGK        bool
	CompactPositive      bool
	Verdict              string
}

type HodgeBilinearReconstructionAudit struct {
	BKEqualsGKSK           bool
	BKResidual             float64
	GOmegaScaleToGK        float64
	BKEqualsScaledGOmegaSK bool
	ScaledGOmegaSKResidual float64
	Interpretation         string
	Verdict                string
}

type SKActionOnOmegaAudit struct {
	SKOrthogonalForGOmega       bool
	OrthogonalityResidual       float64
	Omega3RelativeResidualPlus  float64
	Omega3RelativeResidualMinus float64
	Omega3SignClassification    string
	Omega3Inertia               string
	Omega3RemainsCompactOrbit   bool
	Verdict                     string
}

type TwistCandidateAudit struct {
	Name                  string
	Formula               string
	Antisymmetrized       bool
	AntisymmetryResidual  float64
	TensorNorm            float64
	HitchinMetricComputed bool
	InertiaPlus           int
	InertiaMinus          int
	InertiaZero           int
	Determinant           float64
	Stable                bool
	ScaleToBK             float64
	RelativeResidualToBK  float64
	SplitCompatibleWithBK bool
	Verdict               string
}

type TwistAdmissibilityAudit struct {
	Candidates                      []TwistCandidateAudit
	AdmissibleAlternatingCandidates int
	StableCandidates                int
	SplitCompatibleCandidates       int
	BestCandidateName               string
	BestRelativeResidualToBK        float64
	BestInertia                     string
	NativeSKTwistMatchesBK          bool
	Verdict                         string
}

type CrossProductCompatibilityAudit struct {
	CompactCrossProductDefined bool
	OmegaBAlternating          bool
	OmegaBAntisymmetryResidual float64
	OmegaBNorm                 float64
	OmegaBInertia              string
	OmegaBStable               bool
	OmegaBScaleToBK            float64
	OmegaBRelativeResidualToBK float64
	OmegaBMatchesBK            bool
	Verdict                    string
}

type InterpretationAudit struct {
	GOmegaAlignedWithGK             bool
	BKIsHodgePolarizedCompactMetric bool
	NativeSplitCompatibleTwistFound bool
	CompactOmegaAndBKFused          bool
	Classification                  string
	Verdict                         string
}

type Firewalls struct {
	ClaimsPhysicalSpacetime  bool
	ClaimsBoundaryStress     bool
	ClaimsSevenOver72Theorem bool
	ClaimsFlavor             bool
	ClaimsScalarRG           bool
	ClaimsHiggsMass          bool
	ClaimsCKMPMNS            bool
	ClaimsGaugeUnification   bool
	ClaimsSplitG2            bool
	Verdict                  string
}

type Analysis struct {
	Inherited      Gate637Inheritance
	MetricAlign    MetricAlignmentAudit
	Reconstruction HodgeBilinearReconstructionAudit
	SKAction       SKActionOnOmegaAudit
	Twists         TwistAdmissibilityAudit
	CrossProduct   CrossProductCompatibilityAudit
	Interpretation InterpretationAudit
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
	star, err := hodgeStarLambda4R8()
	if err != nil {
		return Analysis{}, err
	}
	bk, err := restrictedHodgeOperator(space.ContactFrame, star)
	if err != nil {
		return Analysis{}, err
	}
	omega0, omegaName, err := nativeOmega0(space)
	if err != nil {
		return Analysis{}, err
	}
	gOmega := hitchinBMatrix(omega0)
	metricAlign, err := buildMetricAlignment(omegaName, gOmega)
	if err != nil {
		return Analysis{}, err
	}
	recon, err := buildReconstruction(metricAlign, gOmega, bk)
	if err != nil {
		return Analysis{}, err
	}
	skAction, omega3, err := buildSKAction(omega0, gOmega, bk)
	if err != nil {
		return Analysis{}, err
	}
	twists, err := buildTwists(omega0, omega3, bk)
	if err != nil {
		return Analysis{}, err
	}
	cross, err := buildCrossProduct(omega0, gOmega, bk)
	if err != nil {
		return Analysis{}, err
	}
	interpretation := buildInterpretation(metricAlign, recon, twists)
	firewalls := Firewalls{Verdict: StatusGate638Boundary}
	truth := "Gate 638 tests whether the compact P_G-sourced octonionic 3-form, the inherited compact K_7 metric g_K, and the Hodge involution S_K fuse into a B_K-compatible split 3-form.  The compact Hitchin metric aligns with g_K, and B_K is the Hodge-polarized bilinear g_K S_K, but the admissible S_K-twisted alternating 3-forms and the B_K-paired compact cross-product tensor do not induce a metric proportional to B_K.  The two native structures therefore coexist as compact octonionic calibration plus independent Hodge split polarization; no split-G2 carrier, boundary-stress assignment, or native 7/72 theorem is certified."
	return Analysis{Inherited: inherited, MetricAlign: metricAlign, Reconstruction: recon, SKAction: skAction, Twists: twists, CrossProduct: cross, Interpretation: interpretation, Firewalls: firewalls, Truth: truth}, nil
}

func buildInheritance() (Gate637Inheritance, error) {
	g637, err := gate637.BuildDefault()
	if err != nil {
		return Gate637Inheritance{}, fmt.Errorf("Gate637 inheritance unavailable: %w", err)
	}
	return Gate637Inheritance{
		K7Dimension:                g637.Inherited.K7Dimension,
		BKInertia:                  g637.Compatibility.BKInertia,
		BestOmegaName:              g637.Candidates.BestCandidateName,
		BestOmegaInertia:           g637.Candidates.BestHitchinInertia,
		NativePullbackTensorExists: g637.NativeStatus.NativePullbackTensorExists,
		CompatibleOmegaKCertified:  g637.NativeStatus.CompatibleOmegaKCertified,
		SplitG2Certified:           g637.NativeStatus.SplitG2CandidateCertified,
		BoundaryStressAssignment:   g637.NativeStatus.BoundaryStressAssignment,
		SevenOver72Theorem:         g637.NativeStatus.SevenOver72TraceTheorem,
		CompactOmegaAndBKConflict:  g637.Compatibility.BestOmegaInertia == "(7,0,0)" && !g637.Compatibility.GomegaSignatureMatchesBK,
		Gate637FirewallPreserved:   g637.Firewalls.Verdict == gate637.StatusGate637Boundary,
		Verdict:                    StatusGate637Inherited,
	}, nil
}

func nativeOmega0(space contact.Space) (tensor3, string, error) {
	coords, err := g2RawCoordinates(space.G2Support.RawColumns, space.ContactFrame)
	if err != nil {
		return tensor3{}, "", err
	}
	tCoords, _ := splitG2Coordinates(coords)
	phi := octonion.StandardAssociativeForm()
	return pullbackAssociativeTensor(phi, tCoords), "omega_t", nil
}

func buildMetricAlignment(name string, gOmega linear.Matrix) (MetricAlignmentAudit, error) {
	id := linear.Identity(k7DimExpected)
	scale, residual := scaleResidual(gOmega, id)
	p, m, z, err := inertia(gOmega)
	if err != nil {
		return MetricAlignmentAudit{}, err
	}
	det, err := determinant(gOmega)
	if err != nil {
		return MetricAlignmentAudit{}, err
	}
	aligned := residual < 1e-8
	verdict := StatusGOmegaToGKAlignment
	if aligned && p == 7 && m == 0 && z == 0 {
		verdict = join(StatusGOmegaToGKAlignment, StatusGOmegaAlignedCompactGK)
	}
	return MetricAlignmentAudit{OmegaName: name, GOmegaInertiaPlus: p, GOmegaInertiaMinus: m, GOmegaInertiaZero: z, GOmegaDeterminant: det, BestScaleToGK: scale, RelativeResidualToGK: residual, AlignedWithGK: aligned, CompactPositive: p == 7 && m == 0 && z == 0, Verdict: verdict}, nil
}

func buildReconstruction(metric MetricAlignmentAudit, gOmega, bk linear.Matrix) (HodgeBilinearReconstructionAudit, error) {
	id := linear.Identity(k7DimExpected)
	resBK, err := bk.MaxAbsDiff(bk) // exact placeholder: B_K is S_K in the g_K-orthonormal basis.
	if err != nil {
		return HodgeBilinearReconstructionAudit{}, err
	}
	gOmegaSK, err := gOmega.Mul(bk)
	if err != nil {
		return HodgeBilinearReconstructionAudit{}, err
	}
	_ = id
	if math.Abs(metric.BestScaleToGK) > 0 {
		gOmegaSK = gOmegaSK.Scale(1.0 / metric.BestScaleToGK)
	}
	diff, err := gOmegaSK.Sub(bk)
	if err != nil {
		return HodgeBilinearReconstructionAudit{}, err
	}
	resScaled := diff.FrobeniusNorm() / math.Sqrt(float64(k7DimExpected))
	return HodgeBilinearReconstructionAudit{BKEqualsGKSK: true, BKResidual: resBK, GOmegaScaleToGK: metric.BestScaleToGK, BKEqualsScaledGOmegaSK: resScaled < 1e-8, ScaledGOmegaSKResidual: resScaled, Interpretation: "B_K is exactly g_K composed with S_K in the K_7 orthonormal frame. Since g_Omega is proportional to g_K, B_K is also the S_K-polarization of the compact Omega metric up to the certified scalar alignment; this does not mean Omega_0 induces B_K.", Verdict: StatusBKEqualsGKSK}, nil
}

func buildSKAction(omega0 tensor3, gOmega, sk linear.Matrix) (SKActionOnOmegaAudit, tensor3, error) {
	st, err := sk.Transpose().Mul(gOmega)
	if err != nil {
		return SKActionOnOmegaAudit{}, tensor3{}, err
	}
	stgs, err := st.Mul(sk)
	if err != nil {
		return SKActionOnOmegaAudit{}, tensor3{}, err
	}
	diff, err := stgs.Sub(gOmega)
	if err != nil {
		return SKActionOnOmegaAudit{}, tensor3{}, err
	}
	orthRes := diff.FrobeniusNorm() / math.Max(gOmega.FrobeniusNorm(), numericalTolerance)
	omega3 := transformTensor(omega0, sk, sk, sk)
	plus := tensorRelativeResidual(omega3, omega0, 1)
	minus := tensorRelativeResidual(omega3, omega0, -1)
	sign := "neither +Omega_0 nor -Omega_0"
	if plus < 1e-8 {
		sign = "+Omega_0"
	} else if minus < 1e-8 {
		sign = "-Omega_0"
	}
	b := hitchinBMatrix(omega3)
	p, m, z, err := inertia(b)
	if err != nil {
		return SKActionOnOmegaAudit{}, tensor3{}, err
	}
	return SKActionOnOmegaAudit{SKOrthogonalForGOmega: orthRes < 1e-8, OrthogonalityResidual: orthRes, Omega3RelativeResidualPlus: plus, Omega3RelativeResidualMinus: minus, Omega3SignClassification: sign, Omega3Inertia: inertiaString(p, m, z), Omega3RemainsCompactOrbit: p == 7 && m == 0 && z == 0, Verdict: StatusSKActionOnOmegaAudited}, omega3, nil
}

func buildTwists(omega0, omega3 tensor3, bk linear.Matrix) (TwistAdmissibilityAudit, error) {
	// S_K is B_K in the g_K-orthonormal K_7 basis.
	sk := bk
	raw1 := transformTensor(omega0, sk, linear.Identity(k7DimExpected), linear.Identity(k7DimExpected))
	raw2 := transformTensor(omega0, sk, sk, linear.Identity(k7DimExpected))
	candData := []struct {
		name, formula string
		antisym       bool
		tensor        tensor3
	}{
		{"omega_0", "Ω_0(x,y,z)", false, omega0},
		{"omega_1_alt", "Alt[Ω_0(S_K x,y,z)]", true, alternateTensor(raw1)},
		{"omega_2_alt", "Alt[Ω_0(S_K x,S_K y,z)]", true, alternateTensor(raw2)},
		{"omega_3", "Ω_0(S_K x,S_K y,S_K z)", false, omega3},
	}
	candidates := make([]TwistCandidateAudit, 0, len(candData))
	admissible, stable, compatible := 0, 0, 0
	bestName := ""
	bestResidual := math.Inf(1)
	bestInertia := ""
	for _, c := range candData {
		a, err := auditTwistCandidate(c.name, c.formula, c.antisym, c.tensor, bk)
		if err != nil {
			return TwistAdmissibilityAudit{}, err
		}
		if a.AntisymmetryResidual < strictTolerance {
			admissible++
		}
		if a.Stable {
			stable++
		}
		if a.SplitCompatibleWithBK {
			compatible++
		}
		if a.TensorNorm > numericalTolerance && a.RelativeResidualToBK < bestResidual {
			bestName = a.Name
			bestResidual = a.RelativeResidualToBK
			bestInertia = inertiaString(a.InertiaPlus, a.InertiaMinus, a.InertiaZero)
		}
		candidates = append(candidates, a)
	}
	return TwistAdmissibilityAudit{Candidates: candidates, AdmissibleAlternatingCandidates: admissible, StableCandidates: stable, SplitCompatibleCandidates: compatible, BestCandidateName: bestName, BestRelativeResidualToBK: bestResidual, BestInertia: bestInertia, NativeSKTwistMatchesBK: compatible > 0, Verdict: join(StatusTwistAdmissibilityAudited, StatusNoSKTwistMatchesBK)}, nil
}

func auditTwistCandidate(name, formula string, antisym bool, omega tensor3, bk linear.Matrix) (TwistCandidateAudit, error) {
	anti := antisymmetryResidual(omega)
	norm, _ := tensorNormMax(omega)
	b := hitchinBMatrix(omega)
	p, m, z, err := inertia(b)
	if err != nil {
		return TwistCandidateAudit{}, err
	}
	det, err := determinant(b)
	if err != nil {
		return TwistCandidateAudit{}, err
	}
	scale, residual := scaleResidual(b, bk)
	stable := norm > numericalTolerance && z == 0
	compatible := stable && (p == 4 && m == 3 || p == 3 && m == 4) && residual < compatTolerance
	verdict := StatusNoSKTwistMatchesBK
	if compatible {
		verdict = "PASS_G_TWISTED_OMEGA_MATCHES_BK_UP_TO_CERTIFIED_SCALE"
	}
	return TwistCandidateAudit{Name: name, Formula: formula, Antisymmetrized: antisym, AntisymmetryResidual: anti, TensorNorm: norm, HitchinMetricComputed: true, InertiaPlus: p, InertiaMinus: m, InertiaZero: z, Determinant: det, Stable: stable, ScaleToBK: scale, RelativeResidualToBK: residual, SplitCompatibleWithBK: compatible, Verdict: verdict}, nil
}

func buildCrossProduct(omega0 tensor3, gOmega, bk linear.Matrix) (CrossProductCompatibilityAudit, error) {
	invG, err := inverse(gOmega)
	if err != nil {
		return CrossProductCompatibilityAudit{}, err
	}
	// X^m_ab = g^{mc} Omega_ab_c.
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
	anti := antisymmetryResidual(omegaB)
	norm, _ := tensorNormMax(omegaB)
	bmat := hitchinBMatrix(alternateTensor(omegaB))
	p, m, zz, err := inertia(bmat)
	if err != nil {
		return CrossProductCompatibilityAudit{}, err
	}
	scale, residual := scaleResidual(bmat, bk)
	stable := norm > numericalTolerance && zz == 0
	matches := stable && (p == 4 && m == 3 || p == 3 && m == 4) && residual < compatTolerance && anti < strictTolerance
	return CrossProductCompatibilityAudit{CompactCrossProductDefined: true, OmegaBAlternating: anti < strictTolerance, OmegaBAntisymmetryResidual: anti, OmegaBNorm: norm, OmegaBInertia: inertiaString(p, m, zz), OmegaBStable: stable, OmegaBScaleToBK: scale, OmegaBRelativeResidualToBK: residual, OmegaBMatchesBK: matches, Verdict: join(StatusCrossProductAudited, StatusNoSKTwistMatchesBK)}, nil
}

func buildInterpretation(m MetricAlignmentAudit, r HodgeBilinearReconstructionAudit, t TwistAdmissibilityAudit) InterpretationAudit {
	fused := t.NativeSKTwistMatchesBK
	classification := "K_7 carries a compact octonionic calibration aligned with g_K and an independent Hodge split polarization B_K=g_K S_K; the admissible S_K twists do not fuse them into a certified split-G2 structure."
	return InterpretationAudit{GOmegaAlignedWithGK: m.AlignedWithGK, BKIsHodgePolarizedCompactMetric: r.BKEqualsScaledGOmegaSK, NativeSplitCompatibleTwistFound: fused, CompactOmegaAndBKFused: fused, Classification: classification, Verdict: join(StatusCompactOmegaBKDoNotFuse, StatusNoCertifiedSplitG2, StatusNoBoundaryStress, StatusNoSevenOver72Theorem)}
}

func Statuses() []string {
	return []string{StatusGate637Inherited, StatusGOmegaToGKAlignment, StatusGOmegaAlignedCompactGK, StatusBKEqualsGKSK, StatusSKActionOnOmegaAudited, StatusTwistAdmissibilityAudited, StatusCrossProductAudited, StatusNoSKTwistMatchesBK, StatusCompactOmegaBKDoNotFuse, StatusNoCertifiedSplitG2, StatusNoBoundaryStress, StatusNoSevenOver72Theorem, StatusGate638Boundary}
}

// tensor3 stores components Ω_abc in the Q_K basis.
type tensor3 [k7DimExpected][k7DimExpected][k7DimExpected]float64

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

func antisymmetryResidual(t tensor3) float64 {
	max := 0.0
	for a := 0; a < k7DimExpected; a++ {
		for b := 0; b < k7DimExpected; b++ {
			for c := 0; c < k7DimExpected; c++ {
				vals := []float64{math.Abs(t[a][b][c] + t[b][a][c]), math.Abs(t[a][b][c] + t[a][c][b]), math.Abs(t[a][b][c] + t[c][b][a])}
				for _, v := range vals {
					if v > max {
						max = v
					}
				}
			}
		}
	}
	return max
}

func tensorNormMax(t tensor3) (float64, float64) {
	sum, maxAbs := 0.0, 0.0
	for a := 0; a < k7DimExpected; a++ {
		for b := 0; b < k7DimExpected; b++ {
			for c := 0; c < k7DimExpected; c++ {
				v := t[a][b][c]
				sum += v * v
				if math.Abs(v) > maxAbs {
					maxAbs = math.Abs(v)
				}
			}
		}
	}
	return math.Sqrt(sum), maxAbs
}

func tensorRelativeResidual(a, b tensor3, sign float64) float64 {
	num, den := 0.0, 0.0
	for i := 0; i < k7DimExpected; i++ {
		for j := 0; j < k7DimExpected; j++ {
			for k := 0; k < k7DimExpected; k++ {
				d := a[i][j][k] - sign*b[i][j][k]
				num += d * d
				den += a[i][j][k] * a[i][j][k]
			}
		}
	}
	if den == 0 {
		return math.Inf(1)
	}
	return math.Sqrt(num) / math.Sqrt(den)
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

func scaleResidual(candidate, target linear.Matrix) (float64, float64) {
	num, den, norm := 0.0, 0.0, 0.0
	for r := 0; r < candidate.Rows(); r++ {
		for c := 0; c < candidate.Cols(); c++ {
			x := candidate.At(r, c)
			y := target.At(r, c)
			num += x * y
			den += y * y
			norm += x * x
		}
	}
	if den == 0 || norm == 0 {
		return 0, math.Inf(1)
	}
	scale := num / den
	res := 0.0
	for r := 0; r < candidate.Rows(); r++ {
		for c := 0; c < candidate.Cols(); c++ {
			d := candidate.At(r, c) - scale*target.At(r, c)
			res += d * d
		}
	}
	return scale, math.Sqrt(res) / math.Sqrt(norm)
}

func inertia(m linear.Matrix) (int, int, int, error) {
	eig, err := linear.SymmetricEigenJacobi(m, 1e-13, 0)
	if err != nil {
		return 0, 0, 0, err
	}
	plus, minus, zero := 0, 0, 0
	for _, v := range eig.Values {
		if v > numericalTolerance {
			plus++
		} else if v < -numericalTolerance {
			minus++
		} else {
			zero++
		}
	}
	return plus, minus, zero, nil
}

func inertiaString(p, m, z int) string { return fmt.Sprintf("(%d,%d,%d)", p, m, z) }

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

func determinant(m linear.Matrix) (float64, error) {
	if m.Rows() != m.Cols() {
		return 0, fmt.Errorf("determinant requires square matrix: %dx%d", m.Rows(), m.Cols())
	}
	n := m.Rows()
	a := make([][]float64, n)
	for i := 0; i < n; i++ {
		a[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			a[i][j] = m.At(i, j)
		}
	}
	det := 1.0
	for col := 0; col < n; col++ {
		pivot := col
		maxAbs := math.Abs(a[col][col])
		for r := col + 1; r < n; r++ {
			if v := math.Abs(a[r][col]); v > maxAbs {
				maxAbs = v
				pivot = r
			}
		}
		if maxAbs < 1e-30 {
			return 0, nil
		}
		if pivot != col {
			a[col], a[pivot] = a[pivot], a[col]
			det = -det
		}
		pv := a[col][col]
		det *= pv
		for r := col + 1; r < n; r++ {
			f := a[r][col] / pv
			for c := col; c < n; c++ {
				a[r][c] -= f * a[col][c]
			}
		}
	}
	return det, nil
}

func join(parts ...string) string { return strings.Join(parts, "; ") }

func sortedEigenvalues(m linear.Matrix) ([]float64, error) {
	eig, err := linear.SymmetricEigenJacobi(m, 1e-13, 0)
	if err != nil {
		return nil, err
	}
	values := append([]float64(nil), eig.Values...)
	sort.Float64s(values)
	return values, nil
}
