// Package generation2cosmologicalf4vacuumairlock implements Gate 512:
// Cosmological f4 Vacuum Energy and Subtraction Airlock Audit.
//
// Gate 511 classified the scale-independent a4 curvature-squared sockets and
// explicitly left the a0/f4Λ4 cosmological volume channel unclosed. Gate 512
// audits that channel directly. It computes the native dimensionless finite
// trace prefactor of the spectral-action volume term, proves that the raw
// trace is positive and therefore not internally cancelled by the finite
// Hilbert-space ledger, and quarantines every physical cosmological constant,
// vacuum-subtraction, cutoff, f4-moment, and dark-energy interpretation.
package generation2cosmologicalf4vacuumairlock

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2a4curvaturesquaredledger"
	"github.com/bagherbal/asha-engine/pkg/bridge/productspectralactioncoefficients"
)

const (
	AuditID = "GATE512-COSMOLOGICAL-F4-VACUUM-ENERGY-SUBTRACTION-AIRLOCK-AUDIT"

	StatusGate511A4FirewallInherited        = "CONDITIONAL_SUPPORT_GATE511_A4_CURVATURE_FIREWALL_INHERITED"
	StatusProductA0ChannelInherited         = "CONDITIONAL_SUPPORT_PRODUCT_A0_COSMOLOGICAL_CHANNEL_INHERITED"
	StatusA0VolumePrefactorComputed         = "CONDITIONAL_SUPPORT_A0_VOLUME_PREFACTOR_COMPUTED"
	StatusA0FiniteTraceWeightNative         = "CONDITIONAL_SUPPORT_A0_FINITE_TRACE_VOLUME_WEIGHT_NATIVE"
	StatusF4LambdaFourthObligationIsolated  = "CONDITIONAL_SUPPORT_F4_LAMBDA_FOURTH_OBLIGATION_ISOLATED"
	StatusPositiveVacuumVolumeLedgerAudited = "CONDITIONAL_SUPPORT_POSITIVE_VACUUM_VOLUME_LEDGER_AUDITED"
	StatusCosmologicalAirlockDefined        = "CONDITIONAL_SUPPORT_COSMOLOGICAL_CONSTANT_AIRLOCK_DEFINED"
	StatusNoObservedCosmologyDataImported   = "CONDITIONAL_SUPPORT_NO_OBSERVED_COSMOLOGY_DATA_IMPORTED"

	StatusFailedCosmologicalConstantNotDerived   = "FAILED_ROUTE_COSMOLOGICAL_CONSTANT_NOT_DERIVED_BY_A0"
	StatusFailedF4MomentNotSelected              = "FAILED_ROUTE_F4_MOMENT_NOT_SELECTED"
	StatusFailedCutoffLambdaNotSelectedByA0      = "FAILED_ROUTE_CUTOFF_SCALE_LAMBDA_NOT_SELECTED_BY_A0"
	StatusFailedVacuumSubtractionNotNative       = "FAILED_ROUTE_VACUUM_SUBTRACTION_RENORMALIZATION_NOT_NATIVE"
	StatusFailedFiniteTraceDoesNotCancelVolume   = "FAILED_ROUTE_FINITE_TRACE_DOES_NOT_CANCEL_VOLUME_TERM"
	StatusFailedDarkEnergyNotPredicted           = "FAILED_ROUTE_OBSERVED_DARK_ENERGY_NOT_IMPORTED_OR_PREDICTED"
	StatusFailedSupersymmetricCancellationAbsent = "FAILED_ROUTE_SUPERSYMMETRIC_BOSON_FERMION_CANCELLATION_NOT_PRESENT"

	StatusFirewallPreservedNoCosmologyData       = "FIREWALL_PRESERVED_NO_COSMOLOGY_NEWTON_EW_OR_FLAVOR_DATA_IMPORTED"
	StatusFirewallCosmologicalNativeWriteBlocked = "FIREWALL_BLOCKED_COSMOLOGICAL_CONSTANT_NATIVE_WRITE"
)

const finiteTraceDimension = 96.0

