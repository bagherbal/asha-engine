// Package generation2yukawatracescalarnormalization implements Gate 501:
// Yukawa-Trace Scalar Normalization Airlock Audit.
//
// Gate 500 read off the symbolic product spectral-action scalar kinetic channel
// K_phi |D_phi|^2 with K_phi=f0 a/pi^2, but exposed a=Tr(Y†Y) as the remaining
// normalization obstruction.  Gate 501 asks whether this trace can be promoted
// natively without reopening the sealed Yukawa-texture/13-moduli branch.
//
// The result is an airlock theorem.  The trace a is a legitimate
// basis/rephasing-invariant scalar norm of the finite Yukawa operator, so CKM
// orientation and unphysical diagonalizer phases do not enter scalar
// normalization.  But the numerical value of a is still the sum of squared
// Yukawa singular values.  Since Gate 489 closed the native Yukawa-amplitude
// selector branch, a remains a sealed bridge/environmental scalar-normalization
// coordinate.  Gate 501 therefore preserves the symbolic action form while
// blocking native scalar normalization, canonical I4 metric, VEV, kappa, and W/Z
// mass promotion.
package generation2yukawatracescalarnormalization

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2productspectralactionkineticprojection"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2yukawaselectorairlock"
)

const (
	AuditID = "GATE501-YUKAWA-TRACE-SCALAR-NORMALIZATION-AIRLOCK-AUDIT"

	StatusGate500Inherited                    = "CONDITIONAL_SUPPORT_GATE500_SYMBOLIC_SCALAR_KINETIC_CHANNEL_INHERITED"
	StatusTraceADefined                       = "CONDITIONAL_SUPPORT_YUKAWA_TRACE_A_LEDGER_DEFINED"
	StatusTraceBasisInvariant                 = "CONDITIONAL_SUPPORT_YUKAWA_TRACE_A_IS_BASIS_REPHASING_INVARIANT"
	StatusCKMOrientationDropsOut              = "CONDITIONAL_SUPPORT_CKM_ORIENTATION_DROPS_OUT_OF_SCALAR_NORMALIZATION"
	StatusSymbolicNormalizationAirlockDefined = "CONDITIONAL_SUPPORT_SYMBOLIC_SCALAR_NORMALIZATION_AIRLOCK_DEFINED"
	StatusScalarNormBridgeAccepted            = "CONDITIONAL_SUPPORT_YUKAWA_TRACE_SCALAR_NORM_BRIDGE_ACCEPTED"
	StatusFirewallPreserved                   = "FIREWALL_PRESERVED_NO_YUKAWA_OR_ELECTROWEAK_NUMERICS_IMPORTED"
	StatusRegistryWriteBlocked                = "FIREWALL_BLOCKED_NATIVE_YUKAWA_TRACE_SCALAR_NORMALIZATION_WRITE"

	StatusFailedTraceValueNotNative         = "FAILED_ROUTE_YUKAWA_TRACE_A_VALUE_NOT_NATIVE_WITHOUT_YUKAWA_AMPLITUDE_SELECTOR"
	StatusFailedTraceNotTopologicalInteger  = "FAILED_ROUTE_YUKAWA_TRACE_A_IS_NOT_A_DISCRETE_TOPOLOGICAL_CHARGE_LEDGER"
	StatusFailedScalarKineticRemainsBridge  = "FAILED_ROUTE_SCALAR_KINETIC_NORMALIZATION_REMAINS_BRIDGE_ENVIRONMENTAL"
	StatusFailedCanonicalI4StillNotSelected = "FAILED_ROUTE_CANONICAL_I4_SCALAR_METRIC_STILL_NOT_NATIVE_SELECTED"
	StatusFailedVEVAndWZStillBlocked        = "FAILED_ROUTE_HIGGS_VEV_AND_WZ_MASS_MATRIX_STILL_BLOCKED"
	StatusFailedKappaStillBridge            = "FAILED_ROUTE_KAPPA_U1_SIX_REMAINS_BRIDGE_CANDIDATE_AFTER_TRACE_AIRLOCK"
	StatusGate502RedirectDefined            = "CONDITIONAL_SUPPORT_GATE502_NORMALIZATION_INDEPENDENT_EW_QUOTIENT_REDIRECT_DEFINED"
)

