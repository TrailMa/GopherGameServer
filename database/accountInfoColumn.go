package database

import(
	"errors"
	"strconv"
	"fmt"
)

// An AccountInfoColumn is the representation of an extra column on the users table that you can define. You can define as many
// as you want. The client APIs take in an optional parameter when signing up a client that you use
// to set your AccountInfoColumn(s). When a client logs in, the server callback will provide you with a
// map[string]interface{} of your AccountInfoColumn(s) where the key is the name of the column and the
// interface is the data retrieved from the column.
type AccountInfoColumn struct {
	dataType int
	maxSize int
}

var(
	customAccountInfo map[string]AccountInfoColumn = make(map[string]AccountInfoColumn)
)

// MySQL database data types. Use one of these when making a new AccountInfoColumn or
// CustomTable's columns. The parentheses next to a type indicate it requires a maximum
// size when making a column of that type.
const (
	//NUMERIC TYPES
	DataTypeTinyInt = iota // TINYINT()
	DataTypeSmallInt // SMALLINT()
	DataTypeReal // REAL
	DataTypeMediumInt // MEDIUMINT()
	DataTypeInt // INT()
	DataTypeFloat // FLOAT
	DataTypeDouble // DOUBLE
	DataTypeDecimal // DECIMAL
	DataTypeBigInt // BIGINT()

	//CHARACTER TYPES
	DataTypeChar // CHAR()
	DataTypeVarChar // VARCHAR()
	DataTypeNationalVarChar // NVARCHAR()
	DataTypeJSON // JSON

	//TEXT TYPES
	DataTypeTinyText // TINYTEXT
	DataTypeMediumText // MEDIUMTEXT
	DataTypeText // TEXT()
	DataTypeLongText // LONGTEXT

	//DATE TYPES
	DataTypeDate // DATE
	DataTypeDateTime // DATETIME()
	DataTypeTime // TIME()
	DataTypeTimeStamp // TIMESTAMP()
	DataTypeYear // YEAR()

	//BINARY TYPES
	DataTypeTinyBlob // TINYBLOB
	DataTypeMediumBlob // MEDIUMBLOB
	DataTypeBlob // BLOB()
	DataTypeLongBlob // LONGBLOB
	DataTypeBinary // BINARY()
	DataTypeVarBinary // VARBINARY()

	//OTHER TYPES
	DataTypeBit // BIT()
	DataTypeENUM // ENUM()
	DataTypeSet // SET()
)

var (
	//DATA TYPES THAT REQUIRE A SIZE
	dataTypesSize []string = []string{
					"TINYINT",
					"SMALLINT",
					"MEDIUMINT",
					"INT",
					"BIGINT",
					"CHAR",
					"VARCHAR",
					"NVARCHAR",
					"TEXT",
					"DATETIME",
					"TIME",
					"TIMESTAMP",
					"YEAR",
					"BLOB",
					"BINARY",
					"VARBINARY",
					"BIT",
					"ENUM",
					"SET"}

	//DATA TYPES THAT REQUIRE QUOTES
	dataTypesQuotes []string = []string{
					"CHAR",
					"VARCHAR",
					"NVARCHAR",
					"JSON",
					"TINYTEXT",
					"MEDIUMTEXT",
					"TEXT",
					"LONGTEXT",
					"DATE",
					"DATETIME",
					"TIME",
					"TIMESTAMP",
					"YEAR"}

	//DATA TYPE LITERAL NAME LIST
	dataTypes []string = []string{
					"TINYINT",
					"SMALLINT",
					"REAL",
					"MEDIUMINT",
					"INT",
					"FLOAT",
					"DOUBLE",
					"DECIMAL",
					"BIG INT",
					"CHAR",
					"VARCHAR",
					"NVARCHAR",
					"JSON",
					"TINYTEXT",
					"MEDIUMTEXT",
					"TEXT",
					"LONGTEXT",
					"DATE",
					"DATETIME",
					"TIME",
					"TIMESTAMP",
					"YEAR",
					"TINYBLOB",
					"MEDIUMBLOB",
					"BLOB",
					"LONGBLOB",
					"BINARY",
					"VARBINARY",
					"BIT",
					"ENUM",
					"SET"}
)

// Use this to make a new AccountInfoColumn. You can only make new AccountInfoColumns before starting the server.
func NewAccountInfoColumn(name string, dataType int, maxSize int) error {
	if(serverStarted){
		return errors.New("You can't make a new AccountInfoColumn after the server has started");
	}else if(len(name) == 0){
		return errors.New("database.NewAccountInfoColumn() requires a name");
	}else if(dataType < 0 || dataType > len(dataTypes)-1){
		return errors.New("Incorrect data type");
	}

	if(isSizeDataType(dataType) && maxSize == 0){
		return errors.New("The data type '"+dataTypesSize[dataType]+"' requires a max size");
	}

	customAccountInfo[name] = AccountInfoColumn{dataType: dataType, maxSize: maxSize};

	//
	return nil;
}

//CHECKS IF THE DATA TYPE REQUIRES A MAX SIZE
func isSizeDataType(dataType int) bool {
	for i := 0; i < len(dataTypesSize); i++ {
		if(dataTypes[dataType] == dataTypesSize[i]){
			return true;
		}
	}
	return false;
}

//CHECKS IF THE DATA TYPE REQUIRES QUOTES ON INSERT
func dataTypeNeedsQuotes(dataType int) bool {
	for i := 0; i < len(dataTypesQuotes); i++ {
		if(dataTypes[dataType] == dataTypesQuotes[i]){
			return true;
		}
	}
	return false;
}

//CONVERTS DATA TYPE TO STRING
func convertDataToString(dataType string, data interface{}) (string, error) {
	switch data.(type) {
		case int:
			if(dataType != "INT" && dataType != "TINYINT" && dataType != "MEDIUMINT" && dataType != "BIGINT" && dataType != "SMALLINT"){
				return "", errors.New("Mismatched data types");
			}
			return strconv.Itoa(data.(int)), nil;

		case float32:
			if(dataType != "REAL" && dataType != "FLOAT" && dataType != "DOUBLE" && dataType != "DECIMAL"){
				return "", errors.New("Mismatched data types");
			}
			return fmt.Sprintf("%f", data.(float32)), nil;

		case float64:
			if(dataType != "REAL" && dataType != "FLOAT" && dataType != "DOUBLE" && dataType != "DECIMAL"){
				return "", errors.New("Mismatched data types");
			}
			return strconv.FormatFloat(data.(float64), 'f', -1, 64), nil;

		case string:
			if(dataType != "CHAR" && dataType != "VARCHAR" && dataType != "NVARCHAR" && dataType != "JSON" && dataType != "TEXT" &&
						dataType != "TINYTEXT" && dataType != "MEDIUMTEXT" && dataType != "LONGTEXT" && dataType != "DATE" &&
						dataType != "DATETIME" && dataType != "TIME" && dataType != "TIMESTAMP" && dataType != "YEAR"){
				return "", errors.New("Mismatched data types");
			}
			return data.(string), nil;

		default:
			return "", errors.New("Data type provided isn't supported yet. You can open a ticket at Gopher Game Server's project on GitHub to request SQL support for a data type.");
	}
}
