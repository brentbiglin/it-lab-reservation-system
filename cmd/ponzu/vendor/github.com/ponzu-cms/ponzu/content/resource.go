package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Resource struct {
	item.Item

	DescriptorName    string `json:"descriptor_name"`
	Model             string `json:"model"`
	IncludedMaterials string `json:"included_materials"`
	Available         bool   `json:"available"`
	Location          string `json:"location"`
	QualityCheck      string `json:"quality_check"`
	UTID              string `json:"utid"`
	SerialNumber      string `json:"serial_number"`
	SoftwareVersion   string `json:"software_version"`
	Recommendations   string `json:"recommendations"`
	Other             string `json:"other"`
}

// MarshalEditor writes a buffer of html to edit a Resource within the CMS
// and implements editor.Editable
func (r *Resource) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(r,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Resource field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("DescriptorName", r, map[string]string{
				"label":       "DescriptorName",
				"type":        "text",
				"placeholder": "Enter the Descriptor Name here",
			}),
		},
		editor.Field{
			View: editor.Input("Model", r, map[string]string{
				"label":       "Model",
				"type":        "text",
				"placeholder": "Enter the Model here",
			}),
		},
		editor.Field{
			View: editor.Richtext("IncludedMaterials", r, map[string]string{
				"label":       "IncludedMaterials",
				"placeholder": "Enter the IncludedMaterials here",
			}),
		},
		editor.Field{
			View: editor.Input("Available", r, map[string]string{
				"label":       "Available",
				"type":        "text",
				"placeholder": "Enter the Available here",
			}),
		},
		editor.Field{
			View: editor.Input("Location", r, map[string]string{
				"label":       "Location",
				"type":        "text",
				"placeholder": "Enter the Location here",
			}),
		},
		editor.Field{
			View: editor.Input("QualityCheck", r, map[string]string{
				"label":       "QualityCheck",
				"type":        "text",
				"placeholder": "Enter the Quality Check here",
			}),
		},
		editor.Field{
			View: editor.Input("UTID", r, map[string]string{
				"label":       "UTID",
				"type":        "text",
				"placeholder": "Enter the UTID here",
			}),
		},
		editor.Field{
			View: editor.Input("SerialNumber", r, map[string]string{
				"label":       "SerialNumber",
				"type":        "text",
				"placeholder": "Enter the Serial Number here",
			}),
		},
		editor.Field{
			View: editor.Input("SoftwareVersion", r, map[string]string{
				"label":       "SoftwareVersion",
				"type":        "text",
				"placeholder": "Enter the Software Version here",
			}),
		},
		editor.Field{
			View: editor.Input("Recommendations", r, map[string]string{
				"label":       "Recommendations",
				"type":        "text",
				"placeholder": "Enter the Recommendations here",
			}),
		},
		editor.Field{
			View: editor.Input("Other", r, map[string]string{
				"label":       "Other",
				"type":        "text",
				"placeholder": "Enter the Other here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Resource editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Resource"] = func() interface{} { return new(Resource) }
}

// String defines how a Resource is printed. Update it using more descriptive
// fields from the Resource struct type
func (r *Resource) String() string {
	return fmt.Sprintf("Resource: %s", r.UUID)
}
