package libs

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/cclab-inu/Kunerva/src/types"
	"github.com/rs/zerolog/log"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	createDB()
	CreateTableNetworkPolicyMySQL()
	CreateTableNetworkLogsMySQL()
}

func createDB() {
	db, err := sql.Open(DBDriver, DBUser+":"+DBPass+"@tcp("+DBHost+":"+DBPort+")/")
	if err != nil {
		log.Error().Msg("connection error :" + err.Error())
		return
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + DBName)
	if err != nil {
		log.Error().Msg("database creation error :" + err.Error())
		return
	}

	db.Close()
}

// ConnectMySQL function
func ConnectMySQL() (db *sql.DB) {
	db, err := sql.Open(DBDriver, DBUser+":"+DBPass+"@tcp("+DBHost+":"+DBPort+")/"+DBName)
	for err != nil {
		log.Error().Msg("connection error :" + err.Error())
		time.Sleep(time.Second * 1)
		db, err = sql.Open(DBDriver, DBUser+":"+DBPass+"@tcp("+DBHost+":"+DBPort+")/"+DBName)
	}

	return db
}

// =========== //
// == Table == //
// =========== //

func ClearDBTablesMySQL() error {
	db := ConnectMySQL()
	defer db.Close()

	query := "DELETE FROM " + TableNetworkFlow
	if _, err := db.Query(query); err != nil {
		return err
	}

	query = "DELETE FROM " + TableDiscoveredPolicy
	if _, err := db.Query(query); err != nil {
		return err
	}

	return nil
}

func CreateTableNetworkPolicyMySQL() error {
	db := ConnectMySQL()
	defer db.Close()

	tableName := TableDiscoveredPolicy

	query :=
		"CREATE TABLE IF NOT EXISTS `" + tableName + "` (" +
			"	`id` int NOT NULL AUTO_INCREMENT," +
			"	`apiVersion` varchar(20) DEFAULT NULL," +
			"	`kind` varchar(20) DEFAULT NULL," +
			"	`flow_ids` JSON DEFAULT NULL," +
			"	`name` varchar(50) DEFAULT NULL," +
			"	`cluster_name` varchar(50) DEFAULT NULL," +
			"	`namespace` varchar(50) DEFAULT NULL," +
			"	`type` varchar(10) DEFAULT NULL," +
			"	`rule` varchar(30) DEFAULT NULL," +
			"	`status` varchar(10) DEFAULT NULL," +
			"	`outdated` varchar(50) DEFAULT NULL," +
			"	`spec` JSON DEFAULT NULL," +
			"	`generatedTime` bigint NOT NULL," +
			// "	`updatedTime` bigint NOT NULL," +
			"	PRIMARY KEY (`id`)" +
			"  );"

	if _, err := db.Query(query); err != nil {
		return err
	}

	return nil
}

func CreateTableNetworkLogsMySQL() error {
	db := ConnectMySQL()
	defer db.Close()

	tableName := TableNetworkFlow

	query :=
		"CREATE TABLE IF NOT EXISTS `" + tableName + "` (" +
			"	`id` integer NOT NULL PRIMARY KEY AUTO_INCREMENT," +
			"	`time` INTEGER DEFAULT NULL," +
			"	`verdict` varchar(100) DEFAULT NULL," +
			"	`drop_reason` INTEGER DEFAULT NULL," +
			"	`ethernet` JSON DEFAULT NULL," +
			"	`ip` JSON DEFAULT NULL," +
			"	`l4` JSON DEFAULT NULL," +
			"	`source` JSON DEFAULT NULL," +
			"	`destination` JSON DEFAULT NULL," +
			"	`type` INTEGER DEFAULT NULL," +
			"	`l7` JSON DEFAULT NULL," +
			"	`reply` BOOLEAN," +
			"	`src_cluster_name` varchar(100) DEFAULT NULL," +
			"	`dest_cluster_name` varchar(100) DEFAULT NULL," +
			"	`src_pod_name` varchar(100) DEFAULT NULL," +
			"	`dest_pod_name` varchar(100) DEFAULT NULL," +
			"	`node_name` varchar(100) DEFAULT NULL," +
			"	`event_type` JSON DEFAULT NULL," +
			"	`source_service` JSON DEFAULT NULL," +
			"	`destination_service` JSON DEFAULT NULL," +
			"	`traffic_direction` INTEGER DEFAULT NULL," +
			"	`policy_match_type` INTEGER DEFAULT NULL," +
			"	`trace_observation_point` INTEGER DEFAULT NULL," +
			"	`summary` varchar(1000) DEFAULT NULL" +
			" 	);"

	_, err := db.Query(query)
	return err
}

