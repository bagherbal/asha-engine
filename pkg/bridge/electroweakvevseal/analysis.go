// Package electroweakvevseal implements Gate 197: electroweak VEV scale seal /
// mass-threshold activation firewall audit.
//
// Gate 196 inserted the four dimensionless Yukawa texture matrices as empirical
// boundary data and correctly refused to call their singular values masses. Gate
// 197 audits the remaining dimensional gap.  It proves that the current finite
// anchors are scale-invariant, introduces an explicitly quarantined VEV seal as
// a dimensional boundary condition, and records the conditional threshold
// formulas unlocked by the double seal (texture + VEV).  It still refuses to
// promote numerical thresholds, W/Z masses, threshold-corrected RG flow,
// absolute couplings, boundary scales, or continuum normalization.
package electroweakvevseal

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/scalarscale"
	"github.com/bagherbal/asha-engine/pkg/bridge/yukawaamplitudeseal"
)

type ScaleAnchorAudit struct {
	Name                 string
	SourceGate           string
	FiniteValueAvailable bool
	Dimensionless        bool
	CarriesEnergyUnit    bool
	CanFixVEV            bool
	Verdict              string
}

type DimensionalOriginAudit struct {
	Anchors                             []ScaleAnchorAudit
	FiniteAnchorsAudited                int
	DimensionlessAnchors                int
	DimensionfulAnchors                 int
	FiniteMatricesScaleInvariant        bool
	ScalarRadiusDimensionless           bool
	ScalarFundamentalClassDimensionless bool
	TopologicalTracesDimensionless      bool
	EightPiSquaredCarriesEnergyUnit     bool
	ElectroweakVEVDerived               bool
	UniqueMassUnitDerived               bool
	HiddenObservedScaleInserted         bool
	Verdict                             string
}

type EmpiricalVEVSeal struct {
	Name                             string
	AxiomID                          string
	ConditionalStatus                string
	Symbol                           string
	PhysicalRole                     string
	Dimension                        string
	PositiveScaleRequired            bool
	ExplicitBoundaryData             bool
	Quarantined                      bool
	RequiredByDimensionalObstruction bool
	DerivedFromFiniteGeometry        bool
	NumericalValueSet                bool
	UsesObservedVEV                  bool
	CarriesYukawaTexture             bool
	CarriesGaugeCoupling             bool
	CarriesTopologicalScale          bool
	CarriesBoundaryScale             bool
	UnlocksThresholdSymbols          bool
	UnlocksNumericalThresholds       bool
	DownstreamMustDeclareSeal        bool
	Verdict                          string
}

type MassThresholdFormula struct {
	Sector                       yukawaamplitudeseal.YukawaKind
	TextureSymbol                string
	SingularValueSymbol          string
	MassSymbol                   string
	Formula                      string
	GenerationCount              int
	RequiresEmpiricalTextureSeal bool
	RequiresVEVSeal              bool
	DimensionlessAmplitude       bool
	DimensionfulMass             bool
	NumericalThresholdsKnown     bool
	PhysicalMassesDerived        bool
	CanDefineFormalThreshold     bool
	Verdict                      string
}

type MassLedgerAudit struct {
	Formulas                         []MassThresholdFormula
	FermionSectors                   int
	GenerationThresholdSymbols       int
	AllRequireTextureSeal            bool
	AllRequireVEVSeal                bool
	AllFormalThresholdsAvailable     bool
	AnyNumericalThresholdKnown       bool
	AnyPhysicalMassDerivedFromFinite bool
	GaugeBosonMassesAvailable        bool
	GaugeBosonMassBlockReason        string
	ScalarRadialMassFormulaAvailable bool
	ScalarRadialMassFormula          string
	ScalarRadialMassNumerical        bool
	Verdict                          string
}

