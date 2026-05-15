// Package generation2ckmnullmirror implements Gate 486:
// Universal Null-Mirror & CKM Compression Audit.
//
// Gate 485 proved the null-C3 Koide baseline for a bare colorless mass-shadow:
// the C3 democratic leg and phase-plane leg obey q=3S²-(3/2)R², and q=0 forces
// R/S=sqrt(2). Gate 486 asks whether the analogous shared-null-cone picture can
// collapse the physical CKM quotient from four parameters to two coordinates.
//
// The result is intentionally severe: the shared null cone supports a useful
// bridge coordinate socket for relative sector geometry, but it does not yet
// prove a native CKM 4->2 theorem. CKM lives in a rephasing quotient of up/down
// diagonalization frames. A theorem would need native up/down operators and two
// independent rephasing-invariant polynomial constraints. Neither is present.
package generation2ckmnullmirror

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE486-UNIVERSAL-NULL-MIRROR-CKM-COMPRESSION-AUDIT"

	StatusInheritedKoideBaseline                  = "CONDITIONAL_SUPPORT_GATE485_NULL_C3_KOIDE_BASELINE_INHERITED"
	StatusNullMirrorCoordinateChartFound          = "CONDITIONAL_SUPPORT_NULL_MIRROR_BRIDGE_COORDINATE_CHART_FOUND"
	StatusSharedConeDoesNotForceCKMReduction      = "FAILED_ROUTE_SHARED_NULL_CONE_DOES_NOT_FORCE_PHYSICAL_CKM_4_TO_2"
	StatusRephasingInvariantConstraintsAbsent     = "FAILED_ROUTE_REPHASING_INVARIANT_CKM_CONSTRAINTS_NOT_DERIVED"
	StatusNativeUpDownOperatorsAbsent             = "FAILED_ROUTE_NATIVE_UP_DOWN_DIAGONALIZATION_OPERATORS_ABSENT"
	StatusCKMNativeTheoremNotProven               = "FAILED_ROUTE_NATIVE_CKM_4_TO_2_THEOREM_NOT_PROVEN"
	StatusFirewallBlockedCKMRegistryWrite         = "FIREWALL_BLOCKED_CKM_NATIVE_REGISTRY_WRITE"
	StatusEmpiricalCKMFitRejected                 = "FAILED_ROUTE_EMPIRICAL_CKM_WOLFENSTEIN_FIT_REJECTED"
	StatusGate487InvariantPolynomialSearchDefined = "CONDITIONAL_SUPPORT_GATE487_CKM_INVARIANT_POLYNOMIAL_SEARCH_DEFINED"
)

const (
	NativeFlavorDim                  = 13
	KXYCoeffDim                      = 9
	CKMPhysicalParameterDim          = 4
	ProposedNullMirrorDim            = 2
	RequiredConstraintsForFourToTwo  = 2
	DerivedRephasingInvariantEqsNow  = 0
	SectorNullC3ShapeDimAfterGate485 = 2
)

type Inheritance struct {
	Executed                            bool
	Gate480NullConeNative               bool
	Gate481CommonBaselineCancels        bool
	Gate483ColorTopologyGenerationBlind bool
	Gate484TiltReparametrizationBlocked bool
	Gate485KoideProvenanceAccepted      bool
	Gate485NullC3RatioNativeBaseline    bool
	Gate485MassPhaseMixingSealed        bool
	ObservedCKMImported                 bool
	NativeRegistryClean                 bool
	Verdict                             string
	Reason                              string
}

type NullMirrorGeometry struct {
	Executed                      bool
	SharedNullConeAvailable       bool
	EachSectorNullC3ShapeDim      int
	TwoSectorRawNullShadowDim     int
	CommonBaselineQuotientAllowed bool
	RelativeCoordinateChartDim    int
	RelativeTiltCoordinateNamed   string
	RelativePhaseCoordinateNamed  string
	CoordinateChartBridgeOnly     bool
	CKMEigenbasisMismatchDerived  bool
	CKMFourToTwoForcedByCone      bool
	Verdict                       string
	Failure                       string
	Reason                        string
}

