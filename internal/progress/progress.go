package progress

import "time"

// Progress holds all player progress data.
type Progress struct {
	PlayerName string    `json:"player_name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	Tutorials map[string]TutorialProgress `json:"tutorials"` // key: "1-1", "1-2", etc.
	Missions  map[string]MissionProgress  `json:"missions"`  // key: "1", "2", etc.

	TotalKeystrokes int `json:"total_keystrokes"`
}

// TutorialProgress tracks completion of a single tutorial substep.
type TutorialProgress struct {
	Completed   bool      `json:"completed"`
	CompletedAt time.Time `json:"completed_at,omitempty"`
	BestTime    float64   `json:"best_time,omitempty"`    // seconds
	Keystrokes  int       `json:"keystrokes,omitempty"`
	Attempts    int       `json:"attempts"`
}

// MissionProgress tracks completion of a single mission.
type MissionProgress struct {
	Completed   bool      `json:"completed"`
	CompletedAt time.Time `json:"completed_at,omitempty"`
	BestTime    float64   `json:"best_time,omitempty"`
	Keystrokes  int       `json:"keystrokes,omitempty"`
	Stars       int       `json:"stars"` // 1-3
}

// NewProgress creates a new empty progress.
func NewProgress() *Progress {
	now := time.Now()
	return &Progress{
		CreatedAt: now,
		UpdatedAt: now,
		Tutorials: make(map[string]TutorialProgress),
		Missions:  make(map[string]MissionProgress),
	}
}

// CompletedTutorialCount returns the number of completed tutorials.
func (p *Progress) CompletedTutorialCount() int {
	count := 0
	for _, t := range p.Tutorials {
		if t.Completed {
			count++
		}
	}
	return count
}

// CompletedMissionCount returns the number of completed missions.
func (p *Progress) CompletedMissionCount() int {
	count := 0
	for _, m := range p.Missions {
		if m.Completed {
			count++
		}
	}
	return count
}

// CurrentRank returns the player's current rank.
func (p *Progress) CurrentRank() Rank {
	return CalculateRank(p.CompletedTutorialCount(), p.CompletedMissionCount())
}

// CompleteTutorial marks a tutorial substep as completed.
func (p *Progress) CompleteTutorial(id string, timeSeconds float64, keystrokes int) {
	tp := p.Tutorials[id]
	tp.Attempts++
	if !tp.Completed {
		tp.Completed = true
		tp.CompletedAt = time.Now()
	}
	if tp.BestTime == 0 || timeSeconds < tp.BestTime {
		tp.BestTime = timeSeconds
	}
	if tp.Keystrokes == 0 || keystrokes < tp.Keystrokes {
		tp.Keystrokes = keystrokes
	}
	p.Tutorials[id] = tp
	p.UpdatedAt = time.Now()
}
