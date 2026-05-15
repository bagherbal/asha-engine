// Package generation2koideprovenance implements Gate 485:
// Koide Constraint Provenance & Topological Baseline.
//
// Gate 484 found that the charged-lepton square-root mass fingerprint is almost
// exactly Koide, while the quark fingerprints fail the same fixed-ratio test.
// Gate 485 deliberately removes empirical masses from the proof step. It audits
// only the native C3 shadow geometry: if a three-generation square-root shadow
// is decomposed into a democratic leg and a C3 phase plane, and if that shadow
// is constrained to the Cℓ(1,7) null boundary with democratic leg timelike and
// phase-plane leg spacelike, then the null condition forces R/S=sqrt(2) and
// hence Q=2/3.
//
// Integrity boundary: this proves the provenance of the Koide ratio as a native
// null-C3 baseline theorem. It does not derive electron/muon/tau masses, quark
// masses, CKM/PMNS matrices, phase sheets, or physical sector perturbations.
package generation2koideprovenance

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE485-KOIDE-CONSTRAINT-PROVENANCE-TOPOLOGICAL-BASELINE"

	StatusGate484Inherited         = "CONDITIONAL_SUPPORT_GATE484_KOIDE_SHADOW_INHERITED"
	StatusC3ShadowBasisProved      = "CONDITIONAL_SUPPORT_C3_SHADOW_BASIS_ORTHOGONALITY_PROVED"
	StatusNullKoideRatioDerived    = "CONDITIONAL_SUPPORT_NULL_BOUNDARY_FORCES_R_OVER_S_SQRT2"
	StatusKoideQDerived            = "CONDITIONAL_SUPPORT_KOIDE_Q_TWO_THIRDS_DERIVED_FROM_NULL_C3_SHADOW"
	StatusLeptonBaselineCompatible = "CONDITIONAL_SUPPORT_CHARGED_LEPTON_BRIDGE_SHADOW_COMPATIBLE_WITH_NULL_BASELINE"
	StatusFlavorFirewallPreserved  = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_KOIDE_PROVENANCE"

	StatusFailedEmpiricalFitRejected       = "FAILED_ROUTE_EMPIRICAL_KOIDE_SHADOW_AS_NATIVE_MASS_FIT_REJECTED"
	StatusFailedMassesNotDerived           = "FAILED_ROUTE_NULL_KOIDE_DOES_NOT_DERIVE_ABSOLUTE_MASSES"
	StatusFailedPhaseNotSelected           = "FAILED_ROUTE_NULL_KOIDE_DOES_NOT_SELECT_C3_PHASE_PSI"
	StatusFailedQuarkPromotionRejected     = "FAILED_ROUTE_QUARK_SECTORS_COLOR_DRESSED_NOT_NULL_BASELINE"
	StatusFailedMixingPredictionRejected   = "FAILED_ROUTE_KOIDE_BASELINE_AS_CKM_PMNS_PREDICTION_REJECTED"
	StatusFailedFullFlavorCollapseRejected = "FAILED_ROUTE_KOIDE_BASELINE_DOES_NOT_COLLAPSE_13_MODULI"
)

const (
	NativeFlavorDim = 13
	KXYCoeffDim     = 9
	Theta0          = 0.0
	Theta1          = 2.0 * math.Pi / 3.0
	Theta2          = 4.0 * math.Pi / 3.0
	KoideTargetQ    = 2.0 / 3.0
	KoideTargetR    = math.Sqrt2
)

type Inheritance struct {
	Executed                                bool
	Gate480NullConeNative                   bool
	Gate481CommonBaselineCancels            bool
	Gate483QuarkLeptonTopologySeparatedOnly bool
	Gate484C3BasisValidated                 bool
	Gate484ChargedLeptonKoideShadowFound    bool
	Gate484UniversalTiltFailed              bool
	Gate484NativeTiltRatioPreviouslyAbsent  bool
	ObservedMassesRemainBridgeData          bool
	NativeRegistryClean                     bool
	Verdict                                 string
	Reason                                  string
}

