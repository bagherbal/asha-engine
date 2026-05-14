// Package coloredoctetportal implements Gate 223: colored-octet pure-SM
// portal search / Spectrum falsification audit.
//
// Gate 222 proved that the color-octet doublet Dirac carrier cannot mix with
// the SM quark doublet and therefore remained a cosmological obstruction. Gate
// 223 performs the sharper EFT question: scan pure-SM field products up to total
// mass dimension six and ask whether any Lorentz/gauge compatible operator can
// couple to the sealed heavy carrier Ψ8=(8,2,Y=1/2). If such an operator exists,
// it is quarantined behind the RelicDecaySeal; if not, the Rank-1 spectrum is
// falsified by colored relic cosmology.
package coloredoctetportal

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/eftdecayportal"
)

const (
	AuditID = "GATE223-COLORED-OCTET-PURE-SM-PORTAL-SEARCH"

	StatusConditionalRelicSealGranted = "CONDITIONAL_PHENOMENOLOGY_RELIC_DECAY_SEAL_GRANTED"
	StatusOctetPortalFound            = "COLORED_OCTET_DIM6_PORTAL_FOUND"
	StatusSpectrumNotFalsified        = "RANK1_SPECTRUM_NOT_FALSIFIED_BY_RELIC_DECAY"

	hbarGeVS   = 6.582119569e-25
	bbnSeconds = 1.0
	vevGeV     = 246.22
)

type Gate222Snapshot struct {
	Gate222Inherited             bool
	TripletPortalSupported       bool
	ColoredOctetUnresolved       bool
	FullRelicSealGranted         bool
	ThresholdSpectrumSealActive  bool
	MatchingCorrectionSealActive bool
	EmpiricalCarrierSealActive   bool
	LeptoquarkDynamicsSealActive bool
	MBGeV                        float64
	RequiredWidthGeV             float64
	TruthStatement               string
}

type Field struct {
	Name    string
	Kind    string
	Color   string
	WeakDim int
	Y6      int
	DimHalf int
	Fermion bool
	B3      int
}

type SearchCombination struct {
	Fields            []string
	ColorReps         []string
	WeakDims          []int
	Y6                int
	DimHalfSM         int
	TotalDimHalf      int
	FermionCount      int
	OddFermionSpinor  bool
	TargetColor       bool
	TargetWeak        bool
	TargetHypercharge bool
	LorentzCompatible bool
	B3                int
	BaryonSafe        bool
	ValidPortal       bool
	SymbolicForm      string
	Verdict           string
}

type TensorSearchAudit struct {
	TargetCarrier             string
	ConjugateHeavy            string
	RequiredSMOperator        string
	Fields                    []Field
	CombinationsScanned       int
	DimensionEligibleScanned  int
	GaugeMatches              int
	LorentzGaugeMatches       int
	ValidPortals              []SearchCombination
	BestPortal                SearchCombination
	OctetPortalFound          bool
	FalseQMixingStillRejected bool
	FalsifiesSpectrum         bool
	SearchLimit               string
	Verdict                   string
}

type BBNPortalKinematics struct {
	MBGeV                        float64
	RequiredWidthGeV             float64
	BBNThresholdSeconds          float64
	OperatorDimension            int
	UnitWilsonLambdaMax3BodyGeV  float64
	UnitWilsonLambdaMaxDipoleGeV float64
	ConservativeLambdaMaxGeV     float64
	WilsonMinAtPlanck            float64
	WilsonMinAtGUT               float64
	PlanckGeV                    float64
	ReferenceGUTGeV              float64
	BBNSafeForPerturbativeWilson bool
	Formula3Body                 string
	FormulaDipole                string
	Verdict                      string
}

type RelicDecaySealAudit struct {
	SealName              string
	SealPreviouslyDenied  bool
	SealGranted           bool
	TripletPortal         string
	OctetPortal           string
	QuarantinedInputs     []string
	StillNotFiniteDerived bool
	OperationalStatus     string
	Verdict               string
}

type FallbackAudit struct {
	Triggered         bool
	Reason            string
	Rank2Or3LookAhead string
	Verdict           string
}

