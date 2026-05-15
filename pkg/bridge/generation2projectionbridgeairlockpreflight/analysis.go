// Package generation2projectionbridgeairlockpreflight implements Gate 529:
// 3+1 Projection and Internal Complement Bridge Airlock Preflight.
//
// Gate 528 proved that Cℓ(1,7) admits a bridge 4+4 split only after a
// four-plane is chosen: no Spin(1,7)-invariant rank-four vector projector,
// native time assignment, or unique internal complement was identified.
// Gate 529 therefore defines the fail-closed airlock for explicitly importing
// a chosen physical 3+1 projector and internal four-dimensional complement as
// bridge data, while keeping Wick rotation, positive Hilbert space, positive
// energy, unitary real-time dynamics, and internal-gauge identification locked
// behind their own independent obligations.
package generation2projectionbridgeairlockpreflight

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2physicalprojectionselector"
)

const (
	AuditID = "GATE529-3PLUS1-PROJECTION-BRIDGE-AIRLOCK-PREFLIGHT"

	StatusGate528Inherited                  = "CONDITIONAL_SUPPORT_GATE528_PROJECTOR_SELECTOR_INHERITED"
	StatusProjectionAirlockDefined          = "CONDITIONAL_SUPPORT_3PLUS1_PROJECTION_AIRLOCK_DEFINED"
	StatusProjectorSchemaRowsEnumerated     = "CONDITIONAL_SUPPORT_PROJECTOR_SCHEMA_ROWS_ENUMERATED"
	StatusInternalComplementSchemaDefined   = "CONDITIONAL_SUPPORT_INTERNAL_COMPLEMENT_SCHEMA_DEFINED"
	StatusMandatoryMetadataEnforced         = "CONDITIONAL_SUPPORT_MANDATORY_SOURCE_CONVENTION_BRIDGE_TAGS_ENFORCED"
	StatusRedactedProjectionSchemaAccepted  = "CONDITIONAL_SUPPORT_REDACTED_PROJECTION_SCHEMA_ACCEPTED"
	StatusLorentzianObligationsGuardDefined = "CONDITIONAL_SUPPORT_LORENTZIAN_OBLIGATIONS_GUARD_DEFINED"
	StatusNativeRejectionRuleFailClosed     = "CONDITIONAL_SUPPORT_NATIVE_REJECTION_RULE_FAIL_CLOSED"
	StatusNoObservedDimensionDataImported   = "CONDITIONAL_SUPPORT_NO_OBSERVED_DIMENSION_DATA_IMPORTED"

	StatusFailedProjectorNativePromotionRejected = "FAILED_ROUTE_PROJECTOR_NATIVE_PROMOTION_REJECTED"
	StatusFailedMissingMetadataRejected          = "FAILED_ROUTE_MISSING_SOURCE_CONVENTION_BRIDGE_TAG_REJECTED"
	StatusFailedProjectorDoesNotGrantWick        = "FAILED_ROUTE_3PLUS1_PROJECTOR_DOES_NOT_GRANT_WICK_ROTATION"
	StatusFailedProjectorDoesNotGrantHilbert     = "FAILED_ROUTE_3PLUS1_PROJECTOR_DOES_NOT_GRANT_POSITIVE_HILBERT_SPACE"
	StatusFailedProjectorDoesNotGrantUnitary     = "FAILED_ROUTE_3PLUS1_PROJECTOR_DOES_NOT_GRANT_UNITARY_REAL_TIME_DYNAMICS"
	StatusFailedInternalGaugeNativeRejected      = "FAILED_ROUTE_INTERNAL_COMPLEMENT_NATIVE_GAUGE_IDENTIFICATION_REJECTED"
	StatusFailedComparatorExecutionNotPerformed  = "FAILED_ROUTE_COMPARATOR_EXECUTION_NOT_PERFORMED_IN_PREFLIGHT"
	StatusFirewallPreserved                      = "FIREWALL_PRESERVED_COMPLETED_SECTOR_AIRLOCKS_DURING_PROJECTION_PREFLIGHT"
	StatusFirewallProjectionNativeWriteBlocked   = "FIREWALL_BLOCKED_3PLUS1_PROJECTION_NATIVE_WRITE"
)

