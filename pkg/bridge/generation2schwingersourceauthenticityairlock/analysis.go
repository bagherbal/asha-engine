// Package generation2schwingersourceauthenticityairlock implements Gate 538:
// Schwinger Source Authenticity Comparator Airlock Preflight.
//
// Gate 537 proved that ASHA can parse and algebraically dry-run a complete
// synthetic 19-row Schwinger-function ledger. Gate 538 defines the next hard
// boundary: a non-synthetic physical or constructive source may be considered
// only after provenance, integrity, license, reproducibility, OS certificate,
// Wick/Hamiltonian metadata, uncertainty, and quarantine tags pass an explicit
// authenticity sieve. This gate imports no real correlators and executes no
// physics comparator.
package generation2schwingersourceauthenticityairlock

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2syntheticschwingerledgeradapter"
)

const (
	AuditID = "GATE538-SCHWINGER-SOURCE-AUTHENTICITY-COMPARATOR-AIRLOCK-PREFLIGHT"

	StatusGate537SyntheticAdapterInherited      = "CONDITIONAL_SUPPORT_GATE537_SYNTHETIC_SCHWINGER_ADAPTER_INHERITED"
	StatusSourceAuthenticityAirlockDefined      = "CONDITIONAL_SUPPORT_SCHWINGER_SOURCE_AUTHENTICITY_AIRLOCK_DEFINED"
	StatusAuthenticitySchemaRowsEnumerated      = "CONDITIONAL_SUPPORT_SCHWINGER_AUTHENTICITY_SCHEMA_ROWS_ENUMERATED"
	StatusPhysicalSourceDiscriminatorDefined    = "CONDITIONAL_SUPPORT_PHYSICAL_SOURCE_DISCRIMINATOR_DEFINED"
	StatusSyntheticFixtureRejectedAsPhysical    = "CONDITIONAL_SUPPORT_SYNTHETIC_FIXTURE_REJECTED_AS_PHYSICAL_SOURCE"
	StatusProvenanceIntegritySieveDefined       = "CONDITIONAL_SUPPORT_PROVENANCE_INTEGRITY_SIEVE_DEFINED"
	StatusComparatorExecutionBlockedInPreflight = "CONDITIONAL_SUPPORT_SOURCE_AUTHENTICITY_COMPARATOR_BLOCKED_IN_PREFLIGHT"
	StatusNoRealSchwingerSourceImported         = "CONDITIONAL_SUPPORT_NO_REAL_SCHWINGER_SOURCE_IMPORTED_IN_GATE538"
	StatusNativePromotionRejected               = "CONDITIONAL_SUPPORT_SOURCE_AUTHENTICITY_NATIVE_PROMOTION_REJECTED"

	StatusFailedAuthenticitySchemaNotSchwinger      = "FAILED_ROUTE_AUTHENTICITY_SCHEMA_DOES_NOT_DERIVE_PHYSICAL_SCHWINGER_FUNCTIONS"
	StatusFailedAuthenticitySchemaNotOSProof        = "FAILED_ROUTE_AUTHENTICITY_SCHEMA_DOES_NOT_PROVE_OS_REFLECTION_POSITIVITY"
	StatusFailedAuthenticitySchemaNotWick           = "FAILED_ROUTE_AUTHENTICITY_SCHEMA_DOES_NOT_GRANT_WICK_ROTATION"
	StatusFailedAuthenticitySchemaNotHilbert        = "FAILED_ROUTE_AUTHENTICITY_SCHEMA_DOES_NOT_SELECT_PHYSICAL_HILBERT_SPACE"
	StatusFailedAuthenticitySchemaNotHamiltonian    = "FAILED_ROUTE_AUTHENTICITY_SCHEMA_DOES_NOT_DERIVE_POSITIVE_ENERGY_HAMILTONIAN"
	StatusFailedAuthenticitySchemaNotUnitary        = "FAILED_ROUTE_AUTHENTICITY_SCHEMA_DOES_NOT_GRANT_UNITARY_REAL_TIME_DYNAMICS"
	StatusFailedAuthenticitySchemaNotGlobal         = "FAILED_ROUTE_AUTHENTICITY_SCHEMA_DOES_NOT_GRANT_GLOBAL_HYPERBOLICITY"
	StatusFailedAuthenticitySchemaNotArrow          = "FAILED_ROUTE_AUTHENTICITY_SCHEMA_DOES_NOT_SELECT_ARROW_OF_TIME"
	StatusFailedNoPhysicalSourceInPreflight         = "FAILED_ROUTE_NO_NON_SYNTHETIC_SCHWINGER_SOURCE_IMPORTED_IN_GATE538_PREFLIGHT"
	StatusFailedSyntheticCannotAuthenticateUniverse = "FAILED_ROUTE_SYNTHETIC_LEDGER_CANNOT_AUTHENTICATE_A_PHYSICAL_UNIVERSE"
	StatusFirewallPreserved                         = "FIREWALL_PRESERVED_GATE538_SOURCE_AUTHENTICITY_AIRLOCK_BRIDGE_ONLY"
	StatusFirewallNativeWriteBlocked                = "FIREWALL_BLOCKED_GATE538_REAL_CORRELATION_NATIVE_WRITE"
)

