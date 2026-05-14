// Package contactspectralcutoffpromotion implements Gate 304:
// Contact-Spectral Cutoff Promotion / Canonical Positive Test-Profile
// Construction Audit.
//
// Gate 303 proved that the internal contact-spectral candidate
// zeta_contact(0)=7 satisfies the Gate-302 positivity obligation for f_0, but
// refused to equate that discrete invariant with the continuous heat-kernel
// cutoff moment without a bridge theorem. Gate 304 constructs that bridge at
// the coefficient-source level: for any admissible positive a4-channel moment
// functional Lambda_4 and any positive base profile rho with Lambda_4[rho]>0,
// the normalized profile
//
//	f_contact := zeta_contact(0) * rho / Lambda_4[rho]
//
// is continuous/smooth when rho is, non-negative, finite in the a4 channel, and
// satisfies Lambda_4[f_contact]=zeta_contact(0)=7 exactly. This promotes the
// contact-spectral cutoff value as a sealed source for f_0. It does not derive a
// unique profile shape, a variationally preferred cutoff curve, numerical
// Yukawas, Z_H, Higgs mass/quartic values, absolute gauge couplings, or B-gap
// instanton dynamics.
package contactspectralcutoffpromotion

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE304-CONTACT-SPECTRAL-CUTOFF-PROMOTION-CANONICAL-POSITIVE-TEST-PROFILE-CONSTRUCTION-AUDIT"

	StatusGate303Inherited                      = "CONDITIONAL_SUPPORT_GATE303_CONTACT_F0_CANDIDATE_INHERITED"
	StatusContinuousPositiveProfileFormalized   = "CONDITIONAL_SUPPORT_CONTINUOUS_POSITIVE_TEST_PROFILE_CLASS_FORMALIZED"
	StatusCanonicalPositiveProfileConstructed   = "CONDITIONAL_SUPPORT_CANONICAL_POSITIVE_TEST_PROFILE_CONSTRUCTED"
	StatusDiscreteContinuousMomentMapFormalized = "CONDITIONAL_SUPPORT_DISCRETE_TO_CONTINUOUS_MOMENT_MAP_FORMALIZED"
	StatusPromotionSealActivated                = "CONDITIONAL_SUPPORT_CONTACT_SPECTRAL_CUTOFF_PROMOTION_SEAL_ACTIVATED"
	StatusContactSpectralCutoffPromoted         = "CONDITIONAL_SUPPORT_CONTACT_SPECTRAL_CUTOFF_PROMOTED"
	StatusFirewallsPreserved                    = "CONDITIONAL_SUPPORT_GATE304_DYNAMICAL_FIREWALLS_PRESERVED"

	StatusFailedUniqueCutoffProfileShapeNotDerived         = "FAILED_ROUTE_UNIQUE_CUTOFF_PROFILE_SHAPE_NOT_DERIVED"
	StatusFailedCutoffProfileVariationalPrincipleMissing   = "FAILED_ROUTE_CUTOFF_PROFILE_VARIATIONAL_PRINCIPLE_NOT_DERIVED"
	StatusFailedCutoffProfileHigherMomentsNotLocked        = "FAILED_ROUTE_CUTOFF_PROFILE_HIGHER_MOMENTS_NOT_LOCKED"
	StatusFailedZHNumericalValueStillSealed                = "FAILED_ROUTE_ZH_NUMERICAL_VALUE_STILL_SEALED"
	StatusFailedYukawaAmplitudesStillSealed                = "FAILED_ROUTE_NUMERICAL_YUKAWA_AMPLITUDES_STILL_SEALED"
	StatusFailedHiggsMassQuarticStillFirewalled            = "FAILED_ROUTE_HIGGS_MASS_AND_QUARTIC_STILL_FIREWALLED"
	StatusFailedGaugeCouplingsAbsoluteStillFirewalled      = "FAILED_ROUTE_ABSOLUTE_GAUGE_COUPLINGS_STILL_FIREWALLED"
	StatusFailedBGAPInstantonStillSealed                   = "FAILED_ROUTE_BGAP_INSTANTON_ACTION_STILL_SEALED"
	StatusFailedHeatKernelSubtractionSchemeStillOpen       = "FAILED_ROUTE_HEAT_KERNEL_SUBTRACTION_SCHEME_STILL_OPEN"
	StatusFailedPhysicalProfileShapeNotEmpiricallyTestable = "FAILED_ROUTE_PROFILE_SHAPE_PHYSICAL_UNIQUENESS_NOT_TESTED"
)

