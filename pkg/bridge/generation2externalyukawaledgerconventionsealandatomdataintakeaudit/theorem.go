package generation2externalyukawaledgerconventionsealandatomdataintakeaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-796-EXTERNAL-YUKAWA-LEDGER-CONVENTION-SEAL-ATOM-DATA-INTAKE"
	theoremName = "Gate 796 — External Yukawa Ledger Convention Seal and Atom Data Intake Audit"
)

func Generation2ExternalYukawaLedgerConventionSealAndAtomDataIntakeAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 796 analysis", Passed: false, Detail: err.Error()}}, Notes: []string{StatusFirewallPreservedGate796}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 795 non-identifiability audit", Passed: a.Gate795.Inherited && a.Gate795.Verdict == StatusGate795Inherited, Detail: a.Gate795.Verdict},
			{Name: "define external Yukawa ledger convention seal", Passed: a.ExternalSeal.Defined && containsAll(a.ExternalSeal.Fields, []string{"source_label", "scale_mu", "renormalization_scheme", "sector_singular_values", "neutrino_convention", "validation_against_aggregate"}), Detail: FormatExternalSeal(a.ExternalSeal)},
			{Name: "define circular intake firewall", Passed: a.Circular.Defined && !a.Circular.UsesHiggsBacksolve && containsAll(a.Circular.ForbiddenSources, []string{"N_eff", "C_Higgs", "m_H_pole", "observed Higgs"}) && containsAll(a.Circular.ForbiddenOperations, []string{"choose T", "choose rest atoms", "Higgs pole mass"}), Detail: FormatCircular(a.Circular)},
			{Name: "define trace atom construction rules and color convention", Passed: a.AtomRules.Defined && a.AtomRules.AtomFormula == "x_i=y_i^2" && a.AtomRules.RequiresExactlyOneColorRule && strings.Contains(a.AtomRules.CoefficientColorConvention, "coefficient 3") && strings.Contains(a.AtomRules.RepeatedAtomConvention, "repeated"), Detail: FormatAtomRules(a.AtomRules)},
			{Name: "define Yukawa atom input schema", Passed: a.Schema.Defined && a.Schema.RequiresLabels && containsAll(a.Schema.Fields, []string{"fermion_label", "sector", "generation_label", "y_value", "uncertainty"}) && containsAll(a.Schema.RequiredSectors, []string{"up", "down", "charged_lepton", "neutrino"}), Detail: FormatSchema(a.Schema)},
			{Name: "audit neutrino convention requirement", Passed: a.Neutrino.Audited && a.Neutrino.ExplicitRequired && containsAll(a.Neutrino.AllowedStatuses, []string{"absent", "zero", "Dirac", "Majorana", "unknown"}) && a.Neutrino.Verdict == StatusNeutrinoMustBeExplicit, Detail: a.Neutrino.ActiveStatus},
			{Name: "define aggregate validation rules", Passed: a.Validation.Defined && closeAbs(a.Validation.InheritedA, aInherited, 1e-16) && closeAbs(a.Validation.InheritedB, bInherited, 1e-16) && closeAbs(a.Validation.InheritedNEff, nEffInherited, 5e-16) && a.Validation.MustValidateA && a.Validation.MustValidateB && a.Validation.MustValidateNEff && !a.Validation.SilentRescaleAllowed, Detail: FormatValidation(a.Validation)},
			{Name: "define top-channel selector rules", Passed: a.TopChannel.Defined && a.TopChannel.RequiresTypedTop && !a.TopChannel.MayInferFromNEff && containsAll(a.TopChannel.Formulae, []string{"T=y_t^2", "alpha", "beta", "b/a^2"}), Detail: FormatTopChannel(a.TopChannel)},
			{Name: "define sector contribution outputs but block without ledger", Passed: a.SectorOutputs.Defined && !a.SectorOutputs.ExternalLedgerSupplied && !a.SectorOutputs.CanOutputNow && containsAll(a.SectorOutputs.Outputs, []string{"a_u/a", "b_u/b", "largest atoms", "top dominance", "neutrino"}), Detail: a.SectorOutputs.Verdict},
			{Name: "define scale stability intake rules", Passed: a.Scale.Defined && a.Scale.AllowsSingleScale && a.Scale.AllowsMultiScale && !a.Scale.MultiScaleCertified && a.Scale.Verdict == StatusSingleScaleLocal, Detail: a.Scale.CurrentExternalData},
			{Name: "record Level-B C_Higgs impact", Passed: a.Impact.Recorded && a.Impact.ValidatedExternalWouldImprove && !a.Impact.ValidatedExternalIsNativeYukawa && !a.Impact.CHiggsLevelC && closeAbs(a.Impact.CurrentCHiggs, cHiggsLevelB, 1e-16), Detail: a.Impact.Verdict},
			{Name: "preserve triality firewall", Passed: a.Triality.Preserved && !a.Triality.ExternalLedgerD4Theorem && !a.Triality.ExternalLedgerGeneration && containsAll(a.Triality.RequiresD4Package, []string{"D4TrialityCarrierPackage", "trace-readout", "breaking operator"}), Detail: a.Triality.Verdict},
			{Name: "record branch decision", Passed: a.Branch.Recorded && !a.Branch.ValidatedExternalFound && !a.Branch.ValidationFailed && !a.Branch.NativeYukawaOperators && strings.Contains(a.Branch.Recommended, "External Yukawa Input Request") && containsAll(a.Branch.Alternatives, []string{"Sector Contribution", "Convention Mismatch", "Native Yukawa"}), Detail: FormatBranch(a.Branch)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && !a.Firewalls.ExternalLedgerNativeYukawa && !a.Firewalls.ValidatedAtomsPMNSCKM && !a.Firewalls.SectorDominanceGeneration && !a.Firewalls.TopDominanceTopYukawa && !a.Firewalls.NEffD4Triality && !a.Firewalls.SingleScaleStable && !a.Firewalls.CHiggsLevelC && !a.Firewalls.TreeProxyPoleMass && a.Firewalls.Verdict == StatusFirewallPreservedGate796, Detail: a.Firewalls.Verdict},
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
		notes := append([]string{a.Truth, FormatExternalSeal(a.ExternalSeal), FormatCircular(a.Circular), FormatAtomRules(a.AtomRules), FormatSchema(a.Schema), FormatValidation(a.Validation), FormatTopChannel(a.TopChannel), FormatBranch(a.Branch), a.FinalStatement}, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
