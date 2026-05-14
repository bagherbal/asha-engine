package baryonleptonoperatoraudit

import "testing"

func TestInventoryKeepsLeptoquarkSlotsUnactivated(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault error: %v", err)
	}
	if a.Inventory.U4Dimension != 16 || !a.Inventory.DecompositionComplete {
		t.Fatalf("bad u(4) inventory: %+v", a.Inventory)
	}
	if !a.Inventory.ContainsQuarkLeptonCurrentSlots {
		t.Fatalf("expected u(4) leptoquark current inventory to be present")
	}
	if a.Inventory.LeptoquarkSlotsGaugeActivated || a.Inventory.ContactConnectionHasLeptoquark {
		t.Fatalf("leptoquark slots must remain unactivated: %+v", a.Inventory)
	}
}

func TestBMinusLDoesNotFakeForbidStandardTemplates(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault error: %v", err)
	}
	if !a.OperatorSearch.SMGaugeInvariantTemplatesExist {
		t.Fatalf("expected external SM-gauge-invariant proton-decay templates to be audited")
	}
	for _, tpl := range a.OperatorSearch.Templates {
		if tpl.DeltaBMinusL != 0 {
			t.Fatalf("template %s should preserve B-L in this audit, got Δ(B-L)=%d", tpl.Name, tpl.DeltaBMinusL)
		}
		if tpl.ConstructedByFiniteAlgebra {
			t.Fatalf("no template should be finite-constructed yet: %+v", tpl)
		}
	}
	if a.Conservation.BMinusLForbidsQQQL || a.Inventory.BMinusLConservationAloneForbidsPD {
		t.Fatalf("B-L must not be used as a fake proton-decay firewall")
	}
}

func TestCurrentConnectionStabilityButNoAbsoluteBaryonTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault error: %v", err)
	}
	if !a.Conservation.CurrentConnectionProtonStable || !a.Conservation.ProtonDecayChannelConstructionFailed {
		t.Fatalf("expected current-connection proton stability obstruction: %+v", a.Conservation)
	}
	if a.Conservation.ExactBaryonConservationProved || a.Conservation.ExactLeptonConservationProved {
		t.Fatalf("absolute B/L conservation must not be overclaimed: %+v", a.Conservation)
	}
	if !a.Conservation.LeptoquarkInventoryPreventsAbsoluteNoGo {
		t.Fatalf("u(4) leptoquark inventory should keep the stronger theorem open")
	}
}

func TestTheoremVerifierRecordsFailedRouteWithoutFailedChecks(t *testing.T) {
	res := BaryonLeptonViolatingOperatorBasisAuditTheorem().Run()
	if string(res.Status) != "FAILED_ROUTE" {
		t.Fatalf("status = %s", res.Status)
	}
	for _, chk := range res.Checks {
		if !chk.Passed {
			t.Fatalf("check %q failed: %s", chk.Name, chk.Detail)
		}
	}
}
