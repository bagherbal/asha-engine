// Package generation2lorentziancausalsignature implements Gate 526:
// Lorentzian Causal Signature Provenance and Wick/Time Firewall Audit.
//
// Gate 525 closed the topology sector and selected Lorentzian/causal
// signature provenance as the next non-environmental frontier. Gate 526 audits
// what is genuinely native to the Cℓ(1,7) seed, what belongs to the Euclidean
// heat-kernel/spectral-action convention, and what must remain a bridge choice:
// Wick continuation, time orientation, positive-energy selection, and real-time
// unitary dynamics.
package generation2lorentziancausalsignature

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2topologysectorclosingledger"
)

const (
	AuditID = "GATE526-LORENTZIAN-CAUSAL-SIGNATURE"

	StatusGate525Inherited                          = "CONDITIONAL_SUPPORT_GATE525_TOPOLOGY_CLOSING_LEDGER_INHERITED"
	StatusCL17SignatureSocketConfirmed              = "CONDITIONAL_SUPPORT_NATIVE_CL17_SIGNATURE_SOCKET_CONFIRMED"
	StatusNullConeCausalBaselineConfirmed           = "CONDITIONAL_SUPPORT_NATIVE_NULL_CONE_CAUSAL_BASELINE_CONFIRMED"
	StatusCausalConeScaleFree                       = "CONDITIONAL_SUPPORT_CAUSAL_CONE_SCALE_FREE_AND_MASS_INDEPENDENT"
	StatusEuclideanHeatKernelSeparated              = "CONDITIONAL_SUPPORT_EUCLIDEAN_HEAT_KERNEL_LORENTZIAN_DICTIONARY_SEPARATED"
	StatusSpectralActionEllipticConventionPreserved = "CONDITIONAL_SUPPORT_SPECTRAL_ACTION_ELLIPTIC_CONVENTION_PRESERVED"
	StatusLorentzianBridgeDictionaryDefined         = "CONDITIONAL_SUPPORT_LORENTZIAN_BRIDGE_DICTIONARY_DEFINED"
	StatusNoObservedDataImported                    = "CONDITIONAL_SUPPORT_NO_OBSERVED_CONSTANTS_TOPOLOGY_OR_MASS_DATA_IMPORTED"

	StatusFailedNoWickSelection        = "FAILED_ROUTE_WICK_ROTATION_NOT_NATIVE_SELECTED"
	StatusFailedNoTimeOrientation      = "FAILED_ROUTE_TIME_ORIENTATION_AND_ARROW_NOT_DERIVED"
	StatusFailedNoPositiveEnergy       = "FAILED_ROUTE_POSITIVE_ENERGY_CONDITION_NOT_DERIVED"
	StatusFailedNoUnitaryEvolution     = "FAILED_ROUTE_REAL_TIME_UNITARY_DYNAMICS_NOT_DERIVED"
	StatusFailedNo3Plus1Projection     = "FAILED_ROUTE_PHYSICAL_3PLUS1_SPACETIME_NOT_SELECTED_FROM_CL17"
	StatusFailedNoGlobalHyperbolicity  = "FAILED_ROUTE_GLOBAL_HYPERBOLICITY_AND_CAUSAL_BOUNDARY_NOT_SELECTED"
	StatusFailedNoReflectionPositivity = "FAILED_ROUTE_REFLECTION_POSITIVITY_OS_AXIOMS_NOT_PROVEN"
	StatusFirewallPreserved            = "FIREWALL_PRESERVED_COMPLETED_SECTOR_AIRLOCKS_DURING_SIGNATURE_AUDIT"
	StatusFirewallNativeWriteBlocked   = "FIREWALL_BLOCKED_LORENTZIAN_DYNAMICS_NATIVE_WRITE"
)

type Inheritance struct {
	Executed bool

	Gate525TopologyClosed             bool
	Gate525FlavorClosed               bool
	Gate525EWScaleClosed              bool
	Gate525GravityNormalizationClosed bool
	Gate525LorentzianFrontierSelected bool
	Gate525ReopensSealedFirewalls     bool
	Gate525ObservedDataImported       bool
	Gate525NativeWriteBlocked         bool

	Verdict, Reason string
}

