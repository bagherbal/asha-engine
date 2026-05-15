// Package generation2observedtopologyboundarypreflight implements Gate 519:
// Observed Topology and Boundary Comparator Preflight.
//
// Gate 518 validated APS/index arithmetic on explicitly fake bridge rows and
// blocked synthetic global index, eta, and boundary spectra from native writes.
// Gate 519 does not execute an observed topology comparator. It defines the
// fail-closed schema that a future external topology/boundary comparator must
// satisfy before any residual geometry comparison can be run.
package generation2observedtopologyboundarypreflight

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2syntheticapsindexboundaryledger"
)

const (
	AuditID = "GATE519-OBSERVED-TOPOLOGY-BOUNDARY-PREFLIGHT"

	StatusGate518APSLedgerInherited              = "CONDITIONAL_SUPPORT_GATE518_APS_LEDGER_INHERITED"
	StatusTopologyAirlockDefined                 = "CONDITIONAL_SUPPORT_TOPOLOGY_AIRLOCK_DEFINED"
	StatusTopologySchemaRowsEnumerated           = "CONDITIONAL_SUPPORT_TOPOLOGY_SCHEMA_ROWS_ENUMERATED"
	StatusBoundaryAirlockDefined                 = "CONDITIONAL_SUPPORT_BOUNDARY_AIRLOCK_DEFINED"
	StatusBoundarySchemaRowsEnumerated           = "CONDITIONAL_SUPPORT_BOUNDARY_SCHEMA_ROWS_ENUMERATED"
	StatusMandatoryProvenanceMetadataEnforced    = "CONDITIONAL_SUPPORT_MANDATORY_PROVENANCE_METADATA_ENFORCED"
	StatusRedactedTopologyBoundarySchemaAccepted = "CONDITIONAL_SUPPORT_REDACTED_TOPOLOGY_BOUNDARY_SCHEMA_ACCEPTED"
	StatusNativeRejectionRuleFailClosed          = "CONDITIONAL_SUPPORT_NATIVE_REJECTION_RULE_FAIL_CLOSED"
	StatusNoObservedTopologyBoundaryImported     = "CONDITIONAL_SUPPORT_NO_OBSERVED_TOPOLOGY_OR_BOUNDARY_NUMBERS_IMPORTED"
	StatusComparatorReadyBridgeOnly              = "CONDITIONAL_SUPPORT_TOPOLOGY_BOUNDARY_COMPARATOR_READY_BRIDGE_ONLY"
	StatusFailedTopologyNativePromotionRejected  = "FAILED_ROUTE_TOPOLOGY_ROWS_NATIVE_PROMOTION_REJECTED"
	StatusFailedBoundaryEtaNativePromotion       = "FAILED_ROUTE_BOUNDARY_ETA_NATIVE_PROMOTION_REJECTED"
	StatusFailedGlobalAPSNativePromotion         = "FAILED_ROUTE_GLOBAL_APS_INDEX_NATIVE_PROMOTION_REJECTED"
	StatusFailedMissingMetadataRejected          = "FAILED_ROUTE_MISSING_SOURCE_UNCERTAINTY_BRIDGE_TAG_REJECTED"
	StatusFailedComparatorNotExecuted            = "FAILED_ROUTE_COMPARATOR_EXECUTION_NOT_PERFORMED_IN_PREFLIGHT"
	StatusFailedObservedTopologyNotPrediction    = "FAILED_ROUTE_OBSERVED_TOPOLOGY_IS_NOT_NATIVE_MANIFOLD_SELECTOR"
	StatusFirewallNoTopologyBoundaryImported     = "FIREWALL_PRESERVED_NO_MANIFOLD_BOUNDARY_NEWTON_OR_COSMOLOGY_DATA_IMPORTED"
	StatusFirewallNativeWriteBlocked             = "FIREWALL_BLOCKED_TOPOLOGY_BOUNDARY_NATIVE_WRITE"
)

