package halladministration

type Hall struct {
	Id      int64
	Name    string
	Rows    int64
	Columns int64
}

var halls []Hall

func InitDB() {
	halls = []Hall{
		{Id: int64(1), Name: "ThrillerHall", Rows: 4, Columns: 4},
		{Id: int64(2), Name: "HorrorHall", Rows: 3, Columns: 3},
	}
}

func GetAllHaals() []Hall {
	return halls
}
