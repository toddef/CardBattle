Below is a comprehensive blueprint that breaks down the Card Battler project into increasingly smaller steps. After the iterative breakdown, you’ll find a series of prompt sections in Markdown, each enclosed in triple backticks so you can copy-paste them for your code-generation LLM of choice. Each prompt references the previous steps, follows test-driven best practices, and ensures that no code is left hanging or unused.

1. High-Level Blueprint
1.1 Repository & Project Setup
Monorepo Structure

Create a top-level repository.
Add folders for:
backend (containing microservices, each in its own subfolder)
frontend (Remix-based)
infrastructure (Kubernetes configs, Terraform, Helm chart, etc.)
Basic CI/CD Pipeline

Initialize GitHub Actions or another CI for building, linting, and basic testing.
Prepare Dockerfiles for each microservice and the frontend.
Scaffold Services

Auth Service skeleton.
User Service skeleton.
Deck Service skeleton.
Card Service skeleton.
Lobby Service skeleton.
Game Service skeleton.
Replay Service skeleton.
Frontend Scaffolding

Initialize a Remix app with basic routes (Login, Main Menu, Not Found).
Infrastructure

Set up Terraform scripts for AWS EKS and other resources.
Write a Helm chart that deploys your services to Kubernetes.
Configure ArgoCD for continuous delivery.
1.2 Major Features by Layer
Backend Microservices
Auth Service: Google OAuth, JWT handling.
User Service: Manage profiles, Google avatar, custom username.
Deck Service: Create/edit decks, enforce point limits, store deck data.
Card Service: Admin-manageable card repository.
Lobby Service: Create/join/leave lobbies, track players.
Game Service: Real-time gameplay, round resolution, scoring, game lifecycle.
Replay Service: Store event logs for replays.
Frontend
Remix Loaders/Actions: Server-side data fetching.
WebSocket Integration: Real-time updates in gameplay/lobbies.
UI Pages:
Login
Main Menu
Deck Builder
Lobby Browser
Game UI (real-time)
Profile/Stats
Admin Dashboard
Infrastructure
Kubernetes: Microservices as individual deployments.
PostgreSQL: For relational data (users, decks, games, stats).
MongoDB: For card storage and replays.
Redis: Caching.
NATS or another event bus for cross-service communication.
2. Iterative Breakdown
We’ll now break these tasks into smaller chunks. Each iteration ensures we have testable, incremental progress.

2.1 Initial Skeletons & Testing
Create Monorepo & CI

Initialize Git with the proper folder structure.
Add GitHub Actions config for linting and basic test placeholders.
Backend Microservice Skeletons

Write a minimal main.go (or cmd/main.go) for each microservice.
Create a shared go.mod or one per microservice (depending on your preference).
Define empty REST endpoints or gRPC stubs.
Frontend Skeleton

npx create-remix@latest or similar.
Add placeholder routes /login, /menu, etc.
Infrastructure Scaffolding

Write minimal Terraform to stand up an EKS cluster.
Write a Helm chart that just deploys a placeholder “Hello World” service.
Set up ArgoCD to watch the Git repo.
Basic Tests

Add minimal test files in each microservice (e.g., auth_service_test.go with a single placeholder).
Add a basic end-to-end test script that checks if “Hello World” is running on the cluster.
2.2 Authentication & User Flow
Implement Auth Service

Google OAuth flow.
JWT token generation and validation.
Store user ID tokens in an HTTP-only cookie.
Basic test coverage (unit + integration).
Implement User Service

CRUD for user profiles (username, avatar).
Basic test coverage (unit + integration).
Link Auth Service and User Service so that upon successful OAuth, a user record is created (if new).
Frontend Integration

Add a login page that redirects to Google OAuth.
After OAuth, store the JWT in an HTTP-only cookie.
Verify user can see the main menu once authenticated.
Basic test coverage with Playwright (or Cypress).
2.3 Deck & Card Management
Card Service

Admin endpoints for adding/editing/removing cards.
Card schema in MongoDB.
Basic tests: add card, fetch card, update card.
Deck Service

Endpoint to create a deck for a user, specify card composition.
Enforce point-based limits (any rule you specify).
Basic tests: create deck, list decks, validate deck.
Frontend: Deck Builder

Display available cards from the Card Service.
Allow adding/removing cards with live point limit check.
Save deck to Deck Service.
Basic UI test coverage.
2.4 Lobby & Game
Lobby Service