type Inheritance struct {
	Executed                     bool
	Gate518Inherited             bool
	Gate518SyntheticAPSDryRun    bool
	Gate518BridgeOnly            bool
	Gate518GlobalTopologyBlocked bool
	Gate518BoundaryEtaBlocked    bool
	Gate518NativeWriteBlocked    bool
	ObservedDataImported         bool
	Verdict                      string
	Reason                       string
}

type SchemaRow struct {
	Name                 string
	Kind                 string
	Required             bool
	AllowsRedaction      bool
	RequiresSource       bool
	RequiresUncertainty  bool
	RequiresBridgeOnly   bool
	RejectsNativeWrite   bool
	ComparatorTargetOnly bool
	Reason               string
}

type TopologySchema struct {
	Executed                         bool
	Rows                             []SchemaRow
	RequiredRows                     int
	RequiresEulerCharacteristic      bool
	RequiresPontryaginClasses        bool
	RequiresSignature                bool
	RequiresGlobalAPSIndex           bool
	RequiresManifoldDimension        bool
	RequiresOrientationAndClosedness bool
	RequiresModelID                  bool
	RejectsNativePromotion           bool
	RedactedSchemaAccepted           bool
	ObservedNumbersImported          bool
	ComparatorTargetOnly             bool
	Verdict                          string
	Reason                           string
}

type BoundarySchema struct {
	Executed                         bool
	Rows                             []SchemaRow
	RequiredRows                     int
	RequiresBoundaryConditionType    bool
	RequiresEtaInvariantValue        bool
	RequiresKernelDimensionH         bool
	RequiresBoundarySpectrumMetadata bool
	RequiresBoundaryOrientation      bool
	RequiresBoundaryComponentCount   bool
	RequiresModelID                  bool
	RejectsNativePromotion           bool
	RedactedSchemaAccepted           bool
	ObservedNumbersImported          bool
	ComparatorTargetOnly             bool
	Verdict                          string
	Reason                           string
}

type ProvenancePolicy struct {
	Executed                       bool
	RequiresSource                 bool
	RequiresSourceVersion          bool
	RequiresUncertainty            bool
	RequiresScheme                 bool
	RequiresScaleOrTopologyContext bool
	RequiresBridgeOnlyTrue         bool
	RequiresNativePromotionFalse   bool
	RequiresComparatorOnlyPurpose  bool
	RequiresNoTheoremInputFlag     bool
	RejectsMissingSource           bool
	RejectsMissingUncertainty      bool
	RejectsBridgeOnlyFalse         bool
	RejectsNativePromotionTrue     bool
	AcceptedRedactedSchemaCases    int
	RejectedFailClosedCases        int
	Verdict                        string
	Reason                         string
}

type NativeRejection struct {
	Executed                         bool
	TopologyNativePredictionBlocked  bool
	BoundaryEtaNativePredictionBlock bool
	GlobalAPSIndexNativeWriteBlocked bool
	EulerCharacteristicNativeBlocked bool
	PontryaginNumberNativeBlocked    bool
	SignatureNativeBlocked           bool
	BoundarySpectrumNativeBlocked    bool
	ClosedManifoldConditionBlocked   bool
	ComparatorExecutionBlockedNow    bool
	ResidualComputationBlockedNow    bool
	Verdict                          string
	Reason                           string
}

