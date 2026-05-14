// Package scalarheatkernelsubtraction implements Gate 305:
// Scalar Heat-Kernel Subtraction / Higgs Potential Channel Separation Audit.
//
// Gate 304 promoted the a4-channel cutoff coefficient f_0 to the sealed
// contact-spectral value 7, thereby stabilizing kinetic/gauge normalization at
// the coefficient-source level. Gate 305 moves to the a2 channel. It does not
// compute the Higgs mass. Instead it formalizes the exact subtraction sieve
// needed before the scalar quadratic coefficient can be interpreted physically:
// subtract the field-independent vacuum/background a2 carrier from the raw
// fluctuated a2 coefficient, isolate the scalar-power-2 dynamical remainder,
// and only then divide by the Gate-300/301/302 Z_H normalization. The result is
// still conditional on f_2, the cutoff scale Lambda, subtraction convention,
// numerical Yukawa/amplitude data, and the sign convention that identifies the
// canonical Lorentzian Higgs-potential mass parameter.
package scalarheatkernelsubtraction

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE305-SCALAR-HEAT-KERNEL-SUBTRACTION-HIGGS-POTENTIAL-CHANNEL-SEPARATION-AUDIT"

	StatusGate304Inherited                          = "CONDITIONAL_SUPPORT_GATE304_CONTACT_F0_SEAL_INHERITED"
	StatusRawA2DecompositionFormalized              = "CONDITIONAL_SUPPORT_RAW_A2_DECOMPOSITION_FORMALIZED"
	StatusSubtractionSchemeFormalized               = "CONDITIONAL_SUPPORT_SCALAR_HEAT_KERNEL_SUBTRACTION_SCHEME_FORMALIZED"
	StatusHiggsMassExtractionMapFormalized          = "CONDITIONAL_SUPPORT_HIGGS_MASS_PARAMETER_EXTRACTION_MAP_FORMALIZED"
	StatusF2MomentDependencySieveFormalized         = "CONDITIONAL_SUPPORT_F2_MOMENT_DEPENDENCY_SIEVE_FORMALIZED"
	StatusHiggsPotentialChannelSeparationFormalized = "CONDITIONAL_SUPPORT_HIGGS_POTENTIAL_CHANNEL_SEPARATION_FORMALIZED"
	StatusFirewallsPreserved                        = "CONDITIONAL_SUPPORT_GATE305_DYNAMICAL_FIREWALLS_PRESERVED"

	StatusFailedF2MomentNotLocked                      = "FAILED_ROUTE_F2_MOMENT_NOT_LOCKED"
	StatusFailedCutoffScaleLambdaNotDerived            = "FAILED_ROUTE_CUTOFF_SCALE_LAMBDA_NOT_DERIVED"
	StatusFailedNumericalZHStillSealed                 = "FAILED_ROUTE_NUMERICAL_ZH_VALUE_STILL_SEALED"
	StatusFailedNumericalYukawaAmplitudesStillSealed   = "FAILED_ROUTE_NUMERICAL_YUKAWA_AMPLITUDES_STILL_SEALED"
	StatusFailedHiggsMassNumericalPredictionNotDerived = "FAILED_ROUTE_HIGGS_MASS_NUMERICAL_PREDICTION_NOT_DERIVED"
	StatusFailedHiggsQuarticStillSealed                = "FAILED_ROUTE_HIGGS_QUARTIC_STILL_SEALED"
	StatusFailedSubtractionSchemeNotUnique             = "FAILED_ROUTE_SUBTRACTION_RENORMALIZATION_SCHEME_NOT_UNIQUE"
	StatusFailedVacuumCountertermSelectionNotDerived   = "FAILED_ROUTE_VACUUM_COUNTERTERM_PHYSICAL_SELECTION_NOT_DERIVED"
	StatusFailedHigherMomentsStillOpen                 = "FAILED_ROUTE_CUTOFF_PROFILE_HIGHER_MOMENTS_STILL_OPEN"
	StatusFailedBGapInstantonStillSealed               = "FAILED_ROUTE_BGAP_INSTANTON_ACTION_STILL_SEALED"
)

const (
	sealedF0Value = 7
)