type Inheritance struct {
	Executed bool

	Gate537SyntheticAdapterExecuted bool
	Gate537RowsAccepted             bool
	Gate537MetadataSieveEnforced    bool
	Gate537FinitePlumbingVerified   bool
	Gate537SyntheticOnly            bool
	Gate537PhysicalDataAbsent       bool
	Gate537NativeWriteBlocked       bool
	Gate538AuthenticityRedirect     bool

	Verdict, Reason string
}

type AuthenticityRow struct {
	Name        string
	Required    bool
	BridgeOnly  bool
	Comparator  bool
	NativeWrite bool
	Reason      string
}

type AuthenticitySchema struct {
	Executed bool

	Rows                      []AuthenticityRow
	RequiredRows              int
	BridgeOnlyRows            int
	ComparatorRows            int
	NativeWriteRows           int
	ImmutableSourceRequired   bool
	NonSyntheticRequired      bool
	LicenseRequired           bool
	ChecksumRequired          bool
	ReproducibilityRequired   bool
	MeasureProvenanceRequired bool
	OSCertificateRequired     bool
	WickHamiltonianRequired   bool
	UncertaintyRequired       bool
	QuarantineTagsRequired    bool
	NativePromotionRejected   bool

	Verdict, Reason string
}

type Discriminator struct {
	Executed bool

	SyntheticLedgerRecognized         bool
	SyntheticLedgerAcceptedAsPhysical bool
	NonSyntheticSourceLoaded          bool
	ObservedCorrelationLoaded         bool
	ConstructiveMeasureLoaded         bool
	PhysicalSchwingerAuthenticated    bool
	IntegrityComparatorExecuted       bool
	OSCertificateComparatorExecuted   bool
	WickHamiltonianComparatorExecuted bool
	NativePromotionAttempted          bool
	NativePromotionBlocked            bool

	Verdict, Reason string
}

type Guard struct {
	Executed bool

	ComparatorExecutionPerformed bool
	PhysicalSchwingerImported    bool
	ConstructiveMeasureImported  bool
	ObservedCorrelationImported  bool
	PhysicalSchwingerDerived     bool
	OSPositivityProven           bool
	WickRotationSelected         bool
	PhysicalHilbertSpaceSelected bool
	PositiveHamiltonianDerived   bool
	UnitaryDynamicsDerived       bool
	GlobalCausalitySelected      bool
	ArrowOfTimeSelected          bool

	Verdict, Reason string
}

