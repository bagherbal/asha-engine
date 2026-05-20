package generation2linearresponsefunctionalandtracepairingnormalizationaudit

import (
	"math"
	"strings"
	"testing"
)

func TestInheritanceAndDiscipline(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.Gate689FirstTraceSelectionInherited || !a.Inherited.Gate690ResidualStatusInherited {
		t.Fatalf("missing inheritance: %+v", a.Inherited)
	}
	if a.Inherited.Operator != "R_split = S_split P_K7" || a.Inherited.H72Dimension != h72Dimension || a.Inherited.K7Dimension != k7Dimension {
		t.Fatalf("bad inherited operator/dimensions: %+v", a.Inherited)
	}
	if !a.Inherited.QuadraticResidualClueRetained || a.Inherited.QuadraticCorrectionPromoted || a.Inherited.NativeSpectralExpansionTheorem || a.Inherited.NativeFirstTraceTheorem || a.Inherited.NativeSevenOver72Theorem {
		t.Fatalf("residual theorem promotion leaked into Gate691: %+v", a.Inherited)
	}
	if a.Discipline.ClaimsUniqueFullH72Observer || a.Discipline.ClaimsNativeLinearResponseTheorem || a.Discipline.ClaimsNativeFirstTraceTheorem || a.Discipline.ClaimsNativeSevenOver72 || a.Discipline.PromotesQuadraticResidualCorrection {
		t.Fatalf("discipline firewall violated: %+v", a.Discipline)
	}
}

func TestNormalizedTracePairing(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	want := (7.0 / 72.0) * a.Inherited.SSplit
	if math.Abs(a.Pairing.Value-want) > pairingTolerance || math.Abs(a.Pairing.Value-a.Inherited.F1) > pairingTolerance {
		t.Fatalf("bad pairing value: %+v", a.Pairing)
	}
	if !a.Pairing.EqualsFirstTrace || !a.Pairing.LinearInResponse || !a.Pairing.LinearInSSplit || !a.Pairing.BilinearInArguments {
		t.Fatalf("bad pairing classification: %+v", a.Pairing)
	}
	if !strings.Contains(a.Pairing.Verdict, StatusActiveBridgeRewrittenAsTracePairing) {
		t.Fatalf("missing trace-pairing verdict: %+v", a.Pairing)
	}
}

func TestRolesAndObserverPairings(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Roles.FullObserverTypeCorrect || !a.Roles.ResponseSupportSelected || !a.Roles.BoundaryScalarIsEigenvalue {
		t.Fatalf("bad role classification: %+v", a.Roles)
	}
	if a.Observers.CandidateCount != 5 || a.Observers.PositiveIdentityOnK7Count != 4 {
		t.Fatalf("bad observer count: %+v", a.Observers)
	}
	if !a.Observers.AllPositiveK7ObserversGiveSameValue || !a.Observers.SignedPolarityObserverInactive || a.Observers.FullH72ObserverUnique {
		t.Fatalf("bad observer degeneracy status: %+v", a.Observers)
	}
	if !observerValuesPass(a) {
		t.Fatalf("observer values do not match expected active/signed split: %+v", a.Observers)
	}
}

func TestLinearResponseAndResidualStatus(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.LinearResponse.DBaseLinearInWallCoordinates || !a.LinearResponse.TracePairingLinearInResponse || !a.LinearResponse.TracePairingLinearInSSplit || !a.LinearResponse.MatchesWallCoordinateOrder {
		t.Fatalf("bad linear response status: %+v", a.LinearResponse)
	}
	if a.LinearResponse.NativeLinearResponseFunctionalTheorem {
		t.Fatalf("linear response theorem should remain missing: %+v", a.LinearResponse)
	}
	if math.Abs(a.Residual.E1-8.525834398014336e-10) > residualTolerance {
		t.Fatalf("bad inherited residual: %+v", a.Residual)
	}
	if !a.Residual.QuadraticResidualClueRetained || a.Residual.QuadraticCorrectionPromoted || a.Residual.NativeSpectralExpansionTheorem {
		t.Fatalf("quadratic residual status promoted incorrectly: %+v", a.Residual)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2LinearResponseFunctionalAndTracePairingNormalizationAuditTheorem().Verify()
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
