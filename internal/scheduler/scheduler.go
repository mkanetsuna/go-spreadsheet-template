package scheduler

import (
	"context"
	"fmt"
	"log"

	cloudscheduler "cloud.google.com/go/scheduler/apiv1"
	schedulerpb "google.golang.org/genproto/googleapis/cloud/scheduler/v1"
)

func SetupScheduler(projectID, location, serviceURL string) error {
	ctx := context.Background()
	client, err := cloudscheduler.NewCloudSchedulerClient(ctx)
	if err != nil {
		return fmt.Errorf("NewCloudSchedulerClient: %v", err)
	}
	defer client.Close()

	// Setup hourly job
	if err := setupJob(ctx, client, projectID, location, serviceURL, "hourly-job", "0 * * * *"); err != nil {
		return err
	}

	// Setup daily job
	if err := setupJob(ctx, client, projectID, location, serviceURL, "daily-job", "0 0 * * *"); err != nil {
		return err
	}

	return nil
}

func setupJob(ctx context.Context, client *cloudscheduler.CloudSchedulerClient, projectID, location, serviceURL, jobName, schedule string) error {
    parent := fmt.Sprintf("projects/%s/locations/%s", projectID, location)
    jobID := fmt.Sprintf("%s/jobs/%s", parent, jobName)

    job := &schedulerpb.Job{
        Name: jobID,
        Target: &schedulerpb.Job_HttpTarget{
            HttpTarget: &schedulerpb.HttpTarget{
                Uri:        fmt.Sprintf("%s/scheduled-task", serviceURL),
                HttpMethod: schedulerpb.HttpMethod_POST,
            },
        },
        Schedule: schedule,
        TimeZone: "Etc/UTC",
    }

    req := &schedulerpb.CreateJobRequest{
        Parent: parent,
        Job:    job,
    }

    _, err := client.CreateJob(ctx, req)
    if err != nil {
        // If the job already exists, update it
        updateReq := &schedulerpb.UpdateJobRequest{
            Job: job,
        }
        _, err = client.UpdateJob(ctx, updateReq)
        if err != nil {
            return fmt.Errorf("Failed to update job: %v", err)
        }
        log.Printf("Updated existing job: %s", jobID)
    } else {
        log.Printf("Created new job: %s", jobID)
    }

    return nil
}