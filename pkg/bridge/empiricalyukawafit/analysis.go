// Package empiricalyukawafit implements Gate 264:
// Empirical Yukawa Seal Activation / Texture Amplitude Fit Audit.
//
// Gate 263 exposed the lawful three-term geometric ansatz
//
//	Y_f = alpha*tau_eta + beta*(C+C^T) + gamma*i(C-C^T),
//
// but proved that the finite core has no action functional that selects the
// amplitudes alpha,beta,gamma. Gate 264 activates the EmpiricalYukawaSeal and
// uses representative quark-sector data only as quarantined phenomenological
// stress data. It asks whether the observed flavor structure fits inside this
// restricted shell, without promoting any fitted number to a finite theorem.
package empiricalyukawafit

import (
	"fmt"
	"math"
	"math/cmplx"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/finiteyukawaaction"
	"github.com/bagherbal/asha-engine/pkg/bridge/tauetamixingpartner"
)

const (
	AuditID = "GATE264-EMPIRICAL-YUKAWA-SEAL-ACTIVATION-TEXTURE-AMPLITUDE-FIT-AUDIT"

	StatusGate263Inherited                 = "CONDITIONAL_SUPPORT_GATE263_GEOMETRIC_YUKAWA_ANSATZ_INHERITED"
	StatusEmpiricalSealActivated           = "CONDITIONAL_SUPPORT_EMPIRICAL_YUKAWA_SEAL_ACTIVATED"
	StatusRepresentativeDataIngested       = "CONDITIONAL_SUPPORT_REPRESENTATIVE_QUARK_FLAVOR_DATA_INGESTED"
	StatusAnsatzProjectionCompleted        = "CONDITIONAL_SUPPORT_GEOMETRIC_ANSATZ_PROJECTION_COMPLETED"
	StatusEmpiricalFitEstablished          = "CONDITIONAL_SUPPORT_EMPIRICAL_YUKAWA_FIT_ESTABLISHED"
	StatusFitViolatesAnsatz                = "FAILED_ROUTE_EMPIRICAL_FIT_VIOLATES_GEOMETRIC_ANSATZ"
	StatusThreeParameterUnderfit           = "FAILED_ROUTE_THREE_PARAMETER_TEXTURE_UNDERFITS_QUARK_FLAVOR_DATA"
	StatusFullEmpiricalMatricesStillNeeded = "FAILED_ROUTE_FULL_EMPIRICAL_YUKAWA_MATRICES_STILL_REQUIRED"
	StatusMassesMixingPhenomenological     = "CONDITIONAL_SUPPORT_MASSES_AND_MIXING_MARKED_PHENOMENOLOGICAL"
	StatusCKMPMNSStillSealed               = "FAILED_ROUTE_CKM_PMNS_AND_FERMION_MASSES_REMAIN_EMPIRICAL_SEAL_OUTPUTS"
)

type Complex struct {
	Re float64
	Im float64
}

type Matrix3 [3][3]complex128

type Gate263Inheritance struct {
	GeometricAnsatzAvailable        bool
	DiagonalTauSourceAvailable      bool
	HermitianOffDiagonalBasisExists bool
	FiniteActionCoefficientRule     bool
	PreviousPhysicalTextureDerived  bool
	PreviousCKMPMNSDerived          bool
	PreviousFermionMassesDerived    bool
	CandidateFormula                string
	FreeParameters                  []string
	TauEta                          Matrix3
	RealTrialityBasis               Matrix3
	PhaseTrialityBasis              Matrix3
	BasisOrthogonal                 bool
	BasisNorms                      []float64
	Verdict                         string
}

type EmpiricalSealActivation struct {
	Name                    string
	Activated               bool
	ActivatedByGate         int
	BoundaryDataKind        string
	ExplicitlyQuarantined   bool
	DerivedFromFiniteCore   bool
	NumericalOutputsDerived bool
	PhenomenologicalFitOnly bool
	RewritesGate263NoGo     bool
	AllowsStressFit         bool
	AllowsFinitePrediction  bool
	Verdict                 string
}

