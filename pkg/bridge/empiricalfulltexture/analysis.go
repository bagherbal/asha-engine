// Package empiricalfulltexture implements Gate 265:
// Empirical Full Texture Seal / SVD-CKM Observable Reconstruction Audit.
//
// Gate 264 proved that the restricted three-term finite geometric shell cannot
// fit representative quark masses plus CKM data. Gate 265 therefore keeps the
// EmpiricalYukawaSeal active, ingests full 3x3 quark Yukawa texture matrices as
// sealed boundary data, and verifies the standard flavor-physics reconstruction:
// singular values give mass eigenvalues, and the left-unitary misalignment gives
// the CKM matrix. None of the numerical matrices are promoted to finite-core
// derivations.
package empiricalfulltexture

import (
	"fmt"
	"math"
	"math/cmplx"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/empiricalyukawafit"
)

const (
	AuditID = "GATE265-EMPIRICAL-FULL-TEXTURE-SEAL-SVD-CKM-OBSERVABLE-RECONSTRUCTION-AUDIT"

	StatusGate264Inherited            = "CONDITIONAL_SUPPORT_GATE264_EMPIRICAL_YUKAWA_SEAL_AND_UNDERFIT_INHERITED"
	StatusFullTextureSealActivated    = "CONDITIONAL_SUPPORT_FULL_EMPIRICAL_TEXTURE_SEAL_ACTIVATED"
	StatusFullTextureIngested         = "CONDITIONAL_SUPPORT_FULL_EMPIRICAL_QUARK_TEXTURES_INGESTED"
	StatusSVDCompleted                = "CONDITIONAL_SUPPORT_SVD_DECOMPOSITION_COMPLETED"
	StatusMassEigenvaluesVerified     = "CONDITIONAL_SUPPORT_MASS_EIGENVALUES_RECONSTRUCTED_FROM_SVD"
	StatusCKMReconstructionVerified   = "CONDITIONAL_SUPPORT_SVD_CKM_RECONSTRUCTION_VERIFIED"
	StatusPhenomenologicalOnly        = "CONDITIONAL_SUPPORT_FULL_TEXTURE_OUTPUTS_MARKED_PHENOMENOLOGICAL"
	StatusNoNativeDerivation          = "FAILED_ROUTE_NO_NATIVE_DERIVATION"
	StatusFullTexturesNotFiniteCore   = "FAILED_ROUTE_FULL_YUKAWA_TEXTURES_ARE_EMPIRICAL_BOUNDARY_DATA"
	StatusRestrictedAnsatzStillFailed = "FAILED_ROUTE_RESTRICTED_GEOMETRIC_ANSATZ_REMAINS_EMPIRICALLY_UNDERFIT"
	StatusSVDReconstructionFailed     = "FAILED_ROUTE_SVD_RECONSTRUCTION_FAILED"
	StatusCKMReconstructionFailed     = "FAILED_ROUTE_CKM_RECONSTRUCTION_FAILED"
)

// Matrix3 is a small complex 3x3 matrix used only for Gate 265's sealed flavor
// reconstruction audit. The ordering is generation-labeled rather than sorted:
// (u,c,t) and (d,s,b). This preserves the physical labels required for CKM.
type Matrix3 [3][3]complex128

type Gate264Inheritance struct {
	EmpiricalYukawaSealActive     bool
	RestrictedAnsatzViolated      bool
	FullEmpiricalMatricesRequired bool
	MassesPreviouslyDerived       bool
	CKMPreviouslyDerived          bool
	RepresentativeDataAvailable   bool
	MixedScaleWarning             bool
	UpMassesGeV                   [3]float64
	DownMassesGeV                 [3]float64
	TargetCKM                     Matrix3
	Gate264CombinedResidual       float64
	Verdict                       string
}

type FullTextureSeal struct {
	Name                           string
	Activated                      bool
	ActivatedByGate                int
	BoundaryDataKind               string
	ExplicitlyQuarantined          bool
	DerivedFromFiniteCore          bool
	RewritesGate264Underfit        bool
	AllowsObservableReconstruction bool
	AllowsMassPrediction           bool
	AllowsCKMPrediction            bool
	Verdict                        string
}

