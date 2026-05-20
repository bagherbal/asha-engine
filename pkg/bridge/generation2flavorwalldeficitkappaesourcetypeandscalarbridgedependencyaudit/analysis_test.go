package generation2flavorwalldeficitkappaesourcetypeandscalarbridgedependencyaudit

import (
	"strings"
	"testing"
)

func TestGate746KappaESourceAndReplacement(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate745.Inherited || !a.Gate745.DiagnosticAllowed || !a.Gate745.ExternalPoleIsNotASHA || !a.Gate745.NoIndependentPoleMass {
		t.Fatalf("bad Gate745 inheritance: %+v", a.Gate745)
	}
	if !a.Dependency.StructurallyActive || !a.Dependency.AppearsInBoundaryPolynomial || !a.Dependency.AppearsInRuntimeTransport || !Near(a.Dependency.KappaE, 0.00550355419157456, 1e-18) {
		t.Fatalf("bad kappa_e dependency: %+v", a.Dependency)
	}
	if !Near(a.Orientation.KappaEOrient, 0.00550633006471245, 1e-18) || !Near(a.Orientation.DeltaKappaE, -2.77587313789e-6, 1e-15) || !a.Orientation.CloseButNotExact || a.Orientation.NativeTheorem {
		t.Fatalf("bad orientation candidate: %+v", a.Orientation)
	}
	if !Near(a.Replacement.FWallShift, 4.5080547e-13, 1e-18) || !Near(a.Replacement.RuntimeOrientResidual, -1.3795353970280644e-8, 5e-14) || !a.Replacement.ApproximationOnly {
		t.Fatalf("bad replacement test: %+v", a.Replacement)
	}
}

func TestGate746ResidualAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Residual.Candidates) != 4 || a.Residual.NativeSourceCertified {
		t.Fatalf("bad residual source audit: %+v", a.Residual)
	}
	if a.Firewall.DerivesPMNS || a.Firewall.DerivesCKM || a.Firewall.DerivesFlavorHierarchy || a.Firewall.DerivesYukawaEigenvalues || a.Firewall.DerivesScalarRuntime || a.Firewall.DerivesHiggsMass || !a.Firewall.KappaEStillBridgeSeal {
		t.Fatalf("bad firewall: %+v", a.Firewall)
	}
	res := Generation2FlavorWallDeficitKappaESourceTypeAndScalarBridgeDependencyAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status note %s", want)
		}
	}
}
