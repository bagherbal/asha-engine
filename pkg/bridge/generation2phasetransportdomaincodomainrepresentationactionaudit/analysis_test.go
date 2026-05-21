package generation2phasetransportdomaincodomainrepresentationactionaudit

import (
	"strings"
	"testing"
)

func TestGate904DomainTypedButNoCR2Action(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	d := a.Domain
	if !d.HopfS1Typed || !d.CL17ChiralityTyped || d.HopfActsOnCR2 || d.GammaChiActsOnCR2 || d.NativeDomainActionMap {
		t.Fatalf("bad domain audit: %s", FormatDomain(d))
	}
}

func TestGate904CodomainTypedAsRhoRProjectorPair(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	c := a.Codomain
	if !c.ProjectorPairTyped || !c.OutputsOrderedPair || c.NativeTransportToProjectors || !strings.Contains(c.RightCharacterSplit, "rho_R") {
		t.Fatalf("bad codomain audit: %s", FormatCodomain(c))
	}
}

func TestGate904ActionMapMissing(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	ac := a.Action
	if !ac.PositiveToEPlus || !ac.ConjugateToEMinus || ac.ActionCompatibleWithRhoR || ac.TypedActionOnRightPair {
		t.Fatalf("action incorrectly certified: %s", FormatAction(ac))
	}
}

func TestGate904NonCircularity(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	n := a.NonCircularity
	if !n.RhoRLabelsSockets || n.RhoRExplainsOrdering || n.TransportDefinedByLabels || !n.TargetLabelRestatement || n.NonCircularSourceCertified {
		t.Fatalf("noncircularity leak: %s", FormatNonCircularity(n))
	}
}

func TestGate904AirlockEffectAndMissingObject(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.AirlockEffect.IfTransportSealed || !a.AirlockEffect.OrdersNeutralPuncture || !a.AirlockEffect.CollapsesLocalWounds || a.AirlockEffect.NativeR3Promotion {
		t.Fatalf("bad airlock effect: %s", FormatAirlockEffect(a.AirlockEffect))
	}
	if a.MissingObject.MissingObject != TransportMap || !a.MissingObject.NowFullyTyped || a.MissingObject.NativeMapCertified || !strings.Contains(a.MissingObject.RequiredAction, "C_R^2") {
		t.Fatalf("bad missing object: %s", FormatMissingObject(a.MissingObject))
	}
}

func TestGate904FreezeAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Freeze.Frozen || !a.Freeze.DiagnosticOnly || a.Freeze.CanUpdate || !near(a.Freeze.OperatorNEff, OperatorNEffDiagnostic) || near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) {
		t.Fatalf("bad freeze: %s", FormatFreeze(a.Freeze))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate904Theorem(t *testing.T) {
	res := Generation2PhaseTransportDomainCodomainRepresentationActionAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