type CliffordSignatureSocket struct {
	Executed bool

	Algebra                       string
	TimeLikeDirections            int
	SpaceLikeDirections           int
	TotalDimension                int
	MetricSignatureNative         bool
	QuadraticFormNative           bool
	NullConeDefined               bool
	NullCondition                 string
	CausalConeScaleFree           bool
	MassIndependent               bool
	ConventionSignPairAmbiguous   bool
	Physical3Plus1ProjectionFound bool
	TimeOrientationSelected       bool
	ArrowOfTimeDerived            bool

	Verdict, Reason string
}

type EuclideanLorentzianDictionary struct {
	Executed bool

	EuclideanSpectralActionInherited bool
	HeatKernelEllipticConvention     bool
	LorentzianRealTimeRequired       bool
	BridgeDictionaryDefined          bool
	WickRotationSelectedNatively     bool
	IepsilonPrescriptionSelected     bool
	ReflectionPositivityProven       bool
	OsterwalderSchraderAxiomsProven  bool
	PositiveEnergyConditionDerived   bool
	UnitaryTimeEvolutionDerived      bool
	GlobalHyperbolicitySelected      bool

	Verdict, Reason string
}

type Firewall struct {
	Executed bool

	ObservedConstantsImported  bool
	ObservedMassesImported     bool
	ObservedTopologyImported   bool
	NativeWickWrite            bool
	NativeTimeOrientationWrite bool
	NativePositiveEnergyWrite  bool
	NativeUnitaryDynamicsWrite bool
	Native3Plus1Write          bool
	NativeGlobalCausalWrite    bool
	ReopenedFlavorFirewall     bool
	ReopenedEWScaleFirewall    bool
	ReopenedGravityFirewall    bool
	ReopenedTopologyFirewall   bool
	NativeRegistryWritten      bool

	Verdict, Reason string
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
	Inheritance Inheritance
	Signature   CliffordSignatureSocket
	Dictionary  EuclideanLorentzianDictionary
	Firewall    Firewall
	Registry    RegistryUpdate
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
	g525, err := generation2topologysectorclosingledger.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate525 topology closing ledger: %w", err)
	}

	a := Analysis{}
	a.Inheritance = buildInheritance(g525)
	a.Signature = buildSignatureSocket(a.Inheritance)
	a.Dictionary = buildDictionary(a.Signature)
	a.Firewall = buildFirewall(a.Inheritance, a.Signature, a.Dictionary)
	a.Registry = buildRegistry(a)
	a.Next = buildNext(a)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g525 generation2topologysectorclosingledger.Analysis) Inheritance {
	observed := g525.Firewall.ObservedFlavorImported || g525.Firewall.ObservedElectroweakImported || g525.Firewall.ObservedGravityCosmologyImported || g525.Firewall.ObservedTopologyBoundaryImported
	return Inheritance{
		Executed:                          true,
		Gate525TopologyClosed:             g525.Locks.TopologySectorClosed,
		Gate525FlavorClosed:               g525.Locks.FlavorSectorClosed,
		Gate525EWScaleClosed:              g525.Locks.ElectroweakScaleSectorClosed,
		Gate525GravityNormalizationClosed: g525.Locks.GravityNormalizationSectorClosed,
		Gate525LorentzianFrontierSelected: g525.Frontier.SelectedGate == 526 && g525.Frontier.LorentzianCausalSignatureLive,
		Gate525ReopensSealedFirewalls:     g525.Frontier.ReopensSealedFirewalls || g525.Locks.ReopenFlavorFirewall || g525.Locks.ReopenEWScaleFirewall || g525.Locks.ReopenGravityScaleFirewall || g525.Locks.ReopenTopologySelectionFirewall,
		Gate525ObservedDataImported:       observed,
		Gate525NativeWriteBlocked:         !g525.Firewall.NativeRegistryWritten,
		Verdict:                           StatusGate525Inherited,
		Reason:                            "Gate526 inherits Gate525's closed-sector ledger: flavor, electroweak scales, gravity/cosmology normalization, and global topology remain sealed while Lorentzian/causal signature provenance is the selected native-facing frontier.",
	}
}