type RephasingAudit struct {
	Executed                         bool
	CKMPhysicalQuotientAudited       bool
	CKMRawUnitaryDim                 int
	CKMPhysicalParameterDim          int
	ProposedCompressedDim            int
	RequiredIndependentConstraints   int
	DerivedIndependentConstraints    int
	ModuliRelationsDerived           int
	JarlskogRelationDerived          bool
	UnitarityTriangleRelationDerived bool
	RephasingInvariantConstraintsOK  bool
	CoordinateChartSurvivesRephasing bool
	Verdict                          string
	Reason                           string
}

type NativeOperatorAudit struct {
	Executed                     bool
	NativeUpOperatorDerived      bool
	NativeDownOperatorDerived    bool
	NativeDiagonalizersDerived   bool
	CKMAsUuDaggerUdConstructed   bool
	MassShadowEigenvaluesOnly    bool
	EigenvectorsDeterminedByNull bool
	InvariantPolynomialProduced  bool
	Verdict                      string
	Reason                       string
}

type Firewall struct {
	Executed                      bool
	ObservedCKMImported           bool
	ObservedWolfensteinImported   bool
	ObservedQuarkMassesImported   bool
	CKMMatrixNativePrediction     bool
	CKMFourToTwoNativeWritten     bool
	NullMirrorSocketBridgeWritten bool
	NativeRegistryWritten         bool
	NativeFlavorDimAfter          int
	KXYCoeffDimAfter              int
	Verdict                       string
	Reason                        string
}

type RegistryUpdate struct {
	NativeEntries        []string
	BridgeEntries        []string
	EnvironmentalEntries []string
	FailedRoutes         []string
	OpenTheorems         []string
}

type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance    Inheritance
	Geometry       NullMirrorGeometry
	Rephasing      RephasingAudit
	Operators      NativeOperatorAudit
	Firewall       Firewall
	RegistryUpdate RegistryUpdate
	Next           NextStep
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
	a := Analysis{Inheritance: buildInheritance()}
	a.Geometry = buildNullMirrorGeometry(a.Inheritance)
	a.Rephasing = buildRephasingAudit(a.Geometry)
	a.Operators = buildNativeOperatorAudit(a.Geometry, a.Rephasing)
	a.Firewall = buildFirewall(a)
	a.RegistryUpdate = buildRegistryUpdate(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{
		Executed:                            true,
		Gate480NullConeNative:               true,
		Gate481CommonBaselineCancels:        true,
		Gate483ColorTopologyGenerationBlind: true,
		Gate484TiltReparametrizationBlocked: true,
		Gate485KoideProvenanceAccepted:      true,
		Gate485NullC3RatioNativeBaseline:    true,
		Gate485MassPhaseMixingSealed:        true,
		ObservedCKMImported:                 false,
		NativeRegistryClean:                 true,
		Verdict:                             StatusInheritedKoideBaseline,
		Reason:                              "Gate485 contributes only the null-C3 Koide baseline shape theorem; masses, C3 phase, quark dressing, CKM, PMNS, and physical sector perturbations remain outside native law-space",
	}
}

func buildNullMirrorGeometry(_ Inheritance) NullMirrorGeometry {
	return NullMirrorGeometry{
		Executed:                      true,
		SharedNullConeAvailable:       true,
		EachSectorNullC3ShapeDim:      SectorNullC3ShapeDimAfterGate485,
		TwoSectorRawNullShadowDim:     2 * SectorNullC3ShapeDimAfterGate485,
		CommonBaselineQuotientAllowed: true,
		RelativeCoordinateChartDim:    ProposedNullMirrorDim,
		RelativeTiltCoordinateNamed:   "DeltaAlpha",
		RelativePhaseCoordinateNamed:  "DeltaPhi",
		CoordinateChartBridgeOnly:     true,
		CKMEigenbasisMismatchDerived:  false,
		CKMFourToTwoForcedByCone:      false,
		Verdict:                       StatusNullMirrorCoordinateChartFound,
		Failure:                       StatusSharedConeDoesNotForceCKMReduction,
		Reason:                        "two null-C3 sector shadows can be placed in a two-coordinate relative chart after common-baseline quotienting, but a mass-shadow plane is spectral-shape data, not a native quark eigenbasis frame; the shared cone alone does not construct V_CKM=U_u^† U_d",
	}
}

