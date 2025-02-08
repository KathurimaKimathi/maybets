package gorm_test

import (
	"fmt"
	"log"
	"os"
	"testing"
	"text/template"

	"github.com/KathurimaKimathi/maybets/pkg/maybets/application/helpers"
	"github.com/KathurimaKimathi/maybets/pkg/maybets/infrastructure/database/postgres"
	"github.com/KathurimaKimathi/maybets/pkg/maybets/infrastructure/database/postgres/gorm"
	"github.com/go-testfixtures/testfixtures/v3"
)

var (
	fixtures  *testfixtures.Loader
	testingDB *gorm.DBInstance

	// user variables
	userID  = "6ecbbc80-24c8-421a-9f1a-e14e12678ee0"
	userID2 = "6fd1bc80-24c8-421b-9f1a-e14e12698ef0"
	userID3 = "8ecbbc80-24c8-421a-9f1a-e14e12678ef1"
	userID4 = "4181df12-ca96-4f28-b78b-8e8ad88b25df"
	userID5 = "6ecccc80-24c8-421a-9f1a-e14e13678ef0"
	userID6 = "5ecbbc80-24c8-421a-9f1a-e14e12678ee0"

	// bets
	bet1UserID  = "4ecbbc80-24c8-421a-9f1a-e14e12678ee0"
	bet2UserID  = "5ecbbc80-24b8-421a-9f1a-e14e12678ee0"
	bet3UserID  = "64063dc0-028f-4fbd-85d2-657e453cc40c"
	bet4UserID  = "8ecbbc80-24c8-421a-9f1a-e14e12678ef0"
	bet1UserID2 = "e1e90ea3-fc06-442e-a1ec-251a031c0ca7"
	bet2UserID2 = "723b64b3-e4d6-4416-98b2-18798279e457"
	bet3UserID2 = "839f9a85-bbe6-48e7-a730-42d56a39b532"
	bet4UserID2 = "f186100a-2b6c-4656-9bbd-960492f6bfb4"
	bet5UserID2 = "4aa35fa8-a720-4c6f-9510-86fe4b4addbd"
	bet6UserID2 = "5ecbbc80-24c8-421a-9f1a-e14e12678ef3"
	bet7UserID2 = "5ecbbc80-24c8-421a-9f1a-e14e12678ee4"
	bet1UserID3 = "650b7958-12fd-4fa6-9309-ec11618263ae"
	bet2UserID3 = "d15c3bb1-bc52-44cc-875e-bf7f4d921dee"
	bet1UserID4 = "5ecbbc80-24c8-421a-9f1a-e14e12678ee1"
	bet2UserID4 = "5ecbbc80-24c8-421a-9f1a-e14e12678ef1"
	bet1UserID5 = "5ecbbc80-24c8-421a-9f1a-e14e12678ee2"
	bet2UserID5 = "f933fd4b-1e3c-4ecd-9d7a-82b2790c0543"
	bet2UserID6 = "5ecbbc80-24c8-421a-9f1a-e14e12678ef4"
)

func TestMain(m *testing.M) {
	isLocalDB := helpers.CheckIfCurrentDBIsLocal()
	if !isLocalDB {
		fmt.Println("Cannot run tests. The current database is not a local database.")
		os.Exit(1)
	}

	log.Println("setting up test database")

	var err error

	testingDB, err = gorm.NewDBInstance()
	if err != nil {
		fmt.Println("failed to initialize db:", err)
		os.Exit(1)
	}

	db, err := testingDB.DB.DB()
	if err != nil {
		fmt.Println("failed to initialize db:", err)
		os.Exit(1)
	}

	err = postgres.RunMigrations()
	if err != nil {
		fmt.Println("failed to run migrations:", err)
		os.Exit(1)
	}

	fixtures, err = testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("sqlite"),
		testfixtures.Template(),
		testfixtures.TemplateData(template.FuncMap{
			"test_user_id":      userID,
			"test_user_id2":     userID2,
			"test_user_id3":     userID3,
			"test_user_id4":     userID4,
			"test_user_id5":     userID5,
			"test_user_id6":     userID6,
			"test_bet1user1_id": bet1UserID,
			"test_bet2user1_id": bet2UserID,
			"test_bet3user1_id": bet3UserID,
			"test_bet4user1_id": bet4UserID,
			"test_bet1user2_id": bet1UserID2,
			"test_bet2user2_id": bet2UserID2,
			"test_bet3user2_id": bet3UserID2,
			"test_bet4user2_id": bet4UserID2,
			"test_bet5user2_id": bet5UserID2,
			"test_bet6user2_id": bet6UserID2,
			"test_bet7user2_id": bet7UserID2,
			"test_bet1user3_id": bet1UserID3,
			"test_bet2user3_id": bet2UserID3,
			"test_bet1user4_id": bet1UserID4,
			"test_bet2user4_id": bet2UserID4,
			"test_bet1user5_id": bet1UserID5,
			"test_bet2user5_id": bet2UserID5,
			"test_bet1user6_id": bet2UserID6,
		}),
		testfixtures.Paths(
			"../../../../../../fixtures/bets.yml",
		),
		testfixtures.DangerousSkipTestDatabaseCheck(),
	)
	if err != nil {
		fmt.Println("failed to create fixtures:", err)
		os.Exit(1)
	}

	err = prepareTestDatabase()
	if err != nil {
		fmt.Println("failed to prepare test database:", err)
		os.Exit(1)
	}

	log.Printf("Running tests ...")
	os.Exit(m.Run())
}

func prepareTestDatabase() error {
	if err := fixtures.Load(); err != nil {
		return err
	}

	return nil
}