type Gate304Inheritance struct {
	ContactF0Promoted            bool
	PromotedF0Value              int
	F0Positive                   bool
	KineticNormalizationAnchored bool
	HigherMomentsLocked          bool
	UniqueProfileShapeDerived    bool
	HeatKernelSubtractionClaimed bool
	NumericalZHComputed          bool
	HiggsMassPredictionClaimed   bool
	NumericalYukawasInserted     bool
	Verdict                      string
}

type A2Component struct {
	Name              string
	Symbol            string
	FieldPower        int
	CutoffDependent   bool
	VacuumArtifact    bool
	DynamicalScalar   bool
	Subtracted        bool
	PhysicalCandidate bool
	Status            string
}

type RawA2Decomposition struct {
	RawExpression              string
	VacuumReference            string
	DynamicalRemainder         string
	Components                 []A2Component
	FieldIndependentVacuumSeen bool
	ScalarPower2ChannelSeen    bool
	MixedTermsFirewalled       bool
	NumericalCoefficientUsed   bool
	DecompositionFormalized    bool
	Verdict                    string
}

type SubtractionScheme struct {
	Name                       string
	Formula                    string
	ReferenceBackground        string
	SubtractedVacuumPieces     []string
	RetainedDynamicalPieces    []string
	LinearityRequired          bool
	GaugeCovariant             bool
	BackgroundIndependent      bool
	SchemeUnique               bool
	CountertermPhysicallyFixed bool
	NumericalCountertermUsed   bool
	Formalized                 bool
	Verdict                    string
}

type HiggsMassExtractionMap struct {
	RawQuadraticChannel      string
	SubtractedChannel        string
	ZHInput                  string
	CanonicalRescaling       string
	MassMap                  string
	SignConvention           string
	UsesGate300Normalization bool
	UsesSubtractedA2         bool
	RequiresPositiveZH       bool
	RequiresF2               bool
	RequiresCutoffScale      bool
	RequiresYukawaAmplitudes bool
	NumericalMassComputed    bool
	MapFormalized            bool
	Verdict                  string
}

type F2MomentDependencySieve struct {
	F0Status                    string
	F2Role                      string
	F2LockedByGate304           bool
	SameProfileCouldVaryF2      bool
	RequiresProfileShapeRule    bool
	RequiresCutoffScaleLambda   bool
	CanClaimMassWithoutF2       bool
	PredictivePowerLostIfFreeF2 bool
	DependencyFormalized        bool
	Verdict                     string
}

type ChannelSeparationLedger struct {
	QuadraticChannelIsolated bool
	QuarticChannelTouched    bool
	GaugeKineticDisturbed    bool
	VacuumChannelSubtracted  bool
	A4F0SealPreserved        bool
	A2F2StillOpen            bool
	NoDynamicsOverclaimed    bool
	Verdict                  string
}

type RemainingObligation struct {
	Name, WhyRequired, Status string
	BlocksPrediction          bool
}

type FirewallAudit struct {
	NoNumericalF2Inserted            bool
	NoCutoffScaleInserted            bool
	NoYukawaNumbersInserted          bool
	NoNumericalZHComputed            bool
	NoHiggsMassPredictionClaimed     bool
	NoHiggsQuarticPredictionClaimed  bool
	NoBGapInstantonClaimed           bool
	NoUniqueSubtractionSchemeClaimed bool
	NoProfileHigherMomentLockClaimed bool
	F0SealPreservedOnlyForA4         bool
	FiniteCorePolluted               bool
	Obligations                      []RemainingObligation
	Verdict                          string
}

type Summary struct {
	Gate304Inherited               bool
	RawA2Decomposed                bool
	SubtractionSchemeFormalized    bool
	QuadraticChannelSeparated      bool
	MassMapFormalized              bool
	F2DependencyFormalized         bool
	F0SealPreserved                bool
	NumericalHiggsMassDerived      bool
	PhysicalDynamicsDerived        bool
	FirewallPreserved              bool
	Status, DirectAnswer, NextGate string
}

