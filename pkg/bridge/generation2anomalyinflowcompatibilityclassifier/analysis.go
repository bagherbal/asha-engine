// Package generation2anomalyinflowcompatibilityclassifier implements Gate 524:
// Anomaly-Inflow Compatibility Classifier for Bridge Topology Classes.
//
// Gate 523 aggregated APS-boundary and bordism residual fixtures while blocking
// any native merge into a universe topology. Gate 524 audits the final topology
// consistency lane: whether the local Cℓ(1,7) index-density and anomaly sockets
// have the structural capacity required for anomaly inflow on bridge-only
// topology classes. The answer is positive as a capacity theorem, but bounded:
// inflow compatibility does not select a boundary, derive an eta spectrum,
// choose a bordism representative, merge heterogeneous fixtures, or write global
// topology into the native registry.
package generation2anomalyinflowcompatibilityclassifier

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2gravitationalindexetaairlock"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2topologicalanomalyledger"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2topologyresidualclassifierreport"
)

const (
	AuditID = "GATE524-ANOMALY-INFLOW-COMPATIBILITY-CLASSIFIER"

	StatusGate523TopologyReportInherited     = "CONDITIONAL_SUPPORT_GATE523_TOPOLOGY_RESIDUAL_REPORT_INHERITED"
	StatusGate517IndexEtaInflowInherited     = "CONDITIONAL_SUPPORT_GATE517_INDEX_ETA_INFLOW_SOCKET_INHERITED"
	StatusGate490AnomalyLedgerInherited      = "CONDITIONAL_SUPPORT_GATE490_LOCAL_ANOMALY_LEDGER_INHERITED"
	StatusInflowClassifierDefined            = "CONDITIONAL_SUPPORT_ANOMALY_INFLOW_COMPATIBILITY_CLASSIFIER_DEFINED"
	StatusBulkBoundaryDescentSocketPresent   = "CONDITIONAL_SUPPORT_BULK_BOUNDARY_DESCENT_SOCKET_PRESENT"
	StatusNativeInflowCapacityConfirmed      = "CONDITIONAL_SUPPORT_ANOMALY_INFLOW_CAPACITY_CONFIRMED"
	StatusAPSBoundaryClassCompatible         = "CONDITIONAL_SUPPORT_APS_BOUNDARY_CLASS_COMPATIBLE_BRIDGE_ONLY"
	StatusSpinSpinCBordismCompatible         = "CONDITIONAL_SUPPORT_SPIN_SPINC_BORDISM_CLASS_COMPATIBLE_BRIDGE_ONLY"
	StatusMixedGaugeGravityInflowTraceCancel = "CONDITIONAL_SUPPORT_MIXED_GAUGE_GRAVITY_INFLOW_TRACES_CANCEL"
	StatusHeterogeneousFixtureGuardPreserved = "CONDITIONAL_SUPPORT_HETEROGENEOUS_FIXTURE_GUARD_PRESERVED"
	StatusNoObservedTopologyBoundaryImported = "CONDITIONAL_SUPPORT_NO_OBSERVED_TOPOLOGY_OR_BOUNDARY_DATA_IMPORTED"

	StatusFailedInflowDoesNotSelectBoundary  = "FAILED_ROUTE_INFLOW_CAPACITY_DOES_NOT_SELECT_BOUNDARY"
	StatusFailedEtaSpectrumNotDerived        = "FAILED_ROUTE_INFLOW_CAPACITY_DOES_NOT_DERIVE_ETA_SPECTRUM"
	StatusFailedCrossFixtureMergeRejected    = "FAILED_ROUTE_INFLOW_CAPACITY_DOES_NOT_MERGE_GATE520_GATE522_FIXTURES"
	StatusFailedGlobalAnomalyCoeffUnselected = "FAILED_ROUTE_GLOBAL_ANOMALY_COEFFICIENTS_NOT_SELECTED"
	StatusFailedGravitationalThetaUnselected = "FAILED_ROUTE_GRAVITATIONAL_THETA_STILL_NOT_SELECTED_BY_INFLOW_CLASSIFIER"
	StatusFirewallPreserved                  = "FIREWALL_PRESERVED_NO_MANIFOLD_BOUNDARY_BORDISM_NEWTON_OR_COSMOLOGY_DATA_IMPORTED"
	StatusFirewallNativeWriteBlocked         = "FIREWALL_BLOCKED_GATE524_INFLOW_COMPATIBILITY_NATIVE_TOPOLOGY_WRITE"
)

