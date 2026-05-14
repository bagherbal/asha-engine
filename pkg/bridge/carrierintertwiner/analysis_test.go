package carrierintertwiner

import "testing"

func TestGate255CarrierIntertwinerAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Summary.Gate254Inherited || !a.Summary.SCCarrierKnown || !a.Summary.LocalActionsAudited {
		t.Fatalf("expected predecessor and local carrier data: %s", FormatSummary(a.Summary))
	}
	if a.Summary.CommonCarrierDerived || a.Summary.CarrierIntertwinerDerived || a.Summary.UnifiedLedgerConstructed {
		t.Fatalf("Gate 255 must not construct the missing carrier/unified ledger: %s", FormatSummary(a.Summary))
	}
	if a.Intertwiners.LawfulCommonIntertwiner || a.Intertwiners.JointIntertwiningCandidates != 0 {
		t.Fatalf("unexpected lawful common intertwiner: %s", FormatIntertwiners(a.Intertwiners))
	}
	if a.SO8.T3LSO8Coordinates || a.SO8.YPhiSO8Coordinates || a.TrialityKernel.Q8vCConstructed || a.TrialityKernel.NeutralThreePlaneDerived {
		t.Fatalf("downstream pullback must remain blocked: %s / %s", FormatSO8(a.SO8), FormatTrialityKernel(a.TrialityKernel))
	}
	if a.Firewall.EmbeddedHphiIntoSCByDimension || a.Firewall.EmbeddedLeftDoubletByLabel || a.Firewall.TreatedTensorProductAsSC || a.Firewall.TreatedDirectSumAsIntertwiner || a.Firewall.PollutedFiniteCore {
		t.Fatalf("firewall violation: %s", FormatFirewall(a.Firewall))
	}
}
