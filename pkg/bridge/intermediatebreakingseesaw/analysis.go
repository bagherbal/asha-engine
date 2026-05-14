// Package intermediatebreakingseesaw implements Gate 231:
// IntermediateBreakingSeal activation / Type-I neutrino seesaw preflight audit.
//
// Gate 230 proved that the finite core does not currently derive the Hopf
// instanton, Hopf-action localization map, or hidden order parameter needed to
// make the intermediate scale a finite theorem. Gate 231 therefore activates an
// explicit IntermediateBreakingSeal as a phenomenological boundary condition and
// asks whether the sealed intermediate scale can act as a right-handed neutrino
// Majorana scale in a Type-I seesaw preflight.
//
// The audit intentionally corrects a common arithmetic trap: with v≈246 GeV
// and M_R≈6.65e11 GeV, an order-one Dirac Yukawa gives mν≈91 eV, not 0.09 eV.
// A phenomenologically viable active-neutrino scale requires a small Dirac
// Yukawa, yν≈0.01–0.03, which is permitted only behind the existing empirical
// Yukawa-amplitude firewall and is not a finite derivation.
package intermediatebreakingseesaw

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/finitehopfaction"
)

const (
	AuditID = "GATE231-INTERMEDIATE-BREAKING-SEAL-NEUTRINO-SEESAW-PREFLIGHT-AUDIT"

	StatusIntermediateSealActivated        = "INTERMEDIATE_BREAKING_SEAL_ACTIVATED_PHENOMENOLOGICALLY"
	StatusOrderOneSeesawFailed             = "FAILED_ROUTE_ORDER_ONE_TYPE_I_SEESAW_RESONANCE"
	StatusSmallYukawaConditionalSupport    = "CONDITIONAL_SUPPORT_TYPE_I_SEESAW_WITH_EMPIRICAL_YUKAWA_AMPLITUDE_SEAL"
	StatusNeutrinoMassMatrixNotDerived     = "FAILED_ROUTE_FINITE_NEUTRINO_MASS_MATRIX_DERIVATION"
	StatusFiniteIntermediateStillNotProven = "FINITE_INTERMEDIATE_DYNAMICS_STILL_NOT_DERIVED"
)

const (
	// Inherited sealed hierarchy from Gate 227/230. Gate 231 reads these from
	// Gate 230 when available, but keeps constants here for explicit reference.
	sealedIntermediateScaleGeV = 6.650726476871e11
	sealedMStarGeV             = 1.72179441e17

	// Electroweak VEV used by the previous precision gates. It is empirical
	// boundary data, not a finite-core derivation.
	electroweakVEVGeV = 246.22

	// Neutrino comparison windows used only as phenomenological sanity tests.
	atmosphericScaleEV  = 0.05
	cosmologySumBoundEV = 0.12
	lowerPlausibleEV    = 0.01
	upperPlausibleEV    = 0.10
)

type Gate230Snapshot struct {
	Gate230Inherited                  bool
	HopfResonanceInherited            bool
	FiniteInstantonDerived            bool
	HopfActionMapDerived              bool
	HiddenOrderParameterDerived       bool
	IntermediateSealPreviouslyGranted bool
	IntermediateSealRequired          bool
	MIntGeV                           float64
	MStarGeV                          float64
	BGap                              float64
	HopfScaleGeV                      float64
	TruthStatement                    string
}

type IntermediateBreakingSeal struct {
	Name                         string
	AxiomID                      string
	Active                       bool
	PhenomenologicalBoundary     bool
	FiniteDerived                bool
	Assumption                   string
	ScaleGeV                     float64
	Source                       string
	RequiredBecauseGate230Failed bool
	ReopensPatiSalam             bool
	ReopensLeptoquarkDynamics    bool
	GrantsHiddenOrderParameter   bool
	Verdict                      string
}

