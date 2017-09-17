package integrations

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wilsontamarozzi/panda-api/helpers"
	"github.com/wilsontamarozzi/panda-api/models"
	"github.com/wilsontamarozzi/panda-api/repositories"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	PATH                    = "/"
	DESTINATION_FOLDER      = "temp/"
	FILENAME_EXTRATO        = "Extrato"
	FILENAME_PAX            = "Pax"
	FILENAME_TIPO_MOVIMENTO = "TipoMovimento"
	FILENAME_VENDEDOR       = "Vendedor"
)

var (
	FILES_SELLERS []arquivoVendedor
	USER_REQUEST  string
)

type arquivoRecibo struct {
	numeroRecibo                    string
	dataLancamento                  string
	dataReserva                     string
	dataCancelamento                string
	codigoPessoaContratante         string
	tipoPessoaContratante           string
	codigoCPFCNPJ                   string
	codigoCPFCNPJFilial             string
	codigoCPFCNPJDigito             string
	codigoTipoMovimento             string
	codigoTipoVenda                 string
	codigoExtrato                   string
	codigoPessoaFilial              string
	codigoPessoaVendedor            string
	codigoPessoaIntermediario       string
	valorLancamento                 string
	valorTaxa                       string
	valorDesconto                   string
	valorAbatimento                 string
	valorComissaoCalculada          string
	valorComissaoRetida             string
	valorDeposito                   string
	valorOPFAX                      string
	valorSaldo                      string
	contratanteNome                 string
	contratanteSexo                 string
	contratanteDataNascimento       string
	contratanteTipoEndereco         string
	contratanteNomeEndereco         string
	contratanteNumeroEndereco       string
	contratanteNomeBairro           string
	contratanteDescricaoComplemento string
	contratanteNomeCidade           string
	contratanteUF                   string
	contratanteCEP                  string
	contratanteEmail                string
	contratanteRG                   string
	contratanteIE                   string
	contratanteDDI                  string
	contratanteDDD                  string
	contratanteTelefone             string
	dataEmbarque                    string
	dataRetorno                     string
	codigoExcursao                  string
	tipoRoteiro                     string
	descricaoRoteiro                string
	tipoProduto                     string
	codigoProduto                   string
	descricaoProduto                string
	descricaoTipoProduto            string
	numeroReciboReferencia          string
	vendedorTrocado                 string
	valorComissaoIntermediario      string
}

type arquivoTipoMovimento struct {
	codigoTipoMovimento    string
	descricaoTipoMovimento string
	dataAtualizacao        string
}

type arquivoVendedor struct {
	codigoVendedor       string
	situacaoHabilitacao  string
	codigoFilial         string
	codigoPessoa         string
	tipoPessoa           string
	codigoCPFCNPJ        string
	codigoCNPJFilial     string
	codigoCPFCNPJDigito  string
	nomePessoa           string
	nomeFantasia         string
	dataNascimento       string
	codigoCategoria      string
	descricaoCategoria   string
	dataAtualizacao      string
	tipoEndereco         string
	nomeEndereco         string
	numeroEndereco       string
	nomeBairro           string
	descricaoComplemento string
	nomeCidade           string
	uf                   string
	codigoCEP            string
	email                string
	rg                   string
	ie                   string
	numeroDDI            string
	numeroDDD            string
	numeroTelefone       string
}

type arquivoPAX struct {
	codigoReserva         string
	codigoContratante     string
	codigoPassageiro      string
	numeroRecibo          string
	nomePassageiro        string
	numeroIdade           string
	dataNascimento        string
	passageiroSexo        string
	codigoRG              string
	descricaoOrgaoEmissor string
	codigoUFEmissao       string
	observacao            string
}

type IntegrationCVC struct {
	PersonRepository      repositories.PersonRepository
	SaleRepository        repositories.SaleRepository
	SaleProductRepository repositories.SaleProductRepository
	ProductRepository     repositories.ProductRepository
}

func clientConfigImport() helpers.FTPClient {
	log.Println("Lendo configurações FTP")
	return helpers.FTPClient{
		Address:        "localhost:21",
		Username:       "ftp",
		Password:       "123",
		FtpModePassive: true,
	}
}

