package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SmartContract provides functions for control the cliente
type SmartContract struct {
	contractapi.Contract
}

//Cliente describes basic details of what makes up a cliente
type Cliente struct {
	ClienteId     string `json:"clienteid"`
	NombreCliente string `json:"nombrecliente"`
	Puntos        string `json:"puntos"`
}

type Contrado struct {
	ContradoId     string `json:"contradoid"`
	TipoContrado   string `json:"tipocontrado"` //Producto o Servicio
	NombreContrato string `json:"nombrecontrato"`
	ValorEnPuntos  string `json:"valorenpuntos"`
}

type Adquirir struct {
	AdquirirId string `json:"adquiririd"`
	ClienteId  string `json:"clienteid"`
	ContradoId string `json:"contradoid"`
}

type TransferenciaEntreClientes struct {
	TransferenciaEntreClientesId string `json:"transferenciaentreclientesid"`
	ClienteId_origen             string `json:"clienteid_origen"`
	ClienteId_destino            string `json:"clienteid_destino"`
	MontoTransferido             string `json:"montotransferido"`
}

//Inicio: Se crea contratos
func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	producto := Contrado{ContradoId: "1", TipoContrado: "Producto", NombreContrato: "Computadora", ValorEnPuntos: "100"}

	productoAsJSONBytes, _ := json.Marshal(producto)
	err := stub.PutState(producto.ContradoId, productoAsJSONBytes)
	if err != nil {
		return shim.Error("Failed to producto asset " + producto.NombreContrato)
	}

	servicio := Contrado{ContradoId: "2", TipoContrado: "Producto", NombreContrato: "Computadora", ValorEnPuntos: "100"}

	servicioAsJSONBytes, _ := json.Marshal(servicio)
	errServicio := stub.PutState(servicio.ContradoId, servicioAsJSONBytes)
	if errServicio != nil {
		return shim.Error("Failed to servicio asset " + servicio.NombreContrato)
	}

	cliente := Cliente{ClienteId: "1", NombreCliente: "Fabio Ramos", Puntos: "100"}

	clienteAsJSONBytes, _ := json.Marshal(cliente)
	errCliente := stub.PutState(cliente.ClienteId, clienteAsJSONBytes)
	if errCliente != nil {
		return shim.Error("Failed to cliente asset " + cliente.NombreCliente)
	}

	return shim.Success([]byte("Assets created successfully."))
}

func (s *SmartContract) QueryContrato(ctx contractapi.TransactionContextInterface, contradoId string) (*Contrado, error) {

	contratoAsBytes, err := ctx.GetStub().GetState(contradoId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if contratoAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", contradoId)
	}

	contrato := new(Contrado)

	err = json.Unmarshal(contratoAsBytes, contrato)
	if err != nil {
		return nil, fmt.Errorf("Unmarshal error. %s", err.Error())
	}

	return contrato, nil
}

//Fin: Se crea contratos

//Inicio: chaincode de Clientes

func (s *SmartContract) SaldoActual(ctx contractapi.TransactionContextInterface, clienteId string) (*Cliente, error) {

	clienteAsBytes, err := ctx.GetStub().GetState(clienteId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if clienteAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", clienteId)
	}

	cliente := new(Cliente)

	err = json.Unmarshal(clienteAsBytes, cliente)
	if err != nil {
		return nil, fmt.Errorf("Unmarshal error. %s", err.Error())
	}

	return cliente, nil
}

func (s *SmartContract) HistorialTransacciones(ctx contractapi.TransactionContextInterface, clienteId string) (*Cliente, error) {

	clienteAsBytes, err := ctx.GetStub().GetState(clienteId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if clienteAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", clienteId)
	}

	cliente := new(Cliente)

	err = json.Unmarshal(clienteAsBytes, cliente)
	if err != nil {
		return nil, fmt.Errorf("Unmarshal error. %s", err.Error())
	}

	return cliente, nil
}

func (s *SmartContract) TransferirPuntosEntreClientes(ctx contractapi.TransactionContextInterface, clienteId_origen string, clienteId_destino string, montoTransferido string) error {

	// cliente_origenAsBytes, errOrigen := ctx.GetStub().GetState(clienteId_origen)

	// if errOrigen != nil {
	// 	return errOrigen
	// }

	cliente_destinoAsBytes, errClienteDestino := ctx.GetStub().GetState(clienteId_destino)
	if errClienteDestino != nil {
		return errClienteDestino
	}

	return ctx.GetStub().PutState(clienteId_destino, cliente_destinoAsBytes)
}

func (s *SmartContract) CrearClientes(ctx contractapi.TransactionContextInterface, clienteId string, nombrecliente string, puntos string) error {

	intVar, err := strconv.Atoi(puntos)
	if intVar <= 0 {
		return fmt.Errorf("No se puede crear un cliente con 0 punto. %s", err.Error())
	}

	cliente := Cliente{
		ClienteId:     uuid.New(),
		NombreCliente: nombrecliente,
		Puntos:        puntos,
	}

	clienteAsBytes, err := json.Marshal(cliente)
	if err != nil {
		fmt.Printf("Marshal error: %s", err.Error())
		return err
	}

	return ctx.GetStub().PutState(clienteId, clienteAsBytes)
}

//Fin: chaincode de Clientes

//Inicio: Acciones de organizaciones

func (s *SmartContract) EmitirPuntos(ctx contractapi.TransactionContextInterface, clienteId string, puntos string) error {

	intVar, err := strconv.Atoi(puntos)
	if intVar <= 0 {
		return fmt.Errorf("La cantidad de puntos a emitir tiene que mayor que 0. %s", err.Error())
	}

	cliente := Cliente{
		ClienteId: clienteId,
		Puntos:    puntos,
	}

	clienteAsBytes, err := json.Marshal(cliente)
	if err != nil {
		fmt.Printf("Marshal error: %s", err.Error())
		return err
	}

	return ctx.GetStub().PutState(clienteId, clienteAsBytes)
}

//Fin: Acciones de organizaciones

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create ceibacoins chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting ceibacoins chaincode: %s", err.Error())
	}
}
