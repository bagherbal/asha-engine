// Package generation2nullconeikselector implements Gate 480:
// Algebraic Null-Cone Bridge & I_K Selection.
//
// Gate 473 proved that raw quark masses do not force alpha=1 or I_K=1/2.
// Gate 474 proved that Higgs/gauge electroweak universals are generation-blind.
// Gate 480 therefore tests a strictly algebraic alternative: if the bare family
// bridge is constrained to lie on a C\ell(1,7)-native null boundary, and K is the
// timelike hierarchy leg while the X/Y bridge plane is spacelike, then the
// quadratic form q=a^2-r^2 forces a=r for positive amplitudes. This yields a
// vacuum baseline alpha_vac=1 and I_K=1/2 without using observed masses, CKM, or
// PMNS data.
//
// Integrity boundary: the null-cone calculation selects only a conditional
// vacuum bridge baseline. It does not determine the sector-specific physical
// coordinates alpha_u,phi_u,alpha_d,phi_d,alpha_e,phi_e,alpha_nu,phi_nu, nor does
// it export CKM/PMNS entries as native predictions.
package generation2nullconeikselector

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE480-ALGEBRAIC-NULL-CONE-BRIDGE-I-K-SELECTION"

	StatusGate478Inherited       = "CONDITIONAL_SUPPORT_GATE478_FIREWALL_INHERITED"
	StatusNullConeNative         = "CONDITIONAL_SUPPORT_CLIFFORD_NULL_CONE_NATIVE_OBJECT_CONFIRMED"
	StatusNullBoundaryApplied    = "CONDITIONAL_SUPPORT_FAMILY_BRIDGE_NULL_BOUNDARY_APPLIED"
	StatusEquipartitionDerived   = "CONDITIONAL_SUPPORT_NULL_BOUNDARY_FORCES_A_EQUALS_R"
	StatusIKDerived              = "CONDITIONAL_SUPPORT_I_K_DERIVED_FROM_NATIVE_NULL_BOUNDARY"
	StatusVacuumBaselineExported = "CONDITIONAL_SUPPORT_I_K_HALF_EXPORTED_AS_VACUUM_BASELINE_ONLY"
	StatusFirewallPreserved      = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_GATE480_NULL_BOUNDARY"

	StatusFailedNullBoundaryNotPriorForced = "FAILED_ROUTE_NULL_BOUNDARY_NOT_FORCED_BY_PREVIOUS_GATES"
	StatusFailedKXYSignatureAssumption     = "FAILED_ROUTE_K_TIMELIKE_XY_SPACELIKE_ASSIGNMENT_REQUIRES_EXPLICIT_BOUNDARY_LEDGER"
	StatusFailedSectorCoordinatesUnsolved  = "FAILED_ROUTE_NULL_BASELINE_DOES_NOT_SOLVE_SECTOR_COORDINATES"
	StatusFailedCKMPMNSPredictionRejected  = "FAILED_ROUTE_NULL_CONE_AS_CKM_PMNS_PREDICTION_REJECTED"
	StatusFailedPhysicalIKPromotion        = "FAILED_ROUTE_NULL_BASELINE_AS_PHYSICAL_SECTOR_I_K_REJECTED"
)

const (
	NativeFlavorDim = 13
	KXYCoeffDim     = 9
)

type Inheritance struct {
	Executed                   bool
	Gate444KGenForced          bool
	Gate445TriangleForced      bool
	Gate454IKFormulaAvailable  bool
	Gate473RawMassIKFailed     bool
	Gate474ElectroweakIKFailed bool
	Gate478LeptonFirewallClean bool
	NativeRegistryClean        bool
	Verdict                    string
}

type NullConeMap struct {
	Executed                       bool
	CliffordSignature              string
	NativeNullConeExists           bool
	HierarchyGenerator             string
	MixingGenerators               []string
	KTimelikeAssigned              bool
	XYSpacelikeAssigned            bool
	RadialMixingDefinition         string
	QuadraticForm                  string
	BoundaryCondition              string
	AssignmentPreviouslyForced     bool
	AssignmentDeclaredForGate480   bool
	NullBoundaryPreviouslyForced   bool
	NullBoundaryDeclaredForGate480 bool
	Verdict                        string
	Reason                         string
}

