package mixmining

import (
	"errors"

	"github.com/BorisBorshevsky/timemock"
	"github.com/nymtech/nym-directory/mixmining/mocks"
	"github.com/nymtech/nym-directory/models"
	. "github.com/onsi/ginkgo"
	"gotest.tools/assert"
)

// Some fixtures data to dry up tests a bit

// A slice of IPv4 mix statuses with 2 ups and 1 down during the past day
func twoUpOneDown() []models.PersistedMixStatus {
	db := []models.PersistedMixStatus{}
	var status = persistedStatus()

	status.PubKey = "key1"
	status.IPVersion = "4"
	status.Up = true

	status.Timestamp = minutesAgo(5)
	db = append(db, status)

	status.Timestamp = minutesAgo(10)
	db = append(db, status)

	status.Timestamp = minutesAgo(15)
	status.Up = false
	db = append(db, status)

	return db
}

func persistedStatus() models.PersistedMixStatus {
	mixStatus := status()
	persisted := models.PersistedMixStatus{
		MixStatus: mixStatus,
		Timestamp: Now(),
	}
	return persisted
}

func status() models.MixStatus {
	return models.MixStatus{
		PubKey:    "key1",
		IPVersion: "4",
		Up:        false,
	}
}

// A version of now with a frozen shared clock so we can have determinate time-based tests
func Now() int64 {
	now := timemock.Now()
	timemock.Freeze(now) //time is frozen
	nanos := now.UnixNano()
	return nanos
}