// QueryBaseSimple
var QueryBaseSimple string = "select id,time,traffic_direction,verdict,policy_match_type,drop_reason,event_type,source,destination,ip,l4,l7 from "

// flowScannerToCiliumFlow scans the trafficflow.
func flowScannerToCiliumFlow(results *sql.Rows) ([]map[string]interface{}, error) {
	trafficFlows := []map[string]interface{}{}
	var err error

	for results.Next() {
		var id, time, verdict, policyMatchType, dropReason, direction uint32
		var srcByte, destByte, eventTypeByte []byte
		var ipByte, l4Byte, l7Byte []byte

		err = results.Scan(
			&id,
			&time,
			&direction,
			&verdict,
			&policyMatchType,
			&dropReason,
			&eventTypeByte,
			&srcByte,
			&destByte,
			&ipByte,
			&l4Byte,
			&l7Byte,
		)

		if err != nil {
			log.Error().Msg("Error while scanning traffic flows :" + err.Error())
			return nil, err
		}

		flow := map[string]interface{}{
			"id":                id,
			"time":              time,
			"traffic_direction": direction,
			"verdict":           verdict,
			"policy_match_type": policyMatchType,
			"drop_reason":       dropReason,
			"event_type":        eventTypeByte,
			"source":            srcByte,
			"destination":       destByte,
			"ip":                ipByte,
			"l4":                l4Byte,
			"l7":                l7Byte,
		}

		trafficFlows = append(trafficFlows, flow)
	}

	return trafficFlows, nil
}

// GetTrafficFlowByTime function
func GetTrafficFlowByTime(startTime, endTime int64) ([]map[string]interface{}, error) {
	db := ConnectMySQL()
	defer db.Close()

	QueryBase := QueryBaseSimple + TableNetworkFlow

	rows, err := db.Query(QueryBase+" WHERE time >= ? and time < ?", int(startTime), int(endTime))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return flowScannerToCiliumFlow(rows)
}