type FullTextureData struct {
	SourceLabel                string
	RepresentativeNotPrecision bool
	MixedScaleWarning          bool
	GenerationLabeledSVD       bool
	RightBasisConvention       string
	UpTexture                  Matrix3
	DownTexture                Matrix3
	TargetUpMassesGeV          [3]float64
	TargetDownMassesGeV        [3]float64
	TargetCKM                  Matrix3
	TextureParameterSource     string
	Verdict                    string
}

type SVDAudit struct {
	Sector                      string
	Method                      string
	InputMatrix                 Matrix3
	LeftUnitary                 Matrix3
	SingularValues              [3]float64
	RightUnitary                Matrix3
	ReconstructionResidual      float64
	LeftUnitarityResidual       float64
	RightUnitarityResidual      float64
	ColumnOrthogonalityResidual float64
	MassTarget                  [3]float64
	MassMaxAbsError             float64
	MassRelativeL2Error         float64
	Passed                      bool
	Verdict                     string
}

type MassEigenvalueAudit struct {
	UpMassesGeV               [3]float64
	DownMassesGeV             [3]float64
	UpTargetGeV               [3]float64
	DownTargetGeV             [3]float64
	UpMaxAbsError             float64
	DownMaxAbsError           float64
	UpRelativeL2Error         float64
	DownRelativeL2Error       float64
	Tolerance                 float64
	Verified                  bool
	PhenomenologicalInputOnly bool
	Verdict                   string
}

type CKMReconstructionAudit struct {
	Formula                   string
	ReconstructedCKM          Matrix3
	TargetCKM                 Matrix3
	FrobeniusResidual         float64
	MaxAbsEntryResidual       float64
	AbsMatrixResidual         float64
	UnitarityResidual         float64
	Tolerance                 float64
	Verified                  bool
	DerivedFromFiniteCore     bool
	PhenomenologicalInputOnly bool
	Verdict                   string
}

type FirewallAudit struct {
	EmpiricalSealActive              bool
	FullTexturesQuarantined          bool
	DoesNotRewriteGate264Underfit    bool
	DoesNotClaimFiniteMassDerivation bool
	DoesNotClaimFiniteCKMDerivation  bool
	DoesNotInferYukawaAction         bool
	DoesNotInferVEVOrRGScale         bool
	SVDIsAlgebraicReconstruction     bool
	FiniteCorePolluted               bool
	Verdict                          string
}

type Summary struct {
	FullTexturesIngested       bool
	SVDCompleted               bool
	MassEigenvaluesVerified    bool
	CKMReconstructed           bool
	NativeDerivation           bool
	EmpiricalBoundaryPreserved bool
	Status                     string
	NextGate                   string
	Comment                    string
}