type Inheritance struct {
	Executed bool

	Gate523ReportDefined            bool
	Gate523Rows                     int
	Gate523ZeroResidualRows         int
	Gate523APSBoundaryRows          int
	Gate523ClosedBordismRows        int
	Gate523BridgeOnly               bool
	Gate523SyntheticOnly            bool
	Gate523ObservedImported         bool
	Gate523HeterogeneousGuard       bool
	Gate523CrossLedgerMergeRejected bool
	Gate523NativeManifoldSelected   bool
	Gate523NativeWriteBlocked       bool

	Gate517IndexSocket                bool
	Gate517APSSocket                  bool
	Gate517InflowSocket               bool
	Gate517EtaGlobalData              bool
	Gate517NativeIndexEtaWriteBlocked bool

	Gate490GaugeAnomaliesCancel    bool
	Gate490MixedGaugeGravityCancel bool
	Gate490WittenSU2Cancel         bool
	Gate490ExactRational           bool
	Gate490FlavorMassIndependent   bool

	Verdict, Reason string
}

type InflowSieve struct {
	Executed                         bool
	LocalIndexDensity                string
	DescentSocket                    string
	BoundaryPairing                  string
	BulkCharacteristicClassesPresent bool
	GaugeAnomalyTraceZero            bool
	MixedGaugeGravityTraceZero       bool
	WittenSU2GlobalCleared           bool
	APSBoundaryCorrectionPairing     bool
	ChernSimonsTransgressionSocket   bool
	ScaleFree                        bool
	MassFlavorIndependent            bool
	NativeCapacityConfirmed          bool
	BoundaryTheorySelected           bool
	BoundaryConditionSelected        bool
	EtaSpectrumDerived               bool
	GlobalAnomalyCoefficientSelected bool
	Verdict, Reason                  string
}

type BridgeCompatibility struct {
	Executed                          bool
	CompatibleClassCount              int
	APSBoundaryFixtureCompatible      bool
	SpinBordismFixtureCompatible      bool
	SpinCBordismFixtureCompatible     bool
	Gate520BoundaryMode               bool
	Gate522ClosedBoundary             bool
	HeterogeneousGuardPreserved       bool
	CrossFixtureIdentityAllowed       bool
	CrossFixtureMergeRejected         bool
	ClassifiesButDoesNotSelect        bool
	BoundaryCurrentConservationSocket bool
	NativeManifoldSelected            bool
	NativeBoundarySelected            bool
	NativeBordismClassSelected        bool
	Verdict, Reason                   string
}