type SeesawInput struct {
	Formula                     string
	VEVGeV                      float64
	RightHandedScaleGeV         float64
	OrderOneDiracYukawa         float64
	VEVIsEmpiricalSeal          bool
	MajoranaScaleFromSeal       bool
	DiracYukawaMatrixDerived    bool
	MajoranaMatrixDerived       bool
	MixingAnglesDerived         bool
	UsesObservedNeutrinoMassFit bool
}

type SeesawComputation struct {
	OrderOneMassEV              float64
	OrderOneMassMeV             float64
	OrderOneInPlausibleWindow   bool
	OrderOneAboveCosmologyBound bool
	RatioToAtmosphericScale     float64
	RatioToCosmologySumBound    float64
	YukawaForAtmosphericScale   float64
	YukawaForLowerWindow        float64
	YukawaForUpperWindow        float64
	DiracMassForAtmosphericGeV  float64
	Verdict                     string
}

type BoundCheck struct {
	AtmosphericScaleEV                 float64
	CosmologySumBoundEV                float64
	PlausibleLowerEV                   float64
	PlausibleUpperEV                   float64
	OrderOnePassesOscillationScale     bool
	OrderOnePassesCosmologySumBound    bool
	SmallYukawaCanEnterPlausibleWindow bool
	RequiresEmpiricalYukawaAmplitude   bool
	ExactMixingAnglesDerived           bool
	MassOrderingDerived                bool
	SumMassesComputed                  bool
	Verdict                            string
}

type NeutrinoMatrixAudit struct {
	RightHandedNeutrinoFieldDerived bool
	MajoranaMassOperatorDerived     bool
	DiracYukawaTextureDerived       bool
	FlavorMixingMatrixDerived       bool
	ThreeGenerationRankDerived      bool
	LightNeutrinoMatrixDerived      bool
	OnlyScalePreflightAvailable     bool
	Verdict                         string
}

type FirewallAudit struct {
	UsesOnlySealedIntermediateScale bool
	ActivatesIntermediateSeal       bool
	ClaimsFiniteInstanton           bool
	ClaimsFiniteOrderParameter      bool
	ClaimsFiniteMajoranaMass        bool
	ClaimsFiniteDiracYukawa         bool
	ClaimsNeutrinoMixingAngles      bool
	TunesYukawaToObservedMass       bool
	ReopensPatiSalam                bool
	ReopensLeptoquarkDynamics       bool
	UsesObservedNeutrinoMassAsInput bool
	FiniteCorePolluted              bool
	Verdict                         string
}

type Summary struct {
	IntermediateSealActive           bool
	OrderOneSeesawSupported          bool
	SmallYukawaSeesawConditionallyOK bool
	NeutrinoMatrixDerived            bool
	Status                           string
	NextGate                         string
	Comment                          string
}

type Analysis struct {
	Gate230  Gate230Snapshot
	Seal     IntermediateBreakingSeal
	Input    SeesawInput
	Compute  SeesawComputation
	Bounds   BoundCheck
	Matrix   NeutrinoMatrixAudit
	Firewall FirewallAudit
	Summary  Summary

	TruthStatement string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		g230, err := finitehopfaction.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 230 input: %w", err)
			return
		}
		defaultA, defaultErr = Build(g230)
	})
	return defaultA, defaultErr
}

func Build(g230 finitehopfaction.Analysis) (Analysis, error) {
	s230 := snapshotFromGate230(g230)
	if !s230.Gate230Inherited || !s230.HopfResonanceInherited || s230.MIntGeV <= 0 {
		return Analysis{}, fmt.Errorf("Gate 231 requires Gate 230 inherited Hopf resonance and positive intermediate scale")
	}
	seal := activateSeal(s230)
	input := buildInput(seal)
	comp := computeSeesaw(input)
	bounds := auditBounds(comp)
	matrix := auditMatrix()
	firewall := auditFirewall(seal, input, comp, matrix)
	summary := summarize(seal, comp, bounds, matrix)
	truth := buildTruth(seal, input, comp, bounds, matrix)
	return Analysis{Gate230: s230, Seal: seal, Input: input, Compute: comp, Bounds: bounds, Matrix: matrix, Firewall: firewall, Summary: summary, TruthStatement: truth}, nil
}

