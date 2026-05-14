// Package fermionicroottracesieve implements Gate 352:
// Fermionic Effective Action / Root-Trace (Pfaffian) Sieve.
//
// Gate 352 audits the loophole left by Gate 351.  The bosonic heat-kernel
// expansion only produced even trace invariants, so Koide-type root traces were
// not native there.  This gate asks whether the fermionic Pfaffian effective
// action or the contact/Dixmier trace sector can generate the missing
// Tr(sqrt(Y†Y)) operator.  The answer is kept strict: the Pfaffian gives a
// square-root determinant/product/log-volume structure, not the Koide linear
// root-sum, unless an additional root-trace observable is postulated.
package fermionicroottracesieve

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE352-FERMIONIC-EFFECTIVE-ACTION-ROOT-TRACE-PFAFFIAN-SIEVE"

	StatusFermionicEffectiveActionFormalized = "CONDITIONAL_SUPPORT_FERMIONIC_EFFECTIVE_ACTION_FORMALIZED"
	StatusPfaffianSieveExecuted              = "CONDITIONAL_SUPPORT_PFAFFIAN_ROOT_STRUCTURE_SIEVE_EXECUTED"
	StatusRootTraceOperatorAudited           = "CONDITIONAL_SUPPORT_ROOT_TRACE_OPERATOR_AUDITED"
	StatusDixmierContactTraceAudited         = "CONDITIONAL_SUPPORT_DIXMIER_CONTACT_TRACE_AUDITED"
	StatusKoidePromotionSieveExecuted        = "CONDITIONAL_SUPPORT_KOIDE_PROMOTION_SIEVE_EXECUTED"
	StatusParameterCensusUpdated             = "CONDITIONAL_SUPPORT_PARAMETER_CENSUS_UPDATED"

	StatusTensionPfaffianIsProductNotRootSum = "CONDITIONAL_TENSION_PFAFFIAN_GENERATES_ROOT_DETERMINANT_NOT_ROOT_TRACE_SUM"
	StatusTensionDixmierFiniteMatrixZero     = "CONDITIONAL_TENSION_DIXMIER_TRACE_ON_FINITE_YUKAWA_MATRIX_NOT_NATIVE"
	StatusTensionRootTraceWouldBeNewOperator = "CONDITIONAL_TENSION_ROOT_TRACE_REQUIRES_NEW_NONLOCAL_OBSERVABLE"
	StatusTensionKoideStillEmpirical         = "CONDITIONAL_TENSION_KOIDE_ALIGNMENT_STILL_EMPIRICAL"

	StatusFailedRootTraceNotDerived        = "FAILED_ROUTE_ROOT_TRACE_OPERATOR_NOT_DERIVED"
	StatusFailedFermionicPfaffianNoKoide   = "FAILED_ROUTE_FERMIONIC_PFAFFIAN_DOES_NOT_DERIVE_KOIDE_TRACE"
	StatusFailedDixmierNoYukawaRootTrace   = "FAILED_ROUTE_DIXMIER_TRACE_DOES_NOT_LOCK_YUKAWA_ROOT_TRACE"
	StatusFailedMatrixInvariantNotPromoted = "FAILED_ROUTE_MATRIX_TRACE_INVARIANT_STILL_NOT_PROMOTED"
	StatusFailedNoAdditionalReduction      = "FAILED_ROUTE_NO_ADDITIONAL_VACUUM_PARAMETER_REDUCTION_PROVED"
	StatusFailedSevenCoordinatesNotProved  = "FAILED_ROUTE_SEVEN_VACUUM_COORDINATES_NOT_PROVED"
)

const (
	inheritedGate        = 351
	startingVacuumInputs = 15
)

type Span struct {
	AuditID       string
	InheritedGate int
	AddsFit       bool
	Purpose       string
	Verdict       string
}

type FermionicAction struct {
	Formalized      bool
	MajoranaFormula string
	DiracFormula    string
	EffectiveAction string
	Verdict         string
}

type PfaffianSieve struct {
	Executed                 bool
	Eigenvalues              []float64
	RootDeterminant          float64
	RootTrace                float64
	SumLogsHalf              float64
	KoideK                   float64
	PfaffianCanGenerateKoide bool
	Verdict                  string
}

type RootTraceAudit struct {
	Audited                 bool
	RequiredOperator        string
	BosonicEvenTraceBarrier bool
	FermionicObservable     string
	RootTraceNative         bool
	Verdict                 string
}

type DixmierAudit struct {
	Audited                     bool
	AppliesToInfiniteSpectrum   bool
	FiniteRankYukawaDixmierZero bool
	ContactVolumeF0             float64
	LocksYukawaRootTrace        bool
	Verdict                     string
}

type KoidePromotion struct {
	Executed           bool
	ChargedLeptonK     float64
	KoideTarget        float64
	EmpiricalAlignment bool
	NativePromotion    bool
	RequiredNewObject  string
	Verdict            string
}