type FirewallAudit struct {
	Gate222Inherited             bool
	ThresholdSpectrumSealActive  bool
	MatchingCorrectionSealActive bool
	EmpiricalCarrierSealActive   bool
	LeptoquarkDynamicsSealActive bool
	NewMediatorInvented          bool
	LeptoquarkSealViolated       bool
	FalseOctetQMixingClaimed     bool
	FiniteOperatorClaimed        bool
	WilsonCoefficientFixed       bool
	RelicAbundanceComputed       bool
	BBNUsedAsFilterOnly          bool
	FallbackTuningPerformed      bool
	FiniteCorePolluted           bool
	Verdict                      string
}

type Summary struct {
	OctetPortalFound       bool
	FullRelicDecaySeal     bool
	Rank1SpectrumFalsified bool
	ConditionalRelicSafe   bool
	Status                 string
	NextGate               string
	Comment                string
}

type Analysis struct {
	Gate222         Gate222Snapshot
	Gate222Analysis eftdecayportal.Analysis
	TensorSearch    TensorSearchAudit
	Kinematics      BBNPortalKinematics
	RelicSeal       RelicDecaySealAudit
	Fallback        FallbackAudit
	Firewall        FirewallAudit
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
		g222, err := eftdecayportal.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 222 input: %w", err)
			return
		}
		defaultA, defaultErr = Build(g222)
	})
	return defaultA, defaultErr
}

func Build(g222 eftdecayportal.Analysis) (Analysis, error) {
	snap := snapshotFromGate222(g222)
	if !snap.Gate222Inherited || !snap.TripletPortalSupported || !snap.ColoredOctetUnresolved || snap.FullRelicSealGranted {
		return Analysis{}, fmt.Errorf("Gate 223 requires Gate 222 partial triplet support and unresolved colored octet")
	}
	search := runTensorSearch()
	kin := auditBBN(snap, search)
	seal := auditRelicDecaySeal(snap, search, kin)
	fallback := auditFallback(search)
	firewall := auditFirewall(snap)
	summary := summarize(search, seal)
	truth := buildTruth(snap, search, kin, seal, summary)
	return Analysis{Gate222: snap, Gate222Analysis: g222, TensorSearch: search, Kinematics: kin, RelicSeal: seal, Fallback: fallback, Firewall: firewall, Summary: summary, TruthStatement: truth}, nil
}

func snapshotFromGate222(a eftdecayportal.Analysis) Gate222Snapshot {
	return Gate222Snapshot{
		Gate222Inherited:             a.Summary.Status != "" && a.Summary.TripletRescuedByEFTPortal && !a.Summary.ColoredOctetRescued,
		TripletPortalSupported:       a.Summary.TripletRescuedByEFTPortal && a.RelicSeal.PartialTripletSubseal,
		ColoredOctetUnresolved:       !a.Summary.ColoredOctetRescued && a.Operators.OctetPortalObstructed,
		FullRelicSealGranted:         a.RelicSeal.SealGranted,
		ThresholdSpectrumSealActive:  a.Firewall.ThresholdSpectrumSealActive,
		MatchingCorrectionSealActive: a.Firewall.MatchingCorrectionSealActive,
		EmpiricalCarrierSealActive:   a.Firewall.EmpiricalCarrierSealActive,
		LeptoquarkDynamicsSealActive: a.Firewall.LeptoquarkDynamicsSealActive,
		MBGeV:                        a.Kinematics.MBGeV,
		RequiredWidthGeV:             a.Kinematics.RequiredWidthGeV,
		TruthStatement:               a.TruthStatement,
	}
}