type Firewall struct {
	Executed bool

	NativeSchwingerWrite           bool
	NativeConstructiveMeasureWrite bool
	NativeOSProofWrite             bool
	NativeWickWrite                bool
	NativeHilbertWrite             bool
	NativeHamiltonianWrite         bool
	NativeUnitaryWrite             bool
	NativeGlobalCausalWrite        bool
	NativeTimeArrowWrite           bool
	ReopenedFlavorFirewall         bool
	ReopenedEWScaleFirewall        bool
	ReopenedGravityScaleFirewall   bool
	ReopenedTopologyFirewall       bool
	ReopenedDimensionalFirewall    bool
	ReopenedKreinHilbertFirewall   bool
	NativeRegistryWritten          bool

	Verdict, Reason string
}

type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance   Inheritance
	Schema        AuthenticitySchema
	Discriminator Discriminator
	Guard         Guard
	Firewall      Firewall
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
	g537, err := generation2syntheticschwingerledgeradapter.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate537 synthetic Schwinger adapter: %w", err)
	}
	a := Analysis{}
	a.Inheritance = buildInheritance(g537)
	a.Schema = buildSchema(a.Inheritance)
	a.Discriminator = buildDiscriminator(g537, a.Schema)
	a.Guard = buildGuard(a.Discriminator)
	a.Firewall = buildFirewall(a.Guard)
	a.Next = buildNext(a)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g generation2syntheticschwingerledgeradapter.Analysis) Inheritance {
	return Inheritance{
		Executed:                        true,
		Gate537SyntheticAdapterExecuted: g.Output.Executed && g.Output.SyntheticSchwingerAdapterVerified,
		Gate537RowsAccepted:             g.Import.Loaded && g.Import.AcceptedRows == 19 && g.Output.RequiredRowsMatched == 19,
		Gate537MetadataSieveEnforced:    g.Import.AllRowsBridgeOnly && g.Import.AllRowsNoTheoremInput && g.Import.AllRowsSynthetic && g.Import.AllRowsSourceTagged && g.Import.AllRowsConventionTagged,
		Gate537FinitePlumbingVerified:   g.Output.FiniteSchwingerPlumbingVerified && g.Output.AllSyntheticQuadraticsNonnegative,
		Gate537SyntheticOnly:            g.Firewall.SyntheticFixtureOnly && g.Output.SyntheticRows == 19 && g.Output.ObservedRows == 0,
		Gate537PhysicalDataAbsent:       !g.Firewall.PhysicalSchwingerDataImported && !g.Firewall.ObservedCorrelationDataImported && !g.Firewall.ConstructiveMeasureImported,
		Gate537NativeWriteBlocked:       !g.Firewall.NativeRegistryWritten && !g.Firewall.NativeSchwingerFunctionWrite && !g.Firewall.NativeOSPositivityWrite && !g.Firewall.NativeWickWrite && !g.Firewall.NativeHamiltonianWrite,
		Gate538AuthenticityRedirect:     true,
		Verdict:                         strings.Join([]string{StatusGate537SyntheticAdapterInherited, StatusSourceAuthenticityAirlockDefined}, ";"),
		Reason:                          "Gate537 proves parser and finite synthetic OS/Schwinger plumbing, but its all-synthetic metadata must be rejected as physical provenance.",
	}
}