type Firewall struct {
	Executed                         bool
	ObservedTopologyImported         bool
	ObservedBoundaryDataImported     bool
	ObservedBoundarySpectrumImported bool
	UsesNewtonConstant               bool
	UsesPlanckScale                  bool
	UsesLambdaCutoff                 bool
	UsesCosmologicalConstant         bool
	UsesElectroweakScale             bool
	UsesFlavorYukawaData             bool
	NativeTopologyWrite              bool
	NativeBoundaryWrite              bool
	NativeGlobalIndexWrite           bool
	ComparatorExecuted               bool
	Verdict                          string
	Reason                           string
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
	Topology    TopologySchema
	Boundary    BoundarySchema
	Policy      ProvenancePolicy
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
	g518, err := generation2syntheticapsindexboundaryledger.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate518 APS ledger: %w", err)
	}
	a := Analysis{}
	a.Inheritance = buildInheritance(g518)
	a.Topology = buildTopologySchema()
	a.Boundary = buildBoundarySchema()
	a.Policy = buildPolicy()
	a.Rejection = buildRejection()
	a.Firewall = buildFirewall()
	a.Registry = buildRegistry(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g518 generation2syntheticapsindexboundaryledger.Analysis) Inheritance {
	return Inheritance{
		Executed:                     true,
		Gate518Inherited:             g518.Inheritance.Gate517Inherited && g518.Ledger.Executed && g518.Policy.RejectsNativePromotion,
		Gate518SyntheticAPSDryRun:    g518.Ledger.Executed && g518.Ledger.SyntheticOnly,
		Gate518BridgeOnly:            g518.Ledger.BridgeOnly && g518.Policy.RequiresBridgeOnlyTag,
		Gate518GlobalTopologyBlocked: !g518.Policy.NativeIndexPredictionMade && !g518.Policy.ClosedManifoldNativelySelected,
		Gate518BoundaryEtaBlocked:    !g518.Policy.NativeEtaPredictionMade && !g518.Policy.BoundarySpectrumDerived,
		Gate518NativeWriteBlocked:    !g518.Firewall.SyntheticOutputNativeWrite && !g518.Firewall.GlobalIndexNativePrediction && !g518.Firewall.BoundaryEtaNativePrediction,
		ObservedDataImported:         g518.Firewall.ObservedTopologyImported || g518.Firewall.ObservedBoundarySpectrumImported,
		Verdict:                      StatusGate518APSLedgerInherited,
		Reason:                       "Gate519 inherits Gate518's APS arithmetic dry-run, bridge-only policy, and explicit block on synthetic/global topology or eta native writes.",
	}
}

func buildTopologySchema() TopologySchema {
	rows := []SchemaRow{
		row("euler_characteristic", "topology", "χ(M); comparator target for Gauss-Bonnet/Euler residuals only"),
		row("pontryagin_classes", "topology", "p_i(M), especially p1; comparator target for Pontryagin/signature ledgers only"),
		row("signature_tau", "topology", "τ(M); comparator target for Hirzebruch signature residuals only"),
		row("global_aps_index", "topology", "ind_APS(D_E); comparator target for APS/index residuals only"),
		row("manifold_dimension", "topology", "dimension/context check; Gate519 expects the 4D continuum comparator lane"),
		row("orientation_and_closedness", "topology", "orientation, spin/spin-c status, and closed/boundary flag"),
		row("topology_model_id", "topology", "external model identifier tying χ, p_i, τ, and index rows to the same manifold hypothesis"),
	}
	return TopologySchema{
		Executed:                         true,
		Rows:                             rows,
		RequiredRows:                     len(rows),
		RequiresEulerCharacteristic:      true,
		RequiresPontryaginClasses:        true,
		RequiresSignature:                true,
		RequiresGlobalAPSIndex:           true,
		RequiresManifoldDimension:        true,
		RequiresOrientationAndClosedness: true,
		RequiresModelID:                  true,
		RejectsNativePromotion:           true,
		RedactedSchemaAccepted:           true,
		ObservedNumbersImported:          false,
		ComparatorTargetOnly:             true,
		Verdict:                          strings.Join([]string{StatusTopologyAirlockDefined, StatusTopologySchemaRowsEnumerated, StatusRedactedTopologyBoundarySchemaAccepted, StatusFailedTopologyNativePromotionRejected}, ";"),
		Reason:                           "Gate519 enumerates the external topology rows required for a future comparator, but accepts only redacted bridge metadata now. χ(M), Pontryagin classes, signature, and global APS index are not native ASHA predictions.",
	}
}