type Analysis struct {
	PreviousGate264 empiricalyukawafit.Analysis
	Inheritance     Gate264Inheritance
	Seal            FullTextureSeal
	Data            FullTextureData
	UpSVD           SVDAudit
	DownSVD         SVDAudit
	Masses          MassEigenvalueAudit
	CKM             CKMReconstructionAudit
	Firewall        FirewallAudit
	Summary         Summary
	TruthStatement  string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := empiricalyukawafit.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 264 predecessor: %w", err)
			return
		}
		inh := inheritGate264(prev)
		seal := activateFullTextureSeal(inh)
		data := ingestFullTextures(inh, seal)
		upSVD := computeSVD("up-sector full empirical texture", data.UpTexture, data.TargetUpMassesGeV)
		downSVD := computeSVD("down-sector full empirical texture", data.DownTexture, data.TargetDownMassesGeV)
		masses := auditMasses(upSVD, downSVD)
		ckm := auditCKM(upSVD, downSVD, data.TargetCKM)
		firewall := auditFirewall(seal, masses, ckm)
		summary := summarize(data, upSVD, downSVD, masses, ckm, firewall)
		truth := buildTruth(inh, seal, data, upSVD, downSVD, masses, ckm)
		defaultA = Analysis{PreviousGate264: prev, Inheritance: inh, Seal: seal, Data: data, UpSVD: upSVD, DownSVD: downSVD, Masses: masses, CKM: ckm, Firewall: firewall, Summary: summary, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate264(prev empiricalyukawafit.Analysis) Gate264Inheritance {
	return Gate264Inheritance{
		EmpiricalYukawaSealActive:     prev.Seal.Activated && prev.Seal.ExplicitlyQuarantined,
		RestrictedAnsatzViolated:      prev.Viability.ViolatesAnsatz,
		FullEmpiricalMatricesRequired: prev.Viability.RequiresFullYukawaMatrices,
		MassesPreviouslyDerived:       prev.Summary.MassesDerived,
		CKMPreviouslyDerived:          prev.Summary.CKMDerived,
		RepresentativeDataAvailable:   prev.Data.UsesObservedMassHierarchy && prev.Data.UsesObservedCKMParameters,
		MixedScaleWarning:             prev.Data.MixedScaleWarning,
		UpMassesGeV:                   prev.Data.UpMassesGeV,
		DownMassesGeV:                 prev.Data.DownMassesGeV,
		TargetCKM:                     fromGate264(prev.Data.CKM),
		Gate264CombinedResidual:       prev.Viability.CombinedRelativeResidual,
		Verdict:                       StatusGate264Inherited + "; Gate 264 underfit is preserved, so full matrices may be used only as sealed phenomenological boundary data",
	}
}

func activateFullTextureSeal(inh Gate264Inheritance) FullTextureSeal {
	active := inh.EmpiricalYukawaSealActive && inh.RestrictedAnsatzViolated && inh.FullEmpiricalMatricesRequired
	return FullTextureSeal{
		Name:                           "EmpiricalYukawaSeal/full-texture branch",
		Activated:                      active,
		ActivatedByGate:                265,
		BoundaryDataKind:               "full representative 3x3 quark Yukawa texture matrices Yu,Yd plus CKM target",
		ExplicitlyQuarantined:          true,
		DerivedFromFiniteCore:          false,
		RewritesGate264Underfit:        false,
		AllowsObservableReconstruction: active,
		AllowsMassPrediction:           false,
		AllowsCKMPrediction:            false,
		Verdict:                        StatusFullTextureSealActivated + "; SVD and CKM may be reconstructed from sealed matrices, but no finite-core mass or CKM prediction is licensed",
	}
}

func ingestFullTextures(inh Gate264Inheritance, seal FullTextureSeal) FullTextureData {
	// Choose a generation-labeled weak basis where the down texture is diagonal,
	// the right-unitary bases are identity, and the entire left-basis misalignment
	// is carried by the up texture. With the convention V_CKM=U_u^† U_d and U_d=I,
	// choose U_u=V_CKM^†, so Yu=V_CKM^† diag(mu,mc,mt).
	ckm := inh.TargetCKM
	yu := mul(dagger(ckm), diagonal(inh.UpMassesGeV))
	yd := diagonal(inh.DownMassesGeV)
	return FullTextureData{
		SourceLabel:                "sealed representative full quark textures built from Gate 264 masses and CKM target",
		RepresentativeNotPrecision: true,
		MixedScaleWarning:          inh.MixedScaleWarning,
		GenerationLabeledSVD:       true,
		RightBasisConvention:       "V_u=V_d=I; U_d=I; U_u=V_CKM^dagger; V_CKM=U_u^dagger U_d",
		UpTexture:                  yu,
		DownTexture:                yd,
		TargetUpMassesGeV:          inh.UpMassesGeV,
		TargetDownMassesGeV:        inh.DownMassesGeV,
		TargetCKM:                  ckm,
		TextureParameterSource:     "empirical boundary data under EmpiricalYukawaSeal; not finite-core output",
		Verdict:                    StatusFullTextureIngested + "; textures are full empirical inputs chosen in a transparent weak-basis convention for SVD reconstruction",
	}
}

func computeSVD(sector string, y Matrix3, target [3]float64) SVDAudit {
	// The sealed texture convention makes columns mutually orthogonal. A labeled
	// SVD is therefore obtained by column norms, normalized columns as U, and V=I.
	// This is an actual SVD for these full texture witnesses, not a diagonal-shell
	// projection. We audit orthogonality and reconstruction explicitly.
	var u Matrix3
	var s [3]float64
	for j := 0; j < 3; j++ {
		colNorm2 := 0.0
		for i := 0; i < 3; i++ {
			colNorm2 += cmplx.Abs(y[i][j]) * cmplx.Abs(y[i][j])
		}
		s[j] = math.Sqrt(colNorm2)
		if s[j] > 0 {
			for i := 0; i < 3; i++ {
				u[i][j] = y[i][j] / complex(s[j], 0)
			}
		}
	}
	v := identity()
	recon := mul(mul(u, diagonal(s)), dagger(v))
	recRes := frobNorm(sub(y, recon))
	leftRes := unitarityResidual(u)
	rightRes := unitarityResidual(v)
	colRes := columnOrthogonalityResidual(y)
	maxMass := maxAbsDiff(s, target)
	relMass := relL2Diff(s, target)
	passed := recRes < 1e-9 && leftRes < 1e-9 && rightRes < 1e-12 && colRes < 1e-9 && maxMass < 1e-9
	verdict := StatusSVDCompleted
	if !passed {
		verdict = StatusSVDReconstructionFailed
	}
	return SVDAudit{
		Sector:                      sector,
		Method:                      "generation-labeled column-orthogonal complex SVD: Y=U Sigma V^dagger with V=I",
		InputMatrix:                 y,
		LeftUnitary:                 u,
		SingularValues:              s,
		RightUnitary:                v,
		ReconstructionResidual:      recRes,
		LeftUnitarityResidual:       leftRes,
		RightUnitarityResidual:      rightRes,
		ColumnOrthogonalityResidual: colRes,
		MassTarget:                  target,
		MassMaxAbsError:             maxMass,
		MassRelativeL2Error:         relMass,
		Passed:                      passed,
		Verdict:                     verdict,
	}
}

func auditMasses(up, down SVDAudit) MassEigenvalueAudit {
	tol := 1e-9
	verified := up.Passed && down.Passed && up.MassMaxAbsError < tol && down.MassMaxAbsError < tol
	verdict := StatusMassEigenvaluesVerified
	if !verified {
		verdict = StatusSVDReconstructionFailed
	}
	return MassEigenvalueAudit{
		UpMassesGeV:               up.SingularValues,
		DownMassesGeV:             down.SingularValues,
		UpTargetGeV:               up.MassTarget,
		DownTargetGeV:             down.MassTarget,
		UpMaxAbsError:             up.MassMaxAbsError,
		DownMaxAbsError:           down.MassMaxAbsError,
		UpRelativeL2Error:         up.MassRelativeL2Error,
		DownRelativeL2Error:       down.MassRelativeL2Error,
		Tolerance:                 tol,
		Verified:                  verified,
		PhenomenologicalInputOnly: true,
		Verdict:                   verdict,
	}
}

func auditCKM(up, down SVDAudit, target Matrix3) CKMReconstructionAudit {
	recon := mul(dagger(up.LeftUnitary), down.LeftUnitary)
	resid := frobNorm(sub(recon, target))
	maxEntry := maxAbsEntry(sub(recon, target))
	absResid := absMatrixResidual(recon, target)
	unitRes := unitarityResidual(recon)
	tol := 1e-9
	verified := up.Passed && down.Passed && resid < tol && unitRes < 1e-9
	verdict := StatusCKMReconstructionVerified
	if !verified {
		verdict = StatusCKMReconstructionFailed
	}
	return CKMReconstructionAudit{
		Formula:                   "V_CKM = U_u^dagger U_d from sealed full-texture SVD left-unitary factors",
		ReconstructedCKM:          recon,
		TargetCKM:                 target,
		FrobeniusResidual:         resid,
		MaxAbsEntryResidual:       maxEntry,
		AbsMatrixResidual:         absResid,
		UnitarityResidual:         unitRes,
		Tolerance:                 tol,
		Verified:                  verified,
		DerivedFromFiniteCore:     false,
		PhenomenologicalInputOnly: true,
		Verdict:                   verdict,
	}
}

func auditFirewall(seal FullTextureSeal, masses MassEigenvalueAudit, ckm CKMReconstructionAudit) FirewallAudit {
	return FirewallAudit{
		EmpiricalSealActive:              seal.Activated,
		FullTexturesQuarantined:          seal.ExplicitlyQuarantined && !seal.DerivedFromFiniteCore,
		DoesNotRewriteGate264Underfit:    !seal.RewritesGate264Underfit,
		DoesNotClaimFiniteMassDerivation: masses.PhenomenologicalInputOnly,
		DoesNotClaimFiniteCKMDerivation:  !ckm.DerivedFromFiniteCore && ckm.PhenomenologicalInputOnly,
		DoesNotInferYukawaAction:         true,
		DoesNotInferVEVOrRGScale:         true,
		SVDIsAlgebraicReconstruction:     masses.Verified && ckm.Verified,
		FiniteCorePolluted:               false,
		Verdict:                          strings.Join([]string{StatusPhenomenologicalOnly, StatusNoNativeDerivation, StatusFullTexturesNotFiniteCore}, "; "),
	}
}

func summarize(data FullTextureData, up, down SVDAudit, masses MassEigenvalueAudit, ckm CKMReconstructionAudit, fw FirewallAudit) Summary {
	status := StatusSVDReconstructionFailed
	comment := "sealed full-texture SVD reconstruction failed"
	if up.Passed && down.Passed && masses.Verified && ckm.Verified && !fw.FiniteCorePolluted {
		status = strings.Join([]string{StatusCKMReconstructionVerified, StatusNoNativeDerivation}, "; ")
		comment = "full empirical textures reconstruct masses and CKM by SVD, while all matrices and numerical observables remain quarantined empirical boundary data"
	}
	return Summary{
		FullTexturesIngested:       data.UpTexture != Matrix3{} && data.DownTexture != Matrix3{},
		SVDCompleted:               up.Passed && down.Passed,
		MassEigenvaluesVerified:    masses.Verified,
		CKMReconstructed:           ckm.Verified,
		NativeDerivation:           false,
		EmpiricalBoundaryPreserved: !fw.FiniteCorePolluted && fw.FullTexturesQuarantined,
		Status:                     status,
		NextGate:                   "Gate 266 — Full Empirical Flavor Ledger / Lepton-PMNS and Sector Firewall Extension Audit",
		Comment:                    comment,
	}
}

func buildTruth(inh Gate264Inheritance, seal FullTextureSeal, data FullTextureData, up, down SVDAudit, masses MassEigenvalueAudit, ckm CKMReconstructionAudit) string {
	_ = seal
	_ = up
	_ = down
	return fmt.Sprintf("Gate 265 preserves Gate 264's restricted-ansatz no-go (combined residual %.6g) and activates the full-texture branch of EmpiricalYukawaSeal. In the sealed weak-basis convention %s, SVD reconstructs labeled up masses %v and down masses %v with max errors %.3g/%.3g, and V_CKM=U_u^dagger U_d matches the target with Frobenius residual %.3g. This verifies the algebraic observable reconstruction pipeline only: the full textures, masses, CKM entries, basis convention, and scale choices are empirical boundary data, not finite-core derivations.", inh.Gate264CombinedResidual, data.RightBasisConvention, masses.UpMassesGeV, masses.DownMassesGeV, masses.UpMaxAbsError, masses.DownMaxAbsError, ckm.FrobeniusResidual)
}

func fromGate264(m empiricalyukawafit.Matrix3) Matrix3 {
	var out Matrix3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			out[i][j] = m[i][j]
		}
	}
	return out
}

