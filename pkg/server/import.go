package server

import (
	"context"

	"github.com/selesy/hnjobs"
	pb "github.com/selesy/hnjobs"
	log "github.com/sirupsen/logrus"
)

/*
HNJobsServer ...

Trying out the server:

If reflection is not enabled, the -proto flag is needed as in the first
example.  If reflection is enabled the second (shorter) example can be
used:

grpcurl -plaintext -proto hnjobs.proto -d '{"id":"12345"}' localhost:8080 hnjobs.HNJobs.ImportWhoIsHiring
grpcurl -plaintext -d '{"id":"12345"}' localhost:8080 hnjobs.HNJobs.ImportWhoIsHiring

If reflection is enabled, it's also possible to get a list of the RPC
methods as follows:

grpcurl -plaintext localhost:8080 list
grpcurl -plaintext localhost:8080 list hnjobs.HNJobs
grpcurl -plaintext localhost:8080 describe hnjobs.HNJobs
grpcurl -plaintext localhost:8080 describe hnjobs.ItemReference
*/
type ImportServer struct {
	hnjobs.UnimplementedImportServer
}

func (s ImportServer) AddWhoIsHiring(context.Context, *pb.ItemReference) (*pb.ImportStatus, error) {
	log.Info("AddWhoIsHiring")

	return &pb.ImportStatus{
		ItemStatus: pb.ImportItemStatus_ACCEPTED,
		JobStatus:  pb.ImportJobStatus_ARCHIVED,
	}, nil
}

func (s ImportServer) UpsertJob(context.Context, *hnjobs.ItemReference) (*hnjobs.JobStatus, error) {
	log.Info("UpsertJob")

	return &hnjobs.JobStatus{}, nil
}