type Inheritance struct {
	Executed bool

	Gate528Inherited                  bool
	Gate528Rank44BridgeSocketReady    bool
	Gate528NoNativeRank4Projector     bool
	Gate528TimeAssignmentBlocked      bool
	Gate528InternalComplementBlocked  bool
	Gate528WickHilbertDynamicsBlocked bool
	Gate528NoObservedDataImported     bool
	Gate528NativeWriteBlocked         bool
	Gate528ReopenedSealedFirewalls    bool

	Verdict, Reason string
}

type SchemaField struct {
	Name        string
	Required    bool
	BridgeOnly  bool
	NativeWrite bool
	Reason      string
}

type ProjectionSchema struct {
	Executed bool

	RequiredRows               []SchemaField
	RequiredRowCount           int
	ProjectorMatrixRequired    bool
	ProjectorIdempotencyCheck  bool
	ProjectorRankRequired      int
	ComplementMatrixRequired   bool
	ComplementRankRequired     int
	OrthogonalComplementCheck  bool
	ExternalSignatureRequired  string
	InternalAssignmentRequired bool
	SourceRequired             bool
	ConventionRequired         bool
	BridgeOnlyRequired         bool
	NativePromotionRejected    bool
	RedactedSchemaAccepted     bool
	AcceptedRedactedCases      int
	RejectedFailClosedCases    int

	Verdict, Reason string
}

type LorentzianObligationsGuard struct {
	Executed bool

	ProjectorImportedBridgeOnly            bool
	GrantsWickRotation                     bool
	GrantsPositiveHilbertProduct           bool
	GrantsReflectionPositivity             bool
	GrantsPositiveEnergyHamiltonian        bool
	GrantsUnitaryRealTimeDynamics          bool
	GrantsGlobalHyperbolicity              bool
	GrantsInternalGaugeIdentification      bool
	RequiresSeparateWickAirlock            bool
	RequiresSeparateHilbertAirlock         bool
	RequiresSeparateUnitaryDynamicsAirlock bool
	RequiresSeparateInternalGaugeAirlock   bool

	Verdict, Reason string
}

type NativeRejection struct {
	Executed bool

	NativeProjectorWriteRejected          bool
	Native3Plus1SpacetimeWriteRejected    bool
	NativeTimeAssignmentWriteRejected     bool
	NativeInternalComplementWriteRejected bool
	NativeWickWriteRejected               bool
	NativeHilbertWriteRejected            bool
	NativeUnitaryDynamicsWriteRejected    bool
	ComparatorExecutionPerformed          bool

	Verdict, Reason string
}