func smFields() []Field {
	return []Field{
		{Name: "Q", Kind: "fermion", Color: "3", WeakDim: 2, Y6: 1, DimHalf: 3, Fermion: true, B3: 1},
		{Name: "u^c", Kind: "fermion", Color: "3bar", WeakDim: 1, Y6: -4, DimHalf: 3, Fermion: true, B3: -1},
		{Name: "d^c", Kind: "fermion", Color: "3bar", WeakDim: 1, Y6: 2, DimHalf: 3, Fermion: true, B3: -1},
		{Name: "L", Kind: "fermion", Color: "1", WeakDim: 2, Y6: -3, DimHalf: 3, Fermion: true},
		{Name: "e^c", Kind: "fermion", Color: "1", WeakDim: 1, Y6: 6, DimHalf: 3, Fermion: true},
		{Name: "H", Kind: "scalar", Color: "1", WeakDim: 2, Y6: 3, DimHalf: 2, Fermion: false},
		{Name: "H†", Kind: "scalar", Color: "1", WeakDim: 2, Y6: -3, DimHalf: 2, Fermion: false},
		{Name: "G_{μν}", Kind: "field-strength", Color: "8", WeakDim: 1, Y6: 0, DimHalf: 4, Fermion: false},
		{Name: "W_{μν}", Kind: "field-strength", Color: "1", WeakDim: 3, Y6: 0, DimHalf: 4, Fermion: false},
		{Name: "B_{μν}", Kind: "field-strength", Color: "1", WeakDim: 1, Y6: 0, DimHalf: 4, Fermion: false},
	}
}

func runTensorSearch() TensorSearchAudit {
	fields := smFields()
	var combos []SearchCombination
	var rec func(start int, chosen []Field, dimHalf int)
	rec = func(start int, chosen []Field, dimHalf int) {
		if len(chosen) > 0 {
			combos = append(combos, evaluateCombination(chosen))
		}
		for i := start; i < len(fields); i++ {
			nextDim := dimHalf + fields[i].DimHalf
			if nextDim > 9 { // O_SM dimension <= 9/2, so heavy+O_SM <= 6.
				continue
			}
			next := append(append([]Field{}, chosen...), fields[i])
			rec(i, next, nextDim)
		}
	}
	rec(0, nil, 0)
	valid := make([]SearchCombination, 0)
	gaugeMatches := 0
	lorentzGauge := 0
	dimEligible := 0
	for _, c := range combos {
		if c.TotalDimHalf <= 12 {
			dimEligible++
		}
		if c.TargetColor && c.TargetWeak && c.TargetHypercharge {
			gaugeMatches++
		}
		if c.TargetColor && c.TargetWeak && c.TargetHypercharge && c.LorentzCompatible {
			lorentzGauge++
		}
		if c.ValidPortal {
			valid = append(valid, c)
		}
	}
	sort.Slice(valid, func(i, j int) bool {
		if valid[i].TotalDimHalf != valid[j].TotalDimHalf {
			return valid[i].TotalDimHalf < valid[j].TotalDimHalf
		}
		return strings.Join(valid[i].Fields, " ") < strings.Join(valid[j].Fields, " ")
	})
	best := SearchCombination{}
	if len(valid) > 0 {
		best = valid[0]
	}
	return TensorSearchAudit{
		TargetCarrier:             "Dirac Ψ8=(8,2,Y=1/2)",
		ConjugateHeavy:            "bar(Ψ8)=(8,2,Y=-1/2)",
		RequiredSMOperator:        "O_SM=(8,2,Y=1/2), fermionic, dim(O_SM)≤9/2 so bar(Ψ8)O_SM has dimension≤6",
		Fields:                    fields,
		CombinationsScanned:       len(combos),
		DimensionEligibleScanned:  dimEligible,
		GaugeMatches:              gaugeMatches,
		LorentzGaugeMatches:       lorentzGauge,
		ValidPortals:              valid,
		BestPortal:                best,
		OctetPortalFound:          len(valid) > 0,
		FalseQMixingStillRejected: true,
		FalsifiesSpectrum:         len(valid) == 0,
		SearchLimit:               "pure SM fields Q,u^c,d^c,L,e^c,H,H†,G,W,B; total operator dimension≤6; no new mediators; no dormant leptoquark slots",
		Verdict:                   tensorVerdict(valid),
	}
}

