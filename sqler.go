package SqlEr

import (
	"fmt"

	JsonX "github.com/pschlump/JSONx"

	"github.com/pschlump/Go-FTL/server/sizlib"
	"github.com/pschlump/MiscLib"
	"github.com/pschlump/godebug"
)

type TableColumnsType struct {
	ColumnName      string
	ColumnType      string
	ColumnLen       int
	ColumnPrecision int
	ColumnDefault   string
	ColumnOpts      string // not null, pimary key etc.
}

type TableType struct {
	TableName        string
	TableColumnsName []TableColumnsType
	DDL              []string
	AlterDDL         []string
	Data             []string
}

type SqlType struct {
	Tables []TableType
	// indexes
	// functions/procedures
	// other items?
}

type SqlErType struct {
	FileName string
	Sql      SqlType
}

// xyzzy - find the code that checks for tables existence and validates them
// xyzzy - read in global config - .jx file to connect to d.b.
// xyzzy - connect to d.b.
// xyzzy - disconnect from d.b.
// xyzzy - get meta data from d.b.
// xyzzy - describe tables - in list - into .jx file
// xyzzy - CLI to do this
// xyzzy - transform DDL from old to new format
// xyzzy - transform *Data* from old to new format - with redidual
// xyzzy - extract data in .sql for push to client
// xyzzy - extract/encrypt data
// xyzzy - sanitize data - clean for push to test
// xyzzy - new/old .jx files decribing model

func ValidateInstallModel(fn string) (rv *SqlErType) {
	rv = &SqlErType{
		FileName: fn,
	}
	meta, err := JsonX.UnmarshalFile(fn, &rv.Sql)
	if err != nil {
		fmt.Printf("Uname to parse JsonX file %s, %s\n", fn, err)
		return
	} else {
		fmt.Printf("Raw data: %s\n", JsonX.SVarI(rv.Sql))
	}

	if ValidateData {

		err = JsonX.ValidateValues(&rv.Sql, meta, "", "", "")
		if err != nil {
			// xyzzy
		}

		// Test ValidateRequired - both success and fail, nested, in arrays
		errReq := JsonX.ValidateRequired(&rv.Sql, meta)
		if errReq == nil {
			// xyzzy
		}

		fmt.Printf("err=%s meta=%s\n", err, JsonX.SVarI(meta))
		msg, err := JsonX.ErrorSummary("text", &rv.Sql, meta)
		if err != nil {
			fmt.Printf("%s\n", msg)
		}

	}
	return
}

/*
func (hdlr *TabServer2Type) GetTableInformationSchema(conn *sizlib.MyDb, TableName string) (rv DbTableType, err error) {
	// qry := `SELECT * FROM information_schema.tables WHERE table_schema = $1 ORDER BY table_name`
	fmt.Printf("Validateing [%s]\n", TableName)
	qry := `SELECT * FROM information_schema.tables WHERE table_schema = $1 and table_name = $2`
	data := sizlib.SelData(conn.Db, qry, g_schema, TableName)
	if data == nil || len(data) == 0 {
		fmt.Fprintf(os.Stderr, "%sMessage(90532): Missing table%s%s\n", MiscLib.ColorRed, TableName, MiscLib.ColorReset)
		fmt.Printf("Message(90532): Missing table%s\n", TableName)
		err = errors.New("Missing Table")
		return
	}
	fmt.Fprintf(os.Stderr, "%sFound table: %s%s\n", MiscLib.ColorGreen, TableName, MiscLib.ColorReset)
	fmt.Printf("Table [%s] found in schema [%s]\n", data[0]["table_name"], data[0]["table_schema"])
	rv.TableName = TableName
*/

func (se *SqlErType) RunASQLCommand(conn *sizlib.MyDb, cmd string, cmdNo int) (err error) {
	//data := sizlib.SelData(conn.Db, cmd)
	//_ = data
	err = sizlib.Run1(conn.Db, cmd)
	if err != nil {
		fmt.Printf("%sAt: %s - Error on Command %s:%d:%s%s\n", MiscLib.ColorRed, godebug.LF(), err, cmdNo, JsonX.SVarI(cmd), MiscLib.ColorReset)
	}
	return
}

func (se *SqlErType) RunSliceOfSQLCommands(conn *sizlib.MyDb, cmd []string) (err error) {

	fmt.Printf("%sAt: %s - Running:%s%s\n", MiscLib.ColorCyan, godebug.LF(), JsonX.SVarI(cmd), MiscLib.ColorReset)
	n_err := 0
	for ii, vv := range cmd {
		err = se.RunASQLCommand(conn, vv, ii)
		if err != nil {
			n_err++
		}
	}

	fmt.Printf("n_err=%d\n", n_err)

	return

}

func (se *SqlErType) CheckTables() {
	for ii, vv := range se.Sql.Tables {
		_, _ = ii, vv
	}
}

func IfTableExists(tn string) bool {
	return true
}

func IfTableColumnsExists(tab *TableType) bool {
	return true
}

func CreateTable(tab *TableType) bool {
	return true
}

func AlterTable(tab *TableType) bool {
	return true
}

const ValidateData = false