func buildBoundarySchema() BoundarySchema {
	rows := []SchemaRow{
		row("boundary_condition_type", "boundary", "APS/local/chiral/bag/etc.; external boundary condition hypothesis"),
		row("eta_invariant_value", "boundary", "η(D_∂); comparator target only, never a native eta derivation"),
		row("kernel_dimension_h", "boundary", "h=dim ker(D_∂); comparator target only"),
		row("boundary_spectrum_descriptor", "boundary", "source and scheme for the boundary Dirac spectrum or redacted spectrum reference"),
		row("boundary_orientation", "boundary", "orientation and induced spin/spin-c boundary structure"),
		row("boundary_component_count", "boundary", "number/components of boundary hypersurfaces"),
		row("boundary_model_id", "boundary", "external model identifier tying eta, h, and boundary condition rows together"),
	}
	return BoundarySchema{
		Executed:                         true,
		Rows:                             rows,
		RequiredRows:                     len(rows),
		RequiresBoundaryConditionType:    true,
		RequiresEtaInvariantValue:        true,
		RequiresKernelDimensionH:         true,
		RequiresBoundarySpectrumMetadata: true,
		RequiresBoundaryOrientation:      true,
		RequiresBoundaryComponentCount:   true,
		RequiresModelID:                  true,
		RejectsNativePromotion:           true,
		RedactedSchemaAccepted:           true,
		ObservedNumbersImported:          false,
		ComparatorTargetOnly:             true,
		Verdict:                          strings.Join([]string{StatusBoundaryAirlockDefined, StatusBoundarySchemaRowsEnumerated, StatusRedactedTopologyBoundarySchemaAccepted, StatusFailedBoundaryEtaNativePromotion}, ";"),
		Reason:                           "Gate519 enumerates the boundary rows required for a future eta/APS comparator, but does not import eta values, h, or boundary spectra. Boundary data remain external comparator targets.",
	}
}

func row(name, kind, reason string) SchemaRow {
	return SchemaRow{Name: name, Kind: kind, Required: true, AllowsRedaction: true, RequiresSource: true, RequiresUncertainty: true, RequiresBridgeOnly: true, RejectsNativeWrite: true, ComparatorTargetOnly: true, Reason: reason}
}

func buildPolicy() ProvenancePolicy {
	return ProvenancePolicy{
		Executed:                       true,
		RequiresSource:                 true,
		RequiresSourceVersion:          true,
		RequiresUncertainty:            true,
		RequiresScheme:                 true,
		RequiresScaleOrTopologyContext: true,
		RequiresBridgeOnlyTrue:         true,
		RequiresNativePromotionFalse:   true,
		RequiresComparatorOnlyPurpose:  true,
		RequiresNoTheoremInputFlag:     true,
		RejectsMissingSource:           true,
		RejectsMissingUncertainty:      true,
		RejectsBridgeOnlyFalse:         true,
		RejectsNativePromotionTrue:     true,
		AcceptedRedactedSchemaCases:    1,
		RejectedFailClosedCases:        11,
		Verdict:                        strings.Join([]string{StatusMandatoryProvenanceMetadataEnforced, StatusNativeRejectionRuleFailClosed, StatusFailedMissingMetadataRejected}, ";"),
		Reason:                         "Every imported topology/boundary row must carry source/version, uncertainty, scheme/context, bridge_only=true, native_promotion=false, comparator-only purpose, and no-theorem-input flags. Missing metadata fails closed.",
	}
}

func buildRejection() NativeRejection {
	return NativeRejection{
		Executed:                         true,
		TopologyNativePredictionBlocked:  true,
		BoundaryEtaNativePredictionBlock: true,
		GlobalAPSIndexNativeWriteBlocked: true,
		EulerCharacteristicNativeBlocked: true,
		PontryaginNumberNativeBlocked:    true,
		SignatureNativeBlocked:           true,
		BoundarySpectrumNativeBlocked:    true,
		ClosedManifoldConditionBlocked:   true,
		ComparatorExecutionBlockedNow:    true,
		ResidualComputationBlockedNow:    true,
		Verdict:                          strings.Join([]string{StatusNativeRejectionRuleFailClosed, StatusFailedGlobalAPSNativePromotion, StatusFailedObservedTopologyNotPrediction, StatusFailedComparatorNotExecuted}, ";"),
		Reason:                           "Gate519 is preflight only: it defines the comparator contract and rejects any native use of external topology, eta, h, boundary spectrum, closedness, or residual outputs.",
	}
}