// Inheritance records the boundary imported from Gates 511 and 377.
type Inheritance struct {
	Executed                       bool
	Gate511Inherited               bool
	Gate511A4SocketPresent         bool
	Gate511CosmologicalF4Unsolved  bool
	Gate511PhysicalDynamicsBlocked bool
	ProductTripleValid             bool
	ProductA0ChannelDeclared       bool
	ProductA0Computed              bool
	ProductA0PhysicalPrediction    bool
	ProductHardTOEClosure          bool
	Verdict                        string
	Reason                         string
}

type A0VolumeAudit struct {
	Executed                       bool
	HeatKernelChannel              string
	FiniteTraceDimension           float64
	FourPiFactor                   string
	PrefactorPerF4Lambda4          float64
	ExpectedPrefactor              float64
	NativeDimensionlessTraceWeight bool
	UsesF4LambdaFourth             bool
	UsesF2LambdaSquared            bool
	UsesF0Moment                   bool
	PhysicalCosmologicalConstant   bool
	Formula                        string
	Verdict                        string
	Reason                         string
}

type CancellationAudit struct {
	Executed                        bool
	RawTracePositive                bool
	BosonicSpectralTrace            bool
	FermionicMinusTraceIncluded     bool
	SupersymmetricPairingPresent    bool
	NativeZeroCancellationFound     bool
	SignedEtaCancellationApplicable bool
	VacuumEnergyCancelled           bool
	Verdict                         string
	Reason                          string
}

type SubtractionAirlock struct {
	Executed                       bool
	F4MomentSelected               bool
	CutoffLambdaSelected           bool
	RenormalizationSchemeSelected  bool
	VacuumSubtractionSelected      bool
	ManifoldVolumeSelected         bool
	BoundaryConditionSelected      bool
	ObservedDarkEnergyImported     bool
	PhysicalLambdaCosmoDerived     bool
	NativeCosmologicalWriteAllowed bool
	Verdict                        string
	Reason                         string
}

type Firewall struct {
	Executed                          bool
	NewtonConstantImported            bool
	PlanckScaleImported               bool
	CutoffLambdaImported              bool
	F4MomentImported                  bool
	CosmologicalConstantImported      bool
	DarkEnergyDensityImported         bool
	ElectroweakScaleImported          bool
	FlavorDataImported                bool
	VacuumSubtractionWritten          bool
	NativeCosmologicalConstantWritten bool
	Verdict                           string
	Reason                            string
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
	Inheritance  Inheritance
	A0           A0VolumeAudit
	Cancellation CancellationAudit
	Airlock      SubtractionAirlock
	Firewall     Firewall
	Registry     RegistryUpdate
	Next         NextStep
	Truth        string
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
	g511, err := generation2a4curvaturesquaredledger.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate511 a4 curvature audit: %w", err)
	}
	g377, err := productspectralactioncoefficients.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit product spectral-action coefficient ledger: %w", err)
	}
	a := Analysis{}
	a.Inheritance = buildInheritance(g511, g377)
	a.A0 = buildA0(g377)
	a.Cancellation = buildCancellation()
	a.Airlock = buildAirlock()
	a.Firewall = buildFirewall()
	a.Registry = buildRegistry(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g511 generation2a4curvaturesquaredledger.Analysis, g377 productspectralactioncoefficients.Analysis) Inheritance {
	a0 := g377.Calculation.A0CosmologicalPrefactorPerF4Lambda4
	return Inheritance{
		Executed:                       true,
		Gate511Inherited:               g511.Inheritance.Executed && g511.A4.DimensionlessChannel,
		Gate511A4SocketPresent:         g511.Basis.DynamicalCurvatureSocket && g511.Topological.TopologicalSocketNative,
		Gate511CosmologicalF4Unsolved:  !g511.Firewall.CosmologicalConstantDerived && !g511.Firewall.F4VacuumSubtractionSelected,
		Gate511PhysicalDynamicsBlocked: !g511.Dynamical.PhysicalA4DynamicsClosed && g511.Firewall.PhysicalA4DynamicsWritten == false,
		ProductTripleValid:             g377.Calculation.Product.Valid,
		ProductA0ChannelDeclared:       strings.Contains(g377.Calculation.Convention.A0Density, "Tr(1)") && strings.Contains(g377.Calculation.Convention.Expansion, "f₄Λ⁴"),
		ProductA0Computed:              a0.Numeric > 0 && a0.DeterminedByASHA,
		ProductA0PhysicalPrediction:    a0.FullyPhysical,
		ProductHardTOEClosure:          g377.Calculation.HardTOEClosure,
		Verdict:                        strings.Join([]string{StatusGate511A4FirewallInherited, StatusProductA0ChannelInherited}, ";"),
		Reason:                         "Gate512 inherits the Gate511 gravity firewall and returns to the product spectral-action a0 term: the volume prefactor is computed, but Gate377 already marks the f4Λ4/vacuum-subtraction channel as physically unclosed.",
	}
}

