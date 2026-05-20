// Package generation2strongthresholdsignfieldcontentviabilityaudit implements
// Gate 609: Strong Threshold Sign and Field-Content Viability Audit.
//
// Gate 608 showed that the gauge pairwise meeting scales form a skew triangle
// and that closing the strong residual at Lambda_12, if modeled as a constant
// full-interval one-loop deformation, requires Delta b3 < 0.  Gate 609 audits
// the sign and viability of possible correction origins without introducing
// new fields, claiming threshold existence, or promoting gauge unification.
package generation2strongthresholdsignfieldcontentviabilityaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2gaugemeetingscaletrianglegeometryaudit"
)

const (
	AuditID = "GATE609-STRONG-THRESHOLD-SIGN-FIELD-CONTENT-VIABILITY-AUDIT"

	StatusGate608Inherited             = "PASS_GATE608_STRONG_RESIDUAL_INHERITED"
	StatusRequiredSignClassified       = "PASS_SIGN_OF_REQUIRED_DELTA_B3_CLASSIFIED"
	StatusOrdinaryMatterAudited        = "PASS_ORDINARY_MATTER_SIGN_AUDITED"
	StatusCorrectionOriginsClassified  = "PASS_CORRECTION_ORIGIN_VIABILITY_TABLE_CONSTRUCTED"
	StatusBoundaryThresholdCompatible  = "CONDITIONAL_SUPPORT_BOUNDARY_LOCALIZED_THRESHOLD_SIGN_COMPATIBLE"
	StatusFSABoundarySlotDefined       = "CONDITIONAL_SUPPORT_FINITE_SPECTRAL_ACTION_BOUNDARY_CORRECTION_SLOT_DEFINED"
	StatusSchemeTwoLoopSlotRelevant    = "CONDITIONAL_SUPPORT_SCHEME_AND_TWO_LOOP_SLOTS_SIGN_RELEVANT_BUT_UNCERTIFIED"
	StatusExtraColoredWrongSign        = "FAILED_ROUTE_SIMPLE_EXTRA_COLORED_MATTER_FULL_INTERVAL_HAS_WRONG_SIGN"
	StatusNoNativeStrongThreshold      = "FAILED_ROUTE_NO_NATIVE_STRONG_THRESHOLD_THEOREM"
	StatusNoNativeColorKineticBoundary = "FAILED_ROUTE_NO_NATIVE_COLOR_KINETIC_BOUNDARY_CORRECTION"
	StatusNoExtendedGaugeSectorNative  = "FAILED_ROUTE_EXTENDED_GAUGE_SECTOR_NOT_NATIVE"
	StatusNoGaugeUnificationClaim      = "FAILED_ROUTE_NO_GAUGE_UNIFICATION_CLAIM"
	StatusNoNewParticles               = "FAILED_ROUTE_NO_NEW_PARTICLES_INTRODUCED"
	StatusNoThresholdExistenceClaim    = "FAILED_ROUTE_NO_THRESHOLD_EXISTENCE_CLAIM"
	StatusGate609Boundary              = "FIREWALL_PRESERVED_GATE609_STRONG_THRESHOLD_SIGN_BOUNDARY"
	StatusGate606608Preserved          = "FIREWALL_PRESERVED_GATES606_608_TRANSPORT_BOUNDARIES"
)

const (
	b1SM = 41.0 / 10.0
	b2SM = -19.0 / 6.0
	b3SM = -7.0
)

type InheritedGate608 struct {
	Lambda12GeV             float64
	Lambda13GeV             float64
	Lambda23GeV             float64
	Delta3ThresholdRequired float64
	DeltaAlpha3InvRequired  float64
	DeltaB3Required         float64
	B3SM                    float64
	B3EffectiveDiagnostic   float64
	RelativeB3Deformation   float64
	Verdict                 string
}

type SignConventionRow struct {
	Statement      string
	Equation       string
	RequiredSign   string
	Interpretation string
	Verdict        string
}

type WrongSignMatterAudit struct {
	OrdinaryMatterContributionSign string
	Reason                         string
	RequiredSign                   string
	Conclusion                     string
	Verdict                        string
}

type CorrectionOriginViabilityRow struct {
	Origin         string
	Class          string
	ExpectedSign   string
	SignCompatible bool
	SizeComment    string
	Native         bool
	Certified      bool
	Interpretation string
	Verdict        string
}

type BoundaryThresholdSlotAudit struct {
	SlotName              string
	RequiredDeltaU        float64
	RequiredDeltaAlpha    float64
	SignCompatible        bool
	UniformBetaEquivalent float64
	Interpretation        string
	Verdict               string
}

