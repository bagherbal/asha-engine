package higgsonshellrenormalizationscheme

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Inputs.HighestInheritedGate != inheritedHighestGate || a.Inputs.RequiredRePiGeV2 <= 40 || a.Inputs.FiniteRemainderGeV2 <= 1000 {
		t.Fatalf("bad inputs: %s", FormatInputs(a.Inputs))
	}
}

func TestPVStructureAndBlocks(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if len(a.PV.Definitions) != 3 || a.PV.PoleEquation == "" {
		t.Fatalf("bad PV structure: %s", FormatPVStructure(a.PV))
	}
	if len(a.Blocks.Blocks) != 4 || !allRealFiniteBlocks(a.Blocks) {
		t.Fatalf("bad PV blocks: %s", FormatPVBlocks(a.Blocks))
	}
	if !nearlyEqual(a.Blocks.Blocks[0].B0Finite, -0.550445463802, 1e-9) {
		t.Fatalf("unexpected top B0: %s", FormatPVBlock(a.Blocks.Blocks[0]))
	}
	if !nearlyEqual(a.Blocks.Blocks[1].A0FiniteGeV2, 12194.7503755795, 1e-6) {
		t.Fatalf("unexpected W A0: %s", FormatPVBlock(a.Blocks.Blocks[1]))
	}
}

func TestSchemesAndCounterterm(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !hasOnShellAndMSBar(a.Schemes) {
		t.Fatalf("scheme audit did not include OS and MS-bar lanes: %s", FormatSchemes(a.Schemes))
	}
	if !nearlyEqual(a.Counterterm.RequiredFiniteRemainder, 1035.17147945909, 1e-6) {
		t.Fatalf("unexpected finite remainder: %s", FormatCounterterm(a.Counterterm))
	}
	if a.Counterterm.RemainderOverTarget <= 20 || a.Counterterm.RemainderOverRawAbs <= 1 {
		t.Fatalf("bad counterterm ratios: %s", FormatCounterterm(a.Counterterm))
	}
}

func TestAlignmentAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Alignment.UVBoundaryFixed || !a.Alignment.ContactShapeImmutable || a.Alignment.IRSchemeSelected || a.Alignment.NativeCountertermFound {
		t.Fatalf("bad alignment: %s", FormatAlignment(a.Alignment))
	}
	if !a.Firewalls.NoFullSMCoefficientTable || !a.Firewalls.NoNativeCounterterms || !a.Firewalls.NoExactPoleMassClaim {
		t.Fatalf("bad firewalls: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	statuses := Statuses(a)
	required := []string{StatusPVStructureFormalized, StatusRenormalizationSchemeAudited, StatusCountertermTargetMapped, StatusFailedCountertermsDerived, StatusFailedExactPoleMassClaim}
	for _, req := range required {
		found := false
		for _, got := range statuses {
			if got == req {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("missing status %s in %v", req, statuses)
		}
	}
}

func TestTheoremPasses(t *testing.T) {
	res := OnShellRenormalizationSchemePassarinoVeltmanPoleMatchingAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