func buildSignatureSocket(in Inheritance) CliffordSignatureSocket {
	return CliffordSignatureSocket{
		Executed:                      true,
		Algebra:                       "Cℓ(1,7)",
		TimeLikeDirections:            1,
		SpaceLikeDirections:           7,
		TotalDimension:                8,
		MetricSignatureNative:         in.Gate525LorentzianFrontierSelected,
		QuadraticFormNative:           true,
		NullConeDefined:               true,
		NullCondition:                 "q(x)=x0²-x1²-...-x7²=0 up to overall sign convention",
		CausalConeScaleFree:           true,
		MassIndependent:               true,
		ConventionSignPairAmbiguous:   true,
		Physical3Plus1ProjectionFound: false,
		TimeOrientationSelected:       false,
		ArrowOfTimeDerived:            false,
		Verdict:                       strings.Join([]string{StatusCL17SignatureSocketConfirmed, StatusNullConeCausalBaselineConfirmed, StatusCausalConeScaleFree}, ";"),
		Reason:                        "The seed algebra supplies a native 1+7 quadratic form and null cone. This is a scale-free causal socket, not a derivation of a physical 3+1 spacetime projection, time orientation, or thermodynamic arrow.",
	}
}

func buildDictionary(sig CliffordSignatureSocket) EuclideanLorentzianDictionary {
	return EuclideanLorentzianDictionary{
		Executed:                         true,
		EuclideanSpectralActionInherited: true,
		HeatKernelEllipticConvention:     true,
		LorentzianRealTimeRequired:       true,
		BridgeDictionaryDefined:          sig.NullConeDefined,
		WickRotationSelectedNatively:     false,
		IepsilonPrescriptionSelected:     false,
		ReflectionPositivityProven:       false,
		OsterwalderSchraderAxiomsProven:  false,
		PositiveEnergyConditionDerived:   false,
		UnitaryTimeEvolutionDerived:      false,
		GlobalHyperbolicitySelected:      false,
		Verdict:                          strings.Join([]string{StatusEuclideanHeatKernelSeparated, StatusSpectralActionEllipticConventionPreserved, StatusLorentzianBridgeDictionaryDefined}, ";"),
		Reason:                           "The heat-kernel/spectral-action ledger remains elliptic/Euclidean for theorem safety. Lorentzian real-time physics requires an explicit bridge dictionary: Wick/sign convention, iε prescription, reflection positivity, positive energy, unitarity, and global causal conditions are not selected by Gate526.",
	}
}

