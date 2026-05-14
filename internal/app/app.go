package app

import (
	"fmt"

	"github.com/bagherbal/asha-engine/internal/report"
	"github.com/bagherbal/asha-engine/pkg/bridge/actionscale"
	"github.com/bagherbal/asha-engine/pkg/bridge/betacoeff"
	"github.com/bagherbal/asha-engine/pkg/bridge/condensate"
	"github.com/bagherbal/asha-engine/pkg/bridge/couplingnorm"
	"github.com/bagherbal/asha-engine/pkg/bridge/ewprojection"
	"github.com/bagherbal/asha-engine/pkg/bridge/fieldmap"
	"github.com/bagherbal/asha-engine/pkg/bridge/goldstone"
	"github.com/bagherbal/asha-engine/pkg/bridge/rgflow"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarcomplex"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarscale"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarsu2"
	"github.com/bagherbal/asha-engine/pkg/bridge/threshold"
	"github.com/bagherbal/asha-engine/pkg/bridge/thresholdactivation"
	"github.com/bagherbal/asha-engine/pkg/bridge/thresholdrep"
	"github.com/bagherbal/asha-engine/pkg/clifford"
	"github.com/bagherbal/asha-engine/pkg/dynamics/bsector"
	"github.com/bagherbal/asha-engine/pkg/dynamics/higgspotential"
	"github.com/bagherbal/asha-engine/pkg/dynamics/scalarpotential"
	"github.com/bagherbal/asha-engine/pkg/engine/cache"
	"github.com/bagherbal/asha-engine/pkg/exterior"
	"github.com/bagherbal/asha-engine/pkg/gauge"
	"github.com/bagherbal/asha-engine/pkg/gauge/boundary"
	"github.com/bagherbal/asha-engine/pkg/gauge/connection"
	"github.com/bagherbal/asha-engine/pkg/gauge/higgs"
	"github.com/bagherbal/asha-engine/pkg/gauge/lift"
	"github.com/bagherbal/asha-engine/pkg/geometry/boolean"
	"github.com/bagherbal/asha-engine/pkg/geometry/contact"
	"github.com/bagherbal/asha-engine/pkg/geometry/g2"
	"github.com/bagherbal/asha-engine/pkg/matter"
	"github.com/bagherbal/asha-engine/pkg/matter/action"
	"github.com/bagherbal/asha-engine/pkg/matter/bfbridge"
	"github.com/bagherbal/asha-engine/pkg/matter/bfcurvature"
	"github.com/bagherbal/asha-engine/pkg/matter/bfsource"
	"github.com/bagherbal/asha-engine/pkg/matter/charge"
	"github.com/bagherbal/asha-engine/pkg/matter/electroweak"
	"github.com/bagherbal/asha-engine/pkg/matter/embedding"
	"github.com/bagherbal/asha-engine/pkg/matter/gencurvature"
	"github.com/bagherbal/asha-engine/pkg/matter/generationbreak"
	"github.com/bagherbal/asha-engine/pkg/matter/hyperaudit"
	"github.com/bagherbal/asha-engine/pkg/matter/hypercharge"
	"github.com/bagherbal/asha-engine/pkg/matter/sourceaction"
	"github.com/bagherbal/asha-engine/pkg/matter/sourcemap"
	"github.com/bagherbal/asha-engine/pkg/matter/sourcepotential"
	"github.com/bagherbal/asha-engine/pkg/matter/su2l"
	"github.com/bagherbal/asha-engine/pkg/matter/su2lgauge"
	"github.com/bagherbal/asha-engine/pkg/matter/t3r"
	"github.com/bagherbal/asha-engine/pkg/matter/tensor"
	"github.com/bagherbal/asha-engine/pkg/matter/texture"
	"github.com/bagherbal/asha-engine/pkg/matter/trialityyukawa"
	"github.com/bagherbal/asha-engine/pkg/matter/yukawa"
	"github.com/bagherbal/asha-engine/pkg/matter/yukawaintertwiner"
	"github.com/bagherbal/asha-engine/pkg/phase"
	"github.com/bagherbal/asha-engine/pkg/spinor"
	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Run() error {
	registry := theorem.NewRegistry(
		exterior.GradeStructureTheorem(8),
		clifford.StructureTheorem(clifford.Signature{Positive: 1, Negative: 7}),
		phase.CovariantPhaseSpaceTheorem(4),
		boolean.IncidenceSupportTheorem(8, 3, 4),
		g2.CalibrationSupportTheorem(),
		contact.ContactSpaceTheorem(),
		bsector.ContactVacuumTheorem(),
		gauge.ContactCentralizerTheorem(),
		lift.BooleanCompressionTheorem(),
		boundary.BoundaryFixedClosureTheorem(),
		connection.ProjectedConnectionTheorem(),
		higgs.VacuumMixingTheorem(),
		higgspotential.PotentialCandidateTheorem(),
		spinor.FockSpaceTheorem(),
		matter.FockContactBridgeTheorem(),
		action.RepresentationActionTheorem(),
		embedding.CanonicalEmbeddingTheorem(),
		charge.ChargePolarizationTheorem(),
		tensor.TensorFactorBridgeTheorem(),
		yukawa.IntertwinerSelectionTheorem(),
		electroweak.OperatorSearchTheorem(),
		hypercharge.ScalarHyperchargeBridgeTheorem(),
		t3r.MatterT3RSearchTheorem(),
		hyperaudit.HyperchargeTableAuditTheorem(),
		su2l.DoubletAuditTheorem(),
		su2lgauge.GeneratorAuditTheorem(),
		yukawaintertwiner.GaugeCompatibleYukawaTheorem(),
		trialityyukawa.GenerationTrialityYukawaTheorem(),
		texture.GenerationBreakingTextureSearchTheorem(),
		generationbreak.FiniteGenerationBreakingSearchTheorem(),
		gencurvature.CurvatureOnGenerationCarrierTheorem(),
		bfbridge.ActiveGenerationProjectionBridgeTheorem(),
		bfcurvature.FiniteMaurerCartanCurvatureTheorem(),
		bfsource.BFActionSourceTextureTheorem(),
		sourcemap.SourceTensorSelectionTheorem(),
		sourceaction.SourceTensorActionTheorem(),
		sourcepotential.SymmetryBreakingSourceActionTheorem(),
		scalarpotential.EffectivePotentialTheorem(),
		scalarscale.ScaleBridgeSearchTheorem(),
		actionscale.ActionNormalizationScaleBridgeTheorem(),
		couplingnorm.CouplingNormalizationBridgeTheorem(),
		ewprojection.ElectroweakProjectionTheorem(),
		rgflow.RGBoundaryFlowAuditTheorem(),
		betacoeff.FiniteSpectrumBetaCoefficientAuditTheorem(),
		threshold.ThresholdSpectrumMatchingAuditTheorem(),
		thresholdrep.ThresholdRepresentationAssignmentAuditTheorem(),
		thresholdactivation.ThresholdActivationDecouplingAuditTheorem(),
		fieldmap.FiniteToContinuumFieldMapAuditTheorem(),
		goldstone.GaugeEatingGoldstoneAuditTheorem(),
		scalarsu2.ScalarContactSU2ActionSearchTheorem(),
		scalarcomplex.ScalarComplexQuaternionicStructureSearchTheorem(),
		condensate.CompositeHiggsCondensateDirectionAuditTheorem(),
		cache.RuntimeFixtureCacheTheorem(),
	)

	results := registry.RunAll()
	fmt.Print(report.RenderTerminal(results))
	return nil
}
