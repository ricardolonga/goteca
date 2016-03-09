package main

import (
	"testing"
//	"os"
	"github.com/NeowayLabs/logger"
	"gopkg.in/mgo.v2"
	"github.com/satori/go.uuid"
	"os"
)

var (
	DB_TEST = "goteca_test_" + uuid.NewV1().String()
)

const (
//	MONGO_CONTAINER_NAME = "mongotest"
	MONGO_URL = "localhost:27017"
)

func TestMain(m *testing.M) {
//	client, err := conjure.NewClient()
//	if err != nil {
//		logger.Fatal("Error on getting conjure client.", err)
//	}

//	os.Setenv("MONGO_URL", MONGO_URL)
os.Setenv("GIN_MODE", "release")

//	session := getSession()

//	teardown(/*client*/session)
	//	setup(client)

//	prepareOutputs(session)
//	prepareSections(session)

	logger.Info("Running all tests...")

    m.Run()

	logger.Info("Finished all tests.")

//    teardown(/*client*/session)
}

//func setup(client *conjure.Client) {
//	// http://docs.docker.com/engine/reference/api/docker_remote_api_v1.20/#create-a-container
//
//	mongoDockerSpec := `{
//		"Name": "mongotest",
//		"Config": {
//			"Image": "tutum/mongodb",
//			"Env": ["AUTH=no"]
//		},
//		"HostConfig": {
//			"PortBindings": {
//				"27017/tcp": [{
//					"HostPort": "27000"
//				}],
//				"28017/tcp": [{
//					"HostPort": "28000"
//				}]
//			}
//		}
//	}`
//
//	container, err := client.Run(mongoDockerSpec)
//	if err != nil {
//		logger.Fatal("Error on starting docker container. %s", err)
//	}
//
//	if err = netutil.AwaitReachable(MONGO_URL, 15 * time.Second); err != nil {
//		if err := client.Remove(container.Name); err != nil {
//			logger.Fatal("Mongo not reachable (timeout: 15s), error on stopping Mongo container. %s", err)
//		}
//
//		logger.Fatal("Mongo not reachable (timeout: 15s): %s", err)
//	}
//}

//func prepareOutputs(session *mgo.Session) {
//	logger.Info("Prepare %s collection on %s database.", OUTPUTS_TEST, DB_TEST)
//
//	if err := loadObject(session, DB_TEST, OUTPUTS_TEST, "datasets/config_output_1_empresas.json"); err != nil {
//		logger.Fatal("Error on loading a object into Mongo: %s", err)
//	}
//	//	if err := loadObject(session, DB, OUTPUTS, "datasets/config_output_3_empresas.json"); err != nil {
//	//		logger.Fatal("Error on loading a object into Mongo: %s", err)
//	//	}
//	if err := loadObject(session, DB_TEST, OUTPUTS_TEST, "datasets/config_output_5_criactive_empresas.json"); err != nil {
//		logger.Fatal("Error on loading a object into Mongo: %s", err)
//	}
//}
//
//func prepareSections(session *mgo.Session) {
//	logger.Info("Prepare %s collection on %s database.", SECTIONS_TEST, DB_TEST)
//
//	if err := loadArray(session, DB_TEST, SECTIONS_TEST, "datasets/config_output_1_empresas_sections.json"); err != nil {
//		logger.Fatal("Error on loading objects array into Mongo: %s", err)
//	}
//	//	if err := loadArray(session, DB, SECTIONS, "datasets/config_output_3_empresas_sections.json"); err != nil {
//	//		logger.Fatal("Error on loading objects array into Mongo: %s", err)
//	//	}
//	if err := loadArray(session, DB_TEST, SECTIONS_TEST, "datasets/config_output_5_criactive_empresas_sections.json"); err != nil {
//		logger.Fatal("Error on loading objects array into Mongo: %s", err)
//	}
//}

func teardown(/*client *conjure.Client*/session *mgo.Session) {
//	client.Remove(MONGO_CONTAINER_NAME)
	logger.Info("Drop Mongo's %s database.", DB_TEST)

	session.Refresh()
	session.DB(DB_TEST).DropDatabase()
}