Create a new lobby, list public lobbies.
Join/leave logic.
Real-time state (WebSockets or NATS messages).
Basic tests: create a lobby, user joins, user leaves.
Game Service

Orchestrate game flow:
Start game when two players are ready.
Round resolution logic.
Score tracking.
Persist game state to PostgreSQL for reconnect.
Basic tests: 2 players each play a card, highest card wins, validate scoreboard.
Replay Service (Optional Early Implementation)

Store game events for replays.
Basic retrieval of event logs.
Basic test coverage.
2.5 End-to-End Integration & Admin Dashboard
Integrate All Services

Authentication gating for the entire app.
The main menu can connect to the Lobby.
Lobby triggers a new Game session.
The game references Deck and Card data.
The user can see final results.
Admin Dashboard

Manage Cards (CRUD).
Manage Maintenance Mode.
View game logs (Replay Service).
Basic test coverage.
E2E Testing

Log in, create deck, join lobby, play a game, see final results.
Admin creates new card, tries to add it to a deck, etc.
3. Final Detailed Breakdown
Below is a more granular breakdown of each phase, ensuring each step is small enough for safe implementation and testing but large enough to move the project forward.

Phase 1: Monorepo Initialization

Create top-level Git repo with backend, frontend, infrastructure folders.
Add README with project overview.
Set up GitHub Actions to run a lint check on any Go file (or a minimal test).
Phase 2: Scaffold Microservices

For each microservice (Auth, User, Deck, Card, Lobby, Game, Replay):
Create a cmd/<service_name>/main.go.
Create a pkg/<service_name> folder for internal logic.
Set up a minimal HTTP server (net/http or a framework).
Add a simple “/healthz” endpoint.
Add minimal test (service_test.go) verifying the health endpoint returns 200.
Phase 3: Frontend Scaffold

Generate Remix app in frontend.
Create routes: /login, /menu, and a generic “Not Found” route.
Add a minimal test in frontend/tests/ verifying the homepage loads.
Phase 4: Infrastructure Scaffolding

Create Terraform scripts for:
EKS cluster.
Possibly RDS for PostgreSQL and VPC.
Create a Helm chart for a placeholder “hello” microservice.
Configure ArgoCD to watch the Git repo and deploy.
Add a test that ensures the placeholder is deployed on a PR merge.
Phase 5: Auth Service (TDD)

Write tests for:
Google OAuth endpoint (mock the Google part).
JWT issuance.
Refresh token handling.
Implement code to pass tests.
Integrate with the frontend login page.
Verify with automated tests.
Phase 6: User Service (TDD)

Write tests for:
Creating a user profile.
Editing a user profile.
Retrieving a user profile by ID.
Implement these features.
Integrate with Auth Service (e.g., “on first login, create user record”).
Frontend calls the User Service to show user info on /menu.
Phase 7: Card Service & Deck Service

Card Service
Write tests for CRUD endpoints.
Implement.
Connect to MongoDB (Terraform for MongoDB Atlas or self-hosted?).
Deck Service
Write tests for deck creation (with point limit).
Implement.
Store deck in PostgreSQL.
Integrate with Card Service to validate card IDs.
Phase 8: Deck Builder UI

Write a test verifying deck creation from the UI (Cypress/Playwright).
Implement the UI that fetches card data and builds a deck.
Validate the deck before saving (point limit).
Phase 9: Lobby Service

Write tests for creating/joining/leaving a lobby.
Implement the Lobby Service with real-time updates (WebSocket or NATS).
Frontend: lobby browser page with real-time updates.
Phase 10: Game Service

Write tests for:
Handling 2-player games.
Round resolution logic (highest card wins).
Tie scenario with shuffle-back.
5-round limit.
Implement the game logic.
Integrate with Lobby Service to start a game once both players are ready.
Persist game state to DB for reconnection.
Phase 11: Replay Service (Optional Now or Later)

Write tests to store event logs.
Implement storing and retrieving logs.
Integrate with the Game Service so all round events are published to Replay.
Phase 12: Admin Dashboard

Write tests for card management and maintenance mode toggling.
Implement UI in the frontend.
Verify changes replicate (e.g., disabling a card means it can’t be added to new decks).
Phase 13: Full E2E Testing & Integration

Comprehensive test: create account → build deck → join lobby → play game → view results.
Admin test: create new card, attempt to add to deck, etc.
Phase 14: Polish & Monitoring