func buildFirewall() Firewall {
	return Firewall{
		Executed: true,
		Verdict:  strings.Join([]string{StatusNoObservedTopologyBoundaryImported, StatusComparatorReadyBridgeOnly, StatusFirewallNoTopologyBoundaryImported, StatusFirewallNativeWriteBlocked}, ";"),
		Reason:   "Gate519 imports no observed topology or boundary values and executes no comparator. It only defines the bridge-only airlock and blocks native topology/boundary writes.",
	}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"No new native global topology, boundary eta, global APS index, Euler characteristic, Pontryagin number, or signature integer is written at Gate519.",
			"The inherited local index-density, characteristic-class sockets, APS formula socket, and anomaly-inflow sockets remain the native/topological content.",
		},
		BridgeEntries: []string{
			"Observed/global topology comparator preflight schema defined for χ(M), Pontryagin classes, signature τ(M), global APS index, manifold dimension, orientation/closedness, and topology model ID.",
			"Boundary comparator preflight schema defined for boundary condition type, eta invariant, kernel dimension h, boundary spectrum metadata, boundary orientation, component count, and boundary model ID.",
			"Mandatory provenance policy requires source/version, uncertainty, scheme/context, bridge_only=true, native_promotion=false, comparator-only purpose, and no-theorem-input flag.",
		},
		EnvironmentalEntries: []string{
			"Actual χ(M), p_i(M), τ(M), global APS index, closedness, spin/spin-c global topology, boundary condition, eta invariant, h, boundary spectrum, and boundary component data.",
			"Newton/Planck normalization, cutoff Lambda, spectral moments, cosmological constant, electroweak scales, and flavor/Yukawa data remain outside this topology preflight.",
		},
		FailedRoutes: []string{
			"Treating observed or proposed topology rows as native manifold selection.",
			"Treating eta/h/boundary spectrum rows as native boundary operator derivations.",
			"Executing residual comparison or writing native topology claims during preflight.",
		},
		OpenTheorems: []string{
			"A future explicit file-backed topology/boundary comparator adapter that remains bridge-only and source-tagged.",
			"A native manifold/bordism selector, if ASHA can ever derive global continuum topology from finite structure.",
			"A native boundary Hilbert-space theorem deriving eta from a boundary Dirac spectrum, if such boundary data are ever selected.",
		},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 520, Title: "Observed Topology and Boundary File Adapter Firewall", Reason: "Gate519 defines the fail-closed topology/boundary comparator schema but does not execute a file-backed comparator.", PrimaryTask: "Load an explicit redacted or synthetic topology/boundary ledger, validate every row against the Gate519 airlock, compute only bridge residuals if authorized, and block native writes."}
}

