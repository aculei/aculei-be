package models

type DatasetInfo struct {
	Badger       string `csv:"badger"`
	Buzzard      string `csv:"buzzard"`
	Cat          string `csv:"cat"`
	Cam1         string `csv:"cam_1"`
	Cam2         string `csv:"cam_2"`
	Cam3         string `csv:"cam_3"`
	Cam4         string `csv:"cam_4"`
	Cam5         string `csv:"cam_5"`
	Cam6         string `csv:"cam_6"`
	Cam7         string `csv:"cam_7"`
	Deer         string `csv:"deer"`
	Fox          string `csv:"fox"`
	Hare         string `csv:"hare"`
	Heron        string `csv:"heron"`
	Horse        string `csv:"horse"`
	Mallard      string `csv:"mallard"`
	Marten       string `csv:"marten"`
	PhotosAutumn string `csv:"photos_autumn"`
	PhotosWinter string `csv:"photos_winter"`
	PhotosSpring string `csv:"photos_spring"`
	PhotosSummer string `csv:"photos_summer"`
	Porcupine    string `csv:"porcupine"`
	Squirrel     string `csv:"squirrel"`
	TotalRecords string `csv:"total_records"`
	WildBoar     string `csv:"wild_boar"`
	Wolf         string `csv:"wolf"`
}

type DatasetInfoResponse struct {
	Badger       string `json:"badger"`
	Buzzard      string `json:"buzzard"`
	Cat          string `json:"cat"`
	Cam1         string `json:"cam_1"`
	Cam2         string `json:"cam_2"`
	Cam3         string `json:"cam_3"`
	Cam4         string `json:"cam_4"`
	Cam5         string `json:"cam_5"`
	Cam6         string `json:"cam_6"`
	Cam7         string `json:"cam_7"`
	Deer         string `json:"deer"`
	Fox          string `json:"fox"`
	Hare         string `json:"hare"`
	Heron        string `json:"heron"`
	Horse        string `json:"horse"`
	Mallard      string `json:"mallard"`
	Marten       string `json:"marten"`
	PhotosAutumn string `json:"photos_autumn"`
	PhotosWinter string `json:"photos_winter"`
	PhotosSpring string `json:"photos_spring"`
	PhotosSummer string `json:"photos_summer"`
	Porcupine    string `json:"porcupine"`
	Squirrel     string `json:"squirrel"`
	TotalRecords string `json:"total_records"`
	WildBoar     string `json:"wild_boar"`
	Wolf         string `json:"wolf"`
}
