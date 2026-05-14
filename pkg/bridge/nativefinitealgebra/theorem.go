package nativefinitealgebra

import "github.com/bagherbal/asha-engine/pkg/theorem"

func NativeFiniteAlgebraContactPreservingSubalgebraSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-NATIVE-FINITE-ALGEBRA-CONTACT-PRESERVING-SUBALGEBRA-SEARCH"
	const name = "Native finite algebra derivation / contact-preserving subalgebra search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 236 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 235 complexified carrier is inherited", Passed: a.Previous.Summary.ComplexificationDerived && a.Previous.Summary.AntiLinearJAvailable && !a.Previous.Summary.NativeAlgebraDerived, Detail: a.Previous.TruthStatement},
			{Name: "native mode bookkeeping contains a 1⊕3 split", Passed: a.Split.ModeLevelProjectionExists && a.Split.LeptonLikeGeneratorCount == 1 && a.Split.ColorLikeGeneratorCount == 3 && a.Split.ColorLeptonBookkeepingNative && a.Split.ExtendsToFockBookkeeping, Detail: FormatSplit(a.Split) + " :: " + FormatModes(a.Modes)},
			{Name: "1⊕3 mode commutant supports C⊕M3(C) preflight", Passed: a.Commutant.ColorMatrixAlgebraPreflight && a.Commutant.ComplexSingletPreflight && a.Commutant.ModeProjectionCommutantDimensionC == 10 && !a.Commutant.M3CDerivedAsPhysicalColorGauge && !a.Commutant.LiftToFullExteriorRepresentation && !a.Commutant.MaximalAlgebraOnFullSC, Detail: FormatCommutant(a.Commutant)},
			{Name: "u(1) complex summand is plausible but su(2) is not promoted to H", Passed: a.Contact.U1ComplexSummandPreflight && a.Contact.SU2LieAlgebraAvailable && !a.Contact.SU2ToQuaternionHModuleDerived && !a.Contact.LeftQuaternionicActionDerived && !a.Contact.HGenerated, Detail: FormatContact(a.Contact)},
			{Name: "exact Connes algebra remains un-derived and un-imported", Passed: !a.Algebra.ConnesAlgebraImported && a.Algebra.CPlusM3Preflight && !a.Algebra.QuaternionicHDerived && !a.Algebra.ExactCPlusHPlusM3Derived && !a.Algebra.FaithfulRepresentationOnSC && !a.Algebra.OppositeAlgebraActionDerived && !a.Algebra.OrderOneCalculusReady, Detail: FormatAlgebra(a.Algebra)},
			{Name: "firewall blocks forced Standard Model algebra", Passed: !a.Firewall.ImportedConnesAlgebra && !a.Firewall.InsertedSMGaugeGroup && !a.Firewall.InsertedGaugeMatrices && !a.Firewall.InsertedYukawaOrMassData && !a.Firewall.BGapPromotedToMass && !a.Firewall.ClaimedOrderOne && !a.Firewall.ClaimedSMAlgebraDerivation && !a.Firewall.FiniteCorePolluted, Detail: FormatFirewall(a.Firewall)},
			{Name: "summary records partial native algebra support and binding obstruction", Passed: a.Summary.SplitDerived && a.Summary.CPlusM3Preflight && a.Summary.U1ComplexPreflight && !a.Summary.QuaternionicHDerived && !a.Summary.ExactSMAlgebraDerived && !a.Summary.OrderOneReady, Detail: FormatSummary(a.Summary) + " :: " + a.TruthStatement},
		}
		notes := []string{
			"Gate 236 derives a native 1⊕3 mode-level algebraic preflight: C⊕M3(C) appears as the block commutant of the temporal/spatial split.",
			"This is not yet the Connes algebra: the quaternionic H summand, faithful doubled-space representation, opposite action, and order-one calculus remain un-derived.",
			"The next hard target is to represent the contact-preserving su(2) derivations explicitly on S_C and test whether their associative closure forces a left quaternionic module.",
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
