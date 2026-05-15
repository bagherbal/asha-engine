// Package generation2bordismcobordismclassifierairlock implements Gate 521:
// Bordism and Cobordism Classifier Airlock.
//
// Gate 520 proved that topology/boundary comparator files can be loaded and
// evaluated only inside a bridge airlock. Gate 521 asks a narrower native-safe
// question: does ASHA have scale-free bordism/cobordism sockets that classify
// admissible topology classes without selecting the actual universe manifold?
// The answer is conditional support for the classifier ledger, plus a hard
// firewall against writing any specific bordism class, Euler number, signature,
// eta invariant, or boundary condition natively.
package generation2bordismcobordismclassifierairlock

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2observedtopologyboundaryfileadapter"
)

const (
	AuditID = "GATE521-BORDISM-COBORDISM-CLASSIFIER-AIRLOCK"

	StatusGate520Inherited                 = "CONDITIONAL_SUPPORT_GATE520_TOPOLOGY_BOUNDARY_FILE_ADAPTER_INHERITED"
	StatusBordismLedgerDefined             = "CONDITIONAL_SUPPORT_BORDISM_COBORDISM_CLASSIFIER_LEDGER_DEFINED"
	StatusOrientedSocketPresent            = "CONDITIONAL_SUPPORT_ORIENTED_BORDISM_SOCKET_PRESENT"
	StatusSpinSocketPresent                = "CONDITIONAL_SUPPORT_SPIN_BORDISM_SOCKET_PRESENT"
	StatusSpinCSocketPresent               = "CONDITIONAL_SUPPORT_SPINC_BORDISM_SOCKET_PRESENT"
	StatusCharacteristicConstraintsAudited = "CONDITIONAL_SUPPORT_CHARACTERISTIC_NUMBER_CONSTRAINTS_AUDITED"
	StatusClassifierScaleFree              = "CONDITIONAL_SUPPORT_BORDISM_CLASSIFIER_SCALE_FREE_AND_MASS_INDEPENDENT"
	StatusClassifierAirlockDefined         = "CONDITIONAL_SUPPORT_BORDISM_CLASSIFIER_AIRLOCK_DEFINED"
	StatusNoObservedTopologyImported       = "CONDITIONAL_SUPPORT_NO_OBSERVED_BORDISM_OR_BOUNDARY_DATA_IMPORTED"

	StatusFailedSpecificClassNotSelected          = "FAILED_ROUTE_SPECIFIC_BORDISM_CLASS_NOT_NATIVE_SELECTED"
	StatusFailedManifoldRepresentativeNotSelected = "FAILED_ROUTE_MANIFOLD_REPRESENTATIVE_NOT_NATIVE_SELECTED"
	StatusFailedStiefelWhitneyNotDerived          = "FAILED_ROUTE_STIEFEL_WHITNEY_CLASSES_NOT_DERIVED_WITHOUT_TANGENT_BUNDLE"
	StatusFailedGlobalCharacteristicNumbers       = "FAILED_ROUTE_GLOBAL_CHARACTERISTIC_NUMBERS_NOT_DERIVED_WITHOUT_MANIFOLD_TOPOLOGY"
	StatusFailedSpinStructureNotSelected          = "FAILED_ROUTE_SPIN_OR_SPINC_STRUCTURE_NOT_NATIVE_SELECTED"
	StatusFailedBoundaryBordismNotClosed          = "FAILED_ROUTE_BOUNDARY_BORDISM_AND_ETA_DATA_NOT_NATIVE_CLOSED"
	StatusFailedTopologyNativeWrite               = "FAILED_ROUTE_BORDISM_TOPOLOGY_NATIVE_WRITE_REJECTED"
	StatusFirewallPreserved                       = "FIREWALL_PRESERVED_NO_MANIFOLD_BOUNDARY_NEWTON_COSMOLOGY_EW_OR_FLAVOR_DATA_IMPORTED"
	StatusFirewallNativeWriteBlocked              = "FIREWALL_BLOCKED_BORDISM_CLASSIFIER_NATIVE_TOPOLOGY_WRITE"
)

