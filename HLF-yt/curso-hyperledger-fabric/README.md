# Curso de Hyperledger Fabric LatinoAmerica 2020

Bienvenid@ a nuestro curso de desarrollador blockchain en Hyperledger Fabric ofrecido por [Hyperledger Latinoamérica](https://wiki.hyperledger.org/display/CP/Hyperledger+Latinoamerica "Hyperledger Latinoamérica") y [Business Blockchain](https://www.blockchainempresarial.com/ "Business Blockchain").

Con el fin de gestionar correctamente su experiencia en el curso le pedimos lea detenidamente las siguientes instrucciones:

Primero es indispensable registrarse en nuestro canal de rocket chat por donde estaremos interactuando, para ello use el siguiente enlace (debes tener un ID de Linux Foundation), si ya tienes uno solo te debes autenticarte.
 
https://chat.hyperledger.org/channel/community-latinoamerica

El curso dura aproximadamente 30 horas, distribuidas en workshops prácticos y clases teóricas en vivo on line via zoom, repartidas durante 2 meses. En las diferentes sesiones podrán interactuar con instructores expertos que dictarán el curso en vivo.

Todas las clases serán vía  Zoom en la misma sala de zoom

https://us02web.zoom.us/j/83944607895?pwd=RCtoZjZhaDhhK1hhUVROTWM2bUFqUT09

Las instrucciones para instalar el cliente de zoom estan aquí: https://zoom.us/download#client_4meeting 
El curso inicia el día Jueves 20  de Agosto del 2020  y pueden ver las fechas definidas en el siguiente:
 https://wiki.hyperledger.org/display/CP/Curso++Hyperledger+Fabric
 
Los horarios de las sesiones son los siguientes:

Horario Jueves 
- Perú, México, Colombia, Ecuador 18:00 pm
- Chile, Bolivia 19:00 pm
- Argentina 20:00 horas

Horario sábados 
- Perú, México, Colombia, Ecuador 10:00 am
- Chile, Bolivia 11:00 am
- Argentina 12:00  horas
 
El horario de las clases están orientadas a no interrumpir  actividades laborales o personales de cada participante.

Cada participante del curso debe instalar su propio ambiente de desarrollo local con las instrucciones que se entregarán y revisarán en la primera clase del Jueves 20 de agosto de 2020.

La capacitación tiene una orientación principalmente práctica. Los asistentes trabajan junto al entrenador via online, desarrollando aplicaciones de blockchain en Hyperledger Fabric. 


¡Abrazo y bienvenido!

chmod 777 /home/fabio/HLF_pruebas/ -R
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

> cryptogen generate --config=./crypto-config.yaml
    VERIFICAR CERTIFICADO
    https://www.dondominio.com/products/ssl/tools/ssl-checker/

Primero crear carpeta
> mkdir channel-artifacts
Dentro de la carpeta con el archivo "configtx.yaml", crea el bloque genesis, "ThreeOrgsOrdererGenesis" LN 218 de yaml
> configtxgen -profile ThreeOrgsOrdererGenesis -channelID system-channel -outputBlock ./channel-artifacts/genesis.block
> export CHANNEL_NAME=acmechannel
Crear trasaccion de canal, "acmechannel" NOMBRE DEL CANAL, "ThreeOrgsChannel" LN 232 de yaml
> configtxgen -profile ThreeOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID $CHANNEL_NAME
Crear archivos de anchor peers, "Org1MSP" LN 30 de yaml, HACER ESTO PARA TODOS LOS anchors peers
> configtxgen -profile ThreeOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org1MSP
> configtxgen -profile ThreeOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org2MSP
> configtxgen -profile ThreeOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org3MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org3MSP

Levantar portainer
> docker volume create portainer_data
> docker run -d -p 8000:8000 -p 9000:9000 -v /var/run/docker.sock:/var/run/docker.sock -v portainer_data:/data portainer/portainer
    admin
    Hlf2022.*+

Definiendo variables
> export CHANNEL_NAME=acmechannel
> export VERBOSE=false
> export FABRIC_CFG_PATH=$PWD

Levantar ChoudDB
> CHANNEL_NAME=$CHANNEL_NAME docker-compose -f docker-compose-cli-couchdb.yaml up -d

COnectarse al contenedor "cli" desde portainer y ejecutar los comandos:
> pwd
> export CHANNEL_NAME=marketplace
> peer channel create -o orderer.ceiba.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/ceiba.com/orderers/orderer.ceiba.com/msp/tlscacerts/tlsca.ceiba.com-cert.pem
> cd channel-artifacts/

    PARA CONECTARSE A COUCHDB "5984": Uno de los puertos asignados al couchdb
        http://172.18.159.17:5984/_utils/#login

Conectado las organizaciones al canal, cuando se une una organizacion al canal se crea una BD en su respectivo couchdb
    Se une la organizacion main, por defecto se una la identidad de la primera organizacion
> peer channel join -b ceibachannel.block
    Se une al canal otra organizacion que no es la main 
> CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.ceiba.com/users/Admin@org2.ceiba.com/msp CORE_PEER_ADDRESS=peer0.org2.ceiba.com:7051 CORE_PEER_LOCALMSPID="Org2MSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.ceiba.com/peers/peer0.org2.ceiba.com/tls/ca.crt peer channel join -b ceibachannel.block

> CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.ceiba.com/users/Admin@org3.ceiba.com/msp CORE_PEER_ADDRESS=peer0.org3.ceiba.com:7051 CORE_PEER_LOCALMSPID="Org3MSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.ceiba.com/peers/peer0.org3.ceiba.com/tls/ca.crt peer channel join -b ceibachannel.block

Setear el anchorpeer para cada organizacion
    Anchor peer main
> peer channel update -o orderer.ceiba.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/Org1MSPanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/ceiba.com/orderers/orderer.ceiba.com/msp/tlscacerts/tlsca.ceiba.com-cert.pem

    Setear anchor peer de una organizacion no main
> CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.ceiba.com/users/Admin@org2.ceiba.com/msp CORE_PEER_ADDRESS=peer0.org2.ceiba.com:7051 CORE_PEER_LOCALMSPID="Org2MSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.ceiba.com/peers/peer0.org2.ceiba.com/tls/ca.crt peer channel update -o orderer.ceiba.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/Org2MSPanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/ceiba.com/orderers/orderer.ceiba.com/msp/tlscacerts/tlsca.ceiba.com-cert.pem

> CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.ceiba.com/users/Admin@org3.ceiba.com/msp CORE_PEER_ADDRESS=peer0.org3.ceiba.com:7051 CORE_PEER_LOCALMSPID="Org3MSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.ceiba.com/peers/peer0.org3.ceiba.com/tls/ca.crt peer channel update -o orderer.ceiba.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/Org3MSPanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/ceiba.com/orderers/orderer.ceiba.com/msp/tlscacerts/tlsca.ceiba.com-cert.pem


***********RED BLOCKCAIN LISTA PARA CREAR CONTRATOS INTELIGENTES (CODECHAIN)***********

Aca se empieza a desarrollar el chaincode en golang
Una vez se termine de escribir el chaincode, dentro del contenedor cli dirigirse a "/opt/gopath/src/github.com/chaincode", esta carpeta esta binding a la carpeta "./../chaincode" VER LINEA 133 DE "docker-compose-cli-couchdb.yaml"

Se definen variables
> export CHANNEL_NAME=marketplace
    Nombre del chaincode a desplegar
> export CHAINCODE_NAME=foodcontrol
> export CHAINCODE_VERSION=1
> export CC_RUNTIME_LANGUAGE=golang
> export CC_SRC_PATH="../../../chaincode/$CHAINCODE_NAME/"
    Ruta del CA
> export ORDERER_CA="/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/ceiba.com/orderers/orderer.ceiba.com/msp/tlscacerts/tlsca.ceiba.com-cert.pem"

Ejecutar los comando justo en path inicial del contenedor cli (justo cuando se conecta), se compilan los archivos en un archivo tar.gz, y luego se instala en cada organizacion (ELIMINAR go.sum)
> peer lifecycle chaincode package ${CHAINCODE_NAME}.tar.gz --path ${CC_SRC_PATH} --lang ${CC_RUNTIME_LANGUAGE} --label ${CHAINCODE_NAME}_${CHAINCODE_VERSION} >&log.txt

El empaquetado generado se envia a instalar a todos los peers
Se le manda a todas las organizaciones
> peer lifecycle chaincode install ${CHAINCODE_NAME}.tar.gz
    (siempre es el mismo) COPIAR EL "Chaincode code package identifier:"
        foodcontrol_1:3bb63d5460a2fd7bb991090accf37ea4ce3992a2ce79a17ccd7d21889146c399
        
Instalandolo en Organizaciones no mai, tiene que setar los valores
        > CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.ceiba.com/users/Admin@org2.ceiba.com/msp CORE_PEER_ADDRESS=peer0.org2.ceiba.com:7051 CORE_PEER_LOCALMSPID="Org2MSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.ceiba.com/peers/peer0.org2.ceiba.com/tls/ca.crt peer lifecycle chaincode install ${CHAINCODE_NAME}.tar.gz

> CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.ceiba.com/users/Admin@org3.ceiba.com/msp CORE_PEER_ADDRESS=peer0.org3.ceiba.com:7051 CORE_PEER_LOCALMSPID="Org3MSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.ceiba.com/peers/peer0.org3.ceiba.com/tls/ca.crt peer lifecycle chaincode install ${CHAINCODE_NAME}.tar.gz

Definiendo las politicas de endorzamiento (que organizaciones puede aprovar los chaincode)
> peer lifecycle chaincode approveformyorg --tls --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name $CHAINCODE_NAME --version $CHAINCODE_VERSION --sequence 1 --waitForEvent --signature-policy "OR ('Org1MSP.peer', 'Org3MSP.peer')" --package-id foodcontrol_1:3bb63d5460a2fd7bb991090accf37ea4ce3992a2ce79a17ccd7d21889146c399

    VALIDAR EL ESTADO DE APROBADOS PARA LA NUEVA POLITICA
> peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name $CHAINCODE_NAME --version $CHAINCODE_VERSION --sequence 1 --signature-policy "OR ('Org1MSP.peer', 'Org3MSP.peer')" --output json

Se aprueba por la organizacion 3
> CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.ceiba.com/users/Admin@org3.ceiba.com/msp CORE_PEER_ADDRESS=peer0.org3.ceiba.com:7051 CORE_PEER_LOCALMSPID="Org3MSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.ceiba.com/peers/peer0.org3.ceiba.com/tls/ca.crt peer lifecycle chaincode approveformyorg --tls --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name $CHAINCODE_NAME --version $CHAINCODE_VERSION --sequence 1 --waitForEvent --signature-policy "OR ('Org1MSP.peer', 'Org3MSP.peer')" --package-id foodcontrol_1:3bb63d5460a2fd7bb991090accf37ea4ce3992a2ce79a17ccd7d21889146c399

Se comitea los chaincode
> peer lifecycle chaincode commit -o orderer.ceiba.com:7050 --tls --cafile $ORDERER_CA --peerAddresses peer0.org1.ceiba.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.ceiba.com/peers/peer0.org1.ceiba.com/tls/ca.crt --peerAddresses peer0.org3.ceiba.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.ceiba.com/peers/peer0.org3.ceiba.com/tls/ca.crt --channelID $CHANNEL_NAME --name $CHAINCODE_NAME --version $CHAINCODE_VERSION --sequence 1 --signature-policy "OR ('Org1MSP.peer', 'Org3MSP.peer')"


Consultado los ENDPOINTS, "Set" funcion del chaincode de golang
    POST
> peer chaincode invoke -o orderer.ceiba.com:7050 --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n $CHAINCODE_NAME -c '{"Args":["Set", "did:3", "fabio", "ramos"]}'
> CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.ceiba.com/users/Admin@org3.ceiba.com/msp CORE_PEER_ADDRESS=peer0.org3.ceiba.com:7051 CORE_PEER_LOCALMSPID="Org3MSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.ceiba.com/peers/peer0.org3.ceiba.com/tls/ca.crt peer chaincode invoke -o orderer.ceiba.com:7050 --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n $CHAINCODE_NAME -c '{"Args":["Set", "did:3", "fabio ernesto", "ramos reyes"]}'
    GET
> peer chaincode invoke -o orderer.ceiba.com:7050 --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n $CHAINCODE_NAME -c '{"Args":["Query", "did:3"]}'