func identity() Matrix3 {
	var out Matrix3
	for i := 0; i < 3; i++ {
		out[i][i] = 1
	}
	return out
}

func diagonal(v [3]float64) Matrix3 {
	var out Matrix3
	for i := 0; i < 3; i++ {
		out[i][i] = complex(v[i], 0)
	}
	return out
}

func add(a, b Matrix3) Matrix3 {
	var out Matrix3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			out[i][j] = a[i][j] + b[i][j]
		}
	}
	return out
}

func sub(a, b Matrix3) Matrix3 {
	var out Matrix3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			out[i][j] = a[i][j] - b[i][j]
		}
	}
	return out
}

func mul(a, b Matrix3) Matrix3 {
	var out Matrix3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			var sum complex128
			for k := 0; k < 3; k++ {
				sum += a[i][k] * b[k][j]
			}
			out[i][j] = sum
		}
	}
	return out
}

func dagger(a Matrix3) Matrix3 {
	var out Matrix3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			out[i][j] = cmplx.Conj(a[j][i])
		}
	}
	return out
}

func frobNorm(a Matrix3) float64 {
	acc := 0.0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			v := cmplx.Abs(a[i][j])
			acc += v * v
		}
	}
	return math.Sqrt(acc)
}

func unitarityResidual(u Matrix3) float64 {
	return frobNorm(sub(mul(dagger(u), u), identity()))
}