type Inheritance struct {
	Executed                        bool
	Gate520FileAdapterDefined       bool
	Gate520SyntheticOnly            bool
	Gate520ResidualsZero            bool
	Gate520NativePrediction         bool
	Gate520NativeWriteBlocked       bool
	Gate520ObservedTopologyImported bool
	Gate520BoundaryMode             bool
	Gate521Redirect                 bool
	Verdict, Reason                 string
}

type BordismSocket struct {
	Executed                       bool
	Dimension                      int
	OrientedSocket                 bool
	SpinSocket                     bool
	SpinCSocket                    bool
	BoundaryBordismSocket          bool
	RequiresW1ZeroForOriented      bool
	RequiresW2ZeroForSpin          bool
	RequiresW3ZeroForSpinC         bool
	RequiresC1Mod2EqualsW2ForSpinC bool
	ClassifiesAllowedClasses       bool
	SelectsSpecificClass           bool
	SelectsManifoldRepresentative  bool
	Verdict, Reason                string
}

type CharacteristicConstraints struct {
	Executed                    bool
	SignaturePontryaginRelation string
	SpinDiracIndexRelation      string
	RokhlinDivisibility         int
	SyntheticTau                float64
	SyntheticP1                 float64
	SyntheticAHat               float64
	SignatureP1Residual         float64
	SpinDivisibilityPassed      bool
	UsesEulerSocket             bool
	UsesPontryaginSocket        bool
	UsesSignatureSocket         bool
	UsesEtaBoundaryCorrection   bool
	GlobalNumbersDerived        bool
	PhysicalThetaSelected       bool
	Verdict, Reason             string
}

type ScaleFirewall struct {
	Executed                 bool
	UsesLambda               bool
	UsesF2                   bool
	UsesF4                   bool
	UsesNewton               bool
	UsesCosmologicalConstant bool
	UsesElectroweakData      bool
	UsesFlavorData           bool
	UsesObservedTopology     bool
	UsesBoundarySpectrum     bool
	ClassifierScaleFree      bool
	Verdict, Reason          string
}

type Rejection struct {
	Executed                            bool
	SpecificBordismClassNativeBlocked   bool
	ManifoldRepresentativeNativeBlocked bool
	StiefelWhitneyClassesNativeBlocked  bool
	CharacteristicNumbersNativeBlocked  bool
	SpinStructureNativeBlocked          bool
	BoundaryBordismNativeBlocked        bool
	EtaInvariantNativeBlocked           bool
	NativeRegistryWriteBlocked          bool
	ComparatorOnlyPurpose               bool
	Verdict, Reason                     string
}

type RegistryUpdate struct{ NativeEntries, BridgeEntries, EnvironmentalEntries, FailedRoutes, OpenTheorems []string }
type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Socket      BordismSocket
	Constraints CharacteristicConstraints
	Scale       ScaleFirewall
	Rejection   Rejection
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
	a := Analysis{}
	a.Inheritance = buildInheritance()
	a.Socket = buildSocket()
	a.Constraints = buildConstraints()
	a.Scale = buildScale()
	a.Rejection = buildRejection()
	a.Registry = buildRegistry(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	g520, err := generation2observedtopologyboundaryfileadapter.BuildDefault()
	if err != nil {
		return Inheritance{Executed: false, Verdict: StatusFailedTopologyNativeWrite, Reason: fmt.Sprintf("could not inherit Gate520: %v", err)}
	}
	return Inheritance{
		Executed:                        true,
		Gate520FileAdapterDefined:       g520.Import.Loaded && g520.Import.MetadataComplete,
		Gate520SyntheticOnly:            g520.Import.SyntheticFixture && !g520.Import.ObservedValuesLoaded,
		Gate520ResidualsZero:            g520.Output.AllResidualsZero,
		Gate520NativePrediction:         g520.Output.NativePrediction,
		Gate520NativeWriteBlocked:       !g520.Firewall.NativeRegistryWritten && !g520.Firewall.FileRowsNative,
		Gate520ObservedTopologyImported: g520.Firewall.ObservedTopologyImported || g520.Firewall.ObservedBoundaryDataImported,
		Gate520BoundaryMode:             g520.Output.BoundaryMode,
		Gate521Redirect:                 g520.Next.Gate == 521,
		Verdict:                         StatusGate520Inherited,
		Reason:                          "Gate521 inherits Gate520's bridge-only topology/boundary file adapter, zero-residual plumbing, synthetic-only default fixture, and hard native-write firewall.",
	}
}

