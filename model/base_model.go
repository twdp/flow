package model

type BaseModel struct {
	Name        string `xml:"name,attr"`
	DisplayName string `xml:"displayName,attr"`
}