type Inheritance struct {
	Executed                         bool
	Gate500AuditDefined              bool
	SymbolicScalarKineticReadOff     bool
	CoefficientDependsOnTraceA       bool
	Gate500TraceANativeNumeric       bool
	Gate500NativeScalarKineticClosed bool
	Gate500WZMassBlocked             bool
	NoElectroweakFlavorDataImported  bool
	Verdict                          string
	Reason                           string
}

type TraceDefinition struct {
	Executed                    bool
	TraceSymbol                 string
	Definition                  string
	UsesFiniteYukawaOperator    bool
	PositiveSemidefiniteNorm    bool
	BasisInvariant              bool
	RephasingInvariant          bool
	CKMOrientationIndependent   bool
	DependsOnYukawaSingularVals bool
	DependsOnYukawaAmplitudes   bool
	DiscreteTopologicalCharge   bool
	NativeNumericValueDerived   bool
	Verdict                     string
	Reason                      string
}

type YukawaAirlock struct {
	Executed                         bool
	Gate489AirlockAvailable          bool
	NativeYukawaSelectorBranchClosed bool
	YukawaEntriesEnvironmental       bool
	YukawaNativeSelectorsPassing     int
	RankThreeYukawaMatricesDerived   bool
	CKMOrientationEnvironmental      bool
	TraceARequiresAmplitudeSpectrum  bool
	TraceASealedByFirewall           bool
	Verdict                          string
	Reason                           string
}

type NormalizationDecision struct {
	Executed                       bool
	SymbolicKineticFormAccepted    bool
	TraceABridgeScalarNormAccepted bool
	TraceANativeNumericAccepted    bool
	ScalarKineticCoefficientNative bool
	CanonicalScalarMetricNative    bool
	CanonicalRescalingOnlySymbolic bool
	VacuumOrientationNative        bool
	KappaU1Native                  bool
	WZMassMatrixNative             bool
	Verdict                        string
	Reason                         string
}

type Firewall struct {
	Executed                    bool
	ObservedYukawaImported      bool
	ObservedFermionMassImported bool
	ObservedCKMPMNSImported     bool
	ObservedWMassImported       bool
	ObservedZMassImported       bool
	ObservedHiggsVEVImported    bool
	WeakAngleImported           bool
	GaugeCouplingImported       bool
	NativeTraceAValueWritten    bool
	NativeScalarMetricWritten   bool
	NativeKappaWritten          bool
	NativeWZMassWritten         bool
	Verdict                     string
	Reason                      string
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
	Inheritance   Inheritance
	Trace         TraceDefinition
	YukawaAirlock YukawaAirlock
	Decision      NormalizationDecision
	Firewall      Firewall
	Registry      RegistryUpdate
	Next          NextStep
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
	g500, err := generation2productspectralactionkineticprojection.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate500 product scalar kinetic projection audit: %w", err)
	}
	g489, err := generation2yukawaselectorairlock.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate489 Yukawa selector airlock: %w", err)
	}

	a := Analysis{}
	a.Inheritance = buildInheritance(g500)
	a.Trace = buildTraceDefinition()
	a.YukawaAirlock = buildYukawaAirlock(g489, a.Trace)
	a.Decision = buildDecision(a.Inheritance, a.Trace, a.YukawaAirlock)
	a.Firewall = buildFirewall()
	a.Registry = buildRegistry(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g generation2productspectralactionkineticprojection.Analysis) Inheritance {
	return Inheritance{
		Executed:                         true,
		Gate500AuditDefined:              true,
		SymbolicScalarKineticReadOff:     g.KineticProjection.SymbolicKineticProjectionReadOff && g.KineticProjection.ProductActionContainsDphiSquared,
		CoefficientDependsOnTraceA:       g.KineticProjection.CoefficientDependsOnYukawaTraceA,
		Gate500TraceANativeNumeric:       g.KineticProjection.YukawaTraceANativelyNumeric,
		Gate500NativeScalarKineticClosed: g.Boundary.NativeScalarKineticCoefficientDerived,
		Gate500WZMassBlocked:             !g.Boundary.NativeWZMassMatrixDerived,
		NoElectroweakFlavorDataImported:  g.Firewall.Executed && !g.Firewall.ObservedWMassImported && !g.Firewall.ObservedZMassImported && !g.Firewall.HiggsVEVImported && !g.Firewall.YukawaImported && !g.Firewall.CKMPMNSImported,
		Verdict:                          StatusGate500Inherited,
		Reason:                           "Gate500 supplies the symbolic product scalar kinetic channel but exposes a=Tr(Y†Y) as the coefficient that blocks native scalar normalization.",
	}
}