const (
	contactZeta0Exact = "ζ_contact(0)=7"
	contactZeta0      = 7

	// Under the radial example Lambda_4[f]=∫_0^∞ x^3 f(x) dx and rho=e^{-x^2},
	// Lambda_4[rho]=1/2, so f_contact=14 e^{-x^2}. The audit treats this only
	// as a witness; the theorem itself is the abstract positive-linear-functional
	// normalization rule.
	radialGaussianBaseMomentNumerator   = 1
	radialGaussianBaseMomentDenominator = 2
	radialGaussianScale                 = 14
)

type Gate303Inheritance struct {
	PositiveF0ClassFormalized         bool
	ContactCandidateAvailable         bool
	ContactCandidateExact             string
	ContactCandidateValue             int
	ContactCandidatePositive          bool
	ContactCandidateObservedInput     bool
	FinalSourcePreviouslySelected     bool
	PromotionTheoremPreviouslyDerived bool
	NumericalZHPreviouslyComputed     bool
	PhysicalDynamicsPreviouslyDerived bool
	Verdict                           string
}

type ContinuousProfileFormalization struct {
	ProfileClass                   string
	MomentFunctional               string
	AbstractFunctionalRule         string
	PositivityCondition            string
	RegularityCondition            string
	FiniteMomentCondition          string
	AdmissibleBaseProfileCondition string
	PreservesGate302Sign           bool
	MomentFunctionalPositive       bool
	ObservedInputUsed              bool
	Verdict                        string
}

type CanonicalProfileConstruction struct {
	BaseProfileSymbol            string
	BaseProfileCondition         string
	NormalizationFormula         string
	ContactValue                 int
	ConstructedProfile           string
	ProfileNonNegative           bool
	ProfileContinuous            bool
	ProfileSmooth                bool
	MomentFinite                 bool
	MomentEqualsContactZeta0     bool
	PreservesPositivity          bool
	CanonicalNormalizationRule   bool
	UniqueProfileShapeDerived    bool
	VariationalPreferenceDerived bool
	UsesObservedInput            bool
	RadialGaussianWitness        string
	EvaluationFunctionalWitness  string
	Verdict                      string
}

type DiscreteContinuousMapSieve struct {
	DiscreteInvariant               string
	DiscreteValue                   int
	DiscreteInterpretation          string
	ContinuousTarget                string
	MapFormula                      string
	MapIsAlgebraicallyExact         bool
	MapPreservesSign                bool
	MapUsesPositiveLinearFunctional bool
	MapRequiresNonZeroBaseMoment    bool
	MapIsUniqueAtCoefficientLevel   bool
	MapIsUniqueAtProfileShapeLevel  bool
	ImportsEmpiricalData            bool
	LocksHigherMoments              bool
	Verdict                         string
}

type PromotionSeal struct {
	Name                          string
	ActivatedConditionally        bool
	PromotedF0Exact               string
	PromotedF0Value               int
	F0Positive                    bool
	F0SourceSelectedBySeal        bool
	ContactSpectrumPromoted       bool
	ProfileShapePromoted          bool
	HigherMomentsPromoted         bool
	HeatKernelCoefficientPromoted bool
	NumericalZHComputed           bool
	HiggsPredictionClaimed        bool
	GaugeCouplingAbsoluteClaimed  bool
	BGapInstantonClaimed          bool
	Verdict                       string
}

type RemainingObligation struct {
	Name, WhyRequired, Status string
	BlocksPrediction          bool
}

type FirewallAudit struct {
	NoObservedInputInserted             bool
	NoYukawaNumbersInserted             bool
	NoNumericalZHComputed               bool
	NoHiggsMassQuarticClaimed           bool
	NoAbsoluteGaugeCouplingsClaimed     bool
	NoBGapInstantonClaimed              bool
	NoHeatKernelSubtractionClaimed      bool
	NoUniqueProfileShapeClaimed         bool
	NoHigherMomentLockClaimed           bool
	ContactF0PromotedOnlyAsSealedSource bool
	FiniteCorePolluted                  bool
	Obligations                         []RemainingObligation
	Verdict                             string
}

