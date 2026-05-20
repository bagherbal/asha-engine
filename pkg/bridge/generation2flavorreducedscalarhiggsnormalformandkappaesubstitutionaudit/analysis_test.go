package generation2flavorreducedscalarhiggsnormalformandkappaesubstitutionaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate752FlavorReducedNormalForm(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate751.Inherited || !a.Gate751.TypedNormalFormReady || !a.Gate751.IllegalTermsRejected || !a.Gate751.KappaInsertionRecorded {
		t.Fatalf("bad Gate751 inheritance: %+v", a.Gate751)
	}
	if !strings.Contains(a.KappaERed.Formula, "sin²") || !strings.Contains(a.KappaERed.Formula, "xi_boundary") {
		t.Fatalf("bad kappa_e reduced formula: %+v", a.KappaERed)
	}
	if math.Abs(a.KappaERed.SSplit-0.0012924448188162962) > 1e-18 {
		t.Fatalf("bad S_split: %.18g", a.KappaERed.SSplit)
	}
	if math.Abs(a.KappaERed.KappaEReduced-0.005503554218475772) > 1e-18 {
		t.Fatalf("bad kappa_e_red: %.18g", a.KappaERed.KappaEReduced)
	}
	if math.Abs(a.KappaERed.Residual+2.6901212160646004e-11) > 1e-20 {
		t.Fatalf("bad kappa_e residual: %.18g", a.KappaERed.Residual)
	}
	if !strings.Contains(a.FWallRed.MapType, "Q_boundary -> Q_history") || !strings.Contains(a.FWallRed.Polynomial, "kappa_e_red") {
		t.Fatalf("bad reduced F_wall_3: %+v", a.FWallRed)
	}
	if math.Abs(a.RuntimeRed.RuntimeShift-1.3369860774048448e-13) > 5e-16 {
		t.Fatalf("bad runtime shift: %.18g", a.RuntimeRed.RuntimeShift)
	}
}

func TestGate752SensitivityAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Sensitivity.PK7SSquared-1.624013231638281e-7) > 1e-20 || math.Abs(a.Sensitivity.Agreement) > 1e-16 {
		t.Fatalf("bad sensitivity: %+v", a.Sensitivity)
	}
	if !a.Reduction.BareSealReduced || a.Reduction.NativeFlavorTheorem || len(a.Reduction.Components) != 4 {
		t.Fatalf("bad reduction status: %+v", a.Reduction)
	}
	if !a.Firewalls.KappaERedNativeBlocked || !a.Firewalls.PMNSCKMBlocked || !a.Firewalls.FlavorDeficitBlocked || !a.Firewalls.YukawaBlocked || !a.Firewalls.RuntimePredictionBlocked || !a.Firewalls.HiggsMassBlocked {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
	res := Generation2FlavorReducedScalarHiggsNormalFormAndKappaESubstitutionAuditTheorem().Verify()
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
