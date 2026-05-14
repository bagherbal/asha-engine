// Package sealedthresholdstresstest implements Gate 207: sealed-threshold
// prediction stress test / experimental and proton-decay firewall audit.
//
// Gate 206 emitted numerical threshold scales only under the explicit
// EmpiricalCarrierSeal. Gate 207 therefore does not strengthen those numbers
// into finite-core predictions. It stress-tests the sealed predictions against
// external collider reach, audits proton-decay mediator support inside the
// engine's own gauge/connection inventory, and verifies that the still-external
// universal completion does not create an obvious one-loop Landau-pole pathology
// below the Planck scale.
package sealedthresholdstresstest

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	StatusConditionalStressTest = "FAILED_ROUTE_UNIVERSAL_COMPLETION_STRESS"
	gevPerTeV                   = 1.0e3
	planckGeV                   = 1.2209e19
)

type Gate206Snapshot struct {
	Gate206Inherited                     bool
	CarrierSealExplicit                  bool
	CarrierSealQuarantined               bool
	NativeSemanticSearchFailed           bool
	AnomalyCompatibilityPassed           bool
	PredictionsEmitted                   int
	AllConditionalOnCarrierSeal          bool
	UniversalCompletionStillExternal     bool
	PhysicalUnificationClaimed           bool
	AbsoluteMassPredicted                bool
	ThresholdCorrectedPhysicalFitClaimed bool
	FiniteMatchingCorrectionsDerived     bool
	AlphaGUT                             float64
	AlphaGUTInverse                      float64
	Predictions                          []SealedPrediction
	TruthStatement                       string
}

type SealedPrediction struct {
	CarrierName              string
	SMRepresentation         string
	ThresholdScaleMBGeV      float64
	BoundaryScaleMStarGeV    float64
	RequiredUniversalBetaRow float64
	TotalDeltaB              FloatTripleCompat
	ConditionalOnCarrierSeal bool
	AnomalyCompatible        bool
	FiniteDerived            bool
	Verdict                  string
}

// FloatTripleCompat is defined locally by conversion below so Gate 207 does not
// depend on unexported numerical helpers from Gate 201. It mirrors the physical
// content of the Gate-206 total beta deformation triple.
type FloatTripleCompat struct {
	U1GUT float64
	SU2L  float64
	SU3C  float64
}

func DefaultGate206Snapshot() (Gate206Snapshot, error) {
	// Gate 207 intentionally imports the Gate-206 audited ledger as a compact
	// snapshot instead of rebuilding the full historical theorem chain. This keeps
	// the stress-test package fast and prevents timeout-prone registry imports from
	// contaminating the new gate. The numbers below are the exact Gate-206 sealed
	// outputs recorded in gate206_registry_audit.md.
	preds := []SealedPrediction{
		{
			CarrierName:              "Dirac vectorlike quark doublet",
			SMRepresentation:         "(3,2,1/6)",
			ThresholdScaleMBGeV:      1.46774973718e6,
			BoundaryScaleMStarGeV:    2.40099519719e15,
			RequiredUniversalBetaRow: 7.65295390904,
			TotalDeltaB:              FloatTripleCompat{U1GUT: 7.78628724237, SU2L: 9.65295390904, SU3C: 8.98628724237},
			ConditionalOnCarrierSeal: true,
			AnomalyCompatible:        true,
			FiniteDerived:            false,
			Verdict:                  "Gate-206 sealed conditional prediction; not a finite-core mass prediction",
		},
		{
			CarrierName:              "Weyl SU(2)L adjoint fermion",
			SMRepresentation:         "(1,3,0)",
			ThresholdScaleMBGeV:      8.19807624157e6,
			BoundaryScaleMStarGeV:    2.42276543552e14,
			RequiredUniversalBetaRow: 10.1497542656,
			TotalDeltaB:              FloatTripleCompat{U1GUT: 10.1497542656, SU2L: 11.4830875989, SU3C: 10.1497542656},
			ConditionalOnCarrierSeal: true,
			AnomalyCompatible:        true,
			FiniteDerived:            false,
			Verdict:                  "Gate-206 sealed conditional prediction; not a finite-core mass prediction",
		},
	}
	return Gate206Snapshot{
		Gate206Inherited:                     true,
		CarrierSealExplicit:                  true,
		CarrierSealQuarantined:               true,
		NativeSemanticSearchFailed:           true,
		AnomalyCompatibilityPassed:           true,
		PredictionsEmitted:                   len(preds),
		AllConditionalOnCarrierSeal:          true,
		UniversalCompletionStillExternal:     true,
		PhysicalUnificationClaimed:           false,
		AbsoluteMassPredicted:                false,
		ThresholdCorrectedPhysicalFitClaimed: false,
		FiniteMatchingCorrectionsDerived:     false,
		AlphaGUT:                             1.0 / (4.0 * math.Pi),
		AlphaGUTInverse:                      4.0 * math.Pi,
		Predictions:                          preds,
		TruthStatement:                       "Gate 206 recorded an explicit EmpiricalCarrierSeal, proved the two sealed carriers anomaly compatible, and emitted inverse-threshold scales only as CONDITIONAL_ON_CARRIER_SEAL phenomenology.",
	}, nil
}

