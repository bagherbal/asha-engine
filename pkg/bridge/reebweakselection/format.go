package reebweakselection

import (
	"fmt"
	"strings"
)

func FormatContact(a ContactStructureAudit) string {
	return fmt.Sprintf("dim=%d expected=%d index=%.10f frameResidual=%.3e booleanResidual=%.3e g2Residual=%.3e projector=%t eta=%t deta=%t reeb=%t verdict=%s", a.ContactDimension, a.ExpectedContactDim, a.ContactIndex, a.FrameIsometryResidual, a.BooleanContainment, a.G2Containment, a.ContactProjectorExists, a.EtaOneFormDerived, a.DEtaTwoFormDerived, a.ReebVectorDerived, a.Verdict)
}

func FormatReeb(a ReebVectorAudit) string {
	return fmt.Sprintf("definition=%q etaCond=%q contraction=%q candidate=%t nativeVacuum=%t livesOnK=%t mappedW=%t mappedSpatial=%t manualAxis=%t verdict=%s", a.Definition, a.EtaOfReebCondition, a.ContractionCondition, a.CandidateAvailable, a.NativeFromVacuumStabilizer, a.LivesOnContactK, a.MappedToFockGeneratorW, a.MappedToSpatialFockAxes, a.ManualAxisChoice, a.Verdict)
}

func FormatProjection(a SpatialProjectionAudit) string {
	return fmt.Sprintf("axes=%v pureSpatial=%v map=%s derived=%t components=%v source=%q uniformOrAbsent=%t tagged=%t taggedAxis=%q s3Broken=%t verdict=%s", a.SpatialAxes, a.PureSpatialPlanes, a.ProjectionMapName, a.KToWProjectionDerived, a.ReebComponentsOnSpatialAxes, a.ComponentSource, a.ComponentsAreUniformOrAbsent, a.UniqueSpatialAxisTagged, a.TaggedAxis, a.S3PermutationBroken, a.Verdict)
}

func FormatPlanes(rows []PlaneReebAudit) string {
	parts := make([]string, 0, len(rows))
	for _, p := range rows {
		parts = append(parts, fmt.Sprintf("%s modes=%v complement=%s inherited=%t survivesU1=%t requiresTaggedAxis=%t selected=%t rule=%q verdict=%s", p.Plane, p.ModeIndices, p.ComplementAxis, p.InheritedFromGate240, p.SurvivesU1Twist, p.RequiresTaggedAxis, p.SelectedByReeb, p.SelectionRule, p.Verdict))
	}
	return strings.Join(parts, " | ")
}

func FormatSieve(a ReebSieveAudit) string {
	return fmt.Sprintf("inherited=%v candidates=%d taggedAxis=%q selected=%v rule=%q s3Broken=%t unique=%t verdict=%s", a.InheritedGate240Candidates, a.CandidatePlaneCount, a.TaggedAxis, a.SelectedPlanes, a.HypotheticalRule, a.S3DegeneracyBroken, a.UniqueWeakPlaneSelected, a.Verdict)
}

func FormatWeak(a WeakOutcomeAudit) string {
	return fmt.Sprintf("gate240PureSpatial=%t contact=%t reeb=%t contactToFock=%t uniquePlane=%t physicalLeft=%t globalH=%t orderOne=%t verdict=%s", a.Gate240ReducedToPureSpatial, a.ContactGeometryAvailable, a.ReebSelectorDerived, a.ContactToFockBridgeDerived, a.UniqueWeakPlaneSelected, a.PhysicalLeftHandedDerived, a.GlobalHSummandDerived, a.OrderOneReady, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("forcedReeb=%t importedContactCoords=%t importedSMWeak=%t importedChirality=%t projectorToReeb=%t claimedH=%t claimedOrderOne=%t polluted=%t verdict=%s", a.ForcedReebAxis, a.ImportedContactCoordinates, a.ImportedSMWeakPlane, a.ImportedElectroweakChirality, a.PromotedProjectorToReeb, a.ClaimedGlobalH, a.ClaimedOrderOne, a.FiniteCorePolluted, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("contactK=%t etaDeta=%t reeb=%t projection=%t spatialTagged=%t pureSpatial=%d uniquePlane=%t physicalChirality=%t globalH=%t status=%s next=%q comment=%q", a.ContactKAvailable, a.EtaDEtaDerived, a.ReebVectorDerived, a.ContactToFockProjection, a.SpatialAxisTagged, a.PureSpatialPlanesInherited, a.UniqueWeakPlaneDerived, a.PhysicalChiralityDerived, a.GlobalHDerived, a.Status, a.NextGate, a.Comment)
}
