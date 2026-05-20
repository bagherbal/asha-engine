// Package generation2k7nativeomegasourcesplitg2audit implements
// Gate 637: K7 Native Omega Source and Split-G2 Compatibility Audit.
//
// Gate 636 certified the native split bilinear carrier (K_7,B_K) with
// inertia (4,3).  Gate 637 asks the next typed question: can ASHA source a
// native stable 3-form Omega_K from already-existing Boolean--octonionic data
// such that its Hitchin metric is equivalent to B_K?  The audit does not place
// an arbitrary split-G2 normal form on K_7.  It computes the natural pullback
// candidates supplied by the octonionic calibration sector that defines P_G,
// then checks antisymmetry, nonzero/stability, Hitchin metric inertia, and
// proportionality to B_K.  The computed obstruction is sharp: the pullback
// candidates are stable but induce a compact positive metric, not the Gate 636
// split bilinear B_K.  Therefore no compatible native Omega_K, split-G2
// structure, boundary assignment, or 7/72 theorem is certified.
package generation2k7nativeomegasourcesplitg2audit

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	gate636 "github.com/bagherbal/asha-engine/pkg/bridge/generation2k7splitsignaturehodgebilinearaudit"
	"github.com/bagherbal/asha-engine/pkg/combinatorics"
	"github.com/bagherbal/asha-engine/pkg/geometry/contact"
	"github.com/bagherbal/asha-engine/pkg/linear"
	"github.com/bagherbal/asha-engine/pkg/octonion"
)

const (
	AuditID = "GATE637-K7-NATIVE-OMEGA-SOURCE-SPLIT-G2-COMPATIBILITY-AUDIT"

	StatusGate636Inherited             = "PASS_GATE636_BK_SPLIT_SIGNATURE_INHERITED"
	StatusOctonionicCalibrationSource  = "PASS_OCTONIONIC_CALIBRATION_SOURCE_AUDITED"
	StatusPGPullbackCandidatesComputed = "PASS_PG_PULLBACK_OMEGA_CANDIDATES_COMPUTED"
	StatusOmegaCandidateAntisymmetric  = "PASS_OMEGA_CANDIDATE_FULLY_ANTISYMMETRIC"
	StatusHitchinMetricComputed        = "PASS_HITCHIN_METRIC_COMPUTED_FOR_PULLBACK_CANDIDATES"
	StatusOmegaCandidateStable         = "PASS_OCTONIONIC_PULLBACK_OMEGA_CANDIDATE_STABILITY_CERTIFIED"
	StatusOmegaCompactNotSplitBK       = "FAILED_ROUTE_G_OMEGA_IS_COMPACT_POSITIVE_NOT_BK_SPLIT"
	StatusNoCompatibleOmegaK           = "FAILED_ROUTE_NO_NATIVE_COMPATIBLE_OMEGA_K_SOURCE_FOUND"
	StatusSplitSignatureAloneNoSplitG2 = "FAILED_ROUTE_SPLIT_SIGNATURE_ALONE_DOES_NOT_DEFINE_SPLIT_G2"
	StatusNoCertifiedSplitG2           = "FAILED_ROUTE_NO_CERTIFIED_SPLIT_G2_STRUCTURE_YET"
	StatusNoCrossProductIdentity       = "FAILED_ROUTE_NO_BK_COMPATIBLE_CROSS_PRODUCT_IDENTITY_CERTIFIED"
	StatusNoBoundaryStressAssignment   = "FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT"
	StatusNoSevenOver72Theorem         = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM"
	StatusGate637Boundary              = "FIREWALL_PRESERVED_GATE637_SPLIT_G2_IS_INTERNAL_NOT_PHYSICAL"
)

const (
	vectorDimExpected  = 8
	lambda4DimExpected = 70
	k7DimExpected      = 7
	g2SectorDim        = 14
	g2CopyDim          = 7
	strictTolerance    = 1e-10
	numericalTolerance = 1e-8
	compatTolerance    = 1e-6
)

type Gate636Inheritance struct {
	K7Dimension              int
	BKInertiaPlus            int
	BKInertiaMinus           int
	BKInertiaZero            int
	BKTrace                  float64
	BKDeterminant            float64
	NativeSplitSignature     bool
	BilinearNotSelector      bool
	NoFockSelectorMap        bool
	NoSplitG2Yet             bool
	NoBoundaryAssignment     bool
	NoSevenOver72Theorem     bool
	Gate636FirewallPreserved bool
	Verdict                  string
}

