package handler

import (
	"errors"
	"fmt"
	"grpc-client/pb"
	"grpc-client/request"
	"grpc-client/response"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HttpHandler struct {
	UserClient pb.AttendanceServiceClient
}

func InitHttpHandler(e *echo.Echo, userClient pb.AttendanceServiceClient) {
	handler := &HttpHandler{
		UserClient: userClient,
	}

	route := e.Group("/api")

	route.POST("/checkin", handler.CheckIn)
}

func (h HttpHandler) CheckIn(c echo.Context) error {
	req := new(request.CheckInRequest)
	if err := c.Bind(req); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, errors.New("bad request"))
	}

	pbRequest := &pb.CheckInRequest{Username: req.Username, Datetime: req.Datetime}

	resp, err := h.UserClient.CheckIn(c.Request().Context(), pbRequest)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, errors.New("unable to hit grpc"))
	}

	response := response.CheckInResponse{
		Status:      resp.Status,
		Description: resp.Description,
	}

	return c.JSON(http.StatusOK, response)
}
