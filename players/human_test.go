package players

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"dots/board"
)

func TestNewHuman(t *testing.T) {

	buff := new(bytes.Buffer)
	human := NewHuman(buff)

	if human.writer != os.Stdout || human.reader != buff {
		t.Errorf("Initialising human failed")
	}
}

func TestHumanDoMove(t *testing.T) {

	buff := new(bytes.Buffer)
	human := NewHuman(buff)

	for i := uint(0); i < 64; i++ {

		if (i%8 == 3 || i%8 == 4) && (i/8 == 3 || i/8 == 4) {
			continue
		}

		b := *board.NewBoard()
		for !b.Moves().TestBit(i) {
			b = *board.RandomBoard(20)
		}

		str := fmt.Sprintf("%c%d", 'a'+(i%8), 1+(i/8))
		buff.WriteString(str)

		got := human.DoMove(b)

		expected := b
		expected.DoMove(i)

		if expected != got {
			t.Errorf("Human produced wrong child board\n")
		}

		buff.Reset()
	}

	// sending junk should be ignored
	buff.WriteString("asdf\n")
	buff.WriteString("e6\n")

	expected := *board.NewBoard()
	expected.DoMove(44)
	got := human.DoMove(*board.NewBoard())

	if expected != got {
		t.Errorf("Human produced wrong child board\n")
	}

	buff.Reset()

	// sending invalid move should be ignored
	buff.WriteString("e5\n")
	buff.WriteString("e6\n")

	got = human.DoMove(*board.NewBoard())

	if expected != got {
		t.Errorf("Human produced wrong child board\n")
	}

}