func snapshotFromGate230(a finitehopfaction.Analysis) Gate230Snapshot {
	mint := a.HopfAction.TargetMIntGeV
	if mint <= 0 {
		mint = sealedIntermediateScaleGeV
	}
	return Gate230Snapshot{
		Gate230Inherited:                  a.Summary.Status != "" && a.Gate229.MIntTargetGeV > 0,
		HopfResonanceInherited:            a.Summary.GeometricResonanceInherited,
		FiniteInstantonDerived:            a.Summary.FiniteInstantonDerived,
		HopfActionMapDerived:              a.Summary.HopfActionMapDerived,
		HiddenOrderParameterDerived:       a.Summary.HiddenOrderParameterDerived,
		IntermediateSealPreviouslyGranted: a.Summary.IntermediateSealGranted,
		IntermediateSealRequired:          !a.Summary.IntermediateSealGranted && a.Summary.GeometricResonanceInherited,
		MIntGeV:                           mint,
		MStarGeV:                          a.Gate229.MStarGeV,
		BGap:                              a.Gate229.BGap,
		HopfScaleGeV:                      a.HopfAction.HopfMIntGeV,
		TruthStatement:                    a.TruthStatement,
	}
}

func activateSeal(g Gate230Snapshot) IntermediateBreakingSeal {
	return IntermediateBreakingSeal{
		Name:                         "IntermediateBreakingSeal",
		AxiomID:                      "SEAL-INTERMEDIATE-BREAKING-GATE231",
		Active:                       true,
		PhenomenologicalBoundary:     true,
		FiniteDerived:                false,
		Assumption:                   "A hidden/Hopf-sector order parameter exists at M_int and may set right-handed neutrino Majorana thresholds.",
		ScaleGeV:                     g.MIntGeV,
		Source:                       "Gate 227/229/230 sealed Hopf-geometric intermediate scale",
		RequiredBecauseGate230Failed: !g.FiniteInstantonDerived && !g.HopfActionMapDerived && !g.HiddenOrderParameterDerived,
		ReopensPatiSalam:             false,
		ReopensLeptoquarkDynamics:    false,
		GrantsHiddenOrderParameter:   false,
		Verdict:                      StatusIntermediateSealActivated,
	}
}

func buildInput(seal IntermediateBreakingSeal) SeesawInput {
	return SeesawInput{
		Formula:                     "m_nu ≈ y_nu^2 v^2 / M_R",
		VEVGeV:                      electroweakVEVGeV,
		RightHandedScaleGeV:         seal.ScaleGeV,
		OrderOneDiracYukawa:         1.0,
		VEVIsEmpiricalSeal:          true,
		MajoranaScaleFromSeal:       seal.Active,
		DiracYukawaMatrixDerived:    false,
		MajoranaMatrixDerived:       false,
		MixingAnglesDerived:         false,
		UsesObservedNeutrinoMassFit: false,
	}
}