type NativeASHAStatus struct {
	HasNativeStrongThresholdTheorem bool
	HasNativeColorKineticBoundary   bool
	HasNativeExtraColoredSpectrum   bool
	HasNativeGaugeSectorExtension   bool
	ClaimsUnification               bool
	Statement                       string
	Verdict                         string
}

type Firewalls struct {
	IntroducesNewParticles   bool
	ClaimsThresholdExistence bool
	ClaimsGaugeUnification   bool
	AltersAF                 bool
	DerivesEndpoint          bool
	Verdict                  string
}

type Analysis struct {
	Inherited         InheritedGate608
	SignConventions   []SignConventionRow
	WrongSignMatter   WrongSignMatterAudit
	CorrectionOrigins []CorrectionOriginViabilityRow
	BoundaryThreshold BoundaryThresholdSlotAudit
	NativeStatus      NativeASHAStatus
	Firewalls         Firewalls
	Truth             string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	g608, err := generation2gaugemeetingscaletrianglegeometryaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate608 predecessor: %w", err)
	}
	inh := inherit(g608)
	signs := buildSignConventions(inh)
	wrong := buildWrongSignMatterAudit(inh)
	origins := buildCorrectionOrigins(inh)
	boundary := buildBoundaryThresholdSlot(inh)
	native := buildNativeStatus()
	firewalls := auditFirewalls()
	truth := "Gate 609 classifies the strong residual sign.  The Lambda_12 wound requires a positive inverse-coupling boundary correction or, if smeared over the full interval, a negative Delta b3.  Ordinary full-interval extra colored matter has the wrong one-loop sign, while boundary-localized threshold, scheme/matching, or finite spectral-action color-kinetic correction slots are sign-compatible but uncertified."
	return Analysis{inh, signs, wrong, origins, boundary, native, firewalls, truth}, nil
}

func inherit(a generation2gaugemeetingscaletrianglegeometryaudit.Analysis) InheritedGate608 {
	db3 := a.Inherited.DeltaB3Required
	return InheritedGate608{
		Lambda12GeV:             a.Inherited.Lambda12GeV,
		Lambda13GeV:             a.Inherited.Lambda13GeV,
		Lambda23GeV:             a.Inherited.Lambda23GeV,
		Delta3ThresholdRequired: a.Inherited.Delta3ThresholdRequired,
		DeltaAlpha3InvRequired:  a.Inherited.DeltaAlpha3InvRequired,
		DeltaB3Required:         db3,
		B3SM:                    b3SM,
		B3EffectiveDiagnostic:   b3SM + db3,
		RelativeB3Deformation:   math.Abs(db3 / b3SM),
		Verdict:                 StatusGate608Inherited,
	}
}

func buildSignConventions(i InheritedGate608) []SignConventionRow {
	return []SignConventionRow{
		{
			Statement:      "one-loop beta convention",
			Equation:       "dg_i/dlnmu = b_i g_i^3/(16*pi^2); d(1/g_i^2)/dlnmu = -b_i/(8*pi^2)",
			RequiredSign:   "Delta b3 < 0 for a full-interval deformation",
			Interpretation: fmt.Sprintf("Delta b3=%0.15g makes b3_eff=%0.15g, increasing u3=1/g3^2 at high scale", i.DeltaB3Required, i.B3EffectiveDiagnostic),
			Verdict:        StatusRequiredSignClassified,
		},
		{
			Statement:      "boundary threshold convention",
			Equation:       "1/g3_eff^2 = 1/g3_runtime^2 + delta_3^threshold",
			RequiredSign:   "delta_3^threshold > 0",
			Interpretation: fmt.Sprintf("required delta_3^threshold=%0.15g, equivalent Delta alpha_3^-1=%0.15g", i.Delta3ThresholdRequired, i.DeltaAlpha3InvRequired),
			Verdict:        StatusBoundaryThresholdCompatible,
		},
	}
}

func buildWrongSignMatterAudit(i InheritedGate608) WrongSignMatterAudit {
	return WrongSignMatterAudit{
		OrdinaryMatterContributionSign: "Delta b3 > 0 for additional Weyl/Dirac fermions or complex scalars in colored representations under the standard one-loop non-Abelian convention",
		Reason:                         "matter terms are proportional to positive Dynkin indices T(R), making non-Abelian b_i less negative rather than more negative",
		RequiredSign:                   fmt.Sprintf("Gate 609 requires Delta b3=%0.15g < 0 if modeled over the full interval", i.DeltaB3Required),
		Conclusion:                     "simple full-interval extra colored matter moves the strong running in the wrong direction for the Lambda_12 residual",
		Verdict:                        StatusExtraColoredWrongSign,
	}
}

