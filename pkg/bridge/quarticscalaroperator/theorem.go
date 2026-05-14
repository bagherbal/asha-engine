package quarticscalaroperator

import "github.com/bagherbal/asha-engine/pkg/theorem"

func QuarticScalarOperatorMinimalPolynomialTheorem() theorem.Theorem {
	const id = "BRIDGE-QUARTIC-SCALAR-OPERATOR-MINIMAL-POLYNOMIAL"
	const name = "quartic scalar operator / minimal-polynomial construction on H_Phi"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build quartic scalar operator audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 184 quartic scalar escape hatch is inherited", Passed: a.PreviousGate184.Firewall.QuarticScalarAbstractModuleDerived && !a.PreviousGate184.Firewall.CanonicalScalarActionDerived, Detail: a.PreviousGate184.Firewall.Verdict},
			{Name: "exact quartic primary polynomial is recorded", Passed: a.Polynomial.Degree == 4 && len(a.Polynomial.Coefficients) == 5 && a.Polynomial.Coefficients[0] == "3240" && a.Polynomial.Coefficients[4] == "271", Detail: FormatPolynomial(a.Polynomial)},
			{Name: "companion operator satisfies q4(T)=0 exactly", Passed: a.Companion.Dimension == 4 && a.Companion.PolynomialIdentityZero && a.Companion.CyclicModule && a.Companion.CyclicVectorRank == 4 && a.Companion.BranchFree && a.Companion.UsesExactRationalArithmetic && a.Companion.AbstractQuarticModule, Detail: FormatCompanion(a.Companion)},
			{Name: "quartic trace moment ledger matches Gate 161", Passed: a.Moments.AllMatch && a.Moments.TraceT == "71/30" && a.Moments.TraceT2 == "1471/900" && a.Moments.TraceT3 == "33581/27000" && a.Moments.TraceT4 == "809891/810000", Detail: FormatMoments(a.Moments)},
			{Name: "Gate 37 scalar operator is not the quartic-minimal operator", Passed: a.Gate37Comparison.ActiveDimension == 4 && a.Gate37Comparison.Gate37PairDegenerate && a.Gate37Comparison.MomentShapeMatchesContactTarget && !a.Gate37Comparison.HasQuarticMinimalPolynomial && !a.Gate37Comparison.IdentifiedWithQuarticModule && !a.Gate37Comparison.PhysicalScalarBundleDerived, Detail: FormatGate37(a.Gate37Comparison)},
			{Name: "exact block restriction to physical H_Phi is not yet available", Passed: a.BlockRestriction.ExactOmegaContactAvailable && a.BlockRestriction.QuarticPrimaryProjectorAvailable && !a.BlockRestriction.PhysicalHphiProjectorAvailable && !a.BlockRestriction.CanonicalMapQuarticToHphi && !a.BlockRestriction.ExactBlockRestrictionComputed && !a.BlockRestriction.PromotesCompanionToPhysicalHphi, Detail: FormatBlockRestriction(a.BlockRestriction)},
			{Name: "summary distinguishes abstract module from physical scalar bundle", Passed: a.Summary.CandidatesAudited == 3 && a.Summary.ExactQuarticOperators == 1 && a.Summary.PolynomialIdentitiesVerified == 1 && a.Summary.MomentLedgersMatched == 1 && a.Summary.AbstractModulesDerived == 1 && a.Summary.PhysicalHphiOperatorsWithQuarticMinPoly == 0 && a.Summary.PhysicalScalarBundlesDerived == 0, Detail: FormatSummary(a.Summary)},
			{Name: "firewall preserves nullity and forbids overclaiming physical constants", Passed: !a.Firewall.UsesObservedInputForDerivation && !a.Firewall.UsesBranchDiagonalization && !a.Firewall.UsesArbitraryMatrixFit && a.Firewall.QuarticAbstractOperatorDerived && a.Firewall.QuarticMomentsVerified && a.Firewall.Gate37ScalarOperatorQuadratic && !a.Firewall.ExactBlockRestrictionDerived && !a.Firewall.CanonicalHphiIdentificationDerived && !a.Firewall.PhysicalScalarBundleDerived && !a.Firewall.ChernWeilCarrierDerived && !a.Firewall.HeatKernelMatchingDerived && !a.Firewall.ThresholdCorrectedBetaDerived && !a.Firewall.AbsoluteCouplingPromoted && !a.Firewall.PhysicalConstantsDerived && a.Firewall.StrictNullityBefore == 3 && a.Firewall.StrictNullityAfter == 3 && a.Firewall.ConditionalNullityBefore == 2 && a.Firewall.ConditionalNullityAfter == 2, Detail: FormatFirewall(a.Firewall)},
		}, Notes: []string{
			a.TruthStatement,
			"Gate 185 proves the exact abstract quartic module Q[x]/(q4) without diagonalizing the quartic roots.",
			"The physical scalar bundle is not promoted: Gate 37 gives a pair-degenerate scalar mixing operator, while the quartic companion operator is a separate exact module until a canonical scalar/contact identification is derived.",
		}}
	}}
}