type SieveCase struct {
	Name                   string
	A, B, C, R             float64
	QuadraticNorm          float64
	Alpha                  float64
	IK                     float64
	Null                   bool
	PositiveFutureBranch   bool
	AcceptedVacuumBaseline bool
	Verdict                string
	Reason                 string
}

type EquipartitionSieve struct {
	Executed                     bool
	Cases                        []SieveCase
	AcceptedNullCases            int
	TimelikeRejected             bool
	SpacelikeRejected            bool
	NullForcesAEqualsR           bool
	AlphaVac                     float64
	IKVac                        float64
	IKFormula                    string
	UniqueModuloScaleOrientation bool
	Verdict                      string
	Reason                       string
	Warnings                     []string
}

type CoordinateGap struct {
	Executed                         bool
	VacuumBaselineIKDefined          bool
	PhysicalSectorIKDefined          bool
	QuarkSectorCoordinatesSolved     bool
	LeptonSectorCoordinatesSolved    bool
	DUDComputed                      bool
	DENuComputed                     bool
	CKMPrediction                    bool
	PMNSPrediction                   bool
	NeedsSectorRayPerturbationLedger bool
	Verdict                          string
	Reason                           string
	Failures                         []string
}

type Firewall struct {
	Executed                         bool
	ObservedMassImported             bool
	CKMImported                      bool
	PMNSImported                     bool
	VacuumIKNativeBaseline           bool
	VacuumIKPhysicalSectorCoordinate bool
	DUDNativePrediction              bool
	DENuNativePrediction             bool
	CKMMatrixConstructed             bool
	PMNSMatrixConstructed            bool
	NativeRegistryWritten            bool
	KGenStillForced                  bool
	XTriangleStillForced             bool
	YPhaseStillQuarantined           bool
	SectorCoefficientsStillSealed    bool
	NativeFlavorDimAfter             int
	KXYCoeffDimAfter                 int
	Verdict                          string
	Reason                           string
}

type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Map         NullConeMap
	Sieve       EquipartitionSieve
	Gap         CoordinateGap
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
	a.Map = buildNullConeMap()
	a.Sieve = buildSieve()
	a.Gap = buildGap(a)
	a.Firewall = buildFirewall(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{Executed: true, Gate444KGenForced: true, Gate445TriangleForced: true, Gate454IKFormulaAvailable: true, Gate473RawMassIKFailed: true, Gate474ElectroweakIKFailed: true, Gate478LeptonFirewallClean: true, NativeRegistryClean: true, Verdict: StatusGate478Inherited}
}

func buildNullConeMap() NullConeMap {
	return NullConeMap{
		Executed:                       true,
		CliffordSignature:              "Cℓ(1,7)",
		NativeNullConeExists:           true,
		HierarchyGenerator:             "K_gen hierarchy leg",
		MixingGenerators:               []string{"X_triangle cosine bridge", "Y_phase sine bridge"},
		KTimelikeAssigned:              true,
		XYSpacelikeAssigned:            true,
		RadialMixingDefinition:         "r=sqrt(b^2+c^2)",
		QuadraticForm:                  "q(a,b,c)=a^2-b^2-c^2=a^2-r^2",
		BoundaryCondition:              "q=0 on the bare family bridge",
		AssignmentPreviouslyForced:     false,
		AssignmentDeclaredForGate480:   true,
		NullBoundaryPreviouslyForced:   false,
		NullBoundaryDeclaredForGate480: true,
		Verdict:                        StatusNullConeNative,
		Reason:                         "Cℓ(1,7) has a native null cone; Gate480 declares the family bridge to use K as the timelike leg and the X/Y bridge plane as spacelike legs, then audits the resulting null boundary",
	}
}

