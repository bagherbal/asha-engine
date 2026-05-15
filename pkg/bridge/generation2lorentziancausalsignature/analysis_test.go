package generation2lorentziancausalsignature

import (
	"strings"
	"testing"
)

func TestGate526Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate525TopologyClosed || !a.Inheritance.Gate525FlavorClosed || !a.Inheritance.Gate525EWScaleClosed || !a.Inheritance.Gate525GravityNormalizationClosed || !a.Inheritance.Gate525LorentzianFrontierSelected || a.Inheritance.Gate525ReopensSealedFirewalls || a.Inheritance.Gate525ObservedDataImported || !a.Inheritance.Gate525NativeWriteBlocked {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if a.Signature.TimeLikeDirections != 1 || a.Signature.SpaceLikeDirections != 7 || a.Signature.TotalDimension != 8 || !a.Signature.MetricSignatureNative || !a.Signature.QuadraticFormNative || !a.Signature.NullConeDefined || !a.Signature.CausalConeScaleFree || !a.Signature.MassIndependent || !a.Signature.ConventionSignPairAmbiguous || a.Signature.Physical3Plus1ProjectionFound || a.Signature.TimeOrientationSelected || a.Signature.ArrowOfTimeDerived {
		t.Fatalf("bad signature socket: %+v", a.Signature)
	}
	if !a.Dictionary.EuclideanSpectralActionInherited || !a.Dictionary.HeatKernelEllipticConvention || !a.Dictionary.LorentzianRealTimeRequired || !a.Dictionary.BridgeDictionaryDefined || a.Dictionary.WickRotationSelectedNatively || a.Dictionary.IepsilonPrescriptionSelected || a.Dictionary.ReflectionPositivityProven || a.Dictionary.OsterwalderSchraderAxiomsProven || a.Dictionary.PositiveEnergyConditionDerived || a.Dictionary.UnitaryTimeEvolutionDerived || a.Dictionary.GlobalHyperbolicitySelected {
		t.Fatalf("bad dictionary: %+v", a.Dictionary)
	}
	if a.Firewall.ObservedConstantsImported || a.Firewall.ObservedMassesImported || a.Firewall.ObservedTopologyImported || a.Firewall.NativeWickWrite || a.Firewall.NativeTimeOrientationWrite || a.Firewall.NativePositiveEnergyWrite || a.Firewall.NativeUnitaryDynamicsWrite || a.Firewall.Native3Plus1Write || a.Firewall.NativeGlobalCausalWrite || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
}

func TestGate526Markdown(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	md := Markdown(a)
	for _, want := range []string{"# Gate 526 Registry Audit", StatusCL17SignatureSocketConfirmed, StatusFailedNoWickSelection, StatusFirewallNativeWriteBlocked, "Gate 527"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate526Theorem(t *testing.T) {
	result := Generation2LorentzianCausalSignatureProvenanceAndWickTimeFirewallAuditTheorem().Verify()
	for _, c := range result.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
}
