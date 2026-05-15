package generation2abeliancoefficientselection

import (
	"strings"
	"testing"
)

func TestBuildDefaultGate494AbelianCoefficientSelection(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate493QuadraticFamilyTyped || a.Inheritance.Gate493KappaSelected || a.Inheritance.Gate493GaugeHessianSelected {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Hypercharge.KYConfirmed || !a.Hypercharge.BoundarySin238Confirmed || a.Hypercharge.SelectsKappaU1 || a.Hypercharge.PhysicalWeakMixingAngleDerived {
		t.Fatalf("bad hypercharge trace audit: %+v", a.Hypercharge)
	}
	if !a.Kappa.CompletionFamilyTyped || a.Kappa.TargetKappa != 6 || !a.Kappa.WhiteningSelectsKappa || a.Kappa.ActionSelectsKappa || a.Kappa.UniqueDerivation || a.Kappa.CountResonanceSelected {
		t.Fatalf("bad kappa audit: %+v", a.Kappa)
	}
	if !a.Metric.DiagonalTraceGramAsRepresentationMetric || a.Metric.DiagonalTraceGramAsGaugeKineticHessian || a.Metric.PhysicalGaugeCouplingsDerived {
		t.Fatalf("bad representation metric audit: %+v", a.Metric)
	}
	if !a.Boundary.TraceLedgerAvailable || !a.Boundary.KappaWhiteningCandidateAvailable || a.Boundary.TraceToKappaNativeMapDerived || a.Boundary.NativeKappaSelected || a.Boundary.NativeGaugeHessianSelected || a.Boundary.NativeWeakAngleDerived || a.Boundary.NativeWZMassesDerived {
		t.Fatalf("boundary over-promoted coefficient: %+v", a.Boundary)
	}
	if a.Firewall.WeakAngleImported || a.Firewall.NativeKappaWritten || a.Firewall.NativeGaugeHessianWritten || a.Firewall.NativeWZMassWritten {
		t.Fatalf("firewall leak: %+v", a.Firewall)
	}
}

func TestGate494RenderAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{
		"# Gate 494 Registry Audit",
		StatusHyperchargeTraceNormalizationFound,
		StatusU1CompletionKappaTargetInherited,
		StatusFailedTraceDoesNotSelectKappa,
		StatusKappaRegistryWriteBlocked,
		"k_Y = 5/3",
		"kappa_U1 = 6",
		"Gate 495",
	} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