func (cvc *IntegrationCVC) Import(c *gin.Context) {
	enableLog(true)
	USER_REQUEST = c.MustGet("userRequest").(string)

	ftp := clientConfigImport()
	log.Println("Copiando arquivos FTP")
	err := ftp.DownloadFiles(PATH, DESTINATION_FOLDER)
	if err != nil {
		c.JSON(500, err)
		return
	}

	files := helpers.ListFileFromDiretory(DESTINATION_FOLDER)
	cvc.readAllArquivoVendedor(files)
	//cvc.readAllArquivoPAX(filesPath)
	cvc.readAllArquivoExtrato(files)

	c.JSON(200, "Importado com sucesso!")
}

func (cvc *IntegrationCVC) readArquivoExtrato(filePath string) []arquivoRecibo {
	linesFileExtrato, err := helpers.ReadFileLines(filePath)
	if err != nil {
		log.Fatalf("readLines: %s", err)
		return nil
	}

	log.Println("Carregando dados extrato")
	log.Printf("Quantidade de registros: %v", len(linesFileExtrato))

	var fileReceipts []arquivoRecibo
	for i, line := range linesFileExtrato {
		if !isFirstLine(i) {
			columns := strings.Split(line, ";")
			if isDefaultExtractFile(columns) {
				fileReceipts = append(fileReceipts, arquivoRecibo{
					numeroRecibo:                    columns[0],
					dataLancamento:                  columns[1],
					dataReserva:                     columns[2],
					dataCancelamento:                columns[3],
					codigoPessoaContratante:         columns[4],
					tipoPessoaContratante:           columns[5],
					codigoCPFCNPJ:                   columns[6],
					codigoCPFCNPJFilial:             columns[7],
					codigoCPFCNPJDigito:             columns[8],
					codigoTipoMovimento:             columns[9],
					codigoTipoVenda:                 columns[10],
					codigoExtrato:                   columns[11],
					codigoPessoaFilial:              columns[12],
					codigoPessoaVendedor:            columns[13],
					codigoPessoaIntermediario:       columns[14],
					valorLancamento:                 columns[15],
					valorTaxa:                       columns[16],
					valorDesconto:                   columns[17],
					valorAbatimento:                 columns[18],
					valorComissaoCalculada:          columns[19],
					valorComissaoRetida:             columns[20],
					valorDeposito:                   columns[21],
					valorOPFAX:                      columns[22],
					valorSaldo:                      columns[23],
					contratanteNome:                 columns[24],
					contratanteSexo:                 columns[25],
					contratanteDataNascimento:       columns[26],
					contratanteTipoEndereco:         columns[27],
					contratanteNomeEndereco:         columns[28],
					contratanteNumeroEndereco:       columns[29],
					contratanteNomeBairro:           columns[30],
					contratanteDescricaoComplemento: columns[31],
					contratanteNomeCidade:           columns[32],
					contratanteUF:                   columns[33],
					contratanteCEP:                  columns[34],
					contratanteEmail:                columns[35],
					contratanteRG:                   columns[36],
					contratanteIE:                   columns[37],
					contratanteDDI:                  columns[38],
					contratanteDDD:                  columns[39],
					contratanteTelefone:             columns[40],
					dataEmbarque:                    columns[41],
					dataRetorno:                     columns[42],
					codigoExcursao:                  columns[43],
					tipoRoteiro:                     columns[44],
					descricaoRoteiro:                columns[45],
					tipoProduto:                     columns[46],
					codigoProduto:                   columns[47],
					descricaoProduto:                columns[48],
					descricaoTipoProduto:            columns[49],
					numeroReciboReferencia:          columns[50],
					vendedorTrocado:                 columns[51],
					valorComissaoIntermediario:      columns[52],
				})
			} else {
				log.Printf("Arquivo não está no padrão\n%s", filePath)
			}
		}
	}
	return fileReceipts
}