func buildTraceDefinition() TraceDefinition {
	return TraceDefinition{
		Executed:                    true,
		TraceSymbol:                 "a",
		Definition:                  "a = Tr(Y†Y) = sum_i sigma_i(Y)^2 over the finite Dirac/Yukawa operator",
		UsesFiniteYukawaOperator:    true,
		PositiveSemidefiniteNorm:    true,
		BasisInvariant:              true,
		RephasingInvariant:          true,
		CKMOrientationIndependent:   true,
		DependsOnYukawaSingularVals: true,
		DependsOnYukawaAmplitudes:   true,
		DiscreteTopologicalCharge:   false,
		NativeNumericValueDerived:   false,
		Verdict:                     strings.Join([]string{StatusTraceADefined, StatusTraceBasisInvariant, StatusCKMOrientationDropsOut, StatusFailedTraceValueNotNative}, ";"),
		Reason:                      "The trace a is a legitimate invariant scalar norm of Y, independent of basis phases and CKM orientation, but its value is the squared Yukawa amplitude spectrum; it is not a discrete topological charge or native number.",
	}
}

func buildYukawaAirlock(g generation2yukawaselectorairlock.Analysis, tr TraceDefinition) YukawaAirlock {
	sealed := g.Airlock.NativeYukawaSelectorBranchClosed && g.Airlock.YukawaEntriesEnvironmental && g.Ledger.NativeSelectorsPassing == 0 && tr.DependsOnYukawaAmplitudes
	return YukawaAirlock{
		Executed:                         true,
		Gate489AirlockAvailable:          true,
		NativeYukawaSelectorBranchClosed: g.Airlock.NativeYukawaSelectorBranchClosed,
		YukawaEntriesEnvironmental:       g.Airlock.YukawaEntriesEnvironmental,
		YukawaNativeSelectorsPassing:     g.Ledger.NativeSelectorsPassing,
		RankThreeYukawaMatricesDerived:   g.Variational.RankThreeUpMatrixDerived || g.Variational.RankThreeDownMatrixDerived,
		CKMOrientationEnvironmental:      g.Airlock.CKMOrientationEnvironmental,
		TraceARequiresAmplitudeSpectrum:  tr.DependsOnYukawaSingularVals,
		TraceASealedByFirewall:           sealed,
		Verdict:                          StatusFailedTraceValueNotNative,
		Reason:                           "Gate489 closed the native selector branch for Yukawa amplitudes.  Because a is a spectral norm of those amplitudes, its numeric value cannot be promoted without reopening a sealed flavor theorem lane.",
	}
}

func buildDecision(in Inheritance, tr TraceDefinition, y YukawaAirlock) NormalizationDecision {
	bridge := in.SymbolicScalarKineticReadOff && tr.BasisInvariant && tr.RephasingInvariant && tr.CKMOrientationIndependent
	nativeTrace := tr.NativeNumericValueDerived && !y.TraceASealedByFirewall
	nativeKinetic := bridge && nativeTrace
	return NormalizationDecision{
		Executed:                       true,
		SymbolicKineticFormAccepted:    in.SymbolicScalarKineticReadOff,
		TraceABridgeScalarNormAccepted: bridge,
		TraceANativeNumericAccepted:    nativeTrace,
		ScalarKineticCoefficientNative: nativeKinetic,
		CanonicalScalarMetricNative:    false,
		CanonicalRescalingOnlySymbolic: bridge && !nativeKinetic,
		VacuumOrientationNative:        false,
		KappaU1Native:                  false,
		WZMassMatrixNative:             false,
		Verdict:                        strings.Join([]string{StatusSymbolicNormalizationAirlockDefined, StatusScalarNormBridgeAccepted, StatusFailedScalarKineticRemainsBridge, StatusFailedVEVAndWZStillBlocked}, ";"),
		Reason:                         "The invariant trace a is accepted as a symbolic bridge scalar norm, but not as a native number.  Hence scalar normalization and every W/Z promotion depending on it remain blocked.",
	}
}