type Analysis struct {
	Input       Gate304Inheritance
	A2          RawA2Decomposition
	Subtraction SubtractionScheme
	MassMap     HiggsMassExtractionMap
	F2          F2MomentDependencySieve
	Channels    ChannelSeparationLedger
	Firewalls   FirewallAudit
	Summary     Summary
	Truth       string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	i := inheritGate304()
	a2 := decomposeRawA2(i)
	sub := formalizeSubtraction(a2)
	mass := formalizeMassMap(a2, sub)
	f2 := auditF2Dependency(i, mass)
	channels := separateChannels(i, a2, sub, f2)
	fw := auditFirewalls(i, sub, mass, f2, channels)
	sum := buildSummary(i, a2, sub, mass, f2, channels, fw)
	truth := "Gate 305 formalizes the scalar heat-kernel subtraction algorithm needed before the a2 coefficient can be read as a Higgs mass channel. The gate decomposes raw a2(D_A) into a field-independent vacuum/reference carrier and a scalar-power-2 dynamical remainder, defines Delta a2 := a2(D_A)-a2(D_vac), and maps the normalized mass parameter schematically as mu_H^2 proportional to f_2 Lambda^2 Delta a2_scalar / Z_H with the relevant sign convention. It does not compute a Higgs mass, because f_2, Lambda, numerical Yukawa amplitudes, absolute Z_H, and the physical renormalization/counterterm prescription remain sealed."
	return Analysis{Input: i, A2: a2, Subtraction: sub, MassMap: mass, F2: f2, Channels: channels, Firewalls: fw, Summary: sum, Truth: truth}, nil
}

func inheritGate304() Gate304Inheritance {
	return Gate304Inheritance{
		ContactF0Promoted:            true,
		PromotedF0Value:              sealedF0Value,
		F0Positive:                   sealedF0Value > 0,
		KineticNormalizationAnchored: true,
		HigherMomentsLocked:          false,
		UniqueProfileShapeDerived:    false,
		HeatKernelSubtractionClaimed: false,
		NumericalZHComputed:          false,
		HiggsMassPredictionClaimed:   false,
		NumericalYukawasInserted:     false,
		Verdict:                      strings.Join([]string{StatusGate304Inherited, StatusFailedHigherMomentsStillOpen}, ";"),
	}
}

func decomposeRawA2(i Gate304Inheritance) RawA2Decomposition {
	components := []A2Component{
		{Name: "vacuum/reference finite Dirac trace", Symbol: "a2(D_vac) or a2(D_F)|_{H=0}", FieldPower: 0, CutoffDependent: true, VacuumArtifact: true, DynamicalScalar: false, Subtracted: true, PhysicalCandidate: false, Status: "VACUUM_REFERENCE_SUBTRACTED"},
		{Name: "scalar quadratic dynamical channel", Symbol: "a2_scalar^{(2)}(H)", FieldPower: 2, CutoffDependent: true, VacuumArtifact: false, DynamicalScalar: true, Subtracted: false, PhysicalCandidate: true, Status: "RETAIN_AS_DELTA_A2_SCALAR"},
		{Name: "mixed/background convention residue", Symbol: "a2_mixed_or_counterterm", FieldPower: -1, CutoffDependent: true, VacuumArtifact: true, DynamicalScalar: false, Subtracted: true, PhysicalCandidate: false, Status: StatusFailedVacuumCountertermSelectionNotDerived},
	}
	return RawA2Decomposition{
		RawExpression:              "a2(D_A) = a2(D_vac) + a2_scalar^{(2)}(H_raw) + a2_mixed_or_counterterm",
		VacuumReference:            "D_vac := D_A evaluated on the same finite spectral data with scalar fluctuation H_raw set to zero and no physical scalar displacement",
		DynamicalRemainder:         "Delta a2_scalar := scalar-power-2 projection of [a2(D_A)-a2(D_vac)]",
		Components:                 components,
		FieldIndependentVacuumSeen: true,
		ScalarPower2ChannelSeen:    true,
		MixedTermsFirewalled:       true,
		NumericalCoefficientUsed:   false,
		DecompositionFormalized:    i.ContactF0Promoted && !i.HeatKernelSubtractionClaimed,
		Verdict:                    strings.Join([]string{StatusRawA2DecompositionFormalized, StatusFailedVacuumCountertermSelectionNotDerived}, ";"),
	}
}

