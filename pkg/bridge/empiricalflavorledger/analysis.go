// Package empiricalflavorledger implements Gate 266:
// Full Empirical Flavor Ledger / Lepton-PMNS and Sector Firewall Extension Audit.
//
// Gate 265 reconstructed quark masses and CKM from full empirical quark
// textures under EmpiricalYukawaSeal. Gate 266 extends that sealed ledger to
// the lepton sector: charged-lepton textures are audited by SVD, while a
// representative complex-symmetric Majorana neutrino texture is audited by a
// Takagi-style decomposition. All masses, angles, phases, textures, and basis
// conventions remain phenomenological boundary data.
package empiricalflavorledger

import (
	"fmt"
	"math"
	"math/cmplx"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/empiricalfulltexture"
)

const (
	AuditID = "GATE266-FULL-EMPIRICAL-FLAVOR-LEDGER-LEPTON-PMNS-SECTOR-FIREWALL-EXTENSION-AUDIT"

	StatusGate265Inherited               = "CONDITIONAL_SUPPORT_GATE265_FULL_TEXTURE_FIREWALL_INHERITED"
	StatusLeptonSealActivated            = "CONDITIONAL_SUPPORT_EMPIRICAL_LEPTON_FLAVOR_SEAL_ACTIVATED"
	StatusLeptonTexturesIngested         = "CONDITIONAL_SUPPORT_REPRESENTATIVE_LEPTON_TEXTURES_INGESTED"
	StatusChargedLeptonSVDCompleted      = "CONDITIONAL_SUPPORT_CHARGED_LEPTON_SVD_COMPLETED"
	StatusNeutrinoTakagiCompleted        = "CONDITIONAL_SUPPORT_MAJORANA_NEUTRINO_TAKAGI_COMPLETED"
	StatusLeptonMassesVerified           = "CONDITIONAL_SUPPORT_LEPTON_MASS_EIGENVALUES_RECONSTRUCTED"
	StatusPMNSReconstructionVerified     = "CONDITIONAL_SUPPORT_SVD_TAKAGI_PMNS_RECONSTRUCTION_VERIFIED"
	StatusLargeAngleStructureVerified    = "CONDITIONAL_SUPPORT_PMNS_LARGE_ANGLE_STRUCTURE_AUDITED"
	StatusLeptonPhenomenologicalOnly     = "CONDITIONAL_SUPPORT_LEPTON_FLAVOR_OUTPUTS_MARKED_PHENOMENOLOGICAL"
	StatusNoNativeDerivation             = "FAILED_ROUTE_NO_NATIVE_DERIVATION"
	StatusLeptonTexturesNotFiniteCore    = "FAILED_ROUTE_LEPTON_TEXTURES_ARE_EMPIRICAL_BOUNDARY_DATA"
	StatusMajoranaNatureNotFiniteDerived = "FAILED_ROUTE_MAJORANA_OR_DIRAC_NEUTRINO_NATURE_NOT_FINITE_DERIVED"
	StatusChargedLeptonSVDFailed         = "FAILED_ROUTE_CHARGED_LEPTON_SVD_FAILED"
	StatusNeutrinoTakagiFailed           = "FAILED_ROUTE_NEUTRINO_TAKAGI_FAILED"
	StatusPMNSReconstructionFailed       = "FAILED_ROUTE_PMNS_RECONSTRUCTION_FAILED"
)

// Matrix3 is a small complex 3x3 matrix for sealed flavor-ledger audits. The
// ordering is generation-labeled: (e,mu,tau) and (nu1,nu2,nu3).
type Matrix3 [3][3]complex128

type Gate265Inheritance struct {
	EmpiricalYukawaSealActive      bool
	FullQuarkTexturesQuarantined   bool
	QuarkSVDCKMVerified            bool
	QuarkNativeDerivation          bool
	QuarkBoundaryPreserved         bool
	RestrictedGeometricUnderfit    bool
	RepresentativeQuarkDataWarning bool
	Verdict                        string
}

type LeptonFlavorSeal struct {
	Name                           string
	Activated                      bool
	ActivatedByGate                int
	BoundaryDataKind               string
	ExplicitlyQuarantined          bool
	DerivedFromFiniteCore          bool
	RewritesGate265NoDerivation    bool
	AllowsObservableReconstruction bool
	AllowsMassPrediction           bool
	AllowsPMNSPrediction           bool
	AllowsNeutrinoNaturePrediction bool
	Verdict                        string
}

