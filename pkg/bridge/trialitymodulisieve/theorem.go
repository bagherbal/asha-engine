package trialitymodulisieve

import "github.com/bagherbal/asha-engine/pkg/theorem"

func TrialityDomainAdmissionEquivariantYukawaCentralizerSieveTheorem() theorem.Theorem {
	const id = "BRIDGE-TRIALITY-DOMAIN-ADMISSION-EQUIVARIANT-YUKAWA-CENTRALIZER-SIEVE"
	const name = "Triality Domain-Admission & Equivariant Yukawa Centralizer Sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 393 audit", Passed: false, Detail: err.Error()}}}
		}
		c3 := centralizerByName(a.Centralizer.Cases, "exact C3 label-triality stress test")
		s3 := centralizerByName(a.Centralizer.Cases, "exact S3 label-triality stress test")
		native := scenarioByName(a.Moduli.Scenarios, "native ASHA after Gate 393")
		sealedC3 := scenarioByName(a.Moduli.Scenarios, "sealed exact C3 label-triality")
		sealedS3 := scenarioByName(a.Moduli.Scenarios, "sealed exact S3 label-triality")
		sealedN := scenarioByName(a.Moduli.Scenarios, "sealed N=diag(0,1,2) hierarchy")
		checks := []theorem.Check{
			{Name: "late-gate firewall inheritance is loaded", Passed: a.Inheritance.Executed && a.Inheritance.Gate247TrialityFunctorMissing && a.Inheritance.Gate370NativeMapsCentral && a.Inheritance.Gate371NumberOperatorNonNative && a.Inheritance.Gate372ChargedModuliDim == 13 && a.Inheritance.Gate387FlavorFirewallSealed && a.Inheritance.NoEmpiricalFlavorValuesImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "domain admission rejects manual generation-to-triality relabeling", Passed: a.Domain.Executed && a.Domain.AbstractSpin8TrialityAvailable && !a.Domain.NativeTrialityCarrierFound && !a.Domain.GenerationToTrialityFunctorDerived && !a.Domain.ExplicitNativeThetaAvailable && a.Domain.ExplicitLabelPermutationThetaAvailable && !a.Domain.DomainAdmitted && a.Domain.ManualGenerationRelabelingRejected, Detail: FormatDomain(a.Domain)},
			{Name: "C3 centralizer is computed as six real dimensions of complex circulants", Passed: c3.GeneralComplexRealDim == 6 && c3.HermitianRealDim == 3 && c3.Sealed && c3.AllSectorTexturesCommute && !c3.CKMMisalignmentCapacity && c3.RankResidual < eps, Detail: FormatCentralizerCase(c3)},
			{Name: "S3 centralizer is computed as four real dimensions with 1+2 degeneracy", Passed: s3.GeneralComplexRealDim == 4 && s3.HermitianRealDim == 2 && s3.Sealed && s3.HasOnePlusTwoDegeneracy && s3.AllSectorTexturesCommute && !s3.CKMMisalignmentCapacity && s3.RankResidual < eps, Detail: FormatCentralizerCase(s3)},
			{Name: "number operator is hierarchy-capable but non-native and triality-breaking", Passed: a.Number.Executed && !a.Number.NativeDerived && a.Number.BridgeCompatible && a.Number.SealedExternalExtension && a.Number.CircularIfUsedAsSolution && a.Number.BreaksExactTriality && a.Number.ProducesDiagonalHierarchy && !a.Number.ProducesMixing && !a.Number.ProvidesTwoNoncommutingTextures, Detail: FormatNumber(a.Number)},
			{Name: "native moduli firewall is preserved at thirteen", Passed: native.Native && !native.Conditional && !native.Failed && native.ResultingDim == 13 && a.Moduli.StartingChargedDim == 13 && !a.Moduli.NativeReductionBelow13 && a.Moduli.BestNativeDim == 13, Detail: FormatModuliScenario(native)},
			{Name: "sealed C3 reduction removes CKM capacity", Passed: sealedC3.Conditional && sealedC3.ResultingDim == 9 && sealedC3.DistinctChargedMassesPossible && !sealedC3.CKMMisalignmentPossible && sealedC3.Failed, Detail: FormatModuliScenario(sealedC3)},
			{Name: "sealed S3 reduction is degenerate", Passed: sealedS3.Conditional && sealedS3.ResultingDim == 6 && !sealedS3.DistinctChargedMassesPossible && !sealedS3.CKMMisalignmentPossible && sealedS3.Failed, Detail: FormatModuliScenario(sealedS3)},
			{Name: "sealed N hierarchy is diagonal-only", Passed: sealedN.Conditional && sealedN.ResultingDim == 9 && sealedN.DistinctChargedMassesPossible && !sealedN.CKMMisalignmentPossible && sealedN.Failed, Detail: FormatModuliScenario(sealedN)},
			{Name: "firewalls remain clean", Passed: a.Firewall.Executed && a.Firewall.NoYukawaMassesImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.NoEmpiricalOrderingImported && a.Firewall.NoManualGenerationAssignment && a.Firewall.NoFakeSpin8MatricesInvented && a.Firewall.NoNativeCarrierClaimed && a.Firewall.NoModuliReductionClaimed, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate is native generation-address functor search", Passed: a.Next.Gate == 394 && a.Next.Title != "", Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: checks, Notes: []string{a.Truth}}
	}}
}

func centralizerByName(cases []CentralizerCase, name string) CentralizerCase {
	for _, c := range cases {
		if c.Name == name {
			return c
		}
	}
	return CentralizerCase{}
}

func scenarioByName(scenarios []ModuliScenario, name string) ModuliScenario {
	for _, s := range scenarios {
		if s.Name == name {
			return s
		}
	}
	return ModuliScenario{}
}
