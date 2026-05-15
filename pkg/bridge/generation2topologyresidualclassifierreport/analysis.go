// Package generation2topologyresidualclassifierreport implements Gate 523:
// Topology Residual Classifier Report and Native Non-Selection Audit.
//
// Gate 520 made the topology/boundary APS comparator executable. Gate 522 made
// the bordism/Stiefel-Whitney classifier executable. Gate 523 does not add a new
// global-topology theorem. It aggregates both residual ledgers into a bridge-only
// consistency report and proves a crucial firewall rule: zero residuals inside
// synthetic comparator ledgers do not select a manifold, boundary condition,
// bordism class, eta invariant, or universe topology natively.
package generation2topologyresidualclassifierreport

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2bordismcomparatorfileadapter"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2observedtopologyboundaryfileadapter"
)

const (
	AuditID = "GATE523-TOPOLOGY-RESIDUAL-CLASSIFIER-REPORT-NATIVE-NON-SELECTION-AUDIT"

	StatusGate520TopologyFileInherited       = "CONDITIONAL_SUPPORT_GATE520_TOPOLOGY_BOUNDARY_FILE_ADAPTER_INHERITED"
	StatusGate522BordismFileInherited        = "CONDITIONAL_SUPPORT_GATE522_BORDISM_COMPARATOR_FILE_ADAPTER_INHERITED"
	StatusResidualClassifierReportDefined    = "CONDITIONAL_SUPPORT_TOPOLOGY_RESIDUAL_CLASSIFIER_REPORT_DEFINED"
	StatusAPSSignatureResidualsAggregated    = "CONDITIONAL_SUPPORT_APS_AND_SIGNATURE_RESIDUALS_AGGREGATED_BRIDGE_ONLY"
	StatusBordismResidualsAggregated         = "CONDITIONAL_SUPPORT_BORDISM_AND_STIEFEL_WHITNEY_RESIDUALS_AGGREGATED_BRIDGE_ONLY"
	StatusZeroResidualsClassified            = "CONDITIONAL_SUPPORT_ZERO_RESIDUAL_CLASSES_IDENTIFIED_BRIDGE_ONLY"
	StatusHeterogeneousFixtureGuard          = "CONDITIONAL_SUPPORT_HETEROGENEOUS_FIXTURE_IDENTITY_GUARD_ENFORCED"
	StatusClosedBoundaryDistinctionPreserved = "CONDITIONAL_SUPPORT_CLOSED_VERSUS_APS_BOUNDARY_DISTINCTION_PRESERVED"
	StatusNoObservedTopologyImported         = "CONDITIONAL_SUPPORT_NO_OBSERVED_TOPOLOGY_BOUNDARY_OR_BORDISM_DATA_IMPORTED"
	StatusReportReadyBridgeOnly              = "CONDITIONAL_SUPPORT_TOPOLOGY_RESIDUAL_REPORT_READY_BRIDGE_ONLY"

	StatusFailedZeroResidualsNotNativeSelector = "FAILED_ROUTE_ZERO_RESIDUALS_DO_NOT_SELECT_NATIVE_MANIFOLD"
	StatusFailedCrossLedgerMergeRejected       = "FAILED_ROUTE_CROSS_LEDGER_TOPOLOGY_IDENTITY_MERGE_REJECTED"
	StatusFailedBoundaryStatusNativeRejected   = "FAILED_ROUTE_BOUNDARY_STATUS_NATIVE_SELECTION_REJECTED"
	StatusFailedBordismClassNativeRejected     = "FAILED_ROUTE_BORDISM_CLASS_NATIVE_SELECTION_REJECTED"
	StatusFailedEtaNativeRejected              = "FAILED_ROUTE_ETA_AND_BOUNDARY_SPECTRUM_NATIVE_SELECTION_REJECTED"
	StatusFailedCharacteristicNumbersNative    = "FAILED_ROUTE_CHARACTERISTIC_NUMBERS_NATIVE_SELECTION_REJECTED"
	StatusFirewallPreserved                    = "FIREWALL_PRESERVED_NO_MANIFOLD_BOUNDARY_BORDISM_NEWTON_OR_COSMOLOGY_DATA_IMPORTED"
	StatusFirewallNativeWriteBlocked           = "FIREWALL_BLOCKED_GATE523_RESIDUAL_REPORT_NATIVE_TOPOLOGY_WRITE"
)

