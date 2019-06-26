package main

import (
	"github.com/transcom/nom/pkg/gen/ordersapi/models"
)

type CategorizedRank struct {
	officer  bool
	paygrade models.Rank
}

var paygradeToRank = map[string]CategorizedRank{
	"E01": {officer: false, paygrade: models.RankE1},
	"E02": {officer: false, paygrade: models.RankE2},
	"E03": {officer: false, paygrade: models.RankE3},
	"E04": {officer: false, paygrade: models.RankE4},
	"E05": {officer: false, paygrade: models.RankE5},
	"E06": {officer: false, paygrade: models.RankE6},
	"E07": {officer: false, paygrade: models.RankE7},
	"E08": {officer: false, paygrade: models.RankE8},
	"E09": {officer: false, paygrade: models.RankE9},
	"O01": {officer: true, paygrade: models.RankO1},
	"O02": {officer: true, paygrade: models.RankO2},
	"O03": {officer: true, paygrade: models.RankO3},
	"O04": {officer: true, paygrade: models.RankO4},
	"O05": {officer: true, paygrade: models.RankO5},
	"O06": {officer: true, paygrade: models.RankO6},
	"O07": {officer: true, paygrade: models.RankO7},
	"O08": {officer: true, paygrade: models.RankO8},
	"O09": {officer: true, paygrade: models.RankO9},
	"O10": {officer: true, paygrade: models.RankO10},
	"W01": {officer: true, paygrade: models.RankW1},
	"W02": {officer: true, paygrade: models.RankW2},
	"W03": {officer: true, paygrade: models.RankW3},
	"W04": {officer: true, paygrade: models.RankW4},
	"W05": {officer: true, paygrade: models.RankW5},
}