func buildSchema(in Inheritance) AuthenticitySchema {
	rows := []AuthenticityRow{
		{Name: "immutable_source_identifier", Required: true, BridgeOnly: true, Comparator: true, Reason: "A real Schwinger source must identify a stable publication, dataset, proof object, or construction record."},
		{Name: "non_synthetic_physical_or_constructive_claim", Required: true, BridgeOnly: true, Comparator: true, Reason: "The adapter must distinguish a synthetic dry run from a real constructive or observed correlation source."},
		{Name: "license_and_access_rights", Required: true, BridgeOnly: true, Comparator: true, Reason: "Auditable use requires permission, citation, and access metadata before data can enter the bridge."},
		{Name: "content_checksum_or_proof_hash", Required: true, BridgeOnly: true, Comparator: true, Reason: "The imported object must be integrity-pinned so later comparator results are reproducible."},
		{Name: "construction_or_measure_provenance", Required: true, BridgeOnly: true, Comparator: true, Reason: "Physical Schwinger functions need a construction, measure, lattice ensemble, theorem, or controlled approximation lane."},
		{Name: "renormalization_and_regulator_provenance", Required: true, BridgeOnly: true, Comparator: true, Reason: "Scheme, cutoff, lattice spacing, continuum limit, or regularization conventions must be attached to the source."},
		{Name: "field_label_alignment_with_gate536", Required: true, BridgeOnly: true, Comparator: true, Reason: "Source fields must map into the Gate536 field-algebra row before OS comparison is meaningful."},
		{Name: "euclidean_covariance_certificate_provenance", Required: true, BridgeOnly: true, Comparator: true, Reason: "Covariance cannot be inferred from file shape; a source-level certificate or comparator lane is required."},
		{Name: "os_reflection_positivity_certificate_provenance", Required: true, BridgeOnly: true, Comparator: true, Reason: "Reflection positivity must be certified or compared for the source; schema presence is not proof."},
		{Name: "wick_and_i_epsilon_provenance", Required: true, BridgeOnly: true, Comparator: true, Reason: "Analytic continuation convention must be attributable and cannot be guessed by ASHA."},
		{Name: "hamiltonian_spectrum_domain_certificate", Required: true, BridgeOnly: true, Comparator: true, Reason: "Positive energy requires a domain/spectrum certificate downstream of reconstruction."},
		{Name: "uncertainty_validity_and_reproducibility_ledger", Required: true, BridgeOnly: true, Comparator: true, Reason: "Approximation error, theorem domain, ensemble statistics, or validity window must be explicit."},
		{Name: "bridge_only_no_native_write_quarantine", Required: true, BridgeOnly: true, Comparator: false, Reason: "Even an authenticated physical source remains bridge/environmental data and cannot rewrite native law."},
	}
	bridge, comparator, native := 0, 0, 0
	for _, r := range rows {
		if r.BridgeOnly {
			bridge++
		}
		if r.Comparator {
			comparator++
		}
		if r.NativeWrite {
			native++
		}
	}
	return AuthenticitySchema{
		Executed:                  true,
		Rows:                      rows,
		RequiredRows:              len(rows),
		BridgeOnlyRows:            bridge,
		ComparatorRows:            comparator,
		NativeWriteRows:           native,
		ImmutableSourceRequired:   true,
		NonSyntheticRequired:      true,
		LicenseRequired:           true,
		ChecksumRequired:          true,
		ReproducibilityRequired:   true,
		MeasureProvenanceRequired: true,
		OSCertificateRequired:     true,
		WickHamiltonianRequired:   true,
		UncertaintyRequired:       true,
		QuarantineTagsRequired:    true,
		NativePromotionRejected:   in.Gate537NativeWriteBlocked,
		Verdict:                   strings.Join([]string{StatusAuthenticitySchemaRowsEnumerated, StatusPhysicalSourceDiscriminatorDefined, StatusProvenanceIntegritySieveDefined, StatusNativePromotionRejected}, ";"),
		Reason:                    "Gate538 enumerates the provenance/integrity sieve required before a non-synthetic Schwinger source may be compared; it does not import that source.",
	}
}

func buildDiscriminator(g generation2syntheticschwingerledgeradapter.Analysis, schema AuthenticitySchema) Discriminator {
	return Discriminator{
		Executed:                          true,
		SyntheticLedgerRecognized:         g.Import.AllRowsSynthetic && g.Firewall.SyntheticFixtureOnly,
		SyntheticLedgerAcceptedAsPhysical: false,
		NonSyntheticSourceLoaded:          false,
		ObservedCorrelationLoaded:         false,
		ConstructiveMeasureLoaded:         false,
		PhysicalSchwingerAuthenticated:    false,
		IntegrityComparatorExecuted:       false,
		OSCertificateComparatorExecuted:   false,
		WickHamiltonianComparatorExecuted: false,
		NativePromotionAttempted:          false,
		NativePromotionBlocked:            schema.NativePromotionRejected,
		Verdict:                           strings.Join([]string{StatusSyntheticFixtureRejectedAsPhysical, StatusNoRealSchwingerSourceImported, StatusFailedSyntheticCannotAuthenticateUniverse, StatusFailedNoPhysicalSourceInPreflight}, ";"),
		Reason:                            "The inherited Gate537 ledger is intentionally synthetic; Gate538 recognizes it as a valid dry-run fixture and rejects it as physical provenance.",
	}
}

