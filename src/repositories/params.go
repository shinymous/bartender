package repositories

// func Save(param models.Params) error {
// 	db.DB.Options().DB = 0
// 	err := db.DB.Set(db.DB.Context(), strconv.FormatUint(param.TvID, 10), param, 0).Err()
// 	return err
// }

// func FindByTv(tvID uint64) (models.Params, error) {
// 	db.DB.Options().DB = 0
// 	result, err := db.DB.Get(db.DB.Context(), strconv.FormatUint(tvID, 10)).Result()
// 	if err != nil {
// 		return models.Params{}, err
// 	}
// 	var param models.Params
// 	if err := json.Unmarshal([]byte(result), &param); err != nil {
// 		return models.Params{}, err
// 	}
// 	return param, nil
// }
