package yukawashapeconstraint

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FiniteYukawaAmplitudeTextureScalarShapeConstraintTheorem() theorem.Theorem {
	const id = "BRIDGE-FINITE-YUKAWA-AMPLITUDE-TEXTURE-SCALAR-SHAPE-CONSTRAINT"
	const name = "Finite Yukawa amplitude texture target from the Gate-37 scalar-shape constraint"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Yukawa scalar-shape constraint", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate168 scalar-shape target is inherited", Passed: a.Target.InEightSlotRange && a.Target.InFourClassRange && !a.Target.IntegerEqualSlotMatch && !a.Target.UsesObservedMassInput, Detail: FormatTarget(a.Target)},
			{Name: "unit eight-channel incidence remains rejected", Passed: !a.Candidates[0].MatchesTarget && a.Candidates[0].SlotCount == 8, Detail: FormatCandidate(a.Candidates[0])},
			{Name: "direct Phi± duplication of contact spectrum fails", Passed: !a.Candidates[2].MatchesTarget && a.Candidates[2].RequiresPairCollapse && a.Candidates[2].UsesContactEigenvalues, Detail: FormatCandidate(a.Candidates[2])},
			{Name: "four-class contact-spectrum pattern conditionally matches scalar shape", Passed: a.Best.MatchesTarget && a.Best.SlotCount == 4 && a.Best.UsesContactEigenvalues && a.Best.RequiresPairCollapse && a.Best.RequiresKindAssignment && !a.Best.CanonicalSelected, Detail: FormatCandidate(a.Best)},
			{Name: "Higgs-conjugate pair collapse is necessary but not derived", Passed: a.PairCollapse.FourClassQuotientAvailable && !a.PairCollapse.FourClassQuotientDerived && a.PairCollapse.ConditionalShapeMatch && a.PairCollapse.PairCollapseRequiredForMatch && a.PairCollapse.DirectEightChannelDuplicationFails && a.PairCollapse.KindAssignmentAmbiguity == 6, Detail: FormatPairCollapse(a.PairCollapse)},
			{Name: "scalar shape is only one moment constraint on the generation texture problem", Passed: a.Generation.GenerationCount == 3 && a.Generation.KindTextureMatrices == 4 && a.Generation.TotalGeneralTextureEntries == 36 && a.Generation.ScalarShapeConstraints == 1 && a.Generation.ShapeConstraintOnlyMoment && a.Generation.TextureUnderdetermined && !a.Generation.FermionMassesDerived && !a.Generation.CKMPMNSDerived, Detail: FormatGeneration(a.Generation)},
			{Name: "mass/texture firewall remains closed", Passed: a.Firewall.GaugeRatioClosed && a.Firewall.ScalarShapeTargetAvailable && a.Firewall.ConditionalFourClassMatchFound && !a.Firewall.EightChannelAmplitudeTextureSelected && !a.Firewall.PairCollapseDerived && !a.Firewall.KindAssignmentDerived && !a.Firewall.GenerationTextureDerived && !a.Firewall.YukawaAmplitudesDerived && !a.Firewall.FermionMassesDerived && !a.Firewall.CKMPMNSDerived && !a.Firewall.PhysicalConstantsDerived && a.Firewall.ResidualNullityBefore == 3 && a.Firewall.ResidualNullityAfter == 3, Detail: FormatFirewall(a.Firewall) + " :: " + a.TruthStatement},
			{Name: "required two-level amplitude anisotropy is finite and mild", Passed: a.PairCollapse.SquaredAmplitudeRatio > 1 && a.PairCollapse.SquaredAmplitudeRatio < 2 && a.PairCollapse.AmplitudeRatio > 1 && a.PairCollapse.AmplitudeRatio < 1.3, Detail: fmt.Sprintf("|y_high|^2/|y_low|^2=%.12g, |y_high|/|y_low|=%.12g", a.PairCollapse.SquaredAmplitudeRatio, a.PairCollapse.AmplitudeRatio)},
		}, Notes: []string{
			"Gate 169 does not derive observed fermion masses. It turns the Gate-37 scalar shape into a finite Yukawa moment target.",
			"A conditional four-class target matches exactly because the active contact/Higgs spectrum has two high and two low eigenvalues whose shape is λ_contact=1197/4624.",
			"The match is not yet a theorem of the physical Dirac amplitudes: the scalar-conjugate channel quotient, kind assignment, generation lift, phases, and non-commuting texture operators remain open.",
		}}
	}}
}