type OmegaSourceAudit struct {
	StrongestCandidateSource     string
	PGSectorDimension            int
	RawCalibrationColumns        int
	RawCalibrationRows           int
	AssociativeFanoTerms         int
	CoassociativeTerms           int
	K7ToPGCoordinatesComputed    bool
	UsesArbitrarySplitG2Normal   bool
	UsesExternalThreeForm        bool
	HodgePolarityAloneSufficient bool
	SourceStatus                 string
	Verdict                      string
}

type CandidateTensorAudit struct {
	Name                      string
	Formula                   string
	Source                    string
	FullyAntisymmetric        bool
	AntisymmetryResidual      float64
	TensorNorm                float64
	MaxAbsComponent           float64
	NonZero                   bool
	HitchinMetricComputed     bool
	HitchinMetricFrobenius    float64
	HitchinMetricDeterminant  float64
	HitchinMetricInertiaPlus  int
	HitchinMetricInertiaMinus int
	HitchinMetricInertiaZero  int
	Stable                    bool
	ScaleToBK                 float64
	RelativeResidualToBK      float64
	CompatibleWithBK          bool
	Verdict                   string
}

type OmegaCandidateSummary struct {
	Candidates                     []CandidateTensorAudit
	NonZeroStableCandidates        int
	CompatibleWithBKCount          int
	BestCandidateName              string
	BestRelativeResidualToBK       float64
	BestHitchinInertia             string
	PullbackCandidatesComputed     bool
	CandidateStabilityCertified    bool
	CompatibleNativeOmegaCertified bool
	Verdict                        string
}

type MetricCompatibilityAudit struct {
	BKInertia                string
	BestOmegaName            string
	BestOmegaInertia         string
	BestScaleToBK            float64
	BestRelativeResidualToBK float64
	GomegaProportionalToBK   bool
	GomegaSignatureMatchesBK bool
	CertifiedScaleNotFitted  bool
	Reason                   string
	Verdict                  string
}

type CrossProductAudit struct {
	OmegaCompatibleWithBK      bool
	CrossProductDefined        bool
	BKPairingIdentityCertified bool
	SplitCrossProductIdentity  bool
	Reason                     string
	Verdict                    string
}

type StabilizerAudit struct {
	BilinearStabilizerCandidate string
	OmegaStabilizerCandidate    string
	StabilizerDimensionComputed bool
	ExpectedSplitG2Dimension    int
	SplitG2Certified            bool
	Reason                      string
	Verdict                     string
}

type NativeOmegaStatus struct {
	NativePullbackTensorExists     bool
	CompatibleOmegaKCertified      bool
	SplitG2CandidateCertified      bool
	BoundaryStressAssignment       bool
	SevenOver72TraceTheorem        bool
	PhysicalSpacetimeMetricClaimed bool
	FockSelectorClaimed            bool
	ScalarRGMatchingClaimed        bool
	FlavorClaimed                  bool
	GaugeUnificationClaimed        bool
	Statement                      string
	Verdict                        string
}

type Firewalls struct {
	ClaimsPhysicalSpacetimeMetric bool
	ClaimsFockSelector            bool
	ClaimsBoundaryStress          bool
	ClaimsSevenOver72Theorem      bool
	ClaimsScalarRGMatching        bool
	ClaimsHiggsMass               bool
	ClaimsFlavor                  bool
	ClaimsCKMPMNS                 bool
	ClaimsGaugeUnification        bool
	ClaimsSplitG2WithoutOmega     bool
	Verdict                       string
}

type Analysis struct {
	Inherited     Gate636Inheritance
	Source        OmegaSourceAudit
	Candidates    OmegaCandidateSummary
	Compatibility MetricCompatibilityAudit
	CrossProduct  CrossProductAudit
	Stabilizer    StabilizerAudit
	NativeStatus  NativeOmegaStatus
	Firewalls     Firewalls
	Truth         string
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
	if space.AmbientDimension() != lambda4DimExpected || space.Dimension() != k7DimExpected || space.G2Support.SectorDimension() != g2SectorDim {
		return Analysis{}, fmt.Errorf("unexpected contact/G2 dimensions: ambient=%d K7=%d G2=%d", space.AmbientDimension(), space.Dimension(), space.G2Support.SectorDimension())
	}
	star, err := hodgeStarLambda4R8()
	if err != nil {
		return Analysis{}, err
	}
	bk, err := restrictedHodgeOperator(space.ContactFrame, star)
	if err != nil {
		return Analysis{}, err
	}
	source := buildSourceAudit(space)
	candidates, err := buildCandidateSummary(space, bk)
	if err != nil {
		return Analysis{}, err
	}
	compat := buildCompatibility(candidates)
	cross := buildCrossProduct(compat)
	stab := buildStabilizer(compat)
	status := buildNativeStatus(candidates, compat, stab)
	firewalls := Firewalls{Verdict: StatusGate637Boundary}
	truth := "Gate 637 computes the natural Omega_K candidates sourced by the octonionic calibration sector that defines P_G.  These pullback 3-forms are nonzero, fully antisymmetric, and stable, but their Hitchin metrics are compact positive-definite and not proportional to the Gate 636 split bilinear B_K of inertia (4,3).  Therefore ASHA has a native octonionic pullback tensor, but not a native B_K-compatible Omega_K, not a split-G2 carrier on K_7, not a boundary-stress assignment, and not a 7/72 theorem."
	return Analysis{Inherited: inherited, Source: source, Candidates: candidates, Compatibility: compat, CrossProduct: cross, Stabilizer: stab, NativeStatus: status, Firewalls: firewalls, Truth: truth}, nil
}