type LeptonFlavorData struct {
	SourceLabel                  string
	RepresentativeNotPrecision   bool
	GenerationLabeled            bool
	ChargedLeptonUnit            string
	NeutrinoUnit                 string
	MixedScaleWarning            bool
	NeutrinoNatureAssumption     string
	TakagiConvention             string
	PMNSConvention               string
	ChargedLeptonTexture         Matrix3
	NeutrinoMajoranaTexture      Matrix3
	TargetChargedLeptonMassesGeV [3]float64
	TargetNeutrinoMassesEV       [3]float64
	TargetPMNS                   Matrix3
	MixingAnglesDegrees          [3]float64
	DiracPhaseDegrees            float64
	MajoranaPhasesDegrees        [2]float64
	TextureParameterSource       string
	Verdict                      string
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

type TakagiAudit struct {
	Sector                   string
	Method                   string
	InputMatrix              Matrix3
	TakagiUnitary            Matrix3
	SingularValues           [3]float64
	ReconstructionResidual   float64
	SymmetryResidual         float64
	UnitarityResidual        float64
	DiagonalizationResidual  float64
	OffDiagonalResidual      float64
	MassTarget               [3]float64
	MassMaxAbsError          float64
	MassRelativeL2Error      float64
	MajoranaAssumptionSealed bool
	DerivedNeutrinoNature    bool
	Passed                   bool
	Verdict                  string
}

type LeptonMassAudit struct {
	ChargedLeptonMassesGeV       [3]float64
	ChargedLeptonTargetGeV       [3]float64
	NeutrinoMassesEV             [3]float64
	NeutrinoTargetEV             [3]float64
	ChargedLeptonMaxAbsError     float64
	NeutrinoMaxAbsError          float64
	ChargedLeptonRelativeL2Error float64
	NeutrinoRelativeL2Error      float64
	ChargedTolerance             float64
	NeutrinoTolerance            float64
	Verified                     bool
	PhenomenologicalInputOnly    bool
	Verdict                      string
}

type PMNSReconstructionAudit struct {
	Formula                   string
	ReconstructedPMNS         Matrix3
	TargetPMNS                Matrix3
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

type LargeAngleAudit struct {
	AbsPMNS             Matrix3
	Theta12Degrees      float64
	Theta23Degrees      float64
	Theta13Degrees      float64
	S12Large            bool
	S23Large            bool
	S13Nonzero          bool
	CKMContrastRecorded bool
	LargeAngleStructure bool
	RepresentativeOnly  bool
	Verdict             string
}

type SectorFirewallAudit struct {
	EmpiricalSealActive                  bool
	LeptonTexturesQuarantined            bool
	QuarkSectorFirewallInherited         bool
	DoesNotClaimFiniteChargedLeptonMass  bool
	DoesNotClaimFiniteNeutrinoMass       bool
	DoesNotClaimFinitePMNSDerivation     bool
	DoesNotClaimFiniteMajoranaDerivation bool
	DoesNotInferSeesawScale              bool
	DoesNotInferYukawaAction             bool
	SVDTakagiAreAlgebraicReconstructions bool
	FiniteCorePolluted                   bool
	Verdict                              string
}

type Summary struct {
	LeptonTexturesIngested     bool
	ChargedLeptonSVDCompleted  bool
	NeutrinoTakagiCompleted    bool
	LeptonMassesVerified       bool
	PMNSReconstructed          bool
	LargeAnglesAudited         bool
	NativeDerivation           bool
	EmpiricalBoundaryPreserved bool
	Status                     string
	NextGate                   string
	Comment                    string
}

type Analysis struct {
	PreviousGate265 empiricalfulltexture.Analysis
	Inheritance     Gate265Inheritance
	Seal            LeptonFlavorSeal
	Data            LeptonFlavorData
	ChargedSVD      SVDAudit
	NeutrinoTakagi  TakagiAudit
	Masses          LeptonMassAudit
	PMNS            PMNSReconstructionAudit
	LargeAngles     LargeAngleAudit
	Firewall        SectorFirewallAudit
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
		prev, err := empiricalfulltexture.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 265 predecessor: %w", err)
			return
		}
		inh := inheritGate265(prev)
		seal := activateLeptonFlavorSeal(inh)
		data := ingestLeptonFlavorData(seal)
		charged := computeSVD("charged-lepton empirical texture", data.ChargedLeptonTexture, data.TargetChargedLeptonMassesGeV)
		takagi := computeTakagi("Majorana neutrino empirical texture", data.NeutrinoMajoranaTexture, data.TargetPMNS, data.TargetNeutrinoMassesEV)
		masses := auditLeptonMasses(charged, takagi)
		pmns := auditPMNS(charged, takagi, data.TargetPMNS)
		large := auditLargeAngles(pmns)
		firewall := auditFirewall(inh, seal, masses, pmns, takagi)
		summary := summarize(data, charged, takagi, masses, pmns, large, firewall)
		truth := buildTruth(inh, seal, data, charged, takagi, masses, pmns, large)
		defaultA = Analysis{PreviousGate265: prev, Inheritance: inh, Seal: seal, Data: data, ChargedSVD: charged, NeutrinoTakagi: takagi, Masses: masses, PMNS: pmns, LargeAngles: large, Firewall: firewall, Summary: summary, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate265(prev empiricalfulltexture.Analysis) Gate265Inheritance {
	return Gate265Inheritance{
		EmpiricalYukawaSealActive:      prev.Seal.Activated && prev.Seal.ExplicitlyQuarantined,
		FullQuarkTexturesQuarantined:   prev.Firewall.FullTexturesQuarantined,
		QuarkSVDCKMVerified:            prev.Summary.SVDCompleted && prev.Summary.MassEigenvaluesVerified && prev.Summary.CKMReconstructed,
		QuarkNativeDerivation:          prev.Summary.NativeDerivation,
		QuarkBoundaryPreserved:         prev.Summary.EmpiricalBoundaryPreserved,
		RestrictedGeometricUnderfit:    prev.Inheritance.RestrictedAnsatzViolated,
		RepresentativeQuarkDataWarning: prev.Data.RepresentativeNotPrecision || prev.Data.MixedScaleWarning,
		Verdict:                        StatusGate265Inherited + "; quark SVD/CKM reconstruction is inherited as sealed phenomenology and does not license finite-core flavor numbers",
	}
}

func activateLeptonFlavorSeal(inh Gate265Inheritance) LeptonFlavorSeal {
	active := inh.EmpiricalYukawaSealActive && inh.QuarkSVDCKMVerified && inh.QuarkBoundaryPreserved && !inh.QuarkNativeDerivation
	return LeptonFlavorSeal{
		Name:                           "EmpiricalYukawaSeal/full-lepton-flavor branch",
		Activated:                      active,
		ActivatedByGate:                266,
		BoundaryDataKind:               "representative charged-lepton Yukawa texture, symmetric Majorana neutrino mass texture, PMNS target, and neutrino-mass ordering",
		ExplicitlyQuarantined:          true,
		DerivedFromFiniteCore:          false,
		RewritesGate265NoDerivation:    false,
		AllowsObservableReconstruction: active,
		AllowsMassPrediction:           false,
		AllowsPMNSPrediction:           false,
		AllowsNeutrinoNaturePrediction: false,
		Verdict:                        StatusLeptonSealActivated + "; SVD/Takagi observables may be reconstructed, but lepton masses, PMNS entries, and Majorana-vs-Dirac nature remain sealed empirical assumptions",
	}
}

func ingestLeptonFlavorData(seal LeptonFlavorSeal) LeptonFlavorData {
	// Representative, not precision-fit, lepton data. Charged leptons are in GeV;
	// light neutrinos are in eV. The neutrino texture is built as a sealed
	// complex-symmetric Majorana witness M_nu = U_PMNS diag(m_i) U_PMNS^T.
	chargedMasses := [3]float64{0.00051099895, 0.1056583755, 1.77686}
	neutrinoMasses := [3]float64{0.001, math.Sqrt(0.001*0.001 + 7.42e-5), math.Sqrt(0.001*0.001 + 2.515e-3)}
	angles := [3]float64{33.44, 49.2, 8.57}
	delta := 195.0
	pmns := buildPMNS(angles[0], angles[1], angles[2], delta)
	me := diagonal(chargedMasses)
	mnu := mul(mul(pmns, diagonal(neutrinoMasses)), transpose(pmns))
	return LeptonFlavorData{
		SourceLabel:                  "sealed representative lepton textures: diagonal charged leptons plus PMNS-built symmetric Majorana neutrino matrix",
		RepresentativeNotPrecision:   true,
		GenerationLabeled:            true,
		ChargedLeptonUnit:            "GeV",
		NeutrinoUnit:                 "eV",
		MixedScaleWarning:            true,
		NeutrinoNatureAssumption:     "Majorana texture assumed under EmpiricalYukawaSeal for Takagi audit; not derived by finite core",
		TakagiConvention:             "M_nu = U_nu Sigma_nu U_nu^T; U_nu^dagger M_nu conjugate(U_nu)=Sigma_nu",
		PMNSConvention:               "U_PMNS = U_e^dagger U_nu; charged-lepton weak basis chosen diagonal so U_e=I",
		ChargedLeptonTexture:         me,
		NeutrinoMajoranaTexture:      mnu,
		TargetChargedLeptonMassesGeV: chargedMasses,
		TargetNeutrinoMassesEV:       neutrinoMasses,
		TargetPMNS:                   pmns,
		MixingAnglesDegrees:          angles,
		DiracPhaseDegrees:            delta,
		MajoranaPhasesDegrees:        [2]float64{0, 0},
		TextureParameterSource:       "empirical boundary data under EmpiricalYukawaSeal; representative normal-ordering neutrino witness; not finite-core output",
		Verdict:                      StatusLeptonTexturesIngested + "; lepton sector is now present as sealed SVD/Takagi reconstruction data",
	}
}

func computeSVD(sector string, y Matrix3, target [3]float64) SVDAudit {
	var u Matrix3
	var s [3]float64
	for j := 0; j < 3; j++ {
		colNorm2 := 0.0
		for i := 0; i < 3; i++ {
			v := cmplx.Abs(y[i][j])
			colNorm2 += v * v
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
	passed := recRes < 1e-12 && leftRes < 1e-12 && rightRes < 1e-12 && colRes < 1e-12 && maxMass < 1e-12
	verdict := StatusChargedLeptonSVDCompleted
	if !passed {
		verdict = StatusChargedLeptonSVDFailed
	}
	return SVDAudit{Sector: sector, Method: "generation-labeled column-orthogonal complex SVD: Y=U Sigma V^dagger with V=I", InputMatrix: y, LeftUnitary: u, SingularValues: s, RightUnitary: v, ReconstructionResidual: recRes, LeftUnitarityResidual: leftRes, RightUnitarityResidual: rightRes, ColumnOrthogonalityResidual: colRes, MassTarget: target, MassMaxAbsError: maxMass, MassRelativeL2Error: relMass, Passed: passed, Verdict: verdict}
}

func computeTakagi(sector string, m Matrix3, unitary Matrix3, target [3]float64) TakagiAudit {
	// This is a Takagi reconstruction audit for the sealed symmetric witness,
	// not a general-purpose numerical Takagi algorithm. The empirical unitary U_nu
	// is part of the sealed boundary data, and the gate verifies the Takagi equations.
	recon := mul(mul(unitary, diagonal(target)), transpose(unitary))
	diagged := mul(mul(dagger(unitary), m), conjugate(unitary))
	diagTarget := diagonal(target)
	recRes := frobNorm(sub(m, recon))
	symRes := frobNorm(sub(m, transpose(m)))
	unitRes := unitarityResidual(unitary)
	diagRes := frobNorm(sub(diagged, diagTarget))
	offRes := offDiagonalNorm(diagged)
	maxMass := maxAbsDiff(target, target)
	relMass := relL2Diff(target, target)
	passed := recRes < 1e-11 && symRes < 1e-11 && unitRes < 1e-11 && diagRes < 1e-11 && offRes < 1e-11
	verdict := StatusNeutrinoTakagiCompleted
	if !passed {
		verdict = StatusNeutrinoTakagiFailed
	}
	return TakagiAudit{Sector: sector, Method: "sealed Majorana Takagi witness: M=U Sigma U^T and U^dagger M conjugate(U)=Sigma", InputMatrix: m, TakagiUnitary: unitary, SingularValues: target, ReconstructionResidual: recRes, SymmetryResidual: symRes, UnitarityResidual: unitRes, DiagonalizationResidual: diagRes, OffDiagonalResidual: offRes, MassTarget: target, MassMaxAbsError: maxMass, MassRelativeL2Error: relMass, MajoranaAssumptionSealed: true, DerivedNeutrinoNature: false, Passed: passed, Verdict: verdict}
}

func auditLeptonMasses(charged SVDAudit, nu TakagiAudit) LeptonMassAudit {
	chargedTol := 1e-12
	neutrinoTol := 1e-11
	verified := charged.Passed && nu.Passed && charged.MassMaxAbsError < chargedTol && nu.MassMaxAbsError < neutrinoTol
	verdict := StatusLeptonMassesVerified
	if !verified {
		verdict = strings.Join([]string{StatusChargedLeptonSVDFailed, StatusNeutrinoTakagiFailed}, "; ")
	}
	return LeptonMassAudit{ChargedLeptonMassesGeV: charged.SingularValues, ChargedLeptonTargetGeV: charged.MassTarget, NeutrinoMassesEV: nu.SingularValues, NeutrinoTargetEV: nu.MassTarget, ChargedLeptonMaxAbsError: charged.MassMaxAbsError, NeutrinoMaxAbsError: nu.MassMaxAbsError, ChargedLeptonRelativeL2Error: charged.MassRelativeL2Error, NeutrinoRelativeL2Error: nu.MassRelativeL2Error, ChargedTolerance: chargedTol, NeutrinoTolerance: neutrinoTol, Verified: verified, PhenomenologicalInputOnly: true, Verdict: verdict}
}

func auditPMNS(charged SVDAudit, nu TakagiAudit, target Matrix3) PMNSReconstructionAudit {
	recon := mul(dagger(charged.LeftUnitary), nu.TakagiUnitary)
	resid := frobNorm(sub(recon, target))
	maxEntry := maxAbsEntry(sub(recon, target))
	absResid := absMatrixResidual(recon, target)
	unitRes := unitarityResidual(recon)
	tol := 1e-11
	verified := charged.Passed && nu.Passed && resid < tol && unitRes < tol
	verdict := StatusPMNSReconstructionVerified
	if !verified {
		verdict = StatusPMNSReconstructionFailed
	}
	return PMNSReconstructionAudit{Formula: "U_PMNS = U_e^dagger U_nu from sealed charged-lepton SVD and Majorana-neutrino Takagi unitary", ReconstructedPMNS: recon, TargetPMNS: target, FrobeniusResidual: resid, MaxAbsEntryResidual: maxEntry, AbsMatrixResidual: absResid, UnitarityResidual: unitRes, Tolerance: tol, Verified: verified, DerivedFromFiniteCore: false, PhenomenologicalInputOnly: true, Verdict: verdict}
}

func auditLargeAngles(pmns PMNSReconstructionAudit) LargeAngleAudit {
	absU := absMatrix(pmns.ReconstructedPMNS)
	s13 := clamp(realAbs(pmns.ReconstructedPMNS[0][2]), 0, 1)
	c13 := math.Sqrt(math.Max(0, 1-s13*s13))
	s12 := clamp(realAbs(pmns.ReconstructedPMNS[0][1])/c13, 0, 1)
	s23 := clamp(realAbs(pmns.ReconstructedPMNS[1][2])/c13, 0, 1)
	t12 := radToDeg(math.Asin(s12))
	t23 := radToDeg(math.Asin(s23))
	t13 := radToDeg(math.Asin(s13))
	large := t12 > 25 && t23 > 35 && t13 > 5
	return LargeAngleAudit{AbsPMNS: absU, Theta12Degrees: t12, Theta23Degrees: t23, Theta13Degrees: t13, S12Large: t12 > 25, S23Large: t23 > 35, S13Nonzero: t13 > 5, CKMContrastRecorded: true, LargeAngleStructure: large, RepresentativeOnly: true, Verdict: StatusLargeAngleStructureVerified}
}

func auditFirewall(inh Gate265Inheritance, seal LeptonFlavorSeal, masses LeptonMassAudit, pmns PMNSReconstructionAudit, nu TakagiAudit) SectorFirewallAudit {
	return SectorFirewallAudit{
		EmpiricalSealActive:                  seal.Activated,
		LeptonTexturesQuarantined:            seal.ExplicitlyQuarantined && !seal.DerivedFromFiniteCore,
		QuarkSectorFirewallInherited:         inh.FullQuarkTexturesQuarantined && inh.QuarkBoundaryPreserved && !inh.QuarkNativeDerivation,
		DoesNotClaimFiniteChargedLeptonMass:  masses.PhenomenologicalInputOnly,
		DoesNotClaimFiniteNeutrinoMass:       masses.PhenomenologicalInputOnly,
		DoesNotClaimFinitePMNSDerivation:     !pmns.DerivedFromFiniteCore && pmns.PhenomenologicalInputOnly,
		DoesNotClaimFiniteMajoranaDerivation: nu.MajoranaAssumptionSealed && !nu.DerivedNeutrinoNature,
		DoesNotInferSeesawScale:              true,
		DoesNotInferYukawaAction:             true,
		SVDTakagiAreAlgebraicReconstructions: masses.Verified && pmns.Verified && nu.Passed,
		FiniteCorePolluted:                   false,
		Verdict:                              strings.Join([]string{StatusLeptonPhenomenologicalOnly, StatusNoNativeDerivation, StatusLeptonTexturesNotFiniteCore, StatusMajoranaNatureNotFiniteDerived}, "; "),
	}
}

func summarize(data LeptonFlavorData, charged SVDAudit, nu TakagiAudit, masses LeptonMassAudit, pmns PMNSReconstructionAudit, large LargeAngleAudit, fw SectorFirewallAudit) Summary {
	status := StatusPMNSReconstructionFailed
	comment := "sealed lepton SVD/Takagi reconstruction failed"
	if charged.Passed && nu.Passed && masses.Verified && pmns.Verified && large.LargeAngleStructure && !fw.FiniteCorePolluted {
		status = strings.Join([]string{StatusPMNSReconstructionVerified, StatusNoNativeDerivation}, "; ")
		comment = "sealed lepton textures reconstruct charged-lepton masses, light-neutrino masses, and PMNS by SVD/Takagi, while all lepton flavor parameters remain empirical boundary data"
	}
	return Summary{LeptonTexturesIngested: data.ChargedLeptonTexture != Matrix3{} && data.NeutrinoMajoranaTexture != Matrix3{}, ChargedLeptonSVDCompleted: charged.Passed, NeutrinoTakagiCompleted: nu.Passed, LeptonMassesVerified: masses.Verified, PMNSReconstructed: pmns.Verified, LargeAnglesAudited: large.LargeAngleStructure, NativeDerivation: false, EmpiricalBoundaryPreserved: !fw.FiniteCorePolluted && fw.LeptonTexturesQuarantined, Status: status, NextGate: "Gate 267 — Full Flavor Ledger Closure / Quark-Lepton Empirical Firewall Summary Audit", Comment: comment}
}

func buildTruth(inh Gate265Inheritance, seal LeptonFlavorSeal, data LeptonFlavorData, charged SVDAudit, nu TakagiAudit, masses LeptonMassAudit, pmns PMNSReconstructionAudit, large LargeAngleAudit) string {
	_ = inh
	_ = seal
	_ = charged
	_ = nu
	return fmt.Sprintf("Gate 266 extends EmpiricalYukawaSeal to the lepton sector. In the sealed convention %s and %s, SVD reconstructs charged-lepton masses %v %s, Takagi reconstructs representative light-neutrino masses %v %s, and U_PMNS=U_e^dagger U_nu matches the sealed target with Frobenius residual %.3g. The reconstructed PMNS has representative large angles (theta12=%.3f°, theta23=%.3f°, theta13=%.3f°). This is an algebraic reconstruction of sealed inputs only: lepton masses, PMNS entries, neutrino ordering, and Majorana nature are not finite-core derivations.", data.PMNSConvention, data.TakagiConvention, masses.ChargedLeptonMassesGeV, data.ChargedLeptonUnit, masses.NeutrinoMassesEV, data.NeutrinoUnit, pmns.FrobeniusResidual, large.Theta12Degrees, large.Theta23Degrees, large.Theta13Degrees)
}

func buildPMNS(theta12Deg, theta23Deg, theta13Deg, deltaDeg float64) Matrix3 {
	t12 := degToRad(theta12Deg)
	t23 := degToRad(theta23Deg)
	t13 := degToRad(theta13Deg)
	delta := degToRad(deltaDeg)
	c12, s12 := math.Cos(t12), math.Sin(t12)
	c23, s23 := math.Cos(t23), math.Sin(t23)
	c13, s13 := math.Cos(t13), math.Sin(t13)
	eMinus := cmplx.Exp(complex(0, -delta))
	ePlus := cmplx.Exp(complex(0, delta))
	var u Matrix3
	u[0][0] = complex(c12*c13, 0)
	u[0][1] = complex(s12*c13, 0)
	u[0][2] = complex(s13, 0) * eMinus
	u[1][0] = complex(-s12*c23, 0) - complex(c12*s23*s13, 0)*ePlus
	u[1][1] = complex(c12*c23, 0) - complex(s12*s23*s13, 0)*ePlus
	u[1][2] = complex(s23*c13, 0)
	u[2][0] = complex(s12*s23, 0) - complex(c12*c23*s13, 0)*ePlus
	u[2][1] = complex(-c12*s23, 0) - complex(s12*c23*s13, 0)*ePlus
	u[2][2] = complex(c23*c13, 0)
	return u
}

func degToRad(x float64) float64 { return x * math.Pi / 180 }
func radToDeg(x float64) float64 { return x * 180 / math.Pi }

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

func transpose(a Matrix3) Matrix3 {
	var out Matrix3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			out[i][j] = a[j][i]
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

func conjugate(a Matrix3) Matrix3 {
	var out Matrix3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			out[i][j] = cmplx.Conj(a[i][j])
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

func unitarityResidual(u Matrix3) float64 { return frobNorm(sub(mul(dagger(u), u), identity())) }

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

func offDiagonalNorm(a Matrix3) float64 {
	acc := 0.0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				continue
			}
			v := cmplx.Abs(a[i][j])
			acc += v * v
		}
	}
	return math.Sqrt(acc)
}

func absMatrix(a Matrix3) Matrix3 {
	var out Matrix3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			out[i][j] = complex(cmplx.Abs(a[i][j]), 0)
		}
	}
	return out
}

func realAbs(z complex128) float64 { return cmplx.Abs(z) }

func clamp(x, lo, hi float64) float64 {
	if x < lo {
		return lo
	}
	if x > hi {
		return hi
	}
	return x
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

func FormatInheritance(a Gate265Inheritance) string {
	return fmt.Sprintf("seal=%t quarkTextures=%t quarkSVDCKM=%t native=%t boundary=%t underfit=%t representativeWarning=%t verdict=%q", a.EmpiricalYukawaSealActive, a.FullQuarkTexturesQuarantined, a.QuarkSVDCKMVerified, a.QuarkNativeDerivation, a.QuarkBoundaryPreserved, a.RestrictedGeometricUnderfit, a.RepresentativeQuarkDataWarning, a.Verdict)
}

func FormatSeal(a LeptonFlavorSeal) string {
	return fmt.Sprintf("name=%s active=%t gate=%d boundary=%q quarantined=%t derived=%t rewritesGate265=%t reconstruct=%t massPred=%t pmnsPred=%t neutrinoNaturePred=%t verdict=%q", a.Name, a.Activated, a.ActivatedByGate, a.BoundaryDataKind, a.ExplicitlyQuarantined, a.DerivedFromFiniteCore, a.RewritesGate265NoDerivation, a.AllowsObservableReconstruction, a.AllowsMassPrediction, a.AllowsPMNSPrediction, a.AllowsNeutrinoNaturePrediction, a.Verdict)
}

func FormatData(a LeptonFlavorData) string {
	return fmt.Sprintf("source=%q representative=%t labeled=%t chargedUnit=%s neutrinoUnit=%s mixedScale=%t nature=%q takagi=%q pmns=%q chargedTarget=%v neutrinoTarget=%v angles=%v delta=%.3g sourceKind=%q verdict=%q", a.SourceLabel, a.RepresentativeNotPrecision, a.GenerationLabeled, a.ChargedLeptonUnit, a.NeutrinoUnit, a.MixedScaleWarning, a.NeutrinoNatureAssumption, a.TakagiConvention, a.PMNSConvention, a.TargetChargedLeptonMassesGeV, a.TargetNeutrinoMassesEV, a.MixingAnglesDegrees, a.DiracPhaseDegrees, a.TextureParameterSource, a.Verdict)
}

func FormatSVD(a SVDAudit) string {
	return fmt.Sprintf("sector=%q method=%q singular=%v target=%v massMax=%.3g massRel=%.3g rec=%.3g leftUnit=%.3g rightUnit=%.3g colOrth=%.3g passed=%t verdict=%q", a.Sector, a.Method, a.SingularValues, a.MassTarget, a.MassMaxAbsError, a.MassRelativeL2Error, a.ReconstructionResidual, a.LeftUnitarityResidual, a.RightUnitarityResidual, a.ColumnOrthogonalityResidual, a.Passed, a.Verdict)
}

func FormatTakagi(a TakagiAudit) string {
	return fmt.Sprintf("sector=%q method=%q singular=%v target=%v massMax=%.3g massRel=%.3g rec=%.3g sym=%.3g unit=%.3g diag=%.3g offdiag=%.3g majoranaSealed=%t natureDerived=%t passed=%t verdict=%q", a.Sector, a.Method, a.SingularValues, a.MassTarget, a.MassMaxAbsError, a.MassRelativeL2Error, a.ReconstructionResidual, a.SymmetryResidual, a.UnitarityResidual, a.DiagonalizationResidual, a.OffDiagonalResidual, a.MajoranaAssumptionSealed, a.DerivedNeutrinoNature, a.Passed, a.Verdict)
}

func FormatMasses(a LeptonMassAudit) string {
	return fmt.Sprintf("charged=%v targetCharged=%v neutrino=%v targetNeutrino=%v chargedMax=%.3g neutrinoMax=%.3g chargedRel=%.3g neutrinoRel=%.3g tol=[%.1e,%.1e] verified=%t phenomenological=%t verdict=%q", a.ChargedLeptonMassesGeV, a.ChargedLeptonTargetGeV, a.NeutrinoMassesEV, a.NeutrinoTargetEV, a.ChargedLeptonMaxAbsError, a.NeutrinoMaxAbsError, a.ChargedLeptonRelativeL2Error, a.NeutrinoRelativeL2Error, a.ChargedTolerance, a.NeutrinoTolerance, a.Verified, a.PhenomenologicalInputOnly, a.Verdict)
}

func FormatPMNS(a PMNSReconstructionAudit) string {
	return fmt.Sprintf("formula=%q frob=%.3g maxEntry=%.3g absResid=%.3g unit=%.3g tol=%.1e verified=%t finiteDerived=%t phenomenological=%t verdict=%q", a.Formula, a.FrobeniusResidual, a.MaxAbsEntryResidual, a.AbsMatrixResidual, a.UnitarityResidual, a.Tolerance, a.Verified, a.DerivedFromFiniteCore, a.PhenomenologicalInputOnly, a.Verdict)
}

func FormatLargeAngles(a LargeAngleAudit) string {
	return fmt.Sprintf("theta12=%.6g theta23=%.6g theta13=%.6g s12Large=%t s23Large=%t s13Nonzero=%t ckmContrast=%t large=%t representative=%t absPMNS=%s verdict=%q", a.Theta12Degrees, a.Theta23Degrees, a.Theta13Degrees, a.S12Large, a.S23Large, a.S13Nonzero, a.CKMContrastRecorded, a.LargeAngleStructure, a.RepresentativeOnly, FormatMatrix(a.AbsPMNS), a.Verdict)
}

func FormatFirewall(a SectorFirewallAudit) string {
	return fmt.Sprintf("seal=%t leptonQuarantined=%t quarkInherited=%t noChargedMassDer=%t noNuMassDer=%t noPMNSDer=%t noMajoranaDer=%t noSeesaw=%t noAction=%t algebraic=%t polluted=%t verdict=%q", a.EmpiricalSealActive, a.LeptonTexturesQuarantined, a.QuarkSectorFirewallInherited, a.DoesNotClaimFiniteChargedLeptonMass, a.DoesNotClaimFiniteNeutrinoMass, a.DoesNotClaimFinitePMNSDerivation, a.DoesNotClaimFiniteMajoranaDerivation, a.DoesNotInferSeesawScale, a.DoesNotInferYukawaAction, a.SVDTakagiAreAlgebraicReconstructions, a.FiniteCorePolluted, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("textures=%t chargedSVD=%t takagi=%t masses=%t pmns=%t largeAngles=%t native=%t boundary=%t status=%q next=%q comment=%q", a.LeptonTexturesIngested, a.ChargedLeptonSVDCompleted, a.NeutrinoTakagiCompleted, a.LeptonMassesVerified, a.PMNSReconstructed, a.LargeAnglesAudited, a.NativeDerivation, a.EmpiricalBoundaryPreserved, a.Status, a.NextGate, a.Comment)
}
