package db

import (
	"context"
	"testing"

	"github.com/benmertmiyim/mpi-draw/util"
	"github.com/stretchr/testify/require"
)

var codeset []string

func TestGenerateCodeSet(t *testing.T) {
	genSet, err := util.GenerateCodeSet(6, 1000)
	require.NoError(t, err)

	if err != nil {
		return
	} else {
		codeset = genSet
	}

	require.NotEmpty(t, genSet)

	codes := make([]Code, len(genSet))

	for _, code := range genSet {
		result, err := TestQueries.CreateCode(context.Background(), code)
		require.NoError(t, err)
		require.NotEmpty(t, result)
		require.Equal(t, code, result.Code)
		require.Equal(t, true, result.Active)
		require.Equal(t, false, result.Banned)
		require.Equal(t, false, result.Used)
		codes = append(codes, result)
	}
}

func TestGetCodes(t *testing.T) {
	for _, code := range codeset {
		result, err := TestQueries.GetCode(context.Background(), code)
		require.NoError(t, err)
		require.NotEmpty(t, result)
		require.Equal(t, code, result.Code)
		require.Equal(t, true, result.Active)
		require.Equal(t, false, result.Banned)
		require.Equal(t, false, result.Used)
	}

	err := TestQueries.DeleteCodes(context.Background())
	require.NoError(t, err)
}
