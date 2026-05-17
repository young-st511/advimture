package progressadapter

import (
	"fmt"
	"strings"
	"time"

	"github.com/young-st511/advimture/internal/progress"
	exerciseruntime "github.com/young-st511/advimture/internal/runtime"
	"github.com/young-st511/advimture/internal/scenario"
)

type MissionCompletion struct {
	MissionID  string
	Completed  bool
	Grade      string
	Keystrokes int
	TimeMs     int64
	Attempts   int
}

func MissionCompletionFromScenario(missionID string, state scenario.State, elapsed time.Duration) (MissionCompletion, error) {
	if strings.TrimSpace(missionID) == "" {
		return MissionCompletion{}, fmt.Errorf("mission id is required")
	}
	if state.Status != exerciseruntime.StatusSucceeded {
		return MissionCompletion{}, fmt.Errorf("scenario is not completed")
	}
	if state.Score == nil || !state.Score.Passed {
		return MissionCompletion{}, fmt.Errorf("scenario score is not passing")
	}

	return MissionCompletion{
		MissionID:  missionID,
		Completed:  true,
		Grade:      string(state.Score.Grade),
		Keystrokes: len(state.Runtime.KeyTrace),
		TimeMs:     elapsed.Milliseconds(),
		Attempts:   state.Runtime.Attempts,
	}, nil
}

func ApplyMissionCompletion(current progress.Progress, completion MissionCompletion) progress.Progress {
	next := copyProgress(current)
	if next.Missions == nil {
		next.Missions = make(map[string]progress.MissionProgress)
	}
	if !completion.Completed {
		return next
	}
	next.CompleteMission(completion.MissionID, completion.Grade, completion.Keystrokes, completion.TimeMs)
	return next
}

func copyProgress(current progress.Progress) progress.Progress {
	next := current
	next.Tutorials = copyTutorials(current.Tutorials)
	next.Missions = copyMissions(current.Missions)
	return next
}

func copyTutorials(values map[string]progress.TutorialProgress) map[string]progress.TutorialProgress {
	if values == nil {
		return nil
	}
	next := make(map[string]progress.TutorialProgress, len(values))
	for key, value := range values {
		next[key] = value
	}
	return next
}

func copyMissions(values map[string]progress.MissionProgress) map[string]progress.MissionProgress {
	if values == nil {
		return nil
	}
	next := make(map[string]progress.MissionProgress, len(values))
	for key, value := range values {
		next[key] = value
	}
	return next
}
