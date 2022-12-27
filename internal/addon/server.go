package addon

import (
	"context"
	"errors"
	"github.com/bufbuild/connect-go"
	"net/http"

	"buf.build/gen/go/unifi-blueberry/addon/bufbuild/connect-go/addon/v1alpha1/addonv1alpha1connect"
	addonv1alpha1 "buf.build/gen/go/unifi-blueberry/addon/protocolbuffers/go/addon/v1alpha1"
)

type Server struct{}

func RegisterServer(mux *http.ServeMux) {
	server := &Server{}
	path, handler := addonv1alpha1connect.NewAddonServiceHandler(server)
	mux.Handle(path, handler)
}

func (s *Server) ListAddons(
	_ context.Context,
	_ *connect.Request[addonv1alpha1.ListAddonsRequest],
) (*connect.Response[addonv1alpha1.ListAddonsResponse], error) {

	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("todo"))
}