func (cvc *IntegrationCVC) readAllArquivoExtrato(filesPath []string) {
	var filesExtracts []arquivoRecibo
	for _, filePath := range filesPath {
		filename := strings.Split(filePath, "_")
		switch filename[0] {
		case FILENAME_EXTRATO:
			log.Println("Importando extrato")
			filesExtracts = append(filesExtracts, cvc.readArquivoExtrato(DESTINATION_FOLDER+filePath)...)
		}
	}
	cvc.createAllProducts(filesExtracts)
	cvc.createAllBuyer(filesExtracts)
	//cvc.readAllArquivoPAX(filesPath)
	cvc.createAllSales(filesExtracts)
}

func (cvc *IntegrationCVC) readArquivoVendedor(filePath string) []arquivoVendedor {
	linesFileSeller, err := helpers.ReadFileLines(filePath)
	if err != nil {
		log.Printf("readLines: %s", err)
		return nil
	}

	log.Println("Carrengando dados vendedor")
	log.Printf("Quantidade de registros: %v", len(linesFileSeller))

	var fileSellers []arquivoVendedor
	for i, line := range linesFileSeller {
		if !isFirstLine(i) {
			columns := strings.Split(line, ";")
			if isDefaultSellerFile(columns) {
				fileSellers = append(fileSellers, arquivoVendedor{
					codigoVendedor:       columns[0],
					situacaoHabilitacao:  columns[1],
					codigoFilial:         columns[2],
					codigoPessoa:         columns[3],
					tipoPessoa:           columns[4],
					codigoCPFCNPJ:        columns[5],
					codigoCNPJFilial:     columns[6],
					codigoCPFCNPJDigito:  columns[7],
					nomePessoa:           columns[8],
					nomeFantasia:         columns[9],
					dataNascimento:       columns[10],
					codigoCategoria:      columns[11],
					descricaoCategoria:   columns[12],
					dataAtualizacao:      columns[13],
					tipoEndereco:         columns[14],
					nomeEndereco:         columns[15],
					numeroEndereco:       columns[16],
					nomeBairro:           columns[17],
					descricaoComplemento: columns[18],
					nomeCidade:           columns[19],
					uf:                   columns[20],
					codigoCEP:            columns[21],
					email:                columns[22],
					rg:                   columns[23],
					ie:                   columns[24],
					numeroDDI:            columns[25],
					numeroDDD:            columns[26],
					numeroTelefone:       columns[27],
				})
			} else {
				log.Printf("Arquivo não está no padrão\n%s", filePath)
			}
		}
	}
	return fileSellers
}

func (cvc *IntegrationCVC) readAllArquivoVendedor(filesPath []string) {
	var filesSellers []arquivoVendedor
	for _, filePath := range filesPath {
		filename := strings.Split(filePath, "_")
		switch filename[0] {
		case FILENAME_VENDEDOR:
			log.Println("Importando vendedor")
			filesSellers = append(filesSellers, cvc.readArquivoVendedor(DESTINATION_FOLDER+filePath)...)
		}
	}
	FILES_SELLERS = filesSellers
	cvc.createAllSellers(filesSellers)
}

func (cvc *IntegrationCVC) readArquivoPAX(filePath string) []arquivoPAX {
	linesFilePAX, err := helpers.ReadFileLines(filePath)
	if err != nil {
		log.Fatalf("readLines: %s", err)
		return nil
	}

	log.Println("Carregando dados passageiro")
	log.Printf("Quantidade de registros: %v", len(linesFilePAX))

	var filesPAX []arquivoPAX
	for i, line := range linesFilePAX {
		if !isFirstLine(i) {
			columns := strings.Split(line, ";")
			if isDefaultPAXFile(columns) {
				filesPAX = append(filesPAX, arquivoPAX{
					codigoReserva:         columns[0],
					codigoContratante:     columns[1],
					codigoPassageiro:      columns[2],
					numeroRecibo:          columns[3],
					nomePassageiro:        columns[4],
					numeroIdade:           columns[5],
					dataNascimento:        columns[6],
					passageiroSexo:        columns[7],
					codigoRG:              columns[8],
					descricaoOrgaoEmissor: columns[9],
					codigoUFEmissao:       columns[10],
					observacao:            columns[11],
				})
			} else {
				log.Printf("Arquivo não está no padrão\n%s", filePath)
			}
		}
	}
	return filesPAX
}