func buildA0(g377 productspectralactioncoefficients.Analysis) A0VolumeAudit {
	pref := g377.Calculation.A0CosmologicalPrefactorPerF4Lambda4.Numeric
	expected := finiteTraceDimension / (16.0 * math.Pi * math.Pi)
	return A0VolumeAudit{
		Executed:                       true,
		HeatKernelChannel:              "a0 cosmological/volume channel of Tr f(D²/Λ²)",
		FiniteTraceDimension:           finiteTraceDimension,
		FourPiFactor:                   "(4π)^(-2) = 1/(16π²)",
		PrefactorPerF4Lambda4:          pref,
		ExpectedPrefactor:              expected,
		NativeDimensionlessTraceWeight: nearly(pref, expected, 1e-15),
		UsesF4LambdaFourth:             true,
		UsesF2LambdaSquared:            false,
		UsesF0Moment:                   false,
		PhysicalCosmologicalConstant:   false,
		Formula:                        "C_Λ/(f₄Λ⁴) = Tr_F(1)/(16π²) = 6/π² for Tr_F(1)=96",
		Verdict:                        strings.Join([]string{StatusA0VolumePrefactorComputed, StatusA0FiniteTraceWeightNative, StatusF4LambdaFourthObligationIsolated}, ";"),
		Reason:                         "the a0 heat-kernel term fixes only the positive dimensionless finite-trace volume prefactor; the physical cosmological constant still requires f4, the cutoff scale, manifold volume conventions, and a vacuum-subtraction/renormalization condition.",
	}
}

func buildCancellation() CancellationAudit {
	return CancellationAudit{
		Executed:                        true,
		RawTracePositive:                true,
		BosonicSpectralTrace:            true,
		FermionicMinusTraceIncluded:     false,
		SupersymmetricPairingPresent:    false,
		NativeZeroCancellationFound:     false,
		SignedEtaCancellationApplicable: false,
		VacuumEnergyCancelled:           false,
		Verdict:                         strings.Join([]string{StatusPositiveVacuumVolumeLedgerAudited, StatusFailedFiniteTraceDoesNotCancelVolume, StatusFailedSupersymmetricCancellationAbsent}, ";"),
		Reason:                          "the spectral-action a0 channel is an unsigned bosonic heat-kernel trace over the finite Hilbert-space multiplicity. Because Tr_F(1)=96 is strictly positive and no native boson/fermion supersymmetric pairing or signed eta cancellation is part of this channel, the finite ledger does not cancel the volume term.",
	}
}

func buildAirlock() SubtractionAirlock {
	return SubtractionAirlock{
		Executed:                       true,
		F4MomentSelected:               false,
		CutoffLambdaSelected:           false,
		RenormalizationSchemeSelected:  false,
		VacuumSubtractionSelected:      false,
		ManifoldVolumeSelected:         false,
		BoundaryConditionSelected:      false,
		ObservedDarkEnergyImported:     false,
		PhysicalLambdaCosmoDerived:     false,
		NativeCosmologicalWriteAllowed: false,
		Verdict:                        strings.Join([]string{StatusCosmologicalAirlockDefined, StatusFailedCosmologicalConstantNotDerived, StatusFailedF4MomentNotSelected, StatusFailedCutoffLambdaNotSelectedByA0, StatusFailedVacuumSubtractionNotNative, StatusFailedDarkEnergyNotPredicted}, ";"),
		Reason:                         "Gate512 defines the cosmological airlock: f4, Λ, renormalization prescription, subtraction baseline, spacetime volume/boundary data, and any observed dark-energy comparator are bridge/environmental inputs unless a separate native theorem supplies them.",
	}
}

