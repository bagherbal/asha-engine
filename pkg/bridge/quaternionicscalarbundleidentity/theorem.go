package quaternionicscalarbundleidentity

import (
	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func QuaternionicScalarBundleIdentitySieveTheorem() theorem.Theorem {
	const id = "GATE399-QUATERNIONIC-SCALAR-BUNDLE-IDENTITY-SIEVE"
	const name = "Quaternionic (H) endomorphism / scalar bundle identity sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate399 audit", Passed: false, Detail: err.Error()}}}
		}
		j := findEndomorphism(a.Endomorphisms.Candidates, "left H unit J pair-rotation on H_phi")
		generic := findEndomorphism(a.Endomorphisms.Candidates, "generic single quaternion element")
		sealed := findEndomorphism(a.Endomorphisms.Candidates, "sealed q4 companion operator placed on H_phi")
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{
			{Name: "inherits Gate 398 quartic/H_phi obstruction and scalar/firewall ledgers", Passed: a.Inheritance.Executed && a.Inheritance.Gate398NoCanonicalHphiID && a.Inheritance.Gate398QuarticDim == 4 && a.Inheritance.Gate398HphiDim == 4 && a.Inheritance.Gate385OneFormEdgeSupportDerived && a.Inheritance.Gate372ChargedModuliDim == 13 && a.Inheritance.NoEmpiricalValuesImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "q4 remains an irreducible contact quartic primary", Passed: a.Q4.Degree == 4 && a.Q4.IrreducibleOverQ && a.Q4.BranchFreePrimary && a.Q4.ContactSpectralDatum, Detail: FormatQ4(a.Q4)},
			{Name: "H_phi supports local quaternionic/doublet structure", Passed: a.Module.RealDimension == 4 && a.Module.ComplexDoubletDimension == 2 && a.Module.LocalHExtracted && a.Module.MoritaWeakHAction && a.Module.PairComplexAvailable && a.Module.AbstractQuaternionicTripleAvailable, Detail: FormatModule(a.Module)},
			{Name: "global quaternionic scalar theorem remains unsealed", Passed: !a.Module.GlobalHUnsealed && !a.Module.CanonicalComplexDerived && !a.Module.QuaternionicTripleSelectedByScalar, Detail: FormatModule(a.Module)},
			{Name: "strongest pair-compatible H action is quadratic", Passed: j.Native && j.QuaternionicAction && j.SquaresToMinusIdentity && j.MinimalDegree == 2 && j.CharPolyIsSquareOfQuadratic && !j.Q4ExactMatch && !j.PromotableAsQ4Selector, Detail: FormatEndomorphism(j)},
			{Name: "generic single quaternion element cannot have q4 minimal polynomial", Passed: generic.Native && generic.QuaternionicAction && generic.MinimalDegree == 2 && generic.CharPolyIsSquareOfQuadratic && !generic.Q4ExactMatch && !generic.PromotableAsQ4Selector, Detail: FormatEndomorphism(generic)},
			{Name: "sealed q4 companion stress test is not quaternionic or promotable", Passed: sealed.Sealed && sealed.Circular && sealed.Q4ExactMatch && !sealed.QuaternionicAction && !sealed.CompatibleWithJ && !sealed.CompatibleWithFirstOrder && !sealed.PromotableAsQ4Selector, Detail: FormatEndomorphism(sealed)},
			{Name: "no native q4 scalar selector exists", Passed: a.Endomorphisms.PromotableNativeCount == 0 && a.Endomorphisms.Q4ExactMatchCount == 1, Detail: FormatEndomorphisms(a.Endomorphisms)},
			{Name: "identity and flavor firewall are preserved", Passed: !a.Identity.HphiQuarticIdentified && !a.Identity.OneFormEdgeFunctorDerived && !a.Identity.YukawaCouplingsReduced && a.Identity.ChargedModuliResult == 13 && a.Identity.FlavorFirewallPreserved, Detail: FormatIdentity(a.Identity)},
			{Name: "firewalls remain clean", Passed: a.Firewall.NoObservedMassesImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.NoObservedHiggsInserted && a.Firewall.NoManualQ4HphiID && a.Firewall.NoCompanionOperatorPromoted && a.Firewall.NoYukawaCouplingClaimed && a.Firewall.NoFlavorModuliReductionClaimed, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate searches non-quaternionic identity selectors", Passed: a.Next.Gate == 400 && a.Next.Title != "", Detail: FormatNext(a.Next)},
		}, Notes: []string{a.Truth}}
	}}
}

func findEndomorphism(xs []EndomorphismFingerprint, name string) EndomorphismFingerprint {
	for _, x := range xs {
		if x.Name == name {
			return x
		}
	}
	return EndomorphismFingerprint{Name: name, Reason: "not found", Verdict: "MISSING"}
}
