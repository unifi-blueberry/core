package addon

import (
	"context"
	"github.com/bufbuild/connect-go"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&addonv1alpha1.Addon{})
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&addonv1alpha1.ListAddonsResponse{
		Addons: []*addonv1alpha1.Addon{
			{Name: "podman"},
		},
	})

	return res, nil
}
