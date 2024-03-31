package models

/// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float32
	Data map[string]interface {} // In Go if you are not sure what type of data you are going to send then you can make use of the interface..
	CSRFToken string // CRSF stands for Cross Site Request Forgery Tokens...
	Flash string
	Warning string
	Error string
  }