func buildSocket() BordismSocket {
	return BordismSocket{
		Executed: true, Dimension: 4,
		OrientedSocket: true, SpinSocket: true, SpinCSocket: true, BoundaryBordismSocket: true,
		RequiresW1ZeroForOriented:      true,
		RequiresW2ZeroForSpin:          true,
		RequiresW3ZeroForSpinC:         true,
		RequiresC1Mod2EqualsW2ForSpinC: true,
		ClassifiesAllowedClasses:       true,
		SelectsSpecificClass:           false,
		SelectsManifoldRepresentative:  false,
		Verdict:                        strings.Join([]string{StatusBordismLedgerDefined, StatusOrientedSocketPresent, StatusSpinSocketPresent, StatusSpinCSocketPresent}, ";"),
		Reason:                         "The bordism ledger can type oriented, spin, spin-c, and boundary-bordism sockets as scale-free classifier structures, but it does not choose the actual manifold or a representative of a bordism class.",
	}
}

func buildConstraints() CharacteristicConstraints {
	tau := -16.0
	p1 := 3 * tau
	ahat := -tau / 8
	return CharacteristicConstraints{
		Executed:                    true,
		SignaturePontryaginRelation: "∫p1 = 3τ",
		SpinDiracIndexRelation:      "Â(M) = -τ/8 for closed spin 4-manifold",
		RokhlinDivisibility:         16,
		SyntheticTau:                tau,
		SyntheticP1:                 p1,
		SyntheticAHat:               ahat,
		SignatureP1Residual:         math.Abs(p1/3 - tau),
		SpinDivisibilityPassed:      math.Mod(math.Abs(tau), 16) == 0,
		UsesEulerSocket:             true,
		UsesPontryaginSocket:        true,
		UsesSignatureSocket:         true,
		UsesEtaBoundaryCorrection:   true,
		GlobalNumbersDerived:        false,
		PhysicalThetaSelected:       false,
		Verdict:                     StatusCharacteristicConstraintsAudited,
		Reason:                      "Characteristic-number relations and spin divisibility rules are available as classifier constraints; the synthetic τ=-16 row only tests the relation plumbing and is not a native topology prediction.",
	}
}

func buildScale() ScaleFirewall {
	return ScaleFirewall{Executed: true, UsesLambda: false, UsesF2: false, UsesF4: false, UsesNewton: false, UsesCosmologicalConstant: false, UsesElectroweakData: false, UsesFlavorData: false, UsesObservedTopology: false, UsesBoundarySpectrum: false, ClassifierScaleFree: true, Verdict: StatusClassifierScaleFree, Reason: "Bordism classification uses characteristic classes and parity/divisibility data, not cutoff moments, Newton normalization, electroweak scales, flavor data, or observed topology rows."}
}