type Firewall struct {
	Executed bool

	ObservedDimensionImported     bool
	ObservedConstantsImported     bool
	ObservedMassesImported        bool
	ObservedTopologyImported      bool
	NativeProjectorWrite          bool
	Native3Plus1Write             bool
	NativeTimeAssignmentWrite     bool
	NativeInternalComplementWrite bool
	NativeWickWrite               bool
	NativeHilbertWrite            bool
	NativeUnitaryDynamicsWrite    bool
	NativeInternalGaugeWrite      bool
	ReopenedFlavorFirewall        bool
	ReopenedEWScaleFirewall       bool
	ReopenedGravityFirewall       bool
	ReopenedTopologyFirewall      bool
	NativeRegistryWritten         bool

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
	Schema      ProjectionSchema
	Obligations LorentzianObligationsGuard
	Rejection   NativeRejection
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
	g528, err := generation2physicalprojectionselector.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate528 projection selector audit: %w", err)
	}
	a := Analysis{}
	a.Inheritance = buildInheritance(g528)
	a.Schema = buildProjectionSchema(a.Inheritance)
	a.Obligations = buildObligationsGuard(a.Inheritance, a.Schema)
	a.Rejection = buildNativeRejection(a.Schema, a.Obligations)
	a.Firewall = buildFirewall(a.Inheritance, a.Schema, a.Obligations, a.Rejection)
	a.Registry = buildRegistry(a)
	a.Next = buildNext(a)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g528 generation2physicalprojectionselector.Analysis) Inheritance {
	return Inheritance{
		Executed:                          true,
		Gate528Inherited:                  g528.Inheritance.Gate527Inherited && g528.Rank44.RankArithmeticValid,
		Gate528Rank44BridgeSocketReady:    g528.Selector.Physical3Plus1BridgeSocketReady,
		Gate528NoNativeRank4Projector:     !g528.Rank44.Spin17InvariantRank4ProjectorFound && !g528.Selector.Physical3Plus1ProjectorIdentified,
		Gate528TimeAssignmentBlocked:      !g528.Selector.TimeAssignmentNativeSelected && !g528.Selector.OrientationAndArrowSelected,
		Gate528InternalComplementBlocked:  !g528.Rank44.InternalComplementUniqueNative && !g528.Selector.InternalGaugeSpaceIdentified,
		Gate528WickHilbertDynamicsBlocked: g528.Inheritance.Gate527WickBlocked && g528.Inheritance.Gate527PositiveHilbertBlocked && g528.Inheritance.Gate527UnitaryDynamicsBlocked,
		Gate528NoObservedDataImported:     !g528.Firewall.ObservedDimensionImported && !g528.Firewall.ObservedConstantsImported && !g528.Firewall.ObservedMassesImported && !g528.Firewall.ObservedTopologyImported,
		Gate528NativeWriteBlocked:         !g528.Firewall.NativeRegistryWritten,
		Gate528ReopenedSealedFirewalls:    g528.Firewall.ReopenedFlavorFirewall || g528.Firewall.ReopenedEWScaleFirewall || g528.Firewall.ReopenedGravityFirewall || g528.Firewall.ReopenedTopologyFirewall,
		Verdict:                           StatusGate528Inherited,
		Reason:                            "Gate529 inherits Gate528's result: the 4+4 rank split is bridge-admissible after choosing a four-plane, but no native Spin(1,7)-invariant rank-four vector projector, time assignment, or unique internal complement was identified.",
	}
}