type Summary struct {
	Gate303Inherited                bool
	ContinuousProfileFormalized     bool
	CanonicalPositiveProfileBuilt   bool
	DiscreteContinuousMapFormalized bool
	PromotionSealActivated          bool
	ContactCutoffPromoted           bool
	F0LockedBySeal                  bool
	UniqueProfileShapeDerived       bool
	HigherMomentsLocked             bool
	NumericalZHComputed             bool
	PhysicalDynamicsDerived         bool
	FirewallPreserved               bool
	Status, DirectAnswer, NextGate  string
}

type Analysis struct {
	Input        Gate303Inheritance
	Profile      ContinuousProfileFormalization
	Construction CanonicalProfileConstruction
	Map          DiscreteContinuousMapSieve
	Seal         PromotionSeal
	Firewalls    FirewallAudit
	Summary      Summary
	Truth        string
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
	i := inheritGate303()
	p := formalizeContinuousProfile()
	c := constructCanonicalPositiveProfile(i, p)
	m := auditDiscreteContinuousMap(i, p, c)
	s := activatePromotionSeal(i, p, c, m)
	fw := auditFirewalls(i, c, m, s)
	sum := buildSummary(i, p, c, m, s, fw)
	truth := "Gate 304 constructs the missing coefficient-source bridge from the discrete contact invariant to a continuous heat-kernel cutoff moment. Given a positive a4-channel moment functional Lambda_4 and a positive admissible base profile rho with Lambda_4[rho]>0, the normalized profile f_contact=ζ_contact(0) rho/Lambda_4[rho] is positive and has Lambda_4[f_contact]=7 exactly. This conditionally activates the ContactSpectralCutoffPromotionSeal and promotes f_0 to the sealed value 7. The gate deliberately does not derive a unique cutoff-curve shape, higher cutoff moments, numerical Z_H, Higgs mass/quartic, absolute gauge couplings, or B-gap instanton dynamics."
	return Analysis{Input: i, Profile: p, Construction: c, Map: m, Seal: s, Firewalls: fw, Summary: sum, Truth: truth}, nil
}

func inheritGate303() Gate303Inheritance {
	return Gate303Inheritance{
		PositiveF0ClassFormalized:         true,
		ContactCandidateAvailable:         true,
		ContactCandidateExact:             contactZeta0Exact,
		ContactCandidateValue:             contactZeta0,
		ContactCandidatePositive:          contactZeta0 > 0,
		ContactCandidateObservedInput:     false,
		FinalSourcePreviouslySelected:     false,
		PromotionTheoremPreviouslyDerived: false,
		NumericalZHPreviouslyComputed:     false,
		PhysicalDynamicsPreviouslyDerived: false,
		Verdict:                           StatusGate303Inherited,
	}
}

func formalizeContinuousProfile() ContinuousProfileFormalization {
	return ContinuousProfileFormalization{
		ProfileClass:                   "admissible spectral-action test profile f: [0,∞)→[0,∞), continuous/smooth enough for the selected heat-kernel convention, compactly supported or rapidly decaying",
		MomentFunctional:               "Λ_4[f], the positive a_4-channel cutoff functional; depending on convention it may be evaluation f(0), a radial weighted moment ∫_0^∞ x^3 f(x) dx, or the already-fixed Gate-302 a_4 weight functional",
		AbstractFunctionalRule:         "Λ_4 is required to be positive linear on the admissible cone: f≥0 and f not Λ_4-null imply Λ_4[f]>0",
		PositivityCondition:            "f(x)≥0 for all x and Λ_4[f]>0",
		RegularityCondition:            "smooth/continuous smoothed-step or Gaussian-type profile; sharp cutoffs are approximation data, not canonical proof objects",
		FiniteMomentCondition:          "Λ_4[f]<∞ and all heat-kernel channels used by the gate are finite",
		AdmissibleBaseProfileCondition: "choose rho≥0, rho continuous/smooth, finite in Λ_4, and Λ_4[rho]>0",
		PreservesGate302Sign:           true,
		MomentFunctionalPositive:       true,
		ObservedInputUsed:              false,
		Verdict:                        StatusContinuousPositiveProfileFormalized,
	}
}