var _ = Describe("mixmining.Service", func() {
	var mockDb mocks.IDb
	var status1 models.MixStatus
	var status2 models.MixStatus
	var persisted1 models.PersistedMixStatus
	var persisted2 models.PersistedMixStatus
	var last5minutes []models.PersistedMixStatus

	var serv Service

	status1 = models.MixStatus{
		PubKey:    "key1",
		IPVersion: "4",
		Up:        true,
	}

	persisted1 = models.PersistedMixStatus{
		MixStatus: status1,
		Timestamp: Now(),
	}

	status2 = models.MixStatus{
		PubKey:    "key2",
		IPVersion: "6",
		Up:        true,
	}

	persisted2 = models.PersistedMixStatus{
		MixStatus: status2,
		Timestamp: Now(),
	}

	downer := persisted1
	downer.MixStatus.Up = false

	persistedList := []models.PersistedMixStatus{persisted1, persisted2}
	emptyList := []models.PersistedMixStatus{}

	BeforeEach(func() {
		mockDb = *new(mocks.IDb)
		serv = *NewService(&mockDb)
	})

	Describe("Adding a mix status and creating a new summary report for a node", func() {
		Context("when no statuses have yet been saved", func() {
			It("should add a PersistedMixStatus to the db and save the new report", func() {

				mockDb.On("Add", persisted1)

				serv.CreateMixStatus(status1)
				mockDb.AssertCalled(GinkgoT(), "Add", persisted1)
			})
		})
	})
	Describe("Listing mix statuses", func() {
		Context("when receiving a list request", func() {
			It("should call to the Db", func() {
				mockDb.On("List", persisted1.PubKey, 1000).Return(persistedList)

				result := serv.List(persisted1.PubKey)

				mockDb.AssertCalled(GinkgoT(), "List", persisted1.PubKey, 1000)
				assert.Equal(GinkgoT(), persistedList[0].MixStatus.PubKey, result[0].MixStatus.PubKey)
				assert.Equal(GinkgoT(), persistedList[1].MixStatus.PubKey, result[1].MixStatus.PubKey)
			})
		})
	})

	Describe("Calculating uptime", func() {
		Context("when no statuses exist yet", func() {
			It("should return 0", func() {
				mockDb.On("ListDateRange", "key1", "4", Now(), daysAgo(30)).Return(emptyList)

				uptime := serv.CalculateUptime(persisted1.PubKey, persisted1.IPVersion, daysAgo(30))
				assert.Equal(GinkgoT(), 0, uptime)
			})

		})
		Context("when 2 ups and 1 down exist in the given time period", func() {
			It("should return 66", func() {
				mockDb.On("ListDateRange", "key1", "4", Now(), daysAgo(1)).Return(twoUpOneDown())

				uptime := serv.CalculateUptime("key1", "4", daysAgo(1))
				expected := 66 // percent
				assert.Equal(GinkgoT(), expected, uptime)
			})
		})
	})

	Describe("Saving a mix status report", func() {
		BeforeEach(func() {
			mockDb = *new(mocks.IDb)
			serv = *NewService(&mockDb)

			last5minutes = persistedList
			// lastHour := append(last5minutes, persistedList...)
			// lastDay := append(lastHour, persistedList...)
			// lastWeek := append(lastDay, persistedList...)
			// lastMonth := append(lastWeek, persistedList...)
			// so in the last month we have 10 persisted statuses, all up
		})
		Context("when no statuses exist yet", func() {
			BeforeEach(func() {
				expectedSave := models.MixStatusReport{
					PubKey: downer.PubKey,
				}
				mockDb.On("SaveMixStatusReport", expectedSave)
				mockDb.On("ListDateRange", downer.PubKey, downer.IPVersion, now(), minutesAgo(5)).Return(emptyList)
				mockDb.On("ListDateRange", downer.PubKey, downer.IPVersion, now(), minutesAgo(60)).Return(emptyList)
				mockDb.On("ListDateRange", downer.PubKey, downer.IPVersion, now(), daysAgo(1)).Return(emptyList)
				mockDb.On("ListDateRange", downer.PubKey, downer.IPVersion, now(), daysAgo(30)).Return(emptyList)
				mockDb.On("LoadReport", downer.PubKey).Return(errors.New("Report does not exist"))

			})
			It("should save the initial report, all statuses will be up (or down)", func() {
				result := serv.SaveStatusReport(downer)

				assert.Equal(GinkgoT(), 0, result.IPV4Status.Last5Minutes)
				assert.Equal(GinkgoT(), 0, result.IPV4Status.LastHour)
				assert.Equal(GinkgoT(), 0, result.IPV4Status.LastDay)
				assert.Equal(GinkgoT(), 0, result.IPV4Status.LastWeek)
				assert.Equal(GinkgoT(), 0, result.IPV4Status.LastMonth)
				mockDb.AssertExpectations(GinkgoT())
			})
		})

		Context("when 2 up statuses exist for the last 5 minutes already", func() {
			BeforeEach(func() {
				expectedSave := models.MixStatusReport{
					PubKey: downer.PubKey,
				}

				mockDb.On("SaveMixStatusReport", expectedSave)
				mockDb.On("ListDateRange", downer.PubKey, downer.IPVersion, now(), minutesAgo(5)).Return(last5minutes)
				mockDb.On("ListDateRange", downer.PubKey, downer.IPVersion, now(), minutesAgo(60)).Return(emptyList)
				mockDb.On("ListDateRange", downer.PubKey, downer.IPVersion, now(), daysAgo(1)).Return(emptyList)
				mockDb.On("ListDateRange", downer.PubKey, downer.IPVersion, now(), daysAgo(30)).Return(emptyList)
			})
			It("should save the report", func() {
				initialState := models.MixStatusReport{
					PubKey: downer.PubKey,
					IPV4Status: models.UptimeReport{
						IPVersion:    "4",
						MostRecent:   true,
						Last5Minutes: 100,
					},
					IPV6Status: models.UptimeReport{},
				}

				expectedAfterUpdate := models.MixStatusReport{
					PubKey: downer.PubKey,
					IPV4Status: models.UptimeReport{
						IPVersion:    "4",
						MostRecent:   downer.Up,
						Last5Minutes: 66,
					},
					IPV6Status: models.UptimeReport{},
				}
				mockDb.On("LoadReport", downer.PubKey).Return(initialState)
				mockDb.On("GetStatusReport", persisted1.PubKey).Return(initialState)
				mockDb.On("SaveMixStatusReport", expectedAfterUpdate)
				result := serv.SaveStatusReport(downer)
				assert.Equal(GinkgoT(), expectedAfterUpdate, result)
			})
		})

		Context("when all time periods exist", func() {

		})
	})
})