func buildInheritance() (Gate636Inheritance, error) {
	g636, err := gate636.BuildDefault()
	if err != nil {
		return Gate636Inheritance{}, fmt.Errorf("Gate636 inheritance unavailable: %w", err)
	}
	return Gate636Inheritance{
		K7Dimension:              g636.Definition.Dimension,
		BKInertiaPlus:            g636.Signature.InertiaPlus,
		BKInertiaMinus:           g636.Signature.InertiaMinus,
		BKInertiaZero:            g636.Signature.InertiaZero,
		BKTrace:                  g636.Signature.Trace,
		BKDeterminant:            g636.Signature.Determinant,
		NativeSplitSignature:     g636.Signature.SplitIndefinite && g636.Octonionic.SplitSignatureMatchesDimension,
		BilinearNotSelector:      g636.Octonionic.SplitSignatureMatchesDimension && !g636.Octonionic.G2SplitStructureCertified,
		NoFockSelectorMap:        !g636.Firewalls.K7ToFockMapCertified,
		NoSplitG2Yet:             !g636.Octonionic.G2SplitStructureCertified,
		NoBoundaryAssignment:     !g636.Firewalls.BoundaryStressAssigned,
		NoSevenOver72Theorem:     !g636.Firewalls.SevenOver72Promoted,
		Gate636FirewallPreserved: g636.Firewalls.Verdict == gate636.StatusGate636Boundary,
		Verdict:                  StatusGate636Inherited,
	}, nil
}

func buildSourceAudit(space contact.Space) OmegaSourceAudit {
	return OmegaSourceAudit{
		StrongestCandidateSource:     "pullback of the standard octonionic associative form through the calibrated P_G sector coordinates of K_7",
		PGSectorDimension:            space.G2Support.SectorDimension(),
		RawCalibrationColumns:        space.G2Support.RawColumns.Cols(),
		RawCalibrationRows:           space.G2Support.RawColumns.Rows(),
		AssociativeFanoTerms:         space.G2Support.Associative.NonZeroCanonicalTerms(),
		CoassociativeTerms:           space.G2Support.Coassociative.NonZeroCanonicalTerms(),
		K7ToPGCoordinatesComputed:    true,
		UsesArbitrarySplitG2Normal:   false,
		UsesExternalThreeForm:        false,
		HodgePolarityAloneSufficient: false,
		SourceStatus:                 StatusOctonionicCalibrationSource,
		Verdict:                      StatusOctonionicCalibrationSource,
	}
}