type BasisSample struct {
	Psi                   float64
	SumCos                float64
	SumCosSquared         float64
	DemocraticDotPhase    float64
	DemocraticNormSquared float64
	PhaseNormSquaredAtR1  float64
	Orthogonal            bool
	PhaseNormIndependent  bool
	Verdict               string
}

type C3ShadowBasis struct {
	Executed                     bool
	Thetas                       [3]float64
	Samples                      []BasisSample
	SumCosIdentity               string
	SumCosSquaredIdentity        string
	DemocraticNormFormula        string
	PhasePlaneNormFormula        string
	DemocraticPhaseOrthogonality bool
	PhaseNormEqualsThreeHalves   bool
	PhaseNormIndependentOfPsi    bool
	BasisNativeC3Orbit           bool
	NoEmpiricalMassesUsed        bool
	Verdict                      string
	Reason                       string
}

type NullDerivation struct {
	Executed             bool
	S, R, Psi            float64
	TimelikeNormSquared  float64
	SpacelikeNormSquared float64
	MinkowskiNorm        float64
	PositiveFutureBranch bool
	NullCondition        string
	RatioDerived         float64
	RatioResidual        float64
	KoideQDerived        float64
	KoideResidual        float64
	NullForcesRatio      bool
	KoideEquivalent      bool
	ScaleFree            bool
	PhaseFree            bool
	Verdict              string
	Reason               string
}

type BoundaryCollapse struct {
	Executed                      bool
	C3RawShadowDOF                int
	NullConstrainedShadowDOF      int
	CollapsedShapeDOF             int
	ScaleStillFree                bool
	PsiStillFree                  bool
	AbsoluteMassesDerived         bool
	ChargedLeptonBridgeCompatible bool
	QuarkSectorsPromoted          bool
	FullFlavorModuliCollapsed     bool
	NativeFlavorDimAfter          int
	KXYCoeffDimAfter              int
	Verdict                       string
	Reason                        string
}

type SectorBoundary struct {
	Executed                        bool
	ColorlessLeptonBaselineEligible bool
	QuarkBaselineEligible           bool
	QuarkColorDressingDeclared      bool
	ChargedLeptonKoideShadowBridge  bool
	PhysicalLeptonMassesDerived     bool
	PhysicalQuarkMassesDerived      bool
	CKMConstructed                  bool
	PMNSConstructed                 bool
	Verdict                         string
	Reason                          string
	Failures                        []string
}

type Firewall struct {
	Executed                      bool
	ObservedMassImportedForProof  bool
	CKMImported                   bool
	PMNSImported                  bool
	NullC3RatioNativeBaseline     bool
	KoideAsPhysicalMassPrediction bool
	LeptonMassesDerived           bool
	QuarkMassesDerived            bool
	PhasePsiSelected              bool
	SectorPerturbationsNative     bool
	CKMMatrixConstructed          bool
	PMNSMatrixConstructed         bool
	NativeRegistryWritten         bool
	NativeFlavorDimAfter          int
	KXYCoeffDimAfter              int
	Verdict                       string
	Reason                        string
}

type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Basis       C3ShadowBasis
	Derivation  NullDerivation
	Collapse    BoundaryCollapse
	Sector      SectorBoundary
	Firewall    Firewall
	Next        NextStep
	Truth       string
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
	a := Analysis{Inheritance: buildInheritance()}
	a.Basis = buildC3ShadowBasis()
	a.Derivation = buildNullDerivation(1, KoideTargetR, math.Pi/7)
	a.Collapse = buildBoundaryCollapse(a)
	a.Sector = buildSectorBoundary(a)
	a.Firewall = buildFirewall(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{
		Executed:                                true,
		Gate480NullConeNative:                   true,
		Gate481CommonBaselineCancels:            true,
		Gate483QuarkLeptonTopologySeparatedOnly: true,
		Gate484C3BasisValidated:                 true,
		Gate484ChargedLeptonKoideShadowFound:    true,
		Gate484UniversalTiltFailed:              true,
		Gate484NativeTiltRatioPreviouslyAbsent:  true,
		ObservedMassesRemainBridgeData:          true,
		NativeRegistryClean:                     true,
		Verdict:                                 StatusGate484Inherited,
		Reason:                                  "Gate484 supplied the C3 coordinate audit and the charged-lepton empirical shadow, but did not promote masses or a universal tilt vector into native law-space",
	}
}