func buildRephasingAudit(_ NullMirrorGeometry) RephasingAudit {
	return RephasingAudit{
		Executed:                         true,
		CKMPhysicalQuotientAudited:       true,
		CKMRawUnitaryDim:                 9,
		CKMPhysicalParameterDim:          CKMPhysicalParameterDim,
		ProposedCompressedDim:            ProposedNullMirrorDim,
		RequiredIndependentConstraints:   RequiredConstraintsForFourToTwo,
		DerivedIndependentConstraints:    DerivedRephasingInvariantEqsNow,
		ModuliRelationsDerived:           0,
		JarlskogRelationDerived:          false,
		UnitarityTriangleRelationDerived: false,
		RephasingInvariantConstraintsOK:  false,
		CoordinateChartSurvivesRephasing: false,
		Verdict:                          StatusRephasingInvariantConstraintsAbsent,
		Reason:                           "a physical CKM compression from four to two parameters requires two independent relations among rephasing-invariant data such as |V_ij|, unitarity-triangle shape, J, or native commutator traces; Gate486 derives none",
	}
}

func buildNativeOperatorAudit(_ NullMirrorGeometry, _ RephasingAudit) NativeOperatorAudit {
	return NativeOperatorAudit{
		Executed:                     true,
		NativeUpOperatorDerived:      false,
		NativeDownOperatorDerived:    false,
		NativeDiagonalizersDerived:   false,
		CKMAsUuDaggerUdConstructed:   false,
		MassShadowEigenvaluesOnly:    true,
		EigenvectorsDeterminedByNull: false,
		InvariantPolynomialProduced:  false,
		Verdict:                      StatusNativeUpDownOperatorsAbsent,
		Reason:                       "the null-C3 Koide theorem constrains eigenvalue-shadow shape; CKM needs native up/down operators and their diagonalizing frames, which are not produced by the shared null cone",
	}
}

func buildFirewall(a Analysis) Firewall {
	bridgeSocket := a.Geometry.CoordinateChartBridgeOnly && !a.Geometry.CKMFourToTwoForcedByCone
	return Firewall{
		Executed:                      true,
		ObservedCKMImported:           false,
		ObservedWolfensteinImported:   false,
		ObservedQuarkMassesImported:   false,
		CKMMatrixNativePrediction:     false,
		CKMFourToTwoNativeWritten:     false,
		NullMirrorSocketBridgeWritten: bridgeSocket,
		NativeRegistryWritten:         false,
		NativeFlavorDimAfter:          NativeFlavorDim,
		KXYCoeffDimAfter:              KXYCoeffDim,
		Verdict:                       StatusFirewallBlockedCKMRegistryWrite,
		Reason:                        "Gate486 may record a bridge-only null-mirror coordinate socket, but blocks all native CKM, Wolfenstein, quark-mass, and 4->2 registry writes",
	}
}