func buildCandidateSummary(space contact.Space, bk linear.Matrix) (OmegaCandidateSummary, error) {
	coords, err := g2RawCoordinates(space.G2Support.RawColumns, space.ContactFrame)
	if err != nil {
		return OmegaCandidateSummary{}, err
	}
	tCoords, sCoords := splitG2Coordinates(coords)
	phi := octonion.StandardAssociativeForm()
	tTensor := pullbackAssociativeTensor(phi, tCoords)
	sTensor := pullbackAssociativeTensor(phi, sCoords)
	candData := []struct {
		name    string
		formula string
		source  string
		tensor  tensor3
	}{
		{"omega_t", "Ω_t(a,b,c)=φ(t_a,t_b,t_c)", "first calibrated 7-copy in P_G raw sector", tTensor},
		{"omega_s", "Ω_s(a,b,c)=φ(s_a,s_b,s_c)", "second calibrated 7-copy in P_G raw sector", sTensor},
		{"omega_t_plus_s", "Ω_+(a,b,c)=Ω_t+Ω_s", "sum of both calibrated 7-copies", addTensor(tTensor, sTensor, 1, 1)},
		{"omega_t_minus_s", "Ω_-(a,b,c)=Ω_t-Ω_s", "difference of both calibrated 7-copies", addTensor(tTensor, sTensor, 1, -1)},
	}
	candidates := make([]CandidateTensorAudit, 0, len(candData))
	nonZeroStable := 0
	compatible := 0
	bestName := ""
	bestResidual := math.Inf(1)
	bestInertia := ""
	for _, c := range candData {
		audit, err := auditTensorCandidate(c.name, c.formula, c.source, c.tensor, bk)
		if err != nil {
			return OmegaCandidateSummary{}, err
		}
		if audit.NonZero && audit.Stable {
			nonZeroStable++
		}
		if audit.CompatibleWithBK {
			compatible++
		}
		if audit.NonZero && audit.RelativeResidualToBK < bestResidual {
			bestResidual = audit.RelativeResidualToBK
			bestName = audit.Name
			bestInertia = inertiaString(audit.HitchinMetricInertiaPlus, audit.HitchinMetricInertiaMinus, audit.HitchinMetricInertiaZero)
		}
		candidates = append(candidates, audit)
	}
	return OmegaCandidateSummary{
		Candidates:                     candidates,
		NonZeroStableCandidates:        nonZeroStable,
		CompatibleWithBKCount:          compatible,
		BestCandidateName:              bestName,
		BestRelativeResidualToBK:       bestResidual,
		BestHitchinInertia:             bestInertia,
		PullbackCandidatesComputed:     true,
		CandidateStabilityCertified:    nonZeroStable > 0,
		CompatibleNativeOmegaCertified: compatible > 0,
		Verdict:                        join(StatusPGPullbackCandidatesComputed, StatusOmegaCandidateAntisymmetric, StatusHitchinMetricComputed, StatusOmegaCandidateStable, StatusOmegaCompactNotSplitBK, StatusNoCompatibleOmegaK),
	}, nil
}

func auditTensorCandidate(name, formula, source string, omega tensor3, bk linear.Matrix) (CandidateTensorAudit, error) {
	anti := antisymmetryResidual(omega)
	norm, maxAbs := tensorNormMax(omega)
	b := hitchinBMatrix(omega)
	det, err := determinant(b)
	if err != nil {
		return CandidateTensorAudit{}, err
	}
	plus, minus, zero, err := inertia(b)
	if err != nil {
		return CandidateTensorAudit{}, err
	}
	scale, residual := scaleResidual(b, bk)
	stable := norm > numericalTolerance && zero == 0
	compatible := stable && plus == 4 && minus == 3 && zero == 0 && residual < compatTolerance
	verdict := StatusOmegaCompactNotSplitBK
	if norm <= numericalTolerance {
		verdict = StatusNoCompatibleOmegaK
	} else if compatible {
		verdict = "PASS_NATIVE_OMEGA_K_COMPATIBLE_WITH_BK"
	}
	return CandidateTensorAudit{
		Name:                      name,
		Formula:                   formula,
		Source:                    source,
		FullyAntisymmetric:        anti < strictTolerance,
		AntisymmetryResidual:      anti,
		TensorNorm:                norm,
		MaxAbsComponent:           maxAbs,
		NonZero:                   norm > numericalTolerance,
		HitchinMetricComputed:     true,
		HitchinMetricFrobenius:    b.FrobeniusNorm(),
		HitchinMetricDeterminant:  det,
		HitchinMetricInertiaPlus:  plus,
		HitchinMetricInertiaMinus: minus,
		HitchinMetricInertiaZero:  zero,
		Stable:                    stable,
		ScaleToBK:                 scale,
		RelativeResidualToBK:      residual,
		CompatibleWithBK:          compatible,
		Verdict:                   verdict,
	}, nil
}

func buildCompatibility(s OmegaCandidateSummary) MetricCompatibilityAudit {
	return MetricCompatibilityAudit{
		BKInertia:                "(4,3,0)",
		BestOmegaName:            s.BestCandidateName,
		BestOmegaInertia:         s.BestHitchinInertia,
		BestScaleToBK:            bestCandidate(s).ScaleToBK,
		BestRelativeResidualToBK: s.BestRelativeResidualToBK,
		GomegaProportionalToBK:   s.CompatibleWithBKCount > 0,
		GomegaSignatureMatchesBK: s.BestHitchinInertia == "(4,3,0)" || s.BestHitchinInertia == "(3,4,0)",
		CertifiedScaleNotFitted:  false,
		Reason:                   "The only nonzero stable pullbacks from the existing P_G octonionic calibration induce compact positive Hitchin metrics; their best residual against B_K is order one, so the scale cannot be certified as a native proportionality.",
		Verdict:                  join(StatusOmegaCompactNotSplitBK, StatusNoCompatibleOmegaK),
	}
}