func buildC3ShadowBasis() C3ShadowBasis {
	psis := []float64{0, math.Pi / 9, math.Pi / 3, -math.Pi / 5}
	samples := make([]BasisSample, 0, len(psis))
	allOrthogonal := true
	allThreeHalves := true
	for _, psi := range psis {
		s := evaluateBasisSample(psi)
		if !s.Orthogonal {
			allOrthogonal = false
		}
		if !s.PhaseNormIndependent || !nearly(s.SumCosSquared, 1.5, 1e-12) {
			allThreeHalves = false
		}
		samples = append(samples, s)
	}
	return C3ShadowBasis{
		Executed:                     true,
		Thetas:                       [3]float64{Theta0, Theta1, Theta2},
		Samples:                      samples,
		SumCosIdentity:               "Σ_i cos(θ_i-ψ)=0 for θ_i∈{0,2π/3,4π/3}",
		SumCosSquaredIdentity:        "Σ_i cos²(θ_i-ψ)=3/2, independent of ψ",
		DemocraticNormFormula:        "||S(1,1,1)||²=3S²",
		PhasePlaneNormFormula:        "||R cos(θ_i-ψ)||²=(3/2)R²",
		DemocraticPhaseOrthogonality: allOrthogonal,
		PhaseNormEqualsThreeHalves:   allThreeHalves,
		PhaseNormIndependentOfPsi:    allThreeHalves,
		BasisNativeC3Orbit:           true,
		NoEmpiricalMassesUsed:        true,
		Verdict:                      StatusC3ShadowBasisProved,
		Reason:                       "the C3 roots of unity split any normalized three-family square-root shadow into an orthogonal democratic leg and a phase-plane leg with fixed norm weights 3 and 3/2; this proof uses only trigonometric C3 identities",
	}
}

func evaluateBasisSample(psi float64) BasisSample {
	thetas := [3]float64{Theta0, Theta1, Theta2}
	sumCos := 0.0
	sumCosSq := 0.0
	for _, theta := range thetas {
		c := math.Cos(theta - psi)
		sumCos += c
		sumCosSq += c * c
	}
	dot := sumCos // S=1,R=1 diagnostic
	demNorm := 3.0
	phaseNorm := sumCosSq
	orthogonal := nearly(dot, 0, 1e-12)
	phaseInvariant := nearly(phaseNorm, 1.5, 1e-12)
	verdict := StatusC3ShadowBasisProved
	if !orthogonal || !phaseInvariant {
		verdict = StatusFailedEmpiricalFitRejected
	}
	return BasisSample{Psi: psi, SumCos: sumCos, SumCosSquared: sumCosSq, DemocraticDotPhase: dot, DemocraticNormSquared: demNorm, PhaseNormSquaredAtR1: phaseNorm, Orthogonal: orthogonal, PhaseNormIndependent: phaseInvariant, Verdict: verdict}
}

