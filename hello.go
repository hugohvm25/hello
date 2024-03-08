/*
	esse formato ele identifica o tipo da variavel que foi declarada anteriormente
	sem que precise ser identificada com "%d" para inteiro ou "%f" para float ou "%s" para string
	%d codigo para modificar o de input de usuário serve para identificar o tipo de variavel que esta sendo escolhida
	e "&"" significa o endereço da variavel que acompanha
	fmt.Scanf("%d", &comando) forma mais completa de obter input de usuario;


	## IF

	if valor == 1 {
		etc...
	} else {
		etc...
	}
	no GO sempre deve ser incluida uma expressão booleana para verificar true or false
	no GO usamos "else if"

	## OPÇÃO COM IF - ELSE

	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo logs:")
	} else if comando == 0 {
		fmt.Println("Programa encerrado.")
	} else {
		fmt.Println("Opção incorreta, favor escolher novamente.")
	}


	## ESCREVENDO VARIÁVEIS EM GO

	nome := "Hugo"            => variavel pode ser escrita como var "nome da variavel" tipo da variavel
	var versao float32 = 1.1  => pode ser escrita também com a forma reduzida "nome da variavel" :=


	=================================
	COMENTAR BLOCO VS CODE => CTRL+;
	LIMPAR TERMINAL VS CODE => CTRL+L
	=================================


	## FORMA EXTENSA PARA EXIBIR A VARIÁVEL E SOLICITAR INPUT

	var comando int
	fmt.Scan(&comando)


	## FOR ===  NÃO EXISTE WHILE EM GOLANG ===

	o FOR sem alguma instrução depois da função, tem a propriedade do WHILE e faz um sistema de loop infinito


	## SLICE EM GO

	tem a mesma propriedade do array mas não é limitado pelo tamanho
	Ex de Array simples em GO:
	var nomes [3]string 		 => devemos sempre informa o tamanho do array no colchete e o tipo de variável em seguida
	nome[0] = "hugo"  			 => devemos sempre identificar a posição do dado no colchete da variável
	nome[1] = "pedro"
	nome[2] = "maria"

	Já em Slice, não precisamos informar o tamanho
	Ex de Slice em GO:
	nomes := []string{"hugo", "pedro", "maria"}      => no slice devemos deixar em branco pois é variável o tamanho,
															informar o tipo da variável e informar os dados dentro de chaves

	Podemos incluir dados no slice com APPEND e nesse caso quando estoura o tamanho informado anteriormente nas {}, o GO duplica o tamanho do array AUTOMATICAMENTE
	porque ele não quer que nos preocupemos com o espaço de alocação dos dados e sim com a manipulação dos dados

	## RANGE

	O range tem a propriedade de acessar os dados do slice e, também, informar a posição se for necessário

	sintaxe:
	for i, variável := range slice {}

*/

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

const monitoramentos = 2
const quantidades = 4
const delay = 2

//import que informa ao sistema operacional a saida correta do programa

func main() {

	// chamada das funções modulares
	intro()

	for {

		exibirMenu()

		comando := lerComando()

		switch comando { // no switch em go não tem a função break para interromper a execução
		case 1:
			monitoramentoIndividual()
		case 2:
			adicionarSite()
		case 3:
			monitoramentoSites()
		case 4:
			fmt.Println("Exibindo logs:")
			imprimeLogs()
		case 0:
			fmt.Println("Programa encerrado.")
			os.Exit(0)
			// sempre que usar o "os" deve ser informado a opção de "saída" para informar corretamente ao sistema operacional
		default:
			// retorna a origem semelhante ao else
			fmt.Println("Opção incorreta, favor escolher novamente.")
		}
	}
}

func intro() {

	var nome string
	versao := 1.1

	fmt.Println("Olá, qual seu nome?")
	fmt.Scan(&nome) //input do usuário
	fmt.Println("")

	fmt.Println("Boa noite", nome)
	fmt.Println("Este programa está na versão:", versao)
	fmt.Println("")
}

func exibirMenu() {

	fmt.Println("1 - Monitoramento individual")
	fmt.Println("2 - Adicionar site ao arquivo")
	fmt.Println("3 - Ler arquivo de sites")
	fmt.Println("4 - Exibir logs")
	fmt.Println("0 - Encerrar programa")
	fmt.Println("")
}