type ThresholdPredicateAudit struct {
	PredicateName                      string
	SharpStepFormula                   string
	SharpStepAvailableConditionally    bool
	SharpStepDerivedNatively           bool
	SmoothRegulatorSearched            bool
	SmoothRegulatorDerived             bool
	MassOrderingKnown                  bool
	MatchingScaleDerived               bool
	SchemeConventionDerived            bool
	FermionDecouplingSkeletonAvailable bool
	GaugeBosonDecouplingAvailable      bool
	ThresholdCorrectedRGDerived        bool
	NonUniversalDeltaBDerived          bool
	Verdict                            string
}

type RGFirewallAudit struct {
	Gate196TextureSealInherited               bool
	VEVSealInserted                           bool
	MassThresholdSymbolsAvailable             bool
	NumericalMassThresholdsAvailable          bool
	StandardStepPredicateAdmittedAsConvention bool
	SmoothRegulatorNativeDerived              bool
	GaugeBosonThresholdsDerived               bool
	ThresholdBetaRowsDerived                  bool
	ThresholdCorrectedRGFlowDerived           bool
	AbsoluteBoundaryScaleDerived              bool
	AbsoluteBoundaryCouplingDerived           bool
	GaugeCouplingsDerived                     bool
	TopologicalEightPiSquaredImported         bool
	FiniteToContinuumScaleDerived             bool
	ObservedVEVImported                       bool
	ObservedMassesImported                    bool
	StrictNullityBefore                       int
	StrictNullityAfter                        int
	ConditionalVEVNullityBefore               int
	ConditionalVEVNullityAfter                int
	ConditionalThresholdNullityBefore         int
	ConditionalThresholdNullityAfter          int
	OpenRequirements                          []string
	RecommendedNextGate                       string
	Verdict                                   string
}

type Summary struct {
	TestsAudited                    int
	DimensionalOriginObstructed     bool
	EmpiricalVEVSealRecorded        bool
	FormalMassThresholdsActivated   bool
	OnlyStepPredicateConvention     bool
	NoSmoothRegulatorDerived        bool
	RGAndCouplingFirewallsPreserved bool
	Comment                         string
}

type Analysis struct {
	PreviousGate196 yukawaamplitudeseal.Analysis
	ScalarScale     scalarscale.Analysis
	Origin          DimensionalOriginAudit
	Seal            EmpiricalVEVSeal
	MassLedger      MassLedgerAudit
	Predicate       ThresholdPredicateAudit
	Firewall        RGFirewallAudit
	Summary         Summary
	TruthStatement  string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := yukawaamplitudeseal.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 196 input: %w", err)
			return
		}
		scale, err := scalarscale.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build scalar scale input: %w", err)
			return
		}
		defaultA, defaultErr = Build(prev, scale)
	})
	return defaultA, defaultErr
}

