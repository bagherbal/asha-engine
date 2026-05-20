package generation2fourfoldselectororigintraceaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2FourfoldSelectorOriginAndTraceTransferAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 fourfold selector origin and trace-transfer audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate555 fourfold selector audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "prove general Fock/Witt selector commutator algebra", Passed: a.Selector.AllRowsVerified && a.Selector.CommutatorIdentity == "[S, E_ij] = (s_i - s_j) E_ij" && a.Selector.CommutantDimension == a.Selector.ExpectedDimension, Detail: FormatSelector(a.Selector)},
			{Name: "verify B-L 4->1+3 selector, commutant, and lepton-color bridge deltas", Passed: a.BMinusL.CommutantDimension == 10 && a.BMinusL.ExpectedDimension == 10 && a.BMinusL.AllBridgeDeltasPMFourThird, Detail: FormatBMinusL(a.BMinusL)},
			{Name: "sieve all six weak-plane candidates under B-L", Passed: a.WeakPlane.RejectedCount == 3 && a.WeakPlane.PreservedCount == 3 && !a.WeakPlane.UniqueWeakPlane, Detail: FormatWeakPlane(a.WeakPlane)},
			{Name: "keep tau_eta sealed without a unit-preserving Fock/generation pullback", Passed: !a.TauEta.ExistingUnitPreservingPullback && !a.TauEta.RhoOneIsIdentity && a.TauEta.CanSelectTwoPlusOneIfPulledBack && !a.TauEta.NativeThreeToTwoPlusOne, Detail: FormatTauEta(a.TauEta)},
			{Name: "verify contact quartic regular representation and block carrier-action promotion", Passed: a.Contact.RegularRepresentationUnit && a.Contact.IrreducibleOverQ && a.Contact.NontrivialRationalIdempotents == 0 && !a.Contact.NativeCarrierAction, Detail: FormatContact(a.Contact)},
			{Name: "compare fourfold carriers without promoting dimension matches", Passed: a.Carriers.RowCount == 5 && a.Carriers.NativeRows == 2 && a.Carriers.BlockedRows == 1, Detail: FormatCarriers(a.Carriers)},
			{Name: "preserve trace-transfer/contact-quartic firewalls", Passed: !a.Firewall.DimensionMatchesPromoted && !a.Firewall.TauEtaPromotedToFockSelector && !a.Firewall.TauEtaPromotedToGenerationMap && !a.Firewall.ContactQuarticPromotedToHiggs && !a.Firewall.ContactQuarticPromotedToFlavor && !a.Firewall.ContactQuarticPromotedToYukawa && !a.Firewall.PhysicalMassesImported && !a.Firewall.ObservedYukawasImported && !a.Firewall.NativeRegistryPolluted, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: append(Statuses(), a.Truth, a.Final.NextTheorem)}
	}}
}