func formalizeSubtraction(a2 RawA2Decomposition) SubtractionScheme {
	formalized := a2.DecompositionFormalized && a2.FieldIndependentVacuumSeen && a2.ScalarPower2ChannelSeen
	return SubtractionScheme{
		Name:                       "VacuumReferencedScalarA2Subtraction",
		Formula:                    "Delta a2[H] := Pi_{scalar^2}(a2(D_A[H]) - a2(D_A[0]))",
		ReferenceBackground:        "same finite spectral triple, same cutoff convention, same trace normalization, scalar fluctuation set to zero",
		SubtractedVacuumPieces:     []string{"field-independent a2(D_vac)", "cosmological/reference trace contribution", "scheme-dependent scalar-independent counterterm residue"},
		RetainedDynamicalPieces:    []string{"scalar-power-2 Higgs channel", "terms transforming as |H_raw|^2 under the already-audited gauge/Higgs representation"},
		LinearityRequired:          true,
		GaugeCovariant:             true,
		BackgroundIndependent:      false,
		SchemeUnique:               false,
		CountertermPhysicallyFixed: false,
		NumericalCountertermUsed:   false,
		Formalized:                 formalized,
		Verdict:                    strings.Join([]string{StatusSubtractionSchemeFormalized, StatusFailedSubtractionSchemeNotUnique, StatusFailedVacuumCountertermSelectionNotDerived}, ";"),
	}
}

func formalizeMassMap(a2 RawA2Decomposition, sub SubtractionScheme) HiggsMassExtractionMap {
	formalized := a2.DecompositionFormalized && sub.Formalized && sub.LinearityRequired && sub.GaugeCovariant
	return HiggsMassExtractionMap{
		RawQuadraticChannel:      "raw a2 scalar-power-2 coefficient before vacuum subtraction",
		SubtractedChannel:        "Delta a2_scalar := Pi_{scalar^2}(a2(D_A[H])-a2(D_A[0]))",
		ZHInput:                  "Z_H from Gate 300-302 normalization ledger, with Gate 304 f0=7 only for the a4 kinetic source",
		CanonicalRescaling:       "H_raw = H_phys / sqrt(Z_H)",
		MassMap:                  "mu_H^2 = Sign_L · N_2 · f_2 · Lambda^2 · Delta a2_scalar / Z_H",
		SignConvention:           "Sign_L is the Euclidean-to-Lorentzian potential convention selecting V(H) = -mu_H^2 |H_phys|^2 + lambda |H_phys|^4 or the equivalent declared canonical form",
		UsesGate300Normalization: true,
		UsesSubtractedA2:         true,
		RequiresPositiveZH:       true,
		RequiresF2:               true,
		RequiresCutoffScale:      true,
		RequiresYukawaAmplitudes: true,
		NumericalMassComputed:    false,
		MapFormalized:            formalized,
		Verdict:                  strings.Join([]string{StatusHiggsMassExtractionMapFormalized, StatusFailedF2MomentNotLocked, StatusFailedCutoffScaleLambdaNotDerived, StatusFailedNumericalZHStillSealed, StatusFailedNumericalYukawaAmplitudesStillSealed, StatusFailedHiggsMassNumericalPredictionNotDerived}, ";"),
	}
}

func auditF2Dependency(i Gate304Inheritance, mass HiggsMassExtractionMap) F2MomentDependencySieve {
	return F2MomentDependencySieve{
		F0Status:                    "Gate 304 seals f0=7 for the a4 kinetic/gauge normalization channel only",
		F2Role:                      "f2 multiplies the Lambda^2 a2 channel and therefore the scalar quadratic/Higgs mass map",
		F2LockedByGate304:           i.HigherMomentsLocked,
		SameProfileCouldVaryF2:      true,
		RequiresProfileShapeRule:    true,
		RequiresCutoffScaleLambda:   true,
		CanClaimMassWithoutF2:       false,
		PredictivePowerLostIfFreeF2: true,
		DependencyFormalized:        mass.MapFormalized && mass.RequiresF2,
		Verdict:                     strings.Join([]string{StatusF2MomentDependencySieveFormalized, StatusFailedF2MomentNotLocked, StatusFailedHigherMomentsStillOpen, StatusFailedCutoffScaleLambdaNotDerived}, ";"),
	}
}