func buildRegistryUpdate(_ Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"no new native CKM theorem",
			"inherited native null-C3 baseline remains limited to Gate485 Koide-shape provenance",
		},
		BridgeEntries: []string{
			StatusNullMirrorCoordinateChartFound,
			"DeltaAlpha/DeltaPhi may be used only as a bridge-coordinate socket for future invariant tests",
		},
		EnvironmentalEntries: []string{
			"observed CKM matrix entries remain external comparator data",
			"Wolfenstein parameters, quark masses, and CP phase remain forbidden theorem inputs",
		},
		FailedRoutes: []string{
			StatusSharedConeDoesNotForceCKMReduction,
			StatusRephasingInvariantConstraintsAbsent,
			StatusNativeUpDownOperatorsAbsent,
			StatusCKMNativeTheoremNotProven,
			StatusEmpiricalCKMFitRejected,
		},
		OpenTheorems: []string{
			StatusGate487InvariantPolynomialSearchDefined,
			"search for two native rephasing-invariant polynomial constraints, or prove the null mirror is only a coordinate chart",
		},
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        487,
		Title:       "CKM Rephasing-Invariant Polynomial Constraint Search",
		Reason:      "Gate486 found a bridge-only null-mirror socket but no physical CKM quotient theorem.",
		PrimaryTask: "attempt to derive two independent rephasing-invariant polynomial constraints from native up/down finite operators, without importing CKM values, Wolfenstein parameters, quark masses, or CP phases",
	}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate480NullConeNative || !a.Inheritance.Gate485KoideProvenanceAccepted || !a.Inheritance.Gate485MassPhaseMixingSealed || a.Inheritance.ObservedCKMImported || !a.Inheritance.NativeRegistryClean {
		return fmt.Errorf("Gate486 inheritance invalid: %+v", a.Inheritance)
	}
	if !a.Geometry.Executed || !a.Geometry.SharedNullConeAvailable || a.Geometry.RelativeCoordinateChartDim != ProposedNullMirrorDim || !a.Geometry.CoordinateChartBridgeOnly || a.Geometry.CKMEigenbasisMismatchDerived || a.Geometry.CKMFourToTwoForcedByCone {
		return fmt.Errorf("Gate486 geometry invalid: %+v", a.Geometry)
	}
	if !a.Rephasing.Executed || !a.Rephasing.CKMPhysicalQuotientAudited || a.Rephasing.CKMPhysicalParameterDim != CKMPhysicalParameterDim || a.Rephasing.RequiredIndependentConstraints != RequiredConstraintsForFourToTwo || a.Rephasing.DerivedIndependentConstraints != 0 || a.Rephasing.RephasingInvariantConstraintsOK || a.Rephasing.CoordinateChartSurvivesRephasing {
		return fmt.Errorf("Gate486 rephasing audit invalid: %+v", a.Rephasing)
	}
	if !a.Operators.Executed || a.Operators.NativeUpOperatorDerived || a.Operators.NativeDownOperatorDerived || a.Operators.NativeDiagonalizersDerived || a.Operators.CKMAsUuDaggerUdConstructed || !a.Operators.MassShadowEigenvaluesOnly || a.Operators.EigenvectorsDeterminedByNull || a.Operators.InvariantPolynomialProduced {
		return fmt.Errorf("Gate486 native operator audit invalid: %+v", a.Operators)
	}
	if !a.Firewall.Executed || a.Firewall.ObservedCKMImported || a.Firewall.ObservedWolfensteinImported || a.Firewall.ObservedQuarkMassesImported || a.Firewall.CKMMatrixNativePrediction || a.Firewall.CKMFourToTwoNativeWritten || !a.Firewall.NullMirrorSocketBridgeWritten || a.Firewall.NativeRegistryWritten || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("Gate486 firewall invalid: %+v", a.Firewall)
	}
	return nil
}

func truth(a Analysis) string {
	return fmt.Sprintf("Gate486 result: the shared null cone supports a bridge-only %d-coordinate null-mirror chart (%s,%s), but it does not prove CKM 4->2. Physical CKM compression requires %d independent rephasing-invariant constraints and native up/down diagonalization operators; Gate486 derives %d constraints and therefore blocks the native CKM registry write.", a.Geometry.RelativeCoordinateChartDim, a.Geometry.RelativeTiltCoordinateNamed, a.Geometry.RelativePhaseCoordinateNamed, a.Rephasing.RequiredIndependentConstraints, a.Rephasing.DerivedIndependentConstraints)
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("Gate480Null=%t Gate485Koide=%t nullC3Baseline=%t massPhaseMixingSealed=%t observedCKM=%t clean=%t", x.Gate480NullConeNative, x.Gate485KoideProvenanceAccepted, x.Gate485NullC3RatioNativeBaseline, x.Gate485MassPhaseMixingSealed, x.ObservedCKMImported, x.NativeRegistryClean)
}

func FormatGeometry(x NullMirrorGeometry) string {
	return fmt.Sprintf("shared_cone=%t sector_dim=%d two_sector_dim=%d relative_chart_dim=%d bridge_only=%t ckm_forced=%t", x.SharedNullConeAvailable, x.EachSectorNullC3ShapeDim, x.TwoSectorRawNullShadowDim, x.RelativeCoordinateChartDim, x.CoordinateChartBridgeOnly, x.CKMFourToTwoForcedByCone)
}

func FormatRephasing(x RephasingAudit) string {
	return fmt.Sprintf("CKM_phys_dim=%d proposed_dim=%d required_constraints=%d derived_constraints=%d J=%t moduli_relations=%d quotient_ok=%t", x.CKMPhysicalParameterDim, x.ProposedCompressedDim, x.RequiredIndependentConstraints, x.DerivedIndependentConstraints, x.JarlskogRelationDerived, x.ModuliRelationsDerived, x.RephasingInvariantConstraintsOK)
}

