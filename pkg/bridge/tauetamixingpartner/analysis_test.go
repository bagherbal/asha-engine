package tauetamixingpartner

import "testing"

func TestGate262TauEtaMixingPartnerAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Inheritance.BilinearCarrierDefined || !a.Inheritance.TauEtaActionDerived || a.Inheritance.OffDiagonalComplementDimension != 6 || a.Inheritance.PreviousCanonicalPartnerFound {
		t.Fatalf("bad Gate 261 inheritance: %s", FormatInheritance(a.Inheritance))
	}
	if a.Inventory.CandidateCount != 6 || a.Inventory.MatrixCandidates != 4 || a.Inventory.RawNonCommutingCandidates != 4 || a.Inventory.QualifiedFiniteMixingPartners != 0 {
		t.Fatalf("bad candidate inventory: %s", FormatInventory(a.Inventory))
	}
	if !a.TrialityPartner.PermutationCycleNonCommuting || !a.TrialityPartner.ReflectionNonCommuting || !a.TrialityPartner.HermitianRealPartNonCommuting || !a.TrialityPartner.HermitianImaginaryPartNonCommuting {
		t.Fatalf("triality complement not populated: %s", FormatTriality(a.TrialityPartner))
	}
	if a.TrialityPartner.RawComplementDirectionsTouched != 6 || a.TrialityPartner.HermitianPhaseBasisDimension != 2 || a.TrialityPartner.AnyTrialityMapQualifiedAsAmplitude {
		t.Fatalf("bad triality partner audit: %s", FormatTriality(a.TrialityPartner))
	}
	if !a.FinitePhaseGap.BGapAvailableAsPositiveScale || a.FinitePhaseGap.BGapHasGenerationMatrix || a.FinitePhaseGap.BGapCanPopulateOffDiagonalComplement {
		t.Fatalf("B_gap should remain scalar-only: %s", FormatFinitePhaseGap(a.FinitePhaseGap))
	}
	if !a.FinitePhaseGap.HopfPhasesAvailableAsPhaseLedger || a.FinitePhaseGap.HopfPhaseGenerationMapDerived || a.FinitePhaseGap.HopfCanPopulateOffDiagonalComplement {
		t.Fatalf("Hopf phase should remain representation-free: %s", FormatFinitePhaseGap(a.FinitePhaseGap))
	}
	if !a.PartnerVerdict.RawNonCommutingPartnerExists || !a.PartnerVerdict.RawSelfAdjointOffDiagonalBasisExists || a.PartnerVerdict.QualifiedFiniteMixingPartnerFound || a.PartnerVerdict.PhysicalYukawaTextureDerived {
		t.Fatalf("bad partner verdict: %s", FormatPartnerVerdict(a.PartnerVerdict))
	}
	if !a.Firewall.DoesNotPromoteSymmetryToAmplitude || !a.Firewall.DoesNotUseBGapAsTextureWithoutMap || !a.Firewall.DoesNotUseHopfPhaseWithoutMap || !a.Firewall.EmpiricalYukawaSealInactive || a.Firewall.FiniteCorePolluted {
		t.Fatalf("firewall violation: %s", FormatFirewall(a.Firewall))
	}
}

func TestGate262ExactCommutators(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	byName := map[string]CandidateAudit{}
	for _, c := range a.Candidates {
		byName[c.Name] = c
	}
	if byName["C3_cycle"].CommutatorFrobeniusNormSquared != 26 {
		t.Fatalf("unexpected C3 commutator: %s", FormatCandidate(byName["C3_cycle"]))
	}
	if byName["S3_reflection_23"].CommutatorFrobeniusNormSquared != 18 {
		t.Fatalf("unexpected reflection commutator: %s", FormatCandidate(byName["S3_reflection_23"]))
	}
	if byName["A_triality_real=C+C^T"].CommutatorFrobeniusNormSquared != 52 {
		t.Fatalf("unexpected real Hermitian commutator: %s", FormatCandidate(byName["A_triality_real=C+C^T"]))
	}
	if byName["K_triality_phase=i(C-C^T)"].CommutatorFrobeniusNormSquared != 52 {
		t.Fatalf("unexpected phase Hermitian commutator: %s", FormatCandidate(byName["K_triality_phase=i(C-C^T)"]))
	}
	if !IsSelfAdjoint(byName["A_triality_real=C+C^T"].Matrix) || !IsSelfAdjoint(byName["K_triality_phase=i(C-C^T)"].Matrix) {
		t.Fatalf("Hermitian triality bases must be self-adjoint")
	}
}