func buildFirewall() Firewall {
	return Firewall{
		Executed:                    true,
		ObservedYukawaImported:      false,
		ObservedFermionMassImported: false,
		ObservedCKMPMNSImported:     false,
		ObservedWMassImported:       false,
		ObservedZMassImported:       false,
		ObservedHiggsVEVImported:    false,
		WeakAngleImported:           false,
		GaugeCouplingImported:       false,
		NativeTraceAValueWritten:    false,
		NativeScalarMetricWritten:   false,
		NativeKappaWritten:          false,
		NativeWZMassWritten:         false,
		Verdict:                     StatusFirewallPreserved,
		Reason:                      "No Yukawa, fermion-mass, CKM/PMNS, W/Z, VEV, weak-angle, or gauge-coupling data are imported; no native trace-a, scalar-metric, kappa, or W/Z registry write is made.",
	}
}

func buildRegistry(_ Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"No native numeric value for a=Tr(Y†Y), scalar kinetic coefficient, canonical I4 metric, kappa_U1, Higgs VEV, or W/Z mass matrix is admitted at Gate501.",
		},
		BridgeEntries: []string{
			"a=Tr(Y†Y) is accepted as a basis- and rephasing-invariant symbolic scalar norm of the finite Yukawa operator.",
			"CKM/eigenbasis orientation drops out of scalar normalization because Tr(Y†Y) depends only on singular values.",
			"The scalar kinetic term remains symbolically K_phi=f0 a/pi^2 with canonical rescaling known only in terms of a.",
		},
		EnvironmentalEntries: []string{
			"The numerical Yukawa singular-value spectrum, trace a, scalar normalization, electroweak scale, W/Z/Higgs masses, CKM, PMNS, and continuum couplings remain sealed bridge/environmental data.",
		},
		FailedRoutes: []string{
			StatusFailedTraceValueNotNative,
			StatusFailedTraceNotTopologicalInteger,
			StatusFailedScalarKineticRemainsBridge,
			StatusFailedCanonicalI4StillNotSelected,
			StatusFailedVEVAndWZStillBlocked,
			StatusFailedKappaStillBridge,
		},
		OpenTheorems: []string{
			"Search for scalar-normalization-independent electroweak statements, such as rank, photon nullity, and dimensionless quotient diagnostics, without using a numeric a.",
			"Do not reopen native Yukawa amplitude prediction unless a new finite theorem produces Yukawa singular values without empirical data.",
			"Test whether kappa_U1=6 can be interpreted only as a bridge whitening candidate after scalar trace normalization is sealed.",
		},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 502, Title: "Scalar-Normalization-Independent Electroweak Quotient Audit", Reason: "Gate501 seals the numeric Yukawa trace a, so the only remaining electroweak statements must be invariant under unknown scalar normalization.", PrimaryTask: "audit which electroweak conclusions survive after quotienting out a, f0, VEV, and continuum scale: photon nullity, broken-rank structure, and possible dimensionless Hessian ratios only"}
}

func truth(a Analysis) string {
	if a.Decision.TraceABridgeScalarNormAccepted && !a.Decision.TraceANativeNumericAccepted {
		return "Gate501 proves the scalar-normalization obstruction precisely.  The trace a=Tr(Y†Y) is a valid basis- and rephasing-invariant scalar norm, so CKM orientation does not contaminate the product-action scalar kinetic coefficient.  But its numeric value is the squared Yukawa singular-value spectrum, and Gate489 sealed Yukawa amplitudes as environmental.  Therefore scalar normalization, canonical I4 metric, Higgs VEV, kappa promotion, and physical W/Z masses remain firewalled."
	}
	return "Gate501 did not establish even the symbolic Yukawa-trace scalar normalization airlock."
}