func computeSeesaw(in SeesawInput) SeesawComputation {
	// m[GeV] = y^2 v^2 / M_R; convert GeV to eV with 1 GeV = 1e9 eV.
	mEV := in.OrderOneDiracYukawa * in.OrderOneDiracYukawa * in.VEVGeV * in.VEVGeV / in.RightHandedScaleGeV * 1e9
	yAtm := math.Sqrt(atmosphericScaleEV / mEV)
	yLow := math.Sqrt(lowerPlausibleEV / mEV)
	yHigh := math.Sqrt(upperPlausibleEV / mEV)
	verdict := StatusOrderOneSeesawFailed
	if mEV > lowerPlausibleEV && mEV < upperPlausibleEV {
		verdict = "CONDITIONAL_SUPPORT_ORDER_ONE_TYPE_I_SEESAW_RESONANCE"
	}
	return SeesawComputation{
		OrderOneMassEV:              mEV,
		OrderOneMassMeV:             mEV * 1e-6,
		OrderOneInPlausibleWindow:   mEV > lowerPlausibleEV && mEV < upperPlausibleEV,
		OrderOneAboveCosmologyBound: mEV > cosmologySumBoundEV,
		RatioToAtmosphericScale:     mEV / atmosphericScaleEV,
		RatioToCosmologySumBound:    mEV / cosmologySumBoundEV,
		YukawaForAtmosphericScale:   yAtm,
		YukawaForLowerWindow:        yLow,
		YukawaForUpperWindow:        yHigh,
		DiracMassForAtmosphericGeV:  yAtm * in.VEVGeV,
		Verdict:                     verdict,
	}
}

func auditBounds(c SeesawComputation) BoundCheck {
	return BoundCheck{
		AtmosphericScaleEV:                 atmosphericScaleEV,
		CosmologySumBoundEV:                cosmologySumBoundEV,
		PlausibleLowerEV:                   lowerPlausibleEV,
		PlausibleUpperEV:                   upperPlausibleEV,
		OrderOnePassesOscillationScale:     c.OrderOneInPlausibleWindow,
		OrderOnePassesCosmologySumBound:    !c.OrderOneAboveCosmologyBound,
		SmallYukawaCanEnterPlausibleWindow: c.YukawaForLowerWindow > 0 && c.YukawaForUpperWindow > c.YukawaForLowerWindow,
		RequiresEmpiricalYukawaAmplitude:   !c.OrderOneInPlausibleWindow,
		ExactMixingAnglesDerived:           false,
		MassOrderingDerived:                false,
		SumMassesComputed:                  false,
		Verdict:                            StatusSmallYukawaConditionalSupport,
	}
}

func auditMatrix() NeutrinoMatrixAudit {
	return NeutrinoMatrixAudit{
		RightHandedNeutrinoFieldDerived: false,
		MajoranaMassOperatorDerived:     false,
		DiracYukawaTextureDerived:       false,
		FlavorMixingMatrixDerived:       false,
		ThreeGenerationRankDerived:      false,
		LightNeutrinoMatrixDerived:      false,
		OnlyScalePreflightAvailable:     true,
		Verdict:                         StatusNeutrinoMassMatrixNotDerived,
	}
}

func auditFirewall(seal IntermediateBreakingSeal, in SeesawInput, c SeesawComputation, m NeutrinoMatrixAudit) FirewallAudit {
	return FirewallAudit{
		UsesOnlySealedIntermediateScale: true,
		ActivatesIntermediateSeal:       seal.Active && seal.PhenomenologicalBoundary && !seal.FiniteDerived,
		ClaimsFiniteInstanton:           false,
		ClaimsFiniteOrderParameter:      false,
		ClaimsFiniteMajoranaMass:        m.MajoranaMassOperatorDerived,
		ClaimsFiniteDiracYukawa:         in.DiracYukawaMatrixDerived,
		ClaimsNeutrinoMixingAngles:      in.MixingAnglesDerived,
		TunesYukawaToObservedMass:       false,
		ReopensPatiSalam:                false,
		ReopensLeptoquarkDynamics:       false,
		UsesObservedNeutrinoMassAsInput: false,
		FiniteCorePolluted:              false,
		Verdict:                         "FIREWALLS_PRESERVED_SEESAW_SCALE_PREFLIGHT_ONLY",
	}
}

