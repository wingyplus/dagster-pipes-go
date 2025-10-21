// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    assetCheckSeverity, err := UnmarshalAssetCheckSeverity(bytes)
//    bytes, err = assetCheckSeverity.Marshal()
//
//    pipesContextData, err := UnmarshalPipesContextData(bytes)
//    bytes, err = pipesContextData.Marshal()
//
//    pipesException, err := UnmarshalPipesException(bytes)
//    bytes, err = pipesException.Marshal()
//
//    pipesLogLevel, err := UnmarshalPipesLogLevel(bytes)
//    bytes, err = pipesLogLevel.Marshal()
//
//    pipesMessage, err := UnmarshalPipesMessage(bytes)
//    bytes, err = pipesMessage.Marshal()
//
//    pipesMetadataValue, err := UnmarshalPipesMetadataValue(bytes)
//    bytes, err = pipesMetadataValue.Marshal()

package dagster_pipes

import "bytes"
import "errors"

import "encoding/json"

func UnmarshalAssetCheckSeverity(data []byte) (AssetCheckSeverity, error) {
	var r AssetCheckSeverity
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AssetCheckSeverity) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPipesContextData(data []byte) (PipesContextData, error) {
	var r PipesContextData
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PipesContextData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPipesException(data []byte) (PipesException, error) {
	var r PipesException
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PipesException) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPipesLogLevel(data []byte) (PipesLogLevel, error) {
	var r PipesLogLevel
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PipesLogLevel) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPipesMessage(data []byte) (PipesMessage, error) {
	var r PipesMessage
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PipesMessage) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPipesMetadataValue(data []byte) (PipesMetadataValue, error) {
	var r PipesMetadataValue
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PipesMetadataValue) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// The serializable data passed from the orchestration process to the external process. This
// gets wrapped in a PipesContext.
type PipesContextData struct {
	AssetKeys             []string                         `json:"asset_keys,omitempty"`
	CodeVersionByAssetKey map[string]*string               `json:"code_version_by_asset_key,omitempty"`
	Extras                map[string]interface{}           `json:"extras"`
	JobName               *string                          `json:"job_name,omitempty"`
	PartitionKey          *string                          `json:"partition_key,omitempty"`
	PartitionKeyRange     *PartitionKeyRange               `json:"partition_key_range,omitempty"`
	PartitionTimeWindow   *PartitionTimeWindow             `json:"partition_time_window,omitempty"`
	ProvenanceByAssetKey  map[string]*ProvenanceByAssetKey `json:"provenance_by_asset_key"`
	RetryNumber           int64                            `json:"retry_number"`
	RunID                 string                           `json:"run_id"`
}

type PartitionKeyRange struct {
	End   *string `json:"end,omitempty"`
	Start *string `json:"start,omitempty"`
}

type PartitionTimeWindow struct {
	End   *string `json:"end,omitempty"`
	Start *string `json:"start,omitempty"`
}

type ProvenanceByAssetKey struct {
	CodeVersion       *string           `json:"code_version,omitempty"`
	InputDataVersions map[string]string `json:"input_data_versions,omitempty"`
	IsUserProvided    *bool             `json:"is_user_provided,omitempty"`
}

type PipesException struct {
	// exception that explicitly led to this exception
	Cause *PipesExceptionClass `json:"cause"`
	// exception that being handled when this exception was raised
	Context *ContextClass `json:"context,omitempty"`
	Message *string       `json:"message,omitempty"`
	// class name of Exception object
	Name  *string  `json:"name,omitempty"`
	Stack []string `json:"stack,omitempty"`
}

// exception that being handled when this exception was raised
type ContextClass struct {
	// exception that explicitly led to this exception
	Cause *PipesExceptionClass `json:"cause"`
	// exception that being handled when this exception was raised
	Context *ContextClass `json:"context,omitempty"`
	Message *string       `json:"message,omitempty"`
	// class name of Exception object
	Name  *string  `json:"name,omitempty"`
	Stack []string `json:"stack,omitempty"`
}

type PipesExceptionClass struct {
	// exception that explicitly led to this exception
	Cause *PipesExceptionClass `json:"cause"`
	// exception that being handled when this exception was raised
	Context *ContextClass `json:"context,omitempty"`
	Message *string       `json:"message,omitempty"`
	// class name of Exception object
	Name  *string  `json:"name,omitempty"`
	Stack []string `json:"stack,omitempty"`
}

type PipesMessage struct {
	// The version of the Dagster Pipes protocol
	DagsterPipesVersion string `json:"__dagster_pipes_version"`
	// Event type
	Method Method `json:"method"`
	// Event parameters
	Params map[string]interface{} `json:"params"`
}

type PipesMetadataValue struct {
	RawValue *RawValue `json:"raw_value"`
	Type     *Type     `json:"type,omitempty"`
}

type AssetCheckSeverity string

const (
	AssetCheckSeverityERROR AssetCheckSeverity = "ERROR"
	Warn                    AssetCheckSeverity = "WARN"
)

type PipesLogLevel string

const (
	Critical           PipesLogLevel = "CRITICAL"
	Debug              PipesLogLevel = "DEBUG"
	Info               PipesLogLevel = "INFO"
	PipesLogLevelERROR PipesLogLevel = "ERROR"
	Warning            PipesLogLevel = "WARNING"
)

// Event type
type Method string

const (
	Closed                     Method = "closed"
	Log                        Method = "log"
	Opened                     Method = "opened"
	ReportAssetCheck           Method = "report_asset_check"
	ReportAssetMaterialization Method = "report_asset_materialization"
	ReportCustomMessage        Method = "report_custom_message"
)

type Type string

const (
	Asset      Type = "asset"
	Bool       Type = "bool"
	DagsterRun Type = "dagster_run"
	Float      Type = "float"
	Infer      Type = "__infer__"
	Int        Type = "int"
	JSON       Type = "json"
	Job        Type = "job"
	Md         Type = "md"
	Notebook   Type = "notebook"
	Null       Type = "null"
	Path       Type = "path"
	Text       Type = "text"
	Timestamp  Type = "timestamp"
	URL        Type = "url"
)

type RawValue struct {
	AnythingArray []interface{}
	AnythingMap   map[string]interface{}
	Bool          *bool
	Double        *float64
	Integer       *int64
	String        *string
}

func (x *RawValue) UnmarshalJSON(data []byte) error {
	x.AnythingArray = nil
	x.AnythingMap = nil
	object, err := unmarshalUnion(data, &x.Integer, &x.Double, &x.Bool, &x.String, true, &x.AnythingArray, false, nil, true, &x.AnythingMap, false, nil, true)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *RawValue) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, x.Double, x.Bool, x.String, x.AnythingArray != nil, x.AnythingArray, false, nil, x.AnythingMap != nil, x.AnythingMap, false, nil, true)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