func validate(a Analysis) error {
	checks := []struct {
		ok  bool
		msg string
	}{
		{a.Inheritance.Executed && a.Inheritance.Gate500AuditDefined && a.Inheritance.SymbolicScalarKineticReadOff && a.Inheritance.CoefficientDependsOnTraceA, "Gate500 inheritance missing"},
		{a.Trace.Executed && a.Trace.UsesFiniteYukawaOperator && a.Trace.PositiveSemidefiniteNorm && a.Trace.BasisInvariant && a.Trace.RephasingInvariant && a.Trace.CKMOrientationIndependent, "trace-a invariant definition not established"},
		{a.Trace.DependsOnYukawaSingularVals && a.Trace.DependsOnYukawaAmplitudes && !a.Trace.DiscreteTopologicalCharge && !a.Trace.NativeNumericValueDerived, "trace-a firewall not enforced"},
		{a.YukawaAirlock.Executed && a.YukawaAirlock.NativeYukawaSelectorBranchClosed && a.YukawaAirlock.TraceASealedByFirewall && a.YukawaAirlock.YukawaNativeSelectorsPassing == 0, "Yukawa airlock not inherited"},
		{a.Decision.Executed && a.Decision.SymbolicKineticFormAccepted && a.Decision.TraceABridgeScalarNormAccepted && !a.Decision.TraceANativeNumericAccepted && !a.Decision.ScalarKineticCoefficientNative && !a.Decision.WZMassMatrixNative, "normalization decision over-promoted trace a"},
		{a.Firewall.Executed && !a.Firewall.ObservedYukawaImported && !a.Firewall.ObservedFermionMassImported && !a.Firewall.ObservedCKMPMNSImported && !a.Firewall.ObservedWMassImported && !a.Firewall.ObservedHiggsVEVImported && !a.Firewall.NativeTraceAValueWritten && !a.Firewall.NativeWZMassWritten, "firewall violation"},
		{a.Next.Gate == 502, "Gate502 redirect missing"},
	}
	for _, c := range checks {
		if !c.ok {
			return fmt.Errorf(c.msg)
		}
	}
	return nil
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("gate500=%t symbolic_kinetic=%t coeff_a=%t a_native=%t native_kinetic=%t WZ_blocked=%t no_data=%t verdict=%s reason=%s", x.Gate500AuditDefined, x.SymbolicScalarKineticReadOff, x.CoefficientDependsOnTraceA, x.Gate500TraceANativeNumeric, x.Gate500NativeScalarKineticClosed, x.Gate500WZMassBlocked, x.NoElectroweakFlavorDataImported, x.Verdict, x.Reason)
}

func FormatTraceDefinition(x TraceDefinition) string {
	return fmt.Sprintf("symbol=%s definition=%q finite_Y=%t psd=%t basis=%t rephase=%t CKM_drop=%t singular_values=%t amplitudes=%t topological=%t native_numeric=%t verdict=%s reason=%s", x.TraceSymbol, x.Definition, x.UsesFiniteYukawaOperator, x.PositiveSemidefiniteNorm, x.BasisInvariant, x.RephasingInvariant, x.CKMOrientationIndependent, x.DependsOnYukawaSingularVals, x.DependsOnYukawaAmplitudes, x.DiscreteTopologicalCharge, x.NativeNumericValueDerived, x.Verdict, x.Reason)
}

func FormatYukawaAirlock(x YukawaAirlock) string {
	return fmt.Sprintf("gate489=%t branch_closed=%t env=%t selectors=%d rank3=%t CKM_env=%t trace_requires_spectrum=%t trace_sealed=%t verdict=%s reason=%s", x.Gate489AirlockAvailable, x.NativeYukawaSelectorBranchClosed, x.YukawaEntriesEnvironmental, x.YukawaNativeSelectorsPassing, x.RankThreeYukawaMatricesDerived, x.CKMOrientationEnvironmental, x.TraceARequiresAmplitudeSpectrum, x.TraceASealedByFirewall, x.Verdict, x.Reason)
}

