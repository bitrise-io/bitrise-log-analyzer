package pattern

import (
	"testing"

	"github.com/bitrise-io/go-utils/testutil"
	"github.com/stretchr/testify/require"
)

func Test_Matcher_ProcessText(t *testing.T) {
	t.Log("Empty")
	{
		matcher := NewMatcher([]Model{})
		require.NoError(t, matcher.ProcessText(``))
		require.Equal(t, []Model{}, matcher.Results())
	}

	t.Log("One liner text - single match")
	{
		matcher := NewMatcher([]Model{
			{Lines: []string{"match"}},
			{Lines: []string{"no match"}},
		})
		require.NoError(t, matcher.ProcessText(`this should match`))
		require.Equal(t,
			[]Model{
				{Lines: []string{"match"}},
			},
			matcher.Results())
	}

	t.Log("One liner text - single match - same pattern multiple times - result should only include it once")
	{
		matcher := NewMatcher([]Model{
			{Lines: []string{"match"}},
		})
		require.NoError(t, matcher.ProcessText(`this should match or match and match`))
		require.Equal(t,
			[]Model{
				{Lines: []string{"match"}},
			},
			matcher.Results())
	}

	t.Log("One liner text - multi single line match")
	{
		matcher := NewMatcher([]Model{
			{Lines: []string{"should"}},
			{Lines: []string{"match"}},
		})
		require.NoError(t, matcher.ProcessText(`this should match`))
		testutil.EqualSlicesWithoutOrder(t,
			[]Model{
				{Lines: []string{"should"}},
				{Lines: []string{"match"}},
			},
			matcher.Results())
	}

	t.Log("Multi-line input text - no match")
	{
		matcher := NewMatcher([]Model{
			{Lines: []string{"nothing should match this"}},
		})
		require.NoError(t, matcher.ProcessText(`first line,
second line
and the third line`))
		testutil.EqualSlicesWithoutOrder(t, []Model{}, matcher.Results())
	}

	t.Log("Multi-line input text - multi single line match")
	{
		matcher := NewMatcher([]Model{
			{Lines: []string{"should"}},
			{Lines: []string{"match"}},
		})
		require.NoError(t, matcher.ProcessText(`this should be catched,
as well as
this should match`))
		testutil.EqualSlicesWithoutOrder(t,
			[]Model{
				{Lines: []string{"should"}},
				{Lines: []string{"match"}},
			},
			matcher.Results())
	}

	t.Log("Multi-line input text - multi lines pattern match")
	{
		matcher := NewMatcher([]Model{
			{Lines: []string{
				"this should",
				"as well",
				"match",
			}},
		})
		require.NoError(t, matcher.ProcessText(`this should be catched,
as well as
this should match`))
		testutil.EqualSlicesWithoutOrder(t,
			[]Model{
				{Lines: []string{"this should", "as well", "match"}},
			}, matcher.Results())
	}
}