func constructCanonicalPositiveProfile(i Gate303Inheritance, p ContinuousProfileFormalization) CanonicalProfileConstruction {
	momentOK := i.ContactCandidateAvailable && i.ContactCandidatePositive && p.MomentFunctionalPositive
	return CanonicalProfileConstruction{
		BaseProfileSymbol:            "rho",
		BaseProfileCondition:         "rho is any admissible non-negative nonzero base profile with Λ_4[rho]>0",
		NormalizationFormula:         "f_contact(x) := ζ_contact(0) · rho(x) / Λ_4[rho]",
		ContactValue:                 i.ContactCandidateValue,
		ConstructedProfile:           "f_contact = 7·rho/Λ_4[rho]",
		ProfileNonNegative:           momentOK,
		ProfileContinuous:            true,
		ProfileSmooth:                true,
		MomentFinite:                 momentOK,
		MomentEqualsContactZeta0:     momentOK,
		PreservesPositivity:          momentOK,
		CanonicalNormalizationRule:   momentOK,
		UniqueProfileShapeDerived:    false,
		VariationalPreferenceDerived: false,
		UsesObservedInput:            false,
		RadialGaussianWitness:        fmt.Sprintf("if Λ_4[f]=∫_0^∞ x^3 f(x)dx and rho=e^{-x^2}, then Λ_4[rho]=%d/%d and f_contact=%d e^{-x^2}, so Λ_4[f_contact]=7", radialGaussianBaseMomentNumerator, radialGaussianBaseMomentDenominator, radialGaussianScale),
		EvaluationFunctionalWitness:  "if Λ_4[f]=f(0) and rho(0)>0, then f_contact=7·rho/rho(0), so Λ_4[f_contact]=7",
		Verdict:                      strings.Join([]string{StatusCanonicalPositiveProfileConstructed, StatusFailedUniqueCutoffProfileShapeNotDerived, StatusFailedCutoffProfileVariationalPrincipleMissing}, ";"),
	}
}

func auditDiscreteContinuousMap(i Gate303Inheritance, p ContinuousProfileFormalization, c CanonicalProfileConstruction) DiscreteContinuousMapSieve {
	valid := i.ContactCandidateValue == contactZeta0 && c.MomentEqualsContactZeta0 && c.PreservesPositivity && p.MomentFunctionalPositive
	return DiscreteContinuousMapSieve{
		DiscreteInvariant:               i.ContactCandidateExact,
		DiscreteValue:                   i.ContactCandidateValue,
		DiscreteInterpretation:          "finite contact spectral modal count / zeta dimension of the internal contact carrier",
		ContinuousTarget:                "a_4-channel cutoff coefficient f_0 := Λ_4[f_contact]",
		MapFormula:                      "ζ_contact(0) ↦ f_contact=ζ_contact(0)·rho/Λ_4[rho] ↦ Λ_4[f_contact]=ζ_contact(0)=7",
		MapIsAlgebraicallyExact:         valid,
		MapPreservesSign:                valid,
		MapUsesPositiveLinearFunctional: p.MomentFunctionalPositive,
		MapRequiresNonZeroBaseMoment:    true,
		MapIsUniqueAtCoefficientLevel:   valid,
		MapIsUniqueAtProfileShapeLevel:  false,
		ImportsEmpiricalData:            false,
		LocksHigherMoments:              false,
		Verdict:                         strings.Join([]string{StatusDiscreteContinuousMomentMapFormalized, StatusFailedUniqueCutoffProfileShapeNotDerived, StatusFailedCutoffProfileHigherMomentsNotLocked}, ";"),
	}
}

func activatePromotionSeal(i Gate303Inheritance, p ContinuousProfileFormalization, c CanonicalProfileConstruction, m DiscreteContinuousMapSieve) PromotionSeal {
	activated := i.PositiveF0ClassFormalized && i.ContactCandidatePositive && p.PreservesGate302Sign && c.MomentEqualsContactZeta0 && m.MapIsAlgebraicallyExact && m.MapPreservesSign && !m.ImportsEmpiricalData
	return PromotionSeal{
		Name:                          "ContactSpectralCutoffPromotionSeal",
		ActivatedConditionally:        activated,
		PromotedF0Exact:               "f_0 := Λ_4[f_contact] = ζ_contact(0) = 7",
		PromotedF0Value:               contactZeta0,
		F0Positive:                    activated && contactZeta0 > 0,
		F0SourceSelectedBySeal:        activated,
		ContactSpectrumPromoted:       activated,
		ProfileShapePromoted:          false,
		HigherMomentsPromoted:         false,
		HeatKernelCoefficientPromoted: activated,
		NumericalZHComputed:           false,
		HiggsPredictionClaimed:        false,
		GaugeCouplingAbsoluteClaimed:  false,
		BGapInstantonClaimed:          false,
		Verdict:                       strings.Join([]string{StatusPromotionSealActivated, StatusContactSpectralCutoffPromoted, StatusFailedUniqueCutoffProfileShapeNotDerived, StatusFailedCutoffProfileHigherMomentsNotLocked, StatusFailedZHNumericalValueStillSealed}, ";"),
	}
}

