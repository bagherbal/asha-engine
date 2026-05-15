package generation2topologyresidualclassifierreport

import (
	"strings"
	"testing"
)

func TestGate523Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate520FileLoaded || !a.Inheritance.Gate520APSResidualZero || !a.Inheritance.Gate520SignatureResidualZero || !a.Inheritance.Gate520BoundaryMode || !a.Inheritance.Gate522FileLoaded || !a.Inheritance.Gate522OrientedAdmissible || !a.Inheritance.Gate522SpinAdmissible || !a.Inheritance.Gate522SpinCAdmissible || !a.Inheritance.Gate522CharacteristicResidualsZero || !a.Inheritance.Gate522ClosedBoundary {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if a.Report.Rows != 4 || a.Report.ZeroResidualRows != 4 || a.Report.APSBoundaryRows != 2 || a.Report.ClosedBordismRows != 2 || !a.Report.BridgeOnly || !a.Report.SyntheticOnly || a.Report.ObservedImported || a.Report.NativePrediction || !a.Report.ReportReady || !a.Report.ClassifiesButDoesNotSelect {
		t.Fatalf("bad report: %+v", a.Report)
	}
	if a.Guard.CrossLedgerIdentityAsserted || a.Guard.CrossLedgerIdentityAllowed || !a.Guard.CrossLedgerMergeRejected || !a.Guard.DifferentSyntheticContexts || !a.Guard.BoundaryStatusCompatibleOnlyIfSeparated || !nearly(a.Guard.MergedSignatureResidual, 17, 1e-12) || !nearly(a.Guard.BoundaryComponentResidualIfMerged, 1, 1e-12) || a.Guard.NativeManifoldSelected {
		t.Fatalf("bad guard: %+v", a.Guard)
	}
	if a.Firewall.ObservedTopologyImported || a.Firewall.ObservedBoundaryImported || a.Firewall.ObservedBordismImported || a.Firewall.ObservedTangentBundleImported || a.Firewall.FileResidualsNative || a.Firewall.ReportNative || a.Firewall.ZeroResidualsNativeSelector || a.Firewall.CrossLedgerMergeNative || a.Firewall.BoundaryConditionNativeSelected || a.Firewall.EtaNativeSelected || a.Firewall.BordismClassNativeSelected || a.Firewall.CharacteristicNumbersNativeSelected || a.Firewall.ManifoldRepresentativeNative || a.Firewall.NewtonPlanckCosmologyImported || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
	if a.Next.Gate != 524 {
		t.Fatalf("bad next gate: %+v", a.Next)
	}
}

func TestGate523Markdown(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	md := Markdown(a)
	for _, want := range []string{"# Gate 523 Registry Audit", StatusResidualClassifierReportDefined, StatusHeterogeneousFixtureGuard, StatusFirewallNativeWriteBlocked, "Gate520", "Gate522", "Gate 524"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate523Theorem(t *testing.T) {
	result := Generation2TopologyResidualClassifierReportNativeNonSelectionAuditTheorem().Verify()
	for _, c := range result.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
}
