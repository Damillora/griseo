package migrations

import (
	"encoding/json/v2"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		jsonData := `{
			"createRule": null,
			"deleteRule": null,
			"fields": [
				{
					"autogeneratePattern": "[a-z0-9]{15}",
					"help": "",
					"hidden": false,
					"id": "text3208210256",
					"max": 15,
					"min": 15,
					"name": "id",
					"pattern": "^[a-z0-9]+$",
					"presentable": false,
					"primaryKey": true,
					"required": true,
					"system": true,
					"type": "text"
				},
				{
					"hidden": false,
					"id": "autodate2990389176",
					"name": "created",
					"onCreate": true,
					"onUpdate": false,
					"presentable": false,
					"system": false,
					"type": "autodate"
				},
				{
					"hidden": false,
					"id": "autodate3332085495",
					"name": "updated",
					"onCreate": true,
					"onUpdate": true,
					"presentable": false,
					"system": false,
					"type": "autodate"
				},
				{
					"help": "",
					"hidden": false,
					"id": "date1748787223",
					"max": "",
					"min": "",
					"name": "published",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "date"
				},
				{
					"autogeneratePattern": "",
					"help": "",
					"hidden": false,
					"id": "text724990059",
					"max": 0,
					"min": 0,
					"name": "title",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": true,
					"system": false,
					"type": "text"
				},
				{
					"autogeneratePattern": "",
					"help": "",
					"hidden": false,
					"id": "text2560465762",
					"max": 0,
					"min": 0,
					"name": "slug",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": false,
					"system": false,
					"type": "text"
				},
				{
					"help": "",
					"hidden": false,
					"id": "select2063623452",
					"maxSelect": 0,
					"name": "status",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "select",
					"values": [
						"draft",
						"published",
						"archived"
					]
				},
				{
					"cascadeDelete": false,
					"collectionId": "_pb_users_auth_",
					"help": "",
					"hidden": false,
					"id": "relation3182418120",
					"maxSelect": 0,
					"minSelect": 0,
					"name": "author",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "relation"
				},
				{
					"cascadeDelete": false,
					"collectionId": "pbc_3292755704",
					"help": "",
					"hidden": false,
					"id": "relation105650625",
					"maxSelect": 0,
					"minSelect": 0,
					"name": "category",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "relation"
				},
				{
					"help": "",
					"hidden": false,
					"id": "file534213990",
					"maxSelect": 0,
					"maxSize": 0,
					"mimeTypes": [],
					"name": "feature",
					"presentable": false,
					"protected": false,
					"required": false,
					"system": false,
					"thumbs": [],
					"type": "file"
				},
				{
					"autogeneratePattern": "",
					"help": "",
					"hidden": false,
					"id": "text1591429585",
					"max": 0,
					"min": 0,
					"name": "excerpt",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": false,
					"system": false,
					"type": "text"
				},
				{
					"convertURLs": false,
					"help": "",
					"hidden": false,
					"id": "editor4274335913",
					"maxSize": 0,
					"name": "content",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "editor"
				}
			],
			"id": "pbc_1125843985",
			"indexes": [],
			"listRule": null,
			"name": "posts",
			"system": false,
			"type": "base",
			"updateRule": null,
			"viewRule": null
		}`

		collection := &core.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1125843985")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
