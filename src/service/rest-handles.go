package rlr

import (
	"latoken/relayer-smart-contract/src/models"
)

// Status ...
func (r *RelayerSRV) Status() (*models.RelayerStatus, error) {
	relayerStatus := &models.RelayerStatus{
		Mode:    "",
		Workers: make(map[string]models.WorkerStatus),
	}

	// for k, w := range r.Workers {
	// 	fmt.Println("----- NAME: ", k)
	// 	var worker models.WorkerStatus
	// 	// get block logs
	// 	currentBlockLog, err := r.storage.GetCurrentBlockLog(w.GetChain())
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	worker.SyncHeight = currentBlockLog.Height
	// 	worker.LastBlockFetchedAt = time.Unix(currentBlockLog.CreateTime, 0)

	// 	height, err := w.GetHeight()
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	worker.Height = height

	// 	status, err := w.GetStatus("LA") //TODO
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	worker.Status = status

	// 	relayerStatus.Workers[k] = worker
	// }

	// fmt.Println(relayerStatus)

	return relayerStatus, nil
}

// CreateNewBindRequest ...
func (r *RelayerSRV) CreateNewBindRequest() {}

// GetRelayerAccountBalance ...
func (r *RelayerSRV) GetRelayerAccountBalance() {

}

// GetWorkers ...
func (r *RelayerSRV) GetWorkers() {
	//	r.Workers
}
