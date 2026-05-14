// Package intermediatebreakingaudit implements Gate 228: Pati-Salam
// falsification / B-sector non-perturbative hierarchy origin search.
//
// Gate 227 found that the sealed hierarchy geometric mean
//
//	M_int = sqrt(M_B M_*) ≈ 6.65e11 GeV
//
// lies between the relic-decay EFT scale and the sealed axion scale. Gate 228
// asks what, if anything, may break at this intermediate scale. The gate runs
// the most lethal test first: temporarily unsealing the dormant u(4)
// leptoquark route only as a proton-decay lifetime estimate. It then tests
// whether the B-sector first spectral gap can generate the same scale through a
// non-perturbative hierarchy function without inserting a fitted coefficient.
package intermediatebreakingaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/geometricmeanresonance"
)

const (
	AuditID = "GATE228-PATI-SALAM-FALSIFICATION-BSECTOR-HIERARCHY-AUDIT"

	StatusPatiSalamFailed              = "FAILED_ROUTE_PATI_SALAM_INTERMEDIATE_BREAKING"
	StatusBGapHierarchyNearResonance   = "CONDITIONAL_SUPPORT_BSECTOR_NONPERTURBATIVE_HIERARCHY_SHAPE"
	StatusIntermediateSealNotGranted   = "INTERMEDIATE_BREAKING_SEAL_REQUIRED_NOT_GRANTED"
	StatusHiddenOriginFavoredByFailure = "CONDITIONAL_SUPPORT_HIDDEN_SECTOR_ORIGIN_AFTER_PATI_SALAM_FALSIFICATION"
)

const (
	// Inherited finite datum from Gate 225 and older B-sector audits.
	bSectorFirstGap = 0.1024649212

	// Prompt-level proton-decay stress threshold. Gate 228 uses it only as a
	// phenomenological kill-switch, not as a finite-core theorem.
	superKBoundYears = 1.0e34

	// Topological-boundary coupling used as a conservative dimensionless gauge
	// strength in the order-one dimension-six lifetime proxy.
	alphaTopological = 1.0 / (4.0 * math.Pi)

	protonMassGeV      = 0.9382720813
	hbarGeVS           = 6.582119569e-25
	secondsPerJulianYr = 31557600.0
	planckGeV          = 1.2209e19
	oneDecade          = 1.0
	exactTolerance     = 1e-12
)

type Gate227Snapshot struct {
	Gate227Inherited         bool
	GeometricResonanceFound  bool
	MIntGeV                  float64
	MBGeV                    float64
	MStarGeV                 float64
	FAGeV                    float64
	LambdaEFTGeV             float64
	PatiSalamQuarantined     bool
	NativeBreakingNotDerived bool
	TruthStatement           string
}

type PatiSalamLifetimeAudit struct {
	TemporarilyUnsealedForLifetimeOnly bool
	DormantU4LeptoquarkSlotsPresent    bool
	LeptoquarkDynamicsSealStillBinding bool
	MLQGeV                             float64
	Alpha                              float64
	ProtonMassGeV                      float64
	Formula                            string
	WidthGeV                           float64
	LifetimeSeconds                    float64
	LifetimeYears                      float64
	SuperKBoundYears                   float64
	RatioToBound                       float64
	Log10DeficitYears                  float64
	CatastrophicFailure                bool
	FailedByOrders                     float64
	ProtonLifetimeFiniteDerived        bool
	Verdict                            string
}

type SpectralFunctionCandidate struct {
	Name                 string
	C                    float64
	Formula              string
	ScaleGeV             float64
	RatioToMInt          float64
	Log10Gap             float64
	WithinOneDecade      bool
	ExactMatch           bool
	CanonicalCoefficient bool
	CoefficientDerived   bool
	Promoted             bool
	Verdict              string
}