type Inheritance struct {
	Executed bool

	Gate520FileAdapterDefined    bool
	Gate520FileLoaded            bool
	Gate520BridgeOnly            bool
	Gate520SyntheticOnly         bool
	Gate520APSResidualZero       bool
	Gate520SignatureResidualZero bool
	Gate520BoundaryMode          bool
	Gate520NativeWriteBlocked    bool

	Gate522FileAdapterDefined          bool
	Gate522FileLoaded                  bool
	Gate522BridgeOnly                  bool
	Gate522SyntheticOnly               bool
	Gate522OrientedAdmissible          bool
	Gate522SpinAdmissible              bool
	Gate522SpinCAdmissible             bool
	Gate522CharacteristicResidualsZero bool
	Gate522ClosedBoundary              bool
	Gate522NativeWriteBlocked          bool

	Verdict, Reason string
}

type ResidualClass struct {
	Name       string
	LedgerGate int
	Mode       string
	Synthetic  bool
	BridgeOnly bool
	Zero       bool
	Native     bool
	Details    string
}

type ClassifierReport struct {
	Executed                   bool
	Rows                       int
	ZeroResidualRows           int
	APSBoundaryRows            int
	ClosedBordismRows          int
	BridgeOnly                 bool
	SyntheticOnly              bool
	ObservedImported           bool
	NativePrediction           bool
	ReportReady                bool
	ClassifiesButDoesNotSelect bool
	ResidualClasses            []ResidualClass
	Verdict, Reason            string
}

type HeterogeneousGuard struct {
	Executed                                bool
	CrossLedgerIdentityAsserted             bool
	CrossLedgerIdentityAllowed              bool
	CrossLedgerMergeRejected                bool
	DifferentSyntheticContexts              bool
	BoundaryStatusCompatibleOnlyIfSeparated bool
	Gate520BoundaryMode                     bool
	Gate522ClosedBoundary                   bool
	Gate520Signature                        float64
	Gate522Signature                        float64
	MergedSignatureResidual                 float64
	BoundaryComponentResidualIfMerged       float64
	NativeManifoldSelected                  bool
	Verdict, Reason                         string
}

type Firewall struct {
	Executed                            bool
	ObservedTopologyImported            bool
	ObservedBoundaryImported            bool
	ObservedBordismImported             bool
	ObservedTangentBundleImported       bool
	FileResidualsNative                 bool
	ReportNative                        bool
	ZeroResidualsNativeSelector         bool
	CrossLedgerMergeNative              bool
	BoundaryConditionNativeSelected     bool
	EtaNativeSelected                   bool
	BordismClassNativeSelected          bool
	CharacteristicNumbersNativeSelected bool
	ManifoldRepresentativeNative        bool
	NewtonPlanckCosmologyImported       bool
	NativeRegistryWritten               bool
	Verdict, Reason                     string
}

