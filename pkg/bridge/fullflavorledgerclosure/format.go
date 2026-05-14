package fullflavorledgerclosure

import (
	"fmt"
	"strings"
)

func FormatInheritance(a Gate266Inheritance) string {
	return fmt.Sprintf("seal=%t quarkCKM=%t chargedSVD=%t takagi=%t pmns=%t largePMNS=%t quarkNative=%t leptonNative=%t boundary=%t majoranaDerived=%t representative=%t verdict=%q", a.EmpiricalYukawaSealActive, a.QuarkSVDCKMReconstructed, a.ChargedLeptonSVDReconstructed, a.NeutrinoTakagiReconstructed, a.PMNSReconstructed, a.LargeAngleLeptonStructure, a.QuarkNativeDerivation, a.LeptonNativeDerivation, a.EmpiricalBoundaryPreserved, a.MajoranaNatureFiniteDerived, a.RepresentativeDataOnly, a.Verdict)
}

func FormatGeometricLedger(a GeometricDerivationLedger) string {
	return fmt.Sprintf("items=%d S_C=%t algebra=%t gaugeCharge=%t gen=%t tauSource=%t adTau=%t trialityBasis=%t amps=%t ckmPmns=%t masses=%t verdict=%q", len(a.Items), a.SCarrierSpaceRecorded, a.FiniteAlgebraRecorded, a.GaugeMatterChargeRecorded, a.ThreeGenerationCapacity, a.TauEtaSourceMapRecorded, a.AdTauMixingComplementRecorded, a.TrialityHermitianBasisRecorded, a.YukawaAmplitudeDerived, a.CKMPMNSDerived, a.FermionMassesDerived, a.Verdict)
}

func FormatEmpiricalLedger(a EmpiricalInputLedger) string {
	return fmt.Sprintf("items=%d ssbSeal=%t yukawaSeal=%t weakFrame=%t scalarVEV=%t quark=%t lepton=%t ckm=%t pmns=%t majorana=%t noRewrite=%t verdict=%q", len(a.Items), a.SpontaneousCarrierSealActive, a.EmpiricalYukawaSealActive, a.WeakFrameOrientationSealed, a.ScalarVEVAlignmentSealed, a.QuarkTexturesSealed, a.LeptonTexturesSealed, a.CKMEntriesSealed, a.PMNSEntriesSealed, a.MajoranaChoiceSealed, a.DoesNotRewriteFiniteCore, a.Verdict)
}

func FormatReconstruction(a ReconstructionVerification) string {
	return fmt.Sprintf("quarkCKM=%t chargedSVD=%t takagi=%t pmns=%t svdRecon=%t takagiRecon=%t worksOnData=%t predicts=%t polluted=%t verdict=%q", a.QuarkSVDCKMVerified, a.ChargedLeptonSVDVerified, a.MajoranaTakagiVerified, a.PMNSVerified, a.SVDIsAlgebraicReconstruction, a.TakagiIsAlgebraicReconstruction, a.ObservablePipelineWorksOnData, a.ObservablePipelinePredictsData, a.FiniteCorePolluted, a.Verdict)
}

func FormatFutureCriteria(a FutureTheoremCriteria) string {
	missing := make([]string, 0, len(a.Criteria))
	for _, c := range a.Criteria {
		if c.Required && !c.Satisfied {
			missing = append(missing, c.Name)
		}
	}
	return fmt.Sprintf("criteria=%d missing=[%s] spectral=%t DF=%t heat=%t yukawaMap=%t hopfBGap=%t norm=%t prediction=%t canLift=%t next=%q verdict=%q", len(a.Criteria), strings.Join(missing, "; "), a.RequiresFiniteSpectralAction, a.RequiresCanonicalFiniteDirac, a.RequiresHeatKernelCoefficients, a.RequiresYukawaAmplitudeMap, a.RequiresHopfOrBGapProjection, a.RequiresNormalizationScheme, a.RequiresMassAndMixingPrediction, a.CurrentGateCanLiftSeal, a.RecommendedNextGate, a.Verdict)
}

func FormatFirewall(a FirewallManifest) string {
	return fmt.Sprintf("kinematics=%t dynamicsSealed=%t ssb=%t empirical=%t noMass=%t noCKM=%t noMajorana=%t noSpectral=%t noNewPhysics=%t polluted=%t verdict=%q", a.KinematicsDerived, a.DynamicsSealed, a.SpontaneousCarrierSealPreserved, a.EmpiricalYukawaSealPreserved, a.NoMassPredictionClaim, a.NoCKMPMNSPredictionClaim, a.NoMajoranaNatureClaim, a.NoSpectralActionClaim, a.ClosureDoesNotAddNewPhysics, a.FiniteCorePolluted, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("gate266=%t geometric=%t empirical=%t recon=%t future=%t closed=%t nativeDynamics=%t status=%q next=%q comment=%q", a.Gate266Inherited, a.GeometricLedgerClosed, a.EmpiricalLedgerClosed, a.ReconstructionsVerified, a.FutureCriteriaDefined, a.FullFlavorLedgerClosed, a.NativeFlavorDynamicsDerived, a.Status, a.NextGate, a.Comment)
}

func FormatLedgerItems(items []LedgerItem) string {
	parts := make([]string, 0, len(items))
	for _, it := range items {
		parts = append(parts, fmt.Sprintf("%s[%s|%s]", it.Name, it.SourceGate, it.Status))
	}
	return strings.Join(parts, "; ")
}