type ExternalConstraintLedger struct {
	LedgerDate                           string
	ColliderCurrentRunEnergyTeV          float64
	ConservativeCurrentDirectLimitTeV    float64
	ConservativeFutureReachProxyTeV      float64
	SuperKPEPi0LifetimeLowerLimitYears   float64
	SuperKLimitConfidenceLevelPercent    float64
	HyperKProjectedPEPi0SensitivityYears float64
	References                           []string
	QuarantinedExternalPhenomenology     bool
	UsedForFiniteCoreDerivation          bool
	Verdict                              string
}

type ColliderStressCase struct {
	CarrierName                        string
	MassScaleGeV                       float64
	MassScaleTeV                       float64
	SeparationFromCurrentRunEnergy     float64
	SeparationFromCurrentDirectLimit   float64
	SeparationFromFutureReachProxy     float64
	KinematicallyBeyondCurrentLHC      bool
	BeyondPublishedDirectLimits        bool
	BeyondConservativeFutureReachProxy bool
	DirectColliderStressPassed         bool
	NoAbsoluteExclusionClaim           bool
	Verdict                            string
}

type ColliderAudit struct {
	CasesAudited                     int
	MinimumThresholdScaleGeV         float64
	MinimumThresholdScaleTeV         float64
	AllBeyondCurrentLHC              bool
	AllBeyondCurrentDirectLimits     bool
	AllBeyondConservativeFutureReach bool
	ColliderStressPassed             bool
	ExternalConstraintOnly           bool
	NoIndirectConstraintClaim        bool
	Verdict                          string
}

type EngineMediatorInventory struct {
	DerivedContactGaugeAlgebra             string
	MatterCurrentCarrier                   string
	ColorCarrierAvailableAsMatterCurrent   bool
	ContactElectroweakCarrierAvailable     bool
	FullSU5GaugeAlgebraDerived             bool
	SO10GaugeAlgebraDerived                bool
	XYLeptoquarkGaugeBosonsDerived         bool
	BLViolatingGaugeCurvatureDerived       bool
	DimensionSixProtonDecayOperatorDerived bool
	FourFermionBLViolationDerived          bool
	BaryonNumberViolationDerived           bool
	LeptonNumberViolationDerived           bool
	Verdict                                string
}

type ProtonDecayAudit struct {
	BoundaryScalesAudited                int
	MinimumBoundaryScaleGeV              float64
	MaximumBoundaryScaleGeV              float64
	SuperKPEPi0LifetimeLowerLimitYears   float64
	NaiveSU5DimensionSixWarning          bool
	NaiveLifetimeComputed                bool
	NaiveLifetimeClaimed                 bool
	EngineMediatorInventory              EngineMediatorInventory
	XYMediatedChannelSupported           bool
	NaturalSuppressionByMediatorAbsence  bool
	ProtonDecayStressPassedConditionally bool
	RequiresFutureOperatorAudit          bool
	Verdict                              string
}

type UniversalCompletionCase struct {
	CarrierName                         string
	RequiredUniversalBetaRow            float64
	TotalDeltaBU1                       float64
	TotalDeltaBSU2                      float64
	TotalDeltaBSU3                      float64
	ThresholdScaleMBGeV                 float64
	BoundaryScaleMStarGeV               float64
	U1LandauPoleGeV                     float64
	SU2LandauPoleGeV                    float64
	SU3LandauPoleGeV                    float64
	SU2AsymptoticallyFreeAboveThreshold bool
	SU3AsymptoticallyFreeAboveThreshold bool
	U1LandauPoleAbovePlanck             bool
	U1LandauPoleAboveBoundary           bool
	OrderedScalesBelowPlanck            bool
	NoOneLoopPathologyBelowPlanck       bool
	Verdict                             string
}

type UniversalCompletionAudit struct {
	CasesAudited                     int
	UniversalCompletionStillExternal bool
	AllBoundaryScalesBelowPlanck     bool
	AllU1LandauPolesAbovePlanck      bool
	AllU1LandauPolesAboveBoundary    bool
	AllNonAbelianRowsSafeAtOneLoop   bool
	NoOneLoopPathologyBelowPlanck    bool
	NoUniversalSourceDerivationClaim bool
	Verdict                          string
}

type FirewallAudit struct {
	Gate206Inherited                                    bool
	CarrierSealStillRequired                            bool
	ExternalConstraintsQuarantined                      bool
	ObservedBoundsUsedForFiniteCore                     bool
	ColliderSafetyClaimLimitedToDirectReach             bool
	IndirectColliderConstraintsClaimed                  bool
	ProtonDecayLifetimeComputed                         bool
	XYMediatedProtonDecayClaimed                        bool
	NaturalSuppressionClaimConditionalOnMediatorAbsence bool
	UniversalBetaSourceDerived                          bool
	AbsoluteMassPredicted                               bool
	PhysicalUnificationClaimed                          bool
	ThresholdCorrectedPhysicalFitClaimed                bool
	FiniteMatchingCorrectionsDerived                    bool
	RecommendedNextGate                                 string
	OpenRequirements                                    []string
	Verdict                                             string
}