type QuarkFlavorData struct {
	SourceLabel                string
	RepresentativeNotPrecision bool
	MixedScaleWarning          bool
	UsesObservedMassHierarchy  bool
	UsesObservedCKMParameters  bool
	UpMassesGeV                [3]float64
	DownMassesGeV              [3]float64
	WolfensteinLambda          float64
	WolfensteinA               float64
	WolfensteinRhoBar          float64
	WolfensteinEtaBar          float64
	CKM                        Matrix3
	DataParameterCount         int
	AnsatzQuarkParameterCount  int
	ParameterDeficit           int
	Verdict                    string
}

type ProjectionFit struct {
	Sector                  string
	TargetConvention        string
	TargetFrobeniusNorm     float64
	Alpha                   float64
	Beta                    float64
	Gamma                   float64
	ProjectionFrobeniusNorm float64
	ResidualFrobeniusNorm   float64
	RelativeResidual        float64
	ExactFitTolerance       float64
	FitsExactly             bool
	OffDiagonalEqualityRule string
	TargetOffDiagonalAbs    [3]float64
	EqualOffDiagonalFailure bool
	DiagonalShapeFailure    bool
	Verdict                 string
}

type StructuralViabilityAudit struct {
	QuarkFlavorPhysicalParameters int
	RestrictedAnsatzParameters    int
	ParameterDeficit              int
	SameBasisForAllSectors        bool
	RequiresFullYukawaMatrices    bool
	AnySectorExactFit             bool
	AllSectorsExactFit            bool
	CombinedRelativeResidual      float64
	ViolatesAnsatz                bool
	CKMNumericalFitDerived        bool
	MassSpectrumDerived           bool
	Verdict                       string
}

type FirewallAudit struct {
	EmpiricalSealActive            bool
	ObservedDataQuarantined        bool
	DoesNotRewriteFiniteCore       bool
	DoesNotClaimMassPrediction     bool
	DoesNotClaimCKMPrediction      bool
	DoesNotInferVEVOrThresholds    bool
	DoesNotPromoteProjectionToLaw  bool
	Gate263NoGoPreserved           bool
	FullEmpiricalSealStillRequired bool
	FiniteCorePolluted             bool
	Verdict                        string
}

type Summary struct {
	Gate263Inherited              bool
	EmpiricalSealActivated        bool
	RepresentativeDataIngested    bool
	FitsRestrictedAnsatz          bool
	ViolatesRestrictedAnsatz      bool
	FullEmpiricalMatricesRequired bool
	MassesDerived                 bool
	CKMDerived                    bool
	Status                        string
	NextGate                      string
	Comment                       string
}

