package middleware

import (
	"fmt"
	"github.com/coalalib/coalago"
)

type CoalaMiddleware func(h coalago.CoAPResourceHandler) coalago.CoAPResourceHandler

func CoalaGroup(h coalago.CoAPResourceHandler, ms ...CoalaMiddleware) func(*coalago.CoAPMessage) *coalago.CoAPResourceHandlerResult {
	return func(message *coalago.CoAPMessage) (r *coalago.CoAPResourceHandlerResult) {
		lastRes := h
		for i := len(ms) - 1; i >= 0; i-- {
			lastRes = ms[i](lastRes)

		}
		return lastRes(message)
	}
}

func CoalaMetrics(path string, h coalago.CoAPResourceHandler) (string, coalago.CoAPResourceHandler) {
	return path, func(message *coalago.CoAPMessage) (r *coalago.CoAPResourceHandlerResult) {
		msgMethod := message.GetMethod()

		g := h(message)

		if g == nil {
			return g
		}

		method := ""
		switch msgMethod {
		case 1:
			method = "GET"
		case 2:
			method = "PUT"
		case 3:
			method = "POST"
		case 4:
			method = "DEL"
		}
		status := fmt.Sprint(g.Code)
		if g.Code == 68 || g.Code == 69 {
			status = "ok"
		}
		Stat.WithLabelValues(message.GetSchemeString(), method, path, status).Inc()
		return g
	}
}
