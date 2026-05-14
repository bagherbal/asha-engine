package doubledspacerepresentation

import "github.com/bagherbal/asha-engine/pkg/theorem"

func DoubledSpaceRepresentationOppositeAlgebraActionAssemblyAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-DOUBLED-SPACE-REPRESENTATION-OPPOSITE-ALGEBRA-ACTION-ASSEMBLY-AUDIT"
	const name = "Doubled-Space Representation / Opposite Algebra Action Assembly Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 294 doubled-space representation audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 293 J_swap KO6 candidate is inherited", Passed: a.Input.KOSixLike && a.Input.J2Sign == 1 && a.Input.JGammaSign == -1, Detail: FormatInput(a.Input)},
			{Name: "doubled-space J_swap satisfies KO6 grading signs", Passed: a.JSwap.KOSixLike && a.JSwap.ResidualJ2 < 1e-12 && a.JSwap.ResidualGamma < 1e-12, Detail: FormatJSwap(a.JSwap)},
			{Name: "naive weak/color Q_L action fails direct-sum multiplicativity", Passed: a.NaiveDiagnostic.MultiplicativityResidual > 0.1, Detail: FormatNaive(a.NaiveDiagnostic)},
			{Name: "representation candidates are separated by category", Passed: len(a.Representations) == 3 && !a.Representations[0].Associative && a.Representations[1].Associative && !a.Representations[1].PhysicalSMBimodule, Detail: FormatRepresentations(a.Representations)},
			{Name: "formal J_swap opposite-action formula is available only conditionally", Passed: a.Opposite.JSwapAvailable && !a.Opposite.ConstructedForPhysicalHF && !a.Opposite.ZeroOrderVerified, Detail: FormatOpposite(a.Opposite)},
			{Name: "full order-one theorem is not verified without physical H_F and D_F", Passed: !a.OrderOne.OrderOneVerified && !a.OrderOne.PhysicalRepresentationAvailable && !a.OrderOne.PhysicalDFAvailable, Detail: FormatOrderOne(a.OrderOne)},
			{Name: "firewalls preserve Higgs/B-gap dynamics", Passed: !a.Firewalls.FiniteCorePolluted && a.Firewalls.DoesNotUnlockHiggs && a.Firewalls.DoesNotUnlockBGap, Detail: FormatFirewalls(a.Firewalls)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, FormatSummary(a.Summary)}}
	}}
}