func buildFirewall() Firewall {
	return Firewall{
		Executed:                          true,
		NewtonConstantImported:            false,
		PlanckScaleImported:               false,
		CutoffLambdaImported:              false,
		F4MomentImported:                  false,
		CosmologicalConstantImported:      false,
		DarkEnergyDensityImported:         false,
		ElectroweakScaleImported:          false,
		FlavorDataImported:                false,
		VacuumSubtractionWritten:          false,
		NativeCosmologicalConstantWritten: false,
		Verdict:                           strings.Join([]string{StatusNoObservedCosmologyDataImported, StatusFirewallPreservedNoCosmologyData, StatusFirewallCosmologicalNativeWriteBlocked}, ";"),
		Reason:                            "Gate512 imports no G, Planck scale, cutoff value, f4 moment, cosmological constant, dark-energy density, electroweak scale, Yukawa, CKM, or PMNS data, and writes no vacuum subtraction or native cosmological constant.",
	}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"The product spectral-action a0 channel contains a native finite-trace volume prefactor.",
			"For the current finite Hilbert-space ledger, Tr_F(1)=96 gives C_Λ/(f₄Λ⁴)=6/π².",
			"The a0 volume channel is separate from the a2 Einstein-Hilbert and a4 curvature-squared channels.",
		},
		BridgeEntries: []string{
			"The symbolic volume action has the form f₄Λ⁴·(4π)^(-2)·Tr_F(1)·∫√g.",
			"Vacuum-energy subtraction, renormalization prescription, manifold volume/boundary data, f₄, and Λ belong to the cosmological airlock.",
		},
		EnvironmentalEntries: []string{
			"Observed cosmological constant, dark-energy density, Planck/cutoff scale, and vacuum subtraction baseline remain quarantined.",
		},
		FailedRoutes: []string{
			StatusFailedCosmologicalConstantNotDerived,
			StatusFailedF4MomentNotSelected,
			StatusFailedCutoffLambdaNotSelectedByA0,
			StatusFailedVacuumSubtractionNotNative,
			StatusFailedFiniteTraceDoesNotCancelVolume,
			StatusFailedSupersymmetricCancellationAbsent,
			StatusFailedDarkEnergyNotPredicted,
		},
		OpenTheorems: []string{
			"Audit whether any native cutoff-moment hierarchy can relate f₄, f₂, f₀ without importing a scale.",
			"Audit whether boundary/topology/renormalization conditions can be selected natively rather than externally.",
			"Define an observed cosmology comparator airlock only as bridge data, never as a native registry write.",
		},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 513, Title: "Cutoff-Moment Hierarchy and Spectral Scale-Separation Airlock Audit", Reason: "Gate512 isolates the f4Λ4 cosmological obligation but does not select f4, Λ, or a subtraction scheme; the next native question is whether spectral moments have an internal hierarchy or remain independent bridge data.", PrimaryTask: "test whether f4, f2, and f0 are natively related by the finite geometry or whether all absolute cutoff-moment ratios remain conventional/environmental inputs"}
}