Add structured JSON logs, integrate with Loki.
Add metrics (Prometheus/OpenTelemetry).
Finalize Grafana dashboards.
4. Prompts for a Code-Generation LLM
Below is a sequence of prompts you can use with a code-generation LLM (like ChatGPT or GitHub Copilot). Each prompt includes a context and instructions. They are tagged as code so you can copy and paste them directly. You should run them in order, ensuring each piece of code is integrated into your repository before moving to the next prompt.

Important: The prompts assume you have a standard, empty (or near-empty) repository and are following the steps in the same order. Adjust paths, filenames, or references if your structure differs.

Prompt 1: Monorepo Initialization
text
Copy
Edit
**Context:**
We are creating a monorepo for a card-battler project with multiple Go microservices and a Remix frontend. We want a minimal initial commit.

**Instructions:**
1. Create a new Git repository with the following structure:
   - backend/
   - frontend/
   - infrastructure/
   - .github/workflows/ (for CI)
2. In each directory, add a README.md with a brief description.
3. In the root README.md, add an overview of the project.
4. Create a GitHub Actions workflow file named "ci.yml" in .github/workflows that:
   - Runs on pull_request
   - Checks out the repo
   - Installs Go
   - Runs "go version" to confirm the environment
   - Prints "Hello from CI!"

**Deliverables:**
A set of file definitions and their contents that I can copy into my new repo.
Please include a short explanation of how each file fits into the project.
Prompt 2: Scaffolding Microservices
text
Copy
Edit
**Context:**
We have a monorepo with a basic structure. We want to scaffold our Go microservices. Each microservice will have its own main.go and a minimal health check endpoint.

**Instructions:**
1. For each microservice (Auth, User, Deck, Card, Lobby, Game, Replay):
   - Create a folder under backend/ named <service_name> (e.g., backend/auth).
   - Inside each folder, create a cmd/main.go that starts an HTTP server on port 808X (increment port numbers for each service).
   - Add a /healthz endpoint that returns JSON { "status": "ok" }.
2. Write a minimal Go test for each microservice to verify that /healthz returns 200.
3. Update the GitHub Actions CI to run "go test ./..." at the root, so it covers all services.

**Deliverables:**
- Updated file structure for each microservice.
- main.go files for each service with the code for the healthz endpoint.
- test files for each service verifying the endpoint.
- Explanation on how to run these locally and how to run tests.
Prompt 3: Frontend Scaffold
text
Copy
Edit
**Context:**
We have a monorepo with Go microservices scaffolded. Now we want to set up a Remix frontend in the `frontend` folder. We'll just have a minimal set of routes and a test.

**Instructions:**
1. Initialize a Remix project in the `frontend` folder using TypeScript.
2. Create routes:
   - /login
   - /menu
   - / (root) that shows a "Home Page" message
3. Add a basic test setup (Jest or Playwright) to verify the root route returns a 200 status and the text "Home Page".
4. Explain how to run the frontend in dev mode and how to run tests.

**Deliverables:**
- The relevant package.json, Remix config, routes, and test file(s).
- Explanation of how the front end is structured and how to run it.
Prompt 4: Infrastructure Scaffolding
text
Copy
Edit
**Context:**
We have microservices in Go and a Remix frontend. We want to create infrastructure code in the `infrastructure` folder. We'll use Terraform to provision an EKS cluster on AWS and create a simple Helm chart that deploys a "hello" microservice.

**Instructions:**
1. Create Terraform files (main.tf, variables.tf) to:
   - Stand up an EKS cluster.
   - Create a node group with minimal capacity.
2. Create a minimal Helm chart under infrastructure/helm/hello-chart:
   - Deployment.yaml with a single container that just runs a "hello" container (e.g., public nginx).
   - Service.yaml exposing port 80.
   - A basic values.yaml.
3. Provide instructions on:
   - Setting up AWS credentials.
   - Running `terraform init/plan/apply`.
   - Installing the Helm chart to the new cluster with `helm install`.
4. Explain how to verify the "hello" service is running.

**Deliverables:**
- Terraform file contents (main.tf, variables.tf, outputs.tf if needed).
- Helm chart file contents (Deployment.yaml, Service.yaml, values.yaml).
- Step-by-step instructions for deployment and verification.
Prompt 5: Auth Service (TDD)
text
Copy
Edit
**Context:**
Our microservices are scaffolded. Let’s focus on the Auth Service and implement Google OAuth + JWT. We’ll do test-driven development.

**Instructions:**
1. Write tests first for:
   - A Google OAuth handler that receives a code (mock or stub the Google part).
   - JWT generation and validation.
   - Storing the JWT in an HTTP-only cookie.
