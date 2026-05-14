package reebweakselection

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ReebVectorSpatialIsotropyWeakPlaneSieveAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-REEB-VECTOR-SPATIAL-ISOTROPY-WEAK-PLANE-SIEVE"
	const name = "Reeb vector spatial isotropy break and contact geometry weak-plane sieve audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 241 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 240 leaves exactly three pure-spatial candidate weak planes", Passed: a.Previous.Summary.U1RejectsTemporalPlanes && a.Previous.Summary.PureSpatialPlanesRemain == 3 && !a.Previous.Summary.UniqueWeakPlaneDerived, Detail: a.Previous.TruthStatement},
			{Name: "finite contact K is retrieved but not an eta/deta/Reeb structure", Passed: a.Contact.ContactProjectorExists && a.Contact.ContactDimension == a.Contact.ExpectedContactDim && !a.Contact.EtaOneFormDerived && !a.Contact.DEtaTwoFormDerived && !a.Contact.ReebVectorDerived, Detail: FormatContact(a.Contact)},
			{Name: "Reeb vector would be the correct selector type but is not derived", Passed: !a.Reeb.CandidateAvailable && !a.Reeb.NativeFromVacuumStabilizer && !a.Reeb.MappedToFockGeneratorW && !a.Reeb.MappedToSpatialFockAxes && !a.Reeb.ManualAxisChoice, Detail: FormatReeb(a.Reeb)},
			{Name: "no native projection maps contact K to the three spatial Fock axes", Passed: !a.Projection.KToWProjectionDerived && a.Projection.ComponentsAreUniformOrAbsent && !a.Projection.UniqueSpatialAxisTagged && !a.Projection.S3PermutationBroken, Detail: FormatProjection(a.Projection)},
			{Name: "the three pure-spatial planes remain unselected", Passed: len(a.Planes) == 3 && a.Sieve.CandidatePlaneCount == 3 && len(a.Sieve.SelectedPlanes) == 0 && !a.Sieve.S3DegeneracyBroken && !a.Sieve.UniqueWeakPlaneSelected, Detail: FormatSieve(a.Sieve) + " :: " + FormatPlanes(a.Planes)},
			{Name: "physical weak chirality/global H/order-one calculus remain blocked", Passed: a.Weak.Gate240ReducedToPureSpatial && a.Weak.ContactGeometryAvailable && !a.Weak.ReebSelectorDerived && !a.Weak.ContactToFockBridgeDerived && !a.Weak.UniqueWeakPlaneSelected && !a.Weak.PhysicalLeftHandedDerived && !a.Weak.GlobalHSummandDerived && !a.Weak.OrderOneReady, Detail: FormatWeak(a.Weak)},
			{Name: "firewall preserved: no Reeb axis or weak plane forced", Passed: !a.Firewall.ForcedReebAxis && !a.Firewall.ImportedContactCoordinates && !a.Firewall.ImportedSMWeakPlane && !a.Firewall.ImportedElectroweakChirality && !a.Firewall.PromotedProjectorToReeb && !a.Firewall.ClaimedGlobalH && !a.Firewall.ClaimedOrderOne && !a.Firewall.FiniteCorePolluted, Detail: FormatFirewall(a.Firewall)},
			{Name: "summary records contact-selector obstruction", Passed: a.Summary.ContactKAvailable && !a.Summary.EtaDEtaDerived && !a.Summary.ReebVectorDerived && !a.Summary.ContactToFockProjection && !a.Summary.SpatialAxisTagged && a.Summary.PureSpatialPlanesInherited == 3 && !a.Summary.UniqueWeakPlaneDerived && !a.Summary.PhysicalChiralityDerived && !a.Summary.GlobalHDerived, Detail: FormatSummary(a.Summary) + " :: " + a.TruthStatement},
		}
		notes := []string{
			"Gate 241 tests the proposed Reeb-vector selector directly. The finite contact projector K is available, but the current core has not constructed a contact one-form eta, its exterior derivative, or a Reeb vector R.",
			"A derived Reeb axis would be the right mathematical kind of object: if it tagged one spatial mode, the complementary pure-spatial two-plane would be selected. This remains hypothetical, not derived.",
			"Therefore the final pure-spatial S3 degeneracy, the physical weak plane, the left-handed weak action, and the global quaternionic H summand remain open.",
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