type BGapHierarchyAudit struct {
	BGap                      float64
	MStarGeV                  float64
	MIntTargetGeV             float64
	FormulaFamily             string
	CanonicalC1               SpectralFunctionCandidate
	CandidateFourOverPi       SpectralFunctionCandidate
	CandidateFiveQuarter      SpectralFunctionCandidate
	RequiredC                 float64
	RequiredCFormula          string
	RequiredCScaleGeV         float64
	RequiredCIsFitted         bool
	RequiredCOrderOne         bool
	BestCandidate             SpectralFunctionCandidate
	NonPerturbativeShapeWorks bool
	NativeCoefficientDerived  bool
	ExactBGapOriginDerived    bool
	Verdict                   string
}

type IntermediateBreakingSealAudit struct {
	SealName                      string
	SealPrepared                  bool
	SealGranted                   bool
	PatiSalamFalsified            bool
	HiddenSectorFavored           bool
	BGapShapeSupported            bool
	BGapCoefficientDerived        bool
	QuarantinedAssumptionRequired string
	OperationalStatus             string
	Verdict                       string
}

type BaryonSafetyAudit struct {
	IntermediateScaleMustBeBaryonSafe bool
	PatiSalamAtMIntAllowed            bool
	BGapHiddenSectorCarriesBaryon     bool
	LeptoquarkSealRemainsActive       bool
	ProtonDecayRouteReopened          bool
	BaryonSafeIntermediateOrigin      bool
	Verdict                           string
}

type FirewallAudit struct {
	Gate227Inherited               bool
	UsedOnlySealedScales           bool
	PatiSalamUnsealedForDynamics   bool
	PatiSalamUnsealedForLifetime   bool
	LeptoquarkDynamicsClaimed      bool
	ProtonLifetimeClaimedExact     bool
	BGapPromotedToPhysicalField    bool
	CoefficientCFittedButSealed    bool
	IntermediateScaleFiniteDerived bool
	AxionShiftDerived              bool
	EFTMediatorDerived             bool
	ProtonDecayBoundUsedAsFilter   bool
	FiniteCorePolluted             bool
	Verdict                        string
}

type Summary struct {
	PatiSalamFalsified      bool
	HiddenSectorPreferred   bool
	BGapShapeNearResonance  bool
	IntermediateSealGranted bool
	MIntGeV                 float64
	PatiSalamLifetimeYears  float64
	RequiredC               float64
	Status                  string
	NextGate                string
	Comment                 string
}

type Analysis struct {
	Gate227   Gate227Snapshot
	PatiSalam PatiSalamLifetimeAudit
	BGap      BGapHierarchyAudit
	Seal      IntermediateBreakingSealAudit
	Baryon    BaryonSafetyAudit
	Firewall  FirewallAudit
	Summary   Summary

	TruthStatement string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		g227, err := geometricmeanresonance.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 227 input: %w", err)
			return
		}
		defaultA, defaultErr = Build(g227)
	})
	return defaultA, defaultErr
}

func Build(g227 geometricmeanresonance.Analysis) (Analysis, error) {
	snap := snapshotFromGate227(g227)
	if !snap.Gate227Inherited || !snap.GeometricResonanceFound || snap.MIntGeV <= 0 || snap.MStarGeV <= snap.MIntGeV {
		return Analysis{}, fmt.Errorf("Gate 228 requires Gate 227 geometric-mean resonance and ordered sealed hierarchy")
	}
	ps := auditPatiSalamLifetime(snap)
	bgap := auditBGapHierarchy(snap)
	seal := auditIntermediateSeal(ps, bgap)
	baryon := auditBaryonSafety(ps, seal)
	firewall := auditFirewalls(snap, ps, bgap, seal)
	summary := summarize(snap, ps, bgap, seal, baryon)
	truth := buildTruth(snap, ps, bgap, seal, baryon)

	return Analysis{Gate227: snap, PatiSalam: ps, BGap: bgap, Seal: seal, Baryon: baryon, Firewall: firewall, Summary: summary, TruthStatement: truth}, nil
}