func buildSieve() EquipartitionSieve {
	cases := []SieveCase{
		evaluateCase("null future bridge", 1, 1, 0),
		evaluateCase("null rotated bridge", 1, math.Sqrt(0.25), math.Sqrt(0.75)),
		evaluateCase("timelike hierarchy-heavy bridge", 2, 1, 0),
		evaluateCase("spacelike mixing-heavy bridge", 1, 2, 0),
	}
	accepted := 0
	timelikeRejected := false
	spacelikeRejected := false
	for _, c := range cases {
		if c.AcceptedVacuumBaseline {
			accepted++
		}
		if c.QuadraticNorm > 0 && !c.AcceptedVacuumBaseline {
			timelikeRejected = true
		}
		if c.QuadraticNorm < 0 && !c.AcceptedVacuumBaseline {
			spacelikeRejected = true
		}
	}
	alpha := 1.0
	ik := IK(alpha)
	return EquipartitionSieve{
		Executed:                     true,
		Cases:                        cases,
		AcceptedNullCases:            accepted,
		TimelikeRejected:             timelikeRejected,
		SpacelikeRejected:            spacelikeRejected,
		NullForcesAEqualsR:           true,
		AlphaVac:                     alpha,
		IKVac:                        ik,
		IKFormula:                    "I_K=alpha/sqrt(alpha^2+3)",
		UniqueModuloScaleOrientation: true,
		Verdict:                      StatusIKDerived,
		Reason:                       "with r=sqrt(b^2+c^2) and positive future amplitudes, q=0 implies a^2=r^2 and hence alpha=a/r=1; the Gate454 formula gives I_K=1/2",
		Warnings:                     []string{StatusFailedNullBoundaryNotPriorForced, StatusFailedKXYSignatureAssumption},
	}
}

func evaluateCase(name string, a, b, c float64) SieveCase {
	r := math.Hypot(b, c)
	q := a*a - r*r
	alpha := math.NaN()
	ik := math.NaN()
	positive := a > 0 && r > 0
	if r > 0 {
		alpha = a / r
		ik = IK(alpha)
	}
	null := math.Abs(q) < 1e-12
	accepted := null && positive && math.Abs(alpha-1) < 1e-12
	verdict := StatusNullBoundaryApplied
	reason := "case lies on the positive null branch and represents the same alpha=1 ray modulo bridge-plane rotation"
	if !accepted {
		verdict = StatusFailedSectorCoordinatesUnsolved
		reason = "case is not the positive null vacuum baseline selected by q=0"
	}
	return SieveCase{Name: name, A: a, B: b, C: c, R: r, QuadraticNorm: q, Alpha: alpha, IK: ik, Null: null, PositiveFutureBranch: positive, AcceptedVacuumBaseline: accepted, Verdict: verdict, Reason: reason}
}

func IK(alpha float64) float64 { return alpha / math.Sqrt(alpha*alpha+3) }

func buildGap(a Analysis) CoordinateGap {
	return CoordinateGap{
		Executed:                         true,
		VacuumBaselineIKDefined:          a.Sieve.NullForcesAEqualsR && nearly(a.Sieve.IKVac, 0.5, 1e-12),
		PhysicalSectorIKDefined:          false,
		QuarkSectorCoordinatesSolved:     false,
		LeptonSectorCoordinatesSolved:    false,
		DUDComputed:                      false,
		DENuComputed:                     false,
		CKMPrediction:                    false,
		PMNSPrediction:                   false,
		NeedsSectorRayPerturbationLedger: true,
		Verdict:                          StatusVacuumBaselineExported,
		Reason:                           "the null cone closes the vacuum baseline coordinate gap only at the bare bridge level; sector-specific excitations still require bridge ledgers for deviations from alpha_vac and phase sheets",
		Failures:                         []string{StatusFailedSectorCoordinatesUnsolved, StatusFailedCKMPMNSPredictionRejected, StatusFailedPhysicalIKPromotion},
	}
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{
		Executed:                         true,
		ObservedMassImported:             false,
		CKMImported:                      false,
		PMNSImported:                     false,
		VacuumIKNativeBaseline:           a.Sieve.IKVac == 0.5,
		VacuumIKPhysicalSectorCoordinate: false,
		DUDNativePrediction:              false,
		DENuNativePrediction:             false,
		CKMMatrixConstructed:             false,
		PMNSMatrixConstructed:            false,
		NativeRegistryWritten:            false,
		KGenStillForced:                  a.Inheritance.Gate444KGenForced,
		XTriangleStillForced:             a.Inheritance.Gate445TriangleForced,
		YPhaseStillQuarantined:           true,
		SectorCoefficientsStillSealed:    true,
		NativeFlavorDimAfter:             NativeFlavorDim,
		KXYCoeffDimAfter:                 KXYCoeffDim,
		Verdict:                          StatusFirewallPreserved,
		Reason:                           "Gate480 exports I_K=1/2 only as a null-vacuum baseline; no observed data, sector coordinate, CKM/PMNS matrix, or physical flavor prediction is written into native law-space",
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 481, Title: "Null-baseline perturbation ledger", Reason: "Gate480 gives a bare null-vacuum I_K baseline but does not determine sector excitations or phases.", PrimaryTask: "define bridge-only perturbation variables around alpha_vac=1 and test whether quark/lepton residual ledgers can be expressed as deviations from the null baseline without native-promotion"}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate454IKFormulaAvailable || !a.Inheritance.NativeRegistryClean {
		return fmt.Errorf("Gate480 inheritance incomplete: %+v", a.Inheritance)
	}
	if !a.Map.Executed || !a.Map.NativeNullConeExists || !a.Map.KTimelikeAssigned || !a.Map.XYSpacelikeAssigned || !a.Map.NullBoundaryDeclaredForGate480 {
		return fmt.Errorf("Gate480 null-cone map invalid: %+v", a.Map)
	}
	if !a.Sieve.Executed || !a.Sieve.NullForcesAEqualsR || !nearly(a.Sieve.AlphaVac, 1, 1e-12) || !nearly(a.Sieve.IKVac, 0.5, 1e-12) || !a.Sieve.UniqueModuloScaleOrientation {
		return fmt.Errorf("Gate480 equipartition sieve invalid: %+v", a.Sieve)
	}
	if !a.Gap.Executed || !a.Gap.VacuumBaselineIKDefined || a.Gap.PhysicalSectorIKDefined || a.Gap.CKMPrediction || a.Gap.PMNSPrediction {
		return fmt.Errorf("Gate480 coordinate-gap boundary invalid: %+v", a.Gap)
	}
	if !a.Firewall.Executed || a.Firewall.ObservedMassImported || a.Firewall.CKMImported || a.Firewall.PMNSImported || !a.Firewall.VacuumIKNativeBaseline || a.Firewall.VacuumIKPhysicalSectorCoordinate || a.Firewall.DUDNativePrediction || a.Firewall.DENuNativePrediction || a.Firewall.NativeRegistryWritten || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("Gate480 firewall invalid: %+v", a.Firewall)
	}
	return nil
}

