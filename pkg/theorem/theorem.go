package theorem

type Verifier func() Result

type Theorem struct {
	ID     string
	Name   string
	Layer  Layer
	Status Status
	Verify Verifier
}

func (t Theorem) Run() Result {
	if t.Verify == nil {
		return Result{
			ID:     t.ID,
			Name:   t.Name,
			Layer:  t.Layer,
			Status: OpenTest,
			Checks: []Check{{Name: "verifier", Passed: false, Detail: "no verifier implemented"}},
		}
	}
	return t.Verify()
}
