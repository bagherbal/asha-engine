package generation2higgshessianspectralprojectorandradialeventreplacementaudit

import (
	"math"
	"strings"
	"testing"
)

func closeTo(x, y float64) bool { return math.Abs(x-y) <= 1e-12 }

func TestGate768SpectralProjectorDefinition(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate767.Inherited || a.Gate767.AlignmentSeal != "HistoryLoopHessianRadialAlignmentSeal" || a.Gate767.AlignmentNative || a.Gate767.RankTraceIdentifiesSupport {
		t.Fatalf("bad Gate767 inheritance: %+v", a.Gate767)
	}
	if !strings.Contains(a.Spectral.PotentialLane, "U(2)-invariant") || !strings.Contains(a.Spectral.HessianFormula, "P_hessian") || !a.Spectral.PositiveRadialEigenvalue || a.Spectral.RadialEigenvalueFormula != "2 lambda v^2" || !closeTo(a.Spectral.RadialEigenvalueGeV2, 2*lambdaRuntimeEff*vevGate741GeV*vevGate741GeV) || len(a.Spectral.AngularEigenvalues) != angularZeroModes || a.Spectral.SupportProjector != "P_Hess=supp(H_V(x_0))" || !strings.Contains(a.Spectral.SupportProjectorFormula, "H_V(x_0)/Tr") || a.Spectral.SupportRank != hessianSupportRank || !a.Spectral.EqualsHessianProjector || a.Spectral.NativePotentialTheorem || a.Spectral.NativeVEVTheorem {
		t.Fatalf("bad spectral projector audit: %+v", a.Spectral)
	}
}

func TestGate768ReplacementAndHistoryLoopTrace(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.Replacement.Before, "independent") || !strings.Contains(a.Replacement.After, "P_rad := P_Hess") || !strings.Contains(a.Replacement.ReplacementScope, "potential lane only") || !a.Replacement.IndependentRadialSymbolReduced || !a.Replacement.RequiresSuppliedPotential || !a.Replacement.RequiresSuppliedVacuum || !a.Replacement.HistoryLoopAlignmentStillRequired || a.Replacement.NativeAlignmentTheorem {
		t.Fatalf("bad radial replacement: %+v", a.Replacement)
	}
	if a.HistoryLoop.State != "rho_plus=I_K7+/4" || !strings.Contains(a.HistoryLoop.Projector, "supp") || a.HistoryLoop.Rank != hessianSupportRank || !strings.Contains(a.HistoryLoop.TraceWeightFormula, "rank(P_Hess)/4") || !closeTo(a.HistoryLoop.TraceWeight, 0.25) || !strings.Contains(a.HistoryLoop.LHopfFormula, "supp(H_V") || !closeTo(a.HistoryLoop.LHopf, 1/(8*math.Pi)) || a.HistoryLoop.NativeHistoryLoopTheorem {
		t.Fatalf("bad HistoryLoop support trace: %+v", a.HistoryLoop)
	}
}

func TestGate768ThreeFactorAndUpgrade(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.ThreeFactor.Formula, "lambda_runtime_eff") || !strings.Contains(a.ThreeFactor.SupportForm, "supp(H_V") || !closeTo(a.ThreeFactor.LHopf, 1/(8*math.Pi)) || !closeTo(a.ThreeFactor.CHistory, 1.038025177923625) || !closeTo(a.ThreeFactor.LambdaRuntimeEff, lambdaRuntimeEff) || !a.ThreeFactor.RewritesOnly || a.ThreeFactor.IndependentRuntimeTheorem {
		t.Fatalf("bad three-factor rewrite: %+v", a.ThreeFactor)
	}
	if a.Upgrade.FromType != "supplied rank-one radial projector" || !strings.Contains(a.Upgrade.ToType, "Hessian spectral support") || !a.Upgrade.StrongerThanGate767 || !a.Upgrade.StillBridgeConditional || a.Upgrade.PotentialAndVacuumNative || a.Upgrade.HistoryLoopNative {
		t.Fatalf("bad source-type upgrade: %+v", a.Upgrade)
	}
}

func TestGate768FirewallsAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Obstruction.PotentialDerived || a.Obstruction.VacuumDerived || a.Obstruction.HistoryLoopRuleDerived || a.Obstruction.HistoryLoopUsesHessianSupportDerived || a.Obstruction.PoleMassDerived || a.Obstruction.YukawaDerived {
		t.Fatalf("bad obstruction audit: %+v", a.Obstruction)
	}
	if !a.Firewalls.Audited || a.Firewalls.NativePotentialTheorem || a.Firewalls.NativeVEVTheorem || a.Firewalls.NativeHistoryLoopHessianAlignment || a.Firewalls.NativeHistoryLoopUnitTheorem || a.Firewalls.TreeProxyPoleMassTheorem || a.Firewalls.HiggsMassOrPoleMassTheorem || a.Firewalls.YukawaOperatorOrEigenvalueTheorem {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
	res := Generation2HiggsHessianSpectralProjectorAndRadialEventReplacementAuditTheorem().Verify()
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
