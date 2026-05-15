package generation2projectionbridgeairlockpreflight

import (
	"strings"
	"testing"
)

func TestGate529Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate528Inherited || !a.Inheritance.Gate528Rank44BridgeSocketReady || !a.Inheritance.Gate528NoNativeRank4Projector || !a.Inheritance.Gate528TimeAssignmentBlocked || !a.Inheritance.Gate528InternalComplementBlocked || !a.Inheritance.Gate528WickHilbertDynamicsBlocked || !a.Inheritance.Gate528NoObservedDataImported || !a.Inheritance.Gate528NativeWriteBlocked || a.Inheritance.Gate528ReopenedSealedFirewalls {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if a.Schema.RequiredRowCount != 12 || !a.Schema.ProjectorMatrixRequired || !a.Schema.ProjectorIdempotencyCheck || a.Schema.ProjectorRankRequired != 4 || !a.Schema.ComplementMatrixRequired || a.Schema.ComplementRankRequired != 4 || !a.Schema.OrthogonalComplementCheck || a.Schema.ExternalSignatureRequired != "1+3" || !a.Schema.InternalAssignmentRequired || !a.Schema.SourceRequired || !a.Schema.ConventionRequired || !a.Schema.BridgeOnlyRequired || !a.Schema.NativePromotionRejected || !a.Schema.RedactedSchemaAccepted || a.Schema.AcceptedRedactedCases != 1 || a.Schema.RejectedFailClosedCases < 8 {
		t.Fatalf("bad schema: %+v", a.Schema)
	}
	if !a.Obligations.ProjectorImportedBridgeOnly || a.Obligations.GrantsWickRotation || a.Obligations.GrantsPositiveHilbertProduct || a.Obligations.GrantsReflectionPositivity || a.Obligations.GrantsPositiveEnergyHamiltonian || a.Obligations.GrantsUnitaryRealTimeDynamics || a.Obligations.GrantsGlobalHyperbolicity || a.Obligations.GrantsInternalGaugeIdentification || !a.Obligations.RequiresSeparateWickAirlock || !a.Obligations.RequiresSeparateHilbertAirlock || !a.Obligations.RequiresSeparateUnitaryDynamicsAirlock || !a.Obligations.RequiresSeparateInternalGaugeAirlock {
		t.Fatalf("bad obligations guard: %+v", a.Obligations)
	}
	if !a.Rejection.NativeProjectorWriteRejected || !a.Rejection.Native3Plus1SpacetimeWriteRejected || !a.Rejection.NativeTimeAssignmentWriteRejected || !a.Rejection.NativeInternalComplementWriteRejected || !a.Rejection.NativeWickWriteRejected || !a.Rejection.NativeHilbertWriteRejected || !a.Rejection.NativeUnitaryDynamicsWriteRejected || a.Rejection.ComparatorExecutionPerformed {
		t.Fatalf("bad native rejection: %+v", a.Rejection)
	}
	if a.Firewall.ObservedDimensionImported || a.Firewall.ObservedConstantsImported || a.Firewall.ObservedMassesImported || a.Firewall.ObservedTopologyImported || a.Firewall.NativeProjectorWrite || a.Firewall.Native3Plus1Write || a.Firewall.NativeTimeAssignmentWrite || a.Firewall.NativeInternalComplementWrite || a.Firewall.NativeWickWrite || a.Firewall.NativeHilbertWrite || a.Firewall.NativeUnitaryDynamicsWrite || a.Firewall.NativeInternalGaugeWrite || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
}

func TestGate529Markdown(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	md := Markdown(a)
	for _, want := range []string{"# Gate 529 Registry Audit", StatusProjectionAirlockDefined, StatusFailedProjectorDoesNotGrantWick, StatusFailedProjectorDoesNotGrantHilbert, StatusFirewallProjectionNativeWriteBlocked, "Gate 530"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate529Theorem(t *testing.T) {
	result := Generation2ProjectionBridgeAirlockPreflightTheorem().Verify()
	for _, c := range result.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
}