func separateChannels(i Gate304Inheritance, a2 RawA2Decomposition, sub SubtractionScheme, f2 F2MomentDependencySieve) ChannelSeparationLedger {
	return ChannelSeparationLedger{
		QuadraticChannelIsolated: a2.ScalarPower2ChannelSeen && sub.Formalized,
		QuarticChannelTouched:    false,
		GaugeKineticDisturbed:    false,
		VacuumChannelSubtracted:  a2.FieldIndependentVacuumSeen && len(sub.SubtractedVacuumPieces) > 0,
		A4F0SealPreserved:        i.ContactF0Promoted && i.PromotedF0Value == sealedF0Value,
		A2F2StillOpen:            !f2.F2LockedByGate304,
		NoDynamicsOverclaimed:    !i.HiggsMassPredictionClaimed && !f2.CanClaimMassWithoutF2,
		Verdict:                  strings.Join([]string{StatusHiggsPotentialChannelSeparationFormalized, StatusFailedHiggsQuarticStillSealed, StatusFailedF2MomentNotLocked}, ";"),
	}
}

func auditFirewalls(i Gate304Inheritance, sub SubtractionScheme, mass HiggsMassExtractionMap, f2 F2MomentDependencySieve, channels ChannelSeparationLedger) FirewallAudit {
	obs := []RemainingObligation{
		{"f2 cutoff moment", "a2 is multiplied by f2 Lambda^2; Gate 304 only promoted f0", StatusFailedF2MomentNotLocked, true},
		{"cutoff scale Lambda", "absolute mass dimension requires the physical/unification cutoff scale", StatusFailedCutoffScaleLambdaNotDerived, true},
		{"absolute Z_H", "mass extraction divides by Z_H; Gate 301-304 prove positivity/source class but not a numerical value", StatusFailedNumericalZHStillSealed, true},
		{"Yukawa/amplitude ledger", "Delta a2_scalar depends on finite Dirac/Yukawa amplitudes that remain sealed", StatusFailedNumericalYukawaAmplitudesStillSealed, true},
		{"subtraction/counterterm selection", "vacuum subtraction is formalized but not uniquely selected by a physical renormalization prescription", StatusFailedSubtractionSchemeNotUnique, true},
		{"vacuum counterterm finite part", "physical finite part must be fixed by an internal rule or a declared renormalization condition", StatusFailedVacuumCountertermSelectionNotDerived, true},
		{"quartic potential channel", "lambda_H extraction belongs to the a4 scalar-power-4 channel and is not evaluated in this a2 gate", StatusFailedHiggsQuarticStillSealed, true},
		{"B-gap instanton action", "scalar a2 subtraction does not derive S_inst=(4/pi)/B_gap", StatusFailedBGapInstantonStillSealed, true},
	}
	polluted := i.NumericalYukawasInserted || i.NumericalZHComputed || i.HiggsMassPredictionClaimed || mass.NumericalMassComputed || sub.NumericalCountertermUsed || sub.SchemeUnique || sub.CountertermPhysicallyFixed || f2.F2LockedByGate304 || f2.CanClaimMassWithoutF2 || channels.QuarticChannelTouched || channels.GaugeKineticDisturbed
	return FirewallAudit{
		NoNumericalF2Inserted:            !f2.F2LockedByGate304,
		NoCutoffScaleInserted:            mass.RequiresCutoffScale,
		NoYukawaNumbersInserted:          !i.NumericalYukawasInserted,
		NoNumericalZHComputed:            !i.NumericalZHComputed,
		NoHiggsMassPredictionClaimed:     !i.HiggsMassPredictionClaimed && !mass.NumericalMassComputed,
		NoHiggsQuarticPredictionClaimed:  !channels.QuarticChannelTouched,
		NoBGapInstantonClaimed:           true,
		NoUniqueSubtractionSchemeClaimed: !sub.SchemeUnique && !sub.CountertermPhysicallyFixed,
		NoProfileHigherMomentLockClaimed: !i.HigherMomentsLocked && !f2.F2LockedByGate304,
		F0SealPreservedOnlyForA4:         channels.A4F0SealPreserved && channels.A2F2StillOpen,
		FiniteCorePolluted:               polluted,
		Obligations:                      obs,
		Verdict:                          strings.Join([]string{StatusFirewallsPreserved, StatusFailedF2MomentNotLocked, StatusFailedCutoffScaleLambdaNotDerived, StatusFailedNumericalZHStillSealed, StatusFailedNumericalYukawaAmplitudesStillSealed, StatusFailedHiggsMassNumericalPredictionNotDerived, StatusFailedSubtractionSchemeNotUnique}, ";"),
	}
}