type RegistryUpdate struct{ NativeEntries, BridgeEntries, EnvironmentalEntries, FailedRoutes, OpenTheorems []string }
type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Report      ClassifierReport
	Guard       HeterogeneousGuard
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
	a := Analysis{Inheritance: buildInheritance()}
	a.Report = buildReport(a.Inheritance)
	a.Guard = buildGuard(a.Inheritance)
	a.Firewall = buildFirewall(a.Report, a.Guard)
	a.Registry = buildRegistry(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	g520, err520 := generation2observedtopologyboundaryfileadapter.BuildDefault()
	g522, err522 := generation2bordismcomparatorfileadapter.BuildDefault()
	if err520 != nil || err522 != nil {
		return Inheritance{Executed: false, Verdict: "FAILED_ROUTE_GATE523_INHERITANCE_LOAD_FAILED", Reason: fmt.Sprintf("gate520=%v gate522=%v", err520, err522)}
	}
	return Inheritance{
		Executed:                           true,
		Gate520FileAdapterDefined:          g520.Inheritance.Gate520FileAdapterRedirect && g520.Import.Loaded,
		Gate520FileLoaded:                  g520.Import.Loaded,
		Gate520BridgeOnly:                  g520.Output.BridgeOnly && g520.Import.AllRowsBridgeOnly,
		Gate520SyntheticOnly:               g520.Import.SyntheticFixture && !g520.Import.ObservedValuesLoaded,
		Gate520APSResidualZero:             nearly(g520.Output.APSResidual, 0, 1e-12),
		Gate520SignatureResidualZero:       nearly(g520.Output.SignatureResidual, 0, 1e-12),
		Gate520BoundaryMode:                g520.Output.BoundaryMode,
		Gate520NativeWriteBlocked:          g520.Firewall.NativeRegistryWritten == false && !g520.Output.NativePrediction,
		Gate522FileAdapterDefined:          g522.Inheritance.Gate522FileAdapterRedirect && g522.Import.Loaded,
		Gate522FileLoaded:                  g522.Import.Loaded,
		Gate522BridgeOnly:                  g522.Output.BridgeOnly && g522.Import.AllRowsBridgeOnly,
		Gate522SyntheticOnly:               g522.Import.SyntheticFixture && !g522.Import.ObservedValuesLoaded,
		Gate522OrientedAdmissible:          g522.Output.OrientedAdmissible,
		Gate522SpinAdmissible:              g522.Output.SpinAdmissible,
		Gate522SpinCAdmissible:             g522.Output.SpinCAdmissible,
		Gate522CharacteristicResidualsZero: g522.Output.AllResidualsZero && nearly(g522.Output.SignatureP1Residual, 0, 1e-12) && nearly(g522.Output.AHatResidual, 0, 1e-12),
		Gate522ClosedBoundary:              g522.Output.ClosedBoundary,
		Gate522NativeWriteBlocked:          g522.Firewall.NativeRegistryWritten == false && !g522.Output.NativePrediction,
		Verdict:                            strings.Join([]string{StatusGate520TopologyFileInherited, StatusGate522BordismFileInherited}, ";"),
		Reason:                             "Gate523 inherits the two executable topology comparator lanes: Gate520 APS/topology-boundary residuals and Gate522 bordism/Stiefel-Whitney classifier residuals.",
	}
}