func (cvc *IntegrationCVC) readAllArquivoPAX(filesPath []string) {
	var filesPax []arquivoPAX
	for _, filePath := range filesPath {
		filename := strings.Split(filePath, "_")
		switch filename[0] {
		case FILENAME_PAX:
			log.Println("Importando passageiro")
			filesPax = append(filesPax, cvc.readArquivoPAX(DESTINATION_FOLDER+filePath)...)
		}
	}
	cvc.createAllPassengers(filesPax)
}

func (cvc *IntegrationCVC) createAllSellers(filesSellers []arquivoVendedor) {
	var sellers []models.Person
	for _, fileSeller := range filesSellers {
		sellers = append(sellers, buildSeller(fileSeller))
	}
	log.Println("Salvando vendedores")
	cvc.savePeople(sellers)
}

func (cvc *IntegrationCVC) createAllBuyer(filesExtracts []arquivoRecibo) {
	var buyers []models.Person
	for _, fileReceipt := range filesExtracts {
		buyers = append(buyers, buildBuyer(fileReceipt))
	}
	log.Println("Salvando compradores")
	cvc.savePeople(buyers)
}

func (cvc *IntegrationCVC) createAllSales(filesExtracts []arquivoRecibo) {
	var sales []models.Sale
	for _, fileSale := range filesExtracts {
		sales = append(sales, buildSale(fileSale))
	}
	cvc.saveSale(sales)
}

func (cvc *IntegrationCVC) createAllPassengers(filesPAX []arquivoPAX) {
	var passengers []models.Person
	for _, filePAX := range filesPAX {
		passengers = append(passengers, buildPassenger(filePAX))
	}
	log.Println("Salvando passageiros")
	cvc.savePeople(passengers)
}

func (cvc *IntegrationCVC) createAllProducts(filesExtracts []arquivoRecibo) {
	var products []models.Product
	for _, fileReceipt := range filesExtracts {
		products = append(products, buildProduct(fileReceipt))
	}
	log.Println("Salvando produtos")
	cvc.saveProduct(products)
}

func (cvc *IntegrationCVC) savePeople(people []models.Person) {
	for _, person := range people {
		person.CreatedByUUID = USER_REQUEST
		cvc.PersonRepository.Create(&person)
	}
}

func (cvc *IntegrationCVC) saveSale(sales []models.Sale) {
	log.Println("Salvando vendas")
	for _, sale := range sales {
		//Verifica se já existe vendedor
		if sale.Seller.IdCVC != nil {
			seller := cvc.PersonRepository.GetByIdCVC(*sale.Seller.IdCVC)
			if !seller.IsEmpty() {
				sale.SellerUUID = seller.UUID
			}
		}
		//Verifica se já existe pagante
		if sale.Buyer.IdCVC != nil {
			buyer := cvc.PersonRepository.GetByIdCVC(*sale.Buyer.IdCVC)
			if !buyer.IsEmpty() {
				sale.BuyerUUID = buyer.UUID
			}
		}
		for _, product := range sale.Products {
			//Verifica se já existe produto com o número de recibo
			saleTemp := cvc.SaleProductRepository.GetByDocument(product.Document)
			if saleTemp.IsEmpty() {
				if err := cvc.SaleRepository.Create(&sale); err != nil {
					log.Printf("Venda já existe: %v", sale.SaleDate)
					return
				}

				product.SaleUUID = sale.UUID
				if err := cvc.SaleProductRepository.Create(&product); err != nil {
					log.Printf("Produto já existe: %v", product.Document)
					return
				}
			}
		}
	}
}

func (cvc *IntegrationCVC) saveProduct(products []models.Product) {
	for _, product := range products {
		if !product.IsEmpty() {
			cvc.ProductRepository.Create(&product)
		}
	}
}

