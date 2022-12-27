package core

import (
	"context"
	"github.com/bufbuild/connect-go"
	"net/http"

	"buf.build/gen/go/unifi-blueberry/core/bufbuild/connect-go/core/v1alpha1/corev1alpha1connect"
	corev1alpha1 "buf.build/gen/go/unifi-blueberry/core/protocolbuffers/go/core/v1alpha1"
)

type Server struct{}

func (s *Server) GetVersionInfo(
	_ context.Context,
	_ *connect.Request[corev1alpha1.GetVersionInfoRequest],
) (*connect.Response[corev1alpha1.GetVersionInfoResponse], error) {

	return connect.NewResponse(&corev1alpha1.GetVersionInfoResponse{
		Version:   BuildVersion,
		Platform:  BuildPlatform,
		GitCommit: BuildGitCommit,
		GoVersion: BuildGoVersion,
	}), nil
}

func (s *Server) GetPlatformInfo(
	_ context.Context,
	_ *connect.Request[corev1alpha1.GetPlatformInfoRequest],
) (*connect.Response[corev1alpha1.GetPlatformInfoResponse], error) {

	pi, err := getPlatformInfo()
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(pi), nil
}

func getPlatformInfo() (*corev1alpha1.GetPlatformInfoResponse, error) {
	return &corev1alpha1.GetPlatformInfoResponse{
		Blueberry: &corev1alpha1.BlueberryPlatformInfo{
			Version: BuildVersion,
		},
		Unifi: &corev1alpha1.UnifiPlatformInfo{
			SubsystemId:       "ea15",
			Model:             "Test Device",
			ModelShort:        "TD",
			Family:            "Test Family",
			FamilyShort:       "TF",
			Firmware:          "0.0.0",
			FirmwareDetail:    "0.0.0",
			FirmwareDiscovery: "0.0.0",
			Ipv4:              "127.0.0.1",
			Mac:               "aa:bb:cc:dd:ee",
		},
	}, nil
}

func RegisterServer(mux *http.ServeMux) {
	server := &Server{}
	path, handler := corev1alpha1connect.NewCoreServiceHandler(server)
	mux.Handle(path, handler)
}
