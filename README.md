# TESTING

![](https://raw.githubusercontent.com/gist/brudnak/efd7b887bd7c0441d8bb88ae1c77374a/raw/e96188dae6a84223fef3ff4e834d79f75680a094/gopher-workout.gif)

É muito provável que você já tenha ouvido falar de testes então vamos dar uma molhada nos pés aqui.

O teste em si serve pra você verificar a integridade do seu software e por meio desse garantir que funcione sem ou com o mínimo de efeitos colaterais. Em outras palavras, você precisa garantir que seu código faz o que você diz que ele faz.

Existe uma infinidade de convenções a respeito de testes em várias comunidades, e pode ser que em cada uma delas os testes sejam feitos de formas diferentes.

Meu foco não será na técnica e sim em um tipo especifico de teste, o _Teste de Unidade_ ou _Unit Test_, onde se verifica a integridade da menor unidade dentro do software, uma função por exemplo.

Ao final do desenvolvimento de uma feature é muito comum que se tenha dezenas de testes de unidade inclusive mais código de teste do que de feature, isso porque o teste de unidade é muito simples e rápido de ser escrito.

Por ser rápido e simples costumamos chamar esse tipo de teste de “barato”, pois demanda pouco esforço durante do desenvolvimento.

Podemos mensurar o custo dos tipos de teste através da pirâmide de testes.

![](https://cdn-images-1.medium.com/max/800/1*sTglc-qAzLcZRTl_CkuphQ.png)

Geralmente no meio da pirâmide tem outros tipos, essa é uma versão mais simplificada.

**Teste de Integração:** É definido como um tipo de teste onde os módulos do software são testados como um grupo. Um projeto de software típico consiste em vários módulos de software, codificados por diferentes programadores. O objetivo deste nível de teste é expor defeitos na interação entre esses módulos quando eles são integrados.

**E2E** (End to End) O objetivo do teste de ponta a ponta é testar tanto o software quanto as dependências, integridade de dados e comunicação com outros sistemas, interfaces e bancos de dados para exercer a produção completa como cenário.

Existem ainda mais alguns, aqui eu vou falar mais dos unitários pois quem está aprendendo a programar começa com programas menores então nem faz muito sentido algo como E2E.

Na maioria das empresas é o mais utilizado justamente por ter um custo menor, pelo menos é o que dizem hahaha. Inclusive algumas empresas nem testam…

# Como funciona?

Matemática básica é sempre muito simples de se entender então acredito que a função mais simples de uma calculadora como, soma é um bom começo, se você quiser pode complicar depois, mas não seremos rasos aqui.

A questão aqui é saber quais os dados de entrada e saída, o teste é basicamente uma afirmação do que você deseja como software.

**EX**: _Entrada{2,2},_ _Esperado{4}_ para uma função de soma, é bobinho a principio, mas boa parte das coisas sobre software são assim, as pessoas que estragam tudo criando complexidade desnecessária, muitos iniciantes evitam testes no inicio por isso.

Vamos dizer que alguém escreveu a função **Sum()** e de alguma forma onde temos _Esperado{4}_ o teste nos retorna _Obteve{0}_, provável que a função tenha sido mal escrita e precisa ser corrigida, então depois de mudanças finalmente ela retorna _Obteve{4}_ **PASS**, significa que obtivemos sucesso em fazer com que a nossa função se comportasse como era esperado.

Vamos a um exemplo com código um pouco mais concreto:

Essa é a nossa função de soma, muito simples nada demais.

```go
func Sum(a, b int) int {
	return a + b
}
```

E esse é o nosso teste:

```go
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
```

### Dissecação

Disse que era bem simples e isso na verdade é 10x maior que a função que estamos testando haha, mas agora que você sabe o conceito eu vou explicar a parte prática e tu vai ver que é tranquilo.

Uma das nossas opções, seria algo como isso:

```go
func TestSum(t *testing.T) {
	soma := Sum(10, 10)
	if soma != 20 {
		t.Error("Falhou no teste!")
	}
}
```

Assim é bem mais simples, apesar de ser um teste feito na mesma unidade (**Sum**) temos apenas um bloco de código. Todo teste tem um prefixo [**Test**](https://golang.org/src/testing/testing.go), seguido do nome da função que queremos testar.

[_func TestXxx(*testing.T)_](https://golang.org/pkg/testing/#T)

Um parâmetro com um ponteiro importado do pacote “testing” o [**type T**](https://golang.org/src/testing/testing.go?s=23377:23479#L647) possui três propriedades:


**_common_**: É outra struct que compartilha os elementos em comum entre [**type T**](https://golang.org/src/testing/testing.go?s=23377:23479#L647) e [**type B**](https://golang.org/src/testing/benchmark.go?s=2453:3341#L82) essa segunda serve para funcionalidades de _benchmark_ que falarei sobre futuramente. O **type common** é usado como receiver em várias funções presentes na [**interface TB**](https://golang.org/src/testing/testing.go?s=22280:22856#L610) que agrega todas as funções que utilizam dados de **common**, sendo a maioria de _log_, _error_ e _fatal_.

**_isParallel_**: É um tipo booleano, que sinaliza se os testes estão rodando em paralelo ou não.

**_context_**: Um ponteiro de **testContext** que define testes em paralelo, channels e coisas como onde começam e onde terminam usando o pacote Time.

Esse miolo que é onde entra a especificação do que seria a entrada de dados, a função **Sum** com os parâmetros é declarada em uma variável e em seguida perguntamos, **soma** é diferente de **20**?

```go
	soma := Sum(10, 10)
	if soma != 20
```

Por final temos o que seria a saída, a função [**Error**](https://golang.org/src/testing/testing.go?s=27660:27703#L776) que vem da [**interface TB**](https://golang.org/src/testing/testing.go?s=22280:22856#L610) usa a **type common** como receiver.

_t.Error(“Falhou no teste!”)_


### Analisando o conjunto

Agora que temos uma ideia de como funciona a estrutura de um teste vamos ver com mais detalhes aquele primeiro maior mas por partes, vamos saindo do micro para o macro.

Esse teste possui uma estrutura parecida, podemos ver que a função [**Run**](https://golang.org/src/testing/testing.go?s=37631:37678#L1125) vem da mesma struct que estávamos falando [**type T**](https://golang.org/src/testing/testing.go?s=23377:23479#L647) então é muito provável que ela venha da [**interface TB**](https://golang.org/src/testing/testing.go?s=22280:22856#L610), mas não, ela tem apenas [**type T**](https://golang.org/src/testing/testing.go?s=23377:23479#L647) como receiver existe outra função [**Run**](https://golang.org/src/testing/benchmark.go?s=17060:17107#L596) que tem como receiver [**type B**](https://golang.org/src/testing/benchmark.go?s=2453:3341#L82) mas essa fica no arquivo [_benchmark.go_](https://golang.org/src/testing/benchmark.go) do pacote [_testing_](https://golang.org/src/testing/testing.go).

A estrutura de [**Run**](https://golang.org/src/testing/testing.go?s=37631:37678#L1125) é assim:

_func (t *T) Run(_**_name_** _string,_ **_f_** _func(t *T))_

Name: “Soma dois inteiros”  
f: func TestSum()

Temos um miolo com uma separação mais interessante, aqui fica muito claro que **resultado** é o retorno de **Sum()** e nosso **esperado** é um **int** ao invés de colocar o numero _hardcoded_ numa condicional.


```go
t.Run("Soma dois inteiros", func(t *testing.T) {
	resultado := Sum(2, 2)
	esperado := 4
	verifica(t, resultado, esperado)
})
```

Agora temos **verifica**(), vamos dar uma olhada melhor nela:

```
verifica  :=  func(t *testing.T, resultado, esperado int)
```

**verifica** recebe uma função que tem [**type T**](https://golang.org/src/testing/testing.go?s=23377:23479#L647), **Sum**() e o retorno de **Sum**() como argumentos.

**Helper**() faz parte de **common** então temos duas versões uma com [**type T**](https://golang.org/src/testing/testing.go?s=23377:23479#L647) e outra [**type B**](https://golang.org/src/testing/benchmark.go?s=2453:3341#L82) ela serve para marcar a sua função como auxiliar o que é justamente o caso.

A condicional é usada em **verifica** para que todas as condições de teste que você deseja não precisem aumentar a complexidade.

[**Errorf()**](https://golang.org/src/testing/testing.go?s=27799:27858#L782) indica que temos uma saída formatada, ela funciona como [**Logf**()](https://golang.org/src/testing/testing.go?s=27513:27570#L773) porem tem uma chamada de [**Fail**()](https://golang.org/src/testing/testing.go?s=23815:23838#L670) que serve pra marcar o teste como falho, mas continua a execução para que os demais testes terminem.

E então como argumento de [**Errorf()**](https://golang.org/src/testing/testing.go?s=27799:27858#L782) temos as saídas que irão nos dizer o que o teste nos retornou e o que ele esperava, porém agora numa função auxiliar dedicada a dar a saída de dados de todos os test cases.


```go
verifica := func(t *testing.T, resultado, esperado int) {
	t.Helper()
	if resultado != esperado {
		t.Errorf("resultado '%d esperado '%d'", resultado, esperado)
	}
}
```

O que parecia complicado, já não é mais tão complicado assim, teste analisar o mesmo código agora conhecendo a base.

```go
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
```

Agora que você leu até aqui eu tenho duas coisas pra dizer, uma é que não terminei de falar sobre testes ainda e tem mais conteúdo pela frente, a segunda é que eu gostaria do seu feedback aqui nos comentários ou você pode me chamar no meu [Twitter](https://twitter.com/Vapordev1).

Eu quero saber é a sua senioridade e qual a adequação da minha abordagem e linguagem durante não apenas esse artigo mas também em relação aos demais, assim posso ter uma ideia mais clara de como devo abordar os assuntos por aqui.