func FormatOperators(x NativeOperatorAudit) string {
	return fmt.Sprintf("native_U=%t native_D=%t diagonalizers=%t V=Uu†Ud=%t eigenvalues_only=%t invariant_polynomial=%t", x.NativeUpOperatorDerived, x.NativeDownOperatorDerived, x.NativeDiagonalizersDerived, x.CKMAsUuDaggerUdConstructed, x.MassShadowEigenvaluesOnly, x.InvariantPolynomialProduced)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("observed_CKM=%t wolfenstein=%t quark_masses=%t CKM_native=%t four_to_two_native=%t bridge_socket=%t native_write=%t dims=(%d,%d)", x.ObservedCKMImported, x.ObservedWolfensteinImported, x.ObservedQuarkMassesImported, x.CKMMatrixNativePrediction, x.CKMFourToTwoNativeWritten, x.NullMirrorSocketBridgeWritten, x.NativeRegistryWritten, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 486 Registry Audit — Universal Null-Mirror & CKM Compression Audit\n\n")
	b.WriteString("## Verdict\n\n")
	b.WriteString("```text\n")
	b.WriteString(StatusInheritedKoideBaseline + "\n")
	b.WriteString(StatusNullMirrorCoordinateChartFound + "\n")
	b.WriteString(StatusSharedConeDoesNotForceCKMReduction + "\n")
	b.WriteString(StatusRephasingInvariantConstraintsAbsent + "\n")
	b.WriteString(StatusNativeUpDownOperatorsAbsent + "\n")
	b.WriteString(StatusCKMNativeTheoremNotProven + "\n")
	b.WriteString(StatusFirewallBlockedCKMRegistryWrite + "\n")
	b.WriteString("```\n\n")
	b.WriteString("Gate 486 accepts only a bridge-level null-mirror coordinate chart. It rejects the stronger claim that a shared null cone natively compresses the physical CKM quotient from four parameters to two.\n\n")

	b.WriteString("## Inherited boundary\n\n")
	b.WriteString("Gate 485 proved the null-C3 Koide baseline as a shape theorem:\n\n")
	b.WriteString("```text\n")
	b.WriteString("3S² - (3/2)R² = 0  ⇒  R/S = sqrt(2)  ⇒  Q = 2/3\n")
	b.WriteString("```\n\n")
	b.WriteString("The inherited boundary remains strict: this proves a bare colorless C3 mass-shadow baseline only. It does not derive absolute masses, the C3 phase ψ, quark dressing, CKM, PMNS, or a collapse of the 13 flavor moduli.\n\n")
	b.WriteString("| inherited object | status |\n|---|---|\n")
	b.WriteString(fmt.Sprintf("| Gate 480 null cone | `%t` |\n", a.Inheritance.Gate480NullConeNative))
	b.WriteString(fmt.Sprintf("| Gate 485 Koide provenance | `%t` |\n", a.Inheritance.Gate485KoideProvenanceAccepted))
	b.WriteString(fmt.Sprintf("| observed CKM imported | `%t` |\n", a.Inheritance.ObservedCKMImported))
	b.WriteString(fmt.Sprintf("| native registry clean | `%t` |\n\n", a.Inheritance.NativeRegistryClean))

	b.WriteString("## Topological CKM Audit\n\n")
	b.WriteString("The shared null-cone construction allows the following bridge chart:\n\n")
	b.WriteString("```text\n")
	b.WriteString("sector shadow after Gate485:     (S, ψ)       with R/S fixed by q=0\n")
	b.WriteString("two sectors before quotient:     (S_u, ψ_u; S_d, ψ_d)\n")
	b.WriteString("relative null-mirror socket:     (Δα, Δφ)\n")
	b.WriteString("```\n\n")
	b.WriteString("This is a coordinate socket, not yet a CKM theorem. The null-C3 shadow constrains eigenvalue-shape data. CKM is a mismatch of diagonalizing frames, `V_CKM = U_u^† U_d`. The shared null cone has not produced the native operators whose eigenvectors are `U_u` and `U_d`.\n\n")
	b.WriteString("| audit item | value |\n|---|---:|\n")
	b.WriteString(fmt.Sprintf("| null-C3 shape dimension per sector | `%d` |\n", a.Geometry.EachSectorNullC3ShapeDim))
	b.WriteString(fmt.Sprintf("| two-sector raw null-shadow dimension | `%d` |\n", a.Geometry.TwoSectorRawNullShadowDim))
	b.WriteString(fmt.Sprintf("| proposed relative chart dimension | `%d` |\n", a.Geometry.RelativeCoordinateChartDim))
	b.WriteString(fmt.Sprintf("| CKM eigenbasis mismatch derived | `%t` |\n", a.Geometry.CKMEigenbasisMismatchDerived))
	b.WriteString(fmt.Sprintf("| CKM 4->2 forced by cone | `%t` |\n\n", a.Geometry.CKMFourToTwoForcedByCone))

	b.WriteString("## Rephasing and invariant audit\n\n")
	b.WriteString("Physical CKM data live in a rephasing quotient:\n\n")
	b.WriteString("```text\n")
	b.WriteString("V_CKM ~ D_u V_CKM D_d^†\n")
	b.WriteString("```\n\n")
	b.WriteString("Therefore a genuine 4->2 theorem must produce two independent relations among rephasing-invariant quantities such as `|V_ij|`, unitarity-triangle invariants, the Jarlskog invariant `J`, or native up/down commutator traces. Gate 486 derives none.\n\n")
	b.WriteString("| invariant requirement | count/status |\n|---|---:|\n")
	b.WriteString(fmt.Sprintf("| physical CKM dimension | `%d` |\n", a.Rephasing.CKMPhysicalParameterDim))
	b.WriteString(fmt.Sprintf("| proposed null-mirror dimension | `%d` |\n", a.Rephasing.ProposedCompressedDim))
	b.WriteString(fmt.Sprintf("| required independent invariant constraints | `%d` |\n", a.Rephasing.RequiredIndependentConstraints))
	b.WriteString(fmt.Sprintf("| derived independent invariant constraints | `%d` |\n", a.Rephasing.DerivedIndependentConstraints))
	b.WriteString(fmt.Sprintf("| Jarlskog relation derived | `%t` |\n", a.Rephasing.JarlskogRelationDerived))
	b.WriteString(fmt.Sprintf("| rephasing-invariant compression passed | `%t` |\n\n", a.Rephasing.RephasingInvariantConstraintsOK))

	b.WriteString("## Firewall result\n\n")
	b.WriteString("```text\n")
	b.WriteString(StatusFirewallBlockedCKMRegistryWrite + "\n")
	b.WriteString(StatusEmpiricalCKMFitRejected + "\n")
	b.WriteString("```\n\n")
	b.WriteString("No CKM entries, Wolfenstein parameters, quark masses, or CP phases were imported. The bridge-only null-mirror socket may be recorded for future tests, but CKM 4->2 is not a native registry theorem.\n\n")
	b.WriteString("| firewall item | status |\n|---|---|\n")
	b.WriteString(fmt.Sprintf("| observed CKM imported | `%t` |\n", a.Firewall.ObservedCKMImported))
	b.WriteString(fmt.Sprintf("| Wolfenstein imported | `%t` |\n", a.Firewall.ObservedWolfensteinImported))
	b.WriteString(fmt.Sprintf("| quark masses imported | `%t` |\n", a.Firewall.ObservedQuarkMassesImported))
	b.WriteString(fmt.Sprintf("| CKM native prediction | `%t` |\n", a.Firewall.CKMMatrixNativePrediction))
	b.WriteString(fmt.Sprintf("| CKM 4->2 native write | `%t` |\n", a.Firewall.CKMFourToTwoNativeWritten))
	b.WriteString(fmt.Sprintf("| bridge socket recorded | `%t` |\n", a.Firewall.NullMirrorSocketBridgeWritten))
	b.WriteString(fmt.Sprintf("| native flavor dimension | `%d` |\n", a.Firewall.NativeFlavorDimAfter))
	b.WriteString(fmt.Sprintf("| K/X/Y charged coefficient dimension | `%d` |\n\n", a.Firewall.KXYCoeffDimAfter))

	b.WriteString("## Registry update\n\n")
	writeList := func(title string, xs []string) {
		b.WriteString(title + "\n\n")
		for _, x := range xs {
			b.WriteString("- `" + x + "`\n")
		}
		b.WriteString("\n")
	}
	writeList("Native", a.RegistryUpdate.NativeEntries)
	writeList("Bridge", a.RegistryUpdate.BridgeEntries)
	writeList("Environmental", a.RegistryUpdate.EnvironmentalEntries)
	writeList("Failed route", a.RegistryUpdate.FailedRoutes)
	writeList("Open theorem", a.RegistryUpdate.OpenTheorems)

	b.WriteString("## Next step\n\n")
	b.WriteString(fmt.Sprintf("Gate %d — %s. %s\n\n", a.Next.Gate, a.Next.Title, a.Next.PrimaryTask))

	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}