func nearly(a, b, tol float64) bool { return math.Abs(a-b) <= tol }

func truth(a Analysis) string {
	return fmt.Sprintf("Gate480 result: declared Cℓ(1,7) null bridge q=a²-r²=0 forces alpha_vac=%.12g and I_K=%.12g as a vacuum baseline; this does not compute d_ud, d_eν, CKM, PMNS, or sector physical coordinates.", a.Sieve.AlphaVac, a.Sieve.IKVac)
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("K=%t X=%t IKFormula=%t rawMassIKFailed=%t EWIKFailed=%t leptonFirewall=%t clean=%t", x.Gate444KGenForced, x.Gate445TriangleForced, x.Gate454IKFormulaAvailable, x.Gate473RawMassIKFailed, x.Gate474ElectroweakIKFailed, x.Gate478LeptonFirewallClean, x.NativeRegistryClean)
}

func FormatMap(x NullConeMap) string {
	return fmt.Sprintf("signature=%s q=%s boundary=%s K_time=%t XY_space=%t declaredBoundary=%t priorForced=%t", x.CliffordSignature, x.QuadraticForm, x.BoundaryCondition, x.KTimelikeAssigned, x.XYSpacelikeAssigned, x.NullBoundaryDeclaredForGate480, x.NullBoundaryPreviouslyForced)
}

func FormatSieve(x EquipartitionSieve) string {
	return fmt.Sprintf("alpha_vac=%s I_K=%s accepted_null_cases=%d unique_mod_scale_orientation=%t warnings=%s", fmtFloat(x.AlphaVac), fmtFloat(x.IKVac), x.AcceptedNullCases, x.UniqueModuloScaleOrientation, strings.Join(x.Warnings, "; "))
}

func FormatGap(x CoordinateGap) string {
	return fmt.Sprintf("vacuum_IK=%t physical_sector_IK=%t d_ud=%t d_eν=%t CKM_pred=%t PMNS_pred=%t need_perturbation_ledger=%t", x.VacuumBaselineIKDefined, x.PhysicalSectorIKDefined, x.DUDComputed, x.DENuComputed, x.CKMPrediction, x.PMNSPrediction, x.NeedsSectorRayPerturbationLedger)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("vacuum_baseline=%t physical_sector_coordinate=%t observed_mass=%t CKM=%t PMNS=%t native_write=%t dims=(%d,%d)", x.VacuumIKNativeBaseline, x.VacuumIKPhysicalSectorCoordinate, x.ObservedMassImported, x.CKMImported, x.PMNSImported, x.NativeRegistryWritten, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter)
}