func lerComando() int {
	// a função precisa ter o tipo da variável para identificar qual tipo de input que o usuário vai informar
	var comando int

	fmt.Println("Escolha uma opção:")
	fmt.Scan(&comando) // input do usuário para a escolha da opção do menu
	fmt.Println("A opção escolhida foi", comando)

	return comando
}

func monitoramentoIndividual() {

	var site string

	fmt.Println("Digite o endereço do site que deseja monitorar:")
	fmt.Scan(&site) // input do usuário informando o site para monitoramento
	fmt.Println("Monitorando...")

	for i := 0; i < monitoramentos; i++ {
		/*	estou determinando que seja feita o monitoramento do valor i que for informado no máximo até a
			quantidade da constante definida e incrementar  */
		for i := 0; i < quantidades; i++ {
			// após os monitoramentos terem acontecido, ele vai contar o delay e testar novamente
			teste1(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func monitoramentoSites() {

	fmt.Println("Monitorando...")

	//sites := lerArquivo()

	sites, err := lerArquivo()
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	for i := 0; i < monitoramentos; i++ {
		/*	estou determinando que seja feita o monitoramento do valor i que for informado no máximo até a
			quantidade da constante definida e incrementar  */
		for _, site := range sites {
			// após os monitoramentos terem acontecido, ele vai contar o delay e testar novamente
			teste2(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

}

func teste1(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O site", site, "está funcionando corretamente")
		registraLog(site, true)
	} else {
		fmt.Println("O site", site, "está com problemas.")
		registraLog(site, false)
		fmt.Println("")
	}
}

func teste2(sites string) {

	resp, err := http.Get(sites)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
		return
	}

	if resp.StatusCode == 200 {
		fmt.Println("O site", sites, "está funcionando corretamente")
		registraLog(sites, true)
	} else {
		fmt.Println("O site", sites, "está com problemas.")
		registraLog(sites, false)
		fmt.Println("")
	}
}

// func lerArquivo() []string {

// 	var sites []string

// 	arquivo, err := os.Open("sites.txt") //abre o arquivo de texto

// 	if err != nil {
// 		fmt.Println("Ocorreu um erro", err)
// 	}

// 	leitor := bufio.NewReader(arquivo) //lê o arquivo de texto
// 	for {
// 		// irá fazer a leitura linha a linha do arquivo e printar cada uma no terminal até chegar ao fim do arquivo EOF (end of file) e dar o braak e sair do loop FOR
// 		linha, err := leitor.ReadString('\n')
// 		linha = strings.TrimSpace(linha)
// 		sites = append(sites, linha) // lê o arquivo e cria um slice(array) de sites
// 		if err == io.EOF {
// 			break
// 		}
// 	}

// 	arquivo.Close()
// 	return sites
// }

func lerArquivo() ([]string, error) {

	var linhas []string

	arquivo, err := os.Open("sites.txt") // Abre o arquivo

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}
	defer arquivo.Close() //garante que o arquivo seja fechado após o uso

	leitor := bufio.NewScanner(arquivo) // le cada linha do arquivo
	for leitor.Scan() {
		linhas = append(linhas, leitor.Text())
	}
	return linhas, leitor.Err()
}

func adicionarSite() error {

	var site string

	fmt.Println("Favor informar o site que deseja adicionar ao arquivo?")
	fmt.Scan(&site) // input do usuário que informa o site que será incluído no arquivo

	/*
		Abre o arquivo "sites.txt" para escrita.
		os.OpenFile => é usado para abrir ou criar um arquivo.
		os.O_APPEND => indica que novos dados devem ser adicionados ao final do arquivo sem substituir os dados anteriores,
		os.O_WRONLY => indica que o arquivo será aberto para escrita apenas,
		os.O_CREATE => indica que o arquivo será criado se não existir.
		0644 é um valor octal que define permissões de acesso ao arquivo.
	*/
	arquivo, err := os.OpenFile("sites.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer arquivo.Close()

	_, err = arquivo.WriteString(site + "\n")
	if err != nil {
		return err
	}

	fmt.Println("Site adicionado com sucesso.")
	return nil
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - Online - " + strconv.FormatBool(status) + "\n")
	// para gerar o formato de data e hora, deve consultar a documentação do GO no site para verificar quais os códigos dos formatos
	// https://go.dev/src/time/format.go

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	println(string(arquivo))

}