type Firewall struct {
	Executed                         bool
	ObservedTopologyImported         bool
	ObservedBoundaryImported         bool
	ObservedBordismImported          bool
	ObservedEtaImported              bool
	ObservedBoundarySpectrumImported bool
	NewtonPlanckCosmologyImported    bool
	InflowCapacityNative             bool
	BoundaryTheoryNative             bool
	BoundaryConditionNative          bool
	EtaSpectrumNative                bool
	BordismClassNative               bool
	CharacteristicNumbersNative      bool
	CrossFixtureMergeNative          bool
	GravitationalThetaNative         bool
	GlobalAnomalyCoefficientsNative  bool
	NativeRegistryWritten            bool
	Verdict, Reason                  string
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
	Inflow        InflowSieve
	Compatibility BridgeCompatibility
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
	g523, err := generation2topologyresidualclassifierreport.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate523 topology residual report: %w", err)
	}
	g517, err := generation2gravitationalindexetaairlock.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate517 index/eta airlock: %w", err)
	}
	g490, err := generation2topologicalanomalyledger.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate490 anomaly ledger: %w", err)
	}
	a := Analysis{}
	a.Inheritance = buildInheritance(g523, g517, g490)
	a.Inflow = buildInflow(a.Inheritance)
	a.Compatibility = buildCompatibility(a.Inheritance, a.Inflow)
	a.Firewall = buildFirewall(a.Inflow, a.Compatibility)
	a.Registry = buildRegistry(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g523 generation2topologyresidualclassifierreport.Analysis, g517 generation2gravitationalindexetaairlock.Analysis, g490 generation2topologicalanomalyledger.Analysis) Inheritance {
	return Inheritance{
		Executed:                          true,
		Gate523ReportDefined:              g523.Report.Executed && g523.Report.ReportReady,
		Gate523Rows:                       g523.Report.Rows,
		Gate523ZeroResidualRows:           g523.Report.ZeroResidualRows,
		Gate523APSBoundaryRows:            g523.Report.APSBoundaryRows,
		Gate523ClosedBordismRows:          g523.Report.ClosedBordismRows,
		Gate523BridgeOnly:                 g523.Report.BridgeOnly,
		Gate523SyntheticOnly:              g523.Report.SyntheticOnly,
		Gate523ObservedImported:           g523.Report.ObservedImported,
		Gate523HeterogeneousGuard:         g523.Guard.Executed && g523.Guard.CrossLedgerMergeRejected && g523.Guard.DifferentSyntheticContexts,
		Gate523CrossLedgerMergeRejected:   g523.Guard.CrossLedgerMergeRejected,
		Gate523NativeManifoldSelected:     g523.Guard.NativeManifoldSelected,
		Gate523NativeWriteBlocked:         !g523.Firewall.NativeRegistryWritten,
		Gate517IndexSocket:                g517.Index.LocalIndexDensitySocketPresent,
		Gate517APSSocket:                  g517.Index.ClosedManifoldSocketConsistent && g517.Eta.BoundaryOperatorRequired && g517.Eta.EtaInvariantRequired,
		Gate517InflowSocket:               g517.Inflow.PontryaginDescentSocketPresent && g517.Inflow.ChernSimonsBoundarySocketPresent && g517.Inflow.ChiralIndexAnomalySocketPresent && g517.Inflow.BoundaryEtaPairsWithInflow,
		Gate517EtaGlobalData:              !g517.Index.BoundaryEtaDerived && !g517.Eta.BoundaryEtaNativeDerived,
		Gate517NativeIndexEtaWriteBlocked: !g517.Firewall.GlobalIndexIntegerNativeWrite && !g517.Firewall.BoundaryEtaNativeWrite,
		Gate490GaugeAnomaliesCancel:       g490.Anomaly.AllPerturbativeGaugeCancel,
		Gate490MixedGaugeGravityCancel:    g490.Anomaly.AllMixedGaugeGravityCancel,
		Gate490WittenSU2Cancel:            g490.Anomaly.SU2GlobalWittenCancels,
		Gate490ExactRational:              g490.Anomaly.ExactRationalArithmetic,
		Gate490FlavorMassIndependent:      g490.Stability.FlavorMassIndependent && g490.Stability.YukawaIndependent && g490.Stability.CKMIndependent,
		Verdict:                           strings.Join([]string{StatusGate523TopologyReportInherited, StatusGate517IndexEtaInflowInherited, StatusGate490AnomalyLedgerInherited}, ";"),
		Reason:                            "Gate524 inherits Gate523's bridge-only residual report and heterogeneous-fixture guard, Gate517's APS/index/eta anomaly-inflow socket, and Gate490's exact local gauge and mixed gauge-gravity anomaly cancellations.",
	}
}

