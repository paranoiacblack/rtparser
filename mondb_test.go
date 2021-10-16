package mondb

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUnmarshalMonster(t *testing.T) {
	// First entry on https://account.ragnaroktravels.com/mondrops.
	data := []byte(`{"Name":"Alarm","item1":["Needle of Alarm","needle of alarm","1095"],"percent1":5335,"item2":["Clip","clip","2607"],"percent2":1,"item3":["Skull","skull","7005"],"percent3":1500,"item4":["Magnifier","magnifier","611"],"percent4":1300,"item5":["Oridecon","oridecon","984"],"percent5":105,"item6":["Key of Clock Tower","key of clock tower","7026"],"percent6":20,"item7":["Zargon","zargon","912"],"percent7":1500,"item8":["Alarm Card","alarm card","4244"],"percent8":1,"aRan":1,"LV":58,"HP":10647,"SP":0,"str":1,"int":10,"vit":72,"dex":85,"agi":62,"luk":45,"atk1":480,"atk2":120,"def":15,"exp":3987,"jexp":2300,"inc":58,"as":10,"es":12,"Mspeed":300,"rechargeTime":1020,"attackedMT":768,"attackMT":500,"property":60,"scale":1,"class":0,"race":0,"mdef":15,"tamingitem":"0","fooditem":"0","db_name":["ALARM","ALARM"]}`)

	var m Monster
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}

	// Quick check: marshal Monster and compare json output.
	out, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	if !json.Valid(out) {
		t.Fatalf("Marshalled monster produces invalid JSON: %s", data)
	}

	if !cmp.Equal(data, out) {
		t.Errorf("Unmarshalled output doesn't match input data: %s", cmp.Diff(data, out))
	}
}

// A subset of monsters used in the following tests, arbitrarily chosen.
var (
	Alarm       = Monster{Name: "Alarm", Info: Info{Property: Neutral3, Race: Formless, Size: Medium}}
	Cornutus    = Monster{Name: "Cornutus", Info: Info{Property: Water1, Race: Fish, Size: Small}}
	Deniro      = Monster{Name: "Deniro", Info: Info{Property: Earth1, Race: Insect, Size: Small}}
	Picky       = Monster{Name: "Picky", Info: Info{Property: Fire1, Race: Brute, Size: Small}}
	RotarZairo  = Monster{Name: "Rotar Zairo", Info: Info{Property: Wind2, Race: Formless, Size: Large}}
	ThiefBugEgg = Monster{Name: "Thief Bug Egg", Info: Info{Property: Shadow1, Race: Insect, Size: Small}}
	Zealotus    = Monster{Name: "Zealotus", Info: Info{Property: Neutral3, Race: DemiHuman, Size: Medium}}
)

func TestFilter(t *testing.T) {
	monsters := []Monster{Alarm, Cornutus, Deniro, Picky, RotarZairo, ThiefBugEgg, Zealotus}

	type testData struct {
		name   string
		filter func(Monster) bool
		want   []Monster
	}

	tests := []testData{
		{
			name:   "Allow all",
			filter: func(Monster) bool { return true },
			want:   monsters,
		},
		{
			name:   "Deny all",
			filter: func(Monster) bool { return false },
		},
		{
			name: "No filter",
			want: monsters,
		},
		{
			name:   "Exact name",
			filter: func(m Monster) bool { return m.Name == "Zealotus" || m.Name == "Cornutus" },
			want:   []Monster{Cornutus, Zealotus},
		},
		{
			name:   "Name search",
			filter: func(m Monster) bool { return strings.Contains(m.Name, "iro") },
			want:   []Monster{Deniro, RotarZairo},
		},
		{
			name:   "Exact property",
			filter: func(m Monster) bool { return m.Property == Neutral3 },
			want:   []Monster{Alarm, Zealotus},
		},
		{
			name: "Property search",
			filter: func(m Monster) bool {
				switch m.Property {
				case Water1, Fire1, Earth1, Shadow1:
					return true
				}
				return false
			},
			want: []Monster{Cornutus, Deniro, Picky, ThiefBugEgg},
		},
		{
			name:   "Exact race",
			filter: func(m Monster) bool { return m.Race == Formless },
			want:   []Monster{Alarm, RotarZairo},
		},
		{
			name:   "Race search",
			filter: func(m Monster) bool { return m.Race == Brute || m.Race == DemiHuman },
			want:   []Monster{Picky, Zealotus},
		},
		{
			name:   "Exact size",
			filter: func(m Monster) bool { return m.Size == Large },
			want:   []Monster{RotarZairo},
		},
	}

	for _, test := range tests {
		got := Filter(monsters, test.filter)
		if diff := cmp.Diff(got, test.want); diff != "" {
			t.Errorf("Filter(%s): unexpected results %s", test.name, diff)
		}
	}
}

func TestSort(t *testing.T) {
	monsters := []Monster{Alarm, Cornutus, Deniro, Picky, RotarZairo, ThiefBugEgg, Zealotus}

	type testData struct {
		name string
		fn   CompareFn
		// Note: default direction is Ascending.
		direction Direction
		want      []Monster
	}

	tests := []testData{
		{
			name: "Sort by name",
			fn:   ByName,
			want: monsters,
		},
		{
			name:      "Sort by name descending",
			fn:        ByName,
			direction: Descending,
			want:      []Monster{Zealotus, ThiefBugEgg, RotarZairo, Picky, Deniro, Cornutus, Alarm},
		},
		{
			name: "Sort by element property name",
			fn:   ByElement,
			want: []Monster{Deniro, Picky, Alarm, Zealotus, ThiefBugEgg, Cornutus, RotarZairo},
		},
		{
			name:      "Sort by element property name descending",
			fn:        ByElement,
			direction: Descending,
			want:      []Monster{RotarZairo, Cornutus, ThiefBugEgg, Alarm, Zealotus, Picky, Deniro},
		},
		{
			name: "Sort by race",
			fn:   ByRace,
			want: []Monster{Picky, Zealotus, Cornutus, Alarm, RotarZairo, Deniro, ThiefBugEgg},
		},
		{
			name:      "Sort by race descending",
			fn:        ByRace,
			direction: Descending,
			want:      []Monster{Deniro, ThiefBugEgg, Alarm, RotarZairo, Cornutus, Zealotus, Picky},
		},
	}

	for _, test := range tests {
		got := make([]Monster, len(monsters))
		copy(got, monsters)
		Sort(got, test.fn, test.direction)
		if diff := cmp.Diff(got, test.want); diff != "" {
			t.Errorf("Sort(%s) unexpected order %s", test.name, diff)
		}
	}
}