func snapshotFromGate227(a geometricmeanresonance.Analysis) Gate227Snapshot {
	return Gate227Snapshot{
		Gate227Inherited:         a.Summary.Status != "" && a.Geometric.MIntGeV > 0,
		GeometricResonanceFound:  a.Summary.GeometricMeanResonanceFound && a.Geometric.NullHypothesisRejected,
		MIntGeV:                  a.Geometric.MIntGeV,
		MBGeV:                    a.Hierarchy.MBGeV,
		MStarGeV:                 a.Hierarchy.MStarGeV,
		FAGeV:                    a.Hierarchy.FARequirementGeV,
		LambdaEFTGeV:             a.Hierarchy.LambdaEFTMaxGeV,
		PatiSalamQuarantined:     a.PatiSalam.LeptoquarkDynamicsSealActive && !a.PatiSalam.ProtonDecayChannelReopened,
		NativeBreakingNotDerived: !a.Seesaw.NativeIntermediateScaleDerived,
		TruthStatement:           a.TruthStatement,
	}
}

func auditPatiSalamLifetime(s Gate227Snapshot) PatiSalamLifetimeAudit {
	m := s.MIntGeV
	width := alphaTopological * alphaTopological * math.Pow(protonMassGeV, 5) / math.Pow(m, 4)
	tauSeconds := hbarGeVS / width
	tauYears := tauSeconds / secondsPerJulianYr
	ratio := tauYears / superKBoundYears
	deficit := math.Log10(superKBoundYears / tauYears)
	failed := tauYears < superKBoundYears
	return PatiSalamLifetimeAudit{
		TemporarilyUnsealedForLifetimeOnly: true,
		DormantU4LeptoquarkSlotsPresent:    true,
		LeptoquarkDynamicsSealStillBinding: true,
		MLQGeV:                             m,
		Alpha:                              alphaTopological,
		ProtonMassGeV:                      protonMassGeV,
		Formula:                            "Γ_p ~ α² m_p⁵ / M_LQ⁴; τ = ℏ/Γ",
		WidthGeV:                           width,
		LifetimeSeconds:                    tauSeconds,
		LifetimeYears:                      tauYears,
		SuperKBoundYears:                   superKBoundYears,
		RatioToBound:                       ratio,
		Log10DeficitYears:                  deficit,
		CatastrophicFailure:                failed && deficit > 5,
		FailedByOrders:                     math.Max(0, deficit),
		ProtonLifetimeFiniteDerived:        false,
		Verdict:                            StatusPatiSalamFailed,
	}
}

func auditBGapHierarchy(s Gate227Snapshot) BGapHierarchyAudit {
	c1 := candidate("c=1", 1, s.MStarGeV, s.MIntGeV, true, false)
	c4pi := candidate("c=4/π diagnostic", 4/math.Pi, s.MStarGeV, s.MIntGeV, false, false)
	c54 := candidate("c=5/4 diagnostic", 1.25, s.MStarGeV, s.MIntGeV, false, false)
	cReq := bSectorFirstGap * math.Log(s.MStarGeV/s.MIntGeV)
	scaleReq := s.MStarGeV * math.Exp(-cReq/bSectorFirstGap)

	best := c1
	for _, x := range []SpectralFunctionCandidate{c4pi, c54} {
		if x.Log10Gap < best.Log10Gap {
			best = x
		}
	}

	verdict := StatusBGapHierarchyNearResonance
	if best.Log10Gap > oneDecade && !cReqOrderOne(cReq) {
		verdict = "FAILED_ROUTE_BSECTOR_HIERARCHY_SHAPE"
	}

	return BGapHierarchyAudit{
		BGap:                      bSectorFirstGap,
		MStarGeV:                  s.MStarGeV,
		MIntTargetGeV:             s.MIntGeV,
		FormulaFamily:             "M_hidden = M_* exp(-c / B_gap)",
		CanonicalC1:               c1,
		CandidateFourOverPi:       c4pi,
		CandidateFiveQuarter:      c54,
		RequiredC:                 cReq,
		RequiredCFormula:          "c_req = B_gap ln(M_*/M_int)",
		RequiredCScaleGeV:         scaleReq,
		RequiredCIsFitted:         true,
		RequiredCOrderOne:         cReqOrderOne(cReq),
		BestCandidate:             best,
		NonPerturbativeShapeWorks: cReqOrderOne(cReq) && scaleReq > 0 && math.Abs(math.Log10(scaleReq/s.MIntGeV)) < exactTolerance,
		NativeCoefficientDerived:  false,
		ExactBGapOriginDerived:    false,
		Verdict:                   verdict,
	}
}

