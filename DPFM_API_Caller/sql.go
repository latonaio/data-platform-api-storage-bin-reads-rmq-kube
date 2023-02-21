package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-storage-bin-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-storage-bin-reads-rmq-kube/DPFM_API_Output_Formatter"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var general *[]dpfm_api_output_formatter.General
	for _, fn := range accepter {
		switch fn {
		case "General":
			func() {
				general = c.General(mtx, input, output, errs, log)
			}()
		default:
		}
		if len(*errs) != 0 {
			break
		}
	}

	data := &dpfm_api_output_formatter.Message{
		General: general,
	}

	return data
}

func (c *DPFMAPICaller) General(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.General {
	businessPartner := input.General.BusinessPartner
	plant := input.General.Plant
	storageLocation := input.General.StorageLocation
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_storage_bin_general_data
		WHERE (BusinessPartner, Plant, StorageLocation) = (?, ?, ?);`, businessPartner, plant, storageLocation,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToGeneral(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
