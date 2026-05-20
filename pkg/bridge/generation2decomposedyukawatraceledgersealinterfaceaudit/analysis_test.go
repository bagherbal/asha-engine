package generation2decomposedyukawatraceledgersealinterfaceaudit

import (
	"strings"
	"testing"
)

func TestGate794SealAndSectorInterfaces(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate793.Inherited || !a.Gate793.NEffBottleneck || !closeAbs(a.Gate793.NEff, 3.0023273474722147, 5e-16) {
		t.Fatalf("bad inheritance: %+v", a.Gate793)
	}
	if !a.Seal.Defined || a.Seal.Name != "DecomposedYukawaTraceLedgerSeal" || !containsAll(a.Seal.Components, []string{"sector_trace_ledger", "trace_atom_ledger", "top_channel_selector", "neutrino_sector_convention", "validation_rules"}) {
		t.Fatalf("bad seal: %s", FormatSeal(a.Seal))
	}
	if !a.Sector.Specified || a.Sector.DataAvailable || !containsAll(a.Sector.QuadraticTraces, []string{"a_u", "a_d", "a_e", "a_nu"}) || !containsAll(a.Sector.QuarticTraces, []string{"b_u", "b_d", "b_e", "b_nu"}) {
		t.Fatalf("bad sector interface: %s", FormatSector(a.Sector))
	}
}

func TestGate794AtomTopNeutrinoAndScaleInterfaces(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Atom.Specified || !a.Atom.ColorRuleRequired || !a.Atom.MixingForbidden || !strings.Contains(a.Atom.ColorConvention, "never both") || !containsAll(a.Atom.Fields, []string{"atom_id", "sector", "squared_singular_value", "convention"}) {
		t.Fatalf("bad atom interface: %s", FormatAtom(a.Atom))
	}
	if !a.TopSelector.Specified || !a.TopSelector.RequiresTypedT || a.TopSelector.MayInvertNEff || !containsAll(a.TopSelector.Formulas, []string{"a_top=3T", "alpha", "beta"}) {
		t.Fatalf("bad top selector: %s", FormatTopSelector(a.TopSelector))
	}
	if !a.Neutrino.Defined || a.Neutrino.ImplicitAllowed || !containsAll(a.Neutrino.AllowedStatuses, []string{"Y_nu absent", "Y_nu zero", "Y_nu Dirac sealed", "Y_nu Majorana-effective", "Y_nu unknown"}) {
		t.Fatalf("bad neutrino firewall: %+v", a.Neutrino)
	}
	if !a.Scale.Specified || a.Scale.Scale != "M_Z" || !strings.Contains(a.Scale.Differential, "2 d ln a - d ln b") || a.Scale.MultiScaleLedger {
		t.Fatalf("bad scale interface: %+v", a.Scale)
	}
}

func TestGate794ValidationOutputImpactBranchAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Validation.Defined || !a.Validation.RejectOnFail || !containsAll(a.Validation.Rules, []string{"a_u+a_d+a_e+a_nu", "sum_i x_i", "N_eff_inherited"}) {
		t.Fatalf("bad validation: %s", FormatValidation(a.Validation))
	}
	if !a.Output.Defined || !a.Output.RequiresData || a.Output.CanComputeNow || !containsAll(a.Output.OutputsIfSupplied, []string{"sector fractions", "top/rest", "largest non-top"}) {
		t.Fatalf("bad output: %+v", a.Output)
	}
	if !a.Triality.Preserved || a.Triality.LedgerImpliesGeneration || a.Triality.LedgerImpliesD4 {
		t.Fatalf("bad triality firewall: %+v", a.Triality)
	}
	if !a.Impact.Recorded || !closeAbs(a.Impact.CYukawa, 0.9992248188812008, 5e-16) || !closeAbs(a.Impact.CHiggs, 1.0372205204048603, 5e-16) || !a.Impact.NEffAggregateSealed || !a.Impact.NEffSectorAuditableIfDataValid || a.Impact.CHiggsLevelC {
		t.Fatalf("bad impact: %s", FormatImpact(a.Impact))
	}
	if !a.Branch.Recorded || !strings.Contains(a.Branch.Recommended, "Yukawa Trace Atom Data Acquisition") || !containsAll(a.Branch.Alternatives, []string{"Sector Contribution", "D4 Triality"}) {
		t.Fatalf("bad branch: %s", FormatBranch(a.Branch))
	}
	if !a.Firewalls.Enforced || a.Firewalls.DecomposedLedgerNativeYukawa || a.Firewalls.SectorTraceGenerationTheorem || a.Firewalls.TopSelectorTopYukawaDerivation || a.Firewalls.ValidatedAtomsPMNSCKM || a.Firewalls.NEffD4Triality || a.Firewalls.ScaleLocalScaleStable || a.Firewalls.CHiggsPoleMass || a.Firewalls.TreeProxyPoleMass {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
}

func TestGate794TheoremStatusesAndFinalStatement(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.FinalStatement, "does not decompose N_eff yet") || !strings.Contains(a.FinalStatement, "DecomposedYukawaTraceLedgerSeal") || !strings.Contains(a.FinalStatement, "Gate 795") {
		t.Fatalf("bad final statement: %s", a.FinalStatement)
	}
	res := Generation2DecomposedYukawaTraceLedgerSealInterfaceAuditTheorem().Verify()
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
