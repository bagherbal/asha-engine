package generation2stabilizerbranchfirstorderoperatorcalculationaudit

import (
	"strings"
	"testing"
)

func TestGate861FirstOrderAttemptedInOrientAlgebra(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.FirstOrder.Attempted || !a.FirstOrder.Gate860Inherited || a.FirstOrder.Algebra != "A_F^orient = C_R plus C_H plus M_3(C)" {
		t.Fatalf("bad first-order target: %s", FormatFirstOrder(a.FirstOrder))
	}
}

func TestGate861NonzeroDRhoIsOneFormNotFailure(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.FirstOrder.DRhoNonzeroAllowedOneForm || !containsAll(a.FirstOrder.Supports, []string{StatusNonzeroDRhoAllowed}) {
		t.Fatalf("D-rho one-form classification missing: %s", FormatFirstOrder(a.FirstOrder))
	}
}

func TestGate861ColorCentralityRemovesM3Pressure(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !colorEdgesCentral(a.Edges) || !a.FirstOrder.ColorObstructionRemoved {
		t.Fatalf("color centrality not installed: %s | %s", FormatEdges(a.Edges), FormatFirstOrder(a.FirstOrder))
	}
}

func TestGate861PunctureAndKernelPreserved(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !leptonAndPunctureOK(a.Edges) || !a.Kernel.PunctureZero || !a.Kernel.KernelPreserved || a.Kernel.KernelSingleton != "h_+ tensor P_1" {
		t.Fatalf("puncture/kernel not preserved: %s | %s", FormatEdges(a.Edges), FormatKernel(a.Kernel))
	}
}

func TestGate861SocketCharacterMatchRemainsSeal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.FirstOrder.SocketCharacterMatchOperatorCertified || !containsAll(a.FirstOrder.Failures, []string{FailureSocketCharacterSeal}) || !containsAll(a.FirstOrder.Supports, []string{SupportCharacterMatchNeeded}) {
		t.Fatalf("socket character firewall not preserved: %s", FormatFirstOrder(a.FirstOrder))
	}
}

func TestGate861NoMagnitudeOrLedgerPromotion(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range a.Edges {
		if e.NumericalValue || e.YukawaMagnitude {
			t.Fatalf("edge overpromoted: %+v", e)
		}
	}
	if !a.Ledger.OfficialFrozen || a.Ledger.AlphaNative || a.Ledger.R3 || a.Ledger.R4 || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 {
		t.Fatalf("ledger overpromoted: %s | %s", FormatLedger(a.Ledger), FormatImpact(a.Impact))
	}
}

func TestGate861Theorem(t *testing.T) {
	res := Generation2StabilizerBranchFirstOrderOperatorCalculationAuditTheorem().Verify()
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
