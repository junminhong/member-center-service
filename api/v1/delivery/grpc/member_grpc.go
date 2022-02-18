package grpc

import (
	"context"
	"github.com/junminhong/member-center-service/api/v1/delivery/grpc/proto"
	"github.com/junminhong/member-center-service/domain"
	"github.com/junminhong/member-center-service/pkg/jwt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type memberHandler struct {
	memberRepo domain.MemberRepository
}

func (memberHandler *memberHandler) VerifyAtomicToken(ctx context.Context, request *proto.AtomicTokenAuthRequest) (*proto.AtomicTokenAuthResponse, error) {
	if !jwt.VerifyAtomicToken(request.AtomicToken) {
		return &proto.AtomicTokenAuthResponse{MemberUUID: ""}, nil
	}
	memberUUID := memberHandler.memberRepo.GetMemberUUIDByAtomicToken(request.AtomicToken)
	if memberUUID == "" {
		return &proto.AtomicTokenAuthResponse{MemberUUID: ""}, nil
	}
	return &proto.AtomicTokenAuthResponse{MemberUUID: memberUUID}, nil
}

func NewMemberGrpc(server *grpc.Server, lis net.Listener, memberRepo domain.MemberRepository) {
	handler := memberHandler{memberRepo: memberRepo}
	proto.RegisterMemberServiceServer(server, &handler)
	if err := server.Serve(lis); err != nil {
		log.Println(err.Error())
	}
}