func summarize(seal IntermediateBreakingSeal, c SeesawComputation, b BoundCheck, m NeutrinoMatrixAudit) Summary {
	statusParts := []string{StatusIntermediateSealActivated, StatusFiniteIntermediateStillNotProven}
	if c.OrderOneInPlausibleWindow && !c.OrderOneAboveCosmologyBound {
		statusParts = append(statusParts, "CONDITIONAL_SUPPORT_ORDER_ONE_TYPE_I_SEESAW_RESONANCE")
	} else {
		statusParts = append(statusParts, StatusOrderOneSeesawFailed)
		if b.SmallYukawaCanEnterPlausibleWindow {
			statusParts = append(statusParts, StatusSmallYukawaConditionalSupport)
		}
	}
	if !m.LightNeutrinoMatrixDerived {
		statusParts = append(statusParts, StatusNeutrinoMassMatrixNotDerived)
	}
	comment := "Gate 231 activates the IntermediateBreakingSeal as a phenomenological boundary, but the order-one Type-I seesaw at M_int gives an active-neutrino mass far above the allowed scale. A viable scale needs a small sealed Dirac neutrino Yukawa, not a finite derivation."
	return Summary{
		IntermediateSealActive:           seal.Active,
		OrderOneSeesawSupported:          c.OrderOneInPlausibleWindow && !c.OrderOneAboveCosmologyBound,
		SmallYukawaSeesawConditionallyOK: b.SmallYukawaCanEnterPlausibleWindow,
		NeutrinoMatrixDerived:            m.LightNeutrinoMatrixDerived,
		Status:                           strings.Join(statusParts, ";"),
		NextGate:                         "derive or seal a neutrino Yukawa/Majorana flavor texture; do not claim exact neutrino masses from M_int alone",
		Comment:                          comment,
	}
}

func buildTruth(seal IntermediateBreakingSeal, in SeesawInput, c SeesawComputation, b BoundCheck, m NeutrinoMatrixAudit) string {
	return fmt.Sprintf("Gate 231 activates %s at M_int=%.12e GeV only as a phenomenological boundary. The scale preflight gives m_nu(y=1)=%.12g eV, which is %.4g times the atmospheric 0.05 eV scale and %.4g times the 0.12 eV cosmology-sum stress bound. Therefore the order-one seesaw resonance fails. A viable 0.05 eV scale would require y_nu≈%.8g (Dirac mass %.8g GeV), which is allowed only behind the empirical Yukawa-amplitude firewall. No RH neutrino field, Majorana matrix, Dirac Yukawa texture, rank, ordering, or PMNS angles are derived.",
		seal.Name, seal.ScaleGeV, c.OrderOneMassEV, c.RatioToAtmosphericScale, c.RatioToCosmologySumBound, c.YukawaForAtmosphericScale, c.DiracMassForAtmosphericGeV)
}

func FormatGate230(g Gate230Snapshot) string {
	return fmt.Sprintf("inherited=%t hopfResonance=%t finiteInstanton=%t hopfActionMap=%t orderParameter=%t priorSealGranted=%t sealRequired=%t M_int=%.12e M_Hopf=%.12e M_*=%.12e B_gap=%.10f",
		g.Gate230Inherited, g.HopfResonanceInherited, g.FiniteInstantonDerived, g.HopfActionMapDerived, g.HiddenOrderParameterDerived, g.IntermediateSealPreviouslyGranted, g.IntermediateSealRequired, g.MIntGeV, g.HopfScaleGeV, g.MStarGeV, g.BGap)
}

func FormatSeal(s IntermediateBreakingSeal) string {
	return fmt.Sprintf("name=%s axiom=%s active=%t phenomenological=%t finiteDerived=%t scale=%.12e requiredByGate230=%t reopensPS=%t reopensLQ=%t grantsOrderParameter=%t verdict=%s assumption=%q",
		s.Name, s.AxiomID, s.Active, s.PhenomenologicalBoundary, s.FiniteDerived, s.ScaleGeV, s.RequiredBecauseGate230Failed, s.ReopensPatiSalam, s.ReopensLeptoquarkDynamics, s.GrantsHiddenOrderParameter, s.Verdict, s.Assumption)
}