func columnOrthogonalityResidual(y Matrix3) float64 {
	maxOff := 0.0
	for a := 0; a < 3; a++ {
		for b := a + 1; b < 3; b++ {
			var inner complex128
			for i := 0; i < 3; i++ {
				inner += cmplx.Conj(y[i][a]) * y[i][b]
			}
			maxOff = math.Max(maxOff, cmplx.Abs(inner))
		}
	}
	return maxOff
}

func maxAbsDiff(a, b [3]float64) float64 {
	m := 0.0
	for i := 0; i < 3; i++ {
		m = math.Max(m, math.Abs(a[i]-b[i]))
	}
	return m
}

func relL2Diff(a, b [3]float64) float64 {
	num := 0.0
	den := 0.0
	for i := 0; i < 3; i++ {
		d := a[i] - b[i]
		num += d * d
		den += b[i] * b[i]
	}
	if den == 0 {
		return math.Inf(1)
	}
	return math.Sqrt(num / den)
}

func maxAbsEntry(a Matrix3) float64 {
	m := 0.0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			m = math.Max(m, cmplx.Abs(a[i][j]))
		}
	}
	return m
}

func absMatrixResidual(a, b Matrix3) float64 {
	acc := 0.0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			d := cmplx.Abs(a[i][j]) - cmplx.Abs(b[i][j])
			acc += d * d
		}
	}
	return math.Sqrt(acc)
}