type Analysis struct {
	PreviousGate263 finiteyukawaaction.Analysis
	Inheritance     Gate263Inheritance
	Seal            EmpiricalSealActivation
	Data            QuarkFlavorData
	Fits            []ProjectionFit
	Viability       StructuralViabilityAudit
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
		prev, err := finiteyukawaaction.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 263 predecessor: %w", err)
			return
		}
		inh := inheritGate263(prev)
		seal := activateSeal(inh)
		data := ingestRepresentativeQuarkData(seal)
		fits := fitSectors(inh, data)
		viability := auditViability(data, fits)
		firewall := auditFirewall(seal, viability)
		summary := summarize(inh, seal, data, viability)
		truth := buildTruth(inh, seal, data, fits, viability)
		defaultA = Analysis{PreviousGate263: prev, Inheritance: inh, Seal: seal, Data: data, Fits: fits, Viability: viability, Firewall: firewall, Summary: summary, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate263(prev finiteyukawaaction.Analysis) Gate263Inheritance {
	tau := diagonal([3]float64{2, -2, 1})
	A := fromGaussian(prev.Inheritance.RealBasisMatrix)
	K := fromGaussian(prev.Inheritance.PhaseBasisMatrix)
	norms := []float64{hsNormSquared(tau), hsNormSquared(A), hsNormSquared(K)}
	orthogonal := math.Abs(hsInner(tau, A)) < 1e-12 && math.Abs(hsInner(tau, K)) < 1e-12 && math.Abs(hsInner(A, K)) < 1e-12
	return Gate263Inheritance{
		GeometricAnsatzAvailable:        prev.Texture.DiagonalTauSourceAvailable && prev.Texture.HermitianOffDiagonalBasisExists,
		DiagonalTauSourceAvailable:      prev.Texture.DiagonalTauSourceAvailable,
		HermitianOffDiagonalBasisExists: prev.Texture.HermitianOffDiagonalBasisExists,
		FiniteActionCoefficientRule:     prev.Texture.FiniteActionCoefficientRule,
		PreviousPhysicalTextureDerived:  prev.Summary.PhysicalYukawaTextureDerived,
		PreviousCKMPMNSDerived:          prev.Summary.CKMPMNSDerived,
		PreviousFermionMassesDerived:    prev.Summary.FermionMassesDerived,
		CandidateFormula:                prev.Texture.CandidateFormula,
		FreeParameters:                  append([]string(nil), prev.Texture.FreeParameters...),
		TauEta:                          tau,
		RealTrialityBasis:               A,
		PhaseTrialityBasis:              K,
		BasisOrthogonal:                 orthogonal,
		BasisNorms:                      norms,
		Verdict:                         StatusGate263Inherited + "; inherited tau_eta, C+C^T, and i(C-C^T) as an orthogonal M3(C) ansatz basis while preserving Gate 263's no-action-functional result",
	}
}

func activateSeal(inh Gate263Inheritance) EmpiricalSealActivation {
	return EmpiricalSealActivation{
		Name:                    "EmpiricalYukawaSeal",
		Activated:               inh.GeometricAnsatzAvailable && !inh.FiniteActionCoefficientRule,
		ActivatedByGate:         264,
		BoundaryDataKind:        "representative quark masses and CKM parameters used as phenomenological stress data",
		ExplicitlyQuarantined:   true,
		DerivedFromFiniteCore:   false,
		NumericalOutputsDerived: false,
		PhenomenologicalFitOnly: true,
		RewritesGate263NoGo:     false,
		AllowsStressFit:         true,
		AllowsFinitePrediction:  false,
		Verdict:                 StatusEmpiricalSealActivated + "; observed flavor data may be used only to test the restricted shell, not to derive masses or mixing from the finite core",
	}
}

func ingestRepresentativeQuarkData(seal EmpiricalSealActivation) QuarkFlavorData {
	// Representative PDG-like quark-sector stress data. Values are deliberately
	// stored as sealed phenomenology, not as finite-core constants. The mixed
	// mass-scale warning prevents interpreting the numbers as an RG-consistent
	// precision fit.
	data := QuarkFlavorData{
		SourceLabel:                "representative observed quark-sector stress data: approximate quark masses plus Wolfenstein CKM parameters",
		RepresentativeNotPrecision: true,
		MixedScaleWarning:          true,
		UsesObservedMassHierarchy:  seal.Activated,
		UsesObservedCKMParameters:  seal.Activated,
		UpMassesGeV:                [3]float64{0.00216, 1.27, 172.57},
		DownMassesGeV:              [3]float64{0.00467, 0.0934, 4.18},
		WolfensteinLambda:          0.22501,
		WolfensteinA:               0.826,
		WolfensteinRhoBar:          0.159,
		WolfensteinEtaBar:          0.352,
		DataParameterCount:         10, // six quark masses + four CKM parameters.
		AnsatzQuarkParameterCount:  6,  // alpha,beta,gamma independently for up and down sectors.
	}
	data.ParameterDeficit = data.DataParameterCount - data.AnsatzQuarkParameterCount
	data.CKM = buildCKM(data.WolfensteinLambda, data.WolfensteinA, data.WolfensteinRhoBar, data.WolfensteinEtaBar)
	data.Verdict = StatusRepresentativeDataIngested + "; data are intentionally scale/method representative and sealed, sufficient only for an ansatz-viability stress test"
	return data
}

func fitSectors(inh Gate263Inheritance, data QuarkFlavorData) []ProjectionFit {
	yu := hermitianLeftTextureProxy(data.CKM, data.UpMassesGeV)
	yd := diagonal(data.DownMassesGeV)
	fits := []ProjectionFit{
		project("up-sector Hermitian left-texture proxy", "Y_u^proxy = V_CKM diag(m_u,m_c,m_t) V_CKM†", yu, inh),
		project("down-sector diagonal weak-basis proxy", "Y_d^proxy = diag(m_d,m_s,m_b)", yd, inh),
	}
	return fits
}

func project(sector, convention string, target Matrix3, inh Gate263Inheritance) ProjectionFit {
	tau := inh.TauEta
	A := inh.RealTrialityBasis
	K := inh.PhaseTrialityBasis
	alpha := hsInner(tau, target) / hsNormSquared(tau)
	beta := hsInner(A, target) / hsNormSquared(A)
	gamma := hsInner(K, target) / hsNormSquared(K)
	projection := add(add(scale(tau, alpha), scale(A, beta)), scale(K, gamma))
	residual := sub(target, projection)
	tNorm := math.Sqrt(hsNormSquared(target))
	pNorm := math.Sqrt(hsNormSquared(projection))
	rNorm := math.Sqrt(hsNormSquared(residual))
	rel := math.Inf(1)
	if tNorm > 0 {
		rel = rNorm / tNorm
	}
	off := offDiagAbs(target)
	equalFail := !nearlyEqual(off[0], off[1], 1e-6) || !nearlyEqual(off[0], off[2], 1e-6)
	diagFail := diagonalShapeFailure(target)
	tol := 1e-6
	verdict := StatusFitViolatesAnsatz
	if rel <= tol {
		verdict = StatusEmpiricalFitEstablished
	}
	return ProjectionFit{
		Sector:                  sector,
		TargetConvention:        convention,
		TargetFrobeniusNorm:     tNorm,
		Alpha:                   alpha,
		Beta:                    beta,
		Gamma:                   gamma,
		ProjectionFrobeniusNorm: pNorm,
		ResidualFrobeniusNorm:   rNorm,
		RelativeResidual:        rel,
		ExactFitTolerance:       tol,
		FitsExactly:             rel <= tol,
		OffDiagonalEqualityRule: "restricted ansatz forces |Y12|=|Y13|=|Y23|=sqrt(beta^2+gamma^2)",
		TargetOffDiagonalAbs:    off,
		EqualOffDiagonalFailure: equalFail,
		DiagonalShapeFailure:    diagFail,
		Verdict:                 verdict,
	}
}

func auditViability(data QuarkFlavorData, fits []ProjectionFit) StructuralViabilityAudit {
	any := false
	all := len(fits) > 0
	sumResidual2 := 0.0
	sumTarget2 := 0.0
	for _, f := range fits {
		any = any || f.FitsExactly
		all = all && f.FitsExactly
		sumResidual2 += f.ResidualFrobeniusNorm * f.ResidualFrobeniusNorm
		sumTarget2 += f.TargetFrobeniusNorm * f.TargetFrobeniusNorm
	}
	combined := math.Inf(1)
	if sumTarget2 > 0 {
		combined = math.Sqrt(sumResidual2 / sumTarget2)
	}
	violates := !all || data.ParameterDeficit > 0
	verdict := StatusEmpiricalFitEstablished
	if violates {
		verdict = strings.Join([]string{StatusFitViolatesAnsatz, StatusThreeParameterUnderfit, StatusFullEmpiricalMatricesStillNeeded}, "; ")
	}
	return StructuralViabilityAudit{
		QuarkFlavorPhysicalParameters: data.DataParameterCount,
		RestrictedAnsatzParameters:    data.AnsatzQuarkParameterCount,
		ParameterDeficit:              data.ParameterDeficit,
		SameBasisForAllSectors:        true,
		RequiresFullYukawaMatrices:    violates,
		AnySectorExactFit:             any,
		AllSectorsExactFit:            all,
		CombinedRelativeResidual:      combined,
		ViolatesAnsatz:                violates,
		CKMNumericalFitDerived:        false,
		MassSpectrumDerived:           false,
		Verdict:                       verdict,
	}
}

func auditFirewall(seal EmpiricalSealActivation, viability StructuralViabilityAudit) FirewallAudit {
	return FirewallAudit{
		EmpiricalSealActive:            seal.Activated,
		ObservedDataQuarantined:        seal.ExplicitlyQuarantined,
		DoesNotRewriteFiniteCore:       !seal.RewritesGate263NoGo,
		DoesNotClaimMassPrediction:     !viability.MassSpectrumDerived,
		DoesNotClaimCKMPrediction:      !viability.CKMNumericalFitDerived,
		DoesNotInferVEVOrThresholds:    true,
		DoesNotPromoteProjectionToLaw:  viability.ViolatesAnsatz,
		Gate263NoGoPreserved:           true,
		FullEmpiricalSealStillRequired: viability.RequiresFullYukawaMatrices,
		FiniteCorePolluted:             false,
		Verdict:                        StatusMassesMixingPhenomenological + "; projection residuals are diagnostic stress-test outputs only and do not alter finite-core theorem status",
	}
}

func summarize(inh Gate263Inheritance, seal EmpiricalSealActivation, data QuarkFlavorData, viability StructuralViabilityAudit) Summary {
	status := StatusEmpiricalFitEstablished
	next := "Gate 265 — Empirical Full Texture Seal / SVD-CKM Observable Reconstruction Audit"
	comment := "the three-term tau_eta/triality ansatz passed the empirical stress test"
	if viability.ViolatesAnsatz {
		status = StatusFitViolatesAnsatz
		comment = "the three-term tau_eta/triality shell is structurally valuable but too restrictive for representative quark masses plus CKM data; full empirical Yukawa matrices remain sealed boundary data"
	}
	return Summary{
		Gate263Inherited:              inh.GeometricAnsatzAvailable,
		EmpiricalSealActivated:        seal.Activated,
		RepresentativeDataIngested:    data.UsesObservedMassHierarchy && data.UsesObservedCKMParameters,
		FitsRestrictedAnsatz:          !viability.ViolatesAnsatz,
		ViolatesRestrictedAnsatz:      viability.ViolatesAnsatz,
		FullEmpiricalMatricesRequired: viability.RequiresFullYukawaMatrices,
		MassesDerived:                 false,
		CKMDerived:                    false,
		Status:                        status,
		NextGate:                      next,
		Comment:                       comment,
	}
}

func buildTruth(inh Gate263Inheritance, seal EmpiricalSealActivation, data QuarkFlavorData, fits []ProjectionFit, viability StructuralViabilityAudit) string {
	_ = inh
	_ = seal
	sectorBits := make([]string, 0, len(fits))
	for _, f := range fits {
		sectorBits = append(sectorBits, fmt.Sprintf("%s residual=%.6g", f.Sector, f.RelativeResidual))
	}
	sort.Strings(sectorBits)
	return fmt.Sprintf("Gate 264 activates the EmpiricalYukawaSeal and stress-tests the derived three-term shell against quarantined representative quark data. The data ledger has %d physical quark-flavor parameters versus %d real ansatz parameters for independent up/down shells. Orthogonal projection into {tau_eta,C+C^T,i(C-C^T)} leaves [%s], so the restricted shell does not fit the empirical flavor structure. This is a sealed phenomenological no-go for the minimal ansatz, not a pollution of the finite core: masses, CKM, VEV, thresholds, and full Yukawa matrices remain empirical boundary data.", data.DataParameterCount, data.AnsatzQuarkParameterCount, strings.Join(sectorBits, "; "))
}

func fromGaussian(m tauetamixingpartner.Matrix3) Matrix3 {
	var out Matrix3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			out[i][j] = complex(float64(m[i][j].Re), float64(m[i][j].Im))
		}
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

func buildCKM(lambda, A, rhobar, etabar float64) Matrix3 {
	s12 := lambda
	s23 := A * lambda * lambda
	rho := rhobar
	eta := etabar
	s13 := A * math.Pow(lambda, 3) * math.Hypot(rho, eta)
	delta := math.Atan2(eta, rho)
	c12 := math.Sqrt(1 - s12*s12)
	c23 := math.Sqrt(1 - s23*s23)
	c13 := math.Sqrt(1 - s13*s13)
	eid := cmplx.Exp(complex(0, delta))
	emid := cmplx.Exp(complex(0, -delta))
	var v Matrix3
	v[0][0] = complex(c12*c13, 0)
	v[0][1] = complex(s12*c13, 0)
	v[0][2] = complex(s13, 0) * emid
	v[1][0] = complex(-s12*c23, 0) - complex(c12*s23*s13, 0)*eid
	v[1][1] = complex(c12*c23, 0) - complex(s12*s23*s13, 0)*eid
	v[1][2] = complex(s23*c13, 0)
	v[2][0] = complex(s12*s23, 0) - complex(c12*c23*s13, 0)*eid
	v[2][1] = complex(-c12*s23, 0) - complex(s12*c23*s13, 0)*eid
	v[2][2] = complex(c23*c13, 0)
	return v
}

func hermitianLeftTextureProxy(v Matrix3, masses [3]float64) Matrix3 {
	d := diagonal(masses)
	return mul(mul(v, d), dagger(v))
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

func scale(a Matrix3, s float64) Matrix3 {
	var out Matrix3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			out[i][j] = complex(s, 0) * a[i][j]
		}
	}
	return out
}

