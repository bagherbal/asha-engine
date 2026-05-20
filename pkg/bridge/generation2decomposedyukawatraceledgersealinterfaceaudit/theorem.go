package generation2decomposedyukawatraceledgersealinterfaceaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-794-DECOMPOSED-YUKAWA-TRACE-LEDGER-SEAL-INTERFACE"
	theoremName = "Gate 794 — DecomposedYukawaTraceLedgerSeal Specification and Data-Interface Audit"
)

func Generation2DecomposedYukawaTraceLedgerSealInterfaceAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 794 analysis", Passed: false, Detail: err.Error()}}, Notes: []string{StatusFirewallPreservedGate794}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 793 decomposed Yukawa trace audit", Passed: a.Gate793.Inherited && a.Gate793.NEffBottleneck && closeAbs(a.Gate793.AggregateA, aSnapshot, 1e-15) && closeAbs(a.Gate793.NEff, nEffSnapshot, 5e-16), Detail: a.Gate793.Verdict},
			{Name: "define decomposed Yukawa trace ledger seal", Passed: a.Seal.Defined && a.Seal.Name == "DecomposedYukawaTraceLedgerSeal" && containsAll(a.Seal.Components, []string{"sector_trace_ledger", "trace_atom_ledger", "top_channel_selector", "neutrino_sector_convention", "validation_rules"}), Detail: FormatSeal(a.Seal)},
			{Name: "specify sector trace interface", Passed: a.Sector.Specified && !a.Sector.DataAvailable && containsAll(a.Sector.QuadraticTraces, []string{"a_u", "a_d", "a_e", "a_nu"}) && containsAll(a.Sector.QuarticTraces, []string{"b_u", "b_d", "b_e", "b_nu"}) && containsAll(a.Sector.RequiredOutputs, []string{"a_sector/a", "b_sector/b"}), Detail: FormatSector(a.Sector)},
			{Name: "specify trace atom ledger interface", Passed: a.Atom.Specified && a.Atom.ColorRuleRequired && a.Atom.MixingForbidden && strings.Contains(a.Atom.ColorConvention, "never both") && containsAll(a.Atom.Fields, []string{"atom_id", "sector", "squared_singular_value", "scale"}) && containsAll(a.Atom.SumRules, []string{"sum_i x_i = a", "sum_i x_i^2 = b"}), Detail: FormatAtom(a.Atom)},
			{Name: "require color multiplicity rule", Passed: a.Atom.ColorRuleRequired && a.Atom.MixingForbidden && a.Atom.Verdict == StatusAtomsUnavailableFromAggregate, Detail: a.Atom.ColorConvention},
			{Name: "specify top channel selector interface", Passed: a.TopSelector.Specified && a.TopSelector.RequiresTypedT && !a.TopSelector.MayInvertNEff && containsAll(a.TopSelector.Formulas, []string{"a_top=3T", "alpha", "beta", "b/a^2"}), Detail: FormatTopSelector(a.TopSelector)},
			{Name: "define neutrino sector convention firewall", Passed: a.Neutrino.Defined && !a.Neutrino.ImplicitAllowed && containsAll(a.Neutrino.AllowedStatuses, []string{"Y_nu absent", "Y_nu zero", "Y_nu Dirac sealed", "Y_nu Majorana-effective", "Y_nu unknown"}), Detail: strings.Join(a.Neutrino.AllowedStatuses, "; ")},
			{Name: "specify scale and normalization interface", Passed: a.Scale.Specified && a.Scale.Scale == "M_Z" && a.Scale.Scheme == "supplied_or_unknown" && strings.Contains(a.Scale.Differential, "d ln N_eff") && !a.Scale.MultiScaleLedger, Detail: a.Scale.Differential},
			{Name: "define aggregate validation rules", Passed: a.Validation.Defined && a.Validation.RejectOnFail && containsAll(a.Validation.Rules, []string{"a_u+a_d+a_e+a_nu", "b_u+b_d+b_e+b_nu", "a^2/b", "sum_i x_i", "sum_i x_i^2"}), Detail: FormatValidation(a.Validation)},
			{Name: "define source output requirements", Passed: a.Output.Defined && a.Output.RequiresData && !a.Output.CanComputeNow && containsAll(a.Output.OutputsIfSupplied, []string{"sector fractions", "top/rest", "N_eff", "largest non-top"}), Detail: strings.Join(a.Output.OutputsIfSupplied, "; ")},
			{Name: "preserve triality and generation firewall", Passed: a.Triality.Preserved && !a.Triality.LedgerImpliesGeneration && !a.Triality.LedgerImpliesD4 && containsAll(a.Triality.RequiredForNativeGeneration, []string{"generation carrier", "trace-readout theorem", "breaking operator"}), Detail: a.Triality.Verdict},
			{Name: "record C_Higgs impact", Passed: a.Impact.Recorded && closeAbs(a.Impact.CYukawa, cYukawaSnapshot, 5e-16) && closeAbs(a.Impact.CHiggs, cHiggsSnapshot, 5e-16) && a.Impact.NEffAggregateSealed && a.Impact.NEffSectorAuditableIfDataValid && !a.Impact.CHiggsLevelC, Detail: FormatImpact(a.Impact)},
			{Name: "record branch decision", Passed: a.Branch.Recorded && !a.Branch.ValidatedLedgerExists && !a.Branch.D4PackageIntroduced && strings.Contains(a.Branch.Recommended, "Yukawa Trace Atom Data Acquisition") && containsAll(a.Branch.Alternatives, []string{"Sector Contribution", "D4 Triality"}), Detail: FormatBranch(a.Branch)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && !a.Firewalls.DecomposedLedgerNativeYukawa && !a.Firewalls.SectorTraceGenerationTheorem && !a.Firewalls.TopSelectorTopYukawaDerivation && !a.Firewalls.ValidatedAtomsPMNSCKM && !a.Firewalls.NEffD4Triality && !a.Firewalls.ScaleLocalScaleStable && !a.Firewalls.CHiggsPoleMass && !a.Firewalls.TreeProxyPoleMass && a.Firewalls.Verdict == StatusFirewallPreservedGate794, Detail: a.Firewalls.Verdict},
		}
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := append([]string{a.Truth, FormatSeal(a.Seal), FormatSector(a.Sector), FormatAtom(a.Atom), FormatTopSelector(a.TopSelector), FormatImpact(a.Impact), FormatBranch(a.Branch), a.FinalStatement}, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