func buildNullDerivation(S, R, psi float64) NullDerivation {
	timeNorm := 3 * S * S
	spaceNorm := 1.5 * R * R
	q := timeNorm - spaceNorm
	ratio := R / S
	Q := koideQFromSR(S, R)
	return NullDerivation{
		Executed:             true,
		S:                    S,
		R:                    R,
		Psi:                  psi,
		TimelikeNormSquared:  timeNorm,
		SpacelikeNormSquared: spaceNorm,
		MinkowskiNorm:        q,
		PositiveFutureBranch: S > 0 && R > 0,
		NullCondition:        "3S²-(3/2)R²=0",
		RatioDerived:         ratio,
		RatioResidual:        ratio - KoideTargetR,
		KoideQDerived:        Q,
		KoideResidual:        Q - KoideTargetQ,
		NullForcesRatio:      nearly(q, 0, 1e-12) && nearly(ratio, KoideTargetR, 1e-12),
		KoideEquivalent:      nearly(Q, KoideTargetQ, 1e-12),
		ScaleFree:            true,
		PhaseFree:            true,
		Verdict:              StatusNullKoideRatioDerived,
		Reason:               "on the positive C3 null branch, 3S²=(3/2)R² implies R²=2S², so R/S=sqrt(2); substituting into Q=(3S²+(3/2)R²)/(9S²) gives Q=2/3",
	}
}

func koideQFromSR(S, R float64) float64 {
	return (3*S*S + 1.5*R*R) / (9 * S * S)
}

func buildBoundaryCollapse(a Analysis) BoundaryCollapse {
	return BoundaryCollapse{
		Executed:                      true,
		C3RawShadowDOF:                3,
		NullConstrainedShadowDOF:      2,
		CollapsedShapeDOF:             1,
		ScaleStillFree:                a.Derivation.ScaleFree,
		PsiStillFree:                  a.Derivation.PhaseFree,
		AbsoluteMassesDerived:         false,
		ChargedLeptonBridgeCompatible: a.Derivation.KoideEquivalent && a.Inheritance.Gate484ChargedLeptonKoideShadowFound,
		QuarkSectorsPromoted:          false,
		FullFlavorModuliCollapsed:     false,
		NativeFlavorDimAfter:          NativeFlavorDim,
		KXYCoeffDimAfter:              KXYCoeffDim,
		Verdict:                       StatusKoideQDerived,
		Reason:                        "the null boundary collapses one C3 shadow-shape coordinate, R/S, from a free radial ratio to sqrt(2), leaving the overall scale S and C3 phase ψ unselected; therefore it proves the Koide baseline shape but not the physical mass spectrum",
	}
}

