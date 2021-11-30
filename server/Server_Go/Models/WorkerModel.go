package Models

type Worker struct {
	
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Worker_type string `json:"worker_type"`
	Antiguedad string `json:"Antiguedad"`
 
}

func (b *Worker) TableName() string {
	return "worker"
}