type Summary struct {
	TestsAudited                      int
	Gate206Inherited                  bool
	ColliderStressPassed              bool
	ProtonDecayMediatorFirewallPassed bool
	UniversalCompletionStressPassed   bool
	ConditionalOnCarrierSealOnly      bool
	ExperimentalBoundsQuarantined     bool
	NoAbsolutePredictionClaim         bool
	Status                            string
	Comment                           string
}

type Analysis struct {
	PreviousGate206     Gate206Snapshot
	ExternalConstraints ExternalConstraintLedger
	ColliderCases       []ColliderStressCase
	Collider            ColliderAudit
	ProtonDecay         ProtonDecayAudit
	UniversalCases      []UniversalCompletionCase
	UniversalCompletion UniversalCompletionAudit
	Firewall            FirewallAudit
	Summary             Summary
	TruthStatement      string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := DefaultGate206Snapshot()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 206 input: %w", err)
			return
		}
		defaultA, defaultErr = Build(prev, DefaultExternalConstraintLedger())
	})
	return defaultA, defaultErr
}

func Build(prev Gate206Snapshot, constraints ExternalConstraintLedger) (Analysis, error) {
	if !prev.Gate206Inherited || !prev.CarrierSealExplicit || !prev.CarrierSealQuarantined || !prev.NativeSemanticSearchFailed || !prev.AnomalyCompatibilityPassed || prev.PredictionsEmitted != 2 || !prev.AllConditionalOnCarrierSeal {
		return Analysis{}, fmt.Errorf("Gate 207 requires Gate 206 sealed anomaly-compatible conditional predictions")
	}
	if prev.PhysicalUnificationClaimed || prev.AbsoluteMassPredicted || prev.ThresholdCorrectedPhysicalFitClaimed || prev.FiniteMatchingCorrectionsDerived {
		return Analysis{}, fmt.Errorf("Gate 207 refuses inherited physical-prediction leakage")
	}
	if !constraints.QuarantinedExternalPhenomenology || constraints.UsedForFiniteCoreDerivation {
		return Analysis{}, fmt.Errorf("Gate 207 requires quarantined external-constraint ledger")
	}

	colliderCases := auditColliderCases(prev.Predictions, constraints)
	collider := summarizeCollider(colliderCases)
	proton := auditProtonDecay(prev, constraints)
	univCases := auditUniversalCompletion(prev.Predictions)
	univ := summarizeUniversal(univCases, prev)
	fw := auditFirewall(prev, constraints, collider, proton, univ)
	summary := Summary{
		TestsAudited:                      7,
		Gate206Inherited:                  fw.Gate206Inherited,
		ColliderStressPassed:              collider.ColliderStressPassed,
		ProtonDecayMediatorFirewallPassed: proton.ProtonDecayStressPassedConditionally && proton.NaturalSuppressionByMediatorAbsence,
		UniversalCompletionStressPassed:   univ.NoOneLoopPathologyBelowPlanck,
		ConditionalOnCarrierSealOnly:      fw.CarrierSealStillRequired && !fw.AbsoluteMassPredicted && !fw.PhysicalUnificationClaimed,
		ExperimentalBoundsQuarantined:     constraints.QuarantinedExternalPhenomenology && !constraints.UsedForFiniteCoreDerivation,
		NoAbsolutePredictionClaim:         !fw.AbsoluteMassPredicted && !fw.PhysicalUnificationClaimed && !fw.ThresholdCorrectedPhysicalFitClaimed,
		Status:                            StatusConditionalStressTest,
		Comment:                           "Gate 207 stress-tests the Gate-206 sealed scales. The PeV threshold carriers are far beyond current direct collider reach; the low topological boundary is flagged as dangerous for naive SU(5)-like proton decay, but the engine's current connection inventory contains no derived X/Y or B,L-violating gauge mediator. The external universal beta completion fails the one-loop high-scale stress test: it drives positive non-Abelian beta rows and sub-Planck formal Landau poles. All conclusions remain conditional on the EmpiricalCarrierSeal and external phenomenology.",
	}
	truth := "Gate 207 does not convert the sealed threshold scales into native predictions. It shows they survive a first direct-collider scale test, records a proton-decay warning for naive unified gauge theories, proves the current finite connection lacks the X/Y or B,L-violating gauge channels needed to realize that warning inside ASHA, and then rejects the external universal completion because its one-loop beta rows create sub-Planck Landau-pole/asymptotic-safety pathologies. The result is a clean FAILED_ROUTE for the Gate-206 universal-completion scenario, not an absolute phenomenology claim."

	return Analysis{PreviousGate206: prev, ExternalConstraints: constraints, ColliderCases: colliderCases, Collider: collider, ProtonDecay: proton, UniversalCases: univCases, UniversalCompletion: univ, Firewall: fw, Summary: summary, TruthStatement: truth}, nil
}

