package handles

import (
	"net/http"

	"github.com/ettoma/star/utils"
)

type endpoint struct {
	URL         string  `json:"url"`
	Method      string  `json:"method"`
	Description string  `json:"description"`
	Fields      []field `json:"required_fields"`
}

type field struct {
	Key  string `json:"key"`
	Type string `json:"type"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	endpoints := []*endpoint{
		{
			URL:         "/",
			Method:      "GET",
			Description: "Get all endpoints",
			Fields: []field{
				{
					"none",
					"",
				},
			},
		}, {
			URL:         "/users",
			Method:      "GET",
			Description: "Get all users",
			Fields: []field{
				{
					"none",
					"",
				},
			},
		}, {
			URL:         "/users/user",
			Method:      "GET",
			Description: "Get user by username or id (only send 1 field)",
			Fields: []field{
				{
					"id",
					"integer",
				}, {
					"username (min. 4 char)",
					"string",
				},
			},
		}, {
			URL:         "/users",
			Method:      "POST",
			Description: "Add a user",
			Fields: []field{

				{
					"name (min. 4 char)",
					"string",
				}, {

					"username (min. 4 char)",
					"string",
				},
			},
		}, {
			URL:         "/users",
			Method:      "DELETE",
			Description: "Delete a user by username or id (send only 1 field)",
			Fields: []field{
				{
					"id",
					"integer",
				}, {
					"username (min. 4 char)",
					"string",
				},
			},
		}, {
			URL:         "/kudos",
			Method:      "GET",
			Description: "Get all kudos",
			Fields: []field{
				{
					"none",
					"",
				},
			},
		}, {
			URL:         "/kudos",
			Method:      "POST",
			Description: "Add a kudos",
			Fields: []field{

				{
					"sender",
					"string",
				}, {

					"receiver",
					"string",
				},
				{

					"content",
					"string",
				},
			},
		}, {
			URL:         "/kudos",
			Method:      "DELETE",
			Description: "Delete a kudos by id",
			Fields: []field{
				{
					"id",
					"integer",
				},
			},
		},
		{
			URL:         "/kudos/user",
			Method:      "GET",
			Description: "Get all kudos for a user",
			Fields: []field{
				{
					"username",
					"string",
				},
			},
		},
	}

	utils.WriteJsonResponse(endpoints, w)

}
