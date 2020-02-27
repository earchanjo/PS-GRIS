package main

import (
	"fmt"
	"context"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

type PortScanner struct {			// Criando nosso scanner 
	ip string
	lock *semaphore.Weighted		// Nosso limitador de sequencia de uso de recursos do programa
}


// --- A funcao ulimit executa o comando ulimit no shell, que retorna o numero máximo de arquivos/programas aberto simultanemente no sistema
func ulimit() int64 {
	out, err := exec.Command("/bin/sh", "-c", "ulimit -n").Output()

	if err != nil {
		panic(err)
	}

	str := strings.TrimSpace(string(out))	// Retorna a string sem os caracteres especiais

	i, err := strconv.ParseInt(str, 10, 64)		// Converte o numero maximo retornado em string para inteiro na base 10 em formato 64bits
	
	if err != nil {
		panic(err)
	}

	return i
}

// --- Funcao que escaneia as portas
// --- Rece o ip, a porta e o timeout como parâmetros

func scan_port(ip string, port int, timeout time.Duration) {
	tgt := fmt.Sprintf("%s:%d", ip, port)			// Junta o ip com a porta nos dois pontos para poder ser usado na conexao
	connection, err := net.DialTimeout("tcp", tgt, timeout)		// Faz a conexao com o ip e porta especificado

	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {		// Analisa o erro, checa se há algum erro que corresponda apenas a sobrecarga de conexoes
		time.Sleep(timeout)				// Se há muitas conexões é dado um time até a proxima a fim de desafogar 
		scan_port(ip, port, timeout)    // Realiza novamente o scan da porta requerida
		
		} else {
			fmt.Println(port, "Closed") 	// Se o erro nao é do tipo de sobrecarga, entao a porta está fechada
		}
		return
	}
	connection.Close()				// Fecha a conexao e atribui que a porta está aberta, pois nao houve erro na conexao
	fmt.Println(port, "Open")

}

// --- Método do PortScanner que inicia o processo de scan
// --- Recebe como parametro a faixa de porta e o timeout
func (ps *PortScanner) start(f, l int, timeout time.Duration) {
	wg := sync.WaitGroup{}

	defer wg.Wait()

	for port := f; port <= l; port++ {
		ps.lock.Acquire(context.TODO(), 1)
		wg.Add(1)

		go func(port int) {
			defer ps.lock.Release(1)
			defer wg.Done()
			scan_port(ps.ip, port, timeout)
			
		}(port)
	}
}


func main() {
	ps := &PortScanner{
		ip: "127.0.0.1",
		lock: semaphore.NewWeighted(ulimit()),
	}

	ps.start(1, 65535, 500*time.Millisecond)
}






