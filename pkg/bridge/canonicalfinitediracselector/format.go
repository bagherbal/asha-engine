package canonicalfinitediracselector

import (
	"fmt"
	"strings"
)

func FormatInheritance(a Gate268Inheritance) string {
	return fmt.Sprintf("scaffold=%t formalDF=%t raw=%t dependence=%t canonicalDF=%t higgs=%t firewall=%t next=%q verdict=%q", a.ScaffoldRetrieved, a.FormalDFFamilyAvailable, a.RawMomentsEvaluated, a.MomentDependenceExposed, a.CanonicalDFDerived, a.HiggsRatioDerived, a.FirewallPreserved, a.RecommendedNextGate, a.Verdict)
}

func FormatDefinition(a OrderOneDefinition) string {
	return fmt.Sprintf("formula=%q algebra=%q D=%q J=%q needsRep=%t needsOpposite=%t allAB=%t defined=%t verdict=%q", a.Formula, a.AlgebraSymbol, a.DiracSymbol, a.RealStructureSymbol, a.RequiresRepresentation, a.RequiresOppositeAction, a.RequiresAllAAndB, a.Defined, a.Verdict)
}

func FormatAlgebra(a AlgebraRepresentationAudit) string {
	return fmt.Sprintf("algebra=%q carrier=%q dim=%d mode=%t fullSC=%t left=%t right=%t opposite=%t physicalJ=%t oneForms=%t toy=%t importedConnes=%t verdict=%q", a.NativeAlgebraName, a.ModeCarrier, a.ModeDimension, a.ModeLevelCPlusM3Available, a.FullSCRepresentationDerived, a.LeftRepresentationDerived, a.RightRepresentationDerived, a.OppositeRepresentationDerived, a.PhysicalJDerived, a.NonVacuousOneFormsAvailable, a.ToyModePreflightAllowed, a.ImportedConnesAlgebra, a.Verdict)
}

func FormatConstraint(r ConstraintRow) string {
	return fmt.Sprintf("%s: %s -> %s; ok=%t; reason=%s", r.Name, r.Before, r.After, r.Satisfied, r.Reason)
}

func FormatConstraints(rows []ConstraintRow) string {
	parts := make([]string, len(rows))
	for i, r := range rows {
		parts[i] = FormatConstraint(r)
	}
	return strings.Join(parts, " | ")
}

func FormatSieve(a GenericDiracBlockAudit) string {
	return fmt.Sprintf("carrier=%q shape=%q paramsC=%d->%d paramsR=%d->%d eliminatedC=%d constraint=%q leakageRemoved=%t colorAnisotropyRemoved=%t family=%q nontrivial=%t fullSC=%t oneFormsVanish=%t canonical=%t verdict=%q", a.Carrier, a.GenericMatrixShape, a.InitialComplexParameters, a.AllowedComplexParameters, a.InitialRealParameters, a.AllowedRealParameters, a.EliminatedComplexParameters, a.OrderOneToyConstraint, a.TemporalSpatialLeakageRemoved, a.ColorAnisotropyRemoved, a.AllowedFamilyFormula, a.SieveNontrivial, a.SievePhysicalOnFullSC, a.OneFormsVanishForAllowedFamily, a.CanonicalBlockSelected, a.Verdict)
}

func FormatMomentRow(r MomentRow) string {
	return fmt.Sprintf("%s x=%.6g y=%.6g sigmas=%v TrD2=%.12g TrD4=%.12g ratio=%.12g allowed=%t canonical=%t comment=%q", r.Name, r.X, r.Y, r.SingularValues, r.TraceD2, r.TraceD4, r.RawRatio, r.OrderOneAllowed, r.Canonical, r.Comment)
}

func FormatMoments(a SpectralMomentReevaluation) string {
	parts := make([]string, len(a.Rows))
	for i, r := range a.Rows {
		parts[i] = FormatMomentRow(r)
	}
	return fmt.Sprintf("computed=%t allAllowed=%t stable=%t amplitudeDependent=%t sdw=%t higgs=%t rows=[%s] verdict=%q", a.MomentsRecomputed, a.AllRowsOrderOneAllowed, a.RawRatioStableAcrossAllowedDF, a.DependsOnSurvivingAmplitudes, a.SeeleyDeWittMapDerived, a.HiggsRatioDerived, strings.Join(parts, " | "), a.Verdict)
}

func FormatCanonical(a CanonicalDFVerdict) string {
	return fmt.Sprintf("sieve=%t unique=%t dimC=%d additionalSelector=%t normalizationOnly=%t norm=%t gauge=%t scalar=%t promotable=%t verdict=%q", a.OrderOneSieveNontrivial, a.UniqueDFSelected, a.SurvivingFamilyDimensionC, a.RequiresAdditionalSelector, a.CouldUseNormalizationOnly, a.NormalizationDerived, a.GaugeProjectionDerived, a.ScalarFluctuationMapDerived, a.PromotableFiniteDiracOperator, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("empiricalSeal=%t ssbSeal=%t noMass=%t noVEV=%t noCutoff=%t noConnes=%t noYukawaFit=%t toyNotPromoted=%t noHiggs=%t polluted=%t verdict=%q", a.EmpiricalYukawaSealPreserved, a.SpontaneousCarrierSealPreserved, a.NoObservedMassInserted, a.NoVEVInserted, a.NoCutoffScaleInserted, a.NoConnesAlgebraImported, a.NoYukawaFitUsed, a.ToySieveNotPromoted, a.NoHiggsPredictionClaim, a.FiniteCorePolluted, a.Verdict)
}

func FormatFuture(a FutureMap) string {
	missing := make([]string, 0, len(a.Obligations))
	for _, o := range a.Obligations {
		if o.Required && !o.Satisfied {
			missing = append(missing, o.Name)
		}
	}
	return fmt.Sprintf("obligations=%d missing=[%s] faithfulSC=%t oppositeJ=%t oneForms=%t weakH=%t canonicalAmp=%t heat=%t next=%q verdict=%q", len(a.Obligations), strings.Join(missing, "; "), a.NeedFaithfulSCRep, a.NeedPhysicalOppositeJ, a.NeedNonVacuousOneForms, a.NeedQuaternionicOrWeakH, a.NeedCanonicalAmplitude, a.NeedHeatKernelMap, a.RecommendedNextGate, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("gate268=%t orderOne=%t modeAlgebra=%t sieve=%t canonicalDF=%t ratioStable=%t higgs=%t firewall=%t status=%q next=%q comment=%q", a.Gate268Inherited, a.OrderOneDefined, a.ModeAlgebraPreflight, a.OrderOneSieveReduced, a.CanonicalDFDerived, a.AllowedMomentRatioStable, a.HiggsRatioDerived, a.FirewallPreserved, a.Status, a.NextGate, a.Comment)
}
