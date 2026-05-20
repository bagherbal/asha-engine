package generation2generationoperatorsealandyukawamatrixsourceminimalityaudit

import (
	"strings"
	"testing"
)

func TestGate806InheritanceAndSealDefinition(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate805NoGo || !strings.Contains(a.Inheritance.FactorizedNormal, "Edge_f ⊗ Y_f") {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Seal.Defined || a.Seal.Name != "GenerationOperatorSeal" {
		t.Fatalf("bad seal name: %s", FormatSeal(a.Seal))
	}
	if !containsAll(a.Seal.Components, []string{"G_gen", "sector Yukawa operators Y_u,Y_d,Y_e,Y_nu", "Hermitian trace operators H_f=Y_f†Y_f", "PMNS/CKM misalignment readouts", "hierarchy/breaking operator", "noncircularity proof"}) {
		t.Fatalf("bad seal components: %s", FormatSeal(a.Seal))
	}
	if !containsAll(a.Seal.Failures, []string{StatusNoNativeGenerationOperator}) {
		t.Fatalf("missing seal failure: %s", FormatSeal(a.Seal))
	}
}

func TestGate806MinimalityAndReadoutLayerSeparation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	minimal := FormatMinimality(a.Minimality)
	for _, want := range []string{"remove G_gen", "remove Y_f", "remove H_f", "remove singular values", "remove diagonalization frames", "remove hierarchy/breaking operator", StatusCannotCompress} {
		if !strings.Contains(minimal, want) {
			t.Fatalf("minimality missing %s: %s", want, minimal)
		}
	}
	if !a.Layers.Separated || !containsAll(a.Layers.Supports, []string{StatusNEffNeedsSpectra, StatusKappaOrientNeedsFrames}) || !containsAll(a.Layers.Failures, []string{StatusSingularNoMixing}) {
		t.Fatalf("bad readout layers: %s", FormatLayers(a.Layers))
	}
}

func TestGate806SourceAuditsAndNormalForm(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !containsAll(a.FST.Supports, []string{StatusFSTSuppliesEdgeDomain}) || !containsAll(a.FST.Failures, []string{StatusFSTNoYF}) {
		t.Fatalf("bad FST source: %s", FormatSource(a.FST))
	}
	if !containsAll(a.TD4.Failures, []string{StatusTD4NoGenOperator, StatusTD4NoTraceAtoms}) {
		t.Fatalf("bad TD4 source: %s", FormatSource(a.TD4))
	}
	if !containsAll(a.Aggregate.Failures, []string{StatusAggregateNoOperator}) {
		t.Fatalf("bad aggregate source: %s", FormatSource(a.Aggregate))
	}
	if !containsAll(a.External.Supports, []string{StatusExternalCanPopulateSeal}) || !containsAll(a.External.Failures, []string{StatusExternalNotNative}) {
		t.Fatalf("bad external source: %s", FormatSource(a.External))
	}
	if !containsAll(a.K7Projective.Failures, []string{StatusK7NotOperator, StatusProjectiveNotSource}) {
		t.Fatalf("bad K7/projective source: %s", FormatSource(a.K7Projective))
	}
	if !a.NormalForm.Recorded || !containsAll(a.NormalForm.Forms, []string{"D_u = Edge_u ⊗ Y_u", "D_d = Edge_d ⊗ Y_d", "D_e = Edge_e ⊗ Y_e", "D_nu = Edge_nu ⊗ Y_nu"}) {
		t.Fatalf("bad normal form: %s", FormatNormal(a.NormalForm))
	}
}

func TestGate806ObstructionsCHiggsBranchAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !containsAll(a.Trace.Failures, []string{StatusCarrierAloneNoSpectra, StatusGenerationsAloneNoNEff, StatusNoTraceAtomsWithoutH}) {
		t.Fatalf("bad trace obstruction: %s", FormatObstruction(a.Trace))
	}
	if !containsAll(a.Mixing.Failures, []string{StatusTraceLedgerNoKappaOrient, StatusNoPMNSCKMFrames}) {
		t.Fatalf("bad mixing obstruction: %s", FormatObstruction(a.Mixing))
	}
	if !containsAll(a.Hierarchy.Failures, []string{StatusNoTopDominanceOperator, StatusNoLightSuppressionOperator, StatusNoNEffMinusThreeSource}) {
		t.Fatalf("bad hierarchy obstruction: %s", FormatObstruction(a.Hierarchy))
	}
	if !a.CHiggs.Preserved || !containsAll(a.CHiggs.Failures, []string{StatusNoCYukawaUpdate, StatusCHiggsLevelB}) {
		t.Fatalf("bad C_Higgs firewall: %s", FormatCHiggs(a.CHiggs))
	}
	if !strings.Contains(a.Branch.Next, "Gate 807") || !strings.Contains(a.Branch.Next, "TraceMagnitudeOperatorSeal") {
		t.Fatalf("bad branch: %+v", a.Branch)
	}
	res := Generation2GenerationOperatorSealAndYukawaMatrixSourceMinimalityAuditTheorem().Verify()
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
			t.Fatalf("missing status %s", want)
		}
	}
}
