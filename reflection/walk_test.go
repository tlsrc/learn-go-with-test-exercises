package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"one field"},
			[]string{"one field"},
		},
		{
			"Struct with 2 string fields",
			struct {
				Name string
				City string
			}{"two fields a", "two fields b"},
			[]string{"two fields a", "two fields b"},
		},
		{
			"Struct with non string fields",
			struct {
				Name string
				Age  int
			}{"name", 42},
			[]string{"name"},
		},
		{
			"Nested structs",
			Person{"name", Profile{42, "city"}},
			[]string{"name", "city"},
		},
		{
			"With pointers",
			&Person{"name", Profile{42, "city"}},
			[]string{"name", "city"},
		},
		{
			"With slices",
			[]Profile{
				{1, "city1"},
				{2, "city2"},
			},
			[]string{"city1", "city2"},
		},
		{
			"With arrays",
			[2]Profile{
				{1, "city1"},
				{2, "city2"},
			},
			[]string{"city1", "city2"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			assertDeepEqual(got, test.ExpectedCalls, t)
		})
	}

	t.Run("with maps", func(t *testing.T) {
		m := map[string]string{
			"k1": "v1", "k2": "v2",
		}

		var got []string
		walk(m, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "v1")
		assertContains(t, got, "v2")
	})

	t.Run("with channels", func(t *testing.T) {
		c := make(chan Profile)
		go func() {
			c <- Profile{1, "city1"}
			c <- Profile{2, "city2"}
			close(c)
		}()

		var got []string
		want := []string{"city1", "city2"}

		walk(c, func(input string) {
			got = append(got, input)
		})

		assertDeepEqual(got, want, t)
	})

	t.Run("with functions", func(t *testing.T) {
		f := func() (Profile, Profile) {
			return Profile{1, "city1"}, Profile{2, "city2"}
		}

		var got []string
		want := []string{"city1", "city2"}

		walk(f, func(input string) {
			got = append(got, input)
		})

		assertDeepEqual(got, want, t)
	})
}

func assertDeepEqual(got, want interface{}, t *testing.T) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()
	found := false
	for _, x := range haystack {
		if x == needle {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Expected %+v to contain %q but it did not", haystack, needle)
	}
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