func buildBuyer(recibo arquivoRecibo) models.Person {
	var homePhone = fmt.Sprintf("(%s) %s", recibo.contratanteDDD, recibo.contratanteTelefone)
	var address = fmt.Sprintf("%s %s", recibo.contratanteTipoEndereco, recibo.contratanteNomeEndereco)
	var cpf *string
	var rg *string
	var birthDate *time.Time
	var idCVC *int

	if recibo.tipoPessoaContratante == "F" {
		if recibo.codigoCPFCNPJ != "" && recibo.codigoCPFCNPJDigito != "" {
			cpfConcat := fmt.Sprintf("%s%s", recibo.codigoCPFCNPJ, recibo.codigoCPFCNPJDigito)
			if len(cpfConcat) == 10 {
				cpfConcat = fmt.Sprintf("%s%s", "0", cpfConcat)
				cpf = &cpfConcat
			}
		}

		if recibo.contratanteRG != "" {
			rg = &recibo.contratanteRG
		}
	}
	if recibo.contratanteDataNascimento != "" {
		dateParsed, err := time.Parse("02/01/2006", recibo.contratanteDataNascimento)
		if err != nil {
			fmt.Println(err)
		} else {
			birthDate = &dateParsed
		}
	}

	if recibo.codigoPessoaContratante != "" {
		idConvert, _ := strconv.Atoi(recibo.codigoPessoaContratante)
		idCVC = &idConvert
	}

	return models.Person{
		IdCVC:      idCVC,
		Name:       recibo.contratanteNome,
		Type:       recibo.tipoPessoaContratante,
		Gender:     recibo.contratanteSexo,
		Number:     recibo.contratanteNumeroEndereco,
		Complement: recibo.contratanteDescricaoComplemento,
		CityName:   recibo.contratanteNomeCidade,
		Zip:        recibo.contratanteCEP,
		Email:      recibo.contratanteEmail,
		Rg:         rg,
		HomePhone:  homePhone,
		Cpf:        cpf,
		Address:    address,
		BirthDate:  birthDate,
	}
}

func buildSeller(vendedor arquivoVendedor) models.Person {
	var homePhone = fmt.Sprintf("(%s) %s", vendedor.numeroDDD, vendedor.numeroTelefone)
	var address = fmt.Sprintf("%s %s", vendedor.tipoEndereco, vendedor.nomeEndereco)
	var cpf *string
	var birthDate *time.Time
	var rg *string
	var idCVC *int

	if vendedor.tipoPessoa == "F" {
		if vendedor.codigoCPFCNPJ != "" && vendedor.codigoCPFCNPJDigito != "" {
			cpfConcat := fmt.Sprintf("%s%s", vendedor.codigoCPFCNPJ, vendedor.codigoCPFCNPJDigito)
			if len(cpfConcat) == 10 {
				cpfConcat = fmt.Sprintf("%s%s", "0", cpfConcat)
				cpf = &cpfConcat
			}
		}

		if vendedor.rg != "" {
			rg = &vendedor.rg
		}
	}

	if vendedor.dataNascimento != "" {
		dateParsed, err := time.Parse("02/01/2006", vendedor.dataNascimento)
		if err != nil {
			fmt.Println(err)
		} else {
			birthDate = &dateParsed
		}
	}

	if vendedor.codigoPessoa != "" {
		idConvert, _ := strconv.Atoi(vendedor.codigoPessoa)
		idCVC = &idConvert
	}

	return models.Person{
		IdCVC:      idCVC,
		Name:       vendedor.nomePessoa,
		Type:       vendedor.tipoPessoa,
		Number:     vendedor.numeroEndereco,
		Complement: vendedor.descricaoComplemento,
		CityName:   vendedor.nomeCidade,
		Zip:        vendedor.codigoCEP,
		Email:      vendedor.email,
		Rg:         rg,
		HomePhone:  homePhone,
		Cpf:        cpf,
		Address:    address,
		BirthDate:  birthDate,
	}
}

func buildPassenger(passageiro arquivoPAX) models.Person {
	var birthDate *time.Time
	var rg *string

	if passageiro.dataNascimento != "" {
		dateParsed, err := time.Parse("02/01/2006", passageiro.dataNascimento)
		if err != nil {
			fmt.Println(err)
		} else {
			birthDate = &dateParsed
		}
	}

	if passageiro.codigoRG != "" {
		rg = &passageiro.codigoRG
	}

	return models.Person{
		Name:         passageiro.nomePassageiro,
		Type:         "F",
		Gender:       passageiro.passageiroSexo,
		Rg:           rg,
		BirthDate:    birthDate,
		Observations: passageiro.observacao,
	}
}

