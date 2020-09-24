package mixmining

import (
	"time"

	"github.com/BorisBorshevsky/timemock"
	"github.com/nymtech/nym-directory/mixmining/mocks"
	"github.com/nymtech/nym-directory/models"
	. "github.com/onsi/ginkgo"
	"gotest.tools/assert"
)

var _ = Describe("mixmining.Service", func() {
	var mockDb mocks.IDb
	var status1 models.MixStatus
	var status2 models.MixStatus
	var persisted1 models.PersistedMixStatus
	var persisted2 models.PersistedMixStatus

	var serv Service
	var now = time.Now()
	timemock.Freeze(now)
	var frozenNow = timemock.Now().UnixNano()
	// set up fixtures
	status1 = models.MixStatus{
		PubKey:    "key1",
		IPVersion: "4",
		Up:        true,
	}

	persisted1 = models.PersistedMixStatus{
		MixStatus: status1,
		Timestamp: frozenNow,
	}

	status2 = models.MixStatus{
		PubKey:    "key2",
		IPVersion: "6",
		Up:        true,
	}

	persisted2 = models.PersistedMixStatus{
		MixStatus: status2,
		Timestamp: frozenNow,
	}

	persistedList := []models.PersistedMixStatus{persisted1, persisted2}

	Describe("Adding a mix status and creating a new summary report for a node", func() {
		Context("when no statuses have yet been saved", func() {
			It("should add a PersistedMixStatus to the db and save the new report", func() {
				mockDb = *new(mocks.IDb)
				serv = *NewService(&mockDb)
				mockDb.On("Add", persisted1)
				expectedStatusReport := models.MixStatusReport{
					PubKey: persisted1.PubKey,
				}
				mockDb.On("SaveMixStatusReport", expectedStatusReport)

				serv.CreateMixStatus(status1)
				mockDb.AssertCalled(GinkgoT(), "Add", persisted1)
				mockDb.AssertCalled(GinkgoT(), "SaveMixStatusReport", expectedStatusReport)
			})
		})
	})
	Describe("Listing mix statuses", func() {
		Context("when receiving a list request", func() {
			It("should call to the Db", func() {
				mockDb = *new(mocks.IDb)
				serv = *NewService(&mockDb)
				mockDb.On("List", persisted1.PubKey, 1000).Return(persistedList)

				result := serv.List(persisted1.PubKey)

				mockDb.AssertCalled(GinkgoT(), "List", persisted1.PubKey, 1000)
				assert.Equal(GinkgoT(), persistedList[0].MixStatus.PubKey, result[0].MixStatus.PubKey)
				assert.Equal(GinkgoT(), persistedList[1].MixStatus.PubKey, result[1].MixStatus.PubKey)
			})
		})
	})

	var _ = Describe("Calculating uptime", func() {
		Context("when no statuses exist yet", func() {
			It("should return 0", func() {
				mockDb = *new(mocks.IDb)
				now := time.Now()
				thirtyDaysAgo := now.Add(time.Duration(-30) * time.Hour * 24)
				mockDb.On("ListDateRange", now, thirtyDaysAgo).Return([]models.PersistedMixStatus{})
				serv = *NewService(&mockDb)

				uptime := serv.CalculateUptime(persisted1.PubKey, persisted1.IPVersion, thirtyDaysAgo.UnixNano())
				assert.Equal(GinkgoT(), 0, uptime)
			})

			Context("for IPv4", func() {
				Context("when 2 ups and 1 down exist in the past day", func() {
					Context("and 3 addition ups and 1 down exist in the past month", func() {
						Context("with 1 IPv6 in the past day and an additional IPv6 in the past month", func() {
							Context("getting range in the past day", func() {
								It("should return 0", func() {
									mockDb = *new(mocks.IDb)
									now := time.Now()
									list := bigDbState()
									mockDb.On("ListDateRange", now, daysAgo(30)).Return(list)
									serv = *NewService(&mockDb)

									uptime := serv.CalculateUptime("key1", "4", daysAgo(1))
									expected := 2 / 3
									assert.Equal(GinkgoT(), expected, uptime)
								})
							})
						})
					})
				})

			})
		})
	})
})

func bigDbState() []models.PersistedMixStatus {
	var db []models.PersistedMixStatus
	var status = newPersistedStatus()

	// IPv4
	// 2 ups and 1 down in last day
	status.IPVersion = "4"
	status.Up = true
	status.Timestamp = minutesAgo(5)
	db = append(db, status)

	status.Timestamp = minutesAgo(10)
	db = append(db, status)

	status.Timestamp = minutesAgo(15)
	status.Up = false
	db = append(db, status)

	// and 3 addition ups and 1 down exist in the past month
	status.Up = true
	status.Timestamp = daysAgo(5)
	db = append(db, status)

	status.Timestamp = daysAgo(10)
	db = append(db, status)

	status.Timestamp = daysAgo(15)
	db = append(db, status)

	status.Up = false
	status.Timestamp = daysAgo(20)
	db = append(db, status)

	// with 1 IPv6 in the past day and an additional IPv6 in the past month
	status.IPVersion = "6"
	status.Timestamp = minutesAgo(10)
	status.Up = true
	db = append(db, status)

	status.IPVersion = "6"
	status.Timestamp = daysAgo(15)
	status.Up = true
	db = append(db, status)

	return db
}

func newPersistedStatus() models.PersistedMixStatus {

	mixStatus := newStatus()
	persisted := models.PersistedMixStatus{
		MixStatus: mixStatus,
		Timestamp: frozenNow(),
	}
	return persisted
}

func newStatus() models.MixStatus {
	return models.MixStatus{
		PubKey:    "key1",
		IPVersion: "4",
		Up:        false,
	}
}

func frozenNow() int64 {
	now := timemock.Now()
	timemock.Freeze(now) //time is frozen
	return now.UnixNano()
}
