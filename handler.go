package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{}
)

func status(c echo.Context) error {
	return c.JSON(http.StatusOK, fmt.Sprintf("%s version %s is up and running!", config.Get("name").(string), config.Get("version").(string)))
}

func getBarcodesSetting(c echo.Context) error {

	ctx := context.Background()
	var barcodesSetting []BarcodesSetting
	descriptionParam := c.QueryParam("description")
	nameParam := c.QueryParam("name")
	activeParam := c.QueryParam("active")

	fmt.Println(descriptionParam)
	//	if
	fmt.Println(len(descriptionParam))
	where := "deleted_at IS NULL"
	var filterQuery string
	if descriptionParam != "" {
		filterQuery = filterQuery + " AND " + "description like '%" + c.QueryParam("description") + "%'"
	}
	if nameParam != "" {
		filterQuery = filterQuery + " AND " + "name = '" + c.QueryParam("name") + "'"
	}
	if activeParam != "" {
		filterQuery = filterQuery + " AND " + "active = '" + c.QueryParam("active") + "'"
	}
	where = where + filterQuery
	fmt.Println(where)
	//descriptionFilter := descriptionParam+"= "+ c.ParamValues(description)

	cnt, err := db.NewSelect().Model(&barcodesSetting).OrderExpr("ordering_number DESC").Where(where).ScanAndCount(ctx)
	if err != nil {
		panic(err)
	}
	if cnt == 0 {
		//barcodesSetting :=
		myslice := []int{}
		//	fmt.Println(")))))))))))))")
		fmt.Println(myslice)
		return c.JSON(http.StatusOK, myslice)

	}
	return c.JSON(http.StatusOK, barcodesSetting)
}

func getBarcodesSettingById(c echo.Context) error {

	barCodeId := c.Param("id")

	ctx := context.Background()
	var barcodesSetting BarcodesSetting
	err := db.NewSelect().Model(&barcodesSetting).Where("id = ?", barCodeId).Scan(ctx)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, barcodesSetting)
}

func deleteBarcodesSettingById(c echo.Context) error {

	barCodeId := c.Param("id")
	fmt.Println(barCodeId)

	ctx := context.Background()
	var barcodesSetting BarcodesSetting
	_, err := db.NewDelete().Model(&barcodesSetting).Where("id = ?", barCodeId).Exec(ctx)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, barcodesSetting)
}

func SaveBarcodesSetting(c echo.Context) error {

	barcodesSetting := &BarcodesSetting{}

	err := c.Bind(barcodesSetting)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &GenericErrorResponse{
			false,
			err.Error(),
			"Invalid data binding for create test params1",
		})
	}
	ctx := context.Background()
	_, err = db.NewInsert().Model(barcodesSetting).ExcludeColumn("updated_at", "id", "created_at", "updated_by").Exec(ctx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &GenericErrorResponse{
			false,
			err.Error(),
			"Invalid data binding for create test params2",
		})
	}

	return c.JSON(http.StatusOK, GenericResponse{
		Success: true,
	})
}

func updateBarcodesSetting(c echo.Context) error {

	barcodesSetting := &BarcodesSetting{}
	barcodeId := c.Param("id")

	err := c.Bind(barcodesSetting)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &GenericErrorResponse{
			false,
			err.Error(),
			"Invalid data binding for create test params1",
		})
	}
	ctx := context.Background()
	_, err = db.NewUpdate().Model(barcodesSetting).
		ExcludeColumn("created_at", "created_by", "updated_at").
		Where("id = ?", barcodeId).
		Exec(ctx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &GenericErrorResponse{
			false,
			err.Error(),
			"Invalid data binding for create test params",
		})
	}

	return c.JSON(http.StatusOK, GenericResponse{
		Success: true,
	})
}

func bulkUpdateBarcodesSetting(c echo.Context) error {

	barcodesSetting := &[]BarcodesSetting{}
	//	barcodeId := c.Param("id")

	err := c.Bind(barcodesSetting)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &GenericErrorResponse{
			false,
			err.Error(),
			"Invalid data binding for create test params1",
		})
	}
	values := db.NewValues(barcodesSetting)
	ctx := context.Background()
	_, err = db.NewUpdate().
		With("_data", values).
		Model((*BarcodesSetting)(nil)).
		TableExpr("_data").
		Set("active = _data.active").
		Set("ordering_number = _data.ordering_number").
		Set("description = _data.description").
		Where("barcodes_setting.id = _data.id").
		Exec(ctx)

		/*
				res, err := db.NewUpdate().
			With("_data", values).
			Model((*Book)(nil)).
			TableExpr("_data").
			Set("title = _data.title").
			Set("text = _data.text").
			Where("book.id = _data.id").
			Exec(ctx)
		*/
	if err != nil {
		return c.JSON(http.StatusBadRequest, &GenericErrorResponse{
			false,
			err.Error(),
			"Invalid data binding for create test params",
		})
	}

	return c.JSON(http.StatusOK, GenericResponse{
		Success: true,
	})
}