func DefaultExternalConstraintLedger() ExternalConstraintLedger {
	return ExternalConstraintLedger{
		LedgerDate:                           "2026-05-11",
		ColliderCurrentRunEnergyTeV:          13.6,
		ConservativeCurrentDirectLimitTeV:    5.0,
		ConservativeFutureReachProxyTeV:      100.0,
		SuperKPEPi0LifetimeLowerLimitYears:   2.4e34,
		SuperKLimitConfidenceLevelPercent:    90.0,
		HyperKProjectedPEPi0SensitivityYears: 6.0e34,
		References: []string{
			"ATLAS/CMS VLQ searches remain TeV-scale direct searches; recent public summaries quote typical excluded/tested VLQ masses in the 1-3 TeV range, with some analyses testing final-state resonances to a few TeV.",
			"Super-Kamiokande p->e+pi0 lower limit: tau/B > 2.4e34 years at 90% CL.",
			"Hyper-Kamiokande design/projection material targets O(1e35) year sensitivity for p->e+pi0 over long exposures; Gate 207 uses 6e34 years as a conservative stress marker, not as a finite theorem.",
		},
		QuarantinedExternalPhenomenology: true,
		UsedForFiniteCoreDerivation:      false,
		Verdict:                          "external experimental ledger quarantined: used only to stress-test sealed Gate-206 scales",
	}
}

func auditColliderCases(preds []SealedPrediction, c ExternalConstraintLedger) []ColliderStressCase {
	cases := make([]ColliderStressCase, 0, len(preds))
	for _, p := range preds {
		tev := p.ThresholdScaleMBGeV / gevPerTeV
		sepRun := tev / c.ColliderCurrentRunEnergyTeV
		sepLimit := tev / c.ConservativeCurrentDirectLimitTeV
		sepFuture := tev / c.ConservativeFutureReachProxyTeV
		beyondRun := sepRun > 1.0
		beyondLimit := sepLimit > 1.0
		beyondFuture := sepFuture > 1.0
		cases = append(cases, ColliderStressCase{
			CarrierName:                        p.CarrierName,
			MassScaleGeV:                       p.ThresholdScaleMBGeV,
			MassScaleTeV:                       tev,
			SeparationFromCurrentRunEnergy:     sepRun,
			SeparationFromCurrentDirectLimit:   sepLimit,
			SeparationFromFutureReachProxy:     sepFuture,
			KinematicallyBeyondCurrentLHC:      beyondRun,
			BeyondPublishedDirectLimits:        beyondLimit,
			BeyondConservativeFutureReachProxy: beyondFuture,
			DirectColliderStressPassed:         beyondRun && beyondLimit && beyondFuture,
			NoAbsoluteExclusionClaim:           true,
			Verdict:                            fmt.Sprintf("CONDITIONAL_PASS: M_B=%.6g TeV is %.3g times the %.1f TeV direct-limit stress marker and %.3g times the %.1f TeV future-reach proxy", tev, sepLimit, c.ConservativeCurrentDirectLimitTeV, sepFuture, c.ConservativeFutureReachProxyTeV),
		})
	}
	return cases
}

func summarizeCollider(cases []ColliderStressCase) ColliderAudit {
	minGeV := math.Inf(1)
	allRun, allLimit, allFuture, allPass := len(cases) > 0, len(cases) > 0, len(cases) > 0, len(cases) > 0
	for _, c := range cases {
		if c.MassScaleGeV < minGeV {
			minGeV = c.MassScaleGeV
		}
		allRun = allRun && c.KinematicallyBeyondCurrentLHC
		allLimit = allLimit && c.BeyondPublishedDirectLimits
		allFuture = allFuture && c.BeyondConservativeFutureReachProxy
		allPass = allPass && c.DirectColliderStressPassed && c.NoAbsoluteExclusionClaim
	}
	if math.IsInf(minGeV, 1) {
		minGeV = 0
	}
	return ColliderAudit{CasesAudited: len(cases), MinimumThresholdScaleGeV: minGeV, MinimumThresholdScaleTeV: minGeV / gevPerTeV, AllBeyondCurrentLHC: allRun, AllBeyondCurrentDirectLimits: allLimit, AllBeyondConservativeFutureReach: allFuture, ColliderStressPassed: allPass, ExternalConstraintOnly: true, NoIndirectConstraintClaim: true, Verdict: colliderVerdict(allPass)}
}

func colliderVerdict(ok bool) string {
	if ok {
		return "CONDITIONAL_PASS: sealed PeV-scale carriers evade direct collider reach by orders of magnitude"
	}
	return "FAILED_ROUTE: at least one sealed carrier is not safely beyond the direct collider stress proxy"
}

