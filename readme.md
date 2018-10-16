---


---

<h1 id="csv-data-importer">.csv Data Importer</h1>
<p>Este serviço, tem por objetivo importar dados de um tipo especifico de arquivo .csv  pré definido (separado por múltiplos espaços / posição de carácter fixa por coluna também pré-definida.</p>
<p>Desenvolvido em Go, com utilização das libs externas “database/sql” e “<a href="http://github.com/lib/pq">github.com/lib/pq</a>” e versão mais recente do Postgres.</p>
<h3 id="instalação-e-configuração">Instalação e configuração:</h3>
<ol>
<li>Go <a href="https://golang.org/dl/">https://golang.org/dl/</a></li>
<li>Importação das go libs externas (cmd/go get /libname/)<br>
go get “database/sql”<br>
go get “<a href="http://github.com/lib/pq">github.com/lib/pq</a>”</li>
<li>Postgres <a href="https://www.postgresql.org/download/">https://www.postgresql.org/download/</a></li>
<li>Para utilização em localhost, durante a instalação, manter configuração de portas default e configurar senha “postgres”.</li>
</ol>
<ul>
<li>Executar db create.sql <a href="https://github.com/willpioli/NEO-PROJECT/blob/master/db%20create.sql">https://github.com/willpioli/NEO-PROJECT/blob/master/db create.sql</a> para criação da db local e fn para validação de cpf em localhost.</li>
</ul>
<ol start="5">
<li>Para utilização em outro dispositivo, deve se instalar o driver <a href="https://github.com/npgsql/Npgsql/releases">https://github.com/npgsql/Npgsql/releases</a> e realizar devida configuração em “main.go” (linhas 26 a 31)</li>
<li>Apontar arquivo .csv em linha 49</li>
</ol>
<h3 id="tipo-do-arquivo">Tipo do arquivo:</h3>
<ol>
<li>.csv</li>
<li>Colunas e posicoes:<br>
cpf := [0:19]<br>
private := [20:30]<br>
incompleto := [31:42]<br>
datadaultimacompra := [43:65]<br>
ticketmedio := [65:85]<br>
ticketdaultimaCompra := [86:110]<br>
lojamaisfrequente := fileScanner.Text()[111:129]<br>
lojadaultimacompra := fileScanner.Text()[130:]</li>
</ol>
<p>Para utilização em outro tipo de arquivo, deve ser refeita o apontamento das posições das colunas, que devem estar devidamente declaradas (linhas 72 a 79) e apontadas no comando sqlstatment_insertloop (linha 82) e execução do mesmo (linha95).</p>