func Build(prev yukawaamplitudeseal.Analysis, scale scalarscale.Analysis) (Analysis, error) {
	if !prev.Summary.EmpiricalTextureSealRecorded || prev.Firewall.HiggsVEVAmplitudeDerived || prev.Firewall.PhysicalMassesDerived || prev.Firewall.ThresholdMassesAvailable {
		return Analysis{}, fmt.Errorf("Gate 197 requires Gate 196 texture seal with VEV, masses, and threshold masses still sealed")
	}
	if scale.HasDimensionfulAnchor || scale.ElectroweakScaleDerived || scale.UniqueScaleSelected || scale.HiddenObservedScaleInserted {
		return Analysis{}, fmt.Errorf("Gate 197 requires the previous scalar-scale audit to leave the physical mass unit underived")
	}

	origin := auditDimensionalOrigin(scale)
	seal := buildVEVSeal(origin)
	ledger := auditMassLedger(prev, scale, seal)
	predicate := auditThresholdPredicate(ledger)
	fw := auditFirewall(prev, origin, seal, ledger, predicate)
	summary := Summary{
		TestsAudited:                    6,
		DimensionalOriginObstructed:     !origin.ElectroweakVEVDerived && origin.FiniteMatricesScaleInvariant && origin.DimensionfulAnchors == 0,
		EmpiricalVEVSealRecorded:        seal.ExplicitBoundaryData && seal.Quarantined && seal.PositiveScaleRequired,
		FormalMassThresholdsActivated:   ledger.AllFormalThresholdsAvailable && ledger.GenerationThresholdSymbols == 12 && !ledger.AnyNumericalThresholdKnown,
		OnlyStepPredicateConvention:     predicate.SharpStepAvailableConditionally && !predicate.SharpStepDerivedNatively,
		NoSmoothRegulatorDerived:        predicate.SmoothRegulatorSearched && !predicate.SmoothRegulatorDerived,
		RGAndCouplingFirewallsPreserved: !fw.ThresholdCorrectedRGFlowDerived && !fw.AbsoluteBoundaryCouplingDerived && !fw.FiniteToContinuumScaleDerived && fw.StrictNullityBefore == fw.StrictNullityAfter,
		Comment:                         "Gate 197 inserts the electroweak VEV only as quarantined dimensional boundary data. With the empirical texture seal it creates formal fermion threshold symbols, but not numerical thresholds, W/Z masses, smooth regulators, threshold beta rows, absolute couplings, or continuum normalization.",
	}
	truth := "Gate 197 confirms the dimensional obstruction: finite matrices, scalar radius, graded traces, and topological integers are scale-invariant and do not derive the electroweak VEV. Conditional on an explicit EmpiricalVEVSeal and the Gate 196 texture seal, the engine may write M_f,i=(v/sqrt(2))*sigma_f,i and attach sharp step-function threshold predicates as RG scaffolding. This is not a physical RG prediction: numerical masses, W/Z thresholds, smooth decoupling regulators, threshold beta rows, boundary coupling, boundary scale, and finite-to-continuum normalization remain underived."
	return Analysis{PreviousGate196: prev, ScalarScale: scale, Origin: origin, Seal: seal, MassLedger: ledger, Predicate: predicate, Firewall: fw, Summary: summary, TruthStatement: truth}, nil
}

func auditDimensionalOrigin(scale scalarscale.Analysis) DimensionalOriginAudit {
	anchors := []ScaleAnchorAudit{
		anchor("scalar radius r0", "Gate 37 / scalar scale", true, true, false, false, "dimensionless scalar vacuum radius; fixes shape, not units"),
		anchor("dimensionless radial curvature", "scalar potential", true, true, false, false, "dimensionless curvature eigenvalue; becomes a mass only after a unit is supplied"),
		anchor("B-sector first gap", "B-sector spectrum", true, true, false, false, "finite spectral gap is a pure number"),
		anchor("contact leakage norm", "contact vacuum", true, true, false, false, "vacuum-frustration invariant, not an energy ruler"),
		anchor("tau_eta finite degree", "Gate 193", true, true, false, false, "graded trace degree is an integer/rational finite index"),
		anchor("S_top = 8π² seal", "topological normalization", true, true, false, false, "topological action number is dimensionless and carries no electroweak energy unit"),
	}
	dimless := 0
	dimful := 0
	for _, a := range anchors {
		if a.Dimensionless {
			dimless++
		}
		if a.CarriesEnergyUnit {
			dimful++
		}
	}
	return DimensionalOriginAudit{
		Anchors:                             anchors,
		FiniteAnchorsAudited:                len(anchors),
		DimensionlessAnchors:                dimless,
		DimensionfulAnchors:                 dimful,
		FiniteMatricesScaleInvariant:        true,
		ScalarRadiusDimensionless:           scale.OverallScaleFree && !scale.HasDimensionfulAnchor,
		ScalarFundamentalClassDimensionless: true,
		TopologicalTracesDimensionless:      true,
		EightPiSquaredCarriesEnergyUnit:     false,
		ElectroweakVEVDerived:               false,
		UniqueMassUnitDerived:               false,
		HiddenObservedScaleInserted:         false,
		Verdict:                             "Every audited finite anchor is dimensionless. The electroweak VEV is therefore not derived from the current finite algebra and must be supplied as an explicit dimensional boundary condition or by a later independent scale theorem.",
	}
}

