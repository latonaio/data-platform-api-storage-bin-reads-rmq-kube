package dpfm_api_output_formatter

import (
	"data-platform-api-storage-bin-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToGeneral(rows *sql.Rows) (*[]General, error) {
	defer rows.Close()
	general := make([]General, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.General{}

		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.StorageLocation,
			&pm.StorageBin,
			&pm.StorageType,
			&pm.XCoordinates,
			&pm.YCoordinates,
			&pm.ZCoordinates,
			&pm.Capacity,
			&pm.CapacityUnit,
			&pm.CapacityWeight,
			&pm.CapacityWeightUnit,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &general, err
		}

		data := pm
		general = append(general, General{
			BusinessPartner:    data.BusinessPartner,
			Plant:              data.Plant,
			StorageLocation:    data.StorageLocation,
			StorageBin:         data.StorageBin,
			StorageType:        data.StorageType,
			XCoordinates:       data.XCoordinates,
			YCoordinates:       data.YCoordinates,
			ZCoordinates:       data.ZCoordinates,
			Capacity:           data.Capacity,
			CapacityUnit:       data.CapacityUnit,
			CapacityWeight:     data.CapacityWeight,
			CapacityWeightUnit: data.CapacityWeightUnit,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &general, nil
	}

	return &general, nil
}
