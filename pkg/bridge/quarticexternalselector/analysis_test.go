package quarticexternalselector

import "testing"

func TestGate160QuarticExternalSelectorFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.ExactRationalOverlapMatrix || !a.ExactCharacteristicCertified || !a.ExactRootIsolationCertified || a.QuarticOrbitRows != 4 || a.CanonicalQuarticBranches != 0 {
		t.Fatalf("Gate 160 must inherit exact unsplit quartic block: %+v", a.Summary)
	}
	if len(a.Candidates) != 5 || a.SourceAudit.CandidatesAudited != 5 || a.SourceAudit.SourcesAvailable != 5 {
		t.Fatalf("expected five available selector candidates, got %+v", a.SourceAudit)
	}
	if a.Candidates[0].Kind != ScalarVacuumOrientation || !a.Candidates[0].SourceAvailable || a.Candidates[0].SourceCanonical || a.Candidates[0].ReachesQuarticBlock || a.Candidates[0].TwoTwoSplit {
		t.Fatalf("scalar vacuum must not canonically split quartic block: %+v", a.Candidates[0])
	}
	if a.Candidates[1].Kind != BrokenGaugeImages || !a.Candidates[1].SourceAvailable || !a.Candidates[1].SourceCanonical || a.Candidates[1].CanonicalMapToContact || a.Candidates[1].ReachesQuarticBlock || a.Candidates[1].TwoTwoSplit {
		t.Fatalf("broken images must lack contact/quartic intertwiner: %+v", a.Candidates[1])
	}
	if a.Candidates[2].Kind != MatterBLCharge || !a.Candidates[2].SourceAvailable || a.Candidates[2].CanonicalMapToContact || a.Candidates[2].ReachesQuarticBlock || a.FockContactKernel.BMinusLPullbackRowsDerived != 0 || a.FockContactKernel.FullOperatorIntertwinersDerived != 0 {
		t.Fatalf("B-L must not reach quartic block without Fock-contact map: %+v", a.Candidates[2])
	}
	action := a.Candidates[3]
	if action.Kind != ActionSecondVariation || !action.SourceAvailable || !action.SourceCanonical || !action.ReachesQuarticBlock || action.ProjectionRank != 4 || action.DistinctEigenvalues != 1 || action.NonDegenerateSpectrum || action.TwoTwoSplit || action.RequiresBranchChoice {
		t.Fatalf("action restriction must be isotropic and non-splitting: %+v", action)
	}
	cross := a.Candidates[4]
	if cross.Kind != RationalQuarticCoupling || !cross.SourceAvailable || !cross.CanonicalMapToQuartic || a.CrossAudit.CrossTermRank != 0 || a.CrossAudit.CrossTermFrobeniusNorm != 0 || a.CrossAudit.ProvidesSelector || cross.TwoTwoSplit {
		t.Fatalf("rational/quartic cross-coupling must vanish: %+v %+v", cross, a.CrossAudit)
	}
	if a.SourceAudit.SuccessfulSelectors != 0 || a.SourceAudit.NonDegenerateSpectra != 0 || a.SourceAudit.TwoTwoSplits != 0 || a.ExternalSelectorRows != 0 || a.CanonicalTwoTwoSplits != 0 || a.BranchBreakingSources != 0 {
		t.Fatalf("no external selector should succeed: %+v", a.SourceAudit)
	}
	if !a.Firewall.FirewallClosed || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 || a.QuarticBlockBetaRows != 0 || a.QuarticZeroBetaRows != 0 || a.ResidualNullityBefore != 3 || a.ResidualNullityAfter != 3 || a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("Gate 160 must preserve firewall and sealed constants: %+v", a.Firewall)
	}
}
