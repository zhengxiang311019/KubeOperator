package repository

import (
	"github.com/KubeOperator/KubeOperator/pkg/db"
	"github.com/KubeOperator/KubeOperator/pkg/model"
)

type ClusterManifestRepository interface {
	List() ([]model.ClusterManifest, error)
	ListByStatus() ([]model.ClusterManifest, error)
	Save(manifest model.ClusterManifest) error
	Get(name string) (model.ClusterManifest, error)
}

func NewClusterManifestRepository() ClusterManifestRepository {
	return &clusterManifestRepository{}
}

type clusterManifestRepository struct {
}

func (c clusterManifestRepository) Get(name string) (model.ClusterManifest, error) {
	var manifest model.ClusterManifest
	err := db.DB.Where(model.ClusterManifest{Name: name}).First(&manifest).Error
	return manifest, err
}

func (c clusterManifestRepository) List() ([]model.ClusterManifest, error) {
	var manifests []model.ClusterManifest
	err := db.DB.Model(model.ClusterManifest{}).Find(&manifests).Error
	return manifests, err
}

func (c clusterManifestRepository) ListByStatus() ([]model.ClusterManifest, error) {
	var manifests []model.ClusterManifest
	err := db.DB.Where(model.ClusterManifest{IsActive: true}).Model(model.ClusterManifest{}).Find(&manifests).Error
	return manifests, err
}

func (c clusterManifestRepository) Save(manifest model.ClusterManifest) error {
	if db.DB.NewRecord(manifest) {
		return db.DB.Create(&manifest).Error
	} else {
		return db.DB.Save(&manifest).Error
	}
}