func buildSummary(i Gate304Inheritance, a2 RawA2Decomposition, sub SubtractionScheme, mass HiggsMassExtractionMap, f2 F2MomentDependencySieve, channels ChannelSeparationLedger, fw FirewallAudit) Summary {
	return Summary{
		Gate304Inherited:            i.ContactF0Promoted && i.F0Positive && i.PromotedF0Value == sealedF0Value,
		RawA2Decomposed:             a2.DecompositionFormalized,
		SubtractionSchemeFormalized: sub.Formalized,
		QuadraticChannelSeparated:   channels.QuadraticChannelIsolated && channels.VacuumChannelSubtracted,
		MassMapFormalized:           mass.MapFormalized && mass.UsesGate300Normalization && mass.UsesSubtractedA2,
		F2DependencyFormalized:      f2.DependencyFormalized && !f2.F2LockedByGate304,
		F0SealPreserved:             channels.A4F0SealPreserved,
		NumericalHiggsMassDerived:   mass.NumericalMassComputed,
		PhysicalDynamicsDerived:     mass.NumericalMassComputed || channels.QuarticChannelTouched,
		FirewallPreserved:           !fw.FiniteCorePolluted && fw.NoNumericalF2Inserted && fw.NoCutoffScaleInserted && fw.NoYukawaNumbersInserted && fw.NoNumericalZHComputed && fw.NoHiggsMassPredictionClaimed && fw.NoHiggsQuarticPredictionClaimed && fw.NoUniqueSubtractionSchemeClaimed && fw.F0SealPreservedOnlyForA4,
		Status:                      strings.Join([]string{StatusSubtractionSchemeFormalized, StatusHiggsMassExtractionMapFormalized, StatusF2MomentDependencySieveFormalized, StatusFirewallsPreserved}, ";"),
		DirectAnswer:                "Gate 305 formalizes the vacuum-referenced a2 subtraction Delta a2 := Pi_scalar^2(a2(D_A[H])-a2(D_A[0])) and the normalized Higgs mass map mu_H^2 proportional to f2 Lambda^2 Delta a2_scalar / Z_H. It separates the physical candidate channel but does not compute the Higgs mass.",
		NextGate:                    "Gate 306 should audit the f2/higher-moment source: either derive a canonical profile shape or define a sealed higher-moment rule compatible with the Gate 304 f0=7 promotion, while preserving the Higgs-mass firewall until Lambda, amplitudes, and renormalization conditions are fixed.",
	}
}

func FormatGate304Inheritance(i Gate304Inheritance) string {
	return fmt.Sprintf("f0Promoted=%t f0=%d positive=%t kineticAnchored=%t higherMoments=%t uniqueShape=%t subtractionClaimed=%t ZH=%t Higgs=%t Yukawa=%t verdict=%s", i.ContactF0Promoted, i.PromotedF0Value, i.F0Positive, i.KineticNormalizationAnchored, i.HigherMomentsLocked, i.UniqueProfileShapeDerived, i.HeatKernelSubtractionClaimed, i.NumericalZHComputed, i.HiggsMassPredictionClaimed, i.NumericalYukawasInserted, i.Verdict)
}

func FormatA2Component(c A2Component) string {
	return fmt.Sprintf("%s symbol=%q power=%d cutoff=%t vacuum=%t dynamic=%t subtracted=%t physical=%t status=%s", c.Name, c.Symbol, c.FieldPower, c.CutoffDependent, c.VacuumArtifact, c.DynamicalScalar, c.Subtracted, c.PhysicalCandidate, c.Status)
}

func FormatRawA2(a RawA2Decomposition) string {
	parts := []string{}
	for _, c := range a.Components {
		parts = append(parts, FormatA2Component(c))
	}
	return fmt.Sprintf("raw=%q reference=%q remainder=%q vacuumSeen=%t scalar2=%t mixedFirewalled=%t numeric=%t formalized=%t components=[%s] verdict=%s", a.RawExpression, a.VacuumReference, a.DynamicalRemainder, a.FieldIndependentVacuumSeen, a.ScalarPower2ChannelSeen, a.MixedTermsFirewalled, a.NumericalCoefficientUsed, a.DecompositionFormalized, strings.Join(parts, " | "), a.Verdict)
}

