// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.0--rc2
// source: tours_service.proto

package tours

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TourClient is the client API for Tour service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TourClient interface {
	// Checkpoints
	GetAllCheckpoints(ctx context.Context, in *Page, opts ...grpc.CallOption) (*PagedCheckpoints, error)
	GetAllCheckpointsByTour(ctx context.Context, in *PageWithId, opts ...grpc.CallOption) (*CheckpointsResponse, error)
	GetCheckpointById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*CheckpointResponse, error)
	CreateCheckpoint(ctx context.Context, in *CreateCheckpointRequest, opts ...grpc.CallOption) (*CheckpointResponse, error)
	UpdateCheckpoint(ctx context.Context, in *UpdateCheckpointRequest, opts ...grpc.CallOption) (*CheckpointResponse, error)
	DeleteCheckpoint(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Blank, error)
	// Equipment
	GetAvailableEquipment(ctx context.Context, in *EquipmentIds, opts ...grpc.CallOption) (*EquipmentsResponse, error)
	GetAllEquipment(ctx context.Context, in *Page, opts ...grpc.CallOption) (*PagedEquipmentsResponse, error)
	GetEquipment(ctx context.Context, in *Id, opts ...grpc.CallOption) (*EquipmentResponse, error)
	CreateEquipment(ctx context.Context, in *EquipmentResponse, opts ...grpc.CallOption) (*EquipmentResponse, error)
	UpdateEquipment(ctx context.Context, in *UpdateEquipmentId, opts ...grpc.CallOption) (*EquipmentResponse, error)
	DeleteEquipment(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Blank, error)
	// Tour
	GetAllTour(ctx context.Context, in *Page, opts ...grpc.CallOption) (*PagedToursResponse, error)
	GetTour(ctx context.Context, in *Id, opts ...grpc.CallOption) (*TourResponse, error)
	CreateTour(ctx context.Context, in *TourResponse, opts ...grpc.CallOption) (*TourResponse, error)
	UpdateTour(ctx context.Context, in *UpdateTourId, opts ...grpc.CallOption) (*TourResponse, error)
	DeleteTour(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Blank, error)
	GetTourByAuthorId(ctx context.Context, in *TourByAuthorId, opts ...grpc.CallOption) (*ToursResponse, error)
	AddTourEquipment(ctx context.Context, in *TourEquipment, opts ...grpc.CallOption) (*Blank, error)
	DeleteTourEquipment(ctx context.Context, in *TourEquipment, opts ...grpc.CallOption) (*Blank, error)
}

type tourClient struct {
	cc grpc.ClientConnInterface
}

func NewTourClient(cc grpc.ClientConnInterface) TourClient {
	return &tourClient{cc}
}