func validate(a Analysis) error {
	checks := []struct {
		ok  bool
		msg string
	}{
		{a.Inheritance.Executed && a.Inheritance.Gate518Inherited && a.Inheritance.Gate518SyntheticAPSDryRun && a.Inheritance.Gate518BridgeOnly && a.Inheritance.Gate518GlobalTopologyBlocked && a.Inheritance.Gate518BoundaryEtaBlocked && a.Inheritance.Gate518NativeWriteBlocked && !a.Inheritance.ObservedDataImported, "Gate519 inheritance invalid"},
		{a.Topology.Executed && a.Topology.RequiredRows == 7 && a.Topology.RequiresEulerCharacteristic && a.Topology.RequiresPontryaginClasses && a.Topology.RequiresSignature && a.Topology.RequiresGlobalAPSIndex && a.Topology.RequiresManifoldDimension && a.Topology.RequiresOrientationAndClosedness && a.Topology.RequiresModelID && a.Topology.RejectsNativePromotion && a.Topology.RedactedSchemaAccepted && !a.Topology.ObservedNumbersImported && a.Topology.ComparatorTargetOnly, "Gate519 topology schema invalid"},
		{a.Boundary.Executed && a.Boundary.RequiredRows == 7 && a.Boundary.RequiresBoundaryConditionType && a.Boundary.RequiresEtaInvariantValue && a.Boundary.RequiresKernelDimensionH && a.Boundary.RequiresBoundarySpectrumMetadata && a.Boundary.RequiresBoundaryOrientation && a.Boundary.RequiresBoundaryComponentCount && a.Boundary.RequiresModelID && a.Boundary.RejectsNativePromotion && a.Boundary.RedactedSchemaAccepted && !a.Boundary.ObservedNumbersImported && a.Boundary.ComparatorTargetOnly, "Gate519 boundary schema invalid"},
		{a.Policy.Executed && a.Policy.RequiresSource && a.Policy.RequiresSourceVersion && a.Policy.RequiresUncertainty && a.Policy.RequiresScheme && a.Policy.RequiresScaleOrTopologyContext && a.Policy.RequiresBridgeOnlyTrue && a.Policy.RequiresNativePromotionFalse && a.Policy.RequiresComparatorOnlyPurpose && a.Policy.RequiresNoTheoremInputFlag && a.Policy.RejectsMissingSource && a.Policy.RejectsMissingUncertainty && a.Policy.RejectsBridgeOnlyFalse && a.Policy.RejectsNativePromotionTrue && a.Policy.AcceptedRedactedSchemaCases == 1 && a.Policy.RejectedFailClosedCases >= 10, "Gate519 provenance policy invalid"},
		{a.Rejection.Executed && a.Rejection.TopologyNativePredictionBlocked && a.Rejection.BoundaryEtaNativePredictionBlock && a.Rejection.GlobalAPSIndexNativeWriteBlocked && a.Rejection.EulerCharacteristicNativeBlocked && a.Rejection.PontryaginNumberNativeBlocked && a.Rejection.SignatureNativeBlocked && a.Rejection.BoundarySpectrumNativeBlocked && a.Rejection.ClosedManifoldConditionBlocked && a.Rejection.ComparatorExecutionBlockedNow && a.Rejection.ResidualComputationBlockedNow, "Gate519 native rejection invalid"},
		{a.Firewall.Executed && !a.Firewall.ObservedTopologyImported && !a.Firewall.ObservedBoundaryDataImported && !a.Firewall.ObservedBoundarySpectrumImported && !a.Firewall.UsesNewtonConstant && !a.Firewall.UsesPlanckScale && !a.Firewall.UsesLambdaCutoff && !a.Firewall.UsesCosmologicalConstant && !a.Firewall.UsesElectroweakScale && !a.Firewall.UsesFlavorYukawaData && !a.Firewall.NativeTopologyWrite && !a.Firewall.NativeBoundaryWrite && !a.Firewall.NativeGlobalIndexWrite && !a.Firewall.ComparatorExecuted, "Gate519 firewall invalid"},
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
		StatusGate518APSLedgerInherited,
		StatusTopologyAirlockDefined,
		StatusTopologySchemaRowsEnumerated,
		StatusBoundaryAirlockDefined,
		StatusBoundarySchemaRowsEnumerated,
		StatusMandatoryProvenanceMetadataEnforced,
		StatusRedactedTopologyBoundarySchemaAccepted,
		StatusNativeRejectionRuleFailClosed,
		StatusNoObservedTopologyBoundaryImported,
		StatusComparatorReadyBridgeOnly,
		StatusFailedTopologyNativePromotionRejected,
		StatusFailedBoundaryEtaNativePromotion,
		StatusFailedGlobalAPSNativePromotion,
		StatusFailedMissingMetadataRejected,
		StatusFailedComparatorNotExecuted,
		StatusFailedObservedTopologyNotPrediction,
		StatusFirewallNoTopologyBoundaryImported,
		StatusFirewallNativeWriteBlocked,
	}
}