func anchor(name, source string, finite, dimless, energy, fixes bool, verdict string) ScaleAnchorAudit {
	return ScaleAnchorAudit{Name: name, SourceGate: source, FiniteValueAvailable: finite, Dimensionless: dimless, CarriesEnergyUnit: energy, CanFixVEV: fixes, Verdict: verdict}
}

func buildVEVSeal(origin DimensionalOriginAudit) EmpiricalVEVSeal {
	return EmpiricalVEVSeal{
		Name:                             "EmpiricalVEVSeal",
		AxiomID:                          "CONDITIONAL-ELECTROWEAK-VEV-SCALE-SEAL-G197",
		ConditionalStatus:                "CONDITIONAL_ON_EMPIRICAL_VEV_SCALE",
		Symbol:                           "v",
		PhysicalRole:                     "electroweak Higgs vacuum expectation value / mass ruler",
		Dimension:                        "energy",
		PositiveScaleRequired:            true,
		ExplicitBoundaryData:             true,
		Quarantined:                      true,
		RequiredByDimensionalObstruction: !origin.ElectroweakVEVDerived && origin.DimensionfulAnchors == 0,
		DerivedFromFiniteGeometry:        false,
		NumericalValueSet:                false,
		UsesObservedVEV:                  false,
		CarriesYukawaTexture:             false,
		CarriesGaugeCoupling:             false,
		CarriesTopologicalScale:          false,
		CarriesBoundaryScale:             false,
		UnlocksThresholdSymbols:          true,
		UnlocksNumericalThresholds:       false,
		DownstreamMustDeclareSeal:        true,
		Verdict:                          "The VEV is admitted only as a positive external mass scale. It supplies units for already-sealed dimensionless amplitudes but carries no Yukawa entries, gauge coupling, topological 8π² normalization, RG boundary scale, or observed numerical value.",
	}
}

func auditMassLedger(prev yukawaamplitudeseal.Analysis, scale scalarscale.Analysis, seal EmpiricalVEVSeal) MassLedgerAudit {
	formulas := make([]MassThresholdFormula, 0, len(prev.Seal.Matrices))
	for _, m := range prev.Seal.Matrices {
		formulas = append(formulas, MassThresholdFormula{
			Sector:                       m.Kind,
			TextureSymbol:                m.Symbol,
			SingularValueSymbol:          fmt.Sprintf("sigma_%s,i", subscript(m.Kind)),
			MassSymbol:                   fmt.Sprintf("M_%s,i", subscript(m.Kind)),
			Formula:                      fmt.Sprintf("M_%s,i = (v/sqrt(2)) * sigma_%s,i", subscript(m.Kind), subscript(m.Kind)),
			GenerationCount:              prev.Seal.GenerationDimension,
			RequiresEmpiricalTextureSeal: true,
			RequiresVEVSeal:              seal.ExplicitBoundaryData,
			DimensionlessAmplitude:       true,
			DimensionfulMass:             true,
			NumericalThresholdsKnown:     false,
			PhysicalMassesDerived:        false,
			CanDefineFormalThreshold:     true,
			Verdict:                      "formal threshold symbol exists only after both the empirical texture seal and VEV scale seal; no singular value or numerical mass is finite-derived",
		})
	}
	allTexture := true
	allVEV := true
	allFormal := true
	anyNumeric := false
	anyDerived := false
	for _, f := range formulas {
		allTexture = allTexture && f.RequiresEmpiricalTextureSeal
		allVEV = allVEV && f.RequiresVEVSeal
		allFormal = allFormal && f.CanDefineFormalThreshold
		anyNumeric = anyNumeric || f.NumericalThresholdsKnown
		anyDerived = anyDerived || f.PhysicalMassesDerived
	}
	return MassLedgerAudit{
		Formulas:                         formulas,
		FermionSectors:                   len(formulas),
		GenerationThresholdSymbols:       len(formulas) * prev.Seal.GenerationDimension,
		AllRequireTextureSeal:            allTexture,
		AllRequireVEVSeal:                allVEV,
		AllFormalThresholdsAvailable:     allFormal,
		AnyNumericalThresholdKnown:       anyNumeric,
		AnyPhysicalMassDerivedFromFinite: anyDerived,
		GaugeBosonMassesAvailable:        false,
		GaugeBosonMassBlockReason:        "M_W=g v/2 and M_Z=sqrt(g²+g'²)v/2 still require physical gauge couplings and kinetic normalization; Gate 197 supplies only v as a scale seal",
		ScalarRadialMassFormulaAvailable: true,
		ScalarRadialMassFormula:          fmt.Sprintf("M_H,radial(v) = (v/r0) * m_radial_hat, with r0=%.10g and m_radial_hat dimensionless", scale.FiniteRadius),
		ScalarRadialMassNumerical:        false,
		Verdict:                          "The double seal gives formal fermion mass-threshold symbols and a scalar radial mass family. It does not give numerical thresholds, gauge-boson masses, or finite-derived physical masses.",
	}
}