2. Implement only enough code to pass these tests.
3. Provide code for the Auth Service (cmd/main.go or pkg/auth) that:
   - Defines a /oauth/google endpoint which simulates exchanging the code for a token.
   - Issues a JWT with a user ID claim.
   - Returns it in an HTTP-only cookie.
4. Explain how to run the tests and demonstrate they pass.

**Deliverables:**
- Go test file showing the TDD approach.
- Implementation code that passes the tests.
- Explanation of any libraries used (like jwt-go or similar).
Prompt 6: User Service Integration
text
Copy
Edit
**Context:**
We have an Auth Service providing JWT tokens. Now we need a User Service that manages user profiles (username, avatar), storing data in PostgreSQL. It should integrate with the Auth Service so that on first login, we create a new user row.

**Instructions:**
1. Write tests for the User Service:
   - Creating a user profile.
   - Editing the profile.
   - Retrieving the profile by ID.
2. Add a PostgreSQL connection (use a DSN from environment variables).
3. Implement the endpoints in the User Service:
   - POST /users to create a user.
   - PATCH /users/{id} to edit user data.
   - GET /users/{id} to retrieve user data.
4. Integrate with the Auth Service:
   - On successful OAuth login, if user ID is not found in DB, create a new user row automatically.
5. Provide instructions to run migrations (if using a library like golang-migrate) and test it.

**Deliverables:**
- Tests for each endpoint (table-driven tests are nice).
- Implementation code for the User Service.
- SQL schema or migrations for the users table.
- Explanation of how to run migrations and tests locally.
Prompt 7: Card & Deck Services
text
Copy
Edit
**Context:**
We want the Card Service to store card data (using MongoDB), and the Deck Service to store deck configurations in PostgreSQL. We need to enforce point-based deck limits.

**Instructions:**
1. For the Card Service:
   - Write tests for creating, reading, updating, deleting cards.
   - Implement in Go with a MongoDB client. Use environment variables for the MongoDB connection string.
   - Add an admin-only check (assume a role claim in JWT) to allow modifications.
2. For the Deck Service:
   - Write tests for deck creation, listing decks, updating decks.
   - Enforce a point limit (e.g., total card points <= 30).
   - Cross-check card IDs with the Card Service to ensure each card is valid/active.
3. Provide the code, tests, and instructions to run them. Indicate how you’d mock/stub calls to Card Service in the Deck Service tests.

**Deliverables:**
- Test files for Card Service and Deck Service.
- Implementation code (handlers + storage logic).
- Explanation of how to run them locally and how to set environment variables for DB connections.
Prompt 8: Frontend Deck Builder
text
Copy
Edit
**Context:**
We now have a Card Service and a Deck Service. We want the Remix frontend to display available cards and allow building a deck. We’ll do a basic TDD approach with integration tests (Playwright/Cypress).

**Instructions:**
1. Write an integration test that:
   - Logs in a user (stub if necessary).
   - Navigates to a “Deck Builder” page.
   - Fetches available cards (mock or stub an API response).
   - Adds cards until the point limit is reached.
   - Attempts to save the deck.
   - Verifies the deck is saved successfully.
2. Implement the Remix route `/deck-builder` that:
   - Uses a loader to fetch cards from the Card Service.
   - Renders them with a button to add them to a “proposed deck” state.
   - Shows the total point cost.
   - Provides a button to save the deck to the Deck Service.
3. Demonstrate the test passing with real or mock data.

**Deliverables:**
- Integration test file.
- Deck Builder route code (loader + UI).
- Explanation of how to run the integration test.
Prompt 9: Lobby Service
text
Copy
Edit
**Context:**
We need a Lobby Service that manages public lobbies, letting players join/leave. We’ll use WebSockets for real-time updates.

**Instructions:**
1. Write tests (unit + integration) for:
   - Creating a lobby (POST /lobbies).
   - Listing lobbies (GET /lobbies).
   - Joining a lobby (websocket or REST).
   - Leaving a lobby.
2. Implement the Lobby Service using Go. Consider storing lobby state in memory (or Redis if you need persistence).
3. Expose a WebSocket endpoint that broadcasts lobby updates (who’s joined, how many spots are left, etc.).
4. Provide a minimal client that demonstrates connecting to the WebSocket and receiving updates.