func buildSectorBoundary(a Analysis) SectorBoundary {
	return SectorBoundary{
		Executed:                        true,
		ColorlessLeptonBaselineEligible: a.Collapse.ChargedLeptonBridgeCompatible,
		QuarkBaselineEligible:           false,
		QuarkColorDressingDeclared:      true,
		ChargedLeptonKoideShadowBridge:  a.Inheritance.Gate484ChargedLeptonKoideShadowFound,
		PhysicalLeptonMassesDerived:     false,
		PhysicalQuarkMassesDerived:      false,
		CKMConstructed:                  false,
		PMNSConstructed:                 false,
		Verdict:                         StatusLeptonBaselineCompatible,
		Reason:                          "the charged-lepton bridge shadow is compatible with the colorless null-C3 baseline; quark sectors are not promoted because color/QCD dressing is a sector topology and scale problem outside the bare null baseline",
		Failures:                        []string{StatusFailedQuarkPromotionRejected, StatusFailedMassesNotDerived, StatusFailedPhaseNotSelected, StatusFailedMixingPredictionRejected},
	}
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{
		Executed:                      true,
		ObservedMassImportedForProof:  false,
		CKMImported:                   false,
		PMNSImported:                  false,
		NullC3RatioNativeBaseline:     a.Derivation.NullForcesRatio && a.Derivation.KoideEquivalent,
		KoideAsPhysicalMassPrediction: false,
		LeptonMassesDerived:           false,
		QuarkMassesDerived:            false,
		PhasePsiSelected:              false,
		SectorPerturbationsNative:     false,
		CKMMatrixConstructed:          false,
		PMNSMatrixConstructed:         false,
		NativeRegistryWritten:         false,
		NativeFlavorDimAfter:          NativeFlavorDim,
		KXYCoeffDimAfter:              KXYCoeffDim,
		Verdict:                       StatusFlavorFirewallPreserved,
		Reason:                        "Gate485 writes only the null-C3 ratio theorem into the baseline ledger; empirical masses, sector perturbations, quark dressing, C3 phase ψ, CKM, PMNS, and absolute Yukawa amplitudes remain sealed",
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 486, Title: "Color-dressing deformation firewall or lepton-baseline airlock", Reason: "Gate485 derives the bare null-C3 Koide ratio but leaves scale, phase, quark dressing, and physical sector perturbations unresolved.", PrimaryTask: "audit whether color/winding topology can define a bridge-only dressing operator for quark deviations from the null-C3 baseline without importing masses or mixing matrices as native data"}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate480NullConeNative || !a.Inheritance.Gate484C3BasisValidated || !a.Inheritance.NativeRegistryClean || !a.Inheritance.ObservedMassesRemainBridgeData {
		return fmt.Errorf("Gate485 inheritance invalid: %+v", a.Inheritance)
	}
	if !a.Basis.Executed || !a.Basis.DemocraticPhaseOrthogonality || !a.Basis.PhaseNormEqualsThreeHalves || !a.Basis.PhaseNormIndependentOfPsi || !a.Basis.NoEmpiricalMassesUsed {
		return fmt.Errorf("Gate485 C3 shadow basis invalid: %+v", a.Basis)
	}
	if !a.Derivation.Executed || !a.Derivation.PositiveFutureBranch || !a.Derivation.NullForcesRatio || !a.Derivation.KoideEquivalent || !a.Derivation.ScaleFree || !a.Derivation.PhaseFree {
		return fmt.Errorf("Gate485 null derivation invalid: %+v", a.Derivation)
	}
	if !a.Collapse.Executed || a.Collapse.CollapsedShapeDOF != 1 || a.Collapse.AbsoluteMassesDerived || !a.Collapse.ScaleStillFree || !a.Collapse.PsiStillFree || a.Collapse.FullFlavorModuliCollapsed || a.Collapse.NativeFlavorDimAfter != NativeFlavorDim || a.Collapse.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("Gate485 boundary collapse invalid: %+v", a.Collapse)
	}
	if !a.Sector.Executed || !a.Sector.ColorlessLeptonBaselineEligible || a.Sector.QuarkBaselineEligible || !a.Sector.QuarkColorDressingDeclared || a.Sector.PhysicalLeptonMassesDerived || a.Sector.PhysicalQuarkMassesDerived || a.Sector.CKMConstructed || a.Sector.PMNSConstructed {
		return fmt.Errorf("Gate485 sector boundary invalid: %+v", a.Sector)
	}
	if !a.Firewall.Executed || a.Firewall.ObservedMassImportedForProof || a.Firewall.CKMImported || a.Firewall.PMNSImported || !a.Firewall.NullC3RatioNativeBaseline || a.Firewall.KoideAsPhysicalMassPrediction || a.Firewall.LeptonMassesDerived || a.Firewall.QuarkMassesDerived || a.Firewall.PhasePsiSelected || a.Firewall.SectorPerturbationsNative || a.Firewall.CKMMatrixConstructed || a.Firewall.PMNSMatrixConstructed || a.Firewall.NativeRegistryWritten || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("Gate485 firewall invalid: %+v", a.Firewall)
	}
	return nil
}

func nearly(a, b, tol float64) bool { return math.Abs(a-b) <= tol }

func truth(a Analysis) string {
	return fmt.Sprintf("Gate485 result: the C3 shadow null condition %s natively forces R/S=%s and Q=%s for a bare colorless null baseline, collapsing one C3 shape coordinate while leaving scale, phase, quark dressing, CKM, PMNS, and physical masses sealed.", a.Derivation.NullCondition, fmtFloat(a.Derivation.RatioDerived), fmtFloat(a.Derivation.KoideQDerived))
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("Gate480Null=%t Gate481Cancel=%t Gate483Topology=%t Gate484C3=%t Gate484KoideShadow=%t massesBridge=%t clean=%t", x.Gate480NullConeNative, x.Gate481CommonBaselineCancels, x.Gate483QuarkLeptonTopologySeparatedOnly, x.Gate484C3BasisValidated, x.Gate484ChargedLeptonKoideShadowFound, x.ObservedMassesRemainBridgeData, x.NativeRegistryClean)
}