func auditFirewalls(i Gate303Inheritance, c CanonicalProfileConstruction, m DiscreteContinuousMapSieve, s PromotionSeal) FirewallAudit {
	obs := []RemainingObligation{
		{"unique cutoff profile shape", "needed before claiming the entire test function is uniquely derived rather than only its a4 coefficient", StatusFailedUniqueCutoffProfileShapeNotDerived, false},
		{"cutoff variational principle", "needed to prefer one base profile rho among the infinitely many positive profiles with the same a4 moment", StatusFailedCutoffProfileVariationalPrincipleMissing, false},
		{"higher cutoff moments", "needed before using the profile to fix f2, f4, cosmological terms, or moment ratios", StatusFailedCutoffProfileHigherMomentsNotLocked, true},
		{"heat-kernel subtraction scheme", "needed before scalar quadratic/quartic channels become physical mass/quartic claims", StatusFailedHeatKernelSubtractionSchemeStillOpen, true},
		{"absolute Z_H value", "requires f0 plus absolute N4 convention and nonzero numerical Yukawa amplitudes", StatusFailedZHNumericalValueStillSealed, true},
		{"Yukawa amplitude ledger", "K_H^raw magnitude depends on sealed Y_u,Y_d,Y_e,Y_ν amplitudes", StatusFailedYukawaAmplitudesStillSealed, true},
		{"Higgs mass/quartic extraction", "requires normalized Z_H, scalar quadratic/quartic traces, subtraction scheme, and amplitudes", StatusFailedHiggsMassQuarticStillFirewalled, true},
		{"absolute gauge couplings", "f0 helps normalize a4 gauge terms but representation traces and absolute coupling scale are not evaluated here", StatusFailedGaugeCouplingsAbsoluteStillFirewalled, true},
		{"B-gap instanton action", "continuous cutoff promotion is heat-kernel bookkeeping and does not derive S_inst=(4/pi)/B_gap", StatusFailedBGAPInstantonStillSealed, true},
	}
	polluted := i.ContactCandidateObservedInput || c.UsesObservedInput || m.ImportsEmpiricalData || s.NumericalZHComputed || s.HiggsPredictionClaimed || s.GaugeCouplingAbsoluteClaimed || s.BGapInstantonClaimed || s.ProfileShapePromoted || s.HigherMomentsPromoted
	return FirewallAudit{
		NoObservedInputInserted:             !i.ContactCandidateObservedInput && !c.UsesObservedInput && !m.ImportsEmpiricalData,
		NoYukawaNumbersInserted:             true,
		NoNumericalZHComputed:               !s.NumericalZHComputed,
		NoHiggsMassQuarticClaimed:           !s.HiggsPredictionClaimed,
		NoAbsoluteGaugeCouplingsClaimed:     !s.GaugeCouplingAbsoluteClaimed,
		NoBGapInstantonClaimed:              !s.BGapInstantonClaimed,
		NoHeatKernelSubtractionClaimed:      true,
		NoUniqueProfileShapeClaimed:         !c.UniqueProfileShapeDerived && !s.ProfileShapePromoted,
		NoHigherMomentLockClaimed:           !m.LocksHigherMoments && !s.HigherMomentsPromoted,
		ContactF0PromotedOnlyAsSealedSource: s.ActivatedConditionally && s.F0SourceSelectedBySeal && s.HeatKernelCoefficientPromoted && !s.ProfileShapePromoted,
		FiniteCorePolluted:                  polluted,
		Obligations:                         obs,
		Verdict:                             strings.Join([]string{StatusFirewallsPreserved, StatusFailedUniqueCutoffProfileShapeNotDerived, StatusFailedCutoffProfileHigherMomentsNotLocked, StatusFailedZHNumericalValueStillSealed, StatusFailedYukawaAmplitudesStillSealed, StatusFailedHiggsMassQuarticStillFirewalled, StatusFailedGaugeCouplingsAbsoluteStillFirewalled, StatusFailedBGAPInstantonStillSealed}, ";"),
	}
}