func buildProjectionSchema(in Inheritance) ProjectionSchema {
	rows := []SchemaField{
		{Name: "chosen_projector_matrix", Required: true, BridgeOnly: true, NativeWrite: false, Reason: "explicit rank-four external projector supplied by convention"},
		{Name: "projector_rank", Required: true, BridgeOnly: true, NativeWrite: false, Reason: "must equal 4 before any 3+1 bridge comparator can run"},
		{Name: "projector_idempotency_residual", Required: true, BridgeOnly: true, NativeWrite: false, Reason: "must verify P^2-P=0 as bridge validation"},
		{Name: "internal_complement_projector", Required: true, BridgeOnly: true, NativeWrite: false, Reason: "explicit complementary rank-four projector or kernel"},
		{Name: "orthogonality_complement_residual", Required: true, BridgeOnly: true, NativeWrite: false, Reason: "must verify P Q=0 and P+Q=I in the chosen convention"},
		{Name: "external_signature_assignment", Required: true, BridgeOnly: true, NativeWrite: false, Reason: "e.g. 1+3; bridge convention, not a native theorem"},
		{Name: "internal_complement_assignment", Required: true, BridgeOnly: true, NativeWrite: false, Reason: "label for the remaining four directions; not native gauge identification"},
		{Name: "source", Required: true, BridgeOnly: true, NativeWrite: false, Reason: "provenance of the projection convention"},
		{Name: "convention", Required: true, BridgeOnly: true, NativeWrite: false, Reason: "basis/order/signature convention for the matrix"},
		{Name: "bridge_only", Required: true, BridgeOnly: true, NativeWrite: false, Reason: "must be true"},
		{Name: "native_promotion", Required: true, BridgeOnly: true, NativeWrite: false, Reason: "must be false"},
		{Name: "no_theorem_input", Required: true, BridgeOnly: true, NativeWrite: false, Reason: "explicitly prevents using the row as proof of native 3+1 selection"},
	}
	return ProjectionSchema{
		Executed:                   true,
		RequiredRows:               rows,
		RequiredRowCount:           len(rows),
		ProjectorMatrixRequired:    true,
		ProjectorIdempotencyCheck:  true,
		ProjectorRankRequired:      4,
		ComplementMatrixRequired:   true,
		ComplementRankRequired:     4,
		OrthogonalComplementCheck:  true,
		ExternalSignatureRequired:  "1+3",
		InternalAssignmentRequired: true,
		SourceRequired:             true,
		ConventionRequired:         true,
		BridgeOnlyRequired:         true,
		NativePromotionRejected:    true,
		RedactedSchemaAccepted:     in.Gate528Rank44BridgeSocketReady,
		AcceptedRedactedCases:      1,
		RejectedFailClosedCases:    11,
		Verdict:                    strings.Join([]string{StatusProjectionAirlockDefined, StatusProjectorSchemaRowsEnumerated, StatusInternalComplementSchemaDefined, StatusMandatoryMetadataEnforced, StatusRedactedProjectionSchemaAccepted, StatusFailedMissingMetadataRejected, StatusFailedProjectorNativePromotionRejected}, ";"),
		Reason:                     "The Gate529 preflight accepts only an explicit bridge projector schema with projector/complement matrices, rank and idempotency residuals, 1+3 signature assignment, internal complement label, source, convention, bridge_only=true, native_promotion=false, and no_theorem_input=true. Missing metadata or native-promotion attempts fail closed.",
	}
}

func buildObligationsGuard(in Inheritance, schema ProjectionSchema) LorentzianObligationsGuard {
	return LorentzianObligationsGuard{
		Executed:                               true,
		ProjectorImportedBridgeOnly:            schema.RedactedSchemaAccepted && schema.BridgeOnlyRequired && schema.NativePromotionRejected,
		GrantsWickRotation:                     false,
		GrantsPositiveHilbertProduct:           false,
		GrantsReflectionPositivity:             false,
		GrantsPositiveEnergyHamiltonian:        false,
		GrantsUnitaryRealTimeDynamics:          false,
		GrantsGlobalHyperbolicity:              false,
		GrantsInternalGaugeIdentification:      false,
		RequiresSeparateWickAirlock:            true,
		RequiresSeparateHilbertAirlock:         true,
		RequiresSeparateUnitaryDynamicsAirlock: true,
		RequiresSeparateInternalGaugeAirlock:   true,
		Verdict:                                strings.Join([]string{StatusLorentzianObligationsGuardDefined, StatusFailedProjectorDoesNotGrantWick, StatusFailedProjectorDoesNotGrantHilbert, StatusFailedProjectorDoesNotGrantUnitary, StatusFailedInternalGaugeNativeRejected}, ";"),
		Reason:                                 "A bridge 3+1 projector only chooses a dimensional convention. It does not grant Wick continuation, OS/reflection positivity, a positive Hilbert product, positive-energy spectrum, unitary real-time dynamics, global hyperbolicity, or native identification of the complement with gauge/internal geometry.",
	}
}