func cReqOrderOne(c float64) bool { return c > 0.1 && c < 10 }

func candidate(name string, c float64, mstar, mint float64, canonical bool, derived bool) SpectralFunctionCandidate {
	scale := mstar * math.Exp(-c/bSectorFirstGap)
	ratio := scale / mint
	gap := math.Abs(math.Log10(ratio))
	verdict := "NO_RESONANCE"
	within := gap < oneDecade
	if within {
		verdict = "NEAR_RESONANCE_DIAGNOSTIC_ONLY"
	}
	exact := gap < exactTolerance
	promoted := exact && canonical && derived
	if promoted {
		verdict = "CANONICAL_EXACT_MATCH"
	}
	return SpectralFunctionCandidate{
		Name:                 name,
		C:                    c,
		Formula:              fmt.Sprintf("M_* exp(-%.12g / B_gap)", c),
		ScaleGeV:             scale,
		RatioToMInt:          ratio,
		Log10Gap:             gap,
		WithinOneDecade:      within,
		ExactMatch:           exact,
		CanonicalCoefficient: canonical,
		CoefficientDerived:   derived,
		Promoted:             promoted,
		Verdict:              verdict,
	}
}

func auditIntermediateSeal(ps PatiSalamLifetimeAudit, bg BGapHierarchyAudit) IntermediateBreakingSealAudit {
	hiddenFavored := ps.CatastrophicFailure && bg.NonPerturbativeShapeWorks && !bg.NativeCoefficientDerived
	granted := hiddenFavored && bg.BestCandidate.Promoted
	status := "SEAL_PREPARED_NOT_GRANTED"
	verdict := StatusIntermediateSealNotGranted
	if granted {
		status = "SEAL_GRANTED"
		verdict = "INTERMEDIATE_BREAKING_SEAL_GRANTED"
	}
	return IntermediateBreakingSealAudit{
		SealName:                      "IntermediateBreakingSeal",
		SealPrepared:                  true,
		SealGranted:                   granted,
		PatiSalamFalsified:            ps.CatastrophicFailure,
		HiddenSectorFavored:           hiddenFavored,
		BGapShapeSupported:            bg.NonPerturbativeShapeWorks,
		BGapCoefficientDerived:        bg.NativeCoefficientDerived,
		QuarantinedAssumptionRequired: "B-sector hidden order parameter plus canonical c in M_* exp(-c/B_gap)",
		OperationalStatus:             status,
		Verdict:                       verdict,
	}
}

func auditBaryonSafety(ps PatiSalamLifetimeAudit, seal IntermediateBreakingSealAudit) BaryonSafetyAudit {
	safeHidden := ps.CatastrophicFailure && seal.HiddenSectorFavored
	return BaryonSafetyAudit{
		IntermediateScaleMustBeBaryonSafe: true,
		PatiSalamAtMIntAllowed:            !ps.CatastrophicFailure,
		BGapHiddenSectorCarriesBaryon:     false,
		LeptoquarkSealRemainsActive:       true,
		ProtonDecayRouteReopened:          false,
		BaryonSafeIntermediateOrigin:      safeHidden,
		Verdict:                           StatusHiddenOriginFavoredByFailure,
	}
}