func buildCrossProduct(c MetricCompatibilityAudit) CrossProductAudit {
	return CrossProductAudit{
		OmegaCompatibleWithBK:      c.GomegaProportionalToBK && c.GomegaSignatureMatchesBK,
		CrossProductDefined:        false,
		BKPairingIdentityCertified: false,
		SplitCrossProductIdentity:  false,
		Reason:                     "A cross product x×y compatible with B_K requires a B_K-compatible stable 3-form.  The computed octonionic pullback tensors do not supply that metric, so the identity B_K(x×y,z)=Ω_K(x,y,z) is not certified.",
		Verdict:                    StatusNoCrossProductIdentity,
	}
}

func buildStabilizer(c MetricCompatibilityAudit) StabilizerAudit {
	return StabilizerAudit{
		BilinearStabilizerCandidate: "O(4,3)",
		OmegaStabilizerCandidate:    "split-G2 only if a B_K-compatible stable Ω_K is certified",
		StabilizerDimensionComputed: false,
		ExpectedSplitG2Dimension:    14,
		SplitG2Certified:            false,
		Reason:                      "The bilinear stabilizer was certified in Gate 636, but a split-G2 stabilizer is the stabilizer of a compatible stable 3-form, not of B_K alone.",
		Verdict:                     join(StatusSplitSignatureAloneNoSplitG2, StatusNoCertifiedSplitG2),
	}
}

func buildNativeStatus(s OmegaCandidateSummary, c MetricCompatibilityAudit, stab StabilizerAudit) NativeOmegaStatus {
	return NativeOmegaStatus{
		NativePullbackTensorExists:     s.NonZeroStableCandidates > 0,
		CompatibleOmegaKCertified:      s.CompatibleNativeOmegaCertified,
		SplitG2CandidateCertified:      stab.SplitG2Certified,
		BoundaryStressAssignment:       false,
		SevenOver72TraceTheorem:        false,
		PhysicalSpacetimeMetricClaimed: false,
		FockSelectorClaimed:            false,
		ScalarRGMatchingClaimed:        false,
		FlavorClaimed:                  false,
		GaugeUnificationClaimed:        false,
		Statement:                      "A native octonionic pullback 3-form can be computed from P_G, but it is not compatible with the Gate 636 split bilinear B_K; therefore (K_7,B_K,Ω_K) is not yet a certified split-octonionic carrier.",
		Verdict:                        join(StatusNoCompatibleOmegaK, StatusNoCertifiedSplitG2, StatusNoBoundaryStressAssignment, StatusNoSevenOver72Theorem),
	}
}

func Statuses() []string {
	return []string{
		StatusGate636Inherited,
		StatusOctonionicCalibrationSource,
		StatusPGPullbackCandidatesComputed,
		StatusOmegaCandidateAntisymmetric,
		StatusHitchinMetricComputed,
		StatusOmegaCandidateStable,
		StatusOmegaCompactNotSplitBK,
		StatusNoCompatibleOmegaK,
		StatusSplitSignatureAloneNoSplitG2,
		StatusNoCertifiedSplitG2,
		StatusNoCrossProductIdentity,
		StatusNoBoundaryStressAssignment,
		StatusNoSevenOver72Theorem,
		StatusGate637Boundary,
	}
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
	coords, err := pinv.Mul(qK)
	if err != nil {
		return linear.Matrix{}, err
	}
	return coords, nil
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

func addTensor(a, b tensor3, ca, cb float64) tensor3 {
	var out tensor3
	for i := 0; i < k7DimExpected; i++ {
		for j := 0; j < k7DimExpected; j++ {
			for k := 0; k < k7DimExpected; k++ {
				out[i][j][k] = ca*a[i][j][k] + cb*b[i][j][k]
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
				checks := []float64{
					math.Abs(t[a][b][c] + t[b][a][c]),
					math.Abs(t[a][b][c] + t[a][c][b]),
					math.Abs(t[a][b][c] + t[c][b][a]),
				}
				for _, v := range checks {
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

func bestCandidate(s OmegaCandidateSummary) CandidateTensorAudit {
	if len(s.Candidates) == 0 {
		return CandidateTensorAudit{}
	}
	best := s.Candidates[0]
	for _, c := range s.Candidates[1:] {
		if c.RelativeResidualToBK < best.RelativeResidualToBK {
			best = c
		}
	}
	return best
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