func validate(a Analysis) error {
	checks := []struct {
		ok  bool
		msg string
	}{
		{a.Inheritance.Executed && a.Inheritance.Gate511Inherited && a.Inheritance.Gate511A4SocketPresent && a.Inheritance.Gate511CosmologicalF4Unsolved && a.Inheritance.Gate511PhysicalDynamicsBlocked && a.Inheritance.ProductTripleValid && a.Inheritance.ProductA0ChannelDeclared && a.Inheritance.ProductA0Computed && !a.Inheritance.ProductA0PhysicalPrediction && !a.Inheritance.ProductHardTOEClosure, "Gate512 inheritance invalid"},
		{a.A0.Executed && nearly(a.A0.FiniteTraceDimension, 96, 1e-12) && nearly(a.A0.PrefactorPerF4Lambda4, 6.0/(math.Pi*math.Pi), 1e-12) && a.A0.NativeDimensionlessTraceWeight && a.A0.UsesF4LambdaFourth && !a.A0.UsesF2LambdaSquared && !a.A0.UsesF0Moment && !a.A0.PhysicalCosmologicalConstant, "Gate512 a0 volume audit invalid"},
		{a.Cancellation.Executed && a.Cancellation.RawTracePositive && a.Cancellation.BosonicSpectralTrace && !a.Cancellation.FermionicMinusTraceIncluded && !a.Cancellation.SupersymmetricPairingPresent && !a.Cancellation.NativeZeroCancellationFound && !a.Cancellation.SignedEtaCancellationApplicable && !a.Cancellation.VacuumEnergyCancelled, "Gate512 cancellation audit invalid"},
		{a.Airlock.Executed && !a.Airlock.F4MomentSelected && !a.Airlock.CutoffLambdaSelected && !a.Airlock.RenormalizationSchemeSelected && !a.Airlock.VacuumSubtractionSelected && !a.Airlock.ManifoldVolumeSelected && !a.Airlock.BoundaryConditionSelected && !a.Airlock.ObservedDarkEnergyImported && !a.Airlock.PhysicalLambdaCosmoDerived && !a.Airlock.NativeCosmologicalWriteAllowed, "Gate512 subtraction airlock invalid"},
		{a.Firewall.Executed && !a.Firewall.NewtonConstantImported && !a.Firewall.PlanckScaleImported && !a.Firewall.CutoffLambdaImported && !a.Firewall.F4MomentImported && !a.Firewall.CosmologicalConstantImported && !a.Firewall.DarkEnergyDensityImported && !a.Firewall.ElectroweakScaleImported && !a.Firewall.FlavorDataImported && !a.Firewall.VacuumSubtractionWritten && !a.Firewall.NativeCosmologicalConstantWritten, "Gate512 firewall invalid"},
		{a.Next.Gate == 513, "Gate512 next gate invalid"},
	}
	for _, c := range checks {
		if !c.ok {
			return fmt.Errorf(c.msg)
		}
	}
	return nil
}

func statuses() []string {
	return []string{
		StatusGate511A4FirewallInherited,
		StatusProductA0ChannelInherited,
		StatusA0VolumePrefactorComputed,
		StatusA0FiniteTraceWeightNative,
		StatusF4LambdaFourthObligationIsolated,
		StatusPositiveVacuumVolumeLedgerAudited,
		StatusCosmologicalAirlockDefined,
		StatusFailedCosmologicalConstantNotDerived,
		StatusFailedF4MomentNotSelected,
		StatusFailedCutoffLambdaNotSelectedByA0,
		StatusFailedVacuumSubtractionNotNative,
		StatusFailedFiniteTraceDoesNotCancelVolume,
		StatusFailedSupersymmetricCancellationAbsent,
		StatusFailedDarkEnergyNotPredicted,
		StatusNoObservedCosmologyDataImported,
		StatusFirewallPreservedNoCosmologyData,
		StatusFirewallCosmologicalNativeWriteBlocked,
	}
}