func FormatMatrix(m Matrix3) string {
	rows := make([]string, 0, 3)
	for i := 0; i < 3; i++ {
		cells := make([]string, 0, 3)
		for j := 0; j < 3; j++ {
			z := m[i][j]
			if math.Abs(imag(z)) < 1e-12 {
				cells = append(cells, fmt.Sprintf("%.8g", real(z)))
			} else {
				cells = append(cells, fmt.Sprintf("%.8g%+.8gi", real(z), imag(z)))
			}
		}
		rows = append(rows, "["+strings.Join(cells, ", ")+"]")
	}
	return "[" + strings.Join(rows, "; ") + "]"
}

func FormatInheritance(a Gate264Inheritance) string {
	return fmt.Sprintf("seal=%t underfit=%t fullMatrices=%t prevMasses=%t prevCKM=%t data=%t mixedScale=%t gate264Residual=%.12g up=%v down=%v verdict=%q", a.EmpiricalYukawaSealActive, a.RestrictedAnsatzViolated, a.FullEmpiricalMatricesRequired, a.MassesPreviouslyDerived, a.CKMPreviouslyDerived, a.RepresentativeDataAvailable, a.MixedScaleWarning, a.Gate264CombinedResidual, a.UpMassesGeV, a.DownMassesGeV, a.Verdict)
}

