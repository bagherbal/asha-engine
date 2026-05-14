package axionphenomenologyseal

import "github.com/bagherbal/asha-engine/pkg/theorem"

func AxionPhenomenologySealBGapMisalignmentScaleAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-AXION-PHENOMENOLOGY-SEAL-B-GAP-MISALIGNMENT-AUDIT"
	const name = "AxionPhenomenologySeal / B-Gap misalignment relic density scale audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 226 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 225 dark-matter obstruction is inherited", Passed: a.Gate225.Gate225Inherited && a.Gate225.HeavySectorDMAbsenceBinding && a.Gate225.NativeALPFailed && a.Gate225.NativeContactDMFailed && a.Gate225.BGapValue > 0 && a.Gate225.ContactModeCount == 7, Detail: FormatGate225(a.Gate225)},
			{Name: "AxionPhenomenologySeal is explicit and non-native", Passed: a.Seal.Active && a.Seal.ID == AxionPhenomenologySealID && a.Seal.ConditionalOnBGapALP && a.Seal.GrantsShiftSymmetry && a.Seal.GrantsTopologicalCoupling && a.Seal.GrantsMisalignmentMechanics && !a.Seal.NativeALPDerived && !a.Seal.NativeFAObserved && !a.Seal.NativeAnomalyMapDerived, Detail: FormatSeal(a.Seal)},
			{Name: "misalignment calculation returns a required f_a scale under seal", Passed: a.Misalign.RequiredFAGeV > 0 && a.Misalign.ThetaI == defaultThetaI && !a.Misalign.NativeScaleDerived && !a.Misalign.BGapUsedAsTheta && !a.Misalign.BGapUsedAsMass, Detail: FormatMisalignment(a.Misalign)},
			{Name: "required f_a does not resonate with existing ASHA hierarchy", Passed: !a.Resonance.ResonanceFound && a.Resonance.ClosestLog10Gap > 1.0 && len(a.Resonance.Comparisons) == 3, Detail: FormatResonance(a.Resonance)},
			{Name: "B-gap-as-theta variant is diagnostic only and rejected as noncanonical", Passed: a.BGapTheta.Evaluated && !a.BGapTheta.Promoted && a.BGapTheta.RequiredFAGeV > 0 && a.BGapTheta.ClosestLog10Gap > 1.0, Detail: FormatBGapThetaDiagnostic(a.BGapTheta)},
			{Name: "dark matter is parameterized only under seal, not derived from the finite core", Passed: a.DM.ALPOmegaComputedUnderSeal && a.DM.ALPAccountsForDMUnderSeal && a.DM.NativeDMStillAbsent && a.DM.ObservedDMUsedAsTargetOnly && !a.DM.DarkMatterDerivedFromFinite && len(a.DM.DeferredNativeTasks) > 0, Detail: FormatDM(a.DM)},
			{Name: "firewalls remain closed", Passed: a.Firewall.Gate225Inherited && a.Firewall.HeavyDMAbsencePreserved && a.Firewall.AxionSealActive && !a.Firewall.ShiftSymmetryFiniteDerived && !a.Firewall.TopologicalCouplingDerived && !a.Firewall.AxionDecayConstantFinite && !a.Firewall.BGapPromotedWithoutSeal && !a.Firewall.ObservedDMUsedToRewriteCore && !a.Firewall.StructuralResonanceOverclaimed && !a.Firewall.ContactModesPromoted && !a.Firewall.FiniteCorePolluted, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Checks: checks, Notes: []string{a.TruthStatement, "CONDITIONAL_PHENOMENOLOGY: Gate 226 can compute a sealed ALP misalignment scale, but it finds no resonance with v, M_B, or M_* and does not derive the axion from finite geometry."}}
	}}
}
