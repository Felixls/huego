package huego_test

import (
	"os"
	"testing"

	"github.com/Felixls/huego"
)

func TestGetRules(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	rules, err := b.GetRules()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Found %d rules", len(rules))
	for _, rule := range rules {
		t.Log(rule)
	}
}

func TestGetRule(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	rules, err := b.GetRules()
	if err != nil {
		t.Fatal(err)
	}
	for _, rule := range rules {
		l, err := b.GetRule(rule.ID)
		if err != nil {
			t.Fatal(err)
		} else {
			t.Log(l)
		}
		break
	}
}

func TestCreateRule(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	conditions := []*huego.Condition{
		{
			Address:  "/sensors/2/state/buttonevent",
			Operator: "eq",
			Value:    "16",
		},
	}
	actions := []*huego.RuleAction{
		{
			Address: "/groups/0/action",
			Method:  "PUT",
			Body:    &huego.State{On: true},
		},
	}
	rule := &huego.Rule{
		Name:       "Huego Test Rule",
		Conditions: conditions,
		Actions:    actions,
	}
	resp, err := b.CreateRule(rule)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Rule created")
		for k, v := range resp.Success {
			t.Logf("%v: %s", k, v)
		}
	}
}

func TestUpdateRule(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 3
	resp, err := b.UpdateRule(id, &huego.Rule{
		Actions: []*huego.RuleAction{
			{
				Address: "/groups/3/action",
				Method:  "PUT",
				Body:    &huego.State{On: true},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Rule %d updated", id)
		for k, v := range resp.Success {
			t.Logf("%v: %s", k, v)
		}
	}
}

func TestDeleteRule(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 3
	err := b.DeleteRule(id)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Rule %d deleted", id)
	}
}