func buildSale(receipt arquivoRecibo) models.Sale {
	products, errs := buildSaleProduct(receipt)
	saleDate, err := time.Parse("02/01/2006", receipt.dataLancamento)
	if err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return models.Sale{}
	}

	return models.Sale{
		SaleDate: saleDate,
		Buyer:    buildBuyer(receipt),
		Seller:   containsSeller(FILES_SELLERS, receipt.codigoPessoaVendedor),
		Products: products,
	}
}

func buildSaleProduct(fileReceipt arquivoRecibo) ([]models.SaleProduct, []error) {
	var err error
	var errs []error
	var dateCancellation *time.Time
	var dateShipment time.Time
	var returnDate time.Time

	if fileReceipt.dataCancelamento != "" {
		dateParsed, err := time.Parse("02/01/2006", fileReceipt.dataCancelamento)
		if err != nil {
			errs = append(errs, err)
		} else {
			dateCancellation = &dateParsed
		}
	}
	if fileReceipt.dataEmbarque != "" {
		dateShipment, err = time.Parse("02/01/2006", fileReceipt.dataEmbarque)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if fileReceipt.dataRetorno != "" {
		returnDate, err = time.Parse("02/01/2006", fileReceipt.dataRetorno)
		if err != nil {
			errs = append(errs, err)
		}
	}

	supplierDiscountValue, err := strconv.ParseFloat(fileReceipt.valorAbatimento, 32)
	if err != nil {
		errs = append(errs, err)
	}

	taxValue, err := strconv.ParseFloat(fileReceipt.valorTaxa, 32)
	if err != nil {
		errs = append(errs, err)
	}

	agencyDiscountValue, err := strconv.ParseFloat(fileReceipt.valorDesconto, 32)
	if err != nil {
		errs = append(errs, err)
	}

	commissionValue, err := strconv.ParseFloat(fileReceipt.valorComissaoCalculada, 32)
	if err != nil {
		errs = append(errs, err)
	}

	productValue, err := strconv.ParseFloat(fileReceipt.valorLancamento, 32)
	if err != nil {
		errs = append(errs, err)
	}

	valueIntermediaryCommission, err := strconv.ParseFloat(fileReceipt.valorComissaoIntermediario, 32)
	if err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return nil, errs
	} else {
		var products []models.SaleProduct
		products = append(products, models.SaleProduct{
			DateCancellation:            dateCancellation,
			DateShipment:                dateShipment,
			ReturnDate:                  returnDate,
			Document:                    fileReceipt.numeroRecibo,
			DescriptionProductType:      fileReceipt.descricaoTipoProduto,
			ProductCode:                 fileReceipt.codigoProduto,
			ProductDescription:          fileReceipt.descricaoProduto,
			ScriptDescription:           fileReceipt.descricaoRoteiro,
			SupplierDiscountValue:       float32(supplierDiscountValue),
			TaxValue:                    float32(taxValue),
			AgencyDiscountValue:         float32(agencyDiscountValue),
			CommissionValue:             float32(commissionValue),
			ProductValue:                float32(productValue),
			ValueIntermediaryCommission: float32(valueIntermediaryCommission),
		})
		return products, nil
	}
}

func buildProduct(fileReceipt arquivoRecibo) models.Product {
	idCVC, _ := strconv.Atoi(fileReceipt.codigoProduto)
	return models.Product{
		IdCVC:       idCVC,
		Description: fileReceipt.descricaoProduto,
	}
}

func containsSeller(fileSellers []arquivoVendedor, sellerCode string) models.Person {
	for _, seller := range fileSellers {
		if seller.codigoPessoa == sellerCode {
			return buildSeller(seller)
		}
	}
	return models.Person{}
}

func isDefaultExtractFile(columns []string) bool {
	return len(columns) == 53
}

func isDefaultSellerFile(columns []string) bool {
	return len(columns) == 28
}

func isDefaultPAXFile(columns []string) bool {
	return len(columns) == 12
}

func isFirstLine(index int) bool {
	if index == 0 {
		return true
	}
	return false
}

func enableLog(option bool) {
	if !option {
		log.SetOutput(ioutil.Discard)
	}
}