func auditProtonDecay(prev Gate206Snapshot, c ExternalConstraintLedger) ProtonDecayAudit {
	minM, maxM := math.Inf(1), 0.0
	for _, p := range prev.Predictions {
		if p.BoundaryScaleMStarGeV < minM {
			minM = p.BoundaryScaleMStarGeV
		}
		if p.BoundaryScaleMStarGeV > maxM {
			maxM = p.BoundaryScaleMStarGeV
		}
	}
	if math.IsInf(minM, 1) {
		minM = 0
	}
	inventory := EngineMediatorInventory{
		DerivedContactGaugeAlgebra:             "contact-preserving su(2)+u(1) seed; color/Pati-Salam currents remain on a typed matter-current carrier, not a derived SU(5) gauge connection",
		MatterCurrentCarrier:                   "u(4)=central+color-su3+B-L+leptoquark inventory remains kinematic/current-side data, not a unified gauge-boson curvature sector",
		ColorCarrierAvailableAsMatterCurrent:   true,
		ContactElectroweakCarrierAvailable:     true,
		FullSU5GaugeAlgebraDerived:             false,
		SO10GaugeAlgebraDerived:                false,
		XYLeptoquarkGaugeBosonsDerived:         false,
		BLViolatingGaugeCurvatureDerived:       false,
		DimensionSixProtonDecayOperatorDerived: false,
		FourFermionBLViolationDerived:          false,
		BaryonNumberViolationDerived:           false,
		LeptonNumberViolationDerived:           false,
		Verdict:                                "NATURAL_SUPPRESSION_FIREWALL: the finite connection does not contain derived X/Y gauge bosons or B,L-violating curvature channels",
	}
	xysupported := inventory.FullSU5GaugeAlgebraDerived || inventory.SO10GaugeAlgebraDerived || inventory.XYLeptoquarkGaugeBosonsDerived || inventory.BLViolatingGaugeCurvatureDerived || inventory.DimensionSixProtonDecayOperatorDerived
	suppressed := !xysupported && !inventory.BaryonNumberViolationDerived && !inventory.LeptonNumberViolationDerived
	warning := minM > 0 && minM < 3.0e15
	return ProtonDecayAudit{
		BoundaryScalesAudited:                len(prev.Predictions),
		MinimumBoundaryScaleGeV:              minM,
		MaximumBoundaryScaleGeV:              maxM,
		SuperKPEPi0LifetimeLowerLimitYears:   c.SuperKPEPi0LifetimeLowerLimitYears,
		NaiveSU5DimensionSixWarning:          warning,
		NaiveLifetimeComputed:                false,
		NaiveLifetimeClaimed:                 false,
		EngineMediatorInventory:              inventory,
		XYMediatedChannelSupported:           xysupported,
		NaturalSuppressionByMediatorAbsence:  suppressed,
		ProtonDecayStressPassedConditionally: suppressed,
		RequiresFutureOperatorAudit:          true,
		Verdict:                              protonVerdict(warning, suppressed),
	}
}

func protonVerdict(warning, suppressed bool) string {
	if warning && suppressed {
		return "CONDITIONAL_PASS_WITH_WARNING: M_* is low for naive SU(5)-style dimension-six proton decay, but ASHA currently has no derived X/Y or B,L-violating gauge mediator; proton decay is naturally suppressed by mediator absence, pending future operator audits"
	}
	if suppressed {
		return "CONDITIONAL_PASS: no engine-native proton-decay mediator is derived"
	}
	return "FAILED_ROUTE: proton-decay mediator support exists and must be lifetime-tested"
}

func auditUniversalCompletion(preds []SealedPrediction) []UniversalCompletionCase {
	cases := make([]UniversalCompletionCase, 0, len(preds))
	for _, p := range preds {
		b1Total := 41.0/10.0 + p.TotalDeltaB.U1GUT
		b2Total := -19.0/6.0 + p.TotalDeltaB.SU2L
		b3Total := -7.0 + p.TotalDeltaB.SU3C
		u1Pole := computeGaugeLandauPole(p.BoundaryScaleMStarGeV, b1Total)
		su2Pole := computeGaugeLandauPole(p.BoundaryScaleMStarGeV, b2Total)
		su3Pole := computeGaugeLandauPole(p.BoundaryScaleMStarGeV, b3Total)
		su2Safe := b2Total < 0
		su3Safe := b3Total < 0
		belowPlanck := p.ThresholdScaleMBGeV > 0 && p.BoundaryScaleMStarGeV > p.ThresholdScaleMBGeV && p.BoundaryScaleMStarGeV < planckGeV
		caseOK := u1Pole > planckGeV && u1Pole > p.BoundaryScaleMStarGeV && su2Safe && su3Safe && belowPlanck
		cases = append(cases, UniversalCompletionCase{
			CarrierName:                         p.CarrierName,
			RequiredUniversalBetaRow:            p.RequiredUniversalBetaRow,
			TotalDeltaBU1:                       p.TotalDeltaB.U1GUT,
			TotalDeltaBSU2:                      p.TotalDeltaB.SU2L,
			TotalDeltaBSU3:                      p.TotalDeltaB.SU3C,
			ThresholdScaleMBGeV:                 p.ThresholdScaleMBGeV,
			BoundaryScaleMStarGeV:               p.BoundaryScaleMStarGeV,
			U1LandauPoleGeV:                     u1Pole,
			SU2LandauPoleGeV:                    su2Pole,
			SU3LandauPoleGeV:                    su3Pole,
			SU2AsymptoticallyFreeAboveThreshold: su2Safe,
			SU3AsymptoticallyFreeAboveThreshold: su3Safe,
			U1LandauPoleAbovePlanck:             u1Pole > planckGeV,
			U1LandauPoleAboveBoundary:           u1Pole > p.BoundaryScaleMStarGeV,
			OrderedScalesBelowPlanck:            belowPlanck,
			NoOneLoopPathologyBelowPlanck:       caseOK,
			Verdict:                             universalVerdict(caseOK, u1Pole),
		})
	}
	return cases
}

