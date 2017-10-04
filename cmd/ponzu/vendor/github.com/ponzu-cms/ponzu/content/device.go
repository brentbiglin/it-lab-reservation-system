package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Device struct {
	item.Item

	DescriptorName    string `json:"descriptor_name"`
	Model             string `json:"model"`
	Photo             string `json:"photo"`
	IncludedMaterials string `json:"included_materials"`
	Availability      string `json:"available"`
	BorrowerName      string `json:"borrower_name"`
	BorrowerEmail     string `json:"borrower_email"`
	Location          string `json:"location"`
	QualityCheck      string `json:"quality_check"`
	UTID              string `json:"utid"`
	SerialNumber      string `json:"serial_number"`
	SoftwareVersion   string `json:"software_version"`
	Recommendations   string `json:"recommendations"`
	Other             string `json:"other"`
}

// MarshalEditor writes a buffer of html to edit a Device within the CMS
// and implements editor.Editable
func (r *Device) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(r,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Device field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("DescriptorName", r, map[string]string{
				"label":       "Device Name",
				"type":        "text",
				"placeholder": "Enter the name of the resource here",
			}),
		},
		editor.Field{
			View: editor.Input("Model", r, map[string]string{
				"label":       "Model",
				"type":        "text",
				"placeholder": "Enter the model number here",
			}),
		},
		editor.Field{
			View: editor.File("Photo", r, map[string]string{
				"label":       "Image Upload",
				"placeholder": "Upload the photo here",
			}),
		},
		editor.Field{
			View: editor.Textarea("IncludedMaterials", r, map[string]string{
				"label":       "Included Materials",
				"placeholder": "Enter the included materials here",
			}),
		},
		editor.Field{
			View: editor.Select("Availability", r, map[string]string{
				"label": "Availability",
			}, map[string]string{
				// "value": "Display Name",
				"Available": "Available",
				"Reserved":  "Reserved",
				"In Use":    "In Use",
				"Due":       "Due",
			}),
		},
		editor.Field{
			View: editor.Input("BorrowerName", r, map[string]string{
				"label":       "Borrower Name",
				"type":        "text",
				"placeholder": "Enter the name of the borrower here",
			}),
		},
		editor.Field{
			View: editor.Input("BorrowerEmail", r, map[string]string{
				"label":       "Borrower Email",
				"type":        "text",
				"placeholder": "Enter the email address of the borrower here",
			}),
		},
		editor.Field{
			View: editor.Select("Location", r, map[string]string{
				"label": "Location",
			}, map[string]string{
				// "value": "Display Name",
				"IT Lab":   "IT Lab - UTA 1.210",
				"IT Suite": "IT Suite - UTA 5.454",
				"AV Room":  "AV Room - UTA 1.210DA",
			}),
		},
		editor.Field{
			View: editor.Select("QualityCheck", r, map[string]string{
				"label": "Quality Check",
			}, map[string]string{
				// "value": "Display Name",
				"Excellent": "Excellent",
				"Good":      "Good",
				"Fair":      "Fair",
				"Poor":      "Poor",
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
				"label":       "Serial Number",
				"type":        "text",
				"placeholder": "Enter the serial number here",
			}),
		},
		editor.Field{
			View: editor.Input("SoftwareVersion", r, map[string]string{
				"label":       "Software Version",
				"type":        "text",
				"placeholder": "Enter the software version here",
			}),
		},
		editor.Field{
			View: editor.Input("Recommendations", r, map[string]string{
				"label":       "Recommendations",
				"type":        "text",
				"placeholder": "Enter any recommendations here",
			}),
		},
		editor.Field{
			View: editor.Input("Other", r, map[string]string{
				"label":       "Other Information",
				"type":        "text",
				"placeholder": "Enter any other relevant information here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Device editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Device"] = func() interface{} { return new(Device) }
}

// String defines how a Device is printed. Update it using more descriptive
// fields from the Device struct type
func (r *Device) String() string {
	return r.Model
}