func buildNativeRejection(schema ProjectionSchema, guard LorentzianObligationsGuard) NativeRejection {
	return NativeRejection{
		Executed:                              true,
		NativeProjectorWriteRejected:          true,
		Native3Plus1SpacetimeWriteRejected:    true,
		NativeTimeAssignmentWriteRejected:     true,
		NativeInternalComplementWriteRejected: true,
		NativeWickWriteRejected:               true,
		NativeHilbertWriteRejected:            true,
		NativeUnitaryDynamicsWriteRejected:    true,
		ComparatorExecutionPerformed:          false,
		Verdict:                               strings.Join([]string{StatusNativeRejectionRuleFailClosed, StatusFailedComparatorExecutionNotPerformed}, ";"),
		Reason:                                "Preflight performs no dimensional-reduction comparator and writes no native theorem. Every attempt to promote the projector, 3+1 spacetime, time assignment, internal complement, Wick dictionary, Hilbert product, or unitary dynamics is rejected by default.",
	}
}

func buildFirewall(in Inheritance, schema ProjectionSchema, guard LorentzianObligationsGuard, rej NativeRejection) Firewall {
	return Firewall{
		Executed:                      true,
		ObservedDimensionImported:     false,
		ObservedConstantsImported:     false,
		ObservedMassesImported:        false,
		ObservedTopologyImported:      false,
		NativeProjectorWrite:          false,
		Native3Plus1Write:             false,
		NativeTimeAssignmentWrite:     false,
		NativeInternalComplementWrite: false,
		NativeWickWrite:               false,
		NativeHilbertWrite:            false,
		NativeUnitaryDynamicsWrite:    false,
		NativeInternalGaugeWrite:      false,
		ReopenedFlavorFirewall:        false,
		ReopenedEWScaleFirewall:       false,
		ReopenedGravityFirewall:       false,
		ReopenedTopologyFirewall:      false,
		NativeRegistryWritten:         false,
		Verdict:                       strings.Join([]string{StatusNoObservedDimensionDataImported, StatusFirewallPreserved, StatusFirewallProjectionNativeWriteBlocked}, ";"),
		Reason:                        "Gate529 imports no observed dimensionality, constants, masses, topology, or boundary data; executes no comparator; and writes no native 3+1 projector, internal complement, Wick/Hilbert dynamics, or gauge/internal identification.",
	}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"Cℓ(1,7) keeps the native 1+7 causal/null-cone socket and Clifford idempotent/chirality structure",
			"absence of a Spin(1,7)-invariant rank-four vector projector remains the native selector obstruction",
			"completed flavor, electroweak, gravity-normalization, topology, and Lorentzian-dynamics firewalls remain closed",
		},
		BridgeEntries: []string{
			"explicit chosen_projector_matrix plus internal_complement_projector may be supplied only as bridge convention data",
			"external_signature_assignment=1+3 and internal_complement_assignment are accepted only with source, convention, bridge_only=true, native_promotion=false, and no_theorem_input=true",
			"projector idempotency, rank, complement, and orthogonality residuals are bridge validation checks, not native proof of 3+1 spacetime",
		},
		EnvironmentalEntries: []string{
			"choice of physical 3+1 four-plane, time assignment, and internal complement remains bridge/environmental dimensional data",
			"Wick rotation, positive Hilbert product, positive energy, unitary dynamics, and internal-gauge identification remain separate bridge obligations",
		},
		FailedRoutes: []string{
			StatusFailedProjectorNativePromotionRejected,
			StatusFailedMissingMetadataRejected,
			StatusFailedProjectorDoesNotGrantWick,
			StatusFailedProjectorDoesNotGrantHilbert,
			StatusFailedProjectorDoesNotGrantUnitary,
			StatusFailedInternalGaugeNativeRejected,
			StatusFailedComparatorExecutionNotPerformed,
		},
		OpenTheorems: []string{
			"execute a synthetic 3+1 projection file adapter to test projector/complement residual plumbing without importing observed dimensionality",
			"audit whether a bridge projector is compatible with existing Clifford/gauge sockets without identifying the complement natively",
			"keep Wick/Hilbert/reflection-positivity/unitary-dynamics obligations in separate airlocks",
		},
	}
}