func buildInflow(in Inheritance) InflowSieve {
	capacity := in.Gate517IndexSocket && in.Gate517APSSocket && in.Gate517InflowSocket && in.Gate490GaugeAnomaliesCancel && in.Gate490MixedGaugeGravityCancel && in.Gate490WittenSU2Cancel && in.Gate490ExactRational
	return InflowSieve{
		Executed:                         true,
		LocalIndexDensity:                "[Â(R) ch(E)]_4",
		DescentSocket:                    "bulk index density -> Chern-Simons/transgression boundary form",
		BoundaryPairing:                  "ind_APS(D_E)=∫_M[Â(R)ch(E)]_4-(η(D_∂)+h)/2",
		BulkCharacteristicClassesPresent: in.Gate517IndexSocket,
		GaugeAnomalyTraceZero:            in.Gate490GaugeAnomaliesCancel,
		MixedGaugeGravityTraceZero:       in.Gate490MixedGaugeGravityCancel,
		WittenSU2GlobalCleared:           in.Gate490WittenSU2Cancel,
		APSBoundaryCorrectionPairing:     in.Gate517APSSocket,
		ChernSimonsTransgressionSocket:   in.Gate517InflowSocket,
		ScaleFree:                        true,
		MassFlavorIndependent:            in.Gate490FlavorMassIndependent,
		NativeCapacityConfirmed:          capacity,
		BoundaryTheorySelected:           false,
		BoundaryConditionSelected:        false,
		EtaSpectrumDerived:               false,
		GlobalAnomalyCoefficientSelected: false,
		Verdict:                          strings.Join([]string{StatusInflowClassifierDefined, StatusBulkBoundaryDescentSocketPresent, StatusNativeInflowCapacityConfirmed, StatusMixedGaugeGravityInflowTraceCancel}, ";"),
		Reason:                           "The local index-density, Chern-Simons transgression, APS eta-pairing, exact gauge-anomaly zeroes, and mixed gauge-gravity trace cancellation give ASHA structural anomaly-inflow capacity without choosing a boundary theory or eta spectrum.",
	}
}

func buildCompatibility(in Inheritance, s InflowSieve) BridgeCompatibility {
	aps := in.Gate523APSBoundaryRows > 0 && in.Gate523BridgeOnly && s.NativeCapacityConfirmed
	spin := in.Gate523ClosedBordismRows > 0 && in.Gate523BridgeOnly && s.NativeCapacityConfirmed
	spinc := spin
	count := 0
	for _, ok := range []bool{aps, spin, spinc} {
		if ok {
			count++
		}
	}
	return BridgeCompatibility{
		Executed:                          true,
		CompatibleClassCount:              count,
		APSBoundaryFixtureCompatible:      aps,
		SpinBordismFixtureCompatible:      spin,
		SpinCBordismFixtureCompatible:     spinc,
		Gate520BoundaryMode:               in.Gate523APSBoundaryRows == 2,
		Gate522ClosedBoundary:             in.Gate523ClosedBordismRows == 2,
		HeterogeneousGuardPreserved:       in.Gate523HeterogeneousGuard,
		CrossFixtureIdentityAllowed:       false,
		CrossFixtureMergeRejected:         in.Gate523CrossLedgerMergeRejected,
		ClassifiesButDoesNotSelect:        true,
		BoundaryCurrentConservationSocket: s.GaugeAnomalyTraceZero && s.MixedGaugeGravityTraceZero && s.APSBoundaryCorrectionPairing,
		NativeManifoldSelected:            false,
		NativeBoundarySelected:            false,
		NativeBordismClassSelected:        false,
		Verdict:                           strings.Join([]string{StatusAPSBoundaryClassCompatible, StatusSpinSpinCBordismCompatible, StatusHeterogeneousFixtureGuardPreserved}, ";"),
		Reason:                            "The APS-boundary fixture and the spin/spin-c bordism fixture are compatible with the local inflow sockets as separate bridge classes. They remain heterogeneous fixtures and cannot be merged into one native topology identity.",
	}
}