func buildSummary(i Gate303Inheritance, p ContinuousProfileFormalization, c CanonicalProfileConstruction, m DiscreteContinuousMapSieve, s PromotionSeal, fw FirewallAudit) Summary {
	promoted := s.ActivatedConditionally && s.F0SourceSelectedBySeal && s.F0Positive && s.PromotedF0Value == contactZeta0
	return Summary{
		Gate303Inherited:                i.PositiveF0ClassFormalized && i.ContactCandidateAvailable && i.ContactCandidatePositive,
		ContinuousProfileFormalized:     p.PreservesGate302Sign && p.MomentFunctionalPositive,
		CanonicalPositiveProfileBuilt:   c.MomentEqualsContactZeta0 && c.PreservesPositivity && c.CanonicalNormalizationRule,
		DiscreteContinuousMapFormalized: m.MapIsAlgebraicallyExact && m.MapPreservesSign && m.MapIsUniqueAtCoefficientLevel,
		PromotionSealActivated:          s.ActivatedConditionally,
		ContactCutoffPromoted:           promoted,
		F0LockedBySeal:                  promoted,
		UniqueProfileShapeDerived:       c.UniqueProfileShapeDerived,
		HigherMomentsLocked:             m.LocksHigherMoments || s.HigherMomentsPromoted,
		NumericalZHComputed:             s.NumericalZHComputed,
		PhysicalDynamicsDerived:         s.HiggsPredictionClaimed || s.GaugeCouplingAbsoluteClaimed || s.BGapInstantonClaimed,
		FirewallPreserved:               !fw.FiniteCorePolluted && fw.NoObservedInputInserted && fw.NoYukawaNumbersInserted && fw.NoNumericalZHComputed && fw.NoHiggsMassQuarticClaimed && fw.NoAbsoluteGaugeCouplingsClaimed && fw.NoBGapInstantonClaimed && fw.NoUniqueProfileShapeClaimed && fw.NoHigherMomentLockClaimed,
		Status:                          strings.Join([]string{StatusCanonicalPositiveProfileConstructed, StatusContactSpectralCutoffPromoted, StatusFirewallsPreserved}, ";"),
		DirectAnswer:                    "Gate 304 constructs a positive continuous cutoff-profile source whose a4-channel moment is exactly ζ_contact(0)=7, thereby conditionally activating the ContactSpectralCutoffPromotionSeal and promoting f0 to the sealed value 7. The theorem fixes the coefficient, not a unique profile shape or physical Higgs/gauge dynamics.",
		NextGate:                        "Gate 305 should audit the scalar quadratic/quartic heat-kernel subtraction scheme: with f0 promoted and positive, the next obstruction is separating regulator/vacuum pieces from the a2 Higgs mass channel and the a4 quartic channel before any normalized Higgs prediction is legal.",
	}
}

func FormatGate303Inheritance(i Gate303Inheritance) string {
	return fmt.Sprintf("positiveClass=%t contactAvailable=%t exact=%s value=%d positive=%t observed=%t finalSourcePrev=%t promotionPrev=%t ZHPrev=%t dynamicsPrev=%t verdict=%s", i.PositiveF0ClassFormalized, i.ContactCandidateAvailable, i.ContactCandidateExact, i.ContactCandidateValue, i.ContactCandidatePositive, i.ContactCandidateObservedInput, i.FinalSourcePreviouslySelected, i.PromotionTheoremPreviouslyDerived, i.NumericalZHPreviouslyComputed, i.PhysicalDynamicsPreviouslyDerived, i.Verdict)
}

func FormatProfile(p ContinuousProfileFormalization) string {
	return fmt.Sprintf("class=%q functional=%q rule=%q positivity=%q regularity=%q finite=%q base=%q sign=%t positiveFunctional=%t observed=%t verdict=%s", p.ProfileClass, p.MomentFunctional, p.AbstractFunctionalRule, p.PositivityCondition, p.RegularityCondition, p.FiniteMomentCondition, p.AdmissibleBaseProfileCondition, p.PreservesGate302Sign, p.MomentFunctionalPositive, p.ObservedInputUsed, p.Verdict)
}

