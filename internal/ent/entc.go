//go:build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := entc.Generate(
		"./database/schema",
		&gen.Config{
			Target:  "./ent/",
			Schema:  "us-soccer-go-test/internal/database/schema",
			Package: "us-soccer-go-test/internal/ent",
			Features: []gen.Feature{
				gen.FeaturePrivacy,
				gen.FeatureEntQL,
				gen.FeatureSnapshot,
				gen.FeatureUpsert,
				gen.FeatureModifier,
			},
		},
	)
	checkError(err)
}