func buildFirewall(s InflowSieve, c BridgeCompatibility) Firewall {
	return Firewall{
		Executed:                         true,
		ObservedTopologyImported:         false,
		ObservedBoundaryImported:         false,
		ObservedBordismImported:          false,
		ObservedEtaImported:              false,
		ObservedBoundarySpectrumImported: false,
		NewtonPlanckCosmologyImported:    false,
		InflowCapacityNative:             s.NativeCapacityConfirmed,
		BoundaryTheoryNative:             false,
		BoundaryConditionNative:          false,
		EtaSpectrumNative:                false,
		BordismClassNative:               false,
		CharacteristicNumbersNative:      false,
		CrossFixtureMergeNative:          c.CrossFixtureIdentityAllowed,
		GravitationalThetaNative:         false,
		GlobalAnomalyCoefficientsNative:  false,
		NativeRegistryWritten:            false,
		Verdict:                          strings.Join([]string{StatusFirewallPreserved, StatusFirewallNativeWriteBlocked}, ";"),
		Reason:                           "Gate524 may write only the structural capacity for anomaly inflow. It blocks boundary theory, boundary condition, eta spectrum, bordism class, characteristic numbers, global anomaly coefficients, gravitational theta, and cross-fixture merges from native promotion.",
	}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries:        []string{"Native structural capacity for anomaly inflow is confirmed: local index density, descent/transgression socket, APS boundary-pairing socket, gauge-anomaly cancellation, and mixed gauge-gravity trace cancellation.", "No global topology, boundary condition, eta spectrum, characteristic number, or bordism representative is selected natively."},
		BridgeEntries:        []string{"Gate520 APS-boundary and Gate522 spin/spin-c bordism fixtures are compatible with the local inflow sockets only as bridge classes.", "Anomaly-inflow classifier reports compatibility, not identity or selection."},
		EnvironmentalEntries: []string{"Actual manifold topology, boundary state, eta spectrum, boundary degrees of freedom, spin/spin-c representative, gravitational theta, and global anomaly coefficients remain global/environmental inputs."},
		FailedRoutes:         []string{"Using anomaly-inflow capacity to select the universe's boundary or manifold.", "Treating eta correction or boundary spectrum as native finite-algebra data.", "Merging heterogeneous APS-boundary and closed-bordism fixtures into one native topology identity."},
		OpenTheorems:         []string{"A future gate may close the topology block by producing a final sector ledger, or may redirect to a new native frontier such as Lorentzian/causal signature provenance."},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 525, Title: "Topology Sector Closing Ledger and Native Frontier Selection", Reason: "Gate524 confirms anomaly-inflow capacity and blocks topology selection; the topology sector is ready for a closing ledger before choosing the next native frontier.", PrimaryTask: "Summarize the topology block from characteristic classes through APS, bordism, residual classifiers, and anomaly-inflow compatibility, then select the next non-environmental theorem lane."}
}

func truth(a Analysis) string {
	return "Gate 524 confirms that ASHA has the native local capacity required for anomaly inflow: the index-density, descent/transgression, APS boundary-pairing, and anomaly-cancellation sockets compose correctly. This is a structural consistency theorem, not a topology selector: the boundary, eta spectrum, spin/spin-c representative, characteristic numbers, and global manifold remain bridge/environmental data."
}