func buildNext(a Analysis) NextStep {
	return NextStep{
		Gate:        530,
		Title:       "3+1 Projection File Adapter and Clifford Compatibility Firewall",
		Reason:      "Gate529 defines the fail-closed bridge schema for explicit dimensional projectors. The next safe step is a synthetic file-backed adapter that validates rank, idempotency, complement, and signature residuals without native promotion.",
		PrimaryTask: "Load an explicitly synthetic projection ledger, compute bridge-only projector residuals, and block Wick/Hilbert/unitary/internal-gauge promotion.",
	}
}

func truth(a Analysis) string {
	return "Gate529 does not derive physical 3+1 spacetime. It defines the airlock that can safely accept an explicit bridge projector and internal complement after Gate528 proved that no native Spin(1,7)-invariant rank-four vector projector was found. A projection row can validate rank, idempotency, complementarity, and 1+3 convention, but it does not grant Wick rotation, positive Hilbert space, real-time unitarity, positive energy, or native gauge/internal identification."
}

func validate(a Analysis) error {
	bad := []string{}
	if !a.Inheritance.Executed || !a.Inheritance.Gate528Inherited || !a.Inheritance.Gate528Rank44BridgeSocketReady || !a.Inheritance.Gate528NoNativeRank4Projector || !a.Inheritance.Gate528TimeAssignmentBlocked || !a.Inheritance.Gate528InternalComplementBlocked || !a.Inheritance.Gate528WickHilbertDynamicsBlocked || !a.Inheritance.Gate528NoObservedDataImported || !a.Inheritance.Gate528NativeWriteBlocked || a.Inheritance.Gate528ReopenedSealedFirewalls {
		bad = append(bad, "bad inheritance")
	}
	if !a.Schema.Executed || a.Schema.RequiredRowCount != len(a.Schema.RequiredRows) || a.Schema.RequiredRowCount != 12 || !a.Schema.ProjectorMatrixRequired || !a.Schema.ProjectorIdempotencyCheck || a.Schema.ProjectorRankRequired != 4 || !a.Schema.ComplementMatrixRequired || a.Schema.ComplementRankRequired != 4 || !a.Schema.OrthogonalComplementCheck || a.Schema.ExternalSignatureRequired != "1+3" || !a.Schema.InternalAssignmentRequired || !a.Schema.SourceRequired || !a.Schema.ConventionRequired || !a.Schema.BridgeOnlyRequired || !a.Schema.NativePromotionRejected || !a.Schema.RedactedSchemaAccepted || a.Schema.AcceptedRedactedCases != 1 || a.Schema.RejectedFailClosedCases < 8 {
		bad = append(bad, "bad projection schema")
	}
	if !a.Obligations.Executed || !a.Obligations.ProjectorImportedBridgeOnly || a.Obligations.GrantsWickRotation || a.Obligations.GrantsPositiveHilbertProduct || a.Obligations.GrantsReflectionPositivity || a.Obligations.GrantsPositiveEnergyHamiltonian || a.Obligations.GrantsUnitaryRealTimeDynamics || a.Obligations.GrantsGlobalHyperbolicity || a.Obligations.GrantsInternalGaugeIdentification || !a.Obligations.RequiresSeparateWickAirlock || !a.Obligations.RequiresSeparateHilbertAirlock || !a.Obligations.RequiresSeparateUnitaryDynamicsAirlock || !a.Obligations.RequiresSeparateInternalGaugeAirlock {
		bad = append(bad, "bad Lorentzian obligations guard")
	}
	if !a.Rejection.Executed || !a.Rejection.NativeProjectorWriteRejected || !a.Rejection.Native3Plus1SpacetimeWriteRejected || !a.Rejection.NativeTimeAssignmentWriteRejected || !a.Rejection.NativeInternalComplementWriteRejected || !a.Rejection.NativeWickWriteRejected || !a.Rejection.NativeHilbertWriteRejected || !a.Rejection.NativeUnitaryDynamicsWriteRejected || a.Rejection.ComparatorExecutionPerformed {
		bad = append(bad, "bad native rejection rule")
	}
	if !a.Firewall.Executed || a.Firewall.ObservedDimensionImported || a.Firewall.ObservedConstantsImported || a.Firewall.ObservedMassesImported || a.Firewall.ObservedTopologyImported || a.Firewall.NativeProjectorWrite || a.Firewall.Native3Plus1Write || a.Firewall.NativeTimeAssignmentWrite || a.Firewall.NativeInternalComplementWrite || a.Firewall.NativeWickWrite || a.Firewall.NativeHilbertWrite || a.Firewall.NativeUnitaryDynamicsWrite || a.Firewall.NativeInternalGaugeWrite || a.Firewall.ReopenedFlavorFirewall || a.Firewall.ReopenedEWScaleFirewall || a.Firewall.ReopenedGravityFirewall || a.Firewall.ReopenedTopologyFirewall || a.Firewall.NativeRegistryWritten {
		bad = append(bad, "firewall violation")
	}
	if len(bad) > 0 {
		return fmt.Errorf(strings.Join(bad, "; "))
	}
	return nil
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("%s: inherited=%t rank44_bridge_ready=%t no_native_rank4_projector=%t time_blocked=%t internal_blocked=%t wick_hilbert_unitary_blocked=%t no_observed=%t native_blocked=%t reopens_firewalls=%t; %s", x.Verdict, x.Gate528Inherited, x.Gate528Rank44BridgeSocketReady, x.Gate528NoNativeRank4Projector, x.Gate528TimeAssignmentBlocked, x.Gate528InternalComplementBlocked, x.Gate528WickHilbertDynamicsBlocked, x.Gate528NoObservedDataImported, x.Gate528NativeWriteBlocked, x.Gate528ReopenedSealedFirewalls, x.Reason)
}