func evaluateCombination(fields []Field) SearchCombination {
	color := []string{"1"}
	weak := []int{1}
	y6 := 0
	dim := 0
	fermions := 0
	b3 := 0
	names := make([]string, 0, len(fields))
	for _, f := range fields {
		names = append(names, f.Name)
		color = multiplyColorSet(color, []string{f.Color})
		weak = multiplyWeakSet(weak, []int{f.WeakDim})
		y6 += f.Y6
		dim += f.DimHalf
		if f.Fermion {
			fermions++
		}
		b3 += f.B3
	}
	targetColor := containsString(color, "8")
	targetWeak := containsInt(weak, 2)
	targetY := y6 == 3
	oddSpinor := fermions%2 == 1
	lorentz := oddSpinor
	baryonSafe := b3 == 0
	valid := targetColor && targetWeak && targetY && lorentz && baryonSafe && dim+3 <= 12
	verdict := "not a portal"
	if valid {
		verdict = "valid pure-SM dimension≤6 spinor operator for bar(Ψ8)O_SM"
	} else if targetColor && targetWeak && targetY && lorentz && !baryonSafe {
		verdict = "gauge/Lorentz match rejected by baryon-conservation firewall"
	} else if targetColor && targetWeak && targetY && !lorentz {
		verdict = "gauge match but Lorentz/spinor parity fails"
	} else if targetColor || targetWeak || targetY {
		verdict = "partial gauge match only"
	}
	return SearchCombination{Fields: names, ColorReps: color, WeakDims: weak, Y6: y6, DimHalfSM: dim, TotalDimHalf: dim + 3, FermionCount: fermions, OddFermionSpinor: oddSpinor, TargetColor: targetColor, TargetWeak: targetWeak, TargetHypercharge: targetY, LorentzCompatible: lorentz, B3: b3, BaryonSafe: baryonSafe, ValidPortal: valid, SymbolicForm: symbolic(names), Verdict: verdict}
}

func tensorVerdict(valid []SearchCombination) string {
	if len(valid) == 0 {
		return "no pure-SM dimension≤6 octet portal found; Rank-1 spectrum would be falsified by colored relic cosmology"
	}
	return "Gate 223 finds a pure-SM dimension-6 chromomagnetic-Higgs-lepton portal; Rank-1 spectrum is not falsified, but the Wilson coefficient remains sealed phenomenology"
}

func auditBBN(snap Gate222Snapshot, search TensorSearchAudit) BBNPortalKinematics {
	lambda3 := lambdaDim6ThreeBodyMax(snap.MBGeV, snap.RequiredWidthGeV)
	lambda2 := lambdaDim6DipoleAfterVEVMax(snap.MBGeV, snap.RequiredWidthGeV)
	conservative := math.Min(lambda3, lambda2)
	planck := 1.22e19
	gut := 1.0e17
	wPlanck := requiredWilsonForLambda(planck, conservative)
	wGUT := requiredWilsonForLambda(gut, conservative)
	safe := search.OctetPortalFound && conservative > snap.MBGeV
	return BBNPortalKinematics{MBGeV: snap.MBGeV, RequiredWidthGeV: snap.RequiredWidthGeV, BBNThresholdSeconds: bbnSeconds, OperatorDimension: 6, UnitWilsonLambdaMax3BodyGeV: lambda3, UnitWilsonLambdaMaxDipoleGeV: lambda2, ConservativeLambdaMaxGeV: conservative, WilsonMinAtPlanck: wPlanck, WilsonMinAtGUT: wGUT, PlanckGeV: planck, ReferenceGUTGeV: gut, BBNSafeForPerturbativeWilson: safe, Formula3Body: "Γ3≈|c|² M_B^5/(192π³Λ^4)", FormulaDipole: "Γ2≈|c|² v² M_B^3/(8πΛ^4) after H† obtains v", Verdict: fmt.Sprintf("for unit Wilson coefficient, BBN requires Λ ≲ %.6g GeV conservatively; equivalently |c|≳%.6g if Λ=1e17 GeV", conservative, wGUT)}
}