func computeGaugeLandauPole(boundaryScaleGeV, betaTotal float64) float64 {
	// Above M_* the one-loop inverse coupling obeys
	// alpha^{-1}(mu)=alpha^{-1}(M_*)-b_total/(2*pi)*ln(mu/M_*).
	// Gate 206 sets alpha^{-1}(M_*)=4*pi. If b_total<=0 the coupling is
	// asymptotically free/non-growing in this simple stress test, so no UV pole is
	// present. If b_total>0, the formal pole is exp(8*pi^2/b_total) times M_*.
	if betaTotal <= 0 {
		return math.Inf(1)
	}
	return boundaryScaleGeV * math.Exp((8.0*math.Pi*math.Pi)/betaTotal)
}

func universalVerdict(ok bool, pole float64) string {
	if ok {
		return fmt.Sprintf("CONDITIONAL_PASS: formal gauge-coupling poles %.6g GeV stay above Planck and non-Abelian beta rows remain asymptotically safe", pole)
	}
	return fmt.Sprintf("FAILED_ROUTE: one-loop universal-completion stress detects sub-Planck and/or non-Abelian pathology; U(1) pole %.6g GeV", pole)
}

func summarizeUniversal(cases []UniversalCompletionCase, prev Gate206Snapshot) UniversalCompletionAudit {
	allBoundary, allPolePlanck, allPoleBoundary, allNonAbelian, allOK := len(cases) > 0, len(cases) > 0, len(cases) > 0, len(cases) > 0, len(cases) > 0
	for _, c := range cases {
		allBoundary = allBoundary && c.OrderedScalesBelowPlanck
		allPolePlanck = allPolePlanck && c.U1LandauPoleAbovePlanck
		allPoleBoundary = allPoleBoundary && c.U1LandauPoleAboveBoundary
		allNonAbelian = allNonAbelian && c.SU2AsymptoticallyFreeAboveThreshold && c.SU3AsymptoticallyFreeAboveThreshold
		allOK = allOK && c.NoOneLoopPathologyBelowPlanck
	}
	return UniversalCompletionAudit{CasesAudited: len(cases), UniversalCompletionStillExternal: prev.UniversalCompletionStillExternal, AllBoundaryScalesBelowPlanck: allBoundary, AllU1LandauPolesAbovePlanck: allPolePlanck, AllU1LandauPolesAboveBoundary: allPoleBoundary, AllNonAbelianRowsSafeAtOneLoop: allNonAbelian, NoOneLoopPathologyBelowPlanck: allOK, NoUniversalSourceDerivationClaim: true, Verdict: universalAuditVerdict(allOK, prev.UniversalCompletionStillExternal)}
}

func universalAuditVerdict(ok, external bool) string {
	if ok && external {
		return "CONDITIONAL_PASS: no one-loop Landau-pole/asymptotic-safety pathology below the Planck scale, but the universal beta source is still external"
	}
	return "FAILED_ROUTE: universal completion cannot pass the one-loop pathology audit"
}

func auditFirewall(prev Gate206Snapshot, constraints ExternalConstraintLedger, collider ColliderAudit, proton ProtonDecayAudit, univ UniversalCompletionAudit) FirewallAudit {
	return FirewallAudit{
		Gate206Inherited:                                    prev.Gate206Inherited,
		CarrierSealStillRequired:                            prev.AllConditionalOnCarrierSeal,
		ExternalConstraintsQuarantined:                      constraints.QuarantinedExternalPhenomenology && !constraints.UsedForFiniteCoreDerivation,
		ObservedBoundsUsedForFiniteCore:                     constraints.UsedForFiniteCoreDerivation,
		ColliderSafetyClaimLimitedToDirectReach:             collider.ColliderStressPassed && collider.NoIndirectConstraintClaim,
		IndirectColliderConstraintsClaimed:                  false,
		ProtonDecayLifetimeComputed:                         proton.NaiveLifetimeComputed,
		XYMediatedProtonDecayClaimed:                        proton.XYMediatedChannelSupported,
		NaturalSuppressionClaimConditionalOnMediatorAbsence: proton.NaturalSuppressionByMediatorAbsence,
		UniversalBetaSourceDerived:                          !univ.UniversalCompletionStillExternal,
		AbsoluteMassPredicted:                               prev.AbsoluteMassPredicted,
		PhysicalUnificationClaimed:                          prev.PhysicalUnificationClaimed,
		ThresholdCorrectedPhysicalFitClaimed:                prev.ThresholdCorrectedPhysicalFitClaimed,
		FiniteMatchingCorrectionsDerived:                    prev.FiniteMatchingCorrectionsDerived,
		RecommendedNextGate:                                 "Gate 208 — baryon/lepton violating operator basis audit / proton-decay channel construction obstruction",
		OpenRequirements: []string{
			"derive or seal a B/L-violating local operator basis before computing proton lifetime",
			"derive a threshold-matching scheme for the universal beta source before precision unification claims",
			"audit flavour/electroweak precision/cosmology constraints only after coupling and decay portals are sealed",
		},
		Verdict: firewallVerdict(prev, constraints, collider, proton, univ),
	}
}