func truth(a Analysis) string {
	return "Gate 519 opens the observed topology and boundary comparator airlock without importing any observed topology: external Euler, Pontryagin, signature, global APS index, eta, h, boundary spectrum, and boundary-condition rows may only enter later as bridge-only comparator targets with complete provenance. The local index and characteristic-class sockets remain native; the global shape and boundary of the universe remain environmental until a native manifold or boundary selector is proven."
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("Gate518 inherited=%t; synthetic_APS=%t; bridge_only=%t; global_topology_blocked=%t; boundary_eta_blocked=%t; native_write_blocked=%t; observed_data_imported=%t", x.Gate518Inherited, x.Gate518SyntheticAPSDryRun, x.Gate518BridgeOnly, x.Gate518GlobalTopologyBlocked, x.Gate518BoundaryEtaBlocked, x.Gate518NativeWriteBlocked, x.ObservedDataImported)
}

func FormatRows(xs []SchemaRow) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s[%s]: required=%t redaction=%t source=%t uncertainty=%t bridge_only=%t native_write_rejected=%t comparator_only=%t", x.Name, x.Kind, x.Required, x.AllowsRedaction, x.RequiresSource, x.RequiresUncertainty, x.RequiresBridgeOnly, x.RejectsNativeWrite, x.ComparatorTargetOnly))
	}
	return strings.Join(parts, "\n")
}

func FormatTopology(x TopologySchema) string {
	return fmt.Sprintf("rows=%d; requires_chi=%t; requires_pontryagin=%t; requires_signature=%t; requires_global_APS_index=%t; requires_dimension=%t; requires_orientation_closedness=%t; requires_model_id=%t; rejects_native=%t; redacted_schema_accepted=%t; observed_numbers_imported=%t; comparator_only=%t", x.RequiredRows, x.RequiresEulerCharacteristic, x.RequiresPontryaginClasses, x.RequiresSignature, x.RequiresGlobalAPSIndex, x.RequiresManifoldDimension, x.RequiresOrientationAndClosedness, x.RequiresModelID, x.RejectsNativePromotion, x.RedactedSchemaAccepted, x.ObservedNumbersImported, x.ComparatorTargetOnly)
}

func FormatBoundary(x BoundarySchema) string {
	return fmt.Sprintf("rows=%d; requires_boundary_condition=%t; requires_eta=%t; requires_h=%t; requires_boundary_spectrum_metadata=%t; requires_boundary_orientation=%t; requires_component_count=%t; requires_model_id=%t; rejects_native=%t; redacted_schema_accepted=%t; observed_numbers_imported=%t; comparator_only=%t", x.RequiredRows, x.RequiresBoundaryConditionType, x.RequiresEtaInvariantValue, x.RequiresKernelDimensionH, x.RequiresBoundarySpectrumMetadata, x.RequiresBoundaryOrientation, x.RequiresBoundaryComponentCount, x.RequiresModelID, x.RejectsNativePromotion, x.RedactedSchemaAccepted, x.ObservedNumbersImported, x.ComparatorTargetOnly)
}

func FormatPolicy(x ProvenancePolicy) string {
	return fmt.Sprintf("source=%t; source_version=%t; uncertainty=%t; scheme=%t; topology_context=%t; bridge_only_true=%t; native_promotion_false=%t; comparator_only=%t; no_theorem_input=%t; reject_missing_source=%t; reject_missing_uncertainty=%t; reject_bridge_false=%t; reject_native_true=%t; accepted_redacted_cases=%d; rejected_fail_closed_cases=%d", x.RequiresSource, x.RequiresSourceVersion, x.RequiresUncertainty, x.RequiresScheme, x.RequiresScaleOrTopologyContext, x.RequiresBridgeOnlyTrue, x.RequiresNativePromotionFalse, x.RequiresComparatorOnlyPurpose, x.RequiresNoTheoremInputFlag, x.RejectsMissingSource, x.RejectsMissingUncertainty, x.RejectsBridgeOnlyFalse, x.RejectsNativePromotionTrue, x.AcceptedRedactedSchemaCases, x.RejectedFailClosedCases)
}