func FormatSchema(x ProjectionSchema) string {
	return fmt.Sprintf("%s: required_rows=%d projector_matrix=%t idempotency=%t projector_rank=%d complement_matrix=%t complement_rank=%d complement_orthogonality=%t external_signature=%s internal_assignment=%t source=%t convention=%t bridge_only=%t native_promotion_rejected=%t accepted_redacted_cases=%d rejected_fail_closed_cases=%d; %s", x.Verdict, x.RequiredRowCount, x.ProjectorMatrixRequired, x.ProjectorIdempotencyCheck, x.ProjectorRankRequired, x.ComplementMatrixRequired, x.ComplementRankRequired, x.OrthogonalComplementCheck, x.ExternalSignatureRequired, x.InternalAssignmentRequired, x.SourceRequired, x.ConventionRequired, x.BridgeOnlyRequired, x.NativePromotionRejected, x.AcceptedRedactedCases, x.RejectedFailClosedCases, x.Reason)
}

func FormatObligations(x LorentzianObligationsGuard) string {
	return fmt.Sprintf("%s: bridge_projector=%t grants_Wick=%t grants_Hilbert=%t grants_reflection_positivity=%t grants_positive_energy=%t grants_unitary=%t grants_global_hyperbolicity=%t grants_internal_gauge=%t separate_Wick=%t separate_Hilbert=%t separate_unitary=%t separate_internal_gauge=%t; %s", x.Verdict, x.ProjectorImportedBridgeOnly, x.GrantsWickRotation, x.GrantsPositiveHilbertProduct, x.GrantsReflectionPositivity, x.GrantsPositiveEnergyHamiltonian, x.GrantsUnitaryRealTimeDynamics, x.GrantsGlobalHyperbolicity, x.GrantsInternalGaugeIdentification, x.RequiresSeparateWickAirlock, x.RequiresSeparateHilbertAirlock, x.RequiresSeparateUnitaryDynamicsAirlock, x.RequiresSeparateInternalGaugeAirlock, x.Reason)
}

