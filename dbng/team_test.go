package dbng_test

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/concourse/atc"
	"github.com/concourse/atc/dbng"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Team", func() {
	var (
		otherTeam dbng.Team
	)

	BeforeEach(func() {
		otherTeam, err = teamFactory.CreateTeam("some-other-team")
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("SaveWorker", func() {
		var (
			team      dbng.Team
			otherTeam dbng.Team
			atcWorker atc.Worker
			err       error
		)

		BeforeEach(func() {
			postgresRunner.Truncate()
			team, err = teamFactory.CreateTeam("team")
			Expect(err).NotTo(HaveOccurred())

			otherTeam, err = teamFactory.CreateTeam("some-other-team")
			Expect(err).NotTo(HaveOccurred())
			atcWorker = atc.Worker{
				GardenAddr:       "some-garden-addr",
				BaggageclaimURL:  "some-bc-url",
				HTTPProxyURL:     "some-http-proxy-url",
				HTTPSProxyURL:    "some-https-proxy-url",
				NoProxy:          "some-no-proxy",
				ActiveContainers: 140,
				ResourceTypes: []atc.WorkerResourceType{
					{
						Type:    "some-resource-type",
						Image:   "some-image",
						Version: "some-version",
					},
					{
						Type:    "other-resource-type",
						Image:   "other-image",
						Version: "other-version",
					},
				},
				Platform:  "some-platform",
				Tags:      atc.Tags{"some", "tags"},
				Name:      "some-name",
				StartTime: 55,
			}
		})

		Context("the worker already exists", func() {
			Context("the worker is not in stalled state", func() {
				Context("the team_id of the new worker is the same", func() {
					BeforeEach(func() {
						_, err := team.SaveWorker(atcWorker, 5*time.Minute)
						Expect(err).NotTo(HaveOccurred())
					})
					It("overwrites all the data", func() {
						atcWorker.GardenAddr = "new-garden-addr"
						savedWorker, err := team.SaveWorker(atcWorker, 5*time.Minute)
						Expect(err).NotTo(HaveOccurred())
						Expect(savedWorker.Name()).To(Equal("some-name"))
						Expect(*savedWorker.GardenAddr()).To(Equal("new-garden-addr"))
						Expect(savedWorker.State()).To(Equal(dbng.WorkerStateRunning))
					})
				})
				Context("the team_id of the new worker is different", func() {
					BeforeEach(func() {
						_, err = otherTeam.SaveWorker(atcWorker, 5*time.Minute)
						Expect(err).NotTo(HaveOccurred())
					})
					It("errors", func() {
						_, err = team.SaveWorker(atcWorker, 5*time.Minute)
						Expect(err).To(HaveOccurred())
					})
				})
			})
		})
	})

	Describe("Workers", func() {
		var (
			team      dbng.Team
			otherTeam dbng.Team
			atcWorker atc.Worker
			err       error
		)

		BeforeEach(func() {
			postgresRunner.Truncate()
			team, err = teamFactory.CreateTeam("team")
			Expect(err).NotTo(HaveOccurred())

			otherTeam, err = teamFactory.CreateTeam("some-other-team")
			Expect(err).NotTo(HaveOccurred())
			atcWorker = atc.Worker{
				GardenAddr:       "some-garden-addr",
				BaggageclaimURL:  "some-bc-url",
				HTTPProxyURL:     "some-http-proxy-url",
				HTTPSProxyURL:    "some-https-proxy-url",
				NoProxy:          "some-no-proxy",
				ActiveContainers: 140,
				ResourceTypes: []atc.WorkerResourceType{
					{
						Type:    "some-resource-type",
						Image:   "some-image",
						Version: "some-version",
					},
					{
						Type:    "other-resource-type",
						Image:   "other-image",
						Version: "other-version",
					},
				},
				Platform:  "some-platform",
				Tags:      atc.Tags{"some", "tags"},
				Name:      "some-name",
				StartTime: 55,
			}
		})

		Context("when there are global workers and workers for the team", func() {
			BeforeEach(func() {
				_, err = team.SaveWorker(atcWorker, 0)
				Expect(err).NotTo(HaveOccurred())

				atcWorker.Name = "some-new-worker"
				atcWorker.GardenAddr = "some-other-garden-addr"
				atcWorker.BaggageclaimURL = "some-other-bc-url"
				_, err = workerFactory.SaveWorker(atcWorker, 0)
				Expect(err).NotTo(HaveOccurred())
			})

			It("finds them without error", func() {
				workers, err := team.Workers()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(workers)).To(Equal(2))

				Expect(workers[0].Name()).To(Equal("some-name"))
				Expect(*workers[0].GardenAddr()).To(Equal("some-garden-addr"))
				Expect(*workers[0].BaggageclaimURL()).To(Equal("some-bc-url"))

				Expect(workers[1].Name()).To(Equal("some-new-worker"))
				Expect(*workers[1].GardenAddr()).To(Equal("some-other-garden-addr"))
				Expect(*workers[1].BaggageclaimURL()).To(Equal("some-other-bc-url"))
			})
		})

		Context("when there are workers for another team", func() {
			BeforeEach(func() {
				atcWorker.Name = "some-other-team-worker"
				atcWorker.GardenAddr = "some-other-garden-addr"
				atcWorker.BaggageclaimURL = "some-other-bc-url"
				_, err = otherTeam.SaveWorker(atcWorker, 5*time.Minute)
				Expect(err).NotTo(HaveOccurred())
			})

			It("does not find the other team workers", func() {
				workers, err := team.Workers()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(workers)).To(Equal(0))
			})
		})

		Context("when there are no workers", func() {
			It("returns an error", func() {
				workers, err := workerFactory.Workers()
				Expect(err).NotTo(HaveOccurred())
				Expect(workers).To(BeEmpty())
			})
		})
	})

	Describe("FindContainerByHandle", func() {
		var createdContainer dbng.CreatedContainer

		BeforeEach(func() {
			build, err := defaultPipeline.CreateJobBuild("some-job")
			Expect(err).NotTo(HaveOccurred())

			creatingContainer, err := defaultTeam.CreateBuildContainer(defaultWorker.Name(), build.ID(), atc.PlanID("some-job"), dbng.ContainerMetadata{Type: "task", StepName: "some-task"})
			Expect(err).NotTo(HaveOccurred())

			createdContainer, err = creatingContainer.Created()
			Expect(err).NotTo(HaveOccurred())
		})

		Context("when worker is no longer in database", func() {
			BeforeEach(func() {
				err = defaultWorker.Delete()
				Expect(err).NotTo(HaveOccurred())
			})

			It("the container goes away from the db", func() {
				_, found, err := defaultTeam.FindContainerByHandle(createdContainer.Handle())
				Expect(err).NotTo(HaveOccurred())
				Expect(found).To(BeFalse())
			})
		})

		It("finds a container for the team", func() {
			container, found, err := defaultTeam.FindContainerByHandle(createdContainer.Handle())
			Expect(err).NotTo(HaveOccurred())
			Expect(found).To(BeTrue())
			Expect(container).ToNot(BeNil())
			Expect(container.Handle()).To(Equal(createdContainer.Handle()))
		})

		It("does not find container for another team", func() {
			_, found, err := otherTeam.FindContainerByHandle(createdContainer.Handle())
			Expect(err).NotTo(HaveOccurred())
			Expect(found).To(BeFalse())
		})
	})

	Describe("FindWorkerForResourceCheckContainer", func() {
		var resourceConfig *dbng.UsedResourceConfig

		BeforeEach(func() {
			resourceConfig, err = resourceConfigFactory.FindOrCreateResourceConfig(
				logger,
				dbng.ForResource{defaultResource.ID},
				"some-base-resource-type",
				atc.Source{"some": "source"},
				atc.VersionedResourceTypes{},
			)
			Expect(err).NotTo(HaveOccurred())
		})

		Context("when there is a creating container", func() {
			BeforeEach(func() {
				_, err := defaultTeam.CreateResourceCheckContainer(defaultWorker.Name(), resourceConfig, dbng.ContainerMetadata{Type: "check"})
				Expect(err).NotTo(HaveOccurred())
			})

			It("returns it", func() {
				worker, found, err := defaultTeam.FindWorkerForResourceCheckContainer(resourceConfig)
				Expect(err).NotTo(HaveOccurred())
				Expect(found).To(BeTrue())
				Expect(worker).NotTo(BeNil())
				Expect(worker.Name()).To(Equal(defaultWorker.Name()))
			})

			It("does not find container for another team", func() {
				worker, found, err := otherTeam.FindWorkerForResourceCheckContainer(resourceConfig)
				Expect(err).NotTo(HaveOccurred())
				Expect(found).To(BeFalse())
				Expect(worker).To(BeNil())
			})
		})

		Context("when there is a created container", func() {
			var originalCreatedContainer dbng.CreatedContainer

			BeforeEach(func() {
				creatingContainer, err := defaultTeam.CreateResourceCheckContainer(defaultWorker.Name(), resourceConfig, dbng.ContainerMetadata{Type: "check"})
				Expect(err).NotTo(HaveOccurred())
				originalCreatedContainer, err = creatingContainer.Created()
				Expect(err).NotTo(HaveOccurred())
			})

			It("returns it", func() {
				worker, found, err := defaultTeam.FindWorkerForResourceCheckContainer(resourceConfig)
				Expect(err).NotTo(HaveOccurred())
				Expect(found).To(BeTrue())
				Expect(worker).NotTo(BeNil())
				Expect(worker.Name()).To(Equal(defaultWorker.Name()))
			})

			It("does not find container for another team", func() {
				worker, found, err := otherTeam.FindWorkerForResourceCheckContainer(resourceConfig)
				Expect(err).NotTo(HaveOccurred())
				Expect(found).To(BeFalse())
				Expect(worker).To(BeNil())
			})

			Context("when container is expired", func() {
				BeforeEach(func() {
					_, err := psql.Update("containers").
						Set("best_if_used_by", sq.Expr("NOW() - '1 second'::INTERVAL")).
						Where(sq.Eq{"id": originalCreatedContainer.ID()}).
						RunWith(dbConn).Exec()
					Expect(err).NotTo(HaveOccurred())
				})

				It("does not find it", func() {
					worker, found, err := defaultTeam.FindWorkerForResourceCheckContainer(resourceConfig)
					Expect(err).NotTo(HaveOccurred())
					Expect(found).To(BeFalse())
					Expect(worker).To(BeNil())
				})
			})
		})

		Context("when there is no container", func() {
			It("returns nil", func() {
				worker, found, err := defaultTeam.FindWorkerForResourceCheckContainer(resourceConfig)
				Expect(err).NotTo(HaveOccurred())
				Expect(found).To(BeFalse())
				Expect(worker).To(BeNil())
			})
		})
	})

	Describe("FindResourceCheckContainerOnWorker", func() {
		var resourceConfig *dbng.UsedResourceConfig

		BeforeEach(func() {
			resourceConfig, err = resourceConfigFactory.FindOrCreateResourceConfig(
				logger,
				dbng.ForResource{defaultResource.ID},
				"some-base-resource-type",
				atc.Source{"some": "source"},
				atc.VersionedResourceTypes{},
			)
			Expect(err).NotTo(HaveOccurred())
		})

		Context("when there is a creating container", func() {
			BeforeEach(func() {
				_, err := defaultTeam.CreateResourceCheckContainer(defaultWorker.Name(), resourceConfig, dbng.ContainerMetadata{Type: "check"})
				Expect(err).NotTo(HaveOccurred())
			})

			It("returns it", func() {
				creatingContainer, createdContainer, err := defaultTeam.FindResourceCheckContainerOnWorker(defaultWorker.Name(), resourceConfig)
				Expect(err).NotTo(HaveOccurred())
				Expect(createdContainer).To(BeNil())
				Expect(creatingContainer).NotTo(BeNil())
			})

			It("does not find container for another team", func() {
				creatingContainer, createdContainer, err := otherTeam.FindResourceCheckContainerOnWorker(defaultWorker.Name(), resourceConfig)
				Expect(err).NotTo(HaveOccurred())
				Expect(creatingContainer).To(BeNil())
				Expect(createdContainer).To(BeNil())
			})
		})

		Context("when there is a created container", func() {
			var originalCreatedContainer dbng.CreatedContainer

			BeforeEach(func() {
				creatingContainer, err := defaultTeam.CreateResourceCheckContainer(defaultWorker.Name(), resourceConfig, dbng.ContainerMetadata{Type: "check"})
				Expect(err).NotTo(HaveOccurred())
				originalCreatedContainer, err = creatingContainer.Created()
				Expect(err).NotTo(HaveOccurred())
			})

			It("returns it", func() {
				creatingContainer, createdContainer, err := defaultTeam.FindResourceCheckContainerOnWorker(defaultWorker.Name(), resourceConfig)
				Expect(err).NotTo(HaveOccurred())
				Expect(createdContainer).NotTo(BeNil())
				Expect(creatingContainer).To(BeNil())
			})

			It("does not find container for another team", func() {
				creatingContainer, createdContainer, err := otherTeam.FindResourceCheckContainerOnWorker(defaultWorker.Name(), resourceConfig)
				Expect(err).NotTo(HaveOccurred())
				Expect(creatingContainer).To(BeNil())
				Expect(createdContainer).To(BeNil())
			})

			Context("when container is expired", func() {
				BeforeEach(func() {
					_, err := psql.Update("containers").
						Set("best_if_used_by", sq.Expr("NOW() - '1 second'::INTERVAL")).
						Where(sq.Eq{"id": originalCreatedContainer.ID()}).
						RunWith(dbConn).Exec()
					Expect(err).NotTo(HaveOccurred())
				})

				It("does not find it", func() {
					creatingContainer, createdContainer, err := defaultTeam.FindResourceCheckContainerOnWorker(defaultWorker.Name(), resourceConfig)
					Expect(err).NotTo(HaveOccurred())
					Expect(creatingContainer).To(BeNil())
					Expect(createdContainer).To(BeNil())
				})
			})
		})

		Context("when there is no container", func() {
			It("returns nil", func() {
				creatingContainer, createdContainer, err := defaultTeam.FindResourceCheckContainerOnWorker(defaultWorker.Name(), resourceConfig)
				Expect(err).NotTo(HaveOccurred())
				Expect(createdContainer).To(BeNil())
				Expect(creatingContainer).To(BeNil())
			})
		})
	})

	Describe("FindWorkerForBuildContainer", func() {
		var containerMetadata dbng.ContainerMetadata

		BeforeEach(func() {
			containerMetadata = dbng.ContainerMetadata{
				Type:     "task",
				StepName: "some-task",
			}
		})

		Context("when there is a creating container", func() {
			BeforeEach(func() {
				_, err := defaultTeam.CreateBuildContainer(defaultWorker.Name(), defaultBuild.ID(), "some-plan", containerMetadata)
				Expect(err).ToNot(HaveOccurred())
			})

			It("returns it", func() {
				worker, found, err := defaultTeam.FindWorkerForBuildContainer(defaultBuild.ID(), "some-plan")
				Expect(err).NotTo(HaveOccurred())
				Expect(found).To(BeTrue())
				Expect(worker).NotTo(BeNil())
				Expect(worker.Name()).To(Equal(defaultWorker.Name()))
			})

			It("does not find container for another team", func() {
				worker, found, err := otherTeam.FindWorkerForBuildContainer(defaultBuild.ID(), "some-plan")
				Expect(err).NotTo(HaveOccurred())
				Expect(found).To(BeFalse())
				Expect(worker).To(BeNil())
			})
		})

		Context("when there is a created container", func() {
			BeforeEach(func() {
				creatingContainer, err := defaultTeam.CreateBuildContainer(defaultWorker.Name(), defaultBuild.ID(), "some-plan", containerMetadata)
				Expect(err).NotTo(HaveOccurred())
				_, err = creatingContainer.Created()
				Expect(err).NotTo(HaveOccurred())
			})

			It("returns it", func() {
				worker, found, err := defaultTeam.FindWorkerForBuildContainer(defaultBuild.ID(), "some-plan")
				Expect(err).NotTo(HaveOccurred())
				Expect(found).To(BeTrue())
				Expect(worker).NotTo(BeNil())
				Expect(worker.Name()).To(Equal(defaultWorker.Name()))
			})

			It("does not find container for another team", func() {
				worker, found, err := otherTeam.FindWorkerForBuildContainer(defaultBuild.ID(), "some-plan")
				Expect(err).NotTo(HaveOccurred())
				Expect(found).To(BeFalse())
				Expect(worker).To(BeNil())
			})
		})

		Context("when there is no container", func() {
			It("returns nil", func() {
				worker, found, err := defaultTeam.FindWorkerForBuildContainer(defaultBuild.ID(), "some-plan")
				Expect(err).NotTo(HaveOccurred())
				Expect(found).To(BeFalse())
				Expect(worker).To(BeNil())
			})
		})
	})

	Describe("FindBuildContainerOnWorker", func() {
		var containerMetadata dbng.ContainerMetadata

		BeforeEach(func() {
			containerMetadata = dbng.ContainerMetadata{
				Type:     "task",
				StepName: "some-task",
			}
		})

		Context("when there is a creating container", func() {
			BeforeEach(func() {
				_, err := defaultTeam.CreateBuildContainer(defaultWorker.Name(), defaultBuild.ID(), "some-plan", containerMetadata)
				Expect(err).ToNot(HaveOccurred())
			})

			It("returns it", func() {
				creatingContainer, createdContainer, err := defaultTeam.FindBuildContainerOnWorker(defaultWorker.Name(), defaultBuild.ID(), "some-plan")
				Expect(err).NotTo(HaveOccurred())
				Expect(createdContainer).To(BeNil())
				Expect(creatingContainer).NotTo(BeNil())
			})

			It("does not find container for another team", func() {
				creatingContainer, createdContainer, err := otherTeam.FindBuildContainerOnWorker(defaultWorker.Name(), defaultBuild.ID(), "some-plan")
				Expect(err).NotTo(HaveOccurred())
				Expect(creatingContainer).To(BeNil())
				Expect(createdContainer).To(BeNil())
			})
		})

		Context("when there is a created container", func() {
			BeforeEach(func() {
				creatingContainer, err := defaultTeam.CreateBuildContainer(defaultWorker.Name(), defaultBuild.ID(), "some-plan", containerMetadata)
				Expect(err).NotTo(HaveOccurred())
				_, err = creatingContainer.Created()
				Expect(err).NotTo(HaveOccurred())
			})

			It("returns it", func() {
				creatingContainer, createdContainer, err := defaultTeam.FindBuildContainerOnWorker(defaultWorker.Name(), defaultBuild.ID(), "some-plan")
				Expect(err).NotTo(HaveOccurred())
				Expect(createdContainer).NotTo(BeNil())
				Expect(creatingContainer).To(BeNil())
			})

			It("does not find container for another team", func() {
				creatingContainer, createdContainer, err := otherTeam.FindBuildContainerOnWorker(defaultWorker.Name(), defaultBuild.ID(), "some-plan")
				Expect(err).NotTo(HaveOccurred())
				Expect(creatingContainer).To(BeNil())
				Expect(createdContainer).To(BeNil())
			})
		})

		Context("when there is no container", func() {
			It("returns nil", func() {
				creatingContainer, createdContainer, err := defaultTeam.FindBuildContainerOnWorker(defaultWorker.Name(), defaultBuild.ID(), "some-plan")
				Expect(err).NotTo(HaveOccurred())
				Expect(createdContainer).To(BeNil())
				Expect(creatingContainer).To(BeNil())
			})
		})
	})
})