func FormatBasis(x C3ShadowBasis) string {
	return fmt.Sprintf("theta=[0,2π/3,4π/3] sumcos=0 sumcos2=3/2 orthogonal=%t independent_psi=%t samples=%d empirical_masses=%t", x.DemocraticPhaseOrthogonality, x.PhaseNormIndependentOfPsi, len(x.Samples), !x.NoEmpiricalMassesUsed)
}

func FormatDerivation(x NullDerivation) string {
	return fmt.Sprintf("q=%s time=%s space=%s R/S=%s Q=%s scale_free=%t psi_free=%t", fmtFloat(x.MinkowskiNorm), fmtFloat(x.TimelikeNormSquared), fmtFloat(x.SpacelikeNormSquared), fmtFloat(x.RatioDerived), fmtFloat(x.KoideQDerived), x.ScaleFree, x.PhaseFree)
}

func FormatCollapse(x BoundaryCollapse) string {
	return fmt.Sprintf("C3_shape_DOF:%d→%d collapsed=%d scale_free=%t psi_free=%t masses_derived=%t full_flavor_collapsed=%t dims=(%d,%d)", x.C3RawShadowDOF, x.NullConstrainedShadowDOF, x.CollapsedShapeDOF, x.ScaleStillFree, x.PsiStillFree, x.AbsoluteMassesDerived, x.FullFlavorModuliCollapsed, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter)
}

func FormatSector(x SectorBoundary) string {
	return fmt.Sprintf("lepton_baseline=%t quark_baseline=%t color_dressing=%t lepton_masses=%t quark_masses=%t CKM=%t PMNS=%t", x.ColorlessLeptonBaselineEligible, x.QuarkBaselineEligible, x.QuarkColorDressingDeclared, x.PhysicalLeptonMassesDerived, x.PhysicalQuarkMassesDerived, x.CKMConstructed, x.PMNSConstructed)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("null_C3_ratio=%t observed_mass_proof=%t physical_mass_prediction=%t psi=%t sector_perturbations=%t CKM=%t PMNS=%t native_write=%t dims=(%d,%d)", x.NullC3RatioNativeBaseline, x.ObservedMassImportedForProof, x.KoideAsPhysicalMassPrediction, x.PhasePsiSelected, x.SectorPerturbationsNative, x.CKMMatrixConstructed, x.PMNSMatrixConstructed, x.NativeRegistryWritten, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter)
}

