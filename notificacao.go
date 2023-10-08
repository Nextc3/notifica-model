package notifica_model

type Asset struct {
	Id                  int    `json:"id"`
	DocType             string `json:"docType"`
	DataNascimento      string `json:"dataNascimento"`
	Sexo                string `json:"sexo"`
	Endereco            string `json:"endereco"`
	Bairro              string `json:"bairro"`
	Cidade              string `json:"cidade"`
	Estado              string `json:"estado"`
	Pais                string `json:"pais"`
	Doenca              string `json:"doenca"`
	DataInicioSintomas  string `json:"dataInicioSintomas"`
	DataDiagnostico     string `json:"dataDiagnostico"`
	DataNotificacao     string `json:"dataNotificacao"`
	InformacoesClinicas string `json:"informacoesClinicas"`
}
