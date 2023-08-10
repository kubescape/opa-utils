package v2

import (
	"time"

	"github.com/francoispqt/gojay"
)

/*
  responsible on fast unmarshaling of various COMMON posture report v2 structure for basic validation

*/
// UnmarshalJSONObject - File inside a pkg
func (r *PostureReport) UnmarshalJSONObject(dec *gojay.Decoder, key string) (err error) {

	switch key {
	case "customerGUID":
		err = dec.String(&(r.CustomerGUID))

	case "clusterName":
		err = dec.String(&(r.ClusterName))

	case "reportGUID":
		err = dec.String(&(r.ReportID))
	case "jobID":
		err = dec.String(&(r.JobID))
	case "generationTime":
		err = dec.Time(&(r.ReportGenerationTime), time.RFC3339)
		r.ReportGenerationTime = r.ReportGenerationTime.Local()
	case "metadata":
		err = dec.Object(&(r.Metadata))
	}
	return err

}

// func (files *PkgFiles) UnmarshalJSONArray(dec *gojay.Decoder) error {
// 	lae := PackageFile{}
// 	if err := dec.Object(&lae); err != nil {
// 		return err
// 	}

// 	*files = append(*files, lae)
// 	return nil
// }

func (file *PostureReport) NKeys() int {
	return 0
}

// UnmarshalJSONObject unmarshals incoming JSON data into a Metadata object
func (m *Metadata) UnmarshalJSONObject(dec *gojay.Decoder, key string) (err error) {

	switch key {
	case "scanMetadata":
		err = dec.Object(&(m.ScanMetadata))

	case "clusterMetadata":
		err = dec.Object(&(m.ClusterMetadata))

	case "targetMetadata":
		err = dec.Object(&(m.ContextMetadata))
	}

	return err

}

func (file *Metadata) NKeys() int {
	return 0
}

func (c *ContextMetadata) UnmarshalJSONObject(dec *gojay.Decoder, key string) (err error) {
	switch key {
	case "clusterContextMetadata":
		clusterMetadata := &ClusterMetadata{}
		if err = dec.Object(clusterMetadata); err == nil {
			c.ClusterContextMetadata = clusterMetadata
		}
	case "gitRepoContextMetadata":
		repoContextMetadata := &RepoContextMetadata{}
		if err = dec.Object(repoContextMetadata); err == nil {
			c.RepoContextMetadata = repoContextMetadata
		}
	}
	return err
}

func (c *ContextMetadata) NKeys() int {
	return 0
}

func (c *RepoContextMetadata) UnmarshalJSONObject(dec *gojay.Decoder, key string) (err error) {
	switch key {
	case "provider":
		err = dec.String(&(c.Provider))
	case "repo":
		err = dec.String(&(c.Repo))
	case "owner":
		err = dec.String(&(c.Owner))
	case "branch":
		err = dec.String(&(c.Branch))
	case "remoteURL":
		err = dec.String(&(c.RemoteURL))
	}
	return err
}

func (c *RepoContextMetadata) NKeys() int {
	return 0
}

// UnmarshalJSONObject unmarshals incoming JSON data into a ScanMetadata object
func (m *ScanMetadata) UnmarshalJSONObject(dec *gojay.Decoder, key string) (err error) {

	switch key {
	case "format": // string
		err = dec.String(&(m.Format))
	case "formats":
		err = dec.SliceString(&(m.Formats))
	case "excludedNamespaces": // []string
		err = dec.SliceString(&(m.ExcludedNamespaces))
	case "includeNamespaces": // []string
		err = dec.SliceString(&(m.IncludeNamespaces))
	case "failThreshold": // float32
		err = dec.Float32(&(m.FailThreshold))
	case "submit": // bool
		err = dec.Bool(&(m.Submit))
	case "hostScanner": // bool
		err = dec.Bool(&(m.HostScanner))
	case "logger": // string
		err = dec.String(&(m.Logger))
	case "targetType": // string
		err = dec.String(&(m.TargetType))
	case "targetNames": // []string
		err = dec.SliceString(&(m.TargetNames))
	case "useExceptions": // string
		err = dec.String(&(m.UseExceptions))
	case "controlsInputs": // string
		err = dec.String(&(m.ControlsInputs))
	case "verboseMode": // bool
		err = dec.Bool(&(m.VerboseMode))
	}
	return err

}

func (file *ScanMetadata) NKeys() int {
	return 0
}

// UnmarshalJSONObject unmarshals incoming JSON data into a ClusterMetadata object
func (m *ClusterMetadata) UnmarshalJSONObject(dec *gojay.Decoder, key string) (err error) {

	switch key {

	case "numberOfWorkerNodes": //int
		err = dec.Int(&(m.NumberOfWorkerNodes))

	case "cloudProvider": //string
		err = dec.String(&(m.CloudProvider))

	case "contextName": //string
		err = dec.String(&(m.ContextName))

	}
	return err

}

func (file *ClusterMetadata) NKeys() int {
	return 0
}