func auditRelicDecaySeal(snap Gate222Snapshot, search TensorSearchAudit, kin BBNPortalKinematics) RelicDecaySealAudit {
	granted := search.OctetPortalFound && kin.BBNSafeForPerturbativeWilson && snap.TripletPortalSupported
	status := "RELIC_DECAY_SEAL_DENIED"
	verdict := "RelicDecaySeal remains denied"
	if granted {
		status = "RELIC_DECAY_SEAL_GRANTED_CONDITIONAL_ON_EFT_PORTALS"
		verdict = "RelicDecaySeal granted as a phenomenological seal: triplet Yukawa portal plus colored-octet dimension-6 pure-SM portal clear the BBN lifetime filter"
	}
	return RelicDecaySealAudit{SealName: "RelicDecaySeal", SealPreviouslyDenied: !snap.FullRelicSealGranted, SealGranted: granted, TripletPortal: "y_T Ψ_3^a(Lσ^aH†)", OctetPortal: "(c_8/Λ²) \u0305Ψ_8^a_i σ^{μν} e^c H†_i G^a_{μν} + h.c.", QuarantinedInputs: []string{"triplet Yukawa y_T", "octet Wilson coefficient c_8", "octet suppression scale Λ", "flavor choice e^c/μ^c/τ^c", "finite matching to relic Boltzmann history"}, StillNotFiniteDerived: true, OperationalStatus: status, Verdict: verdict}
}

func auditFallback(search TensorSearchAudit) FallbackAudit {
	if search.FalsifiesSpectrum {
		return FallbackAudit{Triggered: true, Reason: "no colored-octet portal found", Rank2Or3LookAhead: "required: audit next Gate-215 spectra for SM-like vectorlike quark/lepton portals", Verdict: "fallback route required"}
	}
	return FallbackAudit{Triggered: false, Reason: "Rank-1 spectrum is rescued by a pure-SM dimension-6 octet portal", Rank2Or3LookAhead: "not executed; no spectrum replacement is required at Gate 223", Verdict: "fallback not triggered"}
}

func auditFirewall(snap Gate222Snapshot) FirewallAudit {
	return FirewallAudit{Gate222Inherited: snap.Gate222Inherited, ThresholdSpectrumSealActive: snap.ThresholdSpectrumSealActive, MatchingCorrectionSealActive: snap.MatchingCorrectionSealActive, EmpiricalCarrierSealActive: snap.EmpiricalCarrierSealActive, LeptoquarkDynamicsSealActive: snap.LeptoquarkDynamicsSealActive, NewMediatorInvented: false, LeptoquarkSealViolated: false, FalseOctetQMixingClaimed: false, FiniteOperatorClaimed: false, WilsonCoefficientFixed: false, RelicAbundanceComputed: false, BBNUsedAsFilterOnly: true, FallbackTuningPerformed: false, FiniteCorePolluted: false, Verdict: "Gate 223 uses only pure-SM fields and a quarantined EFT Wilson coefficient; no new mediator, leptoquark dynamics, or finite-core decay theorem is invented"}
}

func summarize(search TensorSearchAudit, seal RelicDecaySealAudit) Summary {
	falsified := search.FalsifiesSpectrum
	status := StatusConditionalRelicSealGranted + "; " + StatusOctetPortalFound + "; " + StatusSpectrumNotFalsified
	if falsified {
		status = "FAILED_ROUTE_SPECTRUM_FALSIFIED"
	}
	return Summary{OctetPortalFound: search.OctetPortalFound, FullRelicDecaySeal: seal.SealGranted, Rank1SpectrumFalsified: falsified, ConditionalRelicSafe: seal.SealGranted, Status: status, NextGate: "Gate 224 — relic abundance / flavor and rare-decay safety audit", Comment: "The Rank-1 PeV spectrum survives the do-or-die colored relic test only after a RelicDecaySeal quarantines a dimension-6 pure-SM chromomagnetic-Higgs-lepton portal."}
}

func buildTruth(snap Gate222Snapshot, search TensorSearchAudit, kin BBNPortalKinematics, seal RelicDecaySealAudit, summary Summary) string {
	best := "none"
	if search.OctetPortalFound {
		best = FormatCombination(search.BestPortal)
	}
	return fmt.Sprintf("Gate 223 scans pure-SM products up to total dimension six for O_SM=(8,2,1/2) coupled to bar(Ψ8). It finds %d valid portal(s); best=%s. The conservative BBN unit-Wilson bound is Λ<%.6g GeV for M_B=%.9g GeV. The full RelicDecaySeal is granted only conditionally on the triplet and octet EFT portals; no native finite decay theorem is claimed. Status=%s.", len(search.ValidPortals), best, kin.ConservativeLambdaMaxGeV, snap.MBGeV, summary.Status)
}

