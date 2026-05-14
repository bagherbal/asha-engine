package generationaddressfunctor

import "testing"

func TestNativeCandidatesRemainCentral(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Candidates.NativeNoncentralCount != 0 || a.Candidates.CentralNativeCount < 3 {
		t.Fatalf("bad candidate audit:\n%s", FormatCandidateAudit(a.Candidates))
	}
	for _, name := range []string{"identity generation broadcast", "Morita edge uniform incidence", "inner-fluctuation one-form uniform support"} {
		c := findCandidate(a.Candidates.Candidates, name)
		if !c.Native || !c.Central || c.NonCentral {
			t.Fatalf("native candidate should be central: %s\n%s", name, FormatCandidate(c))
		}
	}
}

func TestSealedCandidatesAreNotPromoted(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"abstract triality branch cycle", "protected contact anisotropy spurion", "Fock number ladder N"} {
		c := findCandidate(a.Candidates.Candidates, name)
		if c.Native || !c.Sealed || !c.NonCentral {
			t.Fatalf("sealed noncentral candidate not classified correctly: %s\n%s", name, FormatCandidate(c))
		}
	}
}

func TestNumberOperatorIsHierarchyOnly(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Number.Native || !a.Number.SealedExternalExtension || !a.Number.CircularIfUsedAsSolution || !a.Number.BreaksExactTriality || !a.Number.ProducesHierarchy || a.Number.ProducesMixing {
		t.Fatalf("bad number operator audit:\n%s", FormatNumber(a.Number))
	}
}

func TestNoNativeNoncommutingTextureCapacity(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.TextureCapacity.NativeNoncommutingPairs != 0 || a.TextureCapacity.CKMCapacityNative || a.TextureCapacity.MaxNativeCommutatorNorm >= eps {
		t.Fatalf("native texture capacity should be absent:\n%s", FormatTexture(a.TextureCapacity))
	}
	if a.TextureCapacity.SealedNoncommutingPairs == 0 || a.TextureCapacity.MaxSealedCommutatorNorm <= eps {
		t.Fatalf("sealed stress-test noncommutation should be visible but quarantined:\n%s", FormatTexture(a.TextureCapacity))
	}
}

func TestModuliFirewallPreserved(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Moduli.StartingChargedDim != 13 || a.Moduli.NativeReductionBelow13 || a.Moduli.BestNativeDim != 13 {
		t.Fatalf("bad moduli audit:\n%s", FormatModuli(a.Moduli))
	}
}

func TestStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	statuses := Statuses(a)
	required := []string{
		StatusGate393Inherited,
		StatusFailedGenerationAddressCentral,
		StatusFailedDiagonalOnlyNoCKM,
		StatusFailedCircularTauOrNInsertion,
		StatusFailedNoNativeNoncommutingPair,
		StatusFirewallPreserved13Moduli,
	}
	for _, req := range required {
		found := false
		for _, got := range statuses {
			if got == req {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("missing status %s in %v", req, statuses)
		}
	}
}

func TestTheoremPassesAsFailedRouteAudit(t *testing.T) {
	res := NativeGenerationAddressFunctorTrialityMoritaEdgeIncidenceTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem audit failed:\n%s", res.Details())
	}
}

func TestRenderMarkdown(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderMarkdown(a)
	if len(md) < 1000 || !contains(md, "Gate 394") || !contains(md, "FIREWALL_PRESERVED_13_MODULI") {
		t.Fatalf("markdown audit looks incomplete")
	}
}

func contains(s, sub string) bool {
	return len(sub) == 0 || (len(s) >= len(sub) && index(s, sub) >= 0)
}

func index(s, sub string) int {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}