func buildGuard(d Discriminator) Guard {
	return Guard{
		Executed: true,
		Verdict:  strings.Join([]string{StatusComparatorExecutionBlockedInPreflight, StatusFailedAuthenticitySchemaNotSchwinger, StatusFailedAuthenticitySchemaNotOSProof, StatusFailedAuthenticitySchemaNotWick, StatusFailedAuthenticitySchemaNotHilbert, StatusFailedAuthenticitySchemaNotHamiltonian, StatusFailedAuthenticitySchemaNotUnitary, StatusFailedAuthenticitySchemaNotGlobal, StatusFailedAuthenticitySchemaNotArrow}, ";"),
		Reason:   "Gate538 defines source-authenticity acceptance criteria only; no non-synthetic Schwinger source is loaded and no physical comparator executes.",
	}
}

func buildFirewall(g Guard) Firewall {
	return Firewall{
		Executed: true,
		Verdict:  strings.Join([]string{StatusFirewallPreserved, StatusFirewallNativeWriteBlocked, StatusFailedAuthenticitySchemaNotSchwinger, StatusFailedAuthenticitySchemaNotOSProof, StatusFailedAuthenticitySchemaNotWick, StatusFailedAuthenticitySchemaNotHilbert, StatusFailedAuthenticitySchemaNotHamiltonian, StatusFailedAuthenticitySchemaNotUnitary, StatusFailedAuthenticitySchemaNotGlobal, StatusFailedAuthenticitySchemaNotArrow}, ";"),
		Reason:   "Authenticity metadata can authorize future bridge comparison, but it cannot promote Schwinger functions, OS proof, Wick continuation, Hilbert space, Hamiltonian, dynamics, causality, or time orientation into native ASHA law.",
	}
}

func buildNext(a Analysis) NextStep {
	return NextStep{Gate: 539, Title: "Synthetic Source-Authenticity Ledger Adapter Dry Run", Reason: "Gate538 defines the provenance/integrity discriminator. The safe next step is a synthetic authenticity ledger that satisfies the schema while still being rejected as native physics.", PrimaryTask: "Load a fake non-observed source-authentication ledger, check checksums/provenance/quarantine tags, and block every physical/native promotion path."}
}

func truth(a Analysis) string {
	return "Gate538 separates parser completeness from source authenticity: ASHA can now state what would make a Schwinger import auditable, but no real Schwinger source, constructive measure, OS proof, Wick map, Hilbert space, Hamiltonian, unitarity, global causality, or time arrow is imported or derived."
}

