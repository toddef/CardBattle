# Card Battler - Developer Specification

## 1. Overview
**Card Battler** is a two-player card game where players compete by playing cards in simultaneous turns. The game includes deck-building mechanics, a lobby system for matchmaking, real-time WebSocket-based gameplay, and admin controls for card management.

This project is structured as a **monorepo** containing:
- **Backend (Golang) microservices**
- **Frontend (Remix framework)**
- **Infrastructure (Kubernetes + ArgoCD + Terraform for AWS EKS deployment)**

---

## 2. Core Features

### 2.1 Player Features
- Google OAuth authentication (JWT-based, stored in HTTP-only cookies).
- Profile management (custom username, Google avatar).
- Deck building (up to 10 decks, point-based deck limits).
- Lobby system (players browse/join public lobbies).
- Real-time, WebSocket-based gameplay.
- Automated turn resolution (cards played simultaneously).
- Reconnect support (game state is persisted).
- Game history (last 10 matches stored with basic results).
- Post-game summary (final score, opponent name).
- In-game chat (shared between players and spectators).
- Matchmaking (manual lobby selection, no automatic matchmaking).

### 2.2 Game Mechanics
- Each player starts with a **shuffled deck** and a **hand of 5 cards**.
- Players **simultaneously play a card**, which is then revealed.
- The highest card wins the round, earning points equal to both card values.
- Special card abilities can be introduced in the future.
- If a round is tied, the played cards are **shuffled back into decks**, and the round repeats (max 3 retries).
- The game ends after **5 rounds**, and the **highest total score wins**.
- If the game is tied at the end, it is **marked as a draw**.

### 2.3 Admin Features
- **Admin dashboard** to manage cards, decks, and game rules.
- Add/edit/remove cards in real time (changes apply to all decks).
- Ability to **disable specific cards** from play.
- View game history and replays (event-based replays).
- **Enable maintenance mode**, preventing new games from starting while allowing active games to finish.

---

## 3. Architecture

### 3.1 Backend Microservices (Golang)
Each microservice runs independently and communicates via **gRPC, REST, and NATS messaging**.

| **Service**        | **Responsibilities** |
|--------------------|---------------------|
| **Auth Service**   | Handles Google OAuth, JWT issuance, and user authentication. |
| **User Service**   | Manages player profiles, usernames, and stats. |
| **Deck Service**   | Manages deck creation, validation, and storage. |
| **Card Service**   | Stores and manages available cards (admin-controlled). |
| **Lobby Service**  | Handles lobby creation, player/spectator tracking. |
| **Game Service**   | Manages active games, enforces rules, and processes turns. |
| **Replay Service** | Stores past game events for replays. |

### 3.2 Frontend (Remix Framework + React)
- **Server-driven rendering** with Remix loaders/actions.
- **WebSockets for real-time updates** (gameplay, chat).
- **Pages:**
  - Login
  - Main menu (navigate to other pages)
  - Deck builder
  - Card browser
  - Lobby browser
  - Game UI
  - Profile/stats
  - Admin dashboard

### 3.3 Database & Storage
- **PostgreSQL** → Structured data (users, decks, games, stats).
- **MongoDB** → Flexible storage (cards, replays).
- **Redis** → Caching layer for frequently accessed data.

---

## 4. Infrastructure & Deployment

### 4.1 Kubernetes Setup
- **Microservices deployed as independent containers**.
- **Single Helm chart** for the entire system.
- **Kubernetes Ingress Controller** for routing traffic.
- **Horizontal Pod Autoscaler (HPA)** for automatic scaling.

### 4.2 CI/CD (GitHub Actions + ArgoCD)
- GitHub Actions for **CI pipeline** (build, test, push images).
- ArgoCD for **automated rollouts and rollbacks** on **tagged releases**.
- Terraform for provisioning **AWS EKS, RDS (PostgreSQL), ElastiCache (Redis), and MongoDB Atlas**.

---

## 5. Error Handling & Recovery

### 5.1 Game State & Desync Handling
- **Game state is persisted in the database** to allow reconnects.
- **Regular client-server sync checks** to prevent desync.
- **Logged desync events** for debugging.

### 5.2 Server Failures & Recovery
- If a server failure occurs, the **game reverts to the last valid state**.
- If recovery fails, the game is **marked as abandoned**.

### 5.3 Invalid Moves & Cheating Prevention
- The **backend validates all game actions**.
- If a player attempts an **invalid action**, they **lose the round**.
- **Invalid actions are logged** for analysis.

---

## 6. Testing Strategy

### 6.1 Unit Tests
- Core logic for **game mechanics, deck validation, and authentication**.
- Mock database queries and API responses.

### 6.2 Integration Tests
- Validate interservice communication via **gRPC and REST**.
- Ensure **deck selection, game flow, and scoring logic** function correctly.

### 6.3 End-to-End (E2E) Tests
- Simulated **full game flow** from login to game completion.
- Test **real-time interactions** (WebSockets, lobby updates).

### 6.4 Performance & Load Testing (Future Consideration)
- Not included in the initial version but can be added later.

---

## 7. Security & Compliance

### 7.1 Data Privacy & Account Management
- **Minimal user data stored** (Google OAuth handles authentication).
- **Players can delete their account**, which removes all associated data.

### 7.2 API Security
- **Central API Gateway** for authentication and authorization.
- **Admin-only actions require special roles**.

### 7.3 Logging & Monitoring
- **Structured JSON logs**, collected via **Loki**.
- **Prometheus + OpenTelemetry** for performance monitoring.
- **Grafana dashboards** for observability.

---

## 8. Future Considerations
- **Feature flags** (dynamically enable/disable new features).
- **Deck sharing** (generate/import deck codes).
- **Game variants and custom rule sets**.
- **Cosmetic customization (card backs, themes, avatars beyond Google OAuth)**.
- **Advanced analytics (player trends, win rates, card effectiveness tracking)**.

---

## Final Notes
This specification provides a **complete roadmap for development**. The **initial focus** will be on implementing **core gameplay, authentication, and infrastructure**, followed by iterative feature improvements.

### 💡 Next Steps:
✅ Set up **repo structure and initial scaffolding**.  
✅ Implement **authentication and backend services**.  
✅ Develop **frontend UI components**.  
✅ Deploy **Kubernetes infrastructure** and test CI/CD.

🚀 **Time to build!**
