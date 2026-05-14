package leptoquarkdynamicsseal

import "testing"

func TestNativeLeptoquarkSlotsStayKinematic(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault error: %v", err)
	}
	if a.Dynamics.SlotsAudited != 6 || a.Dynamics.OffDiagonalDimension != 6 {
		t.Fatalf("expected six leptoquark slots, got %+v", a.Dynamics)
	}
	if !a.Dynamics.AllSlotsKinematicOnly || a.Dynamics.AnyDynamicMediator {
		t.Fatalf("slots must remain kinematic-only: %+v", a.Dynamics)
	}
	for _, slot := range a.Dynamics.Slots {
		if slot.GaugeCurvatureDerived || slot.FiniteActionDerived || slot.LocalFieldMapDerived || slot.PropagatorDerived || slot.MassScaleDerived || slot.CouplingCoefficientKnown || slot.DynamicMediator {
			t.Fatalf("slot was accidentally activated: %+v", slot)
		}
	}
}

func TestSealQuarantinesDynamicsAndOperators(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault error: %v", err)
	}
	if !a.Seal.Active || !a.Seal.Conditional {
		t.Fatalf("seal should be active and conditional: %+v", a.Seal)
	}
	if !a.Seal.ForbidsGaugeActivation || !a.Seal.ForbidsPropagatorUse || !a.Seal.ForbidsOperatorCoefficient || !a.Seal.ForbidsLifetimeFormula {
		t.Fatalf("seal does not quarantine all dangerous semantics: %+v", a.Seal)
	}
	if !a.Operators.SealedObstruction || a.Operators.AnyTemplateConstructible || a.Operators.AnySuppressionScaleComputed {
		t.Fatalf("operators should remain obstructed under seal: %+v", a.Operators)
	}
}

func TestBMinusLPreservingTemplatesDoNotBecomeConstructible(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault error: %v", err)
	}
	if !a.Operators.AllTemplatesBMinusLPreserving {
		t.Fatalf("expected standard templates to preserve B-L")
	}
	for _, tpl := range a.Operators.Templates {
		if tpl.DeltaBMinusL != 0 || !tpl.BMinusLPreserving {
			t.Fatalf("template should preserve B-L: %+v", tpl)
		}
		if tpl.ConstructibleUnderSeal || tpl.SuppressionScaleComputed {
			t.Fatalf("template must remain blocked under seal: %+v", tpl)
		}
	}
}

func TestSealedConnectionBaryonConservationIsConditional(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault error: %v", err)
	}
	if !a.Conservation.SealedConnectionBaryonConservation || !a.Conservation.ConditionalOnLeptoquarkSeal {
		t.Fatalf("expected conditional sealed theorem: %+v", a.Conservation)
	}
	if a.Conservation.AbsoluteUnsealedBaryonTheorem {
		t.Fatalf("must not overclaim unsealed absolute baryon conservation: %+v", a.Conservation)
	}
	if !a.Conservation.ProtonLifetimeStrictlyObstructed {
		t.Fatalf("proton lifetime should remain obstructed")
	}
}

func TestTheoremVerifierRecordsPhenomenologyWithoutFailedChecks(t *testing.T) {
	res := LeptoquarkDynamicsSealBaryonConservationTheorem().Run()
	if string(res.Status) != "PHENOMENOLOGY" {
		t.Fatalf("status = %s", res.Status)
	}
	for _, chk := range res.Checks {
		if !chk.Passed {
			t.Fatalf("check %q failed: %s", chk.Name, chk.Detail)
		}
	}
}