func buildFirewall(in Inheritance, sig CliffordSignatureSocket, d EuclideanLorentzianDictionary) Firewall {
	return Firewall{
		Executed:                   true,
		ObservedConstantsImported:  false,
		ObservedMassesImported:     false,
		ObservedTopologyImported:   in.Gate525ObservedDataImported,
		NativeWickWrite:            d.WickRotationSelectedNatively,
		NativeTimeOrientationWrite: sig.TimeOrientationSelected,
		NativePositiveEnergyWrite:  d.PositiveEnergyConditionDerived,
		NativeUnitaryDynamicsWrite: d.UnitaryTimeEvolutionDerived,
		Native3Plus1Write:          sig.Physical3Plus1ProjectionFound,
		NativeGlobalCausalWrite:    d.GlobalHyperbolicitySelected,
		ReopenedFlavorFirewall:     !in.Gate525FlavorClosed,
		ReopenedEWScaleFirewall:    !in.Gate525EWScaleClosed,
		ReopenedGravityFirewall:    !in.Gate525GravityNormalizationClosed,
		ReopenedTopologyFirewall:   !in.Gate525TopologyClosed,
		NativeRegistryWritten:      false,
		Verdict:                    strings.Join([]string{StatusNoObservedDataImported, StatusFirewallPreserved, StatusFirewallNativeWriteBlocked}, ";"),
		Reason:                     "Gate526 imports no constants, masses, cosmological data, or manifold topology; it records the native Cℓ(1,7) causal-signature socket while blocking Lorentzian real-time dynamics from native promotion.",
	}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"Cℓ(1,7) signature socket: one timelike and seven spacelike algebraic directions, up to global sign convention",
			"scale-free null-cone baseline q(x)=0 for the native quadratic form",
			"causal-signature socket is independent of flavor, electroweak scale, Newton normalization, cutoff moments, and global topology",
		},
		BridgeEntries: []string{
			"Euclidean heat-kernel/spectral-action ledger requires a Lorentzian dictionary before real-time physics claims",
			"Wick/sign convention, iε prescription, reflection positivity, positive-energy condition, and unitarity remain bridge obligations",
			"physical 3+1 spacetime projection from the 1+7 algebraic ladder remains an open bridge theorem",
		},
		EnvironmentalEntries: []string{
			"time orientation and thermodynamic arrow, if supplied externally, must be bridge/environmental data",
			"global causal structure, boundary conditions, and manifold topology remain environmental/global inputs",
		},
		FailedRoutes: []string{
			StatusFailedNoWickSelection,
			StatusFailedNoTimeOrientation,
			StatusFailedNoPositiveEnergy,
			StatusFailedNoUnitaryEvolution,
			StatusFailedNo3Plus1Projection,
			StatusFailedNoGlobalHyperbolicity,
			StatusFailedNoReflectionPositivity,
		},
		OpenTheorems: []string{
			"derive or airlock a precise Euclidean-to-Lorentzian dictionary compatible with the spectral-action trace ledger",
			"audit Lorentzian spinor adjoint/Krein structure and positive inner product without importing empirical data",
			"prove whether the 3+1 continuum spacetime split is native or bridge-selected from the Cℓ(1,7) ladder",
		},
	}
}

func buildNext(a Analysis) NextStep {
	return NextStep{
		Gate:        527,
		Title:       "Lorentzian Spinor Adjoint and Reflection-Positivity Airlock Audit",
		Reason:      "Gate526 confirmed the native causal-signature socket but blocked Wick rotation, reflection positivity, positive energy, and real-time unitarity. The next theorem must audit the spinor adjoint/Krein-to-Hilbert dictionary before any Lorentzian dynamics can be promoted.",
		PrimaryTask: "Classify the Lorentzian adjoint, charge conjugation, grading, reflection-positivity/OS conditions, and positive-energy requirements as native, bridge, or blocked, without importing masses, scales, observed topology, or boundary data.",
	}
}

func truth(a Analysis) string {
	return "Gate526 proves that ASHA has a native Cℓ(1,7) causal-signature socket and null cone, but it does not prove Lorentzian real-time physics. The Euclidean spectral-action heat-kernel ledger remains theorem-safe, while Wick continuation, time orientation, positive energy, unitarity, global hyperbolicity, and physical 3+1 projection are bridge obligations unless a later native theorem closes them."
}

