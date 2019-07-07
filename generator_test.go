package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func executeGenerateTests(g Generator, is *require.Assertions, t *testing.T) {
	testCases := []struct {
		title  string
		lenght int
	}{
		{
			title:  "null size",
			lenght: 0,
		},
		{
			title:  "test len 1",
			lenght: 1,
		},
		{
			title:  "test len 100",
			lenght: 100,
		},
		{
			title:  "test len 300",
			lenght: 300,
		},
		{
			title:  "test negativ",
			lenght: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			got, err := g.Generate(tc.lenght)

			if tc.lenght >= 0 {
				is.NoError(err)
				is.Len(got, tc.lenght)
			} else {
				is.Error(err)
				is.EqualError(err, ErrNegativValue.Error())
				is.Len(got, 0)
			}
		})
	}
}

func TestGenerate(t *testing.T) {

	is := require.New(t)

	generators := []Generator{
		NewASCII(),
	}

	for _, generator := range generators {
		executeGenerateTests(generator, is, t)
	}

}
