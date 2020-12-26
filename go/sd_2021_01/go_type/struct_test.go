package go_type_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

type secret struct {
	ID        string
	CreatedAt time.Time
	token     string
}

func (s *secret) Read(p []byte) (int, error) {
	return bytes.NewBuffer(p).WriteString(s.token)
}

func NewSecret() io.Reader {
	return &secret{
		ID:        "dummy",
		CreatedAt: time.Now(),
		token:     "dummy",
	}
}

type Chip struct {
	Number int
}

type Card struct {
	text string
	Chip
	Number int
}

func (c *Chip) Scan() int {
	fmt.Println(c.Number)
	return c.Number
}

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func TestStruct(t *testing.T) {
	t.Run("private field hasn't been included json", func(t *testing.T) {
		s := NewSecret()
		err := json.NewEncoder(os.Stdout).Encode(s)
		if err != nil {
			t.Fatal("failed to json encode, error ", err)
		}

		_, err = json.Marshal(s)
		if err != nil {
			t.Fatal("Failed to marshal json ", s, " error ", err)
		}
	})

	t.Run("struct embed and method receiver", func(t *testing.T) {
		c := Card{
			text: "Credit",
			Chip: Chip{
				Number: 1000,
			},
			Number: 2000,
		}

		// Scan() メソッドのレシーバは、 Card ではなく、 Chip になる
		// Card.Number -> 2000
		// Chip.Number -> 1000
		if c.Scan() != 1000 {
			t.Fatal("expected 1000 but got ", c.Scan())
		}
	})

	t.Run("method can be used as a value", func(t *testing.T) {
		h := Hex(16)
		functionAsAValue := h.String
		if functionAsAValue() != "10" {
			t.Fatal("expected 10 but got ", functionAsAValue())
		}

		functionAsAStatement := Hex.String
		if functionAsAStatement(100) != "64" {
			t.Fatal("expected 64 but got ", functionAsAStatement(100))
		}
	})
}
