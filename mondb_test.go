package mondb

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUnmarshalMonster(t *testing.T) {
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