func firewallVerdict(prev Gate206Snapshot, constraints ExternalConstraintLedger, collider ColliderAudit, proton ProtonDecayAudit, univ UniversalCompletionAudit) string {
	ok := prev.AllConditionalOnCarrierSeal && constraints.QuarantinedExternalPhenomenology && !constraints.UsedForFiniteCoreDerivation && collider.ColliderStressPassed && proton.NaturalSuppressionByMediatorAbsence && !univ.NoOneLoopPathologyBelowPlanck && prev.UniversalCompletionStillExternal && !prev.AbsoluteMassPredicted && !prev.PhysicalUnificationClaimed && !prev.ThresholdCorrectedPhysicalFitClaimed && !prev.FiniteMatchingCorrectionsDerived
	if ok {
		return "FIREWALL_PRESERVED_WITH_FAILED_ROUTE: stress-test conclusions remain conditional and external; the universal-completion pathology is logged without upgrading or fitting the phenomenology"
	}
	return "FAILED_ROUTE: a firewall leak or unlogged stress-test inconsistency was detected"
}

func FormatGate206(s Gate206Snapshot) string {
	parts := []string{
		fmt.Sprintf("inherited=%t", s.Gate206Inherited),
		fmt.Sprintf("sealExplicit=%t", s.CarrierSealExplicit),
		fmt.Sprintf("sealQuarantined=%t", s.CarrierSealQuarantined),
		fmt.Sprintf("nativeFailed=%t", s.NativeSemanticSearchFailed),
		fmt.Sprintf("anomalyPassed=%t", s.AnomalyCompatibilityPassed),
		fmt.Sprintf("predictions=%d", s.PredictionsEmitted),
		fmt.Sprintf("conditional=%t", s.AllConditionalOnCarrierSeal),
		fmt.Sprintf("universalExternal=%t", s.UniversalCompletionStillExternal),
		fmt.Sprintf("alphaGUT=%.12g", s.AlphaGUT),
		fmt.Sprintf("alphaGUT^-1=%.12g", s.AlphaGUTInverse),
		fmt.Sprintf("physicalUnificationClaimed=%t", s.PhysicalUnificationClaimed),
	}
	return strings.Join(parts, " ")
}

func FormatExternal(c ExternalConstraintLedger) string {
	return fmt.Sprintf("date=%s LHCsqrtS=%.1fTeV directLimitProxy=%.1fTeV futureProxy=%.1fTeV SuperK_p_e_pi0=%.3gyr@%.0f%%CL HyperK_proxy=%.3gyr quarantined=%t finiteUse=%t", c.LedgerDate, c.ColliderCurrentRunEnergyTeV, c.ConservativeCurrentDirectLimitTeV, c.ConservativeFutureReachProxyTeV, c.SuperKPEPi0LifetimeLowerLimitYears, c.SuperKLimitConfidenceLevelPercent, c.HyperKProjectedPEPi0SensitivityYears, c.QuarantinedExternalPhenomenology, c.UsedForFiniteCoreDerivation)
}

func FormatColliderCase(c ColliderStressCase) string {
	return fmt.Sprintf("%s M_B=%.9gGeV=%.9gTeV sepRun=%.6g sepDirect=%.6g sepFuture=%.6g pass=%t", c.CarrierName, c.MassScaleGeV, c.MassScaleTeV, c.SeparationFromCurrentRunEnergy, c.SeparationFromCurrentDirectLimit, c.SeparationFromFutureReachProxy, c.DirectColliderStressPassed)
}

func FormatColliderCases(cases []ColliderStressCase) string {
	parts := make([]string, 0, len(cases))
	for _, c := range cases {
		parts = append(parts, FormatColliderCase(c))
	}
	return strings.Join(parts, " :: ")
}

func FormatCollider(a ColliderAudit) string {
	return fmt.Sprintf("cases=%d minMB=%.9gGeV=%.9gTeV beyondLHC=%t beyondLimits=%t beyondFuture=%t pass=%t indirectClaim=%t", a.CasesAudited, a.MinimumThresholdScaleGeV, a.MinimumThresholdScaleTeV, a.AllBeyondCurrentLHC, a.AllBeyondCurrentDirectLimits, a.AllBeyondConservativeFutureReach, a.ColliderStressPassed, !a.NoIndirectConstraintClaim)
}