func mul(a, b Matrix3) Matrix3 {
	var out Matrix3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sum := complex(0, 0)
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

func hsInner(a, b Matrix3) float64 {
	sum := complex(0, 0)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sum += cmplx.Conj(a[i][j]) * b[i][j]
		}
	}
	return real(sum)
}

func hsNormSquared(a Matrix3) float64 {
	return hsInner(a, a)
}

func offDiagAbs(a Matrix3) [3]float64 {
	return [3]float64{cmplx.Abs(a[0][1]), cmplx.Abs(a[0][2]), cmplx.Abs(a[1][2])}
}

func diagonalShapeFailure(a Matrix3) bool {
	// If beta=gamma=0, the ansatz diagonal lies on span{(2,-2,1)}. A positive
	// diagonal hierarchy cannot exactly live there unless the projected residual
	// vanishes. This diagnostic is intentionally simple and only flags obvious
	// incompatible positive diagonal proxies.
	off := offDiagAbs(a)
	diagonalOnly := off[0] < 1e-12 && off[1] < 1e-12 && off[2] < 1e-12
	if !diagonalOnly {
		return false
	}
	d := [3]float64{real(a[0][0]), real(a[1][1]), real(a[2][2])}
	alpha := (2*d[0] - 2*d[1] + d[2]) / 9.0
	return math.Abs(d[0]-2*alpha) > 1e-9 || math.Abs(d[1]+2*alpha) > 1e-9 || math.Abs(d[2]-alpha) > 1e-9
}

func nearlyEqual(a, b, tol float64) bool {
	scale := math.Max(1, math.Max(math.Abs(a), math.Abs(b)))
	return math.Abs(a-b) <= tol*scale
}