**Deliverables:**
- Lobby Service code (handlers, data model).
- Test files covering each scenario.
- Explanation of how to connect to the WebSocket endpoint and verify real-time updates.
Prompt 10: Game Service
text
Copy
Edit
**Context:**
We have lobbies. Once a lobby has two players ready, we move to the Game Service. The Game Service manages a 2-player card battle with 5 rounds. Highest card wins, ties cause replay, etc.

**Instructions:**
1. Write TDD tests:
   - Start game with 2 players and each has a deck of 5 cards.
   - Each round: players choose a card, the system compares them, points are assigned.
   - On tie: shuffle cards back in, retry (up to 3 times).
   - Game ends after 5 rounds or deck exhaustion. Highest total points wins.
2. Implement the Game Service with a persistent game state (PostgreSQL).
3. Provide a real-time method for the frontend to send chosen cards (WebSocket or REST).
4. Demonstrate test coverage and explain how to run an end-to-end local test.

**Deliverables:**
- Game Service code.
- Test files for each rule scenario.
- Explanation of data structures used to store game state and handle concurrency.
Prompt 11: Replay Service (If desired at this stage)
text
Copy
Edit
**Context:**
Optional now or later. This service stores game events for replays.

**Instructions:**
1. Write tests for storing events (round start, round result, final game result).
2. Implement a minimal gRPC or REST to accept events from the Game Service.
3. Provide an endpoint for retrieving a game’s event log.
4. Show the test coverage.

**Deliverables:**
- Replay Service code.
- Test files.
- Explanation of how the Game Service calls the Replay Service.
Prompt 12: Admin Dashboard
text
Copy
Edit
**Context:**
We want an admin dashboard to manage cards (enable/disable, edit) and toggle maintenance mode.

**Instructions:**
1. Write tests for:
   - Admin can view all cards, update a card, disable a card.
   - Admin can toggle maintenance mode, which prevents new games from starting but allows existing ones to finish.
2. Implement the UI in Remix at `/admin`.
3. Connect to the Card Service (for card edits) and possibly a Config Service or environment variable to store “maintenance mode”.
4. Show test results and how to run them.

**Deliverables:**
- Admin Dashboard route code.
- Admin-related service code (if needed).
- Test files (integration + any unit tests).
Prompt 13: Full E2E Testing
text
Copy
Edit
**Context:**
We want a final test that covers the entire flow: login → deck creation → lobby join → game → results.

**Instructions:**
1. Write a Playwright/Cypress test scenario:
   - User logs in via Auth Service (stub or real).
   - User creates a deck using the Deck Builder.
   - User navigates to the Lobby, joins a game.
   - User plays a few rounds in the Game Service.
   - Test verifies final results are shown.
2. Integrate these E2E tests into CI so they run automatically on push.
3. Provide instructions on how to run them locally and in CI.

**Deliverables:**
- E2E test file(s).
- Explanation of test steps and pass criteria.
Prompt 14: Observability & Monitoring
text
Copy
Edit
**Context:**
As a final step, we want to add structured logging, metrics (Prometheus/OpenTelemetry), and set up Grafana dashboards.

**Instructions:**
1. Update each microservice to emit structured JSON logs (Zap or logrus).
2. Add a /metrics endpoint (Prometheus format).
3. Provide instructions on how to deploy Prometheus and Grafana in the cluster.
4. Explain how to create a Grafana dashboard that visualizes request counts and latencies.

**Deliverables:**
- Code changes for each microservice enabling JSON logs and /metrics.
- Terraform or Helm additions for Prometheus & Grafana deployment.
- Basic instructions on adding a Grafana dashboard for those metrics.
Prompt 15: Wrap-Up & Final Integration
text
Copy
Edit
**Context:**
We now have everything built out. Let’s finalize and ensure no code is left hanging.

**Instructions:**
1. Review all microservices and confirm they are integrated:
   - Auth → User → Deck → Card → Lobby → Game → Replay (optional).
   - Admin Dashboard integrated.
2. Provide final code snippets or references showing how each microservice is connected.
3. Summarize any environment variables needed and how to run the entire stack locally via Docker Compose or Kubernetes.
4. Explain final steps for production deployment with ArgoCD.

**Deliverables:**
- A short final “integration guide” that ties all services together.
- Docker Compose or K8s-based examples of how to run everything together.
- Confirmation that the entire system passes the E2E tests.
Conclusion
By following this step-by-step outline and using each prompt in order, you’ll incrementally build out the Card Battler project with safe, test-driven progress. Each prompt ensures best practices, minimal breakage, and forward momentum. Happy coding!