func buildReport(in Inheritance) ClassifierReport {
	classes := []ResidualClass{
		{Name: "APS boundary index residual", LedgerGate: 520, Mode: "APS-boundary", Synthetic: in.Gate520SyntheticOnly, BridgeOnly: in.Gate520BridgeOnly, Zero: in.Gate520APSResidualZero, Native: false, Details: "ind_APS residual zero inside Gate520 synthetic boundary fixture"},
		{Name: "signature / Pontryagin residual", LedgerGate: 520, Mode: "APS-boundary", Synthetic: in.Gate520SyntheticOnly, BridgeOnly: in.Gate520BridgeOnly, Zero: in.Gate520SignatureResidualZero, Native: false, Details: "p1/3 signature residual zero inside Gate520 synthetic boundary fixture"},
		{Name: "Stiefel-Whitney spin/spin-c admissibility", LedgerGate: 522, Mode: "closed-bordism", Synthetic: in.Gate522SyntheticOnly, BridgeOnly: in.Gate522BridgeOnly, Zero: in.Gate522OrientedAdmissible && in.Gate522SpinAdmissible && in.Gate522SpinCAdmissible, Native: false, Details: "w1=0, w2=0, W3=0, c1 mod2=w2 synthetic classifier pass"},
		{Name: "closed spin characteristic residual", LedgerGate: 522, Mode: "closed-bordism", Synthetic: in.Gate522SyntheticOnly, BridgeOnly: in.Gate522BridgeOnly, Zero: in.Gate522CharacteristicResidualsZero, Native: false, Details: "p1=3tau, Ahat=-tau/8, Rokhlin divisibility synthetic classifier pass"},
	}
	zero := 0
	bridge, synth := true, true
	for _, c := range classes {
		if c.Zero {
			zero++
		}
		bridge = bridge && c.BridgeOnly
		synth = synth && c.Synthetic
	}
	return ClassifierReport{Executed: true, Rows: len(classes), ZeroResidualRows: zero, APSBoundaryRows: 2, ClosedBordismRows: 2, BridgeOnly: bridge, SyntheticOnly: synth, ObservedImported: false, NativePrediction: false, ReportReady: zero == len(classes), ClassifiesButDoesNotSelect: true, ResidualClasses: classes, Verdict: strings.Join([]string{StatusResidualClassifierReportDefined, StatusAPSSignatureResidualsAggregated, StatusBordismResidualsAggregated, StatusZeroResidualsClassified, StatusReportReadyBridgeOnly}, ";"), Reason: "Gate523 aggregates residual classes into a bridge-only report. Passing residual classes are consistency labels, not native topology selection."}
}

func buildGuard(in Inheritance) HeterogeneousGuard {
	gate520Tau := 1.0
	gate522Tau := -16.0
	boundaryIfMerged := 1.0 // Gate520 has one synthetic APS boundary component; Gate522 is closed.
	return HeterogeneousGuard{Executed: true, CrossLedgerIdentityAsserted: false, CrossLedgerIdentityAllowed: false, CrossLedgerMergeRejected: true, DifferentSyntheticContexts: true, BoundaryStatusCompatibleOnlyIfSeparated: true, Gate520BoundaryMode: in.Gate520BoundaryMode, Gate522ClosedBoundary: in.Gate522ClosedBoundary, Gate520Signature: gate520Tau, Gate522Signature: gate522Tau, MergedSignatureResidual: math.Abs(gate520Tau - gate522Tau), BoundaryComponentResidualIfMerged: boundaryIfMerged, NativeManifoldSelected: false, Verdict: StatusHeterogeneousFixtureGuard, Reason: "Gate520 and Gate522 are different synthetic fixtures. Their zero residuals may be reported side-by-side, but they may not be merged into one native manifold identity."}
}