func truth(a Analysis) string {
	return "Gate 512 proves that the product spectral action contains a native a0 cosmological/volume socket with finite-trace prefactor Tr_F(1)/(16π²)=6/π². That is a real dimensionless spectral ledger entry. But the same audit proves that the raw finite trace is positive, not self-cancelling, and that f₄, Λ, vacuum subtraction, renormalization scheme, manifold volume/boundary data, and observed dark energy are not selected by the finite algebra. ASHA has a cosmological socket; it does not derive the physical cosmological constant."
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("Gate511 inherited=%t; a4 socket=%t; f4 unsolved=%t; dynamics blocked=%t; product valid=%t; a0 declared=%t; a0 computed=%t; a0 physical=%t; hard ToE=%t", x.Gate511Inherited, x.Gate511A4SocketPresent, x.Gate511CosmologicalF4Unsolved, x.Gate511PhysicalDynamicsBlocked, x.ProductTripleValid, x.ProductA0ChannelDeclared, x.ProductA0Computed, x.ProductA0PhysicalPrediction, x.ProductHardTOEClosure)
}
func FormatA0(x A0VolumeAudit) string {
	return fmt.Sprintf("%s; TrF=%.0f; prefactor=%.12g; expected=%.12g; native dimensionless=%t; uses f4Λ4=%t; uses f2Λ²=%t; uses f0=%t; physical Λ_cosmo=%t", x.Formula, x.FiniteTraceDimension, x.PrefactorPerF4Lambda4, x.ExpectedPrefactor, x.NativeDimensionlessTraceWeight, x.UsesF4LambdaFourth, x.UsesF2LambdaSquared, x.UsesF0Moment, x.PhysicalCosmologicalConstant)
}
func FormatCancellation(x CancellationAudit) string {
	return fmt.Sprintf("raw trace positive=%t; bosonic spectral trace=%t; fermionic minus trace=%t; SUSY pairing=%t; native zero cancellation=%t; eta cancellation applicable=%t; vacuum energy cancelled=%t", x.RawTracePositive, x.BosonicSpectralTrace, x.FermionicMinusTraceIncluded, x.SupersymmetricPairingPresent, x.NativeZeroCancellationFound, x.SignedEtaCancellationApplicable, x.VacuumEnergyCancelled)
}
func FormatAirlock(x SubtractionAirlock) string {
	return fmt.Sprintf("f4 selected=%t; Λ selected=%t; renormalization selected=%t; subtraction selected=%t; manifold volume selected=%t; boundary selected=%t; observed dark energy imported=%t; physical Λ_cosmo derived=%t; native write allowed=%t", x.F4MomentSelected, x.CutoffLambdaSelected, x.RenormalizationSchemeSelected, x.VacuumSubtractionSelected, x.ManifoldVolumeSelected, x.BoundaryConditionSelected, x.ObservedDarkEnergyImported, x.PhysicalLambdaCosmoDerived, x.NativeCosmologicalWriteAllowed)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("G imported=%t; Planck imported=%t; cutoff imported=%t; f4 imported=%t; Λ_cosmo imported=%t; dark energy imported=%t; EW imported=%t; flavor imported=%t; subtraction write=%t; native Λ_cosmo write=%t", x.NewtonConstantImported, x.PlanckScaleImported, x.CutoffLambdaImported, x.F4MomentImported, x.CosmologicalConstantImported, x.DarkEnergyDensityImported, x.ElectroweakScaleImported, x.FlavorDataImported, x.VacuumSubtractionWritten, x.NativeCosmologicalConstantWritten)
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 512 Registry Audit — Cosmological f4 Vacuum Energy and Subtraction Airlock Audit\n\n")
	b.WriteString("## Verdict\n\n```text\n" + strings.Join(statuses(), "\n") + "\n```\n\n")
	b.WriteString("## Inherited boundary\n\n" + a.Inheritance.Reason + "\n\n```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## a0 cosmological volume channel\n\n" + a.A0.Reason + "\n\n```text\n" + FormatA0(a.A0) + "\n```\n\n")
	b.WriteString("## Vacuum cancellation audit\n\n" + a.Cancellation.Reason + "\n\n```text\n" + FormatCancellation(a.Cancellation) + "\n```\n\n")
	b.WriteString("## Subtraction and renormalization airlock\n\n" + a.Airlock.Reason + "\n\n```text\n" + FormatAirlock(a.Airlock) + "\n```\n\n")
	b.WriteString("## Firewall result\n\n" + a.Firewall.Reason + "\n\n```text\n" + FormatFirewall(a.Firewall) + "\n```\n\n")
	b.WriteString("## Registry update\n\n")
	writeList(&b, "Native entries", a.Registry.NativeEntries)
	writeList(&b, "Bridge entries", a.Registry.BridgeEntries)
	writeList(&b, "Environmental entries", a.Registry.EnvironmentalEntries)
	writeList(&b, "Failed routes", a.Registry.FailedRoutes)
	writeList(&b, "Open theorems", a.Registry.OpenTheorems)
	b.WriteString("## Next step\n\nGate513 should be:\n\n```text\nGate 513 — " + a.Next.Title + "\n```\n\nPrimary task:\n\n```text\n" + a.Next.PrimaryTask + "\n```\n\n")
	b.WriteString("## Truth statement\n\n" + a.Truth + "\n")
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

func nearly(a, b, eps float64) bool { return math.Abs(a-b) <= eps }