func FormatSubtraction(s SubtractionScheme) string {
	return fmt.Sprintf("name=%s formula=%q reference=%q subtracted=[%s] retained=[%s] linear=%t gaugeCovariant=%t backgroundIndependent=%t unique=%t countertermFixed=%t numericCounterterm=%t formalized=%t verdict=%s", s.Name, s.Formula, s.ReferenceBackground, strings.Join(s.SubtractedVacuumPieces, "/"), strings.Join(s.RetainedDynamicalPieces, "/"), s.LinearityRequired, s.GaugeCovariant, s.BackgroundIndependent, s.SchemeUnique, s.CountertermPhysicallyFixed, s.NumericalCountertermUsed, s.Formalized, s.Verdict)
}

func FormatMassMap(m HiggsMassExtractionMap) string {
	return fmt.Sprintf("raw=%q delta=%q ZH=%q rescale=%q map=%q sign=%q gate300=%t subtracted=%t positiveZH=%t f2=%t Lambda=%t Yukawa=%t numeric=%t formalized=%t verdict=%s", m.RawQuadraticChannel, m.SubtractedChannel, m.ZHInput, m.CanonicalRescaling, m.MassMap, m.SignConvention, m.UsesGate300Normalization, m.UsesSubtractedA2, m.RequiresPositiveZH, m.RequiresF2, m.RequiresCutoffScale, m.RequiresYukawaAmplitudes, m.NumericalMassComputed, m.MapFormalized, m.Verdict)
}

func FormatF2(f F2MomentDependencySieve) string {
	return fmt.Sprintf("f0=%q f2Role=%q f2Locked=%t variableF2=%t profileRule=%t Lambda=%t massWithoutF2=%t predictiveLoss=%t formalized=%t verdict=%s", f.F0Status, f.F2Role, f.F2LockedByGate304, f.SameProfileCouldVaryF2, f.RequiresProfileShapeRule, f.RequiresCutoffScaleLambda, f.CanClaimMassWithoutF2, f.PredictivePowerLostIfFreeF2, f.DependencyFormalized, f.Verdict)
}

func FormatChannels(c ChannelSeparationLedger) string {
	return fmt.Sprintf("quadratic=%t quarticTouched=%t gaugeDisturbed=%t vacuumSubtracted=%t f0Preserved=%t f2Open=%t noOverclaim=%t verdict=%s", c.QuadraticChannelIsolated, c.QuarticChannelTouched, c.GaugeKineticDisturbed, c.VacuumChannelSubtracted, c.A4F0SealPreserved, c.A2F2StillOpen, c.NoDynamicsOverclaimed, c.Verdict)
}

func FormatObligation(o RemainingObligation) string {
	return fmt.Sprintf("%s required=%q status=%s blocks=%t", o.Name, o.WhyRequired, o.Status, o.BlocksPrediction)
}

func FormatFirewalls(f FirewallAudit) string {
	obs := []string{}
	for _, o := range f.Obligations {
		obs = append(obs, FormatObligation(o))
	}
	return fmt.Sprintf("noF2=%t noLambda=%t noYukawa=%t noZH=%t noMass=%t noQuartic=%t noBGap=%t noUniqueSubtraction=%t noHigherMoments=%t f0OnlyA4=%t polluted=%t obligations=[%s] verdict=%s", f.NoNumericalF2Inserted, f.NoCutoffScaleInserted, f.NoYukawaNumbersInserted, f.NoNumericalZHComputed, f.NoHiggsMassPredictionClaimed, f.NoHiggsQuarticPredictionClaimed, f.NoBGapInstantonClaimed, f.NoUniqueSubtractionSchemeClaimed, f.NoProfileHigherMomentLockClaimed, f.F0SealPreservedOnlyForA4, f.FiniteCorePolluted, strings.Join(obs, " | "), f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("gate304=%t rawA2=%t subtraction=%t quadratic=%t massMap=%t f2=%t f0=%t numericMass=%t dynamics=%t firewall=%t status=%s answer=%q next=%q", s.Gate304Inherited, s.RawA2Decomposed, s.SubtractionSchemeFormalized, s.QuadraticChannelSeparated, s.MassMapFormalized, s.F2DependencyFormalized, s.F0SealPreserved, s.NumericalHiggsMassDerived, s.PhysicalDynamicsDerived, s.FirewallPreserved, s.Status, s.DirectAnswer, s.NextGate)
}
