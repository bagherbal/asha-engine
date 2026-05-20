package generation2externalyukawaledgerconventionsealandatomdataintakeaudit

import (
	"strings"
	"testing"
)

func TestGate796ExternalSealAndCircularFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate795.Inherited {
		t.Fatalf("bad inheritance: %+v", a.Gate795)
	}
	if !a.ExternalSeal.Defined || !containsAll(a.ExternalSeal.Fields, []string{"source_label", "scale_mu", "renormalization_scheme", "sector_singular_values", "neutrino_convention", "validation_against_aggregate"}) {
		t.Fatalf("bad external seal: %s", FormatExternalSeal(a.ExternalSeal))
	}
	if !a.Circular.Defined || a.Circular.UsesHiggsBacksolve || !containsAll(a.Circular.ForbiddenSources, []string{"N_eff", "C_Higgs", "m_H_pole", "observed Higgs"}) || !containsAll(a.Circular.ForbiddenOperations, []string{"choose T", "choose rest atoms", "Higgs pole mass"}) {
		t.Fatalf("bad circular firewall: %s", FormatCircular(a.Circular))
	}
}

func TestGate796AtomSchemaNeutrinoAndValidation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.AtomRules.Defined || a.AtomRules.AtomFormula != "x_i=y_i^2" || !a.AtomRules.RequiresExactlyOneColorRule || !strings.Contains(a.AtomRules.CoefficientColorConvention, "coefficient 3") || !strings.Contains(a.AtomRules.RepeatedAtomConvention, "repeated") {
		t.Fatalf("bad atom rules: %s", FormatAtomRules(a.AtomRules))
	}
	if !a.Schema.Defined || !a.Schema.RequiresLabels || !containsAll(a.Schema.Fields, []string{"fermion_label", "sector", "generation_label", "y_value", "uncertainty"}) || !containsAll(a.Schema.RequiredSectors, []string{"up", "down", "charged_lepton", "neutrino"}) {
		t.Fatalf("bad schema: %s", FormatSchema(a.Schema))
	}
	if !a.Neutrino.Audited || !a.Neutrino.ExplicitRequired || !containsAll(a.Neutrino.AllowedStatuses, []string{"absent", "zero", "Dirac", "Majorana", "unknown"}) {
		t.Fatalf("bad neutrino audit: %+v", a.Neutrino)
	}
	if !a.Validation.Defined || !a.Validation.MustValidateA || !a.Validation.MustValidateB || !a.Validation.MustValidateNEff || a.Validation.SilentRescaleAllowed || !closeAbs(a.Validation.InheritedNEff, nEffInherited, 5e-16) {
		t.Fatalf("bad validation: %s", FormatValidation(a.Validation))
	}
}

func TestGate796TopSectorScaleImpactBranchAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.TopChannel.Defined || !a.TopChannel.RequiresTypedTop || a.TopChannel.MayInferFromNEff || !containsAll(a.TopChannel.Formulae, []string{"T=y_t^2", "alpha", "beta", "b/a^2"}) {
		t.Fatalf("bad top channel: %s", FormatTopChannel(a.TopChannel))
	}
	if !a.SectorOutputs.Defined || a.SectorOutputs.ExternalLedgerSupplied || a.SectorOutputs.CanOutputNow || !containsAll(a.SectorOutputs.Outputs, []string{"a_u/a", "b_u/b", "largest atoms", "top dominance", "neutrino"}) {
		t.Fatalf("bad sector outputs: %+v", a.SectorOutputs)
	}
	if !a.Scale.Defined || !a.Scale.AllowsSingleScale || !a.Scale.AllowsMultiScale || a.Scale.MultiScaleCertified {
		t.Fatalf("bad scale rules: %+v", a.Scale)
	}
	if !a.Impact.Recorded || !a.Impact.ValidatedExternalWouldImprove || a.Impact.ValidatedExternalIsNativeYukawa || a.Impact.CHiggsLevelC || !closeAbs(a.Impact.CurrentCHiggs, cHiggsLevelB, 1e-16) {
		t.Fatalf("bad impact: %+v", a.Impact)
	}
	if !a.Triality.Preserved || a.Triality.ExternalLedgerD4Theorem || a.Triality.ExternalLedgerGeneration || !containsAll(a.Triality.RequiresD4Package, []string{"D4TrialityCarrierPackage", "trace-readout", "breaking operator"}) {
		t.Fatalf("bad triality firewall: %+v", a.Triality)
	}
	if !a.Branch.Recorded || !strings.Contains(a.Branch.Recommended, "External Yukawa Input Request") || !containsAll(a.Branch.Alternatives, []string{"Sector Contribution", "Convention Mismatch", "Native Yukawa"}) {
		t.Fatalf("bad branch: %s", FormatBranch(a.Branch))
	}
	if !a.Firewalls.Enforced || a.Firewalls.ExternalLedgerNativeYukawa || a.Firewalls.ValidatedAtomsPMNSCKM || a.Firewalls.SectorDominanceGeneration || a.Firewalls.TopDominanceTopYukawa || a.Firewalls.NEffD4Triality || a.Firewalls.SingleScaleStable || a.Firewalls.CHiggsLevelC || a.Firewalls.TreeProxyPoleMass {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
}

func TestGate796TheoremStatusesAndFinalStatement(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.FinalStatement, "does not import") || !strings.Contains(a.FinalStatement, "color-counted exactly once") || !strings.Contains(a.FinalStatement, "aggregate sealed participation") {
		t.Fatalf("bad final statement: %s", a.FinalStatement)
	}
	res := Generation2ExternalYukawaLedgerConventionSealAndAtomDataIntakeAuditTheorem().Verify()
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
