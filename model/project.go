package model

import (
	"encoding/json"
	"log"
	"time"

	"github.com/xujiajun/nutsdb"
)

type Project struct {
	ID           string
	Name         string
	Introduction string
	GitURL       string
	Owner        string
	ModifiedBy   string
	ModifiedAt   time.Time
}

const ProjectBucket = "project"

func (p *Project) Save() {
	p.ModifiedAt = time.Now()
	if err := db.Update(putKVFunc(ProjectBucket, []byte(p.ID), p.toJSON())); err != nil {
		log.Fatal(err)
	}
}

func GetAllProjects() []Project {
	var projects []Project
	if err := db.View(
		func(tx *nutsdb.Tx) error {
			entries, err := tx.GetAll(ProjectBucket)
			if err != nil {
				return err
			}
			projects = make([]Project, len(entries))
			for i, entry := range entries {
				var p Project
				json.Unmarshal(entry.Value, &p)
				projects[i] = p
			}
			return nil
		}); err != nil {
		log.Println(err)
	}
	return projects
}

func (p *Project) toJSON() []byte {
	json, _ := json.Marshal(p)
	return json
}
