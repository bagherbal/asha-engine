package contactkindassignment

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ContactSpectrumToFermionKindAssignmentTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-SPECTRUM-FERMION-KIND-ASSIGNMENT-AUDIT"
	const name = "contact-spectrum-to-fermion-kind assignment obstruction"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact-kind assignment audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "contact scalar-shape target has two high and two low weights", Passed: a.Target.ExactShape == "1197/4624" && a.Target.HighMultiplicity == 2 && a.Target.LowMultiplicity == 2 && a.Target.RequiresKindAssignment && !a.Target.UsesObservedInput, Detail: FormatTarget(a.Target)},
			{Name: "four fermion-kind signatures are available", Passed: len(a.KindSignatures) == 4 && a.Previous.KindQuotient.FourKindSupportQuotientVisible, Detail: FormatKindSignatures(a.KindSignatures)},
			{Name: "finite data gives multiple canonical 2+2 partitions", Passed: a.AssignmentAudit.CanonicalPartitionsFound == 2 && a.AssignmentAudit.MultipleIncompatiblePartitions && !a.AssignmentAudit.UniqueContactKindAssignment, Detail: FormatAssignmentAudit(a.AssignmentAudit) + " :: " + FormatPartitions(a.Partitions)},
			{Name: "no candidate ties the contact high eigenspace to a fermion-kind pair", Passed: a.AssignmentAudit.ContactTiedAssignmentsFound == 0 && a.AssignmentAudit.CanonicalOrientedAssignmentsFound == 0 && a.AssignmentAudit.SurvivingBranchChoices == 6, Detail: FormatAssignmentAudit(a.AssignmentAudit)},
			{Name: "Gate-169 scalar-shape closure remains conditional", Passed: a.Consequence.Gate170FourKindQuotientVisible && a.Consequence.ConditionalShapeStillValid && !a.Consequence.ContactKindAssignmentDerived && !a.Consequence.ScalarShapeClosed && !a.Consequence.AmplitudeTextureSelected, Detail: FormatConsequence(a.Consequence)},
			{Name: "mass and physical-constant firewall remains closed", Passed: a.Firewall.GaugeRatioClosed && a.Firewall.ScalarShapeTargetAvailable && a.Firewall.FourKindSupportQuotientVisible && !a.Firewall.ContactKindAssignmentDerived && !a.Firewall.FourAmplitudeClassQuotientDerived && !a.Firewall.YukawaAmplitudesDerived && !a.Firewall.GenerationTextureDerived && !a.Firewall.FermionMassesDerived && !a.Firewall.CKMPMNSDerived && !a.Firewall.PhysicalConstantsDerived && a.Firewall.ResidualNullityBefore == 3 && a.Firewall.ResidualNullityAfter == 3, Detail: FormatFirewall(a.Firewall) + " :: " + a.TruthStatement},
		}, Notes: []string{
			"Gate 171 separates kind partitions from contact-weight assignment: scalar branch/T3 and color/B-L are both canonical partitions, but they are incompatible selectors.",
			"No currently derived finite object maps the contact high-weight pair to a unique pair of fermion kinds.",
			"The scalar-shape match from Gate 169 remains a finite Yukawa moment target, not a selected mass texture.",
		}}
	}}
}