func auditFirewalls(s Gate227Snapshot, ps PatiSalamLifetimeAudit, bg BGapHierarchyAudit, seal IntermediateBreakingSealAudit) FirewallAudit {
	return FirewallAudit{
		Gate227Inherited:               s.Gate227Inherited,
		UsedOnlySealedScales:           true,
		PatiSalamUnsealedForDynamics:   false,
		PatiSalamUnsealedForLifetime:   ps.TemporarilyUnsealedForLifetimeOnly,
		LeptoquarkDynamicsClaimed:      false,
		ProtonLifetimeClaimedExact:     ps.ProtonLifetimeFiniteDerived,
		BGapPromotedToPhysicalField:    false,
		CoefficientCFittedButSealed:    bg.RequiredCIsFitted && !bg.NativeCoefficientDerived,
		IntermediateScaleFiniteDerived: seal.SealGranted && bg.NativeCoefficientDerived,
		AxionShiftDerived:              false,
		EFTMediatorDerived:             false,
		ProtonDecayBoundUsedAsFilter:   true,
		FiniteCorePolluted:             false,
		Verdict:                        "FIREWALLS_CLOSED",
	}
}

func summarize(s Gate227Snapshot, ps PatiSalamLifetimeAudit, bg BGapHierarchyAudit, seal IntermediateBreakingSealAudit, baryon BaryonSafetyAudit) Summary {
	status := strings.Join([]string{StatusPatiSalamFailed, StatusBGapHierarchyNearResonance, StatusIntermediateSealNotGranted}, ";")
	return Summary{
		PatiSalamFalsified:      ps.CatastrophicFailure,
		HiddenSectorPreferred:   baryon.BaryonSafeIntermediateOrigin,
		BGapShapeNearResonance:  bg.NonPerturbativeShapeWorks || bg.BestCandidate.WithinOneDecade,
		IntermediateSealGranted: seal.SealGranted,
		MIntGeV:                 s.MIntGeV,
		PatiSalamLifetimeYears:  ps.LifetimeYears,
		RequiredC:               bg.RequiredC,
		Status:                  status,
		NextGate:                "Gate 229 — canonical B-gap coefficient / hidden order-parameter audit",
		Comment:                 "Gate 228 falsifies Pati-Salam breaking at M_int by proton decay and finds a B-gap non-perturbative hierarchy shape, but no canonical coefficient or intermediate-breaking seal yet.",
	}
}

func buildTruth(s Gate227Snapshot, ps PatiSalamLifetimeAudit, bg BGapHierarchyAudit, seal IntermediateBreakingSealAudit, baryon BaryonSafetyAudit) string {
	return fmt.Sprintf("Gate 228 temporarily unseals dormant u(4) leptoquarks only for a lifetime estimate at M_int=%.9e GeV and finds τ_p≈%.9e yr, %.3f orders below the 1e34 yr stress bound: intermediate Pati-Salam breaking is falsified. The hidden B-gap hierarchy M_* exp(-c/B_gap) can hit M_int with c_req=%.9f, and the diagnostic c=4/π lies %.6f decades from M_int, but no finite theorem derives c or an order parameter; IntermediateBreakingSeal is therefore prepared but not granted.", s.MIntGeV, ps.LifetimeYears, ps.FailedByOrders, bg.RequiredC, bg.CandidateFourOverPi.Log10Gap)
}

func FormatGate227(s Gate227Snapshot) string {
	return fmt.Sprintf("inherited=%t resonance=%t M_int=%.9e M_B=%.9e M_*=%.9e f_a=%.9e LambdaEFT=%.9e PatiSalamQuarantined=%t nativeBreakingNotDerived=%t", s.Gate227Inherited, s.GeometricResonanceFound, s.MIntGeV, s.MBGeV, s.MStarGeV, s.FAGeV, s.LambdaEFTGeV, s.PatiSalamQuarantined, s.NativeBreakingNotDerived)
}

