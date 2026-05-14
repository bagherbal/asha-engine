package branchprojector

import "github.com/bagherbal/asha-engine/pkg/theorem"

func BranchwiseQuadraticIdempotentScalarProjectorTheorem() theorem.Theorem {
	const id = "BRIDGE-BRANCHWISE-QUADRATIC-IDEMPOTENT-SCALAR-PROJECTOR-AUDIT"
	const name = "branchwise quadratic idempotent / scalar-projector construction audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build branchwise scalar-projector audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 187 resolvent-vacuum orbit is inherited without selector promotion", Passed: a.PreviousGate187.Firewall.ResolventVacuumAlgebraDerived && a.PreviousGate187.Firewall.DegenerateVacuumOrbitDerived && a.PreviousGate187.Firewall.SpontaneousBranchDataQuarantined && !a.PreviousGate187.Firewall.CanonicalTwoPlusTwoSelectorDerived && !a.PreviousGate187.Firewall.PhysicalScalarBundleDerived, Detail: a.PreviousGate187.Firewall.Verdict},
			{Name: "branchwise extension records the exact missing quadratic adjunction", Passed: a.BaseField.ResolventRootSelectsPartition && a.BaseField.QuadraticAdjunctionLabelsPairFactors && a.BaseField.DoesNotAdjoinIndividualRoots && a.BaseField.InvariantUnderFactorSwap, Detail: FormatBaseField(a.BaseField)},
			{Name: "two monic quadratic factors are constructed without linear root factors", Passed: a.Factors.FactorizationVerified && a.Factors.FactorsMonicQuadratic && a.Factors.FactorsCoprime && a.Factors.NoLinearRootFactorsConstructed && a.Factors.OnlyTwoPlusTwoSplitConstructed && a.Factors.FactorSwapInvolutionPreserved, Detail: FormatFactors(a.Factors)},
			{Name: "Bezout identity is certified exactly over the branchwise field", Passed: a.Bezout.IdentityVerified && a.Bezout.UsesExtendedEuclideanAlgorithm && a.Bezout.ExactArithmetic && a.Bezout.NoNumericRootApproximation, Detail: FormatBezout(a.Bezout)},
			{Name: "complementary 2D scalar projectors are exact idempotents", Passed: a.Projectors.ProjectorPairConstructed && a.Projectors.ProjectorSumIdentity && a.Projectors.ProjectorsIdempotent && a.Projectors.ProjectorsOrthogonal && a.Projectors.TraceTwoEach && a.Projectors.DimensionTwoEach && a.Projectors.IndividualRootProjectors == 0 && !a.Projectors.PhysicalScalarBundleDerived && !a.Projectors.CanonicalBranchSelected, Detail: FormatProjectors(a.Projectors)},
			{Name: "Higgs carrier is opened only as a conditional scalar-projector pair", Passed: a.HiggsCarrier.Gate37PairDegenerate && a.HiggsCarrier.BranchwiseTwoPlusTwoProjectors && a.HiggsCarrier.ConditionalScalarProjectorDerived && !a.HiggsCarrier.CanonicalScalarProjectorDerived && !a.HiggsCarrier.PhysicalScalarBundleDerived && a.HiggsCarrier.ScalarBundleMapRequiresNextBridge && !a.HiggsCarrier.ChernWeilReady && !a.HiggsCarrier.HeatKernelReady && !a.HiggsCarrier.ThresholdRowsReady, Detail: FormatHiggsCarrier(a.HiggsCarrier)},
			{Name: "summary records conditional projector success, not physical scalar-bundle completion", Passed: a.Summary.TestsAudited == 6 && a.Summary.ResolventVacuumInherited && a.Summary.BranchwiseExtensionConstructed && a.Summary.QuadraticFactorsConstructed && a.Summary.BezoutIdentityConstructed && a.Summary.ProjectorPairConstructed && a.Summary.ConditionalScalarProjectorDerived && !a.Summary.PhysicalScalarBundleDerived && a.Summary.IndividualRootProjectors == 0, Detail: FormatSummary(a.Summary)},
			{Name: "firewall forbids root diagonalization and physical-constant promotion", Passed: !a.Firewall.UsesObservedInputForDerivation && !a.Firewall.UsesNumericRootApproximation && !a.Firewall.UsesIndividualRootDiagonalization && !a.Firewall.UsesArbitraryPairingChoice && a.Firewall.ResolventVacuumInherited && a.Firewall.SpontaneousBranchDataQuarantined && a.Firewall.QuadraticAdjunctionRecorded && a.Firewall.BranchwiseQuadraticFactorsDerived && a.Firewall.BranchwiseProjectorPairDerived && a.Firewall.ConditionalScalarProjectorDerived && !a.Firewall.CanonicalUniqueBranchDerived && !a.Firewall.CanonicalScalarProjectorDerived && !a.Firewall.PhysicalScalarBundleDerived && !a.Firewall.ChernWeilCarrierDerived && !a.Firewall.HeatKernelMatchingDerived && !a.Firewall.ThresholdCorrectedBetaDerived && !a.Firewall.AbsoluteCouplingPromoted && !a.Firewall.PhysicalConstantsDerived && a.Firewall.StrictNullityBefore == 3 && a.Firewall.StrictNullityAfter == 3 && a.Firewall.ConditionalNullityBefore == 2 && a.Firewall.ConditionalNullityAfter == 1, Detail: FormatFirewall(a.Firewall)},
		}, Notes: []string{
			a.TruthStatement,
			"Gate 188 corrects the naive overclaim that Q(z) alone writes ordered quadratic factors: Q(z) selects the partition, while eta labels the two factors and is swapped by eta -> -eta.",
			"The physical scalar bundle is not yet derived; the next lawful gate is the scalar-bundle map/H_Phi identification audit.",
		}}
	}}
}