type ParameterCensus struct {
	StartingVacuumInputs   int
	AdditionalReduction    int
	RemainingVacuumInputs  int
	SevenSealTargetReached bool
	Verdict                string
}

type Summary struct {
	Executed              bool
	RootTracePromoted     bool
	RemainingVacuumInputs int
	OneLine               string
	Status                string
}

type Analysis struct {
	Span            Span
	FermionicAction FermionicAction
	Pfaffian        PfaffianSieve
	RootTrace       RootTraceAudit
	Dixmier         DixmierAudit
	KoidePromotion  KoidePromotion
	Census          ParameterCensus
	Summary         Summary
	Truth           string
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
	span := compileSpan()
	fermionic := formalizeFermionicAction()
	pf := auditPfaffianSieve([]float64{0.51099895, 105.6583755, 1776.86})
	root := auditRootTraceOperator()
	dixmier := auditDixmierContactTrace()
	koide := auditKoidePromotion(pf)
	census := updateCensus(koide)
	summary := compileSummary(census, koide)
	truth := "Gate 352 closes the Koide loophole opened by Gate 351: the fermionic Pfaffian supplies a square-root determinant / half-log action, but not the linear root-trace sum Tr(sqrt(Y†Y)) required by Koide.  The contact/Dixmier trace also does not act as a native finite Yukawa root-trace.  Therefore no additional vacuum-parameter reduction is promoted."
	return Analysis{Span: span, FermionicAction: fermionic, Pfaffian: pf, RootTrace: root, Dixmier: dixmier, KoidePromotion: koide, Census: census, Summary: summary, Truth: truth}, nil
}

func compileSpan() Span {
	return Span{AuditID: AuditID, InheritedGate: inheritedGate, AddsFit: false, Purpose: "audit whether the fermionic Pfaffian or contact/Dixmier sector natively generates Koide-type root-trace invariants", Verdict: StatusFermionicEffectiveActionFormalized}
}

func formalizeFermionicAction() FermionicAction {
	return FermionicAction{
		Formalized:      true,
		MajoranaFormula: "Z_F,Majorana = ∫dχ exp[-1/2 χ^T A χ] = pf(A)",
		DiracFormula:    "Z_F,Dirac = det(D)",
		EffectiveAction: "Γ_F = -log pf(A) = -1/2 Tr log(A^T A) on the nonzero spectral support",
		Verdict:         StatusFermionicEffectiveActionFormalized,
	}
}

func auditPfaffianSieve(vals []float64) PfaffianSieve {
	product := 1.0
	logsum := 0.0
	for _, v := range vals {
		if v > 0 {
			product *= v
			logsum += math.Log(v)
		}
	}
	rootDet := math.Sqrt(product)
	rootTrace := 0.0
	for _, v := range vals {
		rootTrace += math.Sqrt(v)
	}
	k := koideK(vals)
	return PfaffianSieve{
		Executed:                 true,
		Eigenvalues:              append([]float64(nil), vals...),
		RootDeterminant:          rootDet,
		RootTrace:                rootTrace,
		SumLogsHalf:              0.5 * logsum,
		KoideK:                   k,
		PfaffianCanGenerateKoide: false,
		Verdict:                  strings.Join([]string{StatusPfaffianSieveExecuted, StatusTensionPfaffianIsProductNotRootSum, StatusFailedFermionicPfaffianNoKoide}, ";"),
	}
}

func auditRootTraceOperator() RootTraceAudit {
	return RootTraceAudit{
		Audited:                 true,
		RequiredOperator:        "Tr(sqrt(Y†Y)) = Tr(|Y|), a linear root-singular-value observable",
		BosonicEvenTraceBarrier: true,
		FermionicObservable:     "log pf(D) = 1/2 Tr log(D†D), sensitive to product/log-volume, not root-sum",
		RootTraceNative:         false,
		Verdict:                 strings.Join([]string{StatusRootTraceOperatorAudited, StatusTensionRootTraceWouldBeNewOperator, StatusFailedRootTraceNotDerived}, ";"),
	}
}

func auditDixmierContactTrace() DixmierAudit {
	return DixmierAudit{
		Audited:                     true,
		AppliesToInfiniteSpectrum:   true,
		FiniteRankYukawaDixmierZero: true,
		ContactVolumeF0:             7,
		LocksYukawaRootTrace:        false,
		Verdict:                     strings.Join([]string{StatusDixmierContactTraceAudited, StatusTensionDixmierFiniteMatrixZero, StatusFailedDixmierNoYukawaRootTrace}, ";"),
	}
}

func auditKoidePromotion(p PfaffianSieve) KoidePromotion {
	target := 2.0 / 3.0
	empirical := math.Abs(p.KoideK-target) < 1e-4
	return KoidePromotion{
		Executed:           true,
		ChargedLeptonK:     p.KoideK,
		KoideTarget:        target,
		EmpiricalAlignment: empirical,
		NativePromotion:    false,
		RequiredNewObject:  "a native nonlocal root-trace/absolute-Dirac observable Tr(|Y|), or an independent characteristic-polynomial theorem",
		Verdict:            strings.Join([]string{StatusKoidePromotionSieveExecuted, StatusTensionKoideStillEmpirical, StatusFailedMatrixInvariantNotPromoted}, ";"),
	}
}

