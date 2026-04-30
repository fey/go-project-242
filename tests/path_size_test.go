package tests

import (
	"code"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPath(t *testing.T) {
	res, _ := code.GetPathSize("../testdata", false, false, false)
	require.Equal(t, "32248B", res)

	res, _ = code.GetPathSize("../testdata/files/included_folder/9.txt", false, false, false)
	require.Equal(t, "9B", res)

	res, _ = code.GetPathSize("../testdata/files/included_folder/999.txt", false, false, false)
	require.Equal(t, "", res)

	res, _ = code.GetPathSize("../testdata", true, false, false)
	require.Equal(t, "31.5KB", res)

	res, _ = code.GetPathSize("../testdata", false, true, false)
	require.Equal(t, "39170B", res)

	res, _ = code.GetPathSize("../testdata", true, true, false)
	require.Equal(t, "38.3KB", res)

	res, _ = code.GetPathSize("../testdata", false, false, true)
	require.Equal(t, "64505B", res)

	res, _ = code.GetPathSize("../testdata", true, false, true)
	require.Equal(t, "63.0KB", res)

	res, _ = code.GetPathSize("../testdata", false, true, true)
	require.Equal(t, "78376B", res)

	res, _ = code.GetPathSize("../testdata", true, true, true)
	require.Equal(t, "76.5KB", res)
}
