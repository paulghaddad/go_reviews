package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Profile Profile
}

type Profile struct {
	Age int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct{
		Name string
		Input interface{}
		ExpectedCalls []string
	}{
			{
			"Struct with one field",
			struct {
				Name string
			}{"Paul"},
			[]string{ "Paul" },
		},
			{
			"Struct with two fields",
			struct {
				Name string
				City string
			}{"Paul", "St. Augustine"},
			[]string{ "Paul", "St. Augustine" },
		},
			{
			"Struct with non-string fields",
			struct {
				Name string
				Age int
			}{"Paul", 38},
			[]string{ "Paul"},
		},
		{
			"Nested fields",
			Person{
				"Paul",
				Profile{38, "St. Augustine"},
			},
			[]string{"Paul", "St. Augustine"},
		},
		{
			"Pointers",
			&Person{
				"Paul",
				Profile{38, "St. Augustine"},
			},
			[]string{"Paul", "St. Augustine"},
		},
		{
			"Slices",
			[]Profile{
				{38, "St. Augustine"},
				{37, "Boston"},
			},
			[]string{"St. Augustine", "Boston"},
		},
		{
			"Arrays",
			[2]Profile{
				{38, "St. Augustine"},
				{37, "Boston"},
			},
			[]string{"St. Augustine", "Boston"},
		},
	}


	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %q; want: %q", got, test.ExpectedCalls)
			}
		})
	}


	t.Run("With maps", func(t *testing.T) {
		aMap := map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
		}

		var got []string
		Walk(aMap, func(input string) {
			got = append(got, input)
		})
		
		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("With channels", func(t *testing.T) {
		aChannel := make(chan Profile)
		
		go func() {
			aChannel <- Profile{38, "St. Augustine"}
			aChannel <- Profile{37, "Boston"}
			close(aChannel)
		}()

		var got []string
		want := []string{"St. Augustine", "Boston"}

		Walk(aChannel, func(input string) {
			got = append(got, input)
		})
		
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v; want: %v", got, want)
		}
	})

	t.Run("With function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{38, "St. Augustine"}, Profile{37, "Boston"}
		}

		var got []string
		want := []string{"St. Augustine", "Boston"}

		Walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v; want: %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string)  {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