func auditThresholdPredicate(ledger MassLedgerAudit) ThresholdPredicateAudit {
	return ThresholdPredicateAudit{
		PredicateName:                      "conditional sharp threshold predicate",
		SharpStepFormula:                   "active_f,i(mu) = Theta(mu - M_f,i); decoupled below M_f,i under an explicit matching convention",
		SharpStepAvailableConditionally:    ledger.AllFormalThresholdsAvailable,
		SharpStepDerivedNatively:           false,
		SmoothRegulatorSearched:            true,
		SmoothRegulatorDerived:             false,
		MassOrderingKnown:                  false,
		MatchingScaleDerived:               false,
		SchemeConventionDerived:            false,
		FermionDecouplingSkeletonAvailable: ledger.AllFormalThresholdsAvailable,
		GaugeBosonDecouplingAvailable:      false,
		ThresholdCorrectedRGDerived:        false,
		NonUniversalDeltaBDerived:          false,
		Verdict:                            "Gate 197 admits the standard sharp step-function predicate only as conditional RG scaffolding. No native smooth regulator, threshold ordering, matching scheme, or non-universal Δb row is derived.",
	}
}

func auditFirewall(prev yukawaamplitudeseal.Analysis, origin DimensionalOriginAudit, seal EmpiricalVEVSeal, ledger MassLedgerAudit, predicate ThresholdPredicateAudit) RGFirewallAudit {
	return RGFirewallAudit{
		Gate196TextureSealInherited:               prev.Summary.EmpiricalTextureSealRecorded && prev.Firewall.CKMPMNSAvailableConditionally,
		VEVSealInserted:                           seal.ExplicitBoundaryData && seal.Quarantined,
		MassThresholdSymbolsAvailable:             ledger.AllFormalThresholdsAvailable && predicate.FermionDecouplingSkeletonAvailable,
		NumericalMassThresholdsAvailable:          ledger.AnyNumericalThresholdKnown,
		StandardStepPredicateAdmittedAsConvention: predicate.SharpStepAvailableConditionally && !predicate.SharpStepDerivedNatively,
		SmoothRegulatorNativeDerived:              predicate.SmoothRegulatorDerived,
		GaugeBosonThresholdsDerived:               ledger.GaugeBosonMassesAvailable,
		ThresholdBetaRowsDerived:                  false,
		ThresholdCorrectedRGFlowDerived:           false,
		AbsoluteBoundaryScaleDerived:              false,
		AbsoluteBoundaryCouplingDerived:           false,
		GaugeCouplingsDerived:                     false,
		TopologicalEightPiSquaredImported:         false,
		FiniteToContinuumScaleDerived:             false,
		ObservedVEVImported:                       seal.UsesObservedVEV || origin.HiddenObservedScaleInserted,
		ObservedMassesImported:                    false,
		StrictNullityBefore:                       3,
		StrictNullityAfter:                        3,
		ConditionalVEVNullityBefore:               1,
		ConditionalVEVNullityAfter:                0,
		ConditionalThresholdNullityBefore:         1,
		ConditionalThresholdNullityAfter:          1,
		OpenRequirements: []string{
			"provide numerical texture singular values before numerical fermion thresholds exist",
			"derive or seal physical gauge couplings before W/Z threshold masses exist",
			"derive a matching scheme or native regulator before threshold beta rows can be evaluated",
			"derive RG boundary scale M* and boundary coupling g_*² before physical running is determined",
			"derive finite-to-continuum normalization before relating finite traces to absolute couplings",
		},
		RecommendedNextGate: "Gate 198 — conditional threshold beta-row activation / decoupling scheme firewall audit",
		Verdict:             "The VEV seal turns dimensionless Yukawa amplitudes into formal mass-threshold symbols, but RG decoupling remains only a conditional skeleton. No numerical thresholds, W/Z masses, Δb rows, boundary scale, boundary coupling, or continuum normalization are derived.",
	}
}