func FormatDecision(x NormalizationDecision) string {
	return fmt.Sprintf("symbolic=%t trace_bridge=%t trace_native=%t kinetic_native=%t I4=%t rescale_symbolic=%t vacuum=%t kappa=%t WZ=%t verdict=%s reason=%s", x.SymbolicKineticFormAccepted, x.TraceABridgeScalarNormAccepted, x.TraceANativeNumericAccepted, x.ScalarKineticCoefficientNative, x.CanonicalScalarMetricNative, x.CanonicalRescalingOnlySymbolic, x.VacuumOrientationNative, x.KappaU1Native, x.WZMassMatrixNative, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("Yukawa=%t fermion_mass=%t CKM_PMNS=%t W=%t Z=%t VEV=%t weak_angle=%t gauge=%t native_a=%t native_metric=%t native_kappa=%t native_WZ=%t verdict=%s reason=%s", x.ObservedYukawaImported, x.ObservedFermionMassImported, x.ObservedCKMPMNSImported, x.ObservedWMassImported, x.ObservedZMassImported, x.ObservedHiggsVEVImported, x.WeakAngleImported, x.GaugeCouplingImported, x.NativeTraceAValueWritten, x.NativeScalarMetricWritten, x.NativeKappaWritten, x.NativeWZMassWritten, x.Verdict, x.Reason)
}

func FormatRegistry(x RegistryUpdate) string {
	return fmt.Sprintf("native=[%s] bridge=[%s] environmental=[%s] failed=[%s] open=[%s]", strings.Join(x.NativeEntries, "; "), strings.Join(x.BridgeEntries, "; "), strings.Join(x.EnvironmentalEntries, "; "), strings.Join(x.FailedRoutes, "; "), strings.Join(x.OpenTheorems, "; "))
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 501 Registry Audit — Yukawa-Trace Scalar Normalization Airlock Audit\n\n")
	b.WriteString("## Verdict\n\n")
	for _, s := range []string{
		StatusSymbolicNormalizationAirlockDefined,
		StatusTraceADefined,
		StatusTraceBasisInvariant,
		StatusCKMOrientationDropsOut,
		StatusScalarNormBridgeAccepted,
		StatusFailedTraceValueNotNative,
		StatusFailedTraceNotTopologicalInteger,
		StatusFailedScalarKineticRemainsBridge,
		StatusFailedCanonicalI4StillNotSelected,
		StatusFailedVEVAndWZStillBlocked,
		StatusFailedKappaStillBridge,
		StatusFirewallPreserved,
		StatusRegistryWriteBlocked,
		StatusGate502RedirectDefined,
	} {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n## Inherited boundary\n\n")
	b.WriteString("Gate500 read off the product-action scalar kinetic channel:\n\n```text\nK_phi = f0 a / pi^2\na = Tr(Y†Y)\nS ⊃ K_phi |D_mu phi|^2\n```\n\n")
	b.WriteString("It did not derive a native scalar kinetic coefficient, canonical scalar metric, Higgs VEV, kappa_U1, or W/Z mass matrix. Gate489 already closed the native Yukawa selector branch.\n\n")
	b.WriteString("## Yukawa-trace invariant audit\n\n")
	b.WriteString("The trace is a well-defined Hilbert-Schmidt norm of the finite Yukawa operator:\n\n```text\na = Tr(Y†Y) = sum_i sigma_i(Y)^2\n```\n\n")
	b.WriteString("This gives a real positive-semidefinite scalar norm. It is invariant under basis changes and quark/lepton rephasings, and it does not depend on CKM/PMNS eigenbasis orientation. That is the positive part of Gate501.\n\n")
	b.WriteString("But the same formula shows the obstruction: the value of `a` is the Yukawa singular-value amplitude spectrum. It is not a discrete charge trace like an anomaly ledger, not a representation dimension, and not a topological integer.\n\n")
	b.WriteString("## Scalar-normalization decision\n\n")
	b.WriteString("Accepted:\n\n```text\nDΦ†DΦ action form: symbolic bridge\na = Tr(Y†Y): invariant symbolic scalar norm\nCKM/PMNS orientation: drops out of scalar normalization\ncanonical rescaling: symbolic in a\n```\n\n")
	b.WriteString("Rejected:\n\n```text\na native numeric: false\nscalar kinetic coefficient native: false\ncanonical I4 scalar metric native: false\nHiggs VEV native: false\nkappa_U1 = 6 native: false\nphysical W/Z mass matrix native: false\n```\n\n")
	b.WriteString("## Firewall result\n\n")
	b.WriteString("No empirical Yukawa entries, fermion masses, CKM/PMNS matrices, W/Z/Higgs values, Higgs VEV, weak angle, or gauge couplings enter this gate. No native write is made for `a`, scalar normalization, kappa, or W/Z masses.\n\n")
	b.WriteString("## Registry update\n\n")
	writeList(&b, "Native", a.Registry.NativeEntries)
	writeList(&b, "Bridge", a.Registry.BridgeEntries)
	writeList(&b, "Environmental", a.Registry.EnvironmentalEntries)
	writeList(&b, "Failed routes", a.Registry.FailedRoutes)
	writeList(&b, "Open theorems", a.Registry.OpenTheorems)
	b.WriteString("## Next step\n\n")
	b.WriteString("Gate502 should be:\n\n```text\nGate 502 — Scalar-Normalization-Independent Electroweak Quotient Audit\n```\n\n")
	b.WriteString("Primary task:\n\n```text\n")
	b.WriteString(a.Next.PrimaryTask)
	b.WriteString("\n```\n\n")
	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}

func writeList(b *strings.Builder, title string, xs []string) {
	b.WriteString("### " + title + "\n\n")
	if len(xs) == 0 {
		b.WriteString("- None.\n\n")
		return
	}
	for _, x := range xs {
		b.WriteString("- " + x + "\n")
	}
	b.WriteString("\n")
}
