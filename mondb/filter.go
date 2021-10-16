// Monster Filters
package mondb

import "sort"

// Filter filters monsters based on a user-provided filter function.
// If the filter is nil, returns monsters as is.
// If the filter excludes all results, returns nil.
func Filter(monsters []Monster, filter func(Monster) bool) []Monster {
	if filter == nil {
		return monsters
	}

	var filteredMonsters []Monster
	for _, m := range monsters {
		if filter(m) {
			filteredMonsters = append(filteredMonsters, m)
		}
	}
	return filteredMonsters
}

// Sort sorts monsters using a given comparison function in direction order.
// Sorting is stable. In other words, sorting monsters does not change the relative order of monsters.
// For example, if monsters are arranged alphabetically and then
// sorted by size, monsters of the same size will be arranged alphabetically.
func Sort(monsters []Monster, fn CompareFn, direction Direction) {
	ms := &monsterSorter{
		monsters:  monsters,
		compareFn: fn,
	}
	if direction == Descending {
		sort.Stable(sort.Reverse(ms))
		return
	}
	sort.Stable(ms)
}

type monsterSorter struct {
	monsters  []Monster
	compareFn CompareFn
}

func (ms *monsterSorter) Len() int {
	return len(ms.monsters)
}

func (ms *monsterSorter) Swap(i, j int) {
	ms.monsters[i], ms.monsters[j] = ms.monsters[j], ms.monsters[i]
}

func (ms *monsterSorter) Less(i, j int) bool {
	return ms.compareFn(ms.monsters[i], ms.monsters[j])
}

type Direction int

const (
	Ascending = iota
	Descending
)

// CompareFn compares two monsters for sorting purposes.
// This will be used as a custom sort.Less function.
type CompareFn func(m1, m2 Monster) bool

// ByName compares monsters by name.
func ByName(m1, m2 Monster) bool {
	return m1.Name < m2.Name
}

// ByElement compares monsters by elemental property name.
func ByElement(m1, m2 Monster) bool {
	return m1.Property.String() < m2.Property.String()
}

// ByRace compares monsters by race.
func ByRace(m1, m2 Monster) bool {
	return m1.Race.String() < m2.Race.String()
}

// BySize compares monsters by size.
func BySize(m1, m2 Monster) bool {
	return m1.Size < m2.Size
}

// ByBaseExp compares monsters by base experience.
func ByBaseExp(m1, m2 Monster) bool {
	return m1.Base < m2.Base
}

// ByBaseExpPerHP compares monsters by base experience per health point.
func ByBaseExpPerHP(m1, m2 Monster) bool {
	if m1.HP == 0 || m2.HP == 0 {
		return ByBaseExp(m1, m2)
	}
	return float64(m1.Base)/float64(m1.HP) < float64(m2.Base)/float64(m2.HP)
}

// ByJobExp compares monsters by job experience.
func ByJobExp(m1, m2 Monster) bool {
	return m1.Job < m2.Job
}

// ByJobExpPerHP compares monsters by job experience per health point.
func ByJobExpPerHP(m1, m2 Monster) bool {
	if m1.HP == 0 || m2.HP == 0 {
		return ByJobExp(m1, m2)
	}
	return float64(m1.Job)/float64(m1.HP) < float64(m2.Job)/float64(m2.HP)
}

// ByHP compares monsters by HP.
func ByHP(m1, m2 Monster) bool {
	return m1.HP < m2.HP
}

// ByLevel compares monsters by Level.
func ByLevel(m1, m2 Monster) bool {
	return m1.Level < m2.Level
}

// ByAttackRange compares monsters by AttackRange.
func ByAttackRange(m1, m2 Monster) bool {
	return m1.AttackRange < m2.AttackRange
}

// maxHit calculates maximum HIT required for 100%HIT.
func maxHit(m Monster) int {
	// According to https://irowiki.org/classic/Attacks:
	// A hit has a 80 + AttackerHit - DefenderFlee"% chance of occuring.
	// According to https://irowiki.org/classic/FLEE:
	// Actual Flee = Level + AGI for monsters (no skill or item bonus).
	// So for 100%HIT, 80 + AttackerHit - (Level + AGI) = 100
	// or, AttackerHit = 20 + Level + AGI.
	return 20 + m.Level + m.Agi
}

// ByMaxHitRate compares monsters by 100% HIT rate.
func ByMaxHitRate(m1, m2 Monster) bool {
	return maxHit(m1) < maxHit(m2)
}

// maxDodge calculates maximum FLEE required for 95%Flee.
func maxDodge(m Monster) int {
	// According to https://irowiki.org/classic/FLEE:
	// Dodge Rate(%) = 100% - (AttackerHit + 80 - DefendersFlee)
	// where AttackerHit for monsters is BaseLevel + DEX.
	// For 95%FLEE, 95 = 100 - (AttackerHit + 80 - DefendersFlee)
	// or, DefendersFlee = 75 + Level + DEX.
	return 75 + m.Level + m.Dex
}

// ByMaxDodgeRate compares monsters by 95% FLEE rate.
func ByMaxDodgeRate(m1, m2 Monster) bool {
	return maxDodge(m1) < maxDodge(m2)
}

// ByDefense compares monsters by DEF.
func ByDefense(m1, m2 Monster) bool {
	return m1.Def < m2.Def
}

// ByMagicDefense compares monsters by MDEF.
func ByMagicDefense(m1, m2 Monster) bool {
	return m1.Mdef < m2.Mdef
}