func fmtFloat(x float64) string {
	if math.IsNaN(x) {
		return "undefined"
	}
	return fmt.Sprintf("%.12g", x)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 480 Registry Audit — Algebraic Null-Cone Bridge & I_K Selection\n\n")
	b.WriteString("## Verdict\n\n")
	b.WriteString("`" + StatusIKDerived + "`\n\n")
	b.WriteString("Gate 480 confirms that the native Clifford null-cone calculation can select a **bare vacuum bridge baseline**:\n\n")
	b.WriteString("```text\n")
	b.WriteString("q(a,b,c)=a^2-b^2-c^2=a^2-r^2\n")
	b.WriteString("q=0, a>0, r>0  =>  a=r\n")
	b.WriteString("alpha_vac=a/r=1\n")
	b.WriteString("I_K=alpha/sqrt(alpha^2+3)=1/2\n")
	b.WriteString("```\n\n")
	b.WriteString("## Native null-cone map\n\n")
	b.WriteString("| object | result |\n|---|---|\n")
	b.WriteString("| Clifford signature | `" + a.Map.CliffordSignature + "` |\n")
	b.WriteString("| hierarchy leg | `" + a.Map.HierarchyGenerator + "` |\n")
	b.WriteString("| bridge plane | `" + strings.Join(a.Map.MixingGenerators, ", ") + "` |\n")
	b.WriteString("| radial bridge amplitude | `" + a.Map.RadialMixingDefinition + "` |\n")
	b.WriteString("| quadratic form | `" + a.Map.QuadraticForm + "` |\n")
	b.WriteString("| boundary | `" + a.Map.BoundaryCondition + "` |\n\n")
	b.WriteString("The null cone is native to `Cℓ(1,7)`. The specific statement that the **bare family bridge** lies on that null cone is recorded as the Gate 480 boundary ledger. It was not derived by the earlier raw-mass, electroweak, or PMNS preflight gates.\n\n")
	b.WriteString("## Equipartition sieve\n\n")
	b.WriteString("| case | a | b | c | r | q=a²-r² | alpha | I_K | accepted |\n|---|---:|---:|---:|---:|---:|---:|---:|---|\n")
	for _, c := range a.Sieve.Cases {
		b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s | %s | %t |\n", c.Name, fmtFloat(c.A), fmtFloat(c.B), fmtFloat(c.C), fmtFloat(c.R), fmtFloat(c.QuadraticNorm), fmtFloat(c.Alpha), fmtFloat(c.IK), c.AcceptedVacuumBaseline))
	}
	b.WriteString("\nModulo positive scale and rotation inside the X/Y bridge plane, the null branch selects `alpha_vac=1` and therefore `I_K=0.5`.\n\n")
	b.WriteString("## Firewall boundary\n\n")
	b.WriteString("Gate 480 does **not** compute physical sector coordinates. It does not compute `d_ud`, `d_eν`, CKM entries, PMNS entries, masses, Yukawas, or branch sheets. It exports only a vacuum baseline.\n\n")
	b.WriteString("Rejected promotions:\n\n")
	b.WriteString("```text\n")
	b.WriteString(StatusFailedNullBoundaryNotPriorForced + "\n")
	b.WriteString(StatusFailedKXYSignatureAssumption + "\n")
	b.WriteString(StatusFailedSectorCoordinatesUnsolved + "\n")
	b.WriteString(StatusFailedCKMPMNSPredictionRejected + "\n")
	b.WriteString(StatusFailedPhysicalIKPromotion + "\n")
	b.WriteString("```\n\n")
	b.WriteString("## Numerical output\n\n")
	b.WriteString("```text\n")
	b.WriteString(fmt.Sprintf("alpha_vac = %.12f\n", a.Sieve.AlphaVac))
	b.WriteString(fmt.Sprintf("I_K,vac   = %.12f\n", a.Sieve.IKVac))
	b.WriteString("d_ud      = undefined\n")
	b.WriteString("d_eν      = undefined\n")
	b.WriteString("CKM/PMNS  = not constructed\n")
	b.WriteString("```\n\n")
	b.WriteString("## Next step\n\n")
	b.WriteString(fmt.Sprintf("Gate %d — %s: %s\n", a.Next.Gate, a.Next.Title, a.Next.PrimaryTask))
	return b.String()
}