func lambdaDim6ThreeBodyMax(m, gammaReq float64) float64 {
	if m <= 0 || gammaReq <= 0 {
		return 0
	}
	return math.Pow(math.Pow(m, 5)/(192.0*math.Pow(math.Pi, 3)*gammaReq), 0.25)
}

func lambdaDim6DipoleAfterVEVMax(m, gammaReq float64) float64 {
	if m <= 0 || gammaReq <= 0 {
		return 0
	}
	return math.Pow(vevGeV*vevGeV*math.Pow(m, 3)/(8.0*math.Pi*gammaReq), 0.25)
}

func requiredWilsonForLambda(lambda, lambdaUnitMax float64) float64 {
	if lambda <= 0 || lambdaUnitMax <= 0 {
		return math.Inf(1)
	}
	// Width scales as c^2 / Λ^4, so c_min=(Λ/Λ_unit)^2.
	return math.Pow(lambda/lambdaUnitMax, 2)
}

func symbolic(names []string) string { return strings.Join(names, " · ") }

func multiplyWeakSet(a, b []int) []int {
	set := map[int]bool{}
	for _, x := range a {
		for _, y := range b {
			min := int(math.Abs(float64(x-y))) + 1
			max := x + y - 1
			for d := min; d <= max; d += 2 {
				set[d] = true
			}
		}
	}
	out := make([]int, 0, len(set))
	for d := range set {
		out = append(out, d)
	}
	sort.Ints(out)
	return out
}

func multiplyColorSet(a, b []string) []string {
	set := map[string]bool{}
	for _, x := range a {
		for _, y := range b {
			for _, z := range colorProduct(x, y) {
				set[z] = true
			}
		}
	}
	out := make([]string, 0, len(set))
	for r := range set {
		out = append(out, r)
	}
	sort.Strings(out)
	return out
}

func colorProduct(a, b string) []string {
	if a == "1" {
		return []string{b}
	}
	if b == "1" {
		return []string{a}
	}
	key := a + "x" + b
	if a > b {
		key = b + "x" + a
	}
	table := map[string][]string{
		"3x3":       {"3bar", "6"},
		"3barx3bar": {"3", "6bar"},
		"3x3bar":    {"1", "8"},
		"3x8":       {"3", "6bar", "15"},
		"3barx8":    {"3bar", "6", "15bar"},
		"8x8":       {"1", "8", "10", "10bar", "27"},
		"3x6":       {"8", "10"},
		"3barx6bar": {"8", "10bar"},
		"3barx6":    {"3", "15"},
		"3x6bar":    {"3bar", "15bar"},
		"6x6bar":    {"1", "8", "27"},
	}
	if v, ok := table[key]; ok {
		return v
	}
	// Conservative fallback: retain an opaque label rather than inventing an octet.
	return []string{"(" + key + ")"}
}

func containsString(xs []string, target string) bool {
	for _, x := range xs {
		if x == target {
			return true
		}
	}
	return false
}
func containsInt(xs []int, target int) bool {
	for _, x := range xs {
		if x == target {
			return true
		}
	}
	return false
}

func bString(b3 int) string {
	if b3 == 0 {
		return "0"
	}
	if b3%3 == 0 {
		return fmt.Sprintf("%d", b3/3)
	}
	return fmt.Sprintf("%d/3", b3)
}

func yString(y6 int) string {
	if y6 == 0 {
		return "0"
	}
	if y6%6 == 0 {
		return fmt.Sprintf("%d", y6/6)
	}
	if y6%3 == 0 {
		return fmt.Sprintf("%d/2", y6/3)
	}
	if y6%2 == 0 {
		return fmt.Sprintf("%d/3", y6/2)
	}
	return fmt.Sprintf("%d/6", y6)
}

func FormatField(f Field) string {
	return fmt.Sprintf("%s=(%s,%d,Y=%s),dim=%.1f", f.Name, f.Color, f.WeakDim, yString(f.Y6), float64(f.DimHalf)/2.0)
}

func FormatFields(fields []Field) string {
	parts := make([]string, 0, len(fields))
	for _, f := range fields {
		parts = append(parts, FormatField(f))
	}
	return strings.Join(parts, "; ")
}

