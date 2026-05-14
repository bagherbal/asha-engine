package flavororientationoperator

import "github.com/bagherbal/asha-engine/pkg/theorem"

func FlavorOrientationOperatorTrialityToMassEigenstateTextureAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-FLAVOR-ORIENTATION-OPERATOR-TRIALITY-TO-MASS-EIGENSTATE-TEXTURE-AUDIT"
	const name = "Flavor Orientation Operator / Triality-to-Mass-Eigenstate Texture Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 324 flavor orientation audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "geometric trace basis and physical mass basis are separated", Passed: a.Basis.Formalized && len(a.Basis.GeometricBasis) == 3 && len(a.Basis.MassBasis) == 3 && len(a.Basis.NormalizedSource) == 3, Detail: FormatBasis(a.Basis)},
			{Name: "known doubled/J/bimodule operators do not yet select a generation-space unitary", Passed: a.Operator.SieveFormalized && !a.Operator.JSwapActsOnFlavor && !a.Operator.DoubledSpaceActsOnFlavor && !a.Operator.BimoduleOverlapActsOnFlavor && !a.Operator.InstalledNativeUnitary, Detail: FormatOperator(a.Operator)},
			{Name: "tau_eta nullspace has exact top-boundary suppression capacity", Passed: a.Nullspace.Computed && a.Nullspace.Dimension == 2 && a.Nullspace.AllBasisVectorsOrthogonal && a.Nullspace.TopSuppressionPossible && !a.Nullspace.UniquePhysicalTopVector, Detail: FormatNullspace(a.Nullspace)},
			{Name: "candidate flavor vectors reproduce direct-slot and null-top fractions", Passed: len(a.Candidates) == 5 && a.Candidates[0].TopFraction > 0.44-1e-9 && a.Candidates[2].TopFraction > 0.11-1e-9 && a.Candidates[3].TopFraction < 1e-12 && a.Candidates[4].TopFraction < 1e-12, Detail: FormatCandidate(a.Candidates[3]) + " | " + FormatCandidate(a.Candidates[4])},
			{Name: "null-top rotation recovers Gate-322 envelope only as a non-native candidate", Passed: a.RG.Audited && a.RG.NullRotationPreservesGate322 && !a.RG.NativeJustificationInstalled && !a.RG.PhysicalLaneAuthorized, Detail: FormatRG(a.RG)},
			{Name: "firewalls prevent CKM import, observed top-mass input, and collider-mass claim", Passed: a.Firewalls.NoCKMImported && a.Firewalls.NoObservedTopMassInserted && a.Firewalls.NoFlavorTextureInvented && a.Firewalls.NoPoleMassClaimed && a.Firewalls.NoTwoLoopClaimed && a.Firewalls.NoColliderMassClaimed && !a.Firewalls.FiniteCorePolluted, Detail: FormatFirewalls(a.Firewalls)},
			{Name: "summary records capacity without authorizing physical top lane", Passed: a.Summary.BasisFormalized && a.Summary.FlavorSieveFormalized && a.Summary.NullspaceCapacityProved && !a.Summary.NativeFlavorOperatorDerived && !a.Summary.TopBoundarySuppressionJustified && !a.Summary.Gate322PhysicalLaneAuthorized && a.Summary.FirewallsPreserved && !a.Summary.FinalMassClaimed, Detail: FormatSummary(a.Summary)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, "Gate 324 proves that a null-top flavor orientation exists mathematically, but the native U_flavor/CKM texture is not derived.", "The Gate-322 near-125 GeV transport therefore remains a diagnostic flattened-top envelope until the flavor-orientation operator is selected by the finite geometry."}}
	}}
}
