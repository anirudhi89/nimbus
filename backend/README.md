nimbus/
│
├── cmd/
│   └── nimbus-server/                  # Main entry point for the app
│       └── main.go                     # Starts the server and initializes routes
│
├── config/                             # Configuration files and environment setup
│   └── config.go                       # Configuration loading and parsing (e.g., environment variables)
│
├── internal/                           # Core application code
│   ├── auth/                           # Authentication related logic
│   │   └── auth.go                     # Auth logic (JWT generation, validation)
│   │   └── auth_service.go             # Logic to handle user login, registration, etc.
│   │
│   ├── user/                           # User related logic
│   │   └── user.go                     # User model and methods (e.g., user registration)
│   │   └── user_service.go             # Logic to manage user data (CRUD operations)
│   │
│   ├── subscription/                   # Subscription and billing logic
│   │   └── subscription.go             # Handles subscription plan and billing
│   │   └── subscription_service.go     # Business logic for managing plans and upgrades
│   │
│   ├── aws/                            # AWS integration
│   │   └── aws_integration.go          # AWS IAM and other related service interactions
│   │   └── aws_service.go              # Handles data fetch and logic related to AWS (EC2, EKS)
│   │
│   ├── dashboard/                      # Dashboard creation and management
│   │   └── dashboard.go                # Handles dashboard creation, data handling
│   │   └── dashboard_service.go        # Logic for customizing and updating dashboards
│   │
│   └── alert/                          # Alerting logic
│       └── alert.go                    # Alert model, configurations, and settings
│       └── alert_service.go            # Logic for setting up, sending, and managing alerts
│
├── pkg/                                # Reusable and shared libraries
│   ├── logger/                         # Logging utilities
│   │   └── logger.go                   # Logger setup and helper functions
│   └── utils/                          # Helper functions for various utilities
│       └── utils.go                    # Utility functions (e.g., date formatting, etc.)
│
├── router/                             # Routing and API endpoints
│   └── router.go                       # Routes for handling different API requests
│
├── middleware/                         # Middlewares for error handling, JWT auth, etc.
│   └── auth_middleware.go              # JWT validation middleware
│   └── error_middleware.go             # Error handling middleware
│
├── storage/                            # Database models and interactions (could be MongoDB, SQL, etc.)
│   ├── user_model.go                   # User model for DB interactions
│   ├── subscription_model.go           # Subscription DB interactions
│   └── aws_model.go                    # AWS-related database models for tracking connections
│
├── migrations/                         # Database migration scripts
│   └── migrate.go                      # Logic for managing database schema migrations
│
├── Dockerfile                          # Docker setup for building the backend container
├── go.mod                              # Go modules dependencies
└── README.md                           # Documentation on how to set up and run the backend