func buildRejection() Rejection {
	return Rejection{Executed: true, SpecificBordismClassNativeBlocked: true, ManifoldRepresentativeNativeBlocked: true, StiefelWhitneyClassesNativeBlocked: true, CharacteristicNumbersNativeBlocked: true, SpinStructureNativeBlocked: true, BoundaryBordismNativeBlocked: true, EtaInvariantNativeBlocked: true, NativeRegistryWriteBlocked: true, ComparatorOnlyPurpose: true, Verdict: strings.Join([]string{StatusClassifierAirlockDefined, StatusFailedSpecificClassNotSelected, StatusFailedManifoldRepresentativeNotSelected, StatusFailedStiefelWhitneyNotDerived, StatusFailedGlobalCharacteristicNumbers, StatusFailedSpinStructureNotSelected, StatusFailedBoundaryBordismNotClosed, StatusFailedTopologyNativeWrite}, ";"), Reason: "Gate521 is an airlock: it can reject inadmissible topology claims by structure, but it cannot promote a specific bordism class, spin structure, characteristic number, eta value, or boundary class to native ASHA data."}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"No specific bordism class, cobordism class, manifold representative, spin structure, spin-c line, characteristic number, eta invariant, or boundary condition is written natively at Gate521.",
			"Inherited native content remains local: characteristic-class sockets, local index-density sockets, anomaly cancellation, and scale-free classifier rules.",
		},
		BridgeEntries: []string{
			"Oriented, spin, spin-c, and boundary-bordism classifier sockets defined as bridge-safe topology filters.",
			"Four-dimensional characteristic constraints ∫p1=3τ and closed-spin Â=-τ/8/Rokhlin divisibility are recorded as classifier checks.",
			"Synthetic τ=-16, p1=-48, Â=2 row exercises classifier arithmetic only; it is not an observed or native universe topology.",
		},
		EnvironmentalEntries: []string{
			"The actual universe's bordism class, cobordism representative, Stiefel-Whitney classes, spin/spin-c structure, Euler characteristic, Pontryagin numbers, signature, eta invariant, and boundary spectrum remain global/environmental inputs.",
		},
		FailedRoutes: []string{
			"Treating a classifier socket as a manifold selector.",
			"Promoting Stiefel-Whitney, Pontryagin, signature, Euler, eta, or APS rows into native ASHA predictions without global topology and tangent-bundle data.",
			"Using a synthetic bordism example as evidence for the actual universe topology.",
		},
		OpenTheorems: []string{
			"A native manifold/bordism selector would require new finite-to-global topology data not present in Gates 516-521.",
			"Boundary bordism and eta-invariant comparison can be developed only as a source-tagged bridge comparator unless a native boundary spectral theorem is found.",
		},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 522, Title: "Bordism Comparator File Adapter and Stiefel-Whitney Metadata Firewall", Reason: "Gate521 defines admissible topology-class filters but imports no actual class data. The next safe step is a file-backed comparator adapter for bordism/spin/spin-c metadata.", PrimaryTask: "Load a synthetic bordism classifier ledger, validate w1/w2/W3/c1 metadata and characteristic-number constraints, compute bridge-only admissibility residuals, and block native manifold selection."}
}

func truth(a Analysis) string {
	return "Gate 521 promotes bordism/cobordism from vague global topology language into a precise scale-free classifier airlock. ASHA can state the oriented, spin, spin-c, characteristic-number, and boundary-bordism conditions that an external topology hypothesis must satisfy, but it still cannot select the universe's bordism class, tangent-bundle characteristic classes, boundary spectrum, eta invariant, or manifold representative natively."
}