// GetTrafficFlowByIDTime function
func GetTrafficFlowByIDTime(id, endTime int64) ([]map[string]interface{}, error) {
	db := ConnectMySQL()
	defer db.Close()

	QueryBase := QueryBaseSimple + TableNetworkFlow

	rows, err := db.Query(QueryBase+" WHERE id > ? ORDER BY id ASC ", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return flowScannerToCiliumFlow(rows)
}

// GetTrafficFlow function
func GetTrafficFlow() ([]map[string]interface{}, error) {
	db := ConnectMySQL()
	defer db.Close()

	QueryBase := QueryBaseSimple + TableNetworkFlow

	rows, err := db.Query(QueryBase)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return flowScannerToCiliumFlow(rows)
}

// GetNetworkPoliciesFromMySQL function
func GetNetworkPoliciesFromMySQL(namespace, status string) ([]types.KnoxNetworkPolicy, error) {
	db := ConnectMySQL()
	defer db.Close()

	policies := []types.KnoxNetworkPolicy{}
	var results *sql.Rows
	var err error

	query := "SELECT apiVersion,kind,name,namespace,type,rule,status,outdated,spec,generatedTime FROM " + TableDiscoveredPolicy
	if namespace != "" && status != "" {
		query = query + " WHERE namespace = ? and status = ? "
		results, err = db.Query(query, namespace, status)
	} else if namespace != "" {
		query = query + " WHERE namespace = ? "
		results, err = db.Query(query, namespace)
	} else if status != "" {
		query = query + " WHERE status = ? "
		results, err = db.Query(query, status)
	} else {
		results, err = db.Query(query)
	}

	defer results.Close()

	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	for results.Next() {
		policy := types.KnoxNetworkPolicy{}

		var name, namespace, policyType, rule, status string
		specByte := []byte{}
		spec := types.Spec{}

		if err := results.Scan(
			&policy.APIVersion,
			&policy.Kind,
			&name,
			&namespace,
			&policyType,
			&rule,
			&status,
			&policy.Outdated,
			&specByte,
			&policy.GeneratedTime,
		); err != nil {
			return nil, err
		}

		if err := json.Unmarshal(specByte, &spec); err != nil {
			return nil, err
		}

		policy.Metadata = map[string]string{
			"name":      name,
			"namespace": namespace,
			"type":      policyType,
			"rule":      rule,
			"status":    status,
		}
		policy.Spec = spec

		policies = append(policies, policy)
	}

	return policies, nil
}

// UpdateOutdatedPolicyFromMySQL ...
func UpdateOutdatedPolicyFromMySQL(outdatedPolicy string, latestPolicy string) error {
	db := ConnectMySQL()
	defer db.Close()

	var err error

	// set status -> outdated
	stmt1, err := db.Prepare("UPDATE " + TableDiscoveredPolicy + " SET status=? WHERE name=?")
	if err != nil {
		return err
	}
	defer stmt1.Close()

	_, err = stmt1.Exec("outdated", outdatedPolicy)
	if err != nil {
		return err
	}

	// set outdated -> latest' name
	stmt2, err := db.Prepare("UPDATE " + TableDiscoveredPolicy + " SET outdated=? WHERE name=?")
	if err != nil {
		return err
	}
	defer stmt2.Close()

	_, err = stmt2.Exec(latestPolicy, outdatedPolicy)
	if err != nil {
		return err
	}

	return nil
}

// insertDiscoveredPolicy function
func insertDiscoveredPolicy(db *sql.DB, policy types.KnoxNetworkPolicy) error {
	stmt, err := db.Prepare("INSERT INTO " + TableDiscoveredPolicy + "(apiVersion,kind,name,namespace,type,rule,status,outdated,spec,generatedTime) values(?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	specPointer := &policy.Spec
	spec, err := json.Marshal(specPointer)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(policy.APIVersion,
		policy.Kind,
		policy.Metadata["name"],
		policy.Metadata["namespace"],
		policy.Metadata["type"],
		policy.Metadata["rule"],
		policy.Metadata["status"],
		policy.Outdated,
		spec,
		policy.GeneratedTime)
	if err != nil {
		return err
	}

	return nil
}

// InsertDiscoveredPoliciesToMySQL function
func InsertDiscoveredPoliciesToMySQL(policies []types.KnoxNetworkPolicy) error {
	db := ConnectMySQL()
	defer db.Close()

	for _, policy := range policies {
		if err := insertDiscoveredPolicy(db, policy); err != nil {
			return err
		}
	}

	return nil
}

// InsertNetworkLogsMySQL -- Update existing log with time and count
func InsertNetworkLogsMySQL(netLogs []types.NetworkLogRaw) error {
	var err error = nil
	db := ConnectMySQL()
	defer db.Close()

	for _, netLog := range netLogs {
		if err := insertNetLogMySQL(db, netLog); err != nil {
			log.Error().Msg(err.Error())
		}
	}
	return err
}

func insertNetLogMySQL(db *sql.DB, netLog types.NetworkLogRaw) error {
	var err error

	insertQueryString := `(time,verdict,drop_reason,ip,l4,source,destination,l7,event_type,traffic_direction,policy_match_type) 
		VALUES (?,?,?,?,?,?,?,?,?,?,?)`

	query := "INSERT INTO " + TableNetworkFlow + insertQueryString

	insertStmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer insertStmt.Close()

	// var id, time, verdict, policyMatchType, dropReason, direction uint32
	_, err = insertStmt.Exec(
		netLog.Time,
		netLog.Verdict,
		netLog.DropReason,
		netLog.IP,
		netLog.L4,
		netLog.Source,
		netLog.Destination,
		netLog.L7,
		netLog.EventType,
		netLog.TrafficDirection,
		netLog.PolicyMatchType)

	if err != nil {
		log.Error().Msg(err.Error())
	}

	return err
}
