package notifica_model

type Notificacao struct {
	Id                  int    `json:"id"`
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