func FormatRejection(x NativeRejection) string {
	return fmt.Sprintf("%s: projector_rejected=%t spacetime_rejected=%t time_rejected=%t internal_rejected=%t Wick_rejected=%t Hilbert_rejected=%t unitary_rejected=%t comparator_executed=%t; %s", x.Verdict, x.NativeProjectorWriteRejected, x.Native3Plus1SpacetimeWriteRejected, x.NativeTimeAssignmentWriteRejected, x.NativeInternalComplementWriteRejected, x.NativeWickWriteRejected, x.NativeHilbertWriteRejected, x.NativeUnitaryDynamicsWriteRejected, x.ComparatorExecutionPerformed, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("%s: observed_dimension=%t observed_constants=%t observed_masses=%t observed_topology=%t native_projector=%t native_3plus1=%t native_time=%t native_internal=%t native_Wick=%t native_Hilbert=%t native_unitary=%t native_internal_gauge=%t reopen_flavor=%t reopen_EW=%t reopen_gravity=%t reopen_topology=%t native_write=%t; %s", x.Verdict, x.ObservedDimensionImported, x.ObservedConstantsImported, x.ObservedMassesImported, x.ObservedTopologyImported, x.NativeProjectorWrite, x.Native3Plus1Write, x.NativeTimeAssignmentWrite, x.NativeInternalComplementWrite, x.NativeWickWrite, x.NativeHilbertWrite, x.NativeUnitaryDynamicsWrite, x.NativeInternalGaugeWrite, x.ReopenedFlavorFirewall, x.ReopenedEWScaleFirewall, x.ReopenedGravityFirewall, x.ReopenedTopologyFirewall, x.NativeRegistryWritten, x.Reason)
}

func statuses() []string {
	return []string{
		StatusGate528Inherited,
		StatusProjectionAirlockDefined,
		StatusProjectorSchemaRowsEnumerated,
		StatusInternalComplementSchemaDefined,
		StatusMandatoryMetadataEnforced,
		StatusRedactedProjectionSchemaAccepted,
		StatusLorentzianObligationsGuardDefined,
		StatusNativeRejectionRuleFailClosed,
		StatusNoObservedDimensionDataImported,
		StatusFailedProjectorNativePromotionRejected,
		StatusFailedMissingMetadataRejected,
		StatusFailedProjectorDoesNotGrantWick,
		StatusFailedProjectorDoesNotGrantHilbert,
		StatusFailedProjectorDoesNotGrantUnitary,
		StatusFailedInternalGaugeNativeRejected,
		StatusFailedComparatorExecutionNotPerformed,
		StatusFirewallPreserved,
		StatusFirewallProjectionNativeWriteBlocked,
	}
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 529 Registry Audit — 3+1 Projection and Internal Complement Bridge Airlock Preflight\n\n")
	b.WriteString("## Verdict\n\n```text\n")
	for _, s := range statuses() {
		b.WriteString(s + "\n")
	}
	b.WriteString("```\n\n")
	b.WriteString("## Inherited boundary\n\nGate 529 inherits Gate 528's projector obstruction: 4+4 rank arithmetic is bridge-admissible, but no native Spin(1,7)-invariant rank-four vector projector was found.\n\n```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## Projector airlock schema\n\nThe airlock accepts only explicit, labelled bridge projector rows. Required fields are:\n\n")
	for _, r := range a.Schema.RequiredRows {
		b.WriteString(fmt.Sprintf("- `%s` — required=%t bridge_only=%t native_write=%t; %s\n", r.Name, r.Required, r.BridgeOnly, r.NativeWrite, r.Reason))
	}
	b.WriteString("\n```text\n" + FormatSchema(a.Schema) + "\n```\n\n")
	b.WriteString("## Lorentzian obligations guard\n\nProviding a 3+1 projector does not automatically provide Wick rotation, a positive Hilbert product, reflection positivity, positive energy, unitary dynamics, global hyperbolicity, or native gauge/internal identification.\n\n```text\n" + FormatObligations(a.Obligations) + "\n```\n\n")
	b.WriteString("## Native rejection rule\n\n```text\n" + FormatRejection(a.Rejection) + "\n```\n\n")
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
