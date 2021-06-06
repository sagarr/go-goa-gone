package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Sagar"},
			[]string{"Sagar"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Sagar", "Daryapur"},
			[]string{"Sagar", "Daryapur"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Sagar", 35},
			[]string{"Sagar"},
		},
		{
			"nested struct",
			Person{
				"Sagar",
				Profile{35, "Daryapur"},
			},
			[]string{"Sagar", "Daryapur"},
		},
		{
			"pointers to things",
			&Person{
				"Sagar",
				Profile{35, "Daryapur"},
			},
			[]string{"Sagar", "Daryapur"},
		},
		{
			"slices",
			[]Profile{
				{35, "Sagar"},
				{32, "Shrutika"},
			},
			[]string{"Sagar", "Shrutika"},
		},
		{
			"arrays",
			[2]Profile{
				{35, "Sagar"},
				{32, "Shrutika"},
			},
			[]string{"Sagar", "Shrutika"},
		},
		{
			"maps",
			map[string]string{
				"foo": "bar",
				"baz": "boo",
			},
			[]string{"bar", "boo"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("channels", func(t *testing.T) {
		channel := make(chan Profile)

		go func() {
			channel <- Profile{21, "chinu"}
			channel <- Profile{21, "minu"}
			close(channel)
		}()

		var got []string
		want := []string{"chinu", "minu"}

		Walk(channel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("functions", func(t *testing.T) {
		function := func() (Profile, Profile) {
			return Profile{2, "chinu"}, Profile{2, "minu"}
		}

		var got []string
		want := []string{"chinu", "minu"}

		Walk(function, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
