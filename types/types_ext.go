// Package types contains the core type definitions for the Dagster Pipes protocol.
//
// These types are generated from JSON Schema using quicktype and should not be
// modified directly. They represent the data structures used for communication
// between Dagster and external processes.
//
// The main types include:
//   - PipesContextData: Context information passed from Dagster to the external process
//   - PipesMessage: Messages sent from the external process back to Dagster
//   - PipesMetadataValue: Typed metadata values that can be attached to materializations
//   - PipesException: Exception information for error reporting
//
// Most users should interact with these types through the higher-level APIs in
// the dagster_pipes and metadata packages rather than using them directly.
package types

// NewMessage creates a new PipesMessage with the specified method and parameters.
//
// This is a helper function that ensures the message includes the correct
// Dagster Pipes protocol version. It's used internally by the PipesContext
// methods to construct protocol-compliant messages.
//
// Parameters:
//   - method: The message method type (e.g., types.Opened, types.Closed, types.ReportAssetMaterialization)
//   - params: Optional parameters specific to the message method
//
// Returns a PipesMessage ready to be sent to Dagster.
func NewMessage(method Method, params map[string]any) *PipesMessage {
	return &PipesMessage{
		DagsterPipesVersion: "0.1",
		Method:              method,
		Params:              params,
	}
}
