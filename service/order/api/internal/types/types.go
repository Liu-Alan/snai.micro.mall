// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type CreateRequest struct {
	Uid    int64 `json:"uid"`
	Pid    int64 `json:"pid"`
	Amount int64 `json:"amount"`
	Status int64 `json:"status"`
}

type CreateResponse struct {
	Id int64 `json:"id"`
}

type DetailRequest struct {
	Id int64 `json:"id"`
}

type DetailResponse struct {
	Id     int64 `json:"id"`
	Uid    int64 `json:"uid"`
	Pid    int64 `json:"pid"`
	Amount int64 `json:"amount"`
	Status int64 `json:"status"`
}

type ListRequest struct {
	Uid int64 `json:"uid"`
}

type ListResponse struct {
	Id     int64 `json:"id"`
	Uid    int64 `json:"uid"`
	Pid    int64 `json:"pid"`
	Amount int64 `json:"amount"`
	Status int64 `json:"status"`
}

type RemoveRequest struct {
	Id int64 `json:"id"`
}

type RemoveResponse struct {
}

type UpdateRequest struct {
	Id     int64 `json:"id"`
	Uid    int64 `json:"uid,optional"`
	Pid    int64 `json:"pid,optional"`
	Amount int64 `json:"amount,optional"`
	Status int64 `json:"status,optional"`
}

type UpdateResponse struct {
}