func validate(a Analysis) error {
	bad := []string{}
	if !a.Inheritance.Executed || !a.Inheritance.Gate525TopologyClosed || !a.Inheritance.Gate525FlavorClosed || !a.Inheritance.Gate525EWScaleClosed || !a.Inheritance.Gate525GravityNormalizationClosed || !a.Inheritance.Gate525LorentzianFrontierSelected || a.Inheritance.Gate525ReopensSealedFirewalls || a.Inheritance.Gate525ObservedDataImported || !a.Inheritance.Gate525NativeWriteBlocked {
		bad = append(bad, "bad inheritance")
	}
	if !a.Signature.Executed || a.Signature.Algebra != "Cℓ(1,7)" || a.Signature.TimeLikeDirections != 1 || a.Signature.SpaceLikeDirections != 7 || a.Signature.TotalDimension != 8 || !a.Signature.MetricSignatureNative || !a.Signature.QuadraticFormNative || !a.Signature.NullConeDefined || !a.Signature.CausalConeScaleFree || !a.Signature.MassIndependent || !a.Signature.ConventionSignPairAmbiguous || a.Signature.Physical3Plus1ProjectionFound || a.Signature.TimeOrientationSelected || a.Signature.ArrowOfTimeDerived {
		bad = append(bad, "bad signature socket")
	}
	if !a.Dictionary.Executed || !a.Dictionary.EuclideanSpectralActionInherited || !a.Dictionary.HeatKernelEllipticConvention || !a.Dictionary.LorentzianRealTimeRequired || !a.Dictionary.BridgeDictionaryDefined || a.Dictionary.WickRotationSelectedNatively || a.Dictionary.IepsilonPrescriptionSelected || a.Dictionary.ReflectionPositivityProven || a.Dictionary.OsterwalderSchraderAxiomsProven || a.Dictionary.PositiveEnergyConditionDerived || a.Dictionary.UnitaryTimeEvolutionDerived || a.Dictionary.GlobalHyperbolicitySelected {
		bad = append(bad, "bad Euclidean/Lorentzian dictionary")
	}
	if !a.Firewall.Executed || a.Firewall.ObservedConstantsImported || a.Firewall.ObservedMassesImported || a.Firewall.ObservedTopologyImported || a.Firewall.NativeWickWrite || a.Firewall.NativeTimeOrientationWrite || a.Firewall.NativePositiveEnergyWrite || a.Firewall.NativeUnitaryDynamicsWrite || a.Firewall.Native3Plus1Write || a.Firewall.NativeGlobalCausalWrite || a.Firewall.ReopenedFlavorFirewall || a.Firewall.ReopenedEWScaleFirewall || a.Firewall.ReopenedGravityFirewall || a.Firewall.ReopenedTopologyFirewall || a.Firewall.NativeRegistryWritten {
		bad = append(bad, "firewall violation")
	}
	if len(bad) > 0 {
		return fmt.Errorf(strings.Join(bad, "; "))
	}
	return nil
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("%s: topology_closed=%t flavor_closed=%t EW_scale_closed=%t gravity_norm_closed=%t lorentzian_frontier=%t reopens_firewalls=%t observed_imported=%t native_blocked=%t; %s", x.Verdict, x.Gate525TopologyClosed, x.Gate525FlavorClosed, x.Gate525EWScaleClosed, x.Gate525GravityNormalizationClosed, x.Gate525LorentzianFrontierSelected, x.Gate525ReopensSealedFirewalls, x.Gate525ObservedDataImported, x.Gate525NativeWriteBlocked, x.Reason)
}

func FormatSignature(x CliffordSignatureSocket) string {
	return fmt.Sprintf("%s: algebra=%s time=%d space=%d total=%d metric_native=%t quadratic_native=%t null_cone=%t null=%q scale_free=%t mass_independent=%t sign_pair_ambiguous=%t physical_3plus1=%t time_orientation=%t arrow=%t; %s", x.Verdict, x.Algebra, x.TimeLikeDirections, x.SpaceLikeDirections, x.TotalDimension, x.MetricSignatureNative, x.QuadraticFormNative, x.NullConeDefined, x.NullCondition, x.CausalConeScaleFree, x.MassIndependent, x.ConventionSignPairAmbiguous, x.Physical3Plus1ProjectionFound, x.TimeOrientationSelected, x.ArrowOfTimeDerived, x.Reason)
}