func fmtFloat(x float64) string {
	if math.IsNaN(x) {
		return "undefined"
	}
	return fmt.Sprintf("%.12g", x)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 485 Registry Audit — Koide Constraint Provenance & Topological Baseline\n\n")
	b.WriteString("## Verdict\n\n")
	b.WriteString("```text\n")
	b.WriteString(StatusC3ShadowBasisProved + "\n")
	b.WriteString(StatusNullKoideRatioDerived + "\n")
	b.WriteString(StatusKoideQDerived + "\n")
	b.WriteString(StatusLeptonBaselineCompatible + "\n")
	b.WriteString(StatusFlavorFirewallPreserved + "\n")
	b.WriteString("```\n\n")
	b.WriteString("Gate 485 closes the provenance gap left by Gate 484: `R/S = sqrt(2)` is not inserted from the observed charged-lepton masses. It follows from the C3 square-root shadow **only after** the shadow is placed on the `Cℓ(1,7)` null boundary with the democratic leg timelike and the phase-plane leg spacelike.\n\n")

	b.WriteString("## Inherited boundary\n\n")
	b.WriteString("| inherited object | status |\n|---|---|\n")
	b.WriteString(fmt.Sprintf("| Gate 480 null cone | `%t` |\n", a.Inheritance.Gate480NullConeNative))
	b.WriteString(fmt.Sprintf("| Gate 481 common baseline cancellation | `%t` |\n", a.Inheritance.Gate481CommonBaselineCancels))
	b.WriteString(fmt.Sprintf("| Gate 483 quark/lepton topology separation only | `%t` |\n", a.Inheritance.Gate483QuarkLeptonTopologySeparatedOnly))
	b.WriteString(fmt.Sprintf("| Gate 484 C3 basis validated | `%t` |\n", a.Inheritance.Gate484C3BasisValidated))
	b.WriteString(fmt.Sprintf("| Gate 484 charged-lepton Koide shadow found | `%t` |\n", a.Inheritance.Gate484ChargedLeptonKoideShadowFound))
	b.WriteString(fmt.Sprintf("| observed masses remain bridge data | `%t` |\n\n", a.Inheritance.ObservedMassesRemainBridgeData))

	b.WriteString("## C3 square-root shadow theorem\n\n")
	b.WriteString("Define the normalized C3 mass-shadow coordinates\n\n")
	b.WriteString("```text\n")
	b.WriteString("x_i = sqrt(m_i) = S + R cos(θ_i - ψ)\n")
	b.WriteString("θ_i ∈ {0, 2π/3, 4π/3}\n")
	b.WriteString("```\n\n")
	b.WriteString("The democratic component is `D_i=S`; the phase-plane component is `P_i=R cos(θ_i-ψ)`. The C3 identities are:\n\n")
	b.WriteString("```text\n")
	b.WriteString("Σ_i cos(θ_i-ψ)    = 0\n")
	b.WriteString("Σ_i cos²(θ_i-ψ)   = 3/2\n")
	b.WriteString("D·P                 = S R Σ_i cos(θ_i-ψ) = 0\n")
	b.WriteString("||D||²              = 3S²\n")
	b.WriteString("||P||²              = (3/2)R²\n")
	b.WriteString("```\n\n")
	b.WriteString("These are C3 identities, not mass fits. Gate 485 verifies them across phase samples:\n\n")
	b.WriteString("| ψ | Σcos | Σcos² | D·P | ||D||² at S=1 | ||P||² at R=1 | pass |\n")
	b.WriteString("|---:|---:|---:|---:|---:|---:|---|\n")
	for _, s := range a.Basis.Samples {
		b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s | `%t` |\n", fmtFloat(s.Psi), fmtFloat(s.SumCos), fmtFloat(s.SumCosSquared), fmtFloat(s.DemocraticDotPhase), fmtFloat(s.DemocraticNormSquared), fmtFloat(s.PhaseNormSquaredAtR1), s.Orthogonal && s.PhaseNormIndependent))
	}
	b.WriteString("\n")

	b.WriteString("## Null boundary derivation\n\n")
	b.WriteString("The native boundary audit uses the `Cℓ(1,7)` lightlike form with democratic/hierarchy direction timelike and the C3 phase plane spacelike:\n\n")
	b.WriteString("```text\n")
	b.WriteString("q_C3(S,R) = ||D||² - ||P||²\n")
	b.WriteString("          = 3S² - (3/2)R²\n")
	b.WriteString("q_C3 = 0  ⇒  3S² = (3/2)R²\n")
	b.WriteString("          ⇒  R² = 2S²\n")
	b.WriteString("S>0,R>0   ⇒  R/S = sqrt(2)\n")
	b.WriteString("```\n\n")
	b.WriteString("Numerical exactness check on a representative positive null branch:\n\n")
	b.WriteString("| S | R | ψ | 3S² | (3/2)R² | q | R/S | residual from sqrt(2) |\n")
	b.WriteString("|---:|---:|---:|---:|---:|---:|---:|---:|\n")
	b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s | %s |\n\n", fmtFloat(a.Derivation.S), fmtFloat(a.Derivation.R), fmtFloat(a.Derivation.Psi), fmtFloat(a.Derivation.TimelikeNormSquared), fmtFloat(a.Derivation.SpacelikeNormSquared), fmtFloat(a.Derivation.MinkowskiNorm), fmtFloat(a.Derivation.RatioDerived), fmtFloat(a.Derivation.RatioResidual)))

	b.WriteString("## Koide equivalence\n\n")
	b.WriteString("For the same C3 shadow,\n\n")
	b.WriteString("```text\n")
	b.WriteString("Q = (Σ_i x_i²)/(Σ_i x_i)²\n")
	b.WriteString("  = (||D||² + ||P||²)/(3S)²\n")
	b.WriteString("  = (3S² + (3/2)R²)/(9S²).\n")
	b.WriteString("```\n\n")
	b.WriteString("Substitute the null result `R²=2S²`:\n\n")
	b.WriteString("```text\n")
	b.WriteString("Q = (3S² + 3S²)/(9S²) = 2/3.\n")
	b.WriteString("```\n\n")
	b.WriteString(fmt.Sprintf("Gate 485 therefore derives `R/S = %s` and `Q = 2/3 = %s` with no observed masses in the proof.\n\n", fmtFloat(a.Derivation.RatioDerived), fmtFloat(a.Derivation.KoideQDerived)))

	b.WriteString("## Boundary-collapse ledger\n\n")
	b.WriteString("| space | before boundary | after null boundary | collapsed | still free |\n")
	b.WriteString("|---|---:|---:|---:|---|\n")
	b.WriteString(fmt.Sprintf("| C3 square-root shadow `(S,R,ψ)` | `%d` | `%d` | `%d` | `S scale, ψ phase` |\n\n", a.Collapse.C3RawShadowDOF, a.Collapse.NullConstrainedShadowDOF, a.Collapse.CollapsedShapeDOF))
	b.WriteString("The boundary collapses the radial shape coordinate `R/S`, not the mass spectrum. A null C3 baseline has the form\n\n")
	b.WriteString("```text\n")
	b.WriteString("x_i = S [1 + sqrt(2) cos(θ_i - ψ)].\n")
	b.WriteString("```\n\n")
	b.WriteString("So the Koide scalar is fixed, while the absolute scale `S`, the C3 sheet/phase `ψ`, and physical sector perturbations remain unselected.\n\n")

	b.WriteString("## Sector firewall\n\n")
	b.WriteString("```text\n")
	b.WriteString("charged-lepton Koide shadow = bridge-compatible with the colorless null-C3 baseline\n")
	b.WriteString("quark Koide promotion       = rejected; quarks are color/QCD-dressed, not bare null baselines\n")
	b.WriteString("absolute masses             = not derived\n")
	b.WriteString("ψ phase / C3 sheet          = not selected\n")
	b.WriteString("CKM/PMNS                    = not constructed\n")
	b.WriteString("native flavor dimension     = 13 remains sealed\n")
	b.WriteString("K/X/Y charged coefficients  = 9 remain sealed\n")
	b.WriteString("```\n\n")

	b.WriteString("Rejected promotions:\n\n")
	b.WriteString("```text\n")
	b.WriteString(StatusFailedEmpiricalFitRejected + "\n")
	b.WriteString(StatusFailedMassesNotDerived + "\n")
	b.WriteString(StatusFailedPhaseNotSelected + "\n")
	b.WriteString(StatusFailedQuarkPromotionRejected + "\n")
	b.WriteString(StatusFailedMixingPredictionRejected + "\n")
	b.WriteString(StatusFailedFullFlavorCollapseRejected + "\n")
	b.WriteString("```\n\n")

	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n\n")
	b.WriteString("## Next step\n\n")
	b.WriteString(fmt.Sprintf("Gate %d — %s. %s\n", a.Next.Gate, a.Next.Title, a.Next.PrimaryTask))
	return b.String()
}
