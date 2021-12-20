package main

import (
	"context"
	"fmt"
	"log"
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

func getStatus(c echo.Context) error {

	ctx := context.Background()
	var status []Status
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

	cnt, err := db.NewSelect().Model(&status).OrderExpr("order_no DESC").Where(where).ScanAndCount(ctx)
	if err != nil {
		panic(err)
	}
	if cnt == 0 {
		//status :=
		myslice := []int{}
		//	fmt.Println(")))))))))))))")
		fmt.Println(myslice)
		return c.JSON(http.StatusOK, myslice)

	}
	return c.JSON(http.StatusOK, status)
}

func getStatusById(c echo.Context) error {

	barCodeId := c.Param("id")

	ctx := context.Background()
	var status Status
	err := db.NewSelect().Model(&status).Where("id = ?", barCodeId).Scan(ctx)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, status)
}

func deleteStatusById(c echo.Context) error {

	barCodeId := c.Param("id")
	fmt.Println(barCodeId)

	ctx := context.Background()
	var status Status
	_, err := db.NewDelete().Model(&status).Where("id = ?", barCodeId).Exec(ctx)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, nil)
}

func SaveStatus(c echo.Context) error {

	status := &Status{}

	err := c.Bind(status)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &GenericErrorResponse{
			false,
			err.Error(),
			"Invalid data binding for create test params1",
		})
	}
	ctx := context.Background()
	_, err = db.NewInsert().Model(status).ExcludeColumn("updated_at", "id", "created_at", "updated_by").Exec(ctx)
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

func updateStatus(c echo.Context) error {

	status := &Status{}
	statusID := c.Param("id")

	err := c.Bind(status)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &GenericErrorResponse{
			false,
			err.Error(),
			"Invalid data binding for create test params1",
		})
	}
	ctx := context.Background()
	_, err = db.NewUpdate().Model(status).
		ExcludeColumn("created_at", "created_by", "updated_at").
		Where("id = ?", statusID).
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

func bulkUpdateStatus(c echo.Context) error {

	status := &[]Status{}
	//	statusID := c.Param("id")

	err := c.Bind(status)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &GenericErrorResponse{
			false,
			err.Error(),
			"Invalid data binding for create test params1",
		})
	}
	values := db.NewValues(status)
	ctx := context.Background()
	_, err = db.NewUpdate().
		With("_data", values).
		Model((*Status)(nil)).
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
