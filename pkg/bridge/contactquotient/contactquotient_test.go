package contactquotient

import "testing"

func TestContactSpectralInvariantQuotientOrbitCollapse(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.ActionObstructionInherited || a.ContactAction.ContactWeightedAutomorphismGroupOrder != 1 || a.ContactAction.FanoAutomorphismGroupOrder != 168 {
		t.Fatalf("expected Gate 119 action obstruction")
	}
	if !a.WeightedSpectrumQuotientDerived || !a.WeightedQuotientCanonical || !a.WeightedQuotientIsIdentity || !a.WeightedQuotientPreservesAllRows || a.WeightedQuotientProducesFanoOrbit {
		t.Fatalf("weighted quotient should be canonical singleton-only: %s", FormatSummary(a.Summary))
	}
	if !a.AnonymousSpectrumQuotientDerived || !a.AnonymousQuotientCanonical || !a.AnonymousQuotientCollapsesAllRows || !a.AnonymousQuotientDestroysSpectralRows || a.AnonymousQuotientRepresentationUsable {
		t.Fatalf("anonymous quotient should collapse and lose row data: %s", FormatSummary(a.Summary))
	}
	if !a.TransportedFanoQuotientPossibleAfterChoice || a.TransportedFanoQuotientCanonical || a.Summary.CompatibleBijectionCount != 5040 || a.Summary.CanonicalTransportedQuotients != 0 {
		t.Fatalf("transported quotient should remain convention-dependent: %s", FormatSummary(a.Summary))
	}
	if !a.QuotientForkObstructionDerived || !a.OrbitCollapseObstructionDerived || !a.SpectralInformationLossDerived || a.RepresentationRowFromQuotientDerived {
		t.Fatalf("expected quotient fork/orbit-collapse obstruction")
	}
	if a.ContactRows != 7 || a.OpenContactRowsAfter != 7 || a.RepresentationCompleteRows != 0 || a.RepresentationOpenRows != 7 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 || a.ThresholdCorrectedBetaDerived {
		t.Fatalf("contact beta firewall should remain closed")
	}
	if a.ResidualNullityAfter != 3 || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived || a.HiddenObservedInputUsed {
		t.Fatalf("physical bridge leaked")
	}
}