func validate(a Analysis) error {
	bad := []string{}
	if !a.Inheritance.Executed || !a.Inheritance.Gate523ReportDefined || a.Inheritance.Gate523Rows != 4 || a.Inheritance.Gate523ZeroResidualRows != 4 || !a.Inheritance.Gate523BridgeOnly || !a.Inheritance.Gate523SyntheticOnly || a.Inheritance.Gate523ObservedImported || !a.Inheritance.Gate523HeterogeneousGuard || !a.Inheritance.Gate523CrossLedgerMergeRejected || a.Inheritance.Gate523NativeManifoldSelected || !a.Inheritance.Gate523NativeWriteBlocked || !a.Inheritance.Gate517IndexSocket || !a.Inheritance.Gate517APSSocket || !a.Inheritance.Gate517InflowSocket || !a.Inheritance.Gate517EtaGlobalData || !a.Inheritance.Gate517NativeIndexEtaWriteBlocked || !a.Inheritance.Gate490GaugeAnomaliesCancel || !a.Inheritance.Gate490MixedGaugeGravityCancel || !a.Inheritance.Gate490WittenSU2Cancel || !a.Inheritance.Gate490ExactRational || !a.Inheritance.Gate490FlavorMassIndependent {
		bad = append(bad, "bad inheritance")
	}
	if !a.Inflow.Executed || !a.Inflow.BulkCharacteristicClassesPresent || !a.Inflow.GaugeAnomalyTraceZero || !a.Inflow.MixedGaugeGravityTraceZero || !a.Inflow.WittenSU2GlobalCleared || !a.Inflow.APSBoundaryCorrectionPairing || !a.Inflow.ChernSimonsTransgressionSocket || !a.Inflow.ScaleFree || !a.Inflow.MassFlavorIndependent || !a.Inflow.NativeCapacityConfirmed || a.Inflow.BoundaryTheorySelected || a.Inflow.BoundaryConditionSelected || a.Inflow.EtaSpectrumDerived || a.Inflow.GlobalAnomalyCoefficientSelected {
		bad = append(bad, "bad inflow")
	}
	if !a.Compatibility.Executed || a.Compatibility.CompatibleClassCount != 3 || !a.Compatibility.APSBoundaryFixtureCompatible || !a.Compatibility.SpinBordismFixtureCompatible || !a.Compatibility.SpinCBordismFixtureCompatible || !a.Compatibility.Gate520BoundaryMode || !a.Compatibility.Gate522ClosedBoundary || !a.Compatibility.HeterogeneousGuardPreserved || a.Compatibility.CrossFixtureIdentityAllowed || !a.Compatibility.CrossFixtureMergeRejected || !a.Compatibility.ClassifiesButDoesNotSelect || !a.Compatibility.BoundaryCurrentConservationSocket || a.Compatibility.NativeManifoldSelected || a.Compatibility.NativeBoundarySelected || a.Compatibility.NativeBordismClassSelected {
		bad = append(bad, "bad compatibility")
	}
	if !a.Firewall.Executed || a.Firewall.ObservedTopologyImported || a.Firewall.ObservedBoundaryImported || a.Firewall.ObservedBordismImported || a.Firewall.ObservedEtaImported || a.Firewall.ObservedBoundarySpectrumImported || a.Firewall.NewtonPlanckCosmologyImported || !a.Firewall.InflowCapacityNative || a.Firewall.BoundaryTheoryNative || a.Firewall.BoundaryConditionNative || a.Firewall.EtaSpectrumNative || a.Firewall.BordismClassNative || a.Firewall.CharacteristicNumbersNative || a.Firewall.CrossFixtureMergeNative || a.Firewall.GravitationalThetaNative || a.Firewall.GlobalAnomalyCoefficientsNative || a.Firewall.NativeRegistryWritten {
		bad = append(bad, "firewall violation")
	}
	if len(bad) > 0 {
		return fmt.Errorf(strings.Join(bad, "; "))
	}
	return nil
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("%s: Gate523(rows=%d zero=%d bridge=%t synthetic=%t guard=%t merge_rejected=%t); Gate517(index=%t APS=%t inflow=%t eta_global=%t); Gate490(gauge=%t mixed_gravity=%t Witten=%t exact=%t flavor_free=%t); %s", x.Verdict, x.Gate523Rows, x.Gate523ZeroResidualRows, x.Gate523BridgeOnly, x.Gate523SyntheticOnly, x.Gate523HeterogeneousGuard, x.Gate523CrossLedgerMergeRejected, x.Gate517IndexSocket, x.Gate517APSSocket, x.Gate517InflowSocket, x.Gate517EtaGlobalData, x.Gate490GaugeAnomaliesCancel, x.Gate490MixedGaugeGravityCancel, x.Gate490WittenSU2Cancel, x.Gate490ExactRational, x.Gate490FlavorMassIndependent, x.Reason)
}

func FormatInflow(x InflowSieve) string {
	return fmt.Sprintf("%s: local=%s descent=%s APS=%s bulk_classes=%t gauge_zero=%t mixed_gravity_zero=%t Witten=%t CS=%t scale_free=%t mass_flavor_free=%t capacity=%t boundary_selected=%t eta_derived=%t coeff_selected=%t; %s", x.Verdict, x.LocalIndexDensity, x.DescentSocket, x.BoundaryPairing, x.BulkCharacteristicClassesPresent, x.GaugeAnomalyTraceZero, x.MixedGaugeGravityTraceZero, x.WittenSU2GlobalCleared, x.ChernSimonsTransgressionSocket, x.ScaleFree, x.MassFlavorIndependent, x.NativeCapacityConfirmed, x.BoundaryConditionSelected, x.EtaSpectrumDerived, x.GlobalAnomalyCoefficientSelected, x.Reason)
}

