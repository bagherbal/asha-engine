package generation2paulimomentweakplaneincidenceaudit

import (
	"fmt"
	"sort"
	"strings"
)

func FormatInherited(a InheritedAudit) string {
	return fmt.Sprintf("pauli=%t hopf=%t scalar4to1plus3=%t moment3split=%t noTransfer=%t verdict=%q", a.Gate560PauliTriplet, a.Gate560HopfIdentity, a.Gate560ScalarFourToOnePlus3, a.Gate560ScalarMomentThreeSplit, a.Gate560NoTransferFunctor, a.Verdict)
}

func FormatSpatialLabels(a SpatialLabelAudit) string {
	return fmt.Sprintf("carrier=%q labels=%v modes=%v BL=%q insideBL=%t nativeMetricOrient=%t basisOnly=%t metricCert=%t orientCert=%t verdict=%q", a.CarrierName, a.Labels, a.CorrespondingFockModes, a.BLSpatialEigenvalue, a.InsideBLSpatialEigenspace, a.NativeOrientedMetricThreeSpace, a.BasisConventionOnly, a.MetricCertificateAvailable, a.OrientationCertificateAvailable, a.Verdict)
}

func FormatIncidence(a WeakPlaneIncidenceAudit) string {
	keys := make([]string, 0, len(a.PlaneToBivector))
	for k := range a.PlaneToBivector {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s->%s", k, a.PlaneToBivector[k]))
	}
	return fmt.Sprintf("bivectors={%s} dim=%d coordinate=%t nativeSelector=%t notational=%t verdict=%q", strings.Join(parts, ", "), a.IncidenceDimension, a.CoordinateBivectorsAvailable, a.NativeIncidenceSelector, a.NotationalOnly, a.Verdict)
}

func FormatHodge(a HodgeAudit) string {
	return fmt.Sprintf("requiresMetricOrient=%t metric=%t orient=%t rules=%v formalNormalPlane=%t nativeHodge=%t verdict=%q", a.RequiresMetricAndOrientation, a.MetricAvailableNatively, a.OrientationAvailableNatively, a.FormalRules, a.FormalNormalSelectsPlane, a.NativeHodgeStarConstructed, a.Verdict)
}

func FormatIntertwiner(a IntertwinerSearchAudit) string {
	return fmt.Sprintf("source=%q targetVector=%q targetBivector=%q toSpatial=%t toIncidence=%t basisIndependent=%t metricCompatible=%t manualSigma3ToS3=%t verdict=%q", a.Source, a.TargetVector, a.TargetBivector, a.MapToSpatialFound, a.MapToIncidenceFound, a.BasisIndependent, a.UnitMetricCompatible, a.ManualSigma3ToS3Assignment, a.Verdict)
}

func FormatPlane(a CanonicalPlaneAudit) string {
	return fmt.Sprintf("intertwiner=%t sigma3=%t nonzeroMu=%t U12=%t U13=%t U23=%t basisDependentOnly=%t verdict=%q", a.IntertwinerExists, a.Sigma3AxisAvailable, a.NonzeroMuAvailable, a.CanonicalU12, a.CanonicalU13, a.CanonicalU23, a.OnlyBasisDependentPlane, a.Verdict)
}

func FormatBL(a BLCompatibilityAudit) string {
	return fmt.Sprintf("insideW=%t mixesLepton=%t BL=%q formalCommutes=%t nontrivial=%t BLSuppliesLabels=%t verdict=%q", a.SelectionInsideWSpatial, a.MixesLeptonSlot, a.BLRestrictedToWSpatial, a.FormalSelectionCommutesWithBL, a.CompatibilityNontrivial, a.BLSuppliesPlaneLabels, a.Verdict)
}

func FormatSpectral(a SpectralTripleAudit) string {
	return fmt.Sprintf("functor=%t gamma=%t J=%t D=%t firstOrder=%t oneFormRelation=%t passed=%t missing=%v verdict=%q", a.IncidenceFunctorFound, a.GradingCheckAvailable, a.JCheckAvailable, a.DCheckAvailable, a.FirstOrderCheckAvailable, a.FiniteOneFormRelationFound, a.CompatibilityPassed, a.MissingData, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("preserved=%t weakIso=%t gaugeBosons=%t photon=%t generation=%t yukawa=%t ckm=%t observed=%t higgs=%t verdict=%q", a.Preserved, a.WeakIsospinIdentified, a.GaugeBosonsIdentified, a.PhotonIdentified, a.GenerationHierarchyIdentified, a.YukawaTextureDerived, a.CKMPMNSDerived, a.ObservedFlavorImported, a.HiggsLanePromoted, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("S_nativeMetric=%t planesNative=%t hodgeNative=%t intertwiner=%t selectsPlane=%t lawfulTransfer=%t next=%q verdict=%q", a.SSpatialNativeOrientedMetric, a.WeakPlanesNativeBivectors, a.HodgeStarNative, a.PauliToIncidenceIntertwiner, a.ScalarMomentSelectsWeakPlane, a.LawfulTransferAvailable, a.MissingNextTheorem, a.Verdict)
}