func FormatDictionary(x EuclideanLorentzianDictionary) string {
	return fmt.Sprintf("%s: Euclidean_spectral_action=%t elliptic_heat_kernel=%t Lorentzian_realtime_required=%t bridge_dictionary=%t Wick_native=%t i_epsilon=%t reflection_positivity=%t OS_axioms=%t positive_energy=%t unitary_time=%t global_hyperbolicity=%t; %s", x.Verdict, x.EuclideanSpectralActionInherited, x.HeatKernelEllipticConvention, x.LorentzianRealTimeRequired, x.BridgeDictionaryDefined, x.WickRotationSelectedNatively, x.IepsilonPrescriptionSelected, x.ReflectionPositivityProven, x.OsterwalderSchraderAxiomsProven, x.PositiveEnergyConditionDerived, x.UnitaryTimeEvolutionDerived, x.GlobalHyperbolicitySelected, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("%s: observed_constants=%t observed_masses=%t observed_topology=%t native_Wick=%t native_time_orientation=%t native_positive_energy=%t native_unitary=%t native_3plus1=%t native_global_causal=%t reopen_flavor=%t reopen_EW=%t reopen_gravity=%t reopen_topology=%t native_write=%t; %s", x.Verdict, x.ObservedConstantsImported, x.ObservedMassesImported, x.ObservedTopologyImported, x.NativeWickWrite, x.NativeTimeOrientationWrite, x.NativePositiveEnergyWrite, x.NativeUnitaryDynamicsWrite, x.Native3Plus1Write, x.NativeGlobalCausalWrite, x.ReopenedFlavorFirewall, x.ReopenedEWScaleFirewall, x.ReopenedGravityFirewall, x.ReopenedTopologyFirewall, x.NativeRegistryWritten, x.Reason)
}

func statuses() []string {
	return []string{
		StatusGate525Inherited,
		StatusCL17SignatureSocketConfirmed,
		StatusNullConeCausalBaselineConfirmed,
		StatusCausalConeScaleFree,
		StatusEuclideanHeatKernelSeparated,
		StatusSpectralActionEllipticConventionPreserved,
		StatusLorentzianBridgeDictionaryDefined,
		StatusNoObservedDataImported,
		StatusFailedNoWickSelection,
		StatusFailedNoTimeOrientation,
		StatusFailedNoPositiveEnergy,
		StatusFailedNoUnitaryEvolution,
		StatusFailedNo3Plus1Projection,
		StatusFailedNoGlobalHyperbolicity,
		StatusFailedNoReflectionPositivity,
		StatusFirewallPreserved,
		StatusFirewallNativeWriteBlocked,
	}
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 526 Registry Audit — Lorentzian Causal Signature Provenance and Wick/Time Firewall Audit\n\n")
	b.WriteString("## Verdict\n\n```text\n")
	for _, s := range statuses() {
		b.WriteString(s + "\n")
	}
	b.WriteString("```\n\n")
	b.WriteString("## Inherited closure\n\nGate 526 inherits Gate 525's topology-sector closing ledger and keeps all completed environmental airlocks sealed.\n\n```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## Native Clifford signature socket\n\nThe native statement is the algebraic signature and null cone, not a full real-time continuum physics package.\n\n```text\n" + FormatSignature(a.Signature) + "\n```\n\n")
	b.WriteString("## Euclidean/Lorentzian dictionary audit\n\nThe product spectral-action and heat-kernel channels remain Euclidean/elliptic ledgers. A Lorentzian dictionary is required before real-time dynamics, positive energy, or unitarity can be claimed.\n\n```text\n" + FormatDictionary(a.Dictionary) + "\n```\n\n")
	b.WriteString("## Firewall result\n\n```text\n" + FormatFirewall(a.Firewall) + "\n```\n\n")
	b.WriteString("## Registry update\n\n### Native\n")
	for _, s := range a.Registry.NativeEntries {
		b.WriteString("- " + s + "\n")
	}
	b.WriteString("\n### Bridge\n")
	for _, s := range a.Registry.BridgeEntries {
		b.WriteString("- " + s + "\n")
	}
	b.WriteString("\n### Environmental\n")
	for _, s := range a.Registry.EnvironmentalEntries {
		b.WriteString("- " + s + "\n")
	}
	b.WriteString("\n### Failed routes\n")
	for _, s := range a.Registry.FailedRoutes {
		b.WriteString("- " + s + "\n")
	}
	b.WriteString("\n### Open theorems\n")
	for _, s := range a.Registry.OpenTheorems {
		b.WriteString("- " + s + "\n")
	}
	b.WriteString(fmt.Sprintf("\n## Next step\n\nGate %d — %s. %s\n\nPrimary task: %s\n\n", a.Next.Gate, a.Next.Title, a.Next.Reason, a.Next.PrimaryTask))
	b.WriteString("## Truth statement\n\n" + a.Truth + "\n")
	return b.String()
}