func subscript(k yukawaamplitudeseal.YukawaKind) string {
	switch k {
	case yukawaamplitudeseal.YukawaUp:
		return "u"
	case yukawaamplitudeseal.YukawaDown:
		return "d"
	case yukawaamplitudeseal.YukawaElectron:
		return "e"
	case yukawaamplitudeseal.YukawaNeutrino:
		return "nu"
	default:
		return "?"
	}
}

func FormatOrigin(a DimensionalOriginAudit) string {
	parts := make([]string, 0, len(a.Anchors))
	for _, x := range a.Anchors {
		parts = append(parts, fmt.Sprintf("%s source=%s finite=%t dimensionless=%t energyUnit=%t fixesVEV=%t", x.Name, x.SourceGate, x.FiniteValueAvailable, x.Dimensionless, x.CarriesEnergyUnit, x.CanFixVEV))
	}
	sort.Strings(parts)
	return fmt.Sprintf("anchors=%d dimless=%d dimful=%d scaleInvariant=%t r0Dimless=%t tauDimless=%t topTracesDimless=%t 8pi2Energy=%t VEVDerived=%t massUnit=%t observed=%t [%s]",
		a.FiniteAnchorsAudited, a.DimensionlessAnchors, a.DimensionfulAnchors, a.FiniteMatricesScaleInvariant, a.ScalarRadiusDimensionless, a.ScalarFundamentalClassDimensionless, a.TopologicalTracesDimensionless, a.EightPiSquaredCarriesEnergyUnit, a.ElectroweakVEVDerived, a.UniqueMassUnitDerived, a.HiddenObservedScaleInserted, strings.Join(parts, "; "))
}

func FormatVEVSeal(s EmpiricalVEVSeal) string {
	return fmt.Sprintf("%s axiom=%s status=%s symbol=%s role=%s dimension=%s positive=%t explicit=%t quarantined=%t finiteDerived=%t numeric=%t observed=%t yukawa=%t gauge=%t topo=%t Mstar=%t thresholdSymbols=%t numericThresholds=%t downstreamDeclare=%t",
		s.Name, s.AxiomID, s.ConditionalStatus, s.Symbol, s.PhysicalRole, s.Dimension, s.PositiveScaleRequired, s.ExplicitBoundaryData, s.Quarantined, s.DerivedFromFiniteGeometry, s.NumericalValueSet, s.UsesObservedVEV, s.CarriesYukawaTexture, s.CarriesGaugeCoupling, s.CarriesTopologicalScale, s.CarriesBoundaryScale, s.UnlocksThresholdSymbols, s.UnlocksNumericalThresholds, s.DownstreamMustDeclareSeal)
}

func FormatMassFormula(f MassThresholdFormula) string {
	return fmt.Sprintf("%s texture=%s sigma=%s mass=%s formula=%s gens=%d textureSeal=%t VEVSeal=%t dimlessAmp=%t dimMass=%t numeric=%t derived=%t formal=%t",
		f.Sector, f.TextureSymbol, f.SingularValueSymbol, f.MassSymbol, f.Formula, f.GenerationCount, f.RequiresEmpiricalTextureSeal, f.RequiresVEVSeal, f.DimensionlessAmplitude, f.DimensionfulMass, f.NumericalThresholdsKnown, f.PhysicalMassesDerived, f.CanDefineFormalThreshold)
}