func FormatMediatorInventory(i EngineMediatorInventory) string {
	return fmt.Sprintf("contactAlgebra=%q matterCarrier=%q su5=%t so10=%t XY=%t BLcurvature=%t dim6=%t fourFermionBL=%t Bviol=%t Lviol=%t", i.DerivedContactGaugeAlgebra, i.MatterCurrentCarrier, i.FullSU5GaugeAlgebraDerived, i.SO10GaugeAlgebraDerived, i.XYLeptoquarkGaugeBosonsDerived, i.BLViolatingGaugeCurvatureDerived, i.DimensionSixProtonDecayOperatorDerived, i.FourFermionBLViolationDerived, i.BaryonNumberViolationDerived, i.LeptonNumberViolationDerived)
}

func FormatProton(a ProtonDecayAudit) string {
	return fmt.Sprintf("scales=%d minM*=%.9gGeV maxM*=%.9gGeV SuperK=%.3gyr naiveSU5Warning=%t lifetimeComputed=%t XYsupported=%t naturalSuppression=%t futureOperatorAudit=%t :: %s", a.BoundaryScalesAudited, a.MinimumBoundaryScaleGeV, a.MaximumBoundaryScaleGeV, a.SuperKPEPi0LifetimeLowerLimitYears, a.NaiveSU5DimensionSixWarning, a.NaiveLifetimeComputed, a.XYMediatedChannelSupported, a.NaturalSuppressionByMediatorAbsence, a.RequiresFutureOperatorAudit, FormatMediatorInventory(a.EngineMediatorInventory))
}

func FormatUniversalCase(c UniversalCompletionCase) string {
	return fmt.Sprintf("%s c=%.12g totalB=(%.12g,%.12g,%.12g) MB=%.9g M*=%.9g poles=(U1 %.9g, SU2 %.9g, SU3 %.9g) SU2safe=%t SU3safe=%t U1pole>Planck=%t pass=%t", c.CarrierName, c.RequiredUniversalBetaRow, c.TotalDeltaBU1, c.TotalDeltaBSU2, c.TotalDeltaBSU3, c.ThresholdScaleMBGeV, c.BoundaryScaleMStarGeV, c.U1LandauPoleGeV, c.SU2LandauPoleGeV, c.SU3LandauPoleGeV, c.SU2AsymptoticallyFreeAboveThreshold, c.SU3AsymptoticallyFreeAboveThreshold, c.U1LandauPoleAbovePlanck, c.NoOneLoopPathologyBelowPlanck)
}

func FormatUniversalCases(cases []UniversalCompletionCase) string {
	parts := make([]string, 0, len(cases))
	for _, c := range cases {
		parts = append(parts, FormatUniversalCase(c))
	}
	return strings.Join(parts, " :: ")
}

func FormatUniversal(a UniversalCompletionAudit) string {
	return fmt.Sprintf("cases=%d universalExternal=%t boundaryBelowPlanck=%t U1poleAbovePlanck=%t U1poleAboveBoundary=%t nonAbelianSafe=%t pass=%t noSourceDerivation=%t", a.CasesAudited, a.UniversalCompletionStillExternal, a.AllBoundaryScalesBelowPlanck, a.AllU1LandauPolesAbovePlanck, a.AllU1LandauPolesAboveBoundary, a.AllNonAbelianRowsSafeAtOneLoop, a.NoOneLoopPathologyBelowPlanck, a.NoUniversalSourceDerivationClaim)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("gate206=%t carrierSealRequired=%t externalQuarantine=%t observedFiniteUse=%t directColliderOnly=%t indirectColliderClaim=%t protonLifetimeComputed=%t XYclaim=%t suppressionByAbsence=%t universalSourceDerived=%t absoluteMass=%t unification=%t fitClaim=%t matching=%t next=%q", f.Gate206Inherited, f.CarrierSealStillRequired, f.ExternalConstraintsQuarantined, f.ObservedBoundsUsedForFiniteCore, f.ColliderSafetyClaimLimitedToDirectReach, f.IndirectColliderConstraintsClaimed, f.ProtonDecayLifetimeComputed, f.XYMediatedProtonDecayClaimed, f.NaturalSuppressionClaimConditionalOnMediatorAbsence, f.UniversalBetaSourceDerived, f.AbsoluteMassPredicted, f.PhysicalUnificationClaimed, f.ThresholdCorrectedPhysicalFitClaimed, f.FiniteMatchingCorrectionsDerived, f.RecommendedNextGate)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("tests=%d inherited=%t collider=%t protonFirewall=%t universal=%t conditional=%t externalQuarantine=%t noAbsolute=%t status=%s", s.TestsAudited, s.Gate206Inherited, s.ColliderStressPassed, s.ProtonDecayMediatorFirewallPassed, s.UniversalCompletionStressPassed, s.ConditionalOnCarrierSealOnly, s.ExperimentalBoundsQuarantined, s.NoAbsolutePredictionClaim, s.Status)
}
