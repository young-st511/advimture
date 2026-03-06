package progress

// Rank represents the player's progression level.
type Rank int

const (
	RankIntern Rank = iota
	RankJunior
	RankSenior
	RankStaff
	RankPrincipal
	RankVimMaster
)

func (r Rank) String() string {
	switch r {
	case RankIntern:
		return "Intern"
	case RankJunior:
		return "Junior"
	case RankSenior:
		return "Senior"
	case RankStaff:
		return "Staff"
	case RankPrincipal:
		return "Principal"
	case RankVimMaster:
		return "Vim Master"
	default:
		return "Intern"
	}
}

// RankRequirements defines the minimum tutorials completed for each rank.
var RankRequirements = map[Rank]int{
	RankIntern:    0,
	RankJunior:    3,
	RankSenior:    6,
	RankStaff:     8,
	RankPrincipal: 10,
	RankVimMaster: 10, // + all missions
}

// CalculateRank determines rank based on completed tutorials and missions.
func CalculateRank(tutorialsCompleted, missionsCompleted int) Rank {
	if tutorialsCompleted >= 10 && missionsCompleted >= 10 {
		return RankVimMaster
	}
	if tutorialsCompleted >= 10 {
		return RankPrincipal
	}
	if tutorialsCompleted >= 8 {
		return RankStaff
	}
	if tutorialsCompleted >= 6 {
		return RankSenior
	}
	if tutorialsCompleted >= 3 {
		return RankJunior
	}
	return RankIntern
}