func (c *tourClient) GetAllCheckpoints(ctx context.Context, in *Page, opts ...grpc.CallOption) (*PagedCheckpoints, error) {
	out := new(PagedCheckpoints)
	err := c.cc.Invoke(ctx, "/Tour/GetAllCheckpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tourClient) GetAllCheckpointsByTour(ctx context.Context, in *PageWithId, opts ...grpc.CallOption) (*CheckpointsResponse, error) {
	out := new(CheckpointsResponse)
	err := c.cc.Invoke(ctx, "/Tour/GetAllCheckpointsByTour", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tourClient) GetCheckpointById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*CheckpointResponse, error) {
	out := new(CheckpointResponse)
	err := c.cc.Invoke(ctx, "/Tour/GetCheckpointById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tourClient) CreateCheckpoint(ctx context.Context, in *CreateCheckpointRequest, opts ...grpc.CallOption) (*CheckpointResponse, error) {
	out := new(CheckpointResponse)
	err := c.cc.Invoke(ctx, "/Tour/CreateCheckpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tourClient) UpdateCheckpoint(ctx context.Context, in *UpdateCheckpointRequest, opts ...grpc.CallOption) (*CheckpointResponse, error) {
	out := new(CheckpointResponse)
	err := c.cc.Invoke(ctx, "/Tour/UpdateCheckpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tourClient) DeleteCheckpoint(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Blank, error) {
	out := new(Blank)
	err := c.cc.Invoke(ctx, "/Tour/DeleteCheckpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tourClient) GetAvailableEquipment(ctx context.Context, in *EquipmentIds, opts ...grpc.CallOption) (*EquipmentsResponse, error) {
	out := new(EquipmentsResponse)
	err := c.cc.Invoke(ctx, "/Tour/GetAvailableEquipment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tourClient) GetAllEquipment(ctx context.Context, in *Page, opts ...grpc.CallOption) (*PagedEquipmentsResponse, error) {
	out := new(PagedEquipmentsResponse)
	err := c.cc.Invoke(ctx, "/Tour/GetAllEquipment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tourClient) GetEquipment(ctx context.Context, in *Id, opts ...grpc.CallOption) (*EquipmentResponse, error) {
	out := new(EquipmentResponse)
	err := c.cc.Invoke(ctx, "/Tour/GetEquipment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tourClient) CreateEquipment(ctx context.Context, in *EquipmentResponse, opts ...grpc.CallOption) (*EquipmentResponse, error) {
	out := new(EquipmentResponse)
	err := c.cc.Invoke(ctx, "/Tour/CreateEquipment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tourClient) UpdateEquipment(ctx context.Context, in *UpdateEquipmentId, opts ...grpc.CallOption) (*EquipmentResponse, error) {
	out := new(EquipmentResponse)
	err := c.cc.Invoke(ctx, "/Tour/UpdateEquipment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tourClient) DeleteEquipment(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Blank, error) {
	out := new(Blank)
	err := c.cc.Invoke(ctx, "/Tour/DeleteEquipment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tourClient) GetAllTour(ctx context.Context, in *Page, opts ...grpc.CallOption) (*PagedToursResponse, error) {
	out := new(PagedToursResponse)
	err := c.cc.Invoke(ctx, "/Tour/GetAllTour", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tourClient) GetTour(ctx context.Context, in *Id, opts ...grpc.CallOption) (*TourResponse, error) {
	out := new(TourResponse)
	err := c.cc.Invoke(ctx, "/Tour/GetTour", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tourClient) CreateTour(ctx context.Context, in *TourResponse, opts ...grpc.CallOption) (*TourResponse, error) {
	out := new(TourResponse)
	err := c.cc.Invoke(ctx, "/Tour/CreateTour", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tourClient) UpdateTour(ctx context.Context, in *UpdateTourId, opts ...grpc.CallOption) (*TourResponse, error) {
	out := new(TourResponse)
	err := c.cc.Invoke(ctx, "/Tour/UpdateTour", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tourClient) DeleteTour(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Blank, error) {
	out := new(Blank)
	err := c.cc.Invoke(ctx, "/Tour/DeleteTour", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tourClient) GetTourByAuthorId(ctx context.Context, in *TourByAuthorId, opts ...grpc.CallOption) (*ToursResponse, error) {
	out := new(ToursResponse)
	err := c.cc.Invoke(ctx, "/Tour/GetTourByAuthorId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tourClient) AddTourEquipment(ctx context.Context, in *TourEquipment, opts ...grpc.CallOption) (*Blank, error) {
	out := new(Blank)
	err := c.cc.Invoke(ctx, "/Tour/AddTourEquipment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tourClient) DeleteTourEquipment(ctx context.Context, in *TourEquipment, opts ...grpc.CallOption) (*Blank, error) {
	out := new(Blank)
	err := c.cc.Invoke(ctx, "/Tour/DeleteTourEquipment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TourServer is the server API for Tour service.
// All implementations must embed UnimplementedTourServer
// for forward compatibility
type TourServer interface {
	// Checkpoints
	GetAllCheckpoints(context.Context, *Page) (*PagedCheckpoints, error)
	GetAllCheckpointsByTour(context.Context, *PageWithId) (*CheckpointsResponse, error)
	GetCheckpointById(context.Context, *Id) (*CheckpointResponse, error)
	CreateCheckpoint(context.Context, *CreateCheckpointRequest) (*CheckpointResponse, error)
	UpdateCheckpoint(context.Context, *UpdateCheckpointRequest) (*CheckpointResponse, error)
	DeleteCheckpoint(context.Context, *Id) (*Blank, error)
	// Equipment
	GetAvailableEquipment(context.Context, *EquipmentIds) (*EquipmentsResponse, error)
	GetAllEquipment(context.Context, *Page) (*PagedEquipmentsResponse, error)
	GetEquipment(context.Context, *Id) (*EquipmentResponse, error)
	CreateEquipment(context.Context, *EquipmentResponse) (*EquipmentResponse, error)
	UpdateEquipment(context.Context, *UpdateEquipmentId) (*EquipmentResponse, error)
	DeleteEquipment(context.Context, *Id) (*Blank, error)
	// Tour
	GetAllTour(context.Context, *Page) (*PagedToursResponse, error)
	GetTour(context.Context, *Id) (*TourResponse, error)
	CreateTour(context.Context, *TourResponse) (*TourResponse, error)
	UpdateTour(context.Context, *UpdateTourId) (*TourResponse, error)
	DeleteTour(context.Context, *Id) (*Blank, error)
	GetTourByAuthorId(context.Context, *TourByAuthorId) (*ToursResponse, error)
	AddTourEquipment(context.Context, *TourEquipment) (*Blank, error)
	DeleteTourEquipment(context.Context, *TourEquipment) (*Blank, error)
	mustEmbedUnimplementedTourServer()
}

// UnimplementedTourServer must be embedded to have forward compatible implementations.
type UnimplementedTourServer struct {
}

func (UnimplementedTourServer) GetAllCheckpoints(context.Context, *Page) (*PagedCheckpoints, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCheckpoints not implemented")
}
func (UnimplementedTourServer) GetAllCheckpointsByTour(context.Context, *PageWithId) (*CheckpointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCheckpointsByTour not implemented")
}
func (UnimplementedTourServer) GetCheckpointById(context.Context, *Id) (*CheckpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCheckpointById not implemented")
}
func (UnimplementedTourServer) CreateCheckpoint(context.Context, *CreateCheckpointRequest) (*CheckpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCheckpoint not implemented")
}
func (UnimplementedTourServer) UpdateCheckpoint(context.Context, *UpdateCheckpointRequest) (*CheckpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCheckpoint not implemented")
}
func (UnimplementedTourServer) DeleteCheckpoint(context.Context, *Id) (*Blank, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCheckpoint not implemented")
}
func (UnimplementedTourServer) GetAvailableEquipment(context.Context, *EquipmentIds) (*EquipmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableEquipment not implemented")
}
func (UnimplementedTourServer) GetAllEquipment(context.Context, *Page) (*PagedEquipmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllEquipment not implemented")
}
func (UnimplementedTourServer) GetEquipment(context.Context, *Id) (*EquipmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEquipment not implemented")
}
func (UnimplementedTourServer) CreateEquipment(context.Context, *EquipmentResponse) (*EquipmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEquipment not implemented")
}
func (UnimplementedTourServer) UpdateEquipment(context.Context, *UpdateEquipmentId) (*EquipmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEquipment not implemented")
}
func (UnimplementedTourServer) DeleteEquipment(context.Context, *Id) (*Blank, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEquipment not implemented")
}
func (UnimplementedTourServer) GetAllTour(context.Context, *Page) (*PagedToursResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTour not implemented")
}
func (UnimplementedTourServer) GetTour(context.Context, *Id) (*TourResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTour not implemented")
}
func (UnimplementedTourServer) CreateTour(context.Context, *TourResponse) (*TourResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTour not implemented")
}
func (UnimplementedTourServer) UpdateTour(context.Context, *UpdateTourId) (*TourResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTour not implemented")
}
func (UnimplementedTourServer) DeleteTour(context.Context, *Id) (*Blank, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTour not implemented")
}
func (UnimplementedTourServer) GetTourByAuthorId(context.Context, *TourByAuthorId) (*ToursResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTourByAuthorId not implemented")
}
func (UnimplementedTourServer) AddTourEquipment(context.Context, *TourEquipment) (*Blank, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTourEquipment not implemented")
}
func (UnimplementedTourServer) DeleteTourEquipment(context.Context, *TourEquipment) (*Blank, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTourEquipment not implemented")
}
func (UnimplementedTourServer) mustEmbedUnimplementedTourServer() {}

// UnsafeTourServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TourServer will
// result in compilation errors.
type UnsafeTourServer interface {
	mustEmbedUnimplementedTourServer()
}

func RegisterTourServer(s grpc.ServiceRegistrar, srv TourServer) {
	s.RegisterService(&Tour_ServiceDesc, srv)
}

func _Tour_GetAllCheckpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Page)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServer).GetAllCheckpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tour/GetAllCheckpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServer).GetAllCheckpoints(ctx, req.(*Page))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tour_GetAllCheckpointsByTour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageWithId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServer).GetAllCheckpointsByTour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tour/GetAllCheckpointsByTour",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServer).GetAllCheckpointsByTour(ctx, req.(*PageWithId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tour_GetCheckpointById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServer).GetCheckpointById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tour/GetCheckpointById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServer).GetCheckpointById(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tour_CreateCheckpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCheckpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServer).CreateCheckpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tour/CreateCheckpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServer).CreateCheckpoint(ctx, req.(*CreateCheckpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tour_UpdateCheckpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCheckpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServer).UpdateCheckpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tour/UpdateCheckpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServer).UpdateCheckpoint(ctx, req.(*UpdateCheckpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tour_DeleteCheckpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServer).DeleteCheckpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tour/DeleteCheckpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServer).DeleteCheckpoint(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tour_GetAvailableEquipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EquipmentIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServer).GetAvailableEquipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tour/GetAvailableEquipment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServer).GetAvailableEquipment(ctx, req.(*EquipmentIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tour_GetAllEquipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Page)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServer).GetAllEquipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tour/GetAllEquipment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServer).GetAllEquipment(ctx, req.(*Page))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tour_GetEquipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServer).GetEquipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tour/GetEquipment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServer).GetEquipment(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tour_CreateEquipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EquipmentResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServer).CreateEquipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tour/CreateEquipment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServer).CreateEquipment(ctx, req.(*EquipmentResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tour_UpdateEquipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEquipmentId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServer).UpdateEquipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tour/UpdateEquipment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServer).UpdateEquipment(ctx, req.(*UpdateEquipmentId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tour_DeleteEquipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServer).DeleteEquipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tour/DeleteEquipment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServer).DeleteEquipment(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tour_GetAllTour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Page)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServer).GetAllTour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tour/GetAllTour",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServer).GetAllTour(ctx, req.(*Page))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tour_GetTour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServer).GetTour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tour/GetTour",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServer).GetTour(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tour_CreateTour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TourResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServer).CreateTour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tour/CreateTour",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServer).CreateTour(ctx, req.(*TourResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tour_UpdateTour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTourId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServer).UpdateTour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tour/UpdateTour",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServer).UpdateTour(ctx, req.(*UpdateTourId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tour_DeleteTour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServer).DeleteTour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tour/DeleteTour",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServer).DeleteTour(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tour_GetTourByAuthorId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TourByAuthorId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServer).GetTourByAuthorId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tour/GetTourByAuthorId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServer).GetTourByAuthorId(ctx, req.(*TourByAuthorId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tour_AddTourEquipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TourEquipment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServer).AddTourEquipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tour/AddTourEquipment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServer).AddTourEquipment(ctx, req.(*TourEquipment))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tour_DeleteTourEquipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TourEquipment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServer).DeleteTourEquipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tour/DeleteTourEquipment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServer).DeleteTourEquipment(ctx, req.(*TourEquipment))
	}
	return interceptor(ctx, in, info, handler)
}

// Tour_ServiceDesc is the grpc.ServiceDesc for Tour service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tour_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Tour",
	HandlerType: (*TourServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllCheckpoints",
			Handler:    _Tour_GetAllCheckpoints_Handler,
		},
		{
			MethodName: "GetAllCheckpointsByTour",
			Handler:    _Tour_GetAllCheckpointsByTour_Handler,
		},
		{
			MethodName: "GetCheckpointById",
			Handler:    _Tour_GetCheckpointById_Handler,
		},
		{
			MethodName: "CreateCheckpoint",
			Handler:    _Tour_CreateCheckpoint_Handler,
		},
		{
			MethodName: "UpdateCheckpoint",
			Handler:    _Tour_UpdateCheckpoint_Handler,
		},
		{
			MethodName: "DeleteCheckpoint",
			Handler:    _Tour_DeleteCheckpoint_Handler,
		},
		{
			MethodName: "GetAvailableEquipment",
			Handler:    _Tour_GetAvailableEquipment_Handler,
		},
		{
			MethodName: "GetAllEquipment",
			Handler:    _Tour_GetAllEquipment_Handler,
		},
		{
			MethodName: "GetEquipment",
			Handler:    _Tour_GetEquipment_Handler,
		},
		{
			MethodName: "CreateEquipment",
			Handler:    _Tour_CreateEquipment_Handler,
		},
		{
			MethodName: "UpdateEquipment",
			Handler:    _Tour_UpdateEquipment_Handler,
		},
		{
			MethodName: "DeleteEquipment",
			Handler:    _Tour_DeleteEquipment_Handler,
		},
		{
			MethodName: "GetAllTour",
			Handler:    _Tour_GetAllTour_Handler,
		},
		{
			MethodName: "GetTour",
			Handler:    _Tour_GetTour_Handler,
		},
		{
			MethodName: "CreateTour",
			Handler:    _Tour_CreateTour_Handler,
		},
		{
			MethodName: "UpdateTour",
			Handler:    _Tour_UpdateTour_Handler,
		},
		{
			MethodName: "DeleteTour",
			Handler:    _Tour_DeleteTour_Handler,
		},
		{
			MethodName: "GetTourByAuthorId",
			Handler:    _Tour_GetTourByAuthorId_Handler,
		},
		{
			MethodName: "AddTourEquipment",
			Handler:    _Tour_AddTourEquipment_Handler,
		},
		{
			MethodName: "DeleteTourEquipment",
			Handler:    _Tour_DeleteTourEquipment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tours_service.proto",
}