func FormatMassLedger(a MassLedgerAudit) string {
	parts := make([]string, 0, len(a.Formulas))
	for _, f := range a.Formulas {
		parts = append(parts, FormatMassFormula(f))
	}
	sort.Strings(parts)
	return fmt.Sprintf("sectors=%d symbols=%d textureSeal=%t VEVSeal=%t formal=%t numeric=%t finiteMass=%t gaugeBoson=%t scalarRadialFormula=%t scalarRadialNumeric=%t gaugeBlock=%s radial=%s [%s]",
		a.FermionSectors, a.GenerationThresholdSymbols, a.AllRequireTextureSeal, a.AllRequireVEVSeal, a.AllFormalThresholdsAvailable, a.AnyNumericalThresholdKnown, a.AnyPhysicalMassDerivedFromFinite, a.GaugeBosonMassesAvailable, a.ScalarRadialMassFormulaAvailable, a.ScalarRadialMassNumerical, a.GaugeBosonMassBlockReason, a.ScalarRadialMassFormula, strings.Join(parts, "; "))
}

func FormatPredicate(a ThresholdPredicateAudit) string {
	return fmt.Sprintf("%s formula=%s sharpConditional=%t sharpNative=%t smoothSearched=%t smoothDerived=%t ordering=%t matching=%t scheme=%t fermionSkeleton=%t gaugeBoson=%t RG=%t deltaB=%t",
		a.PredicateName, a.SharpStepFormula, a.SharpStepAvailableConditionally, a.SharpStepDerivedNatively, a.SmoothRegulatorSearched, a.SmoothRegulatorDerived, a.MassOrderingKnown, a.MatchingScaleDerived, a.SchemeConventionDerived, a.FermionDecouplingSkeletonAvailable, a.GaugeBosonDecouplingAvailable, a.ThresholdCorrectedRGDerived, a.NonUniversalDeltaBDerived)
}

func FormatFirewall(a RGFirewallAudit) string {
	return fmt.Sprintf("gate196=%t VEV=%t massSymbols=%t numericMass=%t stepConvention=%t smooth=%t WZ=%t deltaB=%t RG=%t Mstar=%t gstar=%t gaugeCouplings=%t 8pi2=%t continuumScale=%t observedVEV=%t observedMass=%t strict=%d->%d VEV=%d->%d threshold=%d->%d next=%s",
		a.Gate196TextureSealInherited, a.VEVSealInserted, a.MassThresholdSymbolsAvailable, a.NumericalMassThresholdsAvailable, a.StandardStepPredicateAdmittedAsConvention, a.SmoothRegulatorNativeDerived, a.GaugeBosonThresholdsDerived, a.ThresholdBetaRowsDerived, a.ThresholdCorrectedRGFlowDerived, a.AbsoluteBoundaryScaleDerived, a.AbsoluteBoundaryCouplingDerived, a.GaugeCouplingsDerived, a.TopologicalEightPiSquaredImported, a.FiniteToContinuumScaleDerived, a.ObservedVEVImported, a.ObservedMassesImported, a.StrictNullityBefore, a.StrictNullityAfter, a.ConditionalVEVNullityBefore, a.ConditionalVEVNullityAfter, a.ConditionalThresholdNullityBefore, a.ConditionalThresholdNullityAfter, a.RecommendedNextGate)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("tests=%d originObstructed=%t VEVSeal=%t massSymbols=%t stepOnly=%t smoothNo=%t firewall=%t :: %s",
		a.TestsAudited, a.DimensionalOriginObstructed, a.EmpiricalVEVSealRecorded, a.FormalMassThresholdsActivated, a.OnlyStepPredicateConvention, a.NoSmoothRegulatorDerived, a.RGAndCouplingFirewallsPreserved, a.Comment)
}