func validate(a Analysis) error {
	bad := []string{}
	if !a.Inheritance.Executed || !a.Inheritance.Gate537SyntheticAdapterExecuted || !a.Inheritance.Gate537RowsAccepted || !a.Inheritance.Gate537MetadataSieveEnforced || !a.Inheritance.Gate537FinitePlumbingVerified || !a.Inheritance.Gate537SyntheticOnly || !a.Inheritance.Gate537PhysicalDataAbsent || !a.Inheritance.Gate537NativeWriteBlocked || !a.Inheritance.Gate538AuthenticityRedirect {
		bad = append(bad, "bad Gate537 inheritance")
	}
	if !a.Schema.Executed || a.Schema.RequiredRows != 13 || a.Schema.BridgeOnlyRows != 13 || a.Schema.ComparatorRows != 12 || a.Schema.NativeWriteRows != 0 || !a.Schema.ImmutableSourceRequired || !a.Schema.NonSyntheticRequired || !a.Schema.LicenseRequired || !a.Schema.ChecksumRequired || !a.Schema.ReproducibilityRequired || !a.Schema.MeasureProvenanceRequired || !a.Schema.OSCertificateRequired || !a.Schema.WickHamiltonianRequired || !a.Schema.UncertaintyRequired || !a.Schema.QuarantineTagsRequired || !a.Schema.NativePromotionRejected {
		bad = append(bad, "bad authenticity schema")
	}
	if !a.Discriminator.Executed || !a.Discriminator.SyntheticLedgerRecognized || a.Discriminator.SyntheticLedgerAcceptedAsPhysical || a.Discriminator.NonSyntheticSourceLoaded || a.Discriminator.ObservedCorrelationLoaded || a.Discriminator.ConstructiveMeasureLoaded || a.Discriminator.PhysicalSchwingerAuthenticated || a.Discriminator.IntegrityComparatorExecuted || a.Discriminator.OSCertificateComparatorExecuted || a.Discriminator.WickHamiltonianComparatorExecuted || a.Discriminator.NativePromotionAttempted || !a.Discriminator.NativePromotionBlocked {
		bad = append(bad, "bad source discriminator")
	}
	if !a.Guard.Executed || a.Guard.ComparatorExecutionPerformed || a.Guard.PhysicalSchwingerImported || a.Guard.ConstructiveMeasureImported || a.Guard.ObservedCorrelationImported || a.Guard.PhysicalSchwingerDerived || a.Guard.OSPositivityProven || a.Guard.WickRotationSelected || a.Guard.PhysicalHilbertSpaceSelected || a.Guard.PositiveHamiltonianDerived || a.Guard.UnitaryDynamicsDerived || a.Guard.GlobalCausalitySelected || a.Guard.ArrowOfTimeSelected {
		bad = append(bad, "bad guard")
	}
	if !a.Firewall.Executed || a.Firewall.NativeSchwingerWrite || a.Firewall.NativeConstructiveMeasureWrite || a.Firewall.NativeOSProofWrite || a.Firewall.NativeWickWrite || a.Firewall.NativeHilbertWrite || a.Firewall.NativeHamiltonianWrite || a.Firewall.NativeUnitaryWrite || a.Firewall.NativeGlobalCausalWrite || a.Firewall.NativeTimeArrowWrite || a.Firewall.ReopenedFlavorFirewall || a.Firewall.ReopenedEWScaleFirewall || a.Firewall.ReopenedGravityScaleFirewall || a.Firewall.ReopenedTopologyFirewall || a.Firewall.ReopenedDimensionalFirewall || a.Firewall.ReopenedKreinHilbertFirewall || a.Firewall.NativeRegistryWritten {
		bad = append(bad, "bad firewall")
	}
	if len(bad) > 0 {
		return fmt.Errorf("Gate538 validation failed: %s", strings.Join(bad, "; "))
	}
	return nil
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("%s: adapter=%t rows=%t metadata=%t plumbing=%t synthetic_only=%t physical_absent=%t native_blocked=%t redirect=%t; %s", x.Verdict, x.Gate537SyntheticAdapterExecuted, x.Gate537RowsAccepted, x.Gate537MetadataSieveEnforced, x.Gate537FinitePlumbingVerified, x.Gate537SyntheticOnly, x.Gate537PhysicalDataAbsent, x.Gate537NativeWriteBlocked, x.Gate538AuthenticityRedirect, x.Reason)
}

func FormatSchema(x AuthenticitySchema) string {
	names := []string{}
	for _, row := range x.Rows {
		names = append(names, row.Name)
	}
	return fmt.Sprintf("%s: rows=%d required=%d bridge_only=%d comparator=%d native_write=%d immutable=%t nonsynthetic=%t license=%t checksum=%t reproducibility=%t measure=%t OS=%t wick_hamiltonian=%t uncertainty=%t quarantine=%t native_rejected=%t names=[%s]; %s", x.Verdict, len(x.Rows), x.RequiredRows, x.BridgeOnlyRows, x.ComparatorRows, x.NativeWriteRows, x.ImmutableSourceRequired, x.NonSyntheticRequired, x.LicenseRequired, x.ChecksumRequired, x.ReproducibilityRequired, x.MeasureProvenanceRequired, x.OSCertificateRequired, x.WickHamiltonianRequired, x.UncertaintyRequired, x.QuarantineTagsRequired, x.NativePromotionRejected, strings.Join(names, ","), x.Reason)
}