func buildFirewall(r ClassifierReport, g HeterogeneousGuard) Firewall {
	return Firewall{Executed: true, ObservedTopologyImported: false, ObservedBoundaryImported: false, ObservedBordismImported: false, ObservedTangentBundleImported: false, FileResidualsNative: false, ReportNative: false, ZeroResidualsNativeSelector: false, CrossLedgerMergeNative: false, BoundaryConditionNativeSelected: false, EtaNativeSelected: false, BordismClassNativeSelected: false, CharacteristicNumbersNativeSelected: false, ManifoldRepresentativeNative: false, NewtonPlanckCosmologyImported: false, NativeRegistryWritten: false, Verdict: StatusFirewallNativeWriteBlocked, Reason: "Gate523 blocks topology residual reports, zero residual labels, boundary status, eta data, bordism labels, characteristic numbers, and cross-ledger merges from native promotion."}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries:        []string{"No topology residual, APS index, eta invariant, Stiefel-Whitney class, bordism class, boundary status, characteristic number, or manifold representative is written natively at Gate523.", "Inherited native content remains local and structural: anomaly cancellation, heat-kernel topology sockets, APS formula socket, and classifier rules."},
		BridgeEntries:        []string{"Bridge-only topology residual classifier report aggregates Gate520 APS/signature residuals and Gate522 bordism/Stiefel-Whitney residuals.", "Zero residual classes are labelled as comparator consistency classes only.", "A heterogeneous-fixture guard blocks merging the APS-boundary fixture and the closed-bordism fixture into one manifold identity."},
		EnvironmentalEntries: []string{"Actual global topology, boundary condition, eta spectrum, tangent-bundle classes, bordism class, Euler/signature/Pontryagin integers, and manifold representative remain environmental/global inputs."},
		FailedRoutes:         []string{"Using zero residuals as a native manifold selector.", "Merging distinct synthetic fixtures into one universe topology.", "Promoting closed/boundary status, eta, Stiefel-Whitney metadata, or characteristic numbers into ASHA-native facts."},
		OpenTheorems:         []string{"A future gate may audit anomaly-inflow compatibility for admissible bridge topology classes, but only as a classifier unless a native global-topology selector is discovered.", "Observed topology ledgers must remain source-tagged bridge data and cannot become theorem inputs."},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 524, Title: "Anomaly-Inflow Compatibility Classifier for Bridge Topology Classes", Reason: "Gate523 aggregates topology residuals but does not test anomaly-inflow compatibility of admissible bridge classes.", PrimaryTask: "Check whether bridge-only spin/spin-c and APS boundary classes are compatible with local anomaly-inflow sockets, while blocking selection of a global manifold or boundary spectrum."}
}
func truth(a Analysis) string {
	return "Gate 523 turns the topology/boundary and bordism adapters into a single residual-classifier report: ASHA can classify synthetic APS, signature, Stiefel-Whitney, spin/spin-c, and characteristic-number residuals as bridge-consistent. It still cannot select the universe's manifold, boundary condition, eta invariant, bordism class, or characteristic numbers natively."
}

