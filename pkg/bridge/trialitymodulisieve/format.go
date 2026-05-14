package trialitymodulisieve

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t gate247_functor_missing=%t gate370_native_maps_central=%t gate371_N_non_native=%t gate372_dim=%d gate387_firewall=%t verdict=%s", x.Executed, x.Gate247TrialityFunctorMissing, x.Gate370NativeMapsCentral, x.Gate371NumberOperatorNonNative, x.Gate372ChargedModuliDim, x.Gate387FlavorFirewallSealed, x.Verdict)
}

func FormatDomain(x DomainAdmission) string {
	return fmt.Sprintf("executed=%t abstract_triality=%t native_carrier=%t functor=%t native_theta=%t label_theta=%t admitted=%t rejected_manual=%t missing=[%s] verdict=%s", x.Executed, x.AbstractSpin8TrialityAvailable, x.NativeTrialityCarrierFound, x.GenerationToTrialityFunctorDerived, x.ExplicitNativeThetaAvailable, x.ExplicitLabelPermutationThetaAvailable, x.DomainAdmitted, x.ManualGenerationRelabelingRejected, strings.Join(x.MissingPieces, " | "), x.Verdict)
}

func FormatCentralizerCase(x CentralizerCase) string {
	return fmt.Sprintf("%s constraints=[%s] complex_real_dim=%d hermitian_real_dim=%d canonical=%q singular_values=%d degeneracy_1+2=%t all_commute=%t ckm_capacity=%t native=%t sealed=%t residual=%.3g verdict=%s", x.Name, strings.Join(x.Constraints, " | "), x.GeneralComplexRealDim, x.HermitianRealDim, x.CanonicalForm, x.DistinctSingularValuesGeneric, x.HasOnePlusTwoDegeneracy, x.AllSectorTexturesCommute, x.CKMMisalignmentCapacity, x.Native, x.Sealed, x.RankResidual, x.Verdict)
}

func FormatCentralizer(x CentralizerAudit) string {
	parts := make([]string, 0, len(x.Cases))
	for _, c := range x.Cases {
		parts = append(parts, fmt.Sprintf("%s:%d", c.Name, c.GeneralComplexRealDim))
	}
	return fmt.Sprintf("executed=%t cases={%s} verdict=%s", x.Executed, strings.Join(parts, ", "), x.Verdict)
}

func FormatNumber(x NumberOperatorAudit) string {
	return fmt.Sprintf("executed=%t N=%v status=%q native=%t bridge=%t sealed=%t circular=%t [N,C3]=%.6g [N,R]=%.6g breaks_triality=%t diagonal_hierarchy=%t mixing=%t two_noncommuting=%t verdict=%s", x.Executed, x.Operator, x.Status, x.NativeDerived, x.BridgeCompatible, x.SealedExternalExtension, x.CircularIfUsedAsSolution, x.CommNormWithCycle, x.CommNormWithMirror, x.BreaksExactTriality, x.ProducesDiagonalHierarchy, x.ProducesMixing, x.ProvidesTwoNoncommutingTextures, x.Verdict)
}

func FormatModuliScenario(x ModuliScenario) string {
	return fmt.Sprintf("%s assumption=%q start=%d result=%d distinct_masses=%t ckm=%t q/l_separation=%t native=%t conditional=%t failed=%t reason=%q verdict=%s", x.Name, x.AssumptionClass, x.StartingChargedDim, x.ResultingDim, x.DistinctChargedMassesPossible, x.CKMMisalignmentPossible, x.LeptonQuarkSectorSeparation, x.Native, x.Conditional, x.Failed, x.Reason, x.Verdict)
}

func FormatModuli(x ModuliAudit) string {
	return fmt.Sprintf("executed=%t start=%d native_reduction=%t best_native=%d best_conditional=%d verdict=%s", x.Executed, x.StartingChargedDim, x.NativeReductionBelow13, x.BestNativeDim, x.BestConditionalDim, x.Verdict)
}

func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("executed=%t no_masses=%t no_ckm=%t no_pmns=%t no_empirical_ordering=%t no_manual_assignment=%t no_fake_spin8=%t no_native_carrier_claimed=%t no_moduli_reduction_claimed=%t verdict=%s", x.Executed, x.NoYukawaMassesImported, x.NoCKMImported, x.NoPMNSImported, x.NoEmpiricalOrderingImported, x.NoManualGenerationAssignment, x.NoFakeSpin8MatricesInvented, x.NoNativeCarrierClaimed, x.NoModuliReductionClaimed, x.Verdict)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}