func FormatDiscriminator(x Discriminator) string {
	return fmt.Sprintf("%s: synthetic_recognized=%t synthetic_as_physical=%t nonsynthetic_loaded=%t observed=%t constructive=%t authenticated=%t integrity_cmp=%t OS_cmp=%t wick_hamiltonian_cmp=%t native_attempt=%t native_blocked=%t; %s", x.Verdict, x.SyntheticLedgerRecognized, x.SyntheticLedgerAcceptedAsPhysical, x.NonSyntheticSourceLoaded, x.ObservedCorrelationLoaded, x.ConstructiveMeasureLoaded, x.PhysicalSchwingerAuthenticated, x.IntegrityComparatorExecuted, x.OSCertificateComparatorExecuted, x.WickHamiltonianComparatorExecuted, x.NativePromotionAttempted, x.NativePromotionBlocked, x.Reason)
}

func FormatGuard(x Guard) string {
	return fmt.Sprintf("%s: comparator=%t physical_import=%t constructive_import=%t observed_import=%t Schwinger_derived=%t OS=%t Wick=%t Hilbert=%t Hamiltonian=%t unitary=%t global=%t arrow=%t; %s", x.Verdict, x.ComparatorExecutionPerformed, x.PhysicalSchwingerImported, x.ConstructiveMeasureImported, x.ObservedCorrelationImported, x.PhysicalSchwingerDerived, x.OSPositivityProven, x.WickRotationSelected, x.PhysicalHilbertSpaceSelected, x.PositiveHamiltonianDerived, x.UnitaryDynamicsDerived, x.GlobalCausalitySelected, x.ArrowOfTimeSelected, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("%s: native_Schwinger=%t native_measure=%t native_OS=%t native_Wick=%t native_Hilbert=%t native_Hamiltonian=%t native_unitary=%t native_global=%t native_arrow=%t reopened_flavor=%t reopened_EW=%t reopened_gravity=%t reopened_topology=%t reopened_dimension=%t reopened_Krein=%t native_registry=%t; %s", x.Verdict, x.NativeSchwingerWrite, x.NativeConstructiveMeasureWrite, x.NativeOSProofWrite, x.NativeWickWrite, x.NativeHilbertWrite, x.NativeHamiltonianWrite, x.NativeUnitaryWrite, x.NativeGlobalCausalWrite, x.NativeTimeArrowWrite, x.ReopenedFlavorFirewall, x.ReopenedEWScaleFirewall, x.ReopenedGravityScaleFirewall, x.ReopenedTopologyFirewall, x.ReopenedDimensionalFirewall, x.ReopenedKreinHilbertFirewall, x.NativeRegistryWritten, x.Reason)
}

func statuses() []string {
	return []string{StatusGate537SyntheticAdapterInherited, StatusSourceAuthenticityAirlockDefined, StatusAuthenticitySchemaRowsEnumerated, StatusPhysicalSourceDiscriminatorDefined, StatusSyntheticFixtureRejectedAsPhysical, StatusProvenanceIntegritySieveDefined, StatusComparatorExecutionBlockedInPreflight, StatusNoRealSchwingerSourceImported, StatusNativePromotionRejected, StatusFailedAuthenticitySchemaNotSchwinger, StatusFailedAuthenticitySchemaNotOSProof, StatusFailedAuthenticitySchemaNotWick, StatusFailedAuthenticitySchemaNotHilbert, StatusFailedAuthenticitySchemaNotHamiltonian, StatusFailedAuthenticitySchemaNotUnitary, StatusFailedAuthenticitySchemaNotGlobal, StatusFailedAuthenticitySchemaNotArrow, StatusFailedNoPhysicalSourceInPreflight, StatusFailedSyntheticCannotAuthenticateUniverse, StatusFirewallPreserved, StatusFirewallNativeWriteBlocked}
}
