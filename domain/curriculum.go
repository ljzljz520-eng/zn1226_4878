package domain

type Lesson struct {
	ID, Title, Category string
	Required            bool
	Sequence            int
}
type Curriculum struct {
	ID, Name string
	Lessons  []Lesson
}

func (c Curriculum) Complete(done map[string]bool) bool {
	if len(c.Lessons) == 0 {
		return false
	}
	for _, l := range c.Lessons {
		if l.Required && !done[l.ID] {
			return false
		}
	}
	return true
}
func (c Curriculum) Next(done map[string]bool) (Lesson, bool) {
	for _, l := range c.Lessons {
		if !done[l.ID] {
			return l, true
		}
	}
	return Lesson{}, false
}
func DefaultCurriculum() Curriculum {
	lessons := []Lesson{
		{"safety-01", "Safety induction", "safety", true, 1}, {"privacy-01", "Privacy basics", "privacy", true, 2}, {"tools-01", "Gateway tools", "operations", true, 3}, {"quality-01", "Quality review", "quality", false, 4},
	}
	return Curriculum{ID: "core", Name: "Core training", Lessons: lessons}
}
func SortLessons(ls []Lesson) []Lesson {
	out := append([]Lesson(nil), ls...)
	for i := 0; i < len(out); i++ {
		for j := i + 1; j < len(out); j++ {
			if out[j].Sequence < out[i].Sequence {
				out[i], out[j] = out[j], out[i]
			}
		}
	}
	return out
}