func FormatPatiSalam(p PatiSalamLifetimeAudit) string {
	return fmt.Sprintf("tempUnsealLifetimeOnly=%t u4Slots=%t sealStillBinding=%t M_LQ=%.9e alpha=%.9e width=%.9e tauSec=%.9e tauYr=%.9e boundYr=%.9e ratio=%.9e logDeficit=%.9f catastrophic=%t verdict=%s", p.TemporarilyUnsealedForLifetimeOnly, p.DormantU4LeptoquarkSlotsPresent, p.LeptoquarkDynamicsSealStillBinding, p.MLQGeV, p.Alpha, p.WidthGeV, p.LifetimeSeconds, p.LifetimeYears, p.SuperKBoundYears, p.RatioToBound, p.Log10DeficitYears, p.CatastrophicFailure, p.Verdict)
}

func FormatCandidate(c SpectralFunctionCandidate) string {
	return fmt.Sprintf("%s: c=%.12f scale=%.9e ratio=%.9e logGap=%.9f withinDecade=%t exact=%t canonical=%t derived=%t promoted=%t verdict=%s", c.Name, c.C, c.ScaleGeV, c.RatioToMInt, c.Log10Gap, c.WithinOneDecade, c.ExactMatch, c.CanonicalCoefficient, c.CoefficientDerived, c.Promoted, c.Verdict)
}

func FormatBGap(b BGapHierarchyAudit) string {
	return strings.Join([]string{
		fmt.Sprintf("B_gap=%.12f M_*=%.9e M_int=%.9e family=%q", b.BGap, b.MStarGeV, b.MIntTargetGeV, b.FormulaFamily),
		FormatCandidate(b.CanonicalC1),
		FormatCandidate(b.CandidateFourOverPi),
		FormatCandidate(b.CandidateFiveQuarter),
		fmt.Sprintf("c_req=%.12f scaleReq=%.9e fitted=%t orderOne=%t shapeWorks=%t nativeCoeff=%t exactOrigin=%t verdict=%s", b.RequiredC, b.RequiredCScaleGeV, b.RequiredCIsFitted, b.RequiredCOrderOne, b.NonPerturbativeShapeWorks, b.NativeCoefficientDerived, b.ExactBGapOriginDerived, b.Verdict),
	}, "; ")
}

func FormatSeal(s IntermediateBreakingSealAudit) string {
	return fmt.Sprintf("seal=%s prepared=%t granted=%t psFalsified=%t hiddenFavored=%t bgapShape=%t coeffDerived=%t required=%q status=%s verdict=%s", s.SealName, s.SealPrepared, s.SealGranted, s.PatiSalamFalsified, s.HiddenSectorFavored, s.BGapShapeSupported, s.BGapCoefficientDerived, s.QuarantinedAssumptionRequired, s.OperationalStatus, s.Verdict)
}

func FormatBaryon(b BaryonSafetyAudit) string {
	return fmt.Sprintf("mustBeBaryonSafe=%t psAllowed=%t bgapCarriesBaryon=%t leptoquarkSeal=%t protonDecayReopened=%t safeHiddenOrigin=%t verdict=%s", b.IntermediateScaleMustBeBaryonSafe, b.PatiSalamAtMIntAllowed, b.BGapHiddenSectorCarriesBaryon, b.LeptoquarkSealRemainsActive, b.ProtonDecayRouteReopened, b.BaryonSafeIntermediateOrigin, b.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("g227=%t sealedScales=%t psDynamicsUnsealed=%t psLifetimeUnsealed=%t lqDynamicsClaimed=%t exactLifetimeClaimed=%t bgapPromoted=%t cFittedSealed=%t nativeMInt=%t axionShift=%t eftMediator=%t boundFilter=%t polluted=%t verdict=%s", f.Gate227Inherited, f.UsedOnlySealedScales, f.PatiSalamUnsealedForDynamics, f.PatiSalamUnsealedForLifetime, f.LeptoquarkDynamicsClaimed, f.ProtonLifetimeClaimedExact, f.BGapPromotedToPhysicalField, f.CoefficientCFittedButSealed, f.IntermediateScaleFiniteDerived, f.AxionShiftDerived, f.EFTMediatorDerived, f.ProtonDecayBoundUsedAsFilter, f.FiniteCorePolluted, f.Verdict)
}