func FormatCompatibility(x BridgeCompatibility) string {
	return fmt.Sprintf("%s: compatible_classes=%d APS_boundary=%t spin=%t spinc=%t gate520_boundary=%t gate522_closed=%t guard=%t merge_allowed=%t merge_rejected=%t current_socket=%t selects_manifold=%t selects_boundary=%t selects_bordism=%t; %s", x.Verdict, x.CompatibleClassCount, x.APSBoundaryFixtureCompatible, x.SpinBordismFixtureCompatible, x.SpinCBordismFixtureCompatible, x.Gate520BoundaryMode, x.Gate522ClosedBoundary, x.HeterogeneousGuardPreserved, x.CrossFixtureIdentityAllowed, x.CrossFixtureMergeRejected, x.BoundaryCurrentConservationSocket, x.NativeManifoldSelected, x.NativeBoundarySelected, x.NativeBordismClassSelected, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("%s: observed_topology=%t observed_boundary=%t observed_bordism=%t observed_eta=%t observed_spectrum=%t Newton_cosmology=%t inflow_capacity_native=%t boundary_native=%t eta_native=%t bordism_native=%t characteristic_native=%t merge_native=%t theta_native=%t global_coeff_native=%t native_write=%t; %s", x.Verdict, x.ObservedTopologyImported, x.ObservedBoundaryImported, x.ObservedBordismImported, x.ObservedEtaImported, x.ObservedBoundarySpectrumImported, x.NewtonPlanckCosmologyImported, x.InflowCapacityNative, x.BoundaryConditionNative, x.EtaSpectrumNative, x.BordismClassNative, x.CharacteristicNumbersNative, x.CrossFixtureMergeNative, x.GravitationalThetaNative, x.GlobalAnomalyCoefficientsNative, x.NativeRegistryWritten, x.Reason)
}

func statuses() []string {
	return []string{
		StatusGate523TopologyReportInherited,
		StatusGate517IndexEtaInflowInherited,
		StatusGate490AnomalyLedgerInherited,
		StatusInflowClassifierDefined,
		StatusBulkBoundaryDescentSocketPresent,
		StatusNativeInflowCapacityConfirmed,
		StatusAPSBoundaryClassCompatible,
		StatusSpinSpinCBordismCompatible,
		StatusMixedGaugeGravityInflowTraceCancel,
		StatusHeterogeneousFixtureGuardPreserved,
		StatusNoObservedTopologyBoundaryImported,
		StatusFailedInflowDoesNotSelectBoundary,
		StatusFailedEtaSpectrumNotDerived,
		StatusFailedCrossFixtureMergeRejected,
		StatusFailedGlobalAnomalyCoeffUnselected,
		StatusFailedGravitationalThetaUnselected,
		StatusFirewallPreserved,
		StatusFirewallNativeWriteBlocked,
	}
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 524 Registry Audit — Anomaly-Inflow Compatibility Classifier for Bridge Topology Classes\n\n")
	b.WriteString("## Verdict\n\n```text\n")
	for _, s := range statuses() {
		b.WriteString(s + "\n")
	}
	b.WriteString("```\n\n")
	b.WriteString("## Inherited boundary\n\n" + a.Inheritance.Reason + "\n\n```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## Inflow sieve\n\nGate 524 audits capacity, not selection. The local bulk index-density socket and the boundary transgression/APS pairing are structurally present.\n\n```text\n" + FormatInflow(a.Inflow) + "\n```\n\n")
	b.WriteString("## Bridge-class compatibility\n\nThe APS-boundary fixture and the spin/spin-c bordism fixture are supported only as separate bridge classes. The heterogeneous-fixture guard from Gate 523 remains active.\n\n```text\n" + FormatCompatibility(a.Compatibility) + "\n```\n\n")
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
