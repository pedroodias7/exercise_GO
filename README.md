# exercise_GO

# Go + Kubernetes Cloud-Native DevOps Project

This project is designed to build a cloud-native DevOps system using Go and Kubernetes. It is divided into several phases, each focusing on core concepts such as Go programming, Docker, Kubernetes, and Google Cloud Platform (GCP) services.

## Phase 1: Basic Go Language Exercises

The goal here is to build a solid foundation in Go by focusing on syntax, data structures, and concurrency.

### 1.1 Hello World and Go Modules Setup
**Objective**: Write a simple Hello, World! program.  
**Key Concepts**: `fmt`, creating and managing Go modules.

**Task**:
- Create a Go module (`go mod init`) and write a simple Go program that prints "Hello, World!".
- Commit the code to GitHub.

### 1.2 Basic HTTP Server
**Objective**: Build a simple HTTP server that returns "Hello, Go!" when accessed.  
**Key Concepts**: `net/http` package, building HTTP handlers.

**Task**:
- Create an HTTP server in Go that listens on port 8080 and responds with "Hello, Go!".
- Commit the code to GitHub.

### 1.3 REST API with CRUD Operations
**Objective**: Create a REST API for managing a list of tasks.  
**Key Concepts**: Structs, slices, maps, JSON encoding/decoding, HTTP methods (GET, POST, PUT, DELETE).

**Task**:
- Build a REST API where users can:
  - **GET** a list of tasks.
  - **POST** a new task.
  - **PUT** to update an existing task.
  - **DELETE** a task.
- Store tasks in a simple in-memory slice.
- Commit the code to GitHub.

---

## Phase 2: Go + Kubernetes: Simple Applications

Now, let’s integrate Go applications into Kubernetes and build some microservices.

### 2.1 Dockerize Your Go App
**Objective**: Containerize your Go REST API from 1.3.  
**Key Concepts**: Dockerfile, multi-stage builds, image optimization.

**Task**:
- Write a Dockerfile to containerize your REST API.
- Use a multi-stage build to reduce image size.
- Push the Docker image to Docker Hub.
- Commit the Dockerfile and image link to GitHub.

### 2.2 Deploying the Go App on Kubernetes
**Objective**: Deploy the containerized Go REST API to a Kubernetes cluster.  
**Key Concepts**: Kubernetes manifests (Deployment, Service), `kubectl`, and GCP’s GKE (Google Kubernetes Engine).

**Task**:
- Create Kubernetes manifests for the REST API (Deployment + Service).
- Deploy the app to your GKE cluster.
- Expose the service externally via a LoadBalancer.
- Commit the Kubernetes manifests and deployment instructions to GitHub.

### 2.3 Health Checks and Liveness/Readiness Probes
**Objective**: Implement liveness and readiness probes in your Go app and configure them in Kubernetes.  
**Key Concepts**: Kubernetes probes, HTTP health checks in Go.

**Task**:
- Add a `/healthz` endpoint in your Go API that returns 200 OK.
- Configure Kubernetes liveness and readiness probes for your app.
- Commit updated manifests and code to GitHub.

---

## Phase 3: Go + Kubernetes + GCP

In this phase, we'll focus on cloud-native practices, GCP integration, and microservices.

### 3.1 Implementing Go Logging and Monitoring
**Objective**: Set up structured logging in Go and push logs to GCP’s Stackdriver.  
**Key Concepts**: Go logging packages (`log`, `logrus`), GCP’s Stackdriver.

**Task**:
- Use `logrus` (or a similar logging library) for structured logging in your Go API.
- Export logs to GCP Stackdriver for monitoring.
- Commit the updated code with logging setup to GitHub.

### 3.2 Go + GCP Pub/Sub Integration
**Objective**: Integrate GCP Pub/Sub in your Go service to send and receive messages.  
**Key Concepts**: GCP Pub/Sub, Go client for Pub/Sub, async processing.

**Task**:
- Create a new endpoint `/publish` that publishes a message to a Pub/Sub topic.
- Create a background Go routine that listens to the Pub/Sub subscription.
- Commit the code and documentation for deploying this integration to GitHub.

### 3.3 Microservices with Go and Kubernetes
**Objective**: Break your API into two microservices: TaskService and NotificationService.  
**Key Concepts**: Microservices architecture, service-to-service communication, Kubernetes Services.

**Task**:
- Split the original API into two services:
  - **TaskService**: Manages task creation, updates, etc.
  - **NotificationService**: Listens to Pub/Sub and sends email (mock email service).
- Deploy both services on Kubernetes and ensure they communicate through Kubernetes services.
- Commit the code and manifests for both services to GitHub.

---

## Phase 4: Advanced DevOps with Go

Focus on building advanced cloud-native tools with Go and Kubernetes.

### 4.1 Build a Kubernetes Controller in Go
**Objective**: Write a custom Kubernetes controller in Go.  
**Key Concepts**: Kubernetes Custom Resource Definitions (CRDs), Go `client-go` library.

**Task**:
- Create a custom resource (e.g., TaskScheduler) that automatically schedules tasks.
- Write a Go-based Kubernetes controller that watches the custom resource and manages tasks.
- Deploy the controller in your Kubernetes cluster.
- Commit the controller code and CRD manifests to GitHub.

### 4.2 CI/CD Pipeline for Go and Kubernetes
**Objective**: Automate testing and deployment of your Go microservices to Kubernetes.  
**Key Concepts**: CI/CD pipelines, GitHub Actions, Google Cloud Build.

**Task**:
- Create a GitHub Actions pipeline that:
  - Runs unit tests for your Go code.
  - Builds and pushes the Docker image.
  - Deploys to GKE using `kubectl`.
- Commit the pipeline configuration to GitHub.

---

## Phase 5: Full Cloud-Native Project

In the final phase, we’ll create a full-fledged cloud-native project that integrates everything.

### 5.1 Build a Cloud-Native DevOps Dashboard
**Objective**: Create a simple cloud-native DevOps dashboard in Go.  
**Key Concepts**: Go web frameworks (e.g., Gin or Echo), GCP APIs (for metrics), Kubernetes APIs, CI/CD integration.

**Task**:
- Build a Go-based dashboard that shows:
  - Status of your Kubernetes clusters.
  - Logs and metrics from GCP services.
  - CI/CD pipeline status (integrating GitHub Actions).
- Deploy this dashboard on Kubernetes.
- Commit the full project to GitHub.

---

## Integrating Helm and Terraform

### Objective: Use Helm and Terraform to manage your Go apps and Kubernetes resources.
**Key Concepts**: Helm charts, Terraform GKE provider, Kubernetes resource management.

**Task**:
- Write a Helm chart for one of your Go microservices.
- Use Terraform to provision GKE clusters and deploy your Go apps.
- Commit both Helm and Terraform configurations to GitHub.
