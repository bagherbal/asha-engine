package generation2polecorrectionsealpackageandlevel1cdiagnosticboundaryaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate743PoleCorrectionSealPackage(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate742.Inherited || a.Gate742.CorrectionObject != "Delta_pole" || a.Gate742.CorrectionValueAssigned || !a.Gate742.TreeProxyNotPoleMass || !a.Gate742.Level1CAllowed || !a.Gate742.Level1CRequiresExternalPackage || a.Gate742.Level2Allowed {
		t.Fatalf("bad Gate742 inheritance: %+v", a.Gate742)
	}
	if math.Abs(a.Gate742.TreeProxyGeV-125.38000000298437) > 1e-9 {
		t.Fatalf("bad inherited tree proxy: %.17g", a.Gate742.TreeProxyGeV)
	}
	if !a.Package.FullPackage || a.Package.Count != 9 || !a.Package.HasPoleObservable || !a.Package.HasPoleMassConvention || !a.Package.HasRGScheme || !a.Package.HasRenormalizationScale || !a.Package.HasLoopOrder || !a.Package.HasThresholdCorrection || !a.Package.HasTopYukawaInput || !a.Package.HasGaugeCouplingInput || !a.Package.HasUncertaintyModel || a.Package.DeltaPoleValueAssigned {
		t.Fatalf("bad correction package: %+v", a.Package)
	}
	if !contains(a.Package.Labels, "PoleMassObservableSeal") || !contains(a.Package.Labels, "UncertaintyModelSeal") {
		t.Fatalf("missing correction seals: %+v", a.Package.Labels)
	}
}

func TestGate743MinimalityBoundaryAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Minimality.Minimal || !a.Minimality.AllRequired || a.Minimality.Count != 9 {
		t.Fatalf("bad minimality audit: %+v", a.Minimality)
	}
	for _, item := range a.Minimality.Items {
		if !item.RequiredForLevel1C || item.Seal == "" || item.RemovalEffect == "" {
			t.Fatalf("bad minimality item: %+v", item)
		}
	}
	if !a.Diagnostic.Level1BAllowed || !a.Diagnostic.Level1CAllowed || !a.Diagnostic.Level1CDiagnosticOnly || !a.Diagnostic.Level1CRequiresAll || a.Diagnostic.Level2Allowed {
		t.Fatalf("bad diagnostic boundary: %+v", a.Diagnostic)
	}
	if a.Separation.PoleObservableExternallySupplied || a.Separation.ExternalPoleObservableASHADerived || a.Separation.TreeProxyEqualsPoleObservable || !a.Separation.DiagnosticCanComputeDeltaOnlyWithPackage {
		t.Fatalf("bad separation audit: %+v", a.Separation)
	}
	if a.Firewall.FittedDeltaIsDerivedTheorem || a.Firewall.Level1CDiagnosticIsPrediction || a.Firewall.TreeProxyProximityIsTheorem || a.Firewall.ExternalObservableIsDerivation || a.Firewall.IndependentPolePrediction || !a.Firewall.NoYukawaTheorem {
		t.Fatalf("firewall failed: %+v", a.Firewall)
	}
	res := Generation2PoleCorrectionSealPackageAndLevel1CDiagnosticBoundaryAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
