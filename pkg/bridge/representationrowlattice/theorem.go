package representationrowlattice

import "github.com/bagherbal/asha-engine/pkg/theorem"

func RepresentationRowLatticeCompletionAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-REPRESENTATION-ROW-LATTICE-COMPLETION-AUDIT"
	const name = "representation-row lattice completion / finite heavy-sector basis search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build representation-row lattice audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Checks: []theorem.Check{
			{Name: "Gate 203 failed-route source classification is inherited", Passed: a.PreviousGate203.Gate203Inherited && a.PreviousGate203.Gate203FailedRoutePreserved && a.PreviousGate203.UniversalBetaSourceStillExternal && !a.PreviousGate203.CompleteMultipletSourceFound && !a.PreviousGate203.RegulatorTraceSourceFound && !a.PreviousGate203.PhysicalUnificationClaimed, Detail: FormatGate203(a.PreviousGate203) + " :: " + FormatShapeRequirements(a.PreviousGate203.Gate201NonUniversalShapeRequirements)},
			{Name: "finite rational beta-row grammar is constructed without unbounded representation enumeration", Passed: a.GrammarAudit.FiniteAlphabetDeclared && a.GrammarAudit.UnboundedEnumerationAvoided && a.GrammarAudit.CandidateRowsGenerated == len(a.Rows) && a.GrammarAudit.UniqueRows == len(a.UniqueRows) && a.GrammarAudit.ExactRationalRows == len(a.Rows) && a.GrammarAudit.StandardFormulaRows == len(a.Rows) && a.GrammarAudit.CommonDenominatorLCM > 0, Detail: FormatGrammar(a.GrammarAudit) + " :: " + FormatRows(a.UniqueRows, 8)},
			{Name: "discrete row lattice is an exact rational semigroup with no continuous RG-scale fit", Passed: a.LatticeAudit.GeneratorCount == len(a.UniqueRows) && a.LatticeAudit.IntegerGridEmbedded && a.LatticeAudit.ContainsZeroRow && a.LatticeAudit.SemigroupOnly && a.LatticeAudit.NoContinuousScales && a.LatticeAudit.NoUniversalFit, Detail: FormatLattice(a.LatticeAudit)},
			{Name: "Gate-201 non-universal shapes are exact row-lattice generators", Passed: a.MembershipAudit.ShapesAudited == len(a.PreviousGate203.Gate201NonUniversalShapeRequirements) && a.MembershipAudit.AllGate201ShapesSupported && a.MembershipAudit.DirectGeneratorMatches == len(a.PreviousGate203.Gate201NonUniversalShapeRequirements) && a.MembershipAudit.ConditionalSupportCount == len(a.PreviousGate203.Gate201NonUniversalShapeRequirements) && a.MembershipAudit.UniversalCompletionIgnored, Detail: FormatMembershipAudit(a.MembershipAudit) + " :: " + FormatMemberships(a.Memberships)},
			{Name: "contact partial-overlap modes are not promoted to a finite heavy-sector basis", Passed: a.ContactInventory.ContactPartialOverlapModes == 7 && !a.ContactInventory.ContactModesHaveChargeLabels && !a.ContactInventory.ContactModesHaveGaugeRepSemantics && !a.ContactInventory.ContactModesHaveDynkinIndices && !a.ContactInventory.ContactModesHaveSpinStatistics && !a.ContactInventory.ContactModesHaveMassActivation && !a.ContactInventory.ContactModesHaveDecouplingLaw && !a.ContactInventory.CanonicalMapToRowBasisFound && a.ContactInventory.CandidateRowsAssigned == 0 && !a.ContactInventory.FiniteHeavySectorBasisDerived, Detail: FormatContactInventory(a.ContactInventory)},
			{Name: "universal beta source and continuous threshold lever arm remain external", Passed: a.Firewall.UniversalBetaSourceStillExternal && !a.Firewall.UniversalBetaFitAttempted && !a.Firewall.ContinuousScalesSolved && !a.Firewall.Gate201ShapesPromotedToFinitePrediction && !a.Firewall.ObservedInputsUsedForFiniteDerivation, Detail: FormatFirewall(a.Firewall)},
			{Name: "physical firewalls and nullity are preserved", Passed: !a.Firewall.ContactModesPromotedToBetaRows && !a.Firewall.FockGenerationPromotedToNewThreshold && !a.Firewall.PhysicalUnificationClaimed && !a.Firewall.ThresholdCorrectedPhysicalFitClaimed && !a.Firewall.AbsoluteMassPredicted && !a.Firewall.FiniteMatchingCorrectionsDerived && a.Firewall.StrictNullityBefore == a.Firewall.StrictNullityAfter && a.Firewall.PhysicalPredictionNullityBefore == a.Firewall.PhysicalPredictionNullityAfter && a.Summary.ConditionalSupportLogged && a.Summary.NoPhysicalPredictionClaim, Detail: FormatFirewall(a.Firewall) + " :: " + FormatSummary(a.Summary)},
		}, Notes: []string{
			a.TruthStatement,
			"Gate 204 answer: the Gate-201 non-universal shapes are legal exact rational row-lattice generators, so the shape grammar receives CONDITIONAL_SUPPORT.",
			"This does not repair the universal beta row, derive contact heavy-sector carriers, or claim absolute unification.",
		}}
	}}
}