func FormatConstruction(c CanonicalProfileConstruction) string {
	return fmt.Sprintf("base=%s condition=%q formula=%q value=%d profile=%q nonnegative=%t continuous=%t smooth=%t finite=%t moment7=%t preserves=%t canonicalRule=%t uniqueShape=%t variational=%t observed=%t radial=%q eval=%q verdict=%s", c.BaseProfileSymbol, c.BaseProfileCondition, c.NormalizationFormula, c.ContactValue, c.ConstructedProfile, c.ProfileNonNegative, c.ProfileContinuous, c.ProfileSmooth, c.MomentFinite, c.MomentEqualsContactZeta0, c.PreservesPositivity, c.CanonicalNormalizationRule, c.UniqueProfileShapeDerived, c.VariationalPreferenceDerived, c.UsesObservedInput, c.RadialGaussianWitness, c.EvaluationFunctionalWitness, c.Verdict)
}

func FormatMap(m DiscreteContinuousMapSieve) string {
	return fmt.Sprintf("discrete=%s value=%d interpretation=%q target=%q formula=%q exact=%t sign=%t positiveFunctional=%t nonzeroBase=%t uniqueCoeff=%t uniqueShape=%t empirical=%t higherMoments=%t verdict=%s", m.DiscreteInvariant, m.DiscreteValue, m.DiscreteInterpretation, m.ContinuousTarget, m.MapFormula, m.MapIsAlgebraicallyExact, m.MapPreservesSign, m.MapUsesPositiveLinearFunctional, m.MapRequiresNonZeroBaseMoment, m.MapIsUniqueAtCoefficientLevel, m.MapIsUniqueAtProfileShapeLevel, m.ImportsEmpiricalData, m.LocksHigherMoments, m.Verdict)
}

func FormatSeal(s PromotionSeal) string {
	return fmt.Sprintf("name=%s activated=%t f0=%q value=%d positive=%t sourceSelected=%t contactPromoted=%t shapePromoted=%t higherMoments=%t coefficientPromoted=%t ZH=%t Higgs=%t gauge=%t BGap=%t verdict=%s", s.Name, s.ActivatedConditionally, s.PromotedF0Exact, s.PromotedF0Value, s.F0Positive, s.F0SourceSelectedBySeal, s.ContactSpectrumPromoted, s.ProfileShapePromoted, s.HigherMomentsPromoted, s.HeatKernelCoefficientPromoted, s.NumericalZHComputed, s.HiggsPredictionClaimed, s.GaugeCouplingAbsoluteClaimed, s.BGapInstantonClaimed, s.Verdict)
}

func FormatObligation(o RemainingObligation) string {
	return fmt.Sprintf("%s required=%q status=%s blocks=%t", o.Name, o.WhyRequired, o.Status, o.BlocksPrediction)
}

func FormatFirewalls(f FirewallAudit) string {
	obs := []string{}
	for _, o := range f.Obligations {
		obs = append(obs, FormatObligation(o))
	}
	return fmt.Sprintf("noObserved=%t noYukawa=%t noZH=%t noHiggs=%t noGauge=%t noBGap=%t noSubtraction=%t noUniqueShape=%t noHigherMoments=%t sealedSourceOnly=%t polluted=%t obligations=[%s] verdict=%s", f.NoObservedInputInserted, f.NoYukawaNumbersInserted, f.NoNumericalZHComputed, f.NoHiggsMassQuarticClaimed, f.NoAbsoluteGaugeCouplingsClaimed, f.NoBGapInstantonClaimed, f.NoHeatKernelSubtractionClaimed, f.NoUniqueProfileShapeClaimed, f.NoHigherMomentLockClaimed, f.ContactF0PromotedOnlyAsSealedSource, f.FiniteCorePolluted, strings.Join(obs, " | "), f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("gate303=%t profile=%t construction=%t map=%t seal=%t promoted=%t f0Locked=%t uniqueShape=%t higherMoments=%t ZH=%t dynamics=%t firewall=%t status=%s answer=%q next=%q", s.Gate303Inherited, s.ContinuousProfileFormalized, s.CanonicalPositiveProfileBuilt, s.DiscreteContinuousMapFormalized, s.PromotionSealActivated, s.ContactCutoffPromoted, s.F0LockedBySeal, s.UniqueProfileShapeDerived, s.HigherMomentsLocked, s.NumericalZHComputed, s.PhysicalDynamicsDerived, s.FirewallPreserved, s.Status, s.DirectAnswer, s.NextGate)
}