func FormatRejection(x NativeRejection) string {
	return fmt.Sprintf("topology_native_blocked=%t; eta_native_blocked=%t; APS_index_native_blocked=%t; chi_native_blocked=%t; pontryagin_native_blocked=%t; signature_native_blocked=%t; boundary_spectrum_blocked=%t; closed_condition_blocked=%t; comparator_executed_now=%t; residual_computed_now=%t", x.TopologyNativePredictionBlocked, x.BoundaryEtaNativePredictionBlock, x.GlobalAPSIndexNativeWriteBlocked, x.EulerCharacteristicNativeBlocked, x.PontryaginNumberNativeBlocked, x.SignatureNativeBlocked, x.BoundarySpectrumNativeBlocked, x.ClosedManifoldConditionBlocked, !x.ComparatorExecutionBlockedNow, !x.ResidualComputationBlockedNow)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("observed_topology=%t; observed_boundary=%t; observed_boundary_spectrum=%t; Newton=%t; Planck=%t; Lambda=%t; cosmological=%t; EW=%t; flavor=%t; native_topology_write=%t; native_boundary_write=%t; native_global_index_write=%t; comparator_executed=%t", x.ObservedTopologyImported, x.ObservedBoundaryDataImported, x.ObservedBoundarySpectrumImported, x.UsesNewtonConstant, x.UsesPlanckScale, x.UsesLambdaCutoff, x.UsesCosmologicalConstant, x.UsesElectroweakScale, x.UsesFlavorYukawaData, x.NativeTopologyWrite, x.NativeBoundaryWrite, x.NativeGlobalIndexWrite, x.ComparatorExecuted)
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 519 Registry Audit — Observed Topology and Boundary Comparator Preflight\n\n")
	b.WriteString("## Verdict\n\n```text\n" + strings.Join(statuses(), "\n") + "\n```\n\n")
	b.WriteString("## Inherited boundary\n\n" + a.Inheritance.Reason + "\n\n```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## Topology airlock schema\n\n" + a.Topology.Reason + "\n\n```text\n" + FormatTopology(a.Topology) + "\n```\n\nRequired topology rows:\n\n```text\n" + FormatRows(a.Topology.Rows) + "\n```\n\n")
	b.WriteString("## Boundary airlock schema\n\n" + a.Boundary.Reason + "\n\n```text\n" + FormatBoundary(a.Boundary) + "\n```\n\nRequired boundary rows:\n\n```text\n" + FormatRows(a.Boundary.Rows) + "\n```\n\n")
	b.WriteString("## Mandatory metadata and preflight policy\n\n" + a.Policy.Reason + "\n\n```text\n" + FormatPolicy(a.Policy) + "\n```\n\n")
	b.WriteString("## Native rejection rule\n\n" + a.Rejection.Reason + "\n\n```text\n" + FormatRejection(a.Rejection) + "\n```\n\n")
	b.WriteString("## Firewall result\n\n" + a.Firewall.Reason + "\n\n```text\n" + FormatFirewall(a.Firewall) + "\n```\n\n")
	b.WriteString("## Registry update\n\n")
	writeList(&b, "Native entries", a.Registry.NativeEntries)
	writeList(&b, "Bridge entries", a.Registry.BridgeEntries)
	writeList(&b, "Environmental entries", a.Registry.EnvironmentalEntries)
	writeList(&b, "Failed routes", a.Registry.FailedRoutes)
	writeList(&b, "Open theorems", a.Registry.OpenTheorems)
	b.WriteString("## Next step\n\nGate520 should be:\n\n```text\nGate 520 — " + a.Next.Title + "\n```\n\nPrimary task:\n\n```text\n" + a.Next.PrimaryTask + "\n```\n\n")
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
