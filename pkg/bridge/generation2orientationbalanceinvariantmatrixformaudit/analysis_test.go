package generation2orientationbalanceinvariantmatrixformaudit

import (
	"strings"
	"testing"
)

func TestGate593InvariantMatrixForm(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Final.InvariantFormAvailable || a.Balance.LeftKappa != a.Inherited.KappaObs || a.Balance.RightProjectorMinusCKM != a.Inherited.OrientationCandidate {
		t.Fatalf("expected invariant form to inherit Gate592 relation exactly: balance=%+v inherited=%+v final=%+v", a.Balance, a.Inherited, a.Final)
	}
	if !strings.Contains(a.Balance.EquationProjector, "Tr(P_e U_PMNS P_3^nu U_PMNS^dagger)") {
		t.Fatalf("projector trace form missing: %s", a.Balance.EquationProjector)
	}
}

func TestGate593RejectsNativeOperator(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Availability.AnyNativeBalanceOperator || a.Availability.AnyNativeRootSpectrumMap || a.Availability.AnyNativeFlavorCommutatorMap || a.Target.NativePresent || a.Final.NativeOperatorPresent {
		t.Fatalf("unexpected native operator: availability=%+v target=%+v final=%+v", a.Availability, a.Target, a.Final)
	}
	if !a.RootSpace.Gate352ObstructionPreserved || a.RootSpace.NativeRootTraceOperatorPresent || a.RootSpace.NativeAbsoluteDiracPresent {
		t.Fatalf("Gate352/root-trace firewall failed: %+v", a.RootSpace)
	}
}

func TestGate593TheoremAndStatuses(t *testing.T) {
	th := Generation2OrientationBalanceInvariantMatrixFormAuditTheorem()
	res := th.Verify()
	if !res.Passed() {
		t.Fatalf("theorem checks failed: %+v", res)
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusInvariantBalanceWritten, StatusPMNSProjectorTraceDefined, StatusCKMJarlskogFormsRecorded, StatusNoCrossSectorTraceOperator, StatusSealEnvironmental, StatusGate593Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