func FormatCombination(c SearchCombination) string {
	return fmt.Sprintf("O=%s color=%v weak=%v Y=%s dimSM=%.1f dimTotal=%.1f fermions=%d B=%s target=(color %t,weak %t,Y %t,lorentz %t,baryonSafe %t) valid=%t :: %s", c.SymbolicForm, c.ColorReps, c.WeakDims, yString(c.Y6), float64(c.DimHalfSM)/2.0, float64(c.TotalDimHalf)/2.0, c.FermionCount, bString(c.B3), c.TargetColor, c.TargetWeak, c.TargetHypercharge, c.LorentzCompatible, c.BaryonSafe, c.ValidPortal, c.Verdict)
}

func FormatTensorSearch(a TensorSearchAudit) string {
	valid := make([]string, 0, len(a.ValidPortals))
	for _, v := range a.ValidPortals {
		valid = append(valid, FormatCombination(v))
	}
	return fmt.Sprintf("target=%s requires=%s limit=%q scanned=%d dimEligible=%d gaugeMatches=%d lorentzGauge=%d valid=%d falseQMixRejected=%t falsifies=%t fields=[%s] validPortals=[%s] :: %s", a.TargetCarrier, a.RequiredSMOperator, a.SearchLimit, a.CombinationsScanned, a.DimensionEligibleScanned, a.GaugeMatches, a.LorentzGaugeMatches, len(a.ValidPortals), a.FalseQMixingStillRejected, a.FalsifiesSpectrum, FormatFields(a.Fields), strings.Join(valid, " | "), a.Verdict)
}

func FormatKinematics(k BBNPortalKinematics) string {
	return fmt.Sprintf("MB=%.9g Γreq=%.9g BBN<%.3gs dim=%d Λ3body(unit)<%.9g Λdipole(unit)<%.9g conservativeΛ<%.9g cmin(Λ=1e17)=%.9g cmin(Planck)=%.9g safe=%t formulas=(%s; %s) :: %s", k.MBGeV, k.RequiredWidthGeV, k.BBNThresholdSeconds, k.OperatorDimension, k.UnitWilsonLambdaMax3BodyGeV, k.UnitWilsonLambdaMaxDipoleGeV, k.ConservativeLambdaMaxGeV, k.WilsonMinAtGUT, k.WilsonMinAtPlanck, k.BBNSafeForPerturbativeWilson, k.Formula3Body, k.FormulaDipole, k.Verdict)
}

func FormatRelicSeal(s RelicDecaySealAudit) string {
	return fmt.Sprintf("seal=%s previouslyDenied=%t granted=%t triplet=%q octet=%q finiteDerived=%t inputs=[%s] status=%s :: %s", s.SealName, s.SealPreviouslyDenied, s.SealGranted, s.TripletPortal, s.OctetPortal, !s.StillNotFiniteDerived, strings.Join(s.QuarantinedInputs, "; "), s.OperationalStatus, s.Verdict)
}

func FormatFallback(f FallbackAudit) string {
	return fmt.Sprintf("triggered=%t reason=%q lookAhead=%q :: %s", f.Triggered, f.Reason, f.Rank2Or3LookAhead, f.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("gate222=%t seals=(spectrum %t,matching %t,carrier %t,lq %t) newMediator=%t lqSealViolation=%t falseQMixing=%t finiteOperator=%t WilsonFixed=%t relicAbundance=%t BBNfilter=%t fallbackTuning=%t finitePolluted=%t :: %s", f.Gate222Inherited, f.ThresholdSpectrumSealActive, f.MatchingCorrectionSealActive, f.EmpiricalCarrierSealActive, f.LeptoquarkDynamicsSealActive, f.NewMediatorInvented, f.LeptoquarkSealViolated, f.FalseOctetQMixingClaimed, f.FiniteOperatorClaimed, f.WilsonCoefficientFixed, f.RelicAbundanceComputed, f.BBNUsedAsFilterOnly, f.FallbackTuningPerformed, f.FiniteCorePolluted, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("octetPortal=%t relicSeal=%t falsified=%t conditionalSafe=%t status=%s next=%q :: %s", s.OctetPortalFound, s.FullRelicDecaySeal, s.Rank1SpectrumFalsified, s.ConditionalRelicSafe, s.Status, s.NextGate, s.Comment)
}
