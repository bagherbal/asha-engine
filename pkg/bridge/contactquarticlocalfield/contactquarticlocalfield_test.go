package contactquarticlocalfield

import "testing"

func TestGate156QuarticLocalFieldObstructionKeepsFirewallClosed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.QuarticOrbitRows != 4 || a.QuarticBlockInvariants != 4 || !a.ExactRationalOverlapMatrix || !a.ExactCharacteristicCertified || !a.ExactRootIsolationCertified {
		t.Fatalf("Gate 156 must inherit exact quartic spectral block: %+v", a.Summary)
	}
	if len(a.FieldCandidates) != 5 || a.Summary.DegreeMatchingCandidates != 5 {
		t.Fatalf("Gate 156 must audit five degree-matching local-field candidates: %+v", a.Summary)
	}
	if a.RequirementAudit.BaseSpaceRows != 0 || a.RequirementAudit.LocalSectionRows != 0 || a.RequirementAudit.LorentzRepresentationRows != 0 || a.RequirementAudit.KineticOperatorRows != 0 || a.RequirementAudit.PoleResidueRows != 0 || a.RequirementAudit.SpinStatisticsRows != 0 || a.RequirementAudit.GaugeRepresentationRows != 0 || a.RequirementAudit.HyperchargeRows != 0 || a.RequirementAudit.MassActivationRows != 0 || a.RequirementAudit.DecouplingRows != 0 || a.RequirementAudit.AllRequirementsSatisfied {
		t.Fatalf("Gate 156 must not derive local-field requirements: %+v", a.RequirementAudit)
	}
	if a.SpinAudit.SpinStatisticsComplete != 0 || a.SpinStatisticsRows != 0 {
		t.Fatalf("Gate 156 must keep spin-statistics unresolved: %+v", a.SpinAudit)
	}
	if !a.KineticAudit.FiniteSpectralBlockExact || a.KineticAudit.LocalDifferentialOperatorRows != 0 || a.KineticAudit.PropagatorDenominatorRows != 0 || a.KineticAudit.PositiveResidueRows != 0 || a.KineticAudit.LocalityComplete {
		t.Fatalf("Gate 156 must distinguish finite spectrum from local propagator: %+v", a.KineticAudit)
	}
	if !a.Firewall.FirewallClosed || a.LocalFieldRows != 0 || a.KineticPoleResidueRows != 0 || a.GaugeRepresentationRows != 0 || a.HyperchargeRowsDerived != 0 || a.MassActivationRows != 0 || a.DecouplingRows != 0 || a.QuarticBlockBetaRows != 0 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 {
		t.Fatalf("Gate 156 must keep beta firewall closed: %+v", a.Summary)
	}
	if a.ResidualNullityBefore != 3 || a.ResidualNullityAfter != 3 || a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("Gate 156 must not change physical-flow nullity or import observations: %+v", a.Summary)
	}
}
