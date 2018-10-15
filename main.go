package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

type Consumidor struct {
	cpf string   `json:"CPF"`
	private  string   `json:"Private"`
	incompleto  string   `json:"Incompleto"`
	compras   *Compras `json:"Compras,omitempty"`
}

type Compras struct {
	datadaultimacompra  string `json:"DataUltimacompra"`
	ticketmedio string `json:"TicketMedio"`
	ticketdaultimaCompra string `json:"TicketUltimaCompra"`
	lojamaisfrequente string `json:"LojaMaisFrequente"`
	lojadaultimacompra string `json:"LojaUltimaCompra"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {
	// Opens connection to DB
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
	defer db.Close()

	sqlStatement_truncate := `truncate table base_teste`
	_, err = db.Exec(sqlStatement_truncate)
	
	// Opens input file
	filename := "base_teste_original.txt"
	fileHandle, _ := os.Open(filename)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)


	i := true
	// Iterates over input file
	for fileScanner.Scan() {
		if i {
			i=false
			continue
		}
		////Metodo1
		//// Substitui multiplos espaços por espaço unicoe quebra a string
		//// Considerando o tipo do arquivo, é possivel desta forma, porem se for adicionada uma coluna eg.: "Nome completo" em que espaços fariam parte da string?
		//rp := regexp.MustCompile("( +)") //procura por N" "
		//cleanLine := rp.ReplaceAllString(fileScanner.Text(), " ") //substitui N" " por " "
		//dbValues := strings.Split(cleanLine, " ")
		//

		////Metodo2
		//cpf := string(([]cpf(fileScanner.Text()))[0:19])
		cpf := fileScanner.Text()[0:19]
		private := fileScanner.Text()[20:30]
		incompleto := fileScanner.Text()[31:42]
		datadaultimacompra := fileScanner.Text()[43:65]
		ticketmedio := fileScanner.Text()[65:85]
		ticketdaultimaCompra := fileScanner.Text()[86:110]
		lojamaisfrequente := fileScanner.Text()[111:129]
		lojadaultimacompra := fileScanner.Text()[130:]

		// Inserts values from file into the db
		sqlStatement_insertloop := `
			INSERT INTO base_teste (cpf, private, incompleto, datadaultimacompra, ticketmedio, ticketdaultimacompra, lojamaisfrequente, lojadaultimacompra)
			VALUES (
			CAST(NULLIF(regexp_replace($1, '[^0-9]*','','g'), '') as VARCHAR(30)), --ok
			CASE WHEN $2 LIKE '%TRUE%' THEN true ELSE false END, --ok
			CASE WHEN $3 LIKE '%TRUE%' THEN true ELSE false END, --ok 
			CASE WHEN $4 LIKE '%NULL%' THEN null ELSE to_date(REPLACE($4,' ',''),'YYYY/MM/DD') END, --ok
			CAST(CASE WHEN $5 LIKE '%NULL%'THEN null ELSE REPLACE($5,',','.') END as NUMERIC), --ok
			CAST(CASE WHEN $6 LIKE '%NULL%'THEN null ELSE REPLACE($6,',','.') END as NUMERIC), --ok
			CASE WHEN $7 LIKE '%NULL%' THEN null else NULLIF(regexp_replace($7, '[^0-9]*','','g'), '') END,
			CASE WHEN $8 LIKE '%NULL%' THEN null ELSE NULLIF(regexp_replace($8, '[^0-9]*','','g'), '') END 
			)
`
		_, err = db.Exec(sqlStatement_insertloop,
			////Metodo1
			//dbValues[0],
			//dbValues[1],
			//dbValues[2],
			//dbValues[3],
			//dbValues[4],
			//dbValues[5],
			//dbValues[6],
			//dbValues[7]			)
			////Metodo2
			cpf,
			private,
			incompleto,
			datadaultimacompra,
			ticketmedio,
			ticketdaultimaCompra,
			lojamaisfrequente,
			lojadaultimacompra		)
	}
		if err != nil {
			panic(err)
		}
	//remove os cpf invalidos baseado em fn no postgre, poderia ter sido checado no serviço
	sqlStatement_remove_invalid_cpf := `DELETE FROM base_teste WHERE fn_cnpj_cpf( NULLIF(regexp_replace(CPF, '[^0-9]*','','g'), '')) is FALSE`
	_, err = db.Exec(sqlStatement_remove_invalid_cpf)
	}
