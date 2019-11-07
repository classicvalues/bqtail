package config

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/afs/matcher"
	"testing"
)

func TestRoutes_HasMatch(t *testing.T) {
	var useCases = []struct {
		description string
		Ruleset     []*Rule
		URL         string
		expextTable string
	}{
		{
			description: "suffix match",
			Ruleset: []*Rule{
				{
					When: matcher.Basic{
						Suffix: ".tsv",
					},
					Dest: &Destination{
						Table: "project:dataset:table1",
					},
				},
				{
					When: matcher.Basic{
						Suffix: ".csv",
					},
					Dest: &Destination{
						Table: "project:dataset:table2",
					},
				},
			},

			URL:         "ssh://zz/folder/a.csv",
			expextTable: "project:dataset:table2",
		},
		{
			description: "prefix with long file",
			Ruleset: []*Rule{
				{
					When: matcher.Basic{
						Prefix: "/s",
					},
					Dest: &Destination{
						Table: "project:dataset:table3",
					},
				},
				{
					When: matcher.Basic{
						Prefix: "/g",
					},
					Dest: &Destination{
						Table: "project:dataset:table4",
					},
				},
			},

			URL:         "ssh://zz/folder/a.csv",
			expextTable: "",
		},
	}

	for _, useCase := range useCases {
		rules := &Ruleset{
			Rules: useCase.Ruleset,
		}

		actual := rules.Match(useCase.URL)
		if useCase.expextTable == "" {
			assert.Nil(t, actual, useCase.description)
			continue
		}

		if !assert.NotNil(t, actual, useCase.description) {
			continue
		}

		assert.Equal(t, useCase.expextTable, actual[0].Dest.Table, useCase.description)
	}
}
