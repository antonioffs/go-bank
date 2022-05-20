package main

import (
	"fmt"
	"study/banco/clientes"
	"study/banco/contas"
)

func main() {

	// conta := ContaCorrente{
	// 	titular:       "Felipe",
	// 	numeroAgencia: 123,
	// 	numeroConta:   123456,
	// 	Saldo:         500,
	// }

	// fmt.Println(reflect.TypeOf(conta))

	// fmt.Println(conta)
	// fmt.Println(conta.Sacar(500))
	// fmt.Println(conta.Saldo)

	// fmt.Println(conta)
	// fmt.Println(conta.Depositar(500.0))
	// fmt.Println(conta.Saldo)

	// --------------------------------------------------------

	// contaFelipe := contas.ContaCorrente{
	// 	Titular:       "Felipe",
	// 	NumeroAgencia: 123,
	// 	NumeroConta:   123456,
	// 	Saldo:         500,
	// }

	// contaNatalia := contas.ContaCorrente{
	// 	Titular:       "Natalia",
	// 	NumeroAgencia: 123,
	// 	NumeroConta:   123456,
	// 	Saldo:         1000,
	// }

	// fmt.Println(contaFelipe)
	// fmt.Println(contaNatalia)

	// status := contaNatalia.Transferir(-200, &contaFelipe)

	// fmt.Println(status)
	// fmt.Println(contaFelipe)
	// fmt.Println(contaNatalia)

	// fmt.Println(reflect.TypeOf(conta))

	// --------------------------------------------------------

	// cliente := clientes.Titular{
	// 	Nome:      "Felipe",
	// 	CPF:       "43327439893",
	// 	Profissao: "Dev",
	// }

	// conta := contas.ContaCorrente{
	// 	Titular:       cliente,
	// 	NumeroAgencia: 123,
	// 	NumeroConta:   123456,
	// }

	// conta.Depositar(500)

	// fmt.Println(conta)

	// fmt.Println(conta.ObterSaldo())

	// --------------------------------------------------------

	cliente := clientes.Titular{
		Nome:      "Felipe",
		CPF:       "12312312312",
		Profissao: "Dev",
	}

	conta := contas.ContaCorrente{
		Titular:       cliente,
		NumeroAgencia: 123,
		NumeroConta:   123456,
	}

	conta.Depositar(500)
	fmt.Println(conta.ObterSaldo())
	PagarBoleto(&conta, 200)
	fmt.Println(conta.ObterSaldo())

}

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}