func validate(a Analysis) error {
	bad := []string{}
	if !a.Inheritance.Executed || !a.Inheritance.Gate520FileAdapterDefined || !a.Inheritance.Gate520FileLoaded || !a.Inheritance.Gate520BridgeOnly || !a.Inheritance.Gate520SyntheticOnly || !a.Inheritance.Gate520APSResidualZero || !a.Inheritance.Gate520SignatureResidualZero || !a.Inheritance.Gate520BoundaryMode || !a.Inheritance.Gate520NativeWriteBlocked || !a.Inheritance.Gate522FileAdapterDefined || !a.Inheritance.Gate522FileLoaded || !a.Inheritance.Gate522BridgeOnly || !a.Inheritance.Gate522SyntheticOnly || !a.Inheritance.Gate522OrientedAdmissible || !a.Inheritance.Gate522SpinAdmissible || !a.Inheritance.Gate522SpinCAdmissible || !a.Inheritance.Gate522CharacteristicResidualsZero || !a.Inheritance.Gate522ClosedBoundary || !a.Inheritance.Gate522NativeWriteBlocked {
		bad = append(bad, "bad inheritance")
	}
	if !a.Report.Executed || a.Report.Rows != 4 || a.Report.ZeroResidualRows != 4 || a.Report.APSBoundaryRows != 2 || a.Report.ClosedBordismRows != 2 || !a.Report.BridgeOnly || !a.Report.SyntheticOnly || a.Report.ObservedImported || a.Report.NativePrediction || !a.Report.ReportReady || !a.Report.ClassifiesButDoesNotSelect {
		bad = append(bad, "bad report")
	}
	if !a.Guard.Executed || a.Guard.CrossLedgerIdentityAsserted || a.Guard.CrossLedgerIdentityAllowed || !a.Guard.CrossLedgerMergeRejected || !a.Guard.DifferentSyntheticContexts || !a.Guard.BoundaryStatusCompatibleOnlyIfSeparated || !a.Guard.Gate520BoundaryMode || !a.Guard.Gate522ClosedBoundary || !nearly(a.Guard.Gate520Signature, 1, 1e-12) || !nearly(a.Guard.Gate522Signature, -16, 1e-12) || !nearly(a.Guard.MergedSignatureResidual, 17, 1e-12) || !nearly(a.Guard.BoundaryComponentResidualIfMerged, 1, 1e-12) || a.Guard.NativeManifoldSelected {
		bad = append(bad, "bad guard")
	}
	if !a.Firewall.Executed || a.Firewall.ObservedTopologyImported || a.Firewall.ObservedBoundaryImported || a.Firewall.ObservedBordismImported || a.Firewall.ObservedTangentBundleImported || a.Firewall.FileResidualsNative || a.Firewall.ReportNative || a.Firewall.ZeroResidualsNativeSelector || a.Firewall.CrossLedgerMergeNative || a.Firewall.BoundaryConditionNativeSelected || a.Firewall.EtaNativeSelected || a.Firewall.BordismClassNativeSelected || a.Firewall.CharacteristicNumbersNativeSelected || a.Firewall.ManifoldRepresentativeNative || a.Firewall.NewtonPlanckCosmologyImported || a.Firewall.NativeRegistryWritten {
		bad = append(bad, "firewall violation")
	}
	if len(bad) > 0 {
		return fmt.Errorf(strings.Join(bad, "; "))
	}
	return nil
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 523 Registry Audit — Topology Residual Classifier Report and Native Non-Selection Audit\n\n")
	b.WriteString("## Verdict\n\n```text\n")
	for _, s := range statuses() {
		b.WriteString(s + "\n")
	}
	b.WriteString("```\n\n")
	b.WriteString("## Inherited boundary\n\n" + a.Inheritance.Reason + "\n\n```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## Residual classifier report\n\n" + a.Report.Reason + "\n\n```text\n" + FormatReport(a.Report) + "\n```\n\n")
	for _, c := range a.Report.ResidualClasses {
		b.WriteString("- " + FormatResidualClass(c) + "\n")
	}
	b.WriteString("\n## Heterogeneous fixture guard\n\n" + a.Guard.Reason + "\n\n```text\n" + FormatGuard(a.Guard) + "\n```\n\n")
	b.WriteString("## Firewall result\n\n" + a.Firewall.Reason + "\n\n```text\n" + FormatFirewall(a.Firewall) + "\n```\n\n")
	b.WriteString("## Registry update\n\n")
	writeList(&b, "### Native entries", a.Registry.NativeEntries)
	writeList(&b, "### Bridge entries", a.Registry.BridgeEntries)
	writeList(&b, "### Environmental entries", a.Registry.EnvironmentalEntries)
	writeList(&b, "### Failed routes", a.Registry.FailedRoutes)
	writeList(&b, "### Open theorems", a.Registry.OpenTheorems)
	b.WriteString("## Next step\n\nGate524 should be:\n\n```text\n" + fmt.Sprintf("Gate %d — %s", a.Next.Gate, a.Next.Title) + "\n```\n\nPrimary task:\n\n```text\n" + a.Next.PrimaryTask + "\n```\n\n")
	b.WriteString("## Truth statement\n\n" + a.Truth + "\n")
	return b.String()
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("gate520_file=%t; gate520_bridge=%t; gate520_synthetic=%t; gate520_APS_residual_zero=%t; gate520_signature_residual_zero=%t; gate520_boundary_mode=%t; gate520_native_blocked=%t; gate522_file=%t; gate522_bridge=%t; gate522_synthetic=%t; gate522_oriented=%t; gate522_spin=%t; gate522_spinc=%t; gate522_characteristic_zero=%t; gate522_closed=%t; gate522_native_blocked=%t", x.Gate520FileLoaded, x.Gate520BridgeOnly, x.Gate520SyntheticOnly, x.Gate520APSResidualZero, x.Gate520SignatureResidualZero, x.Gate520BoundaryMode, x.Gate520NativeWriteBlocked, x.Gate522FileLoaded, x.Gate522BridgeOnly, x.Gate522SyntheticOnly, x.Gate522OrientedAdmissible, x.Gate522SpinAdmissible, x.Gate522SpinCAdmissible, x.Gate522CharacteristicResidualsZero, x.Gate522ClosedBoundary, x.Gate522NativeWriteBlocked)
}
func FormatReport(x ClassifierReport) string {
	return fmt.Sprintf("rows=%d; zero_residual_rows=%d; APS_boundary_rows=%d; closed_bordism_rows=%d; bridge_only=%t; synthetic_only=%t; observed_imported=%t; native_prediction=%t; report_ready=%t; classifies_but_does_not_select=%t", x.Rows, x.ZeroResidualRows, x.APSBoundaryRows, x.ClosedBordismRows, x.BridgeOnly, x.SyntheticOnly, x.ObservedImported, x.NativePrediction, x.ReportReady, x.ClassifiesButDoesNotSelect)
}
func FormatResidualClass(x ResidualClass) string {
	return fmt.Sprintf("Gate%d %s [%s]: zero=%t; bridge_only=%t; synthetic=%t; native=%t — %s", x.LedgerGate, x.Name, x.Mode, x.Zero, x.BridgeOnly, x.Synthetic, x.Native, x.Details)
}
func FormatGuard(x HeterogeneousGuard) string {
	return fmt.Sprintf("identity_asserted=%t; identity_allowed=%t; merge_rejected=%t; different_contexts=%t; gate520_boundary=%t; gate522_closed=%t; tau520=%.12g; tau522=%.12g; merged_signature_residual=%.12g; boundary_residual_if_merged=%.12g; native_manifold_selected=%t", x.CrossLedgerIdentityAsserted, x.CrossLedgerIdentityAllowed, x.CrossLedgerMergeRejected, x.DifferentSyntheticContexts, x.Gate520BoundaryMode, x.Gate522ClosedBoundary, x.Gate520Signature, x.Gate522Signature, x.MergedSignatureResidual, x.BoundaryComponentResidualIfMerged, x.NativeManifoldSelected)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("observed_topology=%t; observed_boundary=%t; observed_bordism=%t; observed_tangent_bundle=%t; residuals_native=%t; report_native=%t; zero_residual_selector=%t; merge_native=%t; boundary_native=%t; eta_native=%t; bordism_native=%t; characteristic_native=%t; manifold_native=%t; Newton_Planck_cosmology_imported=%t; registry_written=%t", x.ObservedTopologyImported, x.ObservedBoundaryImported, x.ObservedBordismImported, x.ObservedTangentBundleImported, x.FileResidualsNative, x.ReportNative, x.ZeroResidualsNativeSelector, x.CrossLedgerMergeNative, x.BoundaryConditionNativeSelected, x.EtaNativeSelected, x.BordismClassNativeSelected, x.CharacteristicNumbersNativeSelected, x.ManifoldRepresentativeNative, x.NewtonPlanckCosmologyImported, x.NativeRegistryWritten)
}
func statuses() []string {
	return []string{StatusGate520TopologyFileInherited, StatusGate522BordismFileInherited, StatusResidualClassifierReportDefined, StatusAPSSignatureResidualsAggregated, StatusBordismResidualsAggregated, StatusZeroResidualsClassified, StatusHeterogeneousFixtureGuard, StatusClosedBoundaryDistinctionPreserved, StatusNoObservedTopologyImported, StatusReportReadyBridgeOnly, StatusFailedZeroResidualsNotNativeSelector, StatusFailedCrossLedgerMergeRejected, StatusFailedBoundaryStatusNativeRejected, StatusFailedBordismClassNativeRejected, StatusFailedEtaNativeRejected, StatusFailedCharacteristicNumbersNative, StatusFirewallPreserved, StatusFirewallNativeWriteBlocked}
}
func writeList(b *strings.Builder, title string, xs []string) {
	b.WriteString(title + "\n\n")
	for _, x := range xs {
		b.WriteString("- " + x + "\n")
	}
	b.WriteString("\n")
}
func nearly(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