func FormatSeal(a FullTextureSeal) string {
	return fmt.Sprintf("name=%s active=%t gate=%d boundary=%q quarantined=%t derived=%t rewrites=%t reconstruct=%t massPred=%t ckmPred=%t verdict=%q", a.Name, a.Activated, a.ActivatedByGate, a.BoundaryDataKind, a.ExplicitlyQuarantined, a.DerivedFromFiniteCore, a.RewritesGate264Underfit, a.AllowsObservableReconstruction, a.AllowsMassPrediction, a.AllowsCKMPrediction, a.Verdict)
}

func FormatData(a FullTextureData) string {
	return fmt.Sprintf("source=%q representative=%t mixedScale=%t labeledSVD=%t rightBasis=%q upMassTarget=%v downMassTarget=%v sourceKind=%q verdict=%q", a.SourceLabel, a.RepresentativeNotPrecision, a.MixedScaleWarning, a.GenerationLabeledSVD, a.RightBasisConvention, a.TargetUpMassesGeV, a.TargetDownMassesGeV, a.TextureParameterSource, a.Verdict)
}

func FormatSVD(a SVDAudit) string {
	return fmt.Sprintf("sector=%q method=%q singular=%v target=%v massMax=%.3g massRel=%.3g rec=%.3g leftUnit=%.3g rightUnit=%.3g colOrth=%.3g passed=%t verdict=%q", a.Sector, a.Method, a.SingularValues, a.MassTarget, a.MassMaxAbsError, a.MassRelativeL2Error, a.ReconstructionResidual, a.LeftUnitarityResidual, a.RightUnitarityResidual, a.ColumnOrthogonalityResidual, a.Passed, a.Verdict)
}

func FormatMasses(a MassEigenvalueAudit) string {
	return fmt.Sprintf("up=%v targetUp=%v down=%v targetDown=%v upMax=%.3g downMax=%.3g upRel=%.3g downRel=%.3g tol=%.1e verified=%t phenomenological=%t verdict=%q", a.UpMassesGeV, a.UpTargetGeV, a.DownMassesGeV, a.DownTargetGeV, a.UpMaxAbsError, a.DownMaxAbsError, a.UpRelativeL2Error, a.DownRelativeL2Error, a.Tolerance, a.Verified, a.PhenomenologicalInputOnly, a.Verdict)
}

func FormatCKM(a CKMReconstructionAudit) string {
	return fmt.Sprintf("formula=%q frob=%.3g maxEntry=%.3g absResid=%.3g unit=%.3g tol=%.1e verified=%t finiteDerived=%t phenomenological=%t verdict=%q", a.Formula, a.FrobeniusResidual, a.MaxAbsEntryResidual, a.AbsMatrixResidual, a.UnitarityResidual, a.Tolerance, a.Verified, a.DerivedFromFiniteCore, a.PhenomenologicalInputOnly, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("seal=%t quarantined=%t noRewrite=%t noMassDer=%t noCKMDer=%t noAction=%t noVEV=%t algebraic=%t polluted=%t verdict=%q", a.EmpiricalSealActive, a.FullTexturesQuarantined, a.DoesNotRewriteGate264Underfit, a.DoesNotClaimFiniteMassDerivation, a.DoesNotClaimFiniteCKMDerivation, a.DoesNotInferYukawaAction, a.DoesNotInferVEVOrRGScale, a.SVDIsAlgebraicReconstruction, a.FiniteCorePolluted, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("textures=%t svd=%t masses=%t ckm=%t native=%t boundary=%t status=%q next=%q comment=%q", a.FullTexturesIngested, a.SVDCompleted, a.MassEigenvaluesVerified, a.CKMReconstructed, a.NativeDerivation, a.EmpiricalBoundaryPreserved, a.Status, a.NextGate, a.Comment)
}
