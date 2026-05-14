package empiricalyukawafit

import (
	"fmt"
	"strings"
)

func FormatInheritance(a Gate263Inheritance) string {
	return fmt.Sprintf("ansatz=%t tau=%t basis=%t actionRule=%t previousTexture=%t previousCKM=%t previousMasses=%t orthogonal=%t norms=%v formula=%q free=%s verdict=%q", a.GeometricAnsatzAvailable, a.DiagonalTauSourceAvailable, a.HermitianOffDiagonalBasisExists, a.FiniteActionCoefficientRule, a.PreviousPhysicalTextureDerived, a.PreviousCKMPMNSDerived, a.PreviousFermionMassesDerived, a.BasisOrthogonal, a.BasisNorms, a.CandidateFormula, strings.Join(a.FreeParameters, ","), a.Verdict)
}

func FormatSeal(a EmpiricalSealActivation) string {
	return fmt.Sprintf("name=%s active=%t gate=%d boundary=%q quarantined=%t derived=%t numericalDerived=%t fitOnly=%t rewritesNoGo=%t stress=%t prediction=%t verdict=%q", a.Name, a.Activated, a.ActivatedByGate, a.BoundaryDataKind, a.ExplicitlyQuarantined, a.DerivedFromFiniteCore, a.NumericalOutputsDerived, a.PhenomenologicalFitOnly, a.RewritesGate263NoGo, a.AllowsStressFit, a.AllowsFinitePrediction, a.Verdict)
}

func FormatData(a QuarkFlavorData) string {
	return fmt.Sprintf("source=%q representative=%t mixedScale=%t masses=%t ckm=%t up=%v down=%v wolf=(lambda=%.6g,A=%.6g,rho=%.6g,eta=%.6g) params=%d ansatz=%d deficit=%d verdict=%q", a.SourceLabel, a.RepresentativeNotPrecision, a.MixedScaleWarning, a.UsesObservedMassHierarchy, a.UsesObservedCKMParameters, a.UpMassesGeV, a.DownMassesGeV, a.WolfensteinLambda, a.WolfensteinA, a.WolfensteinRhoBar, a.WolfensteinEtaBar, a.DataParameterCount, a.AnsatzQuarkParameterCount, a.ParameterDeficit, a.Verdict)
}

func FormatFit(a ProjectionFit) string {
	return fmt.Sprintf("sector=%q convention=%q alpha=%.12g beta=%.12g gamma=%.12g targetNorm=%.12g projNorm=%.12g residual=%.12g rel=%.12g tol=%.1e exact=%t offdiag=%v equalFail=%t diagFail=%t verdict=%q", a.Sector, a.TargetConvention, a.Alpha, a.Beta, a.Gamma, a.TargetFrobeniusNorm, a.ProjectionFrobeniusNorm, a.ResidualFrobeniusNorm, a.RelativeResidual, a.ExactFitTolerance, a.FitsExactly, a.TargetOffDiagonalAbs, a.EqualOffDiagonalFailure, a.DiagonalShapeFailure, a.Verdict)
}

func FormatFits(a []ProjectionFit) string {
	out := make([]string, 0, len(a))
	for _, f := range a {
		out = append(out, FormatFit(f))
	}
	return "[" + strings.Join(out, "; ") + "]"
}

func FormatViability(a StructuralViabilityAudit) string {
	return fmt.Sprintf("physicalParams=%d ansatzParams=%d deficit=%d sameBasis=%t fullMatrices=%t anyExact=%t allExact=%t combinedRel=%.12g violates=%t ckmDerived=%t massesDerived=%t verdict=%q", a.QuarkFlavorPhysicalParameters, a.RestrictedAnsatzParameters, a.ParameterDeficit, a.SameBasisForAllSectors, a.RequiresFullYukawaMatrices, a.AnySectorExactFit, a.AllSectorsExactFit, a.CombinedRelativeResidual, a.ViolatesAnsatz, a.CKMNumericalFitDerived, a.MassSpectrumDerived, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("seal=%t quarantined=%t noRewrite=%t noMassPred=%t noCKMPred=%t noVEV=%t noProjectionLaw=%t gate263=%t fullSeal=%t polluted=%t verdict=%q", a.EmpiricalSealActive, a.ObservedDataQuarantined, a.DoesNotRewriteFiniteCore, a.DoesNotClaimMassPrediction, a.DoesNotClaimCKMPrediction, a.DoesNotInferVEVOrThresholds, a.DoesNotPromoteProjectionToLaw, a.Gate263NoGoPreserved, a.FullEmpiricalSealStillRequired, a.FiniteCorePolluted, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("gate263=%t seal=%t data=%t fits=%t violates=%t fullMatrices=%t masses=%t ckm=%t status=%q next=%q comment=%q", a.Gate263Inherited, a.EmpiricalSealActivated, a.RepresentativeDataIngested, a.FitsRestrictedAnsatz, a.ViolatesRestrictedAnsatz, a.FullEmpiricalMatricesRequired, a.MassesDerived, a.CKMDerived, a.Status, a.NextGate, a.Comment)
}
