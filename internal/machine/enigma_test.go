package machine

import (
	"math/rand"
	"testing"

	"github.com/elliottpolk/enigma/internal/plugboard"
	"github.com/elliottpolk/enigma/internal/reflector"
	"github.com/elliottpolk/enigma/internal/rotor"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

type param struct {
	conns []*plugboard.Connection
	m     *Machine
	input string
	want  string
}

func TestEncrypt(t *testing.T) {
	params := map[string]*param{
		"basic": {
			[]*plugboard.Connection{},
			&Machine{
				Left:      rotor.Create("I", 0, 0),
				Middle:    rotor.Create("II", 0, 0),
				Right:     rotor.Create("III", 0, 0),
				Reflector: reflector.Create("B"),
			},
			"ABCDEFGHIJKLMNOPQRSTUVWXYZAAAAAAAAAAAAAAAAAAAAAAAAAABBBBBBBBBBBBBBBBBBBBBBBBBBABCDEFGHIJKLMNOPQRSTUVWXYZ",
			"BJELRQZVJWARXSNBXORSTNCFMEYHCXTGYJFLINHNXSHIUNTHEORXOPLOVFEKAGADSPNPCMHRVZCYECDAZIHVYGPITMSRZKGGHLSRBLHL",
		},
		"varied rotors": {
			[]*plugboard.Connection{},
			&Machine{
				Left:      rotor.Create("VII", 10, 1),
				Middle:    rotor.Create("V", 5, 2),
				Right:     rotor.Create("IV", 12, 3),
				Reflector: reflector.Create("B"),
			},
			"ABCDEFGHIJKLMNOPQRSTUVWXYZAAAAAAAAAAAAAAAAAAAAAAAAAABBBBBBBBBBBBBBBBBBBBBBBBBBABCDEFGHIJKLMNOPQRSTUVWXYZ",
			"FOTYBPKLBZQSGZBOPUFYPFUSETWKNQQHVNHLKJZZZKHUBEJLGVUNIOYSDTEZJQHHAOYYZSENTGXNJCHEDFHQUCGCGJBURNSEDZSEPLQP",
		},
		"long input": {
			[]*plugboard.Connection{},
			&Machine{
				Left:      rotor.Create("III", 3, 11),
				Middle:    rotor.Create("VI", 5, 13),
				Right:     rotor.Create("VIII", 9, 19),
				Reflector: reflector.Create("B"),
			},
			"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			"YJKJMFQKPCUOCKTEZQVXYZJWJFROVJMWJVXRCQYFCUVBRELVHRWGPYGCHVLBVJEVTTYVMWKJFOZHLJEXYXRDBEVEHVXKQSBPYZNIQDCBGTDDWZQWLHIBQNTYPIEBMNINNGMUPPGLSZCBRJULOLNJSOEDLOBXXGEVTKCOTTLDZPHBUFKLWSFSRKOMXKZELBDJNRUDUCOTNCGLIKVKMHHCYDEKFNOECFBWRIEFQQUFXKKGNTSTVHVITVHDFKIJIHOGMDSQUFMZCGGFZMJUKGDNDSNSJKWKENIRQKSUUHJYMIGWWNMIESFRCVIBFSOUCLBYEEHMESHSGFDESQZJLTORNFBIFUWIFJTOPVMFQCFCFPYZOJFQRFQZTTTOECTDOOYTGVKEWPSZGHCTQRPGZQOVTTOIEGGHEFDOVSUQLLGNOOWGLCLOWSISUGSVIHWCMSIUUSBWQIGWEWRKQFQQRZHMQJNKQTJFDIJYHDFCWTHXUOOCVRCVYOHLV",
		},
	}

	for n, p := range params {
		t.Run(n, func(t *testing.T) {
			pb, err := plugboard.NewPlugboard(p.conns)
			require.NoError(t, err)

			p.m.Plugboard = pb

			require.Equal(t, p.want, p.m.Encrypt(p.input))
		})

	}
}

func TestDecrypt(t *testing.T) {
	var (
		rotors = []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII"}
		input  = make([]rune, 1000)
	)

	for i := 0; i < 1000; i++ {
		input[i] = rune(rand.Intn(26) + 65)
	}

	pb, err := plugboard.NewPlugboard([]*plugboard.Connection{})
	require.NoError(t, err)

	for i := 0; i < 10; i++ {
		r := []string{
			rotors[rand.Intn(len(rotors))],
			rotors[rand.Intn(len(rotors))],
			rotors[rand.Intn(len(rotors))],
		}

		sp := []int{
			rand.Intn(26),
			rand.Intn(26),
			rand.Intn(26),
		}

		rs := []int{
			rand.Intn(26),
			rand.Intn(26),
			rand.Intn(26),
		}

		m1 := &Machine{
			Left:      rotor.Create(r[0], sp[0], rs[0]),
			Middle:    rotor.Create(r[1], sp[1], rs[1]),
			Right:     rotor.Create(r[2], sp[2], rs[2]),
			Reflector: reflector.Create("B"),
			Plugboard: pb,
		}

		r1 := m1.Encrypt(string(input))

		m2 := &Machine{
			Left:      rotor.Create(r[0], sp[0], rs[0]),
			Middle:    rotor.Create(r[1], sp[1], rs[1]),
			Right:     rotor.Create(r[2], sp[2], rs[2]),
			Reflector: reflector.Create("B"),
			Plugboard: pb,
		}

		r2 := m2.Encrypt(r1)
		require.Equal(t, string(input), r2)
	}
}

func TestPlugboard(t *testing.T) {
	params := map[string]*param{
		"4 plugs": {
			[]*plugboard.Connection{
				{A: 'A', B: 'C'},
				{A: 'F', B: 'G'},
				{A: 'J', B: 'Y'},
				{A: 'L', B: 'W'},
			},
			&Machine{
				Left:      rotor.Create("I", 0, 0),
				Middle:    rotor.Create("II", 0, 0),
				Right:     rotor.Create("III", 0, 0),
				Reflector: reflector.Create("B"),
			},
			"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			"QREBNMCYZELKQOJCGJVIVGLYEMUPCURPVPUMDIWXPPWROOQEGI",
		},
		"6 plugs": {
			[]*plugboard.Connection{
				{A: 'B', B: 'M'},
				{A: 'D', B: 'H'},
				{A: 'R', B: 'S'},
				{A: 'K', B: 'N'},
				{A: 'G', B: 'Z'},
				{A: 'F', B: 'Q'},
			},
			&Machine{
				Left:      rotor.Create("IV", 0, 0),
				Middle:    rotor.Create("VI", 10, 0),
				Right:     rotor.Create("III", 6, 0),
				Reflector: reflector.Create("B"),
			},
			"WRBHFRROSFHBCHVBENQFAGNYCGCRSTQYAJNROJAKVKXAHGUZHZVKWUTDGMBMSCYQSKABUGRVMIUOWAPKCMHYCRTSDEYTNJLVWNQY",
			"FYTIDQIBHDONUPAUVPNKILDHDJGCWFVMJUFNJSFYZTSPITBURMCJEEAMZAZIJMZAVFCTYTKYORHYDDSXHBLQWPJBMSSWIPSWLENZ",
		},
		"10 plugs": {
			[]*plugboard.Connection{
				{A: 'A', B: 'G'},
				{A: 'H', B: 'R'},
				{A: 'Y', B: 'T'},
				{A: 'K', B: 'I'},
				{A: 'F', B: 'L'},
				{A: 'W', B: 'E'},
				{A: 'N', B: 'M'},
				{A: 'S', B: 'D'},
				{A: 'O', B: 'P'},
				{A: 'Q', B: 'J'},
			},
			&Machine{
				Left:      rotor.Create("I", 0, 5),
				Middle:    rotor.Create("II", 1, 5),
				Right:     rotor.Create("III", 20, 4),
				Reflector: reflector.Create("B"),
			},
			"RNXYAZUYTFNQFMBOLNYNYBUYPMWJUQSBYRHPOIRKQSIKBKEKEAJUNNVGUQDODVFQZHASHMQIHSQXICTSJNAUVZYIHVBBARPJADRH",
			"CFBJTPYXROYGGVTGBUTEBURBXNUZGGRALBNXIQHVBFWPLZQSCEZWTAWCKKPRSWOGNYXLCOTQAWDRRKBCADTKZGPWSTNYIJGLVIUQ",
		},
	}

	for n, p := range params {
		t.Run(n, func(t *testing.T) {
			pb, err := plugboard.NewPlugboard(p.conns)
			require.NoError(t, err)

			p.m.Plugboard = pb

			require.Equal(t, p.want, p.m.Encrypt(p.input))
		})

	}
}
