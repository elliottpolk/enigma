package main

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnmarshal(t *testing.T) {
	in, err := ioutil.ReadFile("../config/config_a.yml")
	require.NoError(t, err)

	//var m map[string]interface{}
	//require.NoError(t, yaml.Unmarshal(in, &m))

	m, err := unmarshal(in)
	require.NoError(t, err)

	fmt.Printf("%+v\n", m)
}
