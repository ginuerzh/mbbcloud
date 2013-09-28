// app
package models

type App struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Icon        string `json:"icon"`
	RUrl        string `json:"url"`
	IUrl        string `json:"curl"`
	AUrl        string `json:"-"`
}