func updateCensus(k KoidePromotion) ParameterCensus {
	reduction := 0
	if k.NativePromotion {
		reduction = 1
	}
	remaining := startingVacuumInputs - reduction
	return ParameterCensus{StartingVacuumInputs: startingVacuumInputs, AdditionalReduction: reduction, RemainingVacuumInputs: remaining, SevenSealTargetReached: remaining == 7, Verdict: StatusParameterCensusUpdated}
}

func compileSummary(c ParameterCensus, k KoidePromotion) Summary {
	promoted := k.NativePromotion
	status := StatusFailedNoAdditionalReduction
	line := "Gate 352 rejects Pfaffian/Dixmier promotion of Koide: root-determinant and root-trace are different invariants."
	if promoted {
		status = StatusKoidePromotionSieveExecuted
		line = "Gate 352 promotes a native root-trace invariant."
	}
	return Summary{Executed: true, RootTracePromoted: promoted, RemainingVacuumInputs: c.RemainingVacuumInputs, OneLine: line, Status: status}
}

func koideK(vals []float64) float64 {
	sum := 0.0
	rootSum := 0.0
	for _, v := range vals {
		sum += v
		if v > 0 {
			rootSum += math.Sqrt(v)
		}
	}
	if rootSum == 0 {
		return math.NaN()
	}
	return sum / (rootSum * rootSum)
}

func FormatSpan(s Span) string {
	return fmt.Sprintf("%s inherits Gate %d; addsFit=%v; verdict=%s", s.AuditID, s.InheritedGate, s.AddsFit, s.Verdict)
}
func FormatFermionicAction(f FermionicAction) string {
	return fmt.Sprintf("formalized=%v; Majorana=%s; effective=%s; verdict=%s", f.Formalized, f.MajoranaFormula, f.EffectiveAction, f.Verdict)
}
func FormatPfaffian(p PfaffianSieve) string {
	return fmt.Sprintf("rootDet=%.12g; rootTrace=%.12g; halfLog=%.12g; K=%.12g; pfaffianCanGenerateKoide=%v; verdict=%s", p.RootDeterminant, p.RootTrace, p.SumLogsHalf, p.KoideK, p.PfaffianCanGenerateKoide, p.Verdict)
}
func FormatRootTrace(r RootTraceAudit) string {
	return fmt.Sprintf("required=%s; bosonicEvenBarrier=%v; fermionicObservable=%s; native=%v; verdict=%s", r.RequiredOperator, r.BosonicEvenTraceBarrier, r.FermionicObservable, r.RootTraceNative, r.Verdict)
}
func FormatDixmier(d DixmierAudit) string {
	return fmt.Sprintf("infiniteSpectrum=%v; finiteRankZero=%v; f0=%.0f; locksRootTrace=%v; verdict=%s", d.AppliesToInfiniteSpectrum, d.FiniteRankYukawaDixmierZero, d.ContactVolumeF0, d.LocksYukawaRootTrace, d.Verdict)
}
func FormatKoidePromotion(k KoidePromotion) string {
	return fmt.Sprintf("K=%.12g target=%.12g empiricalAlignment=%v nativePromotion=%v required=%s; verdict=%s", k.ChargedLeptonK, k.KoideTarget, k.EmpiricalAlignment, k.NativePromotion, k.RequiredNewObject, k.Verdict)
}
func FormatCensus(c ParameterCensus) string {
	return fmt.Sprintf("start=%d reduction=%d remaining=%d seven=%v; verdict=%s", c.StartingVacuumInputs, c.AdditionalReduction, c.RemainingVacuumInputs, c.SevenSealTargetReached, c.Verdict)
}
func FormatSummary(s Summary) string {
	return fmt.Sprintf("executed=%v rootTracePromoted=%v remaining=%d status=%s; %s", s.Executed, s.RootTracePromoted, s.RemainingVacuumInputs, s.Status, s.OneLine)
}

func Statuses(a Analysis) []string {
	statuses := []string{
		a.Span.Verdict,
		a.FermionicAction.Verdict,
		StatusPfaffianSieveExecuted,
		StatusRootTraceOperatorAudited,
		StatusDixmierContactTraceAudited,
		StatusKoidePromotionSieveExecuted,
		a.Census.Verdict,
		StatusTensionPfaffianIsProductNotRootSum,
		StatusTensionDixmierFiniteMatrixZero,
		StatusTensionRootTraceWouldBeNewOperator,
		StatusTensionKoideStillEmpirical,
		StatusFailedRootTraceNotDerived,
		StatusFailedFermionicPfaffianNoKoide,
		StatusFailedDixmierNoYukawaRootTrace,
		StatusFailedMatrixInvariantNotPromoted,
		a.Summary.Status,
		StatusFailedSevenCoordinatesNotProved,
	}
	return statuses
}