func buildCorrectionOrigins(i InheritedGate608) []CorrectionOriginViabilityRow {
	return []CorrectionOriginViabilityRow{
		{"full-interval extra colored matter", "field-content beta deformation", "Delta b3 > 0", false, "wrong sign relative to required Delta b3<0", false, false, "ordinary colored matter would make QCD less asymptotically free over the interval", StatusExtraColoredWrongSign},
		{"full-interval extra electroweak/colorless matter", "field-content beta deformation", "mostly Delta b1/Delta b2 > 0", false, "does not directly supply the required positive strong inverse-coupling jump", false, false, "may move electroweak slopes but does not solve the strong residual at Lambda_12 by itself", StatusNoNativeStrongThreshold},
		{"extended gauge-sector contribution", "gauge self-interaction beta deformation", "can be Delta b3 < 0 in principle", true, "sign-compatible only with an actual gauge-sector theorem", false, false, "would require changing or extending the native gauge sector beyond current A_F data", StatusNoExtendedGaugeSectorNative},
		{"boundary-localized strong threshold", "localized matching jump", "delta_3^threshold > 0", true, fmt.Sprintf("requires delta_u=%0.15g or Delta alpha^-1=%0.15g", i.Delta3ThresholdRequired, i.DeltaAlpha3InvRequired), false, false, "sign-compatible because it shifts u3 upward without a full-interval beta deformation", StatusBoundaryThresholdCompatible},
		{"finite spectral-action color kinetic boundary correction", "boundary color normalization/matching", "delta_3^FSA_boundary > 0", true, "could act as inverse-coupling boundary offset if a theorem supplied it", false, false, "symbolic slot only; no native color kinetic correction theorem currently exists", StatusFSABoundarySlotDefined},
		{"renormalization scheme and low-energy matching", "scheme / endpoint extraction", "not fixed a priori", true, "must be computed with version-pinned extraction and matching", false, false, "relevant transport slot but not certified as large enough or correctly signed", StatusSchemeTwoLoopSlotRelevant},
		{"two-loop SM running", "higher-loop transport curvature", "not reducible to constant Delta b3", true, "may alter the diagnostic residual, but v1 does not include it", false, false, "calculation slot only; not a native threshold theorem", StatusSchemeTwoLoopSlotRelevant},
		{"native ASHA strong threshold", "theorem-required threshold spectrum", "would need delta_3^threshold > 0", false, "not present", false, false, "no native strong-threshold theorem exists in current project data", StatusNoNativeStrongThreshold},
	}
}

func buildBoundaryThresholdSlot(i InheritedGate608) BoundaryThresholdSlotAudit {
	return BoundaryThresholdSlotAudit{
		SlotName:              "delta_3^threshold",
		RequiredDeltaU:        i.Delta3ThresholdRequired,
		RequiredDeltaAlpha:    i.DeltaAlpha3InvRequired,
		SignCompatible:        i.Delta3ThresholdRequired > 0,
		UniformBetaEquivalent: i.DeltaB3Required,
		Interpretation:        "a localized boundary threshold is sign-compatible because it raises 1/g3^2 at Lambda_12 directly; the uniform beta-equivalent Delta b3 is a diagnostic only",
		Verdict:               StatusBoundaryThresholdCompatible,
	}
}

func buildNativeStatus() NativeASHAStatus {
	return NativeASHAStatus{
		HasNativeStrongThresholdTheorem: false,
		HasNativeColorKineticBoundary:   false,
		HasNativeExtraColoredSpectrum:   false,
		HasNativeGaugeSectorExtension:   false,
		ClaimsUnification:               false,
		Statement:                       "Current ASHA data expose the strong residual and native gauge sockets, but do not supply a threshold spectrum, color kinetic boundary correction, extended gauge-sector theorem, or full unification theorem.",
		Verdict:                         StatusNoNativeStrongThreshold,
	}
}

func auditFirewalls() Firewalls {
	return Firewalls{false, false, false, false, false, StatusGate609Boundary}
}

func Statuses() []string {
	return []string{
		StatusGate608Inherited,
		StatusRequiredSignClassified,
		StatusOrdinaryMatterAudited,
		StatusCorrectionOriginsClassified,
		StatusBoundaryThresholdCompatible,
		StatusFSABoundarySlotDefined,
		StatusSchemeTwoLoopSlotRelevant,
		StatusExtraColoredWrongSign,
		StatusNoNativeStrongThreshold,
		StatusNoNativeColorKineticBoundary,
		StatusNoExtendedGaugeSectorNative,
		StatusNoGaugeUnificationClaim,
		StatusNoNewParticles,
		StatusNoThresholdExistenceClaim,
		StatusGate609Boundary,
		StatusGate606608Preserved,
	}
}

func containsOrigin(rows []CorrectionOriginViabilityRow, key string) bool {
	for _, r := range rows {
		if strings.Contains(r.Origin, key) || strings.Contains(r.Class, key) || strings.Contains(r.Verdict, key) {
			return true
		}
	}
	return false
}