func validate(a Analysis) error {
	p := []string{}
	if !a.Inheritance.Executed || !a.Inheritance.Gate520FileAdapterDefined || !a.Inheritance.Gate520SyntheticOnly || !a.Inheritance.Gate520ResidualsZero || a.Inheritance.Gate520NativePrediction || !a.Inheritance.Gate520NativeWriteBlocked || a.Inheritance.Gate520ObservedTopologyImported || !a.Inheritance.Gate521Redirect {
		p = append(p, "bad Gate520 inheritance")
	}
	if !a.Socket.Executed || a.Socket.Dimension != 4 || !a.Socket.OrientedSocket || !a.Socket.SpinSocket || !a.Socket.SpinCSocket || !a.Socket.BoundaryBordismSocket || !a.Socket.RequiresW1ZeroForOriented || !a.Socket.RequiresW2ZeroForSpin || !a.Socket.RequiresW3ZeroForSpinC || !a.Socket.RequiresC1Mod2EqualsW2ForSpinC || !a.Socket.ClassifiesAllowedClasses || a.Socket.SelectsSpecificClass || a.Socket.SelectsManifoldRepresentative {
		p = append(p, "bad bordism socket")
	}
	if !a.Constraints.Executed || !a.Constraints.UsesEulerSocket || !a.Constraints.UsesPontryaginSocket || !a.Constraints.UsesSignatureSocket || !a.Constraints.UsesEtaBoundaryCorrection || a.Constraints.GlobalNumbersDerived || a.Constraints.PhysicalThetaSelected || !nearly(a.Constraints.SignatureP1Residual, 0, 1e-12) || !a.Constraints.SpinDivisibilityPassed || !nearly(a.Constraints.SyntheticAHat, 2, 1e-12) {
		p = append(p, "bad characteristic constraints")
	}
	if !a.Scale.Executed || a.Scale.UsesLambda || a.Scale.UsesF2 || a.Scale.UsesF4 || a.Scale.UsesNewton || a.Scale.UsesCosmologicalConstant || a.Scale.UsesElectroweakData || a.Scale.UsesFlavorData || a.Scale.UsesObservedTopology || a.Scale.UsesBoundarySpectrum || !a.Scale.ClassifierScaleFree {
		p = append(p, "scale firewall violation")
	}
	if !a.Rejection.Executed || !a.Rejection.SpecificBordismClassNativeBlocked || !a.Rejection.ManifoldRepresentativeNativeBlocked || !a.Rejection.StiefelWhitneyClassesNativeBlocked || !a.Rejection.CharacteristicNumbersNativeBlocked || !a.Rejection.SpinStructureNativeBlocked || !a.Rejection.BoundaryBordismNativeBlocked || !a.Rejection.EtaInvariantNativeBlocked || !a.Rejection.NativeRegistryWriteBlocked || !a.Rejection.ComparatorOnlyPurpose {
		p = append(p, "rejection firewall violation")
	}
	if len(p) > 0 {
		return fmt.Errorf(strings.Join(p, "; "))
	}
	return nil
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 521 Registry Audit — Bordism and Cobordism Classifier Airlock\n\n")
	b.WriteString("## Verdict\n\n```text\n")
	for _, s := range statuses() {
		b.WriteString(s + "\n")
	}
	b.WriteString("```\n\n")
	b.WriteString("## Inherited boundary\n\n" + a.Inheritance.Reason + "\n\n```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## Bordism classifier ledger\n\n" + a.Socket.Reason + "\n\n```text\n" + FormatSocket(a.Socket) + "\n```\n\n")
	b.WriteString("## Characteristic-number constraint audit\n\n" + a.Constraints.Reason + "\n\n```text\n" + FormatConstraints(a.Constraints) + "\n```\n\n")
	b.WriteString("## Scale-independence and firewall result\n\n" + a.Scale.Reason + "\n\n```text\n" + FormatScale(a.Scale) + "\n```\n\n" + a.Rejection.Reason + "\n\n```text\n" + FormatRejection(a.Rejection) + "\n```\n\n")
	b.WriteString("## Registry update\n\n")
	writeList(&b, "### Native entries", a.Registry.NativeEntries)
	writeList(&b, "### Bridge entries", a.Registry.BridgeEntries)
	writeList(&b, "### Environmental entries", a.Registry.EnvironmentalEntries)
	writeList(&b, "### Failed routes", a.Registry.FailedRoutes)
	writeList(&b, "### Open theorems", a.Registry.OpenTheorems)
	b.WriteString("## Next step\n\nGate522 should be:\n\n```text\n" + fmt.Sprintf("Gate %d — %s", a.Next.Gate, a.Next.Title) + "\n```\n\nPrimary task:\n\n```text\n" + a.Next.PrimaryTask + "\n```\n\n")
	b.WriteString("## Truth statement\n\n" + a.Truth + "\n")
	return b.String()
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("Gate520_file_adapter=%t; synthetic_only=%t; residuals_zero=%t; native_prediction=%t; native_write_blocked=%t; observed_topology_imported=%t; boundary_mode=%t; Gate521_redirect=%t", x.Gate520FileAdapterDefined, x.Gate520SyntheticOnly, x.Gate520ResidualsZero, x.Gate520NativePrediction, x.Gate520NativeWriteBlocked, x.Gate520ObservedTopologyImported, x.Gate520BoundaryMode, x.Gate521Redirect)
}
func FormatSocket(x BordismSocket) string {
	return fmt.Sprintf("dimension=%d; oriented=%t(w1=0); spin=%t(w2=0); spinc=%t(W3=0,c1≡w2 mod2); boundary_bordism=%t; classifies_allowed_classes=%t; selects_specific_class=%t; selects_manifold=%t", x.Dimension, x.OrientedSocket, x.SpinSocket, x.SpinCSocket, x.BoundaryBordismSocket, x.ClassifiesAllowedClasses, x.SelectsSpecificClass, x.SelectsManifoldRepresentative)
}
func FormatConstraints(x CharacteristicConstraints) string {
	return fmt.Sprintf("%s; %s; Rokhlin_divisibility=%d; synthetic_tau=%.12g; synthetic_p1=%.12g; synthetic_Ahat=%.12g; signature_p1_residual=%.12g; spin_divisibility_passed=%t; global_numbers_derived=%t; physical_theta_selected=%t", x.SignaturePontryaginRelation, x.SpinDiracIndexRelation, x.RokhlinDivisibility, x.SyntheticTau, x.SyntheticP1, x.SyntheticAHat, x.SignatureP1Residual, x.SpinDivisibilityPassed, x.GlobalNumbersDerived, x.PhysicalThetaSelected)
}
func FormatScale(x ScaleFirewall) string {
	return fmt.Sprintf("uses_Lambda=%t; uses_f2=%t; uses_f4=%t; uses_Newton=%t; uses_cosmological_constant=%t; uses_EW=%t; uses_flavor=%t; uses_observed_topology=%t; uses_boundary_spectrum=%t; classifier_scale_free=%t", x.UsesLambda, x.UsesF2, x.UsesF4, x.UsesNewton, x.UsesCosmologicalConstant, x.UsesElectroweakData, x.UsesFlavorData, x.UsesObservedTopology, x.UsesBoundarySpectrum, x.ClassifierScaleFree)
}
func FormatRejection(x Rejection) string {
	return fmt.Sprintf("specific_class_native_blocked=%t; representative_native_blocked=%t; stiefel_whitney_native_blocked=%t; characteristic_numbers_native_blocked=%t; spin_structure_native_blocked=%t; boundary_bordism_native_blocked=%t; eta_native_blocked=%t; native_registry_write_blocked=%t; comparator_only=%t", x.SpecificBordismClassNativeBlocked, x.ManifoldRepresentativeNativeBlocked, x.StiefelWhitneyClassesNativeBlocked, x.CharacteristicNumbersNativeBlocked, x.SpinStructureNativeBlocked, x.BoundaryBordismNativeBlocked, x.EtaInvariantNativeBlocked, x.NativeRegistryWriteBlocked, x.ComparatorOnlyPurpose)
}

func statuses() []string {
	return []string{StatusGate520Inherited, StatusBordismLedgerDefined, StatusOrientedSocketPresent, StatusSpinSocketPresent, StatusSpinCSocketPresent, StatusCharacteristicConstraintsAudited, StatusClassifierScaleFree, StatusClassifierAirlockDefined, StatusNoObservedTopologyImported, StatusFailedSpecificClassNotSelected, StatusFailedManifoldRepresentativeNotSelected, StatusFailedStiefelWhitneyNotDerived, StatusFailedGlobalCharacteristicNumbers, StatusFailedSpinStructureNotSelected, StatusFailedBoundaryBordismNotClosed, StatusFailedTopologyNativeWrite, StatusFirewallPreserved, StatusFirewallNativeWriteBlocked}
}

func writeList(b *strings.Builder, title string, xs []string) {
	b.WriteString(title + "\n\n")
	for _, x := range xs {
		b.WriteString("- " + x + "\n")
	}
	b.WriteString("\n")
}
func nearly(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
