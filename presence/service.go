package presence

import (
	"net"

	"github.com/BorisBorshevsky/timemock"
	"github.com/nymtech/nym-directory/models"
)

type service struct {
	db         IDb
	ipAssigner IPAssigner
}

// IService defines the REST service interface for presence.
type IService interface {
	AddCocoNodePresence(info models.CocoHostInfo, ip string)
	AddMixNodePresence(info models.MixHostInfo, ip string)
	AddMixProviderPresence(info models.MixProviderHostInfo, ip string)
	Topology() models.Topology
}

func NewService(db IDb) *service {
	ipa := ipAssigner{}
	return &service{
		db:         db,
		ipAssigner: &ipa,
	}
}

func (service *service) AddCocoNodePresence(info models.CocoHostInfo, ip string) {
	presence := models.CocoPresence{
		CocoHostInfo: info,
		LastSeen:     timemock.Now().UnixNano(),
	}
	presence.HostInfo.Host, _ = service.ipAssigner.AssignIP(ip, presence.Host)
	service.db.AddCoco(presence)
}

func (service *service) AddMixNodePresence(info models.MixHostInfo, ip string) {
	presence := models.MixNodePresence{
		MixHostInfo: info,
		LastSeen:    timemock.Now().UnixNano(),
	}
	presence.HostInfo.Host, _ = service.ipAssigner.AssignIP(ip, presence.Host)
	service.db.AddMix(presence)
}

func (service *service) AddMixProviderPresence(info models.MixProviderHostInfo, ip string) {
	presence := models.MixProviderPresence{
		MixProviderHostInfo: info,
		LastSeen:            timemock.Now().UnixNano(),
	}
	presence.HostInfo.Host, _ = service.ipAssigner.AssignIP(ip, presence.Host)
	service.db.AddMixProvider(presence)
}

func (service *service) Topology() models.Topology {
	return service.db.Topology()
}

type ipAssigner struct {
}

// IPAssigner compares the realIP (taken from the incoming request to the
// controller) and the self-reported presence IP (taken from the presence report
// data), and tries to report a reasonable IP. Much like the trouble with SUVs
// detailed by Paul Graham (http://www.paulgraham.com/hundred.html), this is a
// gross solution to a gross problem. In our case, the cause of hassle is that AWS
// boxes (a) don't allow applications hosted on them to determine what address
// they're binding to easily, because there are no "real" public IPs
// assigned, and (b) cause the application to explode if you attempt to bind
// to the public IP at all (private IPs do exist and can be bound to).
// So in some cases (like our testnet), the self-reported IP of the
// incoming presence report will likely be incorrect.
//
// If we could, we'd always read from the incoming request - but this has another
// problem: incoming requests don't tell us which port the remote node is
// listening on. So we need to combine self-reported and real IP in that case.
type IPAssigner interface {
	AssignIP(serverReportedIP string, selfReportedHost string) (string, error)
}

func (ipa *ipAssigner) AssignIP(serverReportedIP string, selfReportedHost string) (string, error) {
	var host string
	selfReportedIP, port, err := net.SplitHostPort(selfReportedHost)
	if err != nil {
		return "", err
	}
	if ipa.isLocal(selfReportedIP) {
		host = selfReportedHost // includes a port
	} else {
		host = net.JoinHostPort(serverReportedIP, port)
	}
	return host, nil
}

func (ipa *ipAssigner) isLocal(ip string) bool {
	return ip == "localhost" || net.ParseIP(ip).IsLoopback()
}
