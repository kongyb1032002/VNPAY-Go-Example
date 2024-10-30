package config

import (
	"log"
	"strconv"

	oracle "github.com/godoes/gorm-oracle"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func LoadOracleConfig(cfg *Config) (*gorm.DB, error) {
	options := map[string]string{
		"CONNECTION TIMEOUT": "90",
		"LANGUAGE":           "SIMPLIFIED CHINESE",
		"TERRITORY":          "CHINA",
		"SSL":                "false",
	}
	port, err := strconv.Atoi(cfg.OraclePort) // Sử dụng strconv.Atoi để chuyển đổi
	if err != nil {
		log.Fatalf("Error converting DbPort: %v", err)
	}
	url := oracle.BuildUrl(cfg.OracleHost, port, cfg.OracleDb, cfg.OracleUser, cfg.OraclePwd, options)

	dialector := oracle.New(oracle.Config{
		DSN:                     url,
		IgnoreCase:              false,
		NamingCaseSensitive:     true,
		VarcharSizeIsCharLength: true,

		RowNumberAliasForOracle11: "ROW_NUM",
	})
	db, err := gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase:         true,
			IdentifierMaxLength: 30, // Oracle: 30, PostgreSQL:63, MySQL: 64, SQL Server、SQLite、DM: 128
		},
		PrepareStmt:     false,
		CreateBatchSize: 50,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
