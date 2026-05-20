package generation2colorcolorlessfinitediractensioncableaudit

import (
	"strings"
	"testing"
)

func TestGate598SectorSplitAndEdges(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited.ChargedLeptonSeal != "ChargedLeptonRootChamberSeal" || !a.Inherited.EnvironmentalOnly {
		t.Fatalf("bad inherited Gate597 state: %+v", a.Inherited)
	}
	if a.SectorSplit.InterSectorDFBlock || len(a.SectorSplit.Rows) < 3 {
		t.Fatalf("bad sector split: %+v", a.SectorSplit)
	}
	if len(a.Edges.Edges) != 4 || !allEdgesBlockSeparated(a.Edges) {
		t.Fatalf("bad edge inventory: %+v", a.Edges)
	}
}

func TestGate598CandidateInvariantsAndRootObstruction(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Candidates.AnyNativeTensionCable || !a.Candidates.ColorColorlessStructureVisible {
		t.Fatalf("unexpected candidate verdict: %+v", a.Candidates)
	}
	joined := FormatCandidateTable(a.Candidates)
	for _, want := range []string{StatusPolynomialNoFourthRoot, StatusNativeSpectralActionPowerSumCable, StatusSpectralActionTraceCableNotKoide, StatusDetPfaffianNoRootTrace, StatusQuarkJNatural, StatusNoColorColorlessTensionCableFound} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing candidate status %s", want)
		}
	}
	if a.RootObstruction.Gate596Avoided {
		t.Fatalf("Gate596 obstruction unexpectedly avoided: %+v", a.RootObstruction)
	}
	root := FormatRootLedger(a.RootObstruction)
	for _, want := range []string{StatusNativeSpectralActionPowerSumCable, StatusSpectralActionTraceCableNotKoide, StatusNoHeOneFourth, StatusNoEpsilonHE, StatusNoBFlavZero} {
		if !strings.Contains(root, want) {
			t.Fatalf("missing root status %s", want)
		}
	}
}

func TestGate598TheoremAndFirewalls(t *testing.T) {
	th := Generation2ColorColorlessFiniteDiracTensionCableAuditTheorem()
	res := th.Verify()
	if !res.Passed() {
		t.Fatalf("theorem checks failed: %+v", res)
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusColorColorlessStructureVisible, StatusColorColorlessTraceCableVisible, StatusNativeSpectralActionPowerSumCable, StatusNoRootChamberNativePromotion, StatusNoRootOrientationCableFound, StatusNoColorColorlessTensionCableFound, StatusGate352Preserved, StatusGate596Preserved, StatusGate598Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