func FormatInput(i SeesawInput) string {
	return fmt.Sprintf("formula=%q v=%.5f GeV M_R=%.12e GeV y_order1=%.3f VEVSeal=%t MajoranaScaleFromSeal=%t DiracYukawaMatrixDerived=%t MajoranaMatrixDerived=%t MixingAnglesDerived=%t observedNeutrinoFitInput=%t",
		i.Formula, i.VEVGeV, i.RightHandedScaleGeV, i.OrderOneDiracYukawa, i.VEVIsEmpiricalSeal, i.MajoranaScaleFromSeal, i.DiracYukawaMatrixDerived, i.MajoranaMatrixDerived, i.MixingAnglesDerived, i.UsesObservedNeutrinoMassFit)
}

func FormatComputation(c SeesawComputation) string {
	return fmt.Sprintf("m_nu(y=1)=%.12g eV plausible=%t aboveCosmology=%t ratioAtm=%.8g ratioCosmo=%.8g y_atm=%.10g y_window=[%.10g,%.10g] mD_atm=%.10g GeV verdict=%s",
		c.OrderOneMassEV, c.OrderOneInPlausibleWindow, c.OrderOneAboveCosmologyBound, c.RatioToAtmosphericScale, c.RatioToCosmologySumBound, c.YukawaForAtmosphericScale, c.YukawaForLowerWindow, c.YukawaForUpperWindow, c.DiracMassForAtmosphericGeV, c.Verdict)
}

func FormatBounds(b BoundCheck) string {
	return fmt.Sprintf("atm=%.6g eV cosmologySumStress=%.6g eV plausibleWindow=[%.6g,%.6g] orderOneOsc=%t orderOneCosmo=%t smallYukawaCanFit=%t requiresEmpiricalYukawa=%t exactMixing=%t ordering=%t sumMasses=%t verdict=%s",
		b.AtmosphericScaleEV, b.CosmologySumBoundEV, b.PlausibleLowerEV, b.PlausibleUpperEV, b.OrderOnePassesOscillationScale, b.OrderOnePassesCosmologySumBound, b.SmallYukawaCanEnterPlausibleWindow, b.RequiresEmpiricalYukawaAmplitude, b.ExactMixingAnglesDerived, b.MassOrderingDerived, b.SumMassesComputed, b.Verdict)
}

func FormatMatrix(m NeutrinoMatrixAudit) string {
	return fmt.Sprintf("RHfield=%t MajoranaOp=%t DiracYukawaTexture=%t PMNS=%t rank3=%t lightMatrix=%t scaleOnly=%t verdict=%s",
		m.RightHandedNeutrinoFieldDerived, m.MajoranaMassOperatorDerived, m.DiracYukawaTextureDerived, m.FlavorMixingMatrixDerived, m.ThreeGenerationRankDerived, m.LightNeutrinoMatrixDerived, m.OnlyScalePreflightAvailable, m.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("sealedMIntOnly=%t sealActivated=%t finiteInstanton=%t finiteOrderParameter=%t finiteMajorana=%t finiteYukawa=%t mixingAngles=%t tunedYukawa=%t reopenPS=%t reopenLQ=%t observedNuInput=%t polluted=%t verdict=%s",
		f.UsesOnlySealedIntermediateScale, f.ActivatesIntermediateSeal, f.ClaimsFiniteInstanton, f.ClaimsFiniteOrderParameter, f.ClaimsFiniteMajoranaMass, f.ClaimsFiniteDiracYukawa, f.ClaimsNeutrinoMixingAngles, f.TunesYukawaToObservedMass, f.ReopensPatiSalam, f.ReopensLeptoquarkDynamics, f.UsesObservedNeutrinoMassAsInput, f.FiniteCorePolluted, f.Verdict)
}
