package calculator

import (
	"testing"
)

func TestSum(t *testing.T) {

	verifica := func(t *testing.T, resultado, esperado int) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%d esperado '%d'", resultado, esperado)
		}
	}

	t.Run("Soma dois inteiros", func(t *testing.T) {
		resultado := Sum(2, 2)
		esperado := 4
		verifica(t, resultado, esperado)
	})

	t.Run("Soma dois negativos:", func(t *testing.T) {
		resultado := Sum(-4, -5)
		esperado := -9
		verifica(t, resultado, esperado)
	})

	t.Run("Soma numeros grandes:", func(t *testing.T) {
		resultado := Sum(100000000000000, 455456745455745342)
		esperado := 455556745455745342
		verifica(t, resultado, esperado)
	})
}

func TestMinus(t *testing.T) {

	verifica := func(t *testing.T, resultado, esperado int) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%d esperado '%d'", resultado, esperado)
		}
	}

	t.Run("Subtrai dois inteiros", func(t *testing.T) {
		resultado := Minus(2, 6)
		esperado := -4
		verifica(t, resultado, esperado)
	})

	t.Run("Subtrai dois negativos:", func(t *testing.T) {
		resultado := Minus(-4, -4)
		esperado := 0
		verifica(t, resultado, esperado)
	})

	t.Run("Subtrai dois grandes:", func(t *testing.T) {
		resultado := Minus(9000000000000000, 8888888888888888)
		esperado := 111111111111112
		verifica(t, resultado, esperado)
	})
}

func TestTimes(t *testing.T) {

	verifica := func(t *testing.T, resultado, esperado int) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%d esperado '%d'", resultado, esperado)
		}
	}

	t.Run("Multiplica dois inteiros", func(t *testing.T) {
		resultado := Times(2, 6)
		esperado := 12
		verifica(t, resultado, esperado)
	})

	t.Run("Multiplica dois negativos:", func(t *testing.T) {
		resultado := Times(4, -4)
		esperado := -16
		verifica(t, resultado, esperado)
	})

	t.Run("Multiplica dois numeros grandes:", func(t *testing.T) {
		resultado := Times(9000000000000000, 8888888888888888)
		esperado := 319172002810953728
		verifica(t, resultado, esperado)
	})
}

func TestDivision(t *testing.T) {

	verifica := func(t *testing.T, resultado, esperado int) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%d esperado '%d'", resultado, esperado)
		}
	}

	t.Run("Divide dois inteiros", func(t *testing.T) {
		resultado := Division(10, 2)
		esperado := 5
		verifica(t, resultado, esperado)
	})

	t.Run("Divide dois negativos:", func(t *testing.T) {
		resultado := Division(4, -4)
		esperado := -1
		verifica(t, resultado, esperado)
	})

	t.Run("Divide dois numeros grandes:", func(t *testing.T) {
		resultado := Division(9000000000000000, 8888888888888888)
		esperado := 1
		verifica(t, resultado, esperado)
	